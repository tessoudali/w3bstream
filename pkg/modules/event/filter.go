package event

type filtration struct{}

func (*filtration) filter(e Event) bool {
	return true
}
