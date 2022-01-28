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

// VenusLensAccountLimits is an auto generated low-level Go binding around an user-defined struct.
type VenusLensAccountLimits struct {
	Markets   []common.Address
	Liquidity *big.Int
	Shortfall *big.Int
}

// VenusLensGovProposal is an auto generated low-level Go binding around an user-defined struct.
type VenusLensGovProposal struct {
	ProposalId   *big.Int
	Proposer     common.Address
	Eta          *big.Int
	Targets      []common.Address
	Values       []*big.Int
	Signatures   []string
	Calldatas    [][]byte
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Canceled     bool
	Executed     bool
}

// VenusLensGovReceipt is an auto generated low-level Go binding around an user-defined struct.
type VenusLensGovReceipt struct {
	ProposalId *big.Int
	HasVoted   bool
	Support    bool
	Votes      *big.Int
}

// VenusLensVTokenBalances is an auto generated low-level Go binding around an user-defined struct.
type VenusLensVTokenBalances struct {
	VToken               common.Address
	BalanceOf            *big.Int
	BorrowBalanceCurrent *big.Int
	BalanceOfUnderlying  *big.Int
	TokenBalance         *big.Int
	TokenAllowance       *big.Int
}

// VenusLensVTokenMetadata is an auto generated low-level Go binding around an user-defined struct.
type VenusLensVTokenMetadata struct {
	VToken                   common.Address
	ExchangeRateCurrent      *big.Int
	SupplyRatePerBlock       *big.Int
	BorrowRatePerBlock       *big.Int
	ReserveFactorMantissa    *big.Int
	TotalBorrows             *big.Int
	TotalReserves            *big.Int
	TotalSupply              *big.Int
	TotalCash                *big.Int
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	UnderlyingAssetAddress   common.Address
	VTokenDecimals           *big.Int
	UnderlyingDecimals       *big.Int
}

// VenusLensVTokenUnderlyingPrice is an auto generated low-level Go binding around an user-defined struct.
type VenusLensVTokenUnderlyingPrice struct {
	VToken          common.Address
	UnderlyingPrice *big.Int
}

// VenusLensVenusVotes is an auto generated low-level Go binding around an user-defined struct.
type VenusLensVenusVotes struct {
	BlockNumber *big.Int
	Votes       *big.Int
}

// VenusLensXVSBalanceMetadata is an auto generated low-level Go binding around an user-defined struct.
type VenusLensXVSBalanceMetadata struct {
	Balance  *big.Int
	Votes    *big.Int
	Delegate common.Address
}

// VenusLensXVSBalanceMetadataExt is an auto generated low-level Go binding around an user-defined struct.
type VenusLensXVSBalanceMetadataExt struct {
	Balance   *big.Int
	Votes     *big.Int
	Delegate  common.Address
	Allocated *big.Int
}

// VlensMetaData contains all meta data concerning the Vlens contract.
var VlensMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"contractComptrollerLensInterface\",\"name\":\"comptroller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLimits\",\"outputs\":[{\"components\":[{\"internalType\":\"contractVToken[]\",\"name\":\"markets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfall\",\"type\":\"uint256\"}],\"internalType\":\"structVenusLens.AccountLimits\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractGovernorAlpha\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proposalIds\",\"type\":\"uint256[]\"}],\"name\":\"getGovProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"internalType\":\"structVenusLens.GovProposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractGovernorAlpha\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proposalIds\",\"type\":\"uint256[]\"}],\"name\":\"getGovReceipts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"internalType\":\"structVenusLens.GovReceipt[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractXVS\",\"name\":\"xvs\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"blockNumbers\",\"type\":\"uint32[]\"}],\"name\":\"getVenusVotes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structVenusLens.VenusVotes[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractXVS\",\"name\":\"xvs\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getXVSBalanceMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"internalType\":\"structVenusLens.XVSBalanceMetadata\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractXVS\",\"name\":\"xvs\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerLensInterface\",\"name\":\"comptroller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getXVSBalanceMetadataExt\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocated\",\"type\":\"uint256\"}],\"internalType\":\"structVenusLens.XVSBalanceMetadataExt\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"vTokenBalances\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceOf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowBalanceCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOfUnderlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAllowance\",\"type\":\"uint256\"}],\"internalType\":\"structVenusLens.VTokenBalances\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"vTokenBalancesAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceOf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowBalanceCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOfUnderlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAllowance\",\"type\":\"uint256\"}],\"internalType\":\"structVenusLens.VTokenBalances[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"vTokenMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRateCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCash\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlyingAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vTokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingDecimals\",\"type\":\"uint256\"}],\"internalType\":\"structVenusLens.VTokenMetadata\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"vTokenMetadataAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRateCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCash\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlyingAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vTokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingDecimals\",\"type\":\"uint256\"}],\"internalType\":\"structVenusLens.VTokenMetadata[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"vTokenUnderlyingPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"}],\"internalType\":\"structVenusLens.VTokenUnderlyingPrice\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"vTokenUnderlyingPriceAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"}],\"internalType\":\"structVenusLens.VTokenUnderlyingPrice[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VlensABI is the input ABI used to generate the binding from.
// Deprecated: Use VlensMetaData.ABI instead.
var VlensABI = VlensMetaData.ABI

