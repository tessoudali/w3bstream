package event

import "errors"

var (
	errParamIllegal        = errors.New("param illegal")
	errInternalSystemError = errors.New("internal system error")
)
