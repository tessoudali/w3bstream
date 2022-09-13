package types

type VmContext interface {
	OnVmStart() VmStartStatus
}

type PluginContext interface {
	OnPluginStart()
	OnPluginDone()
}

type TransporterContext interface {
	OnConnected()
	OnEvent()
	OnEventHandled()
}
