package main

import (
	"os"

	"github.com/notional-labs/gaia-analyzer/handler"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/notional-labs/dig/app"
	"github.com/tendermint/spm/cosmoscmd"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)
	rootCmd.AddCommand(
		handler.BlockCommand(),
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
