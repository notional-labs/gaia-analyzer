package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gaiaapp "github.com/cosmos/gaia/v6/app"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func govVoteQueries(node rpcclient.Client, page, limit, proposalID int, voteOption int) ([]*sdk.TxResponse, error) {
	encCfg := gaiaapp.MakeEncodingConfig()

	var events = []string{
		"message.action=/cosmos.gov.v1beta1.MsgVote",
		fmt.Sprint("proposal_vote.proposal_id={}", proposalID),
		fmt.Sprint("proposal_vote.option={}", voteOption),
	}

	query := strings.Join(events, " AND ")
	resTxs, err := node.TxSearch(context.Background(), query, true, &page, &limit, "")

	resBlocks, err := getBlocksForTxResults(node, resTxs.Txs)
	if err != nil {
		return nil, err
	}

	return formatTxResults(encCfg.TxConfig, resTxs.Txs, resBlocks)
}

func getBlocksForTxResults(node rpcclient.Client, resTxs []*ctypes.ResultTx) (map[int64]*ctypes.ResultBlock, error) {
	resBlocks := make(map[int64]*ctypes.ResultBlock)

	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := node.Block(context.Background(), &resTx.Height)
			if err != nil {
				return nil, err
			}

			resBlocks[resTx.Height] = resBlock
		}
	}

	return resBlocks, nil
}

func mkTxResult(txConfig client.TxConfig, resTx *ctypes.ResultTx, resBlock *ctypes.ResultBlock) (*sdk.TxResponse, error) {
	txb, err := txConfig.TxDecoder()(resTx.Tx)
	if err != nil {
		return nil, err
	}
	p, ok := txb.(intoAny)
	if !ok {
		return nil, fmt.Errorf("expecting a type implementing intoAny, got: %T", txb)
	}
	any := p.AsAny()
	return sdk.NewResponseResultTx(resTx, any, resBlock.Block.Time.Format(time.RFC3339)), nil
}

type intoAny interface {
	AsAny() *codectypes.Any
}

// formatTxResults parses the indexed txs into a slice of TxResponse objects.
func formatTxResults(txConfig client.TxConfig, resTxs []*ctypes.ResultTx, resBlocks map[int64]*ctypes.ResultBlock) ([]*sdk.TxResponse, error) {
	var err error
	out := make([]*sdk.TxResponse, len(resTxs))
	for i := range resTxs {
		out[i], err = mkTxResult(txConfig, resTxs[i], resBlocks[resTxs[i].Height])
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}
