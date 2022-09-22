package deploy

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/modules/deploy"
)

type GetInstance struct {
	httpx.MethodGet
	InstanceID uint32 `in:"path" name:"instanceID"`
}

func (r *GetInstance) Path() string { return "/:instanceID" }

func (r *GetInstance) Output(ctx context.Context) (interface{}, error) {
	_, err := validateByInstance(ctx, r.InstanceID)
	if err != nil {
		return nil, err
	}

	return deploy.GetInstanceByInstanceID(ctx, r.InstanceID)
}
