package project

import (
	"github.com/machinefi/Bumblebee/kit/httptransport"
	"github.com/machinefi/Bumblebee/kit/kit"
)

var Root = kit.NewRouter(httptransport.Group("/project"))

func init() {
	Root.Register(kit.NewRouter(&CreateProject{}))
	Root.Register(kit.NewRouter(&GetProjectByProjectID{}))
	Root.Register(kit.NewRouter(&ListProject{}))
	Root.Register(kit.NewRouter(&RemoveProject{}))
}
