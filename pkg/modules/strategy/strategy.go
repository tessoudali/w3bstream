package strategy

import (
	"context"

	confid "github.com/iotexproject/Bumblebee/conf/id"
	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/pkg/enums"
	"github.com/iotexproject/w3bstream/pkg/errors/status"
	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/types"
)

type InstanceHandler struct {
	AppletID   types.SFID
	InstanceID types.SFID
	Handler    string
}

func FindStrategyInstances(ctx context.Context, prjName string, eventType enums.EventType) ([]*InstanceHandler, error) {
	l := types.MustLoggerFromContext(ctx)
	d := types.MustDBExecutorFromContext(ctx)

	_, l = l.Start(ctx, "FindStrategyInstances")
	defer l.End()

	l = l.WithValues("project", prjName, "event_type", eventType.String())

	mProject := &models.Project{ProjectInfo: models.ProjectInfo{Name: prjName}}

	if err := mProject.FetchByName(d); err != nil {
		l.Error(err)
		return nil, status.CheckDatabaseError(err, "FetchProjectByName")
	}

	mStrategy := &models.Strategy{}

	strategies, err := mStrategy.List(d,
		builder.And(
			mStrategy.ColProjectID().Eq(mProject.ProjectID),
			builder.Or(
				mStrategy.ColEventType().Eq(eventType),
				mStrategy.ColEventType().Eq(enums.EVENT_TYPE__ANY),
			),
		),
	)
	if err != nil {
		l.Error(err)
		return nil, status.CheckDatabaseError(err, "ListStrategy")
	}

	if len(strategies) == 0 {
		l.Warn(errors.New("strategy not found"))
		return nil, status.NotFound.StatusErr().WithDesc("not found strategy")
	}
	strategiesMap := make(map[types.SFID]*models.Strategy)
	for i := range strategies {
		strategiesMap[strategies[i].AppletID] = &strategies[i]
	}

	appletIDs := make(types.SFIDs, 0, len(strategies))

	for i := range strategies {
		appletIDs = append(appletIDs, strategies[i].AppletID)
	}

	mInstance := &models.Instance{}

	instances, err := mInstance.List(d,
		builder.And(
			mInstance.ColAppletID().In(appletIDs),
			mInstance.ColState().Eq(enums.INSTANCE_STATE__STARTED),
		),
	)
	if err != nil {
		l.Error(err)
		return nil, status.CheckDatabaseError(err, "ListInstances")
	}

	if len(instances) == 0 {
		l.Warn(errors.New("instance not found"))
		return nil, status.NotFound.StatusErr().WithDesc("not found instance")
	}

	handlers := make([]*InstanceHandler, 0)

	for _, instance := range instances {
		handlers = append(handlers, &InstanceHandler{
			AppletID:   instance.AppletID,
			InstanceID: instance.InstanceID,
			Handler:    strategiesMap[instance.AppletID].Handler,
		})
	}
	return handlers, nil
}

type CreateStrategyBatchReq struct {
	Strategies []CreateStrategyReq `json:"strategies"`
}

type CreateStrategyReq struct {
	models.RelApplet
	models.StrategyInfo
}

func CreateStrategy(ctx context.Context, projectID types.SFID, r *CreateStrategyBatchReq) (err error) {
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)
	idg := confid.MustSFIDGeneratorFromContext(ctx)

	_, l = l.Start(ctx, "CreateStrategy")
	defer l.End()

	//m := &models.Strategy{}
	err = sqlx.NewTasks(d).With(
		func(db sqlx.DBExecutor) error {
			for i := range r.Strategies {
				if err := (&models.Strategy{
					RelStrategy:  models.RelStrategy{StrategyID: idg.MustGenSFID()},
					RelProject:   models.RelProject{ProjectID: projectID},
					RelApplet:    models.RelApplet{AppletID: r.Strategies[i].AppletID},
					StrategyInfo: models.StrategyInfo{EventType: r.Strategies[i].EventType, Handler: r.Strategies[i].Handler},
				}).Create(db); err != nil {
					return err
				}
			}
			return nil
		},
	).Do()

	if err != nil {
		l.Error(err)
		return status.CheckDatabaseError(err, "CreateStrategy")
	}

	return
}

func UpdateStrategy(ctx context.Context, strategyID types.SFID, r *CreateStrategyReq) (err error) {
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)
	m := models.Strategy{RelStrategy: models.RelStrategy{StrategyID: strategyID}}

	_, l = l.Start(ctx, "UpdateStrategy")
	defer l.End()

	err = sqlx.NewTasks(d).With(
		func(db sqlx.DBExecutor) error {
			return m.FetchByStrategyID(d)
		},
		func(db sqlx.DBExecutor) error {
			m.RelApplet = r.RelApplet
			m.StrategyInfo.EventType = r.EventType
			m.StrategyInfo.Handler = r.Handler
			return m.UpdateByStrategyID(d)
		},
	).Do()

	if err != nil {
		l.Error(err)
		return status.CheckDatabaseError(err, "UpdateStrategy")
	}

	return
}
