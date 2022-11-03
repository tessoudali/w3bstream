package models

import (
	"github.com/machinefi/w3bstream/pkg/depends/base/types"
	"github.com/machinefi/w3bstream/pkg/depends/kit/sqlx/datatypes"
)

// WasmResource database model wasm_resource
// @def primary                            ID
// @def unique_index UI_wasm_resource_id   WasmResourceID
// @def unique_index UI_md5                Md5
//
//go:generate toolkit gen model WasmResource --database DB
type WasmResource struct {
	datatypes.PrimaryID
	RelWasmResource
	WasmResourceInfo
	datatypes.OperationTimes
}

type RelWasmResource struct {
	WasmResourceID types.SFID `db:"f_wasm_resource_id" json:"wasmResourceID"`
}

type WasmResourceInfo struct {
	Path   string `db:"f_path,default=''" json:"-"`
	Md5    string `db:"f_md5"  json:"-"`
	RefCnt int    `db:"f_ref_cnt,default='0'"  json:"-"`
}
