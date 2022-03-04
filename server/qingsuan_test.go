package server

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/venus"
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
	"testing"
)

/*
npx ganache-cli -f http://42.3.146.198:21993 -i 56 -m letter write catalog uncle render nest upset piano coyote glide voice fine betray rebuild margin

Available Accounts
==================
(0) 0x732d0c92E44D62fb594E5F49904cE1055b80b1BB (100 ETH)
(1) 0x2035182b4A5d5D1dC36B16373Aab6904B094b121 (100 ETH)
(2) 0x8d314cfB4Ef38351E842e4b0ef2279d983bb7E91 (100 ETH)
(3) 0xE2316DE05b916C6fE20454c3705F89556324c0F0 (100 ETH)
(4) 0xd53b62BaE62b654cDbcb252D887a9B4a5378a4e0 (100 ETH)
(5) 0xD8F42E25429291F1A32b84f08EF3d21173325c11 (100 ETH)
(6) 0x06EC655baB5D14A9ca9a72921e55C1ec75089cAA (100 ETH)
(7) 0x64dB6f0CB438dC4e58E2d591282a37D112Cef7B4 (100 ETH)
(8) 0x29cA632548bBfC6573b5C2512b69Eb757BE4b5f1 (100 ETH)
(9) 0x6D8411E36517ef704c15614bC4f1e9dc088CAdB7 (100 ETH)

Private Keys
==================
(0) 0x5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384
(1) 0xe65b010b0b6862e74bcc248a4a8255258b77f51bac5675840aab6780893c8a9a
(2) 0x45308292e0e0e41aa4e5da89fdab9a53e87d98b78eb121a09f7b4b6d4e8e6b67
(3) 0x313d9d28ffe4923032153686982edfd89232e40cc0f50940c9dd575d492def9a
(4) 0xce7b2ff6a5d163351212689e85af28ce990f54590a7e9db69109409ca5eedef5
(5) 0x8bd778f5313ee00d78adc5100e6a6c0f7602260e8c73abca6eb131464519d64a
(6) 0x23b17d38f49bea2bf3446d1ae138e53b31aa9eefe477507b622e9aecd58e7205
(7) 0x96725d546d31e815c43a0e0913b159febbfa23aa984c47c00ec5e1dfb16a2008
(8) 0xbefbfbd8382e9db4d92ac1ad7a8fdf6e37c1749ff690383a02ca3e51fbd63c09
(9) 0x03a60947ec3d0e8e2b1b428c25c22346554bf89f4052fc79d11662bc49337fb9

*/
/*
height15325212, account:0x9Dbb3fa7F18C72FDF5fa5D1eC630B0F1F191FAE6, repaySmbol:vBUSD, flashLoanFrom:0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16, repayAddress:0x95c78222B3D6e262426483D42CfA53685A67Ab9D, repayValue:40015902181827895653.4091979725, repayAmount:40010220770488707027 seizedSymbol:vBUSD, seizedAddress:0x95c78222B3D6e262426483D42CfA53685A67Ab9D, seizedCTokenAmount:205825257122, seizedUnderlyingTokenAmount:44011242845431854992.1699406601840677, seizedUnderlyingTokenValue:44017492397904663470.1469737995878732
calculateSeizedTokenAmount case1: seizedSymbol == repaySymbol and symbol is a stable coin, account:0x9Dbb3fa7F18C72FDF5fa5D1eC630B0F1F191FAE6, symbol:vBUSD, seizedAmount:44011242845431854992.1699406601840677, returnAmout:40110497013021260138, gasFee:953555625000000000 profit:2.9477441094180513
case1, profitable liquidation catched:&{0x9Dbb3fa7F18C72FDF5fa5D1eC630B0F1F191FAE6 0.9998935931903574 15325033 0001-01-01 00:00:00 +0000 UTC}, profit:2.9477441094180513
*/
//
//func TestOnlineCase1(t *testing.T) {
//	rpcURL := "http://127.0.0.1:8545"
//	addr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
//	c, err := ethclient.Dial(rpcURL)
//	require.NoError(t, err)
//
//	qs, err := NewQingsuan(addr, c)
//	require.NoError(t, err)
//
//	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
//	require.NoError(t, err)
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
//
//	value, ok := big.NewInt(0).SetString("20000000000000000", 10)
//	require.True(t, ok)
//	gasPrice, err := c.SuggestGasPrice(context.Background())
//	require.NoError(t, err)
//
//	chainID, err := c.NetworkID(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//	tx := types.NewTransaction(nonce, addr, value, 300000, gasPrice, nil)
//
//	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = c.SendTransaction(context.Background(), signedTx)
//	require.NoError(t, err)
//	t.Logf("send 1 ether, hash:%v", signedTx.Hash())
//
//	balance, err := c.BalanceAt(context.Background(), addr, nil)
//	require.NoError(t, err)
//	t.Logf("balance:%v\n", balance)
//	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(56))
//	auth.Value = big.NewInt(0) // in wei
//	auth.Nonce = big.NewInt(int64(nonce + 1))
//	auth.GasPrice = gasPrice
//
//	addresses := []common.Address{
//		common.HexToAddress("0x95c78222B3D6e262426483D42CfA53685A67Ab9D"),
//		common.HexToAddress("0x95c78222B3D6e262426483D42CfA53685A67Ab9D"),
//		common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"),
//		common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"),
//		common.HexToAddress("0x9Dbb3fa7F18C72FDF5fa5D1eC630B0F1F191FAE6"),
//	}
//
//	emptyPath := []common.Address{}
//
//	repayAmount, ok := big.NewInt(0).SetString("40011357893447481535", 10)
//	require.True(t, ok)
//
//	owner, err := qs.Owner(nil)
//	require.NoError(t, err)
//	t.Logf("owner:%v", owner)
//
//	tx, err = qs.Qingsuan(auth, big.NewInt(1), common.HexToAddress("0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16"), emptyPath, emptyPath, addresses, repayAmount)
//	require.NoError(t, err)
//	t.Logf("qingsuan hash:%v", tx.Hash())
//}

