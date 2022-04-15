package handler

import (
	"container/heap"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/gaia-analyzer/data"
	appquery "github.com/notional-labs/gaia-analyzer/db-query/app"
	txquery "github.com/notional-labs/gaia-analyzer/db-query/tx"
	"github.com/notional-labs/gaia-analyzer/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

// cal tracked atom sent from a tracked account using blended, not FIFO
// A Tracked account is an account that receives atom from the root account
func calSentTrackedUatomAmount(atomBalance sdk.Int, trackedUatomBalance sdk.Int, sentUatomAmount sdk.Int) sdk.Int {
	return sentUatomAmount.Mul(trackedUatomBalance).Quo(atomBalance)
}

// updates atom balance to Balance
func updateUatomBalance(address string, height int64) sdk.Int {

	uatomBalance, err := appquery.GetUatomBalanceAtHeight(address, height)
	if err != nil {
		panic(err)
	}
	data.Balance[address] = uatomBalance

	return uatomBalance
}

// Apply bank send tx, update tracked atom balance after tx
func handle_tx(tx *types.TxItem) string {
	height := tx.Height
	sender, recipient, sentUatomAmount := ParseBankSendTxEvent(tx.Events)
	// cal amount of tracked atom sent using blended, not FIFO
	senderUatomBalance := updateUatomBalance(sender, height-1)
	sentTrackedUatomAmount := calSentTrackedUatomAmount(senderUatomBalance, data.TrackedBalance[sender], sentUatomAmount)
	// update tracked atom balance for recipient account and sender account
	fmt.Println(senderUatomBalance, data.TrackedBalance[sender], sentUatomAmount)
	fmt.Printf("%s send %s tracked coins (%s coins) to %s at height %d \n", sender, sentTrackedUatomAmount.String(), sentUatomAmount, recipient, height)
	data.TrackedBalance[recipient] = data.TrackedBalance[recipient].Add(sentTrackedUatomAmount)
	data.TrackedBalance[sender] = data.TrackedBalance[sender].Sub(sentTrackedUatomAmount)

	return sender
}

// tracked all atom from a given address; start tracking from a given height; let's call this address root address
func TrackCoinsFromAccount(rootAddress string, startHeight int64) {

	// get all bank send from root address and to root address, push to global tx queue
	txquery.TrackTxsSpendingCoinsFromAccount(rootAddress, startHeight)
	txquery.TrackTxsSendingCoinsToAccount(rootAddress, startHeight)

	data.IsTrackedAccount[rootAddress] = true

	// query chain to get root account atom balance at start height
	atomAmountInThisAccount, err := appquery.GetUatomBalanceAtHeight(rootAddress, startHeight)
	if err != nil {
		panic(err)
	}
	// update atom balance and tracked atom balance of root address
	// tracked atom balance = atom balance at start height since we have to track all atoms from root account balance at that height
	data.Balance[rootAddress] = atomAmountInThisAccount
	data.TrackedBalance[rootAddress] = atomAmountInThisAccount

	for {
		// if tx queue empty, stop
		if len(data.TrackedTxQueue) == 0 {
			break
		}

		// get next tx from tx queue
		// tx queue : priority queue of txs with priority indicator being the tx's height
		tx := heap.Pop(&data.TrackedTxQueue).(*types.TxItem)
		// apply handle tx, output sender of tx
		sender, recepient, _ := ParseBankSendTxEvent(tx.Events)

		TrackAccount(sender, tx.Height)
		TrackAccount(recepient, tx.Height)

		handle_tx(tx)
	}
	fmt.Println(data.TrackedBalance)
}

func TrackAccount(address string, height int64) {
	_, isTracked := data.IsTrackedAccount[address]
	// if this account is not tracked yet, this means this account has not received any tracked atom before this tx
	// query BankSend tx this account sent starting from this tx height and push to tx queue
	if !isTracked {
		data.IsTrackedAccount[address] = true
		data.TrackedBalance[address] = sdk.ZeroInt()
		txquery.TrackTxsSpendingCoinsFromAccount(address, height)
	}
}

// parse events from bank send tx, return sender, recipient and amount sent
func ParseBankSendTxEvent(events *[]abcitypes.Event) (string, string, sdk.Int) {
	var amount sdk.Int
	var sender, recipient string

	for _, event := range *events {
		if event.Type == "transfer" {
			for _, attribute := range event.Attributes {
				switch string(attribute.Key) {
				case "sender":
					sender = string(attribute.Value)
				case "recipient":
					recipient = string(attribute.Value)
				case "amount":
					// get amount from stringify sdk.Coin
					amount, _ = sdk.NewIntFromString(strings.Trim(string(attribute.Value), data.Denom))
				}
			}
		}
	}

	return sender, recipient, amount
}
