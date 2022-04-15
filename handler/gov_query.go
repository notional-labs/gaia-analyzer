package handler

import (
	"fmt"

	txquery "github.com/notional-labs/gaia-analyzer/db-query/tx"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func govVoteQueries(proposalID int) []*abcitypes.TxResult {
	var tmEvents = []string{
		"message.action='/cosmos.gov.v1beta1.MsgVote'",
		fmt.Sprintf("proposal_vote.proposal_id='%d'", proposalID),
	}
	txsRes := txquery.QueryTxs(tmEvents)
	return txsRes
}
