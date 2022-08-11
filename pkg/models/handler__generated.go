// This is a generated source file. DO NOT EDIT
// Source: models/handler__generated.go

package models

import (
	"fmt"

	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
)

var HandlerTable *builder.Table

func init() {
	HandlerTable = Demo.Register(&Handler{})
}

type HandlerIterator struct {
}

func (HandlerIterator) New() interface{} {
	return &Handler{}
}

func (HandlerIterator) Resolve(v interface{}) *Handler {
	return v.(*Handler)
}

func (Handler) TableName() string {
	return "t_handler"
}

func (Handler) TableDesc() []string {
	return []string{
		"Handler handler info",
	}
}

func (Handler) Comments() map[string]string {
	return map[string]string{}
}

func (Handler) ColDesc() map[string][]string {
	return map[string][]string{}
}

func (Handler) ColRel() map[string][]string {
	return map[string][]string{}
}

func (Handler) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (m *Handler) IndexFieldNames() []string {
	return []string{
		"AppletID",
		"HandlerID",
		"ID",
		"Name",
	}
}

func (Handler) UniqueIndexes() builder.Indexes {
	return builder.Indexes{
		"ui_applet_handler": []string{
			"AppletID",
			"Name",
		},
		"ui_handler_id": []string{
			"HandlerID",
		},
	}
}

func (Handler) UniqueIndexUiAppletHandler() string {
	return "ui_applet_handler"
}

func (Handler) UniqueIndexUiHandlerId() string {
	return "ui_handler_id"
}

func (m *Handler) ColID() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldID())
}

func (Handler) FieldID() string {
	return "ID"
}

func (m *Handler) ColAppletID() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldAppletID())
}

func (Handler) FieldAppletID() string {
	return "AppletID"
}

func (m *Handler) ColHandlerID() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldHandlerID())
}

func (Handler) FieldHandlerID() string {
	return "HandlerID"
}

func (m *Handler) ColAddress() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldAddress())
}

func (Handler) FieldAddress() string {
	return "Address"
}

func (m *Handler) ColNetwork() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldNetwork())
}

func (Handler) FieldNetwork() string {
	return "Network"
}

func (m *Handler) ColWasmFile() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldWasmFile())
}

func (Handler) FieldWasmFile() string {
	return "WasmFile"
}

func (m *Handler) ColAbiFile() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldAbiFile())
}

func (Handler) FieldAbiFile() string {
	return "AbiFile"
}

func (m *Handler) ColAbiName() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldAbiName())
}

func (Handler) FieldAbiName() string {
	return "AbiName"
}

func (m *Handler) ColAbiVersion() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldAbiVersion())
}

func (Handler) FieldAbiVersion() string {
	return "AbiVersion"
}

func (m *Handler) ColName() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldName())
}

func (Handler) FieldName() string {
	return "Name"
}

func (m *Handler) ColParams() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldParams())
}

func (Handler) FieldParams() string {
	return "Params"
}

func (m *Handler) CondByValue(db sqlx.DBExecutor) builder.SqlCondition {
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

func (m *Handler) Create(db sqlx.DBExecutor) error {

	_, err := db.Exec(sqlx.InsertToDB(db, m, nil))
	return err
}

func (m *Handler) List(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) ([]Handler, error) {
	var (
		tbl = db.T(m)
		lst = make([]Handler, 0)
	)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("Handler.List")}, adds...)
	err := db.QueryAndScan(builder.Select(nil).From(tbl, adds...), &lst)
	return lst, err
}

func (m *Handler) Count(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) (cnt int64, err error) {
	tbl := db.T(m)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("Handler.List")}, adds...)
	err = db.QueryAndScan(builder.Select(builder.Count()).From(tbl, adds...), &cnt)
	return
}

func (m *Handler) FetchByID(db sqlx.DBExecutor) error {
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
				builder.Comment("Handler.FetchByID"),
			),
		m,
	)
	return err
}

func (m *Handler) FetchByHandlerID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("HandlerID").Eq(m.HandlerID),
					),
				),
				builder.Comment("Handler.FetchByHandlerID"),
			),
		m,
	)
	return err
}

func (m *Handler) FetchByAppletIDAndName(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("AppletID").Eq(m.AppletID),
						tbl.ColByFieldName("Name").Eq(m.Name),
					),
				),
				builder.Comment("Handler.FetchByAppletIDAndName"),
			),
		m,
	)
	return err
}

func (m *Handler) UpdateByIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("ID").Eq(m.ID),
				),
				builder.Comment("Handler.UpdateByIDWithFVs"),
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

func (m *Handler) UpdateByID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByIDWithFVs(db, fvs)
}

func (m *Handler) UpdateByHandlerIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("HandlerID").Eq(m.HandlerID),
				),
				builder.Comment("Handler.UpdateByHandlerIDWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByHandlerID(db)
	}
	return nil
}

func (m *Handler) UpdateByHandlerID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByHandlerIDWithFVs(db, fvs)
}

func (m *Handler) UpdateByAppletIDAndNameWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("AppletID").Eq(m.AppletID),
					tbl.ColByFieldName("Name").Eq(m.Name),
				),
				builder.Comment("Handler.UpdateByAppletIDAndNameWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByAppletIDAndName(db)
	}
	return nil
}

func (m *Handler) UpdateByAppletIDAndName(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByAppletIDAndNameWithFVs(db, fvs)
}

func (m *Handler) Delete(db sqlx.DBExecutor) error {
	_, err := db.Exec(
		builder.Delete().
			From(
				db.T(m),
				builder.Where(m.CondByValue(db)),
				builder.Comment("Handler.Delete"),
			),
	)
	return err
}

func (m *Handler) DeleteByID(db sqlx.DBExecutor) error {
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
				builder.Comment("Handler.DeleteByID"),
			),
	)
	return err
}

func (m *Handler) DeleteByHandlerID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("HandlerID").Eq(m.HandlerID),
					),
				),
				builder.Comment("Handler.DeleteByHandlerID"),
			),
	)
	return err
}

func (m *Handler) DeleteByAppletIDAndName(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("AppletID").Eq(m.AppletID),
						tbl.ColByFieldName("Name").Eq(m.Name),
					),
				),
				builder.Comment("Handler.DeleteByAppletIDAndName"),
			),
	)
	return err
}
