package plog

import (
	"context"
	"time"

	confid "github.com/machinefi/w3bstream/pkg/depends/conf/id"
	"github.com/machinefi/w3bstream/pkg/depends/kit/sqlx"
	"github.com/machinefi/w3bstream/pkg/models"
	"github.com/machinefi/w3bstream/pkg/types"
)

type PersistenceLog struct {
	DB  sqlx.DBExecutor
	Ctx context.Context
	Src string
}

func (p *PersistenceLog) PersLog(logLevel, msg string) {
	idg := confid.MustSFIDGeneratorFromContext(p.Ctx)
	mLog := &models.RuntimeLog{
		RelRuntimeLog: models.RelRuntimeLog{RuntimeLogID: idg.MustGenSFID()},
		RuntimeLogInfo: models.RuntimeLogInfo{
			ProjectName: types.MustProjectFromContext(p.Ctx).ProjectName.Name,
			AppletName:  types.MustAppletFromContext(p.Ctx).Name,
			SourceName:  p.Src,
			InstanceID:  types.MustInstanceFromContext(p.Ctx).InstanceID,
			Level:       logLevel,
			Msg:         msg,
		},
	}
	// TODO fix time delay
	mLog.LogTime.Set(time.Now())
	go p.dbPers(mLog)
}

func (p *PersistenceLog) dbPers(mLog *models.RuntimeLog) {
	l := types.MustLoggerFromContext(p.Ctx)
	if err := mLog.Create(p.DB); err != nil {
		l.Error(err)
	}
}
