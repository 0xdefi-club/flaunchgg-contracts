// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package memecoin_treasury

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

// MemecoinTreasuryMetaData contains all meta data concerning the MemecoinTreasury contract.
var MemecoinTreasuryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"actionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractTreasuryActionManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimFees\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeAction\",\"inputs\":[{\"name\":\"_action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_positionManager\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_actionManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nativeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolKey\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nativeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolKey\",\"inputs\":[],\"outputs\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPositionManager\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ActionExecuted\",\"inputs\":[{\"name\":\"_action\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_poolKey\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"_data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ActionNotApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610ab68061001c5f395ff3fe608060405260043610610062575f3560e01c8063182148ef1461006d57806375d252a5146100f8578063791b98bc146101245780638257d1cc146101435780638f39d13a14610164578063d294f09314610183578063e1758bd814610197575f5ffd5b3661006957005b5f5ffd5b348015610078575f5ffd5b506003546004546005546100b3926001600160a01b03908116928082169262ffffff600160a01b83041692600160b81b90920460020b911685565b604080516001600160a01b039687168152948616602086015262ffffff9093169284019290925260020b606083015291909116608082015260a0015b60405180910390f35b348015610103575f5ffd5b50600154610117906001600160a01b031681565b6040516100ef9190610786565b34801561012f575f5ffd5b50600254610117906001600160a01b031681565b34801561014e575f5ffd5b5061016261015d36600461082d565b6101b5565b005b34801561016f575f5ffd5b5061016261017e3660046108fb565b6102d7565b34801561018e575f5ffd5b506101626106fc565b3480156101a2575f5ffd5b505f54610117906001600160a01b031681565b63409feecd1980546003825580156101eb5760018160011c14303b106101e25763f92ee8a95f526004601cfd5b818160ff1b1b91505b50600180546001600160a01b03199081166001600160a01b03878116919091179092555f805482168684161790558351600380548316918416919091179055602084015160048054604087015160608801519386166001600160b81b031990921691909117600160a01b62ffffff928316021762ffffff60b81b1916600160b81b919093160291909117905560808401516005805483169184169190911790556002805490911691871691909117905580156102d0576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b5050505050565b3068929eee149b4bd2126854036102f55763ab143c065f526004601cfd5b3068929eee149b4bd2126855600154604051633869a50760e11b81526001600160a01b03909116906370d34a0e90610331908590600401610786565b602060405180830381865afa15801561034c573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061037091906109a0565b61038d576040516304130a2160e01b815260040160405180910390fd5b5f80546040805160a0810182526003546001600160a01b039081168252600454808216602084015262ffffff600160a01b82041693830193909352600160b81b90920460020b6060820152600554821660808201526103ee9290911661075c565b6001600160a01b03166302d05d3f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610429573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061044d91906109bf565b90506001600160a01b0381163314610477576040516282b42960e81b815260040160405180910390fd5b6003546004805460405163095ea7b360e01b81526001600160a01b039384169390911691839163095ea7b3916104b29189915f1991016109da565b6020604051808303815f875af11580156104ce573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104f291906109a0565b5060405163095ea7b360e01b81526001600160a01b0382169063095ea7b3906105229088905f19906004016109da565b6020604051808303815f875af115801561053e573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061056291906109a0565b5061056b6106fc565b6040516337f821a560e11b81526001600160a01b03861690636ff0434a9061059a9060039088906004016109f3565b5f604051808303815f87803b1580156105b1575f5ffd5b505af11580156105c3573d5f5f3e3d5ffd5b50505050846001600160a01b03167f32ea44c9b28c6fcd233c31b0c0ef74ef50bdbe347bcbb0775445b50056bbe3e96003866040516106039291906109f3565b60405180910390a260405163095ea7b360e01b81526001600160a01b0383169063095ea7b3906106399088905f906004016109da565b6020604051808303815f875af1158015610655573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061067991906109a0565b5060405163095ea7b360e01b81526001600160a01b0382169063095ea7b3906106a89088905f906004016109da565b6020604051808303815f875af11580156106c4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106e891906109a0565b505050503868929eee149b4bd21268555050565b600254604051630130b65360e61b81523060048201525f60248201526001600160a01b0390911690634c2d94c0906044015f604051808303815f87803b158015610744575f5ffd5b505af1158015610756573d5f5f3e3d5ffd5b50505050565b81515f906001600160a01b0380841691161461077957825161077f565b82602001515b9392505050565b6001600160a01b0391909116815260200190565b6001600160a01b03811681146107ae575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b60405160a081016001600160401b03811182821017156107e7576107e76107b1565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610815576108156107b1565b604052919050565b80356108288161079a565b919050565b5f5f5f5f848603610100811215610842575f5ffd5b853561084d8161079a565b9450602086013561085d8161079a565b9350604086013561086d8161079a565b925060a0605f1982011215610880575f5ffd5b506108896107c5565b60608601356108978161079a565b815260808601356108a78161079a565b602082015260a086013562ffffff811681146108c1575f5ffd5b604082015260c0860135600281900b81146108da575f5ffd5b60608201526108eb60e0870161081d565b6080820152939692955090935050565b5f5f6040838503121561090c575f5ffd5b82356109178161079a565b915060208301356001600160401b03811115610931575f5ffd5b8301601f81018513610941575f5ffd5b80356001600160401b0381111561095a5761095a6107b1565b61096d601f8201601f19166020016107ed565b818152866020838501011115610981575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f602082840312156109b0575f5ffd5b8151801515811461077f575f5ffd5b5f602082840312156109cf575f5ffd5b815161077f8161079a565b6001600160a01b03929092168252602082015260400190565b82546001600160a01b039081168252600184015480821660208085019190915260a082811c62ffffff16604086015260b89290921c600290810b6060860152860154909216608084015260c090830181905283519083018181525f92909183918190870160e087015e5f602082850101526020601f19601f8301168401019150508092505050939250505056fea2646970667358221220595446abb419a17bec525c4be2ea356011c8786e0451e99018e2ad31ee52796a64736f6c634300081b0033",
}

