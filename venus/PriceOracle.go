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

// PriceOracleMetaData contains all meta data concerning the PriceOracle contract.
var PriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"FeedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"PricePosted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"VAI_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"assetPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPriceOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setDirectPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"setFeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceOracleMetaData.ABI instead.
var PriceOracleABI = PriceOracleMetaData.ABI

// PriceOracle is an auto generated Go binding around an Ethereum contract.
type PriceOracle struct {
	PriceOracleCaller     // Read-only binding to the contract
	PriceOracleTransactor // Write-only binding to the contract
	PriceOracleFilterer   // Log filterer for contract events
}

// PriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceOracleSession struct {
	Contract     *PriceOracle      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceOracleCallerSession struct {
	Contract *PriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceOracleTransactorSession struct {
	Contract     *PriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceOracleRaw struct {
	Contract *PriceOracle // Generic contract binding to access the raw methods on
}

// PriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceOracleCallerRaw struct {
	Contract *PriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// PriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceOracleTransactorRaw struct {
	Contract *PriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceOracle creates a new instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracle(address common.Address, backend bind.ContractBackend) (*PriceOracle, error) {
	contract, err := bindPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceOracle{PriceOracleCaller: PriceOracleCaller{contract: contract}, PriceOracleTransactor: PriceOracleTransactor{contract: contract}, PriceOracleFilterer: PriceOracleFilterer{contract: contract}}, nil
}

// NewPriceOracleCaller creates a new read-only instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*PriceOracleCaller, error) {
	contract, err := bindPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleCaller{contract: contract}, nil
}

// NewPriceOracleTransactor creates a new write-only instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceOracleTransactor, error) {
	contract, err := bindPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleTransactor{contract: contract}, nil
}

// NewPriceOracleFilterer creates a new log filterer instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceOracleFilterer, error) {
	contract, err := bindPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceOracleFilterer{contract: contract}, nil
}

// bindPriceOracle binds a generic wrapper to an already deployed contract.
func bindPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracle *PriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOracle.Contract.PriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracle *PriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.Contract.PriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracle *PriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracle.Contract.PriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracle *PriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracle *PriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracle *PriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracle.Contract.contract.Transact(opts, method, params...)
}

// VAIVALUE is a free data retrieval call binding the contract method 0xb9de61e2.
//
// Solidity: function VAI_VALUE() view returns(uint256)
func (_PriceOracle *PriceOracleCaller) VAIVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "VAI_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VAIVALUE is a free data retrieval call binding the contract method 0xb9de61e2.
//
// Solidity: function VAI_VALUE() view returns(uint256)
func (_PriceOracle *PriceOracleSession) VAIVALUE() (*big.Int, error) {
	return _PriceOracle.Contract.VAIVALUE(&_PriceOracle.CallOpts)
}

// VAIVALUE is a free data retrieval call binding the contract method 0xb9de61e2.
//
// Solidity: function VAI_VALUE() view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) VAIVALUE() (*big.Int, error) {
	return _PriceOracle.Contract.VAIVALUE(&_PriceOracle.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PriceOracle *PriceOracleCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PriceOracle *PriceOracleSession) Admin() (common.Address, error) {
	return _PriceOracle.Contract.Admin(&_PriceOracle.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PriceOracle *PriceOracleCallerSession) Admin() (common.Address, error) {
	return _PriceOracle.Contract.Admin(&_PriceOracle.CallOpts)
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) AssetPrices(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "assetPrices", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_PriceOracle *PriceOracleSession) AssetPrices(asset common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.AssetPrices(&_PriceOracle.CallOpts, asset)
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) AssetPrices(asset common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.AssetPrices(&_PriceOracle.CallOpts, asset)
}

// GetFeed is a free data retrieval call binding the contract method 0x3b39a51c.
//
// Solidity: function getFeed(string symbol) view returns(address)
func (_PriceOracle *PriceOracleCaller) GetFeed(opts *bind.CallOpts, symbol string) (common.Address, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "getFeed", symbol)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeed is a free data retrieval call binding the contract method 0x3b39a51c.
//
// Solidity: function getFeed(string symbol) view returns(address)
func (_PriceOracle *PriceOracleSession) GetFeed(symbol string) (common.Address, error) {
	return _PriceOracle.Contract.GetFeed(&_PriceOracle.CallOpts, symbol)
}

// GetFeed is a free data retrieval call binding the contract method 0x3b39a51c.
//
// Solidity: function getFeed(string symbol) view returns(address)
func (_PriceOracle *PriceOracleCallerSession) GetFeed(symbol string) (common.Address, error) {
	return _PriceOracle.Contract.GetFeed(&_PriceOracle.CallOpts, symbol)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address vToken) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) GetUnderlyingPrice(opts *bind.CallOpts, vToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "getUnderlyingPrice", vToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address vToken) view returns(uint256)
