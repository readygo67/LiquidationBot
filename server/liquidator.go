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
	c           *ethclient.Client
	db          *leveldb.DB
	comptroller *venus.Comptroller
	oracle      *venus.Oracle

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

		case pending := <-liq.priorityLiquidationCh:
			fmt.Printf("receive priority liquidation:%v\n", pending)
			pending.Endtime = time.Now().Add(time.Second * PriorityLiquidationTime)
			priorityPendings = append(priorityPendings, pending)

		case pending := <-liq.liquidationCh:
			fmt.Printf("recive liquidation:%v\n", pending)
			pending.Endtime = time.Now().Add(time.Second * NormalLiquidationTime)
			pendings = append(pendings, pending)

		case <-t.C:
			num := len(pendings) + len(priorityPendings)
			fmt.Printf("%vth liquidator, %vpending liquidations\n", count, num)
			count++

			var priorityRemains []*Liquidation
			for _, pending := range priorityPendings {
				if pending.Endtime.After(time.Now()) {
					priorityRemains = append(priorityRemains, pending)
				}
			}
			priorityPendings = priorityRemains
			for _, pending := range priorityPendings {
				fmt.Printf("verify priority liquidation:%v\n", pending)
				//TODO(keep)
			}

			var remains []*Liquidation
			for _, pending := range pendings {
				if pending.Endtime.After(time.Now()) {
					remains = append(remains, pending)
				}
			}

			pendings = remains
			for _, pending := range pendings {
				fmt.Printf("verify pending liquidation:%v\n", pending)
				//TODO(keep)
			}

			t.Reset(time.Second * 3)
		}
	}
}

//
//func (liq *Liquidator) liquidate(liquidation *Liquidation) error {
//	c := liq.c
//	comptroller := liq.comptroller
//	oracle := liq.oracle
//
//	account := liquidation.Address
//
//	errCode, liquidity, shortfall, err := comptroller.GetAccountLiquidity(nil, liquidation.Address)
//	if errCode != nil {
//		fmt.Printf("liquidate, errCode:%v, liquidation:%v\n", errCode, liquidation)
//		return fmt.Errorf("errCode:%v", errCode)
//	}
//	if err != nil {
//		fmt.Printf("liquidate, err:%v, liquidation:%v\n", err, liquidation)
//		return err
//	}
//	if liquidity.Cmp(BigIntZero) == 1 && shortfall.Cmp(BigIntZero) != 1 {
//		fmt.Printf("liquidate, local calculation error, liquidity:%v, shortfall:%v, liqiudation:%v\n", liquidity, shortfall, liquidation)
//		return fmt.Errorf("liquidate, local calculation error, liquidity:%v, shortfall:%v, liqiudation:%v\n", liquidity, shortfall, liquidation)
//	}
//
//	markets, err := comptroller.GetAssetsIn(nil, account)
//	if err != nil {
//		fmt.Printf("liquidator, fail to get GetAssetsIn, err:%v\n", err)
//		return err
//	}
//
//	for _, market := range markets {
//		vbep20, err := venus.NewVbep20(market, c)
//		_, balance, borrow, exchangeRate, err := vbep20.GetAccountSnapshot(nil, account)
//		if err != nil {
//			fmt.Printf("liquidator, fail to get GetAccountSnapshot, err:%v\n", err)
//			return err
//		}
//	}
//
//	return nil
//}
