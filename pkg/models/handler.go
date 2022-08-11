package models

import (
	"database/sql/driver"

	"github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"
)

//go:generate toolkit gen model Handler --database Demo
// Handler handler info
// @def primary                        ID
// @def unique_index UI_handler_id     HandlerID
// @def unique_index UI_applet_handler AppletID Name
type Handler struct {
	datatypes.PrimaryID
	RefApplet
	RefHandler
	HandlerInfo
}

type RefHandler struct {
	HandlerID string `db:"f_handler_id" json:"handlerID"`
}

type HandlerInfo struct {
	Address    string      `db:"f_address"           json:"address"`
	Network    string      `db:"f_network"           json:"network"`
	WasmFile   string      `db:"f_wasm_file"         json:"wasmFile"`
	AbiFile    string      `db:"f_abi_file"          json:"abiFile"`
	AbiName    string      `db:"f_abi_name"          json:"abiName"`
	AbiVersion string      `db:"f_abi_version"       json:"abiVersion"`
	Name       string      `db:"f_name"              json:"name"`
	Params     EventParams `db:"f_params,default=''" json:"params"`
}

type HandlerParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type EventParams []HandlerParam

func (EventParams) DataType(engine string) string { return "TEXT" }

func (p EventParams) Value() (driver.Value, error) {
	return datatypes.JSONValue(p)
}

func (p *EventParams) Scan(src interface{}) error {
	return datatypes.JSONScan(src, p)
}
