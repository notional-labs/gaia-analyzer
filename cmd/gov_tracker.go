package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/notional-labs/gaia-analyzer/handler"
	"github.com/spf13/cobra"
)

func GovTrackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal_id [proposalID]",
		Short: "Get verified data for a the blocks",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			proposalID, err := strconv.Atoi(args[0])

			if err != nil {
				return err
			}

			data := handler.GetGovVoteData(proposalID)
			bs, _ := json.Marshal(data)
			fmt.Println(string(bs))
			return nil
		},
	}

	return cmd

}

func QueryDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trackaccount [address] [start] [end] ",
		Short: "Get data by database",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			// var tmEvents = []string{
			// 	"message.action='/cosmos.bank.v1beta1.MsgSend'",
			// 	// "tx.height>7314636",
			// 	"message.sender='cosmos1000ya26q2cmh399q4c5aaacd9lmmdqp92z6l7q'",
			// 	// fmt.Sprintf("proposal_vote.proposal_id='%d'", proposalID),
			// }

			startBlock, _ := strconv.Atoi(args[1])
			endBlock, _ := strconv.Atoi(args[2])

			handler.ExecuteTrack(args[0], startBlock, endBlock)
			fmt.Println(handler.CoinTracker)
			return nil
		},
	}
	cmd.Flags().String("fsdfdsf", "uatom", "a")

	return cmd
}
