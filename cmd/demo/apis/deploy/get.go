package deploy

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/modules/applet_deploy"
)

type ListDeploy struct {
	httpx.MethodGet
	applet_deploy.ListDeployReq
}

func (r *ListDeploy) Output(ctx context.Context) (interface{}, error) {
	return applet_deploy.ListDeploy(ctx, &r.ListDeployReq)
}
