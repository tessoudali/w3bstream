package web3

type Client interface {
	Tip() (uint64, error)
}
