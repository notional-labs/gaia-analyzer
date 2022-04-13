package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/notional-labs/gaia-analyzer/handler"
	"github.com/spf13/cobra"
)

func TrackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "track_coin [root_address] [start_height]",
		Short: "Get verified data for a the blocks",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			rootAddress := args[0]
			startHeight, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			handler.TrackCoinsFromAccount(clientCtx, rootAddress, startHeight)
			return nil
		},
	}
	cmd.Flags().StringP(flags.FlagNode, "n", "tcp://localhost:26657", "Node to connect to")

	return cmd

}
