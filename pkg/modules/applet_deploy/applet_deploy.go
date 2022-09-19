package applet_deploy

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
	"github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"

	"github.com/iotexproject/w3bstream/pkg/types"

	"github.com/iotexproject/w3bstream/pkg/errors/status"
	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/modules/resource"
	"github.com/iotexproject/w3bstream/pkg/modules/vm"
)

type CreateDeployByAssertReq struct {
	File *multipart.FileHeader `name:"file"`
	Info AssertInfo            `name:"info"`
}

type AssertInfo struct {
	AppletID  string `json:"appletID"`
	AssertMd5 string `json:"assertMd5,omitempty"`
}

type CreateDeployContext struct {
	models.RelApplet
	models.DeployInfo
	Handlers []models.HandlerInfo
}

type CreateDeployReq struct {
	AppletID string `in:"path" name:"appletID"`
	Location string `in:"path" name:"location,omitempty"`
}

type CreateDeployRsp struct {
	models.RelApplet
	models.RelDeploy
	models.DeployInfo
	Handlers []models.HandlerInfo
	datatypes.OperationTimes
}

func CreateDeployByContext(ctx context.Context, c *vm.VM, r *CreateDeployContext) (*CreateDeployRsp, error) {
	deploy := models.AppletDeploy{
		RelDeploy:  models.RelDeploy{DeployID: uuid.New().String()},
		RelApplet:  r.RelApplet,
		DeployInfo: r.DeployInfo,
	}
	applet := models.Applet{RelApplet: models.RelApplet{AppletID: r.AppletID}}

	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)

	err := sqlx.NewTasks(d).With(
		func(db sqlx.DBExecutor) error {
			err := applet.FetchByAppletID(db)
			if err != nil {
				if sqlx.DBErr(err).IsNotFound() {
					return status.NotFound.StatusErr().WithMsg(
						fmt.Sprintf("applet %s not found", r.AppletID),
					)
				}
				return err
			}
			return nil
		},
		func(db sqlx.DBExecutor) error {
			err := deploy.Create(db)
			if err != nil {
				if sqlx.DBErr(err).IsConflict() {
					return status.Conflict.StatusErr().WithMsg(
						fmt.Sprintf(
							"applet: %v version: %v are already exist",
							deploy.AppletID, deploy.Version,
						),
					)
				}
				return err
			}
			return nil
		},
		func(db sqlx.DBExecutor) error {
			for _, h := range r.Handlers {
				err := (&models.Handler{
					RelApplet:   models.RelApplet{AppletID: r.AppletID},
					RelDeploy:   models.RelDeploy{DeployID: deploy.DeployID},
					RelHandler:  models.RelHandler{HandlerID: uuid.New().String()},
					HandlerInfo: models.HandlerInfo{Name: h.Name, Params: h.Params},
				}).Create(db)
				if err != nil {
					if sqlx.DBErr(err).IsConflict() {
						return status.Conflict.StatusErr().WithMsg(fmt.Sprintf(
							"applet: %v version: %v handler: %v are already exist",
							deploy.AppletID, deploy.Version, h.Name,
						))
					}
					return err
				}
			}
			return nil
		},
	).Do()

	if err != nil {
		l.Error(err)
		return nil, err
	}

	vm.Start(
		ctx,
		vm.NewMonitorContext(
			c, applet.AppletID, applet.Name, deploy.Version, r.Handlers...,
		),
	)

	return &CreateDeployRsp{
		RelApplet:      applet.RelApplet,
		RelDeploy:      deploy.RelDeploy,
		DeployInfo:     deploy.DeployInfo,
		OperationTimes: deploy.OperationTimes,
	}, nil
}

func CreateDeployByAssert(ctx context.Context, r *CreateDeployByAssertReq) (*CreateDeployRsp, error) {
	root, filename, err := resource.Upload(ctx, r.File, r.Info.AppletID)
	if err != nil {
		return nil, status.UploadFileFailed.StatusErr().WithDesc(err.Error())
	}
	if r.Info.AssertMd5 != "" {
		if err = resource.CheckMD5(filename, r.Info.AssertMd5); err != nil {
			return nil, status.MD5ChecksumFailed.StatusErr().WithDesc(err.Error())
		}
	}
	dst := filepath.Join(root, uuid.New().String())
	defer os.RemoveAll(filename)
	if err = resource.UnTar(dst, filename); err != nil {
		return nil, status.ExtractFileFailed.StatusErr().WithDesc(err.Error())
	}

	l := types.MustLoggerFromContext(ctx)

	c, err := vm.Load(dst)
	if err != nil {
		l.Error(err)
		return nil, status.LoadVMFailed.StatusErr().WithDesc(err.Error())
	}
	m := c.DataSources[0].Mapping

	hdls, err := LoadHandlers(
		filepath.Join(dst, m.ABIs[0].File),
		m.EventHandlers...,
	)
	if err != nil {
		l.Error(err)
		return nil, status.LoadVMFailed.StatusErr().WithDesc(err.Error())
	}

	rsp, err := CreateDeployByContext(ctx, c, &CreateDeployContext{
		RelApplet: models.RelApplet{AppletID: r.Info.AppletID},
		DeployInfo: models.DeployInfo{
			Location: filepath.Join(root, m.APIVersion),
			Version:  m.APIVersion,
			WasmFile: m.File,
			AbiName:  m.ABIs[0].Name,
			AbiFile:  m.ABIs[0].File,
		},
		Handlers: hdls,
	})
	if err != nil {
		_ = os.RemoveAll(dst)
	} else {
		installed := filepath.Join(root, m.APIVersion)
		_ = os.RemoveAll(installed)
		if err := os.Rename(dst, installed); err != nil {
			return nil, err
		}
	}
	return rsp, err
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
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)

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
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)
	m := &models.AppletDeploy{
		RelApplet:  models.RelApplet{AppletID: appletID},
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
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)
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
