package server

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"github.com/readygo67/LiquidationBot/venus"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMapStructAssignment(t *testing.T) {
	testmap := make(map[string]*TokenInfo)
	tokenInfo := &TokenInfo{
		Price: decimal.Zero,
	}
	testmap["usdt"] = tokenInfo
	testmap["usdt"].Price = decimal.NewFromInt(1)
}

func TestFilterUSDCLiquidateBorrowEvent(t *testing.T) {
	ctx := context.Background()
	//cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	topicLiquidateBorrow := common.HexToHash("0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52")
	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	require.NoError(t, err)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(15803152),
		//ToBlock:   big.NewInt(1563526),
		Addresses: []common.Address{common.HexToAddress("0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8")}, //usdc
		Topics:    [][]common.Hash{{topicLiquidateBorrow}},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	require.NoError(t, err)
	logger.Printf("start Time:%v\n", time.Now())
	for i, log := range logs {
		var eve venus.Vbep20LiquidateBorrow
		err = vbep20Abi.UnpackIntoInterface(&eve, "LiquidateBorrow", log.Data)
		logger.Printf("%v height:%v, txhash:%v, liquidator:%v borrower:%v, repayAmount:%v, collateral:%v, seizedAmount:%v\n", (i + 1), log.BlockNumber, log.TxHash, eve.Liquidator, eve.Borrower, eve.RepayAmount, eve.VTokenCollateral, eve.SeizeTokens)
	}
	logger.Printf("end Time:%v\n", time.Now())
}

func TestFilterSubscribeUSDCLiquidateBorrowEvent(t *testing.T) {
	ctx := context.Background()
	//cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	topicLiquidateBorrow := common.HexToHash("0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52")
	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	require.NoError(t, err)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(15803152),
		ToBlock:   big.NewInt(15603526),
		Addresses: []common.Address{common.HexToAddress("0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8")}, //usdc
		Topics:    [][]common.Hash{{topicLiquidateBorrow}},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	require.NoError(t, err)
	logger.Printf("start Time:%v\n", time.Now())
	for i, log := range logs {
		var eve venus.Vbep20LiquidateBorrow
		err = vbep20Abi.UnpackIntoInterface(&eve, "LiquidateBorrow", log.Data)
		logger.Printf("%v height:%v, txhash:%v, liquidator:%v borrower:%v, repayAmount:%v, collateral:%v, seizedAmount:%v\n", (i + 1), log.BlockNumber, log.TxHash, eve.Liquidator, eve.Borrower, eve.RepayAmount, eve.VTokenCollateral, eve.SeizeTokens)
	}
	logger.Printf("end Time:%v\n", time.Now())
}

func TestFilterAllVTokensLiquidateBorrowEvent(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	syncer := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)

	topicLiquidateBorrow := common.HexToHash("0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52")

	var addresses []common.Address
	for _, token := range syncer.tokens {
		addresses = append(addresses, token.Address)
	}

	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	require.NoError(t, err)
	monitorStartHeight := uint64(15603526)

	for i := 0; i < 10; i++ {
		monitorEndHeight, err := c.BlockNumber(context.Background())
		if err != nil {
			monitorEndHeight = monitorStartHeight
		}
		logger.Printf("%vth sync monitor LiquidationBorrow event, startHeight:%v, endHeight:%v \n", (i + 1), monitorStartHeight, monitorEndHeight)

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(monitorStartHeight)),
			ToBlock:   big.NewInt(int64(monitorEndHeight)),
			Addresses: addresses, //usdc
			Topics:    [][]common.Hash{{topicLiquidateBorrow}},
		}

		logs, err := c.FilterLogs(context.Background(), query)
		if err == nil {
			for _, log := range logs {
				var eve venus.Vbep20LiquidateBorrow
				vbep20Abi.UnpackIntoInterface(&eve, "LiquidateBorrow", log.Data)
				logger.Printf("LiquidateBorrow event happen @ height:%v, txhash:%v, liquidator:%v borrower:%v, repayAmount:%v, collateral:%v, seizedAmount:%v\n", log.BlockNumber, log.TxHash, eve.Liquidator, eve.Borrower, eve.RepayAmount, eve.VTokenCollateral, eve.SeizeTokens)
			}

			monitorStartHeight = monitorEndHeight
		}

		time.Sleep(30 * time.Second)
	}
}

