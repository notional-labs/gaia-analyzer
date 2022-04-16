package handler

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/notional-labs/gaia-analyzer/db-query/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func ExportStateBankAtHeight(height int64) banktypes.GenesisState {
	ctx := app.EmptyApp.NewContext(true, tmproto.Header{Height: app.EmptyApp.LastBlockHeight()})
	return *app.EmptyApp.BankKeeper.ExportGenesis(ctx)
}
