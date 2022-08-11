package global

import (
	"context"
	"os"

	"github.com/iotexproject/Bumblebee/base/consts"
	confapp "github.com/iotexproject/Bumblebee/conf/app"
	confhttp "github.com/iotexproject/Bumblebee/conf/http"
	confjwt "github.com/iotexproject/Bumblebee/conf/jwt"
	conflog "github.com/iotexproject/Bumblebee/conf/log"
	confmqtt "github.com/iotexproject/Bumblebee/conf/mqtt"
	confpostgres "github.com/iotexproject/Bumblebee/conf/postgres"
	"github.com/iotexproject/Bumblebee/kit/kit"
	"github.com/iotexproject/Bumblebee/kit/sqlx/migration"
	"github.com/iotexproject/Bumblebee/x/contextx"

	"github.com/iotexproject/w3bstream/pkg/models"
)

// global vars

var (
	postgres = &confpostgres.Endpoint{Database: models.DB}
	mqtt     = &confmqtt.Broker{}
	server   = &confhttp.Server{}
	jwt      = &confjwt.Jwt{}
	logger   = &conflog.Log{Name: "srv-demo"}
	std      = conflog.Std()

	App *confapp.Ctx
)

func init() {
	name := os.Getenv(consts.EnvProjectName)
	if name == "" {
		name = "srv-applet-mgr"
	}
	App = confapp.New(
		confapp.WithName(name),
		confapp.WithRoot(".."),
		confapp.WithVersion("0.0.1"),
		confapp.WithLogger(conflog.Std()),
	)
	App.Conf(postgres, server, jwt, logger)

	confhttp.RegisterCheckerBy(postgres, mqtt, server)
	std.(conflog.LevelSetter).SetLevel(conflog.InfoLevel)
}

var WithContext = contextx.WithContextCompose(
	WithDatabase(postgres),
	WithLogger(conflog.Std()),
)

func Server() kit.Transport {
	return server.WithContextInjector(WithContext)
}

func Migrate() {
	ctx, log := conflog.StdContext(context.Background())

	log.Start(ctx, "Migrate")
	defer log.End()
	if err := migration.Migrate(postgres.WithContext(ctx), nil); err != nil {
		log.Panic(err)
	}
}
