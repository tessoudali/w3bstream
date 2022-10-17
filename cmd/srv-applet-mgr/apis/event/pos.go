package event

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/modules/event"

	"github.com/iotexproject/w3bstream/pkg/depends/protocol/eventpb"
)

type HandleEvent struct {
	httpx.MethodPost
	ProjectName   string `in:"path" name:"projectName"`
	eventpb.Event `in:"body"`
}

func (r *HandleEvent) Path() string { return "/:projectName" }

func (r *HandleEvent) Output(ctx context.Context) (interface{}, error) {
	return event.OnEventReceived(ctx, r.ProjectName, &r.Event)
}
