package http

import me "github.com/iotexproject/w3bstream/pkg/modules/event"

type event struct {
	projectID   string
	appletID    string
	publisherID string
	data        []byte
	result      chan me.Result
}

func (e *event) Meta() me.MetaData {
	return me.MetaData{
		PublisherID: e.publisherID,
		ProjectID:   e.projectID,
		AppletID:    e.appletID,
	}
}

func (e *event) Raw() []byte {
	return e.data
}

func (e *event) ResultChan() chan<- me.Result {
	return e.result
}
