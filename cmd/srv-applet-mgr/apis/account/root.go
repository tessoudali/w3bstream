package account

import (
	"github.com/iotexproject/Bumblebee/kit/httptransport"
	"github.com/iotexproject/Bumblebee/kit/kit"
)

var Root = kit.NewRouter(httptransport.Group("/account"))

func init() {
	Root.Register(kit.NewRouter(&CreateAccount{}))
	Root.Register(kit.NewRouter(&UpdatePasswordByAccountID{}))
}