func (_PriceOracle *PriceOracleSession) GetUnderlyingPrice(vToken common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.GetUnderlyingPrice(&_PriceOracle.CallOpts, vToken)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address vToken) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) GetUnderlyingPrice(vToken common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.GetUnderlyingPrice(&_PriceOracle.CallOpts, vToken)
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_PriceOracle *PriceOracleCaller) IsPriceOracle(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "isPriceOracle")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_PriceOracle *PriceOracleSession) IsPriceOracle() (bool, error) {
	return _PriceOracle.Contract.IsPriceOracle(&_PriceOracle.CallOpts)
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_PriceOracle *PriceOracleCallerSession) IsPriceOracle() (bool, error) {
	return _PriceOracle.Contract.IsPriceOracle(&_PriceOracle.CallOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_PriceOracle *PriceOracleTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "setAdmin", newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_PriceOracle *PriceOracleSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetAdmin(&_PriceOracle.TransactOpts, newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_PriceOracle *PriceOracleTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetAdmin(&_PriceOracle.TransactOpts, newAdmin)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_PriceOracle *PriceOracleTransactor) SetDirectPrice(opts *bind.TransactOpts, asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "setDirectPrice", asset, price)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_PriceOracle *PriceOracleSession) SetDirectPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetDirectPrice(&_PriceOracle.TransactOpts, asset, price)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_PriceOracle *PriceOracleTransactorSession) SetDirectPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetDirectPrice(&_PriceOracle.TransactOpts, asset, price)
}

// SetFeed is a paid mutator transaction binding the contract method 0x0c607acf.
//
// Solidity: function setFeed(string symbol, address feed) returns()
func (_PriceOracle *PriceOracleTransactor) SetFeed(opts *bind.TransactOpts, symbol string, feed common.Address) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "setFeed", symbol, feed)
}

// SetFeed is a paid mutator transaction binding the contract method 0x0c607acf.
//
// Solidity: function setFeed(string symbol, address feed) returns()
func (_PriceOracle *PriceOracleSession) SetFeed(symbol string, feed common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetFeed(&_PriceOracle.TransactOpts, symbol, feed)
}

// SetFeed is a paid mutator transaction binding the contract method 0x0c607acf.
//
// Solidity: function setFeed(string symbol, address feed) returns()
func (_PriceOracle *PriceOracleTransactorSession) SetFeed(symbol string, feed common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetFeed(&_PriceOracle.TransactOpts, symbol, feed)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address vToken, uint256 underlyingPriceMantissa) returns()
func (_PriceOracle *PriceOracleTransactor) SetUnderlyingPrice(opts *bind.TransactOpts, vToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "setUnderlyingPrice", vToken, underlyingPriceMantissa)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address vToken, uint256 underlyingPriceMantissa) returns()
func (_PriceOracle *PriceOracleSession) SetUnderlyingPrice(vToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetUnderlyingPrice(&_PriceOracle.TransactOpts, vToken, underlyingPriceMantissa)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address vToken, uint256 underlyingPriceMantissa) returns()
func (_PriceOracle *PriceOracleTransactorSession) SetUnderlyingPrice(vToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetUnderlyingPrice(&_PriceOracle.TransactOpts, vToken, underlyingPriceMantissa)
}

