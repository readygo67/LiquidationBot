package server

import (
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

	var startHeight uint64
	exist, err := db.Has(dbm.LastHandledHeightStoreKey(), nil)
	if !exist {
		startHeight = DefaultStartHeigt
	} else {
		bz, err := db.Get(dbm.LastHandledHeightStoreKey(), nil)
		if err != nil {
			return err
		}
		startHeight = big.NewInt(0).SetBytes(bz).Uint64()
	}

	if cfg.Override {
		startHeight = cfg.StartHeihgt
	}

	err = db.Put(dbm.LastHandledHeightStoreKey(), big.NewInt(0).SetUint64(startHeight).Bytes(), nil)
	if err != nil {
		panic(err)
	}

	if cfg.Override {
		startHeight = cfg.StartHeihgt
	}

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	syncer := NewSyncer(client, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
	liquidator := NewLiquidator(client, db, cfg.Comptroller, cfg.Oracle, liquidationCh, priorityliquidationCh)

	syncer.Start()
	liquidator.Start()

	waitExit()
	syncer.Stop()
	liquidator.Stop()
	return nil
}

func waitExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	i := <-c
	log.Printf("Received interrupt[%v], shutting down...\n", i)
}
