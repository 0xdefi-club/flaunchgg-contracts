// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pause_calculator

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

// PauseCalculatorMetaData contains all meta data concerning the PauseCalculator contract.
var PauseCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"determineSwapFee\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"setFlaunchParams\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"trackSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b506103a88061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c806324f9617714610043578063403910851461005b578063edf2ec7914610086575b5f5ffd5b610059610051366004610127565b505050505050565b005b61006e61006936600461027e565b610099565b60405162ffffff909116815260200160405180910390f35b61005961009436600461032b565b6100b3565b5f60405163d93c066560e01b815260040160405180910390fd5b60405163d93c066560e01b815260040160405180910390fd5b6001600160a01b03811681146100e0575f5ffd5b50565b5f5f83601f8401126100f3575f5ffd5b5081356001600160401b03811115610109575f5ffd5b602083019150836020828501011115610120575f5ffd5b9250929050565b5f5f5f5f5f5f86880361016081121561013e575f5ffd5b8735610149816100cc565b965060a0601f198201121561015c575f5ffd5b602088019550606060bf1982011215610173575f5ffd5b5060c08701935061012087013592506101408701356001600160401b0381111561019b575f5ffd5b6101a789828a016100e3565b979a9699509497509295939492505050565b60405160a081016001600160401b03811182821017156101e757634e487b7160e01b5f52604160045260245ffd5b60405290565b803562ffffff811681146101ff575f5ffd5b919050565b5f60608284031215610214575f5ffd5b604051606081016001600160401b038111828210171561024257634e487b7160e01b5f52604160045260245ffd5b60405290508082358015158114610257575f5ffd5b8152602083810135908201526040830135610271816100cc565b6040919091015292915050565b5f5f5f838503610120811215610292575f5ffd5b60a081121561029f575f5ffd5b506102a86101b9565b84356102b3816100cc565b815260208501356102c3816100cc565b60208201526102d4604086016101ed565b604082015260608501358060020b81146102ec575f5ffd5b606082015260808501356102ff816100cc565b608082015292506103138560a08601610204565b915061032261010085016101ed565b90509250925092565b5f5f5f6040848603121561033d575f5ffd5b8335925060208401356001600160401b03811115610359575f5ffd5b610365868287016100e3565b949790965093945050505056fea26469706673582212209000b44e0252c4d69e3e6b62e814ad7405c5cd537958527127b8f206c289e81464736f6c634300081b0033",
}

// PauseCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use PauseCalculatorMetaData.ABI instead.
var PauseCalculatorABI = PauseCalculatorMetaData.ABI

// PauseCalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PauseCalculatorMetaData.Bin instead.
var PauseCalculatorBin = PauseCalculatorMetaData.Bin

// DeployPauseCalculator deploys a new Ethereum contract, binding an instance of PauseCalculator to it.
func DeployPauseCalculator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PauseCalculator, error) {
	parsed, err := PauseCalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PauseCalculatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PauseCalculator{PauseCalculatorCaller: PauseCalculatorCaller{contract: contract}, PauseCalculatorTransactor: PauseCalculatorTransactor{contract: contract}, PauseCalculatorFilterer: PauseCalculatorFilterer{contract: contract}}, nil
}

// PauseCalculator is an auto generated Go binding around an Ethereum contract.
type PauseCalculator struct {
	PauseCalculatorCaller     // Read-only binding to the contract
	PauseCalculatorTransactor // Write-only binding to the contract
	PauseCalculatorFilterer   // Log filterer for contract events
}

// PauseCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PauseCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauseCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PauseCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauseCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PauseCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauseCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PauseCalculatorSession struct {
	Contract     *PauseCalculator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PauseCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PauseCalculatorCallerSession struct {
	Contract *PauseCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PauseCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PauseCalculatorTransactorSession struct {
	Contract     *PauseCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PauseCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PauseCalculatorRaw struct {
	Contract *PauseCalculator // Generic contract binding to access the raw methods on
}

// PauseCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PauseCalculatorCallerRaw struct {
	Contract *PauseCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// PauseCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PauseCalculatorTransactorRaw struct {
	Contract *PauseCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPauseCalculator creates a new instance of PauseCalculator, bound to a specific deployed contract.
func NewPauseCalculator(address common.Address, backend bind.ContractBackend) (*PauseCalculator, error) {
	contract, err := bindPauseCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PauseCalculator{PauseCalculatorCaller: PauseCalculatorCaller{contract: contract}, PauseCalculatorTransactor: PauseCalculatorTransactor{contract: contract}, PauseCalculatorFilterer: PauseCalculatorFilterer{contract: contract}}, nil
}

// NewPauseCalculatorCaller creates a new read-only instance of PauseCalculator, bound to a specific deployed contract.
func NewPauseCalculatorCaller(address common.Address, caller bind.ContractCaller) (*PauseCalculatorCaller, error) {
	contract, err := bindPauseCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PauseCalculatorCaller{contract: contract}, nil
}

// NewPauseCalculatorTransactor creates a new write-only instance of PauseCalculator, bound to a specific deployed contract.
func NewPauseCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*PauseCalculatorTransactor, error) {
	contract, err := bindPauseCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PauseCalculatorTransactor{contract: contract}, nil
}

// NewPauseCalculatorFilterer creates a new log filterer instance of PauseCalculator, bound to a specific deployed contract.
func NewPauseCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*PauseCalculatorFilterer, error) {
	contract, err := bindPauseCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PauseCalculatorFilterer{contract: contract}, nil
}

