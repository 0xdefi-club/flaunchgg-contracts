// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package referral_escrow

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

// ReferralEscrowMetaData contains all meta data concerning the ReferralEscrow contract.
var ReferralEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_nativeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_positionManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allocations\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assignTokens\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimAndSwap\",\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_sqrtPriceX96Limits\",\"type\":\"uint160[]\",\"internalType\":\"uint160[]\"},{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimTokens\",\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"nativeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolSwap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPoolSwap\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setPoolSwap\",\"inputs\":[{\"name\":\"_poolSwap\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensAssigned\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensClaimed\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensSwapped\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_tokenIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_ethOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedTokensAndLimits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b506040516110d53803806110d583398101604081905261002e916100c4565b6001600160a01b03808316608052811660a05261004a33610051565b50506100f5565b638b78c6d81980541561006b57630dc149f05f526004601cfd5b6001600160a01b03909116801560ff1b8117909155805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b80516001600160a01b03811681146100bf575f5ffd5b919050565b5f5f604083850312156100d5575f5ffd5b6100de836100a9565b91506100ec602084016100a9565b90509250929050565b60805160a051610f806101555f395f8181610169015281816105980152610aef01525f818161023601528181610421015281816104db0152818161051201528181610553015281816107a80152818161091901526109550152610f805ff3fe6080604052600436106100ba575f3560e01c806325692962146100c557806333809ad0146100cf57806354d1f13d1461010a578063715018a614610112578063761032ec1461011a57806378c5ae8c14610139578063791b98bc14610158578063853996e41461018b5780638da5cb5b146101cf578063988a18f6146101e7578063b0f4c2ea14610206578063e1758bd814610225578063f04e283e14610258578063f2fde38b1461026b578063fee81cf41461027e575f5ffd5b366100c157005b5f5ffd5b6100cd6102af565b005b3480156100da575f5ffd5b505f546100ed906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100cd6102fb565b6100cd610334565b348015610125575f5ffd5b506100cd610134366004610c9b565b610347565b348015610144575f5ffd5b506100cd610153366004610d04565b610370565b348015610163575f5ffd5b506100ed7f000000000000000000000000000000000000000000000000000000000000000081565b348015610196575f5ffd5b506101c16101a5366004610d85565b600160209081525f928352604080842090915290825290205481565b604051908152602001610101565b3480156101da575f5ffd5b50638b78c6d819546100ed565b3480156101f2575f5ffd5b506100cd610201366004610dbc565b61088c565b348015610211575f5ffd5b506100cd610220366004610e0e565b610ae4565b348015610230575f5ffd5b506100ed7f000000000000000000000000000000000000000000000000000000000000000081565b6100cd610266366004610c9b565b610bc4565b6100cd610279366004610c9b565b610c01565b348015610289575f5ffd5b506101c1610298366004610c9b565b63389a75e1600c9081525f91909152602090205490565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b61033c610c27565b6103455f610c41565b565b61034f610c27565b5f80546001600160a01b0319166001600160a01b0392909216919091179055565b838281111561039257604051639ff37e5d60e01b815260040160405180910390fd5b5f5f5f5f5f5b85811015610791578a8a828181106103b2576103b2610e53565b90506020020160208101906103c79190610c9b565b335f9081526001602090815260408083206001600160a01b03851684529091529020549095509350831561078957335f9081526001602090815260408083206001600160a01b0389811680865291909352908320929092557f000000000000000000000000000000000000000000000000000000000000000016900361044f578392506106fc565b5f5460405163095ea7b360e01b81526001600160a01b038781169263095ea7b39261048292909116908890600401610e67565b6020604051808303815f875af115801561049e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104c29190610e80565b505f80546040805160a081019091526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081168982161093921690631e2817de908085610536577f0000000000000000000000000000000000000000000000000000000000000000610538565b8a5b6001600160a01b0316815260200185610551578a610573565b7f00000000000000000000000000000000000000000000000000000000000000005b6001600160a01b031681526020015f62ffffff168152602001603c60020b81526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250604051806060016040528086151581526020018a6105e090610eb3565b81526020018f8f898181106105f7576105f7610e53565b905060200201602081019061060c9190610c9b565b6001600160a01b03908116909152604080516001600160e01b031960e087901b168152845183166004820152602080860151841660248301528583015162ffffff166044830152606086015160020b6064830152608090950151831660848201528351151560a48201529383015160c485015291909101511660e4820152610104016020604051808303815f875af11580156106aa573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106ce9190610ecd565b9050816106e4576106df8160801d90565b6106ee565b6106ee81600f0b90565b6001600160801b0316945050505b6107068383610ee4565b9150846001600160a01b0316336001600160a01b03165f516020610f2b5f395f51905f52898760405161073a929190610e67565b60405180910390a360408051858152602081018590526001600160a01b0387169133917f25f1d03755df23c30e25db2dbd3891e31ce084bdfbfc46f9fe5e446ee5f9b2d4910160405180910390a35b600101610398565b50604051632e1a7d4d60e01b8152600481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d906024015f604051808303815f87803b1580156107f1575f5ffd5b505af1158015610803573d5f5f3e3d5ffd5b505050505f866001600160a01b0316826040515f6040518083038185875af1925050503d805f8114610850576040519150601f19603f3d011682016040523d82523d5f602084013e610855565b606091505b505090508061087f5760405162461bcd60e51b815260040161087690610efd565b60405180910390fd5b5050505050505050505050565b5f5f5f5b84811015610adc578585828181106108aa576108aa610e53565b90506020020160208101906108bf9190610c9b565b335f9081526001602090815260408083206001600160a01b038516845290915290205490935091508115610ad457335f9081526001602090815260408083206001600160a01b0387811680865291909352908320929092557f0000000000000000000000000000000000000000000000000000000000000000169003610a2957604051632e1a7d4d60e01b8152600481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d906024015f604051808303815f87803b15801561099e575f5ffd5b505af11580156109b0573d5f5f3e3d5ffd5b505050505f846001600160a01b0316836040515f6040518083038185875af1925050503d805f81146109fd576040519150601f19603f3d011682016040523d82523d5f602084013e610a02565b606091505b5050905080610a235760405162461bcd60e51b815260040161087690610efd565b50610a99565b60405163a9059cbb60e01b81526001600160a01b0384169063a9059cbb90610a579087908690600401610e67565b6020604051808303815f875af1158015610a73573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a979190610e80565b505b826001600160a01b0316336001600160a01b03165f516020610f2b5f395f51905f528685604051610acb929190610e67565b60405180910390a35b600101610890565b505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610b2c576040516282b42960e81b815260040160405180910390fd5b8015610bbe576001600160a01b038084165f90815260016020908152604080832093861683529290529081208054839290610b68908490610ee4565b92505081905550816001600160a01b0316836001600160a01b0316857fc733341be84ca8f5e6ee2be63797abe1af7bbe5efcb5c8abe995dc8a51fd4b3784604051610bb591815260200190565b60405180910390a45b50505050565b610bcc610c27565b63389a75e1600c52805f526020600c208054421115610bf257636f5e88185f526004601cfd5b5f9055610bfe81610c41565b50565b610c09610c27565b8060601b610c1e57637448fbae5f526004601cfd5b610bfe81610c41565b638b78c6d819543314610345576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3811560ff1b8217905550565b6001600160a01b0381168114610bfe575f5ffd5b5f60208284031215610cab575f5ffd5b8135610cb681610c87565b9392505050565b5f5f83601f840112610ccd575f5ffd5b5081356001600160401b03811115610ce3575f5ffd5b6020830191508360208260051b8501011115610cfd575f5ffd5b9250929050565b5f5f5f5f5f60608688031215610d18575f5ffd5b85356001600160401b03811115610d2d575f5ffd5b610d3988828901610cbd565b90965094505060208601356001600160401b03811115610d57575f5ffd5b610d6388828901610cbd565b9094509250506040860135610d7781610c87565b809150509295509295909350565b5f5f60408385031215610d96575f5ffd5b8235610da181610c87565b91506020830135610db181610c87565b809150509250929050565b5f5f5f60408486031215610dce575f5ffd5b83356001600160401b03811115610de3575f5ffd5b610def86828701610cbd565b9094509250506020840135610e0381610c87565b809150509250925092565b5f5f5f5f60808587031215610e21575f5ffd5b843593506020850135610e3381610c87565b92506040850135610e4381610c87565b9396929550929360600135925050565b634e487b7160e01b5f52603260045260245ffd5b6001600160a01b03929092168252602082015260400190565b5f60208284031215610e90575f5ffd5b81518015158114610cb6575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b5f600160ff1b8201610ec757610ec7610e9f565b505f0390565b5f60208284031215610edd575f5ffd5b5051919050565b80820180821115610ef757610ef7610e9f565b92915050565b60208082526013908201527211551208151c985b9cd9995c8811985a5b1959606a1b60408201526060019056fe40b4b442b70a20081fa377001b642c7d3618847a45b5b8ed2b938e30a3299ddba2646970667358221220c69c1d783338cdc98526fb6ada5e863c9a92c1f68e78c7d3cc59c361468db71264736f6c634300081b0033",
}

// ReferralEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use ReferralEscrowMetaData.ABI instead.
var ReferralEscrowABI = ReferralEscrowMetaData.ABI

// ReferralEscrowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReferralEscrowMetaData.Bin instead.
var ReferralEscrowBin = ReferralEscrowMetaData.Bin

// DeployReferralEscrow deploys a new Ethereum contract, binding an instance of ReferralEscrow to it.
func DeployReferralEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, _nativeToken common.Address, _positionManager common.Address) (common.Address, *types.Transaction, *ReferralEscrow, error) {
	parsed, err := ReferralEscrowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReferralEscrowBin), backend, _nativeToken, _positionManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReferralEscrow{ReferralEscrowCaller: ReferralEscrowCaller{contract: contract}, ReferralEscrowTransactor: ReferralEscrowTransactor{contract: contract}, ReferralEscrowFilterer: ReferralEscrowFilterer{contract: contract}}, nil
}

// ReferralEscrow is an auto generated Go binding around an Ethereum contract.
type ReferralEscrow struct {
	ReferralEscrowCaller     // Read-only binding to the contract
	ReferralEscrowTransactor // Write-only binding to the contract
	ReferralEscrowFilterer   // Log filterer for contract events
}

// ReferralEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReferralEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReferralEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReferralEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReferralEscrowSession struct {
	Contract     *ReferralEscrow   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReferralEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReferralEscrowCallerSession struct {
	Contract *ReferralEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ReferralEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReferralEscrowTransactorSession struct {
	Contract     *ReferralEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ReferralEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReferralEscrowRaw struct {
	Contract *ReferralEscrow // Generic contract binding to access the raw methods on
}

// ReferralEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReferralEscrowCallerRaw struct {
	Contract *ReferralEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// ReferralEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReferralEscrowTransactorRaw struct {
	Contract *ReferralEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReferralEscrow creates a new instance of ReferralEscrow, bound to a specific deployed contract.
func NewReferralEscrow(address common.Address, backend bind.ContractBackend) (*ReferralEscrow, error) {
	contract, err := bindReferralEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReferralEscrow{ReferralEscrowCaller: ReferralEscrowCaller{contract: contract}, ReferralEscrowTransactor: ReferralEscrowTransactor{contract: contract}, ReferralEscrowFilterer: ReferralEscrowFilterer{contract: contract}}, nil
}

// NewReferralEscrowCaller creates a new read-only instance of ReferralEscrow, bound to a specific deployed contract.
func NewReferralEscrowCaller(address common.Address, caller bind.ContractCaller) (*ReferralEscrowCaller, error) {
	contract, err := bindReferralEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReferralEscrowCaller{contract: contract}, nil
}

// NewReferralEscrowTransactor creates a new write-only instance of ReferralEscrow, bound to a specific deployed contract.
func NewReferralEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*ReferralEscrowTransactor, error) {
	contract, err := bindReferralEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReferralEscrowTransactor{contract: contract}, nil
}

// NewReferralEscrowFilterer creates a new log filterer instance of ReferralEscrow, bound to a specific deployed contract.
func NewReferralEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*ReferralEscrowFilterer, error) {
	contract, err := bindReferralEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReferralEscrowFilterer{contract: contract}, nil
}

