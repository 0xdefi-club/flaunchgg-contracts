// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market_capped_price_v3

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

// MarketCappedPriceV3MetaData contains all meta data concerning the MarketCappedPriceV3 contract.
var MarketCappedPriceV3MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_protocolOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ethToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdcToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MINIMUM_USDC_MARKET_CAP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ethToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFlaunchingFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketCap\",\"inputs\":[{\"name\":\"_initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSqrtPriceX96\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_flipped\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"sqrtPriceX96_\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUniswapV3Pool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setPool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"usdcToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdcToken0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTokenPair\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketCapTooSmall\",\"inputs\":[{\"name\":\"_usdcMarketCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_usdcMarketCapMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161133938038061133983398101604081905261002e916100c5565b6001600160a01b03808316608052811660a05261004a83610052565b505050610105565b638b78c6d81980541561006c57630dc149f05f526004601cfd5b6001600160a01b03909116801560ff1b8117909155805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b80516001600160a01b03811681146100c0575f5ffd5b919050565b5f5f5f606084860312156100d7575f5ffd5b6100e0846100aa565b92506100ee602085016100aa565b91506100fc604085016100aa565b90509250925092565b60805160a0516111f06101495f395f818161010b0152818161064901528181610686015261075501525f8181610210015281816106c3015261070001526111f05ff3fe6080604052600436106100c2575f3560e01c80630c7324d3146100c657806311eac855146100fa57806316f0115b1461014557806325692962146101635780632a81c4f81461016d578063355c73711461019a57806342d8c2c0146101b15780634437152a146101d057806354d1f13d146101ef578063715018a6146101f75780637bf1a627146101ff5780638da5cb5b14610232578063f04e283e1461024a578063f2fde38b1461025d578063fc87f3f614610270578063fee81cf41461028f575b5f5ffd5b3480156100d1575f5ffd5b505f546100e590600160a01b900460ff1681565b60405190151581526020015b60405180910390f35b348015610105575f5ffd5b5061012d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100f1565b348015610150575f5ffd5b505f5461012d906001600160a01b031681565b61016b6102c0565b005b348015610178575f5ffd5b5061018c610187366004610d87565b61030c565b6040519081526020016100f1565b3480156101a5575f5ffd5b5061018c633b9aca0081565b3480156101bc575f5ffd5b5061018c6101cb366004610dd9565b61052c565b3480156101db575f5ffd5b5061016b6101ea366004610e29565b61054e565b61016b6107a5565b61016b6107de565b34801561020a575f5ffd5b5061012d7f000000000000000000000000000000000000000000000000000000000000000081565b34801561023d575f5ffd5b50638b78c6d8195461012d565b61016b610258366004610e29565b6107f1565b61016b61026b366004610e29565b61082e565b34801561027b575f5ffd5b5061012d61028a366004610e44565b610854565b34801561029a575f5ffd5b5061018c6102a9366004610e29565b63389a75e1600c9081525f91909152602090205490565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b5f8061031a83850185610eec565b9050633b9aca00815f01511015610359578051604051630292f1cf60e01b81526004810191909152633b9aca0060248201526044015b60405180910390fd5b6040805160028082526060820183525f92602083019080368337019050509050610708815f8151811061038e5761038e610f2d565b602002602001019063ffffffff16908163ffffffff16815250505f816001815181106103bc576103bc610f2d565b63ffffffff909216602092830291909101909101525f805460405163883bdbfd60e01b81526001600160a01b039091169063883bdbfd90610401908590600401610f41565b5f60405180830381865afa15801561041b573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610442919081019061101e565b5090505f815f8151811061045857610458610f2d565b60200260200101518260018151811061047357610473610f2d565b602002602001015161048591906110fc565b90505f6104946107088361113d565b90505f6104a08261087f565b90505f6104bb6001600160a01b03831680600160601b610b37565b5f805491925090600160a01b900460ff166104ec576104e7670de0b6b3a764000083600160601b610b37565b610503565b610503670de0b6b3a7640000600160601b84610b37565b905061051b885f0151670de0b6b3a764000083610b37565b985050505050505050505b92915050565b5f6103e861053a848461030c565b6105449190611179565b90505b9392505050565b610556610bd3565b5f80546001600160a01b0319166001600160a01b038316908117825560408051630dfe168160e01b81529051630dfe1681916004808201926020929091908290030181865afa1580156105ab573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105cf919061118c565b90505f5f5f9054906101000a90046001600160a01b03166001600160a01b031663d21220a76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610621573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610645919061118c565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316141580156106bb57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031614155b8061073557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161415801561073557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031614155b1561075357604051638686656d60e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316145f60146101000a81548160ff021916908315150217905550505050565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b6107e6610bd3565b6107ef5f610bed565b565b6107f9610bd3565b63389a75e1600c52805f526020600c20805442111561081f57636f5e88185f526004601cfd5b5f905561082b81610bed565b50565b610836610bd3565b8060601b61084b57637448fbae5f526004601cfd5b61082b81610bed565b5f610876610862848461030c565b680a18f07d736b90be55601d1b8615610c33565b95945050505050565b60020b5f60ff82901d80830118620d89e88111156108a8576108a86345c3193d60e11b84610cce565b7001fffcb933bd6fad37aa2d162d1a5940016001821602600160801b1860028216156108e4576ffff97272373d413259a46990580e213a0260801c5b6004821615610903576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615610922576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b6010821615610941576fffcb9843d60f6159c9db58835c9266440260801c5b6020821615610960576fff973b41fa98c081472e6896dfb254c00260801c5b604082161561097f576fff2ea16466c96a3843ec78b326b528610260801c5b608082161561099e576ffe5dee046a99a2a811c461f1969c30530260801c5b6101008216156109be576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b6102008216156109de576ff987a7253ac413176f2b074cf7815e540260801c5b6104008216156109fe576ff3392b0822b70005940c7a398e4b70f30260801c5b610800821615610a1e576fe7159475a2c29b7443b29c7fa6e889d90260801c5b611000821615610a3e576fd097f3bdfd2022b8845ad8f792aa58250260801c5b612000821615610a5e576fa9f746462d870fdf8a65dc1f90e061e50260801c5b614000821615610a7e576f70d869a156d2a1b890bb3df62baf32f70260801c5b618000821615610a9e576f31be135f97d08fd981231505542fcfa60260801c5b62010000821615610abf576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b62020000821615610adf576e5d6af8dedb81196699c329225ee6040260801c5b62040000821615610afe576d2216e584f5fa1ea926041bedfe980260801c5b62080000821615610b1b576b048a170391f7dc42444e8fa20260801c5b5f841315610b27575f19045b63ffffffff0160201c9392505050565b5f838302815f1985870982811083820303915050808411610b56575f5ffd5b805f03610b6857508290049050610547565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b638b78c6d8195433146107ef576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3811560ff1b8217905550565b5f5f84118015610c4257505f83115b610c985760405162461bcd60e51b815260206004820152602160248201527f416d6f756e7473206d7573742062652067726561746572207468616e207a65726044820152606f60f81b6064820152608401610350565b8115610cbc57610cb5610cb084600160c01b87610b37565b610cdd565b9050610547565b610544610cb085600160c01b86610b37565b815f528060020b60045260245ffd5b5f815f03610cec57505f919050565b5f6002610cfa8460016111a7565b610d049190611179565b90508291505b81811015610d3d57905080600281610d228186611179565b610d2c91906111a7565b610d369190611179565b9050610d0a565b50919050565b5f5f83601f840112610d53575f5ffd5b5081356001600160401b03811115610d69575f5ffd5b602083019150836020828501011115610d80575f5ffd5b9250929050565b5f5f60208385031215610d98575f5ffd5b82356001600160401b03811115610dad575f5ffd5b610db985828601610d43565b90969095509350505050565b6001600160a01b038116811461082b575f5ffd5b5f5f5f60408486031215610deb575f5ffd5b8335610df681610dc5565b925060208401356001600160401b03811115610e10575f5ffd5b610e1c86828701610d43565b9497909650939450505050565b5f60208284031215610e39575f5ffd5b813561054781610dc5565b5f5f5f5f60608587031215610e57575f5ffd5b8435610e6281610dc5565b935060208501358015158114610e76575f5ffd5b925060408501356001600160401b03811115610e90575f5ffd5b610e9c87828801610d43565b95989497509550505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715610ee457610ee4610ea8565b604052919050565b5f6020828403128015610efd575f5ffd5b50604051602081016001600160401b0381118282101715610f2057610f20610ea8565b6040529135825250919050565b634e487b7160e01b5f52603260045260245ffd5b602080825282518282018190525f918401906040840190835b81811015610f7e57835163ffffffff16835260209384019390920191600101610f5a565b509095945050505050565b5f6001600160401b03821115610fa157610fa1610ea8565b5060051b60200190565b5f82601f830112610fba575f5ffd5b8151610fcd610fc882610f89565b610ebc565b8082825260208201915060208360051b860101925085831115610fee575f5ffd5b602085015b8381101561101457805161100681610dc5565b835260209283019201610ff3565b5095945050505050565b5f5f6040838503121561102f575f5ffd5b82516001600160401b03811115611044575f5ffd5b8301601f81018513611054575f5ffd5b8051611062610fc882610f89565b8082825260208201915060208360051b850101925087831115611083575f5ffd5b6020840193505b828410156110b35783518060060b81146110a2575f5ffd5b82526020938401939091019061108a565b6020870151909550925050506001600160401b038111156110d2575f5ffd5b6110de85828601610fab565b9150509250929050565b634e487b7160e01b5f52601160045260245ffd5b600682810b9082900b03667fffffffffffff198112667fffffffffffff82131715610526576105266110e8565b634e487b7160e01b5f52601260045260245ffd5b5f8160060b8360060b8061115357611153611129565b667fffffffffffff1982145f1982141615611170576111706110e8565b90059392505050565b5f8261118757611187611129565b500490565b5f6020828403121561119c575f5ffd5b815161054781610dc5565b80820180821115610526576105266110e856fea26469706673582212208d39edf084f2c911e82e6e08f567008f3cca56f09af9931555360129d8328e3964736f6c634300081b0033",
}

