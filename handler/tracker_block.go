package handler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/gaia-analyzer/data"
	"github.com/notional-labs/gaia-analyzer/db-query/app"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

var (
	CoinTracker map[string]sdk.Int
	StartBlock  int
)

func setCoinTracker(address string, blockHeight int) error {
	CoinTracker = make(map[string]sdk.Int)

	amount, err := app.GetUatomBalanceAtHeight(address, int64(blockHeight))
	CoinTracker[address] = amount
	return err
}

func updateCoinTrackerByTx(tx abcitypes.TxResult) {
	var trueEvent abcitypes.Event
	for _, v := range tx.Result.Events {
		if v.Type == "transfer" {
			trueEvent = v
		}
	}
	sender := string(trueEvent.Attributes[1].GetValue())
	recipient := string(trueEvent.Attributes[0].GetValue())
	amountinTx := string(trueEvent.Attributes[2].GetValue())
	if !strings.Contains(amountinTx, data.TrackedDenom) {
		return
	}

	currentTrackedCoin, ok := CoinTracker[string(sender)]

	if !ok {
		return
	}

	re := regexp.MustCompile("[0-9]+")
	a := re.FindAllString(amountinTx, -1)
	tempAmount, _ := strconv.Atoi(a[0])
	amountTransfer := sdk.NewInt(int64(tempAmount))
	if amountTransfer.GTE(currentTrackedCoin) {
		coin, ok := CoinTracker[recipient]
		if ok {
			CoinTracker[recipient] = coin.Add(currentTrackedCoin)
		} else {
			CoinTracker[recipient] = currentTrackedCoin
		}
		delete(CoinTracker, sender)
		fmt.Printf("Tracked coin from %s to %s : %d  \n", sender, recipient, currentTrackedCoin)

	} else {
		coin, ok := CoinTracker[recipient]
		if ok {
			CoinTracker[recipient] = coin.Add(amountTransfer)
		} else {
			CoinTracker[recipient] = amountTransfer
		}
		CoinTracker[sender] = currentTrackedCoin.Sub(amountTransfer)
		fmt.Printf("Tracked coin from %s to %s : %s  \n", sender, recipient, amountTransfer)
	}
}

func updateCoinTrackerByBlock(blockHeight int) {
	txs := QuerySendTxInBlock(blockHeight)
	for _, tx := range txs {
		updateCoinTrackerByTx(*tx)
	}
}

func ExecuteTrack(address string, blockStart int, blockEnd int) {
	setCoinTracker(address, blockStart)
	for i := blockStart; i <= blockEnd; i++ {
		updateCoinTrackerByBlock(i)
	}
}
