// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package venus

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ComptrollerMetaData contains all meta data concerning the Comptroller contract.
var ComptrollerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"ActionProtocolPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedBorrowerVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusSupplyIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedSupplierVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaiMinter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusVAIMintIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedVAIMinterVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributedVAIVaultVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowCap\",\"type\":\"uint256\"}],\"name\":\"NewBorrowCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBorrowCapGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"NewBorrowCapGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"}],\"name\":\"NewTreasuryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"}],\"name\":\"NewTreasuryGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTreasuryPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"NewTreasuryPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVAIControllerInterface\",\"name\":\"oldVAIController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractVAIControllerInterface\",\"name\":\"newVAIController\",\"type\":\"address\"}],\"name\":\"NewVAIController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVAIMintRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVAIMintRate\",\"type\":\"uint256\"}],\"name\":\"NewVAIMintRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseStartBlock_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseInterval_\",\"type\":\"uint256\"}],\"name\":\"NewVAIVaultInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVenusVAIRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVenusVAIRate\",\"type\":\"uint256\"}],\"name\":\"NewVenusVAIRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVenusVAIVaultRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVenusVAIVaultRate\",\"type\":\"uint256\"}],\"name\":\"NewVenusVAIVaultRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"VenusGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"VenusSpeedUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_grantXVS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"_setBorrowCapGuardian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBorrowCaps\",\"type\":\"uint256[]\"}],\"name\":\"_setMarketBorrowCaps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setProtocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"_setTreasuryData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVAIControllerInterface\",\"name\":\"vaiController_\",\"type\":\"address\"}],\"name\":\"_setVAIController\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVAIMintRate\",\"type\":\"uint256\"}],\"name\":\"_setVAIMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"releaseStartBlock_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReleaseAmount_\",\"type\":\"uint256\"}],\"name\":\"_setVAIVaultInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"venusSpeed\",\"type\":\"uint256\"}],\"name\":\"_setVenusSpeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"venusVAIVaultRate_\",\"type\":\"uint256\"}],\"name\":\"_setVenusVAIVaultRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractVToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractVToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getXVSAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastContributorBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateVAICalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVenus\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minReleaseAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintedVAIs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"releaseStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"repayVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMintedVAIOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiController\",\"outputs\":[{\"internalType\":\"contractVAIControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusContributorSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIVaultRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ComptrollerABI is the input ABI used to generate the binding from.
// Deprecated: Use ComptrollerMetaData.ABI instead.
var ComptrollerABI = ComptrollerMetaData.ABI

// Comptroller is an auto generated Go binding around an Ethereum contract.
type Comptroller struct {
	ComptrollerCaller     // Read-only binding to the contract
	ComptrollerTransactor // Write-only binding to the contract
	ComptrollerFilterer   // Log filterer for contract events
}

// ComptrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComptrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComptrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComptrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComptrollerSession struct {
	Contract     *Comptroller      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComptrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComptrollerCallerSession struct {
	Contract *ComptrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ComptrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComptrollerTransactorSession struct {
	Contract     *ComptrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ComptrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComptrollerRaw struct {
	Contract *Comptroller // Generic contract binding to access the raw methods on
}

// ComptrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComptrollerCallerRaw struct {
	Contract *ComptrollerCaller // Generic read-only contract binding to access the raw methods on
}

// ComptrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComptrollerTransactorRaw struct {
	Contract *ComptrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComptroller creates a new instance of Comptroller, bound to a specific deployed contract.
func NewComptroller(address common.Address, backend bind.ContractBackend) (*Comptroller, error) {
	contract, err := bindComptroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Comptroller{ComptrollerCaller: ComptrollerCaller{contract: contract}, ComptrollerTransactor: ComptrollerTransactor{contract: contract}, ComptrollerFilterer: ComptrollerFilterer{contract: contract}}, nil
}

// NewComptrollerCaller creates a new read-only instance of Comptroller, bound to a specific deployed contract.
func NewComptrollerCaller(address common.Address, caller bind.ContractCaller) (*ComptrollerCaller, error) {
	contract, err := bindComptroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerCaller{contract: contract}, nil
}

// NewComptrollerTransactor creates a new write-only instance of Comptroller, bound to a specific deployed contract.
func NewComptrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*ComptrollerTransactor, error) {
	contract, err := bindComptroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerTransactor{contract: contract}, nil
}

// NewComptrollerFilterer creates a new log filterer instance of Comptroller, bound to a specific deployed contract.
func NewComptrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*ComptrollerFilterer, error) {
	contract, err := bindComptroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComptrollerFilterer{contract: contract}, nil
}

// bindComptroller binds a generic wrapper to an already deployed contract.
func bindComptroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComptrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comptroller *ComptrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Comptroller.Contract.ComptrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comptroller *ComptrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comptroller.Contract.ComptrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comptroller *ComptrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comptroller.Contract.ComptrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comptroller *ComptrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Comptroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comptroller *ComptrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comptroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comptroller *ComptrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comptroller.Contract.contract.Transact(opts, method, params...)
}

// BorrowGuardianPaused1 is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) BorrowGuardianPaused1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "_borrowGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused1 is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) BorrowGuardianPaused1() (bool, error) {
	return _Comptroller.Contract.BorrowGuardianPaused1(&_Comptroller.CallOpts)
}

// BorrowGuardianPaused1 is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) BorrowGuardianPaused1() (bool, error) {
	return _Comptroller.Contract.BorrowGuardianPaused1(&_Comptroller.CallOpts)
}

// MintGuardianPaused1 is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) MintGuardianPaused1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "_mintGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused1 is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) MintGuardianPaused1() (bool, error) {
	return _Comptroller.Contract.MintGuardianPaused1(&_Comptroller.CallOpts)
}

