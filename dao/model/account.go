package model

import (
	"github.com/shopspring/decimal"
)

const (
	AccountTable = "account"
)

type Account struct {
	Address           string          `gorm:"primary_key"`
	TotalBalance      decimal.Decimal //totalBalance in USDT
	TotalLoan         decimal.Decimal //totalLoan in USDT
	VUSDCBalance      decimal.Decimal
	VUSDTBalance      decimal.Decimal
	VBUSDBalance      decimal.Decimal
	VSXPBalance       decimal.Decimal
	VBNBBalance       decimal.Decimal
	VXVSBalance       decimal.Decimal
	VBTCBalance       decimal.Decimal
	VETHBalance       decimal.Decimal
	VLTCBalance       decimal.Decimal
	VXRPBalance       decimal.Decimal
	VBCHBalance       decimal.Decimal
	VDOTBalance       decimal.Decimal
	VLINKBalance      decimal.Decimal
	VDAIBalance       decimal.Decimal
	VFILBalance       decimal.Decimal
	VBETHBalance      decimal.Decimal
	VCANBalance       decimal.Decimal
	VADABalance       decimal.Decimal
	VDOGEBalance      decimal.Decimal
	VMATICBalance     decimal.Decimal
	VCAKEBalance      decimal.Decimal
	VAAVEBalance      decimal.Decimal
	VTUSDBalance      decimal.Decimal
	VTRXBalance       decimal.Decimal
	VRESERVE1Balance  decimal.Decimal
	VRESERVE2Balance  decimal.Decimal
	VRESERVE3Balance  decimal.Decimal
	VRESERVE4Balance  decimal.Decimal
	VRESERVE5Balance  decimal.Decimal
	VRESERVE6Balance  decimal.Decimal
	VRESERVE7Balance  decimal.Decimal
	VRESERVE8Balance  decimal.Decimal
	VRESERVE9Balance  decimal.Decimal
	VRESERVE10Balance decimal.Decimal

	VUSDCLoan      decimal.Decimal
	VUSDTLoan      decimal.Decimal
	VBUSDLoan      decimal.Decimal
	VSXPLoan       decimal.Decimal
	VBNBLoan       decimal.Decimal
	VXVSLoan       decimal.Decimal
	VBTCLoan       decimal.Decimal
	VETHLoan       decimal.Decimal
	VLTCLoan       decimal.Decimal
	VXRPLoan       decimal.Decimal
	VBCHLoan       decimal.Decimal
	VDOTLoan       decimal.Decimal
	VLINKLoan      decimal.Decimal
	VDAILoan       decimal.Decimal
	VFILLoan       decimal.Decimal
	VBETHLoan      decimal.Decimal
	VCANLoan       decimal.Decimal
	VADALoan       decimal.Decimal
	VDOGELoan      decimal.Decimal
	VMATICLoan     decimal.Decimal
	VCAKELoan      decimal.Decimal
	VAAVELoan      decimal.Decimal
	VTUSDLoan      decimal.Decimal
	VTRXLoan       decimal.Decimal
	VRESERVE1Loan  decimal.Decimal
	VRESERVE2Loan  decimal.Decimal
	VRESERVE3Loan  decimal.Decimal
	VRESERVE4Loan  decimal.Decimal
	VRESERVE5Loan  decimal.Decimal
	VRESERVE6Loan  decimal.Decimal
	VRESERVE7Loan  decimal.Decimal
	VRESERVE8Loan  decimal.Decimal
	VRESERVE9Loan  decimal.Decimal
	VRESERVE10Loan decimal.Decimal
}

func (Account) TableName() string {
	return AccountTable
}
