package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/readygo67/LiquidationBot/db"
	"github.com/readygo67/LiquidationBot/venus"
	"github.com/shopspring/decimal"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	ConfirmHeight        = 0
	ScanSpan             = 1000
	SyncIntervalBelow1P0 = 3 //in secs
	SyncIntervalBelow1P1 = 6
	SyncIntervalBelow1P5 = 30
	SyncIntervalBelow2P0 = 150
	SyncIntervalAbove2P0 = 300
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

type AssetWithPrice struct {
	Symbol       string
	Decimals     uint8
	Balance      *big.Float
	Loan         *big.Float
	Price        *big.Float
	ExchangeRate *big.Float
}

type AccountInfo struct {
	HealthFactor *big.Float
	Assets       []Asset
}

type FeededPrice struct {
	Address    common.Address
	PriceFloat *big.Float
	Hash       common.Hash
}

type FeededPrices struct {
	Prices []FeededPrice
	Height uint64
}

type Liquidation struct {
	Address      common.Address
	HealthFactor *big.Float
	BlockNumber  uint64
	Endtime      time.Time
}

type semaphore chan struct{}

var (
	ExpScaleFloat, _ = big.NewFloat(0).SetString("1000000000000000000")
	BigFloatZero, _  = big.NewFloat(0).SetString("0")
	BigFloat1P0, _   = big.NewFloat(0).SetString("1.0")
	BigFloat1P1, _   = big.NewFloat(0).SetString("1.1")
	BigFloat1P5, _   = big.NewFloat(0).SetString("1.5")
	BigFloat2P0, _   = big.NewFloat(0).SetString("2.0")
	BigFloat3P0, _   = big.NewFloat(0).SetString("3.0")
)

type Syncer struct {
	c  *ethclient.Client
	db *leveldb.DB

	oracle              *venus.Oracle
	comptroller         *venus.Comptroller
	symbols             map[common.Address]string
	tokens              map[string]TokenInfo
	vbep20s             map[common.Address]*venus.Vbep20
	syncDone            bool
	m                   sync.Mutex
	wg                  sync.WaitGroup
	quitCh              chan struct{}
	forceUpdatePricesCh chan struct{}
	feededPricesCh      chan *FeededPrices

	liquidationCh        chan *Liquidation
	priortyLiquidationCh chan *Liquidation
}

func (s semaphore) Acquire() {
	s <- struct{}{}
}

func (s semaphore) Release() {
	<-s
}

func NewSyncer(
	c *ethclient.Client,
	db *leveldb.DB,
	comptrollerAddress string,
	oracleAddress string,
	feededPricesCh chan *FeededPrices,
	liquidationCh chan *Liquidation,
	priorityLiquationCh chan *Liquidation) *Syncer {

	exist, err := db.Has(dbm.BorrowerNumberKey(), nil)
	if !exist {
		db.Put(dbm.BorrowerNumberKey(), big.NewInt(0).Bytes(), nil)
	}

	comptroller, err := venus.NewComptroller(common.HexToAddress(comptrollerAddress), c)
	if err != nil {
		panic(err)
	}

	oracle, err := venus.NewOracle(common.HexToAddress(oracleAddress), c)
	if err != nil {
		panic(err)
	}

	markets, err := comptroller.GetAllMarkets(nil)
	if err != nil {
		panic(err)
	}

	symbols := make(map[common.Address]string)
	tokens := make(map[string]TokenInfo)
	vbep20s := make(map[common.Address]*venus.Vbep20)

	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(len(markets))

	sem := make(semaphore, runtime.NumCPU())
	for _, market_ := range markets {
		market := market_
		sem.Acquire()
		go func() {
			defer sem.Release()
			defer wg.Done()
			vbep20, err := venus.NewVbep20(market, c)
			if err != nil {
				panic(err)
			}
			symbol, err := vbep20.Symbol(nil)
			if err != nil {
				panic(err)
			}

			marketDetail, err := comptroller.Markets(nil, market)
			if err != nil {
				panic(err)
			}

			price, err := oracle.GetUnderlyingPrice(nil, market)
			if err != nil {
				price = big.NewInt(0)
			}

			collateralFactorFloatExp := big.NewFloat(0).SetInt(marketDetail.CollateralFactorMantissa)
			collateralFactorFloat := big.NewFloat(0).Quo(collateralFactorFloatExp, ExpScaleFloat)
			priceFloatExp := big.NewFloat(0).SetInt(price)
			priceFloat := big.NewFloat(0).Quo(priceFloatExp, ExpScaleFloat)

			token := TokenInfo{
				Address:               market,
				CollateralFactorFloat: collateralFactorFloat,
				PriceFloat:            priceFloat,
			}

			m.Lock()
			//fmt.Printf("market:%v, symbol:%v, token:%+v\n", market, symbol, token)
			tokens[symbol] = token
			symbols[market] = symbol
			vbep20s[market] = vbep20
			m.Unlock()
		}()
	}
	wg.Wait()

	return &Syncer{
		c:                    c,
		db:                   db,
		oracle:               oracle,
		comptroller:          comptroller,
		tokens:               tokens,
		symbols:              symbols,
		vbep20s:              vbep20s,
		m:                    m,
		quitCh:               make(chan struct{}),
		forceUpdatePricesCh:  make(chan struct{}),
		feededPricesCh:       feededPricesCh,
		liquidationCh:        liquidationCh,
		priortyLiquidationCh: priorityLiquationCh,
	}
}