// MintGuardianPaused1 is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) MintGuardianPaused1() (bool, error) {
	return _Comptroller.Contract.MintGuardianPaused1(&_Comptroller.CallOpts)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Comptroller *ComptrollerCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Comptroller *ComptrollerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Comptroller.Contract.AccountAssets(&_Comptroller.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Comptroller *ComptrollerCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Comptroller.Contract.AccountAssets(&_Comptroller.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Comptroller *ComptrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Comptroller *ComptrollerSession) Admin() (common.Address, error) {
	return _Comptroller.Contract.Admin(&_Comptroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Comptroller *ComptrollerCallerSession) Admin() (common.Address, error) {
	return _Comptroller.Contract.Admin(&_Comptroller.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Comptroller *ComptrollerCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Comptroller *ComptrollerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Comptroller.Contract.AllMarkets(&_Comptroller.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Comptroller *ComptrollerCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Comptroller.Contract.AllMarkets(&_Comptroller.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Comptroller *ComptrollerCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "borrowCapGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Comptroller *ComptrollerSession) BorrowCapGuardian() (common.Address, error) {
	return _Comptroller.Contract.BorrowCapGuardian(&_Comptroller.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Comptroller *ComptrollerCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _Comptroller.Contract.BorrowCapGuardian(&_Comptroller.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "borrowCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.BorrowCaps(&_Comptroller.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.BorrowCaps(&_Comptroller.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerCaller) BorrowGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "borrowGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _Comptroller.Contract.BorrowGuardianPaused(&_Comptroller.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerCallerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _Comptroller.Contract.BorrowGuardianPaused(&_Comptroller.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_Comptroller *ComptrollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, vToken common.Address) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "checkMembership", account, vToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_Comptroller *ComptrollerSession) CheckMembership(account common.Address, vToken common.Address) (bool, error) {
	return _Comptroller.Contract.CheckMembership(&_Comptroller.CallOpts, account, vToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_Comptroller *ComptrollerCallerSession) CheckMembership(account common.Address, vToken common.Address) (bool, error) {
	return _Comptroller.Contract.CheckMembership(&_Comptroller.CallOpts, account, vToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Comptroller *ComptrollerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Comptroller *ComptrollerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Comptroller.Contract.CloseFactorMantissa(&_Comptroller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Comptroller.Contract.CloseFactorMantissa(&_Comptroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.ComptrollerImplementation(&_Comptroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.ComptrollerImplementation(&_Comptroller.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Comptroller.Contract.GetAccountLiquidity(&_Comptroller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Comptroller.Contract.GetAccountLiquidity(&_Comptroller.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Comptroller *ComptrollerCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Comptroller *ComptrollerSession) GetAllMarkets() ([]common.Address, error) {
	return _Comptroller.Contract.GetAllMarkets(&_Comptroller.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Comptroller *ComptrollerCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _Comptroller.Contract.GetAllMarkets(&_Comptroller.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Comptroller *ComptrollerCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Comptroller *ComptrollerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Comptroller.Contract.GetAssetsIn(&_Comptroller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Comptroller *ComptrollerCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Comptroller.Contract.GetAssetsIn(&_Comptroller.CallOpts, account)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Comptroller *ComptrollerCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Comptroller *ComptrollerSession) GetBlockNumber() (*big.Int, error) {
	return _Comptroller.Contract.GetBlockNumber(&_Comptroller.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Comptroller.Contract.GetBlockNumber(&_Comptroller.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerCaller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getHypotheticalAccountLiquidity", account, vTokenModify, redeemTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerSession) GetHypotheticalAccountLiquidity(account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Comptroller.Contract.GetHypotheticalAccountLiquidity(&_Comptroller.CallOpts, account, vTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerCallerSession) GetHypotheticalAccountLiquidity(account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Comptroller.Contract.GetHypotheticalAccountLiquidity(&_Comptroller.CallOpts, account, vTokenModify, redeemTokens, borrowAmount)
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_Comptroller *ComptrollerCaller) GetXVSAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getXVSAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_Comptroller *ComptrollerSession) GetXVSAddress() (common.Address, error) {
	return _Comptroller.Contract.GetXVSAddress(&_Comptroller.CallOpts)
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_Comptroller *ComptrollerCallerSession) GetXVSAddress() (common.Address, error) {
	return _Comptroller.Contract.GetXVSAddress(&_Comptroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Comptroller *ComptrollerCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Comptroller *ComptrollerSession) IsComptroller() (bool, error) {
	return _Comptroller.Contract.IsComptroller(&_Comptroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) IsComptroller() (bool, error) {
	return _Comptroller.Contract.IsComptroller(&_Comptroller.CallOpts)
}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) LastContributorBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "lastContributorBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) LastContributorBlock(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.LastContributorBlock(&_Comptroller.CallOpts, arg0)
}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) LastContributorBlock(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.LastContributorBlock(&_Comptroller.CallOpts, arg0)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comptroller *ComptrollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", vTokenBorrowed, vTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comptroller *ComptrollerSession) LiquidateCalculateSeizeTokens(vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comptroller.Contract.LiquidateCalculateSeizeTokens(&_Comptroller.CallOpts, vTokenBorrowed, vTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comptroller *ComptrollerCallerSession) LiquidateCalculateSeizeTokens(vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comptroller.Contract.LiquidateCalculateSeizeTokens(&_Comptroller.CallOpts, vTokenBorrowed, vTokenCollateral, actualRepayAmount)
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comptroller *ComptrollerCaller) LiquidateVAICalculateSeizeTokens(opts *bind.CallOpts, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "liquidateVAICalculateSeizeTokens", vTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comptroller *ComptrollerSession) LiquidateVAICalculateSeizeTokens(vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comptroller.Contract.LiquidateVAICalculateSeizeTokens(&_Comptroller.CallOpts, vTokenCollateral, actualRepayAmount)
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comptroller *ComptrollerCallerSession) LiquidateVAICalculateSeizeTokens(vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comptroller.Contract.LiquidateVAICalculateSeizeTokens(&_Comptroller.CallOpts, vTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Comptroller *ComptrollerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Comptroller *ComptrollerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Comptroller.Contract.LiquidationIncentiveMantissa(&_Comptroller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Comptroller.Contract.LiquidationIncentiveMantissa(&_Comptroller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_Comptroller *ComptrollerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsVenus                  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsListed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CollateralFactorMantissa = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsVenus = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_Comptroller *ComptrollerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	return _Comptroller.Contract.Markets(&_Comptroller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_Comptroller *ComptrollerCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	return _Comptroller.Contract.Markets(&_Comptroller.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Comptroller *ComptrollerCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "maxAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Comptroller *ComptrollerSession) MaxAssets() (*big.Int, error) {
	return _Comptroller.Contract.MaxAssets(&_Comptroller.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) MaxAssets() (*big.Int, error) {
	return _Comptroller.Contract.MaxAssets(&_Comptroller.CallOpts)
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_Comptroller *ComptrollerCaller) MinReleaseAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "minReleaseAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_Comptroller *ComptrollerSession) MinReleaseAmount() (*big.Int, error) {
	return _Comptroller.Contract.MinReleaseAmount(&_Comptroller.CallOpts)
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) MinReleaseAmount() (*big.Int, error) {
	return _Comptroller.Contract.MinReleaseAmount(&_Comptroller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerCaller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "mintGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Comptroller.Contract.MintGuardianPaused(&_Comptroller.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerCallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Comptroller.Contract.MintGuardianPaused(&_Comptroller.CallOpts, arg0)
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) MintVAIGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "mintVAIGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) MintVAIGuardianPaused() (bool, error) {
	return _Comptroller.Contract.MintVAIGuardianPaused(&_Comptroller.CallOpts)
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) MintVAIGuardianPaused() (bool, error) {
	return _Comptroller.Contract.MintVAIGuardianPaused(&_Comptroller.CallOpts)
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) MintedVAIs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "mintedVAIs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) MintedVAIs(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.MintedVAIs(&_Comptroller.CallOpts, arg0)
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) MintedVAIs(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.MintedVAIs(&_Comptroller.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Comptroller *ComptrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Comptroller *ComptrollerSession) Oracle() (common.Address, error) {
	return _Comptroller.Contract.Oracle(&_Comptroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Comptroller *ComptrollerCallerSession) Oracle() (common.Address, error) {
	return _Comptroller.Contract.Oracle(&_Comptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comptroller *ComptrollerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comptroller *ComptrollerSession) PauseGuardian() (common.Address, error) {
	return _Comptroller.Contract.PauseGuardian(&_Comptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comptroller *ComptrollerCallerSession) PauseGuardian() (common.Address, error) {
	return _Comptroller.Contract.PauseGuardian(&_Comptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Comptroller *ComptrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Comptroller *ComptrollerSession) PendingAdmin() (common.Address, error) {
	return _Comptroller.Contract.PendingAdmin(&_Comptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Comptroller *ComptrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _Comptroller.Contract.PendingAdmin(&_Comptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.PendingComptrollerImplementation(&_Comptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.PendingComptrollerImplementation(&_Comptroller.CallOpts)
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) ProtocolPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "protocolPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) ProtocolPaused() (bool, error) {
	return _Comptroller.Contract.ProtocolPaused(&_Comptroller.CallOpts)
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) ProtocolPaused() (bool, error) {
	return _Comptroller.Contract.ProtocolPaused(&_Comptroller.CallOpts)
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_Comptroller *ComptrollerCaller) ReleaseStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "releaseStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_Comptroller *ComptrollerSession) ReleaseStartBlock() (*big.Int, error) {
	return _Comptroller.Contract.ReleaseStartBlock(&_Comptroller.CallOpts)
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) ReleaseStartBlock() (*big.Int, error) {
	return _Comptroller.Contract.ReleaseStartBlock(&_Comptroller.CallOpts)
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) RepayVAIGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "repayVAIGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) RepayVAIGuardianPaused() (bool, error) {
	return _Comptroller.Contract.RepayVAIGuardianPaused(&_Comptroller.CallOpts)
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) RepayVAIGuardianPaused() (bool, error) {
	return _Comptroller.Contract.RepayVAIGuardianPaused(&_Comptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "seizeGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) SeizeGuardianPaused() (bool, error) {
	return _Comptroller.Contract.SeizeGuardianPaused(&_Comptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) SeizeGuardianPaused() (bool, error) {
	return _Comptroller.Contract.SeizeGuardianPaused(&_Comptroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "transferGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) TransferGuardianPaused() (bool, error) {
	return _Comptroller.Contract.TransferGuardianPaused(&_Comptroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) TransferGuardianPaused() (bool, error) {
	return _Comptroller.Contract.TransferGuardianPaused(&_Comptroller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Comptroller *ComptrollerCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Comptroller *ComptrollerSession) TreasuryAddress() (common.Address, error) {
	return _Comptroller.Contract.TreasuryAddress(&_Comptroller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Comptroller *ComptrollerCallerSession) TreasuryAddress() (common.Address, error) {
	return _Comptroller.Contract.TreasuryAddress(&_Comptroller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_Comptroller *ComptrollerCaller) TreasuryGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "treasuryGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_Comptroller *ComptrollerSession) TreasuryGuardian() (common.Address, error) {
	return _Comptroller.Contract.TreasuryGuardian(&_Comptroller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_Comptroller *ComptrollerCallerSession) TreasuryGuardian() (common.Address, error) {
	return _Comptroller.Contract.TreasuryGuardian(&_Comptroller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Comptroller *ComptrollerCaller) TreasuryPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "treasuryPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Comptroller *ComptrollerSession) TreasuryPercent() (*big.Int, error) {
	return _Comptroller.Contract.TreasuryPercent(&_Comptroller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) TreasuryPercent() (*big.Int, error) {
	return _Comptroller.Contract.TreasuryPercent(&_Comptroller.CallOpts)
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_Comptroller *ComptrollerCaller) VaiController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "vaiController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_Comptroller *ComptrollerSession) VaiController() (common.Address, error) {
	return _Comptroller.Contract.VaiController(&_Comptroller.CallOpts)
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_Comptroller *ComptrollerCallerSession) VaiController() (common.Address, error) {
	return _Comptroller.Contract.VaiController(&_Comptroller.CallOpts)
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_Comptroller *ComptrollerCaller) VaiMintRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "vaiMintRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_Comptroller *ComptrollerSession) VaiMintRate() (*big.Int, error) {
	return _Comptroller.Contract.VaiMintRate(&_Comptroller.CallOpts)
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) VaiMintRate() (*big.Int, error) {
	return _Comptroller.Contract.VaiMintRate(&_Comptroller.CallOpts)
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_Comptroller *ComptrollerCaller) VaiVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "vaiVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_Comptroller *ComptrollerSession) VaiVaultAddress() (common.Address, error) {
	return _Comptroller.Contract.VaiVaultAddress(&_Comptroller.CallOpts)
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_Comptroller *ComptrollerCallerSession) VaiVaultAddress() (common.Address, error) {
	return _Comptroller.Contract.VaiVaultAddress(&_Comptroller.CallOpts)
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) VenusAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusAccrued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) VenusAccrued(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.VenusAccrued(&_Comptroller.CallOpts, arg0)
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) VenusAccrued(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.VenusAccrued(&_Comptroller.CallOpts, arg0)
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerCaller) VenusBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusBorrowState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerSession) VenusBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comptroller.Contract.VenusBorrowState(&_Comptroller.CallOpts, arg0)
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerCallerSession) VenusBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comptroller.Contract.VenusBorrowState(&_Comptroller.CallOpts, arg0)
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) VenusBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusBorrowerIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) VenusBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.VenusBorrowerIndex(&_Comptroller.CallOpts, arg0, arg1)
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) VenusBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.VenusBorrowerIndex(&_Comptroller.CallOpts, arg0, arg1)
}

// VenusContributorSpeeds is a free data retrieval call binding the contract method 0xa9046134.
//
// Solidity: function venusContributorSpeeds(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) VenusContributorSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusContributorSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusContributorSpeeds is a free data retrieval call binding the contract method 0xa9046134.
//
// Solidity: function venusContributorSpeeds(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) VenusContributorSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.VenusContributorSpeeds(&_Comptroller.CallOpts, arg0)
}

// VenusContributorSpeeds is a free data retrieval call binding the contract method 0xa9046134.
//
// Solidity: function venusContributorSpeeds(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) VenusContributorSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.VenusContributorSpeeds(&_Comptroller.CallOpts, arg0)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_Comptroller *ComptrollerCaller) VenusInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_Comptroller *ComptrollerSession) VenusInitialIndex() (*big.Int, error) {
	return _Comptroller.Contract.VenusInitialIndex(&_Comptroller.CallOpts)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_Comptroller *ComptrollerCallerSession) VenusInitialIndex() (*big.Int, error) {
	return _Comptroller.Contract.VenusInitialIndex(&_Comptroller.CallOpts)
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_Comptroller *ComptrollerCaller) VenusRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_Comptroller *ComptrollerSession) VenusRate() (*big.Int, error) {
	return _Comptroller.Contract.VenusRate(&_Comptroller.CallOpts)
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) VenusRate() (*big.Int, error) {
	return _Comptroller.Contract.VenusRate(&_Comptroller.CallOpts)
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) VenusSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) VenusSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.VenusSpeeds(&_Comptroller.CallOpts, arg0)
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) VenusSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.VenusSpeeds(&_Comptroller.CallOpts, arg0)
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) VenusSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusSupplierIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) VenusSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.VenusSupplierIndex(&_Comptroller.CallOpts, arg0, arg1)
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) VenusSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.VenusSupplierIndex(&_Comptroller.CallOpts, arg0, arg1)
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerCaller) VenusSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusSupplyState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerSession) VenusSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comptroller.Contract.VenusSupplyState(&_Comptroller.CallOpts, arg0)
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerCallerSession) VenusSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comptroller.Contract.VenusSupplyState(&_Comptroller.CallOpts, arg0)
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_Comptroller *ComptrollerCaller) VenusVAIRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusVAIRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_Comptroller *ComptrollerSession) VenusVAIRate() (*big.Int, error) {
	return _Comptroller.Contract.VenusVAIRate(&_Comptroller.CallOpts)
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) VenusVAIRate() (*big.Int, error) {
	return _Comptroller.Contract.VenusVAIRate(&_Comptroller.CallOpts)
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_Comptroller *ComptrollerCaller) VenusVAIVaultRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "venusVAIVaultRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_Comptroller *ComptrollerSession) VenusVAIVaultRate() (*big.Int, error) {
	return _Comptroller.Contract.VenusVAIVaultRate(&_Comptroller.CallOpts)
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) VenusVAIVaultRate() (*big.Int, error) {
	return _Comptroller.Contract.VenusVAIVaultRate(&_Comptroller.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comptroller *ComptrollerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comptroller *ComptrollerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.Become(&_Comptroller.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comptroller *ComptrollerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.Become(&_Comptroller.TransactOpts, unitroller)
}

// GrantXVS is a paid mutator transaction binding the contract method 0xa7604b41.
//
// Solidity: function _grantXVS(address recipient, uint256 amount) returns()
func (_Comptroller *ComptrollerTransactor) GrantXVS(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_grantXVS", recipient, amount)
}

// GrantXVS is a paid mutator transaction binding the contract method 0xa7604b41.
//
// Solidity: function _grantXVS(address recipient, uint256 amount) returns()
func (_Comptroller *ComptrollerSession) GrantXVS(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.GrantXVS(&_Comptroller.TransactOpts, recipient, amount)
}

// GrantXVS is a paid mutator transaction binding the contract method 0xa7604b41.
//
// Solidity: function _grantXVS(address recipient, uint256 amount) returns()
func (_Comptroller *ComptrollerTransactorSession) GrantXVS(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.GrantXVS(&_Comptroller.TransactOpts, recipient, amount)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Comptroller *ComptrollerTransactor) SetBorrowCapGuardian(opts *bind.TransactOpts, newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setBorrowCapGuardian", newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Comptroller *ComptrollerSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetBorrowCapGuardian(&_Comptroller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Comptroller *ComptrollerTransactorSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetBorrowCapGuardian(&_Comptroller.TransactOpts, newBorrowCapGuardian)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCloseFactor(&_Comptroller.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCloseFactor(&_Comptroller.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetCollateralFactor(opts *bind.TransactOpts, vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setCollateralFactor", vToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerSession) SetCollateralFactor(vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCollateralFactor(&_Comptroller.TransactOpts, vToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetCollateralFactor(vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCollateralFactor(&_Comptroller.TransactOpts, vToken, newCollateralFactorMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comptroller *ComptrollerSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetLiquidationIncentive(&_Comptroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetLiquidationIncentive(&_Comptroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_Comptroller *ComptrollerTransactor) SetMarketBorrowCaps(opts *bind.TransactOpts, vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setMarketBorrowCaps", vTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_Comptroller *ComptrollerSession) SetMarketBorrowCaps(vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMarketBorrowCaps(&_Comptroller.TransactOpts, vTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_Comptroller *ComptrollerTransactorSession) SetMarketBorrowCaps(vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMarketBorrowCaps(&_Comptroller.TransactOpts, vTokens, newBorrowCaps)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetPauseGuardian(opts *bind.TransactOpts, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setPauseGuardian", newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comptroller *ComptrollerSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPauseGuardian(&_Comptroller.TransactOpts, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPauseGuardian(&_Comptroller.TransactOpts, newPauseGuardian)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comptroller *ComptrollerSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPriceOracle(&_Comptroller.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPriceOracle(&_Comptroller.TransactOpts, newOracle)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactor) SetProtocolPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setProtocolPaused", state)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerSession) SetProtocolPaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetProtocolPaused(&_Comptroller.TransactOpts, state)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactorSession) SetProtocolPaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetProtocolPaused(&_Comptroller.TransactOpts, state)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetTreasuryData(opts *bind.TransactOpts, newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setTreasuryData", newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_Comptroller *ComptrollerSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetTreasuryData(&_Comptroller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetTreasuryData(&_Comptroller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetVAIController(opts *bind.TransactOpts, vaiController_ common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setVAIController", vaiController_)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_Comptroller *ComptrollerSession) SetVAIController(vaiController_ common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetVAIController(&_Comptroller.TransactOpts, vaiController_)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetVAIController(vaiController_ common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetVAIController(&_Comptroller.TransactOpts, vaiController_)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetVAIMintRate(opts *bind.TransactOpts, newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setVAIMintRate", newVAIMintRate)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_Comptroller *ComptrollerSession) SetVAIMintRate(newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetVAIMintRate(&_Comptroller.TransactOpts, newVAIMintRate)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetVAIMintRate(newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetVAIMintRate(&_Comptroller.TransactOpts, newVAIMintRate)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_Comptroller *ComptrollerTransactor) SetVAIVaultInfo(opts *bind.TransactOpts, vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setVAIVaultInfo", vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_Comptroller *ComptrollerSession) SetVAIVaultInfo(vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetVAIVaultInfo(&_Comptroller.TransactOpts, vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_Comptroller *ComptrollerTransactorSession) SetVAIVaultInfo(vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetVAIVaultInfo(&_Comptroller.TransactOpts, vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_Comptroller *ComptrollerTransactor) SetVenusSpeed(opts *bind.TransactOpts, vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setVenusSpeed", vToken, venusSpeed)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_Comptroller *ComptrollerSession) SetVenusSpeed(vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetVenusSpeed(&_Comptroller.TransactOpts, vToken, venusSpeed)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_Comptroller *ComptrollerTransactorSession) SetVenusSpeed(vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetVenusSpeed(&_Comptroller.TransactOpts, vToken, venusSpeed)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_Comptroller *ComptrollerTransactor) SetVenusVAIVaultRate(opts *bind.TransactOpts, venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setVenusVAIVaultRate", venusVAIVaultRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_Comptroller *ComptrollerSession) SetVenusVAIVaultRate(venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetVenusVAIVaultRate(&_Comptroller.TransactOpts, venusVAIVaultRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_Comptroller *ComptrollerTransactorSession) SetVenusVAIVaultRate(venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetVenusVAIVaultRate(&_Comptroller.TransactOpts, venusVAIVaultRate_)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SupportMarket(opts *bind.TransactOpts, vToken common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_supportMarket", vToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_Comptroller *ComptrollerSession) SupportMarket(vToken common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SupportMarket(&_Comptroller.TransactOpts, vToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SupportMarket(vToken common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SupportMarket(&_Comptroller.TransactOpts, vToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) BorrowAllowed(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "borrowAllowed", vToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) BorrowAllowed(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowAllowed(&_Comptroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) BorrowAllowed(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowAllowed(&_Comptroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_Comptroller *ComptrollerTransactor) BorrowVerify(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "borrowVerify", vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_Comptroller *ComptrollerSession) BorrowVerify(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowVerify(&_Comptroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_Comptroller *ComptrollerTransactorSession) BorrowVerify(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowVerify(&_Comptroller.TransactOpts, vToken, borrower, borrowAmount)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_Comptroller *ComptrollerTransactor) ClaimVenus(opts *bind.TransactOpts, holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "claimVenus", holder, vTokens)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_Comptroller *ComptrollerSession) ClaimVenus(holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimVenus(&_Comptroller.TransactOpts, holder, vTokens)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) ClaimVenus(holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimVenus(&_Comptroller.TransactOpts, holder, vTokens)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_Comptroller *ComptrollerTransactor) ClaimVenus0(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "claimVenus0", holder)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_Comptroller *ComptrollerSession) ClaimVenus0(holder common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimVenus0(&_Comptroller.TransactOpts, holder)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_Comptroller *ComptrollerTransactorSession) ClaimVenus0(holder common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimVenus0(&_Comptroller.TransactOpts, holder)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_Comptroller *ComptrollerTransactor) ClaimVenus1(opts *bind.TransactOpts, holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "claimVenus1", holders, vTokens, borrowers, suppliers)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_Comptroller *ComptrollerSession) ClaimVenus1(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimVenus1(&_Comptroller.TransactOpts, holders, vTokens, borrowers, suppliers)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_Comptroller *ComptrollerTransactorSession) ClaimVenus1(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimVenus1(&_Comptroller.TransactOpts, holders, vTokens, borrowers, suppliers)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_Comptroller *ComptrollerTransactor) EnterMarkets(opts *bind.TransactOpts, vTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "enterMarkets", vTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_Comptroller *ComptrollerSession) EnterMarkets(vTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.EnterMarkets(&_Comptroller.TransactOpts, vTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_Comptroller *ComptrollerTransactorSession) EnterMarkets(vTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.EnterMarkets(&_Comptroller.TransactOpts, vTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_Comptroller *ComptrollerTransactor) ExitMarket(opts *bind.TransactOpts, vTokenAddress common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "exitMarket", vTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_Comptroller *ComptrollerSession) ExitMarket(vTokenAddress common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ExitMarket(&_Comptroller.TransactOpts, vTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) ExitMarket(vTokenAddress common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ExitMarket(&_Comptroller.TransactOpts, vTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "liquidateBorrowAllowed", vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) LiquidateBorrowAllowed(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowAllowed(&_Comptroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) LiquidateBorrowAllowed(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowAllowed(&_Comptroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "liquidateBorrowVerify", vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerSession) LiquidateBorrowVerify(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowVerify(&_Comptroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) LiquidateBorrowVerify(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowVerify(&_Comptroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) MintAllowed(opts *bind.TransactOpts, vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "mintAllowed", vToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) MintAllowed(vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintAllowed(&_Comptroller.TransactOpts, vToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) MintAllowed(vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintAllowed(&_Comptroller.TransactOpts, vToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Comptroller *ComptrollerTransactor) MintVerify(opts *bind.TransactOpts, vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "mintVerify", vToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Comptroller *ComptrollerSession) MintVerify(vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintVerify(&_Comptroller.TransactOpts, vToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) MintVerify(vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintVerify(&_Comptroller.TransactOpts, vToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactor) RedeemAllowed(opts *bind.TransactOpts, vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "redeemAllowed", vToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comptroller *ComptrollerSession) RedeemAllowed(vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemAllowed(&_Comptroller.TransactOpts, vToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) RedeemAllowed(vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemAllowed(&_Comptroller.TransactOpts, vToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comptroller *ComptrollerTransactor) RedeemVerify(opts *bind.TransactOpts, vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "redeemVerify", vToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comptroller *ComptrollerSession) RedeemVerify(vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemVerify(&_Comptroller.TransactOpts, vToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) RedeemVerify(vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemVerify(&_Comptroller.TransactOpts, vToken, redeemer, redeemAmount, redeemTokens)
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_Comptroller *ComptrollerTransactor) ReleaseToVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "releaseToVault")
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_Comptroller *ComptrollerSession) ReleaseToVault() (*types.Transaction, error) {
	return _Comptroller.Contract.ReleaseToVault(&_Comptroller.TransactOpts)
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_Comptroller *ComptrollerTransactorSession) ReleaseToVault() (*types.Transaction, error) {
	return _Comptroller.Contract.ReleaseToVault(&_Comptroller.TransactOpts)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "repayBorrowAllowed", vToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) RepayBorrowAllowed(vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowAllowed(&_Comptroller.TransactOpts, vToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) RepayBorrowAllowed(vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowAllowed(&_Comptroller.TransactOpts, vToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Comptroller *ComptrollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "repayBorrowVerify", vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Comptroller *ComptrollerSession) RepayBorrowVerify(vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowVerify(&_Comptroller.TransactOpts, vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Comptroller *ComptrollerTransactorSession) RepayBorrowVerify(vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowVerify(&_Comptroller.TransactOpts, vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SeizeAllowed(opts *bind.TransactOpts, vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "seizeAllowed", vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comptroller *ComptrollerSession) SeizeAllowed(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeAllowed(&_Comptroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SeizeAllowed(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeAllowed(&_Comptroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactor) SeizeVerify(opts *bind.TransactOpts, vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "seizeVerify", vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerSession) SeizeVerify(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeVerify(&_Comptroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) SeizeVerify(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeVerify(&_Comptroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetMintedVAIOf(opts *bind.TransactOpts, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "setMintedVAIOf", owner, amount)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_Comptroller *ComptrollerSession) SetMintedVAIOf(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMintedVAIOf(&_Comptroller.TransactOpts, owner, amount)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetMintedVAIOf(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMintedVAIOf(&_Comptroller.TransactOpts, owner, amount)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactor) TransferAllowed(opts *bind.TransactOpts, vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "transferAllowed", vToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comptroller *ComptrollerSession) TransferAllowed(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferAllowed(&_Comptroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) TransferAllowed(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferAllowed(&_Comptroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_Comptroller *ComptrollerTransactor) TransferVerify(opts *bind.TransactOpts, vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "transferVerify", vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_Comptroller *ComptrollerSession) TransferVerify(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferVerify(&_Comptroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) TransferVerify(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferVerify(&_Comptroller.TransactOpts, vToken, src, dst, transferTokens)
}

// ComptrollerActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the Comptroller contract.
type ComptrollerActionPausedIterator struct {
	Event *ComptrollerActionPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerActionPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerActionPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerActionPaused represents a ActionPaused event raised by the Comptroller contract.
type ComptrollerActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) FilterActionPaused(opts *bind.FilterOpts) (*ComptrollerActionPausedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &ComptrollerActionPausedIterator{contract: _Comptroller.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *ComptrollerActionPaused) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerActionPaused)
				if err := _Comptroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) ParseActionPaused(log types.Log) (*ComptrollerActionPaused, error) {
	event := new(ComptrollerActionPaused)
	if err := _Comptroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerActionPaused0Iterator is returned from FilterActionPaused0 and is used to iterate over the raw logs and unpacked data for ActionPaused0 events raised by the Comptroller contract.
type ComptrollerActionPaused0Iterator struct {
	Event *ComptrollerActionPaused0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerActionPaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerActionPaused0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerActionPaused0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerActionPaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerActionPaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerActionPaused0 represents a ActionPaused0 event raised by the Comptroller contract.
type ComptrollerActionPaused0 struct {
	VToken     common.Address
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused0 is a free log retrieval operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) FilterActionPaused0(opts *bind.FilterOpts) (*ComptrollerActionPaused0Iterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return &ComptrollerActionPaused0Iterator{contract: _Comptroller.contract, event: "ActionPaused0", logs: logs, sub: sub}, nil
}

// WatchActionPaused0 is a free log subscription operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) WatchActionPaused0(opts *bind.WatchOpts, sink chan<- *ComptrollerActionPaused0) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerActionPaused0)
				if err := _Comptroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionPaused0 is a log parse operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) ParseActionPaused0(log types.Log) (*ComptrollerActionPaused0, error) {
	event := new(ComptrollerActionPaused0)
	if err := _Comptroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerActionProtocolPausedIterator is returned from FilterActionProtocolPaused and is used to iterate over the raw logs and unpacked data for ActionProtocolPaused events raised by the Comptroller contract.
type ComptrollerActionProtocolPausedIterator struct {
	Event *ComptrollerActionProtocolPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerActionProtocolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerActionProtocolPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerActionProtocolPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerActionProtocolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerActionProtocolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerActionProtocolPaused represents a ActionProtocolPaused event raised by the Comptroller contract.
type ComptrollerActionProtocolPaused struct {
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterActionProtocolPaused is a free log retrieval operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_Comptroller *ComptrollerFilterer) FilterActionProtocolPaused(opts *bind.FilterOpts) (*ComptrollerActionProtocolPausedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "ActionProtocolPaused")
	if err != nil {
		return nil, err
	}
	return &ComptrollerActionProtocolPausedIterator{contract: _Comptroller.contract, event: "ActionProtocolPaused", logs: logs, sub: sub}, nil
}

// WatchActionProtocolPaused is a free log subscription operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_Comptroller *ComptrollerFilterer) WatchActionProtocolPaused(opts *bind.WatchOpts, sink chan<- *ComptrollerActionProtocolPaused) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "ActionProtocolPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerActionProtocolPaused)
				if err := _Comptroller.contract.UnpackLog(event, "ActionProtocolPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionProtocolPaused is a log parse operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_Comptroller *ComptrollerFilterer) ParseActionProtocolPaused(log types.Log) (*ComptrollerActionProtocolPaused, error) {
	event := new(ComptrollerActionProtocolPaused)
	if err := _Comptroller.contract.UnpackLog(event, "ActionProtocolPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerDistributedBorrowerVenusIterator is returned from FilterDistributedBorrowerVenus and is used to iterate over the raw logs and unpacked data for DistributedBorrowerVenus events raised by the Comptroller contract.
type ComptrollerDistributedBorrowerVenusIterator struct {
	Event *ComptrollerDistributedBorrowerVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerDistributedBorrowerVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerDistributedBorrowerVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerDistributedBorrowerVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerDistributedBorrowerVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerDistributedBorrowerVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerDistributedBorrowerVenus represents a DistributedBorrowerVenus event raised by the Comptroller contract.
type ComptrollerDistributedBorrowerVenus struct {
	VToken           common.Address
	Borrower         common.Address
	VenusDelta       *big.Int
	VenusBorrowIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributedBorrowerVenus is a free log retrieval operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_Comptroller *ComptrollerFilterer) FilterDistributedBorrowerVenus(opts *bind.FilterOpts, vToken []common.Address, borrower []common.Address) (*ComptrollerDistributedBorrowerVenusIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "DistributedBorrowerVenus", vTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerDistributedBorrowerVenusIterator{contract: _Comptroller.contract, event: "DistributedBorrowerVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedBorrowerVenus is a free log subscription operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_Comptroller *ComptrollerFilterer) WatchDistributedBorrowerVenus(opts *bind.WatchOpts, sink chan<- *ComptrollerDistributedBorrowerVenus, vToken []common.Address, borrower []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "DistributedBorrowerVenus", vTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerDistributedBorrowerVenus)
				if err := _Comptroller.contract.UnpackLog(event, "DistributedBorrowerVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedBorrowerVenus is a log parse operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_Comptroller *ComptrollerFilterer) ParseDistributedBorrowerVenus(log types.Log) (*ComptrollerDistributedBorrowerVenus, error) {
	event := new(ComptrollerDistributedBorrowerVenus)
	if err := _Comptroller.contract.UnpackLog(event, "DistributedBorrowerVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerDistributedSupplierVenusIterator is returned from FilterDistributedSupplierVenus and is used to iterate over the raw logs and unpacked data for DistributedSupplierVenus events raised by the Comptroller contract.
type ComptrollerDistributedSupplierVenusIterator struct {
	Event *ComptrollerDistributedSupplierVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerDistributedSupplierVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerDistributedSupplierVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerDistributedSupplierVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerDistributedSupplierVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerDistributedSupplierVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerDistributedSupplierVenus represents a DistributedSupplierVenus event raised by the Comptroller contract.
type ComptrollerDistributedSupplierVenus struct {
	VToken           common.Address
	Supplier         common.Address
	VenusDelta       *big.Int
	VenusSupplyIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributedSupplierVenus is a free log retrieval operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_Comptroller *ComptrollerFilterer) FilterDistributedSupplierVenus(opts *bind.FilterOpts, vToken []common.Address, supplier []common.Address) (*ComptrollerDistributedSupplierVenusIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "DistributedSupplierVenus", vTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerDistributedSupplierVenusIterator{contract: _Comptroller.contract, event: "DistributedSupplierVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedSupplierVenus is a free log subscription operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_Comptroller *ComptrollerFilterer) WatchDistributedSupplierVenus(opts *bind.WatchOpts, sink chan<- *ComptrollerDistributedSupplierVenus, vToken []common.Address, supplier []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "DistributedSupplierVenus", vTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerDistributedSupplierVenus)
				if err := _Comptroller.contract.UnpackLog(event, "DistributedSupplierVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedSupplierVenus is a log parse operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_Comptroller *ComptrollerFilterer) ParseDistributedSupplierVenus(log types.Log) (*ComptrollerDistributedSupplierVenus, error) {
	event := new(ComptrollerDistributedSupplierVenus)
	if err := _Comptroller.contract.UnpackLog(event, "DistributedSupplierVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerDistributedVAIMinterVenusIterator is returned from FilterDistributedVAIMinterVenus and is used to iterate over the raw logs and unpacked data for DistributedVAIMinterVenus events raised by the Comptroller contract.
type ComptrollerDistributedVAIMinterVenusIterator struct {
	Event *ComptrollerDistributedVAIMinterVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerDistributedVAIMinterVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerDistributedVAIMinterVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerDistributedVAIMinterVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerDistributedVAIMinterVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerDistributedVAIMinterVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerDistributedVAIMinterVenus represents a DistributedVAIMinterVenus event raised by the Comptroller contract.
type ComptrollerDistributedVAIMinterVenus struct {
	VaiMinter         common.Address
	VenusDelta        *big.Int
	VenusVAIMintIndex *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDistributedVAIMinterVenus is a free log retrieval operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_Comptroller *ComptrollerFilterer) FilterDistributedVAIMinterVenus(opts *bind.FilterOpts, vaiMinter []common.Address) (*ComptrollerDistributedVAIMinterVenusIterator, error) {

	var vaiMinterRule []interface{}
	for _, vaiMinterItem := range vaiMinter {
		vaiMinterRule = append(vaiMinterRule, vaiMinterItem)
	}

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "DistributedVAIMinterVenus", vaiMinterRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerDistributedVAIMinterVenusIterator{contract: _Comptroller.contract, event: "DistributedVAIMinterVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedVAIMinterVenus is a free log subscription operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_Comptroller *ComptrollerFilterer) WatchDistributedVAIMinterVenus(opts *bind.WatchOpts, sink chan<- *ComptrollerDistributedVAIMinterVenus, vaiMinter []common.Address) (event.Subscription, error) {

	var vaiMinterRule []interface{}
	for _, vaiMinterItem := range vaiMinter {
		vaiMinterRule = append(vaiMinterRule, vaiMinterItem)
	}

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "DistributedVAIMinterVenus", vaiMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerDistributedVAIMinterVenus)
				if err := _Comptroller.contract.UnpackLog(event, "DistributedVAIMinterVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedVAIMinterVenus is a log parse operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_Comptroller *ComptrollerFilterer) ParseDistributedVAIMinterVenus(log types.Log) (*ComptrollerDistributedVAIMinterVenus, error) {
	event := new(ComptrollerDistributedVAIMinterVenus)
	if err := _Comptroller.contract.UnpackLog(event, "DistributedVAIMinterVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerDistributedVAIVaultVenusIterator is returned from FilterDistributedVAIVaultVenus and is used to iterate over the raw logs and unpacked data for DistributedVAIVaultVenus events raised by the Comptroller contract.
type ComptrollerDistributedVAIVaultVenusIterator struct {
	Event *ComptrollerDistributedVAIVaultVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerDistributedVAIVaultVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerDistributedVAIVaultVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerDistributedVAIVaultVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerDistributedVAIVaultVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerDistributedVAIVaultVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerDistributedVAIVaultVenus represents a DistributedVAIVaultVenus event raised by the Comptroller contract.
type ComptrollerDistributedVAIVaultVenus struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributedVAIVaultVenus is a free log retrieval operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_Comptroller *ComptrollerFilterer) FilterDistributedVAIVaultVenus(opts *bind.FilterOpts) (*ComptrollerDistributedVAIVaultVenusIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "DistributedVAIVaultVenus")
	if err != nil {
		return nil, err
	}
	return &ComptrollerDistributedVAIVaultVenusIterator{contract: _Comptroller.contract, event: "DistributedVAIVaultVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedVAIVaultVenus is a free log subscription operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_Comptroller *ComptrollerFilterer) WatchDistributedVAIVaultVenus(opts *bind.WatchOpts, sink chan<- *ComptrollerDistributedVAIVaultVenus) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "DistributedVAIVaultVenus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerDistributedVAIVaultVenus)
				if err := _Comptroller.contract.UnpackLog(event, "DistributedVAIVaultVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedVAIVaultVenus is a log parse operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_Comptroller *ComptrollerFilterer) ParseDistributedVAIVaultVenus(log types.Log) (*ComptrollerDistributedVAIVaultVenus, error) {
	event := new(ComptrollerDistributedVAIVaultVenus)
	if err := _Comptroller.contract.UnpackLog(event, "DistributedVAIVaultVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Comptroller contract.
type ComptrollerFailureIterator struct {
	Event *ComptrollerFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerFailure represents a Failure event raised by the Comptroller contract.
type ComptrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comptroller *ComptrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*ComptrollerFailureIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &ComptrollerFailureIterator{contract: _Comptroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comptroller *ComptrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *ComptrollerFailure) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerFailure)
				if err := _Comptroller.contract.UnpackLog(event, "Failure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comptroller *ComptrollerFilterer) ParseFailure(log types.Log) (*ComptrollerFailure, error) {
	event := new(ComptrollerFailure)
	if err := _Comptroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the Comptroller contract.
type ComptrollerMarketEnteredIterator struct {
	Event *ComptrollerMarketEntered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerMarketEntered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerMarketEntered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerMarketEntered represents a MarketEntered event raised by the Comptroller contract.
type ComptrollerMarketEntered struct {
	VToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_Comptroller *ComptrollerFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*ComptrollerMarketEnteredIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &ComptrollerMarketEnteredIterator{contract: _Comptroller.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_Comptroller *ComptrollerFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *ComptrollerMarketEntered) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerMarketEntered)
				if err := _Comptroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_Comptroller *ComptrollerFilterer) ParseMarketEntered(log types.Log) (*ComptrollerMarketEntered, error) {
	event := new(ComptrollerMarketEntered)
	if err := _Comptroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the Comptroller contract.
type ComptrollerMarketExitedIterator struct {
	Event *ComptrollerMarketExited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerMarketExited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerMarketExited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerMarketExited represents a MarketExited event raised by the Comptroller contract.
type ComptrollerMarketExited struct {
	VToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_Comptroller *ComptrollerFilterer) FilterMarketExited(opts *bind.FilterOpts) (*ComptrollerMarketExitedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &ComptrollerMarketExitedIterator{contract: _Comptroller.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_Comptroller *ComptrollerFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *ComptrollerMarketExited) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerMarketExited)
				if err := _Comptroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_Comptroller *ComptrollerFilterer) ParseMarketExited(log types.Log) (*ComptrollerMarketExited, error) {
	event := new(ComptrollerMarketExited)
	if err := _Comptroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerMarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the Comptroller contract.
type ComptrollerMarketListedIterator struct {
	Event *ComptrollerMarketListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerMarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerMarketListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerMarketListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerMarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerMarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerMarketListed represents a MarketListed event raised by the Comptroller contract.
type ComptrollerMarketListed struct {
	VToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_Comptroller *ComptrollerFilterer) FilterMarketListed(opts *bind.FilterOpts) (*ComptrollerMarketListedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &ComptrollerMarketListedIterator{contract: _Comptroller.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_Comptroller *ComptrollerFilterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *ComptrollerMarketListed) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerMarketListed)
				if err := _Comptroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_Comptroller *ComptrollerFilterer) ParseMarketListed(log types.Log) (*ComptrollerMarketListed, error) {
	event := new(ComptrollerMarketListed)
	if err := _Comptroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewBorrowCapIterator is returned from FilterNewBorrowCap and is used to iterate over the raw logs and unpacked data for NewBorrowCap events raised by the Comptroller contract.
type ComptrollerNewBorrowCapIterator struct {
	Event *ComptrollerNewBorrowCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewBorrowCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewBorrowCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewBorrowCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewBorrowCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewBorrowCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewBorrowCap represents a NewBorrowCap event raised by the Comptroller contract.
type ComptrollerNewBorrowCap struct {
	VToken       common.Address
	NewBorrowCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCap is a free log retrieval operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_Comptroller *ComptrollerFilterer) FilterNewBorrowCap(opts *bind.FilterOpts, vToken []common.Address) (*ComptrollerNewBorrowCapIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewBorrowCap", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewBorrowCapIterator{contract: _Comptroller.contract, event: "NewBorrowCap", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCap is a free log subscription operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_Comptroller *ComptrollerFilterer) WatchNewBorrowCap(opts *bind.WatchOpts, sink chan<- *ComptrollerNewBorrowCap, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewBorrowCap", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewBorrowCap)
				if err := _Comptroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBorrowCap is a log parse operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_Comptroller *ComptrollerFilterer) ParseNewBorrowCap(log types.Log) (*ComptrollerNewBorrowCap, error) {
	event := new(ComptrollerNewBorrowCap)
	if err := _Comptroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewBorrowCapGuardianIterator is returned from FilterNewBorrowCapGuardian and is used to iterate over the raw logs and unpacked data for NewBorrowCapGuardian events raised by the Comptroller contract.
type ComptrollerNewBorrowCapGuardianIterator struct {
	Event *ComptrollerNewBorrowCapGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewBorrowCapGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewBorrowCapGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewBorrowCapGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewBorrowCapGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewBorrowCapGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewBorrowCapGuardian represents a NewBorrowCapGuardian event raised by the Comptroller contract.
type ComptrollerNewBorrowCapGuardian struct {
	OldBorrowCapGuardian common.Address
	NewBorrowCapGuardian common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCapGuardian is a free log retrieval operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Comptroller *ComptrollerFilterer) FilterNewBorrowCapGuardian(opts *bind.FilterOpts) (*ComptrollerNewBorrowCapGuardianIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewBorrowCapGuardianIterator{contract: _Comptroller.contract, event: "NewBorrowCapGuardian", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCapGuardian is a free log subscription operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Comptroller *ComptrollerFilterer) WatchNewBorrowCapGuardian(opts *bind.WatchOpts, sink chan<- *ComptrollerNewBorrowCapGuardian) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewBorrowCapGuardian)
				if err := _Comptroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBorrowCapGuardian is a log parse operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Comptroller *ComptrollerFilterer) ParseNewBorrowCapGuardian(log types.Log) (*ComptrollerNewBorrowCapGuardian, error) {
	event := new(ComptrollerNewBorrowCapGuardian)
	if err := _Comptroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the Comptroller contract.
type ComptrollerNewCloseFactorIterator struct {
	Event *ComptrollerNewCloseFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewCloseFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewCloseFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewCloseFactor represents a NewCloseFactor event raised by the Comptroller contract.
type ComptrollerNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comptroller *ComptrollerFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*ComptrollerNewCloseFactorIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewCloseFactorIterator{contract: _Comptroller.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comptroller *ComptrollerFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *ComptrollerNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewCloseFactor)
				if err := _Comptroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comptroller *ComptrollerFilterer) ParseNewCloseFactor(log types.Log) (*ComptrollerNewCloseFactor, error) {
	event := new(ComptrollerNewCloseFactor)
	if err := _Comptroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the Comptroller contract.
type ComptrollerNewCollateralFactorIterator struct {
	Event *ComptrollerNewCollateralFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewCollateralFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewCollateralFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewCollateralFactor represents a NewCollateralFactor event raised by the Comptroller contract.
type ComptrollerNewCollateralFactor struct {
	VToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comptroller *ComptrollerFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*ComptrollerNewCollateralFactorIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewCollateralFactorIterator{contract: _Comptroller.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comptroller *ComptrollerFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *ComptrollerNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewCollateralFactor)
				if err := _Comptroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comptroller *ComptrollerFilterer) ParseNewCollateralFactor(log types.Log) (*ComptrollerNewCollateralFactor, error) {
	event := new(ComptrollerNewCollateralFactor)
	if err := _Comptroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the Comptroller contract.
type ComptrollerNewLiquidationIncentiveIterator struct {
	Event *ComptrollerNewLiquidationIncentive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewLiquidationIncentive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewLiquidationIncentive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the Comptroller contract.
type ComptrollerNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comptroller *ComptrollerFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*ComptrollerNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewLiquidationIncentiveIterator{contract: _Comptroller.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comptroller *ComptrollerFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *ComptrollerNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewLiquidationIncentive)
				if err := _Comptroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comptroller *ComptrollerFilterer) ParseNewLiquidationIncentive(log types.Log) (*ComptrollerNewLiquidationIncentive, error) {
	event := new(ComptrollerNewLiquidationIncentive)
	if err := _Comptroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the Comptroller contract.
type ComptrollerNewPauseGuardianIterator struct {
	Event *ComptrollerNewPauseGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewPauseGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewPauseGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewPauseGuardian represents a NewPauseGuardian event raised by the Comptroller contract.
type ComptrollerNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comptroller *ComptrollerFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*ComptrollerNewPauseGuardianIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewPauseGuardianIterator{contract: _Comptroller.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comptroller *ComptrollerFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *ComptrollerNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewPauseGuardian)
				if err := _Comptroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comptroller *ComptrollerFilterer) ParseNewPauseGuardian(log types.Log) (*ComptrollerNewPauseGuardian, error) {
	event := new(ComptrollerNewPauseGuardian)
	if err := _Comptroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the Comptroller contract.
type ComptrollerNewPriceOracleIterator struct {
	Event *ComptrollerNewPriceOracle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewPriceOracle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewPriceOracle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewPriceOracle represents a NewPriceOracle event raised by the Comptroller contract.
type ComptrollerNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comptroller *ComptrollerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*ComptrollerNewPriceOracleIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewPriceOracleIterator{contract: _Comptroller.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comptroller *ComptrollerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *ComptrollerNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewPriceOracle)
				if err := _Comptroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comptroller *ComptrollerFilterer) ParseNewPriceOracle(log types.Log) (*ComptrollerNewPriceOracle, error) {
	event := new(ComptrollerNewPriceOracle)
	if err := _Comptroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewTreasuryAddressIterator is returned from FilterNewTreasuryAddress and is used to iterate over the raw logs and unpacked data for NewTreasuryAddress events raised by the Comptroller contract.
type ComptrollerNewTreasuryAddressIterator struct {
	Event *ComptrollerNewTreasuryAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewTreasuryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewTreasuryAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewTreasuryAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewTreasuryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewTreasuryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewTreasuryAddress represents a NewTreasuryAddress event raised by the Comptroller contract.
type ComptrollerNewTreasuryAddress struct {
	OldTreasuryAddress common.Address
	NewTreasuryAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryAddress is a free log retrieval operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_Comptroller *ComptrollerFilterer) FilterNewTreasuryAddress(opts *bind.FilterOpts) (*ComptrollerNewTreasuryAddressIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewTreasuryAddressIterator{contract: _Comptroller.contract, event: "NewTreasuryAddress", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryAddress is a free log subscription operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_Comptroller *ComptrollerFilterer) WatchNewTreasuryAddress(opts *bind.WatchOpts, sink chan<- *ComptrollerNewTreasuryAddress) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewTreasuryAddress)
				if err := _Comptroller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryAddress is a log parse operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_Comptroller *ComptrollerFilterer) ParseNewTreasuryAddress(log types.Log) (*ComptrollerNewTreasuryAddress, error) {
	event := new(ComptrollerNewTreasuryAddress)
	if err := _Comptroller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewTreasuryGuardianIterator is returned from FilterNewTreasuryGuardian and is used to iterate over the raw logs and unpacked data for NewTreasuryGuardian events raised by the Comptroller contract.
type ComptrollerNewTreasuryGuardianIterator struct {
	Event *ComptrollerNewTreasuryGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewTreasuryGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewTreasuryGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewTreasuryGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewTreasuryGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewTreasuryGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewTreasuryGuardian represents a NewTreasuryGuardian event raised by the Comptroller contract.
type ComptrollerNewTreasuryGuardian struct {
	OldTreasuryGuardian common.Address
	NewTreasuryGuardian common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryGuardian is a free log retrieval operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_Comptroller *ComptrollerFilterer) FilterNewTreasuryGuardian(opts *bind.FilterOpts) (*ComptrollerNewTreasuryGuardianIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewTreasuryGuardianIterator{contract: _Comptroller.contract, event: "NewTreasuryGuardian", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryGuardian is a free log subscription operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_Comptroller *ComptrollerFilterer) WatchNewTreasuryGuardian(opts *bind.WatchOpts, sink chan<- *ComptrollerNewTreasuryGuardian) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewTreasuryGuardian)
				if err := _Comptroller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryGuardian is a log parse operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_Comptroller *ComptrollerFilterer) ParseNewTreasuryGuardian(log types.Log) (*ComptrollerNewTreasuryGuardian, error) {
	event := new(ComptrollerNewTreasuryGuardian)
	if err := _Comptroller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewTreasuryPercentIterator is returned from FilterNewTreasuryPercent and is used to iterate over the raw logs and unpacked data for NewTreasuryPercent events raised by the Comptroller contract.
type ComptrollerNewTreasuryPercentIterator struct {
	Event *ComptrollerNewTreasuryPercent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewTreasuryPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewTreasuryPercent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewTreasuryPercent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewTreasuryPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewTreasuryPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewTreasuryPercent represents a NewTreasuryPercent event raised by the Comptroller contract.
type ComptrollerNewTreasuryPercent struct {
	OldTreasuryPercent *big.Int
	NewTreasuryPercent *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryPercent is a free log retrieval operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_Comptroller *ComptrollerFilterer) FilterNewTreasuryPercent(opts *bind.FilterOpts) (*ComptrollerNewTreasuryPercentIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewTreasuryPercentIterator{contract: _Comptroller.contract, event: "NewTreasuryPercent", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryPercent is a free log subscription operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_Comptroller *ComptrollerFilterer) WatchNewTreasuryPercent(opts *bind.WatchOpts, sink chan<- *ComptrollerNewTreasuryPercent) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewTreasuryPercent)
				if err := _Comptroller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryPercent is a log parse operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_Comptroller *ComptrollerFilterer) ParseNewTreasuryPercent(log types.Log) (*ComptrollerNewTreasuryPercent, error) {
	event := new(ComptrollerNewTreasuryPercent)
	if err := _Comptroller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewVAIControllerIterator is returned from FilterNewVAIController and is used to iterate over the raw logs and unpacked data for NewVAIController events raised by the Comptroller contract.
type ComptrollerNewVAIControllerIterator struct {
	Event *ComptrollerNewVAIController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewVAIControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewVAIController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewVAIController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewVAIControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewVAIControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewVAIController represents a NewVAIController event raised by the Comptroller contract.
type ComptrollerNewVAIController struct {
	OldVAIController common.Address
	NewVAIController common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewVAIController is a free log retrieval operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_Comptroller *ComptrollerFilterer) FilterNewVAIController(opts *bind.FilterOpts) (*ComptrollerNewVAIControllerIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewVAIController")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewVAIControllerIterator{contract: _Comptroller.contract, event: "NewVAIController", logs: logs, sub: sub}, nil
}

// WatchNewVAIController is a free log subscription operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_Comptroller *ComptrollerFilterer) WatchNewVAIController(opts *bind.WatchOpts, sink chan<- *ComptrollerNewVAIController) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewVAIController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewVAIController)
				if err := _Comptroller.contract.UnpackLog(event, "NewVAIController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIController is a log parse operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_Comptroller *ComptrollerFilterer) ParseNewVAIController(log types.Log) (*ComptrollerNewVAIController, error) {
	event := new(ComptrollerNewVAIController)
	if err := _Comptroller.contract.UnpackLog(event, "NewVAIController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewVAIMintRateIterator is returned from FilterNewVAIMintRate and is used to iterate over the raw logs and unpacked data for NewVAIMintRate events raised by the Comptroller contract.
type ComptrollerNewVAIMintRateIterator struct {
	Event *ComptrollerNewVAIMintRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewVAIMintRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewVAIMintRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewVAIMintRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewVAIMintRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewVAIMintRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewVAIMintRate represents a NewVAIMintRate event raised by the Comptroller contract.
type ComptrollerNewVAIMintRate struct {
	OldVAIMintRate *big.Int
	NewVAIMintRate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewVAIMintRate is a free log retrieval operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_Comptroller *ComptrollerFilterer) FilterNewVAIMintRate(opts *bind.FilterOpts) (*ComptrollerNewVAIMintRateIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewVAIMintRate")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewVAIMintRateIterator{contract: _Comptroller.contract, event: "NewVAIMintRate", logs: logs, sub: sub}, nil
}

// WatchNewVAIMintRate is a free log subscription operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_Comptroller *ComptrollerFilterer) WatchNewVAIMintRate(opts *bind.WatchOpts, sink chan<- *ComptrollerNewVAIMintRate) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewVAIMintRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewVAIMintRate)
				if err := _Comptroller.contract.UnpackLog(event, "NewVAIMintRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIMintRate is a log parse operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_Comptroller *ComptrollerFilterer) ParseNewVAIMintRate(log types.Log) (*ComptrollerNewVAIMintRate, error) {
	event := new(ComptrollerNewVAIMintRate)
	if err := _Comptroller.contract.UnpackLog(event, "NewVAIMintRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewVAIVaultInfoIterator is returned from FilterNewVAIVaultInfo and is used to iterate over the raw logs and unpacked data for NewVAIVaultInfo events raised by the Comptroller contract.
type ComptrollerNewVAIVaultInfoIterator struct {
	Event *ComptrollerNewVAIVaultInfo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewVAIVaultInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewVAIVaultInfo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewVAIVaultInfo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewVAIVaultInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewVAIVaultInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewVAIVaultInfo represents a NewVAIVaultInfo event raised by the Comptroller contract.
type ComptrollerNewVAIVaultInfo struct {
	Vault             common.Address
	ReleaseStartBlock *big.Int
	ReleaseInterval   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewVAIVaultInfo is a free log retrieval operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_Comptroller *ComptrollerFilterer) FilterNewVAIVaultInfo(opts *bind.FilterOpts) (*ComptrollerNewVAIVaultInfoIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewVAIVaultInfo")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewVAIVaultInfoIterator{contract: _Comptroller.contract, event: "NewVAIVaultInfo", logs: logs, sub: sub}, nil
}

// WatchNewVAIVaultInfo is a free log subscription operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_Comptroller *ComptrollerFilterer) WatchNewVAIVaultInfo(opts *bind.WatchOpts, sink chan<- *ComptrollerNewVAIVaultInfo) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewVAIVaultInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewVAIVaultInfo)
				if err := _Comptroller.contract.UnpackLog(event, "NewVAIVaultInfo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIVaultInfo is a log parse operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_Comptroller *ComptrollerFilterer) ParseNewVAIVaultInfo(log types.Log) (*ComptrollerNewVAIVaultInfo, error) {
	event := new(ComptrollerNewVAIVaultInfo)
	if err := _Comptroller.contract.UnpackLog(event, "NewVAIVaultInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewVenusVAIRateIterator is returned from FilterNewVenusVAIRate and is used to iterate over the raw logs and unpacked data for NewVenusVAIRate events raised by the Comptroller contract.
type ComptrollerNewVenusVAIRateIterator struct {
	Event *ComptrollerNewVenusVAIRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewVenusVAIRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewVenusVAIRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewVenusVAIRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewVenusVAIRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewVenusVAIRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewVenusVAIRate represents a NewVenusVAIRate event raised by the Comptroller contract.
type ComptrollerNewVenusVAIRate struct {
	OldVenusVAIRate *big.Int
	NewVenusVAIRate *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewVenusVAIRate is a free log retrieval operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_Comptroller *ComptrollerFilterer) FilterNewVenusVAIRate(opts *bind.FilterOpts) (*ComptrollerNewVenusVAIRateIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewVenusVAIRate")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewVenusVAIRateIterator{contract: _Comptroller.contract, event: "NewVenusVAIRate", logs: logs, sub: sub}, nil
}

// WatchNewVenusVAIRate is a free log subscription operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_Comptroller *ComptrollerFilterer) WatchNewVenusVAIRate(opts *bind.WatchOpts, sink chan<- *ComptrollerNewVenusVAIRate) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewVenusVAIRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewVenusVAIRate)
				if err := _Comptroller.contract.UnpackLog(event, "NewVenusVAIRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVenusVAIRate is a log parse operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_Comptroller *ComptrollerFilterer) ParseNewVenusVAIRate(log types.Log) (*ComptrollerNewVenusVAIRate, error) {
	event := new(ComptrollerNewVenusVAIRate)
	if err := _Comptroller.contract.UnpackLog(event, "NewVenusVAIRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewVenusVAIVaultRateIterator is returned from FilterNewVenusVAIVaultRate and is used to iterate over the raw logs and unpacked data for NewVenusVAIVaultRate events raised by the Comptroller contract.
type ComptrollerNewVenusVAIVaultRateIterator struct {
	Event *ComptrollerNewVenusVAIVaultRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerNewVenusVAIVaultRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewVenusVAIVaultRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerNewVenusVAIVaultRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerNewVenusVAIVaultRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewVenusVAIVaultRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewVenusVAIVaultRate represents a NewVenusVAIVaultRate event raised by the Comptroller contract.
type ComptrollerNewVenusVAIVaultRate struct {
	OldVenusVAIVaultRate *big.Int
	NewVenusVAIVaultRate *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewVenusVAIVaultRate is a free log retrieval operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_Comptroller *ComptrollerFilterer) FilterNewVenusVAIVaultRate(opts *bind.FilterOpts) (*ComptrollerNewVenusVAIVaultRateIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewVenusVAIVaultRate")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewVenusVAIVaultRateIterator{contract: _Comptroller.contract, event: "NewVenusVAIVaultRate", logs: logs, sub: sub}, nil
}

// WatchNewVenusVAIVaultRate is a free log subscription operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_Comptroller *ComptrollerFilterer) WatchNewVenusVAIVaultRate(opts *bind.WatchOpts, sink chan<- *ComptrollerNewVenusVAIVaultRate) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewVenusVAIVaultRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewVenusVAIVaultRate)
				if err := _Comptroller.contract.UnpackLog(event, "NewVenusVAIVaultRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVenusVAIVaultRate is a log parse operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_Comptroller *ComptrollerFilterer) ParseNewVenusVAIVaultRate(log types.Log) (*ComptrollerNewVenusVAIVaultRate, error) {
	event := new(ComptrollerNewVenusVAIVaultRate)
	if err := _Comptroller.contract.UnpackLog(event, "NewVenusVAIVaultRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerVenusGrantedIterator is returned from FilterVenusGranted and is used to iterate over the raw logs and unpacked data for VenusGranted events raised by the Comptroller contract.
type ComptrollerVenusGrantedIterator struct {
	Event *ComptrollerVenusGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerVenusGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerVenusGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerVenusGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerVenusGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerVenusGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerVenusGranted represents a VenusGranted event raised by the Comptroller contract.
type ComptrollerVenusGranted struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVenusGranted is a free log retrieval operation binding the contract event 0xd7fe674cac9eee3998fe3cbd7a6f93c3bc70509d97ec1550a59364be6438147e.
//
// Solidity: event VenusGranted(address recipient, uint256 amount)
func (_Comptroller *ComptrollerFilterer) FilterVenusGranted(opts *bind.FilterOpts) (*ComptrollerVenusGrantedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "VenusGranted")
	if err != nil {
		return nil, err
	}
	return &ComptrollerVenusGrantedIterator{contract: _Comptroller.contract, event: "VenusGranted", logs: logs, sub: sub}, nil
}

// WatchVenusGranted is a free log subscription operation binding the contract event 0xd7fe674cac9eee3998fe3cbd7a6f93c3bc70509d97ec1550a59364be6438147e.
//
// Solidity: event VenusGranted(address recipient, uint256 amount)
func (_Comptroller *ComptrollerFilterer) WatchVenusGranted(opts *bind.WatchOpts, sink chan<- *ComptrollerVenusGranted) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "VenusGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerVenusGranted)
				if err := _Comptroller.contract.UnpackLog(event, "VenusGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVenusGranted is a log parse operation binding the contract event 0xd7fe674cac9eee3998fe3cbd7a6f93c3bc70509d97ec1550a59364be6438147e.
//
// Solidity: event VenusGranted(address recipient, uint256 amount)
func (_Comptroller *ComptrollerFilterer) ParseVenusGranted(log types.Log) (*ComptrollerVenusGranted, error) {
	event := new(ComptrollerVenusGranted)
	if err := _Comptroller.contract.UnpackLog(event, "VenusGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerVenusSpeedUpdatedIterator is returned from FilterVenusSpeedUpdated and is used to iterate over the raw logs and unpacked data for VenusSpeedUpdated events raised by the Comptroller contract.
type ComptrollerVenusSpeedUpdatedIterator struct {
	Event *ComptrollerVenusSpeedUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComptrollerVenusSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerVenusSpeedUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComptrollerVenusSpeedUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComptrollerVenusSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerVenusSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerVenusSpeedUpdated represents a VenusSpeedUpdated event raised by the Comptroller contract.
type ComptrollerVenusSpeedUpdated struct {
	VToken   common.Address
	NewSpeed *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVenusSpeedUpdated is a free log retrieval operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_Comptroller *ComptrollerFilterer) FilterVenusSpeedUpdated(opts *bind.FilterOpts, vToken []common.Address) (*ComptrollerVenusSpeedUpdatedIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "VenusSpeedUpdated", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerVenusSpeedUpdatedIterator{contract: _Comptroller.contract, event: "VenusSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchVenusSpeedUpdated is a free log subscription operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_Comptroller *ComptrollerFilterer) WatchVenusSpeedUpdated(opts *bind.WatchOpts, sink chan<- *ComptrollerVenusSpeedUpdated, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "VenusSpeedUpdated", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerVenusSpeedUpdated)
				if err := _Comptroller.contract.UnpackLog(event, "VenusSpeedUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVenusSpeedUpdated is a log parse operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_Comptroller *ComptrollerFilterer) ParseVenusSpeedUpdated(log types.Log) (*ComptrollerVenusSpeedUpdated, error) {
	event := new(ComptrollerVenusSpeedUpdated)
	if err := _Comptroller.contract.UnpackLog(event, "VenusSpeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
