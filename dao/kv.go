package dao

import "github.com/readygo67/LiquidationBot/dao/model"

type KV struct{}

func (*KV) CreateOrUpdate(a *model.Account) error {
	return DB.Table(model.AccountTable).Save(a).Error
}

func (*KV) GetByAddress(address string) (*model.Account, error) {
	var account model.Account
	err := DB.Table(model.AccountTable).Where("address = ?", address).First(&account).Error
	return &account, err
}
