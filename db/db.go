package dbm

import (
	leveldb "github.com/syndtr/goleveldb/leveldb"
)

const HashLength = 32

// Hash to identify uniqueness
type Hash [HashLength]byte

var (
	KeyLastHandledHeight      = []byte("last_handled_height")
	KeyBorrowerNumber         = []byte("number_of_borrowers")
	BorrowersPrefix           = []byte("borrowers") //prefix with all borrowers
	PricesPrefix              = []byte("prices")
	AccountPrefix             = []byte("account")
	MarketPrefix              = []byte("market")
	LiquidationBelow1P0Prefix = []byte("liquidation_below_1p0")
	LiquidationBelow1P2Prefix = []byte("liquidation_below_1p2")
	LiquidationBelow1P5Prefix = []byte("liquidation_below_1p5")
	LiquidationBelow2P0Prefix = []byte("liquidation_below_2p0")
	LiquidationBelow3P0Prefix = []byte("liquidation_below_3p0")
	LiquidationAbove3P0Prefix = []byte("liquidation_above_3p0")
)

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

func LiquidationBelow1P2StoreKey(address []byte) []byte {
	return append(LiquidationBelow1P2Prefix, address...)
}

func LiquidationBelow1P5StoreKey(address []byte) []byte {
	return append(LiquidationBelow1P5Prefix, address...)
}

func LiquidationBelow2P0StoreKey(address []byte) []byte {
	return append(LiquidationBelow2P0Prefix, address...)
}

func LiquidationBelow3P0StoreKey(address []byte) []byte {
	return append(LiquidationBelow3P0Prefix, address...)
}

func LiquidationAbove3P0StoreKey(address []byte) []byte {
	return append(LiquidationAbove3P0Prefix, address...)
}

func NewDB(path string) (*leveldb.DB, error) {
	return leveldb.OpenFile(path, nil)
}
