package event

type dispatcher struct{}

func (d *dispatcher) dispatch(e Event) error {
	return nil
}
