package status

import (
	"net/http"
)

//go:generate toolkit gen status Error
type Error int

func (Error) ServiceCode() int {
	return 999 * 1e3
}

const (
	// InternalServerError 内部错误
	InternalServerError Error = http.StatusInternalServerError*1e6 + iota + 1
	UploadFileFailed
	ExtractFileFailed
	LoadVMFailed
)

const (
	// @errTalk Unauthorized
	Unauthorized Error = http.StatusUnauthorized*1e6 + iota + 1
)

const (
	// Conflict conflict error
	Conflict Error = http.StatusConflict*1e6 + iota + 1
)

const (
	// BadRequest
	BadRequest Error = http.StatusBadRequest*1e6 + iota + 1
	MD5ChecksumFailed
)

const (
	// NotFound
	NotFound Error = http.StatusNotFound*1e6 + iota + 1
)
