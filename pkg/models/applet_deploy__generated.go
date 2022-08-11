// This is a generated source file. DO NOT EDIT
// Source: models/applet_deploy__generated.go

package models

import (
	"fmt"
	"time"

	"github.com/iotexproject/Bumblebee/base/types"
	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
)

var AppletDeployTable *builder.Table

func init() {
	AppletDeployTable = Demo.Register(&AppletDeploy{})
}

type AppletDeployIterator struct {
}

func (AppletDeployIterator) New() interface{} {
	return &AppletDeploy{}
}

func (AppletDeployIterator) Resolve(v interface{}) *AppletDeploy {
	return v.(*AppletDeploy)
}

func (AppletDeploy) TableName() string {
	return "t_applet_deploy"
}

func (AppletDeploy) TableDesc() []string {
	return []string{
		"AppletDeploy applet deploy info",
	}
}

func (AppletDeploy) Comments() map[string]string {
	return map[string]string{}
}

func (AppletDeploy) ColDesc() map[string][]string {
	return map[string][]string{}
}

func (AppletDeploy) ColRel() map[string][]string {
	return map[string][]string{}
}

func (AppletDeploy) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (AppletDeploy) Indexes() builder.Indexes {
	return builder.Indexes{
		"i_applet_id": []string{
			"AppletID",
		},
	}
}

func (m *AppletDeploy) IndexFieldNames() []string {
	return []string{
		"AppletID",
		"DeployID",
		"ID",
		"Version",
	}
}

func (AppletDeploy) UniqueIndexes() builder.Indexes {
	return builder.Indexes{
		"ui_deploy_id": []string{
			"DeployID",
		},
		"ui_deploy_version": []string{
			"AppletID",
			"Version",
		},
	}
}

func (AppletDeploy) UniqueIndexUiDeployId() string {
	return "ui_deploy_id"
}

func (AppletDeploy) UniqueIndexUiDeployVersion() string {
	return "ui_deploy_version"
}

func (m *AppletDeploy) ColID() *builder.Column {
	return AppletDeployTable.ColByFieldName(m.FieldID())
}

func (AppletDeploy) FieldID() string {
	return "ID"
}

func (m *AppletDeploy) ColAppletID() *builder.Column {
	return AppletDeployTable.ColByFieldName(m.FieldAppletID())
}

func (AppletDeploy) FieldAppletID() string {
	return "AppletID"
}

func (m *AppletDeploy) ColDeployID() *builder.Column {
	return AppletDeployTable.ColByFieldName(m.FieldDeployID())
}

func (AppletDeploy) FieldDeployID() string {
	return "DeployID"
}

func (m *AppletDeploy) ColLocation() *builder.Column {
	return AppletDeployTable.ColByFieldName(m.FieldLocation())
}

func (AppletDeploy) FieldLocation() string {
	return "Location"
}

func (m *AppletDeploy) ColVersion() *builder.Column {
	return AppletDeployTable.ColByFieldName(m.FieldVersion())
}

func (AppletDeploy) FieldVersion() string {
	return "Version"
}

func (m *AppletDeploy) ColWasmFile() *builder.Column {
	return AppletDeployTable.ColByFieldName(m.FieldWasmFile())
}

func (AppletDeploy) FieldWasmFile() string {
	return "WasmFile"
}

func (m *AppletDeploy) ColAbiName() *builder.Column {
	return AppletDeployTable.ColByFieldName(m.FieldAbiName())
}

func (AppletDeploy) FieldAbiName() string {
	return "AbiName"
}

func (m *AppletDeploy) ColAbiFile() *builder.Column {
	return AppletDeployTable.ColByFieldName(m.FieldAbiFile())
}

func (AppletDeploy) FieldAbiFile() string {
	return "AbiFile"
}

func (m *AppletDeploy) ColCreatedAt() *builder.Column {
	return AppletDeployTable.ColByFieldName(m.FieldCreatedAt())
}

func (AppletDeploy) FieldCreatedAt() string {
	return "CreatedAt"
}

