// event proxy

package proxy

import (
	"context"

	"github.com/iotexproject/w3bstream/pkg/modules/event"
	"github.com/iotexproject/w3bstream/pkg/types"
)

func Proxy(ctx context.Context, e event.Event) {
	logger := types.MustLoggerFromContext(ctx)
	func() {
		success := false
		var data []byte
		result, ok := e.(event.EventResult)
		defer func() {
			if ok {
				result.ResultChan() <- event.Result{Success: success, Data: data}
			}
		}()

		res, err := dispatch(e)
		if err != nil {
			logger.Error(err)
			return
		}
		success = true
		data = res
	}()
}
