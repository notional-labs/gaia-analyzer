package handler

import (
	"fmt"

	txquery "github.com/notional-labs/gaia-analyzer/db-query/tx"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func QuerySendTxInBlock(blockHeight int) []*abcitypes.TxResult {
	var tmEvents = []string{
		fmt.Sprintf("tx.height=%d", blockHeight),
		"message.action='/cosmos.bank.v1beta1.MsgSend'",
	}
	txsRes := txquery.QueryTxs(tmEvents)
	return txsRes
}
