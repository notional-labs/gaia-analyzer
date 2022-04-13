package handler

import (
	"github.com/notional-labs/gaia-analyzer/types"
)

var (
	// To check if this account has been tracked. A Tracked account is an account that receives atom from the root account
	IsTrackedAccount map[string]bool = map[string]bool{}
	// Amount of atom a given account has
	AtomBalance map[string]float64 = map[string]float64{}
	// Amount of tracked atom a given account has
	TrackedAtomBalance map[string]float64 = map[string]float64{}
	// A priority of txs with priority indicator being the tx height
	TxQueue types.LowestHeightFirstOutTxQueue = types.LowestHeightFirstOutTxQueue{}
)

// // This example inserts several ints into an IntHeap, checks the minimum,
// // and removes them in order of priority.
// func init() {
// 	heap.Init(&TxQueue)
// }