// Vlens is an auto generated Go binding around an Ethereum contract.
type Vlens struct {
	VlensCaller     // Read-only binding to the contract
	VlensTransactor // Write-only binding to the contract
	VlensFilterer   // Log filterer for contract events
}

// VlensCaller is an auto generated read-only Go binding around an Ethereum contract.
type VlensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VlensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VlensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VlensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VlensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VlensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VlensSession struct {
	Contract     *Vlens            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VlensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VlensCallerSession struct {
	Contract *VlensCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VlensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VlensTransactorSession struct {
	Contract     *VlensTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VlensRaw is an auto generated low-level Go binding around an Ethereum contract.
type VlensRaw struct {
	Contract *Vlens // Generic contract binding to access the raw methods on
}

// VlensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VlensCallerRaw struct {
	Contract *VlensCaller // Generic read-only contract binding to access the raw methods on
}

// VlensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VlensTransactorRaw struct {
	Contract *VlensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVlens creates a new instance of Vlens, bound to a specific deployed contract.
func NewVlens(address common.Address, backend bind.ContractBackend) (*Vlens, error) {
	contract, err := bindVlens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vlens{VlensCaller: VlensCaller{contract: contract}, VlensTransactor: VlensTransactor{contract: contract}, VlensFilterer: VlensFilterer{contract: contract}}, nil
}

// NewVlensCaller creates a new read-only instance of Vlens, bound to a specific deployed contract.
func NewVlensCaller(address common.Address, caller bind.ContractCaller) (*VlensCaller, error) {
	contract, err := bindVlens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VlensCaller{contract: contract}, nil
}

// NewVlensTransactor creates a new write-only instance of Vlens, bound to a specific deployed contract.
func NewVlensTransactor(address common.Address, transactor bind.ContractTransactor) (*VlensTransactor, error) {
	contract, err := bindVlens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VlensTransactor{contract: contract}, nil
}

// NewVlensFilterer creates a new log filterer instance of Vlens, bound to a specific deployed contract.
func NewVlensFilterer(address common.Address, filterer bind.ContractFilterer) (*VlensFilterer, error) {
	contract, err := bindVlens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VlensFilterer{contract: contract}, nil
}

// bindVlens binds a generic wrapper to an already deployed contract.
func bindVlens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VlensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vlens *VlensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vlens.Contract.VlensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vlens *VlensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vlens.Contract.VlensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vlens *VlensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vlens.Contract.VlensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vlens *VlensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vlens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vlens *VlensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vlens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vlens *VlensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vlens.Contract.contract.Transact(opts, method, params...)
}

// GetAccountLimits is a free data retrieval call binding the contract method 0x7dd8f6d9.
//
// Solidity: function getAccountLimits(address comptroller, address account) view returns((address[],uint256,uint256))
func (_Vlens *VlensCaller) GetAccountLimits(opts *bind.CallOpts, comptroller common.Address, account common.Address) (VenusLensAccountLimits, error) {
	var out []interface{}
	err := _Vlens.contract.Call(opts, &out, "getAccountLimits", comptroller, account)

	if err != nil {
		return *new(VenusLensAccountLimits), err
	}

	out0 := *abi.ConvertType(out[0], new(VenusLensAccountLimits)).(*VenusLensAccountLimits)

	return out0, err

}

// GetAccountLimits is a free data retrieval call binding the contract method 0x7dd8f6d9.
//
// Solidity: function getAccountLimits(address comptroller, address account) view returns((address[],uint256,uint256))
func (_Vlens *VlensSession) GetAccountLimits(comptroller common.Address, account common.Address) (VenusLensAccountLimits, error) {
	return _Vlens.Contract.GetAccountLimits(&_Vlens.CallOpts, comptroller, account)
}

// GetAccountLimits is a free data retrieval call binding the contract method 0x7dd8f6d9.
//
// Solidity: function getAccountLimits(address comptroller, address account) view returns((address[],uint256,uint256))
func (_Vlens *VlensCallerSession) GetAccountLimits(comptroller common.Address, account common.Address) (VenusLensAccountLimits, error) {
	return _Vlens.Contract.GetAccountLimits(&_Vlens.CallOpts, comptroller, account)
}

// GetGovProposals is a free data retrieval call binding the contract method 0x96994869.
//
// Solidity: function getGovProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,bool,bool)[])
func (_Vlens *VlensCaller) GetGovProposals(opts *bind.CallOpts, governor common.Address, proposalIds []*big.Int) ([]VenusLensGovProposal, error) {
	var out []interface{}
	err := _Vlens.contract.Call(opts, &out, "getGovProposals", governor, proposalIds)

	if err != nil {
		return *new([]VenusLensGovProposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]VenusLensGovProposal)).(*[]VenusLensGovProposal)

	return out0, err

}

