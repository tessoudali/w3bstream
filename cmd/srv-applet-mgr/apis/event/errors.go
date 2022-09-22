package event

import "errors"

// TODO move status define to pkg/errors/status

var (
	errParamIllegal        = errors.New("param illegal")
	errInternalSystemError = errors.New("internal system error")
)
