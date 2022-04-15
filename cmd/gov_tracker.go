package cmd

import (
	"fmt"
	"strconv"

	appquery "github.com/notional-labs/gaia-analyzer/db-query/app"
	"github.com/notional-labs/gaia-analyzer/handler"
	"github.com/spf13/cobra"
)

func GovTrackCommand() *cobra.Command {
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

func QueryDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "querydata [query_string]",
		Short: "Get data by database",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// var tmEvents = []string{
			// 	"message.action='/cosmos.bank.v1beta1.MsgSend'",
			// 	"tx.height<5",
			// 	// fmt.Sprintf("proposal_vote.proposal_id='%d'", proposalID),
			// }

			data, err := appquery.GetUatomBalanceAtHeight("cosmos1d9725dhaq06mayzfn8ape3kcfn8lmuypquutu6", 2)
			fmt.Println(err.Error())
			fmt.Print(data)
			return nil
		},
	}
	cmd.Flags().String("fsdfdsf", "uatom", "a")

	return cmd
}
