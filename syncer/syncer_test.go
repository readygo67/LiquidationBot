package syncer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"github.com/readygo67/LiquidationBot/venus"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"
)

//Filter with ethereum.FilterQuery
func TestFilterBorrowEvent(t *testing.T) {
	ctx := context.Background()
	rpcURL := "http://42.3.146.198:21993" //"https://bsc-dataseed1.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	vUSDCAddress := common.HexToAddress("0xeca88125a5adbe82614ffc12d0db554e2e2867c8")
	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")

	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	require.NoError(t, err)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(14747566),
		ToBlock:   big.NewInt(14747568),
		Addresses: []common.Address{vUSDCAddress},
		Topics:    [][]common.Hash{{topicBorrow}},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	require.NoError(t, err)
	for _, log := range logs {
		var borrowEvent venus.Vbep20Borrow
		err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", log.Data)
		fmt.Printf("BorrowEvent:%v", borrowEvent)
	}
}

//
//func TestFilterAllCotractsBorrowEvent(t *testing.T) {
//	ctx := context.Background()
//	cfg, err := config.New("../config.yml")
//	rpcURL := "http://42.3.146.198:21993"
//	c, err := ethclient.Dial(rpcURL)
//
//	_, err = c.BlockNumber(ctx)
//	require.NoError(t, err)
//
//	sync := NewSyncer(c, nil, cfg)
//
//	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")
//	var addresses []common.Address
//	name := make(map[string]string)
//	for _, token := range sync.cfg.Tokens {
//		addresses = append(addresses, common.HexToAddress(token.Address))
//		name[strings.ToLower(token.Address)] = token.Name
//		//fmt.Printf("name:%v, address:%v\n", name[token.Address], token.Address)
//	}
//
//	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
//	require.NoError(t, err)
//
//	query := ethereum.FilterQuery{
//		FromBlock: big.NewInt(13019766),
//		ToBlock:   big.NewInt(13019768),
//		Addresses: addresses,
//		Topics:    [][]common.Hash{{topicBorrow}},
//	}
//
//	logs, err := c.FilterLogs(context.Background(), query)
//	require.NoError(t, err)
//	fmt.Printf("start Time:%v\n", time.Now())
//	for i, log := range logs {
//		var borrowEvent venus.Vbep20Borrow
//		err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", log.Data)
//		//fmt.Printf("logAddress:%v\n", log.Address.String())
//		fmt.Printf("%v height:%v, name:%v borrower:%v\n", (i + 1), log.BlockNumber, name[strings.ToLower(log.Address.String())], borrowEvent.Borrower)
//		//jsonLog, err := json.Marshal(log)
//		//require.NoError(t, err)
//		//fmt.Printf("%v log: %s\n", i, jsonLog)
//	}
//	fmt.Printf("end Time:%v\n", time.Now())
//}

