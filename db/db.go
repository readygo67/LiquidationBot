package dbm

import (
	leveldb "github.com/syndtr/goleveldb/leveldb"
)

const HashLength = 32

// Hash to identify uniqueness
type Hash [HashLength]byte

var (
	KeyLastHandledHeight       = []byte("last_handled_height")
	KeyBorrowerNumber          = []byte("number_of_borrowers")
	BorrowersPrefix            = []byte("borrowers") //prefix with all borrowers
	PricesPrefix               = []byte("prices")
	AccountPrefix              = []byte("account")
	MarketPrefix               = []byte("market")
	LiquidationBelow1P0Prefix  = []byte("liquidation_below_1p0")
	LiquidationBelow1P1Prefix  = []byte("liquidation_below_1p1")
	LiquidationBelow1P5Prefix  = []byte("liquidation_below_1p5")
	LiquidationBelow2P0Prefix  = []byte("liquidation_below_2p0")
	LiquidationAbove2P0Prefix  = []byte("liquidation_above_2p0")
	LiquidationNonProfitPrefix = []byte("liquidation_non_profit") //
)

func BorrowerNumberKey() []byte {
	return KeyBorrowerNumber
}

func LastHandledHeightStoreKey() []byte {
	return KeyLastHandledHeight
}

func BorrowersStoreKey(address []byte) []byte {
	return append(BorrowersPrefix, address...)
}

func MarketStoreKey(symbol []byte, account []byte) []byte {
	bz := append(MarketPrefix, symbol...)
	return append(bz, account...)
}

func AccountStoreKey(account []byte) []byte {
	return append(AccountPrefix, account...)
}

func LiquidationBelow1P0StoreKey(address []byte) []byte {
	return append(LiquidationBelow1P0Prefix, address...)
}

func LiquidationBelow1P1StoreKey(address []byte) []byte {
	return append(LiquidationBelow1P1Prefix, address...)
}

func LiquidationBelow1P5StoreKey(address []byte) []byte {
	return append(LiquidationBelow1P5Prefix, address...)
}

func LiquidationBelow2P0StoreKey(address []byte) []byte {
	return append(LiquidationBelow2P0Prefix, address...)
}

func LiquidationAbove2P0StoreKey(address []byte) []byte {
	return append(LiquidationAbove2P0Prefix, address...)
}

func LiquidationNonProfitStoreKey(address []byte) []byte {
	return append(LiquidationNonProfitPrefix, address...)
}

func NewDB(path string) (*leveldb.DB, error) {
	return leveldb.OpenFile(path, nil)
}
