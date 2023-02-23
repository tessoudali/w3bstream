package models

import (
	"github.com/machinefi/w3bstream/pkg/depends/base/types"
	"github.com/machinefi/w3bstream/pkg/depends/kit/sqlx/datatypes"
)

// RuntimeLog database model event
// @def primary                           ID
// @def unique_index UI_runtime_log_id    RuntimeLogID
//
//go:generate toolkit gen model RuntimeLog --database DB
type RuntimeLog struct {
	datatypes.PrimaryID
	RelRuntimeLog
	RuntimeLogInfo
	datatypes.OperationTimes
}

type RelRuntimeLog struct {
	RuntimeLogID types.SFID `db:"f_runtime_log_id" json:"runtimeLogID"`
}

type RuntimeLogInfo struct {
	ProjectName string          `db:"f_project_name" json:"projectName"`
	AppletName  string          `db:"f_applet_name,default=''" json:"appletName"`
	SourceName  string          `db:"f_source_name,default=''" json:"sourceName"`
	InstanceID  types.SFID      `db:"f_instance_id,default='0'" json:"instanceID"`
	Level       string          `db:"f_level,default=''" json:"level"`
	LogTime     types.Timestamp `db:"f_log_time,default='0'" json:"logTime"`
	Msg         string          `db:"f_msg,default=''" json:"msg"`
}
