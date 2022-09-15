package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/iotexproject/Bumblebee/conf/log"

	"github.com/iotexproject/w3bstream/cmd/srv-event-gw/http"
	"github.com/iotexproject/w3bstream/pkg/modules/event"
)

func main() {
	logger := log.Std()
	events := make(chan event.Event, 10)
	proxy := event.NewProxifier(logger)

	go proxy.Proxy(events)
	go http.Run(events, logger)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}
