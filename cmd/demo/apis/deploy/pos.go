package deploy

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/modules/applet_deploy"
)

type CreateDeploy struct {
	httpx.MethodPost
	applet_deploy.CreateDeployReq `in:"body"`
}

func (r *CreateDeploy) Output(ctx context.Context) (interface{}, error) {
	return applet_deploy.CreateDeploy(ctx, &r.CreateDeployReq)
}

type CreateDeployByAssert struct {
	httpx.MethodPost
	applet_deploy.CreateDeployByAssertReq
}

func (r *CreateDeployByAssert) Path() string {
	return "/:appletID/:location"
}

func (r *CreateDeployByAssert) Output(ctx context.Context) (interface{}, error) {
	return applet_deploy.CreateDeployByAssert(ctx, &r.CreateDeployByAssertReq)
}
