package handler

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/notional-labs/gaia-analyzer/data"
	"github.com/notional-labs/gaia-analyzer/types"
)

func TrackCoinsFromAccount(clientCtx client.Context, address string, beginBlock int64) {
	GetBankSendTxsFromAddress(clientCtx, address, beginBlock)
	GetBankSendTxsToAddress(clientCtx, address, beginBlock)

	ThisBalanceAtHeightKey := types.BalanceAtHeightKey{
		Address: address,
		Height:  beginBlock,
	}

	atomAmountInThisAccount := GetAtomBalanceAtHeight(clientCtx, address, beginBlock)

	Balances[ThisBalanceAtHeightKey] = atomAmountInThisAccount
	TrackedCoinBalances[ThisBalanceAtHeightKey] = atomAmountInThisAccount

	for {
		if len(data.TxTimeQueue) == 0 {
			break
		}

	}

}
