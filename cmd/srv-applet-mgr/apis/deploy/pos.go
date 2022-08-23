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
	return "/applet/:appletID/location/:location"
}

func (r *CreateDeploy) Output(ctx context.Context) (interface{}, error) {
	return nil, nil
	// return applet_deploy.CreateDeploy(ctx, &r.CreateDeployReq)
}

type CreateDeployByAssert struct {
	httpx.MethodPost
	applet_deploy.CreateDeployByAssertReq `in:"body" mime:"multipart"`
}

func (r *CreateDeployByAssert) Path() string { return "/assert" }

func (r *CreateDeployByAssert) Output(ctx context.Context) (interface{}, error) {
	return applet_deploy.CreateDeployByAssert(ctx, &r.CreateDeployByAssertReq)
}
