package handler

import (
	"github.com/cosmos/cosmos-sdk/client"
)

func TrackCoinsFromAccount(clientCtx client.Context, address string, beginBlock int64) {
	GetSendTxsAndPushToTxQueue(clientCtx, address, beginBlock)

}
