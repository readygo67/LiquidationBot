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
	ScanSpan             = 10000
	SyncIntervalBelow1P0 = 3 //in secs
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
	BigFloat1P2, _   = big.NewFloat(0).SetString("1.2")
	BigFloat1P5, _   = big.NewFloat(0).SetString("1.5")
	BigFloat2P0, _   = big.NewFloat(0).SetString("2.0")
	BigFloat3P0, _   = big.NewFloat(0).SetString("3.0")
)

type Syncer struct {
	c  *ethclient.Client
	db *leveldb.DB

	oracle      *venus.Oracle
	comptroller *venus.Comptroller
	symbols     map[common.Address]string
	tokens      map[string]TokenInfo
	vbep20s     map[common.Address]*venus.Vbep20

	syncDone            bool
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

			m.Lock()
			//fmt.Printf("symbol:%v, market:%v, token:%v\n", symbol, market, token)
			s.tokens[symbol] = token
			s.symbols[market] = symbol
			m.Unlock()
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
	tokens := s.tokens
	symbols := s.symbols
	ctx := context.Background()

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

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			fmt.Printf("%vth sync below 1.2 start @ %v\n", count, time.Now())
			count++
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

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			fmt.Printf("%vth sync below 3.0 start @ %v\n", count, time.Now())
			count++
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

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			fmt.Printf("%vth sync above 3.0 start @ %v\n", count, time.Now())
			count++
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
	symbols := s.symbols
	tokens := s.tokens
	vbep20s := s.vbep20s

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
	symbols := s.symbols
	tokens := s.tokens
	vbep20s := s.vbep20s

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
	db.Put(dbm.BorrowersStoreKey(accountBytes), accountBytes, nil)
}