// MarketCappedPriceV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketCappedPriceV3MetaData.ABI instead.
var MarketCappedPriceV3ABI = MarketCappedPriceV3MetaData.ABI

// MarketCappedPriceV3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketCappedPriceV3MetaData.Bin instead.
var MarketCappedPriceV3Bin = MarketCappedPriceV3MetaData.Bin

// DeployMarketCappedPriceV3 deploys a new Ethereum contract, binding an instance of MarketCappedPriceV3 to it.
func DeployMarketCappedPriceV3(auth *bind.TransactOpts, backend bind.ContractBackend, _protocolOwner common.Address, _ethToken common.Address, _usdcToken common.Address) (common.Address, *types.Transaction, *MarketCappedPriceV3, error) {
	parsed, err := MarketCappedPriceV3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketCappedPriceV3Bin), backend, _protocolOwner, _ethToken, _usdcToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketCappedPriceV3{MarketCappedPriceV3Caller: MarketCappedPriceV3Caller{contract: contract}, MarketCappedPriceV3Transactor: MarketCappedPriceV3Transactor{contract: contract}, MarketCappedPriceV3Filterer: MarketCappedPriceV3Filterer{contract: contract}}, nil
}

// MarketCappedPriceV3 is an auto generated Go binding around an Ethereum contract.
type MarketCappedPriceV3 struct {
	MarketCappedPriceV3Caller     // Read-only binding to the contract
	MarketCappedPriceV3Transactor // Write-only binding to the contract
	MarketCappedPriceV3Filterer   // Log filterer for contract events
}

