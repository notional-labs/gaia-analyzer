package handler

import (
	"fmt"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	gaiaapp "github.com/cosmos/gaia/v6/app"
	"github.com/notional-labs/gaia-analyzer/types"
)

func UpdateAction(ad *types.AccountData, msg sdktypes.Msg) {
	msgTypeURL := sdktypes.MsgTypeURL(msg)

	switch msgTypeURL {
	case types.DelegateMsgUrl:
		var delegation stakingtypes.Delegation
		encCfg := gaiaapp.MakeEncodingConfig()
		encCfg.Marshaler.Unmarshal([]byte(msg.String()), &delegation)

		delegateAnalyzer := types.Delegation{
			DelegatedValidator: delegation.DelegatorAddress,
			Amount:             delegation.Shares.String(),
		}
		ad.Actions.Delegations = append(ad.Actions.Delegations, delegateAnalyzer)
	}
}

func GetDelegateMsgData(msg sdktypes.Msg) types.Delegation {
	var delegationMsg stakingtypes.MsgDelegate
	encCfg := gaiaapp.MakeEncodingConfig()
	err := encCfg.Marshaler.UnmarshalJSON([]byte(msg), &delegationMsg)

	fmt.Println(err)
	fmt.Println(delegationMsg)

	delegateAnalyzer := types.Delegation{
		DelegatedValidator: delegationMsg.DelegatorAddress,
		Amount:             delegationMsg.Amount.String(),
	}
	return delegateAnalyzer
}
