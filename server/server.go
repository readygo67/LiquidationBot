package server

import (
"context"
"encoding/hex"
"encoding/json"
"fmt"
"github.com/ethereum/go-ethereum"
"github.com/ethereum/go-ethereum/common"
"github.com/ethereum/go-ethereum/core/types"
"github.com/ethereum/go-ethereum/ethclient"
"github.com/stretchr/testify/require"
"math/big"
"strings"
"testing"
)

//ERC20 standard functions signature
/*
   "allowance(address,address)": "dd62ed3e",
   "approve(address,uint256)": "095ea7b3",
   "balanceOf(address)": "70a08231",
   "decimals()": "313ce567",
   "name()": "06fdde03",
   "symbol()": "95d89b41",
   "totalSupply()": "18160ddd",
   "transfer(address,uint256)": "a9059cbb",
   "transferFrom(address,address,uint256)": "23b872dd"
*/

//ERC20Ref functions signature
/*
   "queryRefer(address)": "3fdcd52c",
   "registerReferFor(address,address)": "f4ac13cd"
*/

func TestGetCode(t *testing.T){
	rpcURL := "https://data-seed-prebsc-1-s1.binance.org:8545"
	usdtContractAddr := common.HexToAddress("0xae13d989dac2f0debff460ac112a837c89baa7cd")
	pearToken2ContractAddr := common.HexToAddress("0xD1eE49394Af5e511E568ce500110966FD89608E8")
	ERC20FunctionSignatures := []string{
		"dd62ed3e",
		"095ea7b3",
		"70a08231",
		"313ce567",
		"06fdde03",
		"95d89b41",
		"18160ddd",
		"a9059cbb",
		"23b872dd",
	}
	ERC20RefFunctionSignatures := []string{
		"3fdcd52c",
		"f4ac13cd",
	}

	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err);
	chainID, err := client.ChainID(context.Background());
	fmt.Printf("chainid:%v", chainID);

	//test usdt
	byteCode, err := client.CodeAt(context.Background(),usdtContractAddr, nil)
	code := hex.EncodeToString(byteCode)

	for _, signature := range ERC20FunctionSignatures{
		require.True(t, strings.Contains(code, signature))
	}

	for _, signature := range ERC20RefFunctionSignatures{
		require.False(t, strings.Contains(code, signature))
	}

	//test pear Token
	byteCode, err = client.CodeAt(context.Background(),pearToken2ContractAddr, nil)
	code = hex.EncodeToString(byteCode)

	for _, signature := range ERC20FunctionSignatures{
		require.True(t, strings.Contains(code, signature))
	}

	for _, signature := range ERC20RefFunctionSignatures{
		require.True(t, strings.Contains(code, signature))
	}
}

func TestSubscribeNewBlock(t *testing.T){
	//rpcURL :=  "https://data-seed-prebsc-1-s1.binance.org:8545"
	//rpcURL :=  "wss://data-seed-prebsc-1-s1.binance.org:8546"
	//rpcURL := "wss://bsc-ws-node.nariox.org:443"
	rpcURL := "wss://bsc.getblock.io/mainnet/"
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)



	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	require.NoError(t, err)

	for i :=0; i<=50; i++{
		select {
		case err := <-sub.Err():
			require.NoError(t, err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			require.NoError(t, err)

			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		}
	}
}


func TestSubscribeEvents(t *testing.T){
	//rpcURL :=  "https://data-seed-prebsc-1-s1.binance.org:8545"
	//rpcURL :=  "wss://data-seed-prebsc-1-s1.binance.org:8546"
	rpcURL := "wss://bsc-ws-node.nariox.org:443"
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	usdtContractAddr := common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	transferTopic := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	fromAddr := common.HexToHash("0x00000000000000000000000010ed43c718714eb63d5aa57b78b54704e256024e")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(8745930),
		Addresses: []common.Address{usdtContractAddr},
		Topics: [][]common.Hash{{transferTopic}, {fromAddr}},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query,  logs)
	require.NoError(t, err)

	for i :=0; i<=5; i++{
		select {
		case err := <-sub.Err():
			require.NoError(t, err)
		case vLog := <-logs:
			jsonLog, err := json.Marshal(vLog)
			require.NoError(t, err)
			fmt.Printf("%vlog: %s\n", i, jsonLog)
		}
	}
}


