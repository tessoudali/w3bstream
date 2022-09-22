package instance

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/modules/instance"
)

type CreateInstance struct {
	httpx.MethodPost
	instance.CreateInstanceReq `in:"body"`
}

func (r *CreateInstance) Output(ctx context.Context) (interface{}, error) {
	return instance.CreateInstance(ctx, &r.CreateInstanceReq)
}
