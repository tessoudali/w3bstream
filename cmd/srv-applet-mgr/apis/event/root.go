package event

import (
	"github.com/machinefi/Bumblebee/kit/httptransport"
	"github.com/machinefi/Bumblebee/kit/kit"
)

var Root = kit.NewRouter(httptransport.Group("/event"))

func init() {
	Root.Register(kit.NewRouter(&HandleEvent{}))
}
