package strategy

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/apis/middleware"
	"github.com/iotexproject/w3bstream/pkg/modules/strategy"
	"github.com/iotexproject/w3bstream/pkg/types"
)

type CreateStrategy struct {
	httpx.MethodPost
	ProjectID                       types.SFID `in:"path" name:"projectID"`
	strategy.CreateStrategyBatchReq `in:"body"`
}

func (r *CreateStrategy) Path() string {
	return "/:projectID"
}

func (r *CreateStrategy) Output(ctx context.Context) (interface{}, error) {
	a := middleware.CurrentAccountFromContext(ctx)
	if _, err := a.ValidateProjectPerm(ctx, r.ProjectID); err != nil {
		return nil, err
	}

	return nil, strategy.CreateStrategy(ctx, r.ProjectID, &r.CreateStrategyBatchReq)
}
