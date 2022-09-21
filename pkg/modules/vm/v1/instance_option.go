package v1

import (
	"context"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	conflog "github.com/iotexproject/Bumblebee/conf/log"
	confmqtt "github.com/iotexproject/Bumblebee/conf/mqtt"
	"github.com/iotexproject/Bumblebee/x/mapx"
	"github.com/iotexproject/w3bstream/pkg/types"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
	"github.com/pkg/errors"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

type InstanceOption struct {
	Channel       string
	RuntimeConfig wazero.RuntimeConfig
	Logger        conflog.Logger
	Broker        *confmqtt.Broker
	Client        *confmqtt.Client
}

type InstanceOptionSetter func(o *InstanceOption)

func InstanceOptionWithChannel(ch string) InstanceOptionSetter {
	return func(o *InstanceOption) { o.Channel = ch }
}

func InstanceOptionWithRuntimeConfig(rc wazero.RuntimeConfig) InstanceOptionSetter {
	return func(o *InstanceOption) { o.RuntimeConfig = rc }
}

func InstanceOptionWithLogger(l conflog.Logger) InstanceOptionSetter {
	return func(o *InstanceOption) { o.Logger = l }
}

func InstanceOptionWithMqttBroker(b *confmqtt.Broker) InstanceOptionSetter {
	return func(o *InstanceOption) { o.Broker = b }
}

var (
	DefaultRuntimeConfig = wazero.NewRuntimeConfig().
		WithFeatureBulkMemoryOperations(true).
		WithFeatureNonTrappingFloatToIntConversion(true).
		WithFeatureSignExtensionOps(true).
		WithFeatureMultiValue(true)
	DefaultLogger = conflog.Std()
)

func NewInstance(path string, opts ...InstanceOptionSetter) (uint32, error) {
	ctx := context.Background()

	code, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	opt := &InstanceOption{
		RuntimeConfig: DefaultRuntimeConfig,
		Logger:        DefaultLogger,
	}

	if broker, _ := types.MqttBrokerFromContext(ctx); broker != nil {
		opt.Broker = broker
	}
	if logger, _ := types.LoggerFromContext(ctx); logger != nil {
		opt.Logger = logger
	}

	for _, set := range opts {
		set(opt)
	}

	if opt.Broker != nil && opt.Channel != "" {
		opt.Client, err = opt.Broker.Client(opt.Channel)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("no transporter")
	}

	i := &Instance{
		opt:   opt,
		state: wasm.InstanceState_Created,
		rt:    wazero.NewRuntimeWithConfig(ctx, opt.RuntimeConfig),
	}

	{
		_, err := i.rt.NewModuleBuilder("env").
			ExportFunction("get_data", log).
			ExportFunction("set_data", inc).
			ExportFunction("get_db", get).
			ExportFunction("set_db", nil).
			Instantiate(ctx, i.rt)
		if err != nil {
			return 0, err
		}
	}

	_, err = wasi_snapshot_preview1.Instantiate(ctx, i.rt)
	if err != nil {
		return 0, err
	}

	i.mod, err = i.rt.InstantiateModuleFromBinary(ctx, code)
	if err != nil {
		return 0, err
	}
	i.res = mapx.New[uint32, []byte]()
	i.ctx, i.cancel = context.WithCancel(ctx)

	return AddInstance(i), nil
}

type Instance struct {
	opt    *InstanceOption
	ctx    context.Context
	cancel context.CancelFunc
	state  wasm.InstanceState
	rt     wazero.Runtime
	mod    api.Module
	res    *mapx.Map[uint32, []byte]

	malloc api.Function
	free   api.Function
	start  api.Function

	wasm.ExportsHandler
}

var _ wasm.Instance = (*Instance)(nil)

func (i *Instance) Start() error {
	// start consuming event
	for {
		select {
		case <-i.ctx.Done(): // @todo log
			return i.ctx.Err()
		default:
			err := i.opt.Client.WithTopic(i.opt.Channel).Subscribe(
				func(c mqtt.Client, msg mqtt.Message) {
					code := i.HandleEvent(msg.Payload())
					if code != wasm.ResultStatusCode_OK {
						// @todo log
					}
				},
			)
			if err != nil {
				return err
			}
		}
	}
}

func (i *Instance) Stop() {
	i.state = wasm.InstanceState_Stopped
	i.cancel()
}

func (i *Instance) State() wasm.InstanceState { return i.state }

func (i *Instance) HandleEvent(data []byte) wasm.ResultStatusCode {
	rid := i.AddResource(data)
	defer i.RmvResource(rid)

	results, err := i.start.Call(i.ctx, uint64(rid))
	if err != nil {
		return wasm.ResultStatusCode_Failed
	}
	return wasm.ResultStatusCode(results[0])
}

func (i *Instance) AddResource(data []byte) uint32 {
	id := uuid.New().ID()
	i.res.Store(id, data)
	return id
}

func (i *Instance) GetResource(id uint32) ([]byte, bool) { return i.res.Load(id) }

func (i *Instance) RmvResource(id uint32) { i.res.Remove(id) }

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

// func mapping(hostData interface{}) {
// 	ptr := malloc(sizeof(hostData))
// 	copy(ptr, hostData)
// }
// func eventHandle(data []byte) {
// 	rid := resMgr.Set(data)
// 	defer resMgr.Del(rid)
// 	vm.start(rid)
// }
//
// func getData(rid uint32, data_ptr_addr int32, size_addr int32) (code int32) {
// 	hostData := host.find(rid)           // []byte
// 	vmOffset := vm.malloc(len(hostData)) // []byte
//
// 	buf := mapping(vmOffset, len(hostData))
//
// 	copy(buf, hostData)
// 	copy(data_ptr_addr, vmOffset)
// 	copy(size_addr, len(hostData))
// 	return 0
// }
//
// func getDB(key_data i32, key_size i32, value_ptr_addr i32, value_size_addr i32) (code i32) {
// 	key := mapping(key_data, key_size)
// 	hostData := host.Get(key)
//
// 	vmOffset := vm.malloc(len(hostData)) // []byte
//
// 	buf := mapping(vmOffset, len(hostData))
//
// 	copy(buf, hostData)
// 	copy(value_ptr_addr, vmOffset)
// 	copy(value_size_addr, len(hostData))
// 	return 0
// }
//
// func setData(rid i32, offset i32, size i32) i32 {
// 	buf := []byte{}
// 	copy(buf, offset, size)
//
// 	resMgr.Set(rid, buf)
// }
//
