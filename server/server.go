package server

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const DefaultStartHeigt = uint64(13000000)

//for feedPrice debug
func Start(cfg *config.Config) error {
	fmt.Printf("start for feedPrice test\n")
	cfg, err := config.New("../config.yml")
	rpcURL := "ws://192.168.88.144:28546"
	c, _ := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)
	account := common.HexToAddress("0xDfBe18F35cD3FC6B9CBd3B643b110889635b1Ee9") //0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B
	accountBytes := account.Bytes()

	sync.syncOneAccount(account)

	exist, err := db.Has(dbm.AccountStoreKey(accountBytes), nil)
	if err != nil {
		panic(err)
	}
	if !exist {
		panic(exist)
	}

	bz, err := db.Get(dbm.BorrowersStoreKey(accountBytes), nil)
	if err != nil {
		panic(err)
	}
	if account != common.BytesToAddress(bz) {
		panic("account mismatch")
	}

	accounts := []common.Address{}
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	if len(accounts) != 1 {
		panic("account number mismatch")
	}

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	var info AccountInfo
	err = json.Unmarshal(bz, &info)
	fmt.Printf("info:%+v\n", info.toReadable())

	for _, asset := range info.Assets {
		symbol := asset.Symbol

		bz, err = db.Get(dbm.MarketStoreKey([]byte(symbol), accountBytes), nil)
		if err != nil {
			panic(err)
		}
		if account != common.BytesToAddress(bz) {
			panic("account mismatch")
		}

		prefix := append(dbm.MarketPrefix, []byte(symbol)...)
		accounts = []common.Address{}
		iter = db.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			accounts = append(accounts, common.BytesToAddress(iter.Value()))
		}

		if len(accounts) != 1 {
			panic("account number mismatch")
		}
	}
	key := dbm.LiquidationBelow1P1StoreKey(accountBytes)
	bz, err = db.Get(key, nil)
	if err != nil {
		panic(err)
	}
	if account != common.BytesToAddress(bz) {
		panic("account mismatch")
	}

	//height, err := sync.c.BlockNumber(context.Background())
	//require.NoError(t, err)
	fmt.Printf("enter loop from nowon")
	sync.wg.Add(5)
	go sync.syncMarketsAndPrices()
	go sync.syncLiquidationBelow1P1()
	go sync.monitorTxPool()
	go sync.feedPrices()
	go sync.printConcernedAccountInfo()

	waitExit()

	close(sync.quitCh)
	sync.wg.Wait()

	return nil
}

//
//func Start(cfg *config.Config) error {
//	client, err := ethclient.Dial(cfg.RPCURL)
//	if err != nil {
//		return err
//	}
//
//	db, err := dbm.NewDB(cfg.DB)
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	startHeight := DefaultStartHeigt
//	var storedHeight uint64
//	exist, err := db.Has(dbm.LastHandledHeightStoreKey(), nil)
//	if exist {
//		bz, err := db.Get(dbm.LastHandledHeightStoreKey(), nil)
//		if err != nil {
//			return err
//		}
//		storedHeight = big.NewInt(0).SetBytes(bz).Uint64()
//		startHeight = storedHeight
//	}
//	fmt.Printf("startHeight:%v, storedHeight:%v, configHeight:%v\n", startHeight, storedHeight, cfg.StartHeihgt)
//	if cfg.Override {
//		startHeight = cfg.StartHeihgt
//	}
//	err = db.Put(dbm.LastHandledHeightStoreKey(), big.NewInt(0).SetUint64(startHeight).Bytes(), nil)
//	if err != nil {
//		panic(err)
//	}
//
//	liquidationCh := make(chan *Liquidation, 64)
//	priorityliquidationCh := make(chan *Liquidation, 64)
//	feededPricesCh := make(chan *FeededPrices, 64)
//
//	syncer := NewSyncer(client, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)
//
//	syncer.Start()
//
//	waitExit()
//
//	syncer.Stop()
//	return nil
//}

func waitExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	i := <-c
	log.Printf("Received interrupt[%v], shutting down...\n", i)
}
