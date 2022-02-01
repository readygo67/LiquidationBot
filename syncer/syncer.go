package syncer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/readygo67/LiquidationBot/venus"
	"github.com/syndtr/goleveldb/leveldb/util"
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

const (
	ConfirmHeight        = 0
	ScanSpan             = 10000
	SyncIntervalBelow1P0 = 3
	SyncIntervalBelow1P2 = 9
	SyncIntervalBelow1P5 = 27
	SyncIntervalBelow2P0 = 81
	SyncIntervalBelow3P0 = 243
	SyncIntervalAbove3P0 = 729
)

type TokenInfo struct {
	Address               common.Address
	CollateralFactorFloat *big.Float
	PriceFloat            *big.Float
}

type Asset struct {
	Symbol  string
	Balance *big.Float
	Loan    *big.Float
}

type AccountInfo struct {
	HealthFactor *big.Float
	Assets       []Asset
}

type FeededPrice struct {
	Address    common.Address
	PriceFloat *big.Float
	Height     uint64
	Hash       common.Hash
}

var (
	ExpScaleFloat, _    = big.NewFloat(0).SetString("1000000000000000000")
	VTokenScaleFloat, _ = big.NewFloat(0).SetString("100000000")
	BigFloat1P0, _      = big.NewFloat(0).SetString("1.0")
	BigFloat1P2, _      = big.NewFloat(0).SetString("1.2")
	BigFloat1P5, _      = big.NewFloat(0).SetString("1.5")
	BigFloat2P0, _      = big.NewFloat(0).SetString("2.0")
	BigFloat3P0, _      = big.NewFloat(0).SetString("3.0")
	AddressToSymbol     = make(map[common.Address]string)
	TokenDetail         = make(map[string]TokenInfo)
	Vbep20s             map[string]*venus.Vbep20
)

type Syncer struct {
	c                 *ethclient.Client
	db                *leveldb.DB
	comptroller       string
	oracle            string
	wg                sync.WaitGroup
	quitCh            chan struct{}
	forceUpdatePrices chan struct{}
	feededPrice       chan *FeededPrice
}

func NewSyncer(c *ethclient.Client, db *leveldb.DB, cfg *config.Config) *Syncer {
	return &Syncer{
		c:                 c,
		db:                db,
		comptroller:       cfg.Comptroller,
		oracle:            cfg.Oracle,
		quitCh:            make(chan struct{}),
		forceUpdatePrices: make(chan struct{}, 10),
		feededPrice:       make(chan *FeededPrice, 64),
	}
}

func (s *Syncer) Start() {
	log.Info("syncer start")
	fmt.Println("syncer start")
	err := s.doSyncMarketsAndPrices()
	if err != nil {
		panic(err)
	}

	s.wg.Add(9)
	go s.syncMarketsAndPrices()
	go s.feedPrices()
	go s.syncAllBorrowers()
	go s.syncLiquidationBelow1P0()
	go s.syncLiquidationBelow1P2()
	go s.syncLiquidationBelow1P5()
	go s.syncLiquidationBelow2P0()
	go s.syncLiquidationBelow3P0()
	go s.syncLiquidationAbove3P0()
}

func (s *Syncer) Stop() {
	close(s.quitCh)
	s.wg.Wait()
}

func (s *Syncer) syncMarketsAndPrices() {
	defer s.wg.Done()
	t := time.NewTimer(0)
	defer t.Stop()

	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			s.doSyncMarketsAndPrices()
		case <-s.forceUpdatePrices:
			s.doSyncMarketsAndPrices()
		}
	}
}

