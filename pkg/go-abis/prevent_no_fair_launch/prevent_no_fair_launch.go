// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package prevent_no_fair_launch

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

// PreventNoFairLaunchMetaData contains all meta data concerning the PreventNoFairLaunch contract.
var PreventNoFairLaunchMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_notifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MINIMUM_INITIAL_TOKENS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"notifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"notify\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_key\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"subscribe\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unsubscribe\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"InvalidInitialTokenFairLaunch\",\"inputs\":[{\"name\":\"_invalidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidNotifier\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_validNotifier\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b506040516106ef3803806106ef83398101604081905261002e9161003f565b6001600160a01b031660805261006c565b5f6020828403121561004f575f5ffd5b81516001600160a01b0381168114610065575f5ffd5b9392505050565b6080516106426100ad5f395f8181605e015281816101090152818161013101528181610188015281816101b00152818161026f015261029701526106425ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806309276ea414610059578063334f33231461009d5780634610b70a146100bd5780634ca1f062146100e0578063fcae4484146100f5575b5f5ffd5b6100807f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100af676765c793fa10079d601b1b81565b604051908152602001610094565b6100d06100cb36600461039f565b6100fd565b6040519015158152602001610094565b6100f36100ee3660046103d8565b61017d565b005b6100f3610264565b5f336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461017457337f00000000000000000000000000000000000000000000000000000000000000006040516316ff6b4b60e11b815260040161016b929190610467565b60405180910390fd5b5060015b919050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101ea57337f00000000000000000000000000000000000000000000000000000000000000006040516316ff6b4b60e11b815260040161016b929190610467565b639018191560e01b6001600160e01b031984160161025e575f61020f828401846104a9565b915050676765c793fa10079d601b1b8160600151101561025c576060810151604051630a60f3b960e11b81526004810191909152676765c793fa10079d601b1b602482015260440161016b565b505b50505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146102d157337f00000000000000000000000000000000000000000000000000000000000000006040516316ff6b4b60e11b815260040161016b929190610467565b565b634e487b7160e01b5f52604160045260245ffd5b60405161014081016001600160401b038111828210171561030a5761030a6102d3565b60405290565b5f82601f83011261031f575f5ffd5b8135602083015f806001600160401b0384111561033e5761033e6102d3565b50604051601f19601f85018116603f011681018181106001600160401b038211171561036c5761036c6102d3565b604052838152905080828401871015610383575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f602082840312156103af575f5ffd5b81356001600160401b038111156103c4575f5ffd5b6103d084828501610310565b949350505050565b5f5f5f5f606085870312156103eb575f5ffd5b8435935060208501356001600160e01b031981168114610409575f5ffd5b925060408501356001600160401b03811115610423575f5ffd5b8501601f81018713610433575f5ffd5b80356001600160401b03811115610448575f5ffd5b876020828401011115610459575f5ffd5b949793965060200194505050565b6001600160a01b0392831681529116602082015260400190565b80356001600160a01b0381168114610178575f5ffd5b803562ffffff81168114610178575f5ffd5b5f5f604083850312156104ba575f5ffd5b8235915060208301356001600160401b038111156104d6575f5ffd5b830161014081860312156104e8575f5ffd5b6104f06102e7565b81356001600160401b03811115610505575f5ffd5b61051187828501610310565b82525060208201356001600160401b0381111561052c575f5ffd5b61053887828501610310565b60208301525060408201356001600160401b03811115610556575f5ffd5b61056287828501610310565b604083015250606082810135908201526080808301359082015261058860a08301610481565b60a082015261059960c08301610497565b60c082015260e082810135908201526101008201356001600160401b038111156105c1575f5ffd5b6105cd87828501610310565b610100830152506101208201356001600160401b038111156105ed575f5ffd5b6105f987828501610310565b610120830152508092505050925092905056fea26469706673582212208738ccbeb888039c2b8d94621e34050b21ef5d19c7a2bd110544b1a933d3db5a64736f6c634300081b0033",
}

// PreventNoFairLaunchABI is the input ABI used to generate the binding from.
// Deprecated: Use PreventNoFairLaunchMetaData.ABI instead.
var PreventNoFairLaunchABI = PreventNoFairLaunchMetaData.ABI

// PreventNoFairLaunchBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PreventNoFairLaunchMetaData.Bin instead.
var PreventNoFairLaunchBin = PreventNoFairLaunchMetaData.Bin

// DeployPreventNoFairLaunch deploys a new Ethereum contract, binding an instance of PreventNoFairLaunch to it.
func DeployPreventNoFairLaunch(auth *bind.TransactOpts, backend bind.ContractBackend, _notifier common.Address) (common.Address, *types.Transaction, *PreventNoFairLaunch, error) {
	parsed, err := PreventNoFairLaunchMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PreventNoFairLaunchBin), backend, _notifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PreventNoFairLaunch{PreventNoFairLaunchCaller: PreventNoFairLaunchCaller{contract: contract}, PreventNoFairLaunchTransactor: PreventNoFairLaunchTransactor{contract: contract}, PreventNoFairLaunchFilterer: PreventNoFairLaunchFilterer{contract: contract}}, nil
}

