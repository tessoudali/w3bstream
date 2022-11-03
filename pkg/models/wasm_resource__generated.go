// This is a generated source file. DO NOT EDIT
// Source: models/wasm_resource__generated.go

package models

import (
	"fmt"
	"time"

	"github.com/machinefi/w3bstream/pkg/depends/base/types"
	"github.com/machinefi/w3bstream/pkg/depends/kit/sqlx"
	"github.com/machinefi/w3bstream/pkg/depends/kit/sqlx/builder"
)

var WasmResourceTable *builder.Table

func init() {
	WasmResourceTable = DB.Register(&WasmResource{})
}

type WasmResourceIterator struct {
}

func (*WasmResourceIterator) New() interface{} {
	return &WasmResource{}
}

func (*WasmResourceIterator) Resolve(v interface{}) *WasmResource {
	return v.(*WasmResource)
}

func (*WasmResource) TableName() string {
	return "t_wasm_resource"
}

func (*WasmResource) TableDesc() []string {
	return []string{
		"WasmResource database model wasm_resource",
	}
}

func (*WasmResource) Comments() map[string]string {
	return map[string]string{}
}

func (*WasmResource) ColDesc() map[string][]string {
	return map[string][]string{}
}

func (*WasmResource) ColRel() map[string][]string {
	return map[string][]string{}
}

func (*WasmResource) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (m *WasmResource) IndexFieldNames() []string {
	return []string{
		"ID",
		"Md5",
		"WasmResourceID",
	}
}

func (*WasmResource) UniqueIndexes() builder.Indexes {
	return builder.Indexes{
		"ui_md5": []string{
			"Md5",
		},
		"ui_wasm_resource_id": []string{
			"WasmResourceID",
		},
	}
}

func (*WasmResource) UniqueIndexUIMd5() string {
	return "ui_md5"
}

func (*WasmResource) UniqueIndexUIWasmResourceID() string {
	return "ui_wasm_resource_id"
}

func (m *WasmResource) ColID() *builder.Column {
	return WasmResourceTable.ColByFieldName(m.FieldID())
}

func (*WasmResource) FieldID() string {
	return "ID"
}

func (m *WasmResource) ColWasmResourceID() *builder.Column {
	return WasmResourceTable.ColByFieldName(m.FieldWasmResourceID())
}

func (*WasmResource) FieldWasmResourceID() string {
	return "WasmResourceID"
}

func (m *WasmResource) ColPath() *builder.Column {
	return WasmResourceTable.ColByFieldName(m.FieldPath())
}

func (*WasmResource) FieldPath() string {
	return "Path"
}

func (m *WasmResource) ColMd5() *builder.Column {
	return WasmResourceTable.ColByFieldName(m.FieldMd5())
}

func (*WasmResource) FieldMd5() string {
	return "Md5"
}

func (m *WasmResource) ColRefCnt() *builder.Column {
	return WasmResourceTable.ColByFieldName(m.FieldRefCnt())
}

func (*WasmResource) FieldRefCnt() string {
	return "RefCnt"
}

func (m *WasmResource) ColCreatedAt() *builder.Column {
	return WasmResourceTable.ColByFieldName(m.FieldCreatedAt())
}

func (*WasmResource) FieldCreatedAt() string {
	return "CreatedAt"
}

func (m *WasmResource) ColUpdatedAt() *builder.Column {
	return WasmResourceTable.ColByFieldName(m.FieldUpdatedAt())
}

func (*WasmResource) FieldUpdatedAt() string {
	return "UpdatedAt"
}

func (m *WasmResource) CondByValue(db sqlx.DBExecutor) builder.SqlCondition {
	var (
		tbl  = db.T(m)
		fvs  = builder.FieldValueFromStructByNoneZero(m)
		cond = make([]builder.SqlCondition, 0)
	)

	for _, fn := range m.IndexFieldNames() {
		if v, ok := fvs[fn]; ok {
			cond = append(cond, tbl.ColByFieldName(fn).Eq(v))
			delete(fvs, fn)
		}
	}
	if len(cond) == 0 {
		panic(fmt.Errorf("no field for indexes has value"))
	}
	for fn, v := range fvs {
		cond = append(cond, tbl.ColByFieldName(fn).Eq(v))
	}
	return builder.And(cond...)
}

func (m *WasmResource) Create(db sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt.Set(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt.Set(time.Now())
	}

	_, err := db.Exec(sqlx.InsertToDB(db, m, nil))
	return err
}

func (m *WasmResource) List(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) ([]WasmResource, error) {
	var (
		tbl = db.T(m)
		lst = make([]WasmResource, 0)
	)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("WasmResource.List")}, adds...)
	err := db.QueryAndScan(builder.Select(nil).From(tbl, adds...), &lst)
	return lst, err
}

func (m *WasmResource) Count(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) (cnt int64, err error) {
	tbl := db.T(m)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("WasmResource.List")}, adds...)
	err = db.QueryAndScan(builder.Select(builder.Count()).From(tbl, adds...), &cnt)
	return
}

func (m *WasmResource) FetchByID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("ID").Eq(m.ID),
					),
				),
				builder.Comment("WasmResource.FetchByID"),
			),
		m,
	)
	return err
}

func (m *WasmResource) FetchByMd5(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("Md5").Eq(m.Md5),
					),
				),
				builder.Comment("WasmResource.FetchByMd5"),
			),
		m,
	)
	return err
}

func (m *WasmResource) FetchByWasmResourceID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("WasmResourceID").Eq(m.WasmResourceID),
					),
				),
				builder.Comment("WasmResource.FetchByWasmResourceID"),
			),
		m,
	)
	return err
}

func (m *WasmResource) UpdateByIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("ID").Eq(m.ID),
				),
				builder.Comment("WasmResource.UpdateByIDWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByID(db)
	}
	return nil
}

func (m *WasmResource) UpdateByID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByIDWithFVs(db, fvs)
}

func (m *WasmResource) UpdateByMd5WithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("Md5").Eq(m.Md5),
				),
				builder.Comment("WasmResource.UpdateByMd5WithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByMd5(db)
	}
	return nil
}

func (m *WasmResource) UpdateByMd5(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByMd5WithFVs(db, fvs)
}

func (m *WasmResource) UpdateByWasmResourceIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("WasmResourceID").Eq(m.WasmResourceID),
				),
				builder.Comment("WasmResource.UpdateByWasmResourceIDWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByWasmResourceID(db)
	}
	return nil
}

func (m *WasmResource) UpdateByWasmResourceID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByWasmResourceIDWithFVs(db, fvs)
}

func (m *WasmResource) Delete(db sqlx.DBExecutor) error {
	_, err := db.Exec(
		builder.Delete().
			From(
				db.T(m),
				builder.Where(m.CondByValue(db)),
				builder.Comment("WasmResource.Delete"),
			),
	)
	return err
}

func (m *WasmResource) DeleteByID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("ID").Eq(m.ID),
					),
				),
				builder.Comment("WasmResource.DeleteByID"),
			),
	)
	return err
}

func (m *WasmResource) DeleteByMd5(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("Md5").Eq(m.Md5),
					),
				),
				builder.Comment("WasmResource.DeleteByMd5"),
			),
	)
	return err
}

func (m *WasmResource) DeleteByWasmResourceID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("WasmResourceID").Eq(m.WasmResourceID),
					),
				),
				builder.Comment("WasmResource.DeleteByWasmResourceID"),
			),
	)
	return err
}
