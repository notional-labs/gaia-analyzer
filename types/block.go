package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Block struct {
	Height uint64
	Txs    []sdk.Tx
}