// MarketCappedPriceV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCappedPriceV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketCappedPriceV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketCappedPriceV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketCappedPriceV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketCappedPriceV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketCappedPriceV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketCappedPriceV3Session struct {
	Contract     *MarketCappedPriceV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MarketCappedPriceV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCappedPriceV3CallerSession struct {
	Contract *MarketCappedPriceV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MarketCappedPriceV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketCappedPriceV3TransactorSession struct {
	Contract     *MarketCappedPriceV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MarketCappedPriceV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type MarketCappedPriceV3Raw struct {
	Contract *MarketCappedPriceV3 // Generic contract binding to access the raw methods on
}

// MarketCappedPriceV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCappedPriceV3CallerRaw struct {
	Contract *MarketCappedPriceV3Caller // Generic read-only contract binding to access the raw methods on
}

// MarketCappedPriceV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketCappedPriceV3TransactorRaw struct {
	Contract *MarketCappedPriceV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketCappedPriceV3 creates a new instance of MarketCappedPriceV3, bound to a specific deployed contract.
func NewMarketCappedPriceV3(address common.Address, backend bind.ContractBackend) (*MarketCappedPriceV3, error) {
	contract, err := bindMarketCappedPriceV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceV3{MarketCappedPriceV3Caller: MarketCappedPriceV3Caller{contract: contract}, MarketCappedPriceV3Transactor: MarketCappedPriceV3Transactor{contract: contract}, MarketCappedPriceV3Filterer: MarketCappedPriceV3Filterer{contract: contract}}, nil
}

// NewMarketCappedPriceV3Caller creates a new read-only instance of MarketCappedPriceV3, bound to a specific deployed contract.
func NewMarketCappedPriceV3Caller(address common.Address, caller bind.ContractCaller) (*MarketCappedPriceV3Caller, error) {
	contract, err := bindMarketCappedPriceV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceV3Caller{contract: contract}, nil
}

// NewMarketCappedPriceV3Transactor creates a new write-only instance of MarketCappedPriceV3, bound to a specific deployed contract.
func NewMarketCappedPriceV3Transactor(address common.Address, transactor bind.ContractTransactor) (*MarketCappedPriceV3Transactor, error) {
	contract, err := bindMarketCappedPriceV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceV3Transactor{contract: contract}, nil
}

// NewMarketCappedPriceV3Filterer creates a new log filterer instance of MarketCappedPriceV3, bound to a specific deployed contract.
func NewMarketCappedPriceV3Filterer(address common.Address, filterer bind.ContractFilterer) (*MarketCappedPriceV3Filterer, error) {
	contract, err := bindMarketCappedPriceV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceV3Filterer{contract: contract}, nil
}

// bindMarketCappedPriceV3 binds a generic wrapper to an already deployed contract.
func bindMarketCappedPriceV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketCappedPriceV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketCappedPriceV3 *MarketCappedPriceV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketCappedPriceV3.Contract.MarketCappedPriceV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketCappedPriceV3 *MarketCappedPriceV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.MarketCappedPriceV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketCappedPriceV3 *MarketCappedPriceV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.MarketCappedPriceV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketCappedPriceV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketCappedPriceV3 *MarketCappedPriceV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketCappedPriceV3 *MarketCappedPriceV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMUSDCMARKETCAP is a free data retrieval call binding the contract method 0x355c7371.
//
// Solidity: function MINIMUM_USDC_MARKET_CAP() view returns(uint256)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Caller) MINIMUMUSDCMARKETCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketCappedPriceV3.contract.Call(opts, &out, "MINIMUM_USDC_MARKET_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMUSDCMARKETCAP is a free data retrieval call binding the contract method 0x355c7371.
//
// Solidity: function MINIMUM_USDC_MARKET_CAP() view returns(uint256)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) MINIMUMUSDCMARKETCAP() (*big.Int, error) {
	return _MarketCappedPriceV3.Contract.MINIMUMUSDCMARKETCAP(&_MarketCappedPriceV3.CallOpts)
}

