// This is a generated source file. DO NOT EDIT
// Source: demo/demo__generated.go

package demo

import (
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
)

var DemoTable *builder.Table

func init() {
	DemoTable = DB.Register(&Demo{})
}

type DemoIterator struct {
}

func (DemoIterator) New() interface{} {
	return &Demo{}
}

func (DemoIterator) Resolve(v interface{}) *Demo {
	return v.(*Demo)
}

func (Demo) TableName() string {
	return "t_demo"
}

func (Demo) TableDesc() []string {
	return []string{
		"Demo demo table",
	}
}

func (Demo) Comments() map[string]string {
	return map[string]string{}
}

func (Demo) ColDesc() map[string][]string {
	return map[string][]string{}
}

func (Demo) ColRel() map[string][]string {
	return map[string][]string{}
}

func (Demo) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Demo) Indexes() builder.Indexes {
	return builder.Indexes{
		"i_nickname/BTREE": []string{
			"Name",
		},
		"i_username": []string{
			"Username",
		},
	}
}

func (m *Demo) IndexFieldNames() []string {
	return []string{
		"ID",
		"Name",
		"OrgID",
		"Username",
	}
}

func (Demo) UniqueIndexes() builder.Indexes {
	return builder.Indexes{
		"ui_id_org": []string{
			"ID",
			"OrgID",
			"DeletedAt",
		},
		"ui_name": []string{
			"Name",
			"DeletedAt",
		},
	}
}

func (Demo) UniqueIndexUIIDOrg() string {
	return "ui_id_org"
}

func (Demo) UniqueIndexUIName() string {
	return "ui_name"
}

func (m *Demo) ColID() *builder.Column {
	return DemoTable.ColByFieldName(m.FieldID())
}

func (Demo) FieldID() string {
	return "ID"
}

func (m *Demo) ColName() *builder.Column {
	return DemoTable.ColByFieldName(m.FieldName())
}

func (Demo) FieldName() string {
	return "Name"
}

func (m *Demo) ColNickname() *builder.Column {
	return DemoTable.ColByFieldName(m.FieldNickname())
}

func (Demo) FieldNickname() string {
	return "Nickname"
}

func (m *Demo) ColUsername() *builder.Column {
	return DemoTable.ColByFieldName(m.FieldUsername())
}

func (Demo) FieldUsername() string {
	return "Username"
}

func (m *Demo) ColGender() *builder.Column {
	return DemoTable.ColByFieldName(m.FieldGender())
}

func (Demo) FieldGender() string {
	return "Gender"
}

func (m *Demo) ColBoolean() *builder.Column {
	return DemoTable.ColByFieldName(m.FieldBoolean())
}

func (Demo) FieldBoolean() string {
	return "Boolean"
}

func (m *Demo) ColOrgID() *builder.Column {
	return DemoTable.ColByFieldName(m.FieldOrgID())
}

func (Demo) FieldOrgID() string {
	return "OrgID"
}

func (m *Demo) ColCreatedAt() *builder.Column {
	return DemoTable.ColByFieldName(m.FieldCreatedAt())
}

func (Demo) FieldCreatedAt() string {
	return "CreatedAt"
}

func (m *Demo) ColUpdatedAt() *builder.Column {
	return DemoTable.ColByFieldName(m.FieldUpdatedAt())
}

func (Demo) FieldUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Demo) ColDeletedAt() *builder.Column {
	return DemoTable.ColByFieldName(m.FieldDeletedAt())
}

func (Demo) FieldDeletedAt() string {
	return "DeletedAt"
}