func TestDeployContract(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	expectedAddr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), from)
	require.NoError(t, err)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	require.NoError(t, err)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = 6000000

	addr, tx, qs, err := DeployUniFlashSwap(auth, c)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, addr)
	t.Logf("hash:%v", tx.Hash())

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, from, owner)
}

/*

	calculateSeizedTokenAmount case1: seizedSymbol == repaySymbol and symbol is a stable coin
	height:15509431, account:0x76f8804F869b49D11f0F7EcbA37FfA113281D3AD, repaySymbol:vUSDT, repayUnderlyingAmount:4039960990484884995, seizedSymbol:vUSDT, seizedVTokenAmount:20597993926, seizedUnderlyingAmount:4443957089270753095.4060116453926008, seizedValue:4446703454751922420.8189725605894534, flashLoanReturnAmout:4050086205999884703, remain:393870883270868392.4060116453926008, gasFee:2740680000000000000, profit:-2.3465657045232702
	flashLoanFrom:0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE, path1:<nil>, path2:<nil>, addresses:[0xfD5840Cd36d94D7229439859C0112a4185BC0255 0xfD5840Cd36d94D7229439859C0112a4185BC0255 0x55d398326f99059fF775485246999027B3197955 0x55d398326f99059fF775485246999027B3197955 0x76f8804F869b49D11f0F7EcbA37FfA113281D3AD]	flashLoanFrom := common.HexToAddress("0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE")

	Transaction: 0x29cc2c6472c4520bf1f6ad05bbd2eae9a951f2669b11a29578e33358899a7297
  	Contract created: 0x55c032aef9d353a5f7562ecba115f9b994147cc1
  	Gas usage: 3247721
  	Block Number: 15509549
  	Block Time: Wed Feb 23 2022 22:31:51 GMT+0800 (中国标准时间)



  	Transaction: 0x133d8de8450cc963a7f0855146ec151daa2ad5bf549bcf5d6f09731978ac160f
 	 Gas usage: 752659
 	 Block Number: 15509550
  	Block Time: Wed Feb 23 2022 22:31:54 GMT+0800 (中国标准时间)
*/
func TestOnlineCase1(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	expectedAddr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), from)
	require.NoError(t, err)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	require.NoError(t, err)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = 6000000

	liquidatorAddr, tx, qs, err := DeployUniFlashSwap(auth, c)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, liquidatorAddr)
	t.Logf("hash:%v", tx.Hash())

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, from, owner)

	flashLoanFrom := common.HexToAddress("0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE")
	addresses := []common.Address{
		common.HexToAddress("0xfD5840Cd36d94D7229439859C0112a4185BC0255"),
		common.HexToAddress("0xfD5840Cd36d94D7229439859C0112a4185BC0255"),
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
		common.HexToAddress("0x76f8804F869b49D11f0F7EcbA37FfA113281D3AD"),
	}
	emptyPath := []common.Address{}
	repayAmount, ok := big.NewInt(0).SetString("4039960990484884995", 10)
	require.True(t, ok)

	//nonce, err = c.PendingNonceAt(context.Background(), from)
	//require.NoError(t, err)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.GasLimit = 1800000

	liquidator, err := venus.NewIQingsuan(liquidatorAddr, c)
	if err != nil {
		panic(err)
	}
	syncer := Syncer{
		c:          c,
		liquidator: liquidator,
		PrivateKey: privateKey,
	}
	gasLimit := uint64(5000000)
	//tx, err =qs.Qingsuan(auth, big.NewInt(1), flashLoanFrom, emptyPath, emptyPath, addresses, repayAmount)
	tx, err = syncer.doLiquidation(big.NewInt(1), flashLoanFrom, emptyPath, emptyPath, addresses, repayAmount, gasPrice, gasLimit)
	require.NoError(t, err)
	t.Logf("tx:%+v", tx)
	t.Logf("qingsuan hash:%v", tx.Hash())

	got, pending, err := c.TransactionByHash(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.False(t, pending)
	require.Equal(t, tx.Hash(), got.Hash())
}

