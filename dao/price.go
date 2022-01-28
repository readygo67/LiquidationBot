package dao

import "github.com/readygo67/LiquidationBot/dao/model"

type Prices struct{}

func (*Prices) CreateOrUpdate(a *model.Prices) error {
	return DB.Table(model.PricesTable).Save(a).Error
}

func (*Prices) GetPrices() (*model.Prices, error) {
	var prices model.Prices
	err := DB.Table(model.PricesTable).First(&prices).Error
	return &prices, err
}
