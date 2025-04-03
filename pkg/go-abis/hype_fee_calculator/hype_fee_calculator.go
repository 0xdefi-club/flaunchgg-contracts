// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hype_fee_calculator

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

// HypeFeeCalculatorMetaData contains all meta data concerning the HypeFeeCalculator contract.
var HypeFeeCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_fairLaunch\",\"type\":\"address\",\"internalType\":\"contractFairLaunch\"},{\"name\":\"_nativeToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAXIMUM_FEE_SCALED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINIMUM_FEE_SCALED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"determineSwapFee\",\"inputs\":[{\"name\":\"_poolKey\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"_baseFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[{\"name\":\"swapFee_\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fairLaunch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractFairLaunch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTargetTokensPerSec\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nativeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolInfos\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"totalTokensSold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetTokensPerSec\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFlaunchParams\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trackSwap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"_delta\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"CallerNotPositionManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroTargetTokensPerSec\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051610d38380380610d3883398101604081905261002e916100cc565b6001600160a01b03808316608081905290821660c05260408051631e46e62f60e21b8152905163791b98bc916004808201926020929091908290030181865afa15801561007d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100a19190610104565b6001600160a01b031660a052506101269050565b6001600160a01b03811681146100c9575f5ffd5b50565b5f5f604083850312156100dd575f5ffd5b82516100e8816100b5565b60208401519092506100f9816100b5565b809150509250929050565b5f60208284031215610114575f5ffd5b815161011f816100b5565b9392505050565b60805160a05160c051610bb76101815f395f81816101ab01526102f101525f818161011b015281816101eb015261069601525f818161015a0152818161026a015281816103a70152818161046201526105f70152610bb75ff3fe608060405234801561000f575f5ffd5b506004361061008c575f3560e01c806324f961771461009057806336a9ac40146100a557806337a8087a146100e55780634039108514610103578063791b98bc1461011657806394e1cf9614610155578063be584cbc1461017c578063d0d8e8b214610185578063e1758bd8146101a6578063edf2ec79146101cd575b5f5ffd5b6100a361009e366004610768565b6101e0565b005b6100cb6100b33660046107fa565b5f602081905290815260409020805460019091015482565b604080519283526020830191909152015b60405180910390f35b6100ef6207a12081565b60405162ffffff90911681526020016100dc565b6100ef6101113660046108e6565b610382565b61013d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100dc565b61013d7f000000000000000000000000000000000000000000000000000000000000000081565b6100ef61271081565b6101986101933660046107fa565b6105c5565b6040519081526020016100dc565b61013d7f000000000000000000000000000000000000000000000000000000000000000081565b6100a36101db36600461098e565b61067c565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461022957604051632af6008960e21b815260040160405180910390fd5b5f61024361023c368890038801886109d5565b60a0902090565b5f81815260208190526040908190209051621cc8fb60e81b815260048101839052919250907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690631cc8fb0090602401602060405180830381865afa1580156102b7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102db91906109ef565b6102e657505061037a565b5f6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001661031e60208a018a610a0a565b6001600160a01b03161461033b576103368660801d90565b610345565b61034586600f0b90565b600f0b90505f81126103575780610360565b61036081610a39565b825f015f8282546103719190610a53565b90915550505050505b505050505050565b5f5f61038f8560a0902090565b604051621cc8fb60e81b8152600481018290529091507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690631cc8fb0090602401602060405180830381865afa1580156103f4573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061041891906109ef565b61042557829150506105be565b5f818152602081815260408083208151808301835281548152600190910154928101929092525163727901fb60e11b8152600481018490529091907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063e4f203f69060240160c060405180830381865afa1580156104af573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104d39190610a66565b90505f61070882602001516104e89190610af8565b6104f29042610af8565b9050805f0361050757859450505050506105be565b82515f90610516908390610b0b565b90505f610522866105c5565b90505f818311610535575061271061058a565b5f6105408385610af8565b90505f83826105546127106207a120610b2a565b62ffffff166105639190610b45565b61056d9190610b0b565b61057990612710610a53565b6207a1208082189082110218925050505b60646105aa8261059a8c84610b5c565b62ffffff16808218908210021890565b6105b49190610b0b565b9750505050505050505b9392505050565b5f818152602081905260408120600101548082036106765760405163727901fb60e11b815260048101849052610708907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063e4f203f69060240160c060405180830381865afa158015610644573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106689190610a66565b608001516105be9190610b0b565b92915050565b5f610689828401846107fa565b9050336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106d457604051632af6008960e21b815260040160405180910390fd5b805f036106f45760405163ca3bb9f560e01b815260040160405180910390fd5b5f93845260208490526040909320600101929092555050565b6001600160a01b0381168114610721575f5ffd5b50565b5f5f83601f840112610734575f5ffd5b5081356001600160401b0381111561074a575f5ffd5b602083019150836020828501011115610761575f5ffd5b9250929050565b5f5f5f5f5f5f86880361016081121561077f575f5ffd5b873561078a8161070d565b965060a0601f198201121561079d575f5ffd5b602088019550606060bf19820112156107b4575f5ffd5b5060c08701935061012087013592506101408701356001600160401b038111156107dc575f5ffd5b6107e889828a01610724565b979a9699509497509295939492505050565b5f6020828403121561080a575f5ffd5b5035919050565b803562ffffff81168114610823575f5ffd5b919050565b8060020b8114610721575f5ffd5b5f60a08284031215610846575f5ffd5b60405160a081016001600160401b038111828210171561087457634e487b7160e01b5f52604160045260245ffd5b60405290508082356108858161070d565b815260208301356108958161070d565b60208201526108a660408401610811565b604082015260608301356108b981610828565b606082015260808301356108cc8161070d565b6080919091015292915050565b8015158114610721575f5ffd5b5f5f5f8385036101208112156108fa575f5ffd5b6109048686610836565b93506060609f1982011215610917575f5ffd5b50604051606081016001600160401b038111828210171561094657634e487b7160e01b5f52604160045260245ffd5b60405260a0850135610957816108d9565b815260c0850135602082015260e08501356109718161070d565b604082015291506109856101008501610811565b90509250925092565b5f5f5f604084860312156109a0575f5ffd5b8335925060208401356001600160401b038111156109bc575f5ffd5b6109c886828701610724565b9497909650939450505050565b5f60a082840312156109e5575f5ffd5b6105be8383610836565b5f602082840312156109ff575f5ffd5b81516105be816108d9565b5f60208284031215610a1a575f5ffd5b81356105be8161070d565b634e487b7160e01b5f52601160045260245ffd5b5f600160ff1b8201610a4d57610a4d610a25565b505f0390565b8082018082111561067657610676610a25565b5f60c0828403128015610a77575f5ffd5b5060405160c081016001600160401b0381118282101715610aa657634e487b7160e01b5f52604160045260245ffd5b60409081528351825260208085015190830152830151610ac581610828565b6040820152606083810151908201526080808401519082015260a0830151610aec816108d9565b60a08201529392505050565b8181038181111561067657610676610a25565b5f82610b2557634e487b7160e01b5f52601260045260245ffd5b500490565b62ffffff828116828216039081111561067657610676610a25565b808202811582820484141761067657610676610a25565b62ffffff8181168382160290811690818114610b7a57610b7a610a25565b509291505056fea26469706673582212205b653f09febebc8f1fecefb1f9969299a8832acf9ffa5039fcbe88a60ca1115664736f6c634300081b0033",
}

// HypeFeeCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use HypeFeeCalculatorMetaData.ABI instead.
var HypeFeeCalculatorABI = HypeFeeCalculatorMetaData.ABI

// HypeFeeCalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HypeFeeCalculatorMetaData.Bin instead.
var HypeFeeCalculatorBin = HypeFeeCalculatorMetaData.Bin

// DeployHypeFeeCalculator deploys a new Ethereum contract, binding an instance of HypeFeeCalculator to it.
func DeployHypeFeeCalculator(auth *bind.TransactOpts, backend bind.ContractBackend, _fairLaunch common.Address, _nativeToken common.Address) (common.Address, *types.Transaction, *HypeFeeCalculator, error) {
	parsed, err := HypeFeeCalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HypeFeeCalculatorBin), backend, _fairLaunch, _nativeToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HypeFeeCalculator{HypeFeeCalculatorCaller: HypeFeeCalculatorCaller{contract: contract}, HypeFeeCalculatorTransactor: HypeFeeCalculatorTransactor{contract: contract}, HypeFeeCalculatorFilterer: HypeFeeCalculatorFilterer{contract: contract}}, nil
}

