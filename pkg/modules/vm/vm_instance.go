package vm

import (
	w "github.com/wasmerio/wasmer-go/wasmer"
)

type Wasm struct {
	Code     []byte
	Engine   *w.Engine
	Store    *w.Store
	Module   *w.Module
	Instance *w.Instance
}

type NativeFunc = func(args []w.Value) ([]w.Value, error)

type WasmImportFunc struct {
	Name        string
	InputTypes  []w.ValueKind
	OutputTypes []w.ValueKind
	NativeFunc  NativeFunc
}

type WasmImport struct {
	Namespace string
	Functions []WasmImportFunc
}

func NewWasm(code []byte, imports []WasmImport) (*Wasm, error) {
	engine := w.NewEngine()
	store := w.NewStore(engine)
	module, _ := w.NewModule(store, code)

	importObject := w.NewImportObject()

	for _, v := range imports {
		intoExtern := make(map[string]w.IntoExtern)
		for _, fn := range v.Functions {
			intoExtern[fn.Name] = w.NewFunction(
				store,
				w.NewFunctionType(
					w.NewValueTypes(fn.InputTypes...),
					w.NewValueTypes(fn.OutputTypes...),
				),
				fn.NativeFunc,
			)
		}
		importObject.Register(
			v.Namespace,
			intoExtern,
		)
	}

	instance, e := w.NewInstance(module, importObject)
	if e != nil {
		return nil, e
	}

	return &Wasm{
		Code:     code,
		Engine:   engine,
		Store:    store,
		Module:   module,
		Instance: instance,
	}, nil
}

func (wasm *Wasm) GetFunction(name string) (w.NativeFunction, error) {
	return wasm.Instance.Exports.GetFunction(name)
}

func (wasm *Wasm) ExecuteFunction(name string, args ...interface{}) (interface{}, error) {
	fn, e := wasm.Instance.Exports.GetFunction(name)
	if e != nil {
		return nil, e
	}
	return fn(args...)
}

func (wasm *Wasm) GetMemory(name string) ([]byte, error) {
	memory, e := wasm.Instance.Exports.GetMemory(name)
	if e != nil {
		return nil, e
	}
	return memory.Data(), nil
}
