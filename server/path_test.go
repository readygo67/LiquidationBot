package server

import (
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
)

func TestCheckFlashLoanPool(t *testing.T) {
	//ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "ws://42.3.146.198:21994" //"https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)
	for symbol, pools := range sync.flashLoanPools {
		t.Logf("%v pools:%v\n", symbol, pools)

		bep20, _ := venus.NewBep20(VAIAddress, c)
		if symbol != "VAI" {
			bep20, _ = venus.NewBep20(sync.tokens[symbol].UnderlyingAddress, c)
		}
		var balances []*big.Int
		for _, pair := range pools {
			balance, _ := bep20.BalanceOf(nil, pair)
			balances = append(balances, balance)
		}

		for i := 0; i < len(balances)-1; i++ {
			for j := i + 1; j < len(balances); j++ {
				require.True(t, balances[i].Cmp(balances[j]) != -1)
			}
		}
	}

	for pair, paths := range sync.paths {
		t.Logf("%v paths:%v\n", pair, paths)
	}
}

func TestBuildSwapPath(t *testing.T) {
	//ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)

	for pair, paths := range sync.paths {
		t.Logf("%v:%v\n", pair, paths)

		symobls := strings.Split(pair, ":")
		_, amountOut, _ := sync.selectPathWithMaxAmountOut(symobls[0], symobls[1], decimal.New(1000, 18))
		//require.NoError(t, err)
		//require.Equal(t, paths[0], selected)
		_, amountIn, _ := sync.selectPathWithMinAmountIn(symobls[0], symobls[1], amountOut)
		t.Logf("amountIn:%v, amountOut:%v\n", amountIn, amountOut)
	}
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
