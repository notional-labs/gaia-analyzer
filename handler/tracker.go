package handler

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/notional-labs/gaia-analyzer/types"
)

func GetGovVoteData(clientCtx client.Context, proposalId int) map[string]types.Vote {
	voteResult := make(map[string]types.Vote)

	txs := getAllGovTxsByQuery(clientCtx, 100, proposalId)

	for _, txRes := range txs {
		voterAddress := string(txRes.Events[0].Attributes[0].GetValue())
		_, ok := voteResult[voterAddress]

		if ok && txRes.Height < voteResult[voterAddress].Height {
			continue
		}

		if txRes.Code != 0 {
			continue
		}

		if len(txRes.Events) == 6 {
			voteResult[voterAddress] = types.Vote{
				Option:     getOption(string(txRes.Events[4].Attributes[0].GetValue())),
				ProposalId: proposalId,
				Height:     txRes.Height,
				TxHash:     txRes.TxHash,
			}
			continue

		}

		voteResult[voterAddress] = types.Vote{
			Option:     getOption(string(txRes.Events[8].Attributes[0].GetValue())),
			ProposalId: proposalId,
			Height:     txRes.Height,
			TxHash:     txRes.TxHash,
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
