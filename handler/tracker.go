package handler

import (
	"container/heap"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/notional-labs/gaia-analyzer/types"
)

func calSentTrackedAtomAmount(atomBalance float64, trackedAtomBalance float64, sentAtomAmount float64) float64 {
	return sentAtomAmount * trackedAtomBalance / atomBalance
}

func updateAtomBalance(clientCtx client.Context, address string, height int64) float64 {
	balance, ok := AtomBalances[address]
	if !ok {
		AtomBalances[address] = float64(GetAtomBalanceAtHeight(clientCtx, address, height))
	}
	return balance
}

func handle_tx(clientCtx client.Context, tx *types.TimeTx) string {
	height := tx.Tx.Height

	sender, recipient, sentAtomAmount := ParseBankSendTxEvent(tx.Tx)

	senderAtomBalance := updateAtomBalance(clientCtx, sender, height)

	sentTrackedAtomAmount := calSentTrackedAtomAmount(senderAtomBalance, TrackedAtomBalances[sender], sentAtomAmount)

	TrackedAtomBalances[recipient] = TrackedAtomBalances[recipient] + sentTrackedAtomAmount

	return sender
}

func TrackCoinsFromAccount(clientCtx client.Context, address string, beginBlock int64) {
	GetBankSendTxsFromAddress(clientCtx, address, beginBlock)
	GetBankSendTxsToAddress(clientCtx, address, beginBlock)

	atomAmountInThisAccount := GetAtomBalanceAtHeight(clientCtx, address, beginBlock)

	AtomBalances[address] = atomAmountInThisAccount
	TrackedAtomBalances[address] = atomAmountInThisAccount

	for {
		// if queue empty, stop
		if len(TxQueue) == 0 {
			break
		}

		// get next tx from tx queue
		// tx queue : priority queue of txs with priority indicator being the tx's height
		tx := heap.Pop(&TxQueue).(*types.TimeTx)
		// apply tx logic, output sender of tx
		sender := handle_tx(clientCtx, tx)

		_, isTracked := IsTrackedAccount[sender]
		// if this account is not tracked yet, this means this account has not received any tracked atom before this tx
		// query BankSend tx this account sent starting from this tx height and push to tx queue
		if !isTracked {
			TrackedAtomBalances[address] = 0
			GetBankSendTxsFromAddress(clientCtx, sender, tx.Tx.Height)
		}
	}
	fmt.Printf("%+v", TrackedAtomBalances)
}
