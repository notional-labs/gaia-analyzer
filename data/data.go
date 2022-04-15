package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/gaia-analyzer/types"
)

var (
	// To check if this account has been tracked. A Tracked account is an account that receives atom from the root account
	IsTrackedAccount map[string]bool = map[string]bool{}
	// Amount of atom a given account has
	UatomBalance map[string]sdk.Int = map[string]sdk.Int{}
	// Amount of tracked atom a given account has
	TrackedUatomBalance map[string]sdk.Int = map[string]sdk.Int{}
	// A priority of txs with priority indicator being the tx height
	TxQueue types.LowestHeightFirstOutTxQueue = types.LowestHeightFirstOutTxQueue{}

	TrackedDenom string
)
