// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fee_distributor

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

// FeeDistributorFeeDistribution is an auto generated low-level Go binding around an user-defined struct.
type FeeDistributorFeeDistribution struct {
	SwapFee  *big.Int
	Referrer *big.Int
	Protocol *big.Int
	Active   bool
}

// FeeDistributorMetaData contains all meta data concerning the FeeDistributor contract.
var FeeDistributorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MAX_PROTOCOL_ALLOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fairLaunchFeeCalculator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeCalculator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplit\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"bidWall_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creator_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"protocol_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flayGovernance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeCalculator\",\"inputs\":[{\"name\":\"_isFairLaunch\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolFeeDistribution\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"feeDistribution_\",\"type\":\"tuple\",\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nativeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"referralEscrow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractReferralEscrow\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setFairLaunchFeeCalculator\",\"inputs\":[{\"name\":\"_feeCalculator\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeCalculator\",\"inputs\":[{\"name\":\"_feeCalculator\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeDistribution\",\"inputs\":[{\"name\":\"_feeDistribution\",\"type\":\"tuple\",\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoolFeeDistribution\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_feeDistribution\",\"type\":\"tuple\",\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProtocolFeeDistribution\",\"inputs\":[{\"name\":\"_protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReferralEscrow\",\"inputs\":[{\"name\":\"_referralEscrow\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_unwrap\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CreatorFeeAllocationUpdated\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_allocation\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_payee\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FairLaunchFeeCalculatorUpdated\",\"inputs\":[{\"name\":\"_feeCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeCalculatorUpdated\",\"inputs\":[{\"name\":\"_feeCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeDistributionUpdated\",\"inputs\":[{\"name\":\"_feeDistribution\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeDistributionUpdated\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_feeDistribution\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReferralEscrowUpdated\",\"inputs\":[{\"name\":\"_referralEscrow\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReferrerFeePaid\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotCreator\",\"inputs\":[{\"name\":\"_caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProtocolFeeInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecipientZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferrerFeeInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SwapFeeInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// FeeDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeDistributorMetaData.ABI instead.
var FeeDistributorABI = FeeDistributorMetaData.ABI

// FeeDistributor is an auto generated Go binding around an Ethereum contract.
type FeeDistributor struct {
	FeeDistributorCaller     // Read-only binding to the contract
	FeeDistributorTransactor // Write-only binding to the contract
	FeeDistributorFilterer   // Log filterer for contract events
}

// FeeDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeDistributorSession struct {
	Contract     *FeeDistributor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeDistributorCallerSession struct {
	Contract *FeeDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FeeDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeDistributorTransactorSession struct {
	Contract     *FeeDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FeeDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeDistributorRaw struct {
	Contract *FeeDistributor // Generic contract binding to access the raw methods on
}

// FeeDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeDistributorCallerRaw struct {
	Contract *FeeDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// FeeDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeDistributorTransactorRaw struct {
	Contract *FeeDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeDistributor creates a new instance of FeeDistributor, bound to a specific deployed contract.
func NewFeeDistributor(address common.Address, backend bind.ContractBackend) (*FeeDistributor, error) {
	contract, err := bindFeeDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeDistributor{FeeDistributorCaller: FeeDistributorCaller{contract: contract}, FeeDistributorTransactor: FeeDistributorTransactor{contract: contract}, FeeDistributorFilterer: FeeDistributorFilterer{contract: contract}}, nil
}

// NewFeeDistributorCaller creates a new read-only instance of FeeDistributor, bound to a specific deployed contract.
func NewFeeDistributorCaller(address common.Address, caller bind.ContractCaller) (*FeeDistributorCaller, error) {
	contract, err := bindFeeDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorCaller{contract: contract}, nil
}

// NewFeeDistributorTransactor creates a new write-only instance of FeeDistributor, bound to a specific deployed contract.
func NewFeeDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeDistributorTransactor, error) {
	contract, err := bindFeeDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorTransactor{contract: contract}, nil
}

// NewFeeDistributorFilterer creates a new log filterer instance of FeeDistributor, bound to a specific deployed contract.
func NewFeeDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeDistributorFilterer, error) {
	contract, err := bindFeeDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorFilterer{contract: contract}, nil
}

// bindFeeDistributor binds a generic wrapper to an already deployed contract.
func bindFeeDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeDistributor *FeeDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeDistributor.Contract.FeeDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeDistributor *FeeDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.Contract.FeeDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeDistributor *FeeDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeDistributor.Contract.FeeDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeDistributor *FeeDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeDistributor *FeeDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeDistributor *FeeDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeDistributor.Contract.contract.Transact(opts, method, params...)
}

// MAXPROTOCOLALLOCATION is a free data retrieval call binding the contract method 0xc29c945b.
//
// Solidity: function MAX_PROTOCOL_ALLOCATION() view returns(uint24)
func (_FeeDistributor *FeeDistributorCaller) MAXPROTOCOLALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "MAX_PROTOCOL_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPROTOCOLALLOCATION is a free data retrieval call binding the contract method 0xc29c945b.
//
// Solidity: function MAX_PROTOCOL_ALLOCATION() view returns(uint24)
func (_FeeDistributor *FeeDistributorSession) MAXPROTOCOLALLOCATION() (*big.Int, error) {
	return _FeeDistributor.Contract.MAXPROTOCOLALLOCATION(&_FeeDistributor.CallOpts)
}

