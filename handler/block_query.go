package handler

import (
	"fmt"

	txquery "github.com/notional-labs/gaia-analyzer/db-query/tx"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func QueryBlock(blockHeight int) []*abcitypes.TxResult {
	var tmEvents = []string{
		fmt.Sprintf("block.height='%d'", blockHeight),
	}
	txsRes := txquery.QueryTxs(tmEvents)
	return txsRes
}
