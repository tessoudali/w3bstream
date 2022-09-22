package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Note: In TinyGo "//export" on a func is actually an import!

// main is required for TinyGo to compile to Wasm.
func main() {}

// _log is a wasm import, prints a string
//
//go:wasm-module env
//export log
func _log(ptr uint32, size uint32)

//go:wasm-module env
//export get_data
func _get_data(rid uint32, ptr uint32, size uint32)

//export start
func _start(rid uint32) {
	log(fmt.Sprintf("wasm >> received: %d", rid))
	message, err := getData(rid)
	if err != nil {
		log("wasm >> error" + err.Error())
	}
	log("wasm >> " + message)
}

// log a message to the console using _log.
func log(message string) {
	ptr, size := stringToPtr(message)
	_log(ptr, size)
}

func getData(rid uint32) (string, error) {
	addr := uintptr(unsafe.Pointer(new(uint32)))
	size := uintptr(unsafe.Pointer(new(uint32)))

	_get_data(rid, uint32(addr), uint32(size))

	log(fmt.Sprintf("wasm >> addr=%d", addr))
	log(fmt.Sprintf("wasm >> size=%d", size))

	vAddr := (*uint32)(unsafe.Pointer(addr))
	vSize := (*uint32)(unsafe.Pointer(size))

	log(fmt.Sprintf("wasm >> *vaddr=%d", *vAddr))
	log(fmt.Sprintf("wasm >> *vsize=%d", *vSize))

	return ptrToString(*vAddr, *vSize), nil
}

func ptrToString(ptr uint32, size uint32) string {
	return *(*string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  uintptr(size),
		Cap:  uintptr(size),
	}))
}

func stringToPtr(s string) (uint32, uint32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}
