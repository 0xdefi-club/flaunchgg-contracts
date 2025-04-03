// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package public_treasury_escrow

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

// PublicTreasuryEscrowMetaData contains all meta data concerning the PublicTreasuryEscrow contract.
var PublicTreasuryEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_flaunch\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeAction\",\"inputs\":[{\"name\":\"_action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"flaunch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractFlaunch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"originalOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipBurned\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reclaim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TreasuryEscrowed\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TreasuryOwnershipBurned\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TreasuryReclaimed\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOriginalOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipBurned\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161094e38038061094e83398101604081905261002e9161003f565b6001600160a01b031660805261006c565b5f6020828403121561004f575f5ffd5b81516001600160a01b0381168114610065575f5ffd5b9392505050565b6080516108a86100a65f395f818160bf0152818161014d0152818161021701528181610402015281816104c301526105f101526108a85ff3fe608060405234801561000f575f5ffd5b5060043610610081575f3560e01c806317d70f7c146100855780634e71d92d146100a0578063716344dc146100aa57806380e9071b146100b257806387211ceb146100ba5780638f39d13a146100ee578063a14be0cd14610101578063f1fffdcb14610125578063fe4b84df14610138575b5f5ffd5b61008d5f5481565b6040519081526020015b60405180910390f35b6100a861014b565b005b6100a86102e5565b6100a8610383565b6100e17f000000000000000000000000000000000000000000000000000000000000000081565b60405161009791906106cf565b6100a86100fc36600461070e565b6104aa565b60015461011590600160a01b900460ff1681565b6040519015158152602001610097565b6001546100e1906001600160a01b031681565b6100a86101463660046107d1565b610592565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663791b98bc6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101a7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101cb91906107e8565b6001600160a01b0316634c2d94c0600160149054906101000a900460ff166101fe576001546001600160a01b0316610288565b5f54604051630f38f95560e41b815260048101919091527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063f38f955090602401602060405180830381865afa158015610264573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061028891906107e8565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152600160248201526044015f604051808303815f87803b1580156102cd575f5ffd5b505af11580156102df573d5f5f3e3d5ffd5b50505050565b600154600160a01b900460ff161561031057604051637b5e483b60e11b815260040160405180910390fd5b6001546001600160a01b0316331461033b576040516323a7681d60e01b815260040160405180910390fd5b61034361014b565b6001805460ff60a01b1916600160a01b1790555f805460405190917f4625c0844db4984e389cd9976540fa8cfe06d31074f44d6d64a5a6a9eb0f472f91a2565b600154600160a01b900460ff16156103ae57604051637b5e483b60e11b815260040160405180910390fd5b6001546001600160a01b031633146103d9576040516323a7681d60e01b815260040160405180910390fd5b6103e161014b565b600180546001600160a01b03191690555f546040516323b872dd60e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316916323b872dd9161044291309133919060040161080a565b5f604051808303815f87803b158015610459575f5ffd5b505af115801561046b573d5f5f3e3d5ffd5b505050505f547f0bf8c148a6d29eb4919cbe7a77945c7564fdae47deee7233f4c5f47578b52366336040516104a091906106cf565b60405180910390a2565b5f54604051630f38f95560e41b815260048101919091527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063f38f955090602401602060405180830381865afa158015610510573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061053491906107e8565b6001600160a01b0316638f39d13a83836040518363ffffffff1660e01b815260040161056192919061082e565b5f604051808303815f87803b158015610578575f5ffd5b505af115801561058a573d5f5f3e3d5ffd5b505050505050565b63409feecd1980546003825580156105c85760018160011c14303b106105bf5763f92ee8a95f526004601cfd5b818160ff1b1b91505b505f829055600180546001600160a01b031916339081179091556040516323b872dd60e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316916323b872dd9161063191903090879060040161080a565b5f604051808303815f87803b158015610648575f5ffd5b505af115801561065a573d5f5f3e3d5ffd5b50505050817f1eabb06fa8a0515cb88851a6a9a5e489c988f2ad2bd177b2adeff7ada775c7ff3360405161068e91906106cf565b60405180910390a280156106cb576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b5050565b6001600160a01b0391909116815260200190565b6001600160a01b03811681146106f7575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561071f575f5ffd5b823561072a816106e3565b915060208301356001600160401b03811115610744575f5ffd5b8301601f81018513610754575f5ffd5b80356001600160401b0381111561076d5761076d6106fa565b604051601f8201601f19908116603f011681016001600160401b038111828210171561079b5761079b6106fa565b6040528181528282016020018710156107b2575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f602082840312156107e1575f5ffd5b5035919050565b5f602082840312156107f8575f5ffd5b8151610803816106e3565b9392505050565b6001600160a01b039384168152919092166020820152604081019190915260600190565b60018060a01b0383168152604060208201525f82518060408401528060208501606085015e5f606082850101526060601f19601f830116840101915050939250505056fea26469706673582212200c76eece0ab3dcd46486df10c75dacb47690c166e566669a7d41e4065de5d53564736f6c634300081b0033",
}

// PublicTreasuryEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicTreasuryEscrowMetaData.ABI instead.
var PublicTreasuryEscrowABI = PublicTreasuryEscrowMetaData.ABI

// PublicTreasuryEscrowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicTreasuryEscrowMetaData.Bin instead.
var PublicTreasuryEscrowBin = PublicTreasuryEscrowMetaData.Bin

// DeployPublicTreasuryEscrow deploys a new Ethereum contract, binding an instance of PublicTreasuryEscrow to it.
func DeployPublicTreasuryEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, _flaunch common.Address) (common.Address, *types.Transaction, *PublicTreasuryEscrow, error) {
	parsed, err := PublicTreasuryEscrowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicTreasuryEscrowBin), backend, _flaunch)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicTreasuryEscrow{PublicTreasuryEscrowCaller: PublicTreasuryEscrowCaller{contract: contract}, PublicTreasuryEscrowTransactor: PublicTreasuryEscrowTransactor{contract: contract}, PublicTreasuryEscrowFilterer: PublicTreasuryEscrowFilterer{contract: contract}}, nil
}

// PublicTreasuryEscrow is an auto generated Go binding around an Ethereum contract.
type PublicTreasuryEscrow struct {
	PublicTreasuryEscrowCaller     // Read-only binding to the contract
	PublicTreasuryEscrowTransactor // Write-only binding to the contract
	PublicTreasuryEscrowFilterer   // Log filterer for contract events
}

// PublicTreasuryEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicTreasuryEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicTreasuryEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicTreasuryEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicTreasuryEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicTreasuryEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicTreasuryEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicTreasuryEscrowSession struct {
	Contract     *PublicTreasuryEscrow // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PublicTreasuryEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicTreasuryEscrowCallerSession struct {
	Contract *PublicTreasuryEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PublicTreasuryEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicTreasuryEscrowTransactorSession struct {
	Contract     *PublicTreasuryEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PublicTreasuryEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicTreasuryEscrowRaw struct {
	Contract *PublicTreasuryEscrow // Generic contract binding to access the raw methods on
}

// PublicTreasuryEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicTreasuryEscrowCallerRaw struct {
	Contract *PublicTreasuryEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// PublicTreasuryEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicTreasuryEscrowTransactorRaw struct {
	Contract *PublicTreasuryEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicTreasuryEscrow creates a new instance of PublicTreasuryEscrow, bound to a specific deployed contract.
func NewPublicTreasuryEscrow(address common.Address, backend bind.ContractBackend) (*PublicTreasuryEscrow, error) {
	contract, err := bindPublicTreasuryEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicTreasuryEscrow{PublicTreasuryEscrowCaller: PublicTreasuryEscrowCaller{contract: contract}, PublicTreasuryEscrowTransactor: PublicTreasuryEscrowTransactor{contract: contract}, PublicTreasuryEscrowFilterer: PublicTreasuryEscrowFilterer{contract: contract}}, nil
}

// NewPublicTreasuryEscrowCaller creates a new read-only instance of PublicTreasuryEscrow, bound to a specific deployed contract.
func NewPublicTreasuryEscrowCaller(address common.Address, caller bind.ContractCaller) (*PublicTreasuryEscrowCaller, error) {
	contract, err := bindPublicTreasuryEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicTreasuryEscrowCaller{contract: contract}, nil
}

// NewPublicTreasuryEscrowTransactor creates a new write-only instance of PublicTreasuryEscrow, bound to a specific deployed contract.
func NewPublicTreasuryEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicTreasuryEscrowTransactor, error) {
	contract, err := bindPublicTreasuryEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicTreasuryEscrowTransactor{contract: contract}, nil
}

