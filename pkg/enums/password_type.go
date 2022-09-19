package enums

//go:generate toolkit gen enum PasswordType
type PasswordType uint8

const (
	PASSWORD_TYPE_UNKNOWN PasswordType = iota
	PASSWORD_TYPE__LOGIN
	PASSWORD_TYPE__PERSONAL_TOKEN
)
