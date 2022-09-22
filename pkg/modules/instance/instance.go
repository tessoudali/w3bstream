package instance

import (
	"context"

	"github.com/google/uuid"

	"github.com/iotexproject/w3bstream/pkg/enums"
	"github.com/iotexproject/w3bstream/pkg/modules/applet"
	"github.com/iotexproject/w3bstream/pkg/modules/vm"
	"github.com/iotexproject/w3bstream/pkg/types"

	"github.com/iotexproject/w3bstream/pkg/models"
)

type CreateInstanceReq struct {
	AppletID string `json:"appletID"`
}

func CreateInstance(ctx context.Context, r *CreateInstanceReq) (*models.Instance, error) {
	instanceID := uuid.New().String()
	d := types.MustDBExecutorFromContext(ctx)
	app, err := applet.GetAppletByID(ctx, r.AppletID)
	if err != nil {
		return nil, err
	}
	// TODO db tx
	vmID, err := vm.NewInstance(app.Path)
	if err != nil {
		return nil, err
	}
	// TODO when to start vm
	if err := vm.StartInstance(vmID); err != nil {
		return nil, err
	}

	m := &models.Instance{
		RelInstance:  models.RelInstance{InstanceID: instanceID, InstanceVMID: vmID},
		RelApplet:    models.RelApplet{AppletID: r.AppletID},
		InstanceInfo: models.InstanceInfo{Path: "", State: enums.INSTANCE_STATE__STARTED},
	}

	if err := m.Create(d); err != nil {
		return nil, err
	}
	return m, nil
}

func GetInstanceByAppletID(ctx context.Context, appletID string) ([]*models.Instance, error) {
	// TODO need help
	return nil, nil
}
