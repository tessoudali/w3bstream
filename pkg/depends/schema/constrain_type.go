package schema

import "fmt"

//go:generate toolkit gen enum ConstrainType
type ConstrainType int8

const (
	CONSTRAIN_TYPE_UNKNOWN ConstrainType = iota
	CONSTRAIN_TYPE__AUTOINCREMENT
	CONSTRAIN_TYPE__NOT_NULL
	CONSTRAIN_TYPE__DEFAULT
)

func (v ConstrainType) Tag(driver string, value interface{}) string {
	switch v {
	case CONSTRAIN_TYPE__AUTOINCREMENT:
		return "autoincrement"
	case CONSTRAIN_TYPE__DEFAULT:
		return "default=" + fmt.Sprintf(`'%v'`, value)
	default:
		return ""
	}
}
