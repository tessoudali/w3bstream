package main

import (
	"flag"
	"fmt"

	confapp "github.com/iotexproject/Bumblebee/conf/app"
	conflog "github.com/iotexproject/Bumblebee/conf/log"
	confmqtt "github.com/iotexproject/Bumblebee/conf/mqtt"
)

var (
	mqtt   = &confmqtt.Broker{}
	logger = conflog.Std()

	App *confapp.Ctx
)

func init() {
	App = confapp.New(
		confapp.WithName("mock-mqtt-client"),
		confapp.WithLogger(logger),
		confapp.WithRoot("."),
	)
	App.Conf(mqtt)

	flag.StringVar(&id, "i", "", "publish client id")
	flag.StringVar(&topic, "t", "", "publish topic")
	flag.StringVar(&payload, "c", "", "publish content")
	flag.Parse()
}

var (
	id      string
	topic   string
	payload string
)

func main() {
	c, err := mqtt.Client(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.WithTopic(topic).Publish(payload)
	if err != nil {
		fmt.Println(err)
		return
	}
}
