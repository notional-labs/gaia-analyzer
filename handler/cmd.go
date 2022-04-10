package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// BlockCommand returns the verified block data for a given heights
func BlockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block [start] [end] [num_process]",
		Short: "Get verified data for a the blocks",
		Args:  cobra.MaximumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			startTime := time.Now()
			clientCtx, err := client.GetClientQueryContext(cmd)
			node, err := clientCtx.GetNode()

			if err != nil {
				return err
			}

			if err != nil {
				return err
			}
			var start, end, numberProcess int
			// optional height
			if len(args) > 0 {
				start, err = strconv.Atoi(args[0])
				if err != nil {
					return err
				}

				end, err = strconv.Atoi(args[1])

				if err != nil {
					return err
				}

				numberProcess, err = strconv.Atoi(args[2])

				if err != nil {
					return err
				}
			}

			output, err := getBlocks(node, start, end, numberProcess)
			if err != nil {
				return err
			}
			for _, v := range output {
				msg, _ := DecodeTx(v)
				fmt.Println(types.MsgTypeURL(msg[0]))
				fmt.Println(msg)
			}
			fmt.Println(time.Now().Sub(startTime))
			return nil
		},
	}

	cmd.Flags().StringP(flags.FlagNode, "n", "tcp://localhost:26657", "Node to connect to")

	return cmd
}
