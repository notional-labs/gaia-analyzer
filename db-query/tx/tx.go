package tx

import (
	"context"
	"path/filepath"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmquery "github.com/tendermint/tendermint/libs/pubsub/query"
	"github.com/tendermint/tendermint/state/txindex/kv"
	dbm "github.com/tendermint/tm-db"
)

var (
	TxIndexer *kv.TxIndex
)

// db, err := sdk.NewLevelDB("dig_leveldb_testing", dir)
func OpenTxDB(rootDir string) (dbm.DB, error) {
	dataDir := filepath.Join(rootDir, "data")
	a, err := sdk.NewLevelDB("tx_index", dataDir)
	return a, err
}

func InitTxIndexer(rootDir string) {

	store, err := OpenTxDB(rootDir)

	if err != nil {
		panic(err)
	}

	TxIndexer = kv.NewTxIndex(
		store,
	)
}

func QueryTxs(tmEvents []string) []*abcitypes.TxResult {
	// var tmEvents = []string{
	// 	"message.action='/cosmos.gov.v1beta1.MsgVote'",
	// 	fmt.Sprintf("proposal_vote.proposal_id='%d'", proposalID),
	// }
	//tm events like that

	query := strings.Join(tmEvents, " AND ")

	q, err := tmquery.New(query)
	if err != nil {
		panic(err)
	}

	results, err := TxIndexer.Search(context.Background(), q)
	if err != nil {
		panic(err)
	}

	return results
}

// // query txs event and push to global tx queue
// func QueryTxsAndTrackTxEvents(tmEvents []string) {
// 	// var tmEvents = []string{
// 	// 	"message.action='/cosmos.gov.v1beta1.MsgVote'",
// 	// 	fmt.Sprintf("proposal_vote.proposal_id='%d'", proposalID),
// 	// }
// 	//tm events like that
// 	query := strings.Join(tmEvents, " AND ")

// 	q, err := tmquery.New(query)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err != nil {
// 		panic(err)
// 	}

// 	results, err := TxIndexer.Search(context.Background(), q)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, r := range results {
// 		txItem := types.TxItem{
// 			Height: r.Height,
// 			Events: &r.Result.Events,
// 		}
// 		heap.Push(&data.TrackedTxQueue, txItem)
// 	}
// }
