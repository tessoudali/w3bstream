// This is a generated source file. DO NOT EDIT
// Source: models/event__generated.go

package models

import (
	"fmt"
	"time"

	"github.com/iotexproject/Bumblebee/base/types"
	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
)

var EventTable *builder.Table

func init() {
	EventTable = DB.Register(&Event{})
}

type EventIterator struct {
}

func (EventIterator) New() interface{} {
	return &Event{}
}

func (EventIterator) Resolve(v interface{}) *Event {
	return v.(*Event)
}

func (*Event) TableName() string {
	return "t_event"
}

func (*Event) TableDesc() []string {
	return []string{
		"Event database model demo",
	}
}

func (*Event) Comments() map[string]string {
	return map[string]string{}
}

func (*Event) ColDesc() map[string][]string {
	return map[string][]string{}
}

func (*Event) ColRel() map[string][]string {
	return map[string][]string{}
}

func (*Event) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (m *Event) IndexFieldNames() []string {
	return []string{
		"EventID",
		"ID",
	}
}

func (*Event) UniqueIndexes() builder.Indexes {
	return builder.Indexes{
		"ui_event_id": []string{
			"EventID",
		},
	}
}

func (*Event) UniqueIndexUIEventID() string {
	return "ui_event_id"
}

func (m *Event) ColID() *builder.Column {
	return EventTable.ColByFieldName(m.FieldID())
}

func (*Event) FieldID() string {
	return "ID"
}

func (m *Event) ColEventID() *builder.Column {
	return EventTable.ColByFieldName(m.FieldEventID())
}

func (*Event) FieldEventID() string {
	return "EventID"
}

func (m *Event) ColAppletID() *builder.Column {
	return EventTable.ColByFieldName(m.FieldAppletID())
}

func (*Event) FieldAppletID() string {
	return "AppletID"
}

func (m *Event) ColHandlerID() *builder.Column {
	return EventTable.ColByFieldName(m.FieldHandlerID())
}

func (*Event) FieldHandlerID() string {
	return "HandlerID"
}

func (m *Event) ColCreatedAt() *builder.Column {
	return EventTable.ColByFieldName(m.FieldCreatedAt())
}

func (*Event) FieldCreatedAt() string {
	return "CreatedAt"
}

func (m *Event) ColUpdatedAt() *builder.Column {
	return EventTable.ColByFieldName(m.FieldUpdatedAt())
}

func (*Event) FieldUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Event) CondByValue(db sqlx.DBExecutor) builder.SqlCondition {
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

func (m *Event) Create(db sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt.Set(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt.Set(time.Now())
	}

	_, err := db.Exec(sqlx.InsertToDB(db, m, nil))
	return err
}

func (m *Event) List(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) ([]Event, error) {
	var (
		tbl = db.T(m)
		lst = make([]Event, 0)
	)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("Event.List")}, adds...)
	err := db.QueryAndScan(builder.Select(nil).From(tbl, adds...), &lst)
	return lst, err
}

func (m *Event) Count(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) (cnt int64, err error) {
	tbl := db.T(m)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("Event.List")}, adds...)
	err = db.QueryAndScan(builder.Select(builder.Count()).From(tbl, adds...), &cnt)
	return
}

func (m *Event) FetchByID(db sqlx.DBExecutor) error {
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
				builder.Comment("Event.FetchByID"),
			),
		m,
	)
	return err
}

func (m *Event) FetchByEventID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("EventID").Eq(m.EventID),
					),
				),
				builder.Comment("Event.FetchByEventID"),
			),
		m,
	)
	return err
}

func (m *Event) UpdateByIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

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
				builder.Comment("Event.UpdateByIDWithFVs"),
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

func (m *Event) UpdateByID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByIDWithFVs(db, fvs)
}

func (m *Event) UpdateByEventIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("EventID").Eq(m.EventID),
				),
				builder.Comment("Event.UpdateByEventIDWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByEventID(db)
	}
	return nil
}

func (m *Event) UpdateByEventID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByEventIDWithFVs(db, fvs)
}

func (m *Event) Delete(db sqlx.DBExecutor) error {
	_, err := db.Exec(
		builder.Delete().
			From(
				db.T(m),
				builder.Where(m.CondByValue(db)),
				builder.Comment("Event.Delete"),
			),
	)
	return err
}

func (m *Event) DeleteByID(db sqlx.DBExecutor) error {
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
				builder.Comment("Event.DeleteByID"),
			),
	)
	return err
}

func (m *Event) DeleteByEventID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("EventID").Eq(m.EventID),
					),
				),
				builder.Comment("Event.DeleteByEventID"),
			),
	)
	return err
}
