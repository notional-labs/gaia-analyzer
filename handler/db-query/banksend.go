package dbquery

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
)

func QueryBankSendAtomFromAddress(ctx client.Context, rootDir string, sender string, height int64) {
	senderEvent := fmt.Sprintf("%s='%s'", "message.sender", sender)
	bankSendEvent := fmt.Sprintf("%s='%s'", "message.action", "/cosmos.bank.v1beta1.MsgSend")

	tmEvents := []string{senderEvent, bankSendEvent}

	QueryTxs(ctx, rootDir, tmEvents)
}
