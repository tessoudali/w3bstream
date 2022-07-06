package config

import (
	"log"

	"github.com/go-courier/envconf"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/migration"
	"github.com/iotexproject/w3bstream/base/confpostgres"
	"github.com/iotexproject/w3bstream/base/examples/modelgen/pkg/models"
)

// global vars

var (
	db *confpostgres.PostgresEndpoint

	databaseURI = "postgres://pguser:pgpassword@127.0.0.1:5432/sqlxdemo?sslmode=disable"
)

func init() {
	db = &confpostgres.PostgresEndpoint{Database: models.DB}
	db.SetDefaults()
	ep, err := envconf.ParseEndpoint(databaseURI)
	if err != nil {
		panic(err)
	}
	db.Endpoint = *ep
	db.Init()
	log.Print("database initialized")

	if err = Migrate(db); err != nil {
		log.Panic(err)
	}
	log.Print("models migrated")
}

func Migrate(db sqlx.DBExecutor) error {
	if err := migration.Migrate(db, nil); err != nil {
		panic(err)
	}
	return nil
}

func DB() sqlx.DBExecutor { return db }
