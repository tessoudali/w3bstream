package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/iotexproject/w3bstream/base/examples/modelgen/cmd/config"
	"github.com/iotexproject/w3bstream/base/examples/modelgen/pkg/models"

	"github.com/google/uuid"
)

func main() {
	db := config.DB()
	id := uuid.New()

	rcd := models.Model{
		RefModel: models.RefModel{ModelID: uint64(id.ID())},
		BaseModel: models.BaseModel{
			FieldString: "field_string",
			FieldJsonContent: &models.EmbedModel{
				SomeFieldInt:   100,
				SomeFieldFloat: 100.11,
			},
			UnionIndexField1: int64(id.ID()),
			UnionIndexField2: id.String(),
		},
		OperationTimes: models.OperationTimes{},
	}

	// create
	if err := rcd.Create(db); err != nil {
		log.Panic("create record: ", err)
	}
	bytes, _ := json.MarshalIndent(rcd, "", "  ")
	fmt.Println(string(bytes))
	log.Printf("record created: %v\n", rcd.ModelID)

	// update
	rcd.FieldJsonContent.SomeFieldInt = 101
	rcd.FieldJsonContent.SomeFieldFloat = 101.111
	if err := rcd.UpdateByModelIDWithStruct(db); err != nil {
		log.Panic("update record: ", err)
	}
	bytes, _ = json.MarshalIndent(rcd, "", "  ")
	fmt.Println(string(bytes))
	log.Printf("record created: %v\n", rcd.ModelID)

	// delete
	if err := rcd.DeleteByModelID(db); err != nil {
		log.Panic(err)
	}
	log.Printf("record deleted: %v\n", rcd.ModelID)
}
