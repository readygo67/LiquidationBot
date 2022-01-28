package dao

import "github.com/readygo67/LiquidationBot/dao/model"

type Account struct{}

func (*Account) CreateOrUpdate(key []byte, value []byte) error {
	kv := model.KV{key, value}
	return DB.Table(model.KVTable).Save(kv).Error
}

func (*Account) GetByKey(key []byte) ([]byte, error) {
	var kv model.KV
	err := DB.Table(model.KVTable).Where("key = ?", key).First(&kv).Error
	return kv.Value, err
}
