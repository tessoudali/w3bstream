package vm

import (
	conflog "github.com/iotexproject/Bumblebee/conf/log"
	"github.com/tetratelabs/wazero"

	"github.com/iotexproject/w3bstream/pkg/types/wasm"
)

type InstanceOption struct {
	RuntimeConfig   wazero.RuntimeConfig
	Logger          conflog.Logger
	Tasks           TaskReader
	OnStatusChanged func() // should call this when instance runtime interrupted
}

type InstanceOptionSetter func(o *InstanceOption)

func InstanceOptionWithRuntimeConfig(rc wazero.RuntimeConfig) InstanceOptionSetter {
	return func(o *InstanceOption) { o.RuntimeConfig = rc }
}

func InstanceOptionWithLogger(l conflog.Logger) InstanceOptionSetter {
	return func(o *InstanceOption) { o.Logger = l }
}

var (
	DefaultRuntimeConfig = wazero.NewRuntimeConfig().
				WithFeatureBulkMemoryOperations(true).
				WithFeatureNonTrappingFloatToIntConversion(true).
				WithFeatureSignExtensionOps(true).
				WithFeatureMultiValue(true)
	DefaultLogger               = conflog.Std()
	DefaultInstanceOptionSetter = func(o *InstanceOption) {
		o.Logger, o.RuntimeConfig = DefaultLogger, DefaultRuntimeConfig
	}
)

type TaskReader interface {
	Wait() <-chan Task
}

type Task struct {
	Payload []byte
	Res     chan<- EventHandleResult
}

type EventHandleResult struct {
	Response []byte
	Code     wasm.ResultStatusCode
}