// PreventNoFairLaunch is an auto generated Go binding around an Ethereum contract.
type PreventNoFairLaunch struct {
	PreventNoFairLaunchCaller     // Read-only binding to the contract
	PreventNoFairLaunchTransactor // Write-only binding to the contract
	PreventNoFairLaunchFilterer   // Log filterer for contract events
}

// PreventNoFairLaunchCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreventNoFairLaunchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreventNoFairLaunchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreventNoFairLaunchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreventNoFairLaunchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreventNoFairLaunchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreventNoFairLaunchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreventNoFairLaunchSession struct {
	Contract     *PreventNoFairLaunch // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PreventNoFairLaunchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreventNoFairLaunchCallerSession struct {
	Contract *PreventNoFairLaunchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PreventNoFairLaunchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreventNoFairLaunchTransactorSession struct {
	Contract     *PreventNoFairLaunchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PreventNoFairLaunchRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreventNoFairLaunchRaw struct {
	Contract *PreventNoFairLaunch // Generic contract binding to access the raw methods on
}

// PreventNoFairLaunchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreventNoFairLaunchCallerRaw struct {
	Contract *PreventNoFairLaunchCaller // Generic read-only contract binding to access the raw methods on
}

// PreventNoFairLaunchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreventNoFairLaunchTransactorRaw struct {
	Contract *PreventNoFairLaunchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreventNoFairLaunch creates a new instance of PreventNoFairLaunch, bound to a specific deployed contract.
func NewPreventNoFairLaunch(address common.Address, backend bind.ContractBackend) (*PreventNoFairLaunch, error) {
	contract, err := bindPreventNoFairLaunch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PreventNoFairLaunch{PreventNoFairLaunchCaller: PreventNoFairLaunchCaller{contract: contract}, PreventNoFairLaunchTransactor: PreventNoFairLaunchTransactor{contract: contract}, PreventNoFairLaunchFilterer: PreventNoFairLaunchFilterer{contract: contract}}, nil
}

// NewPreventNoFairLaunchCaller creates a new read-only instance of PreventNoFairLaunch, bound to a specific deployed contract.
func NewPreventNoFairLaunchCaller(address common.Address, caller bind.ContractCaller) (*PreventNoFairLaunchCaller, error) {
	contract, err := bindPreventNoFairLaunch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreventNoFairLaunchCaller{contract: contract}, nil
}

// NewPreventNoFairLaunchTransactor creates a new write-only instance of PreventNoFairLaunch, bound to a specific deployed contract.
func NewPreventNoFairLaunchTransactor(address common.Address, transactor bind.ContractTransactor) (*PreventNoFairLaunchTransactor, error) {
	contract, err := bindPreventNoFairLaunch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreventNoFairLaunchTransactor{contract: contract}, nil
}

// NewPreventNoFairLaunchFilterer creates a new log filterer instance of PreventNoFairLaunch, bound to a specific deployed contract.
func NewPreventNoFairLaunchFilterer(address common.Address, filterer bind.ContractFilterer) (*PreventNoFairLaunchFilterer, error) {
	contract, err := bindPreventNoFairLaunch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreventNoFairLaunchFilterer{contract: contract}, nil
}

