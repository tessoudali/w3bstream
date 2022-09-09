package models

import "github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"

// Applet database model demo
// @def primary                     ID
// @def unique_index UI_applet_name Name
// @def unique_index UI_applet_id   AppletID
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
	Name     string `db:"f_name"      json:"name"`
	AssetLoc string `db:"f_asset_loc" json:"assetLoc"`
}
