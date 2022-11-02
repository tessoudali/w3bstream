package models

import (
	"github.com/machinefi/Bumblebee/base/types"
	"github.com/machinefi/Bumblebee/kit/sqlx/datatypes"
)

// WasmResource database model wasm_resource
// @def primary                            ID
// @def unique_index UI_wasm_resource_id   WasmResourceID
// @def unique_index UI_project_path_md5   ProjectID Path Md5
//
//go:generate toolkit gen model WasmResource --database DB
type WasmResource struct {
	datatypes.PrimaryID
	RelProject
	RelWasmResource
	WasmResourceInfo
	datatypes.OperationTimes
}

type RelWasmResource struct {
	WasmResourceID types.SFID `db:"f_wasm_resource_id" json:"wasmResourceID"`
}

type WasmResourceInfo struct {
	Name string `db:"f_name" json:"wasmName"`
	Path string `db:"f_path" json:"-"`
	Md5  string `db:"f_md5"  json:"-"`
}
