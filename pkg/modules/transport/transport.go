// transport for input/output

package transport

type Transporter interface {
	OnConnected() // tcp
	OnEvent()
	OnEventDone()
}

type TransporterIO interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type EventHeader interface {
}

type Event interface {
	Route() Route // -> filter
	Raw() []byte  // vm
}

type Route struct{}

type Filter func([]byte) (namespace string, applet string, raw []byte, err error)

type Dispatch interface{}
