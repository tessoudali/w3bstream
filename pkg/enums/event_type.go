package enums

//go:generate toolkit gen enum EventType
type EventType int

const (
	EVENT_TYPE_UNKNOWN EventType = iota
	EVENT_TYPE__EXAMP1
	EVENT_TYPE__EXAMP2
	EVENT_TYPE__EXAMP3
	EVENT_TYPE__EXAMP4
	EVENT_TYPE__EXAMP5
	EVENT_TYPE__EXAMP6
)

const EVENT_TYPE__ANY = EVENT_TYPE_UNKNOWN + 0x7FFFFFFF // any event type
