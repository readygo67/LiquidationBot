package syncer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"github.com/readygo67/LiquidationBot/venus"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"
)

func TestScaningBlockWithManualParse(t *testing.T) {
	ctx := context.Background()
	rpcURL := "http://42.3.146.198:21993" //"https://bsc-dataseed1.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	height := big.NewInt(14747566)

	blk, err := c.BlockByNumber(ctx, height)
	require.NoError(t, err)

	vUSDCAddress := common.HexToAddress("0xeca88125a5adbe82614ffc12d0db554e2e2867c8")
	topicBorrow := "0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80"
	var borrowers []common.Address

	for _, tx := range blk.Transactions() {
		if tx.To() == nil {
			continue
		}
		if tx.To().String() == vUSDCAddress.String() {
			hash := tx.Hash()
			receipt, err := c.TransactionReceipt(context.TODO(), hash)
			require.NoError(t, err)

			for _, receiptlog := range receipt.Logs {
				if receiptlog.Removed {
					continue
				}

				if receiptlog.Address.String() == vUSDCAddress.String() {
					if receiptlog.Topics[0].String() == topicBorrow {
						t.Logf("receipt:%+v", receiptlog)
						t.Logf("borrower:%+v", common.BytesToAddress(receiptlog.Data[12:32]))

						borrowers = append(borrowers, common.BytesToAddress(receiptlog.Data[12:32]))
					}

				}
			}
		}
	}
	t.Logf("borrowers:%v", borrowers)
}

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

func TestFilterAllCotractsBorrowEvent(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New("../config.yml")
	rpcURL := "http://42.3.146.198:21993"
	c, err := ethclient.Dial(rpcURL)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	sync := NewSyncer(c, nil, cfg)

	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")
	var addresses []common.Address
	name := make(map[string]string)

	//build all concerned address
	for _, token := range sync.tokens {
		addresses = append(addresses, common.HexToAddress(token.Address))
		name[strings.ToLower(token.Address)] = token.Name
		//fmt.Printf("name:%v, address:%v\n", name[token.Address], token.Address)
	}

	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	require.NoError(t, err)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(13019766),
		ToBlock:   big.NewInt(13019768),
		Addresses: addresses,
		Topics:    [][]common.Hash{{topicBorrow}},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	require.NoError(t, err)
	fmt.Printf("start Time:%v\n", time.Now())
	for i, log := range logs {
		var borrowEvent venus.Vbep20Borrow
		err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", log.Data)
		//fmt.Printf("logAddress:%v\n", log.Address.String())
		fmt.Printf("%v height:%v, name:%v borrower:%v\n", (i + 1), log.BlockNumber, name[strings.ToLower(log.Address.String())], borrowEvent.Borrower)

		//jsonLog, err := json.Marshal(log)
		//require.NoError(t, err)
		//fmt.Printf("%v log: %s\n", i, jsonLog)
	}
	fmt.Printf("end Time:%v\n", time.Now())
}

//Filter with ethereum.SubscribeFilterLogs,
func TestSubscribeBorrowEvent(t *testing.T) {
	t.Skip("this case need websocket support")
	ctx := context.Background()
	rpcURL := "wss://bsc-ws-node.nariox.org:443" //"ws://42.3.146.198:21994" //"https://bsc-dataseed1.binance.org"
	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	_, err = c.BlockNumber(ctx)
	require.NoError(t, err)

	vbep20Abi, err := abi.JSON(strings.NewReader(venus.Vbep20MetaData.ABI))
	require.NoError(t, err)

	vUSDCAddress := common.HexToAddress("0xeca88125a5adbe82614ffc12d0db554e2e2867c8")
	topicBorrow := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(14747566),
		ToBlock:   big.NewInt(14747568),
		Addresses: []common.Address{vUSDCAddress},
		Topics:    [][]common.Hash{{topicBorrow}},
	}

	logs := make(chan types.Log, 10)
	sub, err := c.SubscribeFilterLogs(context.Background(), query, logs)
	require.NoError(t, err)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			fmt.Println(vLog.BlockNumber)
			var borrowEvent venus.Vbep20Borrow
			err = vbep20Abi.UnpackIntoInterface(&borrowEvent, "Borrow", vLog.Data)
			require.NoError(t, err)
		}
	}
}

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
