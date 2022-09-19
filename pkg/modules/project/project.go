// project management

package project

import (
	"context"

	"github.com/google/uuid"
	"github.com/iotexproject/w3bstream/pkg/errors/status"

	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/apis/middleware"
	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/types"
)

type CreateProjectReq struct {
	models.ProjectInfo
}

func CreateProject(ctx context.Context, r *CreateProjectReq) (*models.Project, error) {
	d := types.MustDBExecutorFromContext(ctx)
	ca := middleware.CurrentAccountFromContext(ctx)

	m := &models.Project{
		RelProject:  models.RelProject{ProjectID: uuid.New().String()},
		RelAccount:  models.RelAccount{AccountID: ca.AccountID},
		ProjectInfo: r.ProjectInfo,
	}

	if err := m.Create(d); err != nil {
		return nil, err
	}

	return m, nil
}

func GetAndValidateProjectPerm(ctx context.Context, prjID string) (*models.Project, error) {
	d := types.MustDBExecutorFromContext(ctx)
	ca := middleware.CurrentAccountFromContext(ctx)
	m := &models.Project{RelProject: models.RelProject{ProjectID: prjID}}

	if err := m.FetchByProjectID(d); err != nil {
		return nil, err
	}
	if ca.AccountID != m.AccountID {
		return nil, status.Unauthorized.StatusErr().WithDesc("project permission deny")
	}
	return m, nil
}
