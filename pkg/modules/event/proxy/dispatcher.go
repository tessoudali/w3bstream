// event dispatch

package proxy

import (
	"context"
	"fmt"

	"github.com/iotexproject/w3bstream/pkg/modules/event"
	"github.com/iotexproject/w3bstream/pkg/modules/instance"
	"github.com/iotexproject/w3bstream/pkg/modules/vm"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
)

func dispatch(ctx context.Context, e event.Event) ([]byte, error) {
	ins, err := instance.GetInstanceByAppletID(ctx, e.Meta().AppletID)
	if err != nil {
		return nil, err
	}
	consumer := vm.GetConsumer(ins[0].InstanceVMID)
	res, code := consumer.HandleEvent(e.Raw())
	if code == wasm.ResultStatusCode_Failed {
		return nil, fmt.Errorf("wasm failed, error code %v", code)
	}
	return res, nil
}
