package venus

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBasics(t *testing.T) {
	ctx := context.Background()
	rpcURL := "http://42.3.146.198:21993" //"https://bsc-dataseed1.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)
	//fmt.Println(height)

	//unitroller basic
	unitrollerAddress := common.HexToAddress("0xfD36E2c2a6789Db23113685031d7F16329158384")
	unitroller, err := NewUnitroller(unitrollerAddress, c)
	comptrollerAddress, err := unitroller.ComptrollerImplementation(nil)
	require.Equal(t, "0xB49416B2Fb86EEd9152f6a53c02Bf34c965E8436", comptrollerAddress.String())
	pendingComptroolerImplementation, err := unitroller.PendingComptrollerImplementation(nil)
	require.Equal(t, "0x0000000000000000000000000000000000000000", pendingComptroolerImplementation.String())

	//comtroller basic
	comptrollerAddress = common.HexToAddress("0xfD36E2c2a6789Db23113685031d7F16329158384")
	comptroller, err := NewComptroller(comptrollerAddress, c)

	account := common.HexToAddress("0x9bd3e72f2f1ca05ad8d4489ec870bf2478b10397")
	assetsIns, err := comptroller.GetAssetsIn(nil, account)
	require.NotEqual(t, "0x0000000000000000000000000000000000000000", assetsIns[0].String())
	//logger.Printf("%+v\n", assetsIns)

	_, liquidity, shorfall, err := comptroller.GetAccountLiquidity(nil, account)
	logger.Printf("liquidity:%v, shortfall:%v\n", liquidity, shorfall)

	//vbep20 basic
	vUSDCAddress := common.HexToAddress("0xeca88125a5adbe82614ffc12d0db554e2e2867c8")
	vusdc, err := NewVbep20(vUSDCAddress, c)
	name, err := vusdc.Symbol(nil)
	require.Equal(t, "vUSDC", name)

	_, balance, loan, exchangeRate, err := vusdc.GetAccountSnapshot(nil, account)
	logger.Printf("balance:%v, loan:%v, exchangeRate:%v\n", balance, loan, exchangeRate)
}
