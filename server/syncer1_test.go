package server

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb/util"
	"os"
	"testing"
)

func TestSyncOneAccountWithFeededPrices2(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)
	account := common.HexToAddress("0xDfBe18F35cD3FC6B9CBd3B643b110889635b1Ee9") //0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B
	accountBytes := account.Bytes()
	err = sync.syncOneAccount(account)
	require.NoError(t, err)

	exist, err := db.Has(dbm.AccountStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.True(t, exist)

	bz, err := db.Get(dbm.BorrowersStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	accounts := []common.Address{}
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	var info AccountInfo
	err = json.Unmarshal(bz, &info)
	t.Logf("info:%+v\n", info.toReadable())

	for _, asset := range info.Assets {
		symbol := asset.Symbol

		bz, err = db.Get(dbm.MarketStoreKey([]byte(symbol), accountBytes), nil)
		require.NoError(t, err)
		require.Equal(t, account, common.BytesToAddress(bz))

		prefix := append(dbm.MarketPrefix, []byte(symbol)...)
		accounts = []common.Address{}
		iter = db.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			accounts = append(accounts, common.BytesToAddress(iter.Value()))
		}
		require.Equal(t, 1, len(accounts))
	}

	key := getLiquidationKey(info.MaxLoanValue, info.HealthFactor, accountBytes)
	bz, err = db.Get(key, nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	//height, err := sync.c.BlockNumber(context.Background())
	//require.NoError(t, err)

	sync.wg.Add(5)
	go sync.syncMarketsAndPrices()
	go sync.syncLiquidationBelow1P1()
	go sync.monitorTxPool()
	go sync.feedPrices()
	go sync.printConcernedAccountInfo()

	waitExit()
	
	close(sync.quitCh)
	sync.wg.Wait()

}
