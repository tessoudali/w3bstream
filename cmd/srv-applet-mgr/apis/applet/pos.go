package applet

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/modules/applet"
)

type CreateAndDeployApplet struct {
	httpx.MethodPost
	applet.CreateAndDeployReq `in:"body" mime:"multipart"`
}

func (r *CreateAndDeployApplet) Output(ctx context.Context) (interface{}, error) {
	return applet.CreateAndDeployApplet(ctx, &r.CreateAndDeployReq)
}
