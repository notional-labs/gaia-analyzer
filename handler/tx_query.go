package handler

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func GetSendTxs(clientCtx client.Context, sender string, beginHeight int64) ([]*sdk.TxResponse, error) {

	senderEvent := fmt.Sprintf("%s='%s'", "message.sender", sender)

	bankSendEvent := fmt.Sprintf("%s='%s'", "message.action", "/cosmos.bank.v1beta1.MsgSend")

	tmEvents := []string{senderEvent, bankSendEvent}

	page := 1
	limit := 100

	txs, err := authtx.QueryTxsByEvents(clientCtx.WithHeight(beginHeight), tmEvents, page, limit, "")
	if err != nil {
		return nil, err
	}
	return txs.Txs, nil
}

func GetAtomBalanceAtHeight(clientCtx client.Context, addressStr string, height int64) (uint64, error) {
	queryClient := banktypes.NewQueryClient(clientCtx.WithHeight(height))
	addr, err := sdk.AccAddressFromBech32(addressStr)
	if err != nil {
		return 0, err
	}
	params := types.NewQueryBalanceRequest(addr, "uatom")
	res, err := queryClient.Balance(context.Background(), params)
	return res.Balance.Amount.Uint64(), nil
}
