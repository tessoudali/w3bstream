package account

import (
	"context"

	"github.com/iotexproject/Bumblebee/kit/httptransport/httpx"

	"github.com/iotexproject/w3bstream/pkg/modules/account"
)

type CreateAccount struct {
	httpx.MethodPost
	account.CreateAccountByUsernameReq `in:"body"`
}

func (r *CreateAccount) Output(ctx context.Context) (interface{}, error) {
	return account.CreateAccount(ctx, &r.CreateAccountByUsernameReq)
}
