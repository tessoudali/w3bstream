package models

import (
	"github.com/iotexproject/Bumblebee/kit/sqlx"
)

var DB = sqlx.NewDatabase("demo").WithSchema("applet_management")
