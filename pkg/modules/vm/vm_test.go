package vm_test

import (
	"path/filepath"
	"runtime"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/iotexproject/w3bstream/pkg/modules/vm"
)

func TestNewWasm(t *testing.T) {
	_, current, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(current), "../testdata/simple")

	w, err := vm.Load(root)
	NewWithT(t).Expect(err).To(BeNil())

	sum, e := w.ExecuteFunction("add", 1, 2)
	NewWithT(t).Expect(e).To(BeNil())

	v, ok := sum.(int32)
	NewWithT(t).Expect(ok).To(BeTrue())
	NewWithT(t).Expect(v).To(Equal(int32(3)))

	result, e := w.ExecuteFunction("hello")
	NewWithT(t).Expect(e).To(BeNil())
	NewWithT(t).Expect(result).To(BeNil())
}
