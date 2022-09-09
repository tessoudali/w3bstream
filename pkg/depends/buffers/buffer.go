package buffers

// TODO buffer pool for wasm vms

// Object wasm resource
type Object struct {
	data     []byte
	id       string
	size     int
	capacity int

}

type Pool interface {
	Get(size int) *Object
	Put(*Object)
}
