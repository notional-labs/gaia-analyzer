package handler

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	dbquery "github.com/notional-labs/gaia-analyzer/handler/db-query"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func govVoteQueries(clientCtx client.Context, proposalID int) []*ctypes.ResultTx {
	var tmEvents = []string{
		"message.action='/cosmos.gov.v1beta1.MsgVote'",
		fmt.Sprintf("proposal_vote.proposal_id='%d'", proposalID),
	}
	txsRes := dbquery.QueryTxs(clientCtx, "/home/vuong/.dig", tmEvents)
	return txsRes
}