// MINIMUMUSDCMARKETCAP is a free data retrieval call binding the contract method 0x355c7371.
//
// Solidity: function MINIMUM_USDC_MARKET_CAP() view returns(uint256)
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerSession) MINIMUMUSDCMARKETCAP() (*big.Int, error) {
	return _MarketCappedPriceV3.Contract.MINIMUMUSDCMARKETCAP(&_MarketCappedPriceV3.CallOpts)
}

// EthToken is a free data retrieval call binding the contract method 0x7bf1a627.
//
// Solidity: function ethToken() view returns(address)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Caller) EthToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketCappedPriceV3.contract.Call(opts, &out, "ethToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthToken is a free data retrieval call binding the contract method 0x7bf1a627.
//
// Solidity: function ethToken() view returns(address)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) EthToken() (common.Address, error) {
	return _MarketCappedPriceV3.Contract.EthToken(&_MarketCappedPriceV3.CallOpts)
}

// EthToken is a free data retrieval call binding the contract method 0x7bf1a627.
//
// Solidity: function ethToken() view returns(address)
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerSession) EthToken() (common.Address, error) {
	return _MarketCappedPriceV3.Contract.EthToken(&_MarketCappedPriceV3.CallOpts)
}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0x42d8c2c0.
//
// Solidity: function getFlaunchingFee(address , bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Caller) GetFlaunchingFee(opts *bind.CallOpts, arg0 common.Address, _initialPriceParams []byte) (*big.Int, error) {
	var out []interface{}
	err := _MarketCappedPriceV3.contract.Call(opts, &out, "getFlaunchingFee", arg0, _initialPriceParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0x42d8c2c0.
//
// Solidity: function getFlaunchingFee(address , bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) GetFlaunchingFee(arg0 common.Address, _initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPriceV3.Contract.GetFlaunchingFee(&_MarketCappedPriceV3.CallOpts, arg0, _initialPriceParams)
}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0x42d8c2c0.
//
// Solidity: function getFlaunchingFee(address , bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerSession) GetFlaunchingFee(arg0 common.Address, _initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPriceV3.Contract.GetFlaunchingFee(&_MarketCappedPriceV3.CallOpts, arg0, _initialPriceParams)
}

