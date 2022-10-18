package wazero_test

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/iotexproject/w3bstream/pkg/modules/vm"
	"github.com/iotexproject/w3bstream/pkg/modules/vm/common"
	"github.com/iotexproject/w3bstream/pkg/modules/vm/wazero"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
)

// //go:embed ../../../examples/log/log.wasm
var wasmLogCode []byte

// //go:embed ../../../examples/gjson/parse_json.wasm
var wasmGJsonCode []byte

// //go:embed ../../../examples/easyjson/parse_json.wasm
var wasmEasyJsonCode []byte

// //go:embed ../../../examples/word_count/word_count.wasm
var wasmWordCountCode []byte

// //go:embed ../../../examples/word_count_v2/word_count_v2.wasm
var wasmWordCountV2Code []byte

func init() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	root := filepath.Join(wd, "../../../../examples")
	fmt.Println(root)

	var err error
	wasmLogCode, err = os.ReadFile(filepath.Join(root, "log/log.wasm"))
	if err != nil {
		panic(err)
	}

	wasmGJsonCode, err = os.ReadFile(filepath.Join(root, "gjson/parse_json.wasm"))
	if err != nil {
		panic(err)
	}
	wasmEasyJsonCode, err = os.ReadFile(filepath.Join(root, "easyjson/parse_json.wasm"))
	if err != nil {
		panic(err)
	}
	wasmWordCountCode, err = os.ReadFile(filepath.Join(root, "word_count/word_count.wasm"))
	if err != nil {
		panic(err)
	}
	wasmWordCountV2Code, err = os.ReadFile(filepath.Join(root, "word_count_v2/word_count_v2.wasm"))
	if err != nil {
		panic(err)
	}
}

func TestInstance_LogWASM(t *testing.T) {
	i, err := wazero.NewInstanceByCode(wasmLogCode, common.DefaultInstanceOptionSetter)
	NewWithT(t).Expect(err).To(BeNil())
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	NewWithT(t).Expect(err).To(BeNil())
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("start", []byte("IoTeX"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	_, code = i.HandleEvent("not_exported", []byte("IoTeX"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_UnexportedHandler))
}

func TestInstance_GJsonWASM(t *testing.T) {
	i, err := wazero.NewInstanceByCode(wasmGJsonCode, common.DefaultInstanceOptionSetter)
	NewWithT(t).Expect(err).To(BeNil())
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	NewWithT(t).Expect(err).To(BeNil())
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("start", []byte(`{
  "name": {"first": "Tom", "last": "Anderson", "age": 39},
  "friends": [
    {"first_name": "Dale", "last_name": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first_name": "Roger", "last_name": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first_name": "Jane", "last_name": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))
}

func TestInstance_EasyJsonWASM(t *testing.T) {
	i, err := wazero.NewInstanceByCode(wasmEasyJsonCode, common.DefaultInstanceOptionSetter)
	NewWithT(t).Expect(err).To(BeNil())
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	NewWithT(t).Expect(err).To(BeNil())
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("start", []byte(`{"id":11,"student_name":"Tom","student_school":
								{"school_name":"MIT","school_addr":"xz"},
								"birthday":"2017-08-04T20:58:07.9894603+08:00"}`))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))
}

func TestInstance_WordCount(t *testing.T) {
	i, err := wazero.NewInstanceByCode(wasmWordCountCode, common.DefaultInstanceOptionSetter)
	NewWithT(t).Expect(err).To(BeNil())
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	NewWithT(t).Expect(err).To(BeNil())
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("start", []byte("a b c d a"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	NewWithT(t).Expect(i.Get("a")).To(Equal(int32(2)))
	NewWithT(t).Expect(i.Get("b")).To(Equal(int32(1)))
	NewWithT(t).Expect(i.Get("c")).To(Equal(int32(1)))
	NewWithT(t).Expect(i.Get("d")).To(Equal(int32(1)))

	_, code = i.HandleEvent("start", []byte("a b c d a"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	NewWithT(t).Expect(i.Get("a")).To(Equal(int32(4)))
	NewWithT(t).Expect(i.Get("b")).To(Equal(int32(2)))
	NewWithT(t).Expect(i.Get("c")).To(Equal(int32(2)))
	NewWithT(t).Expect(i.Get("d")).To(Equal(int32(2)))
}

func TestInstance_WordCountV2(t *testing.T) {
	i, err := wazero.NewInstanceByCode(wasmWordCountV2Code, common.DefaultInstanceOptionSetter)
	NewWithT(t).Expect(err).To(BeNil())
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	NewWithT(t).Expect(err).To(BeNil())
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("count", []byte("a b c d a"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	NewWithT(t).Expect(i.Get("a")).To(Equal(int32(2)))
	NewWithT(t).Expect(i.Get("b")).To(Equal(int32(1)))
	NewWithT(t).Expect(i.Get("c")).To(Equal(int32(1)))
	NewWithT(t).Expect(i.Get("d")).To(Equal(int32(1)))

	_, code = i.HandleEvent("count", []byte("a b c d a"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	NewWithT(t).Expect(i.Get("a")).To(Equal(int32(4)))
	NewWithT(t).Expect(i.Get("b")).To(Equal(int32(2)))
	NewWithT(t).Expect(i.Get("c")).To(Equal(int32(2)))
	NewWithT(t).Expect(i.Get("d")).To(Equal(int32(2)))

	_, unique := i.HandleEvent("unique", nil)
	NewWithT(t).Expect(unique).To(Equal(wasm.ResultStatusCode(4)))
}