func (s *Syncer) Start() {
	log.Info("server start")
	fmt.Println("server start")

	s.wg.Add(8)
	go s.syncMarketsAndPrices()
	go s.feedPrices()
	go s.syncAllBorrowers()
	go s.syncLiquidationBelow1P0()
	go s.syncLiquidationBelow1P1()
	go s.syncLiquidationBelow1P5()
	go s.syncLiquidationBelow2P0()
	go s.syncLiquidationAbove2P0()
}

func (s *Syncer) Stop() {
	close(s.quitCh)
	s.wg.Wait()
}

func (s *Syncer) syncMarketsAndPrices() {
	defer s.wg.Done()
	t := time.NewTimer(0)
	defer t.Stop()

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			fmt.Printf("%v th sync markers and prices @ %v\n", count, time.Now())
			count++
			s.doSyncMarketsAndPrices()
			t.Reset(time.Second * 3)
		case <-s.forceUpdatePricesCh:
			s.doSyncMarketsAndPrices()
		}
	}
}

func (s *Syncer) doSyncMarketsAndPrices() {
	comptroller := s.comptroller
	oracle := s.oracle
	c := s.c

	markets, err := comptroller.GetAllMarkets(nil)
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(markets))

	sem := make(semaphore, runtime.NumCPU())
	for _, market_ := range markets {
		market := market_
		sem.Acquire()
		go func() {
			defer sem.Release()
			defer wg.Done()
			vbep20, err := venus.NewVbep20(market, c)
			if err != nil {
				return
			}
			symbol, err := vbep20.Symbol(nil)
			if err != nil {
				return
			}

			marketDetail, err := comptroller.Markets(nil, market)
			if err != nil {
				return
			}

			price, err := oracle.GetUnderlyingPrice(nil, market)
			if err != nil {
				price = big.NewInt(0)
			}

			collateralFactorFloatExp := big.NewFloat(0).SetInt(marketDetail.CollateralFactorMantissa)
			collateralFactorFloat := big.NewFloat(0).Quo(collateralFactorFloatExp, ExpScaleFloat)
			priceFloatExp := big.NewFloat(0).SetInt(price)
			priceFloat := big.NewFloat(0).Quo(priceFloatExp, ExpScaleFloat)

			token := TokenInfo{
				Address:               market,
				CollateralFactorFloat: collateralFactorFloat,
				PriceFloat:            priceFloat,
			}

			s.m.Lock()
			//fmt.Printf("symbol:%v, market:%v, token:%v\n", symbol, market, token)
			s.tokens[symbol] = token
			s.symbols[market] = symbol
			if s.vbep20s[market] == nil {
				s.vbep20s[market] = vbep20
			}
			s.m.Unlock()
		}()
	}
	wg.Wait()
}

