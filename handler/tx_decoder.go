package handler

import (
	"github.com/cosmos/cosmos-sdk/types"
	gaiaapp "github.com/cosmos/gaia/v6/app"
)

func DecodeTx(txBytes []byte) ([]types.Msg, error) {
	encCfg := gaiaapp.MakeEncodingConfig()
	tx, err := encCfg.TxConfig.TxDecoder()(txBytes)

	if err != nil {
		return nil, err
	}

	// json, err := encCfg.TxConfig.TxJSONEncoder()(tx)
	return tx.GetMsgs(), nil
}
