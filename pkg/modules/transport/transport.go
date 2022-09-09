// transport for input/output

package transport

type Transporter interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type Filter func([]byte) (namespace string, applet string, raw []byte, err error)

// TCP | UDP | WS | MQTT.....

type Dispatch interface{}
