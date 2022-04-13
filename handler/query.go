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

// query txs that send atom from a given address since a given height. Push those txs to global tx queue
func GetBankSendAtomFromAddress(clientCtx client.Context, sender string, beginHeight int64) {
	fmt.Println(beginHeight)
	// events use for query
	senderEvent := fmt.Sprintf("%s='%s'", "message.sender", sender)
	bankSendEvent := fmt.Sprintf("%s='%s'", "message.action", "/cosmos.bank.v1beta1.MsgSend")
	tmEvents := []string{senderEvent, bankSendEvent}

	page := 1
	limit := 10

	// query txs by events
	txs, err := authtx.QueryTxsByEvents(clientCtx.WithHeight(beginHeight), tmEvents, page, limit, "")
	if err != nil {
		panic(err)
	}

	// push to tx queue
	// A priority of txs with priority indicator being the tx height
	for _, tx := range txs.Txs {
		if IsBankSendAtomTx(tx) {
			TxItemWrapper := types.TxItemWrapper{
				Tx: tx,
			}
			heap.Push(&TxQueue, &TxItemWrapper)
			fmt.Println(ParseBankSendTxEvent(tx))
			fmt.Print(tx.Height)
		}
	}
}

// query txs that send atom to a given address since a given height. Push those txs to global tx queue
func GetBankSendAtomToAddress(clientCtx client.Context, receiver string, beginHeight int64) {
	fmt.Println(beginHeight)
	// events use for query
	senderEvent := fmt.Sprintf("%s='%s'", "coin_received.receiver", receiver)
	bankSendEvent := fmt.Sprintf("%s='%s'", "message.action", "/cosmos.bank.v1beta1.MsgSend")
	tmEvents := []string{senderEvent, bankSendEvent}

	page := 1
	limit := 10

	// query txs by events
	txs, err := authtx.QueryTxsByEvents(clientCtx.WithHeight(beginHeight), tmEvents, page, limit, "")
	if err != nil {
		panic(err)
	}

	// push to tx queue
	// A priority of txs with priority indicator being the tx height
	for _, tx := range txs.Txs {
		if IsBankSendAtomTx(tx) {
			TxItemWrapper := types.TxItemWrapper{
				Tx: tx,
			}
			heap.Push(&TxQueue, &TxItemWrapper)
			fmt.Println(ParseBankSendTxEvent(tx))
			fmt.Print(tx.Height)
		}
	}
}

// get atom amount a given account has at a given height
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