/*
43 height:14782956, name:vBUSD borrower:0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3
44 height:14783053, name:vBUSD borrower:0xc528045078Ff53eA289fA42aF3e12D8eF901cABD
45 height:14783060, name:vBUSD borrower:0xF2ddE5689B0e13344231D3459B03432E97a48E0c
46 height:14783144, name:vBUSD borrower:0xe572DC5871d62D180519057f2bE27AdFA31c86b6
47 height:14783238, name:vBUSD borrower:0x8ac41aED58cCc43308fD68662aDF3c2B02A8FaCd
48 height:14783287, name:vBUSD borrower:0x6666F2Be71e21444991BF8949B62255e68486666
49 height:14783313, name:vUSDT borrower:0xc086342EbD2a7130966CdF8cD949c15cCcbA8616
50 height:14783466, name:vBUSD borrower:0x390F523604b11e6E71F23B9F8A2055d0316F48CD
51 height:14783633, name:vBUSD borrower:0x1E60b4954e6BE92D9560AfFb4788200d586B9A49
52 height:14783753, name:vUSDC borrower:0xC9f678D97CfE754D5CC16261727442A84878791F
53 height:14783771, name:vBUSD borrower:0x6666F2Be71e21444991BF8949B62255e68486666
54 height:14783857, name:vCAKE borrower:0x584f17edC1139bAD77291221d34b949DAa418554
55 height:14783959, name:vBUSD borrower:0xc528045078Ff53eA289fA42aF3e12D8eF901cABD
56 height:14784062, name:vBUSD borrower:0x6e8494A879BF37fB3Aa0B1a775A48e475A488207
57 height:14784101, name:vUSDT borrower:0x05Fded65c4eEF241294E87380CA36a933b38C7Ce
58 height:14784104, name:vFIL borrower:0x6FD2234B199A3B8D7b6e3F502dd8Ea43A6699A97
59 height:14784263, name:vBNB borrower:0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3
60 height:14784272, name:vUSDT borrower:0xB0cC4E7D7d03c468F98Ccb8513Edcb01FdD8888A
61 height:14784286, name:vBUSD borrower:0xEb9b830138326bdD351724Ec56132a1AdEe7E3C8
62 height:14784308, name:vUSDC borrower:0xC06A1b2B350c564ea25bC00c929cC5D13C03ea2f
63 height:14784313, name:vBUSD borrower:0x67d43DF7D9DaF4f4e1CC3005D9D15F7026Cd0A79
64 height:14784317, name:vCAKE borrower:0x45cf1bDFBE5332322a10659e2cAa1AB998F71588
*/

func TestCalculateHealthFactor(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	comptroller, err := venus.NewComptroller(common.HexToAddress(cfg.Comptroller), c)
	require.NoError(t, err)
	oracle, err := venus.NewOracle(common.HexToAddress(cfg.Oracle), c)
	require.NoError(t, err)

	markets, err := comptroller.GetAllMarkets(nil)
	require.NoError(t, err)

	//closeFactor, err := comptroller.CloseFactorMantissa(nil)
	//require.NoError(t, err)

	var addresses []common.Address
	name := make(map[string]string)

	ExpScale, _ := big.NewInt(0).SetString("1000000000000000000", 10)

	for _, market := range markets {
		addresses = append(addresses, market)
		token, err := venus.NewVbep20(market, c)
		require.NoError(t, err)

		symbol, err := token.Symbol(nil)
		name[strings.ToLower(market.String())] = symbol
		//fmt.Printf("name:%v, address:%v\n", name[strings.ToLower(market.String())], market.String())
	}

	accounts := []string{
		"0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3",
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

		totalCollateral := big.NewInt(0)
		totalLoan := big.NewInt(0)
		mintVAIS, err := comptroller.MintedVAIs(nil, common.HexToAddress(account))
		fmt.Printf("mintVAI:%v\n", mintVAIS)
		require.NoError(t, err)

		for _, asset := range assets {
			fmt.Printf("asset:%v\n", asset)
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

			multiplier := big.NewInt(0).Mul(collateralFactor, exchangeRate)
			multiplier = big.NewInt(0).Div(multiplier, ExpScale)
			multiplier = big.NewInt(0).Mul(multiplier, price)
			multiplier = big.NewInt(0).Div(multiplier, ExpScale)

			collateral := big.NewInt(0).Mul(balance, multiplier)
			collateral = big.NewInt(0).Div(collateral, ExpScale)
			totalCollateral = big.NewInt(0).Add(totalCollateral, collateral)

			loan := big.NewInt(0).Mul(borrow, price)
			loan = big.NewInt(0).Div(loan, ExpScale)
			totalLoan = big.NewInt(0).Add(totalLoan, loan)
		}

		totalLoan = big.NewInt(0).Add(totalLoan, mintVAIS)
		fmt.Printf("totalCollateral:%v, totalLoan:%v\n", totalCollateral, totalLoan)
		healthFactor := big.NewInt(0).Div(totalLoan, totalCollateral)
		fmt.Printf("healthFactor：%v\n", healthFactor)

		calculatedLiquidity := big.NewInt(0)
		calculatedShortfall := big.NewInt(0)
		if totalLoan.Cmp(totalCollateral) == 1 {
			calculatedShortfall = big.NewInt(0).Sub(totalLoan, totalCollateral)
		} else {
			calculatedLiquidity = big.NewInt(0).Sub(totalCollateral, totalLoan)
		}

		fmt.Printf("liquidity:%v, calculatedLiquidity:%v\n", liquidity, calculatedLiquidity)
		fmt.Printf("shortfall:%v, calculatedShortfall:%v\n", shortfall, calculatedShortfall)
	}
}

