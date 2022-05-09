package types

type Message struct {
	Type           uint8
	Sender         string
	Nonce          uint64 // prevent replay
	Data           []byte
	Authentication []byte
}

func (msg *Message) Hash() []byte {
	// TODO: generate message hash with Type, Sender, Nonce, Data
	return nil
}
