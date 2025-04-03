// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dynamic_fee_calculator

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

// DynamicFeeCalculatorMetaData contains all meta data concerning the DynamicFeeCalculator contract.
var DynamicFeeCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_positionManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"determineSwapFee\",\"inputs\":[{\"name\":\"_poolKey\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"_baseFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[{\"name\":\"swapFee_\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolInfos\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"currentFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"lastFeeDecreaseTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token1SoFar\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFlaunchParams\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trackSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"_delta\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"CallerNotPositionManager\",\"inputs\":[]}]",
	Bin: "0x60a0604052348015600e575f5ffd5b50604051610908380380610908833981016040819052602b91603b565b6001600160a01b03166080526066565b5f60208284031215604a575f5ffd5b81516001600160a01b0381168114605f575f5ffd5b9392505050565b6080516108846100845f395f818160f4015261014c01526108845ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806324f961771461005957806336a9ac401461006e57806340391085146100c8578063791b98bc146100ef578063edf2ec791461012e575b5f5ffd5b61006c6100673660046104b8565b610141565b005b6100a261007c36600461054a565b5f6020819052908152604090208054600182015460029092015462ffffff909116919083565b6040805162ffffff90941684526020840192909252908201526060015b60405180910390f35b6100db6100d6366004610620565b61036e565b60405162ffffff90911681526020016100bf565b6101167f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100bf565b61006c61013c3660046106cc565b505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461018a57604051632af6008960e21b815260040160405180910390fd5b5f6101a461019d36889003880188610713565b60a0902090565b5f81815260208190526040812060018101549293509190036101c7574260018201555b5f600a62ffffff16600a8360010154426101e19190610741565b6101eb9190610754565b6101f59190610773565b90505f6102018261078a565b90505f61022961021460208c018c6107a4565b61022460408d0160208e016107a4565b610422565b61023c5761023788600f0b90565b610246565b6102468860801d90565b600f0b846002015461025891906107bf565b90505f600361027169021e19e0c9bab240000084610754565b61027b9190610773565b905061028781846107d2565b925060645f841261029857836102a1565b6102a18461078a565b13156103605783156102d857600a6102b98186610754565b6102c39190610773565b85600101546102d291906107bf565b60018601555b60036102ee69021e19e0c9bab240000083610773565b6102f89190610754565b6103029083610741565b60028087019190915585545f9161033c916103349161032b91889162ffffff909116900b6107d2565b62061a80610438565b61271061044f565b865490915062ffffff16811461035e57855462ffffff191662ffffff82161786555b505b505050505050505050505050565b5f5f5f5f61037d8760a0902090565b815260208082019290925260409081015f9081208251606081018452815462ffffff16815260018201549481018590526002909101549281019290925290925090600a9081906103cd9042610741565b6103d79190610754565b6103e19190610773565b9050606461040e82845f015160020b6103fa91906107f9565b610409600288900b606461081f565b61044f565b6104189190610754565b9695505050505050565b6001600160a01b03808216908316115b92915050565b5f8183126104465781610448565b825b9392505050565b5f8183136104465781610448565b6001600160a01b0381168114610471575f5ffd5b50565b5f5f83601f840112610484575f5ffd5b5081356001600160401b0381111561049a575f5ffd5b6020830191508360208285010111156104b1575f5ffd5b9250929050565b5f5f5f5f5f5f8688036101608112156104cf575f5ffd5b87356104da8161045d565b965060a0601f19820112156104ed575f5ffd5b602088019550606060bf1982011215610504575f5ffd5b5060c08701935061012087013592506101408701356001600160401b0381111561052c575f5ffd5b61053889828a01610474565b979a9699509497509295939492505050565b5f6020828403121561055a575f5ffd5b5035919050565b803562ffffff81168114610573575f5ffd5b919050565b5f60a08284031215610588575f5ffd5b60405160a081016001600160401b03811182821017156105b657634e487b7160e01b5f52604160045260245ffd5b60405290508082356105c78161045d565b815260208301356105d78161045d565b60208201526105e860408401610561565b604082015260608301358060020b8114610600575f5ffd5b606082015260808301356106138161045d565b6080919091015292915050565b5f5f5f838503610120811215610634575f5ffd5b61063e8686610578565b93506060609f1982011215610651575f5ffd5b50604051606081016001600160401b038111828210171561068057634e487b7160e01b5f52604160045260245ffd5b60405260a08501358015158114610695575f5ffd5b815260c0850135602082015260e08501356106af8161045d565b604082015291506106c36101008501610561565b90509250925092565b5f5f5f604084860312156106de575f5ffd5b8335925060208401356001600160401b038111156106fa575f5ffd5b61070686828701610474565b9497909650939450505050565b5f60a08284031215610723575f5ffd5b6104488383610578565b634e487b7160e01b5f52601160045260245ffd5b818103818111156104325761043261072d565b5f8261076e57634e487b7160e01b5f52601260045260245ffd5b500490565b80820281158282048414176104325761043261072d565b5f600160ff1b820161079e5761079e61072d565b505f0390565b5f602082840312156107b4575f5ffd5b81356104488161045d565b808201808211156104325761043261072d565b8082018281125f8312801582168215821617156107f1576107f161072d565b505092915050565b8181035f8312801583831316838312821617156108185761081861072d565b5092915050565b8082025f8212600160ff1b8414161561083a5761083a61072d565b81810583148215176104325761043261072d56fea2646970667358221220ec87b00b32518269653b00116c352283c916919df9768b562080a0f3d6d5b7ef64736f6c634300081b0033",
}

// DynamicFeeCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use DynamicFeeCalculatorMetaData.ABI instead.
var DynamicFeeCalculatorABI = DynamicFeeCalculatorMetaData.ABI

// DynamicFeeCalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DynamicFeeCalculatorMetaData.Bin instead.
var DynamicFeeCalculatorBin = DynamicFeeCalculatorMetaData.Bin

// DeployDynamicFeeCalculator deploys a new Ethereum contract, binding an instance of DynamicFeeCalculator to it.
func DeployDynamicFeeCalculator(auth *bind.TransactOpts, backend bind.ContractBackend, _positionManager common.Address) (common.Address, *types.Transaction, *DynamicFeeCalculator, error) {
	parsed, err := DynamicFeeCalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DynamicFeeCalculatorBin), backend, _positionManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DynamicFeeCalculator{DynamicFeeCalculatorCaller: DynamicFeeCalculatorCaller{contract: contract}, DynamicFeeCalculatorTransactor: DynamicFeeCalculatorTransactor{contract: contract}, DynamicFeeCalculatorFilterer: DynamicFeeCalculatorFilterer{contract: contract}}, nil
}

