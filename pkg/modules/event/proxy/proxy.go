// event proxy

package proxy

import (
	"context"

	"github.com/iotexproject/w3bstream/pkg/types"
)

type Event interface {
	Meta() MetaData
	Raw() []byte
}

type MetaData struct {
	PublisherID types.SFID
	ProjectID   types.SFID
	AppletID    types.SFID
	Handler     string // optional
}

type Result struct {
	Success bool
	Data    []byte
}

type EventResult interface {
	ResultChan() chan<- Result
}

func Proxy(ctx context.Context, e Event) {
	logger := types.MustLoggerFromContext(ctx)
	func() {
		success := false
		var data []byte
		result, ok := e.(EventResult)
		defer func() {
			if ok {
				result.ResultChan() <- Result{Success: success, Data: data}
			}
		}()

		// adapt
		res, err := dispatch(ctx, e)
		if err != nil {
			logger.Error(err)
			return
		}
		success = true
		data = res
	}()
}
