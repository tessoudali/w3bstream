package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {}

// _log is a host interface for wasm importing
//
//go:wasm-module env
//export log
func _log(ptr uint32, size uint32)

// add is wasm import, for add word count
//
//go:wasm-module env
//export add
func _add(ptr uint32, size uint32) (code int32)

// get is wasm import, for get word count
//
//go:wasm-module env
//export get
func _get(ptr uint32, size uint32) (count int32)

// countWords wasm count words, accept a string pointer
//
//export countWords
func countWords(ptr, size uint32) int32 {
	str := ptrToString(ptr, size)
	words := strings.Split(str, " ")

	for _, word := range words {
		if code := add(word); code != 0 {
			return code
		}
		if _, ok := records[word]; ok {
			records[word] = 1
		} else {
			records[word]++
		}
		log(fmt.Sprintf("%s: %d", word, get(word)))
	}
	// msg, _ := json.Marshal(records)
	log(str)
	return 0
}

var records = make(map[string]int)

// add add word count
func add(key string) (code int32) {
	ptr, size := stringToPtr(key)
	return _add(ptr, size)
}

// log log string using _log
func log(str string) {
	ptr, size := stringToPtr("wasm >> " + str)
	_log(ptr, size)
}

func get(key string) (count int32) {
	ptr, size := stringToPtr(key)
	return _get(ptr, size)
}

// ptrToString returns a string from wasm compatible numeric types representing its pointer and length.
func ptrToString(ptr uint32, size uint32) string {
	return *(*string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  uintptr(size), // requires uintptr https://github.com/tinygo-org/tinygo/issues/1284
		Cap:  uintptr(size),
	}))
}

// stringToPtr returns a pointer and size pair for the given string
func stringToPtr(s string) (uint32, uint32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}
