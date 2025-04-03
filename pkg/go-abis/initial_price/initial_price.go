// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package initial_price

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

// InitialPriceInitialSqrtPriceX96 is an auto generated low-level Go binding around an user-defined struct.
type InitialPriceInitialSqrtPriceX96 struct {
	Unflipped *big.Int
	Flipped   *big.Int
}

// InitialPriceMetaData contains all meta data concerning the InitialPrice contract.
var InitialPriceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_flaunchFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_protocolOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getFlaunchingFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketCap\",\"inputs\":[{\"name\":\"_initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSqrtPriceX96\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_flipped\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setSqrtPriceX96\",\"inputs\":[{\"name\":\"_sqrtPriceX96\",\"type\":\"tuple\",\"internalType\":\"structInitialPrice.InitialSqrtPriceX96\",\"components\":[{\"name\":\"unflipped\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"flipped\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"InitialSqrtPriceX96Updated\",\"inputs\":[{\"name\":\"_unflipped\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"_flipped\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161083038038061083083398101604081905261002e9161009b565b608082905261003c81610043565b50506100d5565b638b78c6d81980541561005d57630dc149f05f526004601cfd5b6001600160a01b03909116801560ff1b8117909155805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b5f5f604083850312156100ac575f5ffd5b825160208401519092506001600160a01b03811681146100ca575f5ffd5b809150509250929050565b6080516107436100ed5f395f6102a601526107435ff3fe60806040526004361061008b575f3560e01c8063256929621461008f5780632a81c4f81461009957806342d8c2c0146100cb57806354d1f13d146100ea578063715018a6146100f25780638da5cb5b146100fa578063bf41930714610126578063f04e283e14610145578063f2fde38b14610158578063fc87f3f61461016b578063fee81cf41461018a575b5f5ffd5b6100976101bb565b005b3480156100a4575f5ffd5b506100b86100b336600461055f565b610207565b6040519081526020015b60405180910390f35b3480156100d6575f5ffd5b506100b86100e53660046105b1565b6102a4565b6100976102cd565b610097610306565b348015610105575f5ffd5b50638b78c6d819545b6040516001600160a01b0390911681526020016100c2565b348015610131575f5ffd5b50610097610140366004610601565b610319565b61009761015336600461066b565b610390565b61009761016636600461066b565b6103cd565b348015610176575f5ffd5b5061010e610185366004610686565b6103f3565b348015610195575f5ffd5b506100b86101a436600461066b565b63389a75e1600c9081525f91909152602090205490565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b5f5f610215335f86866103f3565b90506001600160801b036001600160a01b03821611610268575f6102426001600160a01b038316806106ea565b905061025f600160c01b680a18f07d736b90be55601d1b8361041f565b9250505061029e565b5f6102816001600160a01b03831680600160401b61041f565b905061025f600160801b680a18f07d736b90be55601d1b8361041f565b92915050565b7f00000000000000000000000000000000000000000000000000000000000000005b9392505050565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b61030e6104bb565b6103175f6104d5565b565b6103216104bb565b80515f80546001600160a01b039283166001600160a01b0319918216811790925560208085015160018054919095169216821790935560408051928352928201527f0cad32396951ca6426584f761c48652d83ef223b6ff5cfa713435bd2adbbc356910160405180910390a150565b6103986104bb565b63389a75e1600c52805f526020600c2080544211156103be57636f5e88185f526004601cfd5b5f90556103ca816104d5565b50565b6103d56104bb565b8060601b6103ea57637448fbae5f526004601cfd5b6103ca816104d5565b5f83610409575f546001600160a01b0316610416565b6001546001600160a01b03165b95945050505050565b5f838302815f198587098281108382030391505080841161043e575f5ffd5b805f03610450575082900490506102c6565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b638b78c6d819543314610317576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3811560ff1b8217905550565b5f5f83601f84011261052b575f5ffd5b5081356001600160401b03811115610541575f5ffd5b602083019150836020828501011115610558575f5ffd5b9250929050565b5f5f60208385031215610570575f5ffd5b82356001600160401b03811115610585575f5ffd5b6105918582860161051b565b90969095509350505050565b6001600160a01b03811681146103ca575f5ffd5b5f5f5f604084860312156105c3575f5ffd5b83356105ce8161059d565b925060208401356001600160401b038111156105e8575f5ffd5b6105f48682870161051b565b9497909650939450505050565b5f6040828403128015610612575f5ffd5b50604080519081016001600160401b038111828210171561064157634e487b7160e01b5f52604160045260245ffd5b604052823561064f8161059d565b8152602083013561065f8161059d565b60208201529392505050565b5f6020828403121561067b575f5ffd5b81356102c68161059d565b5f5f5f5f60608587031215610699575f5ffd5b84356106a48161059d565b9350602085013580151581146106b8575f5ffd5b925060408501356001600160401b038111156106d2575f5ffd5b6106de8782880161051b565b95989497509550505050565b808202811582820484141761029e57634e487b7160e01b5f52601160045260245ffdfea264697066735822122085facb22a949d5527077b91bc67da18364113f6f297234dad8ec29ff88ab45c464736f6c634300081b0033",
}

