package models

import "github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"

// AppletDeploy applet deploy info
// @def primary                        ID
// @def index        I_applet_id       AppletID
// @def unique_index UI_deploy_id      DeployID
// @def unique_index UI_deploy_version AppletID Version
//
//go:generate toolkit gen model AppletDeploy --database DB
type AppletDeploy struct {
	datatypes.PrimaryID
	RelApplet
	RelDeploy
	DeployInfo
	datatypes.OperationTimes
}

type RelDeploy struct {
	DeployID string `db:"f_deploy_id" json:"deployID"`
}

type DeployInfo struct {
	Location string `db:"f_location,default=''" json:"location"`
	Version  string `db:"f_version,default=''"  json:"version"`
	WasmFile string `db:"f_wasm_file"           json:"wasmFile"`
	AbiName  string `db:"f_abi_loc"             json:"abiName"`
	AbiFile  string `db:"f_abi_file"            json:"abiFile"`
}
