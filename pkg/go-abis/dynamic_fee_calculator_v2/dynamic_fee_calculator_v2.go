// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dynamic_fee_calculator_v2

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

// DynamicFeeCalculatorV2MetaData contains all meta data concerning the DynamicFeeCalculatorV2 contract.
var DynamicFeeCalculatorV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_positionManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nativeToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"determineSwapFee\",\"inputs\":[{\"name\":\"_poolKey\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"_baseFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[{\"name\":\"swapFee_\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nativeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolInfos\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"currentFeeScaled\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"lastFeeIncreaseTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accumulatorWeightedVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accumulatorLastUpdateTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFlaunchParams\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trackSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"_delta\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"CallerNotPositionManager\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b506040516109b43803806109b483398101604081905261002e91610060565b6001600160a01b039182166080521660a052610091565b80516001600160a01b038116811461005b575f5ffd5b919050565b5f5f60408385031215610071575f5ffd5b61007a83610045565b915061008860208401610045565b90509250929050565b60805160a0516108f46100c05f395f818161014901526101fb01525f818161010a015261018901526108f45ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c806324f961771461006457806336a9ac401461007957806340391085146100de578063791b98bc14610105578063e1758bd814610144578063edf2ec791461016b575b5f5ffd5b61007761007236600461055e565b61017e565b005b6100b36100873660046105f0565b5f60208190529081526040902080546001820154600283015460039093015462ffffff90921692909184565b6040805162ffffff909516855260208501939093529183015260608201526080015b60405180910390f35b6100f16100ec3660046106c6565b610412565b60405162ffffff90911681526020016100d5565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100d5565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b610077610179366004610772565b505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101c757604051632af6008960e21b815260040160405180910390fd5b5f6101e16101da368890038801886107b9565b60a0902090565b5f8181526020818152604082209293506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690610228908a018a6107da565b6001600160a01b031614610245576102408660801d90565b61024f565b61024f86600f0b90565b600f0b90505f5f8212610262578161026b565b61026b82610809565b90505f83600301544261027e9190610823565b9050610e10811115610295575f600285015561030e565b65fca32dc55c718102670de0b6b3a764000003805f19816102b8576102b861083c565b04856002015411156102e8577812725dd1d243aba0e75fe645cc4873f9e65afe688c928e1f21600286015561030c565b670de0b6b3a764000081866002015402816103055761030561083c565b0460028601555b505b815f190384600201541115610329575f196002850155610334565b600284018054830190555b42600385015560028401546a0422ca8b0a00a42500000011610405575f680a18f07d736b90be55601d1b85600201541061037257506207a1206103de565b5f6a0422ca8b0a00a425000000866002015461038e9190610823565b90505f6103b26a0422ca8b0a00a425000000680a18f07d736b90be55601d1b610823565b90506103cd826103c76127106207a120610823565b836104d6565b6103d990612710610850565b925050505b845462ffffff16811461040357845462ffffff191662ffffff82161785554260018601555b505b5050505050505050505050565b5f5f61041f8560a0902090565b5f818152602081905260408120600101549192509061043e9042610823565b90505f610e10821061045357506127106104a1565b5f8381526020819052604081205462ffffff1690610e108461047761271085610823565b6104819190610863565b61048b919061087a565b905061049c8162ffffff8416610823565b925050505b60646104c1826104b18884610899565b62ffffff16808218908210021890565b6104cb919061087a565b979650505050505050565b828202831584820484141782026104f45763ad251c275f526004601cfd5b81810615159190040192915050565b6001600160a01b0381168114610517575f5ffd5b50565b5f5f83601f84011261052a575f5ffd5b5081356001600160401b03811115610540575f5ffd5b602083019150836020828501011115610557575f5ffd5b9250929050565b5f5f5f5f5f5f868803610160811215610575575f5ffd5b873561058081610503565b965060a0601f1982011215610593575f5ffd5b602088019550606060bf19820112156105aa575f5ffd5b5060c08701935061012087013592506101408701356001600160401b038111156105d2575f5ffd5b6105de89828a0161051a565b979a9699509497509295939492505050565b5f60208284031215610600575f5ffd5b5035919050565b803562ffffff81168114610619575f5ffd5b919050565b5f60a0828403121561062e575f5ffd5b60405160a081016001600160401b038111828210171561065c57634e487b7160e01b5f52604160045260245ffd5b604052905080823561066d81610503565b8152602083013561067d81610503565b602082015261068e60408401610607565b604082015260608301358060020b81146106a6575f5ffd5b606082015260808301356106b981610503565b6080919091015292915050565b5f5f5f8385036101208112156106da575f5ffd5b6106e4868661061e565b93506060609f19820112156106f7575f5ffd5b50604051606081016001600160401b038111828210171561072657634e487b7160e01b5f52604160045260245ffd5b60405260a0850135801515811461073b575f5ffd5b815260c0850135602082015260e085013561075581610503565b604082015291506107696101008501610607565b90509250925092565b5f5f5f60408486031215610784575f5ffd5b8335925060208401356001600160401b038111156107a0575f5ffd5b6107ac8682870161051a565b9497909650939450505050565b5f60a082840312156107c9575f5ffd5b6107d3838361061e565b9392505050565b5f602082840312156107ea575f5ffd5b81356107d381610503565b634e487b7160e01b5f52601160045260245ffd5b5f600160ff1b820161081d5761081d6107f5565b505f0390565b81810381811115610836576108366107f5565b92915050565b634e487b7160e01b5f52601260045260245ffd5b80820180821115610836576108366107f5565b8082028115828204841417610836576108366107f5565b5f8261089457634e487b7160e01b5f52601260045260245ffd5b500490565b62ffffff81811683821602908116908181146108b7576108b76107f5565b509291505056fea2646970667358221220055d2c412cdc518eaa9783c21bccbde9a7d0f77492ec9032fd9c127c7289fd8964736f6c634300081b0033",
}

// DynamicFeeCalculatorV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use DynamicFeeCalculatorV2MetaData.ABI instead.
var DynamicFeeCalculatorV2ABI = DynamicFeeCalculatorV2MetaData.ABI

// DynamicFeeCalculatorV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DynamicFeeCalculatorV2MetaData.Bin instead.
var DynamicFeeCalculatorV2Bin = DynamicFeeCalculatorV2MetaData.Bin

// DeployDynamicFeeCalculatorV2 deploys a new Ethereum contract, binding an instance of DynamicFeeCalculatorV2 to it.
func DeployDynamicFeeCalculatorV2(auth *bind.TransactOpts, backend bind.ContractBackend, _positionManager common.Address, _nativeToken common.Address) (common.Address, *types.Transaction, *DynamicFeeCalculatorV2, error) {
	parsed, err := DynamicFeeCalculatorV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DynamicFeeCalculatorV2Bin), backend, _positionManager, _nativeToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DynamicFeeCalculatorV2{DynamicFeeCalculatorV2Caller: DynamicFeeCalculatorV2Caller{contract: contract}, DynamicFeeCalculatorV2Transactor: DynamicFeeCalculatorV2Transactor{contract: contract}, DynamicFeeCalculatorV2Filterer: DynamicFeeCalculatorV2Filterer{contract: contract}}, nil
}