// InitialPriceABI is the input ABI used to generate the binding from.
// Deprecated: Use InitialPriceMetaData.ABI instead.
var InitialPriceABI = InitialPriceMetaData.ABI

// InitialPriceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InitialPriceMetaData.Bin instead.
var InitialPriceBin = InitialPriceMetaData.Bin

// DeployInitialPrice deploys a new Ethereum contract, binding an instance of InitialPrice to it.
func DeployInitialPrice(auth *bind.TransactOpts, backend bind.ContractBackend, _flaunchFee *big.Int, _protocolOwner common.Address) (common.Address, *types.Transaction, *InitialPrice, error) {
	parsed, err := InitialPriceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InitialPriceBin), backend, _flaunchFee, _protocolOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InitialPrice{InitialPriceCaller: InitialPriceCaller{contract: contract}, InitialPriceTransactor: InitialPriceTransactor{contract: contract}, InitialPriceFilterer: InitialPriceFilterer{contract: contract}}, nil
}

// InitialPrice is an auto generated Go binding around an Ethereum contract.
type InitialPrice struct {
	InitialPriceCaller     // Read-only binding to the contract
	InitialPriceTransactor // Write-only binding to the contract
	InitialPriceFilterer   // Log filterer for contract events
}

// InitialPriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitialPriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitialPriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitialPriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitialPriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitialPriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitialPriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitialPriceSession struct {
	Contract     *InitialPrice     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitialPriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitialPriceCallerSession struct {
	Contract *InitialPriceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// InitialPriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitialPriceTransactorSession struct {
	Contract     *InitialPriceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InitialPriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitialPriceRaw struct {
	Contract *InitialPrice // Generic contract binding to access the raw methods on
}

// InitialPriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitialPriceCallerRaw struct {
	Contract *InitialPriceCaller // Generic read-only contract binding to access the raw methods on
}

// InitialPriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitialPriceTransactorRaw struct {
	Contract *InitialPriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitialPrice creates a new instance of InitialPrice, bound to a specific deployed contract.
func NewInitialPrice(address common.Address, backend bind.ContractBackend) (*InitialPrice, error) {
	contract, err := bindInitialPrice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InitialPrice{InitialPriceCaller: InitialPriceCaller{contract: contract}, InitialPriceTransactor: InitialPriceTransactor{contract: contract}, InitialPriceFilterer: InitialPriceFilterer{contract: contract}}, nil
}

// NewInitialPriceCaller creates a new read-only instance of InitialPrice, bound to a specific deployed contract.
func NewInitialPriceCaller(address common.Address, caller bind.ContractCaller) (*InitialPriceCaller, error) {
	contract, err := bindInitialPrice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitialPriceCaller{contract: contract}, nil
}