func (s *Syncer) feedPrices() {
	defer s.wg.Done()
	for {
		select {
		case <-s.quitCh:
			return
		case feededPrices := <-s.feededPricesCh:
			s.doFeededPrices(feededPrices)
		}
	}
}

func (s *Syncer) doFeededPrices(feededPrices *FeededPrices) {
	db := s.db
	var accounts []common.Address
	exist := make(map[common.Address]bool)

	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(len(feededPrices.Prices))

	sem := make(semaphore, runtime.NumCPU())
	for _, feededPrice := range feededPrices.Prices {
		sem.Acquire()
		symbol := s.symbols[feededPrice.Address]
		prefix := append(dbm.MarketPrefix, []byte(symbol)...)

		go func() {
			sem.Release()
			wg.Done()
			iter := db.NewIterator(util.BytesPrefix(prefix), nil)
			for iter.Next() {
				account := common.BytesToAddress(iter.Value())
				if exist[account] {
					continue
				}
				m.Lock()
				exist[account] = true
				accounts = append(accounts, account)
				m.Unlock()
			}
			iter.Release()
		}()
	}
	wg.Wait()

	wg.Add(len(accounts))
	for _, account := range accounts {
		sem.Acquire()
		go func() {
			sem.Release()
			wg.Done()
			s.syncOneAccountWithFeededPrices(account, feededPrices)
		}()
	}
	wg.Wait()
}

func (s *Syncer) syncAllBorrowers() {
	defer s.wg.Done()
	db := s.db
	c := s.c
	ctx := context.Background()
	s.m.Lock()
	symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	s.m.Unlock()

	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")
	vbep20Abi, _ := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))

	var logs []types.Log
	var addresses []common.Address

	for _, token := range tokens {
		addresses = append(addresses, token.Address)
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
				fmt.Printf("syncAllBorrowers, fail to filter logs, err:%v\n", err)
				goto EndWithoutUpdateHeight
			}

			for i, log := range logs {
				var borrowEvent venus.Vbep20Borrow
				err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", log.Data)
				fmt.Printf("%v height:%v, name:%v account:%v\n", (i + 1), log.BlockNumber, symbols[log.Address], borrowEvent.Borrower)

				account := borrowEvent.Borrower
				err := s.syncOneAccountWithIncreaseAccountNumber(account)
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

func (s *Syncer) syncLiquidationBelow1P0() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			fmt.Printf("%vth sync below 1.0 start @ %v\n", count, time.Now())
			count++

			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P0Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			iter.Release()

			s.syncAccounts(accounts)
			t.Reset(time.Second * SyncIntervalBelow1P0)
		}
	}
}

func (s *Syncer) syncLiquidationBelow1P1() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			fmt.Printf("%vth sync below 1.1 start @ %v\n", count, time.Now())
			count++

			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P1Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			iter.Release()

			s.syncAccounts(accounts)
			t.Reset(time.Second * SyncIntervalBelow1P1)
		}
	}
}

func (s *Syncer) syncLiquidationBelow1P5() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			fmt.Printf("%vth sync below 1.5 start @ %v\n", count, time.Now())
			count++

			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P5Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			iter.Release()

			s.syncAccounts(accounts)
			t.Reset(time.Second * SyncIntervalBelow1P5)
		}
	}
}

func (s *Syncer) syncLiquidationBelow2P0() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			fmt.Printf("%vth sync below 2.0 start @ %v\n", count, time.Now())
			count++

			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow2P0Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			iter.Release()

			s.syncAccounts(accounts)
			t.Reset(time.Second * SyncIntervalBelow2P0)
		}
	}
}

func (s *Syncer) syncLiquidationAbove2P0() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			fmt.Printf("%vth sync above 2.0 start @ %v\n", count, time.Now())
			count++

			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationAbove2P0Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			iter.Release()

			s.syncAccounts(accounts)
			t.Reset(time.Second * SyncIntervalAbove2P0)
		}
	}
}

