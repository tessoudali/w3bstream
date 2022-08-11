package main

import (
	"sync"
	"time"

	"github.com/iotexproject/Bumblebee/kit/kit"

	"github.com/iotexproject/w3bstream/cmd/demo/apis"
	"github.com/iotexproject/w3bstream/cmd/demo/global"
)

var app = global.App

func main() {
	app.AddCommand("migrate", func(args ...string) {
		global.Migrate()
	})

	// TODO should add batch routines/daemons to app context
	app.Execute(func(args ...string) {
		BatchRun(
			func() {
				kit.Run(apis.RouterRoot, global.Server())
			},
		)
	})
}

func BatchRun(commands ...func()) {
	wg := &sync.WaitGroup{}

	for i := range commands {
		cmd := commands[i]
		wg.Add(1)

		go func() {
			defer wg.Done()
			cmd()
			time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait()
}