// NewInitialPriceTransactor creates a new write-only instance of InitialPrice, bound to a specific deployed contract.
func NewInitialPriceTransactor(address common.Address, transactor bind.ContractTransactor) (*InitialPriceTransactor, error) {
	contract, err := bindInitialPrice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitialPriceTransactor{contract: contract}, nil
}

// NewInitialPriceFilterer creates a new log filterer instance of InitialPrice, bound to a specific deployed contract.
func NewInitialPriceFilterer(address common.Address, filterer bind.ContractFilterer) (*InitialPriceFilterer, error) {
	contract, err := bindInitialPrice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitialPriceFilterer{contract: contract}, nil
}

// bindInitialPrice binds a generic wrapper to an already deployed contract.
func bindInitialPrice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InitialPriceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitialPrice *InitialPriceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InitialPrice.Contract.InitialPriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitialPrice *InitialPriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitialPrice.Contract.InitialPriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitialPrice *InitialPriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitialPrice.Contract.InitialPriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitialPrice *InitialPriceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InitialPrice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitialPrice *InitialPriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitialPrice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitialPrice *InitialPriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitialPrice.Contract.contract.Transact(opts, method, params...)
}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0x42d8c2c0.
//
// Solidity: function getFlaunchingFee(address , bytes ) view returns(uint256)
func (_InitialPrice *InitialPriceCaller) GetFlaunchingFee(opts *bind.CallOpts, arg0 common.Address, arg1 []byte) (*big.Int, error) {
	var out []interface{}
	err := _InitialPrice.contract.Call(opts, &out, "getFlaunchingFee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0x42d8c2c0.
//
// Solidity: function getFlaunchingFee(address , bytes ) view returns(uint256)
func (_InitialPrice *InitialPriceSession) GetFlaunchingFee(arg0 common.Address, arg1 []byte) (*big.Int, error) {
	return _InitialPrice.Contract.GetFlaunchingFee(&_InitialPrice.CallOpts, arg0, arg1)
}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0x42d8c2c0.
//
// Solidity: function getFlaunchingFee(address , bytes ) view returns(uint256)
func (_InitialPrice *InitialPriceCallerSession) GetFlaunchingFee(arg0 common.Address, arg1 []byte) (*big.Int, error) {
	return _InitialPrice.Contract.GetFlaunchingFee(&_InitialPrice.CallOpts, arg0, arg1)
}

// GetMarketCap is a free data retrieval call binding the contract method 0x2a81c4f8.
//
// Solidity: function getMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_InitialPrice *InitialPriceCaller) GetMarketCap(opts *bind.CallOpts, _initialPriceParams []byte) (*big.Int, error) {
	var out []interface{}
	err := _InitialPrice.contract.Call(opts, &out, "getMarketCap", _initialPriceParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCap is a free data retrieval call binding the contract method 0x2a81c4f8.
//
// Solidity: function getMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_InitialPrice *InitialPriceSession) GetMarketCap(_initialPriceParams []byte) (*big.Int, error) {
	return _InitialPrice.Contract.GetMarketCap(&_InitialPrice.CallOpts, _initialPriceParams)
}

// GetMarketCap is a free data retrieval call binding the contract method 0x2a81c4f8.
//
// Solidity: function getMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_InitialPrice *InitialPriceCallerSession) GetMarketCap(_initialPriceParams []byte) (*big.Int, error) {
	return _InitialPrice.Contract.GetMarketCap(&_InitialPrice.CallOpts, _initialPriceParams)
}

// GetSqrtPriceX96 is a free data retrieval call binding the contract method 0xfc87f3f6.
//
// Solidity: function getSqrtPriceX96(address , bool _flipped, bytes ) view returns(uint160)
func (_InitialPrice *InitialPriceCaller) GetSqrtPriceX96(opts *bind.CallOpts, arg0 common.Address, _flipped bool, arg2 []byte) (*big.Int, error) {
	var out []interface{}
	err := _InitialPrice.contract.Call(opts, &out, "getSqrtPriceX96", arg0, _flipped, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSqrtPriceX96 is a free data retrieval call binding the contract method 0xfc87f3f6.
//
// Solidity: function getSqrtPriceX96(address , bool _flipped, bytes ) view returns(uint160)
func (_InitialPrice *InitialPriceSession) GetSqrtPriceX96(arg0 common.Address, _flipped bool, arg2 []byte) (*big.Int, error) {
	return _InitialPrice.Contract.GetSqrtPriceX96(&_InitialPrice.CallOpts, arg0, _flipped, arg2)
}

// GetSqrtPriceX96 is a free data retrieval call binding the contract method 0xfc87f3f6.
//
// Solidity: function getSqrtPriceX96(address , bool _flipped, bytes ) view returns(uint160)
func (_InitialPrice *InitialPriceCallerSession) GetSqrtPriceX96(arg0 common.Address, _flipped bool, arg2 []byte) (*big.Int, error) {
	return _InitialPrice.Contract.GetSqrtPriceX96(&_InitialPrice.CallOpts, arg0, _flipped, arg2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_InitialPrice *InitialPriceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InitialPrice.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_InitialPrice *InitialPriceSession) Owner() (common.Address, error) {
	return _InitialPrice.Contract.Owner(&_InitialPrice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_InitialPrice *InitialPriceCallerSession) Owner() (common.Address, error) {
	return _InitialPrice.Contract.Owner(&_InitialPrice.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_InitialPrice *InitialPriceCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InitialPrice.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_InitialPrice *InitialPriceSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _InitialPrice.Contract.OwnershipHandoverExpiresAt(&_InitialPrice.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_InitialPrice *InitialPriceCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _InitialPrice.Contract.OwnershipHandoverExpiresAt(&_InitialPrice.CallOpts, pendingOwner)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_InitialPrice *InitialPriceTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitialPrice.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_InitialPrice *InitialPriceSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _InitialPrice.Contract.CancelOwnershipHandover(&_InitialPrice.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_InitialPrice *InitialPriceTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _InitialPrice.Contract.CancelOwnershipHandover(&_InitialPrice.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_InitialPrice *InitialPriceTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _InitialPrice.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_InitialPrice *InitialPriceSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _InitialPrice.Contract.CompleteOwnershipHandover(&_InitialPrice.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_InitialPrice *InitialPriceTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _InitialPrice.Contract.CompleteOwnershipHandover(&_InitialPrice.TransactOpts, pendingOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_InitialPrice *InitialPriceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitialPrice.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_InitialPrice *InitialPriceSession) RenounceOwnership() (*types.Transaction, error) {
	return _InitialPrice.Contract.RenounceOwnership(&_InitialPrice.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_InitialPrice *InitialPriceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InitialPrice.Contract.RenounceOwnership(&_InitialPrice.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_InitialPrice *InitialPriceTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitialPrice.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_InitialPrice *InitialPriceSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _InitialPrice.Contract.RequestOwnershipHandover(&_InitialPrice.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_InitialPrice *InitialPriceTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _InitialPrice.Contract.RequestOwnershipHandover(&_InitialPrice.TransactOpts)
}

// SetSqrtPriceX96 is a paid mutator transaction binding the contract method 0xbf419307.
//
// Solidity: function setSqrtPriceX96((uint160,uint160) _sqrtPriceX96) returns()
func (_InitialPrice *InitialPriceTransactor) SetSqrtPriceX96(opts *bind.TransactOpts, _sqrtPriceX96 InitialPriceInitialSqrtPriceX96) (*types.Transaction, error) {
	return _InitialPrice.contract.Transact(opts, "setSqrtPriceX96", _sqrtPriceX96)
}

// SetSqrtPriceX96 is a paid mutator transaction binding the contract method 0xbf419307.
//
// Solidity: function setSqrtPriceX96((uint160,uint160) _sqrtPriceX96) returns()
func (_InitialPrice *InitialPriceSession) SetSqrtPriceX96(_sqrtPriceX96 InitialPriceInitialSqrtPriceX96) (*types.Transaction, error) {
	return _InitialPrice.Contract.SetSqrtPriceX96(&_InitialPrice.TransactOpts, _sqrtPriceX96)
}

// SetSqrtPriceX96 is a paid mutator transaction binding the contract method 0xbf419307.
//
// Solidity: function setSqrtPriceX96((uint160,uint160) _sqrtPriceX96) returns()
func (_InitialPrice *InitialPriceTransactorSession) SetSqrtPriceX96(_sqrtPriceX96 InitialPriceInitialSqrtPriceX96) (*types.Transaction, error) {
	return _InitialPrice.Contract.SetSqrtPriceX96(&_InitialPrice.TransactOpts, _sqrtPriceX96)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_InitialPrice *InitialPriceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InitialPrice.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_InitialPrice *InitialPriceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InitialPrice.Contract.TransferOwnership(&_InitialPrice.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_InitialPrice *InitialPriceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InitialPrice.Contract.TransferOwnership(&_InitialPrice.TransactOpts, newOwner)
}

// InitialPriceInitialSqrtPriceX96UpdatedIterator is returned from FilterInitialSqrtPriceX96Updated and is used to iterate over the raw logs and unpacked data for InitialSqrtPriceX96Updated events raised by the InitialPrice contract.
type InitialPriceInitialSqrtPriceX96UpdatedIterator struct {
	Event *InitialPriceInitialSqrtPriceX96Updated // Event containing the contract specifics and raw log

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
func (it *InitialPriceInitialSqrtPriceX96UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitialPriceInitialSqrtPriceX96Updated)
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
		it.Event = new(InitialPriceInitialSqrtPriceX96Updated)
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
func (it *InitialPriceInitialSqrtPriceX96UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitialPriceInitialSqrtPriceX96UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitialPriceInitialSqrtPriceX96Updated represents a InitialSqrtPriceX96Updated event raised by the InitialPrice contract.
type InitialPriceInitialSqrtPriceX96Updated struct {
	Unflipped *big.Int
	Flipped   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitialSqrtPriceX96Updated is a free log retrieval operation binding the contract event 0x0cad32396951ca6426584f761c48652d83ef223b6ff5cfa713435bd2adbbc356.
//
// Solidity: event InitialSqrtPriceX96Updated(uint160 _unflipped, uint160 _flipped)
func (_InitialPrice *InitialPriceFilterer) FilterInitialSqrtPriceX96Updated(opts *bind.FilterOpts) (*InitialPriceInitialSqrtPriceX96UpdatedIterator, error) {

	logs, sub, err := _InitialPrice.contract.FilterLogs(opts, "InitialSqrtPriceX96Updated")
	if err != nil {
		return nil, err
	}
	return &InitialPriceInitialSqrtPriceX96UpdatedIterator{contract: _InitialPrice.contract, event: "InitialSqrtPriceX96Updated", logs: logs, sub: sub}, nil
}

// WatchInitialSqrtPriceX96Updated is a free log subscription operation binding the contract event 0x0cad32396951ca6426584f761c48652d83ef223b6ff5cfa713435bd2adbbc356.
//
// Solidity: event InitialSqrtPriceX96Updated(uint160 _unflipped, uint160 _flipped)
func (_InitialPrice *InitialPriceFilterer) WatchInitialSqrtPriceX96Updated(opts *bind.WatchOpts, sink chan<- *InitialPriceInitialSqrtPriceX96Updated) (event.Subscription, error) {

	logs, sub, err := _InitialPrice.contract.WatchLogs(opts, "InitialSqrtPriceX96Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitialPriceInitialSqrtPriceX96Updated)
				if err := _InitialPrice.contract.UnpackLog(event, "InitialSqrtPriceX96Updated", log); err != nil {
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

// ParseInitialSqrtPriceX96Updated is a log parse operation binding the contract event 0x0cad32396951ca6426584f761c48652d83ef223b6ff5cfa713435bd2adbbc356.
//
// Solidity: event InitialSqrtPriceX96Updated(uint160 _unflipped, uint160 _flipped)
func (_InitialPrice *InitialPriceFilterer) ParseInitialSqrtPriceX96Updated(log types.Log) (*InitialPriceInitialSqrtPriceX96Updated, error) {
	event := new(InitialPriceInitialSqrtPriceX96Updated)
	if err := _InitialPrice.contract.UnpackLog(event, "InitialSqrtPriceX96Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitialPriceOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the InitialPrice contract.
type InitialPriceOwnershipHandoverCanceledIterator struct {
	Event *InitialPriceOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *InitialPriceOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitialPriceOwnershipHandoverCanceled)
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
		it.Event = new(InitialPriceOwnershipHandoverCanceled)
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
func (it *InitialPriceOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitialPriceOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitialPriceOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the InitialPrice contract.
type InitialPriceOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_InitialPrice *InitialPriceFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*InitialPriceOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _InitialPrice.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InitialPriceOwnershipHandoverCanceledIterator{contract: _InitialPrice.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_InitialPrice *InitialPriceFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *InitialPriceOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _InitialPrice.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitialPriceOwnershipHandoverCanceled)
				if err := _InitialPrice.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_InitialPrice *InitialPriceFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*InitialPriceOwnershipHandoverCanceled, error) {
	event := new(InitialPriceOwnershipHandoverCanceled)
	if err := _InitialPrice.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitialPriceOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the InitialPrice contract.
type InitialPriceOwnershipHandoverRequestedIterator struct {
	Event *InitialPriceOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *InitialPriceOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitialPriceOwnershipHandoverRequested)
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
		it.Event = new(InitialPriceOwnershipHandoverRequested)
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
func (it *InitialPriceOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitialPriceOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitialPriceOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the InitialPrice contract.
type InitialPriceOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_InitialPrice *InitialPriceFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*InitialPriceOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _InitialPrice.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InitialPriceOwnershipHandoverRequestedIterator{contract: _InitialPrice.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_InitialPrice *InitialPriceFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *InitialPriceOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _InitialPrice.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitialPriceOwnershipHandoverRequested)
				if err := _InitialPrice.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_InitialPrice *InitialPriceFilterer) ParseOwnershipHandoverRequested(log types.Log) (*InitialPriceOwnershipHandoverRequested, error) {
	event := new(InitialPriceOwnershipHandoverRequested)
	if err := _InitialPrice.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitialPriceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InitialPrice contract.
type InitialPriceOwnershipTransferredIterator struct {
	Event *InitialPriceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InitialPriceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitialPriceOwnershipTransferred)
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
		it.Event = new(InitialPriceOwnershipTransferred)
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
func (it *InitialPriceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitialPriceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitialPriceOwnershipTransferred represents a OwnershipTransferred event raised by the InitialPrice contract.
type InitialPriceOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_InitialPrice *InitialPriceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*InitialPriceOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InitialPrice.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InitialPriceOwnershipTransferredIterator{contract: _InitialPrice.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_InitialPrice *InitialPriceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InitialPriceOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InitialPrice.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitialPriceOwnershipTransferred)
				if err := _InitialPrice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_InitialPrice *InitialPriceFilterer) ParseOwnershipTransferred(log types.Log) (*InitialPriceOwnershipTransferred, error) {
	event := new(InitialPriceOwnershipTransferred)
	if err := _InitialPrice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