func (s *Syncer) syncAccounts(accounts []common.Address) {
	var wg sync.WaitGroup
	wg.Add(len(accounts))
	sem := make(semaphore, runtime.NumCPU())

	for _, account_ := range accounts {
		account := account_
		sem.Acquire()
		go func() {
			defer sem.Release()
			defer wg.Done()
			s.syncOneAccount(account)
		}()
	}
	wg.Wait()
}

func (s *Syncer) syncOneAccount(account common.Address) error {
	ctx := context.Background()
	comptroller := s.comptroller

	s.m.Lock()
	symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	vbep20s := s.vbep20s
	s.m.Unlock()

	totalCollateral := big.NewFloat(0)
	totalLoan := big.NewFloat(0)

	var assets []Asset
	mintVAIS, err := comptroller.MintedVAIs(nil, account)
	if err != nil {
		fmt.Printf("syncOneAccount, fail to get MintedVAIs, err:%v\n", err)
		return err
	}

	mintVAISFloatExp := big.NewFloat(0).SetInt(mintVAIS)
	mintVAISFloat := big.NewFloat(0).Quo(mintVAISFloatExp, ExpScaleFloat)
	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil {
		fmt.Printf("syncOneAccount, fail to get GetAssetsIn, err:%v\n", err)
		return err
	}

	for _, market := range markets {
		//fmt.Printf("market:%v\n", market)
		_, balance, borrow, exchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil {
			fmt.Printf("syncOneAccount, fail to get GetAccountSnapshot, err:%v\n", err)
			return err
		}

		exchangeRateFloatExp := big.NewFloat(0).SetInt(exchangeRate)
		exchangeRateFloat := big.NewFloat(0).Quo(exchangeRateFloatExp, ExpScaleFloat)

		symbol := symbols[market]
		collateralFactorFloat := tokens[symbol].CollateralFactorFloat
		priceFloat := tokens[symbol].PriceFloat

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
	healthFactor := big.NewFloat(100)
	if totalLoan.Cmp(BigFloatZero) == 1 {
		healthFactor = big.NewFloat(0).Quo(totalCollateral, totalLoan)
	}
	fmt.Printf("accout:%v, healthFactor:%v, totalCollateral:%v, totalLoan:%v\n", account, healthFactor, totalCollateral, totalLoan)
	//update market table and account table
	info := AccountInfo{
		HealthFactor: healthFactor,
		Assets:       assets,
	}
	s.updateDB(account, info)
	if healthFactor.Cmp(big.NewFloat(1)) != 1 {
		blockNumber, _ := s.c.BlockNumber(ctx)
		liquidation := &Liquidation{
			Address:      account,
			HealthFactor: healthFactor,
			BlockNumber:  blockNumber,
		}
		s.liquidationCh <- liquidation
	}
	return nil
}

func (s *Syncer) syncOneAccountWithFeededPrices(account common.Address, feededPrices *FeededPrices) error {
	ctx := context.Background()
	comptroller := s.comptroller

	s.m.Lock()
	symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	vbep20s := s.vbep20s
	s.m.Unlock()

	totalCollateral := big.NewFloat(0)
	totalLoan := big.NewFloat(0)

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
		_, balance, borrow, exchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil {
			return err
		}

		exchangeRateFloatExp := big.NewFloat(0).SetInt(exchangeRate)
		exchangeRateFloat := big.NewFloat(0).Quo(exchangeRateFloatExp, ExpScaleFloat)

		symbol := symbols[market]
		collateralFactorFloat := tokens[symbol].CollateralFactorFloat
		priceFloat := tokens[symbol].PriceFloat
		//use feeded prices
		for _, feededPrice := range feededPrices.Prices {
			if tokens[symbol].Address == feededPrice.Address {
				priceFloat = feededPrice.PriceFloat
			}
		}

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
	}

	totalLoan = big.NewFloat(0).Add(totalLoan, mintVAISFloat)
	healthFactor := big.NewFloat(100)
	if totalLoan.Cmp(BigFloatZero) == 1 {
		healthFactor = big.NewFloat(0).Quo(totalCollateral, totalLoan)
	}
	fmt.Printf("account:%v, healthFactor:%v, totalCollateral:%v, totalLoan:%v \n", account, healthFactor, totalCollateral, totalLoan)

	if healthFactor.Cmp(big.NewFloat(1)) != 1 {
		blockNumber, _ := s.c.BlockNumber(ctx)
		liquidation := &Liquidation{
			Address:      account,
			HealthFactor: healthFactor,
			BlockNumber:  blockNumber,
		}
		s.priortyLiquidationCh <- liquidation
	}
	return nil
}

