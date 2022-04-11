package handler

import (
	"math"

	"github.com/notional-labs/gaia-analyzer/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

func GetTaintedAccounts(node rpcclient.Client, proposalId int) (voteResult map[string]types.Vote) {

	yesTxs, err := govVoteQueries(node, 1, math.MaxInt32, proposalId, 1)

	for _, txRes := range yesTxs {
		voterAddress := string(txRes.Events[0].Attributes[2].GetValue())
		voteResult[voterAddress] = types.Vote{
			Option:     "Yes",
			ProposalId: proposalId,
			Height:     txRes.Height,
			TxHash:     txRes.TxHash,
		}
	}

	noTxs, err := govVoteQueries(node, 1, math.MaxInt32, proposalId, 2)

	for _, txRes := range noTxs {
		voterAddress := string(txRes.Events[0].Attributes[2].GetValue())
		_, ok := voteResult[voterAddress]

		if ok && txRes.Height < voteResult[voterAddress].Height {
			continue
		}

		voteResult[voterAddress] = types.Vote{
			Option:     "No",
			ProposalId: proposalId,
			Height:     txRes.Height,
			TxHash:     txRes.TxHash,
		}
	}

	abstainTxs, err := govVoteQueries(node, 1, math.MaxInt32, proposalId, 3)
	for _, txRes := range abstainTxs {
		voterAddress := string(txRes.Events[0].Attributes[2].GetValue())
		_, ok := voteResult[voterAddress]

		if ok && txRes.Height < voteResult[voterAddress].Height {
			continue
		}

		voteResult[voterAddress] = types.Vote{
			Option:     "Abstain",
			ProposalId: proposalId,
			Height:     txRes.Height,
			TxHash:     txRes.TxHash,
		}
	}
	if err != nil {
		return nil
	}
	return voteResult
}
