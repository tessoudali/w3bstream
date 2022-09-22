package deploy

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"
	"github.com/iotexproject/w3bstream/pkg/modules/deploy"
)

type CreateInstance struct {
	httpx.MethodPost
	deploy.CreateInstanceReq
}

func (r *CreateInstance) Path() string {
	return "/:projectID/:appletID"
}

func (r *CreateInstance) Output(ctx context.Context) (interface{}, error) {
	// TODO project permission
	return deploy.CreateInstance(ctx, &r.CreateInstanceReq)
}
