package server

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/readygo67/LiquidationBot/venus"
	"github.com/syndtr/goleveldb/leveldb"
	"sync"
	"time"
)

const (
	NormalLiquidationTime   = 18  // in secs
	PriorityLiquidationTime = 180 //in secs
)

type Liquidator struct {
	c                     *ethclient.Client
	db                    *leveldb.DB
	comptroller           *venus.Comptroller
	oracle                *venus.Oracle
	wg                    sync.WaitGroup
	quitCh                chan struct{}
	liquidationCh         chan *Liquidation
	priorityLiquidationCh chan *Liquidation
}

func NewLiquidator(c *ethclient.Client, db *leveldb.DB, comptrollerAddress string, oracleAddress string, liquidationCh chan *Liquidation, priorityLiquidationCh chan *Liquidation) *Liquidator {
	comptroller, err := venus.NewComptroller(common.HexToAddress(comptrollerAddress), c)
	if err != nil {
		panic(err)
	}

	oracle, err := venus.NewOracle(common.HexToAddress(oracleAddress), c)
	if err != nil {
		panic(err)
	}

	return &Liquidator{
		c:                     c,
		db:                    db,
		comptroller:           comptroller,
		oracle:                oracle,
		quitCh:                make(chan struct{}),
		liquidationCh:         liquidationCh,
		priorityLiquidationCh: priorityLiquidationCh,
	}
}

func (liq *Liquidator) Start() {
	liq.wg.Add(1)
	log.Info("server start")
	fmt.Println("server start")
	go liq.Run()
}

func (liq *Liquidator) Stop() {
	close(liq.quitCh)
	liq.wg.Wait()
}

func (liq *Liquidator) Run() {
	defer liq.wg.Done()

	var pendings []*Liquidation
	var priorityPendings []*Liquidation
	count := 1

	t := time.NewTimer(0)
	defer t.Stop()

	for {
		select {
		case <-liq.quitCh:
			return
		case pending := <-liq.liquidationCh:
			fmt.Printf("recive liquidation:%v\n", pending)
			pending.Endtime = time.Now().Add(time.Second * NormalLiquidationTime)
			pendings = append(pendings, pending)

		case pending := <-liq.priorityLiquidationCh:
			fmt.Printf("receive priority liquidation:%v\n", pending)
			pending.Endtime = time.Now().Add(time.Second * PriorityLiquidationTime)
			priorityPendings = append(priorityPendings, pending)

		case <-t.C:
			num := len(pendings) + len(priorityPendings)
			fmt.Printf("%vth liquidator, %vpending liquidations\n", count, num)
			count++
			for i, pending := range priorityPendings {
				if pending.Endtime.Before(time.Now()) {
					priorityPendings = append(priorityPendings[:i], priorityPendings[i+1:]...)
					continue
				}
				fmt.Printf("verify priority liquidation:%v", pending)
				//TODO(keep)
			}

			for i, pending := range pendings {
				if pending.Endtime.Before(time.Now()) {
					pendings = append(pendings[:i], pendings[i+1:]...)
					continue
				}
				fmt.Printf("verify priority liquidation:%v", pending)
				//TODO(keep)
			}
			t.Reset(time.Second * 3)
		}
	}
}
