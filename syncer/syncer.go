package syncer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/readygo67/LiquidationBot/venus"
	"strings"

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
const ScanSpan = 10000

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
	db := s.db
	c := s.c
	ctx := context.Background()

	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")
	vbep20Abi, _ := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	var logs []types.Log
	var addresses []common.Address
	name := make(map[string]string)
	for _, token := range s.tokens {
		addresses = append(addresses, common.HexToAddress(token.Address))
		name[strings.ToLower(token.Address)] = token.Name
	}

	t := time.NewTimer(0)
	defer t.Stop()

	for {
		select {
		case <-s.quitCh:
			return

		case <-t.C:
			currentHeight, err := c.BlockNumber(ctx)
			if err != nil {
				t.Reset(time.Second * 3)
				continue
			}

			bz, err := db.Get(dbm.LastHandledHeightStoreKey(), nil)
			if err != nil {
				t.Reset(time.Millisecond * 20)
				continue
			}
			lastHandledHeight := big.NewInt(0).SetBytes(bz).Uint64()

			startHeight := lastHandledHeight + 1
			endHeight := currentHeight
			if startHeight+ConfirmHeight >= currentHeight {
				t.Reset(time.Second * 3)
				continue
			}

			if currentHeight-lastHandledHeight >= ScanSpan {
				endHeight = startHeight + ScanSpan - 1
			}
			fmt.Printf("startHeight:%v, endHeight:%v\n", startHeight, endHeight)
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(startHeight)),
				ToBlock:   big.NewInt(int64(endHeight)),
				Addresses: addresses,
				Topics:    [][]common.Hash{{topicBorrow}},
			}

			logs, err = c.FilterLogs(context.Background(), query)
			if err != nil {
				goto EndWithoutUpdateHeight
			}

			for i, log := range logs {
				var borrowEvent venus.Vbep20Borrow
				err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", log.Data)
				fmt.Printf("%v height:%v, name:%v borrower:%v\n", (i + 1), log.BlockNumber, name[strings.ToLower(log.Address.String())], borrowEvent.Borrower)

				borrowerBytes := borrowEvent.Borrower.Bytes()
				exist, err := db.Has(dbm.BorrowersStoreKey(borrowerBytes), nil)
				if err != nil {
					goto EndWithoutUpdateHeight
				}

				if exist {
					continue
				}

				byteCode, err := c.CodeAt(ctx, log.Address, big.NewInt(int64(log.BlockNumber)))
				if len(byteCode) > 0 {
					//a smart contract
					continue
				}

				err = db.Put(dbm.BorrowersStoreKey(borrowerBytes), borrowerBytes, nil)
				if err != nil {
					goto EndWithoutUpdateHeight
				}

				bz, err := db.Get(dbm.KeyBorrowerNumber, nil)
				num := big.NewInt(0).SetBytes(bz).Int64()
				if err != nil {
					goto EndWithoutUpdateHeight
				}

				num += 1
				err = db.Put(dbm.KeyBorrowerNumber, big.NewInt(num).Bytes(), nil)
				if err != nil {
					goto EndWithoutUpdateHeight
				}
			}

			err = db.Put(dbm.LastHandledHeightStoreKey(), big.NewInt(int64(endHeight)).Bytes(), nil)
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
