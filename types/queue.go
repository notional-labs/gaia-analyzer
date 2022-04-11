// This example demonstrates a priority queue built using the heap interface.
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TimeTx struct {
	Tx *sdk.TxResponse

	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type TxTimeQueue []*TimeTx

func (pq TxTimeQueue) Len() int { return len(pq) }

func (pq TxTimeQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Tx.Height > pq[j].Tx.Height
}

func (pq TxTimeQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *TxTimeQueue) Push(item *TimeTx) {
	n := len(*pq)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *TxTimeQueue) Pop() *TimeTx {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// // update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *SimpleTx, value string, priority int) {
// 	item.Tx =
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }
