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

// VbnbMetaData contains all meta data concerning the Vbnb contract.
var VbnbMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"admin_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"MintBehalf\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"RedeemFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isVToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"contractVToken\",\"name\":\"vTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"repayBorrow\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VbnbABI is the input ABI used to generate the binding from.
// Deprecated: Use VbnbMetaData.ABI instead.
var VbnbABI = VbnbMetaData.ABI

// Vbnb is an auto generated Go binding around an Ethereum contract.
type Vbnb struct {
	VbnbCaller     // Read-only binding to the contract
	VbnbTransactor // Write-only binding to the contract
	VbnbFilterer   // Log filterer for contract events
}

// VbnbCaller is an auto generated read-only Go binding around an Ethereum contract.
type VbnbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VbnbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VbnbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VbnbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VbnbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VbnbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VbnbSession struct {
	Contract     *Vbnb             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VbnbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VbnbCallerSession struct {
	Contract *VbnbCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VbnbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VbnbTransactorSession struct {
	Contract     *VbnbTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VbnbRaw is an auto generated low-level Go binding around an Ethereum contract.
type VbnbRaw struct {
	Contract *Vbnb // Generic contract binding to access the raw methods on
}

// VbnbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VbnbCallerRaw struct {
	Contract *VbnbCaller // Generic read-only contract binding to access the raw methods on
}

// VbnbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VbnbTransactorRaw struct {
	Contract *VbnbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVbnb creates a new instance of Vbnb, bound to a specific deployed contract.
func NewVbnb(address common.Address, backend bind.ContractBackend) (*Vbnb, error) {
	contract, err := bindVbnb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vbnb{VbnbCaller: VbnbCaller{contract: contract}, VbnbTransactor: VbnbTransactor{contract: contract}, VbnbFilterer: VbnbFilterer{contract: contract}}, nil
}

// NewVbnbCaller creates a new read-only instance of Vbnb, bound to a specific deployed contract.
func NewVbnbCaller(address common.Address, caller bind.ContractCaller) (*VbnbCaller, error) {
	contract, err := bindVbnb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VbnbCaller{contract: contract}, nil
}

// NewVbnbTransactor creates a new write-only instance of Vbnb, bound to a specific deployed contract.
func NewVbnbTransactor(address common.Address, transactor bind.ContractTransactor) (*VbnbTransactor, error) {
	contract, err := bindVbnb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VbnbTransactor{contract: contract}, nil
}

// NewVbnbFilterer creates a new log filterer instance of Vbnb, bound to a specific deployed contract.
func NewVbnbFilterer(address common.Address, filterer bind.ContractFilterer) (*VbnbFilterer, error) {
	contract, err := bindVbnb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VbnbFilterer{contract: contract}, nil
}

// bindVbnb binds a generic wrapper to an already deployed contract.
func bindVbnb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VbnbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vbnb *VbnbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vbnb.Contract.VbnbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vbnb *VbnbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbnb.Contract.VbnbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vbnb *VbnbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vbnb.Contract.VbnbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vbnb *VbnbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vbnb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vbnb *VbnbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbnb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vbnb *VbnbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vbnb.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Vbnb *VbnbCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Vbnb *VbnbSession) AccrualBlockNumber() (*big.Int, error) {
	return _Vbnb.Contract.AccrualBlockNumber(&_Vbnb.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Vbnb *VbnbCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _Vbnb.Contract.AccrualBlockNumber(&_Vbnb.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vbnb *VbnbCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vbnb *VbnbSession) Admin() (common.Address, error) {
	return _Vbnb.Contract.Admin(&_Vbnb.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vbnb *VbnbCallerSession) Admin() (common.Address, error) {
	return _Vbnb.Contract.Admin(&_Vbnb.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vbnb *VbnbCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vbnb *VbnbSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Vbnb.Contract.Allowance(&_Vbnb.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vbnb *VbnbCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Vbnb.Contract.Allowance(&_Vbnb.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vbnb *VbnbCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vbnb *VbnbSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Vbnb.Contract.BalanceOf(&_Vbnb.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vbnb *VbnbCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Vbnb.Contract.BalanceOf(&_Vbnb.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Vbnb *VbnbCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Vbnb *VbnbSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Vbnb.Contract.BorrowBalanceStored(&_Vbnb.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Vbnb *VbnbCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Vbnb.Contract.BorrowBalanceStored(&_Vbnb.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Vbnb *VbnbCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Vbnb *VbnbSession) BorrowIndex() (*big.Int, error) {
	return _Vbnb.Contract.BorrowIndex(&_Vbnb.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Vbnb *VbnbCallerSession) BorrowIndex() (*big.Int, error) {
	return _Vbnb.Contract.BorrowIndex(&_Vbnb.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Vbnb *VbnbCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Vbnb *VbnbSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Vbnb.Contract.BorrowRatePerBlock(&_Vbnb.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Vbnb *VbnbCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Vbnb.Contract.BorrowRatePerBlock(&_Vbnb.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Vbnb *VbnbCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Vbnb *VbnbSession) Comptroller() (common.Address, error) {
	return _Vbnb.Contract.Comptroller(&_Vbnb.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Vbnb *VbnbCallerSession) Comptroller() (common.Address, error) {
	return _Vbnb.Contract.Comptroller(&_Vbnb.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vbnb *VbnbCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vbnb *VbnbSession) Decimals() (uint8, error) {
	return _Vbnb.Contract.Decimals(&_Vbnb.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vbnb *VbnbCallerSession) Decimals() (uint8, error) {
	return _Vbnb.Contract.Decimals(&_Vbnb.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Vbnb *VbnbCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Vbnb *VbnbSession) ExchangeRateStored() (*big.Int, error) {
	return _Vbnb.Contract.ExchangeRateStored(&_Vbnb.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Vbnb *VbnbCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Vbnb.Contract.ExchangeRateStored(&_Vbnb.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Vbnb *VbnbCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Vbnb *VbnbSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Vbnb.Contract.GetAccountSnapshot(&_Vbnb.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Vbnb *VbnbCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Vbnb.Contract.GetAccountSnapshot(&_Vbnb.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Vbnb *VbnbCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Vbnb *VbnbSession) GetCash() (*big.Int, error) {
	return _Vbnb.Contract.GetCash(&_Vbnb.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Vbnb *VbnbCallerSession) GetCash() (*big.Int, error) {
	return _Vbnb.Contract.GetCash(&_Vbnb.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Vbnb *VbnbCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Vbnb *VbnbSession) InterestRateModel() (common.Address, error) {
	return _Vbnb.Contract.InterestRateModel(&_Vbnb.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Vbnb *VbnbCallerSession) InterestRateModel() (common.Address, error) {
	return _Vbnb.Contract.InterestRateModel(&_Vbnb.CallOpts)
}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_Vbnb *VbnbCaller) IsVToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "isVToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_Vbnb *VbnbSession) IsVToken() (bool, error) {
	return _Vbnb.Contract.IsVToken(&_Vbnb.CallOpts)
}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_Vbnb *VbnbCallerSession) IsVToken() (bool, error) {
	return _Vbnb.Contract.IsVToken(&_Vbnb.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vbnb *VbnbCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vbnb *VbnbSession) Name() (string, error) {
	return _Vbnb.Contract.Name(&_Vbnb.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vbnb *VbnbCallerSession) Name() (string, error) {
	return _Vbnb.Contract.Name(&_Vbnb.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Vbnb *VbnbCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Vbnb *VbnbSession) PendingAdmin() (common.Address, error) {
	return _Vbnb.Contract.PendingAdmin(&_Vbnb.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Vbnb *VbnbCallerSession) PendingAdmin() (common.Address, error) {
	return _Vbnb.Contract.PendingAdmin(&_Vbnb.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Vbnb *VbnbCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Vbnb *VbnbSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Vbnb.Contract.ReserveFactorMantissa(&_Vbnb.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Vbnb *VbnbCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Vbnb.Contract.ReserveFactorMantissa(&_Vbnb.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Vbnb *VbnbCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Vbnb *VbnbSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Vbnb.Contract.SupplyRatePerBlock(&_Vbnb.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Vbnb *VbnbCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Vbnb.Contract.SupplyRatePerBlock(&_Vbnb.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vbnb *VbnbCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vbnb *VbnbSession) Symbol() (string, error) {
	return _Vbnb.Contract.Symbol(&_Vbnb.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vbnb *VbnbCallerSession) Symbol() (string, error) {
	return _Vbnb.Contract.Symbol(&_Vbnb.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Vbnb *VbnbCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Vbnb *VbnbSession) TotalBorrows() (*big.Int, error) {
	return _Vbnb.Contract.TotalBorrows(&_Vbnb.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Vbnb *VbnbCallerSession) TotalBorrows() (*big.Int, error) {
	return _Vbnb.Contract.TotalBorrows(&_Vbnb.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Vbnb *VbnbCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Vbnb *VbnbSession) TotalReserves() (*big.Int, error) {
	return _Vbnb.Contract.TotalReserves(&_Vbnb.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Vbnb *VbnbCallerSession) TotalReserves() (*big.Int, error) {
	return _Vbnb.Contract.TotalReserves(&_Vbnb.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vbnb *VbnbCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vbnb.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vbnb *VbnbSession) TotalSupply() (*big.Int, error) {
	return _Vbnb.Contract.TotalSupply(&_Vbnb.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vbnb *VbnbCallerSession) TotalSupply() (*big.Int, error) {
	return _Vbnb.Contract.TotalSupply(&_Vbnb.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Vbnb *VbnbTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Vbnb *VbnbSession) AcceptAdmin() (*types.Transaction, error) {
	return _Vbnb.Contract.AcceptAdmin(&_Vbnb.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Vbnb *VbnbTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Vbnb.Contract.AcceptAdmin(&_Vbnb.TransactOpts)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Vbnb *VbnbTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Vbnb *VbnbSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.ReduceReserves(&_Vbnb.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Vbnb *VbnbTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.ReduceReserves(&_Vbnb.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Vbnb *VbnbTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Vbnb *VbnbSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.SetComptroller(&_Vbnb.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Vbnb *VbnbTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.SetComptroller(&_Vbnb.TransactOpts, newComptroller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Vbnb *VbnbTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Vbnb *VbnbSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.SetInterestRateModel(&_Vbnb.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Vbnb *VbnbTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.SetInterestRateModel(&_Vbnb.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Vbnb *VbnbTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Vbnb *VbnbSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.SetPendingAdmin(&_Vbnb.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Vbnb *VbnbTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.SetPendingAdmin(&_Vbnb.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Vbnb *VbnbTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Vbnb *VbnbSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.SetReserveFactor(&_Vbnb.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Vbnb *VbnbTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.SetReserveFactor(&_Vbnb.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Vbnb *VbnbTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Vbnb *VbnbSession) AccrueInterest() (*types.Transaction, error) {
	return _Vbnb.Contract.AccrueInterest(&_Vbnb.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Vbnb *VbnbTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Vbnb.Contract.AccrueInterest(&_Vbnb.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vbnb *VbnbTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vbnb *VbnbSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.Approve(&_Vbnb.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vbnb *VbnbTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.Approve(&_Vbnb.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Vbnb *VbnbTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Vbnb *VbnbSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.BalanceOfUnderlying(&_Vbnb.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Vbnb *VbnbTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.BalanceOfUnderlying(&_Vbnb.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Vbnb *VbnbTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Vbnb *VbnbSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.Borrow(&_Vbnb.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Vbnb *VbnbTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.Borrow(&_Vbnb.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Vbnb *VbnbTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Vbnb *VbnbSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.BorrowBalanceCurrent(&_Vbnb.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Vbnb *VbnbTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.BorrowBalanceCurrent(&_Vbnb.TransactOpts, account)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Vbnb *VbnbTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Vbnb *VbnbSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Vbnb.Contract.ExchangeRateCurrent(&_Vbnb.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Vbnb *VbnbTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Vbnb.Contract.ExchangeRateCurrent(&_Vbnb.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Vbnb *VbnbTransactor) Initialize(opts *bind.TransactOpts, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "initialize", comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Vbnb *VbnbSession) Initialize(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Vbnb.Contract.Initialize(&_Vbnb.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Vbnb *VbnbTransactorSession) Initialize(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Vbnb.Contract.Initialize(&_Vbnb.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xaae40a2a.
//
// Solidity: function liquidateBorrow(address borrower, address vTokenCollateral) payable returns()
func (_Vbnb *VbnbTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "liquidateBorrow", borrower, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xaae40a2a.
//
// Solidity: function liquidateBorrow(address borrower, address vTokenCollateral) payable returns()
func (_Vbnb *VbnbSession) LiquidateBorrow(borrower common.Address, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.LiquidateBorrow(&_Vbnb.TransactOpts, borrower, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xaae40a2a.
//
// Solidity: function liquidateBorrow(address borrower, address vTokenCollateral) payable returns()
func (_Vbnb *VbnbTransactorSession) LiquidateBorrow(borrower common.Address, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.LiquidateBorrow(&_Vbnb.TransactOpts, borrower, vTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Vbnb *VbnbTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Vbnb *VbnbSession) Mint() (*types.Transaction, error) {
	return _Vbnb.Contract.Mint(&_Vbnb.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Vbnb *VbnbTransactorSession) Mint() (*types.Transaction, error) {
	return _Vbnb.Contract.Mint(&_Vbnb.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Vbnb *VbnbTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Vbnb *VbnbSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.Redeem(&_Vbnb.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Vbnb *VbnbTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.Redeem(&_Vbnb.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Vbnb *VbnbTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Vbnb *VbnbSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.RedeemUnderlying(&_Vbnb.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Vbnb *VbnbTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.RedeemUnderlying(&_Vbnb.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x4e4d9fea.
//
// Solidity: function repayBorrow() payable returns()
func (_Vbnb *VbnbTransactor) RepayBorrow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "repayBorrow")
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x4e4d9fea.
//
// Solidity: function repayBorrow() payable returns()
func (_Vbnb *VbnbSession) RepayBorrow() (*types.Transaction, error) {
	return _Vbnb.Contract.RepayBorrow(&_Vbnb.TransactOpts)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x4e4d9fea.
//
// Solidity: function repayBorrow() payable returns()
func (_Vbnb *VbnbTransactorSession) RepayBorrow() (*types.Transaction, error) {
	return _Vbnb.Contract.RepayBorrow(&_Vbnb.TransactOpts)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xe5974619.
//
// Solidity: function repayBorrowBehalf(address borrower) payable returns()
func (_Vbnb *VbnbTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "repayBorrowBehalf", borrower)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xe5974619.
//
// Solidity: function repayBorrowBehalf(address borrower) payable returns()
func (_Vbnb *VbnbSession) RepayBorrowBehalf(borrower common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.RepayBorrowBehalf(&_Vbnb.TransactOpts, borrower)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xe5974619.
//
// Solidity: function repayBorrowBehalf(address borrower) payable returns()
func (_Vbnb *VbnbTransactorSession) RepayBorrowBehalf(borrower common.Address) (*types.Transaction, error) {
	return _Vbnb.Contract.RepayBorrowBehalf(&_Vbnb.TransactOpts, borrower)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Vbnb *VbnbTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Vbnb *VbnbSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.Seize(&_Vbnb.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Vbnb *VbnbTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.Seize(&_Vbnb.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Vbnb *VbnbTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Vbnb *VbnbSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Vbnb.Contract.TotalBorrowsCurrent(&_Vbnb.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Vbnb *VbnbTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Vbnb.Contract.TotalBorrowsCurrent(&_Vbnb.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Vbnb *VbnbTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Vbnb *VbnbSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.Transfer(&_Vbnb.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Vbnb *VbnbTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.Transfer(&_Vbnb.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Vbnb *VbnbTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbnb.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Vbnb *VbnbSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.TransferFrom(&_Vbnb.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Vbnb *VbnbTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vbnb.Contract.TransferFrom(&_Vbnb.TransactOpts, src, dst, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vbnb *VbnbTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Vbnb.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vbnb *VbnbSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Vbnb.Contract.Fallback(&_Vbnb.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vbnb *VbnbTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Vbnb.Contract.Fallback(&_Vbnb.TransactOpts, calldata)
}

// VbnbAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Vbnb contract.
type VbnbAccrueInterestIterator struct {
	Event *VbnbAccrueInterest // Event containing the contract specifics and raw log

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
func (it *VbnbAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbAccrueInterest)
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
		it.Event = new(VbnbAccrueInterest)
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
func (it *VbnbAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbAccrueInterest represents a AccrueInterest event raised by the Vbnb contract.
type VbnbAccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Vbnb *VbnbFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*VbnbAccrueInterestIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &VbnbAccrueInterestIterator{contract: _Vbnb.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Vbnb *VbnbFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *VbnbAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbAccrueInterest)
				if err := _Vbnb.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Vbnb *VbnbFilterer) ParseAccrueInterest(log types.Log) (*VbnbAccrueInterest, error) {
	event := new(VbnbAccrueInterest)
	if err := _Vbnb.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vbnb contract.
type VbnbApprovalIterator struct {
	Event *VbnbApproval // Event containing the contract specifics and raw log

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
func (it *VbnbApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbApproval)
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
		it.Event = new(VbnbApproval)
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
func (it *VbnbApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbApproval represents a Approval event raised by the Vbnb contract.
type VbnbApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Vbnb *VbnbFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VbnbApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VbnbApprovalIterator{contract: _Vbnb.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Vbnb *VbnbFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VbnbApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbApproval)
				if err := _Vbnb.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Vbnb *VbnbFilterer) ParseApproval(log types.Log) (*VbnbApproval, error) {
	event := new(VbnbApproval)
	if err := _Vbnb.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Vbnb contract.
type VbnbBorrowIterator struct {
	Event *VbnbBorrow // Event containing the contract specifics and raw log

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
func (it *VbnbBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbBorrow)
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
		it.Event = new(VbnbBorrow)
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
func (it *VbnbBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbBorrow represents a Borrow event raised by the Vbnb contract.
type VbnbBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbnb *VbnbFilterer) FilterBorrow(opts *bind.FilterOpts) (*VbnbBorrowIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &VbnbBorrowIterator{contract: _Vbnb.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbnb *VbnbFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *VbnbBorrow) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbBorrow)
				if err := _Vbnb.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbnb *VbnbFilterer) ParseBorrow(log types.Log) (*VbnbBorrow, error) {
	event := new(VbnbBorrow)
	if err := _Vbnb.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Vbnb contract.
type VbnbFailureIterator struct {
	Event *VbnbFailure // Event containing the contract specifics and raw log

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
func (it *VbnbFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbFailure)
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
		it.Event = new(VbnbFailure)
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
func (it *VbnbFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbFailure represents a Failure event raised by the Vbnb contract.
type VbnbFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Vbnb *VbnbFilterer) FilterFailure(opts *bind.FilterOpts) (*VbnbFailureIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &VbnbFailureIterator{contract: _Vbnb.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Vbnb *VbnbFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *VbnbFailure) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbFailure)
				if err := _Vbnb.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_Vbnb *VbnbFilterer) ParseFailure(log types.Log) (*VbnbFailure, error) {
	event := new(VbnbFailure)
	if err := _Vbnb.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the Vbnb contract.
type VbnbLiquidateBorrowIterator struct {
	Event *VbnbLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *VbnbLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbLiquidateBorrow)
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
		it.Event = new(VbnbLiquidateBorrow)
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
func (it *VbnbLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbLiquidateBorrow represents a LiquidateBorrow event raised by the Vbnb contract.
type VbnbLiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	VTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_Vbnb *VbnbFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*VbnbLiquidateBorrowIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &VbnbLiquidateBorrowIterator{contract: _Vbnb.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_Vbnb *VbnbFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *VbnbLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbLiquidateBorrow)
				if err := _Vbnb.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_Vbnb *VbnbFilterer) ParseLiquidateBorrow(log types.Log) (*VbnbLiquidateBorrow, error) {
	event := new(VbnbLiquidateBorrow)
	if err := _Vbnb.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Vbnb contract.
type VbnbMintIterator struct {
	Event *VbnbMint // Event containing the contract specifics and raw log

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
func (it *VbnbMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbMint)
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
		it.Event = new(VbnbMint)
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
func (it *VbnbMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbMint represents a Mint event raised by the Vbnb contract.
type VbnbMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Vbnb *VbnbFilterer) FilterMint(opts *bind.FilterOpts) (*VbnbMintIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &VbnbMintIterator{contract: _Vbnb.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Vbnb *VbnbFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *VbnbMint) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbMint)
				if err := _Vbnb.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Vbnb *VbnbFilterer) ParseMint(log types.Log) (*VbnbMint, error) {
	event := new(VbnbMint)
	if err := _Vbnb.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbMintBehalfIterator is returned from FilterMintBehalf and is used to iterate over the raw logs and unpacked data for MintBehalf events raised by the Vbnb contract.
type VbnbMintBehalfIterator struct {
	Event *VbnbMintBehalf // Event containing the contract specifics and raw log

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
func (it *VbnbMintBehalfIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbMintBehalf)
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
		it.Event = new(VbnbMintBehalf)
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
func (it *VbnbMintBehalfIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbMintBehalfIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbMintBehalf represents a MintBehalf event raised by the Vbnb contract.
type VbnbMintBehalf struct {
	Payer      common.Address
	Receiver   common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintBehalf is a free log retrieval operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_Vbnb *VbnbFilterer) FilterMintBehalf(opts *bind.FilterOpts) (*VbnbMintBehalfIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "MintBehalf")
	if err != nil {
		return nil, err
	}
	return &VbnbMintBehalfIterator{contract: _Vbnb.contract, event: "MintBehalf", logs: logs, sub: sub}, nil
}

// WatchMintBehalf is a free log subscription operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_Vbnb *VbnbFilterer) WatchMintBehalf(opts *bind.WatchOpts, sink chan<- *VbnbMintBehalf) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "MintBehalf")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbMintBehalf)
				if err := _Vbnb.contract.UnpackLog(event, "MintBehalf", log); err != nil {
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

// ParseMintBehalf is a log parse operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_Vbnb *VbnbFilterer) ParseMintBehalf(log types.Log) (*VbnbMintBehalf, error) {
	event := new(VbnbMintBehalf)
	if err := _Vbnb.contract.UnpackLog(event, "MintBehalf", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Vbnb contract.
type VbnbNewAdminIterator struct {
	Event *VbnbNewAdmin // Event containing the contract specifics and raw log

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
func (it *VbnbNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbNewAdmin)
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
		it.Event = new(VbnbNewAdmin)
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
func (it *VbnbNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbNewAdmin represents a NewAdmin event raised by the Vbnb contract.
type VbnbNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Vbnb *VbnbFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*VbnbNewAdminIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &VbnbNewAdminIterator{contract: _Vbnb.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Vbnb *VbnbFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *VbnbNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbNewAdmin)
				if err := _Vbnb.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Vbnb *VbnbFilterer) ParseNewAdmin(log types.Log) (*VbnbNewAdmin, error) {
	event := new(VbnbNewAdmin)
	if err := _Vbnb.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the Vbnb contract.
type VbnbNewComptrollerIterator struct {
	Event *VbnbNewComptroller // Event containing the contract specifics and raw log

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
func (it *VbnbNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbNewComptroller)
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
		it.Event = new(VbnbNewComptroller)
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
func (it *VbnbNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbNewComptroller represents a NewComptroller event raised by the Vbnb contract.
type VbnbNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Vbnb *VbnbFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*VbnbNewComptrollerIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &VbnbNewComptrollerIterator{contract: _Vbnb.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Vbnb *VbnbFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *VbnbNewComptroller) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbNewComptroller)
				if err := _Vbnb.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Vbnb *VbnbFilterer) ParseNewComptroller(log types.Log) (*VbnbNewComptroller, error) {
	event := new(VbnbNewComptroller)
	if err := _Vbnb.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the Vbnb contract.
type VbnbNewMarketInterestRateModelIterator struct {
	Event *VbnbNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *VbnbNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbNewMarketInterestRateModel)
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
		it.Event = new(VbnbNewMarketInterestRateModel)
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
func (it *VbnbNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the Vbnb contract.
type VbnbNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Vbnb *VbnbFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*VbnbNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &VbnbNewMarketInterestRateModelIterator{contract: _Vbnb.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Vbnb *VbnbFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *VbnbNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbNewMarketInterestRateModel)
				if err := _Vbnb.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Vbnb *VbnbFilterer) ParseNewMarketInterestRateModel(log types.Log) (*VbnbNewMarketInterestRateModel, error) {
	event := new(VbnbNewMarketInterestRateModel)
	if err := _Vbnb.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Vbnb contract.
type VbnbNewPendingAdminIterator struct {
	Event *VbnbNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *VbnbNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbNewPendingAdmin)
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
		it.Event = new(VbnbNewPendingAdmin)
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
func (it *VbnbNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbNewPendingAdmin represents a NewPendingAdmin event raised by the Vbnb contract.
type VbnbNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Vbnb *VbnbFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*VbnbNewPendingAdminIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &VbnbNewPendingAdminIterator{contract: _Vbnb.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Vbnb *VbnbFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *VbnbNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbNewPendingAdmin)
				if err := _Vbnb.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Vbnb *VbnbFilterer) ParseNewPendingAdmin(log types.Log) (*VbnbNewPendingAdmin, error) {
	event := new(VbnbNewPendingAdmin)
	if err := _Vbnb.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the Vbnb contract.
type VbnbNewReserveFactorIterator struct {
	Event *VbnbNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *VbnbNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbNewReserveFactor)
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
		it.Event = new(VbnbNewReserveFactor)
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
func (it *VbnbNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbNewReserveFactor represents a NewReserveFactor event raised by the Vbnb contract.
type VbnbNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Vbnb *VbnbFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*VbnbNewReserveFactorIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &VbnbNewReserveFactorIterator{contract: _Vbnb.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Vbnb *VbnbFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *VbnbNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbNewReserveFactor)
				if err := _Vbnb.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Vbnb *VbnbFilterer) ParseNewReserveFactor(log types.Log) (*VbnbNewReserveFactor, error) {
	event := new(VbnbNewReserveFactor)
	if err := _Vbnb.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Vbnb contract.
type VbnbRedeemIterator struct {
	Event *VbnbRedeem // Event containing the contract specifics and raw log

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
func (it *VbnbRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbRedeem)
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
		it.Event = new(VbnbRedeem)
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
func (it *VbnbRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbRedeem represents a Redeem event raised by the Vbnb contract.
type VbnbRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Vbnb *VbnbFilterer) FilterRedeem(opts *bind.FilterOpts) (*VbnbRedeemIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &VbnbRedeemIterator{contract: _Vbnb.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Vbnb *VbnbFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *VbnbRedeem) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbRedeem)
				if err := _Vbnb.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Vbnb *VbnbFilterer) ParseRedeem(log types.Log) (*VbnbRedeem, error) {
	event := new(VbnbRedeem)
	if err := _Vbnb.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbRedeemFeeIterator is returned from FilterRedeemFee and is used to iterate over the raw logs and unpacked data for RedeemFee events raised by the Vbnb contract.
type VbnbRedeemFeeIterator struct {
	Event *VbnbRedeemFee // Event containing the contract specifics and raw log

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
func (it *VbnbRedeemFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbRedeemFee)
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
		it.Event = new(VbnbRedeemFee)
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
func (it *VbnbRedeemFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbRedeemFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbRedeemFee represents a RedeemFee event raised by the Vbnb contract.
type VbnbRedeemFee struct {
	Redeemer     common.Address
	FeeAmount    *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeemFee is a free log retrieval operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_Vbnb *VbnbFilterer) FilterRedeemFee(opts *bind.FilterOpts) (*VbnbRedeemFeeIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "RedeemFee")
	if err != nil {
		return nil, err
	}
	return &VbnbRedeemFeeIterator{contract: _Vbnb.contract, event: "RedeemFee", logs: logs, sub: sub}, nil
}

// WatchRedeemFee is a free log subscription operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_Vbnb *VbnbFilterer) WatchRedeemFee(opts *bind.WatchOpts, sink chan<- *VbnbRedeemFee) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "RedeemFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbRedeemFee)
				if err := _Vbnb.contract.UnpackLog(event, "RedeemFee", log); err != nil {
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

// ParseRedeemFee is a log parse operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_Vbnb *VbnbFilterer) ParseRedeemFee(log types.Log) (*VbnbRedeemFee, error) {
	event := new(VbnbRedeemFee)
	if err := _Vbnb.contract.UnpackLog(event, "RedeemFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the Vbnb contract.
type VbnbRepayBorrowIterator struct {
	Event *VbnbRepayBorrow // Event containing the contract specifics and raw log

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
func (it *VbnbRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbRepayBorrow)
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
		it.Event = new(VbnbRepayBorrow)
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
func (it *VbnbRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbRepayBorrow represents a RepayBorrow event raised by the Vbnb contract.
type VbnbRepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbnb *VbnbFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*VbnbRepayBorrowIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &VbnbRepayBorrowIterator{contract: _Vbnb.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbnb *VbnbFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *VbnbRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbRepayBorrow)
				if err := _Vbnb.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Vbnb *VbnbFilterer) ParseRepayBorrow(log types.Log) (*VbnbRepayBorrow, error) {
	event := new(VbnbRepayBorrow)
	if err := _Vbnb.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the Vbnb contract.
type VbnbReservesAddedIterator struct {
	Event *VbnbReservesAdded // Event containing the contract specifics and raw log

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
func (it *VbnbReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbReservesAdded)
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
		it.Event = new(VbnbReservesAdded)
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
func (it *VbnbReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbReservesAdded represents a ReservesAdded event raised by the Vbnb contract.
type VbnbReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Vbnb *VbnbFilterer) FilterReservesAdded(opts *bind.FilterOpts) (*VbnbReservesAddedIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &VbnbReservesAddedIterator{contract: _Vbnb.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Vbnb *VbnbFilterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *VbnbReservesAdded) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbReservesAdded)
				if err := _Vbnb.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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

// ParseReservesAdded is a log parse operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Vbnb *VbnbFilterer) ParseReservesAdded(log types.Log) (*VbnbReservesAdded, error) {
	event := new(VbnbReservesAdded)
	if err := _Vbnb.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the Vbnb contract.
type VbnbReservesReducedIterator struct {
	Event *VbnbReservesReduced // Event containing the contract specifics and raw log

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
func (it *VbnbReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbReservesReduced)
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
		it.Event = new(VbnbReservesReduced)
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
func (it *VbnbReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbReservesReduced represents a ReservesReduced event raised by the Vbnb contract.
type VbnbReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Vbnb *VbnbFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*VbnbReservesReducedIterator, error) {

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &VbnbReservesReducedIterator{contract: _Vbnb.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Vbnb *VbnbFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *VbnbReservesReduced) (event.Subscription, error) {

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbReservesReduced)
				if err := _Vbnb.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Vbnb *VbnbFilterer) ParseReservesReduced(log types.Log) (*VbnbReservesReduced, error) {
	event := new(VbnbReservesReduced)
	if err := _Vbnb.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VbnbTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vbnb contract.
type VbnbTransferIterator struct {
	Event *VbnbTransfer // Event containing the contract specifics and raw log

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
func (it *VbnbTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VbnbTransfer)
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
		it.Event = new(VbnbTransfer)
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
func (it *VbnbTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VbnbTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VbnbTransfer represents a Transfer event raised by the Vbnb contract.
type VbnbTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Vbnb *VbnbFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VbnbTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vbnb.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VbnbTransferIterator{contract: _Vbnb.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Vbnb *VbnbFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VbnbTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vbnb.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VbnbTransfer)
				if err := _Vbnb.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Vbnb *VbnbFilterer) ParseTransfer(log types.Log) (*VbnbTransfer, error) {
	event := new(VbnbTransfer)
	if err := _Vbnb.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
