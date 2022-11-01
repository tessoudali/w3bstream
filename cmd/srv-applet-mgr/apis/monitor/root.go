package monitor

import (
	"github.com/machinefi/Bumblebee/kit/httptransport"
	"github.com/machinefi/Bumblebee/kit/kit"
)

var Root = kit.NewRouter(httptransport.Group("/monitor"))

func init() {
	Root.Register(kit.NewRouter(&CreateMonitor{}))
	Root.Register(kit.NewRouter(&RemoveMonitor{}))
}
