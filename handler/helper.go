package handler

import (
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// parse events from bank send tx, return sender, recipient and amount sent
func ParseBankSendTxEvent(tx *sdk.TxResponse) (string, string, float64) {
	var amount float64
	var sender, recipient string

	for _, event := range tx.Events {
		if event.Type == "transfer" {
			for _, attribute := range event.Attributes {
				switch string(attribute.Key) {
				case "sender":
					sender = string(attribute.Value)
				case "recipient":
					recipient = string(attribute.Value)
				case "amount":
					// get amount from stringify sdk.Coin
					amount, _ = strconv.ParseFloat(strings.Trim(string(attribute.Value), "uatom"), 64)
				}
			}
		}
	}

	return sender, recipient, amount
}

// check if this bank send tx is bank send atom
func IsBankSendAtomTx(tx *sdk.TxResponse) bool {

	for _, event := range tx.Events {
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
