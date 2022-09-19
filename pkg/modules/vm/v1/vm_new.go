package v1

import (
	"context"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"

	"github.com/iotexproject/w3bstream/pkg/types"
)

type Instance struct {
	rc     wazero.Runtime
	ctx    context.Context
	mod    api.Module
	malloc api.Function
	free   api.Function
	start  api.Function
	topic  string
}

func NewInstance(c context.Context, loc string, topic string) (*Instance, error) {
	ctx := context.Background()

	content, err := os.ReadFile(loc)
	if err != nil {
		return nil, err
	}

	// new wasm runtime.
	r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfig().
		WithFeatureBulkMemoryOperations(true).
		WithFeatureNonTrappingFloatToIntConversion(true).
		WithFeatureSignExtensionOps(true).WithFeatureMultiValue(true))

	// exports
	{
		_, err := r.NewModuleBuilder("env").
			ExportFunction("log", log).
			ExportFunction("inc", inc).
			ExportFunction("get", get).
			Instantiate(ctx, r)
		if err != nil {
			return nil, err
		}
	}

	if _, err := wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
		panic(err)
	}

	mod, err := r.InstantiateModuleFromBinary(ctx, content)
	if err != nil {
		panic(err)
	}

	ins := &Instance{
		ctx:    ctx,
		rc:     r,
		mod:    mod,
		start:  mod.ExportedFunction("start"),
		malloc: mod.ExportedFunction("malloc"),
		free:   mod.ExportedFunction("free"),
		topic:  topic,
	}

	AddInstance(ins)
	go ins.Start(c)
	return ins, nil
}

func (i *Instance) Start(ctx context.Context) {
	defer RemoveInstance(i)
	defer i.rc.Close(i.ctx)
	l := types.MustLoggerFromContext(ctx)
	b := types.MustMqttBrokerFromContext(ctx)

	l.Info("%s started", i.topic)

	cli, err := b.Client(i.topic)
	if err != nil {
		l.Error(err)
	}
	for {
		err = cli.WithTopic(i.topic).Subscribe(
			func(c mqtt.Client, msg mqtt.Message) {
				data := msg.Payload()
				length := uint64(len(data))

				results, err := i.malloc.Call(i.ctx, length)
				if err != nil {
					l.Error(err)
					return
				}
				ptr := results[0]
				defer i.free.Call(i.ctx, ptr)

				if !i.mod.Memory().Write(i.ctx, uint32(ptr), data) {
					l.Error(fmt.Errorf("Memory.Write(%d, %d) out of range of memory size %d",
						ptr, length, i.mod.Memory().Size(i.ctx)))
					return
				}

				_, err = i.start.Call(i.ctx, ptr, length)
				if err != nil {
					l.Error(err)
					return
				}
			},
		)
		if err != nil {
			break
		}
	}
}

func log(ctx context.Context, m api.Module, offset, size uint32) {
	buf, ok := m.Memory().Read(ctx, offset, size)
	if !ok {
		panic(fmt.Sprintf("Memory.Read(%d,%d) out of range)", offset, size))
	}
	fmt.Println(string(buf))
}

var words = make(map[string]int32)

func inc(ctx context.Context, m api.Module, offset, size uint32, delta int32) (code int32) {
	buf, ok := m.Memory().Read(ctx, offset, size)
	if !ok {
		return 1
	}
	str := string(buf)
	if _, ok := words[str]; !ok {
		words[str] = delta
	} else {
		words[str] = words[str] + delta
	}
	return 0
}

func get(ctx context.Context, m api.Module, offset, size uint32) (value int32) {
	buf, ok := m.Memory().Read(ctx, offset, size)
	if !ok {
		return 0
	}
	str := string(buf)
	if _, ok := words[str]; !ok {
		return 0
	}
	return words[str]
}
