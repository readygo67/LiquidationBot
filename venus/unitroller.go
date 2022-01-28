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

// UnitrollerMetaData contains all meta data concerning the Unitroller contract.
var UnitrollerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingImplementation\",\"type\":\"address\"}],\"name\":\"NewPendingImplementation\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptImplementation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingImplementation\",\"type\":\"address\"}],\"name\":\"_setPendingImplementation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UnitrollerABI is the input ABI used to generate the binding from.
// Deprecated: Use UnitrollerMetaData.ABI instead.
var UnitrollerABI = UnitrollerMetaData.ABI

// Unitroller is an auto generated Go binding around an Ethereum contract.
type Unitroller struct {
	UnitrollerCaller     // Read-only binding to the contract
	UnitrollerTransactor // Write-only binding to the contract
	UnitrollerFilterer   // Log filterer for contract events
}

// UnitrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnitrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnitrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnitrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnitrollerSession struct {
	Contract     *Unitroller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnitrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnitrollerCallerSession struct {
	Contract *UnitrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UnitrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnitrollerTransactorSession struct {
	Contract     *UnitrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UnitrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnitrollerRaw struct {
	Contract *Unitroller // Generic contract binding to access the raw methods on
}

// UnitrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnitrollerCallerRaw struct {
	Contract *UnitrollerCaller // Generic read-only contract binding to access the raw methods on
}

// UnitrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnitrollerTransactorRaw struct {
	Contract *UnitrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnitroller creates a new instance of Unitroller, bound to a specific deployed contract.
func NewUnitroller(address common.Address, backend bind.ContractBackend) (*Unitroller, error) {
	contract, err := bindUnitroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unitroller{UnitrollerCaller: UnitrollerCaller{contract: contract}, UnitrollerTransactor: UnitrollerTransactor{contract: contract}, UnitrollerFilterer: UnitrollerFilterer{contract: contract}}, nil
}

// NewUnitrollerCaller creates a new read-only instance of Unitroller, bound to a specific deployed contract.
func NewUnitrollerCaller(address common.Address, caller bind.ContractCaller) (*UnitrollerCaller, error) {
	contract, err := bindUnitroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnitrollerCaller{contract: contract}, nil
}

// NewUnitrollerTransactor creates a new write-only instance of Unitroller, bound to a specific deployed contract.
func NewUnitrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*UnitrollerTransactor, error) {
	contract, err := bindUnitroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnitrollerTransactor{contract: contract}, nil
}

// NewUnitrollerFilterer creates a new log filterer instance of Unitroller, bound to a specific deployed contract.
func NewUnitrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*UnitrollerFilterer, error) {
	contract, err := bindUnitroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnitrollerFilterer{contract: contract}, nil
}

