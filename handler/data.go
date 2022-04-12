package handler

import (
	"container/heap"

	"github.com/notional-labs/gaia-analyzer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	TxsOf [][]byte

	Balances map[types.BalanceAtHeightKey]uint64

	TrackedCoinBalances map[types.BalanceAtHeightKey]uint64

	AddedBlocks map[int64]bool

	TxsAtBlock map[int64][]*sdk.TxResponse

	BlockQueue *types.IntHeap = &types.IntHeap{}
)

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func init() {
	heap.Init(BlockQueue)
}
