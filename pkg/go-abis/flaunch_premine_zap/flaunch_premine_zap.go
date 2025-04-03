// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flaunch_premine_zap

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

// PositionManagerFlaunchParams is an auto generated low-level Go binding around an user-defined struct.
type PositionManagerFlaunchParams struct {
	Name                   string
	Symbol                 string
	TokenUri               string
	InitialTokenFairLaunch *big.Int
	PremineAmount          *big.Int
	Creator                common.Address
	CreatorFeeAllocation   *big.Int
	FlaunchAt              *big.Int
	InitialPriceParams     []byte
	FeeCalculatorParams    []byte
}

// FlaunchPremineZapMetaData contains all meta data concerning the FlaunchPremineZap contract.
var FlaunchPremineZapMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_positionManager\",\"type\":\"address\",\"internalType\":\"contractPositionManager\"},{\"name\":\"_flaunchContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_flETH\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolSwap\",\"type\":\"address\",\"internalType\":\"contractPoolSwap\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"calculateFee\",\"inputs\":[{\"name\":\"_premineAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_slippage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"ethRequired_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFLETH\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flaunch\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple\",\"internalType\":\"structPositionManager.FlaunchParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialTokenFairLaunch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"premineAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creatorFeeAllocation\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"flaunchAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeCalculatorParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"memecoin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ethSpent_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"flaunchContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractFlaunch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolSwap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPoolSwap\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPositionManager\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b506040516112bf3803806112bf83398101604081905261002f91610068565b6001600160a01b0393841660805291831660a052821660c0521660e0526100c4565b6001600160a01b0381168114610065575f5ffd5b50565b5f5f5f5f6080858703121561007b575f5ffd5b845161008681610051565b602086015190945061009781610051565b60408601519093506100a881610051565b60608601519092506100b981610051565b939692955090935050565b60805160a05160c05160e05161116f6101505f395f8181609d015281816108880152818161097801526109e801525f8181610170015281816106cd015281816107cd015281816108bb015261094b01525f61011001525f818160dd0152818161019601528181610227015281816102fb015281816103e40152818161045201526105c1015261116f5ff3fe608060405260043610610057575f3560e01c80632a12e4881461006257806333809ad01461008c578063791b98bc146100cc57806384aa1da0146100ff578063bbce359414610132578063f6baab4d1461015f575f5ffd5b3661005e57005b5f5ffd5b610075610070366004610b29565b610192565b604051610083929190610b67565b60405180910390f35b348015610097575f5ffd5b506100bf7f000000000000000000000000000000000000000000000000000000000000000081565b6040516100839190610b80565b3480156100d7575f5ffd5b506100bf7f000000000000000000000000000000000000000000000000000000000000000081565b34801561010a575f5ffd5b506100bf7f000000000000000000000000000000000000000000000000000000000000000081565b34801561013d575f5ffd5b5061015161014c366004610b94565b6102e9565b604051908152602001610083565b34801561016a575f5ffd5b506100bf7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632a12e48834856040518363ffffffff1660e01b81526004016101e19190610cbb565b60206040518083038185885af11580156101fd573d5f5f3e3d5ffd5b50505050506040513d601f19601f820116820180604052508101906102229190610dea565b91505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166355d1cb60846040518263ffffffff1660e01b81526004016102719190610b80565b60a060405180830381865afa15801561028c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102b09190610e05565b90506080840135156102d0576102cd8160808601353031866106ad565b91505b303180156102e2576102e23382610836565b5050915091565b5f5f680a18f07d736b90be55601d1b867f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ff3575e487876040518363ffffffff1660e01b8152600401610347929190610eae565b602060405180830381865afa158015610362573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103869190610ec9565b6103909190610ef4565b61039a9190610f11565b6040805160a0810182525f80825260208201819052818301819052603c606083015260808201819052915163071c4ddb60e41b815260016004820152929350916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906371c4ddb090602401602060405180830381865afa158015610429573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061044d9190610dea565b90505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663b3b4279561048a8560a0902090565b6040518263ffffffff1660e01b81526004016104a891815260200190565b608060405180830381865afa1580156104c3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104e79190610f3f565b5190506001600160a01b038216156105a857816001600160a01b031663403910858460405180606001604052805f151581526020016105258e610853565b815273fffd8963efd1fc6a506488495d951d5263988d266020909101526040516001600160e01b031960e085901b1681526105669291908690600401611034565b602060405180830381865afa158015610581573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105a59190611064565b90505b60405163ec876d7160e01b815284906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063ec876d71906105f8908b908b90600401610eae565b602060405180830381865afa158015610613573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106379190610ec9565b610641919061107f565b945062ffffff8116156106775761271061066062ffffff831686610ef4565b61066a9190610f11565b610674908661107f565b94505b87156106a15761271061068a8987610ef4565b6106949190610f11565b61069e908661107f565b94505b50505050949350505050565b60405163b6b55f2560e01b81525f60048201819052906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b6b55f259085906024015f604051808303818588803b158015610710575f5ffd5b505af1158015610722573d5f5f3e3d5ffd5b5050505050610732858585610871565b60405163a9059cbb60e01b81529091506001600160a01b0383169063a9059cbb906107639033908890600401610b67565b6020604051808303815f875af115801561077f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a39190611092565b505f6107af82856110ab565b9050801561082d57604051632e1a7d4d60e01b8152600481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d906024015f604051808303815f87803b158015610816575f5ffd5b505af1158015610828573d5f5f3e3d5ffd5b505050505b50949350505050565b5f385f3884865af161084f5763b12d13eb5f526004601cfd5b5050565b805f81121561086c5761086c6393dafdf160e01b610b21565b919050565b8251604051636eb1769f60e11b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811660248301525f927f0000000000000000000000000000000000000000000000000000000000000000821691168114159184919063dd62ed3e906044016020604051808303815f875af115801561090a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061092e9190610ec9565b10156109e55760405163095ea7b360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063095ea7b3906109a3907f0000000000000000000000000000000000000000000000000000000000000000905f1990600401610b67565b6020604051808303815f875af11580156109bf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109e39190611092565b505b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631e2817de876040518060600160405280861515158152602001610a338a610853565b81526020018615610a6257610a5d600173fffd8963efd1fc6a506488495d951d5263988d266110be565b610a72565b610a726401000276a360016110dd565b6001600160a01b03168152506040518363ffffffff1660e01b8152600401610a9b9291906110fc565b6020604051808303815f875af1158015610ab7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610adb9190610ec9565b90508115610afb57610aed81600f0b90565b610af690611118565b610b0e565b610b058160801d90565b610b0e90611118565b6001600160801b03169695505050505050565b805f5260045ffd5b5f60208284031215610b39575f5ffd5b81356001600160401b03811115610b4e575f5ffd5b82016101408185031215610b60575f5ffd5b9392505050565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0391909116815260200190565b5f5f5f5f60608587031215610ba7575f5ffd5b843593506020850135925060408501356001600160401b03811115610bca575f5ffd5b8501601f81018713610bda575f5ffd5b80356001600160401b03811115610bef575f5ffd5b876020828401011115610c00575f5ffd5b949793965060200194505050565b5f5f8335601e19843603018112610c23575f5ffd5b83016020810192503590506001600160401b03811115610c41575f5ffd5b803603821315610c4f575f5ffd5b9250929050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b6001600160a01b0381168114610c92575f5ffd5b50565b803561086c81610c7e565b62ffffff81168114610c92575f5ffd5b803561086c81610ca0565b602081525f610cca8384610c0e565b6101406020850152610ce161016085018284610c56565b915050610cf16020850185610c0e565b848303601f19016040860152610d08838284610c56565b92505050610d196040850185610c0e565b848303601f19016060860152610d30838284610c56565b606087013560808781019190915287013560a080880191909152909350610d5b925086019050610c95565b6001600160a01b03811660c085015250610d7760c08501610cb0565b62ffffff811660e08501525060e084013561010084810191909152610d9e90850185610c0e565b848303601f1901610120860152610db6838284610c56565b92505050610dc8610120850185610c0e565b848303601f1901610140860152610de0838284610c56565b9695505050505050565b5f60208284031215610dfa575f5ffd5b8151610b6081610c7e565b5f60a0828403128015610e16575f5ffd5b5060405160a081016001600160401b0381118282101715610e4557634e487b7160e01b5f52604160045260245ffd5b6040528251610e5381610c7e565b81526020830151610e6381610c7e565b60208201526040830151610e7681610ca0565b60408201526060830151600281900b8114610e8f575f5ffd5b60608201526080830151610ea281610c7e565b60808201529392505050565b602081525f610ec1602083018486610c56565b949350505050565b5f60208284031215610ed9575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b8082028115828204841417610f0b57610f0b610ee0565b92915050565b5f82610f2b57634e487b7160e01b5f52601260045260245ffd5b500490565b8051801515811461086c575f5ffd5b5f6080828403128015610f50575f5ffd5b50604051608081016001600160401b0381118282101715610f7f57634e487b7160e01b5f52604160045260245ffd5b6040528251610f8d81610ca0565b81526020830151610f9d81610ca0565b60208201526040830151610fb081610ca0565b6040820152610fc160608401610f30565b60608201529392505050565b80516001600160a01b03908116835260208083015182169084015260408083015162ffffff169084015260608083015160020b9084015260809182015116910152565b805115158252602080820151908301526040908101516001600160a01b0316910152565b61012081016110438286610fcd565b61105060a0830185611010565b62ffffff8316610100830152949350505050565b5f60208284031215611074575f5ffd5b8151610b6081610ca0565b80820180821115610f0b57610f0b610ee0565b5f602082840312156110a2575f5ffd5b610b6082610f30565b81810381811115610f0b57610f0b610ee0565b6001600160a01b038281168282160390811115610f0b57610f0b610ee0565b6001600160a01b038181168382160190811115610f0b57610f0b610ee0565b610100810161110b8285610fcd565b610b6060a0830184611010565b5f600f82900b6001607f1b810161113157611131610ee0565b5f039291505056fea2646970667358221220c7da11ae8d7423e764239fca6007f1ede0dca03158ba686f16593addaddc750864736f6c634300081b0033",
}

