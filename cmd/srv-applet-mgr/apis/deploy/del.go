package deploy

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/modules/applet_deploy"
)

type RemoveDeployByAppletIDAndVersion struct {
	httpx.MethodDelete
	AppletID string `in:"path" name:"appletID"`
	Version  string `in:"path" name:"version"`
}

func (r *RemoveDeployByAppletIDAndVersion) Path() string {
	return "/applet/:appletID/version/:version"
}

func (r *RemoveDeployByAppletIDAndVersion) Output(ctx context.Context) (interface{}, error) {
	return nil, applet_deploy.RemoveDeployByAppletIDAndVersion(
		ctx, r.AppletID, r.Version,
	)
}

type RemoveDeployByDeployID struct {
	httpx.MethodDelete
	DeployID string `in:"path" name:"deployID"`
}

func (r *RemoveDeployByDeployID) Path() string {
	return "/deploy/:deployID"
}

func (r *RemoveDeployByDeployID) Output(ctx context.Context) (interface{}, error) {
	return nil, applet_deploy.RemoveDeployByDeployID(ctx, r.DeployID)
}
