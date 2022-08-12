package main

import (
	"context"
	"sync"
	"time"

	"github.com/iotexproject/Bumblebee/kit/kit"

	"github.com/iotexproject/w3bstream/pkg/modules/applet_deploy"

	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/apis"
	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/global"
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
			func() {
				ctx := context.Background()
				ctx = global.WithDatabaseContext(ctx)
				ctx = global.WithMqttContext(ctx)
				ctx = global.WithLoggerContext(ctx)
				if err := applet_deploy.StartAppletVMs(ctx); err != nil {
					panic(err)
				}
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
