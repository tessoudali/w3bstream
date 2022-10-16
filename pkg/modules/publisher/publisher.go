package publisher

import (
	"context"

	confid "github.com/iotexproject/Bumblebee/conf/id"
	"github.com/iotexproject/Bumblebee/conf/jwt"

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
	j := jwt.MustConfFromContext(ctx)
	idg := confid.MustSFIDGeneratorFromContext(ctx)

	// TODO generate token, maybe use public key
	token, err := j.GenerateTokenByPayload(projectID)
	if err != nil {
		return nil, status.InternalServerError.StatusErr().WithDesc(err.Error())
	}

	m := &models.Publisher{
		RelProject:    models.RelProject{ProjectID: projectID},
		RelPublisher:  models.RelPublisher{PublisherID: idg.MustGenSFID()},
		PublisherInfo: models.PublisherInfo{Name: r.Name, Key: r.Key, Token: token},
	}
	if err = m.Create(d); err != nil {
		return nil, err
	}

	return m, nil
}

func GetPublisherByPublisherKey(ctx context.Context, publisherKey string) (*models.Publisher, error) {
	d := types.MustDBExecutorFromContext(ctx)
	m := &models.Publisher{
		PublisherInfo: models.PublisherInfo{Key: publisherKey},
	}

	if err := m.FetchByKey(d); err != nil {
		return nil, status.CheckDatabaseError(err, "GetPublisherByPublisherKey")
	}

	return m, nil
}