// GetMarketCap is a free data retrieval call binding the contract method 0x2a81c4f8.
//
// Solidity: function getMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Caller) GetMarketCap(opts *bind.CallOpts, _initialPriceParams []byte) (*big.Int, error) {
	var out []interface{}
	err := _MarketCappedPriceV3.contract.Call(opts, &out, "getMarketCap", _initialPriceParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCap is a free data retrieval call binding the contract method 0x2a81c4f8.
//
// Solidity: function getMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) GetMarketCap(_initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPriceV3.Contract.GetMarketCap(&_MarketCappedPriceV3.CallOpts, _initialPriceParams)
}

// GetMarketCap is a free data retrieval call binding the contract method 0x2a81c4f8.
//
// Solidity: function getMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerSession) GetMarketCap(_initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPriceV3.Contract.GetMarketCap(&_MarketCappedPriceV3.CallOpts, _initialPriceParams)
}

// GetSqrtPriceX96 is a free data retrieval call binding the contract method 0xfc87f3f6.
//
// Solidity: function getSqrtPriceX96(address , bool _flipped, bytes _initialPriceParams) view returns(uint160 sqrtPriceX96_)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Caller) GetSqrtPriceX96(opts *bind.CallOpts, arg0 common.Address, _flipped bool, _initialPriceParams []byte) (*big.Int, error) {
	var out []interface{}
	err := _MarketCappedPriceV3.contract.Call(opts, &out, "getSqrtPriceX96", arg0, _flipped, _initialPriceParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSqrtPriceX96 is a free data retrieval call binding the contract method 0xfc87f3f6.
//
// Solidity: function getSqrtPriceX96(address , bool _flipped, bytes _initialPriceParams) view returns(uint160 sqrtPriceX96_)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) GetSqrtPriceX96(arg0 common.Address, _flipped bool, _initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPriceV3.Contract.GetSqrtPriceX96(&_MarketCappedPriceV3.CallOpts, arg0, _flipped, _initialPriceParams)
}

