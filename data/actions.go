package data

import (
	"github.com/notional-labs/gaia-analyzer/types"
)

var (
	AccountDataMap map[string][]types.AccountData
	// map all tainted accounts to their score
	TaintedAccounts []string
)
