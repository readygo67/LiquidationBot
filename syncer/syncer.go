package syncer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/readygo67/LiquidationBot/config"
	"github.com/readygo67/LiquidationBot/db"
	"github.com/syndtr/goleveldb/leveldb"
	"math/big"
	"sync"
	"time"
)

const ConfirmHeight = 0

type Syncer struct {
	c      *ethclient.Client
	db     *leveldb.DB
	wg     sync.WaitGroup
	tokens []config.Token
	quitCh chan struct{}
}

func NewSyncer(c *ethclient.Client, db *leveldb.DB, cfg *config.Config) *Syncer {
	return &Syncer{
		c:      c,
		db:     db,
		tokens: cfg.Tokens,
		quitCh: make(chan struct{}),
	}
}

func (s *Syncer) Start() {
	log.Info("syncer start")
	fmt.Println("syncer start")
	s.wg.Add(2)
	go s.scanForAllBorrowers()
	go s.scanLiquidationBelow1P2()
}

func (s *Syncer) Stop() {
	close(s.quitCh)
	s.wg.Wait()
}

func (s *Syncer) scanForAllBorrowers() {
	defer s.wg.Done()

	t := time.NewTimer(0)
	defer t.Stop()

	db := s.db
	c := s.c
	ctx := context.Background()
	topicBorrow := "0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80"

	for {
		select {
		case <-s.quitCh:
			return

		case <-t.C:
			bz, err := db.Get(dbm.LastHandledHeightStoreKey(), nil)
			if err != nil {
				t.Reset(time.Second * 3)
				continue
			}

			currentHeight, err := c.BlockNumber(ctx)
			if err != nil {
				t.Reset(time.Second * 3)
				continue
			}

			lastHandledHeight := big.NewInt(0).SetBytes(bz)
			height := big.NewInt(0).Add(lastHandledHeight, big.NewInt(1))
			if height.Uint64()+ConfirmHeight >= currentHeight {
				t.Reset(time.Second * 3)
				continue
			}

			log.Info("scanForAllBorrowers", "height to be handled", height, "currentHeight", currentHeight)
			fmt.Printf("scanForAllBorrowers, height:%v, currentHeight:%v\n", height, currentHeight)
			blk, err := c.BlockByNumber(ctx, height)
			if err != nil {
				log.Info("fail to get block by number", "height", height, "error", err)
				t.Reset(time.Second * 3)
				continue
			}

			if blk == nil {
				log.Info("getting block is empty", "height", height)
				t.Reset(time.Second * 3)
				continue
			}

			if len(blk.Transactions()) != 0 {
				for _, tx := range blk.Transactions() {

					if tx.To() == nil { //contract creation
						continue
					}

					for _, contract := range s.tokens {
						if tx.To().String() == contract.Address {
							receipt, err := c.TransactionReceipt(ctx, tx.Hash())
							if err != nil {
								log.Info("fail to get TransactionReceipt", "hash", tx.Hash(), "error", err)
								goto EndWithoutUpdateHeight
							}

							for _, receiptlog := range receipt.Logs {
								if receiptlog.Removed {
									continue
								}

								if receiptlog.Address.String() == contract.Address {
									if receiptlog.Topics[0].String() == topicBorrow {
										borrower := receiptlog.Data[12:32]
										//log.Info("height:%v, contract:%v, borrower:%+v\n", height, contract.Name, common.BytesToAddress(borrower))
										fmt.Printf("height:%v, contract:%v, borrower:%+v\n", height, contract.Name, common.BytesToAddress(borrower))
										exist, err := db.Has(dbm.BorrowersStoreKey(borrower), nil)
										if err != nil {
											goto EndWithoutUpdateHeight
										}

										if exist {
											continue
										}

										err = db.Put(dbm.BorrowersStoreKey(borrower), borrower, nil)
										if err != nil {
											goto EndWithoutUpdateHeight
										}

										bz, err := db.Get(dbm.KeyBorrowerNumber, nil)
										num := big.NewInt(0).SetBytes(bz).Int64()
										fmt.Printf("borrower number:%v\n", num)
										if err != nil {
											goto EndWithoutUpdateHeight
										}

										num += 1
										err = db.Put(dbm.KeyBorrowerNumber, big.NewInt(num).Bytes(), nil)
										if err != nil {
											goto EndWithoutUpdateHeight
										}
									}
								}
							}
						}
					}
				}
			}

			lastHandledHeight = big.NewInt(0).Add(lastHandledHeight, big.NewInt(1))
			err = db.Put(dbm.LastHandledHeightStoreKey(), lastHandledHeight.Bytes(), nil)
			if err != nil {
				goto EndWithoutUpdateHeight
			}

		EndWithoutUpdateHeight:
			t.Reset(time.Millisecond * 20)
		}
	}
}

func (s *Syncer) scanLiquidationBelow1P2() {
	defer s.wg.Done()

	t := time.NewTimer(0)
	defer t.Stop()

	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			t.Reset(time.Second * 50)
		}
	}
}
