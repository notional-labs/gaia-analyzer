package data

import "github.com/notional-labs/gaia-analyzer/types"

var (
	TxsOf [][]byte

	BalanceOf map[types.BalanceQuery]uint64

	BlendedBalance map[types.BalanceQuery]uint64

	TxTimeQueue types.TxTimeQueue
)
