// This is a generated source file. DO NOT EDIT
// Source: models/runtime_log__generated.go

package models

import (
	"fmt"
	"time"

	"github.com/machinefi/w3bstream/pkg/depends/base/types"
	"github.com/machinefi/w3bstream/pkg/depends/kit/sqlx"
	"github.com/machinefi/w3bstream/pkg/depends/kit/sqlx/builder"
)

var RuntimeLogTable *builder.Table

func init() {
	RuntimeLogTable = DB.Register(&RuntimeLog{})
}

type RuntimeLogIterator struct {
}

func (*RuntimeLogIterator) New() interface{} {
	return &RuntimeLog{}
}

func (*RuntimeLogIterator) Resolve(v interface{}) *RuntimeLog {
	return v.(*RuntimeLog)
}

func (*RuntimeLog) TableName() string {
	return "t_runtime_log"
}

func (*RuntimeLog) TableDesc() []string {
	return []string{
		"RuntimeLog database model event",
	}
}

func (*RuntimeLog) Comments() map[string]string {
	return map[string]string{}
}

func (*RuntimeLog) ColDesc() map[string][]string {
	return map[string][]string{}
}

func (*RuntimeLog) ColRel() map[string][]string {
	return map[string][]string{}
}

func (*RuntimeLog) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (m *RuntimeLog) IndexFieldNames() []string {
	return []string{
		"ID",
		"RuntimeLogID",
	}
}

func (*RuntimeLog) UniqueIndexes() builder.Indexes {
	return builder.Indexes{
		"ui_runtime_log_id": []string{
			"RuntimeLogID",
		},
	}
}

func (*RuntimeLog) UniqueIndexUIRuntimeLogID() string {
	return "ui_runtime_log_id"
}

func (m *RuntimeLog) ColID() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldID())
}

func (*RuntimeLog) FieldID() string {
	return "ID"
}

func (m *RuntimeLog) ColRuntimeLogID() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldRuntimeLogID())
}

func (*RuntimeLog) FieldRuntimeLogID() string {
	return "RuntimeLogID"
}

func (m *RuntimeLog) ColProjectName() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldProjectName())
}

func (*RuntimeLog) FieldProjectName() string {
	return "ProjectName"
}

func (m *RuntimeLog) ColAppletName() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldAppletName())
}

func (*RuntimeLog) FieldAppletName() string {
	return "AppletName"
}

func (m *RuntimeLog) ColSourceName() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldSourceName())
}

func (*RuntimeLog) FieldSourceName() string {
	return "SourceName"
}

func (m *RuntimeLog) ColInstanceID() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldInstanceID())
}

func (*RuntimeLog) FieldInstanceID() string {
	return "InstanceID"
}

func (m *RuntimeLog) ColLevel() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldLevel())
}

func (*RuntimeLog) FieldLevel() string {
	return "Level"
}

func (m *RuntimeLog) ColLogTime() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldLogTime())
}

func (*RuntimeLog) FieldLogTime() string {
	return "LogTime"
}

func (m *RuntimeLog) ColMsg() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldMsg())
}

func (*RuntimeLog) FieldMsg() string {
	return "Msg"
}

func (m *RuntimeLog) ColCreatedAt() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldCreatedAt())
}

func (*RuntimeLog) FieldCreatedAt() string {
	return "CreatedAt"
}

func (m *RuntimeLog) ColUpdatedAt() *builder.Column {
	return RuntimeLogTable.ColByFieldName(m.FieldUpdatedAt())
}

func (*RuntimeLog) FieldUpdatedAt() string {
	return "UpdatedAt"
}

func (m *RuntimeLog) CondByValue(db sqlx.DBExecutor) builder.SqlCondition {
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

func (m *RuntimeLog) Create(db sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt.Set(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt.Set(time.Now())
	}

	_, err := db.Exec(sqlx.InsertToDB(db, m, nil))
	return err
}

func (m *RuntimeLog) List(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) ([]RuntimeLog, error) {
	var (
		tbl = db.T(m)
		lst = make([]RuntimeLog, 0)
	)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("RuntimeLog.List")}, adds...)
	err := db.QueryAndScan(builder.Select(nil).From(tbl, adds...), &lst)
	return lst, err
}

func (m *RuntimeLog) Count(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) (cnt int64, err error) {
	tbl := db.T(m)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("RuntimeLog.List")}, adds...)
	err = db.QueryAndScan(builder.Select(builder.Count()).From(tbl, adds...), &cnt)
	return
}

func (m *RuntimeLog) FetchByID(db sqlx.DBExecutor) error {
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
				builder.Comment("RuntimeLog.FetchByID"),
			),
		m,
	)
	return err
}

func (m *RuntimeLog) FetchByRuntimeLogID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("RuntimeLogID").Eq(m.RuntimeLogID),
					),
				),
				builder.Comment("RuntimeLog.FetchByRuntimeLogID"),
			),
		m,
	)
	return err
}

func (m *RuntimeLog) UpdateByIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

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
				builder.Comment("RuntimeLog.UpdateByIDWithFVs"),
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

func (m *RuntimeLog) UpdateByID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByIDWithFVs(db, fvs)
}

func (m *RuntimeLog) UpdateByRuntimeLogIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("RuntimeLogID").Eq(m.RuntimeLogID),
				),
				builder.Comment("RuntimeLog.UpdateByRuntimeLogIDWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByRuntimeLogID(db)
	}
	return nil
}

func (m *RuntimeLog) UpdateByRuntimeLogID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByRuntimeLogIDWithFVs(db, fvs)
}

func (m *RuntimeLog) Delete(db sqlx.DBExecutor) error {
	_, err := db.Exec(
		builder.Delete().
			From(
				db.T(m),
				builder.Where(m.CondByValue(db)),
				builder.Comment("RuntimeLog.Delete"),
			),
	)
	return err
}

func (m *RuntimeLog) DeleteByID(db sqlx.DBExecutor) error {
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
				builder.Comment("RuntimeLog.DeleteByID"),
			),
	)
	return err
}

func (m *RuntimeLog) DeleteByRuntimeLogID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("RuntimeLogID").Eq(m.RuntimeLogID),
					),
				),
				builder.Comment("RuntimeLog.DeleteByRuntimeLogID"),
			),
	)
	return err
}
