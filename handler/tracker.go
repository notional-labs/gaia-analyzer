package handler

import (
	"container/heap"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/notional-labs/gaia-analyzer/types"
)

// cal tracked atom sent from a tracked account using blended, not FIFO
// A Tracked account is an account that receives atom from the root account
func calSentTrackedAtomAmount(atomBalance float64, trackedAtomBalance float64, sentAtomAmount float64) float64 {
	return sentAtomAmount * trackedAtomBalance / atomBalance
}

// updates atom balance to AtomBalance
func updateAtomBalance(clientCtx client.Context, address string, height int64) float64 {
	balance, ok := AtomBalance[address]
	if !ok {
		AtomBalance[address] = float64(GetAtomBalanceAtHeight(clientCtx, address, height))
	}
	return balance
}

// Apply bank send tx, update tracked atom balance after tx
func handle_tx(clientCtx client.Context, tx *types.TxItemWrapper) string {
	height := tx.Tx.Height
	sender, recipient, sentAtomAmount := ParseBankSendTxEvent(tx.Tx)
	// cal amount of tracked atom sent using blended, not FIFO
	senderAtomBalance := updateAtomBalance(clientCtx, sender, height)
	sentTrackedAtomAmount := calSentTrackedAtomAmount(senderAtomBalance, TrackedAtomBalance[sender], sentAtomAmount)
	// update tracked atom balance for recipient account and sender account
	TrackedAtomBalance[recipient] = TrackedAtomBalance[recipient] + sentTrackedAtomAmount
	TrackedAtomBalance[sender] = TrackedAtomBalance[sender] - sentTrackedAtomAmount

	return sender
}

// tracked all atom from a given address; start tracking from a given height; let's call this address root address
func TrackCoinsFromAccount(clientCtx client.Context, rootAddress string, startHeight int64) {

	// get all bank send from root address and to root address, push to global tx queue
	GetBankSendAtomFromAddress(clientCtx, rootAddress, startHeight)
	GetBankSendAtomToAddress(clientCtx, rootAddress, startHeight)
	// query chain to get root account atom balance at start height
	atomAmountInThisAccount := GetAtomBalanceAtHeight(clientCtx, rootAddress, startHeight)
	// update atom balance and tracked atom balance of root address
	// tracked atom balance = atom balance at start height since we have to track all atoms from root account balance at that height
	AtomBalance[rootAddress] = atomAmountInThisAccount
	TrackedAtomBalance[rootAddress] = atomAmountInThisAccount

	for {
		// if tx queue empty, stop
		if len(TxQueue) == 0 {
			break
		}

		// get next tx from tx queue
		// tx queue : priority queue of txs with priority indicator being the tx's height
		tx := heap.Pop(&TxQueue).(*types.TxItemWrapper)
		// apply handle tx, output sender of tx
		sender := handle_tx(clientCtx, tx)

		_, isTracked := IsTrackedAccount[sender]
		// if this account is not tracked yet, this means this account has not received any tracked atom before this tx
		// query BankSend tx this account sent starting from this tx height and push to tx queue
		if !isTracked {
			TrackedAtomBalance[sender] = 0
			GetBankSendAtomFromAddress(clientCtx, sender, tx.Tx.Height)
		}
	}
	fmt.Printf("%+v", TrackedAtomBalance)
}
