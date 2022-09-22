package wasm

type ImportsHandler interface {
	GetDB(keyAddr, keySize, valAddr, valSize uint32) (code int32)
	SetDB()
	GetData()
	SetData()
	Log(level uint32)
}
