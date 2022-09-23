package vm

import (
	"github.com/google/uuid"
	"github.com/iotexproject/Bumblebee/x/mapx"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/pkg/enums"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
)

var instances = mapx.New[uint32, wasm.Instance]()

var (
	ErrNotFound = errors.New("instance not found")
)

func AddInstance(i wasm.Instance) uint32 {
	id := uuid.New().ID()
	instances.Store(id, i)
	return id
}

func DelInstance(id uint32) error {
	i, _ := instances.LoadAndRemove(id)
	if i != nil && i.State() == enums.INSTANCE_STATE__STARTED {
		i.Stop()
	}
	return nil
}

func StartInstance(id uint32) error {
	i, ok := instances.Load(id)
	if !ok {
		return ErrNotFound
	}
	go func() {
		if err := i.Start(); err != nil {
		}
	}()
	return nil
}

func StopInstance(id uint32) error {
	i, ok := instances.Load(id)
	if !ok {
		return ErrNotFound
	}
	i.Stop()
	return nil
}

func GetInstanceState(id uint32) (enums.InstanceState, bool) {
	i, ok := instances.Load(id)
	if !ok {
		return enums.INSTANCE_STATE_UNKNOWN, false
	}
	return i.State(), true
}

func GetConsumer(id uint32) wasm.EventConsumer {
	i, ok := instances.Load(id)
	if !ok {
		return nil
	}
	return i.(wasm.EventConsumer)
}
