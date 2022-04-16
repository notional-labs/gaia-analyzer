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

		for _, event := range txRes.Result.Events {
			for _, attribute := range event.Attributes {
				if string(attribute.GetKey()) == "sender" {
					voterAddress = string(attribute.GetValue())
					break
				}
			}

		}
		fmt.Println(voterAddress)
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
	if strings.Contains(rawOption, `option\\\":1`) || strings.Contains(rawOption, `VOTE_OPTION_YES`) {
		return "Yes"
	}
	fmt.Println(rawOption)
	if strings.Contains(rawOption, `option\\\":1`) || strings.Contains(rawOption, `VOTE_OPTION_NO`) {
		return "No"
	}

	if strings.Contains(rawOption, `option\\\":1`) || strings.Contains(rawOption, `VOTE_OPTION_NO_WITH_VETO`) {
		return "NoWithVeto"
	}

	return "Abstain"
}
