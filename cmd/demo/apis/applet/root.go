package applet

import (
	"github.com/iotexproject/Bumblebee/kit/httptransport"
	"github.com/iotexproject/Bumblebee/kit/kit"
)

var Root = kit.NewRouter(httptransport.Group("/applet"))

func init() {
	Root.Register(kit.NewRouter(&CreateApplet{}))
	Root.Register(kit.NewRouter(&ListApplet{}))
}
