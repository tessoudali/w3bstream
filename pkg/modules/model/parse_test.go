package model_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/iotexproject/w3bstream/pkg/modules/model"
	. "github.com/onsi/gomega"
)

var c = &model.Config{}

func init() {
	_, pwd, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(pwd), "demo")
	file := filepath.Join(root, "schema.yaml")

	var err error
	c, err = model.LoadConfigFrom(file)
	if err != nil {
		panic(err)
	}
}

func TestConfig_Schema(t *testing.T) {
	NewWithT(t).Expect(c.Schema()).To(Equal("demo_0.0.1"))
}

func ExampleConfig_SnippetDefs() {
	snippets := c.SnippetDefs()
	for tbl, s := range snippets {
		fmt.Println(tbl)
		fmt.Println(string(s.Bytes()))
	}
	// Output:
	// t_demo
	// // @def primary                       ID
	// // @def index        I_nickname/BTREE Name
	// // @def index        I_username       Username
	// // @def unique_index UI_name          Name
	// // @def unique_index UI_id_org        ID OrgID
}

func ExampleConfig_SnippetStruct() {
	snippets := c.SnippetStruct()
	for tbl, s := range snippets {
		fmt.Println(tbl)
		fmt.Println(string(s.Bytes()))
	}
	// Output:
	// Demo
	// // Demo demo table
	// type Demo struct {
	// ID uint64 `db:"f_id,autoincrement"`
	// Name string `db:"f_name,default=''"`
	// Nickname string `db:"f_nickname,default=''"`
	// Username string `db:"f_username,default=''"`
	// Gender int `db:"f_gender,default='0'"`
	// Boolean bool `db:"f_boolean,default=false"`
	// OrgID uint64 `db:"f_org_id"`
	// CreatedAt types.Timestamp `db:"f_created_at,default='0'"`
	// UpdatedAt types.Timestamp `db:"f_updated_at,default='0'"`
	// DeletedAt types.Timestamp `db:"f_deleted_at,default='0'"`
	// }
}
