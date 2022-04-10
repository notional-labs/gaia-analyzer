package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DelegateMsgUrl   = "/cosmos.staking.v1beta1.MsgDelegate"
	UndelegateMsgUrl = "cosmos.staking.v1beta1.MsgDelegate"
	VoteMsgUrl       = "/cosmos.gov.v1beta1.MsgVote"
	SubMitProposal   = "/cosmos.gov.v1beta1.MsgSubmitProposal"
)

type Block struct {
	Txs []sdk.Tx
}