// FlaunchPremineZapABI is the input ABI used to generate the binding from.
// Deprecated: Use FlaunchPremineZapMetaData.ABI instead.
var FlaunchPremineZapABI = FlaunchPremineZapMetaData.ABI

// FlaunchPremineZapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FlaunchPremineZapMetaData.Bin instead.
var FlaunchPremineZapBin = FlaunchPremineZapMetaData.Bin

// DeployFlaunchPremineZap deploys a new Ethereum contract, binding an instance of FlaunchPremineZap to it.
func DeployFlaunchPremineZap(auth *bind.TransactOpts, backend bind.ContractBackend, _positionManager common.Address, _flaunchContract common.Address, _flETH common.Address, _poolSwap common.Address) (common.Address, *types.Transaction, *FlaunchPremineZap, error) {
	parsed, err := FlaunchPremineZapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FlaunchPremineZapBin), backend, _positionManager, _flaunchContract, _flETH, _poolSwap)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FlaunchPremineZap{FlaunchPremineZapCaller: FlaunchPremineZapCaller{contract: contract}, FlaunchPremineZapTransactor: FlaunchPremineZapTransactor{contract: contract}, FlaunchPremineZapFilterer: FlaunchPremineZapFilterer{contract: contract}}, nil
}

// FlaunchPremineZap is an auto generated Go binding around an Ethereum contract.
type FlaunchPremineZap struct {
	FlaunchPremineZapCaller     // Read-only binding to the contract
	FlaunchPremineZapTransactor // Write-only binding to the contract
	FlaunchPremineZapFilterer   // Log filterer for contract events
}

// FlaunchPremineZapCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlaunchPremineZapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlaunchPremineZapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlaunchPremineZapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlaunchPremineZapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlaunchPremineZapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlaunchPremineZapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlaunchPremineZapSession struct {
	Contract     *FlaunchPremineZap // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FlaunchPremineZapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlaunchPremineZapCallerSession struct {
	Contract *FlaunchPremineZapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FlaunchPremineZapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlaunchPremineZapTransactorSession struct {
	Contract     *FlaunchPremineZapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FlaunchPremineZapRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlaunchPremineZapRaw struct {
	Contract *FlaunchPremineZap // Generic contract binding to access the raw methods on
}

// FlaunchPremineZapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlaunchPremineZapCallerRaw struct {
	Contract *FlaunchPremineZapCaller // Generic read-only contract binding to access the raw methods on
}

// FlaunchPremineZapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlaunchPremineZapTransactorRaw struct {
	Contract *FlaunchPremineZapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlaunchPremineZap creates a new instance of FlaunchPremineZap, bound to a specific deployed contract.
func NewFlaunchPremineZap(address common.Address, backend bind.ContractBackend) (*FlaunchPremineZap, error) {
	contract, err := bindFlaunchPremineZap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlaunchPremineZap{FlaunchPremineZapCaller: FlaunchPremineZapCaller{contract: contract}, FlaunchPremineZapTransactor: FlaunchPremineZapTransactor{contract: contract}, FlaunchPremineZapFilterer: FlaunchPremineZapFilterer{contract: contract}}, nil
}

