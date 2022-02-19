package server

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	UnderlyingAddress  common.Address
	UnderlyingDecimals uint8
	CollateralFactor   decimal.Decimal
	Price              decimal.Decimal
}

type Asset struct {
	Symbol  string
	Balance decimal.Decimal
	Loan    decimal.Decimal
}

type AssetWithPrice struct {
	Symbol           string
	CollateralFactor decimal.Decimal
	Balance          decimal.Decimal
	Collateral       decimal.Decimal
	Loan             decimal.Decimal
	Price            decimal.Decimal
	ExchangeRate     decimal.Decimal
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
	vBNBAddress      = common.HexToAddress("0xA07c5b74C9B40447a954e1466938b865b6BBea36")
	wBNBAddress      = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	VAIAddress       = common.HexToAddress("0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7")
)

type Syncer struct {
	c  *ethclient.Client
	db *leveldb.DB

	oracle               *venus.Oracle
	comptroller          *venus.Comptroller
	pancakeRouter        *venus.IPancakeRouter02
	pancakeFactory       *venus.IPancakeFactory
	closeFactor          decimal.Decimal
	symbols              map[common.Address]string
	tokens               map[string]*TokenInfo
	vbep20s              map[common.Address]*venus.Vbep20
	liquidator           *venus.IQingsuan
	PrivateKey           *ecdsa.PrivateKey
	m                    sync.Mutex
	wg                   sync.WaitGroup
	quitCh               chan struct{}
	forceUpdatePricesCh  chan struct{}
	feededPricesCh       chan *FeededPrices
	liquidationCh        chan *Liquidation
	priortyLiquidationCh chan *Liquidation
}

func init() {

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
	pancakeRouterAddress string,
	liquidatorAddress string,
	privatekey string,
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

	pancakeRouter, err := venus.NewIPancakeRouter02(common.HexToAddress(pancakeRouterAddress), c)
	if err != nil {
		panic(err)
	}

	factoryAddress, err := pancakeRouter.Factory(nil)
	if err != nil {
		panic(err)
	}

	pancakeFactory, err := venus.NewIPancakeFactory(factoryAddress, c)
	if err != nil {
		panic(err)
	}

	liquidator, err := venus.NewIQingsuan(common.HexToAddress(liquidatorAddress), c)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		panic(err)
	}

	markets, err := comptroller.GetAllMarkets(nil)
	if err != nil {
		panic(err)
	}

	symbols := make(map[common.Address]string)
	tokens := make(map[string]*TokenInfo)
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

			bigPrice, err := oracle.GetUnderlyingPrice(nil, market)
			if err != nil {
				bigPrice = big.NewInt(0)
			}
			var underlyingAddress common.Address
			if market == vBNBAddress {
				underlyingAddress = wBNBAddress
			} else {
				underlyingAddress, err = vbep20.Underlying(nil)
				if err != nil {
					panic(err)
				}
			}

			bep20, err := venus.NewVbep20(underlyingAddress, c)
			underlyingDecimals, err := bep20.Decimals(nil)
			if err != nil {
				panic(err)
			}

			collaterFactor := decimal.NewFromBigInt(marketDetail.CollateralFactorMantissa, 0)
			price := decimal.NewFromBigInt(bigPrice, 0)

			token := &TokenInfo{
				Address:            market,
				UnderlyingAddress:  underlyingAddress,
				UnderlyingDecimals: underlyingDecimals,
				CollateralFactor:   collaterFactor,
				Price:              price,
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
		pancakeRouter:        pancakeRouter,
		pancakeFactory:       pancakeFactory,
		closeFactor:          closeFactor,
		tokens:               tokens,
		symbols:              symbols,
		vbep20s:              vbep20s,
		liquidator:           liquidator,
		PrivateKey:           privateKey,
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

	s.wg.Add(9)
	go s.syncMarketsAndPrices()
	go s.feedPrices()
	go s.syncAllBorrowers()
	go s.syncLiquidationBelow1P0()
	go s.syncLiquidationBelow1P1()
	go s.syncLiquidationBelow1P5()
	go s.syncLiquidationBelow2P0()
	go s.syncLiquidationAbove2P0()
	go s.liqudate()
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

			s.m.Lock()
			defer s.m.Unlock()
			if s.tokens[symbol] != nil {
				s.tokens[symbol].CollateralFactor = decimal.NewFromBigInt(marketDetail.CollateralFactorMantissa, 0)
				s.tokens[symbol].Price = decimal.NewFromBigInt(bigPrice, 0)
			} else {
				fmt.Printf("a new symbol:%v added, please restart venusd\n", symbol)
			}

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

			//fmt.Printf("startHeight:%v, endHeight:%v\n", startHeight, endHeight)
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
	//fmt.Printf("accout:%v, healthFactor:%v, totalCollateral:%v, totalLoan:%v\n", account, healthFactor, totalCollateral, totalLoan)
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

func (s *Syncer) liqudate() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case pending := <-s.priortyLiquidationCh:
			fmt.Printf("receive priority liquidation:%v\n", pending)
			s.calculateSeizedTokenAmount(pending)
			//
			//pending.Endtime = time.Now().Add(time.Second * PriorityLiquidationTime)
			//priorityPendings = append(priorityPendings, pending)

		case pending := <-s.liquidationCh:
			fmt.Printf("recive liquidation:%v\n", pending)
			s.calculateSeizedTokenAmount(pending)
			//
			//pending.Endtime = time.Now().Add(time.Second * NormalLiquidationTime)
			//pendings = append(pendings, pending)
		}
	}
}

