// project management

package project

import (
	"context"

	"github.com/google/uuid"

	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/apis/middleware"
	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/types"
)

type CreateProjectReq = models.ProjectInfo

func CreateProject(ctx context.Context, r *CreateProjectReq) (*models.Project, error) {
	d := types.MustDBExecutorFromContext(ctx)
	a := middleware.CurrentAccountFromContext(ctx)

	m := &models.Project{
		RelProject:  models.RelProject{ProjectID: uuid.New().String()},
		RelAccount:  models.RelAccount{AccountID: a.AccountID},
		ProjectInfo: *r,
	}

	if err := m.Create(d); err != nil {
		return nil, err
	}

	return m, nil
}

func DeleteProject(ctx context.Context, prjID string) error {
	// TODO
	return nil
}

type ListProjectReq struct {
	ProjectIDs []string `in:"query" name:"projectIDs"`
	Names      []string `in:"query" name:"names"`
}

type ListProjectRsp struct {
	Data  []models.Project `json:"data"` // Data project data list
	Total int64            `json:""`     // Total project count under current user
}

func ListProject(ctx context.Context, r *ListProjectReq) (*ListProjectRsp, error) {
	// TODO
	return nil, nil
}

type GetProjectRsp struct {
	ProjectID string
	Applets   []models.Applet
}

func GetProjectByProjectID(ctx context.Context, prjID string) (*GetProjectRsp, error) {
	// TODO
	return nil, nil
}
