package http

import me "github.com/iotexproject/w3bstream/pkg/modules/event"

type event struct {
	project   string
	applet    string
	publisher string
	data      []byte
}

func (e *event) Meta() me.MetaData {
	return me.MetaData{
		PublisherID: e.publisher,
		ProjectID:   e.project,
		AppletID:    e.applet,
	}
}
func (e *event) Raw() []byte {
	return e.data
}
