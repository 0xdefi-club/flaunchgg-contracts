// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package notifier

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

// NotifierMetaData contains all meta data concerning the Notifier contract.
var NotifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_protocolOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"notifySubscribers\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_key\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"subscribe\",\"inputs\":[{\"name\":\"_subscriber\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"unsubscribe\",\"inputs\":[{\"name\":\"_subscriber\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Subscription\",\"inputs\":[{\"name\":\"_subscriber\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unsubscription\",\"inputs\":[{\"name\":\"_subscriber\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SubscriptionReverted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b506040516109d33803806109d3833981016040819052602b916084565b600280546001600160a01b031916331790556044816049565b5060af565b6001600160a01b0316638b78c6d819819055805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b5f602082840312156093575f5ffd5b81516001600160a01b038116811460a8575f5ffd5b9392505050565b610917806100bc5f395ff3fe608060405260043610610080575f3560e01c80632569296214610084578063314e79ad1461008e57806354d1f13d146100ad578063715018a6146100b55780637262561c146100bd5780638da5cb5b146100dc578063d5f9a39714610105578063f04e283e14610124578063f2fde38b14610137578063fee81cf41461014a575b5f5ffd5b61008c610189565b005b348015610099575f5ffd5b5061008c6100a836600461070d565b6101d5565b61008c61027c565b61008c6102b5565b3480156100c8575f5ffd5b5061008c6100d736600461078b565b6102c8565b3480156100e7575f5ffd5b50638b78c6d819546040516100fc91906107a4565b60405180910390f35b348015610110575f5ffd5b5061008c61011f3660046107b8565b61036f565b61008c61013236600461078b565b61044e565b61008c61014536600461078b565b610488565b348015610155575f5ffd5b5061017b61016436600461078b565b63389a75e1600c9081525f91909152602090205490565b6040519081526020016100fc565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b6002546001600160a01b031633146101eb575f5ffd5b5f6101f55f6104ae565b90505f5b818110156102745761020b5f826104bd565b6001600160a01b0316634ca1f062878787876040518563ffffffff1660e01b815260040161023c949392919061082e565b5f604051808303815f87803b158015610253575f5ffd5b505af1158015610265573d5f5f3e3d5ffd5b505050508060010190506101f9565b505050505050565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b6102bd6104cf565b6102c65f6104e9565b565b6102d06104cf565b6102da5f82610526565b1561036c576102e95f8261053a565b50806001600160a01b031663fcae44846040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610322575f5ffd5b505af1925050508015610333575060015b507f804bdbbf25e80c5635294ecf375fa16a531a4b7ffa416f74e121ae40bf4a00fe8160405161036391906107a4565b60405180910390a15b50565b6103776104cf565b6103815f8461054e565b15610449576040516323085b8560e11b81526001600160a01b03841690634610b70a906103b49085908590600401610860565b6020604051808303815f875af11580156103d0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103f4919061087b565b61041157604051633bf9fecd60e01b815260040160405180910390fd5b7fa6610d70b57486ac216657a87c0fce7f41758dd77d526daf7974f813a9b2a46a8360405161044091906107a4565b60405180910390a15b505050565b6104566104cf565b63389a75e1600c52805f526020600c20805442111561047c57636f5e88185f526004601cfd5b5f905561036c816104e9565b6104906104cf565b8060601b6104a557637448fbae5f526004601cfd5b61036c816104e9565b5f6104b7825490565b92915050565b5f6104c88383610562565b9392505050565b638b78c6d8195433146102c6576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a355565b5f6104c8836001600160a01b038416610588565b5f6104c8836001600160a01b03841661059f565b5f6104c8836001600160a01b038416610682565b5f825f0182815481106105775761057761089a565b905f5260205f200154905092915050565b5f9081526001919091016020526040902054151590565b5f8181526001830160205260408120548015610679575f6105c16001836108ae565b85549091505f906105d4906001906108ae565b9050808214610633575f865f0182815481106105f2576105f261089a565b905f5260205f200154905080875f0184815481106106125761061261089a565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080610644576106446108cd565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506104b7565b5f9150506104b7565b5f61068d8383610588565b6106c257508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556104b7565b505f6104b7565b5f5f83601f8401126106d9575f5ffd5b5081356001600160401b038111156106ef575f5ffd5b602083019150836020828501011115610706575f5ffd5b9250929050565b5f5f5f5f60608587031215610720575f5ffd5b8435935060208501356001600160e01b03198116811461073e575f5ffd5b925060408501356001600160401b03811115610758575f5ffd5b610764878288016106c9565b95989497509550505050565b80356001600160a01b0381168114610786575f5ffd5b919050565b5f6020828403121561079b575f5ffd5b6104c882610770565b6001600160a01b0391909116815260200190565b5f5f5f604084860312156107ca575f5ffd5b6107d384610770565b925060208401356001600160401b038111156107ed575f5ffd5b6107f9868287016106c9565b9497909650939450505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b84815263ffffffff60e01b84166020820152606060408201525f610856606083018486610806565b9695505050505050565b602081525f610873602083018486610806565b949350505050565b5f6020828403121561088b575f5ffd5b815180151581146104c8575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b818103818111156104b757634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220e7d3fe79eab14863996255c9403a53f57af392f2a80225ae8bdabd164fe2b2aa64736f6c634300081b0033",
}

