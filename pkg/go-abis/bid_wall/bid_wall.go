// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bid_wall

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

// BidWallMetaData contains all meta data concerning the BidWall contract.
var BidWallMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_nativeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_protocolOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"closeBidWall\",\"inputs\":[{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_poolKey\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"_ethSwapAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_currentTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"_nativeIsZero\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isBidWallEnabled\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nativeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolInfo\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"initialized\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"pendingETHFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeSwapFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"position\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"amount0_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pendingEth_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPositionManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setDisabledState\",\"inputs\":[{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"_disable\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSwapFeeThreshold\",\"inputs\":[{\"name\":\"swapFeeThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"BidWallClosed\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_eth\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidWallDeposit\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_added\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_pending\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidWallDisabledStateUpdated\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_disabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidWallInitialized\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_eth\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_tickLower\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"},{\"name\":\"_tickUpper\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidWallRepositioned\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_eth\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_tickLower\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"},{\"name\":\"_tickUpper\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidWallRewardsTransferred\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FixedSwapFeeThresholdUpdated\",\"inputs\":[{\"name\":\"_newSwapFeeThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotCreator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPositionManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161280b38038061280b83398101604081905261002e91610109565b3360a0526001600160a01b0383811660c052821660805267016345785d8a00005f8190556040519081527ff33554d7210318fe10b4d834a6233a2fd789ea1db37278b7e27f433b9e8659c79060200160405180910390a161008e81610096565b505050610149565b638b78c6d8198054156100b057630dc149f05f526004601cfd5b6001600160a01b03909116801560ff1b8117909155805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b80516001600160a01b0381168114610104575f5ffd5b919050565b5f5f5f6060848603121561011b575f5ffd5b610124846100ee565b9250610132602085016100ee565b9150610140604085016100ee565b90509250925092565b60805160a05160c0516125fe61020d5f395f8181610313015281816104ed0152818161060b01528181610a0701528181610aae01528181610b5401528181610c0c0152610d7601525f818161015f015281816103a50152818161051a015281816109c501528181610b8101528181610e7c01526111a601525f81816102e0015281816105a7015281816109080152818161095201528181610fa7015281816116620152818161173c015281816117a70152818161184801526118b601526125fe5ff3fe6080604052600436106100cd575f3560e01c8063088a80c7146100d157806325692962146100f25780634e944d57146100fa57806354d1f13d1461013e578063715018a614610146578063791b98bc1461014e578063870b9c171461018e5780638cebd942146101ad5780638da5cb5b1461023f578063957d1fe114610257578063ad49d66e14610291578063cfbc79fe146102b0578063dc4c90d3146102cf578063e1758bd814610302578063f04e283e14610335578063f2fde38b14610348578063fee81cf41461035b575b5f5ffd5b3480156100dc575f5ffd5b506100f06100eb366004612149565b61039a565b005b6100f061079a565b348015610105575f5ffd5b50610129610114366004612197565b5f9081526001602052604090205460ff161590565b60405190151581526020015b60405180910390f35b6100f06107e6565b6100f061081f565b348015610159575f5ffd5b506101817f000000000000000000000000000000000000000000000000000000000000000081565b60405161013591906121ae565b348015610199575f5ffd5b506100f06101a8366004612197565b610832565b3480156101b8575f5ffd5b506102096101c7366004612197565b600160208190525f918252604090912080549181015460029182015460ff80851694610100810490911693620100008204810b93600160281b909204900b9186565b6040805196151587529415156020870152600293840b94860194909452910b6060840152608083015260a082015260c001610135565b34801561024a575f5ffd5b50638b78c6d81954610181565b348015610262575f5ffd5b50610276610271366004612197565b610874565b60408051938452602084019290925290820152606001610135565b34801561029c575f5ffd5b506100f06102ab3660046121c2565b6109ba565b3480156102bb575f5ffd5b506100f06102ca3660046121dc565b610d70565b3480156102da575f5ffd5b506101817f000000000000000000000000000000000000000000000000000000000000000081565b34801561030d575f5ffd5b506101817f000000000000000000000000000000000000000000000000000000000000000081565b6100f0610343366004612212565b610f29565b6100f0610356366004612212565b610f66565b348015610366575f5ffd5b5061038c610375366004612212565b63389a75e1600c9081525f91909152602090205490565b604051908152602001610135565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103e35760405163041fb8cb60e31b815260040160405180910390fd5b82156107945760a084205f8181526001602052604081206002810180549192879261040f908490612241565b9250508190555084816001015f8282546104299190612241565b9091555050600181015460405183917fa5f00d866fe48379f7d8a625ccea2f92572b5cbd3f234591dfa4c7f845836fd09161046c91898252602082015260400190565b60405180910390a25f5481600101541015610488575050610794565b6001810180545f9182905582549091908190610100900460ff161561058b5783546104cb908a908890620100008104600290810b91600160281b9004900b610f8c565b909250905081156105865760405163a9059cbb60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb90610544907f0000000000000000000000000000000000000000000000000000000000000000908690600401612254565b6020604051808303815f875af1158015610560573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610584919061226d565b505b610599565b835461ff0019166101001784555b5f6105cd6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001687611044565b50509150508760020b8160020b131515871515036105e9578097505b6105fe8a888a6105f98888612241565b6110f6565b811561072e575f61062f8b7f0000000000000000000000000000000000000000000000000000000000000000611229565b90505f816001600160a01b03166361d027b36040518163ffffffff1660e01b8152600401602060405180830381865afa15801561066e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106929190612288565b60405163a9059cbb60e01b81529091506001600160a01b0383169063a9059cbb906106c39084908890600401612254565b6020604051808303815f875af11580156106df573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610703919061226d565b50875f5160206125a95f395f51905f528286604051610723929190612254565b60405180910390a250505b857fba69105014fb62310aea3b6cd667c51d18201a3c7112e2ead21829195f6fa3fc61075a8686612241565b875460408051928352620100008204600290810b6020850152600160281b90920490910b9082015260600160405180910390a25050505050505b50505050565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b610827611255565b6108305f61126f565b565b61083a611255565b5f8190556040518181527ff33554d7210318fe10b4d834a6233a2fd789ea1db37278b7e27f433b9e8659c79060200160405180910390a150565b5f818152600160208181526040808420815160c081018352815460ff808216151583526101008204161515948201859052620100008104600290810b94830194909452600160281b9004830b6060820152938101546080850152015460a0830152829182916108ee57608001515f935083925090506109b3565b604081015160608201515f9161093e916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691899130919066189a591dd85b1b60ca1b6112b5565b509091505f90506109786001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001688611044565b50505090506109a18161098e856040015161130c565b61099b866060015161130c565b856115c4565b60809094015190965092945091925050505b9193909250565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a035760405163041fb8cb60e31b815260040160405180910390fd5b80517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081169116145f610a408360a0902090565b5f8181526001602052604081208054929350918190610100900460ff1615610a95578254610a869087908790620100008104600290810b91600160281b9004900b610f8c565b845461ff001916855590925090505b6001830180545f918290556002850182905590610ad2887f0000000000000000000000000000000000000000000000000000000000000000611229565b90505f816001600160a01b03166361d027b36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b11573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b359190612288565b90508215610bef576040516323b872dd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd90610bad907f000000000000000000000000000000000000000000000000000000000000000090859088906004016122a3565b6020604051808303815f875af1158015610bc9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bed919061226d565b505b8415610c855760405163a9059cbb60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb90610c439084908990600401612254565b6020604051808303815f875af1158015610c5f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c83919061226d565b505b8315610d225760405163a9059cbb60e01b81526001600160a01b0383169063a9059cbb90610cb99084908890600401612254565b6020604051808303815f875af1158015610cd5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cf9919061226d565b50865f5160206125a95f395f51905f528286604051610d19929190612254565b60405180910390a25b867f98332f6bdb1b434256cc2c7b313581d76b7763b9c764bd32b25d461ff42c2bc382610d4f8689612241565b604051610d5d929190612254565b60405180910390a2505050505050505050565b610d9a827f0000000000000000000000000000000000000000000000000000000000000000611229565b6001600160a01b03166302d05d3f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dd5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610df99190612288565b6001600160a01b0316336001600160a01b031614610e2a57604051632a118c8960e01b815260040160405180910390fd5b5f60015f610e398560a0902090565b815260208101919091526040015f20805490915060ff16151582151503610e5f57505050565b8115610edf576040516356a4eb3760e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063ad49d66e90610eb190869060040161230a565b5f604051808303815f87803b158015610ec8575f5ffd5b505af1158015610eda573d5f5f3e3d5ffd5b505050505b805460ff191682151517815560a0832060405183151581527f88e359a3e31b529597eae3a8c8a107e55d6ff85edf8ca19a7368a9be35f012719060200160405180910390a2505050565b610f31611255565b63389a75e1600c52805f526020600c208054421115610f5757636f5e88185f526004601cfd5b5f9055610f638161126f565b50565b610f6e611255565b8060601b610f8357637448fbae5f526004601cfd5b610f638161126f565b5f5f5f610fdb610f9d8860a0902090565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169030888866189a591dd85b1b60ca1b6112b5565b505090505f610ff688878785610ff090612318565b3061165f565b9050866110165761100781600f0b90565b6110118260801d90565b61102a565b6110208160801d90565b61102a82600f0b90565b6001600160801b039182169a911698509650505050505050565b5f5f5f5f5f61105286611948565b604051631e2eaeaf60e01b8152600481018290529091505f906001600160a01b03891690631e2eaeaf90602401602060405180830381865afa15801561109a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110be9190612339565b90506001600160a01b03811695508060a01c60020b945062ffffff8160b81c16935062ffffff8160d01c169250505092959194509250565b5f8361110c57611107600184612350565b611117565b611117836001612375565b90505f5f5f861561116157611130600285900b5f611984565b925061113d603c84612375565b915061115a61114b8461130c565b6111548461130c565b87611a1a565b905061119d565b611170600285900b6001611984565b915061117d603c83612350565b925061119a61118b8461130c565b6111948461130c565b87611a8d565b90505b6111ca888484847f000000000000000000000000000000000000000000000000000000000000000061165f565b505f60015f6111da8b60a0902090565b815260208101919091526040015f20805467ffffffffffff000019166201000062ffffff9687160262ffffff60281b191617600160281b94909516939093029390931790915550505050505050565b81515f906001600160a01b0380841691161461124657825161124c565b82602001515b90505b92915050565b638b78c6d819543314610830576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3811560ff1b8217905550565b60408051602681018390526006810184905260038101859052858152603a600c8201205f9282018390526020820183905290829052819081906112f98a8a83611aca565b919c909b50909950975050505050505050565b60020b5f60ff82901d80830118620d89e8811115611335576113356345c3193d60e11b84611b6d565b7001fffcb933bd6fad37aa2d162d1a5940016001821602600160801b186002821615611371576ffff97272373d413259a46990580e213a0260801c5b6004821615611390576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b60088216156113af576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b60108216156113ce576fffcb9843d60f6159c9db58835c9266440260801c5b60208216156113ed576fff973b41fa98c081472e6896dfb254c00260801c5b604082161561140c576fff2ea16466c96a3843ec78b326b528610260801c5b608082161561142b576ffe5dee046a99a2a811c461f1969c30530260801c5b61010082161561144b576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b61020082161561146b576ff987a7253ac413176f2b074cf7815e540260801c5b61040082161561148b576ff3392b0822b70005940c7a398e4b70f30260801c5b6108008216156114ab576fe7159475a2c29b7443b29c7fa6e889d90260801c5b6110008216156114cb576fd097f3bdfd2022b8845ad8f792aa58250260801c5b6120008216156114eb576fa9f746462d870fdf8a65dc1f90e061e50260801c5b61400082161561150b576f70d869a156d2a1b890bb3df62baf32f70260801c5b61800082161561152b576f31be135f97d08fd981231505542fcfa60260801c5b6201000082161561154c576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b6202000082161561156c576e5d6af8dedb81196699c329225ee6040260801c5b6204000082161561158b576d2216e584f5fa1ea926041bedfe980260801c5b620800008216156115a8576b048a170391f7dc42444e8fa20260801c5b5f8413156115b4575f19045b63ffffffff0160201c9392505050565b5f5f836001600160a01b0316856001600160a01b031611156115e4579293925b846001600160a01b0316866001600160a01b03161161160f57611608858585611b7c565b9150611656565b836001600160a01b0316866001600160a01b0316101561164857611634868585611b7c565b9150611641858785611be5565b9050611656565b611653858585611be5565b90505b94509492505050565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635a6bcfda8760405180608001604052808960020b81526020018860020b815260200187600f0b815260200166189a591dd85b1b60ca1b8152506040518363ffffffff1660e01b81526004016116e192919061239a565b60408051808303815f875af11580156116fc573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061172091906123eb565b5090505f61172e8260801d90565b600f0b12156117915761178c7f0000000000000000000000000000000000000000000000000000000000000000836117668460801d90565b61176f90612318565b89516001600160a01b03169291906001600160801b03165f611c2e565b61182f565b5f61179c8260801d90565b600f0b131561182f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630b0d9c09875f0151846117e38560801d90565b6040518463ffffffff1660e01b81526004016118019392919061240d565b5f604051808303815f87803b158015611818575f5ffd5b505af115801561182a573d5f5f3e3d5ffd5b505050505b5f61183a82600f0b90565b600f0b12156118a05761189b7f00000000000000000000000000000000000000000000000000000000000000008361187284600f0b90565b61187b90612318565b60208a01516001600160a01b03169291906001600160801b03165f611c2e565b61193f565b5f6118ab82600f0b90565b600f0b131561193f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630b0d9c098760200151846118f385600f0b90565b6040518463ffffffff1660e01b81526004016119119392919061240d565b5f604051808303815f87803b158015611928575f5ffd5b505af115801561193a573d5f5f3e3d5ffd5b505050505b95945050505050565b6040515f90611967908390600690602001918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b5f620d89b319600284900b12156119a157620d89b31992506119b7565b620d89b4600284900b13156119b757620d89b492505b6119c2603c8461244d565b60020b5f036119d257508161124f565b603c6119de818561246e565b6119e891906124a6565b90505f8360020b1215611a0357611a00603c82612350565b90505b8161124f57611a13603c82612375565b905061124f565b5f826001600160a01b0316846001600160a01b03161115611a39579192915b5f611a5b856001600160a01b0316856001600160a01b0316600160601b611ef2565b9050611a82611a7d8483611a6f89896124cc565b6001600160a01b0316611ef2565b611f8e565b9150505b9392505050565b5f826001600160a01b0316846001600160a01b03161115611aac579192915b611ac2611a7d83600160601b611a6f88886124cc565b949350505050565b5f5f5f5f611ad88686611fe5565b604051631afeb18d60e11b815260048101829052600360248201529091505f906001600160a01b038916906335fd631a906044015f60405180830381865afa158015611b26573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611b4d91908101906124eb565b60208101516040820151606090920151909a919950975095505050505050565b815f528060020b60045260245ffd5b5f826001600160a01b0316846001600160a01b03161115611b9b579192915b6001600160a01b038416611bdb600160601b600160e01b03606085901b16611bc387876124cc565b6001600160a01b0316866001600160a01b0316611ef2565b611ac29190612595565b5f826001600160a01b0316846001600160a01b03161115611c04579192915b611ac26001600160801b038316611c1b86866124cc565b6001600160a01b0316600160601b611ef2565b8015611cc257836001600160a01b031663f5298aca84611c5d886001600160a01b03166001600160a01b031690565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604481018590526064015f604051808303815f87803b158015611ca7575f5ffd5b505af1158015611cb9573d5f5f3e3d5ffd5b50505050611eeb565b6001600160a01b038516611d3957836001600160a01b03166311da60b4836040518263ffffffff1660e01b815260040160206040518083038185885af1158015611d0e573d5f5f3e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190611d339190612339565b50611eeb565b604051632961046560e21b81526001600160a01b0385169063a584119490611d659088906004016121ae565b5f604051808303815f87803b158015611d7c575f5ffd5b505af1158015611d8e573d5f5f3e3d5ffd5b505050506001600160a01b0383163014611e18576040516323b872dd60e01b81526001600160a01b038616906323b872dd90611dd2908690889087906004016122a3565b6020604051808303815f875af1158015611dee573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e12919061226d565b50611e88565b60405163a9059cbb60e01b81526001600160a01b0386169063a9059cbb90611e469087908690600401612254565b6020604051808303815f875af1158015611e62573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e86919061226d565b505b836001600160a01b03166311da60b46040518163ffffffff1660e01b81526004016020604051808303815f875af1158015611ec5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ee99190612339565b505b5050505050565b5f838302815f1985870982811083820303915050808411611f11575f5ffd5b805f03611f2357508290049050611a86565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b806001600160801b0381168114611fe05760405162461bcd60e51b81526020600482015260126024820152716c6971756964697479206f766572666c6f7760701b604482015260640160405180910390fd5b919050565b5f5f611ff084611948565b90505f611ffe600683612241565b6040805160208101879052908101829052909150606001604051602081830303815290604052805190602001209250505092915050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561207157612071612035565b604052919050565b6001600160a01b0381168114610f63575f5ffd5b8035600281900b8114611fe0575f5ffd5b5f60a082840312156120ae575f5ffd5b60405160a081016001600160401b03811182821017156120d0576120d0612035565b60405290508082356120e181612079565b815260208301356120f181612079565b6020820152604083013562ffffff8116811461210b575f5ffd5b604082015261211c6060840161208d565b6060820152608083013561212f81612079565b6080919091015292915050565b8015158114610f63575f5ffd5b5f5f5f5f610100858703121561215d575f5ffd5b612167868661209e565b935060a0850135925061217c60c0860161208d565b915060e085013561218c8161213c565b939692955090935050565b5f602082840312156121a7575f5ffd5b5035919050565b6001600160a01b0391909116815260200190565b5f60a082840312156121d2575f5ffd5b61124c838361209e565b5f5f60c083850312156121ed575f5ffd5b6121f7848461209e565b915060a08301356122078161213c565b809150509250929050565b5f60208284031215612222575f5ffd5b8135611a8681612079565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561124f5761124f61222d565b6001600160a01b03929092168252602082015260400190565b5f6020828403121561227d575f5ffd5b8151611a868161213c565b5f60208284031215612298575f5ffd5b8151611a8681612079565b6001600160a01b039384168152919092166020820152604081019190915260600190565b80516001600160a01b03908116835260208083015182169084015260408083015162ffffff169084015260608083015160020b9084015260809182015116910152565b60a0810161124f82846122c7565b5f600f82900b6001607f1b81016123315761233161222d565b5f0392915050565b5f60208284031215612349575f5ffd5b5051919050565b600282810b9082900b03627fffff198112627fffff8213171561124f5761124f61222d565b600281810b9083900b01627fffff8113627fffff198212171561124f5761124f61222d565b6123a481846122c7565b8151600290810b60a08301526020830151900b60c0820152604082015160e082015260609091015161010082015261014061012082018190525f9082015261016001919050565b5f5f604083850312156123fc575f5ffd5b505080516020909101519092909150565b6001600160a01b0393841681529190921660208201526001600160801b03909116604082015260600190565b634e487b7160e01b5f52601260045260245ffd5b5f8260020b8061245f5761245f612439565b808360020b0791505092915050565b5f8160020b8360020b8061248457612484612439565b627fffff1982145f198214161561249d5761249d61222d565b90059392505050565b5f8260020b8260020b028060020b91508082146124c5576124c561222d565b5092915050565b6001600160a01b03828116828216039081111561124f5761124f61222d565b5f602082840312156124fb575f5ffd5b81516001600160401b03811115612510575f5ffd5b8201601f81018413612520575f5ffd5b80516001600160401b0381111561253957612539612035565b8060051b61254960208201612049565b91825260208184018101929081019087841115612564575f5ffd5b6020850194505b8385101561258a5784518083526020958601959093509091019061256b565b979650505050505050565b5f826125a3576125a3612439565b50049056fe6f25beea688da23574729528d9040ac00b2a9f98fe04601175b5bed75222d37ba2646970667358221220de9d4bfc6feccedaea14cf8f5936327299730ff4b5a7fd076ae086207ef0c3d664736f6c634300081b0033",
}

