package tx

import (
	"fmt"
	"strconv"
)

// query txs that send atom from a given address since a given height. Push those txs to global tx queue
func GetBankSendUatomFromAddress(sender string, fromHeight int64) {
	// events use for query
	senderEvent := fmt.Sprintf("%s='%s'", "message.sender", sender)
	bankSendEvent := fmt.Sprintf("%s='%s'", "message.action", "/cosmos.bank.v1beta1.MsgSend")
	heightEvent := fmt.Sprintf("%s>%s", "tx.height", strconv.FormatInt(fromHeight, 10))
	tmEvents := []string{senderEvent, bankSendEvent, heightEvent}

	for _, r := range QueryTxs(tmEvents) {
		fmt.Println(r.Height)
		if IsBankSendUatomTx(&r.Result.Events) {
			PushToTxQueue(r)
		}
	}
}

// query txs that send atom to a given address since a given height. Push those txs to global tx queue
func GetBankSendUatomToAddress(receiver string, fromHeight int64) {
	// events use for query
	senderEvent := fmt.Sprintf("%s='%s'", "coin_received.receiver", receiver)
	bankSendEvent := fmt.Sprintf("%s='%s'", "message.action", "/cosmos.bank.v1beta1.MsgSend")
	heightEvent := fmt.Sprintf("%s>%s", "tx.height", strconv.FormatInt(fromHeight, 10))
	tmEvents := []string{senderEvent, bankSendEvent, heightEvent}

	for _, r := range QueryTxs(tmEvents) {
		if IsBankSendUatomTx(&r.Result.Events) {
			PushToTxQueue(r)
		}
	}
}
