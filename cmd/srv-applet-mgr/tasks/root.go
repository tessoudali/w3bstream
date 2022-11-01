package tasks

import "github.com/machinefi/Bumblebee/kit/kit"

var Root = kit.NewRouter()

func init() {
	Root.Register(kit.NewRouter(&HandleEvent{}))
}
