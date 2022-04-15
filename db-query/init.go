package dbquery

import (
	"github.com/notional-labs/gaia-analyzer/db-query/app"
	"github.com/notional-labs/gaia-analyzer/db-query/tx"
)

func Init(rootDir string) {
	tx.InitTxIndexer(rootDir)
	app.InitCommitMultiStoreAndApp(rootDir)
	// app.InitApp()
}
