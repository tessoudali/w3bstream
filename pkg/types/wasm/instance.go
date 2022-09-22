package wasm

type VM interface {
	Name() string
	Init()
	NewModule(code []byte) Module
}

type Module interface {
	Init()
	NewInstance() Instance
	GetABI() []string
}

type Instance interface {
	Start() error
	Stop()
	State() InstanceState
	AddResource([]byte) uint32
	GetResource(uint32) ([]byte, bool)
	RmvResource(uint32)
}

type EventConsumer interface {
	HandleEvent([]byte) ([]byte, ResultStatusCode)
}
