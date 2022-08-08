package apis

import (
	"github.com/iotexproject/Bumblebee/conf/http"
	"github.com/iotexproject/Bumblebee/kit/httptransport"
	"github.com/iotexproject/Bumblebee/kit/kit"

	"github.com/iotexproject/w3bstream/cmd/demo/apis/applet"
)

var (
	RouterRoot = kit.NewRouter(httptransport.Group("/"))
	RouterV0   = kit.NewRouter(httptransport.BasePath("/w3bstream/v0"))
)

func init() {
	RouterRoot.Register(http.LivenessRouter)
	RouterRoot.Register(RouterV0)

	RouterV0.Register(applet.Root)
}