// NewFlaunchPremineZapCaller creates a new read-only instance of FlaunchPremineZap, bound to a specific deployed contract.
func NewFlaunchPremineZapCaller(address common.Address, caller bind.ContractCaller) (*FlaunchPremineZapCaller, error) {
	contract, err := bindFlaunchPremineZap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlaunchPremineZapCaller{contract: contract}, nil
}

// NewFlaunchPremineZapTransactor creates a new write-only instance of FlaunchPremineZap, bound to a specific deployed contract.
func NewFlaunchPremineZapTransactor(address common.Address, transactor bind.ContractTransactor) (*FlaunchPremineZapTransactor, error) {
	contract, err := bindFlaunchPremineZap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlaunchPremineZapTransactor{contract: contract}, nil
}

// NewFlaunchPremineZapFilterer creates a new log filterer instance of FlaunchPremineZap, bound to a specific deployed contract.
func NewFlaunchPremineZapFilterer(address common.Address, filterer bind.ContractFilterer) (*FlaunchPremineZapFilterer, error) {
	contract, err := bindFlaunchPremineZap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlaunchPremineZapFilterer{contract: contract}, nil
}

// bindFlaunchPremineZap binds a generic wrapper to an already deployed contract.
func bindFlaunchPremineZap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlaunchPremineZapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlaunchPremineZap *FlaunchPremineZapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlaunchPremineZap.Contract.FlaunchPremineZapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlaunchPremineZap *FlaunchPremineZapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlaunchPremineZap.Contract.FlaunchPremineZapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlaunchPremineZap *FlaunchPremineZapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlaunchPremineZap.Contract.FlaunchPremineZapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlaunchPremineZap *FlaunchPremineZapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlaunchPremineZap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlaunchPremineZap *FlaunchPremineZapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlaunchPremineZap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlaunchPremineZap *FlaunchPremineZapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlaunchPremineZap.Contract.contract.Transact(opts, method, params...)
}

// CalculateFee is a free data retrieval call binding the contract method 0xbbce3594.
//
// Solidity: function calculateFee(uint256 _premineAmount, uint256 _slippage, bytes _initialPriceParams) view returns(uint256 ethRequired_)
func (_FlaunchPremineZap *FlaunchPremineZapCaller) CalculateFee(opts *bind.CallOpts, _premineAmount *big.Int, _slippage *big.Int, _initialPriceParams []byte) (*big.Int, error) {
	var out []interface{}
	err := _FlaunchPremineZap.contract.Call(opts, &out, "calculateFee", _premineAmount, _slippage, _initialPriceParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFee is a free data retrieval call binding the contract method 0xbbce3594.
//
// Solidity: function calculateFee(uint256 _premineAmount, uint256 _slippage, bytes _initialPriceParams) view returns(uint256 ethRequired_)
func (_FlaunchPremineZap *FlaunchPremineZapSession) CalculateFee(_premineAmount *big.Int, _slippage *big.Int, _initialPriceParams []byte) (*big.Int, error) {
	return _FlaunchPremineZap.Contract.CalculateFee(&_FlaunchPremineZap.CallOpts, _premineAmount, _slippage, _initialPriceParams)
}

// CalculateFee is a free data retrieval call binding the contract method 0xbbce3594.
//
// Solidity: function calculateFee(uint256 _premineAmount, uint256 _slippage, bytes _initialPriceParams) view returns(uint256 ethRequired_)
func (_FlaunchPremineZap *FlaunchPremineZapCallerSession) CalculateFee(_premineAmount *big.Int, _slippage *big.Int, _initialPriceParams []byte) (*big.Int, error) {
	return _FlaunchPremineZap.Contract.CalculateFee(&_FlaunchPremineZap.CallOpts, _premineAmount, _slippage, _initialPriceParams)
}

// FlETH is a free data retrieval call binding the contract method 0xf6baab4d.
//
// Solidity: function flETH() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapCaller) FlETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlaunchPremineZap.contract.Call(opts, &out, "flETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlETH is a free data retrieval call binding the contract method 0xf6baab4d.
//
// Solidity: function flETH() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapSession) FlETH() (common.Address, error) {
	return _FlaunchPremineZap.Contract.FlETH(&_FlaunchPremineZap.CallOpts)
}

// FlETH is a free data retrieval call binding the contract method 0xf6baab4d.
//
// Solidity: function flETH() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapCallerSession) FlETH() (common.Address, error) {
	return _FlaunchPremineZap.Contract.FlETH(&_FlaunchPremineZap.CallOpts)
}

