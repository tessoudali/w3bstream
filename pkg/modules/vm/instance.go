package vm

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/iotexproject/Bumblebee/x/mapx"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"

	"github.com/iotexproject/w3bstream/pkg/enums"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
)

func NewInstance(path string, opts ...InstanceOptionSetter) (string, error) {
	code, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	i, err := NewInstanceByCode(code, opts...)
	if err != nil {
		return "", err
	}
	return AddInstance(i), nil
}

func NewInstanceWithID(path string, by string, opts ...InstanceOptionSetter) error {
	code, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	i, err := NewInstanceByCode(code, opts...)
	if err != nil {
		return err
	}

	AddInstanceByID(by, i)
	return nil
}

func NewInstanceByCode(code []byte, opts ...InstanceOptionSetter) (*Instance, error) {
	ctx := context.Background()
	opt := &InstanceOption{
		RuntimeConfig: DefaultRuntimeConfig,
		Logger:        DefaultLogger,
		Tasks:         &TaskQueue{ch: make(chan *Task)},
	}

	for _, set := range opts {
		set(opt)
	}

	i := &Instance{
		opt:      opt,
		state:    enums.INSTANCE_STATE__CREATED,
		rt:       wazero.NewRuntimeWithConfig(ctx, opt.RuntimeConfig),
		handlers: make(map[string]api.Function),
		db:       make(map[string]int32),
	}

	{
		_, err := i.rt.NewModuleBuilder("env").
			ExportFunction("get_data", i.GetData).
			ExportFunction("set_data", i.SetData).
			ExportFunction("get_db", i.GetDB).
			ExportFunction("set_db", i.SetDB).
			ExportFunction("log", i.Log).
			Instantiate(ctx, i.rt)
		if err != nil {
			return nil, err
		}
	}

	_, err := wasi_snapshot_preview1.Instantiate(ctx, i.rt)
	if err != nil {
		return nil, err
	}

	i.mod, err = i.rt.InstantiateModuleFromBinary(ctx, code)
	if err != nil {
		return nil, err
	}
	i.malloc = i.mod.ExportedFunction("malloc")
	i.free = i.mod.ExportedFunction("free")

	i.res = mapx.New[uint32, []byte]()
	i.ctx, i.cancel = context.WithCancel(ctx)

	return i, nil
}

type Instance struct {
	opt      *InstanceOption
	ctx      context.Context
	cancel   context.CancelFunc
	state    wasm.InstanceState
	rt       wazero.Runtime
	mod      api.Module
	res      *mapx.Map[uint32, []byte]
	malloc   api.Function
	free     api.Function
	handlers map[string]api.Function
	db       map[string]int32
}

var _ wasm.Instance = (*Instance)(nil)

func (i *Instance) Start() error {
	for {
		select {
		case <-i.ctx.Done():
			return i.ctx.Err()
		case task := <-i.opt.Tasks.Wait():
			task.Res <- i.handleEvent(task)
		}
	}
}

func (i *Instance) Stop() {
	i.state = enums.INSTANCE_STATE__STOPPED
	i.cancel()
}

func (i *Instance) State() wasm.InstanceState { return i.state }

func (i *Instance) HandleEvent(fn string, data []byte) ([]byte, wasm.ResultStatusCode) {
	task := &Task{
		Handler: fn,
		Payload: data,
		Res:     make(chan *EventHandleResult),
	}
	i.opt.Tasks.Push(task)

	res := <-task.Res
	return res.Response, res.Code
}

func (i *Instance) handleEvent(t *Task) *EventHandleResult {
	rid := i.AddResource(t.Payload)
	defer i.RmvResource(rid)

	hdl, ok := i.handlers[t.Handler]
	if !ok {
		hdl = i.mod.ExportedFunction(t.Handler)
		if hdl == nil {
			return &EventHandleResult{nil, wasm.ResultStatusCode_UnexportedHandler}
		}
		i.handlers[t.Handler] = hdl
	}

	results, err := hdl.Call(i.ctx, uint64(rid))
	if err != nil {
		return &EventHandleResult{nil, wasm.ResultStatusCode_Failed}
	}

	return &EventHandleResult{nil, wasm.ResultStatusCode(results[0])}
}

func (i *Instance) AddResource(data []byte) uint32 {
	id := uuid.New().ID()
	i.res.Store(id, data)
	return id
}

func (i *Instance) GetResource(id uint32) ([]byte, bool) { return i.res.Load(id) }

func (i *Instance) RmvResource(id uint32) { i.res.Remove(id) }

func (i *Instance) Get(k string) int32 { return i.db[k] }
