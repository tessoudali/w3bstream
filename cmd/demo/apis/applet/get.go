package applet

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"
	"github.com/iotexproject/w3bstream/pkg/modules/applet"
)

type ListApplet struct {
	httpx.MethodGet `summary:"get applet info"`
	applet.ListAppletReq
}

func (r *ListApplet) Output(ctx context.Context) (interface{}, error) {
	return applet.ListApplets(ctx, &r.ListAppletReq)
}
