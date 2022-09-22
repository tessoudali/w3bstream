package v1

import (
	"github.com/google/uuid"
	"github.com/iotexproject/Bumblebee/x/mapx"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
)

var instances = mapx.New[uint32, wasm.Instance]()

func AddInstance(i wasm.Instance) uint32 {
	id := uuid.New().ID()
	instances.Store(id, i)
	return id
}

func DelInstance(id uint32) {
	i, _ := instances.LoadAndRemove(id)
	if i != nil && i.State() == wasm.InstanceState_Started {
		i.Stop()
	}
}

func GetInstance(id uint32) wasm.Instance {
	i, _ := instances.Load(id)
	return i
}

func RunInstance(id uint32) (err error) {
	i := GetInstance(id)
	if i == nil {
		return nil // return not found error
	}
	go func() {
		i.Start()
		i.Stop()
	}()
	return nil
}
