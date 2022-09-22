package v1

import (
	conflog "github.com/iotexproject/Bumblebee/conf/log"
	confmqtt "github.com/iotexproject/Bumblebee/conf/mqtt"
	"github.com/tetratelabs/wazero"
)

type InstanceOption struct {
	Channel       string
	RuntimeConfig wazero.RuntimeConfig
	Logger        conflog.Logger
	Broker        *confmqtt.Broker
	Client        *confmqtt.Client
}

type InstanceOptionSetter func(o *InstanceOption)

func InstanceOptionWithChannel(ch string) InstanceOptionSetter {
	return func(o *InstanceOption) { o.Channel = ch }
}

func InstanceOptionWithRuntimeConfig(rc wazero.RuntimeConfig) InstanceOptionSetter {
	return func(o *InstanceOption) { o.RuntimeConfig = rc }
}

func InstanceOptionWithLogger(l conflog.Logger) InstanceOptionSetter {
	return func(o *InstanceOption) { o.Logger = l }
}

func InstanceOptionWithMqttBroker(b *confmqtt.Broker) InstanceOptionSetter {
	return func(o *InstanceOption) { o.Broker = b }
}

var (
	DefaultRuntimeConfig = wazero.NewRuntimeConfig().
				WithFeatureBulkMemoryOperations(true).
				WithFeatureNonTrappingFloatToIntConversion(true).
				WithFeatureSignExtensionOps(true).
				WithFeatureMultiValue(true)
	DefaultLogger = conflog.Std()
)