func (s *Syncer) syncOneAccountWithIncreaseAccountNumber(account common.Address) error {
	ctx := context.Background()
	c := s.c
	db := s.db

	accountBytes := account.Bytes()
	exist, err := db.Has(dbm.BorrowersStoreKey(accountBytes), nil)
	if err != nil {
		return err
	}

	byteCode, err := c.CodeAt(ctx, account, nil)
	if len(byteCode) > 0 {
		//ignore smart contract
		return nil
	}

	err = s.syncOneAccount(account)
	if err != nil {
		return err
	}

	if !exist {
		//if account not exist in borrowers table, record it into borrowers table and increase borrowers number
		err = db.Put(dbm.BorrowersStoreKey(accountBytes), accountBytes, nil)
		if err != nil {
			return err
		}

		bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
		num := big.NewInt(0).SetBytes(bz).Int64()
		if err != nil {
			return err
		}

		num += 1
		err = db.Put(dbm.BorrowerNumberKey(), big.NewInt(num).Bytes(), nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Syncer) calculateRepayAmount(liquidation *Liquidation) error {
	comptroller := s.comptroller
	oracle := s.oracle
	account := liquidation.Address

	s.m.Lock()
	symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	vbep20s := s.vbep20s
	s.m.Unlock()

	totalCollateral := decimal.NewFromInt(0)
	totalLoan := decimal.NewFromInt(0)

	var assets []AssetWithPrice
	//mintVAIS, err := comptroller.MintedVAIs(nil, account)
	//if err != nil {
	//	fmt.Printf("liquidate, fail to get MintedVAIs, err:%v\n", err)
	//	return err
	//}
	//mintVAISFloatExp := big.NewFloat(0).SetInt(mintVAIS)
	//mintVAISFloat := big.NewFloat(0).Quo(mintVAISFloatExp, ExpScaleFloat)

	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil {
		fmt.Printf("liquidate, fail to get GetAssetsIn, err:%v\n", err)
		return err
	}

	closeFactor, err := comptroller.CloseFactorMantissa(nil)
	if err != nil {
		fmt.Printf("liquidate, fail to get closefactor, err:%v\n", err)
		return err
	}
	
	closeFactorFloatExp := big.NewFloat(0).SetInt(closeFactor)
	closeFactorFloat := big.NewFloat(0).Quo(closeFactorFloatExp, ExpScaleFloat)

	liquidationIncentive, err := comptroller.LiquidationIncentiveMantissa(nil)
	if err != nil {
		fmt.Printf("liquidate, fail to get liquidatino incentive, err:%v\n", err)
		return err
	}
	liquidationIncentiveFloatExp := big.NewFloat(0).SetInt(liquidationIncentive)
	liquidationIncentiveFloat := big.NewFloat(0).Quo(liquidationIncentiveFloatExp, ExpScaleFloat)

	for _, market := range markets {
		//fmt.Printf("market:%v\n", market)
		//underlyingTokenAddress, err := vbep20s[market].Underlying(nil)
		//if err != nil {
		//	fmt.Printf("liquidate, fail to get underlying token, err:%v\n", err)
		//	return err
		//}

		//decimals, err := vbep20s[underlyingTokenAddress].Decimals(nil)
		//if err != nil {
		//	fmt.Printf("liquidate, fail to get decimal, err:%v\n", err)
		//	return err
		//}
		decimals := uint8(18)

		_, balance, borrow, exchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil {
			fmt.Printf("liquidate, fail to get GetAccountSnapshot, err:%v\n", err)
			return err
		}

		price, err := oracle.GetUnderlyingPrice(nil, market)
		if err != nil {
			fmt.Printf("liquidate, fail to get underlying price, err:%v\n", err)
			return err
		}

		priceFloatExp := big.NewFloat(0).SetInt(price)
		priceFloat := big.NewFloat(0).Quo(priceFloatExp, ExpScaleFloat)

		exchangeRateFloatExp := big.NewFloat(0).SetInt(exchangeRate)
		exchangeRateFloat := big.NewFloat(0).Quo(exchangeRateFloatExp, ExpScaleFloat)

		symbol := symbols[market]
		collateralFactorFloat := tokens[symbol].CollateralFactorFloat

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
		assets = append(assets, AssetWithPrice{
			Symbol:       symbol,
			Decimals:     decimals,
			Balance:      collateral,
			Loan:         loan,
			Price:        priceFloat,
			ExchangeRate: exchangeRateFloat,
		})
	}

	for _, asset := range assets {
		fmt.Printf("asset:%+v, addres:%v\n", asset, tokens[asset.Symbol].Address)
	}
	fmt.Printf("totalBalance:%v, totalLoan:%v, shortfall:%v\n", totalCollateral, totalLoan, big.NewFloat(0).Sub(totalLoan, totalCollateral))

	//select the repayed token and seized collateral token
	maxLoanValue := big.NewFloat(0)
	maxLoanSymbol := ""
	repayIndex := 0

	for i, asset := range assets {
		if asset.Loan.Cmp(maxLoanValue) == 1 {
			maxLoanValue = asset.Loan
			maxLoanSymbol = asset.Symbol
			repayIndex = i
		}
	}
	//the closeFactor applies to only single borrowed asset
	maxRepayValue := big.NewFloat(0).Mul(maxLoanValue, closeFactorFloat)
	repaySymbol := maxLoanSymbol

	repayValue := big.NewFloat(0)
	seizedCollateralSymbol := ""
	seizedIndex := 0
	for k, asset := range assets {
		if asset.Balance.Cmp(maxRepayValue) != -1 {
			repayValue = maxRepayValue
			seizedCollateralSymbol = asset.Symbol
			seizedIndex = k
			break
		} else {
			if asset.Balance.Cmp(repayValue) == 1 {
				repayValue = asset.Balance
				seizedCollateralSymbol = asset.Symbol
				seizedIndex = k
			}
		}
	}

	fmt.Printf("repaySmbol:%v repayValue:%v, seizedSymbol:%v\n", repaySymbol, repayValue, seizedCollateralSymbol)

	repayAmountFloat := big.NewFloat(0).Quo(repayValue, assets[repayIndex].Price) //amount of underlyingtoken token
	fmt.Printf("repayValue:%v, price:%v, amount:%v\n", repayValue, assets[repayIndex].Price, repayAmountFloat)

	//repayAmountFloat = big.NewFloat(0).Quo(repayAmountFloat, assets[repayIndex].ExchangeRate) //amount of cToken

	repayAmountDecimal, err := decimal.NewFromString(repayAmountFloat.String())
	if err != nil {
		fmt.Printf("liquidate, fail to covert to repayAmount decinal, err:%v\n", err)
		return err
	}

	repayAmountDecimal = repayAmountDecimal.Mul(decimal.New(1, int32(assets[repayIndex].Decimals)))
	repayAmount := repayAmountDecimal.Truncate(0)
	fmt.Printf("repayAmountDecimal:%v, repayAmount:%v, repayAmountFloat:%v\n", repayAmountDecimal, repayAmount, repayAmountFloat)
	errCode, seizedAmount, err := comptroller.LiquidateCalculateSeizeTokens(nil, tokens[repaySymbol].Address, tokens[seizedCollateralSymbol].Address, repayAmount.BigInt())
	fmt.Printf("errCode:%v, seizedAmount:%v, err:%v\n", errCode, seizedAmount, err)

	seizedAmountExpFloat := big.NewFloat(0).SetInt(seizedAmount)
	seizedUnderlyingTokenAmountExpFloat := big.NewFloat(0).Mul(seizedAmountExpFloat, assets[seizedIndex].ExchangeRate)
	seizedUnderlyingTokenAmountFloat := big.NewFloat(0).Quo(seizedUnderlyingTokenAmountExpFloat, ExpScaleFloat)
	seizedUnderlyingTokenValue := big.NewFloat(0).Mul(seizedUnderlyingTokenAmountFloat, assets[seizedIndex].Price)

	fmt.Printf("seizedUnderlyingTokenValue:%v", seizedUnderlyingTokenValue)

	ratio := big.NewFloat(0).Quo(seizedUnderlyingTokenValue, repayValue)
	fmt.Printf("ratio:%v\n", ratio)

	fmt.Printf("seizedIndex:%v, liquidationIncentiveFloat:%v\n", seizedIndex, liquidationIncentiveFloat)
	//decimals := assets[repayIndex].Decimal
	//DecimalScaleFloat := big.NewFloat(0).SetInt(big.NewInt(0).Exp(10, decimalw))
	//repayAmountFloatExp := big.NewFloat(0).Mul(repayAmountFloat, DecimalScaleFloat)
	//repayAmout := repayAmountFloatExp.
	//
	//seizedValue := big.NewFloat(0).Mul(repayValue, liquidationIncentiveFloat)
	//seizedUnderlyingAmount := big.NewFloat(0).Quo(seizedValue, assets[seizedIndex].Price)
	//seizedAmount := big.NewFloat(0).Quo(seizedUnderlyingAmount, assets[seizedIndex].ExchangeRate)
	//decimalAmount, err := decimal.NewFromString(seizedAmount.String())
	//if err != nil {
	//	fmt.Printf("liquidate, fail to get underlying price, err:%v\n", err)
	//	return err
	//}
	//
	//decimal.New(1, decimals)

	//actualRepayAmount = seizeTokens * priceCollateral * exchangeRate / (liquidationIncentive * priceBorrowed)
	//repaysymbol, targetCollateralAmount, targeCollateralSymbol
	//calculate profit
	//
	//flashLoanFee := ((amount * 25) / 9975) + 1;
	//
	//gasPrice, err := s.c.SuggestGasPrice(ctx)
	//gas := 7000000
	//fee :=
	return nil
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
		} else if healthFactor.Cmp(BigFloat1P1) == -1 {
			db.Delete(dbm.LiquidationBelow1P1StoreKey(accountBytes), nil)
		} else if healthFactor.Cmp(BigFloat1P5) == -1 {
			db.Delete(dbm.LiquidationBelow1P5StoreKey(accountBytes), nil)
		} else if healthFactor.Cmp(BigFloat2P0) == -1 {
			db.Delete(dbm.LiquidationBelow2P0StoreKey(accountBytes), nil)
		} else {
			db.Delete(dbm.LiquidationAbove2P0StoreKey(accountBytes), nil)
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
	} else if healthFactor.Cmp(BigFloat1P1) == -1 {
		db.Put(dbm.LiquidationBelow1P1StoreKey(accountBytes), accountBytes, nil)
	} else if healthFactor.Cmp(BigFloat1P5) == -1 {
		db.Put(dbm.LiquidationBelow1P5StoreKey(accountBytes), accountBytes, nil)
	} else if healthFactor.Cmp(BigFloat2P0) == -1 {
		db.Put(dbm.LiquidationBelow2P0StoreKey(accountBytes), accountBytes, nil)
	} else {
		db.Put(dbm.LiquidationAbove2P0StoreKey(accountBytes), accountBytes, nil)
	}
	bz, _ := json.Marshal(info)
	db.Put(dbm.AccountStoreKey(accountBytes), bz, nil)
	db.Put(dbm.BorrowersStoreKey(accountBytes), accountBytes, nil)
}

func copySymbols(src map[common.Address]string) map[common.Address]string {
	dst := make(map[common.Address]string)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func copyTokens(src map[string]TokenInfo) map[string]TokenInfo {
	dst := make(map[string]TokenInfo)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