// MAXPROTOCOLALLOCATION is a free data retrieval call binding the contract method 0xc29c945b.
//
// Solidity: function MAX_PROTOCOL_ALLOCATION() view returns(uint24)
func (_FeeDistributor *FeeDistributorCallerSession) MAXPROTOCOLALLOCATION() (*big.Int, error) {
	return _FeeDistributor.Contract.MAXPROTOCOLALLOCATION(&_FeeDistributor.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address _recipient) view returns(uint256 _amount)
func (_FeeDistributor *FeeDistributorCaller) Balances(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "balances", _recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address _recipient) view returns(uint256 _amount)
func (_FeeDistributor *FeeDistributorSession) Balances(_recipient common.Address) (*big.Int, error) {
	return _FeeDistributor.Contract.Balances(&_FeeDistributor.CallOpts, _recipient)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address _recipient) view returns(uint256 _amount)
func (_FeeDistributor *FeeDistributorCallerSession) Balances(_recipient common.Address) (*big.Int, error) {
	return _FeeDistributor.Contract.Balances(&_FeeDistributor.CallOpts, _recipient)
}

// FairLaunchFeeCalculator is a free data retrieval call binding the contract method 0xa1dd2d91.
//
// Solidity: function fairLaunchFeeCalculator() view returns(address)
func (_FeeDistributor *FeeDistributorCaller) FairLaunchFeeCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "fairLaunchFeeCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FairLaunchFeeCalculator is a free data retrieval call binding the contract method 0xa1dd2d91.
//
// Solidity: function fairLaunchFeeCalculator() view returns(address)
func (_FeeDistributor *FeeDistributorSession) FairLaunchFeeCalculator() (common.Address, error) {
	return _FeeDistributor.Contract.FairLaunchFeeCalculator(&_FeeDistributor.CallOpts)
}

// FairLaunchFeeCalculator is a free data retrieval call binding the contract method 0xa1dd2d91.
//
// Solidity: function fairLaunchFeeCalculator() view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) FairLaunchFeeCalculator() (common.Address, error) {
	return _FeeDistributor.Contract.FairLaunchFeeCalculator(&_FeeDistributor.CallOpts)
}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_FeeDistributor *FeeDistributorCaller) FeeCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "feeCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_FeeDistributor *FeeDistributorSession) FeeCalculator() (common.Address, error) {
	return _FeeDistributor.Contract.FeeCalculator(&_FeeDistributor.CallOpts)
}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) FeeCalculator() (common.Address, error) {
	return _FeeDistributor.Contract.FeeCalculator(&_FeeDistributor.CallOpts)
}

// FeeSplit is a free data retrieval call binding the contract method 0x4875cbb8.
//
// Solidity: function feeSplit(bytes32 _poolId, uint256 _amount) view returns(uint256 bidWall_, uint256 creator_, uint256 protocol_)
func (_FeeDistributor *FeeDistributorCaller) FeeSplit(opts *bind.CallOpts, _poolId [32]byte, _amount *big.Int) (struct {
	BidWall  *big.Int
	Creator  *big.Int
	Protocol *big.Int
}, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "feeSplit", _poolId, _amount)

	outstruct := new(struct {
		BidWall  *big.Int
		Creator  *big.Int
		Protocol *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BidWall = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Protocol = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FeeSplit is a free data retrieval call binding the contract method 0x4875cbb8.
//
// Solidity: function feeSplit(bytes32 _poolId, uint256 _amount) view returns(uint256 bidWall_, uint256 creator_, uint256 protocol_)
func (_FeeDistributor *FeeDistributorSession) FeeSplit(_poolId [32]byte, _amount *big.Int) (struct {
	BidWall  *big.Int
	Creator  *big.Int
	Protocol *big.Int
}, error) {
	return _FeeDistributor.Contract.FeeSplit(&_FeeDistributor.CallOpts, _poolId, _amount)
}

// FeeSplit is a free data retrieval call binding the contract method 0x4875cbb8.
//
// Solidity: function feeSplit(bytes32 _poolId, uint256 _amount) view returns(uint256 bidWall_, uint256 creator_, uint256 protocol_)
func (_FeeDistributor *FeeDistributorCallerSession) FeeSplit(_poolId [32]byte, _amount *big.Int) (struct {
	BidWall  *big.Int
	Creator  *big.Int
	Protocol *big.Int
}, error) {
	return _FeeDistributor.Contract.FeeSplit(&_FeeDistributor.CallOpts, _poolId, _amount)
}

// FlayGovernance is a free data retrieval call binding the contract method 0xdddfecf5.
//
// Solidity: function flayGovernance() view returns(address)
func (_FeeDistributor *FeeDistributorCaller) FlayGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "flayGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlayGovernance is a free data retrieval call binding the contract method 0xdddfecf5.
//
// Solidity: function flayGovernance() view returns(address)
func (_FeeDistributor *FeeDistributorSession) FlayGovernance() (common.Address, error) {
	return _FeeDistributor.Contract.FlayGovernance(&_FeeDistributor.CallOpts)
}

// FlayGovernance is a free data retrieval call binding the contract method 0xdddfecf5.
//
// Solidity: function flayGovernance() view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) FlayGovernance() (common.Address, error) {
	return _FeeDistributor.Contract.FlayGovernance(&_FeeDistributor.CallOpts)
}

// GetFeeCalculator is a free data retrieval call binding the contract method 0x71c4ddb0.
//
// Solidity: function getFeeCalculator(bool _isFairLaunch) view returns(address)
func (_FeeDistributor *FeeDistributorCaller) GetFeeCalculator(opts *bind.CallOpts, _isFairLaunch bool) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "getFeeCalculator", _isFairLaunch)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeCalculator is a free data retrieval call binding the contract method 0x71c4ddb0.
//
// Solidity: function getFeeCalculator(bool _isFairLaunch) view returns(address)
func (_FeeDistributor *FeeDistributorSession) GetFeeCalculator(_isFairLaunch bool) (common.Address, error) {
	return _FeeDistributor.Contract.GetFeeCalculator(&_FeeDistributor.CallOpts, _isFairLaunch)
}

// GetFeeCalculator is a free data retrieval call binding the contract method 0x71c4ddb0.
//
// Solidity: function getFeeCalculator(bool _isFairLaunch) view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) GetFeeCalculator(_isFairLaunch bool) (common.Address, error) {
	return _FeeDistributor.Contract.GetFeeCalculator(&_FeeDistributor.CallOpts, _isFairLaunch)
}

// GetPoolFeeDistribution is a free data retrieval call binding the contract method 0xb3b42795.
//
// Solidity: function getPoolFeeDistribution(bytes32 _poolId) view returns((uint24,uint24,uint24,bool) feeDistribution_)
func (_FeeDistributor *FeeDistributorCaller) GetPoolFeeDistribution(opts *bind.CallOpts, _poolId [32]byte) (FeeDistributorFeeDistribution, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "getPoolFeeDistribution", _poolId)

	if err != nil {
		return *new(FeeDistributorFeeDistribution), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeDistributorFeeDistribution)).(*FeeDistributorFeeDistribution)

	return out0, err

}