// FlaunchContract is a free data retrieval call binding the contract method 0x84aa1da0.
//
// Solidity: function flaunchContract() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapCaller) FlaunchContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlaunchPremineZap.contract.Call(opts, &out, "flaunchContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlaunchContract is a free data retrieval call binding the contract method 0x84aa1da0.
//
// Solidity: function flaunchContract() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapSession) FlaunchContract() (common.Address, error) {
	return _FlaunchPremineZap.Contract.FlaunchContract(&_FlaunchPremineZap.CallOpts)
}

// FlaunchContract is a free data retrieval call binding the contract method 0x84aa1da0.
//
// Solidity: function flaunchContract() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapCallerSession) FlaunchContract() (common.Address, error) {
	return _FlaunchPremineZap.Contract.FlaunchContract(&_FlaunchPremineZap.CallOpts)
}

// PoolSwap is a free data retrieval call binding the contract method 0x33809ad0.
//
// Solidity: function poolSwap() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapCaller) PoolSwap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlaunchPremineZap.contract.Call(opts, &out, "poolSwap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolSwap is a free data retrieval call binding the contract method 0x33809ad0.
//
// Solidity: function poolSwap() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapSession) PoolSwap() (common.Address, error) {
	return _FlaunchPremineZap.Contract.PoolSwap(&_FlaunchPremineZap.CallOpts)
}

// PoolSwap is a free data retrieval call binding the contract method 0x33809ad0.
//
// Solidity: function poolSwap() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapCallerSession) PoolSwap() (common.Address, error) {
	return _FlaunchPremineZap.Contract.PoolSwap(&_FlaunchPremineZap.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapCaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlaunchPremineZap.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapSession) PositionManager() (common.Address, error) {
	return _FlaunchPremineZap.Contract.PositionManager(&_FlaunchPremineZap.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_FlaunchPremineZap *FlaunchPremineZapCallerSession) PositionManager() (common.Address, error) {
	return _FlaunchPremineZap.Contract.PositionManager(&_FlaunchPremineZap.CallOpts)
}

// Flaunch is a paid mutator transaction binding the contract method 0x2a12e488.
//
// Solidity: function flaunch((string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params) payable returns(address memecoin_, uint256 ethSpent_)
func (_FlaunchPremineZap *FlaunchPremineZapTransactor) Flaunch(opts *bind.TransactOpts, _params PositionManagerFlaunchParams) (*types.Transaction, error) {
	return _FlaunchPremineZap.contract.Transact(opts, "flaunch", _params)
}

// Flaunch is a paid mutator transaction binding the contract method 0x2a12e488.
//
// Solidity: function flaunch((string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params) payable returns(address memecoin_, uint256 ethSpent_)
func (_FlaunchPremineZap *FlaunchPremineZapSession) Flaunch(_params PositionManagerFlaunchParams) (*types.Transaction, error) {
	return _FlaunchPremineZap.Contract.Flaunch(&_FlaunchPremineZap.TransactOpts, _params)
}

// Flaunch is a paid mutator transaction binding the contract method 0x2a12e488.
//
// Solidity: function flaunch((string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params) payable returns(address memecoin_, uint256 ethSpent_)
func (_FlaunchPremineZap *FlaunchPremineZapTransactorSession) Flaunch(_params PositionManagerFlaunchParams) (*types.Transaction, error) {
	return _FlaunchPremineZap.Contract.Flaunch(&_FlaunchPremineZap.TransactOpts, _params)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlaunchPremineZap *FlaunchPremineZapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlaunchPremineZap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlaunchPremineZap *FlaunchPremineZapSession) Receive() (*types.Transaction, error) {
	return _FlaunchPremineZap.Contract.Receive(&_FlaunchPremineZap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlaunchPremineZap *FlaunchPremineZapTransactorSession) Receive() (*types.Transaction, error) {
	return _FlaunchPremineZap.Contract.Receive(&_FlaunchPremineZap.TransactOpts)
}