// NotifierABI is the input ABI used to generate the binding from.
// Deprecated: Use NotifierMetaData.ABI instead.
var NotifierABI = NotifierMetaData.ABI

// NotifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NotifierMetaData.Bin instead.
var NotifierBin = NotifierMetaData.Bin

// DeployNotifier deploys a new Ethereum contract, binding an instance of Notifier to it.
func DeployNotifier(auth *bind.TransactOpts, backend bind.ContractBackend, _protocolOwner common.Address) (common.Address, *types.Transaction, *Notifier, error) {
	parsed, err := NotifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NotifierBin), backend, _protocolOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Notifier{NotifierCaller: NotifierCaller{contract: contract}, NotifierTransactor: NotifierTransactor{contract: contract}, NotifierFilterer: NotifierFilterer{contract: contract}}, nil
}

// Notifier is an auto generated Go binding around an Ethereum contract.
type Notifier struct {
	NotifierCaller     // Read-only binding to the contract
	NotifierTransactor // Write-only binding to the contract
	NotifierFilterer   // Log filterer for contract events
}

// NotifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type NotifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NotifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NotifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NotifierSession struct {
	Contract     *Notifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NotifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NotifierCallerSession struct {
	Contract *NotifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// NotifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NotifierTransactorSession struct {
	Contract     *NotifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NotifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type NotifierRaw struct {
	Contract *Notifier // Generic contract binding to access the raw methods on
}

// NotifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NotifierCallerRaw struct {
	Contract *NotifierCaller // Generic read-only contract binding to access the raw methods on
}

// NotifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NotifierTransactorRaw struct {
	Contract *NotifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNotifier creates a new instance of Notifier, bound to a specific deployed contract.
func NewNotifier(address common.Address, backend bind.ContractBackend) (*Notifier, error) {
	contract, err := bindNotifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Notifier{NotifierCaller: NotifierCaller{contract: contract}, NotifierTransactor: NotifierTransactor{contract: contract}, NotifierFilterer: NotifierFilterer{contract: contract}}, nil
}

// NewNotifierCaller creates a new read-only instance of Notifier, bound to a specific deployed contract.
func NewNotifierCaller(address common.Address, caller bind.ContractCaller) (*NotifierCaller, error) {
	contract, err := bindNotifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NotifierCaller{contract: contract}, nil
}

// NewNotifierTransactor creates a new write-only instance of Notifier, bound to a specific deployed contract.
func NewNotifierTransactor(address common.Address, transactor bind.ContractTransactor) (*NotifierTransactor, error) {
	contract, err := bindNotifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NotifierTransactor{contract: contract}, nil
}

// NewNotifierFilterer creates a new log filterer instance of Notifier, bound to a specific deployed contract.
func NewNotifierFilterer(address common.Address, filterer bind.ContractFilterer) (*NotifierFilterer, error) {
	contract, err := bindNotifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NotifierFilterer{contract: contract}, nil
}

