package tx

import (
	"container/heap"
	"strings"

	"github.com/notional-labs/gaia-analyzer/data"
	"github.com/notional-labs/gaia-analyzer/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

// check if this bank send tx is bank send atom
func IsBankSendUatomTx(events *[]abcitypes.Event) bool {
	for _, event := range *events {
		if event.Type == "transfer" {
			for _, attribute := range event.Attributes {
				if string(attribute.Key) == "amount" {
					// check if denom of coin sent is atom
					amount := string(attribute.Value)
					if strings.Contains(amount, data.TrackedDenom) {
						if len(amount)-len(data.TrackedDenom) > 5 {
							return true
						}
					}
				}
			}
		}
	}

	return false
}

func PushToTxQueue(tx *abcitypes.TxResult) {
	txItem := types.TxItem{
		Height: tx.Height,
		Events: &tx.Result.Events,
	}
	heap.Push(&data.TxQueue, &txItem)
}
