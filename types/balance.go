package types

type BalanceAtHeightKey struct {
	Address string
	Height  int64
}

func NewBalanceAtHeightKey(address string, height int64) *BalanceAtHeightKey {
	return &BalanceAtHeightKey{Address: address, Height: height}
}
