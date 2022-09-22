package wasm

type Data struct {
	offset, size uint32
}

func (d *Data) ToVMString() {}