/*
	calculateSeizedTokenAmount case2: seizedSymbol == repaySymbol and symbol is not a stable coin
	height:15510031, account:0xFAbE4C180b6eDad32eA0Cf56587c54417189e422, repaySymbol:vETH, repayUnderlyingAmount:9838914573278510, seizedSymbol:vETH, seizedVTokenAmount:53561998, seizedUnderlyingAmount:10822806021371275.6449623630456059, seizedValue:29391202096475396244.5818753391497113, flashLoanReturnAmout:9863573507046126, remain:959232514325149.6449623630456059, gasFee:3189539220000000000, profit:-0.5988276616840223
	flashLoanFrom:0x74E4716E431f45807DCF19f284c7aA99F18a4fbc, path1:<nil>, path2:[0x2170Ed0880ac9A755fd29B2688956BD959F933F8 0x55d398326f99059fF775485246999027B3197955], addresses:[0xf508fCD89b8bd15579dc79A6827cB4686A3592c8 0xf508fCD89b8bd15579dc79A6827cB4686A3592c8 0x2170Ed0880ac9A755fd29B2688956BD959F933F8 0x2170Ed0880ac9A755fd29B2688956BD959F933F8 0xFAbE4C180b6eDad32eA0Cf56587c54417189e422]
	flashLoanFrom:0x74E4716E431f45807DCF19f284c7aA99F18a4fbc,
 	path1:<nil>,
	path2:[0x2170Ed0880ac9A755fd29B2688956BD959F933F8 0x55d398326f99059fF775485246999027B3197955],
	addresses:[0xf508fCD89b8bd15579dc79A6827cB4686A3592c8 0xf508fCD89b8bd15579dc79A6827cB4686A3592c8 0x2170Ed0880ac9A755fd29B2688956BD959F933F8 0x2170Ed0880ac9A755fd29B2688956BD959F933F8 0xFAbE4C180b6eDad32eA0Cf56587c54417189e422]

  Transaction: 0x29cc2c6472c4520bf1f6ad05bbd2eae9a951f2669b11a29578e33358899a7297
  Contract created: 0x55c032aef9d353a5f7562ecba115f9b994147cc1
  Gas usage: 3247721
  Block Number: 15510552
  Block Time: Wed Feb 23 2022 23:22:05 GMT+0800 (中国标准时间)


  Transaction: 0xfa85a9f8577fbfbfdd2bde0888e4db490a2372ebdc74df038e612aa3469d820a
  Gas usage: 1382255
  Block Number: 15510553
  Block Time: Wed Feb 23 2022 23:24:58 GMT+0800 (中国标准时间)
*/
func TestOnlineCase2(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	expectedAddr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), from)
	require.NoError(t, err)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	require.NoError(t, err)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = 6000000

	liquidatorAddr, tx, qs, err := DeployUniFlashSwap(auth, c)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, liquidatorAddr)
	t.Logf("hash:%v", tx.Hash())

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, from, owner)

	emptyPath := []common.Address{}
	flashLoanFrom := common.HexToAddress("0x74E4716E431f45807DCF19f284c7aA99F18a4fbc")
	path1 := emptyPath
	path2 := []common.Address{
		common.HexToAddress("0x2170Ed0880ac9A755fd29B2688956BD959F933F8"),
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	}
	addresses := []common.Address{
		common.HexToAddress("0xf508fCD89b8bd15579dc79A6827cB4686A3592c8"),
		common.HexToAddress("0xf508fCD89b8bd15579dc79A6827cB4686A3592c8"),
		common.HexToAddress("0x2170Ed0880ac9A755fd29B2688956BD959F933F8"),
		common.HexToAddress("0x2170Ed0880ac9A755fd29B2688956BD959F933F8"),
		common.HexToAddress("0xFAbE4C180b6eDad32eA0Cf56587c54417189e422"),
	}

	repayAmount, ok := big.NewInt(0).SetString("9838914573278510", 10)
	require.True(t, ok)

	//nonce, err = c.PendingNonceAt(context.Background(), from)
	//require.NoError(t, err)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.GasLimit = 1800000

	liquidator, err := venus.NewIQingsuan(liquidatorAddr, c)
	if err != nil {
		panic(err)
	}
	syncer := Syncer{
		c:          c,
		liquidator: liquidator,
		PrivateKey: privateKey,
	}
	gasLimit := uint64(5000000)
	//tx, err =qs.Qingsuan(auth, big.NewInt(1), flashLoanFrom, emptyPath, emptyPath, addresses, repayAmount)
	tx, err = syncer.doLiquidation(big.NewInt(2), flashLoanFrom, path1, path2, addresses, repayAmount, gasPrice, gasLimit)
	require.NoError(t, err)
	t.Logf("tx:%+v", tx)
	t.Logf("qingsuan hash:%v", tx.Hash())

	got, pending, err := c.TransactionByHash(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.False(t, pending)
	require.Equal(t, tx.Hash(), got.Hash())

	usdt, err := NewIERC20(path2[1], c)
	require.NoError(t, err)
	balance, err := usdt.BalanceOf(nil, liquidatorAddr)
	require.NoError(t, err)
	t.Logf("profit:%v\n", balance)
}

/*
	calculateSeizedTokenAmount case3: seizedSymbol != repaySymbol and seizedSymbol stable coin
	height:15510486, account:0x4F41889788528e213692af181B582519BF4Cd30E, repaySymbol:vUSDT, repayUnderlyingAmount:17110085456794186, seizedSymbol:vBUSD, seizedVTokenAmount:88038164, seizedUnderlyingAmount:18830922503073760.7093036444092328, seizedValue:18829978113479309.057130312122162, flashLoanReturnAmout:17152967876485400, remain:1594969626686350.7093036444092328, gasFee:3782100000000000000, profit:-3.7805051103626354
	flashLoanFrom:0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE, path1:[0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 0x55d398326f99059fF775485246999027B3197955], path2:<nil>, addresses:[0xfD5840Cd36d94D7229439859C0112a4185BC0255 0x95c78222B3D6e262426483D42CfA53685A67Ab9D 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 0x55d398326f99059fF775485246999027B3197955 0x4F41889788528e213692af181B582519BF4Cd30E]
	flashLoanFrom:0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE,
    path1:[0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 0x55d398326f99059fF775485246999027B3197955],
    path2:<nil>,
    addresses:[0xfD5840Cd36d94D7229439859C0112a4185BC0255 0x95c78222B3D6e262426483D42CfA53685A67Ab9D 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 0x55d398326f99059fF775485246999027B3197955 0x4F41889788528e213692af181B582519BF4Cd30E]

  	Transaction: 0x29cc2c6472c4520bf1f6ad05bbd2eae9a951f2669b11a29578e33358899a7297
  	Contract created: 0x55c032aef9d353a5f7562ecba115f9b994147cc1
  	Gas usage: 3247721
	Block Number: 15510815
	Block Time: Wed Feb 23 2022 23:35:14 GMT+0800 (中国标准时间)

	Transaction: 0xb2496ca44e40c3fe2fa1ddca6fbaa69952150f88cedcb5b3fd60b6ed1ca77ba7
	Gas usage: 1069882
	Block Number: 15510816
	Block Time: Wed Feb 23 2022 23:37:12 GMT+0800 (中国标准时间)
*/

