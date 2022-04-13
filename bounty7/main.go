package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	app "github.com/cosmos/gaia/v6/app"
	gaiacmd "github.com/cosmos/gaia/v6/cmd/gaiad/cmd"
	"github.com/notional-labs/gaia-analyzer/handler/cli"
)

func main() {
	rootCmd, _ := gaiacmd.NewRootCmd()

	rootCmd.AddCommand(
		cli.TrackCommand(),
	)

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

// func main() {
// 	rootCmd, _ := cosmoscmd.NewRootCmd(
// 		app.Name,
// 		app.AccountAddressPrefix,
// 		app.DefaultNodeHome,
// 		app.Name,
// 		app.ModuleBasics,
// 		app.New,
// 		// this line is used by starport scaffolding # root/arguments
// 	)

// 	rootCmd.AddCommand(
// 		cmd.GovCommand(),
// 	)

// 	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
// 		os.Exit(1)
// 	}
// }
