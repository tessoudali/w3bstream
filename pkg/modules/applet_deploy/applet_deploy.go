package applet_deploy

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/google/uuid"
	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
	"github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"

	"github.com/iotexproject/w3bstream/cmd/demo/global"
	"github.com/iotexproject/w3bstream/pkg/models"
)

type CreateDeployReq struct {
	models.RefApplet
	models.DeployInfo
}

type CreateDeployRsp struct {
	models.RefApplet
	models.RelDeploy
	models.DeployInfo
	datatypes.OperationTimes
}

func CreateDeploy(ctx context.Context, r *CreateDeployReq) (*CreateDeployRsp, error) {
	deploy := models.AppletDeploy{
		RelDeploy:  models.RelDeploy{DeployID: uuid.New().String()},
		RefApplet:  r.RefApplet,
		DeployInfo: r.DeployInfo,
	}
	applet := models.Applet{RefApplet: models.RefApplet{AppletID: r.AppletID}}

	d := global.DBExecutorFromContext(ctx)
	l := global.LoggerFromContext(ctx)

	err := sqlx.NewTasks(d).With(
		func(db sqlx.DBExecutor) error {
			return applet.FetchByAppletID(db)
		},
		func(db sqlx.DBExecutor) error {
			return deploy.Create(db)
		},
	).Do()

	if err != nil {
		l.Error(err)
		return nil, err
	}
	return &CreateDeployRsp{
		RefApplet:      applet.RefApplet,
		RelDeploy:      deploy.RelDeploy,
		DeployInfo:     deploy.DeployInfo,
		OperationTimes: deploy.OperationTimes,
	}, nil
}

type CreateDeployByAssertReq struct {
	AppletID string `in:"path" name:"appletID"`
	Location string `in:"path" name:"location,omitempty"`
}

func CreateDeployByAssert(ctx context.Context, r *CreateDeployByAssertReq) (*CreateDeployRsp, error) {
	// TODO fetch asserts from loc(ipfs)

	l := global.LoggerFromContext(ctx)

	_, current, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(current), "testdata")
	c, err := LoadConfigFrom(filepath.Join(root, "applet.yaml"))
	if err != nil {
		l.Error(err)
		return nil, err
	}
	m := c.DataSources[0].Mapping

	return CreateDeploy(ctx, &CreateDeployReq{
		RefApplet: models.RefApplet{AppletID: r.AppletID},
		DeployInfo: models.DeployInfo{
			Location: r.Location,
			Version:  m.APIVersion,
			WasmFile: m.File,
			AbiName:  m.ABIs[0].Name,
			AbiFile:  m.ABIs[0].File,
		},
	})
}

type ListDeployReq struct {
	AppletIDs   []string `in:"query" name:"appletID,omitempty"`
	AppletNames []string `in:"query" name:"appletName,omitempty"`
	DeployIDs   []string `in:"query" name:"deployID,omitempty"`
	datatypes.Pager
}

func (r *ListDeployReq) Condition() builder.SqlCondition {
	var (
		cs []builder.SqlCondition
		ma = &models.Applet{}
		md = &models.AppletDeploy{}
	)
	if len(r.AppletIDs) > 0 {
		cs = append(cs, md.ColAppletID().In(r.AppletIDs))
	}
	if len(r.AppletNames) > 0 {
		cs = append(cs, ma.ColName().In(r.AppletNames))
	}
	if len(r.DeployIDs) > 0 {
		cs = append(cs, md.ColDeployID().In(r.DeployIDs))
	}
	return builder.And(cs...)
}

func (r *ListDeployReq) Additions() builder.Additions {
	var (
		ma = &models.Applet{}
		md = &models.AppletDeploy{}
	)
	return builder.Additions{
		builder.OrderBy(builder.DescOrder(ma.ColCreatedAt())),
		builder.OrderBy(builder.DescOrder(md.ColCreatedAt())),
		r.Pager.Addition(),
	}
}

type ListDeployRsp struct {
	Data []struct {
		AppletID       string `db:"f_applet_id"`
		AppletName     string `db:"f_applet_name"`
		DeployID       string `db:"f_deploy_id"`
		AssertLocation string `db:"f_assert_location"`
		WasmFile       string `db:"f_wasm_file"`
		DeployVersion  string `db:"f_deploy_version"`
		AbiName        string `db:"f_abi_name"`
		AbiFile        string `db:"f_abi_file"`
	} `json:"data"`
	Hints int64 `json:"hints"`
}

func ListDeploy(ctx context.Context, r *ListDeployReq) ([]models.AppletDeploy, error) {
	d := global.DBExecutorFromContext(ctx)
	l := global.LoggerFromContext(ctx)

	md := &models.AppletDeploy{}
	ma := &models.Applet{}

	builder.Select(
		builder.MultiWith(
			",",
			builder.Alias(ma.ColAppletID(), `f_applet_id`),
			builder.Alias(ma.ColName(), `f_applet_name`),
			builder.Alias(md.ColLocation(), `f_assert_location`),
			builder.Alias(md.ColDeployID(), `f_deploy_id`),
			builder.Alias(md.ColVersion(), `f_deploy_version`),
			builder.Alias(md.ColAbiName(), `f_abi_name`),
			builder.Alias(md.ColAbiFile(), `f_abi_file`),
		),
	).From(
		d.T(ma),
		append(
			builder.Additions{
				builder.LeftJoin(d.T(md)).On(ma.ColAppletID().Eq(md.ColAppletID())),
				builder.Where(r.Condition()),
			},
			r.Additions()...,
		)...,
	)

	ret, err := (&models.AppletDeploy{}).List(d, nil)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	return ret, nil
}

func RemoveDeployByAppletIDAndVersion(ctx context.Context, appletID, version string) error {
	d := global.DBExecutorFromContext(ctx)
	l := global.LoggerFromContext(ctx)
	m := &models.AppletDeploy{
		RefApplet:  models.RefApplet{AppletID: appletID},
		DeployInfo: models.DeployInfo{Version: version},
	}

	err := m.DeleteByAppletIDAndVersion(d)
	if err != nil {
		l.Error(err)
		return err
	}
	return nil
}

func RemoveDeployByDeployID(ctx context.Context, deployID string) error {
	d := global.DBExecutorFromContext(ctx)
	l := global.LoggerFromContext(ctx)
	m := &models.AppletDeploy{
		RelDeploy: models.RelDeploy{DeployID: deployID},
	}

	err := m.DeleteByDeployID(d)
	if err != nil {
		l.Error(err)
		return err
	}
	return nil
}