func TestOnlineCase3_2(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	expectedAddr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), from)
	require.NoError(t, err)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	require.NoError(t, err)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = 6000000

	liquidatorAddr, tx, qs, err := DeployUniFlashSwap(auth, c)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, liquidatorAddr)
	t.Logf("hash:%v", tx.Hash())

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, from, owner)

	emptyPath := []common.Address{}
	flashLoanFrom := common.HexToAddress("0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE")

	path1 := []common.Address{
		common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"),
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	}
	path2 := emptyPath
	addresses := []common.Address{
		common.HexToAddress("0xfD5840Cd36d94D7229439859C0112a4185BC0255"),
		common.HexToAddress("0x95c78222B3D6e262426483D42CfA53685A67Ab9D"),
		common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"),
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
		common.HexToAddress("0x4F41889788528e213692af181B582519BF4Cd30E"),
	}

	repayAmount, ok := big.NewInt(0).SetString("17110085456794186", 10)
	require.True(t, ok)

	//nonce, err = c.PendingNonceAt(context.Background(), from)
	//require.NoError(t, err)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.GasLimit = 1800000

	liquidator, err := venus.NewIQingsuan(liquidatorAddr, c)
	if err != nil {
		panic(err)
	}
	syncer := Syncer{
		c:          c,
		liquidator: liquidator,
		PrivateKey: privateKey,
	}
	gasLimit := uint64(5000000)
	//tx, err =qs.Qingsuan(auth, big.NewInt(1), flashLoanFrom, emptyPath, emptyPath, addresses, repayAmount)
	tx, err = syncer.doLiquidation(big.NewInt(3), flashLoanFrom, path1, path2, addresses, repayAmount, gasPrice, gasLimit)
	require.NoError(t, err)
	t.Logf("tx:%+v", tx)
	t.Logf("qingsuan hash:%v", tx.Hash())

	got, pending, err := c.TransactionByHash(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.False(t, pending)
	require.Equal(t, tx.Hash(), got.Hash())

	usdt, err := NewIERC20(path1[0], c)
	require.NoError(t, err)
	balance, err := usdt.BalanceOf(nil, liquidatorAddr)
	require.NoError(t, err)
	t.Logf("profit:%v\n", balance)
}

/*
	calculateSeizedTokenAmount case4: seizedSymbol is not stable coin, repaySymbol is stable coin
	height:15511119, account:0x1002C4dB05060e4c1Bac47CeAE3c090984BdE8fC, repaySymbol:vBUSD, repayUnderlyingAmount:29310726539201562054, seizedSymbol:vADA, seizedVTokenAmount:173754577007, seizedUnderlyingAmount:34921065334363953895.9203542686632974, seizedValue:32235671378848660236.1965918233679866, flashLoanReturnAmout:29384187006718357920, remain:2657165985532723821, gasFee:3787050000000000000, profit:-1.1314517423987405
	flashLoanFrom:0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16, path1:[0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56], path2:<nil>, addresses:[0x95c78222B3D6e262426483D42CfA53685A67Ab9D 0x9A0AF7FDb2065Ce470D72664DE73cAE409dA28Ec 0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 0x1002C4dB05060e4c1Bac47CeAE3c090984BdE8fC]
	flashLoanFrom:0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16,
	path1:[0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56],
	path2:<nil>,
	addresses:[0x95c78222B3D6e262426483D42CfA53685A67Ab9D 0x9A0AF7FDb2065Ce470D72664DE73cAE409dA28Ec 0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 0x1002C4dB05060e4c1Bac47CeAE3c090984BdE8fC]

*/
func TestOnlineCase4_2(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	expectedAddr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), from)
	require.NoError(t, err)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	require.NoError(t, err)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = 6000000

	liquidatorAddr, tx, qs, err := DeployUniFlashSwap(auth, c)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, liquidatorAddr)
	t.Logf("hash:%v", tx.Hash())

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, from, owner)

	emptyPath := []common.Address{}
	flashLoanFrom := common.HexToAddress("0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16")

	path1 := []common.Address{
		common.HexToAddress("0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47"),
		common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"),
	}
	path2 := emptyPath
	addresses := []common.Address{
		common.HexToAddress("0x95c78222B3D6e262426483D42CfA53685A67Ab9D"),
		common.HexToAddress("0x9A0AF7FDb2065Ce470D72664DE73cAE409dA28Ec"),
		common.HexToAddress("0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47"),
		common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"),
		common.HexToAddress("0x1002C4dB05060e4c1Bac47CeAE3c090984BdE8fC"),
	}

	repayAmount, ok := big.NewInt(0).SetString("29310726539201562054", 10)
	require.True(t, ok)

	//nonce, err = c.PendingNonceAt(context.Background(), from)
	//require.NoError(t, err)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.GasLimit = 1800000

	liquidator, err := venus.NewIQingsuan(liquidatorAddr, c)
	if err != nil {
		panic(err)
	}
	syncer := Syncer{
		c:          c,
		liquidator: liquidator,
		PrivateKey: privateKey,
	}
	gasLimit := uint64(5000000)
	//tx, err =qs.Qingsuan(auth, big.NewInt(1), flashLoanFrom, emptyPath, emptyPath, addresses, repayAmount)
	tx, err = syncer.doLiquidation(big.NewInt(4), flashLoanFrom, path1, path2, addresses, repayAmount, gasPrice, gasLimit)
	require.NoError(t, err)
	t.Logf("tx:%+v", tx)
	t.Logf("qingsuan hash:%v", tx.Hash())

	got, pending, err := c.TransactionByHash(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.False(t, pending)
	require.Equal(t, tx.Hash(), got.Hash())

	usdt, err := NewIERC20(path1[1], c)
	require.NoError(t, err)
	balance, err := usdt.BalanceOf(nil, liquidatorAddr)
	require.NoError(t, err)
	t.Logf("profit:%v\n", balance)
}