// HypeFeeCalculator is an auto generated Go binding around an Ethereum contract.
type HypeFeeCalculator struct {
	HypeFeeCalculatorCaller     // Read-only binding to the contract
	HypeFeeCalculatorTransactor // Write-only binding to the contract
	HypeFeeCalculatorFilterer   // Log filterer for contract events
}

// HypeFeeCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type HypeFeeCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HypeFeeCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HypeFeeCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HypeFeeCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HypeFeeCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HypeFeeCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HypeFeeCalculatorSession struct {
	Contract     *HypeFeeCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HypeFeeCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HypeFeeCalculatorCallerSession struct {
	Contract *HypeFeeCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// HypeFeeCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HypeFeeCalculatorTransactorSession struct {
	Contract     *HypeFeeCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// HypeFeeCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type HypeFeeCalculatorRaw struct {
	Contract *HypeFeeCalculator // Generic contract binding to access the raw methods on
}

// HypeFeeCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HypeFeeCalculatorCallerRaw struct {
	Contract *HypeFeeCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// HypeFeeCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HypeFeeCalculatorTransactorRaw struct {
	Contract *HypeFeeCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHypeFeeCalculator creates a new instance of HypeFeeCalculator, bound to a specific deployed contract.
func NewHypeFeeCalculator(address common.Address, backend bind.ContractBackend) (*HypeFeeCalculator, error) {
	contract, err := bindHypeFeeCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HypeFeeCalculator{HypeFeeCalculatorCaller: HypeFeeCalculatorCaller{contract: contract}, HypeFeeCalculatorTransactor: HypeFeeCalculatorTransactor{contract: contract}, HypeFeeCalculatorFilterer: HypeFeeCalculatorFilterer{contract: contract}}, nil
}

