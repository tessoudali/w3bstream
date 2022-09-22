package deploy

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"
	"github.com/iotexproject/w3bstream/pkg/modules/deploy"
)

type ControlInstance struct {
	httpx.MethodPut
	deploy.ControlReq
}

func (r *ControlInstance) Path() string { return "/:projectID/:instanceID" }

func (r *ControlInstance) Output(ctx context.Context) (interface{}, error) {
	return nil, deploy.ControlInstance(ctx, r.InstanceID, r.Cmd)
}
