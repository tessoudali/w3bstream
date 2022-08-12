package deploy

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/modules/applet_deploy"
)

type CreateDeploy struct {
	httpx.MethodPost
	applet_deploy.CreateDeployReq
}

func (r *CreateDeploy) Path() string {
	return "/:appletID/:location"
}

func (r *CreateDeploy) Output(ctx context.Context) (interface{}, error) {
	return applet_deploy.CreateDeploy(ctx, &r.CreateDeployReq)
}