// GetPoolFeeDistribution is a free data retrieval call binding the contract method 0xb3b42795.
//
// Solidity: function getPoolFeeDistribution(bytes32 _poolId) view returns((uint24,uint24,uint24,bool) feeDistribution_)
func (_FeeDistributor *FeeDistributorSession) GetPoolFeeDistribution(_poolId [32]byte) (FeeDistributorFeeDistribution, error) {
	return _FeeDistributor.Contract.GetPoolFeeDistribution(&_FeeDistributor.CallOpts, _poolId)
}

// GetPoolFeeDistribution is a free data retrieval call binding the contract method 0xb3b42795.
//
// Solidity: function getPoolFeeDistribution(bytes32 _poolId) view returns((uint24,uint24,uint24,bool) feeDistribution_)
func (_FeeDistributor *FeeDistributorCallerSession) GetPoolFeeDistribution(_poolId [32]byte) (FeeDistributorFeeDistribution, error) {
	return _FeeDistributor.Contract.GetPoolFeeDistribution(&_FeeDistributor.CallOpts, _poolId)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_FeeDistributor *FeeDistributorCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_FeeDistributor *FeeDistributorSession) NativeToken() (common.Address, error) {
	return _FeeDistributor.Contract.NativeToken(&_FeeDistributor.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) NativeToken() (common.Address, error) {
	return _FeeDistributor.Contract.NativeToken(&_FeeDistributor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_FeeDistributor *FeeDistributorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_FeeDistributor *FeeDistributorSession) Owner() (common.Address, error) {
	return _FeeDistributor.Contract.Owner(&_FeeDistributor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_FeeDistributor *FeeDistributorCallerSession) Owner() (common.Address, error) {
	return _FeeDistributor.Contract.Owner(&_FeeDistributor.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_FeeDistributor *FeeDistributorCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_FeeDistributor *FeeDistributorSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _FeeDistributor.Contract.OwnershipHandoverExpiresAt(&_FeeDistributor.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_FeeDistributor *FeeDistributorCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _FeeDistributor.Contract.OwnershipHandoverExpiresAt(&_FeeDistributor.CallOpts, pendingOwner)
}

// ReferralEscrow is a free data retrieval call binding the contract method 0x8f2bdf75.
//
// Solidity: function referralEscrow() view returns(address)
func (_FeeDistributor *FeeDistributorCaller) ReferralEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "referralEscrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReferralEscrow is a free data retrieval call binding the contract method 0x8f2bdf75.
//
// Solidity: function referralEscrow() view returns(address)
func (_FeeDistributor *FeeDistributorSession) ReferralEscrow() (common.Address, error) {
	return _FeeDistributor.Contract.ReferralEscrow(&_FeeDistributor.CallOpts)
}

// ReferralEscrow is a free data retrieval call binding the contract method 0x8f2bdf75.
//
// Solidity: function referralEscrow() view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) ReferralEscrow() (common.Address, error) {
	return _FeeDistributor.Contract.ReferralEscrow(&_FeeDistributor.CallOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_FeeDistributor *FeeDistributorTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_FeeDistributor *FeeDistributorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _FeeDistributor.Contract.CancelOwnershipHandover(&_FeeDistributor.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_FeeDistributor *FeeDistributorTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _FeeDistributor.Contract.CancelOwnershipHandover(&_FeeDistributor.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_FeeDistributor *FeeDistributorTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_FeeDistributor *FeeDistributorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.CompleteOwnershipHandover(&_FeeDistributor.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_FeeDistributor *FeeDistributorTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.CompleteOwnershipHandover(&_FeeDistributor.TransactOpts, pendingOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_FeeDistributor *FeeDistributorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_FeeDistributor *FeeDistributorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeDistributor.Contract.RenounceOwnership(&_FeeDistributor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_FeeDistributor *FeeDistributorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeDistributor.Contract.RenounceOwnership(&_FeeDistributor.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_FeeDistributor *FeeDistributorTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_FeeDistributor *FeeDistributorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _FeeDistributor.Contract.RequestOwnershipHandover(&_FeeDistributor.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_FeeDistributor *FeeDistributorTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _FeeDistributor.Contract.RequestOwnershipHandover(&_FeeDistributor.TransactOpts)
}

// SetFairLaunchFeeCalculator is a paid mutator transaction binding the contract method 0x38b1e700.
//
// Solidity: function setFairLaunchFeeCalculator(address _feeCalculator) returns()
func (_FeeDistributor *FeeDistributorTransactor) SetFairLaunchFeeCalculator(opts *bind.TransactOpts, _feeCalculator common.Address) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "setFairLaunchFeeCalculator", _feeCalculator)
}

// SetFairLaunchFeeCalculator is a paid mutator transaction binding the contract method 0x38b1e700.
//
// Solidity: function setFairLaunchFeeCalculator(address _feeCalculator) returns()
func (_FeeDistributor *FeeDistributorSession) SetFairLaunchFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetFairLaunchFeeCalculator(&_FeeDistributor.TransactOpts, _feeCalculator)
}

// SetFairLaunchFeeCalculator is a paid mutator transaction binding the contract method 0x38b1e700.
//
// Solidity: function setFairLaunchFeeCalculator(address _feeCalculator) returns()
func (_FeeDistributor *FeeDistributorTransactorSession) SetFairLaunchFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetFairLaunchFeeCalculator(&_FeeDistributor.TransactOpts, _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_FeeDistributor *FeeDistributorTransactor) SetFeeCalculator(opts *bind.TransactOpts, _feeCalculator common.Address) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "setFeeCalculator", _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_FeeDistributor *FeeDistributorSession) SetFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetFeeCalculator(&_FeeDistributor.TransactOpts, _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_FeeDistributor *FeeDistributorTransactorSession) SetFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetFeeCalculator(&_FeeDistributor.TransactOpts, _feeCalculator)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0xd3bff717.
//
// Solidity: function setFeeDistribution((uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_FeeDistributor *FeeDistributorTransactor) SetFeeDistribution(opts *bind.TransactOpts, _feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "setFeeDistribution", _feeDistribution)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0xd3bff717.
//
// Solidity: function setFeeDistribution((uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_FeeDistributor *FeeDistributorSession) SetFeeDistribution(_feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetFeeDistribution(&_FeeDistributor.TransactOpts, _feeDistribution)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0xd3bff717.
//
// Solidity: function setFeeDistribution((uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_FeeDistributor *FeeDistributorTransactorSession) SetFeeDistribution(_feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetFeeDistribution(&_FeeDistributor.TransactOpts, _feeDistribution)
}

// SetPoolFeeDistribution is a paid mutator transaction binding the contract method 0x4ed0f0f5.
//
// Solidity: function setPoolFeeDistribution(bytes32 _poolId, (uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_FeeDistributor *FeeDistributorTransactor) SetPoolFeeDistribution(opts *bind.TransactOpts, _poolId [32]byte, _feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "setPoolFeeDistribution", _poolId, _feeDistribution)
}

// SetPoolFeeDistribution is a paid mutator transaction binding the contract method 0x4ed0f0f5.
//
// Solidity: function setPoolFeeDistribution(bytes32 _poolId, (uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_FeeDistributor *FeeDistributorSession) SetPoolFeeDistribution(_poolId [32]byte, _feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetPoolFeeDistribution(&_FeeDistributor.TransactOpts, _poolId, _feeDistribution)
}

// SetPoolFeeDistribution is a paid mutator transaction binding the contract method 0x4ed0f0f5.
//
// Solidity: function setPoolFeeDistribution(bytes32 _poolId, (uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_FeeDistributor *FeeDistributorTransactorSession) SetPoolFeeDistribution(_poolId [32]byte, _feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetPoolFeeDistribution(&_FeeDistributor.TransactOpts, _poolId, _feeDistribution)
}

// SetProtocolFeeDistribution is a paid mutator transaction binding the contract method 0x2423028c.
//
// Solidity: function setProtocolFeeDistribution(uint24 _protocol) returns()
func (_FeeDistributor *FeeDistributorTransactor) SetProtocolFeeDistribution(opts *bind.TransactOpts, _protocol *big.Int) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "setProtocolFeeDistribution", _protocol)
}

// SetProtocolFeeDistribution is a paid mutator transaction binding the contract method 0x2423028c.
//
// Solidity: function setProtocolFeeDistribution(uint24 _protocol) returns()
func (_FeeDistributor *FeeDistributorSession) SetProtocolFeeDistribution(_protocol *big.Int) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetProtocolFeeDistribution(&_FeeDistributor.TransactOpts, _protocol)
}

// SetProtocolFeeDistribution is a paid mutator transaction binding the contract method 0x2423028c.
//
// Solidity: function setProtocolFeeDistribution(uint24 _protocol) returns()
func (_FeeDistributor *FeeDistributorTransactorSession) SetProtocolFeeDistribution(_protocol *big.Int) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetProtocolFeeDistribution(&_FeeDistributor.TransactOpts, _protocol)
}