func TestSubscribeEvents1(t *testing.T){
	//rpcURL :=  "https://data-seed-prebsc-1-s1.binance.org:8545"
	//rpcURL :=  "wss://data-seed-prebsc-1-s1.binance.org:8546"
	rpcURL := "wss://bsc-ws-node.nariox.org:443"
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	contractAddr := common.HexToAddress("0xca143ce32fe78f1f7019d7d551a6402fc5350c73")
	topic := common.HexToHash("0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9")
	//fromAddr := common.HexToHash("0x00000000000000000000000010ed43c718714eb63d5aa57b78b54704e256024e")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(8746437),
		ToBlock:   big.NewInt(8746842),
		Addresses: []common.Address{contractAddr},
		Topics: [][]common.Hash{{topic}},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query,  logs)
	require.NoError(t, err)

	for i :=0; i<=50; i++{
		select {
		case err := <-sub.Err():
			require.NoError(t, err)
		case vLog := <-logs:
			jsonLog, err := json.Marshal(vLog)
			require.NoError(t, err)
			fmt.Printf("%vlog: %s\n", i, jsonLog)
		}
	}
}


func TestFiltereEvents(t *testing.T){
	rpcURL :=  "https://data-seed-prebsc-1-s1.binance.org:8545"
	//rpcURL :=  "wss://data-seed-prebsc-1-s1.binance.org:8546"
	//rpcURL := "wss://bsc-ws-node.nariox.org:443"
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	lemonContractAddr := common.HexToAddress("0x4197CB336f7a13c631688ECD56bEf1A83aE8343a")
	transferTopic := common.HexToHash("0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9")

	paddingFromAddr := common.BytesToHash(common.LeftPadBytes(common.HexToAddress("0xc96d141c9110a8E61eD62caaD8A7c858dB15B82c").Bytes(), 32))

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(10094076),
		ToBlock:   big.NewInt(10140000),
		Addresses: []common.Address{lemonContractAddr},
		Topics: [][]common.Hash{{transferTopic}, {paddingFromAddr}},
	}


	logs, err := client.FilterLogs(context.Background(), query)
	require.NoError(t, err)

	for i, log := range logs {
		jsonLog, err := json.Marshal(log)
		require.NoError(t, err)
		fmt.Printf("%vlog: %s\n", i, jsonLog)
	}
}

func TestFiltereEvent1s(t *testing.T){
	rpcURL :=  "https://bsc-dataseed.binance.org"
	//rpcURL :=  "wss://data-seed-prebsc-1-s1.binance.org:8546"
	//rpcURL := "wss://bsc-ws-node.nariox.org:443"
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	contractAddr := common.HexToAddress("0xca143ce32fe78f1f7019d7d551a6402fc5350c73")
	topic := common.HexToHash("0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9")

	//paddingFromAddr := common.BytesToHash(common.LeftPadBytes(common.HexToAddress("0xc96d141c9110a8E61eD62caaD8A7c858dB15B82c").Bytes(), 32))

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(8746437),
		ToBlock:   big.NewInt(8746842),
		Addresses: []common.Address{contractAddr},
		Topics: [][]common.Hash{{topic},},
	}


	logs, err := client.FilterLogs(context.Background(), query)
	require.NoError(t, err)

	for i, log := range logs {
		jsonLog, err := json.Marshal(log)
		require.NoError(t, err)
		fmt.Printf("%vlog: %s\n", i, jsonLog)
	}
}