package handler

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/notional-labs/gaia-analyzer/db-query/app"
)

func ExportStateBankAtHeight(height int64) banktypes.GenesisState {
	ctx := app.GetQueryContext(height)
	return *app.EmptyApp.BankKeeper.ExportGenesis(ctx)
}