// GetGovProposals is a free data retrieval call binding the contract method 0x96994869.
//
// Solidity: function getGovProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,bool,bool)[])
func (_Vlens *VlensSession) GetGovProposals(governor common.Address, proposalIds []*big.Int) ([]VenusLensGovProposal, error) {
	return _Vlens.Contract.GetGovProposals(&_Vlens.CallOpts, governor, proposalIds)
}

// GetGovProposals is a free data retrieval call binding the contract method 0x96994869.
//
// Solidity: function getGovProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,bool,bool)[])
func (_Vlens *VlensCallerSession) GetGovProposals(governor common.Address, proposalIds []*big.Int) ([]VenusLensGovProposal, error) {
	return _Vlens.Contract.GetGovProposals(&_Vlens.CallOpts, governor, proposalIds)
}

// GetGovReceipts is a free data retrieval call binding the contract method 0x995ed99f.
//
// Solidity: function getGovReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,bool,uint96)[])
func (_Vlens *VlensCaller) GetGovReceipts(opts *bind.CallOpts, governor common.Address, voter common.Address, proposalIds []*big.Int) ([]VenusLensGovReceipt, error) {
	var out []interface{}
	err := _Vlens.contract.Call(opts, &out, "getGovReceipts", governor, voter, proposalIds)

	if err != nil {
		return *new([]VenusLensGovReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new([]VenusLensGovReceipt)).(*[]VenusLensGovReceipt)

	return out0, err

}

// GetGovReceipts is a free data retrieval call binding the contract method 0x995ed99f.
//
// Solidity: function getGovReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,bool,uint96)[])
func (_Vlens *VlensSession) GetGovReceipts(governor common.Address, voter common.Address, proposalIds []*big.Int) ([]VenusLensGovReceipt, error) {
	return _Vlens.Contract.GetGovReceipts(&_Vlens.CallOpts, governor, voter, proposalIds)
}

// GetGovReceipts is a free data retrieval call binding the contract method 0x995ed99f.
//
// Solidity: function getGovReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,bool,uint96)[])
func (_Vlens *VlensCallerSession) GetGovReceipts(governor common.Address, voter common.Address, proposalIds []*big.Int) ([]VenusLensGovReceipt, error) {
	return _Vlens.Contract.GetGovReceipts(&_Vlens.CallOpts, governor, voter, proposalIds)
}

// GetVenusVotes is a free data retrieval call binding the contract method 0xfbd88b46.
//
// Solidity: function getVenusVotes(address xvs, address account, uint32[] blockNumbers) view returns((uint256,uint256)[])
func (_Vlens *VlensCaller) GetVenusVotes(opts *bind.CallOpts, xvs common.Address, account common.Address, blockNumbers []uint32) ([]VenusLensVenusVotes, error) {
	var out []interface{}
	err := _Vlens.contract.Call(opts, &out, "getVenusVotes", xvs, account, blockNumbers)

	if err != nil {
		return *new([]VenusLensVenusVotes), err
	}

	out0 := *abi.ConvertType(out[0], new([]VenusLensVenusVotes)).(*[]VenusLensVenusVotes)

	return out0, err

}

// GetVenusVotes is a free data retrieval call binding the contract method 0xfbd88b46.
//
// Solidity: function getVenusVotes(address xvs, address account, uint32[] blockNumbers) view returns((uint256,uint256)[])
func (_Vlens *VlensSession) GetVenusVotes(xvs common.Address, account common.Address, blockNumbers []uint32) ([]VenusLensVenusVotes, error) {
	return _Vlens.Contract.GetVenusVotes(&_Vlens.CallOpts, xvs, account, blockNumbers)
}

// GetVenusVotes is a free data retrieval call binding the contract method 0xfbd88b46.
//
// Solidity: function getVenusVotes(address xvs, address account, uint32[] blockNumbers) view returns((uint256,uint256)[])
func (_Vlens *VlensCallerSession) GetVenusVotes(xvs common.Address, account common.Address, blockNumbers []uint32) ([]VenusLensVenusVotes, error) {
	return _Vlens.Contract.GetVenusVotes(&_Vlens.CallOpts, xvs, account, blockNumbers)
}

// GetXVSBalanceMetadata is a free data retrieval call binding the contract method 0xf40c2777.
//
// Solidity: function getXVSBalanceMetadata(address xvs, address account) view returns((uint256,uint256,address))
func (_Vlens *VlensCaller) GetXVSBalanceMetadata(opts *bind.CallOpts, xvs common.Address, account common.Address) (VenusLensXVSBalanceMetadata, error) {
	var out []interface{}
	err := _Vlens.contract.Call(opts, &out, "getXVSBalanceMetadata", xvs, account)

	if err != nil {
		return *new(VenusLensXVSBalanceMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(VenusLensXVSBalanceMetadata)).(*VenusLensXVSBalanceMetadata)

	return out0, err

}

// GetXVSBalanceMetadata is a free data retrieval call binding the contract method 0xf40c2777.
//
// Solidity: function getXVSBalanceMetadata(address xvs, address account) view returns((uint256,uint256,address))
func (_Vlens *VlensSession) GetXVSBalanceMetadata(xvs common.Address, account common.Address) (VenusLensXVSBalanceMetadata, error) {
	return _Vlens.Contract.GetXVSBalanceMetadata(&_Vlens.CallOpts, xvs, account)
}

// GetXVSBalanceMetadata is a free data retrieval call binding the contract method 0xf40c2777.
//
// Solidity: function getXVSBalanceMetadata(address xvs, address account) view returns((uint256,uint256,address))
func (_Vlens *VlensCallerSession) GetXVSBalanceMetadata(xvs common.Address, account common.Address) (VenusLensXVSBalanceMetadata, error) {
	return _Vlens.Contract.GetXVSBalanceMetadata(&_Vlens.CallOpts, xvs, account)
}

// VTokenUnderlyingPrice is a free data retrieval call binding the contract method 0x7c84e3b3.
//
// Solidity: function vTokenUnderlyingPrice(address vToken) view returns((address,uint256))
func (_Vlens *VlensCaller) VTokenUnderlyingPrice(opts *bind.CallOpts, vToken common.Address) (VenusLensVTokenUnderlyingPrice, error) {
	var out []interface{}
	err := _Vlens.contract.Call(opts, &out, "vTokenUnderlyingPrice", vToken)

	if err != nil {
		return *new(VenusLensVTokenUnderlyingPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(VenusLensVTokenUnderlyingPrice)).(*VenusLensVTokenUnderlyingPrice)

	return out0, err

}

// VTokenUnderlyingPrice is a free data retrieval call binding the contract method 0x7c84e3b3.
//
// Solidity: function vTokenUnderlyingPrice(address vToken) view returns((address,uint256))
func (_Vlens *VlensSession) VTokenUnderlyingPrice(vToken common.Address) (VenusLensVTokenUnderlyingPrice, error) {
	return _Vlens.Contract.VTokenUnderlyingPrice(&_Vlens.CallOpts, vToken)
}

// VTokenUnderlyingPrice is a free data retrieval call binding the contract method 0x7c84e3b3.
//
// Solidity: function vTokenUnderlyingPrice(address vToken) view returns((address,uint256))
func (_Vlens *VlensCallerSession) VTokenUnderlyingPrice(vToken common.Address) (VenusLensVTokenUnderlyingPrice, error) {
	return _Vlens.Contract.VTokenUnderlyingPrice(&_Vlens.CallOpts, vToken)
}

// VTokenUnderlyingPriceAll is a free data retrieval call binding the contract method 0x1f884fdf.
//
// Solidity: function vTokenUnderlyingPriceAll(address[] vTokens) view returns((address,uint256)[])
func (_Vlens *VlensCaller) VTokenUnderlyingPriceAll(opts *bind.CallOpts, vTokens []common.Address) ([]VenusLensVTokenUnderlyingPrice, error) {
	var out []interface{}
	err := _Vlens.contract.Call(opts, &out, "vTokenUnderlyingPriceAll", vTokens)

	if err != nil {
		return *new([]VenusLensVTokenUnderlyingPrice), err
	}

	out0 := *abi.ConvertType(out[0], new([]VenusLensVTokenUnderlyingPrice)).(*[]VenusLensVTokenUnderlyingPrice)

	return out0, err

}

// VTokenUnderlyingPriceAll is a free data retrieval call binding the contract method 0x1f884fdf.
//
// Solidity: function vTokenUnderlyingPriceAll(address[] vTokens) view returns((address,uint256)[])
func (_Vlens *VlensSession) VTokenUnderlyingPriceAll(vTokens []common.Address) ([]VenusLensVTokenUnderlyingPrice, error) {
	return _Vlens.Contract.VTokenUnderlyingPriceAll(&_Vlens.CallOpts, vTokens)
}

// VTokenUnderlyingPriceAll is a free data retrieval call binding the contract method 0x1f884fdf.
//
// Solidity: function vTokenUnderlyingPriceAll(address[] vTokens) view returns((address,uint256)[])
func (_Vlens *VlensCallerSession) VTokenUnderlyingPriceAll(vTokens []common.Address) ([]VenusLensVTokenUnderlyingPrice, error) {
	return _Vlens.Contract.VTokenUnderlyingPriceAll(&_Vlens.CallOpts, vTokens)
}

// GetXVSBalanceMetadataExt is a paid mutator transaction binding the contract method 0xe09744c6.
//
// Solidity: function getXVSBalanceMetadataExt(address xvs, address comptroller, address account) returns((uint256,uint256,address,uint256))
func (_Vlens *VlensTransactor) GetXVSBalanceMetadataExt(opts *bind.TransactOpts, xvs common.Address, comptroller common.Address, account common.Address) (*types.Transaction, error) {
	return _Vlens.contract.Transact(opts, "getXVSBalanceMetadataExt", xvs, comptroller, account)
}

// GetXVSBalanceMetadataExt is a paid mutator transaction binding the contract method 0xe09744c6.
//
// Solidity: function getXVSBalanceMetadataExt(address xvs, address comptroller, address account) returns((uint256,uint256,address,uint256))
func (_Vlens *VlensSession) GetXVSBalanceMetadataExt(xvs common.Address, comptroller common.Address, account common.Address) (*types.Transaction, error) {
	return _Vlens.Contract.GetXVSBalanceMetadataExt(&_Vlens.TransactOpts, xvs, comptroller, account)
}

// GetXVSBalanceMetadataExt is a paid mutator transaction binding the contract method 0xe09744c6.
//
// Solidity: function getXVSBalanceMetadataExt(address xvs, address comptroller, address account) returns((uint256,uint256,address,uint256))
func (_Vlens *VlensTransactorSession) GetXVSBalanceMetadataExt(xvs common.Address, comptroller common.Address, account common.Address) (*types.Transaction, error) {
	return _Vlens.Contract.GetXVSBalanceMetadataExt(&_Vlens.TransactOpts, xvs, comptroller, account)
}

// VTokenBalances is a paid mutator transaction binding the contract method 0xb3124239.
//
// Solidity: function vTokenBalances(address vToken, address account) returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Vlens *VlensTransactor) VTokenBalances(opts *bind.TransactOpts, vToken common.Address, account common.Address) (*types.Transaction, error) {
	return _Vlens.contract.Transact(opts, "vTokenBalances", vToken, account)
}

// VTokenBalances is a paid mutator transaction binding the contract method 0xb3124239.
//
// Solidity: function vTokenBalances(address vToken, address account) returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Vlens *VlensSession) VTokenBalances(vToken common.Address, account common.Address) (*types.Transaction, error) {
	return _Vlens.Contract.VTokenBalances(&_Vlens.TransactOpts, vToken, account)
}

// VTokenBalances is a paid mutator transaction binding the contract method 0xb3124239.
//
// Solidity: function vTokenBalances(address vToken, address account) returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Vlens *VlensTransactorSession) VTokenBalances(vToken common.Address, account common.Address) (*types.Transaction, error) {
	return _Vlens.Contract.VTokenBalances(&_Vlens.TransactOpts, vToken, account)
}

