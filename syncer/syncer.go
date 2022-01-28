package syncer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/readygo67/LiquidationBot/config"
	"github.com/syndtr/goleveldb/leveldb"
	"math/big"
	"sync"
	"time"
)

type Syncer struct {
	c      *ethclient.Client
	db     *leveldb.DB
	wg     sync.WaitGroup
	tokens []config.Token
	quitCh chan struct{}
}

func NewSyncer(c *ethclient.Client, db *leveldb.DB, cfg config.Config) *Syncer {
	return &Syncer{
		c:      c,
		db:     db,
		tokens: cfg.Tokens,
		quitCh: make(chan struct{}),
	}
}

func (m *Syncer) Start() {
	log.Info("monitor start")
	m.wg.Add(1)
	go m.scanning()
}

func (m *Syncer) Stop() {
	close(m.quitCh)
	m.wg.Wait()
}

func (m *Syncer) scanning() {
	defer m.wg.Done()

	t := time.NewTimer(0)
	defer t.Stop()

	db := m.db
	c := m.c
	ctx := context.Background()
	for {
		select {
		case <-m.quitCh:
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

			fmt.Println("scanning", "height to be handled", height, "currentHeight", currentHeight)
			log.Info("scanning", "height to be handled", height, "currentHeight", currentHeight)
			blk, err := c.BlockByNumber(ctx, height)
			if err != nil {
				log.Info("fail to get block by number", "height", height, "error", err)
				t.Reset(time.Second * 3)
				continue
			}

			if blk == nil {
				fmt.Println("getting block is empty", "height", height)
				log.Info("getting block is empty", "height", height)
				t.Reset(time.Second * 3)
				continue
			}

			if len(blk.Transactions()) != 0 {
				for _, tx := range blk.Transactions() {
					if tx.To() == nil { //contract creation
						continue
					}
					if tx.To().String() == contract {
						receipt, err := c.TransactionReceipt(ctx, tx.Hash())
						if err != nil {
							log.Info("fail to get TransactionReceipt", "hash", tx.Hash(), "error", err)
							goto EndWithoutUpdateHeight
						}

						for _, receiptlog := range receipt.Logs {
							if receiptlog.Removed {
								continue
							}

							if receiptlog.Address.String() == contract {
								if receiptlog.Topics[0].String() == topic0 && receiptlog.Topics[1].String() == topic1 {
									recipient := ethcmn.BytesToAddress(receiptlog.Topics[2].Bytes())
									id := big.NewInt(0).SetBytes(receiptlog.Topics[3].Bytes()).Uint64()
									log.Info("mint ", "receipt", recipient, "id", id)

									exist, err := db.Has(dbm.HandledHashAndIDStoreKey(tx.Hash(), uint(id)), nil)
									if err != nil {
										log.Info("fail to get handled hash and id ", "hash", tx.Hash(), "id", id, "error", err)
										goto EndWithoutUpdateHeight
									}

									if exist {
										continue
									}

									seed := crypto.Keccak256Hash(recipient.Bytes(), tx.GasPrice().Bytes(), blk.Number().Bytes(), big.NewInt(int64(blk.Time())).Bytes(), blk.ParentHash().Bytes(), big.NewInt(int64(id)).Bytes())
									seedUint64 := seed.Big().Uint64()

									err = dbm.AllocateOneWithCheck(db, tx.Hash(), uint(id), recipient, uint(seedUint64))
									if err != nil {
										log.Error("fail to AllocateOneWithCheck", "hash", tx.Hash(), "id", id, "recipient", recipient, "seed", seedUint64, "err", err)
										goto EndWithoutUpdateHeight
									}
									log.Info("success AllocateOneWithCheck", "hash", tx.Hash(), "id", id, "recipient", recipient, "seed", seedUint64)
									fmt.Println("success AllocateOneWithCheck", "hash", tx.Hash(), "id", id, "recipient", recipient, "seed", seedUint64)
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
			t.Reset(time.Millisecond * 50)
		}
	}
}
