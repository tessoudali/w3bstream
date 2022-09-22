// event dispatch

package proxy

import "github.com/iotexproject/w3bstream/pkg/modules/event"

type dispatcher struct {
}

func (d *dispatcher) dispatch(e event.Event) ([]byte, error) {
	return []byte("Success"), nil
}
