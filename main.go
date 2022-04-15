package main

import (

	// "os"

	// "github.com/notional-labs/gaia-analyzer/cmd"

	"os"

	"github.com/notional-labs/gaia-analyzer/cmd"
	"github.com/notional-labs/gaia-analyzer/data"
	dbquery "github.com/notional-labs/gaia-analyzer/db-query"
	"github.com/spf13/cobra"
	// "github.com/spf13/cobra"
)

var (
	RootDirFlag = "root_dir"
)

func main() {

	// dbquery.Init("/Users/khanh/.dig")

	data.TrackedDenom = "stake"
	// data, err := appquery.GetUatomBalanceAtHeight("cosmos1d9725dhaq06mayzfn8ape3kcfn8lmuypquutu6", 2)
	// if err != nil {
	// 	panic(err)
	// }
	// data, err = appquery.GetUatomBalanceAtHeight("cosmos1d9725dhaq06mayzfn8ape3kcfn8lmuypquutu6", 2)
	// if err != nil {
	// 	panic(err)
	// }

	// handler.TrackCoinsFromAccount("cosmos1dq07qh6rc489le9wjlh9p3n5em3u24vwy94lxr", 1)

	// rootCmd := &cobra.Command{
	// 	Use:   "bounty7",
	// 	Short: "bounty7 gnolang solution",

	// 	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {

	// 		// trackedDenom, err := cmd.Flags().GetString(TrackedDenomFlag)
	// 		// if err != nil {
	// 		// 	return err
	// 		// }
	// 		// data.TrackedDenom = trackedDenom

	// 		rootDir, err := cmd.Flags().GetString(RootDirFlag)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		dbquery.Init(rootDir)

	// 		return nil
	// 	},
	// }
	// fmt.Println(data)
	// data, err = appquery.GetUatomBalanceAtHeight("cosmos1d9725dhaq06mayzfn8ape3kcfn8lmuypquutu6", 2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(data)

	// handler.TrackCoinsFromAccount("cosmos1dq07qh6rc489le9wjlh9p3n5em3u24vwy94lxr", 1)

	rootCmd := &cobra.Command{
		Use:   "bounty7",
		Short: "bounty7 gnolang solution",

		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {

			// trackedDenom, err := cmd.Flags().GetString(TrackedDenomFlag)
			// if err != nil {
			// 	return err
			// }
			// data.TrackedDenom = trackedDenom

			rootDir, err := cmd.Flags().GetString(RootDirFlag)
			if err != nil {
				return err
			}
			dbquery.Init(rootDir)

			return nil
		},
	}
	rootCmd.PersistentFlags().String(RootDirFlag, "/home/vuong/.dig", "path of chain data")

	rootCmd.AddCommand(
		cmd.GovTrackCommand(),
		cmd.QueryDatabase(),
		cmd.CoinTrackCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
