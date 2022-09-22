package event

import (
	"context"
	"unicode/utf8"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/depends/unit"
	"github.com/iotexproject/w3bstream/pkg/errors/status"
	me "github.com/iotexproject/w3bstream/pkg/modules/event"
	"github.com/iotexproject/w3bstream/pkg/modules/event/proxy"
)

const (
	strLenLimit   = 50
	dataSizeLimit = 2 * unit.KiB
)

// TODO should define to pkg/depends/protocol/eventpb/event

type RecvEvent struct {
	httpx.MethodPost
	ProjectID string `in:"path" name:"project"`
	AppletID  string `in:"path" name:"applet"`
	Handler   string `in:"path" name:"handler"`
	Publisher string `in:"header" name:"publisher"`
	Data      string `in:"body" name:"data"`
}

func (r *RecvEvent) Path() string {
	return "/project/:project/applet/:applet/handler/:handler"
}

func (r *RecvEvent) Output(ctx context.Context) (interface{}, error) {
	if !check(r.ProjectID, r.AppletID, r.Handler, r.Publisher) {
		return nil, status.BadRequest
	}
	if len(r.Data) > dataSizeLimit {
		return nil, status.BadRequest
	}

	res := make(chan me.Result)
	proxy.Proxy(ctx, &event{
		projectID:   r.ProjectID,
		handler:     r.Handler,
		appletID:    r.AppletID,
		publisherID: r.Publisher,
		data:        []byte(r.Data),
		result:      res,
	})
	// TODO timeout
	result := <-res
	if !result.Success {
		return nil, status.InternalServerError
	}
	return result.Data, nil
}

func check(projectID, appletID, handler, publisher string) bool {
	if l := utf8.RuneCountInString(projectID); l <= 0 || l > strLenLimit {
		return false
	}
	if l := utf8.RuneCountInString(appletID); l <= 0 || l > strLenLimit {
		return false
	}
	if l := utf8.RuneCountInString(publisher); l <= 0 || l > strLenLimit {
		return false
	}
	return utf8.RuneCountInString(handler) <= strLenLimit
}
