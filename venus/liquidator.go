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

// LiquidatorMetaData contains all meta data concerning the Liquidator contract.
var LiquidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"vBnb_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"treasuryPercentMantissa_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokensForTreasury\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokensForLiquidator\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrowedTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercent\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationTreasuryPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_notEntered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractVToken\",\"name\":\"vTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTreasuryPercentMantissa\",\"type\":\"uint256\"}],\"name\":\"setTreasuryPercent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryPercentMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vBnb\",\"outputs\":[{\"internalType\":\"contractVBNB\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LiquidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidatorMetaData.ABI instead.
var LiquidatorABI = LiquidatorMetaData.ABI

// Liquidator is an auto generated Go binding around an Ethereum contract.
type Liquidator struct {
	LiquidatorCaller     // Read-only binding to the contract
	LiquidatorTransactor // Write-only binding to the contract
	LiquidatorFilterer   // Log filterer for contract events
}

// LiquidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidatorSession struct {
	Contract     *Liquidator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidatorCallerSession struct {
	Contract *LiquidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LiquidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidatorTransactorSession struct {
	Contract     *LiquidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LiquidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidatorRaw struct {
	Contract *Liquidator // Generic contract binding to access the raw methods on
}

// LiquidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidatorCallerRaw struct {
	Contract *LiquidatorCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidatorTransactorRaw struct {
	Contract *LiquidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidator creates a new instance of Liquidator, bound to a specific deployed contract.
func NewLiquidator(address common.Address, backend bind.ContractBackend) (*Liquidator, error) {
	contract, err := bindLiquidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Liquidator{LiquidatorCaller: LiquidatorCaller{contract: contract}, LiquidatorTransactor: LiquidatorTransactor{contract: contract}, LiquidatorFilterer: LiquidatorFilterer{contract: contract}}, nil
}

// NewLiquidatorCaller creates a new read-only instance of Liquidator, bound to a specific deployed contract.
func NewLiquidatorCaller(address common.Address, caller bind.ContractCaller) (*LiquidatorCaller, error) {
	contract, err := bindLiquidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidatorCaller{contract: contract}, nil
}

// NewLiquidatorTransactor creates a new write-only instance of Liquidator, bound to a specific deployed contract.
func NewLiquidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidatorTransactor, error) {
	contract, err := bindLiquidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidatorTransactor{contract: contract}, nil
}

// NewLiquidatorFilterer creates a new log filterer instance of Liquidator, bound to a specific deployed contract.
func NewLiquidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidatorFilterer, error) {
	contract, err := bindLiquidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidatorFilterer{contract: contract}, nil
}

// bindLiquidator binds a generic wrapper to an already deployed contract.
func bindLiquidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquidator *LiquidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquidator.Contract.LiquidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquidator *LiquidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidator.Contract.LiquidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquidator *LiquidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquidator.Contract.LiquidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquidator *LiquidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquidator *LiquidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquidator *LiquidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquidator.Contract.contract.Transact(opts, method, params...)
}

// NotEntered is a free data retrieval call binding the contract method 0xd8438ae8.
//
// Solidity: function _notEntered() view returns(bool)
func (_Liquidator *LiquidatorCaller) NotEntered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "_notEntered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NotEntered is a free data retrieval call binding the contract method 0xd8438ae8.
//
// Solidity: function _notEntered() view returns(bool)
func (_Liquidator *LiquidatorSession) NotEntered() (bool, error) {
	return _Liquidator.Contract.NotEntered(&_Liquidator.CallOpts)
}

