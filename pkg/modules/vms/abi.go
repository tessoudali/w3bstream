package vms

type Imports interface {
	Malloc()
	Free()
}
