package models

import (
	"database/sql/driver"

	"github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"

	"github.com/iotexproject/w3bstream/pkg/enums"
)

// Applet database model applet
// @def primary                   ID
// @def unique_index UI_applet_id AppletID
// @def unique_index UI_name      Name
//
//go:generate toolkit gen model Applet --database DB
type Applet struct {
	datatypes.PrimaryID
	RelProject
	RelApplet
	AppletInfo
	datatypes.OperationTimes
}

type RelApplet struct {
	AppletID string `db:"f_applet_id" json:"appletID"`
}

type AppletInfo struct {
	Name   string        `db:"f_name"              json:"name"`
	Path   string        `db:"f_path"              json:"-"`
	Config *AppletConfig `db:"f_config,default=''" json:"config"`
}

type EventHandler struct {
	Type     enums.EventType `json:"type"`
	Handlers []string        `json:"handlers"`
}

var DefaultEventHandler = EventHandler{enums.EVENT_TYPE_UNKNOWN, []string{"start"}}

type AppletConfig []EventHandler

func (AppletConfig) DataType(drv string) string { return "text" }

func (v AppletConfig) Value() (driver.Value, error) {
	return datatypes.JSONValue(v)
}

func (v *AppletConfig) Scan(src interface{}) error {
	return datatypes.JSONScan(src, v)
}
