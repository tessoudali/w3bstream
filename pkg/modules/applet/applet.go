package applet

import (
	"context"

	"github.com/google/uuid"
	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
	"github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"

	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/global"
	"github.com/iotexproject/w3bstream/pkg/errors/status"
	"github.com/iotexproject/w3bstream/pkg/models"
)

type CreateAppletByNameReq struct {
	Name string `json:"name"`
}

func CreateAppletByName(ctx context.Context, req *CreateAppletByNameReq) (*models.Applet, error) {
	applet := &models.Applet{
		RelApplet:  models.RelApplet{AppletID: uuid.New().String()},
		AppletInfo: models.AppletInfo{Name: req.Name},
	}

	d := global.DBExecutorFromContext(ctx)
	l := global.LoggerFromContext(ctx)

	l.Start(ctx, "CreateAppletByName")
	defer l.End()

	err := sqlx.NewTasks(d).With(
		func(db sqlx.DBExecutor) error {
			return applet.Create(db)
		},
		func(db sqlx.DBExecutor) error {
			return applet.FetchByAppletID(db)
		},
	).Do()
	if err != nil {
		l.Error(err)
		if sqlx.DBErr(err).IsConflict() {
			return nil, status.Conflict.StatusErr().WithMsg("create applet conflict")
		}
		return nil, err
	}

	return applet, nil
}

type ListAppletReq struct {
	IDs       []string `in:"query" name:"id,omitempty"`
	AppletIDs []string `in:"query" name:"appletID,omitempty"`
	Names     []string `in:"query" name:"name,omitempty"`
	datatypes.Pager
}

func (r *ListAppletReq) Condition() builder.SqlCondition {
	var (
		m  = &models.Applet{}
		cs []builder.SqlCondition
	)
	if len(r.IDs) > 0 {
		cs = append(cs, m.ColID().In(r.IDs))
	}
	if len(r.AppletIDs) > 0 {
		cs = append(cs, m.ColAppletID().In(r.AppletIDs))
	}
	if len(r.Names) > 0 {
		cs = append(cs, m.ColName().In(r.Names))
	}
	return builder.And(cs...)
}

func (r *ListAppletReq) Additions() builder.Additions {
	m := &models.Applet{}
	return builder.Additions{
		builder.OrderBy(builder.DescOrder(m.ColCreatedAt())),
		r.Pager.Addition(),
	}
}

type ListAppletRsp struct {
	Data  []models.Applet `json:"data"`
	Hints int64           `json:"hints"`
}

func ListApplets(ctx context.Context, r *ListAppletReq) (*ListAppletRsp, error) {
	applet := &models.Applet{}

	d := global.DBExecutorFromContext(ctx)
	l := global.LoggerFromContext(ctx)

	l.Start(ctx, "ListApplets")
	defer l.End()

	applets, err := applet.List(d, r.Condition(), r.Additions()...)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	hints, err := applet.Count(d, r.Condition())
	if err != nil {
		l.Error(err)
		return nil, err
	}
	return &ListAppletRsp{applets, hints}, nil
}

func RemoveApplet(ctx context.Context, appletID string) error {
	d := global.DBExecutorFromContext(ctx)
	l := global.LoggerFromContext(ctx)

	err := sqlx.NewTasks(d).With(
		func(d sqlx.DBExecutor) error {
			return (&models.Applet{
				RelApplet: models.RelApplet{AppletID: appletID},
			}).DeleteByAppletID(d)
		},
		func(d sqlx.DBExecutor) error {
			m := &models.AppletDeploy{}
			_, err := d.Exec(
				builder.Delete().From(
					models.AppletDeployTable,
					builder.Where(m.ColAppletID().Eq(appletID)),
				),
			)
			return err
		},
	).Do()
	if err != nil {
		l.Error(err)
	}
	return err
}