// bindReferralEscrow binds a generic wrapper to an already deployed contract.
func bindReferralEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReferralEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReferralEscrow *ReferralEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReferralEscrow.Contract.ReferralEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReferralEscrow *ReferralEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.ReferralEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReferralEscrow *ReferralEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.ReferralEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReferralEscrow *ReferralEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReferralEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReferralEscrow *ReferralEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReferralEscrow *ReferralEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.contract.Transact(opts, method, params...)
}

// Allocations is a free data retrieval call binding the contract method 0x853996e4.
//
// Solidity: function allocations(address _user, address _token) view returns(uint256 _amount)
func (_ReferralEscrow *ReferralEscrowCaller) Allocations(opts *bind.CallOpts, _user common.Address, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReferralEscrow.contract.Call(opts, &out, "allocations", _user, _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allocations is a free data retrieval call binding the contract method 0x853996e4.
//
// Solidity: function allocations(address _user, address _token) view returns(uint256 _amount)
func (_ReferralEscrow *ReferralEscrowSession) Allocations(_user common.Address, _token common.Address) (*big.Int, error) {
	return _ReferralEscrow.Contract.Allocations(&_ReferralEscrow.CallOpts, _user, _token)
}

// Allocations is a free data retrieval call binding the contract method 0x853996e4.
//
// Solidity: function allocations(address _user, address _token) view returns(uint256 _amount)
func (_ReferralEscrow *ReferralEscrowCallerSession) Allocations(_user common.Address, _token common.Address) (*big.Int, error) {
	return _ReferralEscrow.Contract.Allocations(&_ReferralEscrow.CallOpts, _user, _token)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_ReferralEscrow *ReferralEscrowCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReferralEscrow.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_ReferralEscrow *ReferralEscrowSession) NativeToken() (common.Address, error) {
	return _ReferralEscrow.Contract.NativeToken(&_ReferralEscrow.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_ReferralEscrow *ReferralEscrowCallerSession) NativeToken() (common.Address, error) {
	return _ReferralEscrow.Contract.NativeToken(&_ReferralEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_ReferralEscrow *ReferralEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReferralEscrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_ReferralEscrow *ReferralEscrowSession) Owner() (common.Address, error) {
	return _ReferralEscrow.Contract.Owner(&_ReferralEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_ReferralEscrow *ReferralEscrowCallerSession) Owner() (common.Address, error) {
	return _ReferralEscrow.Contract.Owner(&_ReferralEscrow.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_ReferralEscrow *ReferralEscrowCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReferralEscrow.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_ReferralEscrow *ReferralEscrowSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _ReferralEscrow.Contract.OwnershipHandoverExpiresAt(&_ReferralEscrow.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_ReferralEscrow *ReferralEscrowCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _ReferralEscrow.Contract.OwnershipHandoverExpiresAt(&_ReferralEscrow.CallOpts, pendingOwner)
}

// PoolSwap is a free data retrieval call binding the contract method 0x33809ad0.
//
// Solidity: function poolSwap() view returns(address)
func (_ReferralEscrow *ReferralEscrowCaller) PoolSwap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReferralEscrow.contract.Call(opts, &out, "poolSwap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolSwap is a free data retrieval call binding the contract method 0x33809ad0.
//
// Solidity: function poolSwap() view returns(address)
func (_ReferralEscrow *ReferralEscrowSession) PoolSwap() (common.Address, error) {
	return _ReferralEscrow.Contract.PoolSwap(&_ReferralEscrow.CallOpts)
}

// PoolSwap is a free data retrieval call binding the contract method 0x33809ad0.
//
// Solidity: function poolSwap() view returns(address)
func (_ReferralEscrow *ReferralEscrowCallerSession) PoolSwap() (common.Address, error) {
	return _ReferralEscrow.Contract.PoolSwap(&_ReferralEscrow.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_ReferralEscrow *ReferralEscrowCaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReferralEscrow.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_ReferralEscrow *ReferralEscrowSession) PositionManager() (common.Address, error) {
	return _ReferralEscrow.Contract.PositionManager(&_ReferralEscrow.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_ReferralEscrow *ReferralEscrowCallerSession) PositionManager() (common.Address, error) {
	return _ReferralEscrow.Contract.PositionManager(&_ReferralEscrow.CallOpts)
}

// AssignTokens is a paid mutator transaction binding the contract method 0xb0f4c2ea.
//
// Solidity: function assignTokens(bytes32 _poolId, address _user, address _token, uint256 _amount) returns()
func (_ReferralEscrow *ReferralEscrowTransactor) AssignTokens(opts *bind.TransactOpts, _poolId [32]byte, _user common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReferralEscrow.contract.Transact(opts, "assignTokens", _poolId, _user, _token, _amount)
}

// AssignTokens is a paid mutator transaction binding the contract method 0xb0f4c2ea.
//
// Solidity: function assignTokens(bytes32 _poolId, address _user, address _token, uint256 _amount) returns()
func (_ReferralEscrow *ReferralEscrowSession) AssignTokens(_poolId [32]byte, _user common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.AssignTokens(&_ReferralEscrow.TransactOpts, _poolId, _user, _token, _amount)
}

// AssignTokens is a paid mutator transaction binding the contract method 0xb0f4c2ea.
//
// Solidity: function assignTokens(bytes32 _poolId, address _user, address _token, uint256 _amount) returns()
func (_ReferralEscrow *ReferralEscrowTransactorSession) AssignTokens(_poolId [32]byte, _user common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.AssignTokens(&_ReferralEscrow.TransactOpts, _poolId, _user, _token, _amount)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_ReferralEscrow *ReferralEscrowTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferralEscrow.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_ReferralEscrow *ReferralEscrowSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _ReferralEscrow.Contract.CancelOwnershipHandover(&_ReferralEscrow.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_ReferralEscrow *ReferralEscrowTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _ReferralEscrow.Contract.CancelOwnershipHandover(&_ReferralEscrow.TransactOpts)
}

// ClaimAndSwap is a paid mutator transaction binding the contract method 0x78c5ae8c.
//
// Solidity: function claimAndSwap(address[] _tokens, uint160[] _sqrtPriceX96Limits, address _recipient) returns()
func (_ReferralEscrow *ReferralEscrowTransactor) ClaimAndSwap(opts *bind.TransactOpts, _tokens []common.Address, _sqrtPriceX96Limits []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.contract.Transact(opts, "claimAndSwap", _tokens, _sqrtPriceX96Limits, _recipient)
}

// ClaimAndSwap is a paid mutator transaction binding the contract method 0x78c5ae8c.
//
// Solidity: function claimAndSwap(address[] _tokens, uint160[] _sqrtPriceX96Limits, address _recipient) returns()
func (_ReferralEscrow *ReferralEscrowSession) ClaimAndSwap(_tokens []common.Address, _sqrtPriceX96Limits []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.ClaimAndSwap(&_ReferralEscrow.TransactOpts, _tokens, _sqrtPriceX96Limits, _recipient)
}

// ClaimAndSwap is a paid mutator transaction binding the contract method 0x78c5ae8c.
//
// Solidity: function claimAndSwap(address[] _tokens, uint160[] _sqrtPriceX96Limits, address _recipient) returns()
func (_ReferralEscrow *ReferralEscrowTransactorSession) ClaimAndSwap(_tokens []common.Address, _sqrtPriceX96Limits []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.ClaimAndSwap(&_ReferralEscrow.TransactOpts, _tokens, _sqrtPriceX96Limits, _recipient)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x988a18f6.
//
// Solidity: function claimTokens(address[] _tokens, address _recipient) returns()
func (_ReferralEscrow *ReferralEscrowTransactor) ClaimTokens(opts *bind.TransactOpts, _tokens []common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.contract.Transact(opts, "claimTokens", _tokens, _recipient)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x988a18f6.
//
// Solidity: function claimTokens(address[] _tokens, address _recipient) returns()
func (_ReferralEscrow *ReferralEscrowSession) ClaimTokens(_tokens []common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.ClaimTokens(&_ReferralEscrow.TransactOpts, _tokens, _recipient)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x988a18f6.
//
// Solidity: function claimTokens(address[] _tokens, address _recipient) returns()
func (_ReferralEscrow *ReferralEscrowTransactorSession) ClaimTokens(_tokens []common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.ClaimTokens(&_ReferralEscrow.TransactOpts, _tokens, _recipient)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_ReferralEscrow *ReferralEscrowTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_ReferralEscrow *ReferralEscrowSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.CompleteOwnershipHandover(&_ReferralEscrow.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_ReferralEscrow *ReferralEscrowTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.CompleteOwnershipHandover(&_ReferralEscrow.TransactOpts, pendingOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_ReferralEscrow *ReferralEscrowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferralEscrow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_ReferralEscrow *ReferralEscrowSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReferralEscrow.Contract.RenounceOwnership(&_ReferralEscrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_ReferralEscrow *ReferralEscrowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReferralEscrow.Contract.RenounceOwnership(&_ReferralEscrow.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_ReferralEscrow *ReferralEscrowTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferralEscrow.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_ReferralEscrow *ReferralEscrowSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _ReferralEscrow.Contract.RequestOwnershipHandover(&_ReferralEscrow.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_ReferralEscrow *ReferralEscrowTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _ReferralEscrow.Contract.RequestOwnershipHandover(&_ReferralEscrow.TransactOpts)
}

// SetPoolSwap is a paid mutator transaction binding the contract method 0x761032ec.
//
// Solidity: function setPoolSwap(address _poolSwap) returns()
func (_ReferralEscrow *ReferralEscrowTransactor) SetPoolSwap(opts *bind.TransactOpts, _poolSwap common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.contract.Transact(opts, "setPoolSwap", _poolSwap)
}

// SetPoolSwap is a paid mutator transaction binding the contract method 0x761032ec.
//
// Solidity: function setPoolSwap(address _poolSwap) returns()
func (_ReferralEscrow *ReferralEscrowSession) SetPoolSwap(_poolSwap common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.SetPoolSwap(&_ReferralEscrow.TransactOpts, _poolSwap)
}

// SetPoolSwap is a paid mutator transaction binding the contract method 0x761032ec.
//
// Solidity: function setPoolSwap(address _poolSwap) returns()
func (_ReferralEscrow *ReferralEscrowTransactorSession) SetPoolSwap(_poolSwap common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.SetPoolSwap(&_ReferralEscrow.TransactOpts, _poolSwap)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_ReferralEscrow *ReferralEscrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_ReferralEscrow *ReferralEscrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.TransferOwnership(&_ReferralEscrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_ReferralEscrow *ReferralEscrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReferralEscrow.Contract.TransferOwnership(&_ReferralEscrow.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ReferralEscrow *ReferralEscrowTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferralEscrow.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ReferralEscrow *ReferralEscrowSession) Receive() (*types.Transaction, error) {
	return _ReferralEscrow.Contract.Receive(&_ReferralEscrow.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ReferralEscrow *ReferralEscrowTransactorSession) Receive() (*types.Transaction, error) {
	return _ReferralEscrow.Contract.Receive(&_ReferralEscrow.TransactOpts)
}

// ReferralEscrowOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the ReferralEscrow contract.
type ReferralEscrowOwnershipHandoverCanceledIterator struct {
	Event *ReferralEscrowOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *ReferralEscrowOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralEscrowOwnershipHandoverCanceled)
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
		it.Event = new(ReferralEscrowOwnershipHandoverCanceled)
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
func (it *ReferralEscrowOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralEscrowOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralEscrowOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the ReferralEscrow contract.
type ReferralEscrowOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_ReferralEscrow *ReferralEscrowFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*ReferralEscrowOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ReferralEscrow.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReferralEscrowOwnershipHandoverCanceledIterator{contract: _ReferralEscrow.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_ReferralEscrow *ReferralEscrowFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *ReferralEscrowOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ReferralEscrow.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralEscrowOwnershipHandoverCanceled)
				if err := _ReferralEscrow.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_ReferralEscrow *ReferralEscrowFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*ReferralEscrowOwnershipHandoverCanceled, error) {
	event := new(ReferralEscrowOwnershipHandoverCanceled)
	if err := _ReferralEscrow.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralEscrowOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the ReferralEscrow contract.
type ReferralEscrowOwnershipHandoverRequestedIterator struct {
	Event *ReferralEscrowOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *ReferralEscrowOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralEscrowOwnershipHandoverRequested)
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
		it.Event = new(ReferralEscrowOwnershipHandoverRequested)
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
func (it *ReferralEscrowOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralEscrowOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralEscrowOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the ReferralEscrow contract.
type ReferralEscrowOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_ReferralEscrow *ReferralEscrowFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*ReferralEscrowOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ReferralEscrow.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReferralEscrowOwnershipHandoverRequestedIterator{contract: _ReferralEscrow.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_ReferralEscrow *ReferralEscrowFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *ReferralEscrowOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ReferralEscrow.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralEscrowOwnershipHandoverRequested)
				if err := _ReferralEscrow.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_ReferralEscrow *ReferralEscrowFilterer) ParseOwnershipHandoverRequested(log types.Log) (*ReferralEscrowOwnershipHandoverRequested, error) {
	event := new(ReferralEscrowOwnershipHandoverRequested)
	if err := _ReferralEscrow.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralEscrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ReferralEscrow contract.
type ReferralEscrowOwnershipTransferredIterator struct {
	Event *ReferralEscrowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReferralEscrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralEscrowOwnershipTransferred)
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
		it.Event = new(ReferralEscrowOwnershipTransferred)
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
func (it *ReferralEscrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralEscrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralEscrowOwnershipTransferred represents a OwnershipTransferred event raised by the ReferralEscrow contract.
type ReferralEscrowOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_ReferralEscrow *ReferralEscrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*ReferralEscrowOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReferralEscrow.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReferralEscrowOwnershipTransferredIterator{contract: _ReferralEscrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_ReferralEscrow *ReferralEscrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReferralEscrowOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReferralEscrow.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralEscrowOwnershipTransferred)
				if err := _ReferralEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ReferralEscrow *ReferralEscrowFilterer) ParseOwnershipTransferred(log types.Log) (*ReferralEscrowOwnershipTransferred, error) {
	event := new(ReferralEscrowOwnershipTransferred)
	if err := _ReferralEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralEscrowTokensAssignedIterator is returned from FilterTokensAssigned and is used to iterate over the raw logs and unpacked data for TokensAssigned events raised by the ReferralEscrow contract.
type ReferralEscrowTokensAssignedIterator struct {
	Event *ReferralEscrowTokensAssigned // Event containing the contract specifics and raw log

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
func (it *ReferralEscrowTokensAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralEscrowTokensAssigned)
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
		it.Event = new(ReferralEscrowTokensAssigned)
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
func (it *ReferralEscrowTokensAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralEscrowTokensAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralEscrowTokensAssigned represents a TokensAssigned event raised by the ReferralEscrow contract.
type ReferralEscrowTokensAssigned struct {
	PoolId [32]byte
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensAssigned is a free log retrieval operation binding the contract event 0xc733341be84ca8f5e6ee2be63797abe1af7bbe5efcb5c8abe995dc8a51fd4b37.
//
// Solidity: event TokensAssigned(bytes32 indexed _poolId, address indexed _user, address indexed _token, uint256 _amount)
func (_ReferralEscrow *ReferralEscrowFilterer) FilterTokensAssigned(opts *bind.FilterOpts, _poolId [][32]byte, _user []common.Address, _token []common.Address) (*ReferralEscrowTokensAssignedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _ReferralEscrow.contract.FilterLogs(opts, "TokensAssigned", _poolIdRule, _userRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &ReferralEscrowTokensAssignedIterator{contract: _ReferralEscrow.contract, event: "TokensAssigned", logs: logs, sub: sub}, nil
}

// WatchTokensAssigned is a free log subscription operation binding the contract event 0xc733341be84ca8f5e6ee2be63797abe1af7bbe5efcb5c8abe995dc8a51fd4b37.
//
// Solidity: event TokensAssigned(bytes32 indexed _poolId, address indexed _user, address indexed _token, uint256 _amount)
func (_ReferralEscrow *ReferralEscrowFilterer) WatchTokensAssigned(opts *bind.WatchOpts, sink chan<- *ReferralEscrowTokensAssigned, _poolId [][32]byte, _user []common.Address, _token []common.Address) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _ReferralEscrow.contract.WatchLogs(opts, "TokensAssigned", _poolIdRule, _userRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralEscrowTokensAssigned)
				if err := _ReferralEscrow.contract.UnpackLog(event, "TokensAssigned", log); err != nil {
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

// ParseTokensAssigned is a log parse operation binding the contract event 0xc733341be84ca8f5e6ee2be63797abe1af7bbe5efcb5c8abe995dc8a51fd4b37.
//
// Solidity: event TokensAssigned(bytes32 indexed _poolId, address indexed _user, address indexed _token, uint256 _amount)
func (_ReferralEscrow *ReferralEscrowFilterer) ParseTokensAssigned(log types.Log) (*ReferralEscrowTokensAssigned, error) {
	event := new(ReferralEscrowTokensAssigned)
	if err := _ReferralEscrow.contract.UnpackLog(event, "TokensAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralEscrowTokensClaimedIterator is returned from FilterTokensClaimed and is used to iterate over the raw logs and unpacked data for TokensClaimed events raised by the ReferralEscrow contract.
type ReferralEscrowTokensClaimedIterator struct {
	Event *ReferralEscrowTokensClaimed // Event containing the contract specifics and raw log

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
func (it *ReferralEscrowTokensClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralEscrowTokensClaimed)
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
		it.Event = new(ReferralEscrowTokensClaimed)
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
func (it *ReferralEscrowTokensClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralEscrowTokensClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralEscrowTokensClaimed represents a TokensClaimed event raised by the ReferralEscrow contract.
type ReferralEscrowTokensClaimed struct {
	User      common.Address
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensClaimed is a free log retrieval operation binding the contract event 0x40b4b442b70a20081fa377001b642c7d3618847a45b5b8ed2b938e30a3299ddb.
//
// Solidity: event TokensClaimed(address indexed _user, address _recipient, address indexed _token, uint256 _amount)
func (_ReferralEscrow *ReferralEscrowFilterer) FilterTokensClaimed(opts *bind.FilterOpts, _user []common.Address, _token []common.Address) (*ReferralEscrowTokensClaimedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _ReferralEscrow.contract.FilterLogs(opts, "TokensClaimed", _userRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &ReferralEscrowTokensClaimedIterator{contract: _ReferralEscrow.contract, event: "TokensClaimed", logs: logs, sub: sub}, nil
}

// WatchTokensClaimed is a free log subscription operation binding the contract event 0x40b4b442b70a20081fa377001b642c7d3618847a45b5b8ed2b938e30a3299ddb.
//
// Solidity: event TokensClaimed(address indexed _user, address _recipient, address indexed _token, uint256 _amount)
func (_ReferralEscrow *ReferralEscrowFilterer) WatchTokensClaimed(opts *bind.WatchOpts, sink chan<- *ReferralEscrowTokensClaimed, _user []common.Address, _token []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _ReferralEscrow.contract.WatchLogs(opts, "TokensClaimed", _userRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralEscrowTokensClaimed)
				if err := _ReferralEscrow.contract.UnpackLog(event, "TokensClaimed", log); err != nil {
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

// ParseTokensClaimed is a log parse operation binding the contract event 0x40b4b442b70a20081fa377001b642c7d3618847a45b5b8ed2b938e30a3299ddb.
//
// Solidity: event TokensClaimed(address indexed _user, address _recipient, address indexed _token, uint256 _amount)
func (_ReferralEscrow *ReferralEscrowFilterer) ParseTokensClaimed(log types.Log) (*ReferralEscrowTokensClaimed, error) {
	event := new(ReferralEscrowTokensClaimed)
	if err := _ReferralEscrow.contract.UnpackLog(event, "TokensClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralEscrowTokensSwappedIterator is returned from FilterTokensSwapped and is used to iterate over the raw logs and unpacked data for TokensSwapped events raised by the ReferralEscrow contract.
type ReferralEscrowTokensSwappedIterator struct {
	Event *ReferralEscrowTokensSwapped // Event containing the contract specifics and raw log

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
func (it *ReferralEscrowTokensSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralEscrowTokensSwapped)
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
		it.Event = new(ReferralEscrowTokensSwapped)
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
func (it *ReferralEscrowTokensSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralEscrowTokensSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralEscrowTokensSwapped represents a TokensSwapped event raised by the ReferralEscrow contract.
type ReferralEscrowTokensSwapped struct {
	User    common.Address
	Token   common.Address
	TokenIn *big.Int
	EthOut  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokensSwapped is a free log retrieval operation binding the contract event 0x25f1d03755df23c30e25db2dbd3891e31ce084bdfbfc46f9fe5e446ee5f9b2d4.
//
// Solidity: event TokensSwapped(address indexed _user, address indexed _token, uint256 _tokenIn, uint256 _ethOut)
func (_ReferralEscrow *ReferralEscrowFilterer) FilterTokensSwapped(opts *bind.FilterOpts, _user []common.Address, _token []common.Address) (*ReferralEscrowTokensSwappedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _ReferralEscrow.contract.FilterLogs(opts, "TokensSwapped", _userRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &ReferralEscrowTokensSwappedIterator{contract: _ReferralEscrow.contract, event: "TokensSwapped", logs: logs, sub: sub}, nil
}

// WatchTokensSwapped is a free log subscription operation binding the contract event 0x25f1d03755df23c30e25db2dbd3891e31ce084bdfbfc46f9fe5e446ee5f9b2d4.
//
// Solidity: event TokensSwapped(address indexed _user, address indexed _token, uint256 _tokenIn, uint256 _ethOut)
func (_ReferralEscrow *ReferralEscrowFilterer) WatchTokensSwapped(opts *bind.WatchOpts, sink chan<- *ReferralEscrowTokensSwapped, _user []common.Address, _token []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _ReferralEscrow.contract.WatchLogs(opts, "TokensSwapped", _userRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralEscrowTokensSwapped)
				if err := _ReferralEscrow.contract.UnpackLog(event, "TokensSwapped", log); err != nil {
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

// ParseTokensSwapped is a log parse operation binding the contract event 0x25f1d03755df23c30e25db2dbd3891e31ce084bdfbfc46f9fe5e446ee5f9b2d4.
//
// Solidity: event TokensSwapped(address indexed _user, address indexed _token, uint256 _tokenIn, uint256 _ethOut)
func (_ReferralEscrow *ReferralEscrowFilterer) ParseTokensSwapped(log types.Log) (*ReferralEscrowTokensSwapped, error) {
	event := new(ReferralEscrowTokensSwapped)
	if err := _ReferralEscrow.contract.UnpackLog(event, "TokensSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
