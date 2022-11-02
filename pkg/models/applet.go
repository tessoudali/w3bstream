package models

import (
	"github.com/machinefi/Bumblebee/base/types"
	"github.com/machinefi/Bumblebee/kit/sqlx/datatypes"
)

// Applet database model applet
// @def primary                          ID
// @def unique_index UI_applet_id        AppletID
// @def unique_index UI_project_name     ProjectID Name
// @def unique_index UI_project_resource ProjectID WasmResourceID
//
//go:generate toolkit gen model Applet --database DB
type Applet struct {
	datatypes.PrimaryID
	RelProject
	RelApplet
	RelWasmResource
	AppletInfo
	datatypes.OperationTimes
}

type RelApplet struct {
	AppletID types.SFID `db:"f_applet_id" json:"appletID"`
}

type AppletInfo struct {
	Name string `db:"f_name" json:"name"`
}
