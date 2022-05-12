package vm

import (
	w "github.com/wasmerio/wasmer-go/wasmer"
)

type Wasm struct {
	code	 []byte
	engine	 *w.Engine
	store	 *w.Store
	module	 *w.Module
	instance *w.Instance
}

type NativeFunc = func(args []w.Value) ([]w.Value, error)

type WasmImport struct {
	name			string
	inputTypes		[]w.ValueKind
	outputTypes		[]w.ValueKind
	nativeFunc		NativeFunc
}

func NewWasm(code []byte, imports []WasmImport) (*Wasm, error) {
	var instance *w.Instance

	engine := w.NewEngine()
	store := w.NewStore(engine)
	module, _ := w.NewModule(store, code)

	importObject := w.NewImportObject()

	intoExtern := make(map[string]w.IntoExtern)
	for _, v := range imports {
		intoExtern[v.name] = w.NewFunction(
			store,
			w.NewFunctionType(w.NewValueTypes(v.inputTypes...), w.NewValueTypes(v.outputTypes...)),
			v.nativeFunc,
		)
	}

	importObject.Register(
		"env",
		intoExtern,
	)

	instance, e := w.NewInstance(module, importObject)
	if e != nil {
		return nil, e
	}

	return &Wasm{
		code: 	  code,
		engine:   engine,
		store: 	  store,
		module:   module,
		instance: instance,
	}, nil
}