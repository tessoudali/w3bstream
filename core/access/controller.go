package access

type (
	// TODO: define primary system permissions
	Authentication interface {
		Type() uint64
	}
	Control interface {
		Check(sender string, nonce uint64, hash []byte, authentication []byte) (Authentication, error)
	}
)