func TestFilterAllVTokensLiquidateBorrowEvent1(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	syncer := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)

	topicLiquidateBorrow := common.HexToHash("0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52")

	var addresses []common.Address
	for _, token := range syncer.tokens {
		addresses = append(addresses, token.Address)
	}

	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	require.NoError(t, err)
	monitorStartHeight := uint64(15633526)

	monitorEndHeight, err := c.BlockNumber(context.Background())
	if err != nil {
		monitorEndHeight = monitorStartHeight
	}
	logger.Printf("sync monitor LiquidationBorrow event, startHeight:%v, endHeight:%v \n", monitorStartHeight, monitorEndHeight)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(monitorStartHeight)),
		ToBlock:   big.NewInt(int64(monitorEndHeight)),
		Addresses: addresses, //usdc
		Topics:    [][]common.Hash{{topicLiquidateBorrow}},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	if err == nil {
		for _, log := range logs {
			var eve venus.Vbep20LiquidateBorrow
			vbep20Abi.UnpackIntoInterface(&eve, "LiquidateBorrow", log.Data)
			logger.Printf("LiquidateBorrow event happen @ height:%v, txhash:%v, liquidator:%v borrower:%v, repayAmount:%v, collateral:%v, seizedAmount:%v\n", log.BlockNumber, log.TxHash, eve.Liquidator, eve.Borrower, eve.RepayAmount, eve.VTokenCollateral, eve.SeizeTokens)
		}

		monitorStartHeight = monitorEndHeight
	}
}

func TestMonitorPricesInTxPool(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	t.Logf("before MonitorTxPoolLoop\n")
	sync.wg.Add(2)
	go sync.MonitorTxPoolLoop()
	go func() {
		defer sync.wg.Done()
		for {
			select {
			case <-sync.quitCh:
				return
			case data := <-sync.feededPriceCh:
				logger.Printf("feedPrice:%v\n", data)
			}
		}
	}()
	t.Logf("sleep 5 minutes\n")
	time.Sleep(300 * time.Second)
	close(sync.quitCh)
}

func TestAddressEqual(t *testing.T) {
	account1 := common.HexToAddress("0x26a27B56308FaB4ffE9ad5C80BB0C3Da9152e833")
	account2 := common.HexToAddress("0x26a27B56308FaB4ffE9ad5C80BB0C3Da9152e833")
	account3 := common.HexToAddress("0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8")
	require.True(t, account1 == account2)
	require.False(t, account1 == account3)
}

func TestSelectPathWithMinAmountIn(t *testing.T) {
	cfg, err := config.New("../config.yml")
	rpcURL := "wss://shy-fragrant-moon.bsc.quiknode.pro/7d01e3d9cc6adce754a7e7cf19bf875b031fbaa2/"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	pancakeRouter, err := venus.NewIPancakeRouter02(common.HexToAddress(cfg.PancakeRouter), c)
	require.NoError(t, err)
	paths := make(map[string][][]common.Address)

	/*[ [0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d 0x55d398326f99059fF775485246999027B3197955]]
	 */
	//paths["vUSDC:vUSDT"] = [][]common.Address{
	//	{common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"), common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"), common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")},
	//	{common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"), common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"), common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")},
	//	{common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"), common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")},
	//}

	// path_test.go:67: VAI:vBETH:[[0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x250632378E573c6Be1AC2f97Fcdf00515d0Aa91B] [0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 0x250632378E573c6Be1AC2f97Fcdf00515d0Aa91B]]
	paths["VAI:vBETH"] = [][]common.Address{
		{common.HexToAddress("0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7"), common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"), common.HexToAddress("0x250632378E573c6Be1AC2f97Fcdf00515d0Aa91B")},
		{common.HexToAddress("0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7"), common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"), common.HexToAddress("0x250632378E573c6Be1AC2f97Fcdf00515d0Aa91B")},
	}

	sync := Syncer{
		pancakeRouter: pancakeRouter,
		paths:         paths,
	}

	for pair, paths1 := range sync.paths {
		t.Logf("%v:%v\n", pair, paths1)

		symobls := strings.Split(pair, ":")
		selected, amountOut, err := sync.selectPathWithMaxAmountOut(symobls[0], symobls[1], decimal.New(1000, 18))
		require.NoError(t, err)
		require.Equal(t, paths1[0], selected)
		selected, amountIn, err := sync.selectPathWithMinAmountIn(symobls[0], symobls[1], amountOut)
		t.Logf("amountIn:%v, amountOut:%v\n", amountIn, amountOut)
	}
}
