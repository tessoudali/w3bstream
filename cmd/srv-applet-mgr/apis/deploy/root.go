package deploy

import (
	"github.com/iotexproject/Bumblebee/kit/httptransport"
	"github.com/iotexproject/Bumblebee/kit/kit"
)

var Root = kit.NewRouter(httptransport.Group("/deploy"))

func init() {
	Root.Register(kit.NewRouter(&CreateDeploy{}))
	Root.Register(kit.NewRouter(&ListDeploy{}))
	Root.Register(kit.NewRouter(&RemoveDeployByAppletIDAndVersion{}))
	Root.Register(kit.NewRouter(&RemoveDeployByDeployID{}))
}
