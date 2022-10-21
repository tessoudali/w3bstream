package wasm

import (
	"context"

	"github.com/iotexproject/w3bstream/pkg/enums"
)

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
	Start(context.Context) error
	Stop()
	State() enums.InstanceState
	AddResource(context.Context, []byte) uint32
	RmvResource(context.Context, uint32)
	GetResource(uint32) ([]byte, bool)
	Get(k string) int32

	EventConsumer
}

type EventConsumer interface {
	HandleEvent(ctx context.Context, handler string, payload []byte) ([]byte, ResultStatusCode, error)
}

type KVStore interface {
	Get(string) int32
}

type ETHClientConfig struct {
	PrivateKey    string `env:""`
	ChainEndpoint string `env:""`
}