// BidWallABI is the input ABI used to generate the binding from.
// Deprecated: Use BidWallMetaData.ABI instead.
var BidWallABI = BidWallMetaData.ABI

// BidWallBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BidWallMetaData.Bin instead.
var BidWallBin = BidWallMetaData.Bin

// DeployBidWall deploys a new Ethereum contract, binding an instance of BidWall to it.
func DeployBidWall(auth *bind.TransactOpts, backend bind.ContractBackend, _nativeToken common.Address, _poolManager common.Address, _protocolOwner common.Address) (common.Address, *types.Transaction, *BidWall, error) {
	parsed, err := BidWallMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BidWallBin), backend, _nativeToken, _poolManager, _protocolOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BidWall{BidWallCaller: BidWallCaller{contract: contract}, BidWallTransactor: BidWallTransactor{contract: contract}, BidWallFilterer: BidWallFilterer{contract: contract}}, nil
}

// BidWall is an auto generated Go binding around an Ethereum contract.
type BidWall struct {
	BidWallCaller     // Read-only binding to the contract
	BidWallTransactor // Write-only binding to the contract
	BidWallFilterer   // Log filterer for contract events
}

// BidWallCaller is an auto generated read-only Go binding around an Ethereum contract.
type BidWallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BidWallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BidWallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BidWallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BidWallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BidWallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BidWallSession struct {
	Contract     *BidWall          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BidWallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BidWallCallerSession struct {
	Contract *BidWallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BidWallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BidWallTransactorSession struct {
	Contract     *BidWallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BidWallRaw is an auto generated low-level Go binding around an Ethereum contract.
type BidWallRaw struct {
	Contract *BidWall // Generic contract binding to access the raw methods on
}

// BidWallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BidWallCallerRaw struct {
	Contract *BidWallCaller // Generic read-only contract binding to access the raw methods on
}

// BidWallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BidWallTransactorRaw struct {
	Contract *BidWallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBidWall creates a new instance of BidWall, bound to a specific deployed contract.
func NewBidWall(address common.Address, backend bind.ContractBackend) (*BidWall, error) {
	contract, err := bindBidWall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BidWall{BidWallCaller: BidWallCaller{contract: contract}, BidWallTransactor: BidWallTransactor{contract: contract}, BidWallFilterer: BidWallFilterer{contract: contract}}, nil
}

// NewBidWallCaller creates a new read-only instance of BidWall, bound to a specific deployed contract.
func NewBidWallCaller(address common.Address, caller bind.ContractCaller) (*BidWallCaller, error) {
	contract, err := bindBidWall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BidWallCaller{contract: contract}, nil
}

// NewBidWallTransactor creates a new write-only instance of BidWall, bound to a specific deployed contract.
func NewBidWallTransactor(address common.Address, transactor bind.ContractTransactor) (*BidWallTransactor, error) {
	contract, err := bindBidWall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BidWallTransactor{contract: contract}, nil
}

// NewBidWallFilterer creates a new log filterer instance of BidWall, bound to a specific deployed contract.
func NewBidWallFilterer(address common.Address, filterer bind.ContractFilterer) (*BidWallFilterer, error) {
	contract, err := bindBidWall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BidWallFilterer{contract: contract}, nil
}

// bindBidWall binds a generic wrapper to an already deployed contract.
func bindBidWall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BidWallMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BidWall *BidWallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BidWall.Contract.BidWallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BidWall *BidWallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BidWall.Contract.BidWallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BidWall *BidWallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BidWall.Contract.BidWallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BidWall *BidWallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BidWall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BidWall *BidWallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BidWall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BidWall *BidWallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BidWall.Contract.contract.Transact(opts, method, params...)
}

// IsBidWallEnabled is a free data retrieval call binding the contract method 0x4e944d57.
//
// Solidity: function isBidWallEnabled(bytes32 _poolId) view returns(bool)
func (_BidWall *BidWallCaller) IsBidWallEnabled(opts *bind.CallOpts, _poolId [32]byte) (bool, error) {
	var out []interface{}
	err := _BidWall.contract.Call(opts, &out, "isBidWallEnabled", _poolId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBidWallEnabled is a free data retrieval call binding the contract method 0x4e944d57.
//
// Solidity: function isBidWallEnabled(bytes32 _poolId) view returns(bool)
func (_BidWall *BidWallSession) IsBidWallEnabled(_poolId [32]byte) (bool, error) {
	return _BidWall.Contract.IsBidWallEnabled(&_BidWall.CallOpts, _poolId)
}

// IsBidWallEnabled is a free data retrieval call binding the contract method 0x4e944d57.
//
// Solidity: function isBidWallEnabled(bytes32 _poolId) view returns(bool)
func (_BidWall *BidWallCallerSession) IsBidWallEnabled(_poolId [32]byte) (bool, error) {
	return _BidWall.Contract.IsBidWallEnabled(&_BidWall.CallOpts, _poolId)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_BidWall *BidWallCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BidWall.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_BidWall *BidWallSession) NativeToken() (common.Address, error) {
	return _BidWall.Contract.NativeToken(&_BidWall.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_BidWall *BidWallCallerSession) NativeToken() (common.Address, error) {
	return _BidWall.Contract.NativeToken(&_BidWall.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_BidWall *BidWallCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BidWall.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_BidWall *BidWallSession) Owner() (common.Address, error) {
	return _BidWall.Contract.Owner(&_BidWall.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_BidWall *BidWallCallerSession) Owner() (common.Address, error) {
	return _BidWall.Contract.Owner(&_BidWall.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_BidWall *BidWallCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BidWall.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_BidWall *BidWallSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _BidWall.Contract.OwnershipHandoverExpiresAt(&_BidWall.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_BidWall *BidWallCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _BidWall.Contract.OwnershipHandoverExpiresAt(&_BidWall.CallOpts, pendingOwner)
}

// PoolInfo is a free data retrieval call binding the contract method 0x8cebd942.
//
// Solidity: function poolInfo(bytes32 _poolId) view returns(bool disabled, bool initialized, int24 tickLower, int24 tickUpper, uint256 pendingETHFees, uint256 cumulativeSwapFees)
func (_BidWall *BidWallCaller) PoolInfo(opts *bind.CallOpts, _poolId [32]byte) (struct {
	Disabled           bool
	Initialized        bool
	TickLower          *big.Int
	TickUpper          *big.Int
	PendingETHFees     *big.Int
	CumulativeSwapFees *big.Int
}, error) {
	var out []interface{}
	err := _BidWall.contract.Call(opts, &out, "poolInfo", _poolId)

	outstruct := new(struct {
		Disabled           bool
		Initialized        bool
		TickLower          *big.Int
		TickUpper          *big.Int
		PendingETHFees     *big.Int
		CumulativeSwapFees *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Disabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Initialized = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.TickLower = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TickUpper = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PendingETHFees = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CumulativeSwapFees = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x8cebd942.
//
// Solidity: function poolInfo(bytes32 _poolId) view returns(bool disabled, bool initialized, int24 tickLower, int24 tickUpper, uint256 pendingETHFees, uint256 cumulativeSwapFees)
func (_BidWall *BidWallSession) PoolInfo(_poolId [32]byte) (struct {
	Disabled           bool
	Initialized        bool
	TickLower          *big.Int
	TickUpper          *big.Int
	PendingETHFees     *big.Int
	CumulativeSwapFees *big.Int
}, error) {
	return _BidWall.Contract.PoolInfo(&_BidWall.CallOpts, _poolId)
}

// PoolInfo is a free data retrieval call binding the contract method 0x8cebd942.
//
// Solidity: function poolInfo(bytes32 _poolId) view returns(bool disabled, bool initialized, int24 tickLower, int24 tickUpper, uint256 pendingETHFees, uint256 cumulativeSwapFees)
func (_BidWall *BidWallCallerSession) PoolInfo(_poolId [32]byte) (struct {
	Disabled           bool
	Initialized        bool
	TickLower          *big.Int
	TickUpper          *big.Int
	PendingETHFees     *big.Int
	CumulativeSwapFees *big.Int
}, error) {
	return _BidWall.Contract.PoolInfo(&_BidWall.CallOpts, _poolId)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_BidWall *BidWallCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BidWall.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_BidWall *BidWallSession) PoolManager() (common.Address, error) {
	return _BidWall.Contract.PoolManager(&_BidWall.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_BidWall *BidWallCallerSession) PoolManager() (common.Address, error) {
	return _BidWall.Contract.PoolManager(&_BidWall.CallOpts)
}

// Position is a free data retrieval call binding the contract method 0x957d1fe1.
//
// Solidity: function position(bytes32 _poolId) view returns(uint256 amount0_, uint256 amount1_, uint256 pendingEth_)
func (_BidWall *BidWallCaller) Position(opts *bind.CallOpts, _poolId [32]byte) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	PendingEth *big.Int
}, error) {
	var out []interface{}
	err := _BidWall.contract.Call(opts, &out, "position", _poolId)

	outstruct := new(struct {
		Amount0    *big.Int
		Amount1    *big.Int
		PendingEth *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PendingEth = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Position is a free data retrieval call binding the contract method 0x957d1fe1.
//
// Solidity: function position(bytes32 _poolId) view returns(uint256 amount0_, uint256 amount1_, uint256 pendingEth_)
func (_BidWall *BidWallSession) Position(_poolId [32]byte) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	PendingEth *big.Int
}, error) {
	return _BidWall.Contract.Position(&_BidWall.CallOpts, _poolId)
}

// Position is a free data retrieval call binding the contract method 0x957d1fe1.
//
// Solidity: function position(bytes32 _poolId) view returns(uint256 amount0_, uint256 amount1_, uint256 pendingEth_)
func (_BidWall *BidWallCallerSession) Position(_poolId [32]byte) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	PendingEth *big.Int
}, error) {
	return _BidWall.Contract.Position(&_BidWall.CallOpts, _poolId)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_BidWall *BidWallCaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BidWall.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_BidWall *BidWallSession) PositionManager() (common.Address, error) {
	return _BidWall.Contract.PositionManager(&_BidWall.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_BidWall *BidWallCallerSession) PositionManager() (common.Address, error) {
	return _BidWall.Contract.PositionManager(&_BidWall.CallOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_BidWall *BidWallTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BidWall.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_BidWall *BidWallSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _BidWall.Contract.CancelOwnershipHandover(&_BidWall.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_BidWall *BidWallTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _BidWall.Contract.CancelOwnershipHandover(&_BidWall.TransactOpts)
}

// CloseBidWall is a paid mutator transaction binding the contract method 0xad49d66e.
//
// Solidity: function closeBidWall((address,address,uint24,int24,address) _key) returns()
func (_BidWall *BidWallTransactor) CloseBidWall(opts *bind.TransactOpts, _key PoolKey) (*types.Transaction, error) {
	return _BidWall.contract.Transact(opts, "closeBidWall", _key)
}

// CloseBidWall is a paid mutator transaction binding the contract method 0xad49d66e.
//
// Solidity: function closeBidWall((address,address,uint24,int24,address) _key) returns()
func (_BidWall *BidWallSession) CloseBidWall(_key PoolKey) (*types.Transaction, error) {
	return _BidWall.Contract.CloseBidWall(&_BidWall.TransactOpts, _key)
}

// CloseBidWall is a paid mutator transaction binding the contract method 0xad49d66e.
//
// Solidity: function closeBidWall((address,address,uint24,int24,address) _key) returns()
func (_BidWall *BidWallTransactorSession) CloseBidWall(_key PoolKey) (*types.Transaction, error) {
	return _BidWall.Contract.CloseBidWall(&_BidWall.TransactOpts, _key)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_BidWall *BidWallTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _BidWall.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_BidWall *BidWallSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _BidWall.Contract.CompleteOwnershipHandover(&_BidWall.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_BidWall *BidWallTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _BidWall.Contract.CompleteOwnershipHandover(&_BidWall.TransactOpts, pendingOwner)
}

// Deposit is a paid mutator transaction binding the contract method 0x088a80c7.
//
// Solidity: function deposit((address,address,uint24,int24,address) _poolKey, uint256 _ethSwapAmount, int24 _currentTick, bool _nativeIsZero) returns()
func (_BidWall *BidWallTransactor) Deposit(opts *bind.TransactOpts, _poolKey PoolKey, _ethSwapAmount *big.Int, _currentTick *big.Int, _nativeIsZero bool) (*types.Transaction, error) {
	return _BidWall.contract.Transact(opts, "deposit", _poolKey, _ethSwapAmount, _currentTick, _nativeIsZero)
}

// Deposit is a paid mutator transaction binding the contract method 0x088a80c7.
//
// Solidity: function deposit((address,address,uint24,int24,address) _poolKey, uint256 _ethSwapAmount, int24 _currentTick, bool _nativeIsZero) returns()
func (_BidWall *BidWallSession) Deposit(_poolKey PoolKey, _ethSwapAmount *big.Int, _currentTick *big.Int, _nativeIsZero bool) (*types.Transaction, error) {
	return _BidWall.Contract.Deposit(&_BidWall.TransactOpts, _poolKey, _ethSwapAmount, _currentTick, _nativeIsZero)
}

// Deposit is a paid mutator transaction binding the contract method 0x088a80c7.
//
// Solidity: function deposit((address,address,uint24,int24,address) _poolKey, uint256 _ethSwapAmount, int24 _currentTick, bool _nativeIsZero) returns()
func (_BidWall *BidWallTransactorSession) Deposit(_poolKey PoolKey, _ethSwapAmount *big.Int, _currentTick *big.Int, _nativeIsZero bool) (*types.Transaction, error) {
	return _BidWall.Contract.Deposit(&_BidWall.TransactOpts, _poolKey, _ethSwapAmount, _currentTick, _nativeIsZero)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_BidWall *BidWallTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BidWall.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_BidWall *BidWallSession) RenounceOwnership() (*types.Transaction, error) {
	return _BidWall.Contract.RenounceOwnership(&_BidWall.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_BidWall *BidWallTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BidWall.Contract.RenounceOwnership(&_BidWall.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_BidWall *BidWallTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BidWall.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_BidWall *BidWallSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _BidWall.Contract.RequestOwnershipHandover(&_BidWall.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_BidWall *BidWallTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _BidWall.Contract.RequestOwnershipHandover(&_BidWall.TransactOpts)
}

// SetDisabledState is a paid mutator transaction binding the contract method 0xcfbc79fe.
//
// Solidity: function setDisabledState((address,address,uint24,int24,address) _key, bool _disable) returns()
func (_BidWall *BidWallTransactor) SetDisabledState(opts *bind.TransactOpts, _key PoolKey, _disable bool) (*types.Transaction, error) {
	return _BidWall.contract.Transact(opts, "setDisabledState", _key, _disable)
}

// SetDisabledState is a paid mutator transaction binding the contract method 0xcfbc79fe.
//
// Solidity: function setDisabledState((address,address,uint24,int24,address) _key, bool _disable) returns()
func (_BidWall *BidWallSession) SetDisabledState(_key PoolKey, _disable bool) (*types.Transaction, error) {
	return _BidWall.Contract.SetDisabledState(&_BidWall.TransactOpts, _key, _disable)
}

// SetDisabledState is a paid mutator transaction binding the contract method 0xcfbc79fe.
//
// Solidity: function setDisabledState((address,address,uint24,int24,address) _key, bool _disable) returns()
func (_BidWall *BidWallTransactorSession) SetDisabledState(_key PoolKey, _disable bool) (*types.Transaction, error) {
	return _BidWall.Contract.SetDisabledState(&_BidWall.TransactOpts, _key, _disable)
}

// SetSwapFeeThreshold is a paid mutator transaction binding the contract method 0x870b9c17.
//
// Solidity: function setSwapFeeThreshold(uint256 swapFeeThreshold) returns()
func (_BidWall *BidWallTransactor) SetSwapFeeThreshold(opts *bind.TransactOpts, swapFeeThreshold *big.Int) (*types.Transaction, error) {
	return _BidWall.contract.Transact(opts, "setSwapFeeThreshold", swapFeeThreshold)
}

// SetSwapFeeThreshold is a paid mutator transaction binding the contract method 0x870b9c17.
//
// Solidity: function setSwapFeeThreshold(uint256 swapFeeThreshold) returns()
func (_BidWall *BidWallSession) SetSwapFeeThreshold(swapFeeThreshold *big.Int) (*types.Transaction, error) {
	return _BidWall.Contract.SetSwapFeeThreshold(&_BidWall.TransactOpts, swapFeeThreshold)
}

// SetSwapFeeThreshold is a paid mutator transaction binding the contract method 0x870b9c17.
//
// Solidity: function setSwapFeeThreshold(uint256 swapFeeThreshold) returns()
func (_BidWall *BidWallTransactorSession) SetSwapFeeThreshold(swapFeeThreshold *big.Int) (*types.Transaction, error) {
	return _BidWall.Contract.SetSwapFeeThreshold(&_BidWall.TransactOpts, swapFeeThreshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_BidWall *BidWallTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BidWall.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_BidWall *BidWallSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BidWall.Contract.TransferOwnership(&_BidWall.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_BidWall *BidWallTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BidWall.Contract.TransferOwnership(&_BidWall.TransactOpts, newOwner)
}

// BidWallBidWallClosedIterator is returned from FilterBidWallClosed and is used to iterate over the raw logs and unpacked data for BidWallClosed events raised by the BidWall contract.
type BidWallBidWallClosedIterator struct {
	Event *BidWallBidWallClosed // Event containing the contract specifics and raw log

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
func (it *BidWallBidWallClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidWallBidWallClosed)
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
		it.Event = new(BidWallBidWallClosed)
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
func (it *BidWallBidWallClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidWallBidWallClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidWallBidWallClosed represents a BidWallClosed event raised by the BidWall contract.
type BidWallBidWallClosed struct {
	PoolId    [32]byte
	Recipient common.Address
	Eth       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidWallClosed is a free log retrieval operation binding the contract event 0x98332f6bdb1b434256cc2c7b313581d76b7763b9c764bd32b25d461ff42c2bc3.
//
// Solidity: event BidWallClosed(bytes32 indexed _poolId, address _recipient, uint256 _eth)
func (_BidWall *BidWallFilterer) FilterBidWallClosed(opts *bind.FilterOpts, _poolId [][32]byte) (*BidWallBidWallClosedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.FilterLogs(opts, "BidWallClosed", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &BidWallBidWallClosedIterator{contract: _BidWall.contract, event: "BidWallClosed", logs: logs, sub: sub}, nil
}

// WatchBidWallClosed is a free log subscription operation binding the contract event 0x98332f6bdb1b434256cc2c7b313581d76b7763b9c764bd32b25d461ff42c2bc3.
//
// Solidity: event BidWallClosed(bytes32 indexed _poolId, address _recipient, uint256 _eth)
func (_BidWall *BidWallFilterer) WatchBidWallClosed(opts *bind.WatchOpts, sink chan<- *BidWallBidWallClosed, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.WatchLogs(opts, "BidWallClosed", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidWallBidWallClosed)
				if err := _BidWall.contract.UnpackLog(event, "BidWallClosed", log); err != nil {
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

// ParseBidWallClosed is a log parse operation binding the contract event 0x98332f6bdb1b434256cc2c7b313581d76b7763b9c764bd32b25d461ff42c2bc3.
//
// Solidity: event BidWallClosed(bytes32 indexed _poolId, address _recipient, uint256 _eth)
func (_BidWall *BidWallFilterer) ParseBidWallClosed(log types.Log) (*BidWallBidWallClosed, error) {
	event := new(BidWallBidWallClosed)
	if err := _BidWall.contract.UnpackLog(event, "BidWallClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BidWallBidWallDepositIterator is returned from FilterBidWallDeposit and is used to iterate over the raw logs and unpacked data for BidWallDeposit events raised by the BidWall contract.
type BidWallBidWallDepositIterator struct {
	Event *BidWallBidWallDeposit // Event containing the contract specifics and raw log

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
func (it *BidWallBidWallDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidWallBidWallDeposit)
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
		it.Event = new(BidWallBidWallDeposit)
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
func (it *BidWallBidWallDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidWallBidWallDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidWallBidWallDeposit represents a BidWallDeposit event raised by the BidWall contract.
type BidWallBidWallDeposit struct {
	PoolId  [32]byte
	Added   *big.Int
	Pending *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBidWallDeposit is a free log retrieval operation binding the contract event 0xa5f00d866fe48379f7d8a625ccea2f92572b5cbd3f234591dfa4c7f845836fd0.
//
// Solidity: event BidWallDeposit(bytes32 indexed _poolId, uint256 _added, uint256 _pending)
func (_BidWall *BidWallFilterer) FilterBidWallDeposit(opts *bind.FilterOpts, _poolId [][32]byte) (*BidWallBidWallDepositIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.FilterLogs(opts, "BidWallDeposit", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &BidWallBidWallDepositIterator{contract: _BidWall.contract, event: "BidWallDeposit", logs: logs, sub: sub}, nil
}

// WatchBidWallDeposit is a free log subscription operation binding the contract event 0xa5f00d866fe48379f7d8a625ccea2f92572b5cbd3f234591dfa4c7f845836fd0.
//
// Solidity: event BidWallDeposit(bytes32 indexed _poolId, uint256 _added, uint256 _pending)
func (_BidWall *BidWallFilterer) WatchBidWallDeposit(opts *bind.WatchOpts, sink chan<- *BidWallBidWallDeposit, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.WatchLogs(opts, "BidWallDeposit", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidWallBidWallDeposit)
				if err := _BidWall.contract.UnpackLog(event, "BidWallDeposit", log); err != nil {
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

// ParseBidWallDeposit is a log parse operation binding the contract event 0xa5f00d866fe48379f7d8a625ccea2f92572b5cbd3f234591dfa4c7f845836fd0.
//
// Solidity: event BidWallDeposit(bytes32 indexed _poolId, uint256 _added, uint256 _pending)
func (_BidWall *BidWallFilterer) ParseBidWallDeposit(log types.Log) (*BidWallBidWallDeposit, error) {
	event := new(BidWallBidWallDeposit)
	if err := _BidWall.contract.UnpackLog(event, "BidWallDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BidWallBidWallDisabledStateUpdatedIterator is returned from FilterBidWallDisabledStateUpdated and is used to iterate over the raw logs and unpacked data for BidWallDisabledStateUpdated events raised by the BidWall contract.
type BidWallBidWallDisabledStateUpdatedIterator struct {
	Event *BidWallBidWallDisabledStateUpdated // Event containing the contract specifics and raw log

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
func (it *BidWallBidWallDisabledStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidWallBidWallDisabledStateUpdated)
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
		it.Event = new(BidWallBidWallDisabledStateUpdated)
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
func (it *BidWallBidWallDisabledStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidWallBidWallDisabledStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidWallBidWallDisabledStateUpdated represents a BidWallDisabledStateUpdated event raised by the BidWall contract.
type BidWallBidWallDisabledStateUpdated struct {
	PoolId   [32]byte
	Disabled bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBidWallDisabledStateUpdated is a free log retrieval operation binding the contract event 0x88e359a3e31b529597eae3a8c8a107e55d6ff85edf8ca19a7368a9be35f01271.
//
// Solidity: event BidWallDisabledStateUpdated(bytes32 indexed _poolId, bool _disabled)
func (_BidWall *BidWallFilterer) FilterBidWallDisabledStateUpdated(opts *bind.FilterOpts, _poolId [][32]byte) (*BidWallBidWallDisabledStateUpdatedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.FilterLogs(opts, "BidWallDisabledStateUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &BidWallBidWallDisabledStateUpdatedIterator{contract: _BidWall.contract, event: "BidWallDisabledStateUpdated", logs: logs, sub: sub}, nil
}

// WatchBidWallDisabledStateUpdated is a free log subscription operation binding the contract event 0x88e359a3e31b529597eae3a8c8a107e55d6ff85edf8ca19a7368a9be35f01271.
//
// Solidity: event BidWallDisabledStateUpdated(bytes32 indexed _poolId, bool _disabled)
func (_BidWall *BidWallFilterer) WatchBidWallDisabledStateUpdated(opts *bind.WatchOpts, sink chan<- *BidWallBidWallDisabledStateUpdated, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.WatchLogs(opts, "BidWallDisabledStateUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidWallBidWallDisabledStateUpdated)
				if err := _BidWall.contract.UnpackLog(event, "BidWallDisabledStateUpdated", log); err != nil {
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

// ParseBidWallDisabledStateUpdated is a log parse operation binding the contract event 0x88e359a3e31b529597eae3a8c8a107e55d6ff85edf8ca19a7368a9be35f01271.
//
// Solidity: event BidWallDisabledStateUpdated(bytes32 indexed _poolId, bool _disabled)
func (_BidWall *BidWallFilterer) ParseBidWallDisabledStateUpdated(log types.Log) (*BidWallBidWallDisabledStateUpdated, error) {
	event := new(BidWallBidWallDisabledStateUpdated)
	if err := _BidWall.contract.UnpackLog(event, "BidWallDisabledStateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BidWallBidWallInitializedIterator is returned from FilterBidWallInitialized and is used to iterate over the raw logs and unpacked data for BidWallInitialized events raised by the BidWall contract.
type BidWallBidWallInitializedIterator struct {
	Event *BidWallBidWallInitialized // Event containing the contract specifics and raw log

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
func (it *BidWallBidWallInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidWallBidWallInitialized)
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
		it.Event = new(BidWallBidWallInitialized)
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
func (it *BidWallBidWallInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidWallBidWallInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidWallBidWallInitialized represents a BidWallInitialized event raised by the BidWall contract.
type BidWallBidWallInitialized struct {
	PoolId    [32]byte
	Eth       *big.Int
	TickLower *big.Int
	TickUpper *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidWallInitialized is a free log retrieval operation binding the contract event 0x855c0ad063cdf9690cf213083eb2d17055810b644b4ca1cbfd7323cf7a76d3c1.
//
// Solidity: event BidWallInitialized(bytes32 indexed _poolId, uint256 _eth, int24 _tickLower, int24 _tickUpper)
func (_BidWall *BidWallFilterer) FilterBidWallInitialized(opts *bind.FilterOpts, _poolId [][32]byte) (*BidWallBidWallInitializedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.FilterLogs(opts, "BidWallInitialized", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &BidWallBidWallInitializedIterator{contract: _BidWall.contract, event: "BidWallInitialized", logs: logs, sub: sub}, nil
}

// WatchBidWallInitialized is a free log subscription operation binding the contract event 0x855c0ad063cdf9690cf213083eb2d17055810b644b4ca1cbfd7323cf7a76d3c1.
//
// Solidity: event BidWallInitialized(bytes32 indexed _poolId, uint256 _eth, int24 _tickLower, int24 _tickUpper)
func (_BidWall *BidWallFilterer) WatchBidWallInitialized(opts *bind.WatchOpts, sink chan<- *BidWallBidWallInitialized, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.WatchLogs(opts, "BidWallInitialized", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidWallBidWallInitialized)
				if err := _BidWall.contract.UnpackLog(event, "BidWallInitialized", log); err != nil {
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

// ParseBidWallInitialized is a log parse operation binding the contract event 0x855c0ad063cdf9690cf213083eb2d17055810b644b4ca1cbfd7323cf7a76d3c1.
//
// Solidity: event BidWallInitialized(bytes32 indexed _poolId, uint256 _eth, int24 _tickLower, int24 _tickUpper)
func (_BidWall *BidWallFilterer) ParseBidWallInitialized(log types.Log) (*BidWallBidWallInitialized, error) {
	event := new(BidWallBidWallInitialized)
	if err := _BidWall.contract.UnpackLog(event, "BidWallInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BidWallBidWallRepositionedIterator is returned from FilterBidWallRepositioned and is used to iterate over the raw logs and unpacked data for BidWallRepositioned events raised by the BidWall contract.
type BidWallBidWallRepositionedIterator struct {
	Event *BidWallBidWallRepositioned // Event containing the contract specifics and raw log

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
func (it *BidWallBidWallRepositionedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidWallBidWallRepositioned)
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
		it.Event = new(BidWallBidWallRepositioned)
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
func (it *BidWallBidWallRepositionedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidWallBidWallRepositionedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidWallBidWallRepositioned represents a BidWallRepositioned event raised by the BidWall contract.
type BidWallBidWallRepositioned struct {
	PoolId    [32]byte
	Eth       *big.Int
	TickLower *big.Int
	TickUpper *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidWallRepositioned is a free log retrieval operation binding the contract event 0xba69105014fb62310aea3b6cd667c51d18201a3c7112e2ead21829195f6fa3fc.
//
// Solidity: event BidWallRepositioned(bytes32 indexed _poolId, uint256 _eth, int24 _tickLower, int24 _tickUpper)
func (_BidWall *BidWallFilterer) FilterBidWallRepositioned(opts *bind.FilterOpts, _poolId [][32]byte) (*BidWallBidWallRepositionedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.FilterLogs(opts, "BidWallRepositioned", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &BidWallBidWallRepositionedIterator{contract: _BidWall.contract, event: "BidWallRepositioned", logs: logs, sub: sub}, nil
}

// WatchBidWallRepositioned is a free log subscription operation binding the contract event 0xba69105014fb62310aea3b6cd667c51d18201a3c7112e2ead21829195f6fa3fc.
//
// Solidity: event BidWallRepositioned(bytes32 indexed _poolId, uint256 _eth, int24 _tickLower, int24 _tickUpper)
func (_BidWall *BidWallFilterer) WatchBidWallRepositioned(opts *bind.WatchOpts, sink chan<- *BidWallBidWallRepositioned, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.WatchLogs(opts, "BidWallRepositioned", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidWallBidWallRepositioned)
				if err := _BidWall.contract.UnpackLog(event, "BidWallRepositioned", log); err != nil {
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

// ParseBidWallRepositioned is a log parse operation binding the contract event 0xba69105014fb62310aea3b6cd667c51d18201a3c7112e2ead21829195f6fa3fc.
//
// Solidity: event BidWallRepositioned(bytes32 indexed _poolId, uint256 _eth, int24 _tickLower, int24 _tickUpper)
func (_BidWall *BidWallFilterer) ParseBidWallRepositioned(log types.Log) (*BidWallBidWallRepositioned, error) {
	event := new(BidWallBidWallRepositioned)
	if err := _BidWall.contract.UnpackLog(event, "BidWallRepositioned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BidWallBidWallRewardsTransferredIterator is returned from FilterBidWallRewardsTransferred and is used to iterate over the raw logs and unpacked data for BidWallRewardsTransferred events raised by the BidWall contract.
type BidWallBidWallRewardsTransferredIterator struct {
	Event *BidWallBidWallRewardsTransferred // Event containing the contract specifics and raw log

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
func (it *BidWallBidWallRewardsTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidWallBidWallRewardsTransferred)
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
		it.Event = new(BidWallBidWallRewardsTransferred)
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
func (it *BidWallBidWallRewardsTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidWallBidWallRewardsTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidWallBidWallRewardsTransferred represents a BidWallRewardsTransferred event raised by the BidWall contract.
type BidWallBidWallRewardsTransferred struct {
	PoolId    [32]byte
	Recipient common.Address
	Tokens    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidWallRewardsTransferred is a free log retrieval operation binding the contract event 0x6f25beea688da23574729528d9040ac00b2a9f98fe04601175b5bed75222d37b.
//
// Solidity: event BidWallRewardsTransferred(bytes32 indexed _poolId, address _recipient, uint256 _tokens)
func (_BidWall *BidWallFilterer) FilterBidWallRewardsTransferred(opts *bind.FilterOpts, _poolId [][32]byte) (*BidWallBidWallRewardsTransferredIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.FilterLogs(opts, "BidWallRewardsTransferred", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &BidWallBidWallRewardsTransferredIterator{contract: _BidWall.contract, event: "BidWallRewardsTransferred", logs: logs, sub: sub}, nil
}

// WatchBidWallRewardsTransferred is a free log subscription operation binding the contract event 0x6f25beea688da23574729528d9040ac00b2a9f98fe04601175b5bed75222d37b.
//
// Solidity: event BidWallRewardsTransferred(bytes32 indexed _poolId, address _recipient, uint256 _tokens)
func (_BidWall *BidWallFilterer) WatchBidWallRewardsTransferred(opts *bind.WatchOpts, sink chan<- *BidWallBidWallRewardsTransferred, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _BidWall.contract.WatchLogs(opts, "BidWallRewardsTransferred", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidWallBidWallRewardsTransferred)
				if err := _BidWall.contract.UnpackLog(event, "BidWallRewardsTransferred", log); err != nil {
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

// ParseBidWallRewardsTransferred is a log parse operation binding the contract event 0x6f25beea688da23574729528d9040ac00b2a9f98fe04601175b5bed75222d37b.
//
// Solidity: event BidWallRewardsTransferred(bytes32 indexed _poolId, address _recipient, uint256 _tokens)
func (_BidWall *BidWallFilterer) ParseBidWallRewardsTransferred(log types.Log) (*BidWallBidWallRewardsTransferred, error) {
	event := new(BidWallBidWallRewardsTransferred)
	if err := _BidWall.contract.UnpackLog(event, "BidWallRewardsTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BidWallFixedSwapFeeThresholdUpdatedIterator is returned from FilterFixedSwapFeeThresholdUpdated and is used to iterate over the raw logs and unpacked data for FixedSwapFeeThresholdUpdated events raised by the BidWall contract.
type BidWallFixedSwapFeeThresholdUpdatedIterator struct {
	Event *BidWallFixedSwapFeeThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *BidWallFixedSwapFeeThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidWallFixedSwapFeeThresholdUpdated)
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
		it.Event = new(BidWallFixedSwapFeeThresholdUpdated)
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
func (it *BidWallFixedSwapFeeThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidWallFixedSwapFeeThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidWallFixedSwapFeeThresholdUpdated represents a FixedSwapFeeThresholdUpdated event raised by the BidWall contract.
type BidWallFixedSwapFeeThresholdUpdated struct {
	NewSwapFeeThreshold *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFixedSwapFeeThresholdUpdated is a free log retrieval operation binding the contract event 0xf33554d7210318fe10b4d834a6233a2fd789ea1db37278b7e27f433b9e8659c7.
//
// Solidity: event FixedSwapFeeThresholdUpdated(uint256 _newSwapFeeThreshold)
func (_BidWall *BidWallFilterer) FilterFixedSwapFeeThresholdUpdated(opts *bind.FilterOpts) (*BidWallFixedSwapFeeThresholdUpdatedIterator, error) {

	logs, sub, err := _BidWall.contract.FilterLogs(opts, "FixedSwapFeeThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &BidWallFixedSwapFeeThresholdUpdatedIterator{contract: _BidWall.contract, event: "FixedSwapFeeThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchFixedSwapFeeThresholdUpdated is a free log subscription operation binding the contract event 0xf33554d7210318fe10b4d834a6233a2fd789ea1db37278b7e27f433b9e8659c7.
//
// Solidity: event FixedSwapFeeThresholdUpdated(uint256 _newSwapFeeThreshold)
func (_BidWall *BidWallFilterer) WatchFixedSwapFeeThresholdUpdated(opts *bind.WatchOpts, sink chan<- *BidWallFixedSwapFeeThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _BidWall.contract.WatchLogs(opts, "FixedSwapFeeThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidWallFixedSwapFeeThresholdUpdated)
				if err := _BidWall.contract.UnpackLog(event, "FixedSwapFeeThresholdUpdated", log); err != nil {
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

// ParseFixedSwapFeeThresholdUpdated is a log parse operation binding the contract event 0xf33554d7210318fe10b4d834a6233a2fd789ea1db37278b7e27f433b9e8659c7.
//
// Solidity: event FixedSwapFeeThresholdUpdated(uint256 _newSwapFeeThreshold)
func (_BidWall *BidWallFilterer) ParseFixedSwapFeeThresholdUpdated(log types.Log) (*BidWallFixedSwapFeeThresholdUpdated, error) {
	event := new(BidWallFixedSwapFeeThresholdUpdated)
	if err := _BidWall.contract.UnpackLog(event, "FixedSwapFeeThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BidWallOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the BidWall contract.
type BidWallOwnershipHandoverCanceledIterator struct {
	Event *BidWallOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *BidWallOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidWallOwnershipHandoverCanceled)
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
		it.Event = new(BidWallOwnershipHandoverCanceled)
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
func (it *BidWallOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidWallOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidWallOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the BidWall contract.
type BidWallOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_BidWall *BidWallFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*BidWallOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _BidWall.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BidWallOwnershipHandoverCanceledIterator{contract: _BidWall.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_BidWall *BidWallFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *BidWallOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _BidWall.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidWallOwnershipHandoverCanceled)
				if err := _BidWall.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_BidWall *BidWallFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*BidWallOwnershipHandoverCanceled, error) {
	event := new(BidWallOwnershipHandoverCanceled)
	if err := _BidWall.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BidWallOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the BidWall contract.
type BidWallOwnershipHandoverRequestedIterator struct {
	Event *BidWallOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *BidWallOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidWallOwnershipHandoverRequested)
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
		it.Event = new(BidWallOwnershipHandoverRequested)
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
func (it *BidWallOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidWallOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidWallOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the BidWall contract.
type BidWallOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_BidWall *BidWallFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*BidWallOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _BidWall.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BidWallOwnershipHandoverRequestedIterator{contract: _BidWall.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_BidWall *BidWallFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *BidWallOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _BidWall.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidWallOwnershipHandoverRequested)
				if err := _BidWall.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_BidWall *BidWallFilterer) ParseOwnershipHandoverRequested(log types.Log) (*BidWallOwnershipHandoverRequested, error) {
	event := new(BidWallOwnershipHandoverRequested)
	if err := _BidWall.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BidWallOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BidWall contract.
type BidWallOwnershipTransferredIterator struct {
	Event *BidWallOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BidWallOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidWallOwnershipTransferred)
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
		it.Event = new(BidWallOwnershipTransferred)
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
func (it *BidWallOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidWallOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidWallOwnershipTransferred represents a OwnershipTransferred event raised by the BidWall contract.
type BidWallOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_BidWall *BidWallFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*BidWallOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BidWall.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BidWallOwnershipTransferredIterator{contract: _BidWall.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_BidWall *BidWallFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BidWallOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BidWall.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidWallOwnershipTransferred)
				if err := _BidWall.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BidWall *BidWallFilterer) ParseOwnershipTransferred(log types.Log) (*BidWallOwnershipTransferred, error) {
	event := new(BidWallOwnershipTransferred)
	if err := _BidWall.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
