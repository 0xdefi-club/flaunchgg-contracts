// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market_capped_price

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

// PoolKey is an auto generated low-level Go binding around an user-defined struct.
type PoolKey struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

// MarketCappedPriceMetaData contains all meta data concerning the MarketCappedPrice contract.
var MarketCappedPriceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_protocolOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ethToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdcToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MINIMUM_USDC_MARKET_CAP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ethToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"Currency\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFlaunchingFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketCap\",\"inputs\":[{\"name\":\"_initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSqrtPriceX96\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_flipped\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"sqrtPriceX96_\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setPool\",\"inputs\":[{\"name\":\"_poolKey\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"usdcToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"Currency\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdcToken0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketCapTooSmall\",\"inputs\":[{\"name\":\"_usdcMarketCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_usdcMarketCapMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051610ded380380610ded83398101604081905261002e916100cc565b6001600160a01b0380841660805282811660a052811660c05261005084610059565b5050505061011d565b638b78c6d81980541561007357630dc149f05f526004601cfd5b6001600160a01b03909116801560ff1b8117909155805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b80516001600160a01b03811681146100c7575f5ffd5b919050565b5f5f5f5f608085870312156100df575f5ffd5b6100e8856100b1565b93506100f6602086016100b1565b9250610104604086016100b1565b9150610112606086016100b1565b905092959194509250565b60805160a05160c051610c9161015c5f395f818161011001526104d401525f81816101ec015261051401525f818161025601526103980152610c915ff3fe6080604052600436106100cd575f3560e01c80630c7324d3146100d157806311eac855146100ff578063256929621461014a5780632a81c4f814610154578063355c7371146101815780633e0dc34e1461019857806342d8c2c0146101ac57806354d1f13d146101cb578063715018a6146101d35780637bf1a627146101db5780638207cb481461020e5780638da5cb5b1461022d578063dc4c90d314610245578063f04e283e14610278578063f2fde38b1461028b578063fc87f3f61461029e578063fee81cf4146102bd575b5f5ffd5b3480156100dc575f5ffd5b506001546100ea9060ff1681565b60405190151581526020015b60405180910390f35b34801561010a575f5ffd5b506101327f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100f6565b6101526102ee565b005b34801561015f575f5ffd5b5061017361016e3660046109b6565b61033a565b6040519081526020016100f6565b34801561018c575f5ffd5b50610173633b9aca0081565b3480156101a3575f5ffd5b506101735f5481565b3480156101b7575f5ffd5b506101736101c6366004610a08565b61044c565b61015261046e565b6101526104a7565b3480156101e6575f5ffd5b506101327f000000000000000000000000000000000000000000000000000000000000000081565b348015610219575f5ffd5b50610152610228366004610a58565b6104ba565b348015610238575f5ffd5b50638b78c6d81954610132565b348015610250575f5ffd5b506101327f000000000000000000000000000000000000000000000000000000000000000081565b610152610286366004610a71565b6105e8565b610152610299366004610a71565b610625565b3480156102a9575f5ffd5b506101326102b8366004610a8c565b61064b565b3480156102c8575f5ffd5b506101736102d7366004610a71565b63389a75e1600c9081525f91909152602090205490565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b5f8061034883850185610b24565b9050633b9aca00815f01511015610387578051604051630292f1cf60e01b81526004810191909152633b9aca0060248201526044015b60405180910390fd5b5f80546103be906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690610676565b50505090505f6103e5826001600160a01b0316836001600160a01b0316600160601b610728565b6001549091505f9060ff166104105761040b670de0b6b3a764000083600160601b610728565b610427565b610427670de0b6b3a7640000600160601b84610728565b905061043f845f0151670de0b6b3a764000083610728565b9450505050505b92915050565b5f6103e861045a848461033a565b6104649190610b71565b90505b9392505050565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b6104af6107c4565b6104b85f6107de565b565b6104c26107c4565b6104f86104d26020830183610a71565b7f0000000000000000000000000000000000000000000000000000000000000000610824565b801561053857506105386105126040830160208401610a71565b7f0000000000000000000000000000000000000000000000000000000000000000610824565b8061056857506105516104d26040830160208401610a71565b801561056857506105686105126020830183610a71565b6105a95760405162461bcd60e51b815260206004820152601260248201527124b73b30b634b2103a37b5b2b7103830b4b960711b604482015260640161037e565b6105c26105bb36839003830183610ba0565b60a0902090565b5f556105d46104d26020830183610a71565b6001805460ff191691151591909117905550565b6105f06107c4565b63389a75e1600c52805f526020600c20805442111561061657636f5e88185f526004601cfd5b5f9055610622816107de565b50565b61062d6107c4565b8060601b61064257637448fbae5f526004601cfd5b610622816107de565b5f61066d610659848461033a565b680a18f07d736b90be55601d1b8615610835565b95945050505050565b5f5f5f5f5f610684866108d0565b604051631e2eaeaf60e01b8152600481018290529091505f906001600160a01b03891690631e2eaeaf90602401602060405180830381865afa1580156106cc573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106f09190610c25565b90506001600160a01b03811695508060a01c60020b945062ffffff8160b81c16935062ffffff8160d01c169250505092959194509250565b5f838302815f1985870982811083820303915050808411610747575f5ffd5b805f0361075957508290049050610467565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b638b78c6d8195433146104b8576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3811560ff1b8217905550565b6001600160a01b0390811691161490565b5f5f8411801561084457505f83115b61089a5760405162461bcd60e51b815260206004820152602160248201527f416d6f756e7473206d7573742062652067726561746572207468616e207a65726044820152606f60f81b606482015260840161037e565b81156108be576108b76108b284600160c01b87610728565b61090c565b9050610467565b6104646108b285600160c01b86610728565b6040515f906108ef908390600690602001918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b5f815f0361091b57505f919050565b5f6002610929846001610c3c565b6109339190610b71565b90508291505b8181101561096c579050806002816109518186610b71565b61095b9190610c3c565b6109659190610b71565b9050610939565b50919050565b5f5f83601f840112610982575f5ffd5b5081356001600160401b03811115610998575f5ffd5b6020830191508360208285010111156109af575f5ffd5b9250929050565b5f5f602083850312156109c7575f5ffd5b82356001600160401b038111156109dc575f5ffd5b6109e885828601610972565b90969095509350505050565b6001600160a01b0381168114610622575f5ffd5b5f5f5f60408486031215610a1a575f5ffd5b8335610a25816109f4565b925060208401356001600160401b03811115610a3f575f5ffd5b610a4b86828701610972565b9497909650939450505050565b5f60a0828403128015610a69575f5ffd5b509092915050565b5f60208284031215610a81575f5ffd5b8135610467816109f4565b5f5f5f5f60608587031215610a9f575f5ffd5b8435610aaa816109f4565b935060208501358015158114610abe575f5ffd5b925060408501356001600160401b03811115610ad8575f5ffd5b610ae487828801610972565b95989497509550505050565b60405160a081016001600160401b0381118282101715610b1e57634e487b7160e01b5f52604160045260245ffd5b60405290565b5f6020828403128015610b35575f5ffd5b50604051602081016001600160401b0381118282101715610b6457634e487b7160e01b5f52604160045260245ffd5b6040529135825250919050565b5f82610b8b57634e487b7160e01b5f52601260045260245ffd5b500490565b8035610b9b816109f4565b919050565b5f60a0828403128015610bb1575f5ffd5b50610bba610af0565b8235610bc5816109f4565b81526020830135610bd5816109f4565b6020820152604083013562ffffff81168114610bef575f5ffd5b60408201526060830135600281900b8114610c08575f5ffd5b6060820152610c1960808401610b90565b60808201529392505050565b5f60208284031215610c35575f5ffd5b5051919050565b8082018082111561044657634e487b7160e01b5f52601160045260245ffdfea264697066735822122091b29b26077de99b85ab448d8818c98c4eec1195bf4c0f1428e76053ece56dca64736f6c634300081b0033",
}

// MarketCappedPriceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketCappedPriceMetaData.ABI instead.
var MarketCappedPriceABI = MarketCappedPriceMetaData.ABI

// MarketCappedPriceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketCappedPriceMetaData.Bin instead.
var MarketCappedPriceBin = MarketCappedPriceMetaData.Bin

// DeployMarketCappedPrice deploys a new Ethereum contract, binding an instance of MarketCappedPrice to it.
func DeployMarketCappedPrice(auth *bind.TransactOpts, backend bind.ContractBackend, _protocolOwner common.Address, _poolManager common.Address, _ethToken common.Address, _usdcToken common.Address) (common.Address, *types.Transaction, *MarketCappedPrice, error) {
	parsed, err := MarketCappedPriceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketCappedPriceBin), backend, _protocolOwner, _poolManager, _ethToken, _usdcToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketCappedPrice{MarketCappedPriceCaller: MarketCappedPriceCaller{contract: contract}, MarketCappedPriceTransactor: MarketCappedPriceTransactor{contract: contract}, MarketCappedPriceFilterer: MarketCappedPriceFilterer{contract: contract}}, nil
}

// MarketCappedPrice is an auto generated Go binding around an Ethereum contract.
type MarketCappedPrice struct {
	MarketCappedPriceCaller     // Read-only binding to the contract
	MarketCappedPriceTransactor // Write-only binding to the contract
	MarketCappedPriceFilterer   // Log filterer for contract events
}

// MarketCappedPriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCappedPriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketCappedPriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketCappedPriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketCappedPriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketCappedPriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketCappedPriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketCappedPriceSession struct {
	Contract     *MarketCappedPrice // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MarketCappedPriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCappedPriceCallerSession struct {
	Contract *MarketCappedPriceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MarketCappedPriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketCappedPriceTransactorSession struct {
	Contract     *MarketCappedPriceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MarketCappedPriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketCappedPriceRaw struct {
	Contract *MarketCappedPrice // Generic contract binding to access the raw methods on
}

// MarketCappedPriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCappedPriceCallerRaw struct {
	Contract *MarketCappedPriceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketCappedPriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketCappedPriceTransactorRaw struct {
	Contract *MarketCappedPriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketCappedPrice creates a new instance of MarketCappedPrice, bound to a specific deployed contract.
func NewMarketCappedPrice(address common.Address, backend bind.ContractBackend) (*MarketCappedPrice, error) {
	contract, err := bindMarketCappedPrice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPrice{MarketCappedPriceCaller: MarketCappedPriceCaller{contract: contract}, MarketCappedPriceTransactor: MarketCappedPriceTransactor{contract: contract}, MarketCappedPriceFilterer: MarketCappedPriceFilterer{contract: contract}}, nil
}

// NewMarketCappedPriceCaller creates a new read-only instance of MarketCappedPrice, bound to a specific deployed contract.
func NewMarketCappedPriceCaller(address common.Address, caller bind.ContractCaller) (*MarketCappedPriceCaller, error) {
	contract, err := bindMarketCappedPrice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceCaller{contract: contract}, nil
}

// NewMarketCappedPriceTransactor creates a new write-only instance of MarketCappedPrice, bound to a specific deployed contract.
func NewMarketCappedPriceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketCappedPriceTransactor, error) {
	contract, err := bindMarketCappedPrice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceTransactor{contract: contract}, nil
}

// NewMarketCappedPriceFilterer creates a new log filterer instance of MarketCappedPrice, bound to a specific deployed contract.
func NewMarketCappedPriceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketCappedPriceFilterer, error) {
	contract, err := bindMarketCappedPrice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceFilterer{contract: contract}, nil
}

// bindMarketCappedPrice binds a generic wrapper to an already deployed contract.
func bindMarketCappedPrice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketCappedPriceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketCappedPrice *MarketCappedPriceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketCappedPrice.Contract.MarketCappedPriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketCappedPrice *MarketCappedPriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.MarketCappedPriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketCappedPrice *MarketCappedPriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.MarketCappedPriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketCappedPrice *MarketCappedPriceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketCappedPrice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketCappedPrice *MarketCappedPriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketCappedPrice *MarketCappedPriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMUSDCMARKETCAP is a free data retrieval call binding the contract method 0x355c7371.
//
// Solidity: function MINIMUM_USDC_MARKET_CAP() view returns(uint256)
func (_MarketCappedPrice *MarketCappedPriceCaller) MINIMUMUSDCMARKETCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "MINIMUM_USDC_MARKET_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMUSDCMARKETCAP is a free data retrieval call binding the contract method 0x355c7371.
//
// Solidity: function MINIMUM_USDC_MARKET_CAP() view returns(uint256)
func (_MarketCappedPrice *MarketCappedPriceSession) MINIMUMUSDCMARKETCAP() (*big.Int, error) {
	return _MarketCappedPrice.Contract.MINIMUMUSDCMARKETCAP(&_MarketCappedPrice.CallOpts)
}

// MINIMUMUSDCMARKETCAP is a free data retrieval call binding the contract method 0x355c7371.
//
// Solidity: function MINIMUM_USDC_MARKET_CAP() view returns(uint256)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) MINIMUMUSDCMARKETCAP() (*big.Int, error) {
	return _MarketCappedPrice.Contract.MINIMUMUSDCMARKETCAP(&_MarketCappedPrice.CallOpts)
}

