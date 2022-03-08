package server

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestComptrollerCaller_LiquidateCalculateSeizeTokens(t *testing.T) {
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

	syncer := NewSyncer(c, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey, feededPricesCh, liquidationCh, priorityliquidationCh)
	comptroller := syncer.comptroller

	//opts := &bind.CallOpts{
	//	BlockNumber: big.NewInt(15839356),
	//}

	repayAmount, _ := decimal.NewFromString("1057682925462918391813")
	t.Logf("vUSDT:%v, vXVS:%v, vCAKE:%v", syncer.tokens["vUSDT"].Address, syncer.tokens["vXVS"].Address, syncer.tokens["vCAKE"].Address)

	_, bigVXVSSeizedAmount, err := comptroller.LiquidateCalculateSeizeTokens(nil, syncer.tokens["vUSDT"].Address, syncer.tokens["vXVS"].Address, repayAmount.BigInt())
	xvsExchangeRate, _ := decimal.NewFromString("201107934228759268020040113")
	xvsPrice := decimal.NewFromInt(8320000000000000000)
	xvsValue := decimal.NewFromBigInt(bigVXVSSeizedAmount, 0).Mul(xvsPrice).Mul(xvsExchangeRate).Div(decimal.New(1, 54))
	t.Logf("xvsSeizedAmount:%v, xvsSeizedValue:%v\n", bigVXVSSeizedAmount, xvsValue)

	_, bigVCAKESeizedAmount, err := comptroller.LiquidateCalculateSeizeTokens(nil, syncer.tokens["vUSDT"].Address, syncer.tokens["vCAKE"].Address, repayAmount.BigInt())
	cakeExchangeRate, _ := decimal.NewFromString("235572059761730000654372979")
	cakePrice := decimal.NewFromInt(6010000000000000000)
	cakeValue := decimal.NewFromBigInt(bigVCAKESeizedAmount, 0).Mul(cakePrice).Mul(cakeExchangeRate).Div(decimal.New(1, 54))

	t.Logf("cakeSeizedAmount:%v, cakeSeizedValue:%v\n", bigVCAKESeizedAmount, cakeValue)

}
