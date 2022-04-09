package blockexplore

import (
	"github.com/google/orderedcode"
	dbm "github.com/tendermint/tm-db"
)

const (
	prefixBlockPart = int64(1)
)

func blockPartKey(height int64, partIndex int, db dbm.DB) []byte {
	key, err := orderedcode.Append(nil, prefixBlockPart, height, int64(partIndex))
	if err != nil {
		panic(err)
	}
	return key
}
