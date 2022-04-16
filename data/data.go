package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/gaia-analyzer/types"
)

var (
	// To check if this account has been tracked. A Tracked account is an account that receives atom from the root account
	IsTrackedAccount map[string]bool = map[string]bool{}
	// Amount of coin a given account has
	Balance map[string]uint64 = map[string]uint64{}
	// Amount of tracked coin a given account has
	TrackedBalance map[string]sdk.Int = map[string]sdk.Int{}
	// A priority of txs with priority indicator being the tx height
	TrackedTxQueue types.LowestHeightFirstOutEventQueue = types.LowestHeightFirstOutEventQueue{}

	Denom string

	NeutralAccount map[string]bool = map[string]bool{}

	BondedTrackedBalance map[string]sdk.Int = map[string]sdk.Int{}

	BondedBalance map[string]uint64 = map[string]uint64{}
)
