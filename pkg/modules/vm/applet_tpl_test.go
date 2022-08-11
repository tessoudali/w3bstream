package vm

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
	w "github.com/wasmerio/wasmer-go/wasmer"
)

var (
	conf *Config
)

func init() {
	path, err := filepath.Abs("./testdata/build/applet.yaml")
	if err != nil {
		panic(err)
	}
	conf, err = LoadConfigFrom(path)
	if err != nil {
		panic(err)
	}
}

func TestNewWasm2(t *testing.T) {
	var wasm *Wasm
	var filename = filepath.Join("./testdata/build", conf.DataSources[0].Mapping.File)

	code, err := ioutil.ReadFile(filename)
	NewWithT(t).Expect(err).To(BeNil())

	imports := []WasmImport{
		{
			namespace: "env",
			functions: []WasmImportFunc{
				{
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
				{
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
		}, {
			namespace: "conversion",
		},
	}

	wasm, e := NewWasm(code, imports)
	NewWithT(t).Expect(e).To(BeNil())

	// sum, e := wasm.ExecuteFunction("add", 1, 2)
	// NewWithT(t).Expect(e).To(BeNil())

	// v, ok := sum.(int32)
	// NewWithT(t).Expect(ok).To(BeTrue())
	// NewWithT(t).Expect(v).To(Equal(int32(3)))

	_, e = wasm.ExecuteFunction("hello")
	NewWithT(t).Expect(e).To(BeNil())
	handlerNewGravatar, err := wasm.instance.Exports.GetFunction("handlerNewGravatar")
	NewWithT(t).Expect(err).To(BeNil())
	NewWithT(t).Expect(handlerNewGravatar).NotTo(BeNil())

	handleUpdatedGravatar, err := wasm.instance.Exports.GetFunction("handleUpdatedGravatar")
	NewWithT(t).Expect(err).To(BeNil())
	NewWithT(t).Expect(handleUpdatedGravatar).NotTo(BeNil())
}
