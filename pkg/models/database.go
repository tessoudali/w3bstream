package models

import (
	"github.com/iotexproject/Bumblebee/kit/sqlx"
)

var Demo = sqlx.NewDatabase("demo").WithSchema("applet_management")
