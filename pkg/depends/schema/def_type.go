package schema

//go:generate toolkit gen enum DefType
type DefType int8

const (
	DEF_TYPE_UNKNOWN DefType = iota
	DEF_TYPE__PRIMARY
	DEF_TYPE__INDEX
	DEF_TYPE__UNIQUE_INDEX
)
