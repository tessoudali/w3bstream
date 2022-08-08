package vm

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	w "github.com/wasmerio/wasmer-go/wasmer"
)

func TestNewWasm(t *testing.T) {
	var wasm *Wasm

	require := require.New(t)
	code, _ := ioutil.ReadFile("release.wasm")

	imports := []WasmImport{
		WasmImport{
			namespace: "env",
			functions: []WasmImportFunc{
				WasmImportFunc{
					name: "log",
					inputTypes: []w.ValueKind{
						w.I32,
					},
					outputTypes: []w.ValueKind{},
					nativeFunc: func(args []w.Value) ([]w.Value, error) {
						data, e := wasm.GetMemory("memory")
						if e != nil {
							return nil, e
						}
						fmt.Println(string(data[args[0].I32():]))
						return []w.Value{}, nil
					},
				},
				WasmImportFunc{
					name: "abort",
					inputTypes: []w.ValueKind{
						w.I32,
						w.I32,
						w.I32,
						w.I32,
					},
					outputTypes: []w.ValueKind{},
					nativeFunc: func(args []w.Value) ([]w.Value, error) {
						// TODO
						return []w.Value{}, nil
					},
				},
			},
		},
	}

	wasm, e := NewWasm(code, imports)
	require.NoError(e)

	sum, e := wasm.ExecuteFunction("add", 1, 2)
	require.NoError(e)

	v, ok := sum.(int32)
	require.Equal(ok, true)
	require.Equal(v, int32(3))

	_, e = wasm.ExecuteFunction("hello")
	require.NoError(e)
}