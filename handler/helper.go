package handler

import (
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	gaiaapp "github.com/cosmos/gaia/v6/app"
)

func DecodeTx(txBytes []byte) ([]sdk.Msg, error) {
	encCfg := gaiaapp.MakeEncodingConfig()
	tx, err := encCfg.TxConfig.TxDecoder()(txBytes)

	if err != nil {
		return nil, err
	}

	// json, err := encCfg.TxConfig.TxJSONEncoder()(tx)
	return tx.GetMsgs(), nil
}

func ParseBankSendTxEvent(tx *sdk.TxResponse) (string, string, float64) {
	var amount float64
	var sender, recipient string
	for _, event := range tx.Events {
		if event.Type == "transfer" {
			for _, attribute := range event.Attributes {
				switch string(attribute.Key) {
				case "sender":
					sender = string(attribute.Value)
				case "recipient":
					recipient = string(attribute.Value)
				case "amount":
					amount, _ = strconv.ParseFloat(strings.Trim(string(attribute.Value), "uatom"), 64)
				}
			}
		}
	}

	return sender, recipient, amount

}
