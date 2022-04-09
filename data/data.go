package data

import (
	"github.com/notional-labs/gaia-analyzer/types"
)

var (
	// map account to its data
	AccountDataMap map[string][]types.AccountData
	// map all tainted accounts to their score
	TaintedAccounts []string
	// send from account to account for how many atoms, use to track atom sent by tainted accounts
	SendData map[types.SendFromTo]uint64
)
