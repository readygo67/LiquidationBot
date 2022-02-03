package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"github.com/readygo67/LiquidationBot/venus"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"
)

var syncer *Syncer

//func TestMain(m *testing.M){
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	if err != nil{
//		panic(err)
//	}
//
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle)
//
//}

func TestGetvAAVEUnderlyingPrice(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	oracle, err := venus.NewOracle(common.HexToAddress(cfg.Oracle), c)
	if err != nil {
		panic(err)
	}
	//fail to get vAAVE prices @ 20220201
	_, err = oracle.GetUnderlyingPrice(nil, common.HexToAddress("0x26DA28954763B92139ED49283625ceCAf52C6f94"))
	require.Equal(t, "execution reverted: REF_DATA_NOT_AVAILABLE", err.Error())
}

func TestNewSyncer(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
	verifyTokens(t, sync)

	bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
	require.NoError(t, err)

	num := big.NewInt(0).SetBytes(bz)
	require.Equal(t, int64(0), num.Int64())
}

func TestDoSyncMarketsAndPrices(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
	t.Logf("begin do sync markets and prices\n")

	sync.doSyncMarketsAndPrices()
	verifyTokens(t, sync)
}

func TestSyncMarketsAndPrices(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
	t.Logf("begin sync markets and prices\n")
	sync.wg.Add(1)
	go sync.syncMarketsAndPrices()

	time.Sleep(time.Second * 60)
	close(sync.quitCh)
	sync.wg.Wait()
	verifyTokens(t, sync)
}

func TestFilterAllCotractsBorrowEvent(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, nil, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)

	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")
	var addresses []common.Address
	name := make(map[string]string)
	for _, token := range sync.tokens {
		addresses = append(addresses, token.Address)
	}

	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	require.NoError(t, err)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(13000000),
		ToBlock:   big.NewInt(13002000),
		Addresses: addresses,
		Topics:    [][]common.Hash{{topicBorrow}},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	require.NoError(t, err)
	fmt.Printf("start Time:%v\n", time.Now())
	for i, log := range logs {
		var borrowEvent venus.Vbep20Borrow
		err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", log.Data)
		fmt.Printf("%v height:%v, name:%v borrower:%v\n", (i + 1), log.BlockNumber, name[strings.ToLower(log.Address.String())], borrowEvent.Borrower)
	}
	fmt.Printf("end Time:%v\n", time.Now())
}

func TestCalculateHealthFactorInFloat(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
	comptroller := sync.comptroller
	oracle := sync.oracle

	accounts := []string{
		"0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B",
		//"0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3",
		//"0xc528045078Ff53eA289fA42aF3e12D8eF901cABD",
		//"0xF2ddE5689B0e13344231D3459B03432E97a48E0c",
	}

	for _, account := range accounts {
		fmt.Printf("address:%v\n", account)
		_, liquidity, shortfall, err := comptroller.GetAccountLiquidity(nil, common.HexToAddress(account))
		require.NoError(t, err)

		assets, err := comptroller.GetAssetsIn(nil, common.HexToAddress(account))
		fmt.Printf("assets:%v\n", assets)
		require.NoError(t, err)

		totalCollateral := big.NewFloat(0)
		totalLoan := big.NewFloat(0)
		mintVAIS, err := comptroller.MintedVAIs(nil, common.HexToAddress(account))

		mintVAISFloatExp := big.NewFloat(0).SetInt(mintVAIS)
		mintVAISFloat := big.NewFloat(0).Quo(mintVAISFloatExp, ExpScaleFloat)

		for _, asset := range assets {
			//fmt.Printf("asset:%v\n", asset)
			marketInfo, err := comptroller.Markets(nil, asset)
			collateralFactor := marketInfo.CollateralFactorMantissa
			require.NoError(t, err)

			token, err := venus.NewVbep20(asset, c)
			require.NoError(t, err)

			_, balance, borrow, exchangeRate, err := token.GetAccountSnapshot(nil, common.HexToAddress(account))

			price, err := oracle.GetUnderlyingPrice(nil, asset)
			if price == big.NewInt(0) {
				continue
			}
			fmt.Printf("collateralFactor:%v, price:%v, exchangeRate:%v, balance:%v, borrow:%v\n", collateralFactor, price, exchangeRate, balance, borrow)

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

			borrowFloatExp := big.NewFloat(0).SetInt(borrow)
			borrowFloat := big.NewFloat(0).Quo(borrowFloatExp, ExpScaleFloat)

			fmt.Printf("collateralFactor:%v, price:%v, exchangeRate:%v, balance:%v, borrow:%v\n", collateralFactorFloat, priceFloat, exchangeRateFloat, balanceFloat, borrowFloat)

			collateral := big.NewFloat(0).Mul(balanceFloat, multiplier)
			totalCollateral = big.NewFloat(0).Add(totalCollateral, collateral)

			loan := big.NewFloat(0).Mul(borrowFloat, priceFloat)
			totalLoan = big.NewFloat(0).Add(totalLoan, loan)
		}

		totalLoan = big.NewFloat(0).Add(totalLoan, mintVAISFloat)
		fmt.Printf("totalCollateral:%v, totalLoan:%v\n", totalCollateral.String(), totalLoan)
		healthFactor := big.NewFloat(100)
		if totalLoan.Cmp(BigFloatZero) == 1 {
			healthFactor = big.NewFloat(0).Quo(totalCollateral, totalLoan)
		}

		fmt.Printf("healthFactor：%v\n", healthFactor)
		calculatedLiquidity := big.NewFloat(0)
		calculatedShortfall := big.NewFloat(0)
		if totalLoan.Cmp(totalCollateral) == 1 {
			calculatedShortfall = big.NewFloat(0).Sub(totalLoan, totalCollateral)
		} else {
			calculatedLiquidity = big.NewFloat(0).Sub(totalCollateral, totalLoan)
		}

		fmt.Printf("liquidity:%v, calculatedLiquidity:%v\n", liquidity.String(), calculatedLiquidity.String())
		fmt.Printf("shortfall:%v, calculatedShortfall:%v\n", shortfall, calculatedShortfall)
	}
}

