package server

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"log"
	"math/big"
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

func TestOnlineCase1(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	addr := common.HexToAddress("0x55C032aEf9D353a5f7562ecba115f9B994147Cc1")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	qs, err := NewQingsuan(addr, c)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("5c694ea694c13fbe230763060a16c6b5c039f8a08f2ae606e5447866f593d384")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)

	value, ok := big.NewInt(0).SetString("20000000000000000", 10)
	require.True(t, ok)
	gasPrice, err := c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, addr, value, 300000, gasPrice, nil)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = c.SendTransaction(context.Background(), signedTx)
	require.NoError(t, err)
	t.Logf("send 1 ether, hash:%v", signedTx.Hash())

	balance, err := c.BalanceAt(context.Background(), addr, nil)
	require.NoError(t, err)
	t.Logf("balance:%v\n", balance)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(56))
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce + 1))
	auth.GasPrice = gasPrice

	addresses := []common.Address{
		common.HexToAddress("0x95c78222B3D6e262426483D42CfA53685A67Ab9D"),
		common.HexToAddress("0x95c78222B3D6e262426483D42CfA53685A67Ab9D"),
		common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"),
		common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"),
		common.HexToAddress("0x9Dbb3fa7F18C72FDF5fa5D1eC630B0F1F191FAE6"),
	}

	emptyPath := []common.Address{}

	repayAmount, ok := big.NewInt(0).SetString("40011357893447481535", 10)
	require.True(t, ok)

	owner, err := qs.Owner(nil)
	require.NoError(t, err)
	t.Logf("owner:%v", owner)

	tx, err = qs.Qingsuan(auth, big.NewInt(1), common.HexToAddress("0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16"), emptyPath, emptyPath, addresses, repayAmount)
	require.NoError(t, err)
	t.Logf("qingsuan hash:%v", tx.Hash())
}
