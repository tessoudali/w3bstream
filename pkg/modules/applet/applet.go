package applet

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
	"github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"

	"github.com/iotexproject/w3bstream/pkg/modules/project"
	"github.com/iotexproject/w3bstream/pkg/modules/resource"
	v1 "github.com/iotexproject/w3bstream/pkg/modules/vm/v1"
	"github.com/iotexproject/w3bstream/pkg/types"

	"github.com/iotexproject/w3bstream/pkg/errors/status"
	"github.com/iotexproject/w3bstream/pkg/models"
)

type CreateAndDeployReq struct {
	File *multipart.FileHeader `name:"file"`
	Info `name:"info"`
}

type Info struct {
	ProjectID  string `json:"projectID"`
	AppletName string `json:"appletName"`
}

func CreateAndDeployApplet(ctx context.Context, r *CreateAndDeployReq) (*models.Applet, error) {
	// handle upload
	appletID := uuid.New().String()
	_, filename, err := resource.Upload(ctx, r.File, appletID)
	if err != nil {
		return nil, err
	}

	d := types.MustDBExecutorFromContext(ctx)

	prj, err := project.GetAndValidateProjectPerm(ctx, r.ProjectID)
	if err != nil {
		return nil, err
	}

	m := &models.Applet{
		RelProject: models.RelProject{ProjectID: r.ProjectID},
		RelApplet:  models.RelApplet{AppletID: appletID},
		AppletInfo: models.AppletInfo{Name: r.AppletName, AssetLoc: filename},
	}

	if err := m.Create(d); err != nil {
		defer os.RemoveAll(filename)
		return nil, err
	}

	instanceID, err := v1.NewInstance(m.AssetLoc,
		v1.InstanceOptionWithChannel(prj.Name+"@"+m.Name),
		v1.InstanceOptionWithLogger(types.MustLoggerFromContext(ctx)),
		v1.InstanceOptionWithMqttBroker(types.MustMqttBrokerFromContext(ctx)),
	)
	if err != nil {
		defer os.RemoveAll(filename)
		return nil, err
	}
	err = v1.RunInstance(instanceID)
	if err != nil {
		defer os.RemoveAll(filename)
		return nil, err
	}

	return m, nil
}

// CreateApplet (prjID ,apple info)
// QueryInsByAppletID()...
// CreateIns... (apple + version)
// StopIns (insID)
// QueryInsByInsId(insID)

type CreateAppletByNameReq struct {
	Name string `json:"name"`
}

func CreateAppletByName(ctx context.Context, req *CreateAppletByNameReq) (*models.Applet, error) {
	applet := &models.Applet{
		RelApplet:  models.RelApplet{AppletID: uuid.New().String()},
		AppletInfo: models.AppletInfo{Name: req.Name},
	}

	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)

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

	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)

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
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)

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
