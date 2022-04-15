package cmd

import (
	"fmt"
	"strconv"

	"github.com/notional-labs/gaia-analyzer/handler"
	"github.com/spf13/cobra"
)

func CoinTrackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal_id [num_process]",
		Short: "Get verified data for a the blocks",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			proposalID, err := strconv.Atoi(args[0])

			if err != nil {
				return err
			}

			data := handler.GetGovVoteData(proposalID)
			fmt.Print(data)
			return nil
		},
	}

	return cmd

}
