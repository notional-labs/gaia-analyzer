package app

import (
	"fmt"

	"github.com/notional-labs/gaia-analyzer/data"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// We use the baseapp.QueryRouter here to do inter-module state querying.
// PLEASE DO NOT REPLICATE THIS PATTERN IN YOUR OWN APP.
func GetUatomBalanceAtHeight(address string, height int64) (sdk.Int, error) {
	ctx := GetQueryContext(height)

	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	coin := EmptyApp.BankKeeper.GetBalance(ctx, accAddress, data.TrackedDenom)
	fmt.Printf("Tracked coin of address %s is : %d", address, coin.Amount.Int64())
	fmt.Println("===================")
	return coin.Amount, nil
}
