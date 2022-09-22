package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed testdata/word_count.wasm
var code []byte

func main() {
	ctx := context.Background()

	// new wasm runtime.
	r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfig().
		WithFeatureBulkMemoryOperations(true).
		WithFeatureNonTrappingFloatToIntConversion(true).
		WithFeatureSignExtensionOps(true).WithFeatureMultiValue(true))
	defer r.Close(ctx)

	// exports
	{
		_, err := r.NewModuleBuilder("env").
			ExportFunction("log", log).
			ExportFunction("inc", inc).
			ExportFunction("get", get).
			Instantiate(ctx, r)
		if err != nil {
			panic(err)
		}
	}

	if _, err := wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
		panic(err)
	}

	mod, err := r.InstantiateModuleFromBinary(ctx, code)
	if err != nil {
		panic(err)
	}

	// fns
	var (
		counter = mod.ExportedFunction("start")
		_       = mod.ExportedFunction("greet")
		malloc  = mod.ExportedFunction("malloc")
		free    = mod.ExportedFunction("free")
	)

	// event handle
	{
		str := os.Args[1]
		strlen := uint64(len(str))

		results, err := malloc.Call(ctx, strlen)
		if err != nil {
			panic(err)
		}
		ptr := results[0]
		defer free.Call(ctx, ptr)
		// free.Call(ctx, ptr2)

		if !mod.Memory().Write(ctx, uint32(ptr), []byte(str)) {
			panic(fmt.Sprintf("Memory.Write(%d, %d) out of range of memory size %d",
				ptr, strlen, mod.Memory().Size(ctx)))
		}

		results, err = counter.Call(ctx, ptr, strlen)
		if err != nil {
			panic(err)
		}

		msg, _ := json.Marshal(words)
		fmt.Println("host >> " + string(msg))
		fmt.Println(results)
	}
}

func log(ctx context.Context, m api.Module, offset, size uint32) {
	buf, ok := m.Memory().Read(ctx, offset, size)
	if !ok {
		panic(fmt.Sprintf("Memory.Read(%d,%d) out of range)", offset, size))
	}
	fmt.Println(string(buf))
}

var words = make(map[string]int32)

func inc(ctx context.Context, m api.Module, offset, size uint32, delta int32) (code int32) {
	buf, ok := m.Memory().Read(ctx, offset, size)
	if !ok {
		return 1
	}
	str := string(buf)
	if _, ok := words[str]; !ok {
		words[str] = delta
	} else {
		words[str] = words[str] + delta
	}
	return 0
}

func get(ctx context.Context, m api.Module, offset, size uint32) (count int32) {
	buf, ok := m.Memory().Read(ctx, offset, size)
	if !ok {
		return 1
	}
	str := string(buf)
	if _, ok := words[str]; !ok {
		return 0
	}
	return words[str]
}