// DynamicFeeCalculator is an auto generated Go binding around an Ethereum contract.
type DynamicFeeCalculator struct {
	DynamicFeeCalculatorCaller     // Read-only binding to the contract
	DynamicFeeCalculatorTransactor // Write-only binding to the contract
	DynamicFeeCalculatorFilterer   // Log filterer for contract events
}

// DynamicFeeCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DynamicFeeCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicFeeCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DynamicFeeCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicFeeCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DynamicFeeCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicFeeCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DynamicFeeCalculatorSession struct {
	Contract     *DynamicFeeCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DynamicFeeCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DynamicFeeCalculatorCallerSession struct {
	Contract *DynamicFeeCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DynamicFeeCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DynamicFeeCalculatorTransactorSession struct {
	Contract     *DynamicFeeCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DynamicFeeCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DynamicFeeCalculatorRaw struct {
	Contract *DynamicFeeCalculator // Generic contract binding to access the raw methods on
}

// DynamicFeeCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DynamicFeeCalculatorCallerRaw struct {
	Contract *DynamicFeeCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// DynamicFeeCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DynamicFeeCalculatorTransactorRaw struct {
	Contract *DynamicFeeCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDynamicFeeCalculator creates a new instance of DynamicFeeCalculator, bound to a specific deployed contract.
func NewDynamicFeeCalculator(address common.Address, backend bind.ContractBackend) (*DynamicFeeCalculator, error) {
	contract, err := bindDynamicFeeCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DynamicFeeCalculator{DynamicFeeCalculatorCaller: DynamicFeeCalculatorCaller{contract: contract}, DynamicFeeCalculatorTransactor: DynamicFeeCalculatorTransactor{contract: contract}, DynamicFeeCalculatorFilterer: DynamicFeeCalculatorFilterer{contract: contract}}, nil
}

// NewDynamicFeeCalculatorCaller creates a new read-only instance of DynamicFeeCalculator, bound to a specific deployed contract.
func NewDynamicFeeCalculatorCaller(address common.Address, caller bind.ContractCaller) (*DynamicFeeCalculatorCaller, error) {
	contract, err := bindDynamicFeeCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DynamicFeeCalculatorCaller{contract: contract}, nil
}

// NewDynamicFeeCalculatorTransactor creates a new write-only instance of DynamicFeeCalculator, bound to a specific deployed contract.
func NewDynamicFeeCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DynamicFeeCalculatorTransactor, error) {
	contract, err := bindDynamicFeeCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DynamicFeeCalculatorTransactor{contract: contract}, nil
}

// NewDynamicFeeCalculatorFilterer creates a new log filterer instance of DynamicFeeCalculator, bound to a specific deployed contract.
func NewDynamicFeeCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DynamicFeeCalculatorFilterer, error) {
	contract, err := bindDynamicFeeCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DynamicFeeCalculatorFilterer{contract: contract}, nil
}

