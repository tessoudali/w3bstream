package applet

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"
	"github.com/iotexproject/w3bstream/pkg/modules/applet"
)

type UpdateApplet struct {
	httpx.MethodPut
	AppletID string `in:"path" name:"appletID"`
	applet.CreateAppletReq
}

func (r *UpdateApplet) Path() string { return "/:appletID" }

func (r *UpdateApplet) Output(ctx context.Context) (interface{}, error) {
	// TODO
	return nil, nil
}

type UpdateAppletAndDeploy struct {
}
