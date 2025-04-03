// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fee_exemptions

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

// FeeExemptionsFeeExemption is an auto generated low-level Go binding around an user-defined struct.
type FeeExemptionsFeeExemption struct {
	FlatFee *big.Int
	Enabled bool
}

// FeeExemptionsMetaData contains all meta data concerning the FeeExemptions contract.
var FeeExemptionsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_protocolOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"feeExemption\",\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeExemptions.FeeExemption\",\"components\":[{\"name\":\"flatFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFeeExemption\",\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setFeeExemption\",\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_flatFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"BeneficiaryFeeRemoved\",\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeneficiaryFeeSet\",\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_flatFee\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeExemptionInvalid\",\"inputs\":[{\"name\":\"_invalidFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"_maxFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoBeneficiaryExemption\",\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051610654380380610654833981016040819052602b91608e565b6032816037565b5060b9565b638b78c6d819805415605057630dc149f05f526004601cfd5b6001600160a01b03909116801560ff1b8117909155805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b5f60208284031215609d575f5ffd5b81516001600160a01b038116811460b2575f5ffd5b9392505050565b61058e806100c65f395ff3fe608060405260043610610080575f3560e01c8063176a57811461008457806325692962146100a5578063360507b6146100ad57806354d1f13d146100cc578063715018a6146100d45780638da5cb5b146100dc578063f04e283e14610105578063f2fde38b14610118578063f6a23d981461012b578063fee81cf4146101b7575b5f5ffd5b34801561008f575f5ffd5b506100a361009e3660046104e8565b6101f6565b005b6100a36102cc565b3480156100b8575f5ffd5b506100a36100c7366004610524565b610318565b6100a36103be565b6100a36103f7565b3480156100e7575f5ffd5b50638b78c6d819546040516100fc9190610544565b60405180910390f35b6100a3610113366004610524565b61040a565b6100a3610126366004610524565b610447565b348015610136575f5ffd5b50610195610145366004610524565b604080518082019091525f8082526020820152506001600160a01b03165f908152602081815260409182902082518084019093525462ffffff811683526301000000900460ff1615159082015290565b60408051825162ffffff168152602092830151151592810192909252016100fc565b3480156101c2575f5ffd5b506101e86101d1366004610524565b63389a75e1600c9081525f91909152602090205490565b6040519081526020016100fc565b6101fe61046d565b620f424062ffffff8216111561023c57604051636b4f872960e11b815262ffffff82166004820152620f424060248201526044015b60405180910390fd5b60408051808201825262ffffff838116808352600160208085019182526001600160a01b0388165f818152808352879020955186549351151563010000000263ffffffff19909416951694909417919091179093558351918252918101919091527f39ed56833ed68b7ad31bc19f04583da5de5c8e83ed9a66115f9066117b5c1920910160405180910390a15050565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b61032061046d565b6001600160a01b0381165f908152602081905260409020546301000000900460ff1661036157806040516338f3c55360e11b81526004016102339190610544565b6001600160a01b0381165f9081526020819052604090819020805463ffffffff19169055517f91891b9cc37faa20e065508e849de7e93346686dcf16e1f34b01f0fa87d01ccb906103b3908390610544565b60405180910390a150565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b6103ff61046d565b6104085f610487565b565b61041261046d565b63389a75e1600c52805f526020600c20805442111561043857636f5e88185f526004601cfd5b5f905561044481610487565b50565b61044f61046d565b8060601b61046457637448fbae5f526004601cfd5b61044481610487565b638b78c6d819543314610408576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3811560ff1b8217905550565b80356001600160a01b03811681146104e3575f5ffd5b919050565b5f5f604083850312156104f9575f5ffd5b610502836104cd565b9150602083013562ffffff81168114610519575f5ffd5b809150509250929050565b5f60208284031215610534575f5ffd5b61053d826104cd565b9392505050565b6001600160a01b039190911681526020019056fea264697066735822122078192e1c8245465cee8a5025c1346e7aa428505242576a1614a6cd5e6450db9c64736f6c634300081b0033",
}

// FeeExemptionsABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeExemptionsMetaData.ABI instead.
var FeeExemptionsABI = FeeExemptionsMetaData.ABI

// FeeExemptionsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeeExemptionsMetaData.Bin instead.
var FeeExemptionsBin = FeeExemptionsMetaData.Bin

// DeployFeeExemptions deploys a new Ethereum contract, binding an instance of FeeExemptions to it.
func DeployFeeExemptions(auth *bind.TransactOpts, backend bind.ContractBackend, _protocolOwner common.Address) (common.Address, *types.Transaction, *FeeExemptions, error) {
	parsed, err := FeeExemptionsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeExemptionsBin), backend, _protocolOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeExemptions{FeeExemptionsCaller: FeeExemptionsCaller{contract: contract}, FeeExemptionsTransactor: FeeExemptionsTransactor{contract: contract}, FeeExemptionsFilterer: FeeExemptionsFilterer{contract: contract}}, nil
}

// FeeExemptions is an auto generated Go binding around an Ethereum contract.
type FeeExemptions struct {
	FeeExemptionsCaller     // Read-only binding to the contract
	FeeExemptionsTransactor // Write-only binding to the contract
	FeeExemptionsFilterer   // Log filterer for contract events
}

// FeeExemptionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeExemptionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeExemptionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeExemptionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeExemptionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeExemptionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeExemptionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeExemptionsSession struct {
	Contract     *FeeExemptions    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeExemptionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeExemptionsCallerSession struct {
	Contract *FeeExemptionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FeeExemptionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeExemptionsTransactorSession struct {
	Contract     *FeeExemptionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FeeExemptionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeExemptionsRaw struct {
	Contract *FeeExemptions // Generic contract binding to access the raw methods on
}

// FeeExemptionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeExemptionsCallerRaw struct {
	Contract *FeeExemptionsCaller // Generic read-only contract binding to access the raw methods on
}

// FeeExemptionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeExemptionsTransactorRaw struct {
	Contract *FeeExemptionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeExemptions creates a new instance of FeeExemptions, bound to a specific deployed contract.
func NewFeeExemptions(address common.Address, backend bind.ContractBackend) (*FeeExemptions, error) {
	contract, err := bindFeeExemptions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeExemptions{FeeExemptionsCaller: FeeExemptionsCaller{contract: contract}, FeeExemptionsTransactor: FeeExemptionsTransactor{contract: contract}, FeeExemptionsFilterer: FeeExemptionsFilterer{contract: contract}}, nil
}

// NewFeeExemptionsCaller creates a new read-only instance of FeeExemptions, bound to a specific deployed contract.
func NewFeeExemptionsCaller(address common.Address, caller bind.ContractCaller) (*FeeExemptionsCaller, error) {
	contract, err := bindFeeExemptions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeExemptionsCaller{contract: contract}, nil
}

// NewFeeExemptionsTransactor creates a new write-only instance of FeeExemptions, bound to a specific deployed contract.
func NewFeeExemptionsTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeExemptionsTransactor, error) {
	contract, err := bindFeeExemptions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeExemptionsTransactor{contract: contract}, nil
}

// NewFeeExemptionsFilterer creates a new log filterer instance of FeeExemptions, bound to a specific deployed contract.
func NewFeeExemptionsFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeExemptionsFilterer, error) {
	contract, err := bindFeeExemptions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeExemptionsFilterer{contract: contract}, nil
}

// bindFeeExemptions binds a generic wrapper to an already deployed contract.
func bindFeeExemptions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeExemptionsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeExemptions *FeeExemptionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeExemptions.Contract.FeeExemptionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeExemptions *FeeExemptionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeExemptions.Contract.FeeExemptionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeExemptions *FeeExemptionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeExemptions.Contract.FeeExemptionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeExemptions *FeeExemptionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeExemptions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeExemptions *FeeExemptionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeExemptions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeExemptions *FeeExemptionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeExemptions.Contract.contract.Transact(opts, method, params...)
}

