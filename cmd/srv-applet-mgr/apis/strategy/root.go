package strategy

import (
	"github.com/iotexproject/Bumblebee/kit/httptransport"
	"github.com/iotexproject/Bumblebee/kit/kit"
)

var Root = kit.NewRouter(httptransport.Group("/strategy"))

func init() {
	Root.Register(kit.NewRouter(&CreateStrategy{}))
	Root.Register(kit.NewRouter(&UpdateStrategy{}))
	Root.Register(kit.NewRouter(&GetStrategy{}))
	Root.Register(kit.NewRouter(&ListStrategy{}))
	Root.Register(kit.NewRouter(&RemoveStrategy{}))
}
