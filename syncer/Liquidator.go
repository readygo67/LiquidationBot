package syncer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/readygo67/LiquidationBot/config"
	"github.com/syndtr/goleveldb/leveldb"
	"sync"
)

type Liquidator struct {
	c             *ethclient.Client
	db            *leveldb.DB
	comptroller   string
	oracle        string
	wg            sync.WaitGroup
	quitCh        chan struct{}
	liquidationCh chan common.Address
}

func NewLiquidator(c *ethclient.Client, db *leveldb.DB, cfg *config.Config) *Liquidator {
	return &Liquidator{
		c:             c,
		db:            db,
		comptroller:   cfg.Comptroller,
		oracle:        cfg.Oracle,
		quitCh:        make(chan struct{}),
		liquidationCh: make(chan common.Address, 64),
	}
}

func (liq *Liquidator) Start() {
	liq.wg.Add(1)
	log.Info("syncer start")
	fmt.Println("syncer start")
	go liq.Run()
}

func (liq *Liquidator) Stop() {
	close(liq.quitCh)
	liq.wg.Wait()
}

func (liq *Liquidator) Run() {
	defer liq.wg.Done()
	for {
		select {
		case <-liq.quitCh:
			return
		case account := <-liq.liquidationCh:
			fmt.Printf("liqudate:%v", account)
			//liquidate(account)
		}
	}
}
