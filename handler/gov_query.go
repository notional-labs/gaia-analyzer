package handler

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

func govVoteQueries(clientCtx client.Context, page, limit, proposalID int) ([]*sdk.TxResponse, int) {
	var tmEvents = []string{
		"message.action='/cosmos.gov.v1beta1.MsgVote'",
		fmt.Sprintf("proposal_vote.proposal_id='%d'", proposalID),
	}
	txs, err := authtx.QueryTxsByEvents(clientCtx, tmEvents, page, limit, "")

	if err != nil {
		return nil, 0
	}
	fmt.Println(txs.TotalCount)
	return txs.Txs, int(txs.TotalCount)
}

func getAllGovTxsByQuery(clientCtx client.Context, limit, proposalID int) []*sdk.TxResponse {
	var txs []*sdk.TxResponse
	var page = 1

	txs, totalPage := govVoteQueries(clientCtx, page, limit, proposalID)
	for page <= totalPage {
		page = page + 1
		newTxs, _ := govVoteQueries(clientCtx, page, limit, proposalID)
		txs = append(txs, newTxs...)
	}
	return txs
}
