package models

import "github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"

// Event database model demo
// @def primary                     ID
// @def unique_index UI_event_id    EventID
//
//go:generate toolkit gen model Event --database DB
type Event struct {
	datatypes.PrimaryID
	RefEventID
	EventInfo
	datatypes.OperationTimes
}

type RefEventID struct {
	EventID string `db:"f_event_id" json:"eventID"`
}

type EventInfo struct {
	RelProject
	RelApplet
	RelHandler
}
