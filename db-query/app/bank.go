package app

import (
	"github.com/notional-labs/gaia-analyzer/data"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetUatomBalanceAtHeight(address string, height int64) (sdk.Int, error) {
	ctx := GetQueryContext(height)

	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	coin := EmptyApp.BankKeeper.GetBalance(*ctx, accAddress, data.Denom)
	return coin.Amount, nil
}
