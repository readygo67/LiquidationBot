// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package server

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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202aaa4aedd5bd3c66c8f598575e3ce25c153f5e6fb8fa35c93c647a849cbef4bd64736f6c634300080b0033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ComptrollerMetaData contains all meta data concerning the Comptroller contract.
var ComptrollerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"claimVenus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimVenus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimVenus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaiMinter\",\"type\":\"address\"}],\"name\":\"distributeVAIMinterVenus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getXVSAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastContributorBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateVAICalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVenus\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minReleaseAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintedVAIs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseToVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repayVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMintedVAIOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaiController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaiMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaiVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusContributorSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"venusInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"venusRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"venusVAIRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"venusVAIVaultRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dce15449": "accountAssets(address,uint256)",
		"f851a440": "admin()",
		"52d84d1e": "allMarkets(uint256)",
		"da3d454c": "borrowAllowed(address,address,uint256)",
		"21af4569": "borrowCapGuardian()",
		"4a584432": "borrowCaps(address)",
		"6d154ea5": "borrowGuardianPaused(address)",
		"5c778605": "borrowVerify(address,address,uint256)",
		"929fe9a1": "checkMembership(address,address)",
		"adcd5fb9": "claimVenus(address)",
		"86df31ee": "claimVenus(address,address[])",
		"d09c54ba": "claimVenus(address[],address[],bool,bool)",
		"e8755446": "closeFactorMantissa()",
		"bb82aa5e": "comptrollerImplementation()",
		"f4b8d5fe": "distributeVAIMinterVenus(address)",
		"c2998238": "enterMarkets(address[])",
		"ede4edd0": "exitMarket(address)",
		"5ec88c79": "getAccountLiquidity(address)",
		"b0772d0b": "getAllMarkets()",
		"abfceffc": "getAssetsIn(address)",
		"42cbb15c": "getBlockNumber()",
		"4e79238f": "getHypotheticalAccountLiquidity(address,address,uint256,uint256)",
		"bf32442d": "getXVSAddress()",
		"007e3dd2": "isComptroller()",
		"bea6b8b8": "lastContributorBlock(address)",
		"5fc7e71e": "liquidateBorrowAllowed(address,address,address,address,uint256)",
		"47ef3b3b": "liquidateBorrowVerify(address,address,address,address,uint256,uint256)",
		"c488847b": "liquidateCalculateSeizeTokens(address,address,uint256)",
		"a78dc775": "liquidateVAICalculateSeizeTokens(address,uint256)",
		"4ada90af": "liquidationIncentiveMantissa()",
		"8e8f294b": "markets(address)",
		"94b2294b": "maxAssets()",
		"0db4b4e5": "minReleaseAmount()",
		"4ef4c3e1": "mintAllowed(address,address,uint256)",
		"731f0c2b": "mintGuardianPaused(address)",
		"4088c73e": "mintVAIGuardianPaused()",
		"41c728b9": "mintVerify(address,address,uint256,uint256)",
		"2bc7e29e": "mintedVAIs(address)",
		"7dc0d1d0": "oracle()",
		"24a3d622": "pauseGuardian()",
		"26782247": "pendingAdmin()",
		"dcfbc0c7": "pendingComptrollerImplementation()",
		"425fad58": "protocolPaused()",
		"eabe7d91": "redeemAllowed(address,address,uint256)",
		"51dff989": "redeemVerify(address,address,uint256,uint256)",
		"719f701b": "releaseStartBlock()",
		"ddfd287e": "releaseToVault()",
		"24008a62": "repayBorrowAllowed(address,address,address,uint256)",
		"1ededc91": "repayBorrowVerify(address,address,address,uint256,uint256)",
		"76551383": "repayVAIGuardianPaused()",
		"d02f7351": "seizeAllowed(address,address,address,address,uint256)",
		"ac0b0bb7": "seizeGuardianPaused()",
		"6d35bf91": "seizeVerify(address,address,address,address,uint256)",
		"fd51a3ad": "setMintedVAIOf(address,uint256)",
		"bdcdc258": "transferAllowed(address,address,address,uint256)",
		"87f76303": "transferGuardianPaused()",
		"6a56947e": "transferVerify(address,address,address,uint256)",
		"c5f956af": "treasuryAddress()",
		"b2eafc39": "treasuryGuardian()",
		"04ef9d58": "treasuryPercent()",
		"9254f5e5": "vaiController()",
		"bec04f72": "vaiMintRate()",
		"7d172bd5": "vaiVaultAddress()",
		"8a7dc165": "venusAccrued(address)",
		"e37d4b79": "venusBorrowState(address)",
		"08e0225c": "venusBorrowerIndex(address,address)",
		"a9046134": "venusContributorSpeeds(address)",
		"c5b4db55": "venusInitialIndex()",
		"879c2e1d": "venusRate()",
		"1abcaa77": "venusSpeeds(address)",
		"41a18d2c": "venusSupplierIndex(address,address)",
		"b8324c7c": "venusSupplyState(address)",
		"399cc80c": "venusVAIRate()",
		"fa6331d8": "venusVAIVaultRate()",
	},
}

// ComptrollerABI is the input ABI used to generate the binding from.
// Deprecated: Use ComptrollerMetaData.ABI instead.
var ComptrollerABI = ComptrollerMetaData.ABI

// Deprecated: Use ComptrollerMetaData.Sigs instead.
// ComptrollerFuncSigs maps the 4-byte function signature to its string representation.
var ComptrollerFuncSigs = ComptrollerMetaData.Sigs

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

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_Comptroller *ComptrollerTransactor) DistributeVAIMinterVenus(opts *bind.TransactOpts, vaiMinter common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "distributeVAIMinterVenus", vaiMinter)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_Comptroller *ComptrollerSession) DistributeVAIMinterVenus(vaiMinter common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.DistributeVAIMinterVenus(&_Comptroller.TransactOpts, vaiMinter)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_Comptroller *ComptrollerTransactorSession) DistributeVAIMinterVenus(vaiMinter common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.DistributeVAIMinterVenus(&_Comptroller.TransactOpts, vaiMinter)
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

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Session) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Session) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20CallerSession) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Session) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20CallerSession) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakeCalleeMetaData contains all meta data concerning the IPancakeCallee contract.
var IPancakeCalleeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"pancakeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"84800812": "pancakeCall(address,uint256,uint256,bytes)",
	},
}

// IPancakeCalleeABI is the input ABI used to generate the binding from.
// Deprecated: Use IPancakeCalleeMetaData.ABI instead.
var IPancakeCalleeABI = IPancakeCalleeMetaData.ABI

// Deprecated: Use IPancakeCalleeMetaData.Sigs instead.
// IPancakeCalleeFuncSigs maps the 4-byte function signature to its string representation.
var IPancakeCalleeFuncSigs = IPancakeCalleeMetaData.Sigs

// IPancakeCallee is an auto generated Go binding around an Ethereum contract.
type IPancakeCallee struct {
	IPancakeCalleeCaller     // Read-only binding to the contract
	IPancakeCalleeTransactor // Write-only binding to the contract
	IPancakeCalleeFilterer   // Log filterer for contract events
}

// IPancakeCalleeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakeCalleeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeCalleeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakeCalleeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeCalleeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakeCalleeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeCalleeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakeCalleeSession struct {
	Contract     *IPancakeCallee   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakeCalleeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakeCalleeCallerSession struct {
	Contract *IPancakeCalleeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IPancakeCalleeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakeCalleeTransactorSession struct {
	Contract     *IPancakeCalleeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IPancakeCalleeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakeCalleeRaw struct {
	Contract *IPancakeCallee // Generic contract binding to access the raw methods on
}

// IPancakeCalleeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakeCalleeCallerRaw struct {
	Contract *IPancakeCalleeCaller // Generic read-only contract binding to access the raw methods on
}

// IPancakeCalleeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakeCalleeTransactorRaw struct {
	Contract *IPancakeCalleeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakeCallee creates a new instance of IPancakeCallee, bound to a specific deployed contract.
func NewIPancakeCallee(address common.Address, backend bind.ContractBackend) (*IPancakeCallee, error) {
	contract, err := bindIPancakeCallee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakeCallee{IPancakeCalleeCaller: IPancakeCalleeCaller{contract: contract}, IPancakeCalleeTransactor: IPancakeCalleeTransactor{contract: contract}, IPancakeCalleeFilterer: IPancakeCalleeFilterer{contract: contract}}, nil
}

// NewIPancakeCalleeCaller creates a new read-only instance of IPancakeCallee, bound to a specific deployed contract.
func NewIPancakeCalleeCaller(address common.Address, caller bind.ContractCaller) (*IPancakeCalleeCaller, error) {
	contract, err := bindIPancakeCallee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeCalleeCaller{contract: contract}, nil
}

// NewIPancakeCalleeTransactor creates a new write-only instance of IPancakeCallee, bound to a specific deployed contract.
func NewIPancakeCalleeTransactor(address common.Address, transactor bind.ContractTransactor) (*IPancakeCalleeTransactor, error) {
	contract, err := bindIPancakeCallee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeCalleeTransactor{contract: contract}, nil
}

// NewIPancakeCalleeFilterer creates a new log filterer instance of IPancakeCallee, bound to a specific deployed contract.
func NewIPancakeCalleeFilterer(address common.Address, filterer bind.ContractFilterer) (*IPancakeCalleeFilterer, error) {
	contract, err := bindIPancakeCallee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakeCalleeFilterer{contract: contract}, nil
}

// bindIPancakeCallee binds a generic wrapper to an already deployed contract.
func bindIPancakeCallee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPancakeCalleeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeCallee *IPancakeCalleeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeCallee.Contract.IPancakeCalleeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeCallee *IPancakeCalleeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeCallee.Contract.IPancakeCalleeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeCallee *IPancakeCalleeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeCallee.Contract.IPancakeCalleeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeCallee *IPancakeCalleeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeCallee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeCallee *IPancakeCalleeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeCallee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeCallee *IPancakeCalleeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeCallee.Contract.contract.Transact(opts, method, params...)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IPancakeCallee *IPancakeCalleeTransactor) PancakeCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IPancakeCallee.contract.Transact(opts, "pancakeCall", sender, amount0, amount1, data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IPancakeCallee *IPancakeCalleeSession) PancakeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IPancakeCallee.Contract.PancakeCall(&_IPancakeCallee.TransactOpts, sender, amount0, amount1, data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IPancakeCallee *IPancakeCalleeTransactorSession) PancakeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IPancakeCallee.Contract.PancakeCall(&_IPancakeCallee.TransactOpts, sender, amount0, amount1, data)
}

// IPancakeFactoryMetaData contains all meta data concerning the IPancakeFactory contract.
var IPancakeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1e3dd18b": "allPairs(uint256)",
		"574f2ba3": "allPairsLength()",
		"c9c65396": "createPair(address,address)",
		"017e7e58": "feeTo()",
		"094b7415": "feeToSetter()",
		"e6a43905": "getPair(address,address)",
		"f46901ed": "setFeeTo(address)",
		"a2e74af6": "setFeeToSetter(address)",
	},
}

// IPancakeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IPancakeFactoryMetaData.ABI instead.
var IPancakeFactoryABI = IPancakeFactoryMetaData.ABI

// Deprecated: Use IPancakeFactoryMetaData.Sigs instead.
// IPancakeFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IPancakeFactoryFuncSigs = IPancakeFactoryMetaData.Sigs

// IPancakeFactory is an auto generated Go binding around an Ethereum contract.
type IPancakeFactory struct {
	IPancakeFactoryCaller     // Read-only binding to the contract
	IPancakeFactoryTransactor // Write-only binding to the contract
	IPancakeFactoryFilterer   // Log filterer for contract events
}

// IPancakeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakeFactorySession struct {
	Contract     *IPancakeFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakeFactoryCallerSession struct {
	Contract *IPancakeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IPancakeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakeFactoryTransactorSession struct {
	Contract     *IPancakeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IPancakeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakeFactoryRaw struct {
	Contract *IPancakeFactory // Generic contract binding to access the raw methods on
}

// IPancakeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakeFactoryCallerRaw struct {
	Contract *IPancakeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IPancakeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakeFactoryTransactorRaw struct {
	Contract *IPancakeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakeFactory creates a new instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactory(address common.Address, backend bind.ContractBackend) (*IPancakeFactory, error) {
	contract, err := bindIPancakeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactory{IPancakeFactoryCaller: IPancakeFactoryCaller{contract: contract}, IPancakeFactoryTransactor: IPancakeFactoryTransactor{contract: contract}, IPancakeFactoryFilterer: IPancakeFactoryFilterer{contract: contract}}, nil
}

// NewIPancakeFactoryCaller creates a new read-only instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactoryCaller(address common.Address, caller bind.ContractCaller) (*IPancakeFactoryCaller, error) {
	contract, err := bindIPancakeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryCaller{contract: contract}, nil
}

// NewIPancakeFactoryTransactor creates a new write-only instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IPancakeFactoryTransactor, error) {
	contract, err := bindIPancakeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryTransactor{contract: contract}, nil
}

// NewIPancakeFactoryFilterer creates a new log filterer instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IPancakeFactoryFilterer, error) {
	contract, err := bindIPancakeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryFilterer{contract: contract}, nil
}

