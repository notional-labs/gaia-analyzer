package handler

import (
	"container/heap"

	"github.com/notional-labs/gaia-analyzer/types"
)

var (
	IsTrackedAccount map[string]bool

	AtomBalances map[string]float64

	TrackedAtomBalances map[string]float64

	TxQueue types.TxTimeQueue = make(types.TxTimeQueue, 100)
)

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func init() {
	heap.Init(&TxQueue)
}
