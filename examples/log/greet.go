package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

// greetWasm was compiled using `tinygo build -o greet.wasm -scheduler=none --no-debug -target=wasi greet.go`
//
//go:embed testdata/test.wasm
var code []byte

// main shows how to interact with a WebAssembly function that was compiled
// from TinyGo.
//
// See README.md for a full description.
func main() {
	ctx := context.Background()

	c := wazero.NewRuntimeConfig().
		WithFeatureBulkMemoryOperations(true).
		WithFeatureNonTrappingFloatToIntConversion(true).
		WithFeatureSignExtensionOps(true)

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntimeWithConfig(ctx, c)
	defer r.Close(ctx) // This closes everything this Runtime created.

	_, err := r.NewModuleBuilder("env").
		ExportFunction("log", logString).
		ExportFunction("get_data_2", getData).
		// ExportFunction("set_uint32", setUint32).
		Instantiate(ctx, r)

	// _, err := r.NewModuleBuilder("env").
	// 	ExportFunction("log", logString).
	// 	ExportFunction("get_data", getData).
	// 	ExportFunction("set_uint32", setUint32).
	// 	Instantiate(ctx, r)

	if err != nil {
		log.Panicln(err)
	}

	if _, err = wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
		log.Panicln(err)
	}

	mod, err := r.InstantiateModuleFromBinary(ctx, code)
	if err != nil {
		log.Panicln(err)
	}

	arg = "xxxxx"

	start := mod.ExportedFunction("start")
	_, err = start.Call(ctx, 100)
	if err != nil {
		log.Panicln(err)
	}

	// startSetUint32 := mod.ExportedFunction("startSetUint32")
	// _, err = startSetUint32.Call(ctx, 100)
	// if err != nil {
	// 	log.Panicln(err)
	// }

	// startLogByRid := mod.ExportedFunction("startLogByRid")
	// _, err = startLogByRid.Call(ctx, 101)
	// if err != nil {
	// 	log.Panicln(err)
	// }
}

var arg string // os.Arg[1]

func getData(ctx context.Context, m api.Module, rid int32, vmDataAddrPtr, vmDataSizePtr int32) {
	malloc := m.ExportedFunction("malloc")

	if malloc == nil {
		log.Panicln("malloc ")
	}

	data := arg
	size := len(data)
	result, err := malloc.Call(ctx, uint64(size))
	if err != nil {
		log.Panicln(err, result)
	}
	// addr := results[0]

	// // The pointer is a linear memory offset, which is where we write the name.
	// if !m.Memory().Write(ctx, uint32(addr), []byte(data)) {
	// 	log.Panicf(
	// 		"Memory.Write(%d, %d) out of range of memory size %d",
	// 		addr, size, m.Memory().Size(ctx),
	// 	)
	// }

	// fmt.Printf("host >> addr=%d\n", addr)
	// fmt.Printf("host >> size=%d\n", size)
	// fmt.Printf("host >> vmDataAddrPtr=%d\n", vmDataAddrPtr)
	// fmt.Printf("host >> vmDataSizePtr=%d\n", vmDataSizePtr)

	// m.Memory().WriteUint32Le(ctx, uint32(vmDataAddrPtr), uint32(addr))
	// m.Memory().WriteUint32Le(ctx, uint32(vmDataSizePtr), uint32(size))

	fmt.Println("host >> get_data returned")
}

func setUint32(ctx context.Context, m api.Module, addr, value uint32) int32 {
	m.Memory().WriteUint32Le(ctx, addr, value)
	return 0
}

func logString(ctx context.Context, m api.Module, offset, byteCount uint32) {
	buf, ok := m.Memory().Read(ctx, uint32(offset), uint32(byteCount))
	if !ok {
		log.Panicf("Memory.Read(%d, %d) out of range", offset, byteCount)
	}
	fmt.Println(string(buf))
}
