package model

const (
	KVTable = "kv"
)

const (
	KeyLastHandledHeight        = "LastHandledHeight"
	KeyTimeIntervalBelow1P1     = "KeyTimeIntervalBelow1P1"
	KeyTimeIntervalBelow1P2     = "KeyTimeIntervalBelow1P2"
	KeyTimeIntervalBelow1P5     = "KeyTimeIntervalBelow1P5"
	KeyTimeIntervalBelow2       = "KeyTimeIntervalBelow2"
	KeyTimeIntervalBelow1Above2 = "KeyTimeIntervalBelow2"
)

type KV struct {
	Key   []byte `gorm:"primary_key"`
	Value []byte
}

func (KV) TableName() string {
	return KVTable
}