// EthToken is a free data retrieval call binding the contract method 0x7bf1a627.
//
// Solidity: function ethToken() view returns(address)
func (_MarketCappedPrice *MarketCappedPriceCaller) EthToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "ethToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthToken is a free data retrieval call binding the contract method 0x7bf1a627.
//
// Solidity: function ethToken() view returns(address)
func (_MarketCappedPrice *MarketCappedPriceSession) EthToken() (common.Address, error) {
	return _MarketCappedPrice.Contract.EthToken(&_MarketCappedPrice.CallOpts)
}

// EthToken is a free data retrieval call binding the contract method 0x7bf1a627.
//
// Solidity: function ethToken() view returns(address)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) EthToken() (common.Address, error) {
	return _MarketCappedPrice.Contract.EthToken(&_MarketCappedPrice.CallOpts)
}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0x42d8c2c0.
//
// Solidity: function getFlaunchingFee(address , bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPrice *MarketCappedPriceCaller) GetFlaunchingFee(opts *bind.CallOpts, arg0 common.Address, _initialPriceParams []byte) (*big.Int, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "getFlaunchingFee", arg0, _initialPriceParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0x42d8c2c0.
//
// Solidity: function getFlaunchingFee(address , bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPrice *MarketCappedPriceSession) GetFlaunchingFee(arg0 common.Address, _initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPrice.Contract.GetFlaunchingFee(&_MarketCappedPrice.CallOpts, arg0, _initialPriceParams)
}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0x42d8c2c0.
//
// Solidity: function getFlaunchingFee(address , bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) GetFlaunchingFee(arg0 common.Address, _initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPrice.Contract.GetFlaunchingFee(&_MarketCappedPrice.CallOpts, arg0, _initialPriceParams)
}

// GetMarketCap is a free data retrieval call binding the contract method 0x2a81c4f8.
//
// Solidity: function getMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPrice *MarketCappedPriceCaller) GetMarketCap(opts *bind.CallOpts, _initialPriceParams []byte) (*big.Int, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "getMarketCap", _initialPriceParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCap is a free data retrieval call binding the contract method 0x2a81c4f8.
//
// Solidity: function getMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPrice *MarketCappedPriceSession) GetMarketCap(_initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPrice.Contract.GetMarketCap(&_MarketCappedPrice.CallOpts, _initialPriceParams)
}

// GetMarketCap is a free data retrieval call binding the contract method 0x2a81c4f8.
//
// Solidity: function getMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) GetMarketCap(_initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPrice.Contract.GetMarketCap(&_MarketCappedPrice.CallOpts, _initialPriceParams)
}

