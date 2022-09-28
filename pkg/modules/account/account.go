package account

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"
	"github.com/iotexproject/w3bstream/pkg/depends/util"
	"github.com/iotexproject/w3bstream/pkg/enums"
	"github.com/iotexproject/w3bstream/pkg/errors/status"
	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/types"
)

type CreateAccountByUsernameReq struct {
	Username string `json:"username"`
}

func CreateAccount(ctx context.Context, r *CreateAccountByUsernameReq) (*models.Account, error) {
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)

	accountID := uuid.New().String()
	m := &models.Account{
		RelAccount: models.RelAccount{AccountID: accountID},
		AccountInfo: models.AccountInfo{
			Username:     r.Username,
			IdentityType: enums.ACCOUNT_IDENTITY_TYPE__USERNAME,
			State:        enums.ACCOUNT_STATE__ENABLED,
			Password: models.AccountPassword{
				Type: enums.PASSWORD_TYPE__LOGIN,
				Password: hashOfAccountPassword(
					accountID,
					string(util.GenRandomPassword(8, 3)),
				),
			},
		},
	}

	l.Start(ctx, "CreateAccountByUsername")
	defer l.End()

	if err := m.Create(d); err != nil {
		l.Error(err)
		return nil, status.CheckDatabaseError(err, "CreateAccount")
	}
	return m, nil
}

func hashOfAccountPassword(accountID string, password string) string {
	return string(toMD5(toMD5([]byte(fmt.Sprintf("%s-%s", accountID, password)))))
}

func toMD5(src []byte) []byte {
	m := md5.New()
	_, _ = m.Write(src)
	cipherStr := m.Sum(nil)
	return []byte(hex.EncodeToString(cipherStr))
}

func ValidateAccountByLogin(ctx context.Context, username, password string) (*models.Account, error) {
	d := types.MustDBExecutorFromContext(ctx)

	m := &models.Account{}
	m.Username = username

	if err := m.FetchByUsername(d); err != nil {
		return nil, status.CheckDatabaseError(err, "FetchByUsername")
	}
	if m.Password.Password == hashOfAccountPassword(m.AccountID, password) {
		return m, nil
	}
	return nil, status.Unauthorized.StatusErr().WithDesc("invalid password")
}

func GetAccountByAccountID(ctx context.Context, accountID string) (*models.Account, error) {
	d := types.MustDBExecutorFromContext(ctx)
	m := &models.Account{RelAccount: models.RelAccount{AccountID: accountID}}
	err := m.FetchByAccountID(d)
	if err != nil {
		return nil, status.CheckDatabaseError(err, "FetchByAccountID")
	}
	return m, err
}

func CreateAdminIfNotExist(ctx context.Context) (string, error) {
	d := types.MustDBExecutorFromContext(ctx)

	accountID := uuid.New().String()
	password := string(util.GenRandomPassword(8, 3))
	m := &models.Account{
		RelAccount: models.RelAccount{AccountID: accountID},
		AccountInfo: models.AccountInfo{
			Username:     "admin",
			IdentityType: enums.ACCOUNT_IDENTITY_TYPE__BUILTIN,
			State:        enums.ACCOUNT_STATE__ENABLED,
			Password: models.AccountPassword{
				Type: enums.PASSWORD_TYPE__LOGIN,
				Password: hashOfAccountPassword(
					accountID,
					password,
				),
				Scope: "admin",
				Desc:  "builtin password",
			},
		},
	}

	results := make([]models.Account, 0)
	err := d.QueryAndScan(builder.Select(nil).
		From(
			d.T(m),
			builder.Where(
				builder.And(
					m.ColUsername().Eq("admin"),
					m.ColIdentityType().Eq(enums.ACCOUNT_IDENTITY_TYPE__BUILTIN),
				),
			),
		), &results)
	if err != nil {
		return "", status.CheckDatabaseError(err, "FetchAdminAccount")
	}
	if len(results) > 0 {
		return results[0].Password.Password, nil
	}
	if err = m.Create(d); err != nil {
		return "", err
	}
	return password, nil
}
