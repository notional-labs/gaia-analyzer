package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/notional-labs/gaia-analyzer/handler"
	"github.com/spf13/cobra"
)

func BankTrackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bankstate [height]",
		Short: "Export bank state at block",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			height, err := strconv.Atoi(args[0])

			if err != nil {
				return err
			}

			data := handler.ExportStateBankAtHeight(int64(height))
			bs, _ := json.Marshal(data)
			fmt.Println(string(bs))
			return nil
		},
	}

	return cmd
}