// bindPauseCalculator binds a generic wrapper to an already deployed contract.
func bindPauseCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PauseCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauseCalculator *PauseCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauseCalculator.Contract.PauseCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauseCalculator *PauseCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseCalculator.Contract.PauseCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauseCalculator *PauseCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauseCalculator.Contract.PauseCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauseCalculator *PauseCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauseCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauseCalculator *PauseCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauseCalculator *PauseCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauseCalculator.Contract.contract.Transact(opts, method, params...)
}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) , (bool,int256,uint160) , uint24 ) pure returns(uint24)
func (_PauseCalculator *PauseCalculatorCaller) DetermineSwapFee(opts *bind.CallOpts, arg0 PoolKey, arg1 IPoolManagerSwapParams, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PauseCalculator.contract.Call(opts, &out, "determineSwapFee", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) , (bool,int256,uint160) , uint24 ) pure returns(uint24)
func (_PauseCalculator *PauseCalculatorSession) DetermineSwapFee(arg0 PoolKey, arg1 IPoolManagerSwapParams, arg2 *big.Int) (*big.Int, error) {
	return _PauseCalculator.Contract.DetermineSwapFee(&_PauseCalculator.CallOpts, arg0, arg1, arg2)
}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) , (bool,int256,uint160) , uint24 ) pure returns(uint24)
func (_PauseCalculator *PauseCalculatorCallerSession) DetermineSwapFee(arg0 PoolKey, arg1 IPoolManagerSwapParams, arg2 *big.Int) (*big.Int, error) {
	return _PauseCalculator.Contract.DetermineSwapFee(&_PauseCalculator.CallOpts, arg0, arg1, arg2)
}

// SetFlaunchParams is a free data retrieval call binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 , bytes ) pure returns()
func (_PauseCalculator *PauseCalculatorCaller) SetFlaunchParams(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte) error {
	var out []interface{}
	err := _PauseCalculator.contract.Call(opts, &out, "setFlaunchParams", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetFlaunchParams is a free data retrieval call binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 , bytes ) pure returns()
func (_PauseCalculator *PauseCalculatorSession) SetFlaunchParams(arg0 [32]byte, arg1 []byte) error {
	return _PauseCalculator.Contract.SetFlaunchParams(&_PauseCalculator.CallOpts, arg0, arg1)
}

// SetFlaunchParams is a free data retrieval call binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 , bytes ) pure returns()
func (_PauseCalculator *PauseCalculatorCallerSession) SetFlaunchParams(arg0 [32]byte, arg1 []byte) error {
	return _PauseCalculator.Contract.SetFlaunchParams(&_PauseCalculator.CallOpts, arg0, arg1)
}

// TrackSwap is a free data retrieval call binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) pure returns()
func (_PauseCalculator *PauseCalculatorCaller) TrackSwap(opts *bind.CallOpts, arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) error {
	var out []interface{}
	err := _PauseCalculator.contract.Call(opts, &out, "trackSwap", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return err
	}

	return err

}

// TrackSwap is a free data retrieval call binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) pure returns()
func (_PauseCalculator *PauseCalculatorSession) TrackSwap(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) error {
	return _PauseCalculator.Contract.TrackSwap(&_PauseCalculator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// TrackSwap is a free data retrieval call binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) , (bool,int256,uint160) , int256 , bytes ) pure returns()
func (_PauseCalculator *PauseCalculatorCallerSession) TrackSwap(arg0 common.Address, arg1 PoolKey, arg2 IPoolManagerSwapParams, arg3 *big.Int, arg4 []byte) error {
	return _PauseCalculator.Contract.TrackSwap(&_PauseCalculator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}
