package vm_test

import (
	_ "embed"
	"testing"

	"github.com/iotexproject/w3bstream/pkg/modules/vm"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
	. "github.com/onsi/gomega"
)

//go:embed testdata/log/log.wasm
var wasmLogCode []byte

//go:embed testdata/word_count/word_count.wasm
var wasmWordCountCode []byte

//go:embed testdata/word_count_v2/word_count_v2.wasm
var wasmWordCountV2Code []byte

func TestInstance_LogWASM(t *testing.T) {
	id, err := vm.NewInstanceByCode(wasmLogCode, vm.DefaultInstanceOptionSetter)
	NewWithT(t).Expect(err).To(BeNil())

	i := vm.GetConsumer(id)
	NewWithT(t).Expect(i).NotTo(BeNil())

	err = vm.StartInstance(id)
	NewWithT(t).Expect(err).To(BeNil())

	defer vm.StopInstance(id)

	_, code := i.HandleEvent("start", []byte("IoTeX"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	_, code = i.HandleEvent("not_exported", []byte("IoTeX"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_UnexportedHandler))
}

func TestInstance_WordCount(t *testing.T) {
	id, err := vm.NewInstanceByCode(wasmWordCountCode, vm.DefaultInstanceOptionSetter)
	NewWithT(t).Expect(err).To(BeNil())

	i := vm.GetConsumer(id)
	NewWithT(t).Expect(i).NotTo(BeNil())

	err = vm.StartInstance(id)
	NewWithT(t).Expect(err).To(BeNil())

	defer vm.StopInstance(id)

	// store, ok := i.(wasm.KVStore)
	// NewWithT(t).Expect(ok).To(BeTrue())

	_, code := i.HandleEvent("handler", []byte("qqqqq"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	// _, code := i.HandleEvent("start", []byte("a b c d a"))
	// NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	// NewWithT(t).Expect(store.Get("a")).To(Equal(int32(2)))
	// NewWithT(t).Expect(store.Get("b")).To(Equal(int32(1)))
	// NewWithT(t).Expect(store.Get("c")).To(Equal(int32(1)))
	// NewWithT(t).Expect(store.Get("d")).To(Equal(int32(1)))

	// _, code = i.HandleEvent("start", []byte("a b c d a"))
	// NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	// NewWithT(t).Expect(store.Get("a")).To(Equal(int32(4)))
	// NewWithT(t).Expect(store.Get("b")).To(Equal(int32(2)))
	// NewWithT(t).Expect(store.Get("c")).To(Equal(int32(2)))
	// NewWithT(t).Expect(store.Get("d")).To(Equal(int32(2)))
}

func TestInstance_WordCountV2(t *testing.T) {
	id, err := vm.NewInstanceByCode(wasmWordCountV2Code, vm.DefaultInstanceOptionSetter)
	NewWithT(t).Expect(err).To(BeNil())

	i := vm.GetConsumer(id)
	NewWithT(t).Expect(i).NotTo(BeNil())

	err = vm.StartInstance(id)
	NewWithT(t).Expect(err).To(BeNil())

	defer vm.StopInstance(id)

	store, ok := i.(wasm.KVStore)
	NewWithT(t).Expect(ok).To(BeTrue())

	_, code := i.HandleEvent("count", []byte("a b c d a"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	NewWithT(t).Expect(store.Get("a")).To(Equal(int32(2)))
	NewWithT(t).Expect(store.Get("b")).To(Equal(int32(1)))
	NewWithT(t).Expect(store.Get("c")).To(Equal(int32(1)))
	NewWithT(t).Expect(store.Get("d")).To(Equal(int32(1)))

	_, code = i.HandleEvent("count", []byte("a b c d a"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	NewWithT(t).Expect(store.Get("a")).To(Equal(int32(4)))
	NewWithT(t).Expect(store.Get("b")).To(Equal(int32(2)))
	NewWithT(t).Expect(store.Get("c")).To(Equal(int32(2)))
	NewWithT(t).Expect(store.Get("d")).To(Equal(int32(2)))

	_, unique := i.HandleEvent("unique", nil)
	NewWithT(t).Expect(unique).To(Equal(wasm.ResultStatusCode(4)))
}