func (s *Syncer) doSyncMarketsAndPrices() error {
	comptroller, err := venus.NewComptroller(common.HexToAddress(s.comptroller), s.c)
	if err != nil {
		return err
	}

	oracle, err := venus.NewOracle(common.HexToAddress(s.oracle), s.c)
	if err != nil {
		return err
	}

	markets, err := comptroller.GetAllMarkets(nil)
	if err != nil {
		return err
	}

	for _, market := range markets {
		vbep20, err := venus.NewVbep20(market, s.c)
		if err != nil {
			return err
		}
		symbol, err := vbep20.Symbol(nil)
		if err != nil {
			return err
		}

		marketDetail, err := comptroller.Markets(nil, market)
		if err != nil {
			return err
		}

		price, err := oracle.GetUnderlyingPrice(nil, market)
		if err != nil {
			return err
		}

		collateralFactorFloatExp := big.NewFloat(0).SetInt(marketDetail.CollateralFactorMantissa)
		collateralFactorFloat := big.NewFloat(0).Quo(collateralFactorFloatExp, ExpScaleFloat)

		priceFloatExp := big.NewFloat(0).SetInt(price)
		priceFloat := big.NewFloat(0).Quo(priceFloatExp, ExpScaleFloat)

		tokenDetail := TokenInfo{
			Address:               market,
			CollateralFactorFloat: collateralFactorFloat,
			PriceFloat:            priceFloat,
		}

		TokenDetail[symbol] = tokenDetail
		AddressToSymbol[market] = symbol
		Vbep20s[symbol] = vbep20
	}
	return nil
}

func (s *Syncer) feedPrices() {
	defer s.wg.Done()
	for {
		select {
		case <-s.quitCh:
			return
		case feededPrice := <-s.feededPrice:
			s.doFeededPrices(feededPrice)
		}
	}
}

func (s *Syncer) doFeededPrices(feededPrice *FeededPrice) {
	//TODD)(keep)
}

