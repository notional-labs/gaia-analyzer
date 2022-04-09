package handler

import (
	"github.com/notional-labs/gaia-analyzer/types"
)

var (
	IsWhaleAddress map[string]bool = map[string]bool{}
)

type Scorer func(a *types.AccountData)

func AccountInactivity(a *types.AccountData) {
	if a.AccountMetadata.NumTxs == 0 {
		a.Score -= 1
	}
}

func AccountGovParticipation(a *types.AccountData) {
	if a.AccountMetadata.NumVotes == 0 {
		a.Score -= 1
	}
}

func AccountRelatedToWhale(a *types.AccountData) {
	_, ok := IsWhaleAddress[a.Address]
	if ok {
		a.Score -= 9999999
	}
}
