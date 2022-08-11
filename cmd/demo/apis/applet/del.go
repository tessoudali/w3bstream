package applet

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/modules/applet"
)

type RemoveAppletByAppletID struct {
	httpx.MethodDelete
	AppletID string `in:"path" name:"appletID"`
}

func (r *RemoveAppletByAppletID) Path() string {
	return "/:appletID"
}

func (r *RemoveAppletByAppletID) Output(ctx context.Context) (interface{}, error) {
	return nil, applet.RemoveApplet(ctx, r.AppletID)
}