// NotEntered is a free data retrieval call binding the contract method 0xd8438ae8.
//
// Solidity: function _notEntered() view returns(bool)
func (_Liquidator *LiquidatorCallerSession) NotEntered() (bool, error) {
	return _Liquidator.Contract.NotEntered(&_Liquidator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Liquidator *LiquidatorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Liquidator *LiquidatorSession) Admin() (common.Address, error) {
	return _Liquidator.Contract.Admin(&_Liquidator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Liquidator *LiquidatorCallerSession) Admin() (common.Address, error) {
	return _Liquidator.Contract.Admin(&_Liquidator.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Liquidator *LiquidatorCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Liquidator *LiquidatorSession) PendingAdmin() (common.Address, error) {
	return _Liquidator.Contract.PendingAdmin(&_Liquidator.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Liquidator *LiquidatorCallerSession) PendingAdmin() (common.Address, error) {
	return _Liquidator.Contract.PendingAdmin(&_Liquidator.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Liquidator *LiquidatorCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Liquidator *LiquidatorSession) Treasury() (common.Address, error) {
	return _Liquidator.Contract.Treasury(&_Liquidator.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Liquidator *LiquidatorCallerSession) Treasury() (common.Address, error) {
	return _Liquidator.Contract.Treasury(&_Liquidator.CallOpts)
}

// TreasuryPercentMantissa is a free data retrieval call binding the contract method 0x8e3525fc.
//
// Solidity: function treasuryPercentMantissa() view returns(uint256)
func (_Liquidator *LiquidatorCaller) TreasuryPercentMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "treasuryPercentMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryPercentMantissa is a free data retrieval call binding the contract method 0x8e3525fc.
//
// Solidity: function treasuryPercentMantissa() view returns(uint256)
func (_Liquidator *LiquidatorSession) TreasuryPercentMantissa() (*big.Int, error) {
	return _Liquidator.Contract.TreasuryPercentMantissa(&_Liquidator.CallOpts)
}

// TreasuryPercentMantissa is a free data retrieval call binding the contract method 0x8e3525fc.
//
// Solidity: function treasuryPercentMantissa() view returns(uint256)
func (_Liquidator *LiquidatorCallerSession) TreasuryPercentMantissa() (*big.Int, error) {
	return _Liquidator.Contract.TreasuryPercentMantissa(&_Liquidator.CallOpts)
}

// VBnb is a free data retrieval call binding the contract method 0x5ba3792d.
//
// Solidity: function vBnb() view returns(address)
func (_Liquidator *LiquidatorCaller) VBnb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "vBnb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VBnb is a free data retrieval call binding the contract method 0x5ba3792d.
//
// Solidity: function vBnb() view returns(address)
func (_Liquidator *LiquidatorSession) VBnb() (common.Address, error) {
	return _Liquidator.Contract.VBnb(&_Liquidator.CallOpts)
}

// VBnb is a free data retrieval call binding the contract method 0x5ba3792d.
//
// Solidity: function vBnb() view returns(address)
func (_Liquidator *LiquidatorCallerSession) VBnb() (common.Address, error) {
	return _Liquidator.Contract.VBnb(&_Liquidator.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns()
func (_Liquidator *LiquidatorTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns()
func (_Liquidator *LiquidatorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Liquidator.Contract.AcceptAdmin(&_Liquidator.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns()
func (_Liquidator *LiquidatorTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Liquidator.Contract.AcceptAdmin(&_Liquidator.TransactOpts)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (_Liquidator *LiquidatorTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (_Liquidator *LiquidatorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.SetPendingAdmin(&_Liquidator.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (_Liquidator *LiquidatorTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.SetPendingAdmin(&_Liquidator.TransactOpts, newPendingAdmin)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address vToken, address borrower, uint256 repayAmount, address vTokenCollateral) payable returns()
func (_Liquidator *LiquidatorTransactor) LiquidateBorrow(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "liquidateBorrow", vToken, borrower, repayAmount, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address vToken, address borrower, uint256 repayAmount, address vTokenCollateral) payable returns()
func (_Liquidator *LiquidatorSession) LiquidateBorrow(vToken common.Address, borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.LiquidateBorrow(&_Liquidator.TransactOpts, vToken, borrower, repayAmount, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address vToken, address borrower, uint256 repayAmount, address vTokenCollateral) payable returns()
func (_Liquidator *LiquidatorTransactorSession) LiquidateBorrow(vToken common.Address, borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.LiquidateBorrow(&_Liquidator.TransactOpts, vToken, borrower, repayAmount, vTokenCollateral)
}

// SetTreasuryPercent is a paid mutator transaction binding the contract method 0x89a2bc25.
//
// Solidity: function setTreasuryPercent(uint256 newTreasuryPercentMantissa) returns()
func (_Liquidator *LiquidatorTransactor) SetTreasuryPercent(opts *bind.TransactOpts, newTreasuryPercentMantissa *big.Int) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "setTreasuryPercent", newTreasuryPercentMantissa)
}

// SetTreasuryPercent is a paid mutator transaction binding the contract method 0x89a2bc25.
//
// Solidity: function setTreasuryPercent(uint256 newTreasuryPercentMantissa) returns()
func (_Liquidator *LiquidatorSession) SetTreasuryPercent(newTreasuryPercentMantissa *big.Int) (*types.Transaction, error) {
	return _Liquidator.Contract.SetTreasuryPercent(&_Liquidator.TransactOpts, newTreasuryPercentMantissa)
}

// SetTreasuryPercent is a paid mutator transaction binding the contract method 0x89a2bc25.
//
// Solidity: function setTreasuryPercent(uint256 newTreasuryPercentMantissa) returns()
func (_Liquidator *LiquidatorTransactorSession) SetTreasuryPercent(newTreasuryPercentMantissa *big.Int) (*types.Transaction, error) {
	return _Liquidator.Contract.SetTreasuryPercent(&_Liquidator.TransactOpts, newTreasuryPercentMantissa)
}

// LiquidatorLiquidateBorrowedTokensIterator is returned from FilterLiquidateBorrowedTokens and is used to iterate over the raw logs and unpacked data for LiquidateBorrowedTokens events raised by the Liquidator contract.
type LiquidatorLiquidateBorrowedTokensIterator struct {
	Event *LiquidatorLiquidateBorrowedTokens // Event containing the contract specifics and raw log

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
func (it *LiquidatorLiquidateBorrowedTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorLiquidateBorrowedTokens)
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
		it.Event = new(LiquidatorLiquidateBorrowedTokens)
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
func (it *LiquidatorLiquidateBorrowedTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorLiquidateBorrowedTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorLiquidateBorrowedTokens represents a LiquidateBorrowedTokens event raised by the Liquidator contract.
type LiquidatorLiquidateBorrowedTokens struct {
	Liquidator               common.Address
	Borrower                 common.Address
	RepayAmount              *big.Int
	VTokenCollateral         common.Address
	SeizeTokensForTreasury   *big.Int
	SeizeTokensForLiquidator *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrowedTokens is a free log retrieval operation binding the contract event 0xd71542567e4876e606f2a1dc04c1b26f3cf39cb975eb0a65aef3c98842410583.
//
// Solidity: event LiquidateBorrowedTokens(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokensForTreasury, uint256 seizeTokensForLiquidator)
func (_Liquidator *LiquidatorFilterer) FilterLiquidateBorrowedTokens(opts *bind.FilterOpts) (*LiquidatorLiquidateBorrowedTokensIterator, error) {

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "LiquidateBorrowedTokens")
	if err != nil {
		return nil, err
	}
	return &LiquidatorLiquidateBorrowedTokensIterator{contract: _Liquidator.contract, event: "LiquidateBorrowedTokens", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrowedTokens is a free log subscription operation binding the contract event 0xd71542567e4876e606f2a1dc04c1b26f3cf39cb975eb0a65aef3c98842410583.
//
// Solidity: event LiquidateBorrowedTokens(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokensForTreasury, uint256 seizeTokensForLiquidator)
func (_Liquidator *LiquidatorFilterer) WatchLiquidateBorrowedTokens(opts *bind.WatchOpts, sink chan<- *LiquidatorLiquidateBorrowedTokens) (event.Subscription, error) {

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "LiquidateBorrowedTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorLiquidateBorrowedTokens)
				if err := _Liquidator.contract.UnpackLog(event, "LiquidateBorrowedTokens", log); err != nil {
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

// ParseLiquidateBorrowedTokens is a log parse operation binding the contract event 0xd71542567e4876e606f2a1dc04c1b26f3cf39cb975eb0a65aef3c98842410583.
//
// Solidity: event LiquidateBorrowedTokens(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokensForTreasury, uint256 seizeTokensForLiquidator)
func (_Liquidator *LiquidatorFilterer) ParseLiquidateBorrowedTokens(log types.Log) (*LiquidatorLiquidateBorrowedTokens, error) {
	event := new(LiquidatorLiquidateBorrowedTokens)
	if err := _Liquidator.contract.UnpackLog(event, "LiquidateBorrowedTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Liquidator contract.
type LiquidatorNewAdminIterator struct {
	Event *LiquidatorNewAdmin // Event containing the contract specifics and raw log

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
func (it *LiquidatorNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorNewAdmin)
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
		it.Event = new(LiquidatorNewAdmin)
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
func (it *LiquidatorNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorNewAdmin represents a NewAdmin event raised by the Liquidator contract.
type LiquidatorNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Liquidator *LiquidatorFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*LiquidatorNewAdminIterator, error) {

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &LiquidatorNewAdminIterator{contract: _Liquidator.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Liquidator *LiquidatorFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *LiquidatorNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorNewAdmin)
				if err := _Liquidator.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Liquidator *LiquidatorFilterer) ParseNewAdmin(log types.Log) (*LiquidatorNewAdmin, error) {
	event := new(LiquidatorNewAdmin)
	if err := _Liquidator.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorNewLiquidationTreasuryPercentIterator is returned from FilterNewLiquidationTreasuryPercent and is used to iterate over the raw logs and unpacked data for NewLiquidationTreasuryPercent events raised by the Liquidator contract.
type LiquidatorNewLiquidationTreasuryPercentIterator struct {
	Event *LiquidatorNewLiquidationTreasuryPercent // Event containing the contract specifics and raw log

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
func (it *LiquidatorNewLiquidationTreasuryPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorNewLiquidationTreasuryPercent)
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
		it.Event = new(LiquidatorNewLiquidationTreasuryPercent)
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
func (it *LiquidatorNewLiquidationTreasuryPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorNewLiquidationTreasuryPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorNewLiquidationTreasuryPercent represents a NewLiquidationTreasuryPercent event raised by the Liquidator contract.
type LiquidatorNewLiquidationTreasuryPercent struct {
	OldPercent *big.Int
	NewPercent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationTreasuryPercent is a free log retrieval operation binding the contract event 0x0e9a1641744b21c49e4183ab9ce941d420e05a05a446c71875fb07814f362c1a.
//
// Solidity: event NewLiquidationTreasuryPercent(uint256 oldPercent, uint256 newPercent)
func (_Liquidator *LiquidatorFilterer) FilterNewLiquidationTreasuryPercent(opts *bind.FilterOpts) (*LiquidatorNewLiquidationTreasuryPercentIterator, error) {

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "NewLiquidationTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return &LiquidatorNewLiquidationTreasuryPercentIterator{contract: _Liquidator.contract, event: "NewLiquidationTreasuryPercent", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationTreasuryPercent is a free log subscription operation binding the contract event 0x0e9a1641744b21c49e4183ab9ce941d420e05a05a446c71875fb07814f362c1a.
//
// Solidity: event NewLiquidationTreasuryPercent(uint256 oldPercent, uint256 newPercent)
func (_Liquidator *LiquidatorFilterer) WatchNewLiquidationTreasuryPercent(opts *bind.WatchOpts, sink chan<- *LiquidatorNewLiquidationTreasuryPercent) (event.Subscription, error) {

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "NewLiquidationTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorNewLiquidationTreasuryPercent)
				if err := _Liquidator.contract.UnpackLog(event, "NewLiquidationTreasuryPercent", log); err != nil {
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

// ParseNewLiquidationTreasuryPercent is a log parse operation binding the contract event 0x0e9a1641744b21c49e4183ab9ce941d420e05a05a446c71875fb07814f362c1a.
//
// Solidity: event NewLiquidationTreasuryPercent(uint256 oldPercent, uint256 newPercent)
func (_Liquidator *LiquidatorFilterer) ParseNewLiquidationTreasuryPercent(log types.Log) (*LiquidatorNewLiquidationTreasuryPercent, error) {
	event := new(LiquidatorNewLiquidationTreasuryPercent)
	if err := _Liquidator.contract.UnpackLog(event, "NewLiquidationTreasuryPercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Liquidator contract.
type LiquidatorNewPendingAdminIterator struct {
	Event *LiquidatorNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *LiquidatorNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorNewPendingAdmin)
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
		it.Event = new(LiquidatorNewPendingAdmin)
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
func (it *LiquidatorNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorNewPendingAdmin represents a NewPendingAdmin event raised by the Liquidator contract.
type LiquidatorNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Liquidator *LiquidatorFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*LiquidatorNewPendingAdminIterator, error) {

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &LiquidatorNewPendingAdminIterator{contract: _Liquidator.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Liquidator *LiquidatorFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *LiquidatorNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorNewPendingAdmin)
				if err := _Liquidator.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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
func (_Liquidator *LiquidatorFilterer) ParseNewPendingAdmin(log types.Log) (*LiquidatorNewPendingAdmin, error) {
	event := new(LiquidatorNewPendingAdmin)
	if err := _Liquidator.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
