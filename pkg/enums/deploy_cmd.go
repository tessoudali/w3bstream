package enums

//go:generate toolkit gen enum DeployCmd
type DeployCmd uint8

const (
	DEPLOY_CMD_UNKNOWN DeployCmd = iota
	DEPLOY_CMD__CREATE
	DEPLOY_CMD__START
	DEPLOY_CMD__STOP
	DEPLOY_CMD__REMOVE
)
