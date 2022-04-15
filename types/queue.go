// This example demonstrates a priority queue built using the heap interface.
package types

import (
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

// tx item to put into queue
type TxItem struct {
	Height int64

	Events *[]abcitypes.Event

	index int // The index of the item in the heap.
}

// A priority of txs with priority indicator being the tx height
type LowestHeightFirstOutTxQueue []*TxItem

func (q LowestHeightFirstOutTxQueue) Len() int { return len(q) }

func (q LowestHeightFirstOutTxQueue) Less(i, j int) bool {
	return q[i].Height < q[j].Height
}

func (q LowestHeightFirstOutTxQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *LowestHeightFirstOutTxQueue) Push(x any) {
	n := len(*q)
	item := x.(*TxItem)
	item.index = n
	*q = append(*q, item)
}

// pop return the tx with lowest height
func (q *LowestHeightFirstOutTxQueue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}