// bindIPancakeFactory binds a generic wrapper to an already deployed contract.
func bindIPancakeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPancakeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeFactory *IPancakeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeFactory.Contract.IPancakeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeFactory *IPancakeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.IPancakeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeFactory *IPancakeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.IPancakeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeFactory *IPancakeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeFactory *IPancakeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeFactory *IPancakeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IPancakeFactory *IPancakeFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IPancakeFactory *IPancakeFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IPancakeFactory.Contract.AllPairs(&_IPancakeFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IPancakeFactory *IPancakeFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IPancakeFactory.Contract.AllPairs(&_IPancakeFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IPancakeFactory *IPancakeFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IPancakeFactory *IPancakeFactorySession) AllPairsLength() (*big.Int, error) {
	return _IPancakeFactory.Contract.AllPairsLength(&_IPancakeFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IPancakeFactory *IPancakeFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _IPancakeFactory.Contract.AllPairsLength(&_IPancakeFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IPancakeFactory *IPancakeFactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IPancakeFactory *IPancakeFactorySession) FeeTo() (common.Address, error) {
	return _IPancakeFactory.Contract.FeeTo(&_IPancakeFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IPancakeFactory *IPancakeFactoryCallerSession) FeeTo() (common.Address, error) {
	return _IPancakeFactory.Contract.FeeTo(&_IPancakeFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IPancakeFactory *IPancakeFactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IPancakeFactory *IPancakeFactorySession) FeeToSetter() (common.Address, error) {
	return _IPancakeFactory.Contract.FeeToSetter(&_IPancakeFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IPancakeFactory *IPancakeFactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _IPancakeFactory.Contract.FeeToSetter(&_IPancakeFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPancakeFactory *IPancakeFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPancakeFactory *IPancakeFactorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IPancakeFactory.Contract.GetPair(&_IPancakeFactory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPancakeFactory *IPancakeFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IPancakeFactory.Contract.GetPair(&_IPancakeFactory.CallOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IPancakeFactory *IPancakeFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IPancakeFactory *IPancakeFactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.CreatePair(&_IPancakeFactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IPancakeFactory *IPancakeFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.CreatePair(&_IPancakeFactory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IPancakeFactory *IPancakeFactoryTransactor) SetFeeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.contract.Transact(opts, "setFeeTo", arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IPancakeFactory *IPancakeFactorySession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.SetFeeTo(&_IPancakeFactory.TransactOpts, arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IPancakeFactory *IPancakeFactoryTransactorSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.SetFeeTo(&_IPancakeFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IPancakeFactory *IPancakeFactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.contract.Transact(opts, "setFeeToSetter", arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IPancakeFactory *IPancakeFactorySession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.SetFeeToSetter(&_IPancakeFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IPancakeFactory *IPancakeFactoryTransactorSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.SetFeeToSetter(&_IPancakeFactory.TransactOpts, arg0)
}

// IPancakeFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the IPancakeFactory contract.
type IPancakeFactoryPairCreatedIterator struct {
	Event *IPancakeFactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *IPancakeFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakeFactoryPairCreated)
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
		it.Event = new(IPancakeFactoryPairCreated)
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
func (it *IPancakeFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakeFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakeFactoryPairCreated represents a PairCreated event raised by the IPancakeFactory contract.
type IPancakeFactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IPancakeFactory *IPancakeFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*IPancakeFactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _IPancakeFactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryPairCreatedIterator{contract: _IPancakeFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IPancakeFactory *IPancakeFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *IPancakeFactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _IPancakeFactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakeFactoryPairCreated)
				if err := _IPancakeFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IPancakeFactory *IPancakeFactoryFilterer) ParsePairCreated(log types.Log) (*IPancakeFactoryPairCreated, error) {
	event := new(IPancakeFactoryPairCreated)
	if err := _IPancakeFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairMetaData contains all meta data concerning the IPancakePair contract.
var IPancakePairMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3644e515": "DOMAIN_SEPARATOR()",
		"ba9a7a56": "MINIMUM_LIQUIDITY()",
		"30adf81f": "PERMIT_TYPEHASH()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"89afcb44": "burn(address)",
		"313ce567": "decimals()",
		"c45a0155": "factory()",
		"0902f1ac": "getReserves()",
		"485cc955": "initialize(address,address)",
		"7464fc3d": "kLast()",
		"6a627842": "mint(address)",
		"06fdde03": "name()",
		"7ecebe00": "nonces(address)",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"5909c0d5": "price0CumulativeLast()",
		"5a3d5493": "price1CumulativeLast()",
		"bc25cf77": "skim(address)",
		"022c0d9f": "swap(uint256,uint256,address,bytes)",
		"95d89b41": "symbol()",
		"fff6cae9": "sync()",
		"0dfe1681": "token0()",
		"d21220a7": "token1()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IPancakePairABI is the input ABI used to generate the binding from.
// Deprecated: Use IPancakePairMetaData.ABI instead.
var IPancakePairABI = IPancakePairMetaData.ABI

// Deprecated: Use IPancakePairMetaData.Sigs instead.
// IPancakePairFuncSigs maps the 4-byte function signature to its string representation.
var IPancakePairFuncSigs = IPancakePairMetaData.Sigs

// IPancakePair is an auto generated Go binding around an Ethereum contract.
type IPancakePair struct {
	IPancakePairCaller     // Read-only binding to the contract
	IPancakePairTransactor // Write-only binding to the contract
	IPancakePairFilterer   // Log filterer for contract events
}

// IPancakePairCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakePairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakePairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakePairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakePairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakePairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakePairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakePairSession struct {
	Contract     *IPancakePair     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakePairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakePairCallerSession struct {
	Contract *IPancakePairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IPancakePairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakePairTransactorSession struct {
	Contract     *IPancakePairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPancakePairRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakePairRaw struct {
	Contract *IPancakePair // Generic contract binding to access the raw methods on
}

// IPancakePairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakePairCallerRaw struct {
	Contract *IPancakePairCaller // Generic read-only contract binding to access the raw methods on
}

// IPancakePairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakePairTransactorRaw struct {
	Contract *IPancakePairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakePair creates a new instance of IPancakePair, bound to a specific deployed contract.
func NewIPancakePair(address common.Address, backend bind.ContractBackend) (*IPancakePair, error) {
	contract, err := bindIPancakePair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakePair{IPancakePairCaller: IPancakePairCaller{contract: contract}, IPancakePairTransactor: IPancakePairTransactor{contract: contract}, IPancakePairFilterer: IPancakePairFilterer{contract: contract}}, nil
}

// NewIPancakePairCaller creates a new read-only instance of IPancakePair, bound to a specific deployed contract.
func NewIPancakePairCaller(address common.Address, caller bind.ContractCaller) (*IPancakePairCaller, error) {
	contract, err := bindIPancakePair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakePairCaller{contract: contract}, nil
}

// NewIPancakePairTransactor creates a new write-only instance of IPancakePair, bound to a specific deployed contract.
func NewIPancakePairTransactor(address common.Address, transactor bind.ContractTransactor) (*IPancakePairTransactor, error) {
	contract, err := bindIPancakePair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakePairTransactor{contract: contract}, nil
}

// NewIPancakePairFilterer creates a new log filterer instance of IPancakePair, bound to a specific deployed contract.
func NewIPancakePairFilterer(address common.Address, filterer bind.ContractFilterer) (*IPancakePairFilterer, error) {
	contract, err := bindIPancakePair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakePairFilterer{contract: contract}, nil
}

// bindIPancakePair binds a generic wrapper to an already deployed contract.
func bindIPancakePair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPancakePairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakePair *IPancakePairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakePair.Contract.IPancakePairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakePair *IPancakePairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakePair.Contract.IPancakePairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakePair *IPancakePairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakePair.Contract.IPancakePairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakePair *IPancakePairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakePair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakePair *IPancakePairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakePair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakePair *IPancakePairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakePair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IPancakePair *IPancakePairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IPancakePair *IPancakePairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IPancakePair.Contract.DOMAINSEPARATOR(&_IPancakePair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IPancakePair *IPancakePairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IPancakePair.Contract.DOMAINSEPARATOR(&_IPancakePair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_IPancakePair *IPancakePairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_IPancakePair *IPancakePairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _IPancakePair.Contract.MINIMUMLIQUIDITY(&_IPancakePair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _IPancakePair.Contract.MINIMUMLIQUIDITY(&_IPancakePair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IPancakePair *IPancakePairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IPancakePair *IPancakePairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IPancakePair.Contract.PERMITTYPEHASH(&_IPancakePair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IPancakePair *IPancakePairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IPancakePair.Contract.PERMITTYPEHASH(&_IPancakePair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IPancakePair *IPancakePairCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IPancakePair *IPancakePairSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.Allowance(&_IPancakePair.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.Allowance(&_IPancakePair.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.BalanceOf(&_IPancakePair.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.BalanceOf(&_IPancakePair.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IPancakePair *IPancakePairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IPancakePair *IPancakePairSession) Decimals() (uint8, error) {
	return _IPancakePair.Contract.Decimals(&_IPancakePair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IPancakePair *IPancakePairCallerSession) Decimals() (uint8, error) {
	return _IPancakePair.Contract.Decimals(&_IPancakePair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IPancakePair *IPancakePairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IPancakePair *IPancakePairSession) Factory() (common.Address, error) {
	return _IPancakePair.Contract.Factory(&_IPancakePair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IPancakePair *IPancakePairCallerSession) Factory() (common.Address, error) {
	return _IPancakePair.Contract.Factory(&_IPancakePair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_IPancakePair *IPancakePairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_IPancakePair *IPancakePairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _IPancakePair.Contract.GetReserves(&_IPancakePair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_IPancakePair *IPancakePairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _IPancakePair.Contract.GetReserves(&_IPancakePair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_IPancakePair *IPancakePairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_IPancakePair *IPancakePairSession) KLast() (*big.Int, error) {
	return _IPancakePair.Contract.KLast(&_IPancakePair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) KLast() (*big.Int, error) {
	return _IPancakePair.Contract.KLast(&_IPancakePair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IPancakePair *IPancakePairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IPancakePair *IPancakePairSession) Name() (string, error) {
	return _IPancakePair.Contract.Name(&_IPancakePair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IPancakePair *IPancakePairCallerSession) Name() (string, error) {
	return _IPancakePair.Contract.Name(&_IPancakePair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.Nonces(&_IPancakePair.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IPancakePair.Contract.Nonces(&_IPancakePair.CallOpts, owner)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairSession) Price0CumulativeLast() (*big.Int, error) {
	return _IPancakePair.Contract.Price0CumulativeLast(&_IPancakePair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _IPancakePair.Contract.Price0CumulativeLast(&_IPancakePair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairSession) Price1CumulativeLast() (*big.Int, error) {
	return _IPancakePair.Contract.Price1CumulativeLast(&_IPancakePair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _IPancakePair.Contract.Price1CumulativeLast(&_IPancakePair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IPancakePair *IPancakePairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IPancakePair *IPancakePairSession) Symbol() (string, error) {
	return _IPancakePair.Contract.Symbol(&_IPancakePair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IPancakePair *IPancakePairCallerSession) Symbol() (string, error) {
	return _IPancakePair.Contract.Symbol(&_IPancakePair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IPancakePair *IPancakePairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IPancakePair *IPancakePairSession) Token0() (common.Address, error) {
	return _IPancakePair.Contract.Token0(&_IPancakePair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IPancakePair *IPancakePairCallerSession) Token0() (common.Address, error) {
	return _IPancakePair.Contract.Token0(&_IPancakePair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IPancakePair *IPancakePairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IPancakePair *IPancakePairSession) Token1() (common.Address, error) {
	return _IPancakePair.Contract.Token1(&_IPancakePair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IPancakePair *IPancakePairCallerSession) Token1() (common.Address, error) {
	return _IPancakePair.Contract.Token1(&_IPancakePair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IPancakePair *IPancakePairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPancakePair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IPancakePair *IPancakePairSession) TotalSupply() (*big.Int, error) {
	return _IPancakePair.Contract.TotalSupply(&_IPancakePair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IPancakePair *IPancakePairCallerSession) TotalSupply() (*big.Int, error) {
	return _IPancakePair.Contract.TotalSupply(&_IPancakePair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.Approve(&_IPancakePair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.Approve(&_IPancakePair.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Burn(&_IPancakePair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Burn(&_IPancakePair.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_IPancakePair *IPancakePairTransactor) Initialize(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "initialize", arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_IPancakePair *IPancakePairSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Initialize(&_IPancakePair.TransactOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_IPancakePair *IPancakePairTransactorSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Initialize(&_IPancakePair.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_IPancakePair *IPancakePairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_IPancakePair *IPancakePairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Mint(&_IPancakePair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_IPancakePair *IPancakePairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Mint(&_IPancakePair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IPancakePair *IPancakePairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IPancakePair *IPancakePairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakePair.Contract.Permit(&_IPancakePair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IPancakePair *IPancakePairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakePair.Contract.Permit(&_IPancakePair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IPancakePair *IPancakePairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IPancakePair *IPancakePairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Skim(&_IPancakePair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IPancakePair *IPancakePairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _IPancakePair.Contract.Skim(&_IPancakePair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IPancakePair *IPancakePairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IPancakePair *IPancakePairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IPancakePair.Contract.Swap(&_IPancakePair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IPancakePair *IPancakePairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IPancakePair.Contract.Swap(&_IPancakePair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_IPancakePair *IPancakePairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_IPancakePair *IPancakePairSession) Sync() (*types.Transaction, error) {
	return _IPancakePair.Contract.Sync(&_IPancakePair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_IPancakePair *IPancakePairTransactorSession) Sync() (*types.Transaction, error) {
	return _IPancakePair.Contract.Sync(&_IPancakePair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.Transfer(&_IPancakePair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.Transfer(&_IPancakePair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.TransferFrom(&_IPancakePair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IPancakePair *IPancakePairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPancakePair.Contract.TransferFrom(&_IPancakePair.TransactOpts, from, to, value)
}

// IPancakePairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IPancakePair contract.
type IPancakePairApprovalIterator struct {
	Event *IPancakePairApproval // Event containing the contract specifics and raw log

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
func (it *IPancakePairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairApproval)
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
		it.Event = new(IPancakePairApproval)
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
func (it *IPancakePairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairApproval represents a Approval event raised by the IPancakePair contract.
type IPancakePairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IPancakePair *IPancakePairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IPancakePairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IPancakePairApprovalIterator{contract: _IPancakePair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IPancakePair *IPancakePairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IPancakePairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairApproval)
				if err := _IPancakePair.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IPancakePair *IPancakePairFilterer) ParseApproval(log types.Log) (*IPancakePairApproval, error) {
	event := new(IPancakePairApproval)
	if err := _IPancakePair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IPancakePair contract.
type IPancakePairBurnIterator struct {
	Event *IPancakePairBurn // Event containing the contract specifics and raw log

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
func (it *IPancakePairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairBurn)
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
		it.Event = new(IPancakePairBurn)
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
func (it *IPancakePairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairBurn represents a Burn event raised by the IPancakePair contract.
type IPancakePairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_IPancakePair *IPancakePairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IPancakePairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPancakePairBurnIterator{contract: _IPancakePair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_IPancakePair *IPancakePairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IPancakePairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairBurn)
				if err := _IPancakePair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_IPancakePair *IPancakePairFilterer) ParseBurn(log types.Log) (*IPancakePairBurn, error) {
	event := new(IPancakePairBurn)
	if err := _IPancakePair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IPancakePair contract.
type IPancakePairMintIterator struct {
	Event *IPancakePairMint // Event containing the contract specifics and raw log

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
func (it *IPancakePairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairMint)
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
		it.Event = new(IPancakePairMint)
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
func (it *IPancakePairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairMint represents a Mint event raised by the IPancakePair contract.
type IPancakePairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*IPancakePairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &IPancakePairMintIterator{contract: _IPancakePair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IPancakePairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairMint)
				if err := _IPancakePair.contract.UnpackLog(event, "Mint", log); err != nil {
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
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPancakePair *IPancakePairFilterer) ParseMint(log types.Log) (*IPancakePairMint, error) {
	event := new(IPancakePairMint)
	if err := _IPancakePair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the IPancakePair contract.
type IPancakePairSwapIterator struct {
	Event *IPancakePairSwap // Event containing the contract specifics and raw log

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
func (it *IPancakePairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairSwap)
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
		it.Event = new(IPancakePairSwap)
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
func (it *IPancakePairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairSwap represents a Swap event raised by the IPancakePair contract.
type IPancakePairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_IPancakePair *IPancakePairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IPancakePairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPancakePairSwapIterator{contract: _IPancakePair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_IPancakePair *IPancakePairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *IPancakePairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairSwap)
				if err := _IPancakePair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_IPancakePair *IPancakePairFilterer) ParseSwap(log types.Log) (*IPancakePairSwap, error) {
	event := new(IPancakePairSwap)
	if err := _IPancakePair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the IPancakePair contract.
type IPancakePairSyncIterator struct {
	Event *IPancakePairSync // Event containing the contract specifics and raw log

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
func (it *IPancakePairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairSync)
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
		it.Event = new(IPancakePairSync)
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
func (it *IPancakePairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairSync represents a Sync event raised by the IPancakePair contract.
type IPancakePairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_IPancakePair *IPancakePairFilterer) FilterSync(opts *bind.FilterOpts) (*IPancakePairSyncIterator, error) {

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &IPancakePairSyncIterator{contract: _IPancakePair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_IPancakePair *IPancakePairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *IPancakePairSync) (event.Subscription, error) {

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairSync)
				if err := _IPancakePair.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_IPancakePair *IPancakePairFilterer) ParseSync(log types.Log) (*IPancakePairSync, error) {
	event := new(IPancakePairSync)
	if err := _IPancakePair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakePairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IPancakePair contract.
type IPancakePairTransferIterator struct {
	Event *IPancakePairTransfer // Event containing the contract specifics and raw log

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
func (it *IPancakePairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPancakePairTransfer)
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
		it.Event = new(IPancakePairTransfer)
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
func (it *IPancakePairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPancakePairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPancakePairTransfer represents a Transfer event raised by the IPancakePair contract.
type IPancakePairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IPancakePair *IPancakePairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IPancakePairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPancakePairTransferIterator{contract: _IPancakePair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IPancakePair *IPancakePairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IPancakePairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPancakePair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPancakePairTransfer)
				if err := _IPancakePair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IPancakePair *IPancakePairFilterer) ParseTransfer(log types.Log) (*IPancakePairTransfer, error) {
	event := new(IPancakePairTransfer)
	if err := _IPancakePair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPancakeRouter01MetaData contains all meta data concerning the IPancakeRouter01 contract.
var IPancakeRouter01MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ad5c4648": "WETH()",
		"e8e33700": "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)",
		"f305d719": "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
		"c45a0155": "factory()",
		"85f8c259": "getAmountIn(uint256,uint256,uint256)",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"1f00ca74": "getAmountsIn(uint256,address[])",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"ad615dec": "quote(uint256,uint256,uint256)",
		"baa2abde": "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)",
		"02751cec": "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
		"ded9382a": "removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"2195995c": "removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"fb3bdb41": "swapETHForExactTokens(uint256,address[],address,uint256)",
		"7ff36ab5": "swapExactETHForTokens(uint256,address[],address,uint256)",
		"18cbafe5": "swapExactTokensForETH(uint256,uint256,address[],address,uint256)",
		"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
		"4a25d94a": "swapTokensForExactETH(uint256,uint256,address[],address,uint256)",
		"8803dbee": "swapTokensForExactTokens(uint256,uint256,address[],address,uint256)",
	},
}

// IPancakeRouter01ABI is the input ABI used to generate the binding from.
// Deprecated: Use IPancakeRouter01MetaData.ABI instead.
var IPancakeRouter01ABI = IPancakeRouter01MetaData.ABI

// Deprecated: Use IPancakeRouter01MetaData.Sigs instead.
// IPancakeRouter01FuncSigs maps the 4-byte function signature to its string representation.
var IPancakeRouter01FuncSigs = IPancakeRouter01MetaData.Sigs

// IPancakeRouter01 is an auto generated Go binding around an Ethereum contract.
type IPancakeRouter01 struct {
	IPancakeRouter01Caller     // Read-only binding to the contract
	IPancakeRouter01Transactor // Write-only binding to the contract
	IPancakeRouter01Filterer   // Log filterer for contract events
}

// IPancakeRouter01Caller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakeRouter01Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter01Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakeRouter01Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter01Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakeRouter01Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter01Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakeRouter01Session struct {
	Contract     *IPancakeRouter01 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakeRouter01CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakeRouter01CallerSession struct {
	Contract *IPancakeRouter01Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IPancakeRouter01TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakeRouter01TransactorSession struct {
	Contract     *IPancakeRouter01Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IPancakeRouter01Raw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakeRouter01Raw struct {
	Contract *IPancakeRouter01 // Generic contract binding to access the raw methods on
}

// IPancakeRouter01CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakeRouter01CallerRaw struct {
	Contract *IPancakeRouter01Caller // Generic read-only contract binding to access the raw methods on
}

// IPancakeRouter01TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakeRouter01TransactorRaw struct {
	Contract *IPancakeRouter01Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakeRouter01 creates a new instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01(address common.Address, backend bind.ContractBackend) (*IPancakeRouter01, error) {
	contract, err := bindIPancakeRouter01(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01{IPancakeRouter01Caller: IPancakeRouter01Caller{contract: contract}, IPancakeRouter01Transactor: IPancakeRouter01Transactor{contract: contract}, IPancakeRouter01Filterer: IPancakeRouter01Filterer{contract: contract}}, nil
}

// NewIPancakeRouter01Caller creates a new read-only instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01Caller(address common.Address, caller bind.ContractCaller) (*IPancakeRouter01Caller, error) {
	contract, err := bindIPancakeRouter01(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01Caller{contract: contract}, nil
}

// NewIPancakeRouter01Transactor creates a new write-only instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01Transactor(address common.Address, transactor bind.ContractTransactor) (*IPancakeRouter01Transactor, error) {
	contract, err := bindIPancakeRouter01(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01Transactor{contract: contract}, nil
}

// NewIPancakeRouter01Filterer creates a new log filterer instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01Filterer(address common.Address, filterer bind.ContractFilterer) (*IPancakeRouter01Filterer, error) {
	contract, err := bindIPancakeRouter01(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01Filterer{contract: contract}, nil
}

// bindIPancakeRouter01 binds a generic wrapper to an already deployed contract.
func bindIPancakeRouter01(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPancakeRouter01ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter01 *IPancakeRouter01Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter01.Contract.IPancakeRouter01Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter01 *IPancakeRouter01Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.IPancakeRouter01Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter01 *IPancakeRouter01Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.IPancakeRouter01Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter01 *IPancakeRouter01CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter01.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter01 *IPancakeRouter01TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter01 *IPancakeRouter01TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01Session) WETH() (common.Address, error) {
	return _IPancakeRouter01.Contract.WETH(&_IPancakeRouter01.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) WETH() (common.Address, error) {
	return _IPancakeRouter01.Contract.WETH(&_IPancakeRouter01.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01Session) Factory() (common.Address, error) {
	return _IPancakeRouter01.Contract.Factory(&_IPancakeRouter01.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) Factory() (common.Address, error) {
	return _IPancakeRouter01.Contract.Factory(&_IPancakeRouter01.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter01 *IPancakeRouter01Caller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter01 *IPancakeRouter01Session) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountIn(&_IPancakeRouter01.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountIn(&_IPancakeRouter01.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter01 *IPancakeRouter01Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter01 *IPancakeRouter01Session) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountOut(&_IPancakeRouter01.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountOut(&_IPancakeRouter01.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountsIn(&_IPancakeRouter01.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountsIn(&_IPancakeRouter01.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountsOut(&_IPancakeRouter01.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountsOut(&_IPancakeRouter01.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Caller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Session) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.Quote(&_IPancakeRouter01.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IPancakeRouter01.Contract.Quote(&_IPancakeRouter01.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.AddLiquidity(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.AddLiquidity(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01Session) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.AddLiquidityETH(&_IPancakeRouter01.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.AddLiquidityETH(&_IPancakeRouter01.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidity(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidity(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01Session) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityETH(&_IPancakeRouter01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityETH(&_IPancakeRouter01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) RemoveLiquidityETHWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01Session) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityETHWithPermit(&_IPancakeRouter01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityETHWithPermit(&_IPancakeRouter01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityWithPermit(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.RemoveLiquidityWithPermit(&_IPancakeRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapETHForExactTokens", amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapETHForExactTokens(&_IPancakeRouter01.TransactOpts, amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapETHForExactTokens(&_IPancakeRouter01.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactETHForTokens(&_IPancakeRouter01.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactETHForTokens(&_IPancakeRouter01.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactTokensForETH(&_IPancakeRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactTokensForETH(&_IPancakeRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactTokensForTokens(&_IPancakeRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapExactTokensForTokens(&_IPancakeRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapTokensForExactETH(&_IPancakeRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapTokensForExactETH(&_IPancakeRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapTokensForExactTokens(&_IPancakeRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.SwapTokensForExactTokens(&_IPancakeRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// IPancakeRouter02MetaData contains all meta data concerning the IPancakeRouter02 contract.
var IPancakeRouter02MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ad5c4648": "WETH()",
		"e8e33700": "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)",
		"f305d719": "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
		"c45a0155": "factory()",
		"85f8c259": "getAmountIn(uint256,uint256,uint256)",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"1f00ca74": "getAmountsIn(uint256,address[])",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"ad615dec": "quote(uint256,uint256,uint256)",
		"baa2abde": "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)",
		"02751cec": "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
		"af2979eb": "removeLiquidityETHSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256)",
		"ded9382a": "removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"5b0d5984": "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"2195995c": "removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"fb3bdb41": "swapETHForExactTokens(uint256,address[],address,uint256)",
		"7ff36ab5": "swapExactETHForTokens(uint256,address[],address,uint256)",
		"b6f9de95": "swapExactETHForTokensSupportingFeeOnTransferTokens(uint256,address[],address,uint256)",
		"18cbafe5": "swapExactTokensForETH(uint256,uint256,address[],address,uint256)",
		"791ac947": "swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
		"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
		"5c11d795": "swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
		"4a25d94a": "swapTokensForExactETH(uint256,uint256,address[],address,uint256)",
		"8803dbee": "swapTokensForExactTokens(uint256,uint256,address[],address,uint256)",
	},
}

// IPancakeRouter02ABI is the input ABI used to generate the binding from.
// Deprecated: Use IPancakeRouter02MetaData.ABI instead.
var IPancakeRouter02ABI = IPancakeRouter02MetaData.ABI

// Deprecated: Use IPancakeRouter02MetaData.Sigs instead.
// IPancakeRouter02FuncSigs maps the 4-byte function signature to its string representation.
var IPancakeRouter02FuncSigs = IPancakeRouter02MetaData.Sigs

// IPancakeRouter02 is an auto generated Go binding around an Ethereum contract.
type IPancakeRouter02 struct {
	IPancakeRouter02Caller     // Read-only binding to the contract
	IPancakeRouter02Transactor // Write-only binding to the contract
	IPancakeRouter02Filterer   // Log filterer for contract events
}

// IPancakeRouter02Caller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakeRouter02Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter02Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakeRouter02Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter02Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakeRouter02Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter02Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakeRouter02Session struct {
	Contract     *IPancakeRouter02 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakeRouter02CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakeRouter02CallerSession struct {
	Contract *IPancakeRouter02Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IPancakeRouter02TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakeRouter02TransactorSession struct {
	Contract     *IPancakeRouter02Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IPancakeRouter02Raw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakeRouter02Raw struct {
	Contract *IPancakeRouter02 // Generic contract binding to access the raw methods on
}

// IPancakeRouter02CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakeRouter02CallerRaw struct {
	Contract *IPancakeRouter02Caller // Generic read-only contract binding to access the raw methods on
}

// IPancakeRouter02TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakeRouter02TransactorRaw struct {
	Contract *IPancakeRouter02Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakeRouter02 creates a new instance of IPancakeRouter02, bound to a specific deployed contract.
func NewIPancakeRouter02(address common.Address, backend bind.ContractBackend) (*IPancakeRouter02, error) {
	contract, err := bindIPancakeRouter02(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter02{IPancakeRouter02Caller: IPancakeRouter02Caller{contract: contract}, IPancakeRouter02Transactor: IPancakeRouter02Transactor{contract: contract}, IPancakeRouter02Filterer: IPancakeRouter02Filterer{contract: contract}}, nil
}

// NewIPancakeRouter02Caller creates a new read-only instance of IPancakeRouter02, bound to a specific deployed contract.
func NewIPancakeRouter02Caller(address common.Address, caller bind.ContractCaller) (*IPancakeRouter02Caller, error) {
	contract, err := bindIPancakeRouter02(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter02Caller{contract: contract}, nil
}

// NewIPancakeRouter02Transactor creates a new write-only instance of IPancakeRouter02, bound to a specific deployed contract.
func NewIPancakeRouter02Transactor(address common.Address, transactor bind.ContractTransactor) (*IPancakeRouter02Transactor, error) {
	contract, err := bindIPancakeRouter02(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter02Transactor{contract: contract}, nil
}

// NewIPancakeRouter02Filterer creates a new log filterer instance of IPancakeRouter02, bound to a specific deployed contract.
func NewIPancakeRouter02Filterer(address common.Address, filterer bind.ContractFilterer) (*IPancakeRouter02Filterer, error) {
	contract, err := bindIPancakeRouter02(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter02Filterer{contract: contract}, nil
}

// bindIPancakeRouter02 binds a generic wrapper to an already deployed contract.
func bindIPancakeRouter02(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPancakeRouter02ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter02 *IPancakeRouter02Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter02.Contract.IPancakeRouter02Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter02 *IPancakeRouter02Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.IPancakeRouter02Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter02 *IPancakeRouter02Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.IPancakeRouter02Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter02 *IPancakeRouter02CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter02.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter02 *IPancakeRouter02TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter02 *IPancakeRouter02TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02Session) WETH() (common.Address, error) {
	return _IPancakeRouter02.Contract.WETH(&_IPancakeRouter02.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) WETH() (common.Address, error) {
	return _IPancakeRouter02.Contract.WETH(&_IPancakeRouter02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02Session) Factory() (common.Address, error) {
	return _IPancakeRouter02.Contract.Factory(&_IPancakeRouter02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) Factory() (common.Address, error) {
	return _IPancakeRouter02.Contract.Factory(&_IPancakeRouter02.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter02 *IPancakeRouter02Caller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter02 *IPancakeRouter02Session) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountIn(&_IPancakeRouter02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountIn(&_IPancakeRouter02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter02 *IPancakeRouter02Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter02 *IPancakeRouter02Session) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountOut(&_IPancakeRouter02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountOut(&_IPancakeRouter02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountsIn(&_IPancakeRouter02.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountsIn(&_IPancakeRouter02.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountsOut(&_IPancakeRouter02.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter02.Contract.GetAmountsOut(&_IPancakeRouter02.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Caller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter02.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Session) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.Quote(&_IPancakeRouter02.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02CallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IPancakeRouter02.Contract.Quote(&_IPancakeRouter02.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.AddLiquidity(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.AddLiquidity(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02Session) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.AddLiquidityETH(&_IPancakeRouter02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.AddLiquidityETH(&_IPancakeRouter02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidity(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidity(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETH(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETH(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidityETHWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHWithPermit(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHWithPermit(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityWithPermit(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.RemoveLiquidityWithPermit(&_IPancakeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapETHForExactTokens", amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapETHForExactTokens(&_IPancakeRouter02.TransactOpts, amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapETHForExactTokens(&_IPancakeRouter02.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactETHForTokens(&_IPancakeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactETHForTokens(&_IPancakeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForETH(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForETH(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IPancakeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapTokensForExactETH(&_IPancakeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapTokensForExactETH(&_IPancakeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapTokensForExactTokens(&_IPancakeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPancakeRouter02 *IPancakeRouter02TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter02.Contract.SwapTokensForExactTokens(&_IPancakeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// IVAIMetaData contains all meta data concerning the IVAI contract.
var IVAIMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractVTokenInterface\",\"name\":\"vTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateVAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"11b3d5e7": "liquidateVAI(address,uint256,address)",
	},
}

// IVAIABI is the input ABI used to generate the binding from.
// Deprecated: Use IVAIMetaData.ABI instead.
var IVAIABI = IVAIMetaData.ABI

// Deprecated: Use IVAIMetaData.Sigs instead.
// IVAIFuncSigs maps the 4-byte function signature to its string representation.
var IVAIFuncSigs = IVAIMetaData.Sigs

// IVAI is an auto generated Go binding around an Ethereum contract.
type IVAI struct {
	IVAICaller     // Read-only binding to the contract
	IVAITransactor // Write-only binding to the contract
	IVAIFilterer   // Log filterer for contract events
}

// IVAICaller is an auto generated read-only Go binding around an Ethereum contract.
type IVAICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVAITransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVAITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVAIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVAIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVAISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVAISession struct {
	Contract     *IVAI             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVAICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVAICallerSession struct {
	Contract *IVAICaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IVAITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVAITransactorSession struct {
	Contract     *IVAITransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVAIRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVAIRaw struct {
	Contract *IVAI // Generic contract binding to access the raw methods on
}

// IVAICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVAICallerRaw struct {
	Contract *IVAICaller // Generic read-only contract binding to access the raw methods on
}

// IVAITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVAITransactorRaw struct {
	Contract *IVAITransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVAI creates a new instance of IVAI, bound to a specific deployed contract.
func NewIVAI(address common.Address, backend bind.ContractBackend) (*IVAI, error) {
	contract, err := bindIVAI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVAI{IVAICaller: IVAICaller{contract: contract}, IVAITransactor: IVAITransactor{contract: contract}, IVAIFilterer: IVAIFilterer{contract: contract}}, nil
}

// NewIVAICaller creates a new read-only instance of IVAI, bound to a specific deployed contract.
func NewIVAICaller(address common.Address, caller bind.ContractCaller) (*IVAICaller, error) {
	contract, err := bindIVAI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVAICaller{contract: contract}, nil
}

// NewIVAITransactor creates a new write-only instance of IVAI, bound to a specific deployed contract.
func NewIVAITransactor(address common.Address, transactor bind.ContractTransactor) (*IVAITransactor, error) {
	contract, err := bindIVAI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVAITransactor{contract: contract}, nil
}

// NewIVAIFilterer creates a new log filterer instance of IVAI, bound to a specific deployed contract.
func NewIVAIFilterer(address common.Address, filterer bind.ContractFilterer) (*IVAIFilterer, error) {
	contract, err := bindIVAI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVAIFilterer{contract: contract}, nil
}

// bindIVAI binds a generic wrapper to an already deployed contract.
func bindIVAI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVAIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVAI *IVAIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVAI.Contract.IVAICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVAI *IVAIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVAI.Contract.IVAITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVAI *IVAIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVAI.Contract.IVAITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVAI *IVAICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVAI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVAI *IVAITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVAI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVAI *IVAITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVAI.Contract.contract.Transact(opts, method, params...)
}

// LiquidateVAI is a paid mutator transaction binding the contract method 0x11b3d5e7.
//
// Solidity: function liquidateVAI(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256, uint256)
func (_IVAI *IVAITransactor) LiquidateVAI(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _IVAI.contract.Transact(opts, "liquidateVAI", borrower, repayAmount, vTokenCollateral)
}

// LiquidateVAI is a paid mutator transaction binding the contract method 0x11b3d5e7.
//
// Solidity: function liquidateVAI(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256, uint256)
func (_IVAI *IVAISession) LiquidateVAI(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _IVAI.Contract.LiquidateVAI(&_IVAI.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// LiquidateVAI is a paid mutator transaction binding the contract method 0x11b3d5e7.
//
// Solidity: function liquidateVAI(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256, uint256)
func (_IVAI *IVAITransactorSession) LiquidateVAI(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _IVAI.Contract.LiquidateVAI(&_IVAI.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// IVBNBMetaData contains all meta data concerning the IVBNB contract.
var IVBNBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isVToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repayBorrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e9c714f2": "_acceptAdmin()",
		"601a0bf1": "_reduceReserves(uint256)",
		"4576b5db": "_setComptroller(address)",
		"f2b3abbd": "_setInterestRateModel(address)",
		"b71d1a0c": "_setPendingAdmin(address)",
		"fca7820b": "_setReserveFactor(uint256)",
		"6c540baf": "accrualBlockNumber()",
		"a6afed95": "accrueInterest()",
		"f851a440": "admin()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"3af9e669": "balanceOfUnderlying(address)",
		"c5ebeaec": "borrow(uint256)",
		"17bfdfbc": "borrowBalanceCurrent(address)",
		"95dd9193": "borrowBalanceStored(address)",
		"aa5af0fd": "borrowIndex()",
		"f8f9da28": "borrowRatePerBlock()",
		"5fe3b567": "comptroller()",
		"313ce567": "decimals()",
		"bd6d894d": "exchangeRateCurrent()",
		"182df0f5": "exchangeRateStored()",
		"c37f68e2": "getAccountSnapshot(address)",
		"3b1d21a2": "getCash()",
		"99d8c1b4": "initialize(address,address,uint256,string,string,uint8)",
		"f3fdb15a": "interestRateModel()",
		"3d9ea3a1": "isVToken()",
		"aae40a2a": "liquidateBorrow(address,address)",
		"1249c58b": "mint()",
		"06fdde03": "name()",
		"26782247": "pendingAdmin()",
		"db006a75": "redeem(uint256)",
		"852a12e3": "redeemUnderlying(uint256)",
		"4e4d9fea": "repayBorrow()",
		"e5974619": "repayBorrowBehalf(address)",
		"173b9904": "reserveFactorMantissa()",
		"b2a02ff1": "seize(address,address,uint256)",
		"ae9d70b0": "supplyRatePerBlock()",
		"95d89b41": "symbol()",
		"47bd3718": "totalBorrows()",
		"73acee98": "totalBorrowsCurrent()",
		"8f840ddd": "totalReserves()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IVBNBABI is the input ABI used to generate the binding from.
// Deprecated: Use IVBNBMetaData.ABI instead.
var IVBNBABI = IVBNBMetaData.ABI

// Deprecated: Use IVBNBMetaData.Sigs instead.
// IVBNBFuncSigs maps the 4-byte function signature to its string representation.
var IVBNBFuncSigs = IVBNBMetaData.Sigs

// IVBNB is an auto generated Go binding around an Ethereum contract.
type IVBNB struct {
	IVBNBCaller     // Read-only binding to the contract
	IVBNBTransactor // Write-only binding to the contract
	IVBNBFilterer   // Log filterer for contract events
}

// IVBNBCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVBNBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVBNBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVBNBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVBNBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVBNBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVBNBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVBNBSession struct {
	Contract     *IVBNB            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVBNBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVBNBCallerSession struct {
	Contract *IVBNBCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IVBNBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVBNBTransactorSession struct {
	Contract     *IVBNBTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVBNBRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVBNBRaw struct {
	Contract *IVBNB // Generic contract binding to access the raw methods on
}

// IVBNBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVBNBCallerRaw struct {
	Contract *IVBNBCaller // Generic read-only contract binding to access the raw methods on
}

// IVBNBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVBNBTransactorRaw struct {
	Contract *IVBNBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVBNB creates a new instance of IVBNB, bound to a specific deployed contract.
func NewIVBNB(address common.Address, backend bind.ContractBackend) (*IVBNB, error) {
	contract, err := bindIVBNB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVBNB{IVBNBCaller: IVBNBCaller{contract: contract}, IVBNBTransactor: IVBNBTransactor{contract: contract}, IVBNBFilterer: IVBNBFilterer{contract: contract}}, nil
}

// NewIVBNBCaller creates a new read-only instance of IVBNB, bound to a specific deployed contract.
func NewIVBNBCaller(address common.Address, caller bind.ContractCaller) (*IVBNBCaller, error) {
	contract, err := bindIVBNB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVBNBCaller{contract: contract}, nil
}

// NewIVBNBTransactor creates a new write-only instance of IVBNB, bound to a specific deployed contract.
func NewIVBNBTransactor(address common.Address, transactor bind.ContractTransactor) (*IVBNBTransactor, error) {
	contract, err := bindIVBNB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVBNBTransactor{contract: contract}, nil
}

// NewIVBNBFilterer creates a new log filterer instance of IVBNB, bound to a specific deployed contract.
func NewIVBNBFilterer(address common.Address, filterer bind.ContractFilterer) (*IVBNBFilterer, error) {
	contract, err := bindIVBNB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVBNBFilterer{contract: contract}, nil
}

// bindIVBNB binds a generic wrapper to an already deployed contract.
func bindIVBNB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVBNBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVBNB *IVBNBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVBNB.Contract.IVBNBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVBNB *IVBNBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVBNB.Contract.IVBNBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVBNB *IVBNBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVBNB.Contract.IVBNBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVBNB *IVBNBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVBNB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVBNB *IVBNBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVBNB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVBNB *IVBNBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVBNB.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_IVBNB *IVBNBCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_IVBNB *IVBNBSession) AccrualBlockNumber() (*big.Int, error) {
	return _IVBNB.Contract.AccrualBlockNumber(&_IVBNB.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_IVBNB *IVBNBCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _IVBNB.Contract.AccrualBlockNumber(&_IVBNB.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IVBNB *IVBNBCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IVBNB *IVBNBSession) Admin() (common.Address, error) {
	return _IVBNB.Contract.Admin(&_IVBNB.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IVBNB *IVBNBCallerSession) Admin() (common.Address, error) {
	return _IVBNB.Contract.Admin(&_IVBNB.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IVBNB *IVBNBCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IVBNB *IVBNBSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IVBNB.Contract.Allowance(&_IVBNB.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IVBNB *IVBNBCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IVBNB.Contract.Allowance(&_IVBNB.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IVBNB *IVBNBCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IVBNB *IVBNBSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IVBNB.Contract.BalanceOf(&_IVBNB.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IVBNB *IVBNBCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IVBNB.Contract.BalanceOf(&_IVBNB.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_IVBNB *IVBNBCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_IVBNB *IVBNBSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _IVBNB.Contract.BorrowBalanceStored(&_IVBNB.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_IVBNB *IVBNBCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _IVBNB.Contract.BorrowBalanceStored(&_IVBNB.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_IVBNB *IVBNBCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_IVBNB *IVBNBSession) BorrowIndex() (*big.Int, error) {
	return _IVBNB.Contract.BorrowIndex(&_IVBNB.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_IVBNB *IVBNBCallerSession) BorrowIndex() (*big.Int, error) {
	return _IVBNB.Contract.BorrowIndex(&_IVBNB.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_IVBNB *IVBNBCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_IVBNB *IVBNBSession) BorrowRatePerBlock() (*big.Int, error) {
	return _IVBNB.Contract.BorrowRatePerBlock(&_IVBNB.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_IVBNB *IVBNBCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _IVBNB.Contract.BorrowRatePerBlock(&_IVBNB.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_IVBNB *IVBNBCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_IVBNB *IVBNBSession) Comptroller() (common.Address, error) {
	return _IVBNB.Contract.Comptroller(&_IVBNB.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_IVBNB *IVBNBCallerSession) Comptroller() (common.Address, error) {
	return _IVBNB.Contract.Comptroller(&_IVBNB.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IVBNB *IVBNBCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IVBNB *IVBNBSession) Decimals() (uint8, error) {
	return _IVBNB.Contract.Decimals(&_IVBNB.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IVBNB *IVBNBCallerSession) Decimals() (uint8, error) {
	return _IVBNB.Contract.Decimals(&_IVBNB.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_IVBNB *IVBNBCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_IVBNB *IVBNBSession) ExchangeRateStored() (*big.Int, error) {
	return _IVBNB.Contract.ExchangeRateStored(&_IVBNB.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_IVBNB *IVBNBCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _IVBNB.Contract.ExchangeRateStored(&_IVBNB.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_IVBNB *IVBNBCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "getAccountSnapshot", account)

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
func (_IVBNB *IVBNBSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IVBNB.Contract.GetAccountSnapshot(&_IVBNB.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_IVBNB *IVBNBCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IVBNB.Contract.GetAccountSnapshot(&_IVBNB.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_IVBNB *IVBNBCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_IVBNB *IVBNBSession) GetCash() (*big.Int, error) {
	return _IVBNB.Contract.GetCash(&_IVBNB.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_IVBNB *IVBNBCallerSession) GetCash() (*big.Int, error) {
	return _IVBNB.Contract.GetCash(&_IVBNB.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_IVBNB *IVBNBCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_IVBNB *IVBNBSession) InterestRateModel() (common.Address, error) {
	return _IVBNB.Contract.InterestRateModel(&_IVBNB.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_IVBNB *IVBNBCallerSession) InterestRateModel() (common.Address, error) {
	return _IVBNB.Contract.InterestRateModel(&_IVBNB.CallOpts)
}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_IVBNB *IVBNBCaller) IsVToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "isVToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_IVBNB *IVBNBSession) IsVToken() (bool, error) {
	return _IVBNB.Contract.IsVToken(&_IVBNB.CallOpts)
}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_IVBNB *IVBNBCallerSession) IsVToken() (bool, error) {
	return _IVBNB.Contract.IsVToken(&_IVBNB.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IVBNB *IVBNBCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IVBNB *IVBNBSession) Name() (string, error) {
	return _IVBNB.Contract.Name(&_IVBNB.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IVBNB *IVBNBCallerSession) Name() (string, error) {
	return _IVBNB.Contract.Name(&_IVBNB.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_IVBNB *IVBNBCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_IVBNB *IVBNBSession) PendingAdmin() (common.Address, error) {
	return _IVBNB.Contract.PendingAdmin(&_IVBNB.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_IVBNB *IVBNBCallerSession) PendingAdmin() (common.Address, error) {
	return _IVBNB.Contract.PendingAdmin(&_IVBNB.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_IVBNB *IVBNBCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_IVBNB *IVBNBSession) ReserveFactorMantissa() (*big.Int, error) {
	return _IVBNB.Contract.ReserveFactorMantissa(&_IVBNB.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_IVBNB *IVBNBCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _IVBNB.Contract.ReserveFactorMantissa(&_IVBNB.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_IVBNB *IVBNBCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_IVBNB *IVBNBSession) SupplyRatePerBlock() (*big.Int, error) {
	return _IVBNB.Contract.SupplyRatePerBlock(&_IVBNB.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_IVBNB *IVBNBCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _IVBNB.Contract.SupplyRatePerBlock(&_IVBNB.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IVBNB *IVBNBCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IVBNB *IVBNBSession) Symbol() (string, error) {
	return _IVBNB.Contract.Symbol(&_IVBNB.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IVBNB *IVBNBCallerSession) Symbol() (string, error) {
	return _IVBNB.Contract.Symbol(&_IVBNB.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_IVBNB *IVBNBCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_IVBNB *IVBNBSession) TotalBorrows() (*big.Int, error) {
	return _IVBNB.Contract.TotalBorrows(&_IVBNB.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_IVBNB *IVBNBCallerSession) TotalBorrows() (*big.Int, error) {
	return _IVBNB.Contract.TotalBorrows(&_IVBNB.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_IVBNB *IVBNBCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_IVBNB *IVBNBSession) TotalReserves() (*big.Int, error) {
	return _IVBNB.Contract.TotalReserves(&_IVBNB.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_IVBNB *IVBNBCallerSession) TotalReserves() (*big.Int, error) {
	return _IVBNB.Contract.TotalReserves(&_IVBNB.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IVBNB *IVBNBCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVBNB.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IVBNB *IVBNBSession) TotalSupply() (*big.Int, error) {
	return _IVBNB.Contract.TotalSupply(&_IVBNB.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IVBNB *IVBNBCallerSession) TotalSupply() (*big.Int, error) {
	return _IVBNB.Contract.TotalSupply(&_IVBNB.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_IVBNB *IVBNBTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_IVBNB *IVBNBSession) AcceptAdmin() (*types.Transaction, error) {
	return _IVBNB.Contract.AcceptAdmin(&_IVBNB.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_IVBNB *IVBNBTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _IVBNB.Contract.AcceptAdmin(&_IVBNB.TransactOpts)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_IVBNB *IVBNBTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_IVBNB *IVBNBSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.ReduceReserves(&_IVBNB.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.ReduceReserves(&_IVBNB.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_IVBNB *IVBNBTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_IVBNB *IVBNBSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.SetComptroller(&_IVBNB.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.SetComptroller(&_IVBNB.TransactOpts, newComptroller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_IVBNB *IVBNBTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_IVBNB *IVBNBSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.SetInterestRateModel(&_IVBNB.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.SetInterestRateModel(&_IVBNB.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_IVBNB *IVBNBTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_IVBNB *IVBNBSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.SetPendingAdmin(&_IVBNB.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.SetPendingAdmin(&_IVBNB.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_IVBNB *IVBNBTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_IVBNB *IVBNBSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.SetReserveFactor(&_IVBNB.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.SetReserveFactor(&_IVBNB.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_IVBNB *IVBNBTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_IVBNB *IVBNBSession) AccrueInterest() (*types.Transaction, error) {
	return _IVBNB.Contract.AccrueInterest(&_IVBNB.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_IVBNB *IVBNBTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _IVBNB.Contract.AccrueInterest(&_IVBNB.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IVBNB *IVBNBTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IVBNB *IVBNBSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.Approve(&_IVBNB.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IVBNB *IVBNBTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.Approve(&_IVBNB.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_IVBNB *IVBNBTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_IVBNB *IVBNBSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.BalanceOfUnderlying(&_IVBNB.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.BalanceOfUnderlying(&_IVBNB.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_IVBNB *IVBNBTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_IVBNB *IVBNBSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.Borrow(&_IVBNB.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.Borrow(&_IVBNB.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_IVBNB *IVBNBTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_IVBNB *IVBNBSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.BorrowBalanceCurrent(&_IVBNB.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.BorrowBalanceCurrent(&_IVBNB.TransactOpts, account)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_IVBNB *IVBNBTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_IVBNB *IVBNBSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _IVBNB.Contract.ExchangeRateCurrent(&_IVBNB.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_IVBNB *IVBNBTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _IVBNB.Contract.ExchangeRateCurrent(&_IVBNB.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_IVBNB *IVBNBTransactor) Initialize(opts *bind.TransactOpts, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "initialize", comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_IVBNB *IVBNBSession) Initialize(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _IVBNB.Contract.Initialize(&_IVBNB.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_IVBNB *IVBNBTransactorSession) Initialize(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _IVBNB.Contract.Initialize(&_IVBNB.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xaae40a2a.
//
// Solidity: function liquidateBorrow(address borrower, address vTokenCollateral) payable returns()
func (_IVBNB *IVBNBTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "liquidateBorrow", borrower, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xaae40a2a.
//
// Solidity: function liquidateBorrow(address borrower, address vTokenCollateral) payable returns()
func (_IVBNB *IVBNBSession) LiquidateBorrow(borrower common.Address, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.LiquidateBorrow(&_IVBNB.TransactOpts, borrower, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xaae40a2a.
//
// Solidity: function liquidateBorrow(address borrower, address vTokenCollateral) payable returns()
func (_IVBNB *IVBNBTransactorSession) LiquidateBorrow(borrower common.Address, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.LiquidateBorrow(&_IVBNB.TransactOpts, borrower, vTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_IVBNB *IVBNBTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_IVBNB *IVBNBSession) Mint() (*types.Transaction, error) {
	return _IVBNB.Contract.Mint(&_IVBNB.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_IVBNB *IVBNBTransactorSession) Mint() (*types.Transaction, error) {
	return _IVBNB.Contract.Mint(&_IVBNB.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_IVBNB *IVBNBTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_IVBNB *IVBNBSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.Redeem(&_IVBNB.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.Redeem(&_IVBNB.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_IVBNB *IVBNBTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_IVBNB *IVBNBSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.RedeemUnderlying(&_IVBNB.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.RedeemUnderlying(&_IVBNB.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x4e4d9fea.
//
// Solidity: function repayBorrow() payable returns()
func (_IVBNB *IVBNBTransactor) RepayBorrow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "repayBorrow")
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x4e4d9fea.
//
// Solidity: function repayBorrow() payable returns()
func (_IVBNB *IVBNBSession) RepayBorrow() (*types.Transaction, error) {
	return _IVBNB.Contract.RepayBorrow(&_IVBNB.TransactOpts)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x4e4d9fea.
//
// Solidity: function repayBorrow() payable returns()
func (_IVBNB *IVBNBTransactorSession) RepayBorrow() (*types.Transaction, error) {
	return _IVBNB.Contract.RepayBorrow(&_IVBNB.TransactOpts)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xe5974619.
//
// Solidity: function repayBorrowBehalf(address borrower) payable returns()
func (_IVBNB *IVBNBTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "repayBorrowBehalf", borrower)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xe5974619.
//
// Solidity: function repayBorrowBehalf(address borrower) payable returns()
func (_IVBNB *IVBNBSession) RepayBorrowBehalf(borrower common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.RepayBorrowBehalf(&_IVBNB.TransactOpts, borrower)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xe5974619.
//
// Solidity: function repayBorrowBehalf(address borrower) payable returns()
func (_IVBNB *IVBNBTransactorSession) RepayBorrowBehalf(borrower common.Address) (*types.Transaction, error) {
	return _IVBNB.Contract.RepayBorrowBehalf(&_IVBNB.TransactOpts, borrower)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_IVBNB *IVBNBTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_IVBNB *IVBNBSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.Seize(&_IVBNB.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_IVBNB *IVBNBTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.Seize(&_IVBNB.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_IVBNB *IVBNBTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_IVBNB *IVBNBSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _IVBNB.Contract.TotalBorrowsCurrent(&_IVBNB.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_IVBNB *IVBNBTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _IVBNB.Contract.TotalBorrowsCurrent(&_IVBNB.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_IVBNB *IVBNBTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_IVBNB *IVBNBSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.Transfer(&_IVBNB.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_IVBNB *IVBNBTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.Transfer(&_IVBNB.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_IVBNB *IVBNBTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IVBNB.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_IVBNB *IVBNBSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.TransferFrom(&_IVBNB.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_IVBNB *IVBNBTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IVBNB.Contract.TransferFrom(&_IVBNB.TransactOpts, src, dst, amount)
}

// IWETHMetaData contains all meta data concerning the IWETH contract.
var IWETHMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d0e30db0": "deposit()",
		"a9059cbb": "transfer(address,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
	},
}

// IWETHABI is the input ABI used to generate the binding from.
// Deprecated: Use IWETHMetaData.ABI instead.
var IWETHABI = IWETHMetaData.ABI

// Deprecated: Use IWETHMetaData.Sigs instead.
// IWETHFuncSigs maps the 4-byte function signature to its string representation.
var IWETHFuncSigs = IWETHMetaData.Sigs

// IWETH is an auto generated Go binding around an Ethereum contract.
type IWETH struct {
	IWETHCaller     // Read-only binding to the contract
	IWETHTransactor // Write-only binding to the contract
	IWETHFilterer   // Log filterer for contract events
}

// IWETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETHSession struct {
	Contract     *IWETH            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETHCallerSession struct {
	Contract *IWETHCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETHTransactorSession struct {
	Contract     *IWETHTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWETHRaw struct {
	Contract *IWETH // Generic contract binding to access the raw methods on
}

// IWETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETHCallerRaw struct {
	Contract *IWETHCaller // Generic read-only contract binding to access the raw methods on
}

// IWETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETHTransactorRaw struct {
	Contract *IWETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH creates a new instance of IWETH, bound to a specific deployed contract.
func NewIWETH(address common.Address, backend bind.ContractBackend) (*IWETH, error) {
	contract, err := bindIWETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH{IWETHCaller: IWETHCaller{contract: contract}, IWETHTransactor: IWETHTransactor{contract: contract}, IWETHFilterer: IWETHFilterer{contract: contract}}, nil
}

// NewIWETHCaller creates a new read-only instance of IWETH, bound to a specific deployed contract.
func NewIWETHCaller(address common.Address, caller bind.ContractCaller) (*IWETHCaller, error) {
	contract, err := bindIWETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHCaller{contract: contract}, nil
}

// NewIWETHTransactor creates a new write-only instance of IWETH, bound to a specific deployed contract.
func NewIWETHTransactor(address common.Address, transactor bind.ContractTransactor) (*IWETHTransactor, error) {
	contract, err := bindIWETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHTransactor{contract: contract}, nil
}

// NewIWETHFilterer creates a new log filterer instance of IWETH, bound to a specific deployed contract.
func NewIWETHFilterer(address common.Address, filterer bind.ContractFilterer) (*IWETHFilterer, error) {
	contract, err := bindIWETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETHFilterer{contract: contract}, nil
}

// bindIWETH binds a generic wrapper to an already deployed contract.
func bindIWETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH *IWETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH.Contract.IWETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH *IWETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.Contract.IWETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH *IWETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH.Contract.IWETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH *IWETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH *IWETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH *IWETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHSession) Deposit() (*types.Transaction, error) {
	return _IWETH.Contract.Deposit(&_IWETH.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHTransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH.Contract.Deposit(&_IWETH.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Transfer(&_IWETH.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Transfer(&_IWETH.TransactOpts, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Withdraw(&_IWETH.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Withdraw(&_IWETH.TransactOpts, arg0)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeLibraryMetaData contains all meta data concerning the PancakeLibrary contract.
var PancakeLibraryMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122056908bfe6247bf6ea815fd460cac22b8408adb3007c2678ad3f09e1b8f62908664736f6c634300080b0033",
}

// PancakeLibraryABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeLibraryMetaData.ABI instead.
var PancakeLibraryABI = PancakeLibraryMetaData.ABI

// PancakeLibraryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PancakeLibraryMetaData.Bin instead.
var PancakeLibraryBin = PancakeLibraryMetaData.Bin

// DeployPancakeLibrary deploys a new Ethereum contract, binding an instance of PancakeLibrary to it.
func DeployPancakeLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PancakeLibrary, error) {
	parsed, err := PancakeLibraryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PancakeLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PancakeLibrary{PancakeLibraryCaller: PancakeLibraryCaller{contract: contract}, PancakeLibraryTransactor: PancakeLibraryTransactor{contract: contract}, PancakeLibraryFilterer: PancakeLibraryFilterer{contract: contract}}, nil
}

// PancakeLibrary is an auto generated Go binding around an Ethereum contract.
type PancakeLibrary struct {
	PancakeLibraryCaller     // Read-only binding to the contract
	PancakeLibraryTransactor // Write-only binding to the contract
	PancakeLibraryFilterer   // Log filterer for contract events
}

// PancakeLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeLibrarySession struct {
	Contract     *PancakeLibrary   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeLibraryCallerSession struct {
	Contract *PancakeLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PancakeLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeLibraryTransactorSession struct {
	Contract     *PancakeLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PancakeLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeLibraryRaw struct {
	Contract *PancakeLibrary // Generic contract binding to access the raw methods on
}

// PancakeLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeLibraryCallerRaw struct {
	Contract *PancakeLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeLibraryTransactorRaw struct {
	Contract *PancakeLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeLibrary creates a new instance of PancakeLibrary, bound to a specific deployed contract.
func NewPancakeLibrary(address common.Address, backend bind.ContractBackend) (*PancakeLibrary, error) {
	contract, err := bindPancakeLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeLibrary{PancakeLibraryCaller: PancakeLibraryCaller{contract: contract}, PancakeLibraryTransactor: PancakeLibraryTransactor{contract: contract}, PancakeLibraryFilterer: PancakeLibraryFilterer{contract: contract}}, nil
}

// NewPancakeLibraryCaller creates a new read-only instance of PancakeLibrary, bound to a specific deployed contract.
func NewPancakeLibraryCaller(address common.Address, caller bind.ContractCaller) (*PancakeLibraryCaller, error) {
	contract, err := bindPancakeLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeLibraryCaller{contract: contract}, nil
}

// NewPancakeLibraryTransactor creates a new write-only instance of PancakeLibrary, bound to a specific deployed contract.
func NewPancakeLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeLibraryTransactor, error) {
	contract, err := bindPancakeLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeLibraryTransactor{contract: contract}, nil
}

// NewPancakeLibraryFilterer creates a new log filterer instance of PancakeLibrary, bound to a specific deployed contract.
func NewPancakeLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeLibraryFilterer, error) {
	contract, err := bindPancakeLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeLibraryFilterer{contract: contract}, nil
}

// bindPancakeLibrary binds a generic wrapper to an already deployed contract.
func bindPancakeLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeLibrary *PancakeLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeLibrary.Contract.PancakeLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeLibrary *PancakeLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeLibrary.Contract.PancakeLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeLibrary *PancakeLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeLibrary.Contract.PancakeLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeLibrary *PancakeLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeLibrary *PancakeLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeLibrary *PancakeLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeLibrary.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209a8e18c7b5cae66d3c6a6db2c27301ec4db6368215da661bbb7c8e94b64bef9264736f6c634300080b0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// UniFlashSwapMetaData contains all meta data concerning the UniFlashSwap contract.
var UniFlashSwapMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scenarioNo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"repayUnderlyingToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seizedUnderlyingToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"flashLoanReturnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizedUnderlyingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"massProfit\",\"type\":\"uint256\"}],\"name\":\"Scenario\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SeizedUnderlyingTokenAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SeizedVTokenAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"WithdrawETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"mintedVAIS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"pancakeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_situation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_flashLoanFrom\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_path1\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_path2\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_flashLoanAmount\",\"type\":\"uint256\"}],\"name\":\"qingsuan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_flashLoanUnderlyingToken\",\"type\":\"address\"}],\"name\":\"swapOneBNBToFlashLoandUnderlyingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"f064b414": "mintedVAIS(address)",
		"8da5cb5b": "owner()",
		"84800812": "pancakeCall(address,uint256,uint256,bytes)",
		"b14e9a72": "qingsuan(uint256,address,address[],address[],address[],uint256)",
		"715018a6": "renounceOwnership()",
		"89c33748": "swapOneBNBToFlashLoandUnderlyingToken(address)",
		"f2fde38b": "transferOwnership(address)",
		"d9caed12": "withdraw(address,address,uint256)",
		"4782f779": "withdrawETH(address,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b506200001d3362000023565b62000073565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61395c80620000836000396000f3fe60806040526004361061008a5760003560e01c80638da5cb5b116100595780638da5cb5b1461010d578063b14e9a721461013a578063d9caed121461015a578063f064b4141461017a578063f2fde38b146101a857600080fd5b80634782f77914610096578063715018a6146100b857806384800812146100cd57806389c33748146100ed57600080fd5b3661009157005b600080fd5b3480156100a257600080fd5b506100b66100b1366004613219565b6101c8565b005b3480156100c457600080fd5b506100b661031f565b3480156100d957600080fd5b506100b66100e8366004613255565b610355565b3480156100f957600080fd5b506100b66101083660046132eb565b611667565b34801561011957600080fd5b506000546040516001600160a01b0390911681526020015b60405180910390f35b34801561014657600080fd5b506100b661015536600461334d565b6117cb565b34801561016657600080fd5b506100b661017536600461340b565b611ceb565b34801561018657600080fd5b5061019a6101953660046132eb565b611ed7565b604051908152602001610131565b3480156101b457600080fd5b506100b66101c33660046132eb565b611f5c565b6000546001600160a01b031633146101fb5760405162461bcd60e51b81526004016101f29061344c565b60405180910390fd5b6001600160a01b0382166102515760405162461bcd60e51b815260206004820152601b60248201527f746f206d757374206e6f74206265207a65726f2061646472657373000000000060448201526064016101f2565b600081116102a15760405162461bcd60e51b815260206004820152601c60248201527f616d6f756e74206d75737420626967676572207468616e207a65726f0000000060448201526064016101f2565b6040516001600160a01b0383169082156108fc029083906000818181858888f193505050501580156102d7573d6000803e3d6000fd5b50816001600160a01b03167f566e45b1c8057e725bf62796a7f1d37ae294393cab069725a09daddd1af98b798260405161031391815260200190565b60405180910390a25050565b6000546001600160a01b031633146103495760405162461bcd60e51b81526004016101f29061344c565b6103536000611ff7565b565b6103976040518060c001604052806000815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b6103a38284018461352e565b60a08701526080860152606085015260408401526001600160a01b03166020830181905290825233146104005760405162461bcd60e51b815260206004820152600560248201526410b830b4b960d91b60448201526064016101f2565b6001600160a01b03861630146104425760405162461bcd60e51b815260206004820152600760248201526610b9b2b73232b960c91b60448201526064016101f2565b6000816080015160038151811061045b5761045b6135db565b602002602001015190506000826080015160028151811061047e5761047e6135db565b6020026020010151905060006126f78460a00151601961049e9190613607565b6104a89190613626565b8460a001516104b79190613648565b6104c2906001613648565b9050600080600060608760000151600114156106d2576104ff88608001516000815181106104f2576104f26135db565b6020026020010151612047565b61053f5760405162461bcd60e51b81526020600482015260116024820152701896b737ba1039ba30b136329031b7b4b760791b60448201526064016101f2565b8760800151600181518110610556576105566135db565b60200260200101516001600160a01b0316886080015160008151811061057e5761057e6135db565b60200260200101516001600160a01b0316146105cf5760405162461bcd60e51b815260206004820152601060248201526f1896903737ba1039b0b6b29031b7b4b760811b60448201526064016101f2565b61063988608001516000815181106105e9576105e96135db565b60200260200101518960800151600181518110610608576106086135db565b60200260200101518a60800151600481518110610627576106276135db565b60200260200101518b60a001516120bc565b50809450506106858860800151600181518110610658576106586135db565b60200260200101518960800151600281518110610677576106776135db565b602002602001015186612492565b92508483116106c15760405162461bcd60e51b81526020600482015260086024820152676e6f20657874726160c01b60448201526064016101f2565b6106cb8584613660565b9150611571565b8751600214156109465761072088608001516000815181106106f6576106f66135db565b60200260200101516001600160a01b031673a07c5b74c9b40447a954e1466938b865b6bbea361490565b156108c15760a0880151604051632e1a7d4d60e01b8152600481019190915273bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c90632e1a7d4d90602401600060405180830381600087803b15801561077857600080fd5b505af115801561078c573d6000803e3d6000fd5b505050506107aa88608001516000815181106105e9576105e96135db565b50809450506107c98860800151600181518110610658576106586135db565b92508483116108095760405162461bcd60e51b815260206004820152600c60248201526b322e312d6e6f2d657874726160a01b60448201526064016101f2565b73bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c6001600160a01b031663d0e30db0846040518263ffffffff1660e01b81526004016000604051808303818588803b15801561085857600080fd5b505af115801561086c573d6000803e3d6000fd5b50505050506000858461087f9190613660565b9050610890818a6060015130612745565b915081600183516108a19190613660565b815181106108b1576108b16135db565b6020026020010151925050611571565b6108db88608001516000815181106105e9576105e96135db565b50809450506108fa8860800151600181518110610658576106586135db565b925084831161093a5760405162461bcd60e51b815260206004820152600c60248201526b322e322d6e6f2d657874726160a01b60448201526064016101f2565b600061087f8685613660565b875160031415610bbc5761096a88608001516001815181106104f2576104f26135db565b6109c05760405162461bcd60e51b815260206004820152602160248201527f332d7365697a656420746f6b656e206973206e6f7420737461626c6520636f696044820152603760f91b60648201526084016101f2565b6109da88608001516000815181106106f6576106f66135db565b15610b145760a0880151604051632e1a7d4d60e01b8152600481019190915273bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c90632e1a7d4d90602401600060405180830381600087803b158015610a3257600080fd5b505af1158015610a46573d6000803e3d6000fd5b50505050610a6488608001516000815181106105e9576105e96135db565b5080945050610a838860800151600181518110610658576106586135db565b9250610a948589604001513061287f565b905080600081518110610aa957610aa96135db565b60200260200101518311610aee5760405162461bcd60e51b815260206004820152600c60248201526b332e312d6e6f2d657874726160a01b60448201526064016101f2565b80600081518110610b0157610b016135db565b6020026020010151836106cb9190613660565b610b2e88608001516000815181106105e9576105e96135db565b5080945050610b4d8860800151600181518110610658576106586135db565b9250610b5e8589604001513061287f565b905080600081518110610b7357610b736135db565b60200260200101518311610aee5760405162461bcd60e51b815260206004820152601060248201526f332e322d626e622d6e6f2d657874726160801b60448201526064016101f2565b875160041415610e2f57610be088608001516000815181106104f2576104f26135db565b610c2c5760405162461bcd60e51b815260206004820152601f60248201527f342d7265706179546f6b656e206973206e6f7420737461626c6520636f696e0060448201526064016101f2565b610c4688608001516001815181106106f6576106f66135db565b15610d7a57610c6588608001516000815181106105e9576105e96135db565b5080945050610c848860800151600181518110610658576106586135db565b925073bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c6001600160a01b031663d0e30db0846040518263ffffffff1660e01b81526004016000604051808303818588803b158015610cd557600080fd5b505af1158015610ce9573d6000803e3d6000fd5b5050505050610cfd83896040015130612745565b905060008160018351610d109190613660565b81518110610d2057610d206135db565b60200260200101519050858111610d685760405162461bcd60e51b815260206004820152600c60248201526b342e312d6e6f20657874726160a01b60448201526064016101f2565b610d728682613660565b925050611571565b610d9488608001516000815181106105e9576105e96135db565b5080945050610db38860800151600181518110610658576106586135db565b9250610dc483896040015130612745565b905060008160018351610dd79190613660565b81518110610de757610de76135db565b60200260200101519050858111610d685760405162461bcd60e51b815260206004820152600c60248201526b342e322d6e6f20657874726160a01b60448201526064016101f2565b87516005141561110457610e5388608001516000815181106106f6576106f66135db565b15610f8f5760a0880151604051632e1a7d4d60e01b8152600481019190915273bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c90632e1a7d4d90602401600060405180830381600087803b158015610eab57600080fd5b505af1158015610ebf573d6000803e3d6000fd5b50505050610edd88608001516000815181106105e9576105e96135db565b5080945050610efc8860800151600181518110610658576106586135db565b9250610f0d8589604001513061287f565b905080600081518110610f2257610f226135db565b60200260200101518311610f675760405162461bcd60e51b815260206004820152600c60248201526b352e312d6e6f20657874726160a01b60448201526064016101f2565b600081600081518110610f7c57610f7c6135db565b60200260200101518461087f9190613660565b610fa988608001516001815181106106f6576106f66135db565b1561106057610fc888608001516000815181106105e9576105e96135db565b5080945050610fe78860800151600181518110610658576106586135db565b925073bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c6001600160a01b031663d0e30db0846040518263ffffffff1660e01b81526004016000604051808303818588803b15801561103857600080fd5b505af115801561104c573d6000803e3d6000fd5b5050505050610f0d8589604001513061287f565b61107a88608001516000815181106105e9576105e96135db565b50809450506110998860800151600181518110610658576106586135db565b92506110aa8589604001513061287f565b9050806000815181106110bf576110bf6135db565b60200260200101518311610f675760405162461bcd60e51b815260206004820152600c60248201526b352e332d6e6f20657874726160a01b60448201526064016101f2565b8751600614156112e45761112888608001516001815181106104f2576104f26135db565b6111745760405162461bcd60e51b815260206004820181905260248201527f362d7365697a6564546f6b656e206973206e6f7420737461626c6520636f696e60448201526064016101f2565b60006111e08960800151600081518110611190576111906135db565b60200260200101518a608001516001815181106111af576111af6135db565b60200260200101518b608001516004815181106111ce576111ce6135db565b60200260200101518c60a001516120bc565b809250819650505061122f8960800151600181518110611202576112026135db565b60200260200101518a60800151600281518110611221576112216135db565b602002602001015187612492565b935060008960a0015182886112449190613648565b61124e9190613660565b905061125f818b604001513061287f565b5082600081518110611273576112736135db565b602002602001015185116112b55760405162461bcd60e51b8152602060048201526009602482015268362d6e6f657874726160b81b60448201526064016101f2565b826000815181106112c8576112c86135db565b6020026020010151856112db9190613660565b93505050611571565b8751600714156100915761130888608001516001815181106106f6576106f66135db565b156114a95760006113298960800151600081518110611190576111906135db565b809250819650505061134b8960800151600181518110611202576112026135db565b935073bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c6001600160a01b031663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b15801561139c57600080fd5b505af11580156113b0573d6000803e3d6000fd5b505050505060008960a0015182886113c89190613648565b6113d29190613660565b90506113e3818b604001513061287f565b9250826000815181106113f8576113f86135db565b6020026020010151851161143d5760405162461bcd60e51b815260206004820152600c60248201526b372e312d6e6f20657874726160a01b60448201526064016101f2565b600083600081518110611452576114526135db565b6020026020010151866114659190613660565b9050611476818c6060015130612745565b935083600185516114879190613660565b81518110611497576114976135db565b60200260200101519450505050611571565b60006114c58960800151600081518110611190576111906135db565b80925081965050506114e78960800151600181518110611202576112026135db565b935060008960a0015182886114fc9190613648565b6115069190613660565b9050611517818b604001513061287f565b92508260008151811061152c5761152c6135db565b6020026020010151851161143d5760405162461bcd60e51b815260206004820152600c60248201526b372e322d6e6f20657874726160a01b60448201526064016101f2565b602088015160405163a9059cbb60e01b81526001600160a01b039182166004820152602481018790529088169063a9059cbb906044016020604051808303816000875af11580156115c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ea9190613677565b50875160a0808a0151604080519384526001600160a01b03808c166020860152908401919091528816606083015260808201879052810184905260c081018390527fee80f069d389b1636bfe8a115ff1d8ad940b394b4d16bb19c8503003d14dd5b49060e00160405180910390a150505050505050505050505050565b6000546001600160a01b031633146116915760405162461bcd60e51b81526004016101f29061344c565b6000670de0b6b3a7640000905073bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c6001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b1580156116ed57600080fd5b505af1158015611701573d6000803e3d6000fd5b506000935060029250611712915050565b60405190808252806020026020018201604052801561173b578160200160208202803683370190505b50905073bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c81600081518110611766576117666135db565b60200260200101906001600160a01b031690816001600160a01b031681525050828160018151811061179a5761179a6135db565b60200260200101906001600160a01b031690816001600160a01b0316815250506117c5828230612745565b50505050565b600073fd36e2c2a6789db23113685031d7f16329158384635ec88c79858560048181106117fa576117fa6135db565b905060200201602081019061180f91906132eb565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401606060405180830381865afa158015611853573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118779190613699565b92505050600081116118d55760405162461bcd60e51b815260206004820152602160248201527f73686f727466616c6c206d7573742067726561746572207468616e207a6572206044820152600360fc1b60648201526084016101f2565b5060016000848460038181106118ed576118ed6135db565b905060200201602081019061190291906132eb565b6001600160a01b03166001600160a01b03168152602001908152602001600020600084846000818110611937576119376135db565b905060200201602081019061194c91906132eb565b6001600160a01b0316815260208101919091526040016000205460ff16611ade5782826003818110611980576119806135db565b905060200201602081019061199591906132eb565b6001600160a01b031663095ea7b3848460008181106119b6576119b66135db565b90506020020160208101906119cb91906132eb565b6040516001600160e01b031960e084901b1681526001600160a01b03909116600482015260001960248201526044016020604051808303816000875af1158015611a19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a3d9190613677565b50600180600085856003818110611a5657611a566135db565b9050602002016020810190611a6b91906132eb565b6001600160a01b03166001600160a01b03168152602001908152602001600020600085856000818110611aa057611aa06135db565b9050602002016020810190611ab591906132eb565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790555b6000886001600160a01b0316630dfe16816040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b4291906136c7565b90506000896001600160a01b031663d21220a76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ba891906136c7565b90506000826001600160a01b031686866003818110611bc957611bc96135db565b9050602002016020810190611bde91906132eb565b6001600160a01b031614611bf3576000611bf5565b835b90506000826001600160a01b031687876003818110611c1657611c166135db565b9050602002016020810190611c2b91906132eb565b6001600160a01b031614611c40576000611c42565b845b905060008d8d8d8d8d8d8d8d8d604051602001611c679998979695949392919061372d565b60408051601f198184030181529082905263022c0d9f60e01b825291506001600160a01b038e169063022c0d9f90611ca9908690869030908790600401613797565b600060405180830381600087803b158015611cc357600080fd5b505af1158015611cd7573d6000803e3d6000fd5b505050505050505050505050505050505050565b6000546001600160a01b03163314611d155760405162461bcd60e51b81526004016101f29061344c565b6001600160a01b038316611d6b5760405162461bcd60e51b815260206004820152601e60248201527f746f6b656e206d757374206e6f74206265207a65726f2061646472657373000060448201526064016101f2565b6001600160a01b038216611dc15760405162461bcd60e51b815260206004820152601b60248201527f746f206d757374206e6f74206265207a65726f2061646472657373000000000060448201526064016101f2565b60008111611e115760405162461bcd60e51b815260206004820152601c60248201527f616d6f756e74206d75737420626967676572207468616e207a65726f0000000060448201526064016101f2565b60405163a9059cbb60e01b81526001600160a01b0383811660048301526024820183905284169063a9059cbb906044016020604051808303816000875af1158015611e60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e849190613677565b50816001600160a01b0316836001600160a01b03167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb83604051611eca91815260200190565b60405180910390a3505050565b6040516315e3f14f60e11b81526001600160a01b038216600482015260009073fd36e2c2a6789db23113685031d7f1632915838490632bc7e29e90602401602060405180830381865afa158015611f32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f56919061380a565b92915050565b6000546001600160a01b03163314611f865760405162461bcd60e51b81526004016101f29061344c565b6001600160a01b038116611feb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101f2565b611ff481611ff7565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006001600160a01b0382167395c78222b3d6e262426483d42cfa53685a67ab9d148061209057506001600160a01b03821673fd5840cd36d94d7229439859c0112a4185bc0255145b80611f5657506001600160a01b03821673334b3ecb4dca3593bccc3c7ebd1a1c1d1780fbf11492915050565b6040516370a0823160e01b8152306004820152600090819085908290819081906001600160a01b038516906370a0823190602401602060405180830381865afa15801561210d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612131919061380a565b905073a07c5b74c9b40447a954e1466938b865b6bbea366001600160a01b038b1614156121c257604051635572051560e11b81526001600160a01b0389811660048301528a811660248301528b169063aae40a2a9089906044016000604051808303818588803b1580156121a457600080fd5b505af11580156121b8573d6000803e3d6000fd5b5050505050612375565b724065d34c6b18ce4370ced1cebde94865dbfafe6001600160a01b038b1614156122b0576040516311b3d5e760e01b81526001600160a01b0389811660048301526024820189905285811660448301528b16906311b3d5e79060640160408051808303816000875af115801561223c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122609190613823565b909350915082156122ab5760405162461bcd60e51b81526020600482015260156024820152743634b8bab4b230ba32a137b93937bb9032b93937b960591b60448201526064016101f2565b612375565b604051637af1e23160e11b81526001600160a01b0389811660048301526024820189905285811660448301528b919082169063f5e3c462906064016020604051808303816000875af115801561230a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061232e919061380a565b156123735760405162461bcd60e51b81526020600482015260156024820152743634b8bab4b230ba32a137b93937bb9032b93937b960591b60448201526064016101f2565b505b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa1580156123bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123e0919061380a565b905060006123ee8383613660565b9050600081116124405760405162461bcd60e51b815260206004820152601c60248201527f7365697a65642056546f6b656e20616d6f756e74206973207a65726f0000000060448201526064016101f2565b604080518a8152602081018690529081018290527ffde0eb536fecee7fca644722fc2e6255ee0571307ab7f46f7fea23fe302d44799060600160405180910390a19b929a509198505050505050505050565b600080808073a07c5b74c9b40447a954e1466938b865b6bbea366001600160a01b03881614156125705760405163db006a7560e01b8152600481018690524793506001600160a01b0388169063db006a75906024016020604051808303816000875af1158015612506573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061252a919061380a565b156125685760405162461bcd60e51b815260206004820152600e60248201526d3932b232b2b6902127211032b93960911b60448201526064016101f2565b4791506126f3565b6040516370a0823160e01b8152306004820152879087906001600160a01b038216906370a0823190602401602060405180830381865afa1580156125b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125dc919061380a565b60405163db006a7560e01b8152600481018990529095506001600160a01b0383169063db006a75906024016020604051808303816000875af1158015612626573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061264a919061380a565b156126865760405162461bcd60e51b815260206004820152600c60248201526b3932b232b2b69032b93937b960a11b60448201526064016101f2565b6040516370a0823160e01b81523060048201526001600160a01b038216906370a0823190602401602060405180830381865afa1580156126ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126ee919061380a565b935050505b6126fd8383613660565b60408051878152602081018390529192507f2a500c424b9e55ca96ef57f09281232662a42069de49730505b2bbd019412dcf910160405180910390a1925050505b9392505050565b606061276673ca143ce32fe78f1f7019d7d551a6402fc5350c7385856128a0565b90508260008151811061277b5761277b6135db565b60200260200101516001600160a01b031663a9059cbb6127e473ca143ce32fe78f1f7019d7d551a6402fc5350c73866000815181106127bc576127bc6135db565b6020026020010151876001815181106127d7576127d76135db565b6020026020010151612a2b565b836000815181106127f7576127f76135db565b60200260200101516040518363ffffffff1660e01b81526004016128309291906001600160a01b03929092168252602082015260400190565b6020604051808303816000875af115801561284f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128739190613677565b5061273e818484612b03565b606061276673ca143ce32fe78f1f7019d7d551a6402fc5350c738585612ce0565b60606002825110156128f45760405162461bcd60e51b815260206004820152601c60248201527f50616e63616b654c6962726172793a20494e56414c49445f504154480000000060448201526064016101f2565b815167ffffffffffffffff81111561290e5761290e613481565b604051908082528060200260200182016040528015612937578160200160208202803683370190505b509050828160008151811061294e5761294e6135db565b60200260200101818152505060005b6001835161296b9190613660565b811015612a23576000806129be8786858151811061298b5761298b6135db565b6020026020010151878660016129a19190613648565b815181106129b1576129b16135db565b6020026020010151612e61565b915091506129e68484815181106129d7576129d76135db565b60200260200101518383612f37565b846129f2856001613648565b81518110612a0257612a026135db565b60200260200101818152505050508080612a1b90613847565b91505061295d565b509392505050565b6000806000612a3a8585613013565b6040516bffffffffffffffffffffffff19606084811b8216602084015283901b1660348201529193509150869060480160405160208183030381529060405280519060200120604051602001612ae19291906001600160f81b0319815260609290921b6bffffffffffffffffffffffff1916600183015260158201527efb7f630766e6a796048ea87d01acd3068e8ff67d078148a3fa3f4a84f69bd5603582015260550190565b60408051601f1981840301815291905280516020909101209695505050505050565b60005b60018351612b149190613660565b8110156117c557600080848381518110612b3057612b306135db565b602002602001015185846001612b469190613648565b81518110612b5657612b566135db565b6020026020010151915091506000612b6e8383613013565b509050600087612b7f866001613648565b81518110612b8f57612b8f6135db565b60200260200101519050600080836001600160a01b0316866001600160a01b031614612bbd57826000612bc1565b6000835b91509150600060028a51612bd59190613660565b8810612be15788612c16565b612c1673ca143ce32fe78f1f7019d7d551a6402fc5350c73878c612c068c6002613648565b815181106127d7576127d76135db565b9050612c3773ca143ce32fe78f1f7019d7d551a6402fc5350c738888612a2b565b6001600160a01b031663022c0d9f84848460006040519080825280601f01601f191660200182016040528015612c74576020820181803683370190505b506040518563ffffffff1660e01b8152600401612c949493929190613797565b600060405180830381600087803b158015612cae57600080fd5b505af1158015612cc2573d6000803e3d6000fd5b50505050505050505050508080612cd890613847565b915050612b06565b6060600282511015612d345760405162461bcd60e51b815260206004820152601c60248201527f50616e63616b654c6962726172793a20494e56414c49445f504154480000000060448201526064016101f2565b815167ffffffffffffffff811115612d4e57612d4e613481565b604051908082528060200260200182016040528015612d77578160200160208202803683370190505b509050828160018351612d8a9190613660565b81518110612d9a57612d9a6135db565b602002602001018181525050600060018351612db69190613660565b90505b8015612a2357600080612dfc8786612dd2600187613660565b81518110612de257612de26135db565b60200260200101518786815181106129b1576129b16135db565b91509150612e24848481518110612e1557612e156135db565b60200260200101518383613109565b84612e30600186613660565b81518110612e4057612e406135db565b60200260200101818152505050508080612e5990613862565b915050612db9565b6000806000612e708585613013565b509050612e7e868686612a2b565b50600080612e8d888888612a2b565b6001600160a01b0316630902f1ac6040518163ffffffff1660e01b8152600401606060405180830381865afa158015612eca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612eee9190613890565b506001600160701b031691506001600160701b03169150826001600160a01b0316876001600160a01b031614612f25578082612f28565b81815b90999098509650505050505050565b6000808411612f9a5760405162461bcd60e51b815260206004820152602960248201527f50616e63616b654c6962726172793a20494e53554646494349454e545f494e50604482015268155517d05353d5539560ba1b60648201526084016101f2565b600083118015612faa5750600082115b612fc65760405162461bcd60e51b81526004016101f2906138e0565b6000612fd4856126f76131e0565b90506000612fe282856131e0565b90506000612ffc83612ff6886127106131e0565b906131ec565b90506130088183613626565b979650505050505050565b600080826001600160a01b0316846001600160a01b031614156130845760405162461bcd60e51b815260206004820152602360248201527f50616e63616b654c6962726172793a204944454e544943414c5f41444452455360448201526253455360e81b60648201526084016101f2565b826001600160a01b0316846001600160a01b0316106130a45782846130a7565b83835b90925090506001600160a01b0382166131025760405162461bcd60e51b815260206004820152601c60248201527f50616e63616b654c6962726172793a205a45524f5f414444524553530000000060448201526064016101f2565b9250929050565b600080841161316d5760405162461bcd60e51b815260206004820152602a60248201527f50616e63616b654c6962726172793a20494e53554646494349454e545f4f555460448201526914155517d05353d5539560b21b60648201526084016101f2565b60008311801561317d5750600082115b6131995760405162461bcd60e51b81526004016101f2906138e0565b60006131b16127106131ab86886131e0565b906131e0565b905060006131c56126f76131ab86896131f8565b90506131d66001612ff68385613626565b9695505050505050565b600061273e8284613607565b600061273e8284613648565b600061273e8284613660565b6001600160a01b0381168114611ff457600080fd5b6000806040838503121561322c57600080fd5b823561323781613204565b946020939093013593505050565b803561325081613204565b919050565b60008060008060006080868803121561326d57600080fd5b853561327881613204565b94506020860135935060408601359250606086013567ffffffffffffffff808211156132a357600080fd5b818801915088601f8301126132b757600080fd5b8135818111156132c657600080fd5b8960208285010111156132d857600080fd5b9699959850939650602001949392505050565b6000602082840312156132fd57600080fd5b813561273e81613204565b60008083601f84011261331a57600080fd5b50813567ffffffffffffffff81111561333257600080fd5b6020830191508360208260051b850101111561310257600080fd5b600080600080600080600080600060c08a8c03121561336b57600080fd5b8935985060208a013561337d81613204565b975060408a013567ffffffffffffffff8082111561339a57600080fd5b6133a68d838e01613308565b909950975060608c01359150808211156133bf57600080fd5b6133cb8d838e01613308565b909750955060808c01359150808211156133e457600080fd5b506133f18c828d01613308565b9a9d999c50979a9699959894979660a00135949350505050565b60008060006060848603121561342057600080fd5b833561342b81613204565b9250602084013561343b81613204565b929592945050506040919091013590565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126134a857600080fd5b8135602067ffffffffffffffff808311156134c5576134c5613481565b8260051b604051601f19603f830116810181811084821117156134ea576134ea613481565b60405293845285810183019383810192508785111561350857600080fd5b83870191505b848210156130085761351f82613245565b8352918301919083019061350e565b60008060008060008060c0878903121561354757600080fd5b86359550602087013561355981613204565b9450604087013567ffffffffffffffff8082111561357657600080fd5b6135828a838b01613497565b9550606089013591508082111561359857600080fd5b6135a48a838b01613497565b945060808901359150808211156135ba57600080fd5b506135c789828a01613497565b92505060a087013590509295509295509295565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615613621576136216135f1565b500290565b60008261364357634e487b7160e01b600052601260045260246000fd5b500490565b6000821982111561365b5761365b6135f1565b500190565b600082821015613672576136726135f1565b500390565b60006020828403121561368957600080fd5b8151801515811461273e57600080fd5b6000806000606084860312156136ae57600080fd5b8351925060208401519150604084015190509250925092565b6000602082840312156136d957600080fd5b815161273e81613204565b8183526000602080850194508260005b8581101561372257813561370781613204565b6001600160a01b0316875295820195908201906001016136f4565b509495945050505050565b8981526001600160a01b038916602082015260c060408201819052600090613758908301898b6136e4565b828103606084015261376b81888a6136e4565b905082810360808401526137808186886136e4565b9150508260a08301529a9950505050505050505050565b84815260006020858184015260018060a01b038516604084015260806060840152835180608085015260005b818110156137df5785810183015185820160a0015282016137c3565b818111156137f157600060a083870101525b50601f01601f19169290920160a0019695505050505050565b60006020828403121561381c57600080fd5b5051919050565b6000806040838503121561383657600080fd5b505080516020909101519092909150565b600060001982141561385b5761385b6135f1565b5060010190565b600081613871576138716135f1565b506000190190565b80516001600160701b038116811461325057600080fd5b6000806000606084860312156138a557600080fd5b6138ae84613879565b92506138bc60208501613879565b9150604084015163ffffffff811681146138d557600080fd5b809150509250925092565b60208082526026908201527f50616e63616b654c6962726172793a20494e53554646494349454e545f4c495160408201526555494449545960d01b60608201526080019056fea26469706673582212202f6c4ddd56a3049383edc078f4df88d69c1ed03eaa5db8c2522bd79f1a045a1564736f6c634300080b0033",
}

// UniFlashSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use UniFlashSwapMetaData.ABI instead.
var UniFlashSwapABI = UniFlashSwapMetaData.ABI

// Deprecated: Use UniFlashSwapMetaData.Sigs instead.
// UniFlashSwapFuncSigs maps the 4-byte function signature to its string representation.
var UniFlashSwapFuncSigs = UniFlashSwapMetaData.Sigs

// UniFlashSwapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniFlashSwapMetaData.Bin instead.
var UniFlashSwapBin = UniFlashSwapMetaData.Bin

// DeployUniFlashSwap deploys a new Ethereum contract, binding an instance of UniFlashSwap to it.
func DeployUniFlashSwap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniFlashSwap, error) {
	parsed, err := UniFlashSwapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniFlashSwapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniFlashSwap{UniFlashSwapCaller: UniFlashSwapCaller{contract: contract}, UniFlashSwapTransactor: UniFlashSwapTransactor{contract: contract}, UniFlashSwapFilterer: UniFlashSwapFilterer{contract: contract}}, nil
}

// UniFlashSwap is an auto generated Go binding around an Ethereum contract.
type UniFlashSwap struct {
	UniFlashSwapCaller     // Read-only binding to the contract
	UniFlashSwapTransactor // Write-only binding to the contract
	UniFlashSwapFilterer   // Log filterer for contract events
}

// UniFlashSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniFlashSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniFlashSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniFlashSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniFlashSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniFlashSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniFlashSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniFlashSwapSession struct {
	Contract     *UniFlashSwap     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniFlashSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniFlashSwapCallerSession struct {
	Contract *UniFlashSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UniFlashSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniFlashSwapTransactorSession struct {
	Contract     *UniFlashSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UniFlashSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniFlashSwapRaw struct {
	Contract *UniFlashSwap // Generic contract binding to access the raw methods on
}

// UniFlashSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniFlashSwapCallerRaw struct {
	Contract *UniFlashSwapCaller // Generic read-only contract binding to access the raw methods on
}

// UniFlashSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniFlashSwapTransactorRaw struct {
	Contract *UniFlashSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniFlashSwap creates a new instance of UniFlashSwap, bound to a specific deployed contract.
func NewUniFlashSwap(address common.Address, backend bind.ContractBackend) (*UniFlashSwap, error) {
	contract, err := bindUniFlashSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniFlashSwap{UniFlashSwapCaller: UniFlashSwapCaller{contract: contract}, UniFlashSwapTransactor: UniFlashSwapTransactor{contract: contract}, UniFlashSwapFilterer: UniFlashSwapFilterer{contract: contract}}, nil
}

// NewUniFlashSwapCaller creates a new read-only instance of UniFlashSwap, bound to a specific deployed contract.
func NewUniFlashSwapCaller(address common.Address, caller bind.ContractCaller) (*UniFlashSwapCaller, error) {
	contract, err := bindUniFlashSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniFlashSwapCaller{contract: contract}, nil
}

// NewUniFlashSwapTransactor creates a new write-only instance of UniFlashSwap, bound to a specific deployed contract.
func NewUniFlashSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*UniFlashSwapTransactor, error) {
	contract, err := bindUniFlashSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniFlashSwapTransactor{contract: contract}, nil
}

// NewUniFlashSwapFilterer creates a new log filterer instance of UniFlashSwap, bound to a specific deployed contract.
func NewUniFlashSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*UniFlashSwapFilterer, error) {
	contract, err := bindUniFlashSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniFlashSwapFilterer{contract: contract}, nil
}

// bindUniFlashSwap binds a generic wrapper to an already deployed contract.
func bindUniFlashSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniFlashSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniFlashSwap *UniFlashSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniFlashSwap.Contract.UniFlashSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniFlashSwap *UniFlashSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.UniFlashSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniFlashSwap *UniFlashSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.UniFlashSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniFlashSwap *UniFlashSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniFlashSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniFlashSwap *UniFlashSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniFlashSwap *UniFlashSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.contract.Transact(opts, method, params...)
}

// MintedVAIS is a free data retrieval call binding the contract method 0xf064b414.
//
// Solidity: function mintedVAIS(address _account) view returns(uint256)
func (_UniFlashSwap *UniFlashSwapCaller) MintedVAIS(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniFlashSwap.contract.Call(opts, &out, "mintedVAIS", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedVAIS is a free data retrieval call binding the contract method 0xf064b414.
//
// Solidity: function mintedVAIS(address _account) view returns(uint256)
func (_UniFlashSwap *UniFlashSwapSession) MintedVAIS(_account common.Address) (*big.Int, error) {
	return _UniFlashSwap.Contract.MintedVAIS(&_UniFlashSwap.CallOpts, _account)
}

// MintedVAIS is a free data retrieval call binding the contract method 0xf064b414.
//
// Solidity: function mintedVAIS(address _account) view returns(uint256)
func (_UniFlashSwap *UniFlashSwapCallerSession) MintedVAIS(_account common.Address) (*big.Int, error) {
	return _UniFlashSwap.Contract.MintedVAIS(&_UniFlashSwap.CallOpts, _account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniFlashSwap *UniFlashSwapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniFlashSwap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniFlashSwap *UniFlashSwapSession) Owner() (common.Address, error) {
	return _UniFlashSwap.Contract.Owner(&_UniFlashSwap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniFlashSwap *UniFlashSwapCallerSession) Owner() (common.Address, error) {
	return _UniFlashSwap.Contract.Owner(&_UniFlashSwap.CallOpts)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_UniFlashSwap *UniFlashSwapTransactor) PancakeCall(opts *bind.TransactOpts, _sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _UniFlashSwap.contract.Transact(opts, "pancakeCall", _sender, _amount0, _amount1, _data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_UniFlashSwap *UniFlashSwapSession) PancakeCall(_sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.PancakeCall(&_UniFlashSwap.TransactOpts, _sender, _amount0, _amount1, _data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_UniFlashSwap *UniFlashSwapTransactorSession) PancakeCall(_sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.PancakeCall(&_UniFlashSwap.TransactOpts, _sender, _amount0, _amount1, _data)
}

// Qingsuan is a paid mutator transaction binding the contract method 0xb14e9a72.
//
// Solidity: function qingsuan(uint256 _situation, address _flashLoanFrom, address[] _path1, address[] _path2, address[] _tokens, uint256 _flashLoanAmount) returns()
func (_UniFlashSwap *UniFlashSwapTransactor) Qingsuan(opts *bind.TransactOpts, _situation *big.Int, _flashLoanFrom common.Address, _path1 []common.Address, _path2 []common.Address, _tokens []common.Address, _flashLoanAmount *big.Int) (*types.Transaction, error) {
	return _UniFlashSwap.contract.Transact(opts, "qingsuan", _situation, _flashLoanFrom, _path1, _path2, _tokens, _flashLoanAmount)
}

// Qingsuan is a paid mutator transaction binding the contract method 0xb14e9a72.
//
// Solidity: function qingsuan(uint256 _situation, address _flashLoanFrom, address[] _path1, address[] _path2, address[] _tokens, uint256 _flashLoanAmount) returns()
func (_UniFlashSwap *UniFlashSwapSession) Qingsuan(_situation *big.Int, _flashLoanFrom common.Address, _path1 []common.Address, _path2 []common.Address, _tokens []common.Address, _flashLoanAmount *big.Int) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.Qingsuan(&_UniFlashSwap.TransactOpts, _situation, _flashLoanFrom, _path1, _path2, _tokens, _flashLoanAmount)
}

// Qingsuan is a paid mutator transaction binding the contract method 0xb14e9a72.
//
// Solidity: function qingsuan(uint256 _situation, address _flashLoanFrom, address[] _path1, address[] _path2, address[] _tokens, uint256 _flashLoanAmount) returns()
func (_UniFlashSwap *UniFlashSwapTransactorSession) Qingsuan(_situation *big.Int, _flashLoanFrom common.Address, _path1 []common.Address, _path2 []common.Address, _tokens []common.Address, _flashLoanAmount *big.Int) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.Qingsuan(&_UniFlashSwap.TransactOpts, _situation, _flashLoanFrom, _path1, _path2, _tokens, _flashLoanAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniFlashSwap *UniFlashSwapTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniFlashSwap.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniFlashSwap *UniFlashSwapSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniFlashSwap.Contract.RenounceOwnership(&_UniFlashSwap.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniFlashSwap *UniFlashSwapTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniFlashSwap.Contract.RenounceOwnership(&_UniFlashSwap.TransactOpts)
}

// SwapOneBNBToFlashLoandUnderlyingToken is a paid mutator transaction binding the contract method 0x89c33748.
//
// Solidity: function swapOneBNBToFlashLoandUnderlyingToken(address _flashLoanUnderlyingToken) returns()
func (_UniFlashSwap *UniFlashSwapTransactor) SwapOneBNBToFlashLoandUnderlyingToken(opts *bind.TransactOpts, _flashLoanUnderlyingToken common.Address) (*types.Transaction, error) {
	return _UniFlashSwap.contract.Transact(opts, "swapOneBNBToFlashLoandUnderlyingToken", _flashLoanUnderlyingToken)
}

// SwapOneBNBToFlashLoandUnderlyingToken is a paid mutator transaction binding the contract method 0x89c33748.
//
// Solidity: function swapOneBNBToFlashLoandUnderlyingToken(address _flashLoanUnderlyingToken) returns()
func (_UniFlashSwap *UniFlashSwapSession) SwapOneBNBToFlashLoandUnderlyingToken(_flashLoanUnderlyingToken common.Address) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.SwapOneBNBToFlashLoandUnderlyingToken(&_UniFlashSwap.TransactOpts, _flashLoanUnderlyingToken)
}

// SwapOneBNBToFlashLoandUnderlyingToken is a paid mutator transaction binding the contract method 0x89c33748.
//
// Solidity: function swapOneBNBToFlashLoandUnderlyingToken(address _flashLoanUnderlyingToken) returns()
func (_UniFlashSwap *UniFlashSwapTransactorSession) SwapOneBNBToFlashLoandUnderlyingToken(_flashLoanUnderlyingToken common.Address) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.SwapOneBNBToFlashLoandUnderlyingToken(&_UniFlashSwap.TransactOpts, _flashLoanUnderlyingToken)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniFlashSwap *UniFlashSwapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UniFlashSwap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniFlashSwap *UniFlashSwapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.TransferOwnership(&_UniFlashSwap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniFlashSwap *UniFlashSwapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.TransferOwnership(&_UniFlashSwap.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _to, uint256 _amount) returns()
func (_UniFlashSwap *UniFlashSwapTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UniFlashSwap.contract.Transact(opts, "withdraw", _token, _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _to, uint256 _amount) returns()
func (_UniFlashSwap *UniFlashSwapSession) Withdraw(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.Withdraw(&_UniFlashSwap.TransactOpts, _token, _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _to, uint256 _amount) returns()
func (_UniFlashSwap *UniFlashSwapTransactorSession) Withdraw(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.Withdraw(&_UniFlashSwap.TransactOpts, _token, _to, _amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address _to, uint256 _amount) returns()
func (_UniFlashSwap *UniFlashSwapTransactor) WithdrawETH(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UniFlashSwap.contract.Transact(opts, "withdrawETH", _to, _amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address _to, uint256 _amount) returns()
func (_UniFlashSwap *UniFlashSwapSession) WithdrawETH(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.WithdrawETH(&_UniFlashSwap.TransactOpts, _to, _amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address _to, uint256 _amount) returns()
func (_UniFlashSwap *UniFlashSwapTransactorSession) WithdrawETH(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UniFlashSwap.Contract.WithdrawETH(&_UniFlashSwap.TransactOpts, _to, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniFlashSwap *UniFlashSwapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniFlashSwap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniFlashSwap *UniFlashSwapSession) Receive() (*types.Transaction, error) {
	return _UniFlashSwap.Contract.Receive(&_UniFlashSwap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniFlashSwap *UniFlashSwapTransactorSession) Receive() (*types.Transaction, error) {
	return _UniFlashSwap.Contract.Receive(&_UniFlashSwap.TransactOpts)
}

// UniFlashSwapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UniFlashSwap contract.
type UniFlashSwapOwnershipTransferredIterator struct {
	Event *UniFlashSwapOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UniFlashSwapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniFlashSwapOwnershipTransferred)
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
		it.Event = new(UniFlashSwapOwnershipTransferred)
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
func (it *UniFlashSwapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniFlashSwapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniFlashSwapOwnershipTransferred represents a OwnershipTransferred event raised by the UniFlashSwap contract.
type UniFlashSwapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniFlashSwap *UniFlashSwapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniFlashSwapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniFlashSwap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniFlashSwapOwnershipTransferredIterator{contract: _UniFlashSwap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniFlashSwap *UniFlashSwapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniFlashSwapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniFlashSwap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniFlashSwapOwnershipTransferred)
				if err := _UniFlashSwap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniFlashSwap *UniFlashSwapFilterer) ParseOwnershipTransferred(log types.Log) (*UniFlashSwapOwnershipTransferred, error) {
	event := new(UniFlashSwapOwnershipTransferred)
	if err := _UniFlashSwap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniFlashSwapScenarioIterator is returned from FilterScenario and is used to iterate over the raw logs and unpacked data for Scenario events raised by the UniFlashSwap contract.
type UniFlashSwapScenarioIterator struct {
	Event *UniFlashSwapScenario // Event containing the contract specifics and raw log

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
func (it *UniFlashSwapScenarioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniFlashSwapScenario)
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
		it.Event = new(UniFlashSwapScenario)
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
func (it *UniFlashSwapScenarioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniFlashSwapScenarioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniFlashSwapScenario represents a Scenario event raised by the UniFlashSwap contract.
type UniFlashSwapScenario struct {
	ScenarioNo             *big.Int
	RepayUnderlyingToken   common.Address
	RepayAmount            *big.Int
	SeizedUnderlyingToken  common.Address
	FlashLoanReturnAmount  *big.Int
	SeizedUnderlyingAmount *big.Int
	MassProfit             *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterScenario is a free log retrieval operation binding the contract event 0xee80f069d389b1636bfe8a115ff1d8ad940b394b4d16bb19c8503003d14dd5b4.
//
// Solidity: event Scenario(uint256 scenarioNo, address repayUnderlyingToken, uint256 repayAmount, address seizedUnderlyingToken, uint256 flashLoanReturnAmount, uint256 seizedUnderlyingAmount, uint256 massProfit)
func (_UniFlashSwap *UniFlashSwapFilterer) FilterScenario(opts *bind.FilterOpts) (*UniFlashSwapScenarioIterator, error) {

	logs, sub, err := _UniFlashSwap.contract.FilterLogs(opts, "Scenario")
	if err != nil {
		return nil, err
	}
	return &UniFlashSwapScenarioIterator{contract: _UniFlashSwap.contract, event: "Scenario", logs: logs, sub: sub}, nil
}

// WatchScenario is a free log subscription operation binding the contract event 0xee80f069d389b1636bfe8a115ff1d8ad940b394b4d16bb19c8503003d14dd5b4.
//
// Solidity: event Scenario(uint256 scenarioNo, address repayUnderlyingToken, uint256 repayAmount, address seizedUnderlyingToken, uint256 flashLoanReturnAmount, uint256 seizedUnderlyingAmount, uint256 massProfit)
func (_UniFlashSwap *UniFlashSwapFilterer) WatchScenario(opts *bind.WatchOpts, sink chan<- *UniFlashSwapScenario) (event.Subscription, error) {

	logs, sub, err := _UniFlashSwap.contract.WatchLogs(opts, "Scenario")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniFlashSwapScenario)
				if err := _UniFlashSwap.contract.UnpackLog(event, "Scenario", log); err != nil {
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

// ParseScenario is a log parse operation binding the contract event 0xee80f069d389b1636bfe8a115ff1d8ad940b394b4d16bb19c8503003d14dd5b4.
//
// Solidity: event Scenario(uint256 scenarioNo, address repayUnderlyingToken, uint256 repayAmount, address seizedUnderlyingToken, uint256 flashLoanReturnAmount, uint256 seizedUnderlyingAmount, uint256 massProfit)
func (_UniFlashSwap *UniFlashSwapFilterer) ParseScenario(log types.Log) (*UniFlashSwapScenario, error) {
	event := new(UniFlashSwapScenario)
	if err := _UniFlashSwap.contract.UnpackLog(event, "Scenario", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniFlashSwapSeizedUnderlyingTokenAmountIterator is returned from FilterSeizedUnderlyingTokenAmount and is used to iterate over the raw logs and unpacked data for SeizedUnderlyingTokenAmount events raised by the UniFlashSwap contract.
type UniFlashSwapSeizedUnderlyingTokenAmountIterator struct {
	Event *UniFlashSwapSeizedUnderlyingTokenAmount // Event containing the contract specifics and raw log

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
func (it *UniFlashSwapSeizedUnderlyingTokenAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniFlashSwapSeizedUnderlyingTokenAmount)
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
		it.Event = new(UniFlashSwapSeizedUnderlyingTokenAmount)
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
func (it *UniFlashSwapSeizedUnderlyingTokenAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniFlashSwapSeizedUnderlyingTokenAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniFlashSwapSeizedUnderlyingTokenAmount represents a SeizedUnderlyingTokenAmount event raised by the UniFlashSwap contract.
type UniFlashSwapSeizedUnderlyingTokenAmount struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSeizedUnderlyingTokenAmount is a free log retrieval operation binding the contract event 0x2a500c424b9e55ca96ef57f09281232662a42069de49730505b2bbd019412dcf.
//
// Solidity: event SeizedUnderlyingTokenAmount(uint256 arg0, uint256 arg1)
func (_UniFlashSwap *UniFlashSwapFilterer) FilterSeizedUnderlyingTokenAmount(opts *bind.FilterOpts) (*UniFlashSwapSeizedUnderlyingTokenAmountIterator, error) {

	logs, sub, err := _UniFlashSwap.contract.FilterLogs(opts, "SeizedUnderlyingTokenAmount")
	if err != nil {
		return nil, err
	}
	return &UniFlashSwapSeizedUnderlyingTokenAmountIterator{contract: _UniFlashSwap.contract, event: "SeizedUnderlyingTokenAmount", logs: logs, sub: sub}, nil
}

// WatchSeizedUnderlyingTokenAmount is a free log subscription operation binding the contract event 0x2a500c424b9e55ca96ef57f09281232662a42069de49730505b2bbd019412dcf.
//
// Solidity: event SeizedUnderlyingTokenAmount(uint256 arg0, uint256 arg1)
func (_UniFlashSwap *UniFlashSwapFilterer) WatchSeizedUnderlyingTokenAmount(opts *bind.WatchOpts, sink chan<- *UniFlashSwapSeizedUnderlyingTokenAmount) (event.Subscription, error) {

	logs, sub, err := _UniFlashSwap.contract.WatchLogs(opts, "SeizedUnderlyingTokenAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniFlashSwapSeizedUnderlyingTokenAmount)
				if err := _UniFlashSwap.contract.UnpackLog(event, "SeizedUnderlyingTokenAmount", log); err != nil {
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

// ParseSeizedUnderlyingTokenAmount is a log parse operation binding the contract event 0x2a500c424b9e55ca96ef57f09281232662a42069de49730505b2bbd019412dcf.
//
// Solidity: event SeizedUnderlyingTokenAmount(uint256 arg0, uint256 arg1)
func (_UniFlashSwap *UniFlashSwapFilterer) ParseSeizedUnderlyingTokenAmount(log types.Log) (*UniFlashSwapSeizedUnderlyingTokenAmount, error) {
	event := new(UniFlashSwapSeizedUnderlyingTokenAmount)
	if err := _UniFlashSwap.contract.UnpackLog(event, "SeizedUnderlyingTokenAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniFlashSwapSeizedVTokenAmountIterator is returned from FilterSeizedVTokenAmount and is used to iterate over the raw logs and unpacked data for SeizedVTokenAmount events raised by the UniFlashSwap contract.
type UniFlashSwapSeizedVTokenAmountIterator struct {
	Event *UniFlashSwapSeizedVTokenAmount // Event containing the contract specifics and raw log

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
func (it *UniFlashSwapSeizedVTokenAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniFlashSwapSeizedVTokenAmount)
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
		it.Event = new(UniFlashSwapSeizedVTokenAmount)
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
func (it *UniFlashSwapSeizedVTokenAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniFlashSwapSeizedVTokenAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniFlashSwapSeizedVTokenAmount represents a SeizedVTokenAmount event raised by the UniFlashSwap contract.
type UniFlashSwapSeizedVTokenAmount struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSeizedVTokenAmount is a free log retrieval operation binding the contract event 0xfde0eb536fecee7fca644722fc2e6255ee0571307ab7f46f7fea23fe302d4479.
//
// Solidity: event SeizedVTokenAmount(uint256 arg0, uint256 arg1, uint256 arg2)
func (_UniFlashSwap *UniFlashSwapFilterer) FilterSeizedVTokenAmount(opts *bind.FilterOpts) (*UniFlashSwapSeizedVTokenAmountIterator, error) {

	logs, sub, err := _UniFlashSwap.contract.FilterLogs(opts, "SeizedVTokenAmount")
	if err != nil {
		return nil, err
	}
	return &UniFlashSwapSeizedVTokenAmountIterator{contract: _UniFlashSwap.contract, event: "SeizedVTokenAmount", logs: logs, sub: sub}, nil
}

// WatchSeizedVTokenAmount is a free log subscription operation binding the contract event 0xfde0eb536fecee7fca644722fc2e6255ee0571307ab7f46f7fea23fe302d4479.
//
// Solidity: event SeizedVTokenAmount(uint256 arg0, uint256 arg1, uint256 arg2)
func (_UniFlashSwap *UniFlashSwapFilterer) WatchSeizedVTokenAmount(opts *bind.WatchOpts, sink chan<- *UniFlashSwapSeizedVTokenAmount) (event.Subscription, error) {

	logs, sub, err := _UniFlashSwap.contract.WatchLogs(opts, "SeizedVTokenAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniFlashSwapSeizedVTokenAmount)
				if err := _UniFlashSwap.contract.UnpackLog(event, "SeizedVTokenAmount", log); err != nil {
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

// ParseSeizedVTokenAmount is a log parse operation binding the contract event 0xfde0eb536fecee7fca644722fc2e6255ee0571307ab7f46f7fea23fe302d4479.
//
// Solidity: event SeizedVTokenAmount(uint256 arg0, uint256 arg1, uint256 arg2)
func (_UniFlashSwap *UniFlashSwapFilterer) ParseSeizedVTokenAmount(log types.Log) (*UniFlashSwapSeizedVTokenAmount, error) {
	event := new(UniFlashSwapSeizedVTokenAmount)
	if err := _UniFlashSwap.contract.UnpackLog(event, "SeizedVTokenAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniFlashSwapWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the UniFlashSwap contract.
type UniFlashSwapWithdrawIterator struct {
	Event *UniFlashSwapWithdraw // Event containing the contract specifics and raw log

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
func (it *UniFlashSwapWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniFlashSwapWithdraw)
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
		it.Event = new(UniFlashSwapWithdraw)
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
func (it *UniFlashSwapWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniFlashSwapWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniFlashSwapWithdraw represents a Withdraw event raised by the UniFlashSwap contract.
type UniFlashSwapWithdraw struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed arg0, address indexed arg1, uint256 arg2)
func (_UniFlashSwap *UniFlashSwapFilterer) FilterWithdraw(opts *bind.FilterOpts, arg0 []common.Address, arg1 []common.Address) (*UniFlashSwapWithdrawIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _UniFlashSwap.contract.FilterLogs(opts, "Withdraw", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return &UniFlashSwapWithdrawIterator{contract: _UniFlashSwap.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed arg0, address indexed arg1, uint256 arg2)
func (_UniFlashSwap *UniFlashSwapFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *UniFlashSwapWithdraw, arg0 []common.Address, arg1 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _UniFlashSwap.contract.WatchLogs(opts, "Withdraw", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniFlashSwapWithdraw)
				if err := _UniFlashSwap.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed arg0, address indexed arg1, uint256 arg2)
func (_UniFlashSwap *UniFlashSwapFilterer) ParseWithdraw(log types.Log) (*UniFlashSwapWithdraw, error) {
	event := new(UniFlashSwapWithdraw)
	if err := _UniFlashSwap.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniFlashSwapWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the UniFlashSwap contract.
type UniFlashSwapWithdrawETHIterator struct {
	Event *UniFlashSwapWithdrawETH // Event containing the contract specifics and raw log

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
func (it *UniFlashSwapWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniFlashSwapWithdrawETH)
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
		it.Event = new(UniFlashSwapWithdrawETH)
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
func (it *UniFlashSwapWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniFlashSwapWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniFlashSwapWithdrawETH represents a WithdrawETH event raised by the UniFlashSwap contract.
type UniFlashSwapWithdrawETH struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0x566e45b1c8057e725bf62796a7f1d37ae294393cab069725a09daddd1af98b79.
//
// Solidity: event WithdrawETH(address indexed arg0, uint256 arg1)
func (_UniFlashSwap *UniFlashSwapFilterer) FilterWithdrawETH(opts *bind.FilterOpts, arg0 []common.Address) (*UniFlashSwapWithdrawETHIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _UniFlashSwap.contract.FilterLogs(opts, "WithdrawETH", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &UniFlashSwapWithdrawETHIterator{contract: _UniFlashSwap.contract, event: "WithdrawETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0x566e45b1c8057e725bf62796a7f1d37ae294393cab069725a09daddd1af98b79.
//
// Solidity: event WithdrawETH(address indexed arg0, uint256 arg1)
func (_UniFlashSwap *UniFlashSwapFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *UniFlashSwapWithdrawETH, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _UniFlashSwap.contract.WatchLogs(opts, "WithdrawETH", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniFlashSwapWithdrawETH)
				if err := _UniFlashSwap.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0x566e45b1c8057e725bf62796a7f1d37ae294393cab069725a09daddd1af98b79.
//
// Solidity: event WithdrawETH(address indexed arg0, uint256 arg1)
func (_UniFlashSwap *UniFlashSwapFilterer) ParseWithdrawETH(log types.Log) (*UniFlashSwapWithdrawETH, error) {
	event := new(UniFlashSwapWithdrawETH)
	if err := _UniFlashSwap.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceMetaData contains all meta data concerning the VTokenInterface contract.
var VTokenInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"MintBehalf\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"RedeemFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractVTokenInterface\",\"name\":\"vTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3e941010": "_addReserves(uint256)",
		"a6afed95": "accrueInterest()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"3af9e669": "balanceOfUnderlying(address)",
		"c5ebeaec": "borrow(uint256)",
		"17bfdfbc": "borrowBalanceCurrent(address)",
		"95dd9193": "borrowBalanceStored(address)",
		"f8f9da28": "borrowRatePerBlock()",
		"bd6d894d": "exchangeRateCurrent()",
		"182df0f5": "exchangeRateStored()",
		"c37f68e2": "getAccountSnapshot(address)",
		"3b1d21a2": "getCash()",
		"f5e3c462": "liquidateBorrow(address,uint256,address)",
		"a0712d68": "mint(uint256)",
		"db006a75": "redeem(uint256)",
		"852a12e3": "redeemUnderlying(uint256)",
		"0e752702": "repayBorrow(uint256)",
		"2608f818": "repayBorrowBehalf(address,uint256)",
		"b2a02ff1": "seize(address,address,uint256)",
		"ae9d70b0": "supplyRatePerBlock()",
		"73acee98": "totalBorrowsCurrent()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"6f307dc3": "underlying()",
	},
}

// VTokenInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use VTokenInterfaceMetaData.ABI instead.
var VTokenInterfaceABI = VTokenInterfaceMetaData.ABI

// Deprecated: Use VTokenInterfaceMetaData.Sigs instead.
// VTokenInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var VTokenInterfaceFuncSigs = VTokenInterfaceMetaData.Sigs

// VTokenInterface is an auto generated Go binding around an Ethereum contract.
type VTokenInterface struct {
	VTokenInterfaceCaller     // Read-only binding to the contract
	VTokenInterfaceTransactor // Write-only binding to the contract
	VTokenInterfaceFilterer   // Log filterer for contract events
}

// VTokenInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type VTokenInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VTokenInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VTokenInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VTokenInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VTokenInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VTokenInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VTokenInterfaceSession struct {
	Contract     *VTokenInterface  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VTokenInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VTokenInterfaceCallerSession struct {
	Contract *VTokenInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VTokenInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VTokenInterfaceTransactorSession struct {
	Contract     *VTokenInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VTokenInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type VTokenInterfaceRaw struct {
	Contract *VTokenInterface // Generic contract binding to access the raw methods on
}

// VTokenInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VTokenInterfaceCallerRaw struct {
	Contract *VTokenInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// VTokenInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VTokenInterfaceTransactorRaw struct {
	Contract *VTokenInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVTokenInterface creates a new instance of VTokenInterface, bound to a specific deployed contract.
func NewVTokenInterface(address common.Address, backend bind.ContractBackend) (*VTokenInterface, error) {
	contract, err := bindVTokenInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VTokenInterface{VTokenInterfaceCaller: VTokenInterfaceCaller{contract: contract}, VTokenInterfaceTransactor: VTokenInterfaceTransactor{contract: contract}, VTokenInterfaceFilterer: VTokenInterfaceFilterer{contract: contract}}, nil
}

// NewVTokenInterfaceCaller creates a new read-only instance of VTokenInterface, bound to a specific deployed contract.
func NewVTokenInterfaceCaller(address common.Address, caller bind.ContractCaller) (*VTokenInterfaceCaller, error) {
	contract, err := bindVTokenInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceCaller{contract: contract}, nil
}

// NewVTokenInterfaceTransactor creates a new write-only instance of VTokenInterface, bound to a specific deployed contract.
func NewVTokenInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*VTokenInterfaceTransactor, error) {
	contract, err := bindVTokenInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceTransactor{contract: contract}, nil
}

// NewVTokenInterfaceFilterer creates a new log filterer instance of VTokenInterface, bound to a specific deployed contract.
func NewVTokenInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*VTokenInterfaceFilterer, error) {
	contract, err := bindVTokenInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceFilterer{contract: contract}, nil
}

// bindVTokenInterface binds a generic wrapper to an already deployed contract.
func bindVTokenInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VTokenInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VTokenInterface *VTokenInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VTokenInterface.Contract.VTokenInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VTokenInterface *VTokenInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VTokenInterface.Contract.VTokenInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VTokenInterface *VTokenInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VTokenInterface.Contract.VTokenInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VTokenInterface *VTokenInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VTokenInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VTokenInterface *VTokenInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VTokenInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VTokenInterface *VTokenInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VTokenInterface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VTokenInterface.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VTokenInterface.Contract.Allowance(&_VTokenInterface.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VTokenInterface.Contract.Allowance(&_VTokenInterface.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VTokenInterface.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VTokenInterface.Contract.BalanceOf(&_VTokenInterface.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VTokenInterface.Contract.BalanceOf(&_VTokenInterface.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VTokenInterface.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _VTokenInterface.Contract.BorrowBalanceStored(&_VTokenInterface.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _VTokenInterface.Contract.BorrowBalanceStored(&_VTokenInterface.CallOpts, account)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VTokenInterface.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) BorrowRatePerBlock() (*big.Int, error) {
	return _VTokenInterface.Contract.BorrowRatePerBlock(&_VTokenInterface.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _VTokenInterface.Contract.BorrowRatePerBlock(&_VTokenInterface.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VTokenInterface.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) ExchangeRateStored() (*big.Int, error) {
	return _VTokenInterface.Contract.ExchangeRateStored(&_VTokenInterface.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _VTokenInterface.Contract.ExchangeRateStored(&_VTokenInterface.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_VTokenInterface *VTokenInterfaceCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _VTokenInterface.contract.Call(opts, &out, "getAccountSnapshot", account)

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
func (_VTokenInterface *VTokenInterfaceSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _VTokenInterface.Contract.GetAccountSnapshot(&_VTokenInterface.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_VTokenInterface *VTokenInterfaceCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _VTokenInterface.Contract.GetAccountSnapshot(&_VTokenInterface.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VTokenInterface.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) GetCash() (*big.Int, error) {
	return _VTokenInterface.Contract.GetCash(&_VTokenInterface.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCallerSession) GetCash() (*big.Int, error) {
	return _VTokenInterface.Contract.GetCash(&_VTokenInterface.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VTokenInterface.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) SupplyRatePerBlock() (*big.Int, error) {
	return _VTokenInterface.Contract.SupplyRatePerBlock(&_VTokenInterface.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_VTokenInterface *VTokenInterfaceCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _VTokenInterface.Contract.SupplyRatePerBlock(&_VTokenInterface.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_VTokenInterface *VTokenInterfaceCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VTokenInterface.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_VTokenInterface *VTokenInterfaceSession) Underlying() (common.Address, error) {
	return _VTokenInterface.Contract.Underlying(&_VTokenInterface.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_VTokenInterface *VTokenInterfaceCallerSession) Underlying() (common.Address, error) {
	return _VTokenInterface.Contract.Underlying(&_VTokenInterface.CallOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.AddReserves(&_VTokenInterface.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.AddReserves(&_VTokenInterface.TransactOpts, addAmount)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) AccrueInterest() (*types.Transaction, error) {
	return _VTokenInterface.Contract.AccrueInterest(&_VTokenInterface.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _VTokenInterface.Contract.AccrueInterest(&_VTokenInterface.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VTokenInterface *VTokenInterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VTokenInterface *VTokenInterfaceSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Approve(&_VTokenInterface.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VTokenInterface *VTokenInterfaceTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Approve(&_VTokenInterface.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _VTokenInterface.Contract.BalanceOfUnderlying(&_VTokenInterface.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _VTokenInterface.Contract.BalanceOfUnderlying(&_VTokenInterface.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Borrow(&_VTokenInterface.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Borrow(&_VTokenInterface.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _VTokenInterface.Contract.BorrowBalanceCurrent(&_VTokenInterface.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _VTokenInterface.Contract.BorrowBalanceCurrent(&_VTokenInterface.TransactOpts, account)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _VTokenInterface.Contract.ExchangeRateCurrent(&_VTokenInterface.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _VTokenInterface.Contract.ExchangeRateCurrent(&_VTokenInterface.TransactOpts)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _VTokenInterface.Contract.LiquidateBorrow(&_VTokenInterface.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _VTokenInterface.Contract.LiquidateBorrow(&_VTokenInterface.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Mint(&_VTokenInterface.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Mint(&_VTokenInterface.TransactOpts, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Redeem(&_VTokenInterface.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Redeem(&_VTokenInterface.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.RedeemUnderlying(&_VTokenInterface.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.RedeemUnderlying(&_VTokenInterface.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.RepayBorrow(&_VTokenInterface.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.RepayBorrow(&_VTokenInterface.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.RepayBorrowBehalf(&_VTokenInterface.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.RepayBorrowBehalf(&_VTokenInterface.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Seize(&_VTokenInterface.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Seize(&_VTokenInterface.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_VTokenInterface *VTokenInterfaceSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _VTokenInterface.Contract.TotalBorrowsCurrent(&_VTokenInterface.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_VTokenInterface *VTokenInterfaceTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _VTokenInterface.Contract.TotalBorrowsCurrent(&_VTokenInterface.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_VTokenInterface *VTokenInterfaceTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_VTokenInterface *VTokenInterfaceSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Transfer(&_VTokenInterface.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_VTokenInterface *VTokenInterfaceTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.Transfer(&_VTokenInterface.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_VTokenInterface *VTokenInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_VTokenInterface *VTokenInterfaceSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.TransferFrom(&_VTokenInterface.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_VTokenInterface *VTokenInterfaceTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VTokenInterface.Contract.TransferFrom(&_VTokenInterface.TransactOpts, src, dst, amount)
}

// VTokenInterfaceAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the VTokenInterface contract.
type VTokenInterfaceAccrueInterestIterator struct {
	Event *VTokenInterfaceAccrueInterest // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceAccrueInterest)
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
		it.Event = new(VTokenInterfaceAccrueInterest)
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
func (it *VTokenInterfaceAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceAccrueInterest represents a AccrueInterest event raised by the VTokenInterface contract.
type VTokenInterfaceAccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*VTokenInterfaceAccrueInterestIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceAccrueInterestIterator{contract: _VTokenInterface.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceAccrueInterest)
				if err := _VTokenInterface.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseAccrueInterest(log types.Log) (*VTokenInterfaceAccrueInterest, error) {
	event := new(VTokenInterfaceAccrueInterest)
	if err := _VTokenInterface.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VTokenInterface contract.
type VTokenInterfaceApprovalIterator struct {
	Event *VTokenInterfaceApproval // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceApproval)
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
		it.Event = new(VTokenInterfaceApproval)
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
func (it *VTokenInterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceApproval represents a Approval event raised by the VTokenInterface contract.
type VTokenInterfaceApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VTokenInterfaceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceApprovalIterator{contract: _VTokenInterface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceApproval)
				if err := _VTokenInterface.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseApproval(log types.Log) (*VTokenInterfaceApproval, error) {
	event := new(VTokenInterfaceApproval)
	if err := _VTokenInterface.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the VTokenInterface contract.
type VTokenInterfaceBorrowIterator struct {
	Event *VTokenInterfaceBorrow // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceBorrow)
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
		it.Event = new(VTokenInterfaceBorrow)
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
func (it *VTokenInterfaceBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceBorrow represents a Borrow event raised by the VTokenInterface contract.
type VTokenInterfaceBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterBorrow(opts *bind.FilterOpts) (*VTokenInterfaceBorrowIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceBorrowIterator{contract: _VTokenInterface.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceBorrow) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceBorrow)
				if err := _VTokenInterface.contract.UnpackLog(event, "Borrow", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseBorrow(log types.Log) (*VTokenInterfaceBorrow, error) {
	event := new(VTokenInterfaceBorrow)
	if err := _VTokenInterface.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the VTokenInterface contract.
type VTokenInterfaceFailureIterator struct {
	Event *VTokenInterfaceFailure // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceFailure)
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
		it.Event = new(VTokenInterfaceFailure)
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
func (it *VTokenInterfaceFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceFailure represents a Failure event raised by the VTokenInterface contract.
type VTokenInterfaceFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterFailure(opts *bind.FilterOpts) (*VTokenInterfaceFailureIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceFailureIterator{contract: _VTokenInterface.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceFailure) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceFailure)
				if err := _VTokenInterface.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseFailure(log types.Log) (*VTokenInterfaceFailure, error) {
	event := new(VTokenInterfaceFailure)
	if err := _VTokenInterface.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the VTokenInterface contract.
type VTokenInterfaceLiquidateBorrowIterator struct {
	Event *VTokenInterfaceLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceLiquidateBorrow)
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
		it.Event = new(VTokenInterfaceLiquidateBorrow)
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
func (it *VTokenInterfaceLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceLiquidateBorrow represents a LiquidateBorrow event raised by the VTokenInterface contract.
type VTokenInterfaceLiquidateBorrow struct {
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
func (_VTokenInterface *VTokenInterfaceFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*VTokenInterfaceLiquidateBorrowIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceLiquidateBorrowIterator{contract: _VTokenInterface.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceLiquidateBorrow)
				if err := _VTokenInterface.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseLiquidateBorrow(log types.Log) (*VTokenInterfaceLiquidateBorrow, error) {
	event := new(VTokenInterfaceLiquidateBorrow)
	if err := _VTokenInterface.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the VTokenInterface contract.
type VTokenInterfaceMintIterator struct {
	Event *VTokenInterfaceMint // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceMint)
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
		it.Event = new(VTokenInterfaceMint)
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
func (it *VTokenInterfaceMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceMint represents a Mint event raised by the VTokenInterface contract.
type VTokenInterfaceMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterMint(opts *bind.FilterOpts) (*VTokenInterfaceMintIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceMintIterator{contract: _VTokenInterface.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceMint) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceMint)
				if err := _VTokenInterface.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseMint(log types.Log) (*VTokenInterfaceMint, error) {
	event := new(VTokenInterfaceMint)
	if err := _VTokenInterface.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceMintBehalfIterator is returned from FilterMintBehalf and is used to iterate over the raw logs and unpacked data for MintBehalf events raised by the VTokenInterface contract.
type VTokenInterfaceMintBehalfIterator struct {
	Event *VTokenInterfaceMintBehalf // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceMintBehalfIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceMintBehalf)
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
		it.Event = new(VTokenInterfaceMintBehalf)
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
func (it *VTokenInterfaceMintBehalfIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceMintBehalfIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceMintBehalf represents a MintBehalf event raised by the VTokenInterface contract.
type VTokenInterfaceMintBehalf struct {
	Payer      common.Address
	Receiver   common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintBehalf is a free log retrieval operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterMintBehalf(opts *bind.FilterOpts) (*VTokenInterfaceMintBehalfIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "MintBehalf")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceMintBehalfIterator{contract: _VTokenInterface.contract, event: "MintBehalf", logs: logs, sub: sub}, nil
}

// WatchMintBehalf is a free log subscription operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchMintBehalf(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceMintBehalf) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "MintBehalf")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceMintBehalf)
				if err := _VTokenInterface.contract.UnpackLog(event, "MintBehalf", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseMintBehalf(log types.Log) (*VTokenInterfaceMintBehalf, error) {
	event := new(VTokenInterfaceMintBehalf)
	if err := _VTokenInterface.contract.UnpackLog(event, "MintBehalf", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the VTokenInterface contract.
type VTokenInterfaceNewAdminIterator struct {
	Event *VTokenInterfaceNewAdmin // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceNewAdmin)
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
		it.Event = new(VTokenInterfaceNewAdmin)
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
func (it *VTokenInterfaceNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceNewAdmin represents a NewAdmin event raised by the VTokenInterface contract.
type VTokenInterfaceNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*VTokenInterfaceNewAdminIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceNewAdminIterator{contract: _VTokenInterface.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceNewAdmin) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceNewAdmin)
				if err := _VTokenInterface.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseNewAdmin(log types.Log) (*VTokenInterfaceNewAdmin, error) {
	event := new(VTokenInterfaceNewAdmin)
	if err := _VTokenInterface.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the VTokenInterface contract.
type VTokenInterfaceNewPendingAdminIterator struct {
	Event *VTokenInterfaceNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceNewPendingAdmin)
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
		it.Event = new(VTokenInterfaceNewPendingAdmin)
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
func (it *VTokenInterfaceNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceNewPendingAdmin represents a NewPendingAdmin event raised by the VTokenInterface contract.
type VTokenInterfaceNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*VTokenInterfaceNewPendingAdminIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceNewPendingAdminIterator{contract: _VTokenInterface.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceNewPendingAdmin)
				if err := _VTokenInterface.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseNewPendingAdmin(log types.Log) (*VTokenInterfaceNewPendingAdmin, error) {
	event := new(VTokenInterfaceNewPendingAdmin)
	if err := _VTokenInterface.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the VTokenInterface contract.
type VTokenInterfaceNewReserveFactorIterator struct {
	Event *VTokenInterfaceNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceNewReserveFactor)
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
		it.Event = new(VTokenInterfaceNewReserveFactor)
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
func (it *VTokenInterfaceNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceNewReserveFactor represents a NewReserveFactor event raised by the VTokenInterface contract.
type VTokenInterfaceNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*VTokenInterfaceNewReserveFactorIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceNewReserveFactorIterator{contract: _VTokenInterface.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceNewReserveFactor)
				if err := _VTokenInterface.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseNewReserveFactor(log types.Log) (*VTokenInterfaceNewReserveFactor, error) {
	event := new(VTokenInterfaceNewReserveFactor)
	if err := _VTokenInterface.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the VTokenInterface contract.
type VTokenInterfaceRedeemIterator struct {
	Event *VTokenInterfaceRedeem // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceRedeem)
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
		it.Event = new(VTokenInterfaceRedeem)
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
func (it *VTokenInterfaceRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceRedeem represents a Redeem event raised by the VTokenInterface contract.
type VTokenInterfaceRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterRedeem(opts *bind.FilterOpts) (*VTokenInterfaceRedeemIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceRedeemIterator{contract: _VTokenInterface.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceRedeem) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceRedeem)
				if err := _VTokenInterface.contract.UnpackLog(event, "Redeem", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseRedeem(log types.Log) (*VTokenInterfaceRedeem, error) {
	event := new(VTokenInterfaceRedeem)
	if err := _VTokenInterface.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceRedeemFeeIterator is returned from FilterRedeemFee and is used to iterate over the raw logs and unpacked data for RedeemFee events raised by the VTokenInterface contract.
type VTokenInterfaceRedeemFeeIterator struct {
	Event *VTokenInterfaceRedeemFee // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceRedeemFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceRedeemFee)
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
		it.Event = new(VTokenInterfaceRedeemFee)
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
func (it *VTokenInterfaceRedeemFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceRedeemFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceRedeemFee represents a RedeemFee event raised by the VTokenInterface contract.
type VTokenInterfaceRedeemFee struct {
	Redeemer     common.Address
	FeeAmount    *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeemFee is a free log retrieval operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterRedeemFee(opts *bind.FilterOpts) (*VTokenInterfaceRedeemFeeIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "RedeemFee")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceRedeemFeeIterator{contract: _VTokenInterface.contract, event: "RedeemFee", logs: logs, sub: sub}, nil
}

// WatchRedeemFee is a free log subscription operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchRedeemFee(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceRedeemFee) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "RedeemFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceRedeemFee)
				if err := _VTokenInterface.contract.UnpackLog(event, "RedeemFee", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseRedeemFee(log types.Log) (*VTokenInterfaceRedeemFee, error) {
	event := new(VTokenInterfaceRedeemFee)
	if err := _VTokenInterface.contract.UnpackLog(event, "RedeemFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the VTokenInterface contract.
type VTokenInterfaceRepayBorrowIterator struct {
	Event *VTokenInterfaceRepayBorrow // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceRepayBorrow)
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
		it.Event = new(VTokenInterfaceRepayBorrow)
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
func (it *VTokenInterfaceRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceRepayBorrow represents a RepayBorrow event raised by the VTokenInterface contract.
type VTokenInterfaceRepayBorrow struct {
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
func (_VTokenInterface *VTokenInterfaceFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*VTokenInterfaceRepayBorrowIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceRepayBorrowIterator{contract: _VTokenInterface.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceRepayBorrow)
				if err := _VTokenInterface.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseRepayBorrow(log types.Log) (*VTokenInterfaceRepayBorrow, error) {
	event := new(VTokenInterfaceRepayBorrow)
	if err := _VTokenInterface.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the VTokenInterface contract.
type VTokenInterfaceReservesAddedIterator struct {
	Event *VTokenInterfaceReservesAdded // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceReservesAdded)
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
		it.Event = new(VTokenInterfaceReservesAdded)
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
func (it *VTokenInterfaceReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceReservesAdded represents a ReservesAdded event raised by the VTokenInterface contract.
type VTokenInterfaceReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterReservesAdded(opts *bind.FilterOpts) (*VTokenInterfaceReservesAddedIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceReservesAddedIterator{contract: _VTokenInterface.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceReservesAdded) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceReservesAdded)
				if err := _VTokenInterface.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseReservesAdded(log types.Log) (*VTokenInterfaceReservesAdded, error) {
	event := new(VTokenInterfaceReservesAdded)
	if err := _VTokenInterface.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the VTokenInterface contract.
type VTokenInterfaceReservesReducedIterator struct {
	Event *VTokenInterfaceReservesReduced // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceReservesReduced)
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
		it.Event = new(VTokenInterfaceReservesReduced)
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
func (it *VTokenInterfaceReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceReservesReduced represents a ReservesReduced event raised by the VTokenInterface contract.
type VTokenInterfaceReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*VTokenInterfaceReservesReducedIterator, error) {

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceReservesReducedIterator{contract: _VTokenInterface.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceReservesReduced) (event.Subscription, error) {

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceReservesReduced)
				if err := _VTokenInterface.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseReservesReduced(log types.Log) (*VTokenInterfaceReservesReduced, error) {
	event := new(VTokenInterfaceReservesReduced)
	if err := _VTokenInterface.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VTokenInterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VTokenInterface contract.
type VTokenInterfaceTransferIterator struct {
	Event *VTokenInterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *VTokenInterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VTokenInterfaceTransfer)
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
		it.Event = new(VTokenInterfaceTransfer)
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
func (it *VTokenInterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VTokenInterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VTokenInterfaceTransfer represents a Transfer event raised by the VTokenInterface contract.
type VTokenInterfaceTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VTokenInterface *VTokenInterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VTokenInterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VTokenInterface.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VTokenInterfaceTransferIterator{contract: _VTokenInterface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VTokenInterface *VTokenInterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VTokenInterfaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VTokenInterface.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VTokenInterfaceTransfer)
				if err := _VTokenInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_VTokenInterface *VTokenInterfaceFilterer) ParseTransfer(log types.Log) (*VTokenInterfaceTransfer, error) {
	event := new(VTokenInterfaceTransfer)
	if err := _VTokenInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
