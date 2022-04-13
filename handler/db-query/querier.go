package dbquery

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmquery "github.com/tendermint/tendermint/libs/pubsub/query"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/state/txindex/kv"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

// db, err := sdk.NewLevelDB("dig_leveldb_testing", dir)
func openDB(rootDir string) (dbm.DB, error) {
	dataDir := filepath.Join(rootDir, "data")
	a, err := sdk.NewLevelDB("tx_index", dataDir)
	return a, err
}

func QueryBankSendAtomFromAddress(ctx client.Context, rootDir string, sender string) []*sdk.TxResponse {
	senderEvent := fmt.Sprintf("%s='%s'", "message.sender", sender)
	bankSendEvent := fmt.Sprintf("%s='%s'", "message.action", "/cosmos.bank.v1beta1.MsgSend")

	tmEvents := []string{senderEvent, bankSendEvent}

	page := 1
	limit := 10

}

func QueryGovTxs(ctx client.Context, rootDir string, tmEvents []string) []*sdk.TxResponse {
	// var tmEvents = []string{
	// 	"message.action='/cosmos.gov.v1beta1.MsgVote'",
	// 	fmt.Sprintf("proposal_vote.proposal_id='%d'", proposalID),
	// }
	//tm events like that

	store, err := openDB(rootDir)

	if err != nil {
		panic(err)
	}

	query := strings.Join(tmEvents, " AND ")

	txi := kv.NewTxIndex(
		store,
	)
	q, err := tmquery.New(query)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		return nil
	}

	results, err := txi.Search(context.Background(), q)

	resTxs := make([]*ctypes.ResultTx, 0, len(results))

	for _, r := range results {
		var proof tmtypes.TxProof

		resTxs = append(resTxs, &ctypes.ResultTx{
			Hash:     tmtypes.Tx(r.Tx).Hash(),
			Height:   r.Height,
			Index:    r.Index,
			TxResult: r.Result,
			Tx:       r.Tx,
			Proof:    proof,
		})
	}

	resBlocks, err := getBlocksForTxResults(ctx, resTxs)
	if err != nil {
		return nil
	}

	txs, err := formatTxResults(ctx.TxConfig, resTxs, resBlocks)
	if err != nil {
		return nil
	}

	return txs
}
