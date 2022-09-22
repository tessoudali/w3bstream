package applet

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/apis/middleware"

	"github.com/iotexproject/w3bstream/pkg/modules/applet"
)

type CreateApplet struct {
	httpx.MethodPost
	applet.CreateAppletReq `in:"body" mime:"multipart"`
}

func (r *CreateApplet) Output(ctx context.Context) (interface{}, error) {
	ca := middleware.CurrentAccountFromContext(ctx)
	if _, err := ca.ValidateProjectPerm(ctx, r.ProjectID); err != nil {
		return nil, err
	}

	return applet.CreateApplet(ctx, &r.CreateAppletReq)
}