// PriceOracleFeedSetIterator is returned from FilterFeedSet and is used to iterate over the raw logs and unpacked data for FeedSet events raised by the PriceOracle contract.
type PriceOracleFeedSetIterator struct {
	Event *PriceOracleFeedSet // Event containing the contract specifics and raw log

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
func (it *PriceOracleFeedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleFeedSet)
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
		it.Event = new(PriceOracleFeedSet)
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
func (it *PriceOracleFeedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleFeedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleFeedSet represents a FeedSet event raised by the PriceOracle contract.
type PriceOracleFeedSet struct {
	Feed   common.Address
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeedSet is a free log retrieval operation binding the contract event 0xd9e7d1778ca05570ced72c9aeb12a41fcc76f7f57ea25853dea228f8836d0022.
//
// Solidity: event FeedSet(address feed, string symbol)
func (_PriceOracle *PriceOracleFilterer) FilterFeedSet(opts *bind.FilterOpts) (*PriceOracleFeedSetIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "FeedSet")
	if err != nil {
		return nil, err
	}
	return &PriceOracleFeedSetIterator{contract: _PriceOracle.contract, event: "FeedSet", logs: logs, sub: sub}, nil
}

// WatchFeedSet is a free log subscription operation binding the contract event 0xd9e7d1778ca05570ced72c9aeb12a41fcc76f7f57ea25853dea228f8836d0022.
//
// Solidity: event FeedSet(address feed, string symbol)
func (_PriceOracle *PriceOracleFilterer) WatchFeedSet(opts *bind.WatchOpts, sink chan<- *PriceOracleFeedSet) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "FeedSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleFeedSet)
				if err := _PriceOracle.contract.UnpackLog(event, "FeedSet", log); err != nil {
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

// ParseFeedSet is a log parse operation binding the contract event 0xd9e7d1778ca05570ced72c9aeb12a41fcc76f7f57ea25853dea228f8836d0022.
//
// Solidity: event FeedSet(address feed, string symbol)
func (_PriceOracle *PriceOracleFilterer) ParseFeedSet(log types.Log) (*PriceOracleFeedSet, error) {
	event := new(PriceOracleFeedSet)
	if err := _PriceOracle.contract.UnpackLog(event, "FeedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOracleNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the PriceOracle contract.
type PriceOracleNewAdminIterator struct {
	Event *PriceOracleNewAdmin // Event containing the contract specifics and raw log

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
func (it *PriceOracleNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleNewAdmin)
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
		it.Event = new(PriceOracleNewAdmin)
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
func (it *PriceOracleNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleNewAdmin represents a NewAdmin event raised by the PriceOracle contract.
type PriceOracleNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_PriceOracle *PriceOracleFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*PriceOracleNewAdminIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &PriceOracleNewAdminIterator{contract: _PriceOracle.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_PriceOracle *PriceOracleFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *PriceOracleNewAdmin) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleNewAdmin)
				if err := _PriceOracle.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_PriceOracle *PriceOracleFilterer) ParseNewAdmin(log types.Log) (*PriceOracleNewAdmin, error) {
	event := new(PriceOracleNewAdmin)
	if err := _PriceOracle.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOraclePricePostedIterator is returned from FilterPricePosted and is used to iterate over the raw logs and unpacked data for PricePosted events raised by the PriceOracle contract.
type PriceOraclePricePostedIterator struct {
	Event *PriceOraclePricePosted // Event containing the contract specifics and raw log

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
func (it *PriceOraclePricePostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclePricePosted)
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
		it.Event = new(PriceOraclePricePosted)
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
func (it *PriceOraclePricePostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclePricePostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclePricePosted represents a PricePosted event raised by the PriceOracle contract.
type PriceOraclePricePosted struct {
	Asset                  common.Address
	PreviousPriceMantissa  *big.Int
	RequestedPriceMantissa *big.Int
	NewPriceMantissa       *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterPricePosted is a free log retrieval operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_PriceOracle *PriceOracleFilterer) FilterPricePosted(opts *bind.FilterOpts) (*PriceOraclePricePostedIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "PricePosted")
	if err != nil {
		return nil, err
	}
	return &PriceOraclePricePostedIterator{contract: _PriceOracle.contract, event: "PricePosted", logs: logs, sub: sub}, nil
}

// WatchPricePosted is a free log subscription operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_PriceOracle *PriceOracleFilterer) WatchPricePosted(opts *bind.WatchOpts, sink chan<- *PriceOraclePricePosted) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "PricePosted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclePricePosted)
				if err := _PriceOracle.contract.UnpackLog(event, "PricePosted", log); err != nil {
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

// ParsePricePosted is a log parse operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_PriceOracle *PriceOracleFilterer) ParsePricePosted(log types.Log) (*PriceOraclePricePosted, error) {
	event := new(PriceOraclePricePosted)
	if err := _PriceOracle.contract.UnpackLog(event, "PricePosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