func TestCalculateHealthFactor1(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	comptroller, err := venus.NewComptroller(common.HexToAddress(cfg.Comptroller), c)
	require.NoError(t, err)
	oracle, err := venus.NewOracle(common.HexToAddress(cfg.Oracle), c)
	require.NoError(t, err)

	markets, err := comptroller.GetAllMarkets(nil)
	require.NoError(t, err)

	var addresses []common.Address
	name := make(map[string]string)

	for _, market := range markets {
		addresses = append(addresses, market)
		token, err := venus.NewVbep20(market, c)
		require.NoError(t, err)

		symbol, err := token.Symbol(nil)
		name[strings.ToLower(market.String())] = symbol
		//fmt.Printf("name:%v, address:%v\n", name[strings.ToLower(market.String())], market.String())
	}

	accounts := []string{
		"0x332E2Dcd239Bb40d4eb31bcaE213F9F06017a4F3",
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

		fmt.Printf("mintVAI:%v\n", mintVAIS)
		require.NoError(t, err)

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
	}
}

func TestSyncMarketsAndPrices(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg)

	sync.doSyncMarketsAndPrices()

	for symbol, detail := range TokenDetail {
		t.Logf("symbol:%v, detail:%v\n", symbol, detail)
	}

	for address, symbol := range AddressToSymbol {
		t.Logf("address:%v, symbol:%v", address, symbol)
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

////Filter with ethereum.SubscribeFilterLogs,
//func TestSubscribeBorrowEvent(t *testing.T) {
//	t.Skip("this case need websocket support")
//	ctx := context.Background()
//	rpcURL := "wss://bsc-ws-node.nariox.org:443" //"ws://42.3.146.198:21994" //"https://bsc-dataseed1.binance.org"
//	c, err := ethclient.Dial(rpcURL)
//	require.NoError(t, err)
//
//	_, err = c.BlockNumber(ctx)
//	require.NoError(t, err)
//
//	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
//	require.NoError(t, err)
//
//	vUSDCAddress := common.HexToAddress("0xeca88125a5adbe82614ffc12d0db554e2e2867c8")
//	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")
//
//	query := ethereum.FilterQuery{
//		FromBlock: big.NewInt(14747566),
//		ToBlock:   big.NewInt(14747568),
//		Addresses: []common.Address{vUSDCAddress},
//		Topics:    [][]common.Hash{{topicBorrow}},
//	}
//
//	logs := make(chan types.Log, 10)
//	sub, err := c.SubscribeFilterLogs(context.Background(), query, logs)
//	require.NoError(t, err)
//
//	for {
//		select {
//		case err := <-sub.Err():
//			log.Fatal(err)
//		case vLog := <-logs:
//			fmt.Println(vLog) // pointer to event log
//			fmt.Println(vLog.BlockNumber)
//			var borrowEvent venus.Vbep20Borrow
//			err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", vLog.Data)
//			require.NoError(t, err)
//		}
//	}
//}

func TestScanAllBorrowers(t *testing.T) {
	ctx := context.Background()

	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	sync := NewSyncer(c, db, cfg)
	start := big.NewInt(13000000) //big.NewInt(14747565)
	db.Put(dbm.KeyLastHandledHeight, start.Bytes(), nil)
	db.Put(dbm.KeyBorrowerNumber, big.NewInt(0).Bytes(), nil)

	sync.Start()
	time.Sleep(time.Second * 120)
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

	sync := NewSyncer(c, db, cfg)
	star := big.NewInt(int64(height - 5000))
	db.Put(dbm.KeyLastHandledHeight, star.Bytes(), nil)
	db.Put(dbm.KeyBorrowerNumber, big.NewInt(0).Bytes(), nil)

	sync.Start()
	time.Sleep(time.Second * 15)
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
