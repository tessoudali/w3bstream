package storage

type KVStore interface {
	Insert([]byte, []byte) error
	Read([]byte) ([]byte, error)
}
