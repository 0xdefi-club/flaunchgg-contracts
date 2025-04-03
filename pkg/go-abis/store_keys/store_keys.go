// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store_keys

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

// StoreKeysMetaData contains all meta data concerning the StoreKeys contract.
var StoreKeysMetaData = &bind.MetaData{
	ABI: "[]",
}

// StoreKeysABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreKeysMetaData.ABI instead.
var StoreKeysABI = StoreKeysMetaData.ABI

// StoreKeys is an auto generated Go binding around an Ethereum contract.
type StoreKeys struct {
	StoreKeysCaller     // Read-only binding to the contract
	StoreKeysTransactor // Write-only binding to the contract
	StoreKeysFilterer   // Log filterer for contract events
}

// StoreKeysCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreKeysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreKeysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreKeysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreKeysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreKeysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreKeysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreKeysSession struct {
	Contract     *StoreKeys        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreKeysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreKeysCallerSession struct {
	Contract *StoreKeysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StoreKeysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreKeysTransactorSession struct {
	Contract     *StoreKeysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StoreKeysRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreKeysRaw struct {
	Contract *StoreKeys // Generic contract binding to access the raw methods on
}

// StoreKeysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreKeysCallerRaw struct {
	Contract *StoreKeysCaller // Generic read-only contract binding to access the raw methods on
}

// StoreKeysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreKeysTransactorRaw struct {
	Contract *StoreKeysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStoreKeys creates a new instance of StoreKeys, bound to a specific deployed contract.
func NewStoreKeys(address common.Address, backend bind.ContractBackend) (*StoreKeys, error) {
	contract, err := bindStoreKeys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StoreKeys{StoreKeysCaller: StoreKeysCaller{contract: contract}, StoreKeysTransactor: StoreKeysTransactor{contract: contract}, StoreKeysFilterer: StoreKeysFilterer{contract: contract}}, nil
}

// NewStoreKeysCaller creates a new read-only instance of StoreKeys, bound to a specific deployed contract.
func NewStoreKeysCaller(address common.Address, caller bind.ContractCaller) (*StoreKeysCaller, error) {
	contract, err := bindStoreKeys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreKeysCaller{contract: contract}, nil
}

// NewStoreKeysTransactor creates a new write-only instance of StoreKeys, bound to a specific deployed contract.
func NewStoreKeysTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreKeysTransactor, error) {
	contract, err := bindStoreKeys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreKeysTransactor{contract: contract}, nil
}

// NewStoreKeysFilterer creates a new log filterer instance of StoreKeys, bound to a specific deployed contract.
func NewStoreKeysFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreKeysFilterer, error) {
	contract, err := bindStoreKeys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreKeysFilterer{contract: contract}, nil
}

// bindStoreKeys binds a generic wrapper to an already deployed contract.
func bindStoreKeys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreKeysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoreKeys *StoreKeysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StoreKeys.Contract.StoreKeysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoreKeys *StoreKeysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoreKeys.Contract.StoreKeysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoreKeys *StoreKeysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoreKeys.Contract.StoreKeysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoreKeys *StoreKeysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StoreKeys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoreKeys *StoreKeysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoreKeys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoreKeys *StoreKeysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoreKeys.Contract.contract.Transact(opts, method, params...)
}
