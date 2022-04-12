package handler

import (
	"context"
	"fmt"

	"container/heap"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/notional-labs/gaia-analyzer/types"
)

//
func GetBankSendFromAddress(clientCtx client.Context, sender string, beginHeight int64) {

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
		timeTx := types.TimeTx{
			Tx: tx,
		}
		heap.Push(&TxQueue, timeTx)
	}
}

func GetBankSendToAddress(clientCtx client.Context, receiver string, beginHeight int64) {

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
		timeTx := types.TimeTx{
			Tx: tx,
		}
		heap.Push(&TxQueue, timeTx)
	}
}

func GetAtomBalanceAtHeight(clientCtx client.Context, addressStr string, height int64) float64 {
	queryClient := banktypes.NewQueryClient(clientCtx.WithHeight(height))
	addr, err := sdk.AccAddressFromBech32(addressStr)
	if err != nil {
		panic(err)
	}
	params := banktypes.NewQueryBalanceRequest(addr, "uatom")
	res, err := queryClient.Balance(context.Background(), params)
	if err != nil {
		panic(err)
	}
	return float64(res.Balance.Amount.Uint64() / 1000000)
}