// NewHypeFeeCalculatorCaller creates a new read-only instance of HypeFeeCalculator, bound to a specific deployed contract.
func NewHypeFeeCalculatorCaller(address common.Address, caller bind.ContractCaller) (*HypeFeeCalculatorCaller, error) {
	contract, err := bindHypeFeeCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HypeFeeCalculatorCaller{contract: contract}, nil
}

// NewHypeFeeCalculatorTransactor creates a new write-only instance of HypeFeeCalculator, bound to a specific deployed contract.
func NewHypeFeeCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*HypeFeeCalculatorTransactor, error) {
	contract, err := bindHypeFeeCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HypeFeeCalculatorTransactor{contract: contract}, nil
}

// NewHypeFeeCalculatorFilterer creates a new log filterer instance of HypeFeeCalculator, bound to a specific deployed contract.
func NewHypeFeeCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*HypeFeeCalculatorFilterer, error) {
	contract, err := bindHypeFeeCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HypeFeeCalculatorFilterer{contract: contract}, nil
}

// bindHypeFeeCalculator binds a generic wrapper to an already deployed contract.
func bindHypeFeeCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HypeFeeCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HypeFeeCalculator *HypeFeeCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HypeFeeCalculator.Contract.HypeFeeCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HypeFeeCalculator *HypeFeeCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HypeFeeCalculator.Contract.HypeFeeCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HypeFeeCalculator *HypeFeeCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HypeFeeCalculator.Contract.HypeFeeCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HypeFeeCalculator *HypeFeeCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HypeFeeCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HypeFeeCalculator *HypeFeeCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HypeFeeCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HypeFeeCalculator *HypeFeeCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HypeFeeCalculator.Contract.contract.Transact(opts, method, params...)
}

