package syncer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
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

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(14747566),
		ToBlock:   big.NewInt(14747568),
		Addresses: []common.Address{vUSDCAddress},
		Topics:    [][]common.Hash{{topicBorrow}},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	require.NoError(t, err)

	for i, log := range logs {
		jsonLog, err := json.Marshal(log)
		require.NoError(t, err)
		fmt.Printf("%v log: %s\n", i, jsonLog)
	}
}
