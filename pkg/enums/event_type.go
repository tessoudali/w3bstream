package enums

//go:generate toolkit gen enum EventType
type EventType uint8

const (
	EVENT_TYPE_UNKNOWN EventType = iota
)
