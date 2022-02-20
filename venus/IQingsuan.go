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

// IQingsuanMetaData contains all meta data concerning the IQingsuan contract.
var IQingsuanMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_situation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_flashLoanFrom\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_path1\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_path2\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_flashLoanAmount\",\"type\":\"uint256\"}],\"name\":\"qingsuan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IQingsuanABI is the input ABI used to generate the binding from.
// Deprecated: Use IQingsuanMetaData.ABI instead.
var IQingsuanABI = IQingsuanMetaData.ABI

// IQingsuan is an auto generated Go binding around an Ethereum contract.
type IQingsuan struct {
	IQingsuanCaller     // Read-only binding to the contract
	IQingsuanTransactor // Write-only binding to the contract
	IQingsuanFilterer   // Log filterer for contract events
}

// IQingsuanCaller is an auto generated read-only Go binding around an Ethereum contract.
type IQingsuanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQingsuanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IQingsuanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQingsuanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IQingsuanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQingsuanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IQingsuanSession struct {
	Contract     *IQingsuan        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IQingsuanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IQingsuanCallerSession struct {
	Contract *IQingsuanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IQingsuanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IQingsuanTransactorSession struct {
	Contract     *IQingsuanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IQingsuanRaw is an auto generated low-level Go binding around an Ethereum contract.
type IQingsuanRaw struct {
	Contract *IQingsuan // Generic contract binding to access the raw methods on
}

// IQingsuanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IQingsuanCallerRaw struct {
	Contract *IQingsuanCaller // Generic read-only contract binding to access the raw methods on
}

// IQingsuanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IQingsuanTransactorRaw struct {
	Contract *IQingsuanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIQingsuan creates a new instance of IQingsuan, bound to a specific deployed contract.
func NewIQingsuan(address common.Address, backend bind.ContractBackend) (*IQingsuan, error) {
	contract, err := bindIQingsuan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IQingsuan{IQingsuanCaller: IQingsuanCaller{contract: contract}, IQingsuanTransactor: IQingsuanTransactor{contract: contract}, IQingsuanFilterer: IQingsuanFilterer{contract: contract}}, nil
}

// NewIQingsuanCaller creates a new read-only instance of IQingsuan, bound to a specific deployed contract.
func NewIQingsuanCaller(address common.Address, caller bind.ContractCaller) (*IQingsuanCaller, error) {
	contract, err := bindIQingsuan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IQingsuanCaller{contract: contract}, nil
}

// NewIQingsuanTransactor creates a new write-only instance of IQingsuan, bound to a specific deployed contract.
func NewIQingsuanTransactor(address common.Address, transactor bind.ContractTransactor) (*IQingsuanTransactor, error) {
	contract, err := bindIQingsuan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IQingsuanTransactor{contract: contract}, nil
}

// NewIQingsuanFilterer creates a new log filterer instance of IQingsuan, bound to a specific deployed contract.
func NewIQingsuanFilterer(address common.Address, filterer bind.ContractFilterer) (*IQingsuanFilterer, error) {
	contract, err := bindIQingsuan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IQingsuanFilterer{contract: contract}, nil
}

// bindIQingsuan binds a generic wrapper to an already deployed contract.
func bindIQingsuan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IQingsuanABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQingsuan *IQingsuanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQingsuan.Contract.IQingsuanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQingsuan *IQingsuanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQingsuan.Contract.IQingsuanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQingsuan *IQingsuanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQingsuan.Contract.IQingsuanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQingsuan *IQingsuanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQingsuan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQingsuan *IQingsuanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQingsuan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQingsuan *IQingsuanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQingsuan.Contract.contract.Transact(opts, method, params...)
}

// Qingsuan is a paid mutator transaction binding the contract method 0xb14e9a72.
//
// Solidity: function qingsuan(uint256 _situation, address _flashLoanFrom, address[] _path1, address[] _path2, address[] _tokens, uint256 _flashLoanAmount) returns()
func (_IQingsuan *IQingsuanTransactor) Qingsuan(opts *bind.TransactOpts, _situation *big.Int, _flashLoanFrom common.Address, _path1 []common.Address, _path2 []common.Address, _tokens []common.Address, _flashLoanAmount *big.Int) (*types.Transaction, error) {
	return _IQingsuan.contract.Transact(opts, "qingsuan", _situation, _flashLoanFrom, _path1, _path2, _tokens, _flashLoanAmount)
}

// Qingsuan is a paid mutator transaction binding the contract method 0xb14e9a72.
//
// Solidity: function qingsuan(uint256 _situation, address _flashLoanFrom, address[] _path1, address[] _path2, address[] _tokens, uint256 _flashLoanAmount) returns()
func (_IQingsuan *IQingsuanSession) Qingsuan(_situation *big.Int, _flashLoanFrom common.Address, _path1 []common.Address, _path2 []common.Address, _tokens []common.Address, _flashLoanAmount *big.Int) (*types.Transaction, error) {
	return _IQingsuan.Contract.Qingsuan(&_IQingsuan.TransactOpts, _situation, _flashLoanFrom, _path1, _path2, _tokens, _flashLoanAmount)
}

// Qingsuan is a paid mutator transaction binding the contract method 0xb14e9a72.
//
// Solidity: function qingsuan(uint256 _situation, address _flashLoanFrom, address[] _path1, address[] _path2, address[] _tokens, uint256 _flashLoanAmount) returns()
func (_IQingsuan *IQingsuanTransactorSession) Qingsuan(_situation *big.Int, _flashLoanFrom common.Address, _path1 []common.Address, _path2 []common.Address, _tokens []common.Address, _flashLoanAmount *big.Int) (*types.Transaction, error) {
	return _IQingsuan.Contract.Qingsuan(&_IQingsuan.TransactOpts, _situation, _flashLoanFrom, _path1, _path2, _tokens, _flashLoanAmount)
}
