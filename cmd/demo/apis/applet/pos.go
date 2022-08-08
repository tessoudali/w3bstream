package applet

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"
	"github.com/iotexproject/w3bstream/pkg/modules/applet"
)

type CreateApplet struct {
	httpx.MethodPost             `summary:"create applet by name"`
	applet.CreateAppletByNameReq `in:"body"`
}

func (r *CreateApplet) Output(ctx context.Context) (interface{}, error) {
	return applet.CreateAppletByName(ctx, &r.CreateAppletByNameReq)
}
