package handler

import (
	"fmt"
	"strings"

	"github.com/notional-labs/gaia-analyzer/types"
)

func GetGovVoteData(proposalId int) map[string]types.Vote {
	voteResult := make(map[string]types.Vote)

	txs := govVoteQueries(proposalId)

	for _, txRes := range txs {
		var voterAddress string
		for _, v := range txRes.Result.Events[0].Attributes {
			if string(v.GetKey()) == "sender" {
				voterAddress = string(v.GetValue())
				break
			}
		}
		_, ok := voteResult[voterAddress]

		if ok && txRes.Height < voteResult[voterAddress].Height {
			continue
		}

		if txRes.Result.Code != 0 {
			continue
		}

		voteResult[voterAddress] = types.Vote{
			Option:     getOption(string(txRes.Result.String())),
			ProposalId: proposalId,
			Height:     txRes.Height,
			TxHash:     fmt.Sprint(txRes.Result.Code),
		}
	}
	return voteResult
}

// raw option like {\"option\":1,\"weight\":\"1.000000000000000000\"}
func getOption(rawOption string) string {
	if strings.Contains(rawOption, `"option":1`) || strings.Contains(rawOption, `VOTE_OPTION_YES`) {
		return "Yes"
	}

	if strings.Contains(rawOption, `"option":2`) || strings.Contains(rawOption, `VOTE_OPTION_NO`) {
		return "No"
	}

	if strings.Contains(rawOption, `"option":3`) || strings.Contains(rawOption, `VOTE_OPTION_NO_WITH_VETO`) {
		return "NoWithVeto"
	}

	return "Abstain"
}
