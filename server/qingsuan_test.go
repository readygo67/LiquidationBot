package server

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

/*

Available Accounts
==================
Available Accounts
==================
(0) 0x4A36F13E42A82b349b0C399F64a5D91F9Ff195Ad (100 ETH)
(1) 0x4B83df201898f693AD801998466730D68880EE69 (100 ETH)
(2) 0x2DFb31Bce57D253AAccf231030BcA212721D13Fd (100 ETH)
(3) 0xE223a419FC59D490Ac65EB964D45F6081950C6b5 (100 ETH)
(4) 0x109ce855Ac2bD5dF5c317f43bF9e63B38c796Ea3 (100 ETH)
(5) 0x982e51DDE955cD7036888ed3EF560186118A1F28 (100 ETH)
(6) 0x51FB7F783EA24f7042531E1312D71C5090559682 (100 ETH)
(7) 0x28f0935060fd6fb8063D83ADF2E70Bce7af26e99 (100 ETH)
(8) 0xf1bAE3E6002B2704d5C3d9FFF3B161cbaeA56790 (100 ETH)
(9) 0x13CA41e5ea60D1eea6ce974Dc8F6c48529c3b271 (100 ETH)

Private Keys
==================
(0) 0x5878d8f0abb10f30a1b2c2d6d6ee01f029ea3e3ff8e0213bb5c91f43107ffcaa
(1) 0x94fa1536beb6e5dacc43975749ac8c9c20fe6175135d737a9c2dd8ec4cbfdaaa
(2) 0x8a796eba59059de37b4627ab7db78e6cca23db077b7a2c710afcdf6d3d4472ba
(3) 0x4ecf637344196c144f2af955e150260ca0fcda6d2fb6ab59d36e4bbdd2a98105
(4) 0xc90557afa51abb61cf31fa620f7cfa4ed92fa9e1e19a7098a200c827f5ab3227
(5) 0xbb0bb5502943315491345fbdc29fbc326eaf7a680a7d35c835c1b8cd34ba4001
(6) 0x06dd100bccc5bcbe9b93ee3b45eb269b399110fc7a23d261200984eea6b070e2
(7) 0xc69d61cca327ff5a59534fe14f92a118495265ae5485411bf141a10a9955003d
(8) 0x4686f571c7da529835dff1d88d7d225e28000757e6536bb5838bf902f1802b5a
(9) 0x204124d2c6859b6b4c83eab981ef17996ea3a21bd88ab5e6f7161bc19b15a496
*/
/*
height15325212, account:0x9Dbb3fa7F18C72FDF5fa5D1eC630B0F1F191FAE6, repaySmbol:vBUSD, flashLoanFrom:0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16, repayAddress:0x95c78222B3D6e262426483D42CfA53685A67Ab9D, repayValue:40015902181827895653.4091979725, repayAmount:40010220770488707027 seizedSymbol:vBUSD, seizedAddress:0x95c78222B3D6e262426483D42CfA53685A67Ab9D, seizedCTokenAmount:205825257122, seizedUnderlyingTokenAmount:44011242845431854992.1699406601840677, seizedUnderlyingTokenValue:44017492397904663470.1469737995878732
calculateSeizedTokenAmount case1: seizedSymbol == repaySymbol and symbol is a stable coin, account:0x9Dbb3fa7F18C72FDF5fa5D1eC630B0F1F191FAE6, symbol:vBUSD, seizedAmount:44011242845431854992.1699406601840677, returnAmout:40110497013021260138, gasFee:953555625000000000 profit:2.9477441094180513
case1, profitable liquidation catched:&{0x9Dbb3fa7F18C72FDF5fa5D1eC630B0F1F191FAE6 0.9998935931903574 15325033 0001-01-01 00:00:00 +0000 UTC}, profit:2.9477441094180513
*/
func TestOnlineCase1(t *testing.T) {
	rpcURL := "http://127.0.0.1:8545"
	addr := common.HexToAddress("0xED682ae11062F85f62D73274D7E0c656eA4825b4")
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	qs, err := NewQingsuan(addr, c)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("f461544366be72b82515b658b6a00e8fb7ed506ee3945a62349f7b659c784b2e")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	require.NoError(t, err)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(56))
	auth.Value = big.NewInt(0) // in wei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice, err = c.SuggestGasPrice(context.Background())
	require.NoError(t, err)

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

	tx, err := qs.Qingsuan(auth, big.NewInt(1), common.HexToAddress("0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16"), emptyPath, emptyPath, addresses, repayAmount)
	require.NoError(t, err)
	t.Logf("hash:%v", tx.Hash())
}