func (s *Syncer) syncAllBorrowers() {
	defer s.wg.Done()
	db := s.db
	c := s.c
	ctx := context.Background()

	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")
	vbep20Abi, _ := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))

	var logs []types.Log
	var addresses []common.Address

	for _, detail := range TokenDetail {
		addresses = append(addresses, detail.Address)
	}

	comptroller, err := venus.NewComptroller(common.HexToAddress(s.comptroller), s.c)
	if err != nil {
		panic(err)
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
				fmt.Printf("%v height:%v, name:%v borrower:%v\n", (i + 1), log.BlockNumber, AddressToSymbol[log.Address], borrowEvent.Borrower)

				account := borrowEvent.Borrower
				accountBytes := account.Bytes()
				exist, err := db.Has(dbm.BorrowersStoreKey(accountBytes), nil)
				if err != nil {
					goto EndWithoutUpdateHeight
				}

				byteCode, err := c.CodeAt(ctx, log.Address, big.NewInt(int64(log.BlockNumber)))
				if len(byteCode) > 0 {
					//ignore smart contract
					continue
				}

				if !exist {
					//if account not exist in borrowers table, record it into borrowers table and increase borrowers number
					err = db.Put(dbm.BorrowersStoreKey(accountBytes), accountBytes, nil)
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

				totalCollateral := big.NewFloat(0)
				totalLoan := big.NewFloat(0)
				var assets []Asset
				mintVAIS, err := comptroller.MintedVAIs(nil, account)
				if err != nil {
					goto EndWithoutUpdateHeight
				}

				mintVAISFloatExp := big.NewFloat(0).SetInt(mintVAIS)
				mintVAISFloat := big.NewFloat(0).Quo(mintVAISFloatExp, ExpScaleFloat)

				markets, err := comptroller.GetAssetsIn(nil, account)
				if err != nil {
					goto EndWithoutUpdateHeight
				}

				for _, market := range markets {
					symbol := AddressToSymbol[market]
					_, balance, borrow, exchangeRate, err := Vbep20s[symbol].GetAccountSnapshot(nil, account)
					if err != nil {
						goto EndWithoutUpdateHeight
					}

					exchangeRateFloatExp := big.NewFloat(0).SetInt(exchangeRate)
					exchangeRateFloat := big.NewFloat(0).Quo(exchangeRateFloatExp, ExpScaleFloat)

					collateralFactorFloat := TokenDetail[symbol].CollateralFactorFloat
					priceFloat := TokenDetail[symbol].PriceFloat

					multiplier := big.NewFloat(0).Mul(exchangeRateFloat, collateralFactorFloat)
					multiplier = big.NewFloat(0).Mul(multiplier, priceFloat)

					balanceFloatExp := big.NewFloat(0).SetInt(balance)
					balanceFloat := big.NewFloat(0).Quo(balanceFloatExp, ExpScaleFloat)
					collateral := big.NewFloat(0).Mul(balanceFloat, multiplier)
					totalCollateral = big.NewFloat(0).Add(totalCollateral, collateral)

					borrowFloatExp := big.NewFloat(0).SetInt(borrow)
					borrowFloat := big.NewFloat(0).Quo(borrowFloatExp, ExpScaleFloat)
					loan := big.NewFloat(0).Mul(borrowFloat, priceFloat)
					totalLoan = big.NewFloat(0).Add(totalLoan, loan)

					//build account table
					assets = append(assets, Asset{
						Symbol:  symbol,
						Balance: balanceFloat,
						Loan:    borrowFloat,
					})
				}

				totalLoan = big.NewFloat(0).Add(totalLoan, mintVAISFloat)
				fmt.Printf("totalCollateral:%v, totalLoan:%v\n", totalCollateral, totalLoan)
				healthFactor := big.NewFloat(0).Quo(totalCollateral, totalLoan)
				fmt.Printf("healthFactor：%v\n", healthFactor)

				//update market table and account table
				info := AccountInfo{
					HealthFactor: healthFactor,
					Assets:       assets,
				}
				s.updateDB(account, info)
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

func (s *Syncer) updateDB(account common.Address, info AccountInfo) {
	s.deleteAccount(account)
	s.storeAccount(account, info)
}

func (s *Syncer) deleteAccount(account common.Address) {
	db := s.db
	accountBytes := account.Bytes()

	had, _ := db.Has(dbm.AccountStoreKey(accountBytes), nil)
	if had {
		bz, _ := db.Get(dbm.AccountStoreKey(accountBytes), nil)
		var info AccountInfo
		err := json.Unmarshal(bz, &info)
		if err != nil {
			panic(err)
		}

		healthFactor := info.HealthFactor
		assets := info.Assets

		for _, asset := range assets {
			db.Delete(dbm.MarketStoreKey([]byte(asset.Symbol), accountBytes), nil)
		}

		if healthFactor.Cmp(BigFloat1P0) == -1 {
			db.Delete(dbm.LiquidationBelow1P0StoreKey(accountBytes), nil)
		} else if healthFactor.Cmp(BigFloat1P2) == -1 {
			db.Delete(dbm.LiquidationBelow1P2StoreKey(accountBytes), nil)
		} else if healthFactor.Cmp(BigFloat1P5) == -1 {
			db.Delete(dbm.LiquidationBelow1P5StoreKey(accountBytes), nil)
		} else if healthFactor.Cmp(BigFloat2P0) == -1 {
			db.Delete(dbm.LiquidationBelow2P0StoreKey(accountBytes), nil)
		} else if healthFactor.Cmp(BigFloat3P0) == -1 {
			db.Delete(dbm.LiquidationBelow3P0StoreKey(accountBytes), nil)
		} else {
			db.Delete(dbm.LiquidationAbove3P0StoreKey(accountBytes), nil)
		}
		db.Delete(dbm.AccountStoreKey(accountBytes), nil)
	}
}

func (s *Syncer) storeAccount(account common.Address, info AccountInfo) {
	db := s.db
	accountBytes := account.Bytes()
	healthFactor := info.HealthFactor

	for _, asset := range info.Assets {
		db.Put(dbm.MarketStoreKey([]byte(asset.Symbol), accountBytes), accountBytes, nil)
	}

	if healthFactor.Cmp(BigFloat1P0) == -1 {
		db.Put(dbm.LiquidationBelow1P0StoreKey(accountBytes), accountBytes, nil)
	} else if healthFactor.Cmp(BigFloat1P2) == -1 {
		db.Put(dbm.LiquidationBelow1P2StoreKey(accountBytes), accountBytes, nil)
	} else if healthFactor.Cmp(BigFloat1P5) == -1 {
		db.Put(dbm.LiquidationBelow1P5StoreKey(accountBytes), accountBytes, nil)
	} else if healthFactor.Cmp(BigFloat2P0) == -1 {
		db.Put(dbm.LiquidationBelow2P0StoreKey(accountBytes), accountBytes, nil)
	} else if healthFactor.Cmp(BigFloat3P0) == -1 {
		db.Put(dbm.LiquidationBelow3P0StoreKey(accountBytes), accountBytes, nil)
	} else {
		db.Put(dbm.LiquidationAbove3P0StoreKey(accountBytes), accountBytes, nil)
	}
	bz, _ := json.Marshal(info)
	db.Put(dbm.AccountStoreKey(accountBytes), bz, nil)
}

func (s *Syncer) syncLiquidationBelow1P0() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P0Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			s.syncAccounts(accounts)
			iter.Release()
			t.Reset(time.Second * SyncIntervalBelow1P0)
		}
	}
}

