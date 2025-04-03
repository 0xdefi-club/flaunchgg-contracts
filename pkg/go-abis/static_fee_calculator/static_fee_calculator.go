// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package static_fee_calculator

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

// IPoolManagerSwapParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerSwapParams struct {
	ZeroForOne        bool
	AmountSpecified   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// PoolKey is an auto generated low-level Go binding around an user-defined struct.
type PoolKey struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

// StaticFeeCalculatorMetaData contains all meta data concerning the StaticFeeCalculator contract.
var StaticFeeCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"determineSwapFee\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"_baseFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[{\"name\":\"swapFee_\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"setFlaunchParams\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trackSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506103768061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c806324f9617714610043578063403910851461005b578063edf2ec7914610087575b5f5ffd5b6100596100513660046100f5565b505050505050565b005b61006f61006936600461024c565b92915050565b60405162ffffff909116815260200160405180910390f35b6100596100953660046102f9565b505050565b6001600160a01b03811681146100ae575f5ffd5b50565b5f5f83601f8401126100c1575f5ffd5b5081356001600160401b038111156100d7575f5ffd5b6020830191508360208285010111156100ee575f5ffd5b9250929050565b5f5f5f5f5f5f86880361016081121561010c575f5ffd5b87356101178161009a565b965060a0601f198201121561012a575f5ffd5b602088019550606060bf1982011215610141575f5ffd5b5060c08701935061012087013592506101408701356001600160401b03811115610169575f5ffd5b61017589828a016100b1565b979a9699509497509295939492505050565b60405160a081016001600160401b03811182821017156101b557634e487b7160e01b5f52604160045260245ffd5b60405290565b803562ffffff811681146101cd575f5ffd5b919050565b5f606082840312156101e2575f5ffd5b604051606081016001600160401b038111828210171561021057634e487b7160e01b5f52604160045260245ffd5b60405290508082358015158114610225575f5ffd5b815260208381013590820152604083013561023f8161009a565b6040919091015292915050565b5f5f5f838503610120811215610260575f5ffd5b60a081121561026d575f5ffd5b50610276610187565b84356102818161009a565b815260208501356102918161009a565b60208201526102a2604086016101bb565b604082015260608501358060020b81146102ba575f5ffd5b606082015260808501356102cd8161009a565b608082015292506102e18560a086016101d2565b91506102f061010085016101bb565b90509250925092565b5f5f5f6040848603121561030b575f5ffd5b8335925060208401356001600160401b03811115610327575f5ffd5b610333868287016100b1565b949790965093945050505056fea2646970667358221220356c2b37955282cf127a9d9fd3585b0b7658eb6650dede52b9443f1f0e14558664736f6c634300081b0033",
}

// StaticFeeCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use StaticFeeCalculatorMetaData.ABI instead.
var StaticFeeCalculatorABI = StaticFeeCalculatorMetaData.ABI

// StaticFeeCalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StaticFeeCalculatorMetaData.Bin instead.
var StaticFeeCalculatorBin = StaticFeeCalculatorMetaData.Bin

// DeployStaticFeeCalculator deploys a new Ethereum contract, binding an instance of StaticFeeCalculator to it.
func DeployStaticFeeCalculator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StaticFeeCalculator, error) {
	parsed, err := StaticFeeCalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StaticFeeCalculatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StaticFeeCalculator{StaticFeeCalculatorCaller: StaticFeeCalculatorCaller{contract: contract}, StaticFeeCalculatorTransactor: StaticFeeCalculatorTransactor{contract: contract}, StaticFeeCalculatorFilterer: StaticFeeCalculatorFilterer{contract: contract}}, nil
}

// StaticFeeCalculator is an auto generated Go binding around an Ethereum contract.
type StaticFeeCalculator struct {
	StaticFeeCalculatorCaller     // Read-only binding to the contract
	StaticFeeCalculatorTransactor // Write-only binding to the contract
	StaticFeeCalculatorFilterer   // Log filterer for contract events
}

// StaticFeeCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaticFeeCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticFeeCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaticFeeCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticFeeCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaticFeeCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticFeeCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaticFeeCalculatorSession struct {
	Contract     *StaticFeeCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StaticFeeCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaticFeeCalculatorCallerSession struct {
	Contract *StaticFeeCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// StaticFeeCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaticFeeCalculatorTransactorSession struct {
	Contract     *StaticFeeCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// StaticFeeCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaticFeeCalculatorRaw struct {
	Contract *StaticFeeCalculator // Generic contract binding to access the raw methods on
}

// StaticFeeCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaticFeeCalculatorCallerRaw struct {
	Contract *StaticFeeCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// StaticFeeCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaticFeeCalculatorTransactorRaw struct {
	Contract *StaticFeeCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaticFeeCalculator creates a new instance of StaticFeeCalculator, bound to a specific deployed contract.