// bindNotifier binds a generic wrapper to an already deployed contract.
func bindNotifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NotifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Notifier *NotifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Notifier.Contract.NotifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Notifier *NotifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notifier.Contract.NotifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Notifier *NotifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Notifier.Contract.NotifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Notifier *NotifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Notifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Notifier *NotifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Notifier *NotifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Notifier.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Notifier *NotifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Notifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Notifier *NotifierSession) Owner() (common.Address, error) {
	return _Notifier.Contract.Owner(&_Notifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Notifier *NotifierCallerSession) Owner() (common.Address, error) {
	return _Notifier.Contract.Owner(&_Notifier.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Notifier *NotifierCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Notifier.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Notifier *NotifierSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Notifier.Contract.OwnershipHandoverExpiresAt(&_Notifier.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Notifier *NotifierCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Notifier.Contract.OwnershipHandoverExpiresAt(&_Notifier.CallOpts, pendingOwner)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Notifier *NotifierTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notifier.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Notifier *NotifierSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Notifier.Contract.CancelOwnershipHandover(&_Notifier.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Notifier *NotifierTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Notifier.Contract.CancelOwnershipHandover(&_Notifier.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Notifier *NotifierTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _Notifier.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Notifier *NotifierSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Notifier.Contract.CompleteOwnershipHandover(&_Notifier.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Notifier *NotifierTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Notifier.Contract.CompleteOwnershipHandover(&_Notifier.TransactOpts, pendingOwner)
}

// NotifySubscribers is a paid mutator transaction binding the contract method 0x314e79ad.
//
// Solidity: function notifySubscribers(bytes32 _poolId, bytes4 _key, bytes _data) returns()
func (_Notifier *NotifierTransactor) NotifySubscribers(opts *bind.TransactOpts, _poolId [32]byte, _key [4]byte, _data []byte) (*types.Transaction, error) {
	return _Notifier.contract.Transact(opts, "notifySubscribers", _poolId, _key, _data)
}

// NotifySubscribers is a paid mutator transaction binding the contract method 0x314e79ad.
//
// Solidity: function notifySubscribers(bytes32 _poolId, bytes4 _key, bytes _data) returns()
func (_Notifier *NotifierSession) NotifySubscribers(_poolId [32]byte, _key [4]byte, _data []byte) (*types.Transaction, error) {
	return _Notifier.Contract.NotifySubscribers(&_Notifier.TransactOpts, _poolId, _key, _data)
}

// NotifySubscribers is a paid mutator transaction binding the contract method 0x314e79ad.
//
// Solidity: function notifySubscribers(bytes32 _poolId, bytes4 _key, bytes _data) returns()
func (_Notifier *NotifierTransactorSession) NotifySubscribers(_poolId [32]byte, _key [4]byte, _data []byte) (*types.Transaction, error) {
	return _Notifier.Contract.NotifySubscribers(&_Notifier.TransactOpts, _poolId, _key, _data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Notifier *NotifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Notifier *NotifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _Notifier.Contract.RenounceOwnership(&_Notifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Notifier *NotifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Notifier.Contract.RenounceOwnership(&_Notifier.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Notifier *NotifierTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notifier.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Notifier *NotifierSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Notifier.Contract.RequestOwnershipHandover(&_Notifier.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Notifier *NotifierTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Notifier.Contract.RequestOwnershipHandover(&_Notifier.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0xd5f9a397.
//
// Solidity: function subscribe(address _subscriber, bytes _data) returns()
func (_Notifier *NotifierTransactor) Subscribe(opts *bind.TransactOpts, _subscriber common.Address, _data []byte) (*types.Transaction, error) {
	return _Notifier.contract.Transact(opts, "subscribe", _subscriber, _data)
}

// Subscribe is a paid mutator transaction binding the contract method 0xd5f9a397.
//
// Solidity: function subscribe(address _subscriber, bytes _data) returns()
func (_Notifier *NotifierSession) Subscribe(_subscriber common.Address, _data []byte) (*types.Transaction, error) {
	return _Notifier.Contract.Subscribe(&_Notifier.TransactOpts, _subscriber, _data)
}

// Subscribe is a paid mutator transaction binding the contract method 0xd5f9a397.
//
// Solidity: function subscribe(address _subscriber, bytes _data) returns()
func (_Notifier *NotifierTransactorSession) Subscribe(_subscriber common.Address, _data []byte) (*types.Transaction, error) {
	return _Notifier.Contract.Subscribe(&_Notifier.TransactOpts, _subscriber, _data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Notifier *NotifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Notifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Notifier *NotifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Notifier.Contract.TransferOwnership(&_Notifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Notifier *NotifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Notifier.Contract.TransferOwnership(&_Notifier.TransactOpts, newOwner)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x7262561c.
//
// Solidity: function unsubscribe(address _subscriber) returns()
func (_Notifier *NotifierTransactor) Unsubscribe(opts *bind.TransactOpts, _subscriber common.Address) (*types.Transaction, error) {
	return _Notifier.contract.Transact(opts, "unsubscribe", _subscriber)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x7262561c.
//
// Solidity: function unsubscribe(address _subscriber) returns()
func (_Notifier *NotifierSession) Unsubscribe(_subscriber common.Address) (*types.Transaction, error) {
	return _Notifier.Contract.Unsubscribe(&_Notifier.TransactOpts, _subscriber)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x7262561c.
//
// Solidity: function unsubscribe(address _subscriber) returns()
func (_Notifier *NotifierTransactorSession) Unsubscribe(_subscriber common.Address) (*types.Transaction, error) {
	return _Notifier.Contract.Unsubscribe(&_Notifier.TransactOpts, _subscriber)
}

// NotifierOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the Notifier contract.
type NotifierOwnershipHandoverCanceledIterator struct {
	Event *NotifierOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *NotifierOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotifierOwnershipHandoverCanceled)
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
		it.Event = new(NotifierOwnershipHandoverCanceled)
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
func (it *NotifierOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotifierOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotifierOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Notifier contract.
type NotifierOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Notifier *NotifierFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*NotifierOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Notifier.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NotifierOwnershipHandoverCanceledIterator{contract: _Notifier.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Notifier *NotifierFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *NotifierOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Notifier.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotifierOwnershipHandoverCanceled)
				if err := _Notifier.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_Notifier *NotifierFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*NotifierOwnershipHandoverCanceled, error) {
	event := new(NotifierOwnershipHandoverCanceled)
	if err := _Notifier.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotifierOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the Notifier contract.
type NotifierOwnershipHandoverRequestedIterator struct {
	Event *NotifierOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *NotifierOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotifierOwnershipHandoverRequested)
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
		it.Event = new(NotifierOwnershipHandoverRequested)
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
func (it *NotifierOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotifierOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotifierOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Notifier contract.
type NotifierOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Notifier *NotifierFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*NotifierOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Notifier.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NotifierOwnershipHandoverRequestedIterator{contract: _Notifier.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Notifier *NotifierFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *NotifierOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Notifier.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotifierOwnershipHandoverRequested)
				if err := _Notifier.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_Notifier *NotifierFilterer) ParseOwnershipHandoverRequested(log types.Log) (*NotifierOwnershipHandoverRequested, error) {
	event := new(NotifierOwnershipHandoverRequested)
	if err := _Notifier.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Notifier contract.
type NotifierOwnershipTransferredIterator struct {
	Event *NotifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NotifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotifierOwnershipTransferred)
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
		it.Event = new(NotifierOwnershipTransferred)
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
func (it *NotifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotifierOwnershipTransferred represents a OwnershipTransferred event raised by the Notifier contract.
type NotifierOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Notifier *NotifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*NotifierOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Notifier.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NotifierOwnershipTransferredIterator{contract: _Notifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Notifier *NotifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NotifierOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Notifier.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotifierOwnershipTransferred)
				if err := _Notifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Notifier *NotifierFilterer) ParseOwnershipTransferred(log types.Log) (*NotifierOwnershipTransferred, error) {
	event := new(NotifierOwnershipTransferred)
	if err := _Notifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotifierSubscriptionIterator is returned from FilterSubscription and is used to iterate over the raw logs and unpacked data for Subscription events raised by the Notifier contract.
type NotifierSubscriptionIterator struct {
	Event *NotifierSubscription // Event containing the contract specifics and raw log

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
func (it *NotifierSubscriptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotifierSubscription)
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
		it.Event = new(NotifierSubscription)
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
func (it *NotifierSubscriptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotifierSubscriptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotifierSubscription represents a Subscription event raised by the Notifier contract.
type NotifierSubscription struct {
	Subscriber common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubscription is a free log retrieval operation binding the contract event 0xa6610d70b57486ac216657a87c0fce7f41758dd77d526daf7974f813a9b2a46a.
//
// Solidity: event Subscription(address _subscriber)
func (_Notifier *NotifierFilterer) FilterSubscription(opts *bind.FilterOpts) (*NotifierSubscriptionIterator, error) {

	logs, sub, err := _Notifier.contract.FilterLogs(opts, "Subscription")
	if err != nil {
		return nil, err
	}
	return &NotifierSubscriptionIterator{contract: _Notifier.contract, event: "Subscription", logs: logs, sub: sub}, nil
}

// WatchSubscription is a free log subscription operation binding the contract event 0xa6610d70b57486ac216657a87c0fce7f41758dd77d526daf7974f813a9b2a46a.
//
// Solidity: event Subscription(address _subscriber)
func (_Notifier *NotifierFilterer) WatchSubscription(opts *bind.WatchOpts, sink chan<- *NotifierSubscription) (event.Subscription, error) {

	logs, sub, err := _Notifier.contract.WatchLogs(opts, "Subscription")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotifierSubscription)
				if err := _Notifier.contract.UnpackLog(event, "Subscription", log); err != nil {
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

// ParseSubscription is a log parse operation binding the contract event 0xa6610d70b57486ac216657a87c0fce7f41758dd77d526daf7974f813a9b2a46a.
//
// Solidity: event Subscription(address _subscriber)
func (_Notifier *NotifierFilterer) ParseSubscription(log types.Log) (*NotifierSubscription, error) {
	event := new(NotifierSubscription)
	if err := _Notifier.contract.UnpackLog(event, "Subscription", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotifierUnsubscriptionIterator is returned from FilterUnsubscription and is used to iterate over the raw logs and unpacked data for Unsubscription events raised by the Notifier contract.
type NotifierUnsubscriptionIterator struct {
	Event *NotifierUnsubscription // Event containing the contract specifics and raw log

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
func (it *NotifierUnsubscriptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotifierUnsubscription)
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
		it.Event = new(NotifierUnsubscription)
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
func (it *NotifierUnsubscriptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotifierUnsubscriptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotifierUnsubscription represents a Unsubscription event raised by the Notifier contract.
type NotifierUnsubscription struct {
	Subscriber common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnsubscription is a free log retrieval operation binding the contract event 0x804bdbbf25e80c5635294ecf375fa16a531a4b7ffa416f74e121ae40bf4a00fe.
//
// Solidity: event Unsubscription(address _subscriber)
func (_Notifier *NotifierFilterer) FilterUnsubscription(opts *bind.FilterOpts) (*NotifierUnsubscriptionIterator, error) {

	logs, sub, err := _Notifier.contract.FilterLogs(opts, "Unsubscription")
	if err != nil {
		return nil, err
	}
	return &NotifierUnsubscriptionIterator{contract: _Notifier.contract, event: "Unsubscription", logs: logs, sub: sub}, nil
}

// WatchUnsubscription is a free log subscription operation binding the contract event 0x804bdbbf25e80c5635294ecf375fa16a531a4b7ffa416f74e121ae40bf4a00fe.
//
// Solidity: event Unsubscription(address _subscriber)
func (_Notifier *NotifierFilterer) WatchUnsubscription(opts *bind.WatchOpts, sink chan<- *NotifierUnsubscription) (event.Subscription, error) {

	logs, sub, err := _Notifier.contract.WatchLogs(opts, "Unsubscription")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotifierUnsubscription)
				if err := _Notifier.contract.UnpackLog(event, "Unsubscription", log); err != nil {
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

// ParseUnsubscription is a log parse operation binding the contract event 0x804bdbbf25e80c5635294ecf375fa16a531a4b7ffa416f74e121ae40bf4a00fe.
//
// Solidity: event Unsubscription(address _subscriber)
func (_Notifier *NotifierFilterer) ParseUnsubscription(log types.Log) (*NotifierUnsubscription, error) {
	event := new(NotifierUnsubscription)
	if err := _Notifier.contract.UnpackLog(event, "Unsubscription", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
