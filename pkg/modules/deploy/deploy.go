package deploy

import (
	"context"
	"fmt"

	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/pkg/enums"
	"github.com/iotexproject/w3bstream/pkg/errors/status"
	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/modules/vm"
	"github.com/iotexproject/w3bstream/pkg/types"
)

type CreateInstanceRsp struct {
	InstanceID    string              `json:"instanceID"`
	InstanceState enums.InstanceState `json:"instanceState"`
}

func CreateInstance(ctx context.Context, path, appletID string) (*CreateInstanceRsp, error) {
	d := types.MustDBExecutorFromContext(ctx)
	m := &models.Instance{
		RelApplet: models.RelApplet{AppletID: appletID},
	}

	// TODO
	count, err := m.Count(d, m.ColAppletID().Eq(appletID))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, status.InstanceLimit
	}

	m.InstanceID, err = vm.NewInstance(path, vm.DefaultInstanceOptionSetter)
	if err != nil {
		return nil, err
	}
	m.State = enums.INSTANCE_STATE__CREATED
	m.Path = path

	if err = m.Create(d); err != nil {
		_ = vm.DelInstance(m.InstanceID)
		return nil, err
	}

	return &CreateInstanceRsp{
		InstanceID:    m.InstanceID,
		InstanceState: m.State,
	}, nil
}

func ControlInstance(ctx context.Context, instanceID string, cmd enums.DeployCmd) (err error) {
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)
	m := &models.Instance{RelInstance: models.RelInstance{InstanceID: instanceID}}

	l.Start(ctx, "ControlInstance")
	defer l.End()

	defer func() {
		if err != nil {
			l.WithValues("instance", instanceID, "cmd", cmd).Error(err)
		}
	}()

	if err = m.FetchByInstanceID(d); err != nil {
		return err
	}

	switch cmd {
	case enums.DEPLOY_CMD__REMOVE:
		if err = vm.DelInstance(instanceID); err != nil {
			return err
		}
		if err = m.DeleteByInstanceID(d); err != nil {
			return err
		}
	case enums.DEPLOY_CMD__STOP:
		if err = vm.StopInstance(instanceID); err != nil {
			return err
		}
		m.State = enums.INSTANCE_STATE__STOPPED
		if err = m.UpdateByInstanceID(d); err != nil {
			return err
		}
	case enums.DEPLOY_CMD__START:
		if err = vm.StartInstance(instanceID); err != nil {
			return err
		}
		m.State = enums.INSTANCE_STATE__STARTED
		if err = m.UpdateByInstanceID(d); err != nil {
			return err
		}
	case enums.DEPLOY_CMD__REDEPLOY:
		if err = vm.StopInstance(instanceID); err != nil && err != vm.ErrNotFound {
			return err
		}
		if err = vm.DelInstance(instanceID); err != nil {
			return err
		}
		err = vm.NewInstanceWithID(m.Path, instanceID, vm.DefaultInstanceOptionSetter)
		if err != nil {
			return err
		}
		m.State = enums.INSTANCE_STATE__CREATED
		if err = m.UpdateByInstanceID(d); err != nil {
			return err
		}
	}
	return nil
}

func GetInstanceByInstanceID(ctx context.Context, instanceID string) (*models.Instance, error) {
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)
	m := &models.Instance{RelInstance: models.RelInstance{InstanceID: instanceID}}

	l.Start(ctx, "GetInstanceByInstanceID")

	if err := m.FetchByInstanceID(d); err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, status.NotFound.StatusErr().WithDesc("instance not found in db")
		}
		return nil, err
	}

	state, ok := vm.GetInstanceState(instanceID)
	if !ok {
		return nil, status.NotFound.StatusErr().WithDesc("instance not found in mgr")
	}
	if state != m.State {
		l.WithValues("mgr_state", state, "db_state", m.State).
			Warn(errors.New("unmatched"))
		m.State = state
		if err := m.UpdateByInstanceID(d); err != nil {
			return nil, err
		}
	}

	return m, nil
}

func GetInstanceByAppletID(ctx context.Context, appletID string) (ret []models.Instance, err error) {
	d := types.MustDBExecutorFromContext(ctx)
	m := &models.Instance{}

	err = d.QueryAndScan(
		builder.Select(nil).From(
			d.T(m),
			builder.Where(m.ColAppletID().Eq(appletID)),
		),
		&ret,
	)
	return
}

func StartInstances(ctx context.Context) error {
	fmt.Println("---------------")

	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)
	m := &models.Instance{}

	list, err := m.List(d, nil)
	if err != nil {
		return err
	}
	for _, i := range list {
		if i.State == enums.INSTANCE_STATE__CREATED || i.State == enums.INSTANCE_STATE__STARTED {
			err = vm.NewInstanceWithID(i.Path, i.InstanceID, vm.DefaultInstanceOptionSetter)
			if err != nil {
				if err := i.DeleteByInstanceID(d); err != nil {
					return err
				}
				if err := (&models.Applet{RelApplet: models.RelApplet{AppletID: i.AppletID}}).
					DeleteByAppletID(d); err != nil {
					return err
				}
				l.WithValues("instance", i.InstanceID, "applet", i.AppletID).
					Warn(errors.New("start failed and removed"))
				return nil
			}
			m.State = enums.INSTANCE_STATE__CREATED
		}
		if i.State == enums.INSTANCE_STATE__STARTED {
			_ = ControlInstance(ctx, i.InstanceID, enums.DEPLOY_CMD__START)
		}
	}
	return nil
}