func (m *AppletDeploy) ColUpdatedAt() *builder.Column {
	return AppletDeployTable.ColByFieldName(m.FieldUpdatedAt())
}

func (AppletDeploy) FieldUpdatedAt() string {
	return "UpdatedAt"
}

func (m *AppletDeploy) CondByValue(db sqlx.DBExecutor) builder.SqlCondition {
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

func (m *AppletDeploy) Create(db sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt.Set(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt.Set(time.Now())
	}

	_, err := db.Exec(sqlx.InsertToDB(db, m, nil))
	return err
}

func (m *AppletDeploy) List(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) ([]AppletDeploy, error) {
	var (
		tbl = db.T(m)
		lst = make([]AppletDeploy, 0)
	)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("AppletDeploy.List")}, adds...)
	err := db.QueryAndScan(builder.Select(nil).From(tbl, adds...), &lst)
	return lst, err
}

func (m *AppletDeploy) Count(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) (cnt int64, err error) {
	tbl := db.T(m)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("AppletDeploy.List")}, adds...)
	err = db.QueryAndScan(builder.Select(builder.Count()).From(tbl, adds...), &cnt)
	return
}

func (m *AppletDeploy) FetchByID(db sqlx.DBExecutor) error {
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
				builder.Comment("AppletDeploy.FetchByID"),
			),
		m,
	)
	return err
}

func (m *AppletDeploy) FetchByDeployID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("DeployID").Eq(m.DeployID),
					),
				),
				builder.Comment("AppletDeploy.FetchByDeployID"),
			),
		m,
	)
	return err
}

func (m *AppletDeploy) FetchByAppletIDAndVersion(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("AppletID").Eq(m.AppletID),
						tbl.ColByFieldName("Version").Eq(m.Version),
					),
				),
				builder.Comment("AppletDeploy.FetchByAppletIDAndVersion"),
			),
		m,
	)
	return err
}

func (m *AppletDeploy) UpdateByIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

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
				builder.Comment("AppletDeploy.UpdateByIDWithFVs"),
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

func (m *AppletDeploy) UpdateByID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByIDWithFVs(db, fvs)
}

func (m *AppletDeploy) UpdateByDeployIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("DeployID").Eq(m.DeployID),
				),
				builder.Comment("AppletDeploy.UpdateByDeployIDWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByDeployID(db)
	}
	return nil
}

func (m *AppletDeploy) UpdateByDeployID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByDeployIDWithFVs(db, fvs)
}

func (m *AppletDeploy) UpdateByAppletIDAndVersionWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("AppletID").Eq(m.AppletID),
					tbl.ColByFieldName("Version").Eq(m.Version),
				),
				builder.Comment("AppletDeploy.UpdateByAppletIDAndVersionWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByAppletIDAndVersion(db)
	}
	return nil
}

func (m *AppletDeploy) UpdateByAppletIDAndVersion(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByAppletIDAndVersionWithFVs(db, fvs)
}

func (m *AppletDeploy) Delete(db sqlx.DBExecutor) error {
	_, err := db.Exec(
		builder.Delete().
			From(
				db.T(m),
				builder.Where(m.CondByValue(db)),
				builder.Comment("AppletDeploy.Delete"),
			),
	)
	return err
}

func (m *AppletDeploy) DeleteByID(db sqlx.DBExecutor) error {
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
				builder.Comment("AppletDeploy.DeleteByID"),
			),
	)
	return err
}

func (m *AppletDeploy) DeleteByDeployID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("DeployID").Eq(m.DeployID),
					),
				),
				builder.Comment("AppletDeploy.DeleteByDeployID"),
			),
	)
	return err
}

func (m *AppletDeploy) DeleteByAppletIDAndVersion(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("AppletID").Eq(m.AppletID),
						tbl.ColByFieldName("Version").Eq(m.Version),
					),
				),
				builder.Comment("AppletDeploy.DeleteByAppletIDAndVersion"),
			),
	)
	return err
}
