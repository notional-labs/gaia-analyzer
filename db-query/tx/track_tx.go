package tx

import (
	"fmt"
	"strconv"
)

// query txs that spend coin from 1 account. push those txs to tracked tx queue
func TrackTxsSpendingCoinsFromAccount(spender string, fromHeight int64) {
	// events use for query
	spendingEvent := fmt.Sprintf("%s='%s'", "coin_spent.spender", spender)

	heightEvent := fmt.Sprintf("%s>%s", "tx.height", strconv.FormatInt(fromHeight, 10))
	tmEvents := []string{spendingEvent, heightEvent}

	for _, r := range QueryTxs(tmEvents) {
		if IsBankSendUatomTx(&r.Result.Events) {
			PushToTrackedTxQueue(r)
		}
	}
}

// query txs that send atom to a given address since a given height. Push those txs to global tx queue
func TrackTxsSendingCoinsToAccount(receiver string, fromHeight int64) {
	// events use for query
	receivingEvent := fmt.Sprintf("%s='%s'", "coin_received.receiver", receiver)
	notSpendingEvent := fmt.Sprintf("%s!='%s'", "coin_spent.spender", receiver)
	heightEvent := fmt.Sprintf("%s>%s", "tx.height", strconv.FormatInt(fromHeight, 10))
	tmEvents := []string{receivingEvent, heightEvent, notSpendingEvent}

	for _, r := range QueryTxs(tmEvents) {
		if IsBankSendUatomTx(&r.Result.Events) {
			PushToTrackedTxQueue(r)
		}
	}
}
