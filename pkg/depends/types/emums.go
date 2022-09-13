package types

type VmStartStatus uint8

const (
	VmStartStatusOK VmStartStatus = iota + 1
	VmStartStatusFailed
)

type PluginStartStatus uint8

const (
	PluginStartStatusOK PluginStartStatus = iota + 1
	PluginStartStatusFailed
)
