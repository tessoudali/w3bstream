package models

import (
	"database/sql/driver"

	"github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"
)

// Applet database model applet
// @def primary                   ID
// @def unique_index UI_applet_id AppletID
// @def index        I_name       Name
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

type AppletConfig struct{}

func (AppletConfig) DataType(drv string) string { return "text" }

func (v AppletConfig) Value() (driver.Value, error) { return nil, nil }

func (v *AppletConfig) Scan(src interface{}) error { return nil }