// SetReferralEscrow is a paid mutator transaction binding the contract method 0xa87277dd.
//
// Solidity: function setReferralEscrow(address _referralEscrow) returns()
func (_FeeDistributor *FeeDistributorTransactor) SetReferralEscrow(opts *bind.TransactOpts, _referralEscrow common.Address) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "setReferralEscrow", _referralEscrow)
}

// SetReferralEscrow is a paid mutator transaction binding the contract method 0xa87277dd.
//
// Solidity: function setReferralEscrow(address _referralEscrow) returns()
func (_FeeDistributor *FeeDistributorSession) SetReferralEscrow(_referralEscrow common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetReferralEscrow(&_FeeDistributor.TransactOpts, _referralEscrow)
}

// SetReferralEscrow is a paid mutator transaction binding the contract method 0xa87277dd.
//
// Solidity: function setReferralEscrow(address _referralEscrow) returns()
func (_FeeDistributor *FeeDistributorTransactorSession) SetReferralEscrow(_referralEscrow common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.SetReferralEscrow(&_FeeDistributor.TransactOpts, _referralEscrow)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_FeeDistributor *FeeDistributorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_FeeDistributor *FeeDistributorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.TransferOwnership(&_FeeDistributor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_FeeDistributor *FeeDistributorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.TransferOwnership(&_FeeDistributor.TransactOpts, newOwner)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x4c2d94c0.
//
// Solidity: function withdrawFees(address _recipient, bool _unwrap) returns()
func (_FeeDistributor *FeeDistributorTransactor) WithdrawFees(opts *bind.TransactOpts, _recipient common.Address, _unwrap bool) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "withdrawFees", _recipient, _unwrap)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x4c2d94c0.
//
// Solidity: function withdrawFees(address _recipient, bool _unwrap) returns()
func (_FeeDistributor *FeeDistributorSession) WithdrawFees(_recipient common.Address, _unwrap bool) (*types.Transaction, error) {
	return _FeeDistributor.Contract.WithdrawFees(&_FeeDistributor.TransactOpts, _recipient, _unwrap)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x4c2d94c0.
//
// Solidity: function withdrawFees(address _recipient, bool _unwrap) returns()
func (_FeeDistributor *FeeDistributorTransactorSession) WithdrawFees(_recipient common.Address, _unwrap bool) (*types.Transaction, error) {
	return _FeeDistributor.Contract.WithdrawFees(&_FeeDistributor.TransactOpts, _recipient, _unwrap)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeDistributor *FeeDistributorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeDistributor *FeeDistributorSession) Receive() (*types.Transaction, error) {
	return _FeeDistributor.Contract.Receive(&_FeeDistributor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeDistributor *FeeDistributorTransactorSession) Receive() (*types.Transaction, error) {
	return _FeeDistributor.Contract.Receive(&_FeeDistributor.TransactOpts)
}

// FeeDistributorCreatorFeeAllocationUpdatedIterator is returned from FilterCreatorFeeAllocationUpdated and is used to iterate over the raw logs and unpacked data for CreatorFeeAllocationUpdated events raised by the FeeDistributor contract.
type FeeDistributorCreatorFeeAllocationUpdatedIterator struct {
	Event *FeeDistributorCreatorFeeAllocationUpdated // Event containing the contract specifics and raw log

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
func (it *FeeDistributorCreatorFeeAllocationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorCreatorFeeAllocationUpdated)
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
		it.Event = new(FeeDistributorCreatorFeeAllocationUpdated)
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
func (it *FeeDistributorCreatorFeeAllocationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorCreatorFeeAllocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorCreatorFeeAllocationUpdated represents a CreatorFeeAllocationUpdated event raised by the FeeDistributor contract.
type FeeDistributorCreatorFeeAllocationUpdated struct {
	PoolId     [32]byte
	Allocation *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatorFeeAllocationUpdated is a free log retrieval operation binding the contract event 0x6d9755b9ab4c1e27104da15a9b4589ffe57a1222d13193f2f1354ffde5fee10e.
//
// Solidity: event CreatorFeeAllocationUpdated(bytes32 indexed _poolId, uint24 _allocation)
func (_FeeDistributor *FeeDistributorFilterer) FilterCreatorFeeAllocationUpdated(opts *bind.FilterOpts, _poolId [][32]byte) (*FeeDistributorCreatorFeeAllocationUpdatedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "CreatorFeeAllocationUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorCreatorFeeAllocationUpdatedIterator{contract: _FeeDistributor.contract, event: "CreatorFeeAllocationUpdated", logs: logs, sub: sub}, nil
}

// WatchCreatorFeeAllocationUpdated is a free log subscription operation binding the contract event 0x6d9755b9ab4c1e27104da15a9b4589ffe57a1222d13193f2f1354ffde5fee10e.
//
// Solidity: event CreatorFeeAllocationUpdated(bytes32 indexed _poolId, uint24 _allocation)
func (_FeeDistributor *FeeDistributorFilterer) WatchCreatorFeeAllocationUpdated(opts *bind.WatchOpts, sink chan<- *FeeDistributorCreatorFeeAllocationUpdated, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "CreatorFeeAllocationUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorCreatorFeeAllocationUpdated)
				if err := _FeeDistributor.contract.UnpackLog(event, "CreatorFeeAllocationUpdated", log); err != nil {
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

// ParseCreatorFeeAllocationUpdated is a log parse operation binding the contract event 0x6d9755b9ab4c1e27104da15a9b4589ffe57a1222d13193f2f1354ffde5fee10e.
//
// Solidity: event CreatorFeeAllocationUpdated(bytes32 indexed _poolId, uint24 _allocation)
func (_FeeDistributor *FeeDistributorFilterer) ParseCreatorFeeAllocationUpdated(log types.Log) (*FeeDistributorCreatorFeeAllocationUpdated, error) {
	event := new(FeeDistributorCreatorFeeAllocationUpdated)
	if err := _FeeDistributor.contract.UnpackLog(event, "CreatorFeeAllocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the FeeDistributor contract.
type FeeDistributorDepositIterator struct {
	Event *FeeDistributorDeposit // Event containing the contract specifics and raw log

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
func (it *FeeDistributorDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorDeposit)
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
		it.Event = new(FeeDistributorDeposit)
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
func (it *FeeDistributorDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorDeposit represents a Deposit event raised by the FeeDistributor contract.
type FeeDistributorDeposit struct {
	PoolId [32]byte
	Payee  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xc95ddcaddf83340b68d0d44c01b1703f5d28d0611a3fd87e69d79ba7e2ac21d3.
//
// Solidity: event Deposit(bytes32 indexed _poolId, address _payee, address _token, uint256 _amount)
func (_FeeDistributor *FeeDistributorFilterer) FilterDeposit(opts *bind.FilterOpts, _poolId [][32]byte) (*FeeDistributorDepositIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "Deposit", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorDepositIterator{contract: _FeeDistributor.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xc95ddcaddf83340b68d0d44c01b1703f5d28d0611a3fd87e69d79ba7e2ac21d3.
//
// Solidity: event Deposit(bytes32 indexed _poolId, address _payee, address _token, uint256 _amount)
func (_FeeDistributor *FeeDistributorFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *FeeDistributorDeposit, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "Deposit", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorDeposit)
				if err := _FeeDistributor.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xc95ddcaddf83340b68d0d44c01b1703f5d28d0611a3fd87e69d79ba7e2ac21d3.
//
// Solidity: event Deposit(bytes32 indexed _poolId, address _payee, address _token, uint256 _amount)
func (_FeeDistributor *FeeDistributorFilterer) ParseDeposit(log types.Log) (*FeeDistributorDeposit, error) {
	event := new(FeeDistributorDeposit)
	if err := _FeeDistributor.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorFairLaunchFeeCalculatorUpdatedIterator is returned from FilterFairLaunchFeeCalculatorUpdated and is used to iterate over the raw logs and unpacked data for FairLaunchFeeCalculatorUpdated events raised by the FeeDistributor contract.
type FeeDistributorFairLaunchFeeCalculatorUpdatedIterator struct {
	Event *FeeDistributorFairLaunchFeeCalculatorUpdated // Event containing the contract specifics and raw log

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
func (it *FeeDistributorFairLaunchFeeCalculatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorFairLaunchFeeCalculatorUpdated)
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
		it.Event = new(FeeDistributorFairLaunchFeeCalculatorUpdated)
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
func (it *FeeDistributorFairLaunchFeeCalculatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorFairLaunchFeeCalculatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorFairLaunchFeeCalculatorUpdated represents a FairLaunchFeeCalculatorUpdated event raised by the FeeDistributor contract.
type FeeDistributorFairLaunchFeeCalculatorUpdated struct {
	FeeCalculator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFairLaunchFeeCalculatorUpdated is a free log retrieval operation binding the contract event 0x87043577396d39ef835a9eb69bb496c219cc61bdd6e718447add3c06b6cc0844.
//
// Solidity: event FairLaunchFeeCalculatorUpdated(address _feeCalculator)
func (_FeeDistributor *FeeDistributorFilterer) FilterFairLaunchFeeCalculatorUpdated(opts *bind.FilterOpts) (*FeeDistributorFairLaunchFeeCalculatorUpdatedIterator, error) {

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "FairLaunchFeeCalculatorUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeDistributorFairLaunchFeeCalculatorUpdatedIterator{contract: _FeeDistributor.contract, event: "FairLaunchFeeCalculatorUpdated", logs: logs, sub: sub}, nil
}

// WatchFairLaunchFeeCalculatorUpdated is a free log subscription operation binding the contract event 0x87043577396d39ef835a9eb69bb496c219cc61bdd6e718447add3c06b6cc0844.
//
// Solidity: event FairLaunchFeeCalculatorUpdated(address _feeCalculator)
func (_FeeDistributor *FeeDistributorFilterer) WatchFairLaunchFeeCalculatorUpdated(opts *bind.WatchOpts, sink chan<- *FeeDistributorFairLaunchFeeCalculatorUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "FairLaunchFeeCalculatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorFairLaunchFeeCalculatorUpdated)
				if err := _FeeDistributor.contract.UnpackLog(event, "FairLaunchFeeCalculatorUpdated", log); err != nil {
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

// ParseFairLaunchFeeCalculatorUpdated is a log parse operation binding the contract event 0x87043577396d39ef835a9eb69bb496c219cc61bdd6e718447add3c06b6cc0844.
//
// Solidity: event FairLaunchFeeCalculatorUpdated(address _feeCalculator)
func (_FeeDistributor *FeeDistributorFilterer) ParseFairLaunchFeeCalculatorUpdated(log types.Log) (*FeeDistributorFairLaunchFeeCalculatorUpdated, error) {
	event := new(FeeDistributorFairLaunchFeeCalculatorUpdated)
	if err := _FeeDistributor.contract.UnpackLog(event, "FairLaunchFeeCalculatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorFeeCalculatorUpdatedIterator is returned from FilterFeeCalculatorUpdated and is used to iterate over the raw logs and unpacked data for FeeCalculatorUpdated events raised by the FeeDistributor contract.
type FeeDistributorFeeCalculatorUpdatedIterator struct {
	Event *FeeDistributorFeeCalculatorUpdated // Event containing the contract specifics and raw log

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
func (it *FeeDistributorFeeCalculatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorFeeCalculatorUpdated)
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
		it.Event = new(FeeDistributorFeeCalculatorUpdated)
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
func (it *FeeDistributorFeeCalculatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorFeeCalculatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorFeeCalculatorUpdated represents a FeeCalculatorUpdated event raised by the FeeDistributor contract.
type FeeDistributorFeeCalculatorUpdated struct {
	FeeCalculator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeCalculatorUpdated is a free log retrieval operation binding the contract event 0x3e762c7e655633ce63121393b9694f9ca1883d14d18f48f1be55e5dc7a9fb6c1.
//
// Solidity: event FeeCalculatorUpdated(address _feeCalculator)
func (_FeeDistributor *FeeDistributorFilterer) FilterFeeCalculatorUpdated(opts *bind.FilterOpts) (*FeeDistributorFeeCalculatorUpdatedIterator, error) {

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "FeeCalculatorUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeDistributorFeeCalculatorUpdatedIterator{contract: _FeeDistributor.contract, event: "FeeCalculatorUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeCalculatorUpdated is a free log subscription operation binding the contract event 0x3e762c7e655633ce63121393b9694f9ca1883d14d18f48f1be55e5dc7a9fb6c1.
//
// Solidity: event FeeCalculatorUpdated(address _feeCalculator)
func (_FeeDistributor *FeeDistributorFilterer) WatchFeeCalculatorUpdated(opts *bind.WatchOpts, sink chan<- *FeeDistributorFeeCalculatorUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "FeeCalculatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorFeeCalculatorUpdated)
				if err := _FeeDistributor.contract.UnpackLog(event, "FeeCalculatorUpdated", log); err != nil {
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

// ParseFeeCalculatorUpdated is a log parse operation binding the contract event 0x3e762c7e655633ce63121393b9694f9ca1883d14d18f48f1be55e5dc7a9fb6c1.
//
// Solidity: event FeeCalculatorUpdated(address _feeCalculator)
func (_FeeDistributor *FeeDistributorFilterer) ParseFeeCalculatorUpdated(log types.Log) (*FeeDistributorFeeCalculatorUpdated, error) {
	event := new(FeeDistributorFeeCalculatorUpdated)
	if err := _FeeDistributor.contract.UnpackLog(event, "FeeCalculatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorFeeDistributionUpdatedIterator is returned from FilterFeeDistributionUpdated and is used to iterate over the raw logs and unpacked data for FeeDistributionUpdated events raised by the FeeDistributor contract.
type FeeDistributorFeeDistributionUpdatedIterator struct {
	Event *FeeDistributorFeeDistributionUpdated // Event containing the contract specifics and raw log

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
func (it *FeeDistributorFeeDistributionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorFeeDistributionUpdated)
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
		it.Event = new(FeeDistributorFeeDistributionUpdated)
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
func (it *FeeDistributorFeeDistributionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorFeeDistributionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorFeeDistributionUpdated represents a FeeDistributionUpdated event raised by the FeeDistributor contract.
type FeeDistributorFeeDistributionUpdated struct {
	FeeDistribution FeeDistributorFeeDistribution
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributionUpdated is a free log retrieval operation binding the contract event 0xca7619d8d03127c99378decb6d4c002b5b8b8d60918ea28c2300a7811e7c1300.
//
// Solidity: event FeeDistributionUpdated((uint24,uint24,uint24,bool) _feeDistribution)
func (_FeeDistributor *FeeDistributorFilterer) FilterFeeDistributionUpdated(opts *bind.FilterOpts) (*FeeDistributorFeeDistributionUpdatedIterator, error) {

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "FeeDistributionUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeDistributorFeeDistributionUpdatedIterator{contract: _FeeDistributor.contract, event: "FeeDistributionUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeDistributionUpdated is a free log subscription operation binding the contract event 0xca7619d8d03127c99378decb6d4c002b5b8b8d60918ea28c2300a7811e7c1300.
//
// Solidity: event FeeDistributionUpdated((uint24,uint24,uint24,bool) _feeDistribution)
func (_FeeDistributor *FeeDistributorFilterer) WatchFeeDistributionUpdated(opts *bind.WatchOpts, sink chan<- *FeeDistributorFeeDistributionUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "FeeDistributionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorFeeDistributionUpdated)
				if err := _FeeDistributor.contract.UnpackLog(event, "FeeDistributionUpdated", log); err != nil {
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

// ParseFeeDistributionUpdated is a log parse operation binding the contract event 0xca7619d8d03127c99378decb6d4c002b5b8b8d60918ea28c2300a7811e7c1300.
//
// Solidity: event FeeDistributionUpdated((uint24,uint24,uint24,bool) _feeDistribution)
func (_FeeDistributor *FeeDistributorFilterer) ParseFeeDistributionUpdated(log types.Log) (*FeeDistributorFeeDistributionUpdated, error) {
	event := new(FeeDistributorFeeDistributionUpdated)
	if err := _FeeDistributor.contract.UnpackLog(event, "FeeDistributionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the FeeDistributor contract.
type FeeDistributorOwnershipHandoverCanceledIterator struct {
	Event *FeeDistributorOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *FeeDistributorOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorOwnershipHandoverCanceled)
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
		it.Event = new(FeeDistributorOwnershipHandoverCanceled)
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
func (it *FeeDistributorOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the FeeDistributor contract.
type FeeDistributorOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_FeeDistributor *FeeDistributorFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*FeeDistributorOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorOwnershipHandoverCanceledIterator{contract: _FeeDistributor.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_FeeDistributor *FeeDistributorFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *FeeDistributorOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorOwnershipHandoverCanceled)
				if err := _FeeDistributor.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_FeeDistributor *FeeDistributorFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*FeeDistributorOwnershipHandoverCanceled, error) {
	event := new(FeeDistributorOwnershipHandoverCanceled)
	if err := _FeeDistributor.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the FeeDistributor contract.
type FeeDistributorOwnershipHandoverRequestedIterator struct {
	Event *FeeDistributorOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *FeeDistributorOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorOwnershipHandoverRequested)
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
		it.Event = new(FeeDistributorOwnershipHandoverRequested)
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
func (it *FeeDistributorOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the FeeDistributor contract.
type FeeDistributorOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_FeeDistributor *FeeDistributorFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*FeeDistributorOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorOwnershipHandoverRequestedIterator{contract: _FeeDistributor.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_FeeDistributor *FeeDistributorFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *FeeDistributorOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorOwnershipHandoverRequested)
				if err := _FeeDistributor.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_FeeDistributor *FeeDistributorFilterer) ParseOwnershipHandoverRequested(log types.Log) (*FeeDistributorOwnershipHandoverRequested, error) {
	event := new(FeeDistributorOwnershipHandoverRequested)
	if err := _FeeDistributor.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeeDistributor contract.
type FeeDistributorOwnershipTransferredIterator struct {
	Event *FeeDistributorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeeDistributorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorOwnershipTransferred)
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
		it.Event = new(FeeDistributorOwnershipTransferred)
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
func (it *FeeDistributorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorOwnershipTransferred represents a OwnershipTransferred event raised by the FeeDistributor contract.
type FeeDistributorOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_FeeDistributor *FeeDistributorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*FeeDistributorOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorOwnershipTransferredIterator{contract: _FeeDistributor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_FeeDistributor *FeeDistributorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeDistributorOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorOwnershipTransferred)
				if err := _FeeDistributor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeeDistributor *FeeDistributorFilterer) ParseOwnershipTransferred(log types.Log) (*FeeDistributorOwnershipTransferred, error) {
	event := new(FeeDistributorOwnershipTransferred)
	if err := _FeeDistributor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorPoolFeeDistributionUpdatedIterator is returned from FilterPoolFeeDistributionUpdated and is used to iterate over the raw logs and unpacked data for PoolFeeDistributionUpdated events raised by the FeeDistributor contract.
type FeeDistributorPoolFeeDistributionUpdatedIterator struct {
	Event *FeeDistributorPoolFeeDistributionUpdated // Event containing the contract specifics and raw log

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
func (it *FeeDistributorPoolFeeDistributionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorPoolFeeDistributionUpdated)
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
		it.Event = new(FeeDistributorPoolFeeDistributionUpdated)
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
func (it *FeeDistributorPoolFeeDistributionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorPoolFeeDistributionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorPoolFeeDistributionUpdated represents a PoolFeeDistributionUpdated event raised by the FeeDistributor contract.
type FeeDistributorPoolFeeDistributionUpdated struct {
	PoolId          [32]byte
	FeeDistribution FeeDistributorFeeDistribution
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPoolFeeDistributionUpdated is a free log retrieval operation binding the contract event 0xe1b2af81a774e1ebee3ca7c94c1e1a0df1210b236149a2079b8bba8dbc475a28.
//
// Solidity: event PoolFeeDistributionUpdated(bytes32 indexed _poolId, (uint24,uint24,uint24,bool) _feeDistribution)
func (_FeeDistributor *FeeDistributorFilterer) FilterPoolFeeDistributionUpdated(opts *bind.FilterOpts, _poolId [][32]byte) (*FeeDistributorPoolFeeDistributionUpdatedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "PoolFeeDistributionUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorPoolFeeDistributionUpdatedIterator{contract: _FeeDistributor.contract, event: "PoolFeeDistributionUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolFeeDistributionUpdated is a free log subscription operation binding the contract event 0xe1b2af81a774e1ebee3ca7c94c1e1a0df1210b236149a2079b8bba8dbc475a28.
//
// Solidity: event PoolFeeDistributionUpdated(bytes32 indexed _poolId, (uint24,uint24,uint24,bool) _feeDistribution)
func (_FeeDistributor *FeeDistributorFilterer) WatchPoolFeeDistributionUpdated(opts *bind.WatchOpts, sink chan<- *FeeDistributorPoolFeeDistributionUpdated, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "PoolFeeDistributionUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorPoolFeeDistributionUpdated)
				if err := _FeeDistributor.contract.UnpackLog(event, "PoolFeeDistributionUpdated", log); err != nil {
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

// ParsePoolFeeDistributionUpdated is a log parse operation binding the contract event 0xe1b2af81a774e1ebee3ca7c94c1e1a0df1210b236149a2079b8bba8dbc475a28.
//
// Solidity: event PoolFeeDistributionUpdated(bytes32 indexed _poolId, (uint24,uint24,uint24,bool) _feeDistribution)
func (_FeeDistributor *FeeDistributorFilterer) ParsePoolFeeDistributionUpdated(log types.Log) (*FeeDistributorPoolFeeDistributionUpdated, error) {
	event := new(FeeDistributorPoolFeeDistributionUpdated)
	if err := _FeeDistributor.contract.UnpackLog(event, "PoolFeeDistributionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorReferralEscrowUpdatedIterator is returned from FilterReferralEscrowUpdated and is used to iterate over the raw logs and unpacked data for ReferralEscrowUpdated events raised by the FeeDistributor contract.
type FeeDistributorReferralEscrowUpdatedIterator struct {
	Event *FeeDistributorReferralEscrowUpdated // Event containing the contract specifics and raw log

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
func (it *FeeDistributorReferralEscrowUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorReferralEscrowUpdated)
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
		it.Event = new(FeeDistributorReferralEscrowUpdated)
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
func (it *FeeDistributorReferralEscrowUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorReferralEscrowUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorReferralEscrowUpdated represents a ReferralEscrowUpdated event raised by the FeeDistributor contract.
type FeeDistributorReferralEscrowUpdated struct {
	ReferralEscrow common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReferralEscrowUpdated is a free log retrieval operation binding the contract event 0x9c922de256a07b4d188faacda5c1abb2cae12f74f4370d5c2f11efb37a742d70.
//
// Solidity: event ReferralEscrowUpdated(address _referralEscrow)
func (_FeeDistributor *FeeDistributorFilterer) FilterReferralEscrowUpdated(opts *bind.FilterOpts) (*FeeDistributorReferralEscrowUpdatedIterator, error) {

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "ReferralEscrowUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeDistributorReferralEscrowUpdatedIterator{contract: _FeeDistributor.contract, event: "ReferralEscrowUpdated", logs: logs, sub: sub}, nil
}

// WatchReferralEscrowUpdated is a free log subscription operation binding the contract event 0x9c922de256a07b4d188faacda5c1abb2cae12f74f4370d5c2f11efb37a742d70.
//
// Solidity: event ReferralEscrowUpdated(address _referralEscrow)
func (_FeeDistributor *FeeDistributorFilterer) WatchReferralEscrowUpdated(opts *bind.WatchOpts, sink chan<- *FeeDistributorReferralEscrowUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "ReferralEscrowUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorReferralEscrowUpdated)
				if err := _FeeDistributor.contract.UnpackLog(event, "ReferralEscrowUpdated", log); err != nil {
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

// ParseReferralEscrowUpdated is a log parse operation binding the contract event 0x9c922de256a07b4d188faacda5c1abb2cae12f74f4370d5c2f11efb37a742d70.
//
// Solidity: event ReferralEscrowUpdated(address _referralEscrow)
func (_FeeDistributor *FeeDistributorFilterer) ParseReferralEscrowUpdated(log types.Log) (*FeeDistributorReferralEscrowUpdated, error) {
	event := new(FeeDistributorReferralEscrowUpdated)
	if err := _FeeDistributor.contract.UnpackLog(event, "ReferralEscrowUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorReferrerFeePaidIterator is returned from FilterReferrerFeePaid and is used to iterate over the raw logs and unpacked data for ReferrerFeePaid events raised by the FeeDistributor contract.
type FeeDistributorReferrerFeePaidIterator struct {
	Event *FeeDistributorReferrerFeePaid // Event containing the contract specifics and raw log

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
func (it *FeeDistributorReferrerFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorReferrerFeePaid)
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
		it.Event = new(FeeDistributorReferrerFeePaid)
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
func (it *FeeDistributorReferrerFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorReferrerFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorReferrerFeePaid represents a ReferrerFeePaid event raised by the FeeDistributor contract.
type FeeDistributorReferrerFeePaid struct {
	PoolId    [32]byte
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReferrerFeePaid is a free log retrieval operation binding the contract event 0x3ca4a7850462c23d5854e8b3e852626beb21b37354c670c0135ab46f0c4bdc31.
//
// Solidity: event ReferrerFeePaid(bytes32 indexed _poolId, address _recipient, address _token, uint256 _amount)
func (_FeeDistributor *FeeDistributorFilterer) FilterReferrerFeePaid(opts *bind.FilterOpts, _poolId [][32]byte) (*FeeDistributorReferrerFeePaidIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "ReferrerFeePaid", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorReferrerFeePaidIterator{contract: _FeeDistributor.contract, event: "ReferrerFeePaid", logs: logs, sub: sub}, nil
}

// WatchReferrerFeePaid is a free log subscription operation binding the contract event 0x3ca4a7850462c23d5854e8b3e852626beb21b37354c670c0135ab46f0c4bdc31.
//
// Solidity: event ReferrerFeePaid(bytes32 indexed _poolId, address _recipient, address _token, uint256 _amount)
func (_FeeDistributor *FeeDistributorFilterer) WatchReferrerFeePaid(opts *bind.WatchOpts, sink chan<- *FeeDistributorReferrerFeePaid, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "ReferrerFeePaid", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorReferrerFeePaid)
				if err := _FeeDistributor.contract.UnpackLog(event, "ReferrerFeePaid", log); err != nil {
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

// ParseReferrerFeePaid is a log parse operation binding the contract event 0x3ca4a7850462c23d5854e8b3e852626beb21b37354c670c0135ab46f0c4bdc31.
//
// Solidity: event ReferrerFeePaid(bytes32 indexed _poolId, address _recipient, address _token, uint256 _amount)
func (_FeeDistributor *FeeDistributorFilterer) ParseReferrerFeePaid(log types.Log) (*FeeDistributorReferrerFeePaid, error) {
	event := new(FeeDistributorReferrerFeePaid)
	if err := _FeeDistributor.contract.UnpackLog(event, "ReferrerFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the FeeDistributor contract.
type FeeDistributorWithdrawalIterator struct {
	Event *FeeDistributorWithdrawal // Event containing the contract specifics and raw log

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
func (it *FeeDistributorWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorWithdrawal)
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
		it.Event = new(FeeDistributorWithdrawal)
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
func (it *FeeDistributorWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorWithdrawal represents a Withdrawal event raised by the FeeDistributor contract.
type FeeDistributorWithdrawal struct {
	Sender    common.Address
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x342e7ff505a8a0364cd0dc2ff195c315e43bce86b204846ecd36913e117b109e.
//
// Solidity: event Withdrawal(address _sender, address _recipient, address _token, uint256 _amount)
func (_FeeDistributor *FeeDistributorFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*FeeDistributorWithdrawalIterator, error) {

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &FeeDistributorWithdrawalIterator{contract: _FeeDistributor.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x342e7ff505a8a0364cd0dc2ff195c315e43bce86b204846ecd36913e117b109e.
//
// Solidity: event Withdrawal(address _sender, address _recipient, address _token, uint256 _amount)
func (_FeeDistributor *FeeDistributorFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *FeeDistributorWithdrawal) (event.Subscription, error) {

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorWithdrawal)
				if err := _FeeDistributor.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x342e7ff505a8a0364cd0dc2ff195c315e43bce86b204846ecd36913e117b109e.
//
// Solidity: event Withdrawal(address _sender, address _recipient, address _token, uint256 _amount)
func (_FeeDistributor *FeeDistributorFilterer) ParseWithdrawal(log types.Log) (*FeeDistributorWithdrawal, error) {
	event := new(FeeDistributorWithdrawal)
	if err := _FeeDistributor.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
