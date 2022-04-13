package cmd

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/notional-labs/gaia-analyzer/handler"
	dbquery "github.com/notional-labs/gaia-analyzer/handler/db-query"
	"github.com/spf13/cobra"
)

func GovCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal_id [num_process]",
		Short: "Get verified data for a the blocks",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			proposalID, err := strconv.Atoi(args[0])

			if err != nil {
				return err
			}

			data := handler.GetGovVoteData(clientCtx, proposalID)
			fmt.Print(data)
			return nil
		},
	}
	cmd.Flags().StringP(flags.FlagNode, "n", "tcp://localhost:26657", "Node to connect to")

	return cmd

}

func QueryDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "querydata [query_string]",
		Short: "Get data by database",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			// proposalID, err := strconv.Atoi(args[0])

			if err != nil {
				return err
			}

			var tmEvents = []string{
				"message.action='/cosmos.bank.v1beta1.MsgSend'",
				"tx.height<5",
				// fmt.Sprintf("proposal_vote.proposal_id='%d'", proposalID),
			}

			data := dbquery.QueryTxs(clientCtx, "/home/vuong/.dig", tmEvents)
			fmt.Print(data)
			return nil
		},
	}
	cmd.Flags().StringP(flags.FlagNode, "n", "tcp://localhost:26657", "Node to connect to")

	return cmd
}
