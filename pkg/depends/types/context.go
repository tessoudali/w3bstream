package types

import "context"

type VmContext interface {
	OnVmStart() VmStartStatus
	NewPluginContext() PluginContext
}

type PluginContext interface {
	OnPluginStart() PluginStartStatus
	OnPluginDone() bool
}

type TransporterContext interface {
	OnConnected()
	OnEvent()
	OnEventHandled()
	NewTransporter()
}

type Runtime interface {
	Module(name string)
	Instantiate(ctx context.Context, conf interface{}) (mod interface{}, err error)
	Close(ctx context.Context, code uint32) error
}