// MemecoinTreasuryABI is the input ABI used to generate the binding from.
// Deprecated: Use MemecoinTreasuryMetaData.ABI instead.
var MemecoinTreasuryABI = MemecoinTreasuryMetaData.ABI

// MemecoinTreasuryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MemecoinTreasuryMetaData.Bin instead.
var MemecoinTreasuryBin = MemecoinTreasuryMetaData.Bin

// DeployMemecoinTreasury deploys a new Ethereum contract, binding an instance of MemecoinTreasury to it.
func DeployMemecoinTreasury(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MemecoinTreasury, error) {
	parsed, err := MemecoinTreasuryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MemecoinTreasuryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MemecoinTreasury{MemecoinTreasuryCaller: MemecoinTreasuryCaller{contract: contract}, MemecoinTreasuryTransactor: MemecoinTreasuryTransactor{contract: contract}, MemecoinTreasuryFilterer: MemecoinTreasuryFilterer{contract: contract}}, nil
}

// MemecoinTreasury is an auto generated Go binding around an Ethereum contract.
type MemecoinTreasury struct {
	MemecoinTreasuryCaller     // Read-only binding to the contract
	MemecoinTreasuryTransactor // Write-only binding to the contract
	MemecoinTreasuryFilterer   // Log filterer for contract events
}

// MemecoinTreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemecoinTreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemecoinTreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemecoinTreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemecoinTreasuryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemecoinTreasuryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemecoinTreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemecoinTreasurySession struct {
	Contract     *MemecoinTreasury // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemecoinTreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemecoinTreasuryCallerSession struct {
	Contract *MemecoinTreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MemecoinTreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemecoinTreasuryTransactorSession struct {
	Contract     *MemecoinTreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MemecoinTreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemecoinTreasuryRaw struct {
	Contract *MemecoinTreasury // Generic contract binding to access the raw methods on
}

// MemecoinTreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemecoinTreasuryCallerRaw struct {
	Contract *MemecoinTreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// MemecoinTreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemecoinTreasuryTransactorRaw struct {
	Contract *MemecoinTreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemecoinTreasury creates a new instance of MemecoinTreasury, bound to a specific deployed contract.
func NewMemecoinTreasury(address common.Address, backend bind.ContractBackend) (*MemecoinTreasury, error) {
	contract, err := bindMemecoinTreasury(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MemecoinTreasury{MemecoinTreasuryCaller: MemecoinTreasuryCaller{contract: contract}, MemecoinTreasuryTransactor: MemecoinTreasuryTransactor{contract: contract}, MemecoinTreasuryFilterer: MemecoinTreasuryFilterer{contract: contract}}, nil
}

// NewMemecoinTreasuryCaller creates a new read-only instance of MemecoinTreasury, bound to a specific deployed contract.
func NewMemecoinTreasuryCaller(address common.Address, caller bind.ContractCaller) (*MemecoinTreasuryCaller, error) {
	contract, err := bindMemecoinTreasury(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemecoinTreasuryCaller{contract: contract}, nil
}

// NewMemecoinTreasuryTransactor creates a new write-only instance of MemecoinTreasury, bound to a specific deployed contract.
func NewMemecoinTreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*MemecoinTreasuryTransactor, error) {
	contract, err := bindMemecoinTreasury(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemecoinTreasuryTransactor{contract: contract}, nil
}

// NewMemecoinTreasuryFilterer creates a new log filterer instance of MemecoinTreasury, bound to a specific deployed contract.
func NewMemecoinTreasuryFilterer(address common.Address, filterer bind.ContractFilterer) (*MemecoinTreasuryFilterer, error) {
	contract, err := bindMemecoinTreasury(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemecoinTreasuryFilterer{contract: contract}, nil
}

// bindMemecoinTreasury binds a generic wrapper to an already deployed contract.
func bindMemecoinTreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MemecoinTreasuryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemecoinTreasury *MemecoinTreasuryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemecoinTreasury.Contract.MemecoinTreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemecoinTreasury *MemecoinTreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.MemecoinTreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemecoinTreasury *MemecoinTreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.MemecoinTreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemecoinTreasury *MemecoinTreasuryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemecoinTreasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemecoinTreasury *MemecoinTreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemecoinTreasury *MemecoinTreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.contract.Transact(opts, method, params...)
}

// ActionManager is a free data retrieval call binding the contract method 0x75d252a5.
//
// Solidity: function actionManager() view returns(address)
func (_MemecoinTreasury *MemecoinTreasuryCaller) ActionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MemecoinTreasury.contract.Call(opts, &out, "actionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActionManager is a free data retrieval call binding the contract method 0x75d252a5.
//
// Solidity: function actionManager() view returns(address)
func (_MemecoinTreasury *MemecoinTreasurySession) ActionManager() (common.Address, error) {
	return _MemecoinTreasury.Contract.ActionManager(&_MemecoinTreasury.CallOpts)
}

// ActionManager is a free data retrieval call binding the contract method 0x75d252a5.
//
// Solidity: function actionManager() view returns(address)
func (_MemecoinTreasury *MemecoinTreasuryCallerSession) ActionManager() (common.Address, error) {
	return _MemecoinTreasury.Contract.ActionManager(&_MemecoinTreasury.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_MemecoinTreasury *MemecoinTreasuryCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MemecoinTreasury.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_MemecoinTreasury *MemecoinTreasurySession) NativeToken() (common.Address, error) {
	return _MemecoinTreasury.Contract.NativeToken(&_MemecoinTreasury.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_MemecoinTreasury *MemecoinTreasuryCallerSession) NativeToken() (common.Address, error) {
	return _MemecoinTreasury.Contract.NativeToken(&_MemecoinTreasury.CallOpts)
}

// PoolKey is a free data retrieval call binding the contract method 0x182148ef.
//
// Solidity: function poolKey() view returns(address currency0, address currency1, uint24 fee, int24 tickSpacing, address hooks)
func (_MemecoinTreasury *MemecoinTreasuryCaller) PoolKey(opts *bind.CallOpts) (struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}, error) {
	var out []interface{}
	err := _MemecoinTreasury.contract.Call(opts, &out, "poolKey")

	outstruct := new(struct {
		Currency0   common.Address
		Currency1   common.Address
		Fee         *big.Int
		TickSpacing *big.Int
		Hooks       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Currency0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Currency1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TickSpacing = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Hooks = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PoolKey is a free data retrieval call binding the contract method 0x182148ef.
//
// Solidity: function poolKey() view returns(address currency0, address currency1, uint24 fee, int24 tickSpacing, address hooks)
func (_MemecoinTreasury *MemecoinTreasurySession) PoolKey() (struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}, error) {
	return _MemecoinTreasury.Contract.PoolKey(&_MemecoinTreasury.CallOpts)
}

// PoolKey is a free data retrieval call binding the contract method 0x182148ef.
//
// Solidity: function poolKey() view returns(address currency0, address currency1, uint24 fee, int24 tickSpacing, address hooks)
func (_MemecoinTreasury *MemecoinTreasuryCallerSession) PoolKey() (struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}, error) {
	return _MemecoinTreasury.Contract.PoolKey(&_MemecoinTreasury.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_MemecoinTreasury *MemecoinTreasuryCaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MemecoinTreasury.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_MemecoinTreasury *MemecoinTreasurySession) PositionManager() (common.Address, error) {
	return _MemecoinTreasury.Contract.PositionManager(&_MemecoinTreasury.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_MemecoinTreasury *MemecoinTreasuryCallerSession) PositionManager() (common.Address, error) {
	return _MemecoinTreasury.Contract.PositionManager(&_MemecoinTreasury.CallOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_MemecoinTreasury *MemecoinTreasuryTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemecoinTreasury.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_MemecoinTreasury *MemecoinTreasurySession) ClaimFees() (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.ClaimFees(&_MemecoinTreasury.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_MemecoinTreasury *MemecoinTreasuryTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.ClaimFees(&_MemecoinTreasury.TransactOpts)
}

// ExecuteAction is a paid mutator transaction binding the contract method 0x8f39d13a.
//
// Solidity: function executeAction(address _action, bytes _data) returns()
func (_MemecoinTreasury *MemecoinTreasuryTransactor) ExecuteAction(opts *bind.TransactOpts, _action common.Address, _data []byte) (*types.Transaction, error) {
	return _MemecoinTreasury.contract.Transact(opts, "executeAction", _action, _data)
}

// ExecuteAction is a paid mutator transaction binding the contract method 0x8f39d13a.
//
// Solidity: function executeAction(address _action, bytes _data) returns()
func (_MemecoinTreasury *MemecoinTreasurySession) ExecuteAction(_action common.Address, _data []byte) (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.ExecuteAction(&_MemecoinTreasury.TransactOpts, _action, _data)
}

// ExecuteAction is a paid mutator transaction binding the contract method 0x8f39d13a.
//
// Solidity: function executeAction(address _action, bytes _data) returns()
func (_MemecoinTreasury *MemecoinTreasuryTransactorSession) ExecuteAction(_action common.Address, _data []byte) (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.ExecuteAction(&_MemecoinTreasury.TransactOpts, _action, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0x8257d1cc.
//
// Solidity: function initialize(address _positionManager, address _actionManager, address _nativeToken, (address,address,uint24,int24,address) _poolKey) returns()
func (_MemecoinTreasury *MemecoinTreasuryTransactor) Initialize(opts *bind.TransactOpts, _positionManager common.Address, _actionManager common.Address, _nativeToken common.Address, _poolKey PoolKey) (*types.Transaction, error) {
	return _MemecoinTreasury.contract.Transact(opts, "initialize", _positionManager, _actionManager, _nativeToken, _poolKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x8257d1cc.
//
// Solidity: function initialize(address _positionManager, address _actionManager, address _nativeToken, (address,address,uint24,int24,address) _poolKey) returns()
func (_MemecoinTreasury *MemecoinTreasurySession) Initialize(_positionManager common.Address, _actionManager common.Address, _nativeToken common.Address, _poolKey PoolKey) (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.Initialize(&_MemecoinTreasury.TransactOpts, _positionManager, _actionManager, _nativeToken, _poolKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x8257d1cc.
//
// Solidity: function initialize(address _positionManager, address _actionManager, address _nativeToken, (address,address,uint24,int24,address) _poolKey) returns()
func (_MemecoinTreasury *MemecoinTreasuryTransactorSession) Initialize(_positionManager common.Address, _actionManager common.Address, _nativeToken common.Address, _poolKey PoolKey) (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.Initialize(&_MemecoinTreasury.TransactOpts, _positionManager, _actionManager, _nativeToken, _poolKey)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MemecoinTreasury *MemecoinTreasuryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemecoinTreasury.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MemecoinTreasury *MemecoinTreasurySession) Receive() (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.Receive(&_MemecoinTreasury.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MemecoinTreasury *MemecoinTreasuryTransactorSession) Receive() (*types.Transaction, error) {
	return _MemecoinTreasury.Contract.Receive(&_MemecoinTreasury.TransactOpts)
}

// MemecoinTreasuryActionExecutedIterator is returned from FilterActionExecuted and is used to iterate over the raw logs and unpacked data for ActionExecuted events raised by the MemecoinTreasury contract.
type MemecoinTreasuryActionExecutedIterator struct {
	Event *MemecoinTreasuryActionExecuted // Event containing the contract specifics and raw log

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
func (it *MemecoinTreasuryActionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemecoinTreasuryActionExecuted)
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
		it.Event = new(MemecoinTreasuryActionExecuted)
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
func (it *MemecoinTreasuryActionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemecoinTreasuryActionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemecoinTreasuryActionExecuted represents a ActionExecuted event raised by the MemecoinTreasury contract.
type MemecoinTreasuryActionExecuted struct {
	Action  common.Address
	PoolKey PoolKey
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterActionExecuted is a free log retrieval operation binding the contract event 0x32ea44c9b28c6fcd233c31b0c0ef74ef50bdbe347bcbb0775445b50056bbe3e9.
//
// Solidity: event ActionExecuted(address indexed _action, (address,address,uint24,int24,address) _poolKey, bytes _data)
func (_MemecoinTreasury *MemecoinTreasuryFilterer) FilterActionExecuted(opts *bind.FilterOpts, _action []common.Address) (*MemecoinTreasuryActionExecutedIterator, error) {

	var _actionRule []interface{}
	for _, _actionItem := range _action {
		_actionRule = append(_actionRule, _actionItem)
	}

	logs, sub, err := _MemecoinTreasury.contract.FilterLogs(opts, "ActionExecuted", _actionRule)
	if err != nil {
		return nil, err
	}
	return &MemecoinTreasuryActionExecutedIterator{contract: _MemecoinTreasury.contract, event: "ActionExecuted", logs: logs, sub: sub}, nil
}

// WatchActionExecuted is a free log subscription operation binding the contract event 0x32ea44c9b28c6fcd233c31b0c0ef74ef50bdbe347bcbb0775445b50056bbe3e9.
//
// Solidity: event ActionExecuted(address indexed _action, (address,address,uint24,int24,address) _poolKey, bytes _data)
func (_MemecoinTreasury *MemecoinTreasuryFilterer) WatchActionExecuted(opts *bind.WatchOpts, sink chan<- *MemecoinTreasuryActionExecuted, _action []common.Address) (event.Subscription, error) {

	var _actionRule []interface{}
	for _, _actionItem := range _action {
		_actionRule = append(_actionRule, _actionItem)
	}

	logs, sub, err := _MemecoinTreasury.contract.WatchLogs(opts, "ActionExecuted", _actionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemecoinTreasuryActionExecuted)
				if err := _MemecoinTreasury.contract.UnpackLog(event, "ActionExecuted", log); err != nil {
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

// ParseActionExecuted is a log parse operation binding the contract event 0x32ea44c9b28c6fcd233c31b0c0ef74ef50bdbe347bcbb0775445b50056bbe3e9.
//
// Solidity: event ActionExecuted(address indexed _action, (address,address,uint24,int24,address) _poolKey, bytes _data)
func (_MemecoinTreasury *MemecoinTreasuryFilterer) ParseActionExecuted(log types.Log) (*MemecoinTreasuryActionExecuted, error) {
	event := new(MemecoinTreasuryActionExecuted)
	if err := _MemecoinTreasury.contract.UnpackLog(event, "ActionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemecoinTreasuryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MemecoinTreasury contract.
type MemecoinTreasuryInitializedIterator struct {
	Event *MemecoinTreasuryInitialized // Event containing the contract specifics and raw log

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
func (it *MemecoinTreasuryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemecoinTreasuryInitialized)
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
		it.Event = new(MemecoinTreasuryInitialized)
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
func (it *MemecoinTreasuryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemecoinTreasuryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemecoinTreasuryInitialized represents a Initialized event raised by the MemecoinTreasury contract.
type MemecoinTreasuryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MemecoinTreasury *MemecoinTreasuryFilterer) FilterInitialized(opts *bind.FilterOpts) (*MemecoinTreasuryInitializedIterator, error) {

	logs, sub, err := _MemecoinTreasury.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MemecoinTreasuryInitializedIterator{contract: _MemecoinTreasury.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MemecoinTreasury *MemecoinTreasuryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MemecoinTreasuryInitialized) (event.Subscription, error) {

	logs, sub, err := _MemecoinTreasury.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemecoinTreasuryInitialized)
				if err := _MemecoinTreasury.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MemecoinTreasury *MemecoinTreasuryFilterer) ParseInitialized(log types.Log) (*MemecoinTreasuryInitialized, error) {
	event := new(MemecoinTreasuryInitialized)
	if err := _MemecoinTreasury.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
