// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool_swap

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

// PoolSwapMetaData contains all meta data concerning the PoolSwap contract.
var PoolSwapMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"_params\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"_params\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"_referrer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"delta_\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"unlockCallback\",\"inputs\":[{\"name\":\"rawData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b506040516115bd3803806115bd83398101604081905261002e9161003f565b6001600160a01b031660805261006c565b5f6020828403121561004f575f5ffd5b81516001600160a01b0381168114610065575f5ffd5b9392505050565b6080516114ed6100d05f395f818160790152818161010701528181610156015281816101900152818161026d01528181610371015281816103ab015281816107b6015281816108070152818161086c015281816108b7015261090a01526114ed5ff3fe60806040526004361061003e575f3560e01c80631e2817de14610042578063481c6a751461006857806391dd7346146100a8578063f8fff191146100d4575b5f5ffd5b610055610050366004610f88565b6100e7565b6040519081526020015b60405180910390f35b348015610073575f5ffd5b5061009b7f000000000000000000000000000000000000000000000000000000000000000081565b60405161005f9190610fbc565b3480156100b3575f5ffd5b506100c76100c2366004610fd0565b6100fa565b60405161005f919061106a565b6100556100e236600461107c565b610907565b5f6100f383835f610907565b9392505050565b6060336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610130575f5ffd5b5f61013d838501856110c4565b6020810151519091505f9061017e906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016903090610a02565b90505f6101c7308460200151602001517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610a029092919063ffffffff16565b9050811561021c5760405162461bcd60e51b815260206004820152601e60248201527f64656c74614265666f726530206973206e6f7420657175616c20746f2030000060448201526064015b60405180910390fd5b801561026a5760405162461bcd60e51b815260206004820152601e60248201527f64656c74614265666f726531206973206e6f7420657175616c20746f203000006044820152606401610213565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f3cd914c856020015186604001515f6001600160a01b031688606001516001600160a01b0316146102ea5787606001516040516020016102d69190610fbc565b6040516020818303038152906040526102fa565b60405180602001604052805f8152505b6040518463ffffffff1660e01b8152600401610318939291906111af565b6020604051808303815f875af1158015610334573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061035891906111de565b6020850151519091505f90610399906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016903090610a02565b90505f6103e2308760200151602001517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610a029092919063ffffffff16565b604087015151909150156105cf575f86604001516020015112156104eb578560400151602001518212156104685760405162461bcd60e51b815260206004820152604760248201525f5160206114585f395f51905f5260448201525f5160206114785f395f51905f526064820152661958da599a595960ca1b608482015260a401610213565b816104738460801d90565b600f0b146104935760405162461bcd60e51b8152600401610213906111f5565b5f8112156104e65760405162461bcd60e51b815260206004820152602d60248201525f5160206114385f395f51905f5260448201526c06f7220657175616c20746f203609c1b6064820152608401610213565b6107a6565b5f82131561053e5760405162461bcd60e51b815260206004820152602d60248201525f5160206114185f395f51905f5260448201526c657175616c20746f207a65726f60981b6064820152608401610213565b8061054984600f0b90565b600f0b146105695760405162461bcd60e51b815260040161021390611240565b8560400151602001518113156104e65760405162461bcd60e51b8152602060048201526044602482018190525f5160206114985f395f51905f52908201525f5160206113f85f395f51905f52606482015263199a595960e21b608482015260a401610213565b5f86604001516020015112156106c5578560400151602001518112156106475760405162461bcd60e51b815260206004820152604760248201525f5160206114385f395f51905f5260448201525f5160206114785f395f51905f526064820152661958da599a595960ca1b608482015260a401610213565b8061065284600f0b90565b600f0b146106725760405162461bcd60e51b815260040161021390611240565b5f8212156104e65760405162461bcd60e51b815260206004820152602d60248201525f5160206114585f395f51905f5260448201526c06f7220657175616c20746f203609c1b6064820152608401610213565b5f8113156107155760405162461bcd60e51b815260206004820152602a60248201525f5160206114985f395f51905f526044820152690657175616c20746f20360b41b6064820152608401610213565b816107208460801d90565b600f0b146107405760405162461bcd60e51b8152600401610213906111f5565b8560400151602001518213156107a65760405162461bcd60e51b8152602060048201526044602482018190525f5160206114185f395f51905f52908201525f5160206113f85f395f51905f52606482015263199a595960e21b608482015260a401610213565b5f8212156107f75785516107f7907f0000000000000000000000000000000000000000000000000000000000000000906107df8561128b565b60208a0151516001600160a01b03169291905f610a99565b5f81121561084a57855161084a907f0000000000000000000000000000000000000000000000000000000000000000906108308461128b565b6020808b015101516001600160a01b03169291905f610a99565b5f821315610893578551602087015151610893916001600160a01b03909116907f000000000000000000000000000000000000000000000000000000000000000090855f610d44565b5f8113156108de57855160208088015101516108de916001600160a01b03909116907f000000000000000000000000000000000000000000000000000000000000000090845f610d44565b604080516020810185905201604051602081830303815290604052965050505050505092915050565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166348c894916040518060800160405280336001600160a01b03168152602001878152602001868152602001856001600160a01b031681525060405160200161097a91906112b1565b6040516020818303038152906040526040518263ffffffff1660e01b81526004016109a5919061106a565b5f604051808303815f875af11580156109c0573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526109e79190810190611305565b8060200190518101906109fa91906111de565b949350505050565b5f5f6001600160a01b0384165f526001600160a01b03831660205260405f209050846001600160a01b031663f135baaa826040518263ffffffff1660e01b8152600401610a5191815260200190565b602060405180830381865afa158015610a6c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a9091906111de565b95945050505050565b8015610b1257836001600160a01b031663f5298aca84610ac1886001600160a01b0316610dee565b856040518463ffffffff1660e01b8152600401610ae093929190611397565b5f604051808303815f87803b158015610af7575f5ffd5b505af1158015610b09573d5f5f3e3d5ffd5b50505050610d3d565b6001600160a01b038516610b8957836001600160a01b03166311da60b4836040518263ffffffff1660e01b815260040160206040518083038185885af1158015610b5e573d5f5f3e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190610b8391906111de565b50610d3d565b604051632961046560e21b81526001600160a01b0385169063a584119490610bb5908890600401610fbc565b5f604051808303815f87803b158015610bcc575f5ffd5b505af1158015610bde573d5f5f3e3d5ffd5b505050506001600160a01b0383163014610c68576040516323b872dd60e01b81526001600160a01b038616906323b872dd90610c22908690889087906004016113b8565b6020604051808303815f875af1158015610c3e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c6291906113dc565b50610cda565b60405163a9059cbb60e01b81526001600160a01b0385811660048301526024820184905286169063a9059cbb906044016020604051808303815f875af1158015610cb4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cd891906113dc565b505b836001600160a01b03166311da60b46040518163ffffffff1660e01b81526004016020604051808303815f875af1158015610d17573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d3b91906111de565b505b5050505050565b80610d7957604051630b0d9c0960e01b81526001600160a01b03851690630b0d9c0990610ae0908890879087906004016113b8565b836001600160a01b031663156e29f684610d9b886001600160a01b0316610dee565b856040518463ffffffff1660e01b8152600401610dba93929190611397565b5f604051808303815f87803b158015610dd1575f5ffd5b505af1158015610de3573d5f5f3e3d5ffd5b505050505050505050565b6001600160a01b031690565b634e487b7160e01b5f52604160045260245ffd5b60405160a081016001600160401b0381118282101715610e3057610e30610dfa565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610e5e57610e5e610dfa565b604052919050565b6001600160a01b0381168114610e7a575f5ffd5b50565b8035610e8881610e66565b919050565b5f60a08284031215610e9d575f5ffd5b610ea5610e0e565b90508135610eb281610e66565b81526020820135610ec281610e66565b6020820152604082013562ffffff81168114610edc575f5ffd5b60408201526060820135600281900b8114610ef5575f5ffd5b6060820152610f0660808301610e7d565b608082015292915050565b8015158114610e7a575f5ffd5b5f60608284031215610f2e575f5ffd5b604051606081016001600160401b0381118282101715610f5057610f50610dfa565b6040529050808235610f6181610f11565b8152602083810135908201526040830135610f7b81610e66565b6040919091015292915050565b5f5f6101008385031215610f9a575f5ffd5b610fa48484610e8d565b9150610fb38460a08501610f1e565b90509250929050565b6001600160a01b0391909116815260200190565b5f5f60208385031215610fe1575f5ffd5b82356001600160401b03811115610ff6575f5ffd5b8301601f81018513611006575f5ffd5b80356001600160401b0381111561101b575f5ffd5b85602082840101111561102c575f5ffd5b6020919091019590945092505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6100f3602083018461103c565b5f5f5f610120848603121561108f575f5ffd5b6110998585610e8d565b92506110a88560a08601610f1e565b91506101008401356110b981610e66565b809150509250925092565b5f6101408284031280156110d6575f5ffd5b50604051608081016001600160401b03811182821017156110f9576110f9610dfa565b604052823561110781610e66565b81526111168460208501610e8d565b60208201526111288460c08501610f1e565b604082015261012083013561113c81610e66565b60608201529392505050565b80516001600160a01b03908116835260208083015182169084015260408083015162ffffff169084015260608083015160020b9084015260809182015116910152565b805115158252602080820151908301526040908101516001600160a01b0316910152565b6111b98185611148565b6111c660a082018461118b565b6101206101008201525f610a9061012083018461103c565b5f602082840312156111ee575f5ffd5b5051919050565b6020808252602b908201527f64656c74612e616d6f756e74302829206973206e6f7420657175616c20746f2060408201526a064656c74614166746572360ac1b606082015260800190565b6020808252602b908201527f64656c74612e616d6f756e74312829206973206e6f7420657175616c20746f2060408201526a64656c746141667465723160a81b606082015260800190565b5f600160ff1b82016112ab57634e487b7160e01b5f52601160045260245ffd5b505f0390565b81516001600160a01b031681526020808301516101408301916112d690840182611148565b5060408301516112e960c084018261118b565b50606092909201516001600160a01b0316610120919091015290565b5f60208284031215611315575f5ffd5b81516001600160401b0381111561132a575f5ffd5b8201601f8101841361133a575f5ffd5b80516001600160401b0381111561135357611353610dfa565b611366601f8201601f1916602001610e36565b81815285602083850101111561137a575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b6001600160a01b039390931683526020830191909152604082015260600190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b5f602082840312156113ec575f5ffd5b81516100f381610f1156fe657175616c20746f20646174612e706172616d732e616d6f756e74537065636964656c7461416674657230206973206e6f74206c657373207468616e206f722064656c7461416674657231206973206e6f742067726561746572207468616e2064656c7461416674657230206973206e6f742067726561746572207468616e206f7220657175616c20746f20646174612e706172616d732e616d6f756e74537064656c7461416674657231206973206e6f74206c657373207468616e206f7220a264697066735822122064cb7f343306bb14c88b931f9a26275dafd1e8a62ec3cd0c8ceb0bd00367148264736f6c634300081b0033",
}

// PoolSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolSwapMetaData.ABI instead.
var PoolSwapABI = PoolSwapMetaData.ABI

// PoolSwapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolSwapMetaData.Bin instead.
var PoolSwapBin = PoolSwapMetaData.Bin

// DeployPoolSwap deploys a new Ethereum contract, binding an instance of PoolSwap to it.
func DeployPoolSwap(auth *bind.TransactOpts, backend bind.ContractBackend, _manager common.Address) (common.Address, *types.Transaction, *PoolSwap, error) {
	parsed, err := PoolSwapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolSwapBin), backend, _manager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolSwap{PoolSwapCaller: PoolSwapCaller{contract: contract}, PoolSwapTransactor: PoolSwapTransactor{contract: contract}, PoolSwapFilterer: PoolSwapFilterer{contract: contract}}, nil
}

// PoolSwap is an auto generated Go binding around an Ethereum contract.
type PoolSwap struct {
	PoolSwapCaller     // Read-only binding to the contract
	PoolSwapTransactor // Write-only binding to the contract
	PoolSwapFilterer   // Log filterer for contract events
}

// PoolSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSwapSession struct {
	Contract     *PoolSwap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolSwapCallerSession struct {
	Contract *PoolSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PoolSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolSwapTransactorSession struct {
	Contract     *PoolSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PoolSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolSwapRaw struct {
	Contract *PoolSwap // Generic contract binding to access the raw methods on
}

// PoolSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolSwapCallerRaw struct {
	Contract *PoolSwapCaller // Generic read-only contract binding to access the raw methods on
}

// PoolSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolSwapTransactorRaw struct {
	Contract *PoolSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolSwap creates a new instance of PoolSwap, bound to a specific deployed contract.
func NewPoolSwap(address common.Address, backend bind.ContractBackend) (*PoolSwap, error) {
	contract, err := bindPoolSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolSwap{PoolSwapCaller: PoolSwapCaller{contract: contract}, PoolSwapTransactor: PoolSwapTransactor{contract: contract}, PoolSwapFilterer: PoolSwapFilterer{contract: contract}}, nil
}

// NewPoolSwapCaller creates a new read-only instance of PoolSwap, bound to a specific deployed contract.
func NewPoolSwapCaller(address common.Address, caller bind.ContractCaller) (*PoolSwapCaller, error) {
	contract, err := bindPoolSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolSwapCaller{contract: contract}, nil
}

// NewPoolSwapTransactor creates a new write-only instance of PoolSwap, bound to a specific deployed contract.
func NewPoolSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolSwapTransactor, error) {
	contract, err := bindPoolSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolSwapTransactor{contract: contract}, nil
}