// bindDynamicFeeCalculator binds a generic wrapper to an already deployed contract.
func bindDynamicFeeCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DynamicFeeCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DynamicFeeCalculator *DynamicFeeCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DynamicFeeCalculator.Contract.DynamicFeeCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DynamicFeeCalculator *DynamicFeeCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicFeeCalculator.Contract.DynamicFeeCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DynamicFeeCalculator *DynamicFeeCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DynamicFeeCalculator.Contract.DynamicFeeCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DynamicFeeCalculator *DynamicFeeCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DynamicFeeCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DynamicFeeCalculator *DynamicFeeCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicFeeCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DynamicFeeCalculator *DynamicFeeCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DynamicFeeCalculator.Contract.contract.Transact(opts, method, params...)
}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) _poolKey, (bool,int256,uint160) , uint24 _baseFee) view returns(uint24 swapFee_)
func (_DynamicFeeCalculator *DynamicFeeCalculatorCaller) DetermineSwapFee(opts *bind.CallOpts, _poolKey PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DynamicFeeCalculator.contract.Call(opts, &out, "determineSwapFee", _poolKey, arg1, _baseFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) _poolKey, (bool,int256,uint160) , uint24 _baseFee) view returns(uint24 swapFee_)
func (_DynamicFeeCalculator *DynamicFeeCalculatorSession) DetermineSwapFee(_poolKey PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	return _DynamicFeeCalculator.Contract.DetermineSwapFee(&_DynamicFeeCalculator.CallOpts, _poolKey, arg1, _baseFee)
}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) _poolKey, (bool,int256,uint160) , uint24 _baseFee) view returns(uint24 swapFee_)
func (_DynamicFeeCalculator *DynamicFeeCalculatorCallerSession) DetermineSwapFee(_poolKey PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	return _DynamicFeeCalculator.Contract.DetermineSwapFee(&_DynamicFeeCalculator.CallOpts, _poolKey, arg1, _baseFee)
}

