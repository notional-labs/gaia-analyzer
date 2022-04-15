package app

import (
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

var (
	// app that not contain any data, used only to utilize the logic defined in it
	EmptyApp *simapp.SimApp
	// multi store taken from application.db
	Cms sdk.CommitMultiStore
	// Used to access to application state at height
	ContextAtHeight map[int64]*sdk.Context = map[int64]*sdk.Context{}
)

func InitApp() {
	encCfg := simapp.MakeTestEncodingConfig()
	db := dbm.NewMemDB()
	EmptyApp = simapp.NewSimApp(nil, db, nil, true, map[int64]bool{}, "", 0, encCfg, simapp.EmptyAppOptions{})
}

func InitCommitMultiStoreAndApp(rootDir string) {
	appDB, err := OpenAppDB(rootDir)
	if err != nil {
		panic(err)
	}

	encCfg := simapp.MakeTestEncodingConfig()

	EmptyApp = simapp.NewSimApp(nil, appDB, nil, true, map[int64]bool{}, "", 0, encCfg, simapp.EmptyAppOptions{})

	Cms = EmptyApp.CMS()
	// fmt.Printf("%+v\n", EmptyApp.BaseApp.CMS())
}

func OpenAppDB(rootDir string) (dbm.DB, error) {
	dataDir := filepath.Join(rootDir, "data")
	a, err := sdk.NewLevelDB("application", dataDir)
	return a, err
}

func GetQueryContext(height int64) *sdk.Context {
	ctx, ok := ContextAtHeight[height]
	if ok {
		return ctx
	}

	cacheMS, err := Cms.CacheMultiStoreWithVersion(height)
	if err != nil {
		panic(err)
	}

	ctx = sdk.NewRefContext(
		cacheMS, tmproto.Header{}, true, nil,
	)
	ContextAtHeight[height] = ctx
	return ctx
}
