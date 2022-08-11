package models

import "github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"

//go:generate toolkit gen model Applet --database Demo
// Applet database model demo
// @def primary                     ID
// @def unique_index UI_applet_name Name
// @def unique_index UI_applet_id   AppletID
type Applet struct {
	datatypes.PrimaryID
	RefApplet
	AppletInfo
	datatypes.OperationTimes
}

type RefApplet struct {
	AppletID string `db:"f_applet_id" json:"appletID"`
}

type AppletInfo struct {
	Name string `db:"f_name" json:"name"`
}