/*
	calculateSeizedTokenAmount case5: seizedSymbol and repaySymbol are not stable coin
	height:15511317, account:0xB7a8c5a31C606dD31fC8523653c2bD097B5d763C, repaySymbol:vADA, repayUnderlyingAmount:26425417566845053839, seizedSymbol:vBNB, seizedVTokenAmount:328235581, seizedUnderlyingAmount:70748857473175031.3347908061483547, seizedValue:26807573395661329179.370420241583839, flashLoanReturnAmout:26491646683553938660, remain:6256832726105499, gasFee:2841838125000000000, profit:-0.4808708721030534
	flashLoanFrom:0x1E249DF2F58cBef7EAc2b0EE35964ED8311D5623, path1:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47], path2:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x55d398326f99059fF775485246999027B3197955], addresses:[0x9A0AF7FDb2065Ce470D72664DE73cAE409dA28Ec 0xA07c5b74C9B40447a954e1466938b865b6BBea36 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47 0xB7a8c5a31C606dD31fC8523653c2bD097B5d763C]
	flashLoanFrom:0x1E249DF2F58cBef7EAc2b0EE35964ED8311D5623,
	path1:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47],
	path2:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x55d398326f99059fF775485246999027B3197955],
	addresses:[0x9A0AF7FDb2065Ce470D72664DE73cAE409dA28Ec 0xA07c5b74C9B40447a954e1466938b865b6BBea36 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47 0xB7a8c5a31C606dD31fC8523653c2bD097B5d763C]

*/

/*
	calculateSeizedTokenAmount case5: seizedSymbol and repaySymbol are not stable coin
	height:15511242, account:0x809851677262344DdDeD8DDD0A2821DC6f9058A4, repaySymbol:vUSDC, repayUnderlyingAmount:28584842397971850382, seizedSymbol:vBNB, seizedVTokenAmount:384968051, seizedUnderlyingAmount:82977133556456360.6581989530932089, seizedValue:31441010885860603415.6293171647156974, flashLoanReturnAmout:28656483606989323663, remain:7124776972181758, gasFee:2841838125000000000, profit:-0.1539612925026267
	flashLoanFrom:0x2354ef4DF11afacb85a5C7f98B624072ECcddbB1, path1:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d], path2:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x55d398326f99059fF775485246999027B3197955], addresses:[0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8 0xA07c5b74C9B40447a954e1466938b865b6BBea36 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d 0x809851677262344DdDeD8DDD0A2821DC6f9058A4]
	flashLoanFrom:0x2354ef4DF11afacb85a5C7f98B624072ECcddbB1,
	path1:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d],
	path2:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x55d398326f99059fF775485246999027B3197955],
	addresses:[0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8 0xA07c5b74C9B40447a954e1466938b865b6BBea36 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d 0x809851677262344DdDeD8DDD0A2821DC6f9058A4]

  Transaction: 0x29cc2c6472c4520bf1f6ad05bbd2eae9a951f2669b11a29578e33358899a7297
  Contract created: 0x55c032aef9d353a5f7562ecba115f9b994147cc1
  Gas usage: 3247721
  Block Number: 15511470
  Block Time: Thu Feb 24 2022 00:08:15 GMT+0800 (中国标准时间)


  Transaction: 0xf67e200e45fcbfd577a5233753bde9ed21c5f5d61063ac3dd41d7d20124f7cda
  Gas usage: 956613
  Block Number: 15511471
  Block Time: Thu Feb 24 2022 00:10:11 GMT+0800 (中国标准时间)

*/

func TestOnlineCase5_2(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	expectedAddr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), from)
	require.NoError(t, err)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	require.NoError(t, err)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = 6000000

	liquidatorAddr, tx, qs, err := DeployUniFlashSwap(auth, c)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, liquidatorAddr)
	t.Logf("hash:%v", tx.Hash())

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, from, owner)

	//emptyPath := []common.Address{}
	flashLoanFrom := common.HexToAddress("0x2354ef4DF11afacb85a5C7f98B624072ECcddbB1")

	path1 := []common.Address{
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"),
	}
	path2 := []common.Address{
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	}
	addresses := []common.Address{
		common.HexToAddress("0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8"),
		common.HexToAddress("0xA07c5b74C9B40447a954e1466938b865b6BBea36"),
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"),
		common.HexToAddress("0x809851677262344DdDeD8DDD0A2821DC6f9058A4"),
	}

	repayAmount, ok := big.NewInt(0).SetString("28584842397971850382", 10)
	require.True(t, ok)

	//nonce, err = c.PendingNonceAt(context.Background(), from)
	//require.NoError(t, err)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.GasLimit = 1800000

	liquidator, err := venus.NewIQingsuan(liquidatorAddr, c)
	if err != nil {
		panic(err)
	}
	syncer := Syncer{
		c:          c,
		liquidator: liquidator,
		PrivateKey: privateKey,
	}
	gasLimit := uint64(5000000)
	//tx, err =qs.Qingsuan(auth, big.NewInt(1), flashLoanFrom, emptyPath, emptyPath, addresses, repayAmount)
	tx, err = syncer.doLiquidation(big.NewInt(5), flashLoanFrom, path1, path2, addresses, repayAmount, gasPrice, gasLimit)
	require.NoError(t, err)
	t.Logf("tx:%+v", tx)
	t.Logf("qingsuan hash:%v", tx.Hash())

	got, pending, err := c.TransactionByHash(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.False(t, pending)
	require.Equal(t, tx.Hash(), got.Hash())

	usdt, err := NewIERC20(path2[1], c)
	require.NoError(t, err)
	balance, err := usdt.BalanceOf(nil, liquidatorAddr)
	require.NoError(t, err)
	t.Logf("profit:%v\n", balance)
}

