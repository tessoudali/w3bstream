package jwt

import (
	"context"
	"github.com/machinefi/w3bstream/pkg/depends/x/contextx"
	"github.com/machinefi/w3bstream/pkg/depends/x/misc/must"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/machinefi/w3bstream/pkg/depends/base/types"
)

type PublisherJwt struct {
	Issuer  string         `env:""`
	ExpIn   types.Duration `env:""`
	SignKey string         `env:""`
	// Method  SigningMethod  `env:""`
}

func (c *PublisherJwt) SetDefault() {}

func (c *PublisherJwt) Init() {}

func (c *PublisherJwt) GenerateTokenByPayload(payload interface{}) (string, error) {
	claim := &Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(c.ExpIn.Duration())},
			Issuer:    c.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(c.SignKey))
}

func (c *PublisherJwt) ParseToken(v string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(
		v,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(c.SignKey), nil
		},
	)
	if err != nil {
		return nil, InvalidToken.StatusErr().WithDesc(err.Error())
	}
	if t == nil {
		return nil, InvalidToken
	}
	claim, ok := t.Claims.(*Claims)
	if !ok || !t.Valid {
		return nil, InvalidClaim
	}
	return claim, nil
}

type publisherAuth struct{}

func WithPublisherAuthContext(jwt *PublisherJwt) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, publisherAuth{}, jwt)
	}
}

func PublisherAuthFromContext(ctx context.Context) (*PublisherJwt, bool) {
	j, ok := ctx.Value(publisherAuth{}).(*PublisherJwt)
	return j, ok
}

func MustPublisherAuthFromContext(ctx context.Context) *PublisherJwt {
	j, ok := ctx.Value(publisherAuth{}).(*PublisherJwt)
	must.BeTrue(ok)
	return j
}