// VTokenBalancesAll is a paid mutator transaction binding the contract method 0x7c51b642.
//
// Solidity: function vTokenBalancesAll(address[] vTokens, address account) returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_Vlens *VlensTransactor) VTokenBalancesAll(opts *bind.TransactOpts, vTokens []common.Address, account common.Address) (*types.Transaction, error) {
	return _Vlens.contract.Transact(opts, "vTokenBalancesAll", vTokens, account)
}

// VTokenBalancesAll is a paid mutator transaction binding the contract method 0x7c51b642.
//
// Solidity: function vTokenBalancesAll(address[] vTokens, address account) returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_Vlens *VlensSession) VTokenBalancesAll(vTokens []common.Address, account common.Address) (*types.Transaction, error) {
	return _Vlens.Contract.VTokenBalancesAll(&_Vlens.TransactOpts, vTokens, account)
}

// VTokenBalancesAll is a paid mutator transaction binding the contract method 0x7c51b642.
//
// Solidity: function vTokenBalancesAll(address[] vTokens, address account) returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_Vlens *VlensTransactorSession) VTokenBalancesAll(vTokens []common.Address, account common.Address) (*types.Transaction, error) {
	return _Vlens.Contract.VTokenBalancesAll(&_Vlens.TransactOpts, vTokens, account)
}