/*
	calculateSeizedTokenAmount case5: seizedSymbol and repaySymbol are not stable coin
	height:15511152, account:0x2eB71e5335d5328e76fa0755Db27E184Be834D31, repaySymbol:vUSDC, repayUnderlyingAmount:1257831148234280764, seizedSymbol:vCAKE, seizedVTokenAmount:831122592, seizedUnderlyingAmount:194610771405611179.1135772508845391, seizedValue:1382709530836867427.6019663675346503, flashLoanReturnAmout:1260983607252411793, remain:16184184764918663, gasFee:2840287500000000000, profit:-2.7257234558402835
	flashLoanFrom:0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b, path1:[0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d], path2:[0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82 0x55d398326f99059fF775485246999027B3197955], addresses:[0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8 0x86aC3974e2BD0d60825230fa6F355fF11409df5c 0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d 0x2eB71e5335d5328e76fa0755Db27E184Be834D31]

	flashLoanFrom:0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b,
	path1:[0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d],
	path2:[0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82 0x55d398326f99059fF775485246999027B3197955],
	addresses:[0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8 0x86aC3974e2BD0d60825230fa6F355fF11409df5c 0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d 0x2eB71e5335d5328e76fa0755Db27E184Be834D31]

  Transaction: 0x29cc2c6472c4520bf1f6ad05bbd2eae9a951f2669b11a29578e33358899a7297
  Contract created: 0x55c032aef9d353a5f7562ecba115f9b994147cc1
  Gas usage: 3247721
  Block Number: 15511653
  Block Time: Thu Feb 24 2022 00:17:18 GMT+0800 (中国标准时间)

  Transaction: 0xdf327434df10a97d1008b4373b3039b45d333e0fa9d227e76c842a7e7bdf28a2
  Gas usage: 1212872
  Block Number: 15511654
  Block Time: Thu Feb 24 2022 00:19:27 GMT+0800 (中国标准时间)

*/

func TestOnlineCase5_3(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	expectedAddr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), from)
	require.NoError(t, err)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	require.NoError(t, err)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = 6000000

	liquidatorAddr, tx, qs, err := DeployUniFlashSwap(auth, c)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, liquidatorAddr)
	t.Logf("hash:%v", tx.Hash())

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, from, owner)

	//emptyPath := []common.Address{}
	flashLoanFrom := common.HexToAddress("0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b")

	path1 := []common.Address{
		common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"),
		common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"),
	}
	path2 := []common.Address{
		common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"),
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	}
	addresses := []common.Address{
		common.HexToAddress("0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8"),
		common.HexToAddress("0x86aC3974e2BD0d60825230fa6F355fF11409df5c"),
		common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"),
		common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"),
		common.HexToAddress("0x2eB71e5335d5328e76fa0755Db27E184Be834D31"),
	}

	repayAmount, ok := big.NewInt(0).SetString("1257831148234280764", 10)
	require.True(t, ok)

	//nonce, err = c.PendingNonceAt(context.Background(), from)
	//require.NoError(t, err)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.GasLimit = 1800000

	liquidator, err := venus.NewIQingsuan(liquidatorAddr, c)
	if err != nil {
		panic(err)
	}
	syncer := Syncer{
		c:          c,
		liquidator: liquidator,
		PrivateKey: privateKey,
	}
	gasLimit := uint64(5000000)
	//tx, err =qs.Qingsuan(auth, big.NewInt(1), flashLoanFrom, emptyPath, emptyPath, addresses, repayAmount)
	tx, err = syncer.doLiquidation(big.NewInt(5), flashLoanFrom, path1, path2, addresses, repayAmount, gasPrice, gasLimit)
	require.NoError(t, err)
	t.Logf("tx:%+v", tx)
	t.Logf("qingsuan hash:%v", tx.Hash())

	got, pending, err := c.TransactionByHash(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.False(t, pending)
	require.Equal(t, tx.Hash(), got.Hash())

	usdt, err := NewIERC20(path2[1], c)
	require.NoError(t, err)
	balance, err := usdt.BalanceOf(nil, liquidatorAddr)
	require.NoError(t, err)
	t.Logf("profit:%v\n", balance)
}

/*
	calculateSeizedTokenAmount case7: repaySymbol is VAI and seizedSymbol is not stable coin
	height:15511629, account:0x9F7A5885051fB71c4D2a7aB4203446FaCdF65BF7, repaySymbol:VAI, repayUnderlyingAmount:1539941852601276, seizedSymbol:vXVS, seizedVTokenAmount:972986, seizedUnderlyingAmount:195675204479828.4167110143560824, seizedValue:1691286637025464.467895520335951, flashLoanReturnAmout:1543801355991255, remain:135870836858802, gasFee:2832326250000000000, profit:-2.8311387234162779
	flashLoanFrom:0x133ee93FE93320e1182923E1a640912eDE17C90C, path1:[0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63 0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7], path2:[0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63 0x55d398326f99059fF775485246999027B3197955], addresses:[0x004065D34C6b18cE4370ced1CeBDE94865DbFAFE 0x151B1e2635A717bcDc836ECd6FbB62B674FE3E1D 0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63 0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7 0x9F7A5885051fB71c4D2a7aB4203446FaCdF65BF7]

	flashLoanFrom:0x133ee93FE93320e1182923E1a640912eDE17C90C,
	path1:[0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63 0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7],
	path2:[0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63 0x55d398326f99059fF775485246999027B3197955],
	addresses:[0x004065D34C6b18cE4370ced1CeBDE94865DbFAFE 0x151B1e2635A717bcDc836ECd6FbB62B674FE3E1D 0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63 0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7 0x9F7A5885051fB71c4D2a7aB4203446FaCdF65BF7]

  Transaction: 0x29cc2c6472c4520bf1f6ad05bbd2eae9a951f2669b11a29578e33358899a7297
  Contract created: 0x55c032aef9d353a5f7562ecba115f9b994147cc1
  Gas usage: 3247721
  Block Number: 15511801
  Block Time: Thu Feb 24 2022 00:24:47 GMT+0800 (中国标准时间)


  Transaction: 0x75ef8992d38066b8428b42220e6a88aa2121b619fbfde52782de67a1a8c8c496
  Gas usage: 2090456
  Block Number: 15511802
*/

