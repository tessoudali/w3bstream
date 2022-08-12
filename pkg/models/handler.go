package models

import (
	"database/sql/driver"

	"github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"
)

//go:generate toolkit gen model Handler --database DB
// Handler handler info
// @def primary                               ID
// @def unique_index UI_applet_deploy_handler AppletID DeployID HandlerID
type Handler struct {
	datatypes.PrimaryID
	RelApplet
	RelDeploy
	RelHandler
	HandlerInfo
	datatypes.OperationTimes
}

type RelHandler struct {
	HandlerID string `db:"f_handler_id" json:"handlerID"`
}

type HandlerInfo struct {
	Name   string        `db:"f_name"              json:"name"`
	Params HandlerParams `db:"f_params,default=''" json:"params"`
}

type HandlerParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type HandlerParams []HandlerParam

func (HandlerParams) DataType(engine string) string { return "TEXT" }

func (p HandlerParams) Value() (driver.Value, error) {
	return datatypes.JSONValue(p)
}

func (p *HandlerParams) Scan(src interface{}) error {
	return datatypes.JSONScan(src, p)
}