// PoolInfos is a free data retrieval call binding the contract method 0x36a9ac40.
//
// Solidity: function poolInfos(bytes32 _poolId) view returns(uint24 currentFee, uint256 lastFeeDecreaseTime, uint256 token1SoFar)
func (_DynamicFeeCalculator *DynamicFeeCalculatorCaller) PoolInfos(opts *bind.CallOpts, _poolId [32]byte) (struct {
	CurrentFee          *big.Int
	LastFeeDecreaseTime *big.Int
	Token1SoFar         *big.Int
}, error) {
	var out []interface{}
	err := _DynamicFeeCalculator.contract.Call(opts, &out, "poolInfos", _poolId)

	outstruct := new(struct {
		CurrentFee          *big.Int
		LastFeeDecreaseTime *big.Int
		Token1SoFar         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastFeeDecreaseTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Token1SoFar = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfos is a free data retrieval call binding the contract method 0x36a9ac40.
//
// Solidity: function poolInfos(bytes32 _poolId) view returns(uint24 currentFee, uint256 lastFeeDecreaseTime, uint256 token1SoFar)
func (_DynamicFeeCalculator *DynamicFeeCalculatorSession) PoolInfos(_poolId [32]byte) (struct {
	CurrentFee          *big.Int
	LastFeeDecreaseTime *big.Int
	Token1SoFar         *big.Int
}, error) {
	return _DynamicFeeCalculator.Contract.PoolInfos(&_DynamicFeeCalculator.CallOpts, _poolId)
}

// PoolInfos is a free data retrieval call binding the contract method 0x36a9ac40.
//
// Solidity: function poolInfos(bytes32 _poolId) view returns(uint24 currentFee, uint256 lastFeeDecreaseTime, uint256 token1SoFar)
func (_DynamicFeeCalculator *DynamicFeeCalculatorCallerSession) PoolInfos(_poolId [32]byte) (struct {
	CurrentFee          *big.Int
	LastFeeDecreaseTime *big.Int
	Token1SoFar         *big.Int
}, error) {
	return _DynamicFeeCalculator.Contract.PoolInfos(&_DynamicFeeCalculator.CallOpts, _poolId)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_DynamicFeeCalculator *DynamicFeeCalculatorCaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DynamicFeeCalculator.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_DynamicFeeCalculator *DynamicFeeCalculatorSession) PositionManager() (common.Address, error) {
	return _DynamicFeeCalculator.Contract.PositionManager(&_DynamicFeeCalculator.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_DynamicFeeCalculator *DynamicFeeCalculatorCallerSession) PositionManager() (common.Address, error) {
	return _DynamicFeeCalculator.Contract.PositionManager(&_DynamicFeeCalculator.CallOpts)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_DynamicFeeCalculator *DynamicFeeCalculatorTransactor) SetFlaunchParams(opts *bind.TransactOpts, _poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculator.contract.Transact(opts, "setFlaunchParams", _poolId, _params)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_DynamicFeeCalculator *DynamicFeeCalculatorSession) SetFlaunchParams(_poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculator.Contract.SetFlaunchParams(&_DynamicFeeCalculator.TransactOpts, _poolId, _params)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_DynamicFeeCalculator *DynamicFeeCalculatorTransactorSession) SetFlaunchParams(_poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculator.Contract.SetFlaunchParams(&_DynamicFeeCalculator.TransactOpts, _poolId, _params)
}

// TrackSwap is a paid mutator transaction binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) _key, (bool,int256,uint160) , int256 _delta, bytes ) returns()
func (_DynamicFeeCalculator *DynamicFeeCalculatorTransactor) TrackSwap(opts *bind.TransactOpts, arg0 common.Address, _key PoolKey, arg2 IPoolManagerSwapParams, _delta *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculator.contract.Transact(opts, "trackSwap", arg0, _key, arg2, _delta, arg4)
}

// TrackSwap is a paid mutator transaction binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) _key, (bool,int256,uint160) , int256 _delta, bytes ) returns()
func (_DynamicFeeCalculator *DynamicFeeCalculatorSession) TrackSwap(arg0 common.Address, _key PoolKey, arg2 IPoolManagerSwapParams, _delta *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculator.Contract.TrackSwap(&_DynamicFeeCalculator.TransactOpts, arg0, _key, arg2, _delta, arg4)
}

// TrackSwap is a paid mutator transaction binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) _key, (bool,int256,uint160) , int256 _delta, bytes ) returns()
func (_DynamicFeeCalculator *DynamicFeeCalculatorTransactorSession) TrackSwap(arg0 common.Address, _key PoolKey, arg2 IPoolManagerSwapParams, _delta *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculator.Contract.TrackSwap(&_DynamicFeeCalculator.TransactOpts, arg0, _key, arg2, _delta, arg4)
}