func (s *Syncer) syncLiquidationBelow1P2() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P2Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			s.syncAccounts(accounts)
			iter.Release()
			t.Reset(time.Second * SyncIntervalBelow1P2)
		}
	}
}

func (s *Syncer) syncLiquidationBelow1P5() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P5Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			s.syncAccounts(accounts)
			iter.Release()
			t.Reset(time.Second * SyncIntervalBelow1P5)
		}
	}
}

func (s *Syncer) syncLiquidationBelow2P0() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow2P0Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			s.syncAccounts(accounts)
			iter.Release()
			t.Reset(time.Second * SyncIntervalBelow2P0)
		}
	}
}

func (s *Syncer) syncLiquidationBelow3P0() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow3P0Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			s.syncAccounts(accounts)
			iter.Release()
			t.Reset(time.Second * SyncIntervalBelow3P0)
		}
	}
}

func (s *Syncer) syncLiquidationAbove3P0() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationAbove3P0Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			s.syncAccounts(accounts)
			iter.Release()
			t.Reset(time.Second * SyncIntervalAbove3P0)
		}
	}
}

func (s *Syncer) syncAccounts(accounts []common.Address) error {
	comptroller, err := venus.NewComptroller(common.HexToAddress(s.comptroller), s.c)
	if err != nil {
		return err
	}

	for _, account := range accounts {
		totalCollateral := big.NewFloat(0)
		totalLoan := big.NewFloat(0)

		var assets []Asset
		mintVAIS, err := comptroller.MintedVAIs(nil, account)
		if err != nil {
			return err
		}

		mintVAISFloatExp := big.NewFloat(0).SetInt(mintVAIS)
		mintVAISFloat := big.NewFloat(0).Quo(mintVAISFloatExp, ExpScaleFloat)

		markets, err := comptroller.GetAssetsIn(nil, account)
		if err != nil {
			return err
		}

		for _, market := range markets {
			symbol := AddressToSymbol[market]
			_, balance, borrow, exchangeRate, err := Vbep20s[symbol].GetAccountSnapshot(nil, account)
			if err != nil {
				return err
			}

			exchangeRateFloatExp := big.NewFloat(0).SetInt(exchangeRate)
			exchangeRateFloat := big.NewFloat(0).Quo(exchangeRateFloatExp, ExpScaleFloat)

			collateralFactorFloat := TokenDetail[symbol].CollateralFactorFloat
			priceFloat := TokenDetail[symbol].PriceFloat

			multiplier := big.NewFloat(0).Mul(exchangeRateFloat, collateralFactorFloat)
			multiplier = big.NewFloat(0).Mul(multiplier, priceFloat)

			balanceFloatExp := big.NewFloat(0).SetInt(balance)
			balanceFloat := big.NewFloat(0).Quo(balanceFloatExp, ExpScaleFloat)
			collateral := big.NewFloat(0).Mul(balanceFloat, multiplier)
			totalCollateral = big.NewFloat(0).Add(totalCollateral, collateral)

			borrowFloatExp := big.NewFloat(0).SetInt(borrow)
			borrowFloat := big.NewFloat(0).Quo(borrowFloatExp, ExpScaleFloat)
			loan := big.NewFloat(0).Mul(borrowFloat, priceFloat)
			totalLoan = big.NewFloat(0).Add(totalLoan, loan)

			//build account table
			assets = append(assets, Asset{
				Symbol:  symbol,
				Balance: balanceFloat,
				Loan:    borrowFloat,
			})
		}

		totalLoan = big.NewFloat(0).Add(totalLoan, mintVAISFloat)
		fmt.Printf("totalCollateral:%v, totalLoan:%v\n", totalCollateral, totalLoan)
		healthFactor := big.NewFloat(0).Quo(totalCollateral, totalLoan)
		fmt.Printf("healthFactor：%v\n", healthFactor)

		//update market table and account table
		info := AccountInfo{
			HealthFactor: healthFactor,
			Assets:       assets,
		}
		s.updateDB(account, info)
	}
	return nil
}