// VTokenMetadata is a paid mutator transaction binding the contract method 0xaa5dbd23.
//
// Solidity: function vTokenMetadata(address vToken) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256))
func (_Vlens *VlensTransactor) VTokenMetadata(opts *bind.TransactOpts, vToken common.Address) (*types.Transaction, error) {
	return _Vlens.contract.Transact(opts, "vTokenMetadata", vToken)
}

// VTokenMetadata is a paid mutator transaction binding the contract method 0xaa5dbd23.
//
// Solidity: function vTokenMetadata(address vToken) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256))
func (_Vlens *VlensSession) VTokenMetadata(vToken common.Address) (*types.Transaction, error) {
	return _Vlens.Contract.VTokenMetadata(&_Vlens.TransactOpts, vToken)
}

// VTokenMetadata is a paid mutator transaction binding the contract method 0xaa5dbd23.
//
// Solidity: function vTokenMetadata(address vToken) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256))
func (_Vlens *VlensTransactorSession) VTokenMetadata(vToken common.Address) (*types.Transaction, error) {
	return _Vlens.Contract.VTokenMetadata(&_Vlens.TransactOpts, vToken)
}

// VTokenMetadataAll is a paid mutator transaction binding the contract method 0xe0a67f11.
//
// Solidity: function vTokenMetadataAll(address[] vTokens) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256)[])
func (_Vlens *VlensTransactor) VTokenMetadataAll(opts *bind.TransactOpts, vTokens []common.Address) (*types.Transaction, error) {
	return _Vlens.contract.Transact(opts, "vTokenMetadataAll", vTokens)
}

// VTokenMetadataAll is a paid mutator transaction binding the contract method 0xe0a67f11.
//
// Solidity: function vTokenMetadataAll(address[] vTokens) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256)[])
func (_Vlens *VlensSession) VTokenMetadataAll(vTokens []common.Address) (*types.Transaction, error) {
	return _Vlens.Contract.VTokenMetadataAll(&_Vlens.TransactOpts, vTokens)
}

// VTokenMetadataAll is a paid mutator transaction binding the contract method 0xe0a67f11.
//
// Solidity: function vTokenMetadataAll(address[] vTokens) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256)[])
func (_Vlens *VlensTransactorSession) VTokenMetadataAll(vTokens []common.Address) (*types.Transaction, error) {
	return _Vlens.Contract.VTokenMetadataAll(&_Vlens.TransactOpts, vTokens)
}
