package handler

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"container/heap"
)

//
func GetBankSendTxsFromAddress(clientCtx client.Context, sender string, beginHeight int64) {

	senderEvent := fmt.Sprintf("%s='%s'", "message.sender", sender)

	bankSendEvent := fmt.Sprintf("%s='%s'", "message.action", "/cosmos.bank.v1beta1.MsgSend")

	tmEvents := []string{senderEvent, bankSendEvent}

	page := 1
	limit := 100

	// query txs by events
	txs, err := authtx.QueryTxsByEvents(clientCtx.WithHeight(beginHeight), tmEvents, page, limit, "")
	if err != nil {
		panic(err)
	}
	// push to tx queue
	for _, tx := range txs.Txs {
		_, ok := AddedBlocks[tx.Height]
		if !ok {
			AddedBlocks[tx.Height] = true
			heap.Push(BlockQueue, tx.Height)
			TxsAtBlock[tx.Height] = append(TxsAtBlock[tx.Height], tx)

		}
	}

}

func GetBankSendTxsToAddress(clientCtx client.Context, receiver string, beginHeight int64) {

	senderEvent := fmt.Sprintf("%s='%s'", "coin_received.receiver", receiver)

	bankSendEvent := fmt.Sprintf("%s='%s'", "message.action", "/cosmos.bank.v1beta1.MsgSend")

	tmEvents := []string{senderEvent, bankSendEvent}

	page := 1
	limit := 100

	// query txs by events
	txs, err := authtx.QueryTxsByEvents(clientCtx.WithHeight(beginHeight), tmEvents, page, limit, "")
	if err != nil {
		panic(err)
	}

	// push to tx queue
	for _, tx := range txs.Txs {
		_, ok := AddedBlocks[tx.Height]
		if !ok {
			AddedBlocks[tx.Height] = true
			heap.Push(BlockQueue, tx.Height)
			TxsAtBlock[tx.Height] = append(TxsAtBlock[tx.Height], tx)

		}
	}
}

func GetAtomBalanceAtHeight(clientCtx client.Context, addressStr string, height int64) uint64 {
	queryClient := banktypes.NewQueryClient(clientCtx.WithHeight(height))
	addr, err := sdk.AccAddressFromBech32(addressStr)
	if err != nil {
		panic("can't query address balance")
	}
	params := banktypes.NewQueryBalanceRequest(addr, "uatom")
	res, err := queryClient.Balance(context.Background(), params)
	return res.Balance.Amount.Uint64()
}
