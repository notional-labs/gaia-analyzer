package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Block struct {
	Txs []sdk.Tx
}
