package model

import (
	"github.com/shopspring/decimal"
)

const (
	PricesTable = "prices"
)

type Prices struct {
	VUSDCPrice      decimal.Decimal
	VUSDTPrice      decimal.Decimal
	VBUSDPrice      decimal.Decimal
	VSXPPrice       decimal.Decimal
	VBNBPrice       decimal.Decimal
	VXVSPrice       decimal.Decimal
	VBTCPrice       decimal.Decimal
	VETHPrice       decimal.Decimal
	VLTCPrice       decimal.Decimal
	VXRPPrice       decimal.Decimal
	VBCHPrice       decimal.Decimal
	VDOTPrice       decimal.Decimal
	VLINKPrice      decimal.Decimal
	VDAIPrice       decimal.Decimal
	VFILPrice       decimal.Decimal
	VBETHPrice      decimal.Decimal
	VCANPrice       decimal.Decimal
	VADAPrice       decimal.Decimal
	VDOGEPrice      decimal.Decimal
	VMATICPrice     decimal.Decimal
	VCAKEPrice      decimal.Decimal
	VAAVEPrice      decimal.Decimal
	VTUSDPrice      decimal.Decimal
	VTRXPrice       decimal.Decimal
	VRESERVE1Price  decimal.Decimal
	VRESERVE2Price  decimal.Decimal
	VRESERVE3Price  decimal.Decimal
	VRESERVE4Price  decimal.Decimal
	VRESERVE5Price  decimal.Decimal
	VRESERVE6Price  decimal.Decimal
	VRESERVE7Price  decimal.Decimal
	VRESERVE8Price  decimal.Decimal
	VRESERVE9Price  decimal.Decimal
	VRESERVE10Price decimal.Decimal
}

func (Prices) TableName() string {
	return PricesTable
}
