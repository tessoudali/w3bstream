package vm

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/iotexproject/Bumblebee/x/mapx"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/pkg/enums"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
)

var instances = mapx.New[string, wasm.Instance]()

var (
	ErrNotFound = errors.New("instance not found")
)

func AddInstance(i wasm.Instance) string {
	id := uuid.New().String()
	instances.Store(id, i)
	fmt.Printf("--- %s created\n", id)
	return id
}

func changeID(oldID, newID string) {
	i, _ := instances.LoadAndRemove(oldID)
	instances.Store(newID, i)
}

func DelInstance(id string) error {
	i, _ := instances.LoadAndRemove(id)
	if i != nil && i.State() == enums.INSTANCE_STATE__STARTED {
		i.Stop()
	}
	fmt.Printf("--- %s deleted\n", id)
	return nil
}

func StartInstance(id string) error {
	i, ok := instances.Load(id)
	if !ok {
		return ErrNotFound
	}
	go func() {
		if err := i.Start(); err != nil {
		}
	}()

	fmt.Printf("--- %s started\n", id)
	return nil
}

func StopInstance(id string) error {
	i, ok := instances.Load(id)
	if !ok {
		return ErrNotFound
	}
	i.Stop()
	return nil
}

func GetInstanceState(id string) (enums.InstanceState, bool) {
	i, ok := instances.Load(id)
	if !ok {
		return enums.INSTANCE_STATE_UNKNOWN, false
	}
	return i.State(), true
}

func GetConsumer(id string) wasm.EventConsumer {
	i, ok := instances.Load(id)
	if !ok {
		return nil
	}
	return i.(wasm.EventConsumer)
}