// NewPublicTreasuryEscrowFilterer creates a new log filterer instance of PublicTreasuryEscrow, bound to a specific deployed contract.
func NewPublicTreasuryEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicTreasuryEscrowFilterer, error) {
	contract, err := bindPublicTreasuryEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicTreasuryEscrowFilterer{contract: contract}, nil
}

// bindPublicTreasuryEscrow binds a generic wrapper to an already deployed contract.
func bindPublicTreasuryEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicTreasuryEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicTreasuryEscrow *PublicTreasuryEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicTreasuryEscrow.Contract.PublicTreasuryEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicTreasuryEscrow *PublicTreasuryEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.PublicTreasuryEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicTreasuryEscrow *PublicTreasuryEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.PublicTreasuryEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicTreasuryEscrow *PublicTreasuryEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicTreasuryEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.contract.Transact(opts, method, params...)
}

// Flaunch is a free data retrieval call binding the contract method 0x87211ceb.
//
// Solidity: function flaunch() view returns(address)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowCaller) Flaunch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicTreasuryEscrow.contract.Call(opts, &out, "flaunch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Flaunch is a free data retrieval call binding the contract method 0x87211ceb.
//
// Solidity: function flaunch() view returns(address)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowSession) Flaunch() (common.Address, error) {
	return _PublicTreasuryEscrow.Contract.Flaunch(&_PublicTreasuryEscrow.CallOpts)
}

// Flaunch is a free data retrieval call binding the contract method 0x87211ceb.
//
// Solidity: function flaunch() view returns(address)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowCallerSession) Flaunch() (common.Address, error) {
	return _PublicTreasuryEscrow.Contract.Flaunch(&_PublicTreasuryEscrow.CallOpts)
}

// OriginalOwner is a free data retrieval call binding the contract method 0xf1fffdcb.
//
// Solidity: function originalOwner() view returns(address)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowCaller) OriginalOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicTreasuryEscrow.contract.Call(opts, &out, "originalOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OriginalOwner is a free data retrieval call binding the contract method 0xf1fffdcb.
//
// Solidity: function originalOwner() view returns(address)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowSession) OriginalOwner() (common.Address, error) {
	return _PublicTreasuryEscrow.Contract.OriginalOwner(&_PublicTreasuryEscrow.CallOpts)
}

// OriginalOwner is a free data retrieval call binding the contract method 0xf1fffdcb.
//
// Solidity: function originalOwner() view returns(address)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowCallerSession) OriginalOwner() (common.Address, error) {
	return _PublicTreasuryEscrow.Contract.OriginalOwner(&_PublicTreasuryEscrow.CallOpts)
}

// OwnershipBurned is a free data retrieval call binding the contract method 0xa14be0cd.
//
// Solidity: function ownershipBurned() view returns(bool)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowCaller) OwnershipBurned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PublicTreasuryEscrow.contract.Call(opts, &out, "ownershipBurned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OwnershipBurned is a free data retrieval call binding the contract method 0xa14be0cd.
//
// Solidity: function ownershipBurned() view returns(bool)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowSession) OwnershipBurned() (bool, error) {
	return _PublicTreasuryEscrow.Contract.OwnershipBurned(&_PublicTreasuryEscrow.CallOpts)
}

// OwnershipBurned is a free data retrieval call binding the contract method 0xa14be0cd.
//
// Solidity: function ownershipBurned() view returns(bool)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowCallerSession) OwnershipBurned() (bool, error) {
	return _PublicTreasuryEscrow.Contract.OwnershipBurned(&_PublicTreasuryEscrow.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicTreasuryEscrow.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowSession) TokenId() (*big.Int, error) {
	return _PublicTreasuryEscrow.Contract.TokenId(&_PublicTreasuryEscrow.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowCallerSession) TokenId() (*big.Int, error) {
	return _PublicTreasuryEscrow.Contract.TokenId(&_PublicTreasuryEscrow.CallOpts)
}

// BurnOwnership is a paid mutator transaction binding the contract method 0x716344dc.
//
// Solidity: function burnOwnership() returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactor) BurnOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.contract.Transact(opts, "burnOwnership")
}

// BurnOwnership is a paid mutator transaction binding the contract method 0x716344dc.
//
// Solidity: function burnOwnership() returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowSession) BurnOwnership() (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.BurnOwnership(&_PublicTreasuryEscrow.TransactOpts)
}