// NewPoolSwapFilterer creates a new log filterer instance of PoolSwap, bound to a specific deployed contract.
func NewPoolSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolSwapFilterer, error) {
	contract, err := bindPoolSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolSwapFilterer{contract: contract}, nil
}

// bindPoolSwap binds a generic wrapper to an already deployed contract.
func bindPoolSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolSwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolSwap *PoolSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolSwap.Contract.PoolSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolSwap *PoolSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolSwap.Contract.PoolSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolSwap *PoolSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolSwap.Contract.PoolSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolSwap *PoolSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolSwap *PoolSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolSwap *PoolSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolSwap.Contract.contract.Transact(opts, method, params...)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_PoolSwap *PoolSwapCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolSwap.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_PoolSwap *PoolSwapSession) Manager() (common.Address, error) {
	return _PoolSwap.Contract.Manager(&_PoolSwap.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_PoolSwap *PoolSwapCallerSession) Manager() (common.Address, error) {
	return _PoolSwap.Contract.Manager(&_PoolSwap.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x1e2817de.
//
// Solidity: function swap((address,address,uint24,int24,address) _key, (bool,int256,uint160) _params) payable returns(int256)
func (_PoolSwap *PoolSwapTransactor) Swap(opts *bind.TransactOpts, _key PoolKey, _params IPoolManagerSwapParams) (*types.Transaction, error) {
	return _PoolSwap.contract.Transact(opts, "swap", _key, _params)
}

// Swap is a paid mutator transaction binding the contract method 0x1e2817de.
//
// Solidity: function swap((address,address,uint24,int24,address) _key, (bool,int256,uint160) _params) payable returns(int256)
func (_PoolSwap *PoolSwapSession) Swap(_key PoolKey, _params IPoolManagerSwapParams) (*types.Transaction, error) {
	return _PoolSwap.Contract.Swap(&_PoolSwap.TransactOpts, _key, _params)
}

// Swap is a paid mutator transaction binding the contract method 0x1e2817de.
//
// Solidity: function swap((address,address,uint24,int24,address) _key, (bool,int256,uint160) _params) payable returns(int256)
func (_PoolSwap *PoolSwapTransactorSession) Swap(_key PoolKey, _params IPoolManagerSwapParams) (*types.Transaction, error) {
	return _PoolSwap.Contract.Swap(&_PoolSwap.TransactOpts, _key, _params)
}

// Swap0 is a paid mutator transaction binding the contract method 0xf8fff191.
//
// Solidity: function swap((address,address,uint24,int24,address) _key, (bool,int256,uint160) _params, address _referrer) payable returns(int256 delta_)
func (_PoolSwap *PoolSwapTransactor) Swap0(opts *bind.TransactOpts, _key PoolKey, _params IPoolManagerSwapParams, _referrer common.Address) (*types.Transaction, error) {
	return _PoolSwap.contract.Transact(opts, "swap0", _key, _params, _referrer)
}

// Swap0 is a paid mutator transaction binding the contract method 0xf8fff191.
//
// Solidity: function swap((address,address,uint24,int24,address) _key, (bool,int256,uint160) _params, address _referrer) payable returns(int256 delta_)
func (_PoolSwap *PoolSwapSession) Swap0(_key PoolKey, _params IPoolManagerSwapParams, _referrer common.Address) (*types.Transaction, error) {
	return _PoolSwap.Contract.Swap0(&_PoolSwap.TransactOpts, _key, _params, _referrer)
}

// Swap0 is a paid mutator transaction binding the contract method 0xf8fff191.
//
// Solidity: function swap((address,address,uint24,int24,address) _key, (bool,int256,uint160) _params, address _referrer) payable returns(int256 delta_)
func (_PoolSwap *PoolSwapTransactorSession) Swap0(_key PoolKey, _params IPoolManagerSwapParams, _referrer common.Address) (*types.Transaction, error) {
	return _PoolSwap.Contract.Swap0(&_PoolSwap.TransactOpts, _key, _params, _referrer)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes rawData) returns(bytes)
func (_PoolSwap *PoolSwapTransactor) UnlockCallback(opts *bind.TransactOpts, rawData []byte) (*types.Transaction, error) {
	return _PoolSwap.contract.Transact(opts, "unlockCallback", rawData)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes rawData) returns(bytes)
func (_PoolSwap *PoolSwapSession) UnlockCallback(rawData []byte) (*types.Transaction, error) {
	return _PoolSwap.Contract.UnlockCallback(&_PoolSwap.TransactOpts, rawData)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes rawData) returns(bytes)
func (_PoolSwap *PoolSwapTransactorSession) UnlockCallback(rawData []byte) (*types.Transaction, error) {
	return _PoolSwap.Contract.UnlockCallback(&_PoolSwap.TransactOpts, rawData)
}