// GetSqrtPriceX96 is a free data retrieval call binding the contract method 0xfc87f3f6.
//
// Solidity: function getSqrtPriceX96(address , bool _flipped, bytes _initialPriceParams) view returns(uint160 sqrtPriceX96_)
func (_MarketCappedPrice *MarketCappedPriceCaller) GetSqrtPriceX96(opts *bind.CallOpts, arg0 common.Address, _flipped bool, _initialPriceParams []byte) (*big.Int, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "getSqrtPriceX96", arg0, _flipped, _initialPriceParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSqrtPriceX96 is a free data retrieval call binding the contract method 0xfc87f3f6.
//
// Solidity: function getSqrtPriceX96(address , bool _flipped, bytes _initialPriceParams) view returns(uint160 sqrtPriceX96_)
func (_MarketCappedPrice *MarketCappedPriceSession) GetSqrtPriceX96(arg0 common.Address, _flipped bool, _initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPrice.Contract.GetSqrtPriceX96(&_MarketCappedPrice.CallOpts, arg0, _flipped, _initialPriceParams)
}

// GetSqrtPriceX96 is a free data retrieval call binding the contract method 0xfc87f3f6.
//
// Solidity: function getSqrtPriceX96(address , bool _flipped, bytes _initialPriceParams) view returns(uint160 sqrtPriceX96_)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) GetSqrtPriceX96(arg0 common.Address, _flipped bool, _initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPrice.Contract.GetSqrtPriceX96(&_MarketCappedPrice.CallOpts, arg0, _flipped, _initialPriceParams)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_MarketCappedPrice *MarketCappedPriceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_MarketCappedPrice *MarketCappedPriceSession) Owner() (common.Address, error) {
	return _MarketCappedPrice.Contract.Owner(&_MarketCappedPrice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) Owner() (common.Address, error) {
	return _MarketCappedPrice.Contract.Owner(&_MarketCappedPrice.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_MarketCappedPrice *MarketCappedPriceCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_MarketCappedPrice *MarketCappedPriceSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _MarketCappedPrice.Contract.OwnershipHandoverExpiresAt(&_MarketCappedPrice.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _MarketCappedPrice.Contract.OwnershipHandoverExpiresAt(&_MarketCappedPrice.CallOpts, pendingOwner)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(bytes32)
func (_MarketCappedPrice *MarketCappedPriceCaller) PoolId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(bytes32)
func (_MarketCappedPrice *MarketCappedPriceSession) PoolId() ([32]byte, error) {
	return _MarketCappedPrice.Contract.PoolId(&_MarketCappedPrice.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(bytes32)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) PoolId() ([32]byte, error) {
	return _MarketCappedPrice.Contract.PoolId(&_MarketCappedPrice.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_MarketCappedPrice *MarketCappedPriceCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_MarketCappedPrice *MarketCappedPriceSession) PoolManager() (common.Address, error) {
	return _MarketCappedPrice.Contract.PoolManager(&_MarketCappedPrice.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) PoolManager() (common.Address, error) {
	return _MarketCappedPrice.Contract.PoolManager(&_MarketCappedPrice.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_MarketCappedPrice *MarketCappedPriceCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_MarketCappedPrice *MarketCappedPriceSession) UsdcToken() (common.Address, error) {
	return _MarketCappedPrice.Contract.UsdcToken(&_MarketCappedPrice.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) UsdcToken() (common.Address, error) {
	return _MarketCappedPrice.Contract.UsdcToken(&_MarketCappedPrice.CallOpts)
}

// UsdcToken0 is a free data retrieval call binding the contract method 0x0c7324d3.
//
// Solidity: function usdcToken0() view returns(bool)
func (_MarketCappedPrice *MarketCappedPriceCaller) UsdcToken0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MarketCappedPrice.contract.Call(opts, &out, "usdcToken0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsdcToken0 is a free data retrieval call binding the contract method 0x0c7324d3.
//
// Solidity: function usdcToken0() view returns(bool)
func (_MarketCappedPrice *MarketCappedPriceSession) UsdcToken0() (bool, error) {
	return _MarketCappedPrice.Contract.UsdcToken0(&_MarketCappedPrice.CallOpts)
}

// UsdcToken0 is a free data retrieval call binding the contract method 0x0c7324d3.
//
// Solidity: function usdcToken0() view returns(bool)
func (_MarketCappedPrice *MarketCappedPriceCallerSession) UsdcToken0() (bool, error) {
	return _MarketCappedPrice.Contract.UsdcToken0(&_MarketCappedPrice.CallOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_MarketCappedPrice *MarketCappedPriceTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCappedPrice.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_MarketCappedPrice *MarketCappedPriceSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.CancelOwnershipHandover(&_MarketCappedPrice.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_MarketCappedPrice *MarketCappedPriceTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.CancelOwnershipHandover(&_MarketCappedPrice.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_MarketCappedPrice *MarketCappedPriceTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPrice.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_MarketCappedPrice *MarketCappedPriceSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.CompleteOwnershipHandover(&_MarketCappedPrice.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_MarketCappedPrice *MarketCappedPriceTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.CompleteOwnershipHandover(&_MarketCappedPrice.TransactOpts, pendingOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_MarketCappedPrice *MarketCappedPriceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCappedPrice.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_MarketCappedPrice *MarketCappedPriceSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.RenounceOwnership(&_MarketCappedPrice.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_MarketCappedPrice *MarketCappedPriceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.RenounceOwnership(&_MarketCappedPrice.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_MarketCappedPrice *MarketCappedPriceTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCappedPrice.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_MarketCappedPrice *MarketCappedPriceSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.RequestOwnershipHandover(&_MarketCappedPrice.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_MarketCappedPrice *MarketCappedPriceTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.RequestOwnershipHandover(&_MarketCappedPrice.TransactOpts)
}

// SetPool is a paid mutator transaction binding the contract method 0x8207cb48.
//
// Solidity: function setPool((address,address,uint24,int24,address) _poolKey) returns()
func (_MarketCappedPrice *MarketCappedPriceTransactor) SetPool(opts *bind.TransactOpts, _poolKey PoolKey) (*types.Transaction, error) {
	return _MarketCappedPrice.contract.Transact(opts, "setPool", _poolKey)
}

// SetPool is a paid mutator transaction binding the contract method 0x8207cb48.
//
// Solidity: function setPool((address,address,uint24,int24,address) _poolKey) returns()
func (_MarketCappedPrice *MarketCappedPriceSession) SetPool(_poolKey PoolKey) (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.SetPool(&_MarketCappedPrice.TransactOpts, _poolKey)
}

// SetPool is a paid mutator transaction binding the contract method 0x8207cb48.
//
// Solidity: function setPool((address,address,uint24,int24,address) _poolKey) returns()
func (_MarketCappedPrice *MarketCappedPriceTransactorSession) SetPool(_poolKey PoolKey) (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.SetPool(&_MarketCappedPrice.TransactOpts, _poolKey)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_MarketCappedPrice *MarketCappedPriceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPrice.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_MarketCappedPrice *MarketCappedPriceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.TransferOwnership(&_MarketCappedPrice.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_MarketCappedPrice *MarketCappedPriceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPrice.Contract.TransferOwnership(&_MarketCappedPrice.TransactOpts, newOwner)
}

// MarketCappedPriceOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the MarketCappedPrice contract.
type MarketCappedPriceOwnershipHandoverCanceledIterator struct {
	Event *MarketCappedPriceOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *MarketCappedPriceOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCappedPriceOwnershipHandoverCanceled)
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
		it.Event = new(MarketCappedPriceOwnershipHandoverCanceled)
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
func (it *MarketCappedPriceOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCappedPriceOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCappedPriceOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the MarketCappedPrice contract.
type MarketCappedPriceOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_MarketCappedPrice *MarketCappedPriceFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*MarketCappedPriceOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _MarketCappedPrice.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceOwnershipHandoverCanceledIterator{contract: _MarketCappedPrice.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_MarketCappedPrice *MarketCappedPriceFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *MarketCappedPriceOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _MarketCappedPrice.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCappedPriceOwnershipHandoverCanceled)
				if err := _MarketCappedPrice.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_MarketCappedPrice *MarketCappedPriceFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*MarketCappedPriceOwnershipHandoverCanceled, error) {
	event := new(MarketCappedPriceOwnershipHandoverCanceled)
	if err := _MarketCappedPrice.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketCappedPriceOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the MarketCappedPrice contract.
type MarketCappedPriceOwnershipHandoverRequestedIterator struct {
	Event *MarketCappedPriceOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *MarketCappedPriceOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCappedPriceOwnershipHandoverRequested)
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
		it.Event = new(MarketCappedPriceOwnershipHandoverRequested)
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
func (it *MarketCappedPriceOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCappedPriceOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCappedPriceOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the MarketCappedPrice contract.
type MarketCappedPriceOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_MarketCappedPrice *MarketCappedPriceFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*MarketCappedPriceOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _MarketCappedPrice.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceOwnershipHandoverRequestedIterator{contract: _MarketCappedPrice.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_MarketCappedPrice *MarketCappedPriceFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *MarketCappedPriceOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _MarketCappedPrice.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCappedPriceOwnershipHandoverRequested)
				if err := _MarketCappedPrice.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_MarketCappedPrice *MarketCappedPriceFilterer) ParseOwnershipHandoverRequested(log types.Log) (*MarketCappedPriceOwnershipHandoverRequested, error) {
	event := new(MarketCappedPriceOwnershipHandoverRequested)
	if err := _MarketCappedPrice.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketCappedPriceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MarketCappedPrice contract.
type MarketCappedPriceOwnershipTransferredIterator struct {
	Event *MarketCappedPriceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketCappedPriceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCappedPriceOwnershipTransferred)
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
		it.Event = new(MarketCappedPriceOwnershipTransferred)
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
func (it *MarketCappedPriceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCappedPriceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCappedPriceOwnershipTransferred represents a OwnershipTransferred event raised by the MarketCappedPrice contract.
type MarketCappedPriceOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_MarketCappedPrice *MarketCappedPriceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*MarketCappedPriceOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketCappedPrice.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceOwnershipTransferredIterator{contract: _MarketCappedPrice.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_MarketCappedPrice *MarketCappedPriceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketCappedPriceOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketCappedPrice.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCappedPriceOwnershipTransferred)
				if err := _MarketCappedPrice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MarketCappedPrice *MarketCappedPriceFilterer) ParseOwnershipTransferred(log types.Log) (*MarketCappedPriceOwnershipTransferred, error) {
	event := new(MarketCappedPriceOwnershipTransferred)
	if err := _MarketCappedPrice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
