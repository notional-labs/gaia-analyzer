package cmd

import (
	"fmt"
	"strconv"

	"github.com/notional-labs/gaia-analyzer/data"
	"github.com/notional-labs/gaia-analyzer/handler"
	"github.com/spf13/cobra"
)

func CoinTrackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trackcoin [address] [height] [denom]",
		Short: "Get verified data for a the blocks",
		Args:  cobra.MaximumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			height, err := strconv.Atoi(args[1])

			if err != nil {
				return err
			}
			data.TrackedDenom = args[2]

			handler.TrackCoinsFromAccount(args[0], int64(height))
			fmt.Print(data.TrackedUatomBalance)
			return nil
		},
	}
	return cmd
}
