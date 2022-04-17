package app

import (
	"github.com/notional-labs/gaia-analyzer/data"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetUatomBalanceAtHeight(address string, height int64) (uint64, error) {
	ctx := GetQueryContext(height)

	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return 0, err
	}
	coin := EmptyApp.BankKeeper.GetBalance(ctx, accAddress, data.Denom)
	return coin.Amount.Uint64(), nil
}
