package types

type VmStartStatus uint8

const (
	VmStartStatusOK     VmStartStatus = iota + 1
	VmStartStatusFailed VmStartStatus = iota + 1
)
