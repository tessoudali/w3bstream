package deploy

import (
	"context"

	"github.com/google/uuid"
	"github.com/iotexproject/w3bstream/pkg/enums"
	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/modules/vm"
)

type CreateInstanceReq struct {
	ProjectID string `in:"path" name:"projectID"`
	AppletID  string `in:"path" name:"appletID"`
}

type CreateInstanceRsp struct {
	ProjectID     string              `json:"projectID"`
	AppletID      string              `json:"appletID"`
	InstanceID    string              `json:"instanceID,omitempty"`
	InstanceState enums.InstanceState `json:"instanceState"`
}

func CreateInstance(ctx context.Context, r *CreateInstanceReq) (*CreateInstanceRsp, error) {
	// TODO
	return nil, nil
}

type ControlReq struct {
	ProjectID  string          `in:"path"  name:"projectID"`
	InstanceID string          `in:"path"  name:"instanceID"`
	Cmd        enums.DeployCmd `in:"query" name:"cmd"`
}

func ControlInstance(ctx context.Context, instanceID string, cmd enums.DeployCmd) error {
	// TODO
	id, err := uuid.Parse(instanceID)

	switch cmd {
	case enums.DEPLOY_CMD__REMOVE:
		// TODO stop instance and remove rel from database
		err = vm.DelInstance(id.ID())
	case enums.DEPLOY_CMD__STOP:
		err = vm.StopInstance(id.ID())
		// TODO
	case enums.DEPLOY_CMD__START:
		err = vm.StartInstance(id.ID())
		return err
	}
	return nil
}

type GetInstanceRsp struct {
	models.Instance
	State enums.InstanceState `json:"state"`
}

func GetInstance(ctx context.Context, instanceID string) (*GetInstanceRsp, error) {
	// TODO
	_, _ = vm.GetInstanceState(0)
	return nil, nil
}

type ListInstanceReq struct {
	ProjectID string                `in:"path"  name:"projectID"`
	AppletIDs string                `in:"query" name:"appletIDs"`
	Status    []enums.InstanceState `in:"query" name:"states,omitempty"`
}
