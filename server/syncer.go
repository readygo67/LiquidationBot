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
	Address            common.Address
	CollateralFactor   decimal.Decimal
	Price              decimal.Decimal
	UnderlyingDecimals uint8
}

type Asset struct {
	Symbol  string
	Balance decimal.Decimal
	Loan    decimal.Decimal
}

type AssetWithPrice struct {
	Symbol string
	//UnderlyingDecimals uint8
	Balance      decimal.Decimal
	Loan         decimal.Decimal
	Price        decimal.Decimal
	ExchangeRate decimal.Decimal
}

type AccountInfo struct {
	HealthFactor decimal.Decimal
	Assets       []Asset
}

type FeededPrice struct {
	Address common.Address
	Price   decimal.Decimal
	Hash    common.Hash
}

type FeededPrices struct {
	Prices []FeededPrice
	Height uint64
}

type Liquidation struct {
	Address      common.Address
	HealthFactor decimal.Decimal
	BlockNumber  uint64
	Endtime      time.Time
}

type semaphore chan struct{}

var (
	EXPSACLE         = decimal.New(1, 18)
	ExpScaleFloat, _ = big.NewFloat(0).SetString("1000000000000000000")
	BigZero          = big.NewInt(0)
	Decimal1P0, _    = decimal.NewFromString("1.0")
	Decimal1P1, _    = decimal.NewFromString("1.1")
	Decimal1P5, _    = decimal.NewFromString("1.5")
	Decimal2P0, _    = decimal.NewFromString("2.0")
	Decimal3P0, _    = decimal.NewFromString("3.0")
)

