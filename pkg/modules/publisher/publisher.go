package publisher

import (
	"context"

	confid "github.com/iotexproject/Bumblebee/conf/id"
	"github.com/iotexproject/Bumblebee/conf/jwt"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"

	"github.com/iotexproject/w3bstream/pkg/errors/status"
	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/types"
)

type CreatePublisherReq struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

func CreatePublisher(ctx context.Context, projectID types.SFID, r *CreatePublisherReq) (*models.Publisher, error) {
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)
	j := jwt.MustConfFromContext(ctx)
	idg := confid.MustSFIDGeneratorFromContext(ctx)

	_, l = l.Start(ctx, "CreatePublisher")
	defer l.End()

	// TODO generate token, maybe use public key
	token, err := j.GenerateTokenByPayload(projectID)
	if err != nil {
		l.Error(err)
		return nil, status.InternalServerError.StatusErr().WithDesc(err.Error())
	}

	m := &models.Publisher{
		RelProject:    models.RelProject{ProjectID: projectID},
		RelPublisher:  models.RelPublisher{PublisherID: idg.MustGenSFID()},
		PublisherInfo: models.PublisherInfo{Name: r.Name, Key: r.Key, Token: token},
	}
	if err = m.Create(d); err != nil {
		l.Error(err)
		return nil, err
	}

	return m, nil
}

func GetPublisherByPublisherKey(ctx context.Context, publisherKey string) (*models.Publisher, error) {
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)
	m := &models.Publisher{
		PublisherInfo: models.PublisherInfo{Key: publisherKey},
	}

	_, l = l.Start(ctx, "GetPublisherByPublisherKey")

	if err := m.FetchByKey(d); err != nil {
		l.Error(err)
		return nil, status.CheckDatabaseError(err, "GetPublisherByPublisherKey")
	}

	return m, nil
}

type ListPublisherReq struct {
	projectID types.SFID
}

func (r *ListPublisherReq) SetCurrentProject(prjID types.SFID) { r.projectID = prjID }

func (r *ListPublisherReq) Condition() builder.SqlCondition {
	m := &models.Publisher{}
	return m.ColProjectID().Eq(r.projectID)
}

func (r *ListPublisherReq) Additions() builder.Additions { return nil }

type ListPublisherRsp struct {
	Total int64              `json:"total"`
	Data  []models.Publisher `json:"data"`
}

func ListPublisher(ctx context.Context, r *ListPublisherReq) (ret *ListPublisherRsp, err error) {
	l := types.MustLoggerFromContext(ctx)
	d := types.MustDBExecutorFromContext(ctx)

	m := &models.Publisher{}

	_, l = l.Start(ctx, "ListPublisher")
	defer l.End()

	ret = &ListPublisherRsp{}

	ret.Data, err = m.List(d, r.Condition(), r.Additions()...)
	if err != nil {
		l.Error(err)
		return nil, status.CheckDatabaseError(err, "ListPublisher")
	}

	ret.Total, err = m.Count(d, r.Condition())
	if err != nil {
		l.Error(err)
		return nil, status.CheckDatabaseError(err, "ListPublisherCount")
	}
	return ret, nil
}
