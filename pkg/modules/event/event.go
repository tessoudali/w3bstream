package event

import (
	"context"

	"github.com/iotexproject/w3bstream/pkg/depends/protocol/eventpb"
	"github.com/iotexproject/w3bstream/pkg/enums"
	"github.com/iotexproject/w3bstream/pkg/errors/status"
	"github.com/iotexproject/w3bstream/pkg/modules/project"
	"github.com/iotexproject/w3bstream/pkg/modules/publisher"
	"github.com/iotexproject/w3bstream/pkg/modules/strategy"
	"github.com/iotexproject/w3bstream/pkg/modules/vm"
	"github.com/iotexproject/w3bstream/pkg/types"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
)

type HandleEventResult struct {
	InstanceID types.SFID            `json:"instanceID"`
	ResultCode wasm.ResultStatusCode `json:"resultCode"`
}

type HandleEventRsp []HandleEventResult

func OnEventReceived(ctx context.Context, projectName string, r *eventpb.Event) (HandleEventRsp, error) {
	if r.Header != nil && len(r.Header.PubId) > 0 {
		puber, err := publisher.GetPublisherByPublisherKey(ctx, r.Header.PubId)
		if err != nil {
			return nil, err
		}

		prj, err := project.GetProjectByProjectName(ctx, projectName)
		if err != nil {
			return nil, err
		}

		if puber.ProjectID != prj.ProjectID {
			return nil, status.Forbidden.StatusErr().WithDesc("no project permission")
		}
	}

	eventType := enums.EVENT_TYPE__ANY
	if r.Header != nil {
		eventType = enums.EventType(r.Header.EventType)
	}

	instances, err := strategy.FindStrategyInstances(ctx, projectName, eventType)
	if err != nil {
		return nil, err
	}

	ret := make(HandleEventRsp, 0, len(instances))

	for _, v := range instances {
		consumer := vm.GetConsumer(v.InstanceID.String())
		if consumer == nil {
			continue
		}
		// TODO
		_, code := consumer.HandleEvent(v.Handler, []byte(r.Payload))
		ret = append(ret, HandleEventResult{
			InstanceID: v.InstanceID,
			ResultCode: code,
		})
	}
	return ret, nil
}
