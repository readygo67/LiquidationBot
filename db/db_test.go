package dbm

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
	"os"
	"testing"
)

func TestAccessDB(t *testing.T) {
	db, err := InitDB("testdb1")
	require.NoError(t, err)

	defer db.Close()
	defer os.RemoveAll("testdb1")

	db.Put(LastHandledHeightStoreKey(), big.NewInt(12561332).Bytes(), nil)
	bz, _ := db.Get(LastHandledHeightStoreKey(), nil)
	height := big.NewInt(0).SetBytes(bz)
	require.EqualValues(t, 12561332, height.Uint64())

	for i := 0; i < 10; i++ {
		db.Put(BorrowersStoreKey([]byte(fmt.Sprintf("account%v", i))), []byte(fmt.Sprintf("account%v", i)), nil)
	}

	iter0 := db.NewIterator(util.BytesPrefix(BorrowersPrefix), nil)
	defer iter0.Release()
	t.Logf("borrows address")
	for iter0.Next() {
		fmt.Printf("%v\n", string(iter0.Value()))
	}

	for i := 10; i < 20; i++ {
		db.Put(LiquidationBelow1StoreKey([]byte(fmt.Sprintf("account%v", i))), []byte(fmt.Sprintf("account%v", i)), nil)
	}

	iter1 := db.NewIterator(util.BytesPrefix(LiquidationBelow1Prefix), nil)
	defer iter1.Release()
	t.Logf("liquidation below 1 address")
	for iter1.Next() {
		fmt.Printf("%v\n", string(iter1.Value()))
	}
}
