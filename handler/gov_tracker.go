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
		voterAddress := string(txRes.Result.Events[0].Attributes[0].GetValue())
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
		}
	}
	fmt.Println(len(voteResult))
	return voteResult
}

// raw option like {\"option\":1,\"weight\":\"1.000000000000000000\"}
func getOption(rawOption string) string {
	if strings.Contains(rawOption, `"option":1`) {
		return "Yes"
	}

	if strings.Contains(rawOption, `"option":2`) {
		return "No"
	}

	if strings.Contains(rawOption, `"option":3`) {
		return "Abstain"
	}

	return "NoWithVeto"
}
