package plog

import (
	"context"
	"fmt"
	"time"

	base "github.com/machinefi/w3bstream/pkg/depends/base/types"
	confid "github.com/machinefi/w3bstream/pkg/depends/conf/id"
	"github.com/machinefi/w3bstream/pkg/depends/kit/mq"
	"github.com/machinefi/w3bstream/pkg/models"
	"github.com/machinefi/w3bstream/pkg/types"
)

func NewTask(ctx context.Context, logLevel, msg string) *Task {
	return &Task{
		pLog: &models.RuntimeLog{
			RelRuntimeLog: models.RelRuntimeLog{RuntimeLogID: confid.MustSFIDGeneratorFromContext(ctx).MustGenSFID()},
			RuntimeLogInfo: models.RuntimeLogInfo{
				ProjectName: types.MustProjectFromContext(ctx).ProjectName.Name,
				AppletName:  types.MustAppletFromContext(ctx).Name,
				SourceName:  types.MustLogSourceFromContext(ctx),
				InstanceID:  types.MustInstanceFromContext(ctx).InstanceID,
				Level:       logLevel,
				LogTime:     base.AsTimestamp(time.Now()),
				Msg:         msg,
			},
		},
	}
}

type Task struct {
	pLog *models.RuntimeLog
	mq.TaskState
}

var _ mq.Task = (*Task)(nil)

func (t *Task) Subject() string {
	return "DbLogStoring"
}

func (t *Task) Arg() interface{} {
	return t.pLog
}

func (t *Task) ID() string {
	return fmt.Sprintf("%s::%s", t.Subject(), t.pLog.RuntimeLogID)
}
