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

// ILiquidateMetaData contains all meta data concerning the ILiquidate contract.
var ILiquidateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_scenarioNo\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_flashLoanVToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_flashLoanUnderlyingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_flashLoanFrom\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_flashLoanAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seizedVToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_flashLoanReturnAmont\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expectedProfit\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ILiquidateABI is the input ABI used to generate the binding from.
// Deprecated: Use ILiquidateMetaData.ABI instead.
var ILiquidateABI = ILiquidateMetaData.ABI

// ILiquidate is an auto generated Go binding around an Ethereum contract.
type ILiquidate struct {
	ILiquidateCaller     // Read-only binding to the contract
	ILiquidateTransactor // Write-only binding to the contract
	ILiquidateFilterer   // Log filterer for contract events
}

// ILiquidateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILiquidateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILiquidateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILiquidateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILiquidateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILiquidateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILiquidateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILiquidateSession struct {
	Contract     *ILiquidate       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILiquidateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILiquidateCallerSession struct {
	Contract *ILiquidateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ILiquidateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILiquidateTransactorSession struct {
	Contract     *ILiquidateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ILiquidateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILiquidateRaw struct {
	Contract *ILiquidate // Generic contract binding to access the raw methods on
}

// ILiquidateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILiquidateCallerRaw struct {
	Contract *ILiquidateCaller // Generic read-only contract binding to access the raw methods on
}

// ILiquidateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILiquidateTransactorRaw struct {
	Contract *ILiquidateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILiquidate creates a new instance of ILiquidate, bound to a specific deployed contract.
func NewILiquidate(address common.Address, backend bind.ContractBackend) (*ILiquidate, error) {
	contract, err := bindILiquidate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILiquidate{ILiquidateCaller: ILiquidateCaller{contract: contract}, ILiquidateTransactor: ILiquidateTransactor{contract: contract}, ILiquidateFilterer: ILiquidateFilterer{contract: contract}}, nil
}

// NewILiquidateCaller creates a new read-only instance of ILiquidate, bound to a specific deployed contract.
func NewILiquidateCaller(address common.Address, caller bind.ContractCaller) (*ILiquidateCaller, error) {
	contract, err := bindILiquidate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILiquidateCaller{contract: contract}, nil
}

// NewILiquidateTransactor creates a new write-only instance of ILiquidate, bound to a specific deployed contract.
func NewILiquidateTransactor(address common.Address, transactor bind.ContractTransactor) (*ILiquidateTransactor, error) {
	contract, err := bindILiquidate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILiquidateTransactor{contract: contract}, nil
}

// NewILiquidateFilterer creates a new log filterer instance of ILiquidate, bound to a specific deployed contract.
func NewILiquidateFilterer(address common.Address, filterer bind.ContractFilterer) (*ILiquidateFilterer, error) {
	contract, err := bindILiquidate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILiquidateFilterer{contract: contract}, nil
}

// bindILiquidate binds a generic wrapper to an already deployed contract.
func bindILiquidate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILiquidateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILiquidate *ILiquidateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILiquidate.Contract.ILiquidateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILiquidate *ILiquidateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILiquidate.Contract.ILiquidateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILiquidate *ILiquidateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILiquidate.Contract.ILiquidateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILiquidate *ILiquidateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILiquidate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILiquidate *ILiquidateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILiquidate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILiquidate *ILiquidateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILiquidate.Contract.contract.Transact(opts, method, params...)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd5244545.
//
// Solidity: function liquidate(uint256 _scenarioNo, address _flashLoanVToken, address _flashLoanUnderlyingToken, address _flashLoanFrom, uint256 _flashLoanAmount, address _seizedVToken, uint256 _flashLoanReturnAmont, uint256 _expectedProfit) returns()
func (_ILiquidate *ILiquidateTransactor) Liquidate(opts *bind.TransactOpts, _scenarioNo *big.Int, _flashLoanVToken common.Address, _flashLoanUnderlyingToken common.Address, _flashLoanFrom common.Address, _flashLoanAmount *big.Int, _seizedVToken common.Address, _flashLoanReturnAmont *big.Int, _expectedProfit *big.Int) (*types.Transaction, error) {
	return _ILiquidate.contract.Transact(opts, "liquidate", _scenarioNo, _flashLoanVToken, _flashLoanUnderlyingToken, _flashLoanFrom, _flashLoanAmount, _seizedVToken, _flashLoanReturnAmont, _expectedProfit)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd5244545.
//
// Solidity: function liquidate(uint256 _scenarioNo, address _flashLoanVToken, address _flashLoanUnderlyingToken, address _flashLoanFrom, uint256 _flashLoanAmount, address _seizedVToken, uint256 _flashLoanReturnAmont, uint256 _expectedProfit) returns()
func (_ILiquidate *ILiquidateSession) Liquidate(_scenarioNo *big.Int, _flashLoanVToken common.Address, _flashLoanUnderlyingToken common.Address, _flashLoanFrom common.Address, _flashLoanAmount *big.Int, _seizedVToken common.Address, _flashLoanReturnAmont *big.Int, _expectedProfit *big.Int) (*types.Transaction, error) {
	return _ILiquidate.Contract.Liquidate(&_ILiquidate.TransactOpts, _scenarioNo, _flashLoanVToken, _flashLoanUnderlyingToken, _flashLoanFrom, _flashLoanAmount, _seizedVToken, _flashLoanReturnAmont, _expectedProfit)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd5244545.
//
// Solidity: function liquidate(uint256 _scenarioNo, address _flashLoanVToken, address _flashLoanUnderlyingToken, address _flashLoanFrom, uint256 _flashLoanAmount, address _seizedVToken, uint256 _flashLoanReturnAmont, uint256 _expectedProfit) returns()
func (_ILiquidate *ILiquidateTransactorSession) Liquidate(_scenarioNo *big.Int, _flashLoanVToken common.Address, _flashLoanUnderlyingToken common.Address, _flashLoanFrom common.Address, _flashLoanAmount *big.Int, _seizedVToken common.Address, _flashLoanReturnAmont *big.Int, _expectedProfit *big.Int) (*types.Transaction, error) {
	return _ILiquidate.Contract.Liquidate(&_ILiquidate.TransactOpts, _scenarioNo, _flashLoanVToken, _flashLoanUnderlyingToken, _flashLoanFrom, _flashLoanAmount, _seizedVToken, _flashLoanReturnAmont, _expectedProfit)
}
