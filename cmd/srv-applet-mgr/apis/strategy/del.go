package strategy

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/apis/middleware"
	"github.com/iotexproject/w3bstream/pkg/modules/strategy"
)

type RemoveStrategy struct {
	httpx.MethodDelete
	strategy.RemoveStrategyReq
}

func (r *RemoveStrategy) Path() string { return "/:projectID" }

func (r *RemoveStrategy) Output(ctx context.Context) (interface{}, error) {
	a := middleware.CurrentAccountFromContext(ctx)
	if _, err := a.ValidateProjectPerm(ctx, r.ProjectID); err != nil {
		return nil, err
	}

	return nil, strategy.RemoveStrategy(ctx, &r.RemoveStrategyReq)
}
