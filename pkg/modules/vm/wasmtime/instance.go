package wasmtime

import (
	"bytes"
	"context"
	"encoding/binary"

	"github.com/bytecodealliance/wasmtime-go"
	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/iotexproject/Bumblebee/x/mapx"

	"github.com/iotexproject/w3bstream/pkg/enums"
	"github.com/iotexproject/w3bstream/pkg/modules/vm/common"
	"github.com/iotexproject/w3bstream/pkg/types"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
)

func NewInstanceByCode(ctx context.Context, code []byte, opts ...common.InstanceOptionSetter) (wasm.Instance, error) {
	l := types.MustLoggerFromContext(ctx)

	_, l = l.Start(ctx, "NewInstanceByCode")
	defer l.End()

	opt := &common.InstanceOption{
		Logger: common.DefaultLogger,
		Tasks:  &common.TaskQueue{Ch: make(chan *common.Task)},
	}

	for _, set := range opts {
		set(opt)
	}

	res := mapx.New[uint32, []byte]()
	db := make(map[string][]byte)

	vmEngine := wasmtime.NewEngineWithConfig(wasmtime.NewConfig())
	vmStore := wasmtime.NewStore(vmEngine)
	linker := wasmtime.NewLinker(vmEngine)

	var cl *ChainClient
	if ethConf, ok := types.ETHClientConfigFromContext(ctx); ok {

		chain, err := ethclient.Dial(ethConf.ChainEndpoint)
		if err != nil {
			return nil, err
		}
		cl = &ChainClient{
			pvk:   crypto.ToECDSAUnsafe(gethCommon.FromHex(ethConf.PrivateKey)),
			chain: chain,
		}
	}

	ef := ExportFuncs{vmStore, res, db, opt.Logger, cl}
	_ = linker.FuncWrap("env", "ws_get_data", ef.GetData)
	_ = linker.FuncWrap("env", "ws_set_data", ef.SetData)
	_ = linker.FuncWrap("env", "ws_get_db", ef.GetDB)
	_ = linker.FuncWrap("env", "ws_set_db", ef.SetDB)
	_ = linker.FuncWrap("env", "ws_log", ef.Log)
	_ = linker.FuncWrap("env", "ws_send_tx", ef.SendTX)
	_ = linker.FuncWrap("env", "ws_call_contract", ef.CallContract)

	_ = linker.DefineWasi()

	wasiConfig := wasmtime.NewWasiConfig()
	vmStore.SetWasi(wasiConfig)

	vmModule, err := wasmtime.NewModule(vmEngine, code)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	vmInstance, err := linker.Instantiate(vmStore, vmModule)
	if err != nil {
		l.Error(err)
		return nil, err
	}

	cctx, cancel := context.WithCancel(context.Background())

	return &Instance{
		tasks:      opt.Tasks,
		ctx:        cctx,
		cancel:     cancel,
		state:      enums.INSTANCE_STATE__CREATED,
		vmEngine:   vmEngine,
		vmStore:    vmStore,
		vmModule:   vmModule,
		vmInstance: vmInstance,
		res:        res,
		handlers:   make(map[string]*wasmtime.Func),
		db:         db,
	}, nil
}

type Instance struct {
	tasks      *common.TaskQueue
	ctx        context.Context
	cancel     context.CancelFunc
	state      wasm.InstanceState
	vmEngine   *wasmtime.Engine
	vmStore    *wasmtime.Store
	vmModule   *wasmtime.Module
	vmInstance *wasmtime.Instance
	res        *mapx.Map[uint32, []byte]
	handlers   map[string]*wasmtime.Func
	db         map[string][]byte
}

var _ wasm.Instance = (*Instance)(nil)

func (i *Instance) Start(ctx context.Context) error {
	l := types.MustLoggerFromContext(ctx)

	_, l = l.Start(ctx, "instance.Start")
	defer l.End()

	for {
		select {
		case <-i.ctx.Done():
			l.Error(i.ctx.Err())
			return ctx.Err()
		case task := <-i.tasks.Wait():
			task.Res <- i.handleEvent(ctx, task)
		}
	}
}

func (i *Instance) Stop() {
	i.state = enums.INSTANCE_STATE__STOPPED
	i.cancel()
}

func (i *Instance) State() wasm.InstanceState { return i.state }

func (i *Instance) HandleEvent(ctx context.Context, fn string, data []byte) ([]byte, wasm.ResultStatusCode, error) {
	select {
	case <-ctx.Done():
		return nil, -1, ctx.Err()
	default:
		task := &common.Task{
			Handler: fn,
			Payload: data,
			Res:     make(chan *common.EventHandleResult),
		}
		i.tasks.Push(task)

		res := <-task.Res

		return res.Response, res.Code, nil
	}
}

func (i *Instance) handleEvent(ctx context.Context, t *common.Task) *common.EventHandleResult {
	l := types.MustLoggerFromContext(ctx)

	_, l = l.Start(ctx, "instance.handleEvent")
	defer l.End()

	rid := i.AddResource(ctx, t.Payload)
	defer i.RmvResource(ctx, rid)

	hdl, ok := i.handlers[t.Handler]
	if !ok {
		hdl = i.vmInstance.GetFunc(i.vmStore, t.Handler)
		if hdl == nil {
			return &common.EventHandleResult{Code: wasm.ResultStatusCode_UnexportedHandler}
		}
		i.handlers[t.Handler] = hdl
	}

	l.Info("call hanlder:%s", t.Handler)

	result, err := hdl.Call(i.vmStore, int32(rid))
	if err != nil {
		l.Error(err)
		return &common.EventHandleResult{Code: wasm.ResultStatusCode_Failed}
	}

	return &common.EventHandleResult{Code: wasm.ResultStatusCode(result.(int32))}
}

const MaxUint = ^uint32(0)
const MaxInt = int(MaxUint >> 1)

func (i *Instance) AddResource(ctx context.Context, data []byte) uint32 {
	l := types.MustLoggerFromContext(ctx)

	_, l = l.Start(ctx, "instance.AddResource")
	defer l.End()

	var id = int32(uuid.New().ID() % uint32(MaxInt))
	i.res.Store(uint32(id), data)

	l.WithValues("res_id", id, "payload", string(data)).Info("added")

	return uint32(id)
}

func (i *Instance) GetResource(id uint32) ([]byte, bool) { return i.res.Load(id) }

func (i *Instance) RmvResource(ctx context.Context, id uint32) {
	l := types.MustLoggerFromContext(ctx)

	_, l = l.Start(ctx, "instance.RmvResource")
	defer l.End()

	i.res.Remove(id)
	l.WithValues("res_id", id).Info("removed")

}

func (i *Instance) Get(k string) int32 {
	data := i.db[k]
	var ret int32
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return ret
}