func TestStoreAndDeleteAccount(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)

	healthFactor, _ := big.NewFloat(0).SetString("0.9")
	vusdtBalance, _ := big.NewFloat(0).SetString("1000000000.0")
	vusdtLoan, _ := big.NewFloat(0).SetString("0")

	vbtcBalance, _ := big.NewFloat(0).SetString("2.5")
	vbtctLoan, _ := big.NewFloat(0).SetString("0.2")

	vbusdBalance, _ := big.NewFloat(0).SetString("0")
	vbusdtLoan, _ := big.NewFloat(0).SetString("500.23")

	assets := []Asset{
		{
			Symbol:  "vUSDT",
			Balance: vusdtBalance,
			Loan:    vusdtLoan,
		},
		{
			Symbol:  "vBTC",
			Balance: vbtcBalance,
			Loan:    vbtctLoan,
		},
		{
			Symbol:  "vBUSD",
			Balance: vbusdBalance,
			Loan:    vbusdtLoan,
		},
	}
	info := AccountInfo{
		HealthFactor: healthFactor,
		Assets:       assets,
	}

	account := common.HexToAddress("0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3")
	sync.storeAccount(account, info)

	bz, err := db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
	//t.Logf("bz:%v\n", string(bz))
	require.NoError(t, err)

	var got AccountInfo
	err = json.Unmarshal(bz, &got)
	require.NoError(t, err)

	has, err := db.Has(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.True(t, has)

	bz, err = db.Get(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.Equal(t, bz, account.Bytes())

	for _, asset := range assets {
		has, err = db.Has(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.True(t, has)

		bz, err = db.Get(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.Equal(t, bz, account.Bytes())

		prefix := append(dbm.MarketPrefix, []byte(asset.Symbol)...)
		var accounts []common.Address
		iter := db.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			accounts = append(accounts, common.BytesToAddress(iter.Value()))
		}

		require.Equal(t, 1, len(accounts))
	}

	had, err := db.Has(dbm.MarketStoreKey([]byte("vETH"), account.Bytes()), nil)
	require.NoError(t, err)
	require.False(t, had)

	sync.deleteAccount(account)
	has, err = db.Has(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.False(t, has)

	bz, err = db.Get(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	require.Equal(t, leveldb.ErrNotFound, err)
}

func TestStoreAndDeleteAccount1(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)

	healthFactor, _ := big.NewFloat(0).SetString("1.1")
	vusdtBalance, _ := big.NewFloat(0).SetString("1000000000.0")
	vusdtLoan, _ := big.NewFloat(0).SetString("0")

	vbtcBalance, _ := big.NewFloat(0).SetString("2.5")
	vbtctLoan, _ := big.NewFloat(0).SetString("0.2")

	vbusdBalance, _ := big.NewFloat(0).SetString("0")
	vbusdtLoan, _ := big.NewFloat(0).SetString("500.23")

	assets := []Asset{
		{
			Symbol:  "vUSDT",
			Balance: vusdtBalance,
			Loan:    vusdtLoan,
		},
		{
			Symbol:  "vBTC",
			Balance: vbtcBalance,
			Loan:    vbtctLoan,
		},
		{
			Symbol:  "vBUSD",
			Balance: vbusdBalance,
			Loan:    vbusdtLoan,
		},
	}
	info := AccountInfo{
		HealthFactor: healthFactor,
		Assets:       assets,
	}

	account := common.HexToAddress("0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3")
	sync.storeAccount(account, info)

	bz, err := db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
	//t.Logf("bz:%v\n", string(bz))
	require.NoError(t, err)

	var got AccountInfo
	err = json.Unmarshal(bz, &got)
	require.NoError(t, err)

	has, err := db.Has(dbm.LiquidationBelow1P1StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.True(t, has)

	bz, err = db.Get(dbm.LiquidationBelow1P1StoreKey(account.Bytes()), nil)
	require.NoError(t, err)
	require.Equal(t, bz, account.Bytes())

	for _, asset := range assets {
		has, err = db.Has(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.True(t, has)

		bz, err = db.Get(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.Equal(t, bz, account.Bytes())
	}

	had, err := db.Has(dbm.MarketStoreKey([]byte("vETH"), account.Bytes()), nil)
	require.NoError(t, err)
	require.False(t, had)

	sync.deleteAccount(account)

	vsxpBalance, _ := big.NewFloat(0).SetString("236.5")
	vsxpLoan, _ := big.NewFloat(0).SetString("800.23")

	assets = append(assets, Asset{
		Symbol:  "vSXP",
		Balance: vsxpBalance,
		Loan:    vsxpLoan,
	})

	info = AccountInfo{
		HealthFactor: healthFactor,
		Assets:       assets,
	}

	sync.storeAccount(account, info)
	bz, err = db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
	//t.Logf("bz:%v\n", string(bz))
	require.NoError(t, err)

	err = json.Unmarshal(bz, &got)
	require.NoError(t, err)

	for _, asset := range assets {
		//fmt.Printf("symbol:%v\n", asset.Symbol)
		has, err = db.Has(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.True(t, has)

		bz, err = db.Get(dbm.MarketStoreKey([]byte(asset.Symbol), account.Bytes()), nil)
		require.NoError(t, err)
		require.Equal(t, bz, account.Bytes())
	}
}

// 从compound通过getExchangeRateStored方法获得的exchangeRat是乘了10^18的结果，实际使用时需要除10^18,
func TestCalculateExchangeRate(t *testing.T) {
	//exchangeRateStored: 202001285536565656590891932
	//totalSupply: 76384766592957
	//totalBorrow: 2298168762317337651162
	//totalReserver:  4713643651873292071
	//cash: 13136365928522364031146
	borrow, _ := big.NewInt(0).SetString("2298168762317337651162", 10)
	supply, _ := big.NewInt(0).SetString("76384766592957", 10)
	reserve, _ := big.NewInt(0).SetString("4713643651873292071", 10)
	cash, _ := big.NewInt(0).SetString("13136365928522364031146", 10)
	sum := big.NewInt(0).Add(cash, borrow)
	sum = big.NewInt(0).Sub(sum, reserve)
	//rate := big.NewInt(0).Div(sum, supply)
	//fmt.Printf("rate:%v\n", rate)

	ExpScale, _ := big.NewInt(0).SetString("1000000000000000000", 10)
	sumExp := big.NewInt(0).Mul(sum, ExpScale)
	rateExp := big.NewInt(0).Div(sumExp, supply)
	//fmt.Printf("rateExp:%v\n", rateExp)
	require.Equal(t, "202001285536565656590891932", rateExp.String())
}

func TestSyncOneAccount(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
	account := common.HexToAddress("0xF5A008a26c8C06F0e778ac07A0db9a2f42423c84") //0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B
	accountBytes := account.Bytes()
	err = sync.syncOneAccount(account)
	require.NoError(t, err)

	exist, err := db.Has(dbm.AccountStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.True(t, exist)

	bz, err := db.Get(dbm.BorrowersStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	accounts := []common.Address{}
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	bz, err = db.Get(dbm.MarketStoreKey([]byte("vDOGE"), accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	bz, err = db.Get(dbm.MarketStoreKey([]byte("vUSDT"), accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	prefix := append(dbm.MarketPrefix, []byte("vDOGE")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	prefix = append(dbm.MarketPrefix, []byte("vUSDT")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	var info AccountInfo
	err = json.Unmarshal(bz, &info)
	fmt.Printf("info:%v\n", info)
	require.True(t, big.NewFloat(100).Cmp(info.HealthFactor) == 0)

	bz, err = db.Get(dbm.LiquidationAbove2P0StoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))
}

//
//func TestSyncAccounts(t *testing.T) {
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	liquidationCh := make(chan *Liquidation, 64)
//	priorityliquidationCh := make(chan *Liquidation, 64)
//	feededPricesCh := make(chan *FeededPrices, 64)
//
//	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
//	accountStrs := []string{
//		"0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B",
//		"0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3",
//		"0xc528045078Ff53eA289fA42aF3e12D8eF901cABD",
//		"0xF2ddE5689B0e13344231D3459B03432E97a48E0c",
//	}
//	var accounts []common.Address
//	for _, accountStr := range accountStrs {
//		accounts = append(accounts, common.HexToAddress(accountStr))
//	}
//
//	sync.syncAccounts(accounts)
//	require.NoError(t, err)
//
//	count := make(map[string]int)
//
//	for _, account := range accounts {
//		accountBytes := account.Bytes()
//		bz, err := db.Get(dbm.BorrowersStoreKey(accountBytes), nil)
//		require.NoError(t, err)
//		require.Equal(t, account, common.BytesToAddress(bz))
//
//		bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
//		require.NoError(t, err)
//
//		var got AccountInfo
//		err = json.Unmarshal(bz, &got)
//		require.NoError(t, err)
//
//		for _, token := range got.Assets {
//			count[token.Symbol] += 1
//			bz, err = db.Get(dbm.MarketStoreKey([]byte(token.Symbol), accountBytes), nil)
//			require.NoError(t, err)
//			require.Equal(t, account, common.BytesToAddress(bz))
//		}
//	}
//
//	//total account number = 4
//	gotAccounts := []common.Address{}
//	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
//	for iter.Next() {
//		gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
//	}
//	require.Equal(t, 4, len(gotAccounts))
//
//	for symbol, cnt := range count {
//		prefix := append(dbm.MarketPrefix, []byte(symbol)...)
//		gotAccounts = []common.Address{}
//		iter = db.NewIterator(util.BytesPrefix(prefix), nil)
//		for iter.Next() {
//			gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
//		}
//		require.Equal(t, cnt, len(gotAccounts))
//	}
//
//	gotAccounts = []common.Address{}
//	iter = db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P5Prefix), nil)
//	for iter.Next() {
//		gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
//	}
//	require.Equal(t, 2, len(gotAccounts))
//
//	gotAccounts = []common.Address{}
//	iter = db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow3P0Prefix), nil)
//	for iter.Next() {
//		gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
//	}
//	require.Equal(t, 1, len(gotAccounts))
//
//	gotAccounts = []common.Address{}
//	iter = db.NewIterator(util.BytesPrefix(dbm.LiquidationAbove3P0Prefix), nil)
//	for iter.Next() {
//		gotAccounts = append(gotAccounts, common.BytesToAddress(iter.Value()))
//	}
//	require.Equal(t, 1, len(gotAccounts))
//}

func TestSyncOneAccountWithIncreaseAccountNumer(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
	account := common.HexToAddress("0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B")
	accountBytes := account.Bytes()
	err = sync.syncOneAccountWithIncreaseAccountNumber(account)
	require.NoError(t, err)

	bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
	num := big.NewInt(0).SetBytes(bz)
	require.Equal(t, int64(1), num.Int64())

	exist, err := db.Has(dbm.AccountStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.True(t, exist)

	bz, err = db.Get(dbm.BorrowersStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	accounts := []common.Address{}
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	bz, err = db.Get(dbm.MarketStoreKey([]byte("vDOGE"), accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	bz, err = db.Get(dbm.MarketStoreKey([]byte("vUSDT"), accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))

	prefix := append(dbm.MarketPrefix, []byte("vDOGE")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	prefix = append(dbm.MarketPrefix, []byte("vUSDT")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 1, len(accounts))

	bz, err = db.Get(dbm.AccountStoreKey(accountBytes), nil)
	var info AccountInfo
	err = json.Unmarshal(bz, &info)
	fmt.Printf("info:%v\n", info)
	require.True(t, big.NewFloat(100).Cmp(info.HealthFactor) == 0)

	bz, err = db.Get(dbm.LiquidationAbove2P0StoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.Equal(t, account, common.BytesToAddress(bz))
}

func TestSyncOneAccountWithFeededPrices(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
	account := common.HexToAddress("0x03CB27196B92B3b6B8681dC00C30946E0DB0EA7B")
	accountBytes := account.Bytes()

	feededUsdtPrice, _ := big.NewFloat(0).SetString("1.02")
	feededDogePrice, _ := big.NewFloat(0).SetString("0.04")
	feededPrices := &FeededPrices{
		Prices: []FeededPrice{
			{
				Address:    common.HexToAddress("0xec3422Ef92B2fb59e84c8B02Ba73F1fE84Ed8D71"),
				PriceFloat: feededUsdtPrice,
			},
			{
				Address:    common.HexToAddress("0xfD5840Cd36d94D7229439859C0112a4185BC0255"),
				PriceFloat: feededDogePrice,
			},
		},
		Height: 100000,
	}
	err = sync.syncOneAccountWithFeededPrices(account, feededPrices)
	require.NoError(t, err)

	bz, err := db.Get(dbm.BorrowerNumberKey(), nil)
	num := big.NewInt(0).SetBytes(bz)
	require.Equal(t, int64(0), num.Int64())

	exist, err := db.Has(dbm.AccountStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.False(t, exist)

	accounts := []common.Address{}
	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 0, len(accounts))

	exist, err = db.Has(dbm.MarketStoreKey([]byte("vDOGE"), accountBytes), nil)
	require.NoError(t, err)
	require.False(t, exist)

	exist, err = db.Has(dbm.MarketStoreKey([]byte("vUSDT"), accountBytes), nil)
	require.NoError(t, err)
	require.False(t, exist)

	prefix := append(dbm.MarketPrefix, []byte("vDOGE")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 0, len(accounts))

	prefix = append(dbm.MarketPrefix, []byte("vUSDT")...)
	accounts = []common.Address{}
	iter = db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		accounts = append(accounts, common.BytesToAddress(iter.Value()))
	}
	require.Equal(t, 0, len(accounts))

	exist, err = db.Has(dbm.AccountStoreKey(accountBytes), nil)
	require.NoError(t, err)
	require.False(t, exist)
}

//func TestScanAllBorrowers(t *testing.T) {
//	ctx := context.Background()
//
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	db, err := dbm.NewDB("testdb1")
//	require.NoError(t, err)
//	defer db.Close()
//	defer os.RemoveAll("testdb1")
//
//	_, err = c.BlockNumber(ctx)
//	require.NoError(t, err)
//
//	sync := NewSyncer(c, db, cfg)
//	start := big.NewInt(13000000) //big.NewInt(14747565)
//	db.Put(dbm.KeyLastHandledHeight, start.Bytes(), nil)
//	db.Put(dbm.KeyBorrowerNumber, big.NewInt(0).Bytes(), nil)
//
//	sync.Start()
//	time.Sleep(time.Second * 120)
//	sync.Stop()
//
//	bz, err := db.Get(dbm.KeyLastHandledHeight, nil)
//	end := big.NewInt(0).SetBytes(bz)
//	t.Logf("end height:%v\n", end.Int64())
//
//	bz, err = db.Get(dbm.KeyBorrowerNumber, nil)
//	num := big.NewInt(0).SetBytes(bz).Int64()
//	t.Logf("num:%v\n", num)
//
//	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
//	defer iter.Release()
//	t.Logf("borrows address")
//	for iter.Next() {
//		addr := common.BytesToAddress(iter.Value())
//		t.Logf("%v\n", addr.String())
//	}
//}
//
func TestScanAllBorrowers1(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	height, err := c.BlockNumber(ctx)
	require.NoError(t, err)

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, feededPricesCh, liquidationCh, priorityliquidationCh)
	startHeight := big.NewInt(int64(height - 5000))
	db.Put(dbm.KeyLastHandledHeight, startHeight.Bytes(), nil)
	db.Put(dbm.KeyBorrowerNumber, big.NewInt(0).Bytes(), nil)

	sync.Start()
	time.Sleep(time.Second * 60)
	sync.Stop()

	bz, err := db.Get(dbm.KeyLastHandledHeight, nil)
	end := big.NewInt(0).SetBytes(bz)
	t.Logf("end height:%v\n", end.Int64())

	bz, err = db.Get(dbm.KeyBorrowerNumber, nil)
	num := big.NewInt(0).SetBytes(bz).Int64()
	t.Logf("num:%v\n", num)

	iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
	defer iter.Release()
	t.Logf("borrows address")
	for iter.Next() {
		addr := common.BytesToAddress(iter.Value())
		t.Logf("%v\n", addr.String())
	}
}

func verifyTokens(t *testing.T, sync *Syncer) {
	require.Equal(t, common.HexToAddress("0xf508fCD89b8bd15579dc79A6827cB4686A3592c8"), sync.tokens["vETH"].Address)
	require.Equal(t, common.HexToAddress("0xfD5840Cd36d94D7229439859C0112a4185BC0255"), sync.tokens["vUSDT"].Address)
	require.Equal(t, common.HexToAddress("0x61eDcFe8Dd6bA3c891CB9bEc2dc7657B3B422E93"), sync.tokens["vTRX"].Address)
	require.Equal(t, common.HexToAddress("0x08CEB3F4a7ed3500cA0982bcd0FC7816688084c3"), sync.tokens["vTUSD"].Address)
	require.Equal(t, common.HexToAddress("0x26DA28954763B92139ED49283625ceCAf52C6f94"), sync.tokens["vAAVE"].Address)
	require.Equal(t, common.HexToAddress("0x86aC3974e2BD0d60825230fa6F355fF11409df5c"), sync.tokens["vCAKE"].Address)
	require.Equal(t, common.HexToAddress("0x5c9476FcD6a4F9a3654139721c949c2233bBbBc8"), sync.tokens["vMATIC"].Address)
	require.Equal(t, common.HexToAddress("0xec3422Ef92B2fb59e84c8B02Ba73F1fE84Ed8D71"), sync.tokens["vDOGE"].Address)
	require.Equal(t, common.HexToAddress("0x9A0AF7FDb2065Ce470D72664DE73cAE409dA28Ec"), sync.tokens["vADA"].Address)
	require.Equal(t, common.HexToAddress("0xeBD0070237a0713E8D94fEf1B728d3d993d290ef"), sync.tokens["vCAN"].Address)
	require.Equal(t, common.HexToAddress("0x972207A639CC1B374B893cc33Fa251b55CEB7c07"), sync.tokens["vBETH"].Address)
	require.Equal(t, common.HexToAddress("0x334b3eCB4DCa3593BCCC3c7EBD1A1C1d1780FBF1"), sync.tokens["vDAI"].Address)
	require.Equal(t, common.HexToAddress("0x650b940a1033B8A1b1873f78730FcFC73ec11f1f"), sync.tokens["vLINK"].Address)
	require.Equal(t, common.HexToAddress("0x1610bc33319e9398de5f57B33a5b184c806aD217"), sync.tokens["vDOT"].Address)
	require.Equal(t, common.HexToAddress("0x5F0388EBc2B94FA8E123F404b79cCF5f40b29176"), sync.tokens["vBCH"].Address)
	require.Equal(t, common.HexToAddress("0xB248a295732e0225acd3337607cc01068e3b9c10"), sync.tokens["vXRP"].Address)
	require.Equal(t, common.HexToAddress("0x57A5297F2cB2c0AaC9D554660acd6D385Ab50c6B"), sync.tokens["vLTC"].Address)
	require.Equal(t, common.HexToAddress("0x882C173bC7Ff3b7786CA16dfeD3DFFfb9Ee7847B"), sync.tokens["vBTC"].Address)
	require.Equal(t, common.HexToAddress("0xA07c5b74C9B40447a954e1466938b865b6BBea36"), sync.tokens["vBNB"].Address)
	require.Equal(t, common.HexToAddress("0x151B1e2635A717bcDc836ECd6FbB62B674FE3E1D"), sync.tokens["vXVS"].Address)
	require.Equal(t, common.HexToAddress("0x2fF3d0F6990a40261c66E1ff2017aCBc282EB6d0"), sync.tokens["vSXP"].Address)
	require.Equal(t, common.HexToAddress("0x95c78222B3D6e262426483D42CfA53685A67Ab9D"), sync.tokens["vBUSD"].Address)
	require.Equal(t, common.HexToAddress("0xf508fCD89b8bd15579dc79A6827cB4686A3592c8"), sync.tokens["vETH"].Address)
	require.Equal(t, common.HexToAddress("0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8"), sync.tokens["vUSDC"].Address)
}
