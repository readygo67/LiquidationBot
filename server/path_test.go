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

func TestBuildFlashLoanPool(t *testing.T) {
	cfg, err := config.New("../config.yml")
	require.NoError(t, err)
	rpcURL := "ws://42.3.146.198:21994"
	c, err := ethclient.Dial(rpcURL)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)

	for symbol, pairs := range sync.flashLoanPools {
		logger.Printf("%v connection:%v\n", symbol, pairs)
	}

	bep20, err := venus.NewBep20(sync.tokens["vUSDT"].UnderlyingAddress, sync.c)
	require.NoError(t, err)

	for _, pair := range sync.flashLoanPools["vUSDT"] {
		balance, err := bep20.BalanceOf(nil, pair)
		require.NoError(t, err)
		t.Logf("balance:%v\n", balance)
	}
}

func TestCheckFlashLoanPool(t *testing.T) {
	//ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)
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

func TestVerifyFlashPoolSorting(t *testing.T) {
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	bep20, err := venus.NewBep20(common.HexToAddress("0xba2ae424d960c26247dd6c32edc70b295c744c43"), c)

	var pairs []common.Address
	var balances []*big.Int
	pairStrs := []string{"0xac109C8025F272414fd9e2faA805a583708A017f", "0xE27859308ae2424506D1ac7BF5bcb92D6a73e211", "0xA1d7621ADaDB86c83779b94D978c48760F5DcE67", "0x0fA119e6a12e3540c2412f9EdA0221Ffd16a7934", "0xCb0c07f31A3fa4876b61056A21c58316BD1fA331"}
	for _, pairStr := range pairStrs {
		pair := common.HexToAddress(pairStr)
		pairs = append(pairs, pair)

		balance, err := bep20.BalanceOf(nil, pair)
		require.NoError(t, err)
		balances = append(balances, balance)
		t.Logf("pair:%v, balance:%v\n", pair, balance)
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

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)

	for pair, paths := range sync.paths {
		t.Logf("%v:%v\n", pair, paths)

		symobls := strings.Split(pair, ":")
		selected, amountOut, err := sync.selectPathWithMaxAmountOut(symobls[0], symobls[1], decimal.New(1000, 18))
		require.NoError(t, err)
		require.Equal(t, paths[0], selected)
		selected, amountIn, err := sync.selectPathWithMinAmountIn(symobls[0], symobls[1], amountOut)
		t.Logf("amountIn:%v, amountOut:%v\n", amountIn, amountOut)
	}
}

func TestBuildOnePairSwapPath(t *testing.T) {
	//ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)

	srcSymbol := "vETH"
	dstSymbol := "vDOT"
	srcToken := sync.tokens[srcSymbol]
	dstToken := sync.tokens[dstSymbol]
	amountIn := decimal.New(1000, 18).BigInt()
	amountOuts := make([][]*big.Int, 0)

	tmpPaths := make([][]common.Address, 4)
	paths := make([][]common.Address, 0)

	tmpPaths[0] = []common.Address{srcToken.UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[1] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBNB"].UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[2] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBUSD"].UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[3] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vUSDT"].UnderlyingAddress, dstToken.UnderlyingAddress}

	for i := 0; i < 4; i++ {
		tmpAmounOuts, err := sync.pancakeRouter.GetAmountsOut(nil, amountIn, tmpPaths[i])
		t.Logf("path:%v, amounts:%v, err:%v\n", tmpPaths[i], tmpAmounOuts, err)
		if err == nil && tmpAmounOuts[len(tmpPaths[i])-1].Cmp(BigZero) == 1 {
			paths = append(paths, tmpPaths[i])
			amountOuts = append(amountOuts, tmpAmounOuts)

		}
	}

	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			ilen := len(amountOuts[i])
			jLen := len(amountOuts[j])
			if amountOuts[i][ilen-1].Cmp(amountOuts[j][jLen-1]) == -1 {
				amountOuts[i], amountOuts[j] = amountOuts[j], amountOuts[i]
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
	t.Logf("after sorting\n")
	for i := 0; i < len(paths); i++ {
		t.Logf("path:%v, amounts:%v\n", paths[i], amountOuts[i])
	}

	t.Logf("paths:%v", paths)
}

func TestBuildOnePairSwapPath1(t *testing.T) {
	//ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)

	srcSymbol := "vETH"
	dstSymbol := "vBUSD"
	srcToken := sync.tokens[srcSymbol]
	dstToken := sync.tokens[dstSymbol]
	amountIn := decimal.New(1000, 18).BigInt()
	amountOuts := make([][]*big.Int, 0)

	tmpPaths := make([][]common.Address, 4)
	paths := make([][]common.Address, 0)

	tmpPaths[0] = []common.Address{srcToken.UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[1] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBNB"].UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[2] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBUSD"].UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[3] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vUSDT"].UnderlyingAddress, dstToken.UnderlyingAddress}

	for i := 0; i < 4; i++ {
		tmpAmounOuts, err := sync.pancakeRouter.GetAmountsOut(nil, amountIn, tmpPaths[i])
		t.Logf("path:%v, amounts:%v, err:%v\n", tmpPaths[i], tmpAmounOuts, err)
		if err == nil && tmpAmounOuts[len(tmpPaths[i])-1].Cmp(BigZero) == 1 {
			paths = append(paths, tmpPaths[i])
			amountOuts = append(amountOuts, tmpAmounOuts)

		}
	}

	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			ilen := len(amountOuts[i])
			jLen := len(amountOuts[j])
			if amountOuts[i][ilen-1].Cmp(amountOuts[j][jLen-1]) == -1 {
				amountOuts[i], amountOuts[j] = amountOuts[j], amountOuts[i]
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
	t.Logf("after sorting\n")
	for i := 0; i < len(paths); i++ {
		t.Logf("path:%v, amounts:%v\n", paths[i], amountOuts[i])
	}

	t.Logf("paths:%v", paths)
}

func TestBuildOnePairSwapPath2(t *testing.T) {
	//ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)

	srcSymbol := "vETH"
	dstSymbol := "vCAN"
	srcToken := sync.tokens[srcSymbol]
	dstToken := sync.tokens[dstSymbol]
	amountIn := decimal.New(1000, 18).BigInt()
	amountOuts := make([][]*big.Int, 0)

	tmpPaths := make([][]common.Address, 4)
	paths := make([][]common.Address, 0)

	tmpPaths[0] = []common.Address{srcToken.UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[1] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBNB"].UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[2] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBUSD"].UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[3] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vUSDT"].UnderlyingAddress, dstToken.UnderlyingAddress}

	for i := 0; i < 4; i++ {
		tmpAmounOuts, err := sync.pancakeRouter.GetAmountsOut(nil, amountIn, tmpPaths[i])
		t.Logf("path:%v, amounts:%v, err:%v\n", tmpPaths[i], tmpAmounOuts, err)
		if err == nil && tmpAmounOuts[len(tmpPaths[i])-1].Cmp(BigZero) == 1 {
			paths = append(paths, tmpPaths[i])
			amountOuts = append(amountOuts, tmpAmounOuts)

		}
	}

	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			ilen := len(amountOuts[i])
			jLen := len(amountOuts[j])
			if amountOuts[i][ilen-1].Cmp(amountOuts[j][jLen-1]) == -1 {
				amountOuts[i], amountOuts[j] = amountOuts[j], amountOuts[i]
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
	t.Logf("after sorting\n")
	for i := 0; i < len(paths); i++ {
		t.Logf("path:%v, amounts:%v\n", paths[i], amountOuts[i])
	}

	t.Logf("paths:%v", paths)
}

func TestBuildOnePairSwapPath3(t *testing.T) {
	//ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)

	srcSymbol := "vBUSD"
	dstSymbol := "vXVS"
	srcToken := sync.tokens[srcSymbol]
	dstToken := sync.tokens[dstSymbol]
	amountIn := decimal.New(1000, 18).BigInt()
	amountOuts := make([][]*big.Int, 0)

	tmpPaths := make([][]common.Address, 4)
	paths := make([][]common.Address, 0)

	tmpPaths[0] = []common.Address{srcToken.UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[1] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBNB"].UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[2] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBUSD"].UnderlyingAddress, dstToken.UnderlyingAddress}
	tmpPaths[3] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vUSDT"].UnderlyingAddress, dstToken.UnderlyingAddress}

	for i := 0; i < 4; i++ {
		tmpAmounOuts, err := sync.pancakeRouter.GetAmountsOut(nil, amountIn, tmpPaths[i])
		t.Logf("path:%v, amounts:%v, err:%v\n", tmpPaths[i], tmpAmounOuts, err)
		if err == nil && tmpAmounOuts[len(tmpPaths[i])-1].Cmp(BigZero) == 1 {
			paths = append(paths, tmpPaths[i])
			amountOuts = append(amountOuts, tmpAmounOuts)

		}
	}

	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			ilen := len(amountOuts[i])
			jLen := len(amountOuts[j])
			if amountOuts[i][ilen-1].Cmp(amountOuts[j][jLen-1]) == -1 {
				amountOuts[i], amountOuts[j] = amountOuts[j], amountOuts[i]
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
	t.Logf("after sorting\n")
	for i := 0; i < len(paths); i++ {
		t.Logf("path:%v, amounts:%v\n", paths[i], amountOuts[i])
	}

	t.Logf("paths:%v", paths)
}

func TestBuildOnePairSwapPath4(t *testing.T) {
	//ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)

	symbols := []string{"vBUSD", "vXVS", "vCAKE", "vDOT", "vBNB", "vBETH", "vETH"}
	for _, srcSymbol := range symbols {
		for _, dstSymbol := range symbols {
			if dstSymbol == srcSymbol {
				continue
			}
			t.Logf("%v:%v\n", srcSymbol, dstSymbol)
			srcToken := sync.tokens[srcSymbol]
			dstToken := sync.tokens[dstSymbol]
			amountIn := decimal.New(1000, 18).BigInt()
			amountOuts := make([][]*big.Int, 0)

			tmpPaths := make([][]common.Address, 4)
			paths := make([][]common.Address, 0)

			tmpPaths[0] = []common.Address{srcToken.UnderlyingAddress, dstToken.UnderlyingAddress}
			tmpPaths[1] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBNB"].UnderlyingAddress, dstToken.UnderlyingAddress}
			tmpPaths[2] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBUSD"].UnderlyingAddress, dstToken.UnderlyingAddress}
			tmpPaths[3] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vUSDT"].UnderlyingAddress, dstToken.UnderlyingAddress}

			for i := 0; i < 4; i++ {
				tmpAmounOuts, err := sync.pancakeRouter.GetAmountsOut(nil, amountIn, tmpPaths[i])
				//t.Logf("path:%v, amounts:%v, err:%v\n", tmpPaths[i], tmpAmounOuts, err)
				if err == nil && tmpAmounOuts[len(tmpPaths[i])-1].Cmp(BigZero) == 1 {
					paths = append(paths, tmpPaths[i])
					amountOuts = append(amountOuts, tmpAmounOuts)

				}
			}

			for i := 0; i < len(paths)-1; i++ {
				for j := i + 1; j < len(paths); j++ {
					ilen := len(amountOuts[i])
					jLen := len(amountOuts[j])
					if amountOuts[i][ilen-1].Cmp(amountOuts[j][jLen-1]) == -1 {
						amountOuts[i], amountOuts[j] = amountOuts[j], amountOuts[i]
						paths[i], paths[j] = paths[j], paths[i]
					}
				}
			}
			t.Logf("after sorting\n")
			for i := 0; i < len(paths); i++ {
				t.Logf("path:%v, amounts:%v\n", paths[i], amountOuts[i])
			}

			//t.Logf("paths:%v", paths)
		}
	}
}

func TestBuildOnePairSwapPath5(t *testing.T) {
	//ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "https://bsc-dataseed.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	db, err := dbm.NewDB("testdb1")
	require.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	liquidationCh := make(chan *Liquidation, 64)
	priorityliquidationCh := make(chan *Liquidation, 64)
	feededPricesCh := make(chan *FeededPrices, 64)

	sync := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)

	symbols := []string{"vBETH", "vETH"} //vBETH's underlyingToken is BETH, vETH's underlyingToken is ETH, their pancake pair's depth is not enough.
	for _, srcSymbol := range symbols {
		for _, dstSymbol := range symbols {
			if dstSymbol == srcSymbol {
				continue
			}
			t.Logf("%v:%v\n", srcSymbol, dstSymbol)
			srcToken := sync.tokens[srcSymbol]
			dstToken := sync.tokens[dstSymbol]
			amountIn := decimal.New(1000, 18).BigInt()
			amountOuts := make([][]*big.Int, 0)

			tmpPaths := make([][]common.Address, 4)
			paths := make([][]common.Address, 0)

			tmpPaths[0] = []common.Address{srcToken.UnderlyingAddress, dstToken.UnderlyingAddress}
			tmpPaths[1] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBNB"].UnderlyingAddress, dstToken.UnderlyingAddress}
			tmpPaths[2] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vBUSD"].UnderlyingAddress, dstToken.UnderlyingAddress}
			tmpPaths[3] = []common.Address{srcToken.UnderlyingAddress, sync.tokens["vUSDT"].UnderlyingAddress, dstToken.UnderlyingAddress}

			for i := 0; i < 4; i++ {
				tmpAmounOuts, err := sync.pancakeRouter.GetAmountsOut(nil, amountIn, tmpPaths[i])
				//t.Logf("path:%v, amounts:%v, err:%v\n", tmpPaths[i], tmpAmounOuts, err)
				if err == nil && tmpAmounOuts[len(tmpPaths[i])-1].Cmp(BigZero) == 1 {
					paths = append(paths, tmpPaths[i])
					amountOuts = append(amountOuts, tmpAmounOuts)

				}
			}

			for i := 0; i < len(paths)-1; i++ {
				for j := i + 1; j < len(paths); j++ {
					ilen := len(amountOuts[i])
					jLen := len(amountOuts[j])
					if amountOuts[i][ilen-1].Cmp(amountOuts[j][jLen-1]) == -1 {
						amountOuts[i], amountOuts[j] = amountOuts[j], amountOuts[i]
						paths[i], paths[j] = paths[j], paths[i]
					}
				}
			}
			t.Logf("after sorting\n")
			for i := 0; i < len(paths); i++ {
				t.Logf("path:%v, amounts:%v\n", paths[i], amountOuts[i])
			}

			t.Logf("paths:%v", paths)
		}
	}
}
