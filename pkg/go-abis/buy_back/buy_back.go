// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package buy_back

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

// BuyBackMetaData contains all meta data concerning the BuyBack contract.
var BuyBackMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"execute\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// BuyBackABI is the input ABI used to generate the binding from.
// Deprecated: Use BuyBackMetaData.ABI instead.
var BuyBackABI = BuyBackMetaData.ABI

// BuyBack is an auto generated Go binding around an Ethereum contract.
type BuyBack struct {
	BuyBackCaller     // Read-only binding to the contract
	BuyBackTransactor // Write-only binding to the contract
	BuyBackFilterer   // Log filterer for contract events
}

// BuyBackCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyBackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyBackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyBackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyBackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuyBackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyBackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuyBackSession struct {
	Contract     *BuyBack          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyBackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyBackCallerSession struct {
	Contract *BuyBackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BuyBackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyBackTransactorSession struct {
	Contract     *BuyBackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BuyBackRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyBackRaw struct {
	Contract *BuyBack // Generic contract binding to access the raw methods on
}

// BuyBackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyBackCallerRaw struct {
	Contract *BuyBackCaller // Generic read-only contract binding to access the raw methods on
}

// BuyBackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyBackTransactorRaw struct {
	Contract *BuyBackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuyBack creates a new instance of BuyBack, bound to a specific deployed contract.
func NewBuyBack(address common.Address, backend bind.ContractBackend) (*BuyBack, error) {
	contract, err := bindBuyBack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BuyBack{BuyBackCaller: BuyBackCaller{contract: contract}, BuyBackTransactor: BuyBackTransactor{contract: contract}, BuyBackFilterer: BuyBackFilterer{contract: contract}}, nil
}

// NewBuyBackCaller creates a new read-only instance of BuyBack, bound to a specific deployed contract.
func NewBuyBackCaller(address common.Address, caller bind.ContractCaller) (*BuyBackCaller, error) {
	contract, err := bindBuyBack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuyBackCaller{contract: contract}, nil
}

// NewBuyBackTransactor creates a new write-only instance of BuyBack, bound to a specific deployed contract.
func NewBuyBackTransactor(address common.Address, transactor bind.ContractTransactor) (*BuyBackTransactor, error) {
	contract, err := bindBuyBack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuyBackTransactor{contract: contract}, nil
}

// NewBuyBackFilterer creates a new log filterer instance of BuyBack, bound to a specific deployed contract.
func NewBuyBackFilterer(address common.Address, filterer bind.ContractFilterer) (*BuyBackFilterer, error) {
	contract, err := bindBuyBack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuyBackFilterer{contract: contract}, nil
}

// bindBuyBack binds a generic wrapper to an already deployed contract.
func bindBuyBack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BuyBackMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyBack *BuyBackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuyBack.Contract.BuyBackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyBack *BuyBackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBack.Contract.BuyBackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyBack *BuyBackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyBack.Contract.BuyBackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyBack *BuyBackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuyBack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyBack *BuyBackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyBack *BuyBackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyBack.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns()
func (_BuyBack *BuyBackTransactor) Execute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyBack.contract.Transact(opts, "execute")
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns()
func (_BuyBack *BuyBackSession) Execute() (*types.Transaction, error) {
	return _BuyBack.Contract.Execute(&_BuyBack.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns()
func (_BuyBack *BuyBackTransactorSession) Execute() (*types.Transaction, error) {
	return _BuyBack.Contract.Execute(&_BuyBack.TransactOpts)
}