func TestOnlineCase7_2(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	expectedAddr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), from)
	require.NoError(t, err)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	require.NoError(t, err)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = 6000000

	liquidatorAddr, tx, qs, err := DeployUniFlashSwap(auth, c)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, liquidatorAddr)
	t.Logf("hash:%v", tx.Hash())

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, from, owner)

	//emptyPath := []common.Address{}
	flashLoanFrom := common.HexToAddress("0x133ee93FE93320e1182923E1a640912eDE17C90C")

	path1 := []common.Address{
		common.HexToAddress("0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63"),
		common.HexToAddress("0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7"),
	}
	path2 := []common.Address{
		common.HexToAddress("0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63"),
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	}
	addresses := []common.Address{
		common.HexToAddress("0x004065D34C6b18cE4370ced1CeBDE94865DbFAFE"),
		common.HexToAddress("0x151B1e2635A717bcDc836ECd6FbB62B674FE3E1D"),
		common.HexToAddress("0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63"),
		common.HexToAddress("0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7"),
		common.HexToAddress("0x9F7A5885051fB71c4D2a7aB4203446FaCdF65BF7"),
	}

	repayAmount, ok := big.NewInt(0).SetString("1539941852601276", 10)
	require.True(t, ok)

	//nonce, err = c.PendingNonceAt(context.Background(), from)
	//require.NoError(t, err)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.GasLimit = 1800000

	liquidator, err := venus.NewIQingsuan(liquidatorAddr, c)
	if err != nil {
		panic(err)
	}
	syncer := Syncer{
		c:          c,
		liquidator: liquidator,
		PrivateKey: privateKey,
	}
	gasLimit := uint64(5000000)
	//tx, err =qs.Qingsuan(auth, big.NewInt(1), flashLoanFrom, emptyPath, emptyPath, addresses, repayAmount)
	tx, err = syncer.doLiquidation(big.NewInt(7), flashLoanFrom, path1, path2, addresses, repayAmount, gasPrice, gasLimit)
	require.NoError(t, err)
	t.Logf("tx:%+v", tx)
	t.Logf("qingsuan hash:%v", tx.Hash())

	got, pending, err := c.TransactionByHash(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.False(t, pending)
	require.Equal(t, tx.Hash(), got.Hash())

	usdt, err := NewIERC20(path2[1], c)
	require.NoError(t, err)
	balance, err := usdt.BalanceOf(nil, liquidatorAddr)
	require.NoError(t, err)
	t.Logf("profit:%v\n", balance)
}

func TestOnlineCase4_2_1(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	expectedAddr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), from)
	require.NoError(t, err)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	require.NoError(t, err)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = 6000000

	liquidatorAddr, tx, qs, err := DeployUniFlashSwap(auth, c)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, liquidatorAddr)
	t.Logf("hash:%v", tx.Hash())

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, from, owner)

	/*
		calculateSeizedTokenAmount case5: seizedSymbol and repaySymbol are not stable coin
		height:15651152, account:0x19d598Fa8846c46A1D94D7F4024Ed49F357C7181, repaySymbol:vBCH, repayUnderlyingAmount:30169989390770485, seizedSymbol:vBTC, seizedVTokenAmount:1340519, seizedUnderlyingAmount:270869908674815.8770540548427115, seizedValue:10402306489908816815.7462959627478293, flashLoanReturnAmout:30245603399268657, remain:25626196971772, gasFee:9199635732250000000, profit:-8.2162810108999601
		flashLoanFrom:0xAfB3c543EBa8aFcb87b5d552C1142d9a18D375e7,
		path1:[0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x8fF795a6F4D97E7887C79beA79aba5cc76444aDf],
		path2:[0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c 0x55d398326f99059fF775485246999027B3197955],
		addresses:[0x5F0388EBc2B94FA8E123F404b79cCF5f40b29176 0x882C173bC7Ff3b7786CA16dfeD3DFFfb9Ee7847B 0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c 0x8fF795a6F4D97E7887C79beA79aba5cc76444aDf 0x19d598Fa8846c46A1D94D7F4024Ed49F357C7181]
	*/

	//emptyPath := []common.Address{}
	flashLoanFrom := common.HexToAddress("0xAfB3c543EBa8aFcb87b5d552C1142d9a18D375e7")

	path1 := []common.Address{
		common.HexToAddress("0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c"),
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		common.HexToAddress("0x8fF795a6F4D97E7887C79beA79aba5cc76444aDf"),
	}
	path2 := []common.Address{
		common.HexToAddress("0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c"),
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	}
	addresses := []common.Address{
		common.HexToAddress("0x5F0388EBc2B94FA8E123F404b79cCF5f40b29176"),
		common.HexToAddress("0x882C173bC7Ff3b7786CA16dfeD3DFFfb9Ee7847B"),
		common.HexToAddress("0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c"),
		common.HexToAddress("0x8fF795a6F4D97E7887C79beA79aba5cc76444aDf"),
		common.HexToAddress("0x19d598Fa8846c46A1D94D7F4024Ed49F357C7181"),
	}

	repayAmount, ok := big.NewInt(0).SetString("30169989390770485", 10)
	require.True(t, ok)

	//nonce, err = c.PendingNonceAt(context.Background(), from)
	//require.NoError(t, err)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.GasLimit = 1800000

	liquidator, err := venus.NewIQingsuan(liquidatorAddr, c)
	if err != nil {
		panic(err)
	}
	syncer := Syncer{
		c:          c,
		liquidator: liquidator,
		PrivateKey: privateKey,
	}
	gasLimit := uint64(5000000)
	//tx, err =qs.Qingsuan(auth, big.NewInt(1), flashLoanFrom, emptyPath, emptyPath, addresses, repayAmount)
	tx, err = syncer.doLiquidation(big.NewInt(5), flashLoanFrom, path1, path2, addresses, repayAmount, gasPrice, gasLimit)
	require.NoError(t, err)
	t.Logf("tx:%+v", tx)
	t.Logf("qingsuan hash:%v", tx.Hash())

	got, pending, err := c.TransactionByHash(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.False(t, pending)
	require.Equal(t, tx.Hash(), got.Hash())

	usdt, err := NewIERC20(path2[1], c)
	require.NoError(t, err)
	balance, err := usdt.BalanceOf(nil, liquidatorAddr)
	require.NoError(t, err)
	t.Logf("profit:%v\n", balance)
}

