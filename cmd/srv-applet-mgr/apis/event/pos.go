package event

import (
	"context"
	"unicode/utf8"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/depends/unit"
	me "github.com/iotexproject/w3bstream/pkg/modules/event"
	"github.com/iotexproject/w3bstream/pkg/types"
)

const (
	strLenLimit   = 50
	dataSizeLimit = 2 * unit.KiB
)

// TODO should define to pkg/depends/protocol/eventpb/event

type RecvEvent struct {
	httpx.MethodPost
	AppletID  string `in:"path" name:"applet"`
	Handler   string `in:"path" name:"handler"`
	Publisher string `in:"header" name:"publisher"`
	Data      []byte `in:"body" mime:"application/octet-stream"`
}

func (r *RecvEvent) Path() string {
	return "/applet/:applet/handler/:handler"
}

func (r *RecvEvent) Output(ctx context.Context) (interface{}, error) {
	if !check(r.AppletID, r.Handler, r.Publisher) {
		return nil, errParamIllegal
	}
	if len(r.Data) > dataSizeLimit {
		return nil, errParamIllegal
	}
	events := types.MustEventChanFromContext(ctx)

	res := make(chan me.Result)
	events <- &event{
		handler:     r.Handler,
		appletID:    r.AppletID,
		publisherID: r.Publisher,
		data:        r.Data,
		result:      res,
	}
	// TODO timeout
	result := <-res
	if !result.Success {
		return nil, errInternalSystemError
	}
	return result.Data, nil
}

func check(appletID, handler, publisher string) bool {
	if l := utf8.RuneCountInString(appletID); l <= 0 || l > strLenLimit {
		return false
	}
	if l := utf8.RuneCountInString(publisher); l <= 0 || l > strLenLimit {
		return false
	}
	return utf8.RuneCountInString(handler) <= strLenLimit
}