// MAXIMUMFEESCALED is a free data retrieval call binding the contract method 0x37a8087a.
//
// Solidity: function MAXIMUM_FEE_SCALED() view returns(uint24)
func (_HypeFeeCalculator *HypeFeeCalculatorCaller) MAXIMUMFEESCALED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HypeFeeCalculator.contract.Call(opts, &out, "MAXIMUM_FEE_SCALED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMFEESCALED is a free data retrieval call binding the contract method 0x37a8087a.
//
// Solidity: function MAXIMUM_FEE_SCALED() view returns(uint24)
func (_HypeFeeCalculator *HypeFeeCalculatorSession) MAXIMUMFEESCALED() (*big.Int, error) {
	return _HypeFeeCalculator.Contract.MAXIMUMFEESCALED(&_HypeFeeCalculator.CallOpts)
}

// MAXIMUMFEESCALED is a free data retrieval call binding the contract method 0x37a8087a.
//
// Solidity: function MAXIMUM_FEE_SCALED() view returns(uint24)
func (_HypeFeeCalculator *HypeFeeCalculatorCallerSession) MAXIMUMFEESCALED() (*big.Int, error) {
	return _HypeFeeCalculator.Contract.MAXIMUMFEESCALED(&_HypeFeeCalculator.CallOpts)
}

// MINIMUMFEESCALED is a free data retrieval call binding the contract method 0xbe584cbc.
//
// Solidity: function MINIMUM_FEE_SCALED() view returns(uint24)
func (_HypeFeeCalculator *HypeFeeCalculatorCaller) MINIMUMFEESCALED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HypeFeeCalculator.contract.Call(opts, &out, "MINIMUM_FEE_SCALED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMFEESCALED is a free data retrieval call binding the contract method 0xbe584cbc.
//
// Solidity: function MINIMUM_FEE_SCALED() view returns(uint24)
func (_HypeFeeCalculator *HypeFeeCalculatorSession) MINIMUMFEESCALED() (*big.Int, error) {
	return _HypeFeeCalculator.Contract.MINIMUMFEESCALED(&_HypeFeeCalculator.CallOpts)
}

// MINIMUMFEESCALED is a free data retrieval call binding the contract method 0xbe584cbc.
//
// Solidity: function MINIMUM_FEE_SCALED() view returns(uint24)
func (_HypeFeeCalculator *HypeFeeCalculatorCallerSession) MINIMUMFEESCALED() (*big.Int, error) {
	return _HypeFeeCalculator.Contract.MINIMUMFEESCALED(&_HypeFeeCalculator.CallOpts)
}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) _poolKey, (bool,int256,uint160) , uint24 _baseFee) view returns(uint24 swapFee_)
func (_HypeFeeCalculator *HypeFeeCalculatorCaller) DetermineSwapFee(opts *bind.CallOpts, _poolKey PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HypeFeeCalculator.contract.Call(opts, &out, "determineSwapFee", _poolKey, arg1, _baseFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) _poolKey, (bool,int256,uint160) , uint24 _baseFee) view returns(uint24 swapFee_)
func (_HypeFeeCalculator *HypeFeeCalculatorSession) DetermineSwapFee(_poolKey PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	return _HypeFeeCalculator.Contract.DetermineSwapFee(&_HypeFeeCalculator.CallOpts, _poolKey, arg1, _baseFee)
}

// DetermineSwapFee is a free data retrieval call binding the contract method 0x40391085.
//
// Solidity: function determineSwapFee((address,address,uint24,int24,address) _poolKey, (bool,int256,uint160) , uint24 _baseFee) view returns(uint24 swapFee_)
func (_HypeFeeCalculator *HypeFeeCalculatorCallerSession) DetermineSwapFee(_poolKey PoolKey, arg1 IPoolManagerSwapParams, _baseFee *big.Int) (*big.Int, error) {
	return _HypeFeeCalculator.Contract.DetermineSwapFee(&_HypeFeeCalculator.CallOpts, _poolKey, arg1, _baseFee)
}