/**/

func TestOnlineCase5_2_1(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	expectedAddr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), from)
	require.NoError(t, err)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	require.NoError(t, err)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = 6000000

	liquidatorAddr, tx, qs, err := DeployUniFlashSwap(auth, c)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, liquidatorAddr)
	t.Logf("hash:%v", tx.Hash())

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, from, owner)

	/*
		calculateSeizedTokenAmount case5: seizedSymbol and repaySymbol are not stable coin
		height:15653630, account:0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627, repaySymbol:vSXP, repayUnderlyingAmount:15106708392873564962, seizedSymbol:vBNB, seizedVTokenAmount:275212714, seizedUnderlyingAmount:59324530524536340.704315371372273, seizedValue:22362100082250249699.1500252394962684, flashLoanReturnAmout:15144569817417107717, remain:4947454212868305, gasFee:5654178773400000000, profit:-3.7974079246868654
		flashLoanFrom:0x55Dfbc7C21678Ee9eD2d0f1bFEe391263d807719,
		path1:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x47BEAd2563dCBf3bF2c9407fEa4dC236fAbA485A],
		path2:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x55d398326f99059fF775485246999027B3197955],
		addresses:[0x2fF3d0F6990a40261c66E1ff2017aCBc282EB6d0 0xA07c5b74C9B40447a954e1466938b865b6BBea36 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x47BEAd2563dCBf3bF2c9407fEa4dC236fAbA485A 0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627]
	*/

	//emptyPath := []common.Address{}
	flashLoanFrom := common.HexToAddress("0x55Dfbc7C21678Ee9eD2d0f1bFEe391263d807719")

	path1 := []common.Address{
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		common.HexToAddress("0x47BEAd2563dCBf3bF2c9407fEa4dC236fAbA485A"),
	}
	path2 := []common.Address{
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	}
	addresses := []common.Address{
		common.HexToAddress("0x2fF3d0F6990a40261c66E1ff2017aCBc282EB6d0"),
		common.HexToAddress("0xA07c5b74C9B40447a954e1466938b865b6BBea36"),
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		common.HexToAddress("0x47BEAd2563dCBf3bF2c9407fEa4dC236fAbA485A"),
		common.HexToAddress("0x0A88bbE6be0005E46F56aA4145c8FB863f9Df627"),
	}

	repayAmount, ok := big.NewInt(0).SetString("15106708392873564962", 10)
	require.True(t, ok)

	//nonce, err = c.PendingNonceAt(context.Background(), from)
	//require.NoError(t, err)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.GasLimit = 1800000

	liquidator, err := venus.NewIQingsuan(liquidatorAddr, c)
	if err != nil {
		panic(err)
	}
	syncer := Syncer{
		c:          c,
		liquidator: liquidator,
		PrivateKey: privateKey,
	}
	gasLimit := uint64(5000000)
	//tx, err =qs.Qingsuan(auth, big.NewInt(1), flashLoanFrom, emptyPath, emptyPath, addresses, repayAmount)
	tx, err = syncer.doLiquidation(big.NewInt(5), flashLoanFrom, path1, path2, addresses, repayAmount, gasPrice, gasLimit)
	require.NoError(t, err)
	t.Logf("tx:%+v", tx)
	t.Logf("qingsuan hash:%v", tx.Hash())

	got, pending, err := c.TransactionByHash(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.False(t, pending)
	require.Equal(t, tx.Hash(), got.Hash())

	usdt, err := NewIERC20(path2[1], c)
	require.NoError(t, err)
	balance, err := usdt.BalanceOf(nil, liquidatorAddr)
	require.NoError(t, err)
	t.Logf("profit:%v\n", balance)
}

/*
0xb14e9a72000000000000000000000000000000000000000000000000000000000000000400000000000000000000000058f876857a02d6762e0101bb5c46a8c1ed44dc1600000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000005938b81ab1dd45b52600000000000000000000000000000000000000000000000000000000000000020000000000000000000000002170ed0880ac9a755fd29b2688956bd959f933f8000000000000000000000000e9e7cea3dedca5984780bafc599bd69add087d560000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000095c78222b3d6e262426483d42cfa53685a67ab9d000000000000000000000000f508fcd89b8bd15579dc79a6827cb4686a3592c80000000000000000000000002170ed0880ac9a755fd29b2688956bd959f933f8000000000000000000000000e9e7cea3dedca5984780bafc599bd69add087d56000000000000000000000000fcd064b4228c4491a642072b036096bd999e52d9
*/

func TestUnpackInputData(t *testing.T) {
	data, err := hex.DecodeString("b14e9a72000000000000000000000000000000000000000000000000000000000000000400000000000000000000000058f876857a02d6762e0101bb5c46a8c1ed44dc1600000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000005938b81ab1dd45b52600000000000000000000000000000000000000000000000000000000000000020000000000000000000000002170ed0880ac9a755fd29b2688956bd959f933f8000000000000000000000000e9e7cea3dedca5984780bafc599bd69add087d560000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000095c78222b3d6e262426483d42cfa53685a67ab9d000000000000000000000000f508fcd89b8bd15579dc79a6827cb4686a3592c80000000000000000000000002170ed0880ac9a755fd29b2688956bd959f933f8000000000000000000000000e9e7cea3dedca5984780bafc599bd69add087d56000000000000000000000000fcd064b4228c4491a642072b036096bd999e52d9")
	require.NoError(t, err)

	qsAbi, err := abi.JSON(strings.NewReader(venus.IQingsuanMetaData.ABI))
	require.NoError(t, err)

	method, err := qsAbi.MethodById(data[0:4])
	require.NoError(t, err)
	t.Logf("method %v\n", method)

	out := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(out, data[4:])

	t.Logf("output:%v", out)

	result, err := method.Inputs.Unpack(data[4:])
	require.NoError(t, err)

	t.Logf("result:%+v", result)
}
