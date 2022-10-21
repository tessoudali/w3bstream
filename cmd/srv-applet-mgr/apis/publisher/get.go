package publisher

import (
	"context"

	"github.com/iotexproject/Bumblebee/base/types"
	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"
	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/apis/middleware"
	"github.com/iotexproject/w3bstream/pkg/modules/publisher"
)

type ListPublisher struct {
	httpx.MethodGet
	ProjectID types.SFID `name:"projectID" in:"path"`
	publisher.ListPublisherReq
}

func (r *ListPublisher) Path() string { return "/:projectID" }

func (r *ListPublisher) Output(ctx context.Context) (interface{}, error) {
	ca := middleware.CurrentAccountFromContext(ctx)

	if _, err := ca.ValidateProjectPerm(ctx, r.ProjectID); err != nil {
		return nil, err
	}

	r.SetCurrentProject(r.ProjectID)
	return publisher.ListPublisher(ctx, &r.ListPublisherReq)
}
