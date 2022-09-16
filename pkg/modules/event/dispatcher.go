package event

type dispatcher struct {
	apiServerURI string
	vmManagerURI string
}

func (d *dispatcher) dispatch(e Event) ([]byte, error) {
	return []byte("Success"), nil
}