// bindUnitroller binds a generic wrapper to an already deployed contract.
func bindUnitroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnitrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unitroller *UnitrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unitroller.Contract.UnitrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unitroller *UnitrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unitroller.Contract.UnitrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unitroller *UnitrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unitroller.Contract.UnitrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unitroller *UnitrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unitroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unitroller *UnitrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unitroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unitroller *UnitrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unitroller.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Unitroller *UnitrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Unitroller *UnitrollerSession) Admin() (common.Address, error) {
	return _Unitroller.Contract.Admin(&_Unitroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Unitroller *UnitrollerCallerSession) Admin() (common.Address, error) {
	return _Unitroller.Contract.Admin(&_Unitroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _Unitroller.Contract.ComptrollerImplementation(&_Unitroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _Unitroller.Contract.ComptrollerImplementation(&_Unitroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Unitroller *UnitrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Unitroller *UnitrollerSession) PendingAdmin() (common.Address, error) {
	return _Unitroller.Contract.PendingAdmin(&_Unitroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Unitroller *UnitrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _Unitroller.Contract.PendingAdmin(&_Unitroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Unitroller.Contract.PendingComptrollerImplementation(&_Unitroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Unitroller.Contract.PendingComptrollerImplementation(&_Unitroller.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Unitroller *UnitrollerTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Unitroller *UnitrollerSession) AcceptAdmin() (*types.Transaction, error) {
	return _Unitroller.Contract.AcceptAdmin(&_Unitroller.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Unitroller.Contract.AcceptAdmin(&_Unitroller.TransactOpts)
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_Unitroller *UnitrollerTransactor) AcceptImplementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_acceptImplementation")
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_Unitroller *UnitrollerSession) AcceptImplementation() (*types.Transaction, error) {
	return _Unitroller.Contract.AcceptImplementation(&_Unitroller.TransactOpts)
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) AcceptImplementation() (*types.Transaction, error) {
	return _Unitroller.Contract.AcceptImplementation(&_Unitroller.TransactOpts)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Unitroller *UnitrollerTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Unitroller *UnitrollerSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SetPendingAdmin(&_Unitroller.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SetPendingAdmin(&_Unitroller.TransactOpts, newPendingAdmin)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_Unitroller *UnitrollerTransactor) SetPendingImplementation(opts *bind.TransactOpts, newPendingImplementation common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setPendingImplementation", newPendingImplementation)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_Unitroller *UnitrollerSession) SetPendingImplementation(newPendingImplementation common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SetPendingImplementation(&_Unitroller.TransactOpts, newPendingImplementation)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) SetPendingImplementation(newPendingImplementation common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SetPendingImplementation(&_Unitroller.TransactOpts, newPendingImplementation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Unitroller *UnitrollerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Unitroller.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Unitroller *UnitrollerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Unitroller.Contract.Fallback(&_Unitroller.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Unitroller *UnitrollerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Unitroller.Contract.Fallback(&_Unitroller.TransactOpts, calldata)
}

// UnitrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Unitroller contract.
type UnitrollerFailureIterator struct {
	Event *UnitrollerFailure // Event containing the contract specifics and raw log

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
func (it *UnitrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerFailure)
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
		it.Event = new(UnitrollerFailure)
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
func (it *UnitrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerFailure represents a Failure event raised by the Unitroller contract.
type UnitrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Unitroller *UnitrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*UnitrollerFailureIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &UnitrollerFailureIterator{contract: _Unitroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Unitroller *UnitrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *UnitrollerFailure) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerFailure)
				if err := _Unitroller.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_Unitroller *UnitrollerFilterer) ParseFailure(log types.Log) (*UnitrollerFailure, error) {
	event := new(UnitrollerFailure)
	if err := _Unitroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Unitroller contract.
type UnitrollerNewAdminIterator struct {
	Event *UnitrollerNewAdmin // Event containing the contract specifics and raw log

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
func (it *UnitrollerNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewAdmin)
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
		it.Event = new(UnitrollerNewAdmin)
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
func (it *UnitrollerNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewAdmin represents a NewAdmin event raised by the Unitroller contract.
type UnitrollerNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Unitroller *UnitrollerFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*UnitrollerNewAdminIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewAdminIterator{contract: _Unitroller.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Unitroller *UnitrollerFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *UnitrollerNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewAdmin)
				if err := _Unitroller.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Unitroller *UnitrollerFilterer) ParseNewAdmin(log types.Log) (*UnitrollerNewAdmin, error) {
	event := new(UnitrollerNewAdmin)
	if err := _Unitroller.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewImplementationIterator is returned from FilterNewImplementation and is used to iterate over the raw logs and unpacked data for NewImplementation events raised by the Unitroller contract.
type UnitrollerNewImplementationIterator struct {
	Event *UnitrollerNewImplementation // Event containing the contract specifics and raw log

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
func (it *UnitrollerNewImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewImplementation)
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
		it.Event = new(UnitrollerNewImplementation)
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
func (it *UnitrollerNewImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewImplementation represents a NewImplementation event raised by the Unitroller contract.
type UnitrollerNewImplementation struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewImplementation is a free log retrieval operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Unitroller *UnitrollerFilterer) FilterNewImplementation(opts *bind.FilterOpts) (*UnitrollerNewImplementationIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewImplementationIterator{contract: _Unitroller.contract, event: "NewImplementation", logs: logs, sub: sub}, nil
}

// WatchNewImplementation is a free log subscription operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Unitroller *UnitrollerFilterer) WatchNewImplementation(opts *bind.WatchOpts, sink chan<- *UnitrollerNewImplementation) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewImplementation)
				if err := _Unitroller.contract.UnpackLog(event, "NewImplementation", log); err != nil {
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

// ParseNewImplementation is a log parse operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Unitroller *UnitrollerFilterer) ParseNewImplementation(log types.Log) (*UnitrollerNewImplementation, error) {
	event := new(UnitrollerNewImplementation)
	if err := _Unitroller.contract.UnpackLog(event, "NewImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Unitroller contract.
type UnitrollerNewPendingAdminIterator struct {
	Event *UnitrollerNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *UnitrollerNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewPendingAdmin)
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
		it.Event = new(UnitrollerNewPendingAdmin)
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
func (it *UnitrollerNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewPendingAdmin represents a NewPendingAdmin event raised by the Unitroller contract.
type UnitrollerNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Unitroller *UnitrollerFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*UnitrollerNewPendingAdminIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewPendingAdminIterator{contract: _Unitroller.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Unitroller *UnitrollerFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *UnitrollerNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewPendingAdmin)
				if err := _Unitroller.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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
func (_Unitroller *UnitrollerFilterer) ParseNewPendingAdmin(log types.Log) (*UnitrollerNewPendingAdmin, error) {
	event := new(UnitrollerNewPendingAdmin)
	if err := _Unitroller.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewPendingImplementationIterator is returned from FilterNewPendingImplementation and is used to iterate over the raw logs and unpacked data for NewPendingImplementation events raised by the Unitroller contract.
type UnitrollerNewPendingImplementationIterator struct {
	Event *UnitrollerNewPendingImplementation // Event containing the contract specifics and raw log

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
func (it *UnitrollerNewPendingImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewPendingImplementation)
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
		it.Event = new(UnitrollerNewPendingImplementation)
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
func (it *UnitrollerNewPendingImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewPendingImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewPendingImplementation represents a NewPendingImplementation event raised by the Unitroller contract.
type UnitrollerNewPendingImplementation struct {
	OldPendingImplementation common.Address
	NewPendingImplementation common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewPendingImplementation is a free log retrieval operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address oldPendingImplementation, address newPendingImplementation)
func (_Unitroller *UnitrollerFilterer) FilterNewPendingImplementation(opts *bind.FilterOpts) (*UnitrollerNewPendingImplementationIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewPendingImplementation")
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewPendingImplementationIterator{contract: _Unitroller.contract, event: "NewPendingImplementation", logs: logs, sub: sub}, nil
}

// WatchNewPendingImplementation is a free log subscription operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address oldPendingImplementation, address newPendingImplementation)
func (_Unitroller *UnitrollerFilterer) WatchNewPendingImplementation(opts *bind.WatchOpts, sink chan<- *UnitrollerNewPendingImplementation) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewPendingImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewPendingImplementation)
				if err := _Unitroller.contract.UnpackLog(event, "NewPendingImplementation", log); err != nil {
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

// ParseNewPendingImplementation is a log parse operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address oldPendingImplementation, address newPendingImplementation)
func (_Unitroller *UnitrollerFilterer) ParseNewPendingImplementation(log types.Log) (*UnitrollerNewPendingImplementation, error) {
	event := new(UnitrollerNewPendingImplementation)
	if err := _Unitroller.contract.UnpackLog(event, "NewPendingImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
