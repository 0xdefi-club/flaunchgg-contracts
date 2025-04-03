// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_tokens

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
	_ = abi.ConvertType
)

// BurnTokensMetaData contains all meta data concerning the BurnTokens contract.
var BurnTokensMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"execute\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// BurnTokensABI is the input ABI used to generate the binding from.
// Deprecated: Use BurnTokensMetaData.ABI instead.
var BurnTokensABI = BurnTokensMetaData.ABI

// BurnTokens is an auto generated Go binding around an Ethereum contract.
type BurnTokens struct {
	BurnTokensCaller     // Read-only binding to the contract
	BurnTokensTransactor // Write-only binding to the contract
	BurnTokensFilterer   // Log filterer for contract events
}

// BurnTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnTokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnTokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnTokensSession struct {
	Contract     *BurnTokens       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnTokensCallerSession struct {
	Contract *BurnTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BurnTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnTokensTransactorSession struct {
	Contract     *BurnTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BurnTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnTokensRaw struct {
	Contract *BurnTokens // Generic contract binding to access the raw methods on
}

// BurnTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnTokensCallerRaw struct {
	Contract *BurnTokensCaller // Generic read-only contract binding to access the raw methods on
}

// BurnTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnTokensTransactorRaw struct {
	Contract *BurnTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnTokens creates a new instance of BurnTokens, bound to a specific deployed contract.
func NewBurnTokens(address common.Address, backend bind.ContractBackend) (*BurnTokens, error) {
	contract, err := bindBurnTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnTokens{BurnTokensCaller: BurnTokensCaller{contract: contract}, BurnTokensTransactor: BurnTokensTransactor{contract: contract}, BurnTokensFilterer: BurnTokensFilterer{contract: contract}}, nil
}

// NewBurnTokensCaller creates a new read-only instance of BurnTokens, bound to a specific deployed contract.
func NewBurnTokensCaller(address common.Address, caller bind.ContractCaller) (*BurnTokensCaller, error) {
	contract, err := bindBurnTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnTokensCaller{contract: contract}, nil
}

// NewBurnTokensTransactor creates a new write-only instance of BurnTokens, bound to a specific deployed contract.
func NewBurnTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnTokensTransactor, error) {
	contract, err := bindBurnTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnTokensTransactor{contract: contract}, nil
}

// NewBurnTokensFilterer creates a new log filterer instance of BurnTokens, bound to a specific deployed contract.
func NewBurnTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnTokensFilterer, error) {
	contract, err := bindBurnTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnTokensFilterer{contract: contract}, nil
}

// bindBurnTokens binds a generic wrapper to an already deployed contract.
func bindBurnTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnTokensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnTokens *BurnTokensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnTokens.Contract.BurnTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnTokens *BurnTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnTokens.Contract.BurnTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnTokens *BurnTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnTokens.Contract.BurnTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnTokens *BurnTokensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnTokens *BurnTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnTokens *BurnTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnTokens.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns()
func (_BurnTokens *BurnTokensTransactor) Execute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnTokens.contract.Transact(opts, "execute")
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns()
func (_BurnTokens *BurnTokensSession) Execute() (*types.Transaction, error) {
	return _BurnTokens.Contract.Execute(&_BurnTokens.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns()
func (_BurnTokens *BurnTokensTransactorSession) Execute() (*types.Transaction, error) {
	return _BurnTokens.Contract.Execute(&_BurnTokens.TransactOpts)
}
