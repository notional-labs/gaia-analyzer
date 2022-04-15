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
					if strings.Contains(string(attribute.Value), "uatom") {
						return true
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
	heap.Push(&data.TxQueue, txItem)
}