// FeeExemption is a free data retrieval call binding the contract method 0xf6a23d98.
//
// Solidity: function feeExemption(address _beneficiary) view returns((uint24,bool))
func (_FeeExemptions *FeeExemptionsCaller) FeeExemption(opts *bind.CallOpts, _beneficiary common.Address) (FeeExemptionsFeeExemption, error) {
	var out []interface{}
	err := _FeeExemptions.contract.Call(opts, &out, "feeExemption", _beneficiary)

	if err != nil {
		return *new(FeeExemptionsFeeExemption), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeExemptionsFeeExemption)).(*FeeExemptionsFeeExemption)

	return out0, err

}

// FeeExemption is a free data retrieval call binding the contract method 0xf6a23d98.
//
// Solidity: function feeExemption(address _beneficiary) view returns((uint24,bool))
func (_FeeExemptions *FeeExemptionsSession) FeeExemption(_beneficiary common.Address) (FeeExemptionsFeeExemption, error) {
	return _FeeExemptions.Contract.FeeExemption(&_FeeExemptions.CallOpts, _beneficiary)
}

// FeeExemption is a free data retrieval call binding the contract method 0xf6a23d98.
//
// Solidity: function feeExemption(address _beneficiary) view returns((uint24,bool))
func (_FeeExemptions *FeeExemptionsCallerSession) FeeExemption(_beneficiary common.Address) (FeeExemptionsFeeExemption, error) {
	return _FeeExemptions.Contract.FeeExemption(&_FeeExemptions.CallOpts, _beneficiary)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_FeeExemptions *FeeExemptionsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeExemptions.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_FeeExemptions *FeeExemptionsSession) Owner() (common.Address, error) {
	return _FeeExemptions.Contract.Owner(&_FeeExemptions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_FeeExemptions *FeeExemptionsCallerSession) Owner() (common.Address, error) {
	return _FeeExemptions.Contract.Owner(&_FeeExemptions.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_FeeExemptions *FeeExemptionsCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeExemptions.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_FeeExemptions *FeeExemptionsSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _FeeExemptions.Contract.OwnershipHandoverExpiresAt(&_FeeExemptions.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_FeeExemptions *FeeExemptionsCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _FeeExemptions.Contract.OwnershipHandoverExpiresAt(&_FeeExemptions.CallOpts, pendingOwner)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_FeeExemptions *FeeExemptionsTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeExemptions.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_FeeExemptions *FeeExemptionsSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _FeeExemptions.Contract.CancelOwnershipHandover(&_FeeExemptions.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_FeeExemptions *FeeExemptionsTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _FeeExemptions.Contract.CancelOwnershipHandover(&_FeeExemptions.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_FeeExemptions *FeeExemptionsTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _FeeExemptions.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_FeeExemptions *FeeExemptionsSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _FeeExemptions.Contract.CompleteOwnershipHandover(&_FeeExemptions.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_FeeExemptions *FeeExemptionsTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _FeeExemptions.Contract.CompleteOwnershipHandover(&_FeeExemptions.TransactOpts, pendingOwner)
}

// RemoveFeeExemption is a paid mutator transaction binding the contract method 0x360507b6.
//
// Solidity: function removeFeeExemption(address _beneficiary) returns()
func (_FeeExemptions *FeeExemptionsTransactor) RemoveFeeExemption(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _FeeExemptions.contract.Transact(opts, "removeFeeExemption", _beneficiary)
}

// RemoveFeeExemption is a paid mutator transaction binding the contract method 0x360507b6.
//
// Solidity: function removeFeeExemption(address _beneficiary) returns()
func (_FeeExemptions *FeeExemptionsSession) RemoveFeeExemption(_beneficiary common.Address) (*types.Transaction, error) {
	return _FeeExemptions.Contract.RemoveFeeExemption(&_FeeExemptions.TransactOpts, _beneficiary)
}

// RemoveFeeExemption is a paid mutator transaction binding the contract method 0x360507b6.
//
// Solidity: function removeFeeExemption(address _beneficiary) returns()
func (_FeeExemptions *FeeExemptionsTransactorSession) RemoveFeeExemption(_beneficiary common.Address) (*types.Transaction, error) {
	return _FeeExemptions.Contract.RemoveFeeExemption(&_FeeExemptions.TransactOpts, _beneficiary)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_FeeExemptions *FeeExemptionsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeExemptions.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_FeeExemptions *FeeExemptionsSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeExemptions.Contract.RenounceOwnership(&_FeeExemptions.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_FeeExemptions *FeeExemptionsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeExemptions.Contract.RenounceOwnership(&_FeeExemptions.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_FeeExemptions *FeeExemptionsTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeExemptions.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_FeeExemptions *FeeExemptionsSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _FeeExemptions.Contract.RequestOwnershipHandover(&_FeeExemptions.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_FeeExemptions *FeeExemptionsTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _FeeExemptions.Contract.RequestOwnershipHandover(&_FeeExemptions.TransactOpts)
}

// SetFeeExemption is a paid mutator transaction binding the contract method 0x176a5781.
//
// Solidity: function setFeeExemption(address _beneficiary, uint24 _flatFee) returns()
func (_FeeExemptions *FeeExemptionsTransactor) SetFeeExemption(opts *bind.TransactOpts, _beneficiary common.Address, _flatFee *big.Int) (*types.Transaction, error) {
	return _FeeExemptions.contract.Transact(opts, "setFeeExemption", _beneficiary, _flatFee)
}

// SetFeeExemption is a paid mutator transaction binding the contract method 0x176a5781.
//
// Solidity: function setFeeExemption(address _beneficiary, uint24 _flatFee) returns()
func (_FeeExemptions *FeeExemptionsSession) SetFeeExemption(_beneficiary common.Address, _flatFee *big.Int) (*types.Transaction, error) {
	return _FeeExemptions.Contract.SetFeeExemption(&_FeeExemptions.TransactOpts, _beneficiary, _flatFee)
}

// SetFeeExemption is a paid mutator transaction binding the contract method 0x176a5781.
//
// Solidity: function setFeeExemption(address _beneficiary, uint24 _flatFee) returns()
func (_FeeExemptions *FeeExemptionsTransactorSession) SetFeeExemption(_beneficiary common.Address, _flatFee *big.Int) (*types.Transaction, error) {
	return _FeeExemptions.Contract.SetFeeExemption(&_FeeExemptions.TransactOpts, _beneficiary, _flatFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_FeeExemptions *FeeExemptionsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeeExemptions.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_FeeExemptions *FeeExemptionsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeExemptions.Contract.TransferOwnership(&_FeeExemptions.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_FeeExemptions *FeeExemptionsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeExemptions.Contract.TransferOwnership(&_FeeExemptions.TransactOpts, newOwner)
}

// FeeExemptionsBeneficiaryFeeRemovedIterator is returned from FilterBeneficiaryFeeRemoved and is used to iterate over the raw logs and unpacked data for BeneficiaryFeeRemoved events raised by the FeeExemptions contract.
type FeeExemptionsBeneficiaryFeeRemovedIterator struct {
	Event *FeeExemptionsBeneficiaryFeeRemoved // Event containing the contract specifics and raw log

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
func (it *FeeExemptionsBeneficiaryFeeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeExemptionsBeneficiaryFeeRemoved)
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
		it.Event = new(FeeExemptionsBeneficiaryFeeRemoved)
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
func (it *FeeExemptionsBeneficiaryFeeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeExemptionsBeneficiaryFeeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeExemptionsBeneficiaryFeeRemoved represents a BeneficiaryFeeRemoved event raised by the FeeExemptions contract.
type FeeExemptionsBeneficiaryFeeRemoved struct {
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeneficiaryFeeRemoved is a free log retrieval operation binding the contract event 0x91891b9cc37faa20e065508e849de7e93346686dcf16e1f34b01f0fa87d01ccb.
//
// Solidity: event BeneficiaryFeeRemoved(address _beneficiary)
func (_FeeExemptions *FeeExemptionsFilterer) FilterBeneficiaryFeeRemoved(opts *bind.FilterOpts) (*FeeExemptionsBeneficiaryFeeRemovedIterator, error) {

	logs, sub, err := _FeeExemptions.contract.FilterLogs(opts, "BeneficiaryFeeRemoved")
	if err != nil {
		return nil, err
	}
	return &FeeExemptionsBeneficiaryFeeRemovedIterator{contract: _FeeExemptions.contract, event: "BeneficiaryFeeRemoved", logs: logs, sub: sub}, nil
}

// WatchBeneficiaryFeeRemoved is a free log subscription operation binding the contract event 0x91891b9cc37faa20e065508e849de7e93346686dcf16e1f34b01f0fa87d01ccb.
//
// Solidity: event BeneficiaryFeeRemoved(address _beneficiary)
func (_FeeExemptions *FeeExemptionsFilterer) WatchBeneficiaryFeeRemoved(opts *bind.WatchOpts, sink chan<- *FeeExemptionsBeneficiaryFeeRemoved) (event.Subscription, error) {

	logs, sub, err := _FeeExemptions.contract.WatchLogs(opts, "BeneficiaryFeeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeExemptionsBeneficiaryFeeRemoved)
				if err := _FeeExemptions.contract.UnpackLog(event, "BeneficiaryFeeRemoved", log); err != nil {
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

// ParseBeneficiaryFeeRemoved is a log parse operation binding the contract event 0x91891b9cc37faa20e065508e849de7e93346686dcf16e1f34b01f0fa87d01ccb.
//
// Solidity: event BeneficiaryFeeRemoved(address _beneficiary)
func (_FeeExemptions *FeeExemptionsFilterer) ParseBeneficiaryFeeRemoved(log types.Log) (*FeeExemptionsBeneficiaryFeeRemoved, error) {
	event := new(FeeExemptionsBeneficiaryFeeRemoved)
	if err := _FeeExemptions.contract.UnpackLog(event, "BeneficiaryFeeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeExemptionsBeneficiaryFeeSetIterator is returned from FilterBeneficiaryFeeSet and is used to iterate over the raw logs and unpacked data for BeneficiaryFeeSet events raised by the FeeExemptions contract.
type FeeExemptionsBeneficiaryFeeSetIterator struct {
	Event *FeeExemptionsBeneficiaryFeeSet // Event containing the contract specifics and raw log

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
func (it *FeeExemptionsBeneficiaryFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeExemptionsBeneficiaryFeeSet)
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
		it.Event = new(FeeExemptionsBeneficiaryFeeSet)
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
func (it *FeeExemptionsBeneficiaryFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeExemptionsBeneficiaryFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeExemptionsBeneficiaryFeeSet represents a BeneficiaryFeeSet event raised by the FeeExemptions contract.
type FeeExemptionsBeneficiaryFeeSet struct {
	Beneficiary common.Address
	FlatFee     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeneficiaryFeeSet is a free log retrieval operation binding the contract event 0x39ed56833ed68b7ad31bc19f04583da5de5c8e83ed9a66115f9066117b5c1920.
//
// Solidity: event BeneficiaryFeeSet(address _beneficiary, uint24 _flatFee)
func (_FeeExemptions *FeeExemptionsFilterer) FilterBeneficiaryFeeSet(opts *bind.FilterOpts) (*FeeExemptionsBeneficiaryFeeSetIterator, error) {

	logs, sub, err := _FeeExemptions.contract.FilterLogs(opts, "BeneficiaryFeeSet")
	if err != nil {
		return nil, err
	}
	return &FeeExemptionsBeneficiaryFeeSetIterator{contract: _FeeExemptions.contract, event: "BeneficiaryFeeSet", logs: logs, sub: sub}, nil
}

// WatchBeneficiaryFeeSet is a free log subscription operation binding the contract event 0x39ed56833ed68b7ad31bc19f04583da5de5c8e83ed9a66115f9066117b5c1920.
//
// Solidity: event BeneficiaryFeeSet(address _beneficiary, uint24 _flatFee)
func (_FeeExemptions *FeeExemptionsFilterer) WatchBeneficiaryFeeSet(opts *bind.WatchOpts, sink chan<- *FeeExemptionsBeneficiaryFeeSet) (event.Subscription, error) {

	logs, sub, err := _FeeExemptions.contract.WatchLogs(opts, "BeneficiaryFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeExemptionsBeneficiaryFeeSet)
				if err := _FeeExemptions.contract.UnpackLog(event, "BeneficiaryFeeSet", log); err != nil {
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

// ParseBeneficiaryFeeSet is a log parse operation binding the contract event 0x39ed56833ed68b7ad31bc19f04583da5de5c8e83ed9a66115f9066117b5c1920.
//
// Solidity: event BeneficiaryFeeSet(address _beneficiary, uint24 _flatFee)
func (_FeeExemptions *FeeExemptionsFilterer) ParseBeneficiaryFeeSet(log types.Log) (*FeeExemptionsBeneficiaryFeeSet, error) {
	event := new(FeeExemptionsBeneficiaryFeeSet)
	if err := _FeeExemptions.contract.UnpackLog(event, "BeneficiaryFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeExemptionsOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the FeeExemptions contract.
type FeeExemptionsOwnershipHandoverCanceledIterator struct {
	Event *FeeExemptionsOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *FeeExemptionsOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeExemptionsOwnershipHandoverCanceled)
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
		it.Event = new(FeeExemptionsOwnershipHandoverCanceled)
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
func (it *FeeExemptionsOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeExemptionsOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeExemptionsOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the FeeExemptions contract.
type FeeExemptionsOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_FeeExemptions *FeeExemptionsFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*FeeExemptionsOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _FeeExemptions.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeExemptionsOwnershipHandoverCanceledIterator{contract: _FeeExemptions.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_FeeExemptions *FeeExemptionsFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *FeeExemptionsOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _FeeExemptions.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeExemptionsOwnershipHandoverCanceled)
				if err := _FeeExemptions.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_FeeExemptions *FeeExemptionsFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*FeeExemptionsOwnershipHandoverCanceled, error) {
	event := new(FeeExemptionsOwnershipHandoverCanceled)
	if err := _FeeExemptions.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeExemptionsOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the FeeExemptions contract.
type FeeExemptionsOwnershipHandoverRequestedIterator struct {
	Event *FeeExemptionsOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *FeeExemptionsOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeExemptionsOwnershipHandoverRequested)
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
		it.Event = new(FeeExemptionsOwnershipHandoverRequested)
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
func (it *FeeExemptionsOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeExemptionsOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeExemptionsOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the FeeExemptions contract.
type FeeExemptionsOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_FeeExemptions *FeeExemptionsFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*FeeExemptionsOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _FeeExemptions.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeExemptionsOwnershipHandoverRequestedIterator{contract: _FeeExemptions.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_FeeExemptions *FeeExemptionsFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *FeeExemptionsOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _FeeExemptions.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeExemptionsOwnershipHandoverRequested)
				if err := _FeeExemptions.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_FeeExemptions *FeeExemptionsFilterer) ParseOwnershipHandoverRequested(log types.Log) (*FeeExemptionsOwnershipHandoverRequested, error) {
	event := new(FeeExemptionsOwnershipHandoverRequested)
	if err := _FeeExemptions.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeExemptionsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeeExemptions contract.
type FeeExemptionsOwnershipTransferredIterator struct {
	Event *FeeExemptionsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeeExemptionsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeExemptionsOwnershipTransferred)
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
		it.Event = new(FeeExemptionsOwnershipTransferred)
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
func (it *FeeExemptionsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeExemptionsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeExemptionsOwnershipTransferred represents a OwnershipTransferred event raised by the FeeExemptions contract.
type FeeExemptionsOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_FeeExemptions *FeeExemptionsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*FeeExemptionsOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeExemptions.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeExemptionsOwnershipTransferredIterator{contract: _FeeExemptions.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_FeeExemptions *FeeExemptionsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeExemptionsOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeExemptions.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeExemptionsOwnershipTransferred)
				if err := _FeeExemptions.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeeExemptions *FeeExemptionsFilterer) ParseOwnershipTransferred(log types.Log) (*FeeExemptionsOwnershipTransferred, error) {
	event := new(FeeExemptionsOwnershipTransferred)
	if err := _FeeExemptions.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