// BurnOwnership is a paid mutator transaction binding the contract method 0x716344dc.
//
// Solidity: function burnOwnership() returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactorSession) BurnOwnership() (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.BurnOwnership(&_PublicTreasuryEscrow.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowSession) Claim() (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.Claim(&_PublicTreasuryEscrow.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactorSession) Claim() (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.Claim(&_PublicTreasuryEscrow.TransactOpts)
}

// ExecuteAction is a paid mutator transaction binding the contract method 0x8f39d13a.
//
// Solidity: function executeAction(address _action, bytes _data) returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactor) ExecuteAction(opts *bind.TransactOpts, _action common.Address, _data []byte) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.contract.Transact(opts, "executeAction", _action, _data)
}

// ExecuteAction is a paid mutator transaction binding the contract method 0x8f39d13a.
//
// Solidity: function executeAction(address _action, bytes _data) returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowSession) ExecuteAction(_action common.Address, _data []byte) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.ExecuteAction(&_PublicTreasuryEscrow.TransactOpts, _action, _data)
}

// ExecuteAction is a paid mutator transaction binding the contract method 0x8f39d13a.
//
// Solidity: function executeAction(address _action, bytes _data) returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactorSession) ExecuteAction(_action common.Address, _data []byte) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.ExecuteAction(&_PublicTreasuryEscrow.TransactOpts, _action, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _tokenId) returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactor) Initialize(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.contract.Transact(opts, "initialize", _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _tokenId) returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowSession) Initialize(_tokenId *big.Int) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.Initialize(&_PublicTreasuryEscrow.TransactOpts, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _tokenId) returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactorSession) Initialize(_tokenId *big.Int) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.Initialize(&_PublicTreasuryEscrow.TransactOpts, _tokenId)
}

// Reclaim is a paid mutator transaction binding the contract method 0x80e9071b.
//
// Solidity: function reclaim() returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactor) Reclaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicTreasuryEscrow.contract.Transact(opts, "reclaim")
}

// Reclaim is a paid mutator transaction binding the contract method 0x80e9071b.
//
// Solidity: function reclaim() returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowSession) Reclaim() (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.Reclaim(&_PublicTreasuryEscrow.TransactOpts)
}

// Reclaim is a paid mutator transaction binding the contract method 0x80e9071b.
//
// Solidity: function reclaim() returns()
func (_PublicTreasuryEscrow *PublicTreasuryEscrowTransactorSession) Reclaim() (*types.Transaction, error) {
	return _PublicTreasuryEscrow.Contract.Reclaim(&_PublicTreasuryEscrow.TransactOpts)
}

// PublicTreasuryEscrowInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PublicTreasuryEscrow contract.
type PublicTreasuryEscrowInitializedIterator struct {
	Event *PublicTreasuryEscrowInitialized // Event containing the contract specifics and raw log

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
func (it *PublicTreasuryEscrowInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicTreasuryEscrowInitialized)
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
		it.Event = new(PublicTreasuryEscrowInitialized)
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
func (it *PublicTreasuryEscrowInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicTreasuryEscrowInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicTreasuryEscrowInitialized represents a Initialized event raised by the PublicTreasuryEscrow contract.
type PublicTreasuryEscrowInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) FilterInitialized(opts *bind.FilterOpts) (*PublicTreasuryEscrowInitializedIterator, error) {

	logs, sub, err := _PublicTreasuryEscrow.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PublicTreasuryEscrowInitializedIterator{contract: _PublicTreasuryEscrow.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PublicTreasuryEscrowInitialized) (event.Subscription, error) {

	logs, sub, err := _PublicTreasuryEscrow.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicTreasuryEscrowInitialized)
				if err := _PublicTreasuryEscrow.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) ParseInitialized(log types.Log) (*PublicTreasuryEscrowInitialized, error) {
	event := new(PublicTreasuryEscrowInitialized)
	if err := _PublicTreasuryEscrow.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicTreasuryEscrowTreasuryEscrowedIterator is returned from FilterTreasuryEscrowed and is used to iterate over the raw logs and unpacked data for TreasuryEscrowed events raised by the PublicTreasuryEscrow contract.
type PublicTreasuryEscrowTreasuryEscrowedIterator struct {
	Event *PublicTreasuryEscrowTreasuryEscrowed // Event containing the contract specifics and raw log

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
func (it *PublicTreasuryEscrowTreasuryEscrowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicTreasuryEscrowTreasuryEscrowed)
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
		it.Event = new(PublicTreasuryEscrowTreasuryEscrowed)
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
func (it *PublicTreasuryEscrowTreasuryEscrowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicTreasuryEscrowTreasuryEscrowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicTreasuryEscrowTreasuryEscrowed represents a TreasuryEscrowed event raised by the PublicTreasuryEscrow contract.
type PublicTreasuryEscrowTreasuryEscrowed struct {
	TokenId *big.Int
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTreasuryEscrowed is a free log retrieval operation binding the contract event 0x1eabb06fa8a0515cb88851a6a9a5e489c988f2ad2bd177b2adeff7ada775c7ff.
//
// Solidity: event TreasuryEscrowed(uint256 indexed _tokenId, address _sender)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) FilterTreasuryEscrowed(opts *bind.FilterOpts, _tokenId []*big.Int) (*PublicTreasuryEscrowTreasuryEscrowedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _PublicTreasuryEscrow.contract.FilterLogs(opts, "TreasuryEscrowed", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PublicTreasuryEscrowTreasuryEscrowedIterator{contract: _PublicTreasuryEscrow.contract, event: "TreasuryEscrowed", logs: logs, sub: sub}, nil
}

// WatchTreasuryEscrowed is a free log subscription operation binding the contract event 0x1eabb06fa8a0515cb88851a6a9a5e489c988f2ad2bd177b2adeff7ada775c7ff.
//
// Solidity: event TreasuryEscrowed(uint256 indexed _tokenId, address _sender)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) WatchTreasuryEscrowed(opts *bind.WatchOpts, sink chan<- *PublicTreasuryEscrowTreasuryEscrowed, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _PublicTreasuryEscrow.contract.WatchLogs(opts, "TreasuryEscrowed", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicTreasuryEscrowTreasuryEscrowed)
				if err := _PublicTreasuryEscrow.contract.UnpackLog(event, "TreasuryEscrowed", log); err != nil {
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

// ParseTreasuryEscrowed is a log parse operation binding the contract event 0x1eabb06fa8a0515cb88851a6a9a5e489c988f2ad2bd177b2adeff7ada775c7ff.
//
// Solidity: event TreasuryEscrowed(uint256 indexed _tokenId, address _sender)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) ParseTreasuryEscrowed(log types.Log) (*PublicTreasuryEscrowTreasuryEscrowed, error) {
	event := new(PublicTreasuryEscrowTreasuryEscrowed)
	if err := _PublicTreasuryEscrow.contract.UnpackLog(event, "TreasuryEscrowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicTreasuryEscrowTreasuryOwnershipBurnedIterator is returned from FilterTreasuryOwnershipBurned and is used to iterate over the raw logs and unpacked data for TreasuryOwnershipBurned events raised by the PublicTreasuryEscrow contract.
type PublicTreasuryEscrowTreasuryOwnershipBurnedIterator struct {
	Event *PublicTreasuryEscrowTreasuryOwnershipBurned // Event containing the contract specifics and raw log

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
func (it *PublicTreasuryEscrowTreasuryOwnershipBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicTreasuryEscrowTreasuryOwnershipBurned)
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
		it.Event = new(PublicTreasuryEscrowTreasuryOwnershipBurned)
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
func (it *PublicTreasuryEscrowTreasuryOwnershipBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicTreasuryEscrowTreasuryOwnershipBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicTreasuryEscrowTreasuryOwnershipBurned represents a TreasuryOwnershipBurned event raised by the PublicTreasuryEscrow contract.
type PublicTreasuryEscrowTreasuryOwnershipBurned struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTreasuryOwnershipBurned is a free log retrieval operation binding the contract event 0x4625c0844db4984e389cd9976540fa8cfe06d31074f44d6d64a5a6a9eb0f472f.
//
// Solidity: event TreasuryOwnershipBurned(uint256 indexed _tokenId)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) FilterTreasuryOwnershipBurned(opts *bind.FilterOpts, _tokenId []*big.Int) (*PublicTreasuryEscrowTreasuryOwnershipBurnedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _PublicTreasuryEscrow.contract.FilterLogs(opts, "TreasuryOwnershipBurned", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PublicTreasuryEscrowTreasuryOwnershipBurnedIterator{contract: _PublicTreasuryEscrow.contract, event: "TreasuryOwnershipBurned", logs: logs, sub: sub}, nil
}

// WatchTreasuryOwnershipBurned is a free log subscription operation binding the contract event 0x4625c0844db4984e389cd9976540fa8cfe06d31074f44d6d64a5a6a9eb0f472f.
//
// Solidity: event TreasuryOwnershipBurned(uint256 indexed _tokenId)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) WatchTreasuryOwnershipBurned(opts *bind.WatchOpts, sink chan<- *PublicTreasuryEscrowTreasuryOwnershipBurned, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _PublicTreasuryEscrow.contract.WatchLogs(opts, "TreasuryOwnershipBurned", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicTreasuryEscrowTreasuryOwnershipBurned)
				if err := _PublicTreasuryEscrow.contract.UnpackLog(event, "TreasuryOwnershipBurned", log); err != nil {
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

// ParseTreasuryOwnershipBurned is a log parse operation binding the contract event 0x4625c0844db4984e389cd9976540fa8cfe06d31074f44d6d64a5a6a9eb0f472f.
//
// Solidity: event TreasuryOwnershipBurned(uint256 indexed _tokenId)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) ParseTreasuryOwnershipBurned(log types.Log) (*PublicTreasuryEscrowTreasuryOwnershipBurned, error) {
	event := new(PublicTreasuryEscrowTreasuryOwnershipBurned)
	if err := _PublicTreasuryEscrow.contract.UnpackLog(event, "TreasuryOwnershipBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicTreasuryEscrowTreasuryReclaimedIterator is returned from FilterTreasuryReclaimed and is used to iterate over the raw logs and unpacked data for TreasuryReclaimed events raised by the PublicTreasuryEscrow contract.
type PublicTreasuryEscrowTreasuryReclaimedIterator struct {
	Event *PublicTreasuryEscrowTreasuryReclaimed // Event containing the contract specifics and raw log

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
func (it *PublicTreasuryEscrowTreasuryReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicTreasuryEscrowTreasuryReclaimed)
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
		it.Event = new(PublicTreasuryEscrowTreasuryReclaimed)
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
func (it *PublicTreasuryEscrowTreasuryReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicTreasuryEscrowTreasuryReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicTreasuryEscrowTreasuryReclaimed represents a TreasuryReclaimed event raised by the PublicTreasuryEscrow contract.
type PublicTreasuryEscrowTreasuryReclaimed struct {
	TokenId *big.Int
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTreasuryReclaimed is a free log retrieval operation binding the contract event 0x0bf8c148a6d29eb4919cbe7a77945c7564fdae47deee7233f4c5f47578b52366.
//
// Solidity: event TreasuryReclaimed(uint256 indexed _tokenId, address _sender)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) FilterTreasuryReclaimed(opts *bind.FilterOpts, _tokenId []*big.Int) (*PublicTreasuryEscrowTreasuryReclaimedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _PublicTreasuryEscrow.contract.FilterLogs(opts, "TreasuryReclaimed", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PublicTreasuryEscrowTreasuryReclaimedIterator{contract: _PublicTreasuryEscrow.contract, event: "TreasuryReclaimed", logs: logs, sub: sub}, nil
}

// WatchTreasuryReclaimed is a free log subscription operation binding the contract event 0x0bf8c148a6d29eb4919cbe7a77945c7564fdae47deee7233f4c5f47578b52366.
//
// Solidity: event TreasuryReclaimed(uint256 indexed _tokenId, address _sender)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) WatchTreasuryReclaimed(opts *bind.WatchOpts, sink chan<- *PublicTreasuryEscrowTreasuryReclaimed, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _PublicTreasuryEscrow.contract.WatchLogs(opts, "TreasuryReclaimed", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicTreasuryEscrowTreasuryReclaimed)
				if err := _PublicTreasuryEscrow.contract.UnpackLog(event, "TreasuryReclaimed", log); err != nil {
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

// ParseTreasuryReclaimed is a log parse operation binding the contract event 0x0bf8c148a6d29eb4919cbe7a77945c7564fdae47deee7233f4c5f47578b52366.
//
// Solidity: event TreasuryReclaimed(uint256 indexed _tokenId, address _sender)
func (_PublicTreasuryEscrow *PublicTreasuryEscrowFilterer) ParseTreasuryReclaimed(log types.Log) (*PublicTreasuryEscrowTreasuryReclaimed, error) {
	event := new(PublicTreasuryEscrowTreasuryReclaimed)
	if err := _PublicTreasuryEscrow.contract.UnpackLog(event, "TreasuryReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
