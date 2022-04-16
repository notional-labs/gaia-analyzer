// This example demonstrates a priority queue built using the heap interface.
package types

type CoinMovingEvent struct {
	From   string
	To     string
	Amount uint64
}

// tx item to put into queue
type EventItem struct {
	Height int64

	Events *[]*CoinMovingEvent

	index int // The index of the item in the heap.
}

// A priority of txs with priority indicator being the tx height
type LowestHeightFirstOutEventQueue []*EventItem

func (q LowestHeightFirstOutEventQueue) Len() int { return len(q) }

func (q LowestHeightFirstOutEventQueue) Less(i, j int) bool {
	return q[i].Height < q[j].Height
}

func (q LowestHeightFirstOutEventQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *LowestHeightFirstOutEventQueue) Push(x any) {
	n := len(*q)
	item := x.(*EventItem)
	item.index = n
	*q = append(*q, item)
}

// pop return the tx with lowest height
func (q *LowestHeightFirstOutEventQueue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}
