package login

import (
	"github.com/iotexproject/Bumblebee/kit/httptransport"
	"github.com/iotexproject/Bumblebee/kit/kit"
)

var Root = kit.NewRouter(httptransport.Group("/"))

func init() {
	Root.Register(kit.NewRouter(&Login{}))
}
