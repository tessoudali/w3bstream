package enums

//go:generate toolkit gen enum AccountState
type AccountState uint8

const (
	ACCOUNT_STATE_UNKNOWN AccountState = iota
	ACCOUNT_STATE__ENABLED
	ACCOUNT_STATE__DISABLED
)
