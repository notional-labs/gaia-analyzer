package tx

import (
	"container/heap"
	"strings"

	"github.com/notional-labs/gaia-analyzer/data"
	"github.com/notional-labs/gaia-analyzer/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func IsUndelegateTx()

// check if this bank send tx is bank send atom
func IsBankSendUatomTx(events *[]abcitypes.Event) bool {
	for _, event := range *events {
		if event.Type == "coin_spent" {
			for _, attribute := range event.Attributes {
				if string(attribute.Key) == "amount" {
					// check if denom of coin sent is atom
					if strings.Contains(string(attribute.Value), data.Denom) {
						return true
					}
				}
			}
		}
	}

	return false
}

func PushToTrackedTxQueue(tx *abcitypes.TxResult) {
	txItem := types.TxItem{
		Height: tx.Height,
		Events: &tx.Result.Events,
	}
	heap.Push(&data.TrackedTxQueue, &txItem)
}
