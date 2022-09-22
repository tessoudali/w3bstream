package instance

import (
	"github.com/iotexproject/Bumblebee/kit/httptransport"
	"github.com/iotexproject/Bumblebee/kit/kit"
)

var Root = kit.NewRouter(httptransport.Group("/instance"))

func init() {
	Root.Register(kit.NewRouter(&CreateInstance{}))
}
