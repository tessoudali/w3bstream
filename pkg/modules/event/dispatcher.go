package event

type dispatcher struct {
	apiServerURI string
	vmManagerURI string
}

func (d *dispatcher) dispatch(e Event) error {
	id, err := d.vmInstance(e.Meta())
	if err != nil {
		return err
	}
	return d.vmCall(id, e.Raw())
}

func (d *dispatcher) vmInstance(e MetaData) (string, error) {
	return "", nil
}

func (d *dispatcher) vmCall(vmID string, data []byte) error {
	return nil
}
