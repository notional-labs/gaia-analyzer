package tx

import (
	"fmt"
	"strconv"
)

// query txs that spend coin from 1 account. push those txs to tracked tx queue
func TrackTxsTransferingCoinsFromAccount(sender string, fromHeight int64) {
	// events use for query
	sendingEvent := fmt.Sprintf("%s='%s'", "transfer.sender", sender)

	heightEvent := fmt.Sprintf("%s>%s", "tx.height", strconv.FormatInt(fromHeight, 10))
	tmEvents := []string{sendingEvent, heightEvent}

	for _, r := range QueryTxs(tmEvents) {
		if IsBankSendUatomTx(&r.Result.Events) {
			PushToTrackedTxQueue(r)
		}
	}
}

// query txs that send atom to a given address since a given height. Push those txs to tracked tx queue
func TrackTxsTransferingCoinsToAccount(receiver string, fromHeight int64) {
	// events use for query
	receivingEvent := fmt.Sprintf("%s='%s'", "transfer.recipient", receiver)
	notSendingEvent := fmt.Sprintf("%s!='%s'", "transfer.sender", receiver)
	heightEvent := fmt.Sprintf("%s>%s", "tx.height", strconv.FormatInt(fromHeight, 10))
	tmEvents := []string{receivingEvent, heightEvent, notSendingEvent}

	for _, r := range QueryTxs(tmEvents) {
		if IsBankSendUatomTx(&r.Result.Events) {
			PushToTrackedTxQueue(r)
		}
	}
}

// query txs that
