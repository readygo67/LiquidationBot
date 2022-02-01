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
}

var (
	ExpScaleFloat, _    = big.NewFloat(0).SetString("1000000000000000000")
	VTokenScaleFloat, _ = big.NewFloat(0).SetString("100000000")
	BigFloat1P0, _      = big.NewFloat(0).SetString("1.0")
	BigFloat1P2, _      = big.NewFloat(0).SetString("1.2")
	BigFloat1P5, _      = big.NewFloat(0).SetString("1.5")
	BigFloat2P0, _      = big.NewFloat(0).SetString("2.0")
	BigFloat3P0, _      = big.NewFloat(0).SetString("3.0")
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
	LiquidationCh       chan *Liquidation
}

func NewSyncer(c *ethclient.Client, db *leveldb.DB, comptrollerAddress string, oracleAddress string) *Syncer {
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

	for _, market := range markets {
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
			panic(err)
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

		tokens[symbol] = token
		symbols[market] = symbol
		vbep20s[market] = vbep20
	}

	return &Syncer{
		c:                   c,
		db:                  db,
		oracle:              oracle,
		comptroller:         comptroller,
		tokens:              tokens,
		symbols:             symbols,
		vbep20s:             vbep20s,
		quitCh:              make(chan struct{}),
		forceUpdatePricesCh: make(chan struct{}, 10),
		feededPricesCh:      make(chan *FeededPrices, 64),
		LiquidationCh:       make(chan *Liquidation, 64),
	}
}

func (s *Syncer) Start() {
	log.Info("syncer start")
	fmt.Println("syncer start")

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
			log.Info("sync markers and prices")
			err := s.doSyncMarketsAndPrices()
			if err != nil {
				t.Reset(time.Millisecond * 20)
			} else {
				t.Reset(time.Second * 3)
			}
		case <-s.forceUpdatePricesCh:
			s.doSyncMarketsAndPrices()
		}
	}
}

func (s *Syncer) doSyncMarketsAndPrices() error {
	comptroller := s.comptroller
	oracle := s.oracle
	c := s.c

	markets, err := comptroller.GetAllMarkets(nil)
	if err != nil {
		return err
	}

	symbols := make(map[common.Address]string)
	tokens := make(map[string]TokenInfo)
	for _, market := range markets {
		vbep20, err := venus.NewVbep20(market, c)
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

		token := TokenInfo{
			Address:               market,
			CollateralFactorFloat: collateralFactorFloat,
			PriceFloat:            priceFloat,
		}

		tokens[symbol] = token
		symbols[market] = symbol
	}

	s.tokens = tokens
	s.symbols = symbols
	return nil
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

	for _, feededPrice := range feededPrices.Prices {
		symbol := s.symbols[feededPrice.Address]
		prefix := append(dbm.MarketPrefix, []byte(symbol)...)

		iter := db.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			account := common.BytesToAddress(iter.Value())
			if exist[account] {
				continue
			}
			exist[account] = true
			accounts = append(accounts, account)
		}
		iter.Release()
	}

	for _, account := range accounts {
		liquidation, _ := s.syncOneAccountWithFeededPrices(account, feededPrices)
		if liquidation != nil {
			s.LiquidationCh <- liquidation
		}
	}
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
				fmt.Printf("%v height:%v, name:%v borrower:%v\n", (i + 1), log.BlockNumber, symbols[log.Address], borrowEvent.Borrower)

				account := borrowEvent.Borrower
				liquidation, err := s.syncOneAccountWithIncreaseAccountNumber(account)
				if err != nil {
					goto EndWithoutUpdateHeight
				}
				if liquidation != nil {
					s.LiquidationCh <- liquidation
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

func (s *Syncer) syncAccounts(accounts []common.Address) {
	for _, account := range accounts {
		liquidation, _ := s.syncOneAccount(account)
		if liquidation != nil {
			s.LiquidationCh <- liquidation
		}
	}
}

func (s *Syncer) syncOneAccount(account common.Address) (*Liquidation, error) {
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
		return nil, err
	}

	mintVAISFloatExp := big.NewFloat(0).SetInt(mintVAIS)
	mintVAISFloat := big.NewFloat(0).Quo(mintVAISFloatExp, ExpScaleFloat)
	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil {
		return nil, err
	}

	for _, market := range markets {
		_, balance, borrow, exchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil {
			return nil, err
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
	fmt.Printf("totalCollateral:%v, totalLoan:%v\n", totalCollateral, totalLoan)
	healthFactor := big.NewFloat(0).Quo(totalCollateral, totalLoan)
	fmt.Printf("healthFactor：%v\n", healthFactor)

	//update market table and account table
	info := AccountInfo{
		HealthFactor: healthFactor,
		Assets:       assets,
	}
	s.updateDB(account, info)

	if healthFactor.Cmp(big.NewFloat(1)) == 1 {
		return nil, nil
	} else {
		blockNumber, _ := s.c.BlockNumber(ctx)
		liquidation := Liquidation{
			Address:      account,
			HealthFactor: healthFactor,
			BlockNumber:  blockNumber,
		}
		return &liquidation, nil
	}
}

func (s *Syncer) syncOneAccountWithFeededPrices(account common.Address, feedPrices *FeededPrices) (*Liquidation, error) {
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
		return nil, err
	}

	mintVAISFloatExp := big.NewFloat(0).SetInt(mintVAIS)
	mintVAISFloat := big.NewFloat(0).Quo(mintVAISFloatExp, ExpScaleFloat)
	markets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil {
		return nil, err
	}

	for _, market := range markets {
		_, balance, borrow, exchangeRate, err := vbep20s[market].GetAccountSnapshot(nil, account)
		if err != nil {
			return nil, err
		}

		exchangeRateFloatExp := big.NewFloat(0).SetInt(exchangeRate)
		exchangeRateFloat := big.NewFloat(0).Quo(exchangeRateFloatExp, ExpScaleFloat)

		symbol := symbols[market]
		collateralFactorFloat := tokens[symbol].CollateralFactorFloat
		priceFloat := tokens[symbol].PriceFloat
		for _, feededPrice := range feedPrices.Prices {
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

	if healthFactor.Cmp(big.NewFloat(1)) == 1 {
		return nil, nil
	} else {
		blockNumber, _ := s.c.BlockNumber(ctx)
		liquidation := Liquidation{
			Address:      account,
			HealthFactor: healthFactor,
			BlockNumber:  blockNumber,
		}
		return &liquidation, nil
	}
}

func (s *Syncer) syncOneAccountWithIncreaseAccountNumber(account common.Address) (*Liquidation, error) {
	ctx := context.Background()
	c := s.c
	db := s.db

	accountBytes := account.Bytes()
	exist, err := db.Has(dbm.BorrowersStoreKey(accountBytes), nil)
	if err != nil {
		return nil, err
	}

	byteCode, err := c.CodeAt(ctx, account, nil)
	if len(byteCode) > 0 {
		//ignore smart contract
		return nil, nil
	}

	liquidation, err := s.syncOneAccount(account)
	if err != nil {
		return nil, err
	}

	if !exist {
		//if account not exist in borrowers table, record it into borrowers table and increase borrowers number
		err = db.Put(dbm.BorrowersStoreKey(accountBytes), accountBytes, nil)
		if err != nil {
			return nil, err
		}

		bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
		num := big.NewInt(0).SetBytes(bz).Int64()
		if err != nil {
			return nil, err
		}

		num += 1
		err = db.Put(dbm.BorrowerNumberKey(), big.NewInt(num).Bytes(), nil)
		if err != nil {
			return nil, err
		}
	}
	return liquidation, nil
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
