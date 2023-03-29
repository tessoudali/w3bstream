package event

import (
	"context"
	"errors"
	"sync"

	"github.com/machinefi/w3bstream/pkg/depends/conf/jwt"
	"github.com/machinefi/w3bstream/pkg/depends/conf/log"
	"github.com/machinefi/w3bstream/pkg/depends/protocol/eventpb"
	"github.com/machinefi/w3bstream/pkg/enums"
	"github.com/machinefi/w3bstream/pkg/errors/status"
	"github.com/machinefi/w3bstream/pkg/models"
	"github.com/machinefi/w3bstream/pkg/modules/strategy"
	"github.com/machinefi/w3bstream/pkg/modules/vm"
	"github.com/machinefi/w3bstream/pkg/types"
	"github.com/machinefi/w3bstream/pkg/types/wasm"
)

type HandleEventResult struct {
	ProjectName string                   `json:"projectName"`
	PubName     string                   `json:"pubName,omitempty"`
	EventID     string                   `json:"eventID"`
	ErrMsg      string                   `json:"errMsg,omitempty"`
	WasmResults []wasm.EventHandleResult `json:"wasmResults"`
}

type HandleEventReq struct {
	Events []eventpb.Event `json:"events"`
}

func OnEventReceived(ctx context.Context, projectName string, r *eventpb.Event) (ret *HandleEventResult, err error) {
	l := types.MustLoggerFromContext(ctx)

	_, l = l.Start(ctx, "OnEventReceived")
	defer l.End()

	l = l.WithValues("project_name", projectName)

	ret = &HandleEventResult{
		ProjectName: projectName,
		EventID:     r.Header.EventId,
	}

	defer func() {
		if err != nil {
			ret.ErrMsg = err.Error()
		}
	}()

	if err = publisherVerification(ctx, r, l); err != nil {
		l.Error(err)
		return
	}

	eventType := enums.EVENTTYPEDEFAULT
	if r.Header != nil && len(r.Header.EventType) > 0 {
		eventType = r.Header.EventType
	}
	l = l.WithValues("event_type", eventType)
	var handlers []*strategy.InstanceHandler
	handlers, err = strategy.FindStrategyInstances(ctx, projectName, eventType)
	if err != nil {
		l.Error(err)
		return
	}

	l.Info("matched strategies: %d", len(handlers))

	res := make(chan *wasm.EventHandleResult, len(handlers))

	wg := &sync.WaitGroup{}
	for _, v := range handlers {
		i := vm.GetConsumer(v.InstanceID)
		if i == nil {
			res <- &wasm.EventHandleResult{
				InstanceID: v.InstanceID.String(),
				Code:       -1,
				ErrMsg:     "instance not found",
			}
			continue
		}

		wg.Add(1)
		go func(v *strategy.InstanceHandler) {
			defer wg.Done()
			res <- i.HandleEvent(ctx, v.Handler, []byte(r.Payload))
		}(v)
	}
	wg.Wait()
	close(res)

	for v := range res {
		if v == nil {
			continue
		}
		ret.WasmResults = append(ret.WasmResults, *v)
	}
	return ret, nil
}

func publisherVerification(ctx context.Context, r *eventpb.Event, l log.Logger) error {
	if r.Header == nil || len(r.Header.Token) == 0 {
		return errors.New("message token is invalid")
	}

	d := types.MustMgrDBExecutorFromContext(ctx)
	p := types.MustProjectFromContext(ctx)
	pj := jwt.MustPublisherAuthFromContext(ctx)

	claim, err := pj.ParseToken(r.Header.Token)
	if err != nil {
		l.Error(err)
		return err
	}

	v, ok := claim.Payload.(string)
	if !ok {
		l.Error(errors.New("claim of publisher convert string error"))
		return status.InvalidAuthValue
	}
	publisherID := types.SFID(0)
	if err := publisherID.UnmarshalText([]byte(v)); err != nil {
		return status.InvalidAuthPublisherID
	}

	m := &models.Publisher{RelPublisher: models.RelPublisher{PublisherID: publisherID}}
	err = m.FetchByPublisherID(d)
	if err != nil {
		l.Error(err)
		return status.CheckDatabaseError(err, "FetchByPublisherID")
	}

	if m.ProjectID == p.ProjectID {
		return nil
	} else {
		return status.NoProjectPermission
	}

}

func HandleEvents(ctx context.Context, projectName string, r *HandleEventReq) []*HandleEventResult {
	results := make([]*HandleEventResult, 0, len(r.Events))
	for i := range r.Events {
		ret, _ := OnEventReceived(ctx, projectName, &r.Events[i])
		results = append(results, ret)
	}
	return results
}