func NewStaticFeeCalculator(address common.Address, backend bind.ContractBackend) (*StaticFeeCalculator, error) {
	contract, err := bindStaticFeeCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaticFeeCalculator{StaticFeeCalculatorCaller: StaticFeeCalculatorCaller{contract: contract}, StaticFeeCalculatorTransactor: StaticFeeCalculatorTransactor{contract: contract}, StaticFeeCalculatorFilterer: StaticFeeCalculatorFilterer{contract: contract}}, nil
}

// NewStaticFeeCalculatorCaller creates a new read-only instance of StaticFeeCalculator, bound to a specific deployed contract.
func NewStaticFeeCalculatorCaller(address common.Address, caller bind.ContractCaller) (*StaticFeeCalculatorCaller, error) {
	contract, err := bindStaticFeeCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaticFeeCalculatorCaller{contract: contract}, nil
}

// NewStaticFeeCalculatorTransactor creates a new write-only instance of StaticFeeCalculator, bound to a specific deployed contract.
func NewStaticFeeCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*StaticFeeCalculatorTransactor, error) {
	contract, err := bindStaticFeeCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaticFeeCalculatorTransactor{contract: contract}, nil
}

// NewStaticFeeCalculatorFilterer creates a new log filterer instance of StaticFeeCalculator, bound to a specific deployed contract.
func NewStaticFeeCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*StaticFeeCalculatorFilterer, error) {
	contract, err := bindStaticFeeCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaticFeeCalculatorFilterer{contract: contract}, nil
}

// bindStaticFeeCalculator binds a generic wrapper to an already deployed contract.
func bindStaticFeeCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StaticFeeCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaticFeeCalculator *StaticFeeCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaticFeeCalculator.Contract.StaticFeeCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaticFeeCalculator *StaticFeeCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticFeeCalculator.Contract.StaticFeeCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaticFeeCalculator *StaticFeeCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaticFeeCalculator.Contract.StaticFeeCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaticFeeCalculator *StaticFeeCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaticFeeCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaticFeeCalculator *StaticFeeCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticFeeCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaticFeeCalculator *StaticFeeCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaticFeeCalculator.Contract.contract.Transact(opts, method, params...)
}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) , (bool,int256,uint160) , uint24 _baseFee) pure returns(uint24 swapFee_)
func (_StaticFeeCalculator *StaticFeeCalculatorCaller) DetermineSwapFee(opts *bind.CallOpts, arg0 PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaticFeeCalculator.contract.Call(opts, &out, "determineSwapFee", arg0, arg1, _baseFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) , (bool,int256,uint160) , uint24 _baseFee) pure returns(uint24 swapFee_)
func (_StaticFeeCalculator *StaticFeeCalculatorSession) DetermineSwapFee(arg0 PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	return _StaticFeeCalculator.Contract.DetermineSwapFee(&_StaticFeeCalculator.CallOpts, arg0, arg1, _baseFee)
}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) , (bool,int256,uint160) , uint24 _baseFee) pure returns(uint24 swapFee_)
func (_StaticFeeCalculator *StaticFeeCalculatorCallerSession) DetermineSwapFee(arg0 PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	return _StaticFeeCalculator.Contract.DetermineSwapFee(&_StaticFeeCalculator.CallOpts, arg0, arg1, _baseFee)
}

// TrackSwap is a free data retrieval call binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) view returns()
func (_StaticFeeCalculator *StaticFeeCalculatorCaller) TrackSwap(opts *bind.CallOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) error {
	var out []interface{}
	err := _StaticFeeCalculator.contract.Call(opts, &out, "trackSwap", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return err
	}

	return err

}

// TrackSwap is a free data retrieval call binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) view returns()
func (_StaticFeeCalculator *StaticFeeCalculatorSession) TrackSwap(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) error {
	return _StaticFeeCalculator.Contract.TrackSwap(&_StaticFeeCalculator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// TrackSwap is a free data retrieval call binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) view returns()
func (_StaticFeeCalculator *StaticFeeCalculatorCallerSession) TrackSwap(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) error {
	return _StaticFeeCalculator.Contract.TrackSwap(&_StaticFeeCalculator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_StaticFeeCalculator *StaticFeeCalculatorTransactor) SetFlaunchParams(opts *bind.TransactOpts, _poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _StaticFeeCalculator.contract.Transact(opts, "setFlaunchParams", _poolId, _params)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_StaticFeeCalculator *StaticFeeCalculatorSession) SetFlaunchParams(_poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _StaticFeeCalculator.Contract.SetFlaunchParams(&_StaticFeeCalculator.TransactOpts, _poolId, _params)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_StaticFeeCalculator *StaticFeeCalculatorTransactorSession) SetFlaunchParams(_poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _StaticFeeCalculator.Contract.SetFlaunchParams(&_StaticFeeCalculator.TransactOpts, _poolId, _params)
}
