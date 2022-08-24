// This is a generated source file. DO NOT EDIT
// Source: demo/demo_m.go

package demo

import (
	"github.com/iotexproject/Bumblebee/base/types"
)

// @def primary                       ID
// @def index        I_nickname/BTREE Name
// @def index        I_username       Username
// @def unique_index UI_name          Name
// @def unique_index UI_id_org        ID OrgID

// Demo demo table
type Demo struct {
	ID        uint64          `db:"f_id,autoincrement"`
	Name      string          `db:"f_name,default=''"`
	Nickname  string          `db:"f_nickname,default=''"`
	Username  string          `db:"f_username,default=''"`
	Gender    int             `db:"f_gender,default='0'"`
	Boolean   bool            `db:"f_boolean,default=false"`
	OrgID     uint64          `db:"f_org_id"`
	CreatedAt types.Timestamp `db:"f_created_at,default='0'"`
	UpdatedAt types.Timestamp `db:"f_updated_at,default='0'"`
	DeletedAt types.Timestamp `db:"f_deleted_at,default='0'"`
}