// bindPreventNoFairLaunch binds a generic wrapper to an already deployed contract.
func bindPreventNoFairLaunch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PreventNoFairLaunchMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreventNoFairLaunch *PreventNoFairLaunchRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreventNoFairLaunch.Contract.PreventNoFairLaunchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreventNoFairLaunch *PreventNoFairLaunchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreventNoFairLaunch.Contract.PreventNoFairLaunchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreventNoFairLaunch *PreventNoFairLaunchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreventNoFairLaunch.Contract.PreventNoFairLaunchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreventNoFairLaunch *PreventNoFairLaunchCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreventNoFairLaunch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreventNoFairLaunch *PreventNoFairLaunchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreventNoFairLaunch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreventNoFairLaunch *PreventNoFairLaunchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreventNoFairLaunch.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMINITIALTOKENS is a free data retrieval call binding the contract method 0x334f3323.
//
// Solidity: function MINIMUM_INITIAL_TOKENS() view returns(uint256)
func (_PreventNoFairLaunch *PreventNoFairLaunchCaller) MINIMUMINITIALTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PreventNoFairLaunch.contract.Call(opts, &out, "MINIMUM_INITIAL_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMINITIALTOKENS is a free data retrieval call binding the contract method 0x334f3323.
//
// Solidity: function MINIMUM_INITIAL_TOKENS() view returns(uint256)
func (_PreventNoFairLaunch *PreventNoFairLaunchSession) MINIMUMINITIALTOKENS() (*big.Int, error) {
	return _PreventNoFairLaunch.Contract.MINIMUMINITIALTOKENS(&_PreventNoFairLaunch.CallOpts)
}

// MINIMUMINITIALTOKENS is a free data retrieval call binding the contract method 0x334f3323.
//
// Solidity: function MINIMUM_INITIAL_TOKENS() view returns(uint256)
func (_PreventNoFairLaunch *PreventNoFairLaunchCallerSession) MINIMUMINITIALTOKENS() (*big.Int, error) {
	return _PreventNoFairLaunch.Contract.MINIMUMINITIALTOKENS(&_PreventNoFairLaunch.CallOpts)
}

// Notifier is a free data retrieval call binding the contract method 0x09276ea4.
//
// Solidity: function notifier() view returns(address)
func (_PreventNoFairLaunch *PreventNoFairLaunchCaller) Notifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PreventNoFairLaunch.contract.Call(opts, &out, "notifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Notifier is a free data retrieval call binding the contract method 0x09276ea4.
//
// Solidity: function notifier() view returns(address)
func (_PreventNoFairLaunch *PreventNoFairLaunchSession) Notifier() (common.Address, error) {
	return _PreventNoFairLaunch.Contract.Notifier(&_PreventNoFairLaunch.CallOpts)
}

// Notifier is a free data retrieval call binding the contract method 0x09276ea4.
//
// Solidity: function notifier() view returns(address)
func (_PreventNoFairLaunch *PreventNoFairLaunchCallerSession) Notifier() (common.Address, error) {
	return _PreventNoFairLaunch.Contract.Notifier(&_PreventNoFairLaunch.CallOpts)
}

// Notify is a free data retrieval call binding the contract method 0x4ca1f062.
//
// Solidity: function notify(bytes32 , bytes4 _key, bytes _data) view returns()
func (_PreventNoFairLaunch *PreventNoFairLaunchCaller) Notify(opts *bind.CallOpts, arg0 [32]byte, _key [4]byte, _data []byte) error {
	var out []interface{}
	err := _PreventNoFairLaunch.contract.Call(opts, &out, "notify", arg0, _key, _data)

	if err != nil {
		return err
	}

	return err

}

// Notify is a free data retrieval call binding the contract method 0x4ca1f062.
//
// Solidity: function notify(bytes32 , bytes4 _key, bytes _data) view returns()
func (_PreventNoFairLaunch *PreventNoFairLaunchSession) Notify(arg0 [32]byte, _key [4]byte, _data []byte) error {
	return _PreventNoFairLaunch.Contract.Notify(&_PreventNoFairLaunch.CallOpts, arg0, _key, _data)
}

// Notify is a free data retrieval call binding the contract method 0x4ca1f062.
//
// Solidity: function notify(bytes32 , bytes4 _key, bytes _data) view returns()
func (_PreventNoFairLaunch *PreventNoFairLaunchCallerSession) Notify(arg0 [32]byte, _key [4]byte, _data []byte) error {
	return _PreventNoFairLaunch.Contract.Notify(&_PreventNoFairLaunch.CallOpts, arg0, _key, _data)
}

// Subscribe is a free data retrieval call binding the contract method 0x4610b70a.
//
// Solidity: function subscribe(bytes ) view returns(bool)
func (_PreventNoFairLaunch *PreventNoFairLaunchCaller) Subscribe(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _PreventNoFairLaunch.contract.Call(opts, &out, "subscribe", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Subscribe is a free data retrieval call binding the contract method 0x4610b70a.
//
// Solidity: function subscribe(bytes ) view returns(bool)
func (_PreventNoFairLaunch *PreventNoFairLaunchSession) Subscribe(arg0 []byte) (bool, error) {
	return _PreventNoFairLaunch.Contract.Subscribe(&_PreventNoFairLaunch.CallOpts, arg0)
}

// Subscribe is a free data retrieval call binding the contract method 0x4610b70a.
//
// Solidity: function subscribe(bytes ) view returns(bool)
func (_PreventNoFairLaunch *PreventNoFairLaunchCallerSession) Subscribe(arg0 []byte) (bool, error) {
	return _PreventNoFairLaunch.Contract.Subscribe(&_PreventNoFairLaunch.CallOpts, arg0)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_PreventNoFairLaunch *PreventNoFairLaunchTransactor) Unsubscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreventNoFairLaunch.contract.Transact(opts, "unsubscribe")
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_PreventNoFairLaunch *PreventNoFairLaunchSession) Unsubscribe() (*types.Transaction, error) {
	return _PreventNoFairLaunch.Contract.Unsubscribe(&_PreventNoFairLaunch.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_PreventNoFairLaunch *PreventNoFairLaunchTransactorSession) Unsubscribe() (*types.Transaction, error) {
	return _PreventNoFairLaunch.Contract.Unsubscribe(&_PreventNoFairLaunch.TransactOpts)
}