func (s *Syncer) syncOneAccount(account common.Address) (<-chan common.Address, error) {
	comptroller, err := venus.NewComptroller(common.HexToAddress(s.comptroller), s.c)
	if err != nil {
		return nil, err
	}

	totalCollateral := big.NewFloat(0)
	totalLoan := big.NewFloat(0)

	var assets []Asset
	mintVAIS, err := comptroller.MintedVAIs(nil, account)
	if err != nil {
		return nil, err
	}

	mintVAISFloatExp := big.NewFloat(0).SetInt(mintVAIS)
	mintVAISFloat := big.NewFloat(0).Quo(mintVAISFloatExp, ExpScaleFloat)

	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil {
		return nil, err
	}

	for _, market := range markets {
		symbol := AddressToSymbol[market]
		_, balance, borrow, exchangeRate, err := Vbep20s[symbol].GetAccountSnapshot(nil, account)
		if err != nil {
			return nil, err
		}

		exchangeRateFloatExp := big.NewFloat(0).SetInt(exchangeRate)
		exchangeRateFloat := big.NewFloat(0).Quo(exchangeRateFloatExp, ExpScaleFloat)

		collateralFactorFloat := TokenDetail[symbol].CollateralFactorFloat
		priceFloat := TokenDetail[symbol].PriceFloat

		multiplier := big.NewFloat(0).Mul(exchangeRateFloat, collateralFactorFloat)
		multiplier = big.NewFloat(0).Mul(multiplier, priceFloat)

		balanceFloatExp := big.NewFloat(0).SetInt(balance)
		balanceFloat := big.NewFloat(0).Quo(balanceFloatExp, ExpScaleFloat)
		collateral := big.NewFloat(0).Mul(balanceFloat, multiplier)
		totalCollateral = big.NewFloat(0).Add(totalCollateral, collateral)

		borrowFloatExp := big.NewFloat(0).SetInt(borrow)
		borrowFloat := big.NewFloat(0).Quo(borrowFloatExp, ExpScaleFloat)
		loan := big.NewFloat(0).Mul(borrowFloat, priceFloat)
		totalLoan = big.NewFloat(0).Add(totalLoan, loan)

		//build account table
		assets = append(assets, Asset{
			Symbol:  symbol,
			Balance: balanceFloat,
			Loan:    borrowFloat,
		})
	}

	totalLoan = big.NewFloat(0).Add(totalLoan, mintVAISFloat)
	fmt.Printf("totalCollateral:%v, totalLoan:%v\n", totalCollateral, totalLoan)
	healthFactor := big.NewFloat(0).Quo(totalCollateral, totalLoan)
	fmt.Printf("healthFactor：%v\n", healthFactor)

	//update market table and account table
	info := AccountInfo{
		HealthFactor: healthFactor,
		Assets:       assets,
	}
	s.updateDB(account, info)
	if healthFactor.Cmp(big.NewFloat(1)) == -1 || healthFactor.Cmp(big.NewFloat(1)) == 0 {
		// catch a liquidation
		return nil, nil
	} else {
		return nil, nil
	}
}
