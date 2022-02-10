package server

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
)

const DefaultStartHeigt = uint64(13000000)

func Start(cfg *config.Config) error {
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return err
	}

	db, err := dbm.NewDB(cfg.DB)
	if err != nil {
		return err
	}
	defer db.Close()

	startHeight := DefaultStartHeigt
	var storedHeight uint64
	exist, err := db.Has(dbm.LastHandledHeightStoreKey(), nil)
	if exist {
		bz, err := db.Get(dbm.LastHandledHeightStoreKey(), nil)
		if err != nil {
			return err
		}
		storedHeight = big.NewInt(0).SetBytes(bz).Uint64()
		startHeight = storedHeight
	}
	fmt.Printf("startHeight:%v, storedHeight:%v, configHeight:%v\n", startHeight, storedHeight, cfg.StartHeihgt)
	if cfg.Override {
		startHeight = cfg.StartHeihgt
	}
	err = db.Put(dbm.LastHandledHeightStoreKey(), big.NewInt(0).SetUint64(startHeight).Bytes(), nil)
	if err != nil {
		panic(err)
	}

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	syncer := NewSyncer(client, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, feededPricesCh, liquidationCh, priorityliquidationCh)

	syncer.Start()

	waitExit()

	syncer.Stop()
	return nil
}

func waitExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	i := <-c
	log.Printf("Received interrupt[%v], shutting down...\n", i)
}
