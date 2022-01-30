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

type TokenInfo struct {
	Address               common.Address
	CollateralFactorFloat *big.Float
	PriceFloat            *big.Float
}

var (
	ExpScaleFloat, _ = big.NewFloat(0).SetString("1000000000000000000")
	AddressToSymbol  = make(map[common.Address]string)
	TokenDetail      = make(map[string]TokenInfo)
)

type Syncer struct {
	c                *ethclient.Client
	db               *leveldb.DB
	cfg              *config.Config
	wg               sync.WaitGroup
	quitCh           chan struct{}
	forceUpdtePrices chan struct{}
}

func NewSyncer(c *ethclient.Client, db *leveldb.DB, cfg *config.Config) *Syncer {
	return &Syncer{
		c:                c,
		db:               db,
		cfg:              cfg,
		quitCh:           make(chan struct{}),
		forceUpdtePrices: make(chan struct{}),
	}
}

func (s *Syncer) Start() {
	log.Info("syncer start")
	fmt.Println("syncer start")
	err := s.doSyncMarketsAndPrices()
	if err != nil {
		panic(err)
	}

	s.wg.Add(3)
	go s.syncMarketsAndPrices()
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
	for _, detail := range TokenDetail {
		addresses = append(addresses, detail.Address)
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

				//calculate the health factor

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

func (s *Syncer) calculateHealthFactor(account common.Address) (*big.Float, *big.Float, *big.Float, error) {
	comptroller, err := venus.NewComptroller(common.HexToAddress(s.cfg.Comptroller), s.c)
	if err != nil {
		return big.NewFloat(0), big.NewFloat(0), big.NewFloat(0), err
	}

	oracle, err := venus.NewOracle(common.HexToAddress(s.cfg.Oracle), s.c)
	if err != nil {
		return big.NewFloat(0), big.NewFloat(0), big.NewFloat(0), err
	}

	fmt.Printf("account:%v\n", account)
	_, liquidity, shortfall, err := comptroller.GetAccountLiquidity(nil, account)
	if err != nil {
		return big.NewFloat(0), big.NewFloat(0), big.NewFloat(0), err
	}

	assets, err := comptroller.GetAssetsIn(nil, account)
	if err != nil {
		return big.NewFloat(0), big.NewFloat(0), big.NewFloat(0), err
	}

	totalCollateral := big.NewFloat(0)
	totalLoan := big.NewFloat(0)
	mintVAIS, err := comptroller.MintedVAIs(nil, account)

	mintVAISFloatExp := big.NewFloat(0).SetInt(mintVAIS)
	mintVAISFloat := big.NewFloat(0).Quo(mintVAISFloatExp, ExpScaleFloat)

	fmt.Printf("mintVAI:%v\n", mintVAIS)
	for _, asset := range assets {
		//fmt.Printf("asset:%v\n", asset)
		marketInfo, err := comptroller.Markets(nil, asset)
		collateralFactor := marketInfo.CollateralFactorMantissa

		token, err := venus.NewVbep20(asset, s.c)

		_, balance, borrow, exchangeRate, err := token.GetAccountSnapshot(nil, common.HexToAddress(account))

		price, err := oracle.GetUnderlyingPrice(nil, asset)
		if price == big.NewInt(0) {
			continue
		}
		//fmt.Printf("collateralFactor:%v, price:%v, exchangeRate:%v, balance:%v, borrow:%v\n", collateralFactor, price, exchangeRate, balance, borrow)

		exchangeRateFloatExp := big.NewFloat(0).SetInt(exchangeRate)
		exchangeRateFloat := big.NewFloat(0).Quo(exchangeRateFloatExp, ExpScaleFloat)

		collateralFactorFloatExp := big.NewFloat(0).SetInt(collateralFactor)
		collateralFactorFloat := big.NewFloat(0).Quo(collateralFactorFloatExp, ExpScaleFloat)

		priceFloatExp := big.NewFloat(0).SetInt(price)
		priceFloat := big.NewFloat(0).Quo(priceFloatExp, ExpScaleFloat)

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
	fmt.Printf("totalCollateral:%v, totalLoan:%v\n", totalCollateral, totalLoan)
	healthFactor := big.NewFloat(0).Quo(totalCollateral, totalLoan)
	fmt.Printf("healthFactor：%v\n", healthFactor)

	calculatedLiquidity := big.NewFloat(0)
	calculatedShortfall := big.NewFloat(0)
	if totalLoan.Cmp(totalCollateral) == 1 {
		calculatedShortfall = big.NewFloat(0).Sub(totalLoan, totalCollateral)
	} else {
		calculatedLiquidity = big.NewFloat(0).Sub(totalCollateral, totalLoan)
	}

	fmt.Printf("liquidity:%v, calculatedLiquidity:%v\n", liquidity, calculatedLiquidity)
	fmt.Printf("shortfall:%v, calculatedShortfall:%v\n", shortfall, calculatedShortfall)

	return healthFactor, totalCollateral, totalLoan, nil

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
		case <-s.forceUpdtePrices:
			s.doSyncMarketsAndPrices()
		}
	}
}

func (s *Syncer) doSyncMarketsAndPrices() error {
	comptroller, err := venus.NewComptroller(common.HexToAddress(s.cfg.Comptroller), s.c)
	if err != nil {
		return err
	}

	oracle, err := venus.NewOracle(common.HexToAddress(s.cfg.Oracle), s.c)
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
	}
	return nil
}
