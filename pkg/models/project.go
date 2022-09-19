package models

import (
	"github.com/iotexproject/Bumblebee/kit/sqlx/datatypes"

	"github.com/iotexproject/w3bstream/pkg/enums"
)

// Project schema for project information
// @def primary                      ID
// @def unique_index UI_project_id   ProjectID
// @def unique_index UI_name_version Name Version
//
//go:generate toolkit gen model Project --database DB
type Project struct {
	datatypes.PrimaryID
	RelProject
	RelAccount
	ProjectInfo
	ProjectSchema
	datatypes.OperationTimesWithDeleted
}

type RelProject struct {
	ProjectID string `db:"f_project_id" json:"projectID"`
}

type ProjectInfo struct {
	Name     string         `db:"f_name"                 json:"name"`               // Name project name
	Version  string         `db:"f_version"              json:"version"`            // Version project version
	Protocol enums.Protocol `db:"f_protocol,default='0'" json:"protocol,omitempty"` // Protocol project protocol for event publisher
}

type ProjectSchema struct {
	Schema string `db:"f_schema,default=''" json:"-"`
}