// DynamicFeeCalculatorV2 is an auto generated Go binding around an Ethereum contract.
type DynamicFeeCalculatorV2 struct {
	DynamicFeeCalculatorV2Caller     // Read-only binding to the contract
	DynamicFeeCalculatorV2Transactor // Write-only binding to the contract
	DynamicFeeCalculatorV2Filterer   // Log filterer for contract events
}

// DynamicFeeCalculatorV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type DynamicFeeCalculatorV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicFeeCalculatorV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DynamicFeeCalculatorV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicFeeCalculatorV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DynamicFeeCalculatorV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicFeeCalculatorV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DynamicFeeCalculatorV2Session struct {
	Contract     *DynamicFeeCalculatorV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DynamicFeeCalculatorV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DynamicFeeCalculatorV2CallerSession struct {
	Contract *DynamicFeeCalculatorV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DynamicFeeCalculatorV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DynamicFeeCalculatorV2TransactorSession struct {
	Contract     *DynamicFeeCalculatorV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DynamicFeeCalculatorV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type DynamicFeeCalculatorV2Raw struct {
	Contract *DynamicFeeCalculatorV2 // Generic contract binding to access the raw methods on
}

// DynamicFeeCalculatorV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DynamicFeeCalculatorV2CallerRaw struct {
	Contract *DynamicFeeCalculatorV2Caller // Generic read-only contract binding to access the raw methods on
}

// DynamicFeeCalculatorV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DynamicFeeCalculatorV2TransactorRaw struct {
	Contract *DynamicFeeCalculatorV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDynamicFeeCalculatorV2 creates a new instance of DynamicFeeCalculatorV2, bound to a specific deployed contract.
func NewDynamicFeeCalculatorV2(address common.Address, backend bind.ContractBackend) (*DynamicFeeCalculatorV2, error) {
	contract, err := bindDynamicFeeCalculatorV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DynamicFeeCalculatorV2{DynamicFeeCalculatorV2Caller: DynamicFeeCalculatorV2Caller{contract: contract}, DynamicFeeCalculatorV2Transactor: DynamicFeeCalculatorV2Transactor{contract: contract}, DynamicFeeCalculatorV2Filterer: DynamicFeeCalculatorV2Filterer{contract: contract}}, nil
}

// NewDynamicFeeCalculatorV2Caller creates a new read-only instance of DynamicFeeCalculatorV2, bound to a specific deployed contract.
func NewDynamicFeeCalculatorV2Caller(address common.Address, caller bind.ContractCaller) (*DynamicFeeCalculatorV2Caller, error) {
	contract, err := bindDynamicFeeCalculatorV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DynamicFeeCalculatorV2Caller{contract: contract}, nil
}

// NewDynamicFeeCalculatorV2Transactor creates a new write-only instance of DynamicFeeCalculatorV2, bound to a specific deployed contract.
func NewDynamicFeeCalculatorV2Transactor(address common.Address, transactor bind.ContractTransactor) (*DynamicFeeCalculatorV2Transactor, error) {
	contract, err := bindDynamicFeeCalculatorV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DynamicFeeCalculatorV2Transactor{contract: contract}, nil
}

// NewDynamicFeeCalculatorV2Filterer creates a new log filterer instance of DynamicFeeCalculatorV2, bound to a specific deployed contract.
func NewDynamicFeeCalculatorV2Filterer(address common.Address, filterer bind.ContractFilterer) (*DynamicFeeCalculatorV2Filterer, error) {
	contract, err := bindDynamicFeeCalculatorV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DynamicFeeCalculatorV2Filterer{contract: contract}, nil
}

// bindDynamicFeeCalculatorV2 binds a generic wrapper to an already deployed contract.
func bindDynamicFeeCalculatorV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DynamicFeeCalculatorV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DynamicFeeCalculatorV2.Contract.DynamicFeeCalculatorV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicFeeCalculatorV2.Contract.DynamicFeeCalculatorV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DynamicFeeCalculatorV2.Contract.DynamicFeeCalculatorV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DynamicFeeCalculatorV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicFeeCalculatorV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DynamicFeeCalculatorV2.Contract.contract.Transact(opts, method, params...)
}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) _poolKey, (bool,int256,uint160) , uint24 _baseFee) view returns(uint24 swapFee_)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Caller) DetermineSwapFee(opts *bind.CallOpts, _poolKey PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DynamicFeeCalculatorV2.contract.Call(opts, &out, "determineSwapFee", _poolKey, arg1, _baseFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) _poolKey, (bool,int256,uint160) , uint24 _baseFee) view returns(uint24 swapFee_)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Session) DetermineSwapFee(_poolKey PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	return _DynamicFeeCalculatorV2.Contract.DetermineSwapFee(&_DynamicFeeCalculatorV2.CallOpts, _poolKey, arg1, _baseFee)
}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) _poolKey, (bool,int256,uint160) , uint24 _baseFee) view returns(uint24 swapFee_)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2CallerSession) DetermineSwapFee(_poolKey PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	return _DynamicFeeCalculatorV2.Contract.DetermineSwapFee(&_DynamicFeeCalculatorV2.CallOpts, _poolKey, arg1, _baseFee)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Caller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DynamicFeeCalculatorV2.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Session) NativeToken() (common.Address, error) {
	return _DynamicFeeCalculatorV2.Contract.NativeToken(&_DynamicFeeCalculatorV2.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2CallerSession) NativeToken() (common.Address, error) {
	return _DynamicFeeCalculatorV2.Contract.NativeToken(&_DynamicFeeCalculatorV2.CallOpts)
}

// PoolInfos is a free data retrieval call binding the contract method 0x36a9ac40.
//
// Solidity: function poolInfos(bytes32 _poolId) view returns(uint24 currentFeeScaled, uint256 lastFeeIncreaseTime, uint256 accumulatorWeightedVolume, uint256 accumulatorLastUpdateTime)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Caller) PoolInfos(opts *bind.CallOpts, _poolId [32]byte) (struct {
	CurrentFeeScaled          *big.Int
	LastFeeIncreaseTime       *big.Int
	AccumulatorWeightedVolume *big.Int
	AccumulatorLastUpdateTime *big.Int
}, error) {
	var out []interface{}
	err := _DynamicFeeCalculatorV2.contract.Call(opts, &out, "poolInfos", _poolId)

	outstruct := new(struct {
		CurrentFeeScaled          *big.Int
		LastFeeIncreaseTime       *big.Int
		AccumulatorWeightedVolume *big.Int
		AccumulatorLastUpdateTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentFeeScaled = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastFeeIncreaseTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccumulatorWeightedVolume = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccumulatorLastUpdateTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfos is a free data retrieval call binding the contract method 0x36a9ac40.
//
// Solidity: function poolInfos(bytes32 _poolId) view returns(uint24 currentFeeScaled, uint256 lastFeeIncreaseTime, uint256 accumulatorWeightedVolume, uint256 accumulatorLastUpdateTime)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Session) PoolInfos(_poolId [32]byte) (struct {
	CurrentFeeScaled          *big.Int
	LastFeeIncreaseTime       *big.Int
	AccumulatorWeightedVolume *big.Int
	AccumulatorLastUpdateTime *big.Int
}, error) {
	return _DynamicFeeCalculatorV2.Contract.PoolInfos(&_DynamicFeeCalculatorV2.CallOpts, _poolId)
}

// PoolInfos is a free data retrieval call binding the contract method 0x36a9ac40.
//
// Solidity: function poolInfos(bytes32 _poolId) view returns(uint24 currentFeeScaled, uint256 lastFeeIncreaseTime, uint256 accumulatorWeightedVolume, uint256 accumulatorLastUpdateTime)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2CallerSession) PoolInfos(_poolId [32]byte) (struct {
	CurrentFeeScaled          *big.Int
	LastFeeIncreaseTime       *big.Int
	AccumulatorWeightedVolume *big.Int
	AccumulatorLastUpdateTime *big.Int
}, error) {
	return _DynamicFeeCalculatorV2.Contract.PoolInfos(&_DynamicFeeCalculatorV2.CallOpts, _poolId)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Caller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DynamicFeeCalculatorV2.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Session) PositionManager() (common.Address, error) {
	return _DynamicFeeCalculatorV2.Contract.PositionManager(&_DynamicFeeCalculatorV2.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2CallerSession) PositionManager() (common.Address, error) {
	return _DynamicFeeCalculatorV2.Contract.PositionManager(&_DynamicFeeCalculatorV2.CallOpts)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Transactor) SetFlaunchParams(opts *bind.TransactOpts, _poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculatorV2.contract.Transact(opts, "setFlaunchParams", _poolId, _params)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Session) SetFlaunchParams(_poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculatorV2.Contract.SetFlaunchParams(&_DynamicFeeCalculatorV2.TransactOpts, _poolId, _params)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2TransactorSession) SetFlaunchParams(_poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculatorV2.Contract.SetFlaunchParams(&_DynamicFeeCalculatorV2.TransactOpts, _poolId, _params)
}

// TrackSwap is a paid mutator transaction binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) _key, (bool,int256,uint160) , int256 _delta, bytes ) returns()
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Transactor) TrackSwap(opts *bind.TransactOpts, arg0 common.Address, _key PoolKey, arg2 IPoolManagerSwapParams, _delta *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculatorV2.contract.Transact(opts, "trackSwap", arg0, _key, arg2, _delta, arg4)
}

// TrackSwap is a paid mutator transaction binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) _key, (bool,int256,uint160) , int256 _delta, bytes ) returns()
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2Session) TrackSwap(arg0 common.Address, _key PoolKey, arg2 IPoolManagerSwapParams, _delta *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculatorV2.Contract.TrackSwap(&_DynamicFeeCalculatorV2.TransactOpts, arg0, _key, arg2, _delta, arg4)
}

// TrackSwap is a paid mutator transaction binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) _key, (bool,int256,uint160) , int256 _delta, bytes ) returns()
func (_DynamicFeeCalculatorV2 *DynamicFeeCalculatorV2TransactorSession) TrackSwap(arg0 common.Address, _key PoolKey, arg2 IPoolManagerSwapParams, _delta *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DynamicFeeCalculatorV2.Contract.TrackSwap(&_DynamicFeeCalculatorV2.TransactOpts, arg0, _key, arg2, _delta, arg4)
}