// FairLaunch is a free data retrieval call binding the contract method 0x94e1cf96.
//
// Solidity: function fairLaunch() view returns(address)
func (_HypeFeeCalculator *HypeFeeCalculatorCaller) FairLaunch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HypeFeeCalculator.contract.Call(opts, &out, "fairLaunch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FairLaunch is a free data retrieval call binding the contract method 0x94e1cf96.
//
// Solidity: function fairLaunch() view returns(address)
func (_HypeFeeCalculator *HypeFeeCalculatorSession) FairLaunch() (common.Address, error) {
	return _HypeFeeCalculator.Contract.FairLaunch(&_HypeFeeCalculator.CallOpts)
}

// FairLaunch is a free data retrieval call binding the contract method 0x94e1cf96.
//
// Solidity: function fairLaunch() view returns(address)
func (_HypeFeeCalculator *HypeFeeCalculatorCallerSession) FairLaunch() (common.Address, error) {
	return _HypeFeeCalculator.Contract.FairLaunch(&_HypeFeeCalculator.CallOpts)
}

// GetTargetTokensPerSec is a free data retrieval call binding the contract method 0xd0d8e8b2.
//
// Solidity: function getTargetTokensPerSec(bytes32 _poolId) view returns(uint256)
func (_HypeFeeCalculator *HypeFeeCalculatorCaller) GetTargetTokensPerSec(opts *bind.CallOpts, _poolId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _HypeFeeCalculator.contract.Call(opts, &out, "getTargetTokensPerSec", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetTokensPerSec is a free data retrieval call binding the contract method 0xd0d8e8b2.
//
// Solidity: function getTargetTokensPerSec(bytes32 _poolId) view returns(uint256)
func (_HypeFeeCalculator *HypeFeeCalculatorSession) GetTargetTokensPerSec(_poolId [32]byte) (*big.Int, error) {
	return _HypeFeeCalculator.Contract.GetTargetTokensPerSec(&_HypeFeeCalculator.CallOpts, _poolId)
}

// GetTargetTokensPerSec is a free data retrieval call binding the contract method 0xd0d8e8b2.
//
// Solidity: function getTargetTokensPerSec(bytes32 _poolId) view returns(uint256)
func (_HypeFeeCalculator *HypeFeeCalculatorCallerSession) GetTargetTokensPerSec(_poolId [32]byte) (*big.Int, error) {
	return _HypeFeeCalculator.Contract.GetTargetTokensPerSec(&_HypeFeeCalculator.CallOpts, _poolId)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_HypeFeeCalculator *HypeFeeCalculatorCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HypeFeeCalculator.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_HypeFeeCalculator *HypeFeeCalculatorSession) NativeToken() (common.Address, error) {
	return _HypeFeeCalculator.Contract.NativeToken(&_HypeFeeCalculator.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_HypeFeeCalculator *HypeFeeCalculatorCallerSession) NativeToken() (common.Address, error) {
	return _HypeFeeCalculator.Contract.NativeToken(&_HypeFeeCalculator.CallOpts)
}

// PoolInfos is a free data retrieval call binding the contract method 0x36a9ac40.
//
// Solidity: function poolInfos(bytes32 ) view returns(uint256 totalTokensSold, uint256 targetTokensPerSec)
func (_HypeFeeCalculator *HypeFeeCalculatorCaller) PoolInfos(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TotalTokensSold    *big.Int
	TargetTokensPerSec *big.Int
}, error) {
	var out []interface{}
	err := _HypeFeeCalculator.contract.Call(opts, &out, "poolInfos", arg0)

	outstruct := new(struct {
		TotalTokensSold    *big.Int
		TargetTokensPerSec *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalTokensSold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TargetTokensPerSec = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfos is a free data retrieval call binding the contract method 0x36a9ac40.
//
// Solidity: function poolInfos(bytes32 ) view returns(uint256 totalTokensSold, uint256 targetTokensPerSec)
func (_HypeFeeCalculator *HypeFeeCalculatorSession) PoolInfos(arg0 [32]byte) (struct {
	TotalTokensSold    *big.Int
	TargetTokensPerSec *big.Int
}, error) {
	return _HypeFeeCalculator.Contract.PoolInfos(&_HypeFeeCalculator.CallOpts, arg0)
}

// PoolInfos is a free data retrieval call binding the contract method 0x36a9ac40.
//
// Solidity: function poolInfos(bytes32 ) view returns(uint256 totalTokensSold, uint256 targetTokensPerSec)
func (_HypeFeeCalculator *HypeFeeCalculatorCallerSession) PoolInfos(arg0 [32]byte) (struct {
	TotalTokensSold    *big.Int
	TargetTokensPerSec *big.Int
}, error) {
	return _HypeFeeCalculator.Contract.PoolInfos(&_HypeFeeCalculator.CallOpts, arg0)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_HypeFeeCalculator *HypeFeeCalculatorCaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HypeFeeCalculator.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_HypeFeeCalculator *HypeFeeCalculatorSession) PositionManager() (common.Address, error) {
	return _HypeFeeCalculator.Contract.PositionManager(&_HypeFeeCalculator.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_HypeFeeCalculator *HypeFeeCalculatorCallerSession) PositionManager() (common.Address, error) {
	return _HypeFeeCalculator.Contract.PositionManager(&_HypeFeeCalculator.CallOpts)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_HypeFeeCalculator *HypeFeeCalculatorTransactor) SetFlaunchParams(opts *bind.TransactOpts, _poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _HypeFeeCalculator.contract.Transact(opts, "setFlaunchParams", _poolId, _params)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_HypeFeeCalculator *HypeFeeCalculatorSession) SetFlaunchParams(_poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _HypeFeeCalculator.Contract.SetFlaunchParams(&_HypeFeeCalculator.TransactOpts, _poolId, _params)
}

// SetFlaunchParams is a paid mutator transaction binding the contract method 0xedf2ec79.
//
// Solidity: function setFlaunchParams(bytes32 _poolId, bytes _params) returns()
func (_HypeFeeCalculator *HypeFeeCalculatorTransactorSession) SetFlaunchParams(_poolId [32]byte, _params []byte) (*types.Transaction, error) {
	return _HypeFeeCalculator.Contract.SetFlaunchParams(&_HypeFeeCalculator.TransactOpts, _poolId, _params)
}

// TrackSwap is a paid mutator transaction binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) _key, (bool,int256,uint160) , int256 _delta, bytes ) returns()
func (_HypeFeeCalculator *HypeFeeCalculatorTransactor) TrackSwap(opts *bind.TransactOpts, arg0 common.Address, _key PoolKey, arg2 IPoolManagerSwapParams, _delta *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _HypeFeeCalculator.contract.Transact(opts, "trackSwap", arg0, _key, arg2, _delta, arg4)
}

// TrackSwap is a paid mutator transaction binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) _key, (bool,int256,uint160) , int256 _delta, bytes ) returns()
func (_HypeFeeCalculator *HypeFeeCalculatorSession) TrackSwap(arg0 common.Address, _key PoolKey, arg2 IPoolManagerSwapParams, _delta *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _HypeFeeCalculator.Contract.TrackSwap(&_HypeFeeCalculator.TransactOpts, arg0, _key, arg2, _delta, arg4)
}

// TrackSwap is a paid mutator transaction binding the contract method 0x24f96177.
//
// Solidity: function trackSwap(address , (address,address,uint24,int24,address) _key, (bool,int256,uint160) , int256 _delta, bytes ) returns()
func (_HypeFeeCalculator *HypeFeeCalculatorTransactorSession) TrackSwap(arg0 common.Address, _key PoolKey, arg2 IPoolManagerSwapParams, _delta *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _HypeFeeCalculator.Contract.TrackSwap(&_HypeFeeCalculator.TransactOpts, arg0, _key, arg2, _delta, arg4)
}
