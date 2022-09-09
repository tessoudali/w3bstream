// This is a generated source file. DO NOT EDIT
// Source: models/handler__generated.go

package models

import (
	"fmt"
	"time"

	"github.com/iotexproject/Bumblebee/base/types"
	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
)

var HandlerTable *builder.Table

func init() {
	HandlerTable = DB.Register(&Handler{})
}

type HandlerIterator struct {
}

func (HandlerIterator) New() interface{} {
	return &Handler{}
}

func (HandlerIterator) Resolve(v interface{}) *Handler {
	return v.(*Handler)
}

func (*Handler) TableName() string {
	return "t_handler"
}

func (*Handler) TableDesc() []string {
	return []string{
		"Handler handler info",
	}
}

func (*Handler) Comments() map[string]string {
	return map[string]string{}
}

func (*Handler) ColDesc() map[string][]string {
	return map[string][]string{}
}

func (*Handler) ColRel() map[string][]string {
	return map[string][]string{}
}

func (*Handler) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (m *Handler) IndexFieldNames() []string {
	return []string{
		"AppletID",
		"DeployID",
		"HandlerID",
		"ID",
	}
}

func (*Handler) UniqueIndexes() builder.Indexes {
	return builder.Indexes{
		"ui_applet_deploy_handler": []string{
			"AppletID",
			"DeployID",
			"HandlerID",
		},
	}
}

func (*Handler) UniqueIndexUIAppletDeployHandler() string {
	return "ui_applet_deploy_handler"
}

func (m *Handler) ColID() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldID())
}

func (*Handler) FieldID() string {
	return "ID"
}

func (m *Handler) ColAppletID() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldAppletID())
}

func (*Handler) FieldAppletID() string {
	return "AppletID"
}

func (m *Handler) ColDeployID() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldDeployID())
}

func (*Handler) FieldDeployID() string {
	return "DeployID"
}

func (m *Handler) ColHandlerID() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldHandlerID())
}

func (*Handler) FieldHandlerID() string {
	return "HandlerID"
}

func (m *Handler) ColName() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldName())
}

func (*Handler) FieldName() string {
	return "Name"
}

func (m *Handler) ColHandler() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldHandler())
}

func (*Handler) FieldHandler() string {
	return "Handler"
}

func (m *Handler) ColParams() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldParams())
}

func (*Handler) FieldParams() string {
	return "Params"
}

func (m *Handler) ColCreatedAt() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldCreatedAt())
}

func (*Handler) FieldCreatedAt() string {
	return "CreatedAt"
}

func (m *Handler) ColUpdatedAt() *builder.Column {
	return HandlerTable.ColByFieldName(m.FieldUpdatedAt())
}

func (*Handler) FieldUpdatedAt() string {
	return "UpdatedAt"
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

	if m.CreatedAt.IsZero() {
		m.CreatedAt.Set(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt.Set(time.Now())
	}

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

func (m *Handler) FetchByAppletIDAndDeployIDAndHandlerID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("AppletID").Eq(m.AppletID),
						tbl.ColByFieldName("DeployID").Eq(m.DeployID),
						tbl.ColByFieldName("HandlerID").Eq(m.HandlerID),
					),
				),
				builder.Comment("Handler.FetchByAppletIDAndDeployIDAndHandlerID"),
			),
		m,
	)
	return err
}

func (m *Handler) UpdateByIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

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

func (m *Handler) UpdateByAppletIDAndDeployIDAndHandlerIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("AppletID").Eq(m.AppletID),
					tbl.ColByFieldName("DeployID").Eq(m.DeployID),
					tbl.ColByFieldName("HandlerID").Eq(m.HandlerID),
				),
				builder.Comment("Handler.UpdateByAppletIDAndDeployIDAndHandlerIDWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByAppletIDAndDeployIDAndHandlerID(db)
	}
	return nil
}

func (m *Handler) UpdateByAppletIDAndDeployIDAndHandlerID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByAppletIDAndDeployIDAndHandlerIDWithFVs(db, fvs)
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

func (m *Handler) DeleteByAppletIDAndDeployIDAndHandlerID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("AppletID").Eq(m.AppletID),
						tbl.ColByFieldName("DeployID").Eq(m.DeployID),
						tbl.ColByFieldName("HandlerID").Eq(m.HandlerID),
					),
				),
				builder.Comment("Handler.DeleteByAppletIDAndDeployIDAndHandlerID"),
			),
	)
	return err
}
