package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/pkg/errors"
)

func main() {}

//go:wasm-module env
//export ws_log
func _ws_log(ptr uint32, size uint32)

//go:wasm-module env
//export ws_get_data
func _ws_get_data(rid uint32, ptr uint32, size uint32) int32

//go:wasm-module env
//export ws_get_db
func _ws_get_db(kaddr, ksize uint32) (v int32)

//go:wasm-module env
//export ws_set_db
func _ws_set_db(kaddr, ksize uint32, v int32)

//go:wasm-module env
//export ws_send_tx
func _ws_send_tx(kaddr, ksize uint32) (v int32)

var allocs = make(map[uintptr][]byte)

//export alloc
func alloc(size uintptr) unsafe.Pointer {
	buf := make([]byte, size)
	ptr := unsafe.Pointer(&buf[0])
	allocs[uintptr(ptr)] = buf
	return ptr
}

//export start
func _start(rid uint32) int32 {
	log(fmt.Sprintf("wasm >> received: %d", rid))
	str, err := getData(rid)
	if err != nil {
		log("wasm >> error" + err.Error())
		return -1
	}
	log(fmt.Sprintf("wasm >> receive message: %s", str))
	count := getValue("clicks") + 1
	if count%5 == 0 {
		/*
			ptr, size := stringToPtr(fmt.Sprintf(
				`{
					"to": "%s",
					"value": "0",
					"data": "40c10f19000000000000000000000000%s0000000000000000000000000000000000000000000000000de0b6b3a7640000"
				}`,
				"0x1ED83F5AD999262eC06Ed8f3B801e108024b3e9c",
				"97186a21fa8e7955c0f154f960d588c3aca44f14",
			))
			_ws_send_tx(ptr, size)
		*/
		log("wasm << send tx")
		// call _ws_send_tx
	}
	setValue("clicks", count)
	return 0
}

// log a message to the console using _log.
func log(message string) {
	ptr, size := stringToPtr(message)
	_ws_log(ptr, size)
}

func getValue(key string) int32 {
	kaddr, ksize := stringToPtr(key)
	return _ws_get_db(kaddr, ksize)
}

func setValue(key string, value int32) {
	kaddr, ksize := stringToPtr(key)
	_ws_set_db(kaddr, ksize, value)
}

func getData(rid uint32) (string, error) {
	addr := uintptr(unsafe.Pointer(new(uint32)))
	size := uintptr(unsafe.Pointer(new(uint32)))

	code := _ws_get_data(rid, uint32(addr), uint32(size))
	if code != 0 {
		return "", fmt.Errorf("get data failed: [rid:%d] [code:%d]", rid, code)
	}

	log(fmt.Sprintf("wasm.getData addr=%d", addr))
	log(fmt.Sprintf("wasm.getData size=%d", size))

	vAddr := (*uint32)(unsafe.Pointer(addr))
	vSize := (*uint32)(unsafe.Pointer(size))

	log(fmt.Sprintf("wasm.getData *vaddr=%d", *vAddr))
	log(fmt.Sprintf("wasm.getData *vsize=%d", *vSize))

	return ptrToString(*vAddr, *vSize), nil
}

func ptrToString(ptr uint32, size uint32) string {
	return *(*string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  uintptr(size),
		Cap:  uintptr(size),
	}))
}

func ptrToBytes(ptr uint32, size uint32) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  uintptr(size),
		Cap:  uintptr(size),
	}))
}

func ptrToInt32(ptr uint32, size uint32) (v int32, err error) {
	if size != 4 {
		return 0, errors.New("invalid data size")
	}
	buf := ptrToBytes(ptr, size)
	err = binary.Read(bytes.NewReader(buf), binary.LittleEndian, &v)
	return
}

func stringToPtr(s string) (uint32, uint32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}

func int32ToPtr(v int32) (uint32, uint32) {
	return uint32(uintptr(unsafe.Pointer(&v))), 4
}