// GetSqrtPriceX96 is a free data retrieval call binding the contract method 0xfc87f3f6.
//
// Solidity: function getSqrtPriceX96(address , bool _flipped, bytes _initialPriceParams) view returns(uint160 sqrtPriceX96_)
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerSession) GetSqrtPriceX96(arg0 common.Address, _flipped bool, _initialPriceParams []byte) (*big.Int, error) {
	return _MarketCappedPriceV3.Contract.GetSqrtPriceX96(&_MarketCappedPriceV3.CallOpts, arg0, _flipped, _initialPriceParams)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketCappedPriceV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) Owner() (common.Address, error) {
	return _MarketCappedPriceV3.Contract.Owner(&_MarketCappedPriceV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerSession) Owner() (common.Address, error) {
	return _MarketCappedPriceV3.Contract.Owner(&_MarketCappedPriceV3.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Caller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MarketCappedPriceV3.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _MarketCappedPriceV3.Contract.OwnershipHandoverExpiresAt(&_MarketCappedPriceV3.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _MarketCappedPriceV3.Contract.OwnershipHandoverExpiresAt(&_MarketCappedPriceV3.CallOpts, pendingOwner)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Caller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketCappedPriceV3.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) Pool() (common.Address, error) {
	return _MarketCappedPriceV3.Contract.Pool(&_MarketCappedPriceV3.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerSession) Pool() (common.Address, error) {
	return _MarketCappedPriceV3.Contract.Pool(&_MarketCappedPriceV3.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Caller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketCappedPriceV3.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) UsdcToken() (common.Address, error) {
	return _MarketCappedPriceV3.Contract.UsdcToken(&_MarketCappedPriceV3.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerSession) UsdcToken() (common.Address, error) {
	return _MarketCappedPriceV3.Contract.UsdcToken(&_MarketCappedPriceV3.CallOpts)
}

// UsdcToken0 is a free data retrieval call binding the contract method 0x0c7324d3.
//
// Solidity: function usdcToken0() view returns(bool)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Caller) UsdcToken0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MarketCappedPriceV3.contract.Call(opts, &out, "usdcToken0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsdcToken0 is a free data retrieval call binding the contract method 0x0c7324d3.
//
// Solidity: function usdcToken0() view returns(bool)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) UsdcToken0() (bool, error) {
	return _MarketCappedPriceV3.Contract.UsdcToken0(&_MarketCappedPriceV3.CallOpts)
}

// UsdcToken0 is a free data retrieval call binding the contract method 0x0c7324d3.
//
// Solidity: function usdcToken0() view returns(bool)
func (_MarketCappedPriceV3 *MarketCappedPriceV3CallerSession) UsdcToken0() (bool, error) {
	return _MarketCappedPriceV3.Contract.UsdcToken0(&_MarketCappedPriceV3.CallOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Transactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCappedPriceV3.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) CancelOwnershipHandover() (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.CancelOwnershipHandover(&_MarketCappedPriceV3.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3TransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.CancelOwnershipHandover(&_MarketCappedPriceV3.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Transactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPriceV3.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.CompleteOwnershipHandover(&_MarketCappedPriceV3.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3TransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.CompleteOwnershipHandover(&_MarketCappedPriceV3.TransactOpts, pendingOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCappedPriceV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.RenounceOwnership(&_MarketCappedPriceV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.RenounceOwnership(&_MarketCappedPriceV3.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Transactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCappedPriceV3.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) RequestOwnershipHandover() (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.RequestOwnershipHandover(&_MarketCappedPriceV3.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3TransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.RequestOwnershipHandover(&_MarketCappedPriceV3.TransactOpts)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _pool) returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Transactor) SetPool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _MarketCappedPriceV3.contract.Transact(opts, "setPool", _pool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _pool) returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) SetPool(_pool common.Address) (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.SetPool(&_MarketCappedPriceV3.TransactOpts, _pool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _pool) returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3TransactorSession) SetPool(_pool common.Address) (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.SetPool(&_MarketCappedPriceV3.TransactOpts, _pool)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPriceV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.TransferOwnership(&_MarketCappedPriceV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_MarketCappedPriceV3 *MarketCappedPriceV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketCappedPriceV3.Contract.TransferOwnership(&_MarketCappedPriceV3.TransactOpts, newOwner)
}

// MarketCappedPriceV3OwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the MarketCappedPriceV3 contract.
type MarketCappedPriceV3OwnershipHandoverCanceledIterator struct {
	Event *MarketCappedPriceV3OwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *MarketCappedPriceV3OwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCappedPriceV3OwnershipHandoverCanceled)
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
		it.Event = new(MarketCappedPriceV3OwnershipHandoverCanceled)
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
func (it *MarketCappedPriceV3OwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCappedPriceV3OwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCappedPriceV3OwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the MarketCappedPriceV3 contract.
type MarketCappedPriceV3OwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Filterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*MarketCappedPriceV3OwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _MarketCappedPriceV3.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceV3OwnershipHandoverCanceledIterator{contract: _MarketCappedPriceV3.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Filterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *MarketCappedPriceV3OwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _MarketCappedPriceV3.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCappedPriceV3OwnershipHandoverCanceled)
				if err := _MarketCappedPriceV3.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_MarketCappedPriceV3 *MarketCappedPriceV3Filterer) ParseOwnershipHandoverCanceled(log types.Log) (*MarketCappedPriceV3OwnershipHandoverCanceled, error) {
	event := new(MarketCappedPriceV3OwnershipHandoverCanceled)
	if err := _MarketCappedPriceV3.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketCappedPriceV3OwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the MarketCappedPriceV3 contract.
type MarketCappedPriceV3OwnershipHandoverRequestedIterator struct {
	Event *MarketCappedPriceV3OwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *MarketCappedPriceV3OwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCappedPriceV3OwnershipHandoverRequested)
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
		it.Event = new(MarketCappedPriceV3OwnershipHandoverRequested)
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
func (it *MarketCappedPriceV3OwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCappedPriceV3OwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCappedPriceV3OwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the MarketCappedPriceV3 contract.
type MarketCappedPriceV3OwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Filterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*MarketCappedPriceV3OwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _MarketCappedPriceV3.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceV3OwnershipHandoverRequestedIterator{contract: _MarketCappedPriceV3.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Filterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *MarketCappedPriceV3OwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _MarketCappedPriceV3.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCappedPriceV3OwnershipHandoverRequested)
				if err := _MarketCappedPriceV3.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_MarketCappedPriceV3 *MarketCappedPriceV3Filterer) ParseOwnershipHandoverRequested(log types.Log) (*MarketCappedPriceV3OwnershipHandoverRequested, error) {
	event := new(MarketCappedPriceV3OwnershipHandoverRequested)
	if err := _MarketCappedPriceV3.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketCappedPriceV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MarketCappedPriceV3 contract.
type MarketCappedPriceV3OwnershipTransferredIterator struct {
	Event *MarketCappedPriceV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketCappedPriceV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCappedPriceV3OwnershipTransferred)
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
		it.Event = new(MarketCappedPriceV3OwnershipTransferred)
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
func (it *MarketCappedPriceV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCappedPriceV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCappedPriceV3OwnershipTransferred represents a OwnershipTransferred event raised by the MarketCappedPriceV3 contract.
type MarketCappedPriceV3OwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*MarketCappedPriceV3OwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketCappedPriceV3.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketCappedPriceV3OwnershipTransferredIterator{contract: _MarketCappedPriceV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_MarketCappedPriceV3 *MarketCappedPriceV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketCappedPriceV3OwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketCappedPriceV3.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCappedPriceV3OwnershipTransferred)
				if err := _MarketCappedPriceV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MarketCappedPriceV3 *MarketCappedPriceV3Filterer) ParseOwnershipTransferred(log types.Log) (*MarketCappedPriceV3OwnershipTransferred, error) {
	event := new(MarketCappedPriceV3OwnershipTransferred)
	if err := _MarketCappedPriceV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
