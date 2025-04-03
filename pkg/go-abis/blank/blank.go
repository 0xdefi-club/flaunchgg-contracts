// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blank

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

// BlankMetaData contains all meta data concerning the Blank contract.
var BlankMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"execute\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// BlankABI is the input ABI used to generate the binding from.
// Deprecated: Use BlankMetaData.ABI instead.
var BlankABI = BlankMetaData.ABI

// Blank is an auto generated Go binding around an Ethereum contract.
type Blank struct {
	BlankCaller     // Read-only binding to the contract
	BlankTransactor // Write-only binding to the contract
	BlankFilterer   // Log filterer for contract events
}

// BlankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlankSession struct {
	Contract     *Blank            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlankCallerSession struct {
	Contract *BlankCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BlankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlankTransactorSession struct {
	Contract     *BlankTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlankRaw struct {
	Contract *Blank // Generic contract binding to access the raw methods on
}

// BlankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlankCallerRaw struct {
	Contract *BlankCaller // Generic read-only contract binding to access the raw methods on
}

// BlankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlankTransactorRaw struct {
	Contract *BlankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlank creates a new instance of Blank, bound to a specific deployed contract.
func NewBlank(address common.Address, backend bind.ContractBackend) (*Blank, error) {
	contract, err := bindBlank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blank{BlankCaller: BlankCaller{contract: contract}, BlankTransactor: BlankTransactor{contract: contract}, BlankFilterer: BlankFilterer{contract: contract}}, nil
}

// NewBlankCaller creates a new read-only instance of Blank, bound to a specific deployed contract.
func NewBlankCaller(address common.Address, caller bind.ContractCaller) (*BlankCaller, error) {
	contract, err := bindBlank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlankCaller{contract: contract}, nil
}

// NewBlankTransactor creates a new write-only instance of Blank, bound to a specific deployed contract.
func NewBlankTransactor(address common.Address, transactor bind.ContractTransactor) (*BlankTransactor, error) {
	contract, err := bindBlank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlankTransactor{contract: contract}, nil
}

// NewBlankFilterer creates a new log filterer instance of Blank, bound to a specific deployed contract.
func NewBlankFilterer(address common.Address, filterer bind.ContractFilterer) (*BlankFilterer, error) {
	contract, err := bindBlank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlankFilterer{contract: contract}, nil
}

// bindBlank binds a generic wrapper to an already deployed contract.
func bindBlank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blank *BlankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blank.Contract.BlankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blank *BlankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blank.Contract.BlankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blank *BlankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blank.Contract.BlankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blank *BlankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blank *BlankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blank *BlankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blank.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns()
func (_Blank *BlankTransactor) Execute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blank.contract.Transact(opts, "execute")
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns()
func (_Blank *BlankSession) Execute() (*types.Transaction, error) {
	return _Blank.Contract.Execute(&_Blank.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns()
func (_Blank *BlankTransactorSession) Execute() (*types.Transaction, error) {
	return _Blank.Contract.Execute(&_Blank.TransactOpts)
}