func (s *Syncer) calculateSeizedTokenAmount(liquidation *Liquidation) error {
	comptroller := s.comptroller
	oracle := s.oracle
	pancakeRouter := s.pancakeRouter
	pancakeFactory := s.pancakeFactory
	account := liquidation.Address

	s.m.Lock()
	symbols := copySymbols(s.symbols)
	tokens := copyTokens(s.tokens)
	vbep20s := s.vbep20s
	closeFactor := s.closeFactor
	s.m.Unlock()
	//current height
	currentHeight, err := s.c.BlockNumber(context.Background())
	if err != nil {
		fmt.Printf("calculateSeizedTokenAmount, fail to get blocknumber,err:%v\n", err)
		return err
	}
	callOptions := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(currentHeight)),
	}

	errCode, _, shortfall, err := comptroller.GetAccountLiquidity(callOptions, account)
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

	bigMintedVAIS, err := comptroller.MintedVAIs(nil, account)
	if err != nil {
		fmt.Printf("calculateSeizedTokenAmount, fail to get MintedVAIs, account:%v, err:%v\n", account, err)
		return err
	}
	mintedVAIS := decimal.NewFromBigInt(bigMintedVAIS, 0)

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

		bigPrice, err := oracle.GetUnderlyingPrice(callOptions, market)
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

		multiplier := collateralFactor.Mul(exchangeRate).Div(EXPSACLE)
		multiplier = multiplier.Mul(price).Div(EXPSACLE)
		collateralD := balance.Mul(multiplier).Div(EXPSACLE)
		totalCollateral = totalCollateral.Add(collateralD.Truncate(0))

		loan := borrow.Mul(price).Div(EXPSACLE)
		totalLoan = totalLoan.Add(loan.Truncate(0))

		asset := AssetWithPrice{
			Symbol:           symbol,
			CollateralFactor: collateralFactor.Div(EXPSACLE),
			Balance:          balance.Mul(exchangeRate).Mul(price).Div(EXPSACLE).Div(EXPSACLE),
			Collateral:       collateralD,
			Loan:             loan,
			Price:            price,
			ExchangeRate:     exchangeRate,
		}
		//fmt.Printf("asset:%+v, address:%v\n", asset, tokens[asset.Symbol].Address)
		assets = append(assets, asset)
	}
	totalLoan = totalLoan.Add(mintedVAIS)
	fmt.Printf("account:%v, totalCollateralValue:%v, mintedVAISValue:%v, totalLoanValue:%v, calculatedshortfall:%v, shorfall:%v\n", account, totalCollateral.Div(EXPSACLE), mintedVAIS.Div(EXPSACLE), totalLoan.Div(EXPSACLE), (totalLoan.Sub(totalCollateral)), shortfall)
	//select the repayed token and seized collateral token
	maxLoanValue := decimal.NewFromInt(0)
	maxLoanSymbol := ""
	//repayIndex := math.MaxInt32

	for _, asset := range assets {
		if asset.Loan.Cmp(maxLoanValue) == 1 {
			maxLoanValue = asset.Loan
			maxLoanSymbol = asset.Symbol
			//repayIndex = i
		}
	}

	if mintedVAIS.Cmp(maxLoanValue) == 1 {
		maxLoanValue = mintedVAIS
		maxLoanSymbol = "VAI"
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

	var bigSeizedCTokenAmount *big.Int
	var repayAmount decimal.Decimal
	if repaySymbol == "VAI" {
		repayAmount = repayValue.Truncate(0)

		errCode, bigSeizedCTokenAmount, err = comptroller.LiquidateVAICalculateSeizeTokens(callOptions, tokens[seizedSymbol].Address, repayAmount.BigInt())
		if err != nil {
			fmt.Printf("calculateSeizedTokenAmount, fail to get LiquidateVAICalculateSeizeTokens, account:%v, err:%v\n", account, err)
			return err
		}
		if errCode.Cmp(BigZero) != 0 {
			fmt.Printf("calculateSeizedTokenAmount, fail to get LiquidateVAICalculateSeizeTokens, account:%v, errCode:%v\n", account, errCode)
			return fmt.Errorf("%v", errCode)
		}
	} else {
		bigBorrowBalanceStored, err := vbep20s[tokens[repaySymbol].Address].BorrowBalanceStored(callOptions, account)
		if err != nil {
			fmt.Printf("calculateSeizedTokenAmount, fail to get BorrowBalanceStored, account:%v, err:%v\n", account, err)
			return err
		}
		repayAmount = decimal.NewFromBigInt(bigBorrowBalanceStored, 0).Mul(closeFactor).Div(EXPSACLE) //repayValue.Mul(EXPSACLE).Div(assets[repayIndex].Price)
		repayAmount = repayAmount.Truncate(0).Sub(decimal.NewFromInt(1000))                           //to avoid TOO_MUCH_REPAY error

		errCode, bigSeizedCTokenAmount, err = comptroller.LiquidateCalculateSeizeTokens(callOptions, tokens[repaySymbol].Address, tokens[seizedSymbol].Address, repayAmount.BigInt())
		if err != nil {
			fmt.Printf("calculateSeizedTokenAmount, fail to get LiquidateCalculateSeizeTokens, account:%v, err:%v\n", account, err)
			return err
		}
		if errCode.Cmp(BigZero) != 0 {
			fmt.Printf("calculateSeizedTokenAmount, fail to get LiquidateCalculateSeizeTokens, account:%v, errCode:%v\n", account, errCode)
			return fmt.Errorf("%v", errCode)
		}
	}

	seizedCTokenAmount := decimal.NewFromBigInt(bigSeizedCTokenAmount, 0)
	seizedUnderlyingTokenAmount := seizedCTokenAmount.Mul(assets[seizedIndex].ExchangeRate).Div(EXPSACLE)
	seizedUnderlyingTokenValue := seizedUnderlyingTokenAmount.Mul(assets[seizedIndex].Price).Div(EXPSACLE)

	var flashLoanFrom common.Address
	switch repaySymbol {
	case "VAI":
		flashLoanFrom, err = pancakeFactory.GetPair(nil, VAIAddress, tokens["vBNB"].UnderlyingAddress)
		fmt.Printf("height%v, account:%v, repaySmbol:%v, flashLoanFrom:%v, repayAddress:%v, repayValue:%v, repayAmount:%v seizedSymbol:%v, seizedAddress:%v, seizedCTokenAmount:%v, seizedUnderlyingTokenAmount:%v, seizedUnderlyingTokenValue:%v\n", currentHeight, account, repaySymbol, flashLoanFrom, VAIAddress, repayValue, repayAmount, seizedSymbol, tokens[seizedSymbol].Address, tokens[seizedSymbol].UnderlyingAddress, seizedCTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue)

	case "vBNB":
		flashLoanFrom, err = pancakeFactory.GetPair(nil, tokens[repaySymbol].UnderlyingAddress, tokens["vUSDT"].UnderlyingAddress)
		fmt.Printf("height%v, account:%v, repaySmbol:%v, flashLoanFrom:%v, repayAddress:%v, repayUnderlyingAddress:%v, repayValue:%v, repayAmount:%v seizedSymbol:%v, seizedAddress:%v, seizdUnderLyingAddress:%v, seizedCTokenAmount:%v, seizedUnderlyingTokenAmount:%v, seizedUnderlyingTokenValue:%v\n", currentHeight, account, repaySymbol, flashLoanFrom, tokens[repaySymbol].Address, tokens[repaySymbol].UnderlyingAddress, repayValue, repayAmount, seizedSymbol, tokens[seizedSymbol].Address, tokens[seizedSymbol].UnderlyingAddress, seizedCTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue)

	default:
		flashLoanFrom, err = pancakeFactory.GetPair(nil, tokens[repaySymbol].UnderlyingAddress, tokens["vBNB"].UnderlyingAddress)
		fmt.Printf("height%v, account:%v, repaySmbol:%v, flashLoanFrom:%v, repayAddress:%v, repayUnderlyingAddress:%v, repayValue:%v, repayAmount:%v seizedSymbol:%v, seizedAddress:%v, seizdUnderLyingAddress:%v, seizedCTokenAmount:%v, seizedUnderlyingTokenAmount:%v, seizedUnderlyingTokenValue:%v\n", currentHeight, account, repaySymbol, flashLoanFrom, tokens[repaySymbol].Address, tokens[repaySymbol].UnderlyingAddress, repayValue, repayAmount, seizedSymbol, tokens[seizedSymbol].Address, tokens[seizedSymbol].UnderlyingAddress, seizedCTokenAmount, seizedUnderlyingTokenAmount, seizedUnderlyingTokenValue)

	}

	if err != nil {
		fmt.Printf("calculateSeizedTokenAmount, fail to get %v flashLoanFrom, account:%v, err:%v\n", repaySymbol, account, err)
		return err
	}

	ratio := seizedUnderlyingTokenValue.Div(repayValue)
	if ratio.Cmp(decimal.NewFromFloat32(1.11)) == 1 || ratio.Cmp(decimal.NewFromFloat32(1.09)) == -1 {
		fmt.Printf("calculated seizedUnerlyingTokenValue != 1.1, calculateRatio:%v, seizedUnderylingTokenValue:%v, repayValue:%v\n", ratio, seizedUnderlyingTokenValue, repayValue)
		err := fmt.Errorf("calculated seizedUnerlyingTokenValue != 1.1, calculateRatio:%v, seizedUnderylingTokenValue:%v, repayValue:%v\n", ratio, seizedUnderlyingTokenValue, repayValue)
		return err
	}

	flashLoanFeeRatio := decimal.NewFromInt(25).Div(decimal.NewFromInt(9975))
	flashLoanFeeAmount := repayAmount.Mul(flashLoanFeeRatio).Add(decimal.NewFromInt(1))
	flashLoanReturnAmount := repayAmount.Add(flashLoanFeeAmount.Truncate(0))

	bigGasPrice, err := s.c.SuggestGasPrice(context.Background())
	if err != nil {
		bigGasPrice = big.NewInt(5)
	}
	gasPrice := decimal.NewFromBigInt(bigGasPrice, 0).Mul(decimal.NewFromFloat32(1.5)) //x1.5 gasPrice for PGA
	bigGas := decimal.NewFromInt(700000)
	ethPrice := tokens["vBNB"].Price

	if repaySymbol == "VAI" {
		addresses := []common.Address{
			VAIAddress,
			tokens[seizedSymbol].Address,
			tokens[seizedSymbol].UnderlyingAddress,
			VAIAddress,
			account,
		}

		if isStalbeCoin(seizedSymbol) {
			//case6, repay VAI and get stablecoin
			bigGas = decimal.NewFromInt(800000)
			gasFee := decimal.NewFromBigInt(bigGasPrice, 0).Mul(bigGas).Mul(ethPrice).Div(EXPSACLE)

			//paths := s.uniswapPaths[seizedSymbol+repaySymbol]
			path := s.buildVAIPaths(seizedSymbol, repaySymbol, tokens)
			amountsIn, err := pancakeRouter.GetAmountsIn(callOptions, flashLoanReturnAmount.BigInt(), path) //amountsIn[0] is the stablecoin needed.
			if err != nil {
				fmt.Printf("calculateSeizedTokenAmount case6, fail to get GetAmountsIn, account:%v, paths:%v, amountout:%v, err:%v\n", account, path, flashLoanReturnAmount, err)
				return err
			}

			remain := seizedUnderlyingTokenAmount.Sub(decimal.NewFromBigInt(amountsIn[0], 0))
			profit := remain.Mul(tokens[seizedSymbol].Price).Div(EXPSACLE).Sub(gasFee)
			fmt.Printf("calculateSeizedTokenAmount case6: repaySymbol is VAI and seizedSymbol is stable coin, account:%v, seizedsymbol:%v, seizedAmount:%v, repaySymbol:%v, returnAmout:%v, remain:%v, gasFee:%v, profit:%v\n", account, seizedSymbol, seizedUnderlyingTokenAmount, repaySymbol, amountsIn[0], remain, gasFee, profit.Div(EXPSACLE))

			if profit.Cmp(decimal.Zero) == 1 {
				fmt.Printf("case6, profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				s.doLiquidation(big.NewInt(6), flashLoanFrom, path, nil, addresses, repayAmount.BigInt(), gasPrice.BigInt())

			}
		} else {
			//case7,  repay VAI and seizedSymbol is not stable coin. sell partly seizedSymbol to repay symbol, sell remain to usdt
			bigGas = decimal.NewFromInt(1000000)
			gasFee := decimal.NewFromBigInt(bigGasPrice, 0).Mul(bigGas).Mul(ethPrice).Div(EXPSACLE)
			path1 := s.buildVAIPaths(seizedSymbol, repaySymbol, tokens)
			amountIns, err := pancakeRouter.GetAmountsIn(callOptions, flashLoanReturnAmount.BigInt(), path1)
			if err != nil {
				fmt.Printf("calculateSeizedTokenAmount case7, fail to get GetAmountsIn, account:%V, paths:%v, amountout:%v, err:%v\n", account, path1, flashLoanReturnAmount, err)
				return err
			}
			fmt.Printf("calculateSeizedTokenAmount case7, account:%v, paths:%+v, swap %v%v for %v%v\n", account, path1, amountIns[0], strings.TrimPrefix(seizedSymbol, "v"), flashLoanFeeAmount.Truncate(0), strings.TrimPrefix(repaySymbol, "v"))

			// swap the remains to usdt
			remain := seizedUnderlyingTokenAmount.Truncate(0).Sub(decimal.NewFromBigInt(amountIns[0], 0))
			path2 := s.buildPaths(seizedSymbol, "vUSDT", tokens)
			amountsOut, err := pancakeRouter.GetAmountsOut(callOptions, remain.Truncate(0).BigInt(), path2)
			if err != nil {
				fmt.Printf("calculateSeizedTokenAmount case7, fail to get GetAmountsOut, account:%v paths:%v, err:%v\n", account, path2, err)
				return err
			}

			usdtAmount := decimal.NewFromBigInt(amountsOut[len(amountsOut)-1], 0)
			profit := usdtAmount.Mul(tokens["vUSDT"].Price).Div(EXPSACLE).Sub(gasFee)

			fmt.Printf("calculateSeizedTokenAmount case7, account:%v, path:%v, swap %v%v for %vUSDT, profit:%v\n", account, path2, remain, strings.TrimPrefix(seizedSymbol, "v"), amountsOut[len(amountsOut)-1], profit.Div(EXPSACLE))
			if profit.Cmp(decimal.Zero) == 1 {
				fmt.Printf("case7: profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				s.doLiquidation(big.NewInt(7), flashLoanFrom, path1, path2, addresses, repayAmount.BigInt(), gasPrice.BigInt())

			}
		}
		return nil
	}

	addresses := []common.Address{
		tokens[repaySymbol].Address,
		tokens[seizedSymbol].Address,
		tokens[seizedSymbol].UnderlyingAddress,
		tokens[repaySymbol].UnderlyingAddress,
		account,
	}

	if seizedSymbol == repaySymbol {
		if isStalbeCoin(seizedSymbol) {
			//case1, seizedSymbol == repaySymbol and symbol is a stable coin
			bigGas = decimal.NewFromInt(600000)
			gasFee := gasPrice.Mul(bigGas).Mul(ethPrice).Div(EXPSACLE)
			profit := ((seizedUnderlyingTokenAmount.Sub(flashLoanReturnAmount)).Mul(tokens[seizedSymbol].Price).Div(EXPSACLE)).Sub(gasFee)

			fmt.Printf("calculateSeizedTokenAmount case1: seizedSymbol == repaySymbol and symbol is a stable coin, account:%v, symbol:%v, seizedAmount:%v, returnAmout:%v, gasFee:%v profit:%v\n", account, seizedSymbol, seizedUnderlyingTokenAmount, flashLoanReturnAmount, gasFee, profit.Div(EXPSACLE))
			if profit.Cmp(decimal.Zero) == 1 {
				fmt.Printf("case1, profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				s.doLiquidation(big.NewInt(1), flashLoanFrom, nil, nil, addresses, repayAmount.BigInt(), gasPrice.BigInt())

			}
		} else {
			//case2,  seizedSymbol == repaySymbol and symbol is not a stable coin, after return flashloan, sell remain to usdt
			bigGas = decimal.NewFromInt(800000)
			gasFee := gasPrice.Mul(bigGas).Mul(ethPrice).Div(EXPSACLE)
			remain := seizedUnderlyingTokenAmount.Sub(flashLoanReturnAmount)

			path2 := s.buildPaths(seizedSymbol, "vUSDT", tokens)
			amountsOut, err := pancakeRouter.GetAmountsOut(callOptions, remain.Truncate(0).BigInt(), path2)
			if err != nil {
				fmt.Printf("calculateSeizedTokenAmount case2:, fail to get GetAmountsout, account:%v, paths:%v, amountIn:%v, err:%v\n", account, path2, remain.Truncate(0), err)
				return err
			}

			usdtAmount := decimal.NewFromBigInt(amountsOut[len(amountsOut)-1], 0)
			profit := (usdtAmount.Mul(tokens["vUSDT"].Price).Div(EXPSACLE)).Sub(gasFee)
			fmt.Printf("calculateSeizedTokenAmount case2: seizedSymbol == repaySymbol and symbol is not stable coin, account:%v, symbol:%v, seizedAmount:%v, returnAmout:%v, usdtAmount:%v, gasFee:%v, profit:%v\n", account, seizedSymbol, seizedUnderlyingTokenAmount, flashLoanReturnAmount, usdtAmount, gasFee, profit.Div(EXPSACLE))

			if profit.Cmp(decimal.Zero) == 1 {
				fmt.Printf("case2, profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				s.doLiquidation(big.NewInt(2), flashLoanFrom, nil, path2, addresses, repayAmount.BigInt(), gasPrice.BigInt())

			}
		}
		return nil
	}

	if isStalbeCoin(seizedSymbol) {
		//case3, collateral(i.e. seizedSymbol) is stable coin, repaySymbol may or not be a stable coin, sell part of seized symbol to repaySymbol
		bigGas = decimal.NewFromInt(900000)
		gasFee := decimal.NewFromBigInt(bigGasPrice, 0).Mul(bigGas).Mul(ethPrice).Div(EXPSACLE)

		//paths := s.uniswapPaths[seizedSymbol+repaySymbol]
		path1 := s.buildPaths(seizedSymbol, repaySymbol, tokens)
		amountsIn, err := pancakeRouter.GetAmountsIn(callOptions, flashLoanReturnAmount.BigInt(), path1) //amountsIn[0] is the stablecoin needed.
		if err != nil {
			fmt.Printf("calculateSeizedTokenAmount case3, fail to get GetAmountsIn, account:%v, paths:%v, amountout:%v, err:%v\n", account, path1, flashLoanReturnAmount, err)
			return err
		}

		remain := seizedUnderlyingTokenAmount.Sub(decimal.NewFromBigInt(amountsIn[0], 0))
		profit := remain.Mul(tokens[seizedSymbol].Price).Div(EXPSACLE).Sub(gasFee)
		fmt.Printf("calculateSeizedTokenAmount case3: seizedSymbol != repaySymbol and seizedSymbol stable coin, account:%v, seizedsymbol:%v, seizedAmount:%v, repaySymbol:%v, returnAmout:%v, remain:%v, gasFee:%v, profit:%v\n", account, seizedSymbol, seizedUnderlyingTokenAmount, repaySymbol, amountsIn[0], remain, gasFee, profit.Div(EXPSACLE))

		if profit.Cmp(decimal.Zero) == 1 {
			fmt.Printf("case3, profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
			s.doLiquidation(big.NewInt(3), flashLoanFrom, path1, nil, addresses, repayAmount.BigInt(), gasPrice.BigInt())
		}
	} else {
		if isStalbeCoin(repaySymbol) {
			//case4, collateral(i.e. seizedSymbol) is not stable coin, repaySymbol is a stable coin, sell all seizedSymbol to repaySymbol
			bigGas = decimal.NewFromInt(900000)
			gasFee := decimal.NewFromBigInt(bigGasPrice, 0).Mul(bigGas).Mul(ethPrice).Div(EXPSACLE)

			//paths := s.uniswapPaths[seizedSymbol+repaySymbol]
			path1 := s.buildPaths(seizedSymbol, repaySymbol, tokens)
			amountsOut, err := pancakeRouter.GetAmountsOut(callOptions, seizedUnderlyingTokenAmount.Truncate(0).BigInt(), path1)
			if err != nil {
				fmt.Printf("calculateSeizedTokenAmount case4, fail to get GetAmountsIn, account:%V, paths:%v, amountout:%v, err:%v\n", account, path1, flashLoanReturnAmount, err)
				return err
			}

			remain := decimal.NewFromBigInt(amountsOut[len(amountsOut)-1], 0).Sub(flashLoanReturnAmount)
			profit := remain.Mul(tokens[repaySymbol].Price).Div(EXPSACLE).Sub(gasFee)
			fmt.Printf("calculateSeizedTokenAmount case4: seizedSymbol is not stable coin, repaySymbol is stable coin, account:%v repaysymbol:%v, seizedsymbol:%v seizedAmount:%v, amountsOut:%v returnAmout:%v, remain:%v, gasFee:%v, profit:%v\n", account, repaySymbol, seizedSymbol, seizedUnderlyingTokenAmount, amountsOut[len(amountsOut)-1], flashLoanReturnAmount, remain, gasFee, profit.Div(EXPSACLE))

			if profit.Cmp(decimal.Zero) == 1 {
				fmt.Printf("case4, profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				s.doLiquidation(big.NewInt(4), flashLoanFrom, path1, nil, addresses, repayAmount.BigInt(), gasPrice.BigInt())
			}
		} else {
			//case5,  collateral(i.e. seizedSymbol) and repaySymbol are not stable coin. sell partly seizedSymbol to repay symbol, sell remain to usdt
			bigGas = decimal.NewFromInt(1000000)
			gasFee := decimal.NewFromBigInt(bigGasPrice, 0).Mul(bigGas).Mul(ethPrice).Div(EXPSACLE)
			path1 := s.buildPaths(seizedSymbol, repaySymbol, tokens)
			amountIns, err := pancakeRouter.GetAmountsIn(callOptions, flashLoanReturnAmount.BigInt(), path1)
			if err != nil {
				fmt.Printf("calculateSeizedTokenAmount case5, fail to get GetAmountsIn, account:%V, paths:%v, amountout:%v, err:%v\n", account, path1, flashLoanReturnAmount, err)
				return err
			}
			fmt.Printf("calculateSeizedTokenAmount case5, account:%v, paths:%+v, swap %v%v for %v%v\n", account, path1, amountIns[0], strings.TrimPrefix(seizedSymbol, "v"), flashLoanFeeAmount.Truncate(0), strings.TrimPrefix(repaySymbol, "v"))

			// swap the remains to usdt
			remain := seizedUnderlyingTokenAmount.Truncate(0).Sub(decimal.NewFromBigInt(amountIns[0], 0))
			path2 := s.buildPaths(seizedSymbol, "vUSDT", tokens)
			amountsOut, err := pancakeRouter.GetAmountsOut(callOptions, remain.Truncate(0).BigInt(), path2)
			if err != nil {
				fmt.Printf("calculateSeizedTokenAmount case5, fail to get GetAmountsOut, account:%v paths:%v, err:%v\n", account, path2, err)
				return err
			}

			usdtAmount := decimal.NewFromBigInt(amountsOut[len(amountsOut)-1], 0)
			profit := usdtAmount.Mul(tokens["vUSDT"].Price).Div(EXPSACLE).Sub(gasFee)

			fmt.Printf("calculateSeizedTokenAmount case5, account:%v, path:%v, swap %v%v for %vUSDT, profit:%v\n", account, path2, remain, strings.TrimPrefix(seizedSymbol, "v"), amountsOut[len(amountsOut)-1], profit.Div(EXPSACLE))
			if profit.Cmp(decimal.Zero) == 1 {
				fmt.Printf("case5: profitable liquidation catched:%v, profit:%v\n", liquidation, profit.Div(EXPSACLE))
				s.doLiquidation(big.NewInt(5), flashLoanFrom, path1, path2, addresses, repayAmount.BigInt(), gasPrice.BigInt())
			}
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

func copyTokens(src map[string]*TokenInfo) map[string]*TokenInfo {
	dst := make(map[string]*TokenInfo)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func isStalbeCoin(symbol string) bool {
	return (symbol == "vUSDT" || symbol == "vDAI" || symbol == "vBUSD")
}

func (s *Syncer) buildPaths(srcSymbol, dstSymbol string, tokens map[string]*TokenInfo) []common.Address {
	pancakeFactory := s.pancakeFactory
	pair, err := pancakeFactory.GetPair(nil, tokens[srcSymbol].UnderlyingAddress, tokens[dstSymbol].UnderlyingAddress)
	if err != nil || pair.String() == "0x0000000000000000000000000000000000000000" {
		paths := make([]common.Address, 3)
		paths[0] = tokens[srcSymbol].UnderlyingAddress
		paths[1] = tokens["vBNB"].UnderlyingAddress
		paths[2] = tokens[dstSymbol].UnderlyingAddress
		return paths
	}

	paths := make([]common.Address, 2)
	paths[0] = tokens[srcSymbol].UnderlyingAddress
	paths[1] = tokens[dstSymbol].UnderlyingAddress
	return paths
}

func (s *Syncer) buildVAIPaths(srcSymbol, dstSymbol string, tokens map[string]*TokenInfo) []common.Address {
	pancakeFactory := s.pancakeFactory
	if srcSymbol == "VAI" {
		pair, err := pancakeFactory.GetPair(nil, VAIAddress, tokens[dstSymbol].UnderlyingAddress)
		if err != nil || pair.String() == "0x0000000000000000000000000000000000000000" {
			paths := make([]common.Address, 3)
			paths[0] = VAIAddress
			paths[1] = tokens["vBNB"].UnderlyingAddress
			paths[2] = tokens[dstSymbol].UnderlyingAddress
			return paths
		}

		paths := make([]common.Address, 2)
		paths[0] = VAIAddress
		paths[1] = tokens[dstSymbol].UnderlyingAddress
		return paths
	} else if dstSymbol == "VAI" {
		pair, err := pancakeFactory.GetPair(nil, tokens[srcSymbol].UnderlyingAddress, VAIAddress)
		if err != nil || pair.String() == "0x0000000000000000000000000000000000000000" {
			paths := make([]common.Address, 3)
			paths[0] = tokens[srcSymbol].UnderlyingAddress
			paths[1] = tokens["vBNB"].UnderlyingAddress
			paths[2] = VAIAddress
			return paths
		}

		paths := make([]common.Address, 2)
		paths[0] = tokens[srcSymbol].UnderlyingAddress
		paths[1] = VAIAddress
		return paths
	}
	return nil
}

//situcation： 情况 1-7
//ch： 借钱用的pair地址
//path1： 卖的时候的path, seizedSymbol => repaySymbol的path
//path2:  将seizedSymbol => USDT
//tokens：
// Tokens array
// [0] - _flashLoanVToken 要去借的钱（要还给venus的）
// [1] - _seizedVToken 可以赎回来的钱
// [2] - _seizedTokenUnderlying 赎回来的钱的underlying
// [3] - _flashloanTokenUnderlying 借的钱的underlying
// [4] - target 目标账号
//_flashLoanAmount ： 借多少？ 还多少？

func (s *Syncer) doLiquidation(scenarioNo *big.Int, flashLoanFrom common.Address, path1 []common.Address, path2 []common.Address, tokens []common.Address, flashLoanAmount *big.Int, gasPrice *big.Int) error {

	publicKey := s.PrivateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := s.c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(s.PrivateKey, big.NewInt(56))
	auth.Value = big.NewInt(0) // in wei

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	tx, err := s.liquidator.Qingsuan(auth, scenarioNo, flashLoanFrom, path1, path2, tokens, flashLoanAmount)
	if err != nil {
		return err
	}
	fmt.Printf("tx:%v\n", tx.Hash())
	return nil
}
