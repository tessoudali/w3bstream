package wasm

import "github.com/iotexproject/w3bstream/pkg/enums"

var NameVersion = "w3bstream@v0.0.1"

// ResultStatusCode wasm call result code
type ResultStatusCode int32

const (
	ResultStatusCode_OK ResultStatusCode = iota
	// TODO result status define here
	ResultStatusCode_Failed = -1 // reserved for wasm invoke failed
)

type InstanceState = enums.InstanceState
