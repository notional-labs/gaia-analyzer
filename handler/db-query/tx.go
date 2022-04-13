package dbquery

import (
	"context"
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

func QueryTxs(ctx client.Context, rootDir string, tmEvents []string) []*ctypes.ResultTx {
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
		panic(err)
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
	return resTxs
}