type Syncer struct {
	c  *ethclient.Client
	db *leveldb.DB

	oracle              *venus.Oracle
	comptroller         *venus.Comptroller
	closeFactor         decimal.Decimal
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

	bigCloseFactor, err := comptroller.CloseFactorMantissa(nil)
	if err != nil {
		panic(err)
	}
	closeFactor := decimal.NewFromBigInt(bigCloseFactor, 0)

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

			//underlyingAddress, err := vbep20.Underlying(nil)
			//if err != nil {
			//	panic(err)
			//}
			//
			//bep20, err := venus.NewVbep20(underlyingAddress, c)
			//underlyingDecimals, err := bep20.Decimals(nil)
			//if err != nil {
			//	panic(err)
			//}

			marketDetail, err := comptroller.Markets(nil, market)
			if err != nil {
				panic(err)
			}

			bigPrice, err := oracle.GetUnderlyingPrice(nil, market)
			if err != nil {
				bigPrice = big.NewInt(0)
			}

			collaterFactor := decimal.NewFromBigInt(marketDetail.CollateralFactorMantissa, 0)
			price := decimal.NewFromBigInt(bigPrice, 0)

			token := TokenInfo{
				Address: market,
				//UnderlyingDecimals: underlyingDecimals,
				CollateralFactor: collaterFactor,
				Price:            price,
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
		closeFactor:          closeFactor,
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

	bigCloseFactor, err := comptroller.CloseFactorMantissa(nil)
	if err != nil {
		return
	}
	s.closeFactor = decimal.NewFromBigInt(bigCloseFactor, 0)

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

			bigPrice, err := oracle.GetUnderlyingPrice(nil, market)
			if err != nil {
				bigPrice = big.NewInt(0)
			}

			token := TokenInfo{
				Address:          market,
				CollateralFactor: decimal.NewFromBigInt(marketDetail.CollateralFactorMantissa, 0),
				Price:            decimal.NewFromBigInt(bigPrice, 0),
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

	totalCollateral := decimal.NewFromInt(0)
	totalLoan := decimal.NewFromInt(0)

	var assets []Asset
	bigMintedVAIS, err := comptroller.MintedVAIs(nil, account)
	if err != nil {
		fmt.Printf("syncOneAccount, fail to get MintedVAIs, err:%v\n", err)
		return err
	}

	mintedVAIS := decimal.NewFromBigInt(bigMintedVAIS, 0)
	//mintVAISFloatExp := big.NewFloat(0).SetInt(mintVAIS)
	//mintVAISFloat := big.NewFloat(0).Quo(mintVAISFloatExp, ExpScaleFloat)
	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil {
		fmt.Printf("syncOneAccount, fail to get GetAssetsIn, err:%v\n", err)
		return err
	}

	for _, market := range markets {
		errCode, bigBalance, bigBorrow, bigExchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil {
			fmt.Printf("syncOneAccount, fail to get GetAccountSnapshot, err:%v\n", err)
			return err
		}

		if errCode.Cmp(BigZero) != 0 {
			fmt.Printf("syncOneAccount, fail to get GetAccountSnapshot, errCode:%v\n", errCode)
			return err
		}

		if bigBalance.Cmp(BigZero) == 0 && bigBorrow.Cmp(BigZero) == 0 {
			continue
		}

		symbol := symbols[market]
		collateralFactor := tokens[symbol].CollateralFactor
		price := tokens[symbol].Price

		exchangeRate := decimal.NewFromBigInt(bigExchangeRate, 0)
		balance := decimal.NewFromBigInt(bigBalance, 0)
		borrow := decimal.NewFromBigInt(bigBorrow, 0)

		multiplier := collateralFactor.Mul(exchangeRate).Div(EXPSACLE)
		multiplier = multiplier.Mul(price).Div(EXPSACLE)
		collateral := balance.Mul(multiplier).Div(EXPSACLE)
		totalCollateral = totalCollateral.Add(collateral.Truncate(0))

		loan := borrow.Mul(price).Div(EXPSACLE)
		totalLoan = totalLoan.Add(loan.Truncate(0))

		//build account table
		assets = append(assets, Asset{
			Symbol:  symbol,
			Balance: balance,
			Loan:    borrow,
		})
	}

	totalLoan = totalLoan.Add(mintedVAIS)
	healthFactor := decimal.New(100, 0)
	if totalLoan.Cmp(decimal.Zero) == 1 {
		healthFactor = totalCollateral.Div(totalLoan)
	}
	fmt.Printf("accout:%v, healthFactor:%v, totalCollateral:%v, totalLoan:%v\n", account, healthFactor, totalCollateral, totalLoan)
	//update market table and account table
	info := AccountInfo{
		HealthFactor: healthFactor,
		Assets:       assets,
	}
	s.updateDB(account, info)
	if healthFactor.Cmp(decimal.New(1, 0)) != 1 {
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

	totalCollateral := decimal.NewFromInt(0)
	totalLoan := decimal.NewFromInt(0)

	bigMintedVAIS, err := comptroller.MintedVAIs(nil, account)
	if err != nil {
		return err
	}

	mintedVAIS := decimal.NewFromBigInt(bigMintedVAIS, 0)

	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil {
		return err
	}

	for _, market := range markets {
		//fmt.Printf("market:%v\n", market)
		errCode, bigBalance, bigBorrow, bigExchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil {
			fmt.Printf("syncOneAccount, fail to get GetAccountSnapshot, err:%v\n", err)
			return err
		}

		if errCode.Cmp(BigZero) != 0 {
			fmt.Printf("syncOneAccount, fail to get GetAccountSnapshot, errCode:%v\n", errCode)
			return err
		}

		if bigBalance.Cmp(BigZero) == 0 && bigBorrow.Cmp(BigZero) == 0 {
			continue
		}

		symbol := symbols[market]
		collateralFactor := tokens[symbol].CollateralFactor
		price := tokens[symbol].Price

		//apply feeded prices if exist
		for _, feededPrice := range feededPrices.Prices {
			if tokens[symbol].Address == feededPrice.Address {
				price = feededPrice.Price
			}
		}

		exchangeRate := decimal.NewFromBigInt(bigExchangeRate, 0)
		balance := decimal.NewFromBigInt(bigBalance, 0)
		borrow := decimal.NewFromBigInt(bigBorrow, 0)

		multiplier := collateralFactor.Mul(exchangeRate).Div(EXPSACLE)
		multiplier = multiplier.Mul(price).Div(EXPSACLE)
		collateral := balance.Mul(multiplier).Div(EXPSACLE)
		totalCollateral = totalCollateral.Add(collateral.Truncate(0))

		loan := borrow.Mul(price).Div(EXPSACLE)
		totalLoan = totalLoan.Add(loan.Truncate(0))
	}

	totalLoan = totalLoan.Add(mintedVAIS)
	healthFactor := decimal.NewFromInt(100)
	if totalLoan.Cmp(decimal.Zero) == 1 {
		healthFactor = totalCollateral.Div(totalLoan)
	}
	fmt.Printf("account:%v, healthFactor:%v, totalCollateral:%v, totalLoan:%v \n", account, healthFactor, totalCollateral, totalLoan)

	if healthFactor.Cmp(decimal.NewFromInt(1)) != 1 {
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

func (s *Syncer) calculateSeizedTokenAmount(liquidation *Liquidation) error {
	comptroller := s.comptroller
	oracle := s.oracle
	account := liquidation.Address

	s.m.Lock()
	symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	vbep20s := s.vbep20s
	closeFactor := s.closeFactor
	s.m.Unlock()

	errCode, _, shortfall, err := comptroller.GetAccountLiquidity(nil, account)
	if err != nil {
		fmt.Printf("calculateSeizedTokenAmount, fail to get GetAccountLiquidity, account:%v, err:%v\n", account, err)
		return err
	}
	if errCode.Cmp(BigZero) != 0 {
		fmt.Printf("calculateSeizedTokenAmount, fail to get GetAccountLiquidity, account:%v, errCode:%v\n", account, errCode)
		return err
	}

	if shortfall.Cmp(BigZero) != 1 {
		err := fmt.Errorf("calculateSeizedTokenAmount, no shortfall, account:%v", account)
		return err
	}

	totalCollateral := decimal.NewFromInt(0)
	totalLoan := decimal.NewFromInt(0)
	var assets []AssetWithPrice
	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil {
		fmt.Printf("calculateSeizedTokenAmount, fail to get GetAssetsIn, account:%v, err:%v\n", account, err)
		return err
	}

	for _, market := range markets {
		errCode, bigBalance, bigBorrow, bigExchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil {
			fmt.Printf("calculateSeizedTokenAmount, fail to get GetAccountSnapshot, account:%v, err:%v\n", account, err)
			return err
		}

		if errCode.Cmp(BigZero) != 0 {
			fmt.Printf("calculateSeizedTokenAmount, fail to get GetAccountSnapshot, account:%v, errCode:%v\n", account, errCode)
			return err
		}

		if bigBalance.Cmp(BigZero) == 0 && bigBorrow.Cmp(BigZero) == 0 {
			continue
		}

		fmt.Printf("balance:%v, borrow:%v, exchangeRate:%v\n", bigBalance, bigBorrow, bigExchangeRate)
		bigPrice, err := oracle.GetUnderlyingPrice(nil, market)
		if err != nil {
			fmt.Printf("calculateSeizedTokenAmount, fail to get underlying price, account:%v, err:%v\n", account, err)
			return err
		}

		price := decimal.NewFromBigInt(bigPrice, 0)
		exchangeRate := decimal.NewFromBigInt(bigExchangeRate, 0)
		balance := decimal.NewFromBigInt(bigBalance, 0)
		borrow := decimal.NewFromBigInt(bigBorrow, 0)

		symbol := symbols[market]
		collateralFactor := tokens[symbol].CollateralFactor
		//decimals := tokens[symbol].UnderlyingDecimals

		multiplier := collateralFactor.Mul(exchangeRate).Div(EXPSACLE)
		multiplier = multiplier.Mul(price).Div(EXPSACLE)
		collateralD := balance.Mul(multiplier).Div(EXPSACLE)
		totalCollateral = totalCollateral.Add(collateralD.Truncate(0))

		loan := borrow.Mul(price).Div(EXPSACLE)
		totalLoan = totalLoan.Add(loan.Truncate(0))

		asset := AssetWithPrice{
			Symbol: symbol,
			//UnderlyingDecimals: decimals,
			Balance:      collateralD,
			Loan:         loan,
			Price:        price,
			ExchangeRate: exchangeRate,
		}
		fmt.Printf("asset:%+v, address:%v\n", asset, tokens[asset.Symbol].Address)
		assets = append(assets, asset)
	}

	fmt.Printf("account:%v, totalCollateral:%v, totalLoan:%v, calculatedshortfall:%v, shorfall:%v\n", account, totalCollateral, totalLoan, totalLoan.Sub(totalCollateral), shortfall)
	//select the repayed token and seized collateral token
	maxLoanValue := decimal.NewFromInt(0)
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
	maxRepayValue := maxLoanValue.Mul(closeFactor).Div(EXPSACLE)
	repaySymbol := maxLoanSymbol

	repayValue := decimal.NewFromInt(0)
	seizedSymbol := ""
	seizedIndex := 0
	for k, asset := range assets {
		if asset.Balance.Cmp(maxRepayValue) != -1 {
			repayValue = maxRepayValue
			seizedSymbol = asset.Symbol
			seizedIndex = k
			break
		} else {
			if asset.Balance.Cmp(repayValue) == 1 {
				repayValue = asset.Balance
				seizedSymbol = asset.Symbol
				seizedIndex = k
			}
		}
	}

	repayAmout := repayValue.Mul(EXPSACLE).Div(assets[repayIndex].Price)
	repayAmount := repayAmout.Truncate(0)
	errCode, bigSeizedCTokenAmount, err := comptroller.LiquidateCalculateSeizeTokens(nil, tokens[repaySymbol].Address, tokens[seizedSymbol].Address, repayAmount.BigInt())

	seizedCTokenAmount := decimal.NewFromBigInt(bigSeizedCTokenAmount, 0)
	seizedUnderlyingTokenAmount := seizedCTokenAmount.Mul(assets[seizedIndex].ExchangeRate).Div(EXPSACLE)
	seizedUnderlyingTokenValue := seizedUnderlyingTokenAmount.Mul(assets[seizedIndex].Price).Div(EXPSACLE)
	fmt.Printf("account:%v,repaySmbol:%v repayValue:%v, repayAmount:%v seizedSymbol:%v, seizedCTokenAmount:%v, seizedUnderlyingTokenAmount:%v, seizedUnderlyingTokenValue:%v\n", account, repaySymbol, repayValue, repayAmount, seizedSymbol, seizedCTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue)

	ratio := seizedUnderlyingTokenValue.Div(repayValue)
	fmt.Printf("ratio:%v\n", ratio)

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

		if healthFactor.Cmp(Decimal1P0) == -1 {
			db.Delete(dbm.LiquidationBelow1P0StoreKey(accountBytes), nil)
		} else if healthFactor.Cmp(Decimal1P1) == -1 {
			db.Delete(dbm.LiquidationBelow1P1StoreKey(accountBytes), nil)
		} else if healthFactor.Cmp(Decimal1P5) == -1 {
			db.Delete(dbm.LiquidationBelow1P5StoreKey(accountBytes), nil)
		} else if healthFactor.Cmp(Decimal2P0) == -1 {
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

	if healthFactor.Cmp(Decimal1P0) == -1 {
		db.Put(dbm.LiquidationBelow1P0StoreKey(accountBytes), accountBytes, nil)
	} else if healthFactor.Cmp(Decimal1P1) == -1 {
		db.Put(dbm.LiquidationBelow1P1StoreKey(accountBytes), accountBytes, nil)
	} else if healthFactor.Cmp(Decimal1P5) == -1 {
		db.Put(dbm.LiquidationBelow1P5StoreKey(accountBytes), accountBytes, nil)
	} else if healthFactor.Cmp(Decimal2P0) == -1 {
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
