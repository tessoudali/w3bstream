package mock

type Event struct {
	AppletID string  `json:"appletID"`
	Hdl      Handler `json:"handler"`
}

type Handler struct {
	Name    string
	Inputs  []Input
	Outputs []Output
}

type Input struct{}

type Output struct{}
