package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/gaia-analyzer/types"
)

var (
	// To check if this account has been tracked. A Tracked account is an account that receives atom from the root account
	IsTrackedAccount map[string]bool = map[string]bool{}
	// Amount of coin a given account has
	Balance map[string]sdk.Int = map[string]sdk.Int{}
	// Amount of tracked coin a given account has
	TrackedBalance map[string]sdk.Int = map[string]sdk.Int{}
	// A priority of txs with priority indicator being the tx height
	TrackedTxQueue types.LowestHeightFirstOutEventQueue = types.LowestHeightFirstOutEventQueue{}

	Denom string
	// accounts like unbonded pool or bonded pool, coins send in and out of here will not be blended
	NeutralAccount map[string]bool = map[string]bool{}
)
