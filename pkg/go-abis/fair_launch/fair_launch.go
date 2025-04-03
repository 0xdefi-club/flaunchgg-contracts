// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fair_launch

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

// FairLaunchFairLaunchInfo is an auto generated low-level Go binding around an user-defined struct.
type FairLaunchFairLaunchInfo struct {
	StartsAt    *big.Int
	EndsAt      *big.Int
	InitialTick *big.Int
	Revenue     *big.Int
	Supply      *big.Int
	Closed      bool
}

// PoolKey is an auto generated low-level Go binding around an user-defined struct.
type PoolKey struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

// FairLaunchMetaData contains all meta data concerning the FairLaunch contract.
var FairLaunchMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_poolManager\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FAIR_LAUNCH_WINDOW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"closePosition\",\"inputs\":[{\"name\":\"_poolKey\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"_tokenFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nativeIsZero\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFairLaunch.FairLaunchInfo\",\"components\":[{\"name\":\"startsAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endsAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"revenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"closed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createPosition\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_initialTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"_flaunchesAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_initialTokenFairLaunch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFairLaunch.FairLaunchInfo\",\"components\":[{\"name\":\"startsAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endsAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"revenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"closed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fairLaunchInfo\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFairLaunch.FairLaunchInfo\",\"components\":[{\"name\":\"startsAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endsAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"revenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"closed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fillFromPosition\",\"inputs\":[{\"name\":\"_poolKey\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"_amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"_nativeIsZero\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"beforeSwapDelta_\",\"type\":\"int256\",\"internalType\":\"BeforeSwapDelta\"},{\"name\":\"balanceDelta_\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"fairLaunchInfo_\",\"type\":\"tuple\",\"internalType\":\"structFairLaunch.FairLaunchInfo\",\"components\":[{\"name\":\"startsAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endsAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"revenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"closed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"inFairLaunchWindow\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyRevenue\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_revenue\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"poolManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FairLaunchCreated\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_tokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_startsAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_endsAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FairLaunchEnded\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_revenue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_supply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_endedAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotModifyLiquidityDuringFairLaunch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSellTokenDuringFairLaunch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPositionManager\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b50604051611bda380380611bda83398101604081905261002e91610043565b6001600160a01b03166080523360a052610070565b5f60208284031215610053575f5ffd5b81516001600160a01b0381168114610069575f5ffd5b9392505050565b60805160a051611afc6100de5f395f8181610100015281816101a7015281816102da01528181610599015281816107b1015281816108850152818161093b01528181610d340152610db701525f818161015401528181610c3201528181610d130152610d960152611afc5ff3fe608060405234801561000f575f5ffd5b5060043610610081575f3560e01c8063197a866b146100855780631cc8fb001461009a5780635cb0fdf4146100c25780636264c08e146100d9578063791b98bc146100fb578063d37f0cb21461012f578063dc4c90d31461014f578063e4f203f614610176578063e7ea93c714610189575b5f5ffd5b610098610093366004611621565b61019c565b005b6100ad6100a8366004611641565b610253565b60405190151581526020015b60405180910390f35b6100cb61070881565b6040519081526020016100b9565b6100ec6100e736600461173b565b6102c5565b6040516100b9939291906117b8565b6101227f000000000000000000000000000000000000000000000000000000000000000081565b6040516100b991906117d4565b61014261013d3660046117e8565b610586565b6040516100b99190611820565b6101227f000000000000000000000000000000000000000000000000000000000000000081565b610142610184366004611641565b610738565b61014261019736600461173b565b61079e565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101e55760405163041fb8cb60e31b815260040160405180910390fd5b5f811215610221576101f681611842565b5f838152602081905260408120600301805490919061021690849061185c565b9091555061024f9050565b5f81131561024f575f828152602081905260408120600301805483929061024990849061186f565b90915550505b5050565b5f81815260208181526040808320815160c0810183528154808252600183015494820194909452600280830154900b9281019290925260038101546060830152600481015460808301526005015460ff16151560a08201529042108015906102be5750806020015142105b9392505050565b5f5f6102cf6115ec565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103185760405163041fb8cb60e31b815260040160405180910390fd5b60a086205f81815260208190526040812090879003610385576040805160c0810182528254815260018301546020820152600280840154900b91810191909152600382015460608201526004820154608082015260059091015460ff16151560a0820152915061057d9050565b5f5f5f8912156103d75761039889611842565b6002808501549193506103d091900b838a6103b7578c602001516103ba565b8c515b8b6103c6578d51610a3c565b8d60200151610a3c565b905061040a565b50600280830154899161040791900b828a156103f7578c602001516103fa565b8c515b8b156103c6578d51610a3c565b91505b8260040154811115610464575f818460040154670de0b6b3a76400006104309190611882565b61043a91906118ad565b9050670de0b6b3a764000061044f8285611882565b61045991906118ad565b925083600401549150505b5f89126104a15761049c61047782610b1e565b610480906118c0565b61048984610b1e565b6001600160801b031660809190911b1790565b6104bf565b6104bf6104ad83610b1e565b6104b683610b1e565b610489906118c0565b96506104f7886104e0576104d282610b1e565b6104db906118c0565b6104e9565b6104e983610b1e565b896104ad5761048984610b1e565b955081836003015f82825461050c919061186f565b9250508190555080836004015f828254610526919061185c565b90915550506040805160c0810182528454815260018501546020820152600280860154900b91810191909152600384015460608201526004840154608082015260059093015460ff16151560a08401525090925050505b93509350939050565b61058e6115ec565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105d75760405163041fb8cb60e31b815260040160405180910390fd5b5f6105e46107088561186f565b90506040518060c001604052808581526020018281526020018660020b81526020015f81526020018481526020015f15158152505f5f8881526020019081526020015f205f820151815f0155602082015181600101556040820151816002015f6101000a81548162ffffff021916908360020b62ffffff160217905550606082015181600301556080820151816004015560a0820151816005015f6101000a81548160ff021916908315150217905550905050857f0935055505897986a1e75d206c9f6c9ff968893b8820199e8d781008b03550a88486846040516106cb939291906118e1565b60405180910390a250505f8481526020818152604091829020825160c08101845281548152600182015492810192909252600280820154900b92820192909252600382015460608201526004820154608082015260059091015460ff16151560a08201525b949350505050565b6107406115ec565b505f9081526020818152604091829020825160c08101845281548152600182015492810192909252600280820154900b92820192909252600382015460608201526004820154608082015260059091015460ff16151560a082015290565b6107a66115ec565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107ef5760405163041fb8cb60e31b815260040160405180910390fd5b5f5f5f6107fd8760a0902090565b81526020019081526020015f2090505f5f84156108d557600280840154610835915f9161082c910b60016118f7565b60020b90610b3e565b9150610842603c836118f7565b905061085687838386600301546001610bd6565b600280840154620d89b31993506108779160019161082c918391900b61191c565b90506108d0878383896108c07f00000000000000000000000000000000000000000000000000000000000000008d602001516001600160a01b0316610e0d90919063ffffffff16565b6108ca919061185c565b5f610bd6565b610986565b6002808401546108ef9160019161082c918391900b61191c565b90506108fc603c8261191c565b915061090f87838386600301545f610bd6565b600280840154610927915f9161082c910b60016118f7565b9150620d89b49050610986878383896109757f00000000000000000000000000000000000000000000000000000000000000008d5f01516001600160a01b0316610e0d90919063ffffffff16565b61097f919061185c565b6001610bd6565b42600184810182905560058501805460ff1916909117905560a088206003850154600486015460405192937f90423afe5c41b867c789a3a14a8084a4cd82970dfde321694c1ea4d64612df99936109df939291906118e1565b60405180910390a250506040805160c0810182528254815260018301546020820152600280840154900b91810191909152600382015460608201526004820154608082015260059091015460ff16151560a0820152949350505050565b5f5f610a4786610ea1565b90506001600160801b036001600160a01b03821611610aba575f610a746001600160a01b03831680611882565b9050836001600160a01b0316856001600160a01b031610610aa357610a9e600160c01b8783611159565b610ab2565b610ab28187600160c01b611159565b925050610b15565b5f610ad36001600160a01b03831680600160401b611159565b9050836001600160a01b0316856001600160a01b031610610b0257610afd600160801b8783611159565b610b11565b610b118187600160801b611159565b9250505b50949350505050565b5f6001607f1b8210610b3a57610b3a6393dafdf160e01b6111f5565b5090565b5f620d89b319600284900b1215610b5b57620d89b3199250610b71565b620d89b4600284900b1315610b7157620d89b492505b610b7c603c84611941565b60020b5f03610b8c575081610bd0565b603c610b988185611962565b610ba2919061199a565b90505f8360020b1215610bbd57610bba603c8261191c565b90505b81610bd057610bcd603c826118f7565b90505b92915050565b5f81610bfc57610bf7610be886610ea1565b610bf186610ea1565b856111fd565b610c17565b610c17610c0886610ea1565b610c1186610ea1565b85611245565b9050806001600160801b03165f03610c2f5750610e06565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635a6bcfda8860405180608001604052808a60020b81526020018960020b8152602001610c8f876001600160801b0316610b1e565b600f0b81526020015f8152506040518363ffffffff1660e01b8152600401610cb89291906119c0565b60408051808303815f875af1158015610cd3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cf79190611a53565b5090505f610d058260801d90565b600f0b1215610d7d57610d7d7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000610d5d8460801d90565b600f0b610d6990611842565b8a516001600160a01b03169291905f6112a3565b5f610d8882600f0b90565b600f0b1215610e0357610e037f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000610de084600f0b90565b600f0b610dec90611842565b60208b01516001600160a01b03169291905f6112a3565b50505b5050505050565b5f610e20836001600160a01b0316611579565b15610e3657506001600160a01b03811631610bd0565b6040516370a0823160e01b81526001600160a01b038416906370a0823190610e629085906004016117d4565b602060405180830381865afa158015610e7d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bcd9190611a75565b60020b5f60ff82901d80830118620d89e8811115610eca57610eca6345c3193d60e11b84611586565b7001fffcb933bd6fad37aa2d162d1a5940016001821602600160801b186002821615610f06576ffff97272373d413259a46990580e213a0260801c5b6004821615610f25576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615610f44576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b6010821615610f63576fffcb9843d60f6159c9db58835c9266440260801c5b6020821615610f82576fff973b41fa98c081472e6896dfb254c00260801c5b6040821615610fa1576fff2ea16466c96a3843ec78b326b528610260801c5b6080821615610fc0576ffe5dee046a99a2a811c461f1969c30530260801c5b610100821615610fe0576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b610200821615611000576ff987a7253ac413176f2b074cf7815e540260801c5b610400821615611020576ff3392b0822b70005940c7a398e4b70f30260801c5b610800821615611040576fe7159475a2c29b7443b29c7fa6e889d90260801c5b611000821615611060576fd097f3bdfd2022b8845ad8f792aa58250260801c5b612000821615611080576fa9f746462d870fdf8a65dc1f90e061e50260801c5b6140008216156110a0576f70d869a156d2a1b890bb3df62baf32f70260801c5b6180008216156110c0576f31be135f97d08fd981231505542fcfa60260801c5b620100008216156110e1576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b62020000821615611101576e5d6af8dedb81196699c329225ee6040260801c5b62040000821615611120576d2216e584f5fa1ea926041bedfe980260801c5b6208000082161561113d576b048a170391f7dc42444e8fa20260801c5b5f841315611149575f19045b63ffffffff0160201c9392505050565b5f838302815f1985870982811083820303915050808411611178575f5ffd5b805f0361118a575082900490506102be565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b805f5260045ffd5b5f826001600160a01b0316846001600160a01b0316111561121c579192915b61073061124083600160601b6112328888611a8c565b6001600160a01b0316611159565b611595565b5f826001600160a01b0316846001600160a01b03161115611264579192915b5f611286856001600160a01b0316856001600160a01b0316600160601b611159565b905061129a61124084836112328989611a8c565b95945050505050565b801561133757836001600160a01b031663f5298aca846112d2886001600160a01b03166001600160a01b031690565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604481018590526064015f604051808303815f87803b15801561131c575f5ffd5b505af115801561132e573d5f5f3e3d5ffd5b50505050610e06565b611349856001600160a01b0316611579565b156113b757836001600160a01b03166311da60b4836040518263ffffffff1660e01b815260040160206040518083038185885af115801561138c573d5f5f3e3d5ffd5b50505050506040513d601f19601f820116820180604052508101906113b19190611a75565b50610e06565b604051632961046560e21b81526001600160a01b0385169063a5841194906113e39088906004016117d4565b5f604051808303815f87803b1580156113fa575f5ffd5b505af115801561140c573d5f5f3e3d5ffd5b505050506001600160a01b038316301461149e576040516323b872dd60e01b81526001600160a01b0384811660048301528581166024830152604482018490528616906323b872dd906064016020604051808303815f875af1158015611474573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114989190611aab565b50611510565b60405163a9059cbb60e01b81526001600160a01b0385811660048301526024820184905286169063a9059cbb906044016020604051808303815f875af11580156114ea573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061150e9190611aab565b505b836001600160a01b03166311da60b46040518163ffffffff1660e01b81526004016020604051808303815f875af115801561154d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115719190611a75565b505050505050565b6001600160a01b03161590565b815f528060020b60045260245ffd5b806001600160801b03811681146115e75760405162461bcd60e51b81526020600482015260126024820152716c6971756964697479206f766572666c6f7760701b604482015260640160405180910390fd5b919050565b6040518060c001604052805f81526020015f81526020015f60020b81526020015f81526020015f81526020015f151581525090565b5f5f60408385031215611632575f5ffd5b50508035926020909101359150565b5f60208284031215611651575f5ffd5b5035919050565b80356001600160a01b03811681146115e7575f5ffd5b803562ffffff811681146115e7575f5ffd5b8035600281900b81146115e7575f5ffd5b5f60a082840312156116a1575f5ffd5b60405160a081016001600160401b03811182821017156116cf57634e487b7160e01b5f52604160045260245ffd5b6040529050806116de83611658565b81526116ec60208401611658565b60208201526116fd6040840161166e565b604082015261170e60608401611680565b606082015261171f60808401611658565b60808201525092915050565b8015158114611738575f5ffd5b50565b5f5f5f60e0848603121561174d575f5ffd5b6117578585611691565b925060a0840135915060c084013561176e8161172b565b809150509250925092565b8051825260208101516020830152604081015160020b6040830152606081015160608301526080810151608083015260a0810151151560a08301525050565b8381526020810183905261010081016107306040830184611779565b6001600160a01b0391909116815260200190565b5f5f5f5f608085870312156117fb575f5ffd5b8435935061180b60208601611680565b93969395505050506040820135916060013590565b60c08101610bd08284611779565b634e487b7160e01b5f52601160045260245ffd5b5f600160ff1b82016118565761185661182e565b505f0390565b81810381811115610bd057610bd061182e565b80820180821115610bd057610bd061182e565b8082028115828204841417610bd057610bd061182e565b634e487b7160e01b5f52601260045260245ffd5b5f826118bb576118bb611899565b500490565b5f600f82900b6001607f1b81016118d9576118d961182e565b5f0392915050565b9283526020830191909152604082015260600190565b600281810b9083900b01627fffff8113627fffff1982121715610bd057610bd061182e565b600282810b9082900b03627fffff198112627fffff82131715610bd057610bd061182e565b5f8260020b8061195357611953611899565b808360020b0791505092915050565b5f8160020b8360020b8061197857611978611899565b627fffff1982145f19821416156119915761199161182e565b90059392505050565b5f8260020b8260020b028060020b91508082146119b9576119b961182e565b5092915050565b82516001600160a01b03908116825260208085015182169083015260408085015162ffffff169083015260608085015160020b9083015260808085015190911690820152611a3860a0820183805160020b8252602081015160020b602083015260408101516040830152606081015160608301525050565b6101406101208201525f61073061014083015f815260200190565b5f5f60408385031215611a64575f5ffd5b505080516020909101519092909150565b5f60208284031215611a85575f5ffd5b5051919050565b6001600160a01b038281168282160390811115610bd057610bd061182e565b5f60208284031215611abb575f5ffd5b81516102be8161172b56fea2646970667358221220876bb60b472c63a1e8881bba7db4d3e47798616abde1b0901c0da9f7cff936d464736f6c634300081b0033",
}

// FairLaunchABI is the input ABI used to generate the binding from.
// Deprecated: Use FairLaunchMetaData.ABI instead.
var FairLaunchABI = FairLaunchMetaData.ABI

// FairLaunchBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FairLaunchMetaData.Bin instead.
var FairLaunchBin = FairLaunchMetaData.Bin

// DeployFairLaunch deploys a new Ethereum contract, binding an instance of FairLaunch to it.
func DeployFairLaunch(auth *bind.TransactOpts, backend bind.ContractBackend, _poolManager common.Address) (common.Address, *types.Transaction, *FairLaunch, error) {
	parsed, err := FairLaunchMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FairLaunchBin), backend, _poolManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FairLaunch{FairLaunchCaller: FairLaunchCaller{contract: contract}, FairLaunchTransactor: FairLaunchTransactor{contract: contract}, FairLaunchFilterer: FairLaunchFilterer{contract: contract}}, nil
}

// FairLaunch is an auto generated Go binding around an Ethereum contract.
type FairLaunch struct {
	FairLaunchCaller     // Read-only binding to the contract
	FairLaunchTransactor // Write-only binding to the contract
	FairLaunchFilterer   // Log filterer for contract events
}

// FairLaunchCaller is an auto generated read-only Go binding around an Ethereum contract.
type FairLaunchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairLaunchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FairLaunchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairLaunchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FairLaunchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairLaunchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FairLaunchSession struct {
	Contract     *FairLaunch       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FairLaunchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FairLaunchCallerSession struct {
	Contract *FairLaunchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FairLaunchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FairLaunchTransactorSession struct {
	Contract     *FairLaunchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FairLaunchRaw is an auto generated low-level Go binding around an Ethereum contract.
type FairLaunchRaw struct {
	Contract *FairLaunch // Generic contract binding to access the raw methods on
}

// FairLaunchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FairLaunchCallerRaw struct {
	Contract *FairLaunchCaller // Generic read-only contract binding to access the raw methods on
}

// FairLaunchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FairLaunchTransactorRaw struct {
	Contract *FairLaunchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFairLaunch creates a new instance of FairLaunch, bound to a specific deployed contract.
func NewFairLaunch(address common.Address, backend bind.ContractBackend) (*FairLaunch, error) {
	contract, err := bindFairLaunch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FairLaunch{FairLaunchCaller: FairLaunchCaller{contract: contract}, FairLaunchTransactor: FairLaunchTransactor{contract: contract}, FairLaunchFilterer: FairLaunchFilterer{contract: contract}}, nil
}

// NewFairLaunchCaller creates a new read-only instance of FairLaunch, bound to a specific deployed contract.
func NewFairLaunchCaller(address common.Address, caller bind.ContractCaller) (*FairLaunchCaller, error) {
	contract, err := bindFairLaunch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FairLaunchCaller{contract: contract}, nil
}

// NewFairLaunchTransactor creates a new write-only instance of FairLaunch, bound to a specific deployed contract.
func NewFairLaunchTransactor(address common.Address, transactor bind.ContractTransactor) (*FairLaunchTransactor, error) {
	contract, err := bindFairLaunch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FairLaunchTransactor{contract: contract}, nil
}

// NewFairLaunchFilterer creates a new log filterer instance of FairLaunch, bound to a specific deployed contract.
func NewFairLaunchFilterer(address common.Address, filterer bind.ContractFilterer) (*FairLaunchFilterer, error) {
	contract, err := bindFairLaunch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FairLaunchFilterer{contract: contract}, nil
}

// bindFairLaunch binds a generic wrapper to an already deployed contract.
func bindFairLaunch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FairLaunchMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FairLaunch *FairLaunchRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FairLaunch.Contract.FairLaunchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FairLaunch *FairLaunchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairLaunch.Contract.FairLaunchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FairLaunch *FairLaunchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FairLaunch.Contract.FairLaunchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FairLaunch *FairLaunchCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FairLaunch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FairLaunch *FairLaunchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairLaunch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FairLaunch *FairLaunchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FairLaunch.Contract.contract.Transact(opts, method, params...)
}

// FAIRLAUNCHWINDOW is a free data retrieval call binding the contract method 0x5cb0fdf4.
//
// Solidity: function FAIR_LAUNCH_WINDOW() view returns(uint256)
func (_FairLaunch *FairLaunchCaller) FAIRLAUNCHWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairLaunch.contract.Call(opts, &out, "FAIR_LAUNCH_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FAIRLAUNCHWINDOW is a free data retrieval call binding the contract method 0x5cb0fdf4.
//
// Solidity: function FAIR_LAUNCH_WINDOW() view returns(uint256)
func (_FairLaunch *FairLaunchSession) FAIRLAUNCHWINDOW() (*big.Int, error) {
	return _FairLaunch.Contract.FAIRLAUNCHWINDOW(&_FairLaunch.CallOpts)
}

// FAIRLAUNCHWINDOW is a free data retrieval call binding the contract method 0x5cb0fdf4.
//
// Solidity: function FAIR_LAUNCH_WINDOW() view returns(uint256)
func (_FairLaunch *FairLaunchCallerSession) FAIRLAUNCHWINDOW() (*big.Int, error) {
	return _FairLaunch.Contract.FAIRLAUNCHWINDOW(&_FairLaunch.CallOpts)
}

// FairLaunchInfo is a free data retrieval call binding the contract method 0xe4f203f6.
//
// Solidity: function fairLaunchInfo(bytes32 _poolId) view returns((uint256,uint256,int24,uint256,uint256,bool))
func (_FairLaunch *FairLaunchCaller) FairLaunchInfo(opts *bind.CallOpts, _poolId [32]byte) (FairLaunchFairLaunchInfo, error) {
	var out []interface{}
	err := _FairLaunch.contract.Call(opts, &out, "fairLaunchInfo", _poolId)

	if err != nil {
		return *new(FairLaunchFairLaunchInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(FairLaunchFairLaunchInfo)).(*FairLaunchFairLaunchInfo)

	return out0, err

}

// FairLaunchInfo is a free data retrieval call binding the contract method 0xe4f203f6.
//
// Solidity: function fairLaunchInfo(bytes32 _poolId) view returns((uint256,uint256,int24,uint256,uint256,bool))
func (_FairLaunch *FairLaunchSession) FairLaunchInfo(_poolId [32]byte) (FairLaunchFairLaunchInfo, error) {
	return _FairLaunch.Contract.FairLaunchInfo(&_FairLaunch.CallOpts, _poolId)
}

// FairLaunchInfo is a free data retrieval call binding the contract method 0xe4f203f6.
//
// Solidity: function fairLaunchInfo(bytes32 _poolId) view returns((uint256,uint256,int24,uint256,uint256,bool))
func (_FairLaunch *FairLaunchCallerSession) FairLaunchInfo(_poolId [32]byte) (FairLaunchFairLaunchInfo, error) {
	return _FairLaunch.Contract.FairLaunchInfo(&_FairLaunch.CallOpts, _poolId)
}

// InFairLaunchWindow is a free data retrieval call binding the contract method 0x1cc8fb00.
//
// Solidity: function inFairLaunchWindow(bytes32 _poolId) view returns(bool)
func (_FairLaunch *FairLaunchCaller) InFairLaunchWindow(opts *bind.CallOpts, _poolId [32]byte) (bool, error) {
	var out []interface{}
	err := _FairLaunch.contract.Call(opts, &out, "inFairLaunchWindow", _poolId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InFairLaunchWindow is a free data retrieval call binding the contract method 0x1cc8fb00.
//
// Solidity: function inFairLaunchWindow(bytes32 _poolId) view returns(bool)
func (_FairLaunch *FairLaunchSession) InFairLaunchWindow(_poolId [32]byte) (bool, error) {
	return _FairLaunch.Contract.InFairLaunchWindow(&_FairLaunch.CallOpts, _poolId)
}

// InFairLaunchWindow is a free data retrieval call binding the contract method 0x1cc8fb00.
//
// Solidity: function inFairLaunchWindow(bytes32 _poolId) view returns(bool)
func (_FairLaunch *FairLaunchCallerSession) InFairLaunchWindow(_poolId [32]byte) (bool, error) {
	return _FairLaunch.Contract.InFairLaunchWindow(&_FairLaunch.CallOpts, _poolId)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_FairLaunch *FairLaunchCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairLaunch.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_FairLaunch *FairLaunchSession) PoolManager() (common.Address, error) {
	return _FairLaunch.Contract.PoolManager(&_FairLaunch.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_FairLaunch *FairLaunchCallerSession) PoolManager() (common.Address, error) {
	return _FairLaunch.Contract.PoolManager(&_FairLaunch.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_FairLaunch *FairLaunchCaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairLaunch.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_FairLaunch *FairLaunchSession) PositionManager() (common.Address, error) {
	return _FairLaunch.Contract.PositionManager(&_FairLaunch.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_FairLaunch *FairLaunchCallerSession) PositionManager() (common.Address, error) {
	return _FairLaunch.Contract.PositionManager(&_FairLaunch.CallOpts)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xe7ea93c7.
//
// Solidity: function closePosition((address,address,uint24,int24,address) _poolKey, uint256 _tokenFees, bool _nativeIsZero) returns((uint256,uint256,int24,uint256,uint256,bool))
func (_FairLaunch *FairLaunchTransactor) ClosePosition(opts *bind.TransactOpts, _poolKey PoolKey, _tokenFees *big.Int, _nativeIsZero bool) (*types.Transaction, error) {
	return _FairLaunch.contract.Transact(opts, "closePosition", _poolKey, _tokenFees, _nativeIsZero)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xe7ea93c7.
//
// Solidity: function closePosition((address,address,uint24,int24,address) _poolKey, uint256 _tokenFees, bool _nativeIsZero) returns((uint256,uint256,int24,uint256,uint256,bool))
func (_FairLaunch *FairLaunchSession) ClosePosition(_poolKey PoolKey, _tokenFees *big.Int, _nativeIsZero bool) (*types.Transaction, error) {
	return _FairLaunch.Contract.ClosePosition(&_FairLaunch.TransactOpts, _poolKey, _tokenFees, _nativeIsZero)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xe7ea93c7.
//
// Solidity: function closePosition((address,address,uint24,int24,address) _poolKey, uint256 _tokenFees, bool _nativeIsZero) returns((uint256,uint256,int24,uint256,uint256,bool))
func (_FairLaunch *FairLaunchTransactorSession) ClosePosition(_poolKey PoolKey, _tokenFees *big.Int, _nativeIsZero bool) (*types.Transaction, error) {
	return _FairLaunch.Contract.ClosePosition(&_FairLaunch.TransactOpts, _poolKey, _tokenFees, _nativeIsZero)
}

// CreatePosition is a paid mutator transaction binding the contract method 0xd37f0cb2.
//
// Solidity: function createPosition(bytes32 _poolId, int24 _initialTick, uint256 _flaunchesAt, uint256 _initialTokenFairLaunch) returns((uint256,uint256,int24,uint256,uint256,bool))
func (_FairLaunch *FairLaunchTransactor) CreatePosition(opts *bind.TransactOpts, _poolId [32]byte, _initialTick *big.Int, _flaunchesAt *big.Int, _initialTokenFairLaunch *big.Int) (*types.Transaction, error) {
	return _FairLaunch.contract.Transact(opts, "createPosition", _poolId, _initialTick, _flaunchesAt, _initialTokenFairLaunch)
}

// CreatePosition is a paid mutator transaction binding the contract method 0xd37f0cb2.
//
// Solidity: function createPosition(bytes32 _poolId, int24 _initialTick, uint256 _flaunchesAt, uint256 _initialTokenFairLaunch) returns((uint256,uint256,int24,uint256,uint256,bool))
func (_FairLaunch *FairLaunchSession) CreatePosition(_poolId [32]byte, _initialTick *big.Int, _flaunchesAt *big.Int, _initialTokenFairLaunch *big.Int) (*types.Transaction, error) {
	return _FairLaunch.Contract.CreatePosition(&_FairLaunch.TransactOpts, _poolId, _initialTick, _flaunchesAt, _initialTokenFairLaunch)
}

// CreatePosition is a paid mutator transaction binding the contract method 0xd37f0cb2.
//
// Solidity: function createPosition(bytes32 _poolId, int24 _initialTick, uint256 _flaunchesAt, uint256 _initialTokenFairLaunch) returns((uint256,uint256,int24,uint256,uint256,bool))
func (_FairLaunch *FairLaunchTransactorSession) CreatePosition(_poolId [32]byte, _initialTick *big.Int, _flaunchesAt *big.Int, _initialTokenFairLaunch *big.Int) (*types.Transaction, error) {
	return _FairLaunch.Contract.CreatePosition(&_FairLaunch.TransactOpts, _poolId, _initialTick, _flaunchesAt, _initialTokenFairLaunch)
}

// FillFromPosition is a paid mutator transaction binding the contract method 0x6264c08e.
//
// Solidity: function fillFromPosition((address,address,uint24,int24,address) _poolKey, int256 _amountSpecified, bool _nativeIsZero) returns(int256 beforeSwapDelta_, int256 balanceDelta_, (uint256,uint256,int24,uint256,uint256,bool) fairLaunchInfo_)
func (_FairLaunch *FairLaunchTransactor) FillFromPosition(opts *bind.TransactOpts, _poolKey PoolKey, _amountSpecified *big.Int, _nativeIsZero bool) (*types.Transaction, error) {
	return _FairLaunch.contract.Transact(opts, "fillFromPosition", _poolKey, _amountSpecified, _nativeIsZero)
}

// FillFromPosition is a paid mutator transaction binding the contract method 0x6264c08e.
//
// Solidity: function fillFromPosition((address,address,uint24,int24,address) _poolKey, int256 _amountSpecified, bool _nativeIsZero) returns(int256 beforeSwapDelta_, int256 balanceDelta_, (uint256,uint256,int24,uint256,uint256,bool) fairLaunchInfo_)
func (_FairLaunch *FairLaunchSession) FillFromPosition(_poolKey PoolKey, _amountSpecified *big.Int, _nativeIsZero bool) (*types.Transaction, error) {
	return _FairLaunch.Contract.FillFromPosition(&_FairLaunch.TransactOpts, _poolKey, _amountSpecified, _nativeIsZero)
}

// FillFromPosition is a paid mutator transaction binding the contract method 0x6264c08e.
//
// Solidity: function fillFromPosition((address,address,uint24,int24,address) _poolKey, int256 _amountSpecified, bool _nativeIsZero) returns(int256 beforeSwapDelta_, int256 balanceDelta_, (uint256,uint256,int24,uint256,uint256,bool) fairLaunchInfo_)
func (_FairLaunch *FairLaunchTransactorSession) FillFromPosition(_poolKey PoolKey, _amountSpecified *big.Int, _nativeIsZero bool) (*types.Transaction, error) {
	return _FairLaunch.Contract.FillFromPosition(&_FairLaunch.TransactOpts, _poolKey, _amountSpecified, _nativeIsZero)
}

// ModifyRevenue is a paid mutator transaction binding the contract method 0x197a866b.
//
// Solidity: function modifyRevenue(bytes32 _poolId, int256 _revenue) returns()
func (_FairLaunch *FairLaunchTransactor) ModifyRevenue(opts *bind.TransactOpts, _poolId [32]byte, _revenue *big.Int) (*types.Transaction, error) {
	return _FairLaunch.contract.Transact(opts, "modifyRevenue", _poolId, _revenue)
}

// ModifyRevenue is a paid mutator transaction binding the contract method 0x197a866b.
//
// Solidity: function modifyRevenue(bytes32 _poolId, int256 _revenue) returns()
func (_FairLaunch *FairLaunchSession) ModifyRevenue(_poolId [32]byte, _revenue *big.Int) (*types.Transaction, error) {
	return _FairLaunch.Contract.ModifyRevenue(&_FairLaunch.TransactOpts, _poolId, _revenue)
}

// ModifyRevenue is a paid mutator transaction binding the contract method 0x197a866b.
//
// Solidity: function modifyRevenue(bytes32 _poolId, int256 _revenue) returns()
func (_FairLaunch *FairLaunchTransactorSession) ModifyRevenue(_poolId [32]byte, _revenue *big.Int) (*types.Transaction, error) {
	return _FairLaunch.Contract.ModifyRevenue(&_FairLaunch.TransactOpts, _poolId, _revenue)
}

// FairLaunchFairLaunchCreatedIterator is returned from FilterFairLaunchCreated and is used to iterate over the raw logs and unpacked data for FairLaunchCreated events raised by the FairLaunch contract.
type FairLaunchFairLaunchCreatedIterator struct {
	Event *FairLaunchFairLaunchCreated // Event containing the contract specifics and raw log

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
func (it *FairLaunchFairLaunchCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairLaunchFairLaunchCreated)
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
		it.Event = new(FairLaunchFairLaunchCreated)
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
func (it *FairLaunchFairLaunchCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairLaunchFairLaunchCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairLaunchFairLaunchCreated represents a FairLaunchCreated event raised by the FairLaunch contract.
type FairLaunchFairLaunchCreated struct {
	PoolId   [32]byte
	Tokens   *big.Int
	StartsAt *big.Int
	EndsAt   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFairLaunchCreated is a free log retrieval operation binding the contract event 0x0935055505897986a1e75d206c9f6c9ff968893b8820199e8d781008b03550a8.
//
// Solidity: event FairLaunchCreated(bytes32 indexed _poolId, uint256 _tokens, uint256 _startsAt, uint256 _endsAt)
func (_FairLaunch *FairLaunchFilterer) FilterFairLaunchCreated(opts *bind.FilterOpts, _poolId [][32]byte) (*FairLaunchFairLaunchCreatedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FairLaunch.contract.FilterLogs(opts, "FairLaunchCreated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &FairLaunchFairLaunchCreatedIterator{contract: _FairLaunch.contract, event: "FairLaunchCreated", logs: logs, sub: sub}, nil
}

// WatchFairLaunchCreated is a free log subscription operation binding the contract event 0x0935055505897986a1e75d206c9f6c9ff968893b8820199e8d781008b03550a8.
//
// Solidity: event FairLaunchCreated(bytes32 indexed _poolId, uint256 _tokens, uint256 _startsAt, uint256 _endsAt)
func (_FairLaunch *FairLaunchFilterer) WatchFairLaunchCreated(opts *bind.WatchOpts, sink chan<- *FairLaunchFairLaunchCreated, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FairLaunch.contract.WatchLogs(opts, "FairLaunchCreated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairLaunchFairLaunchCreated)
				if err := _FairLaunch.contract.UnpackLog(event, "FairLaunchCreated", log); err != nil {
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

// ParseFairLaunchCreated is a log parse operation binding the contract event 0x0935055505897986a1e75d206c9f6c9ff968893b8820199e8d781008b03550a8.
//
// Solidity: event FairLaunchCreated(bytes32 indexed _poolId, uint256 _tokens, uint256 _startsAt, uint256 _endsAt)
func (_FairLaunch *FairLaunchFilterer) ParseFairLaunchCreated(log types.Log) (*FairLaunchFairLaunchCreated, error) {
	event := new(FairLaunchFairLaunchCreated)
	if err := _FairLaunch.contract.UnpackLog(event, "FairLaunchCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairLaunchFairLaunchEndedIterator is returned from FilterFairLaunchEnded and is used to iterate over the raw logs and unpacked data for FairLaunchEnded events raised by the FairLaunch contract.
type FairLaunchFairLaunchEndedIterator struct {
	Event *FairLaunchFairLaunchEnded // Event containing the contract specifics and raw log

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
func (it *FairLaunchFairLaunchEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairLaunchFairLaunchEnded)
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
		it.Event = new(FairLaunchFairLaunchEnded)
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
func (it *FairLaunchFairLaunchEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairLaunchFairLaunchEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairLaunchFairLaunchEnded represents a FairLaunchEnded event raised by the FairLaunch contract.
type FairLaunchFairLaunchEnded struct {
	PoolId  [32]byte
	Revenue *big.Int
	Supply  *big.Int
	EndedAt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFairLaunchEnded is a free log retrieval operation binding the contract event 0x90423afe5c41b867c789a3a14a8084a4cd82970dfde321694c1ea4d64612df99.
//
// Solidity: event FairLaunchEnded(bytes32 indexed _poolId, uint256 _revenue, uint256 _supply, uint256 _endedAt)
func (_FairLaunch *FairLaunchFilterer) FilterFairLaunchEnded(opts *bind.FilterOpts, _poolId [][32]byte) (*FairLaunchFairLaunchEndedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FairLaunch.contract.FilterLogs(opts, "FairLaunchEnded", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &FairLaunchFairLaunchEndedIterator{contract: _FairLaunch.contract, event: "FairLaunchEnded", logs: logs, sub: sub}, nil
}

// WatchFairLaunchEnded is a free log subscription operation binding the contract event 0x90423afe5c41b867c789a3a14a8084a4cd82970dfde321694c1ea4d64612df99.
//
// Solidity: event FairLaunchEnded(bytes32 indexed _poolId, uint256 _revenue, uint256 _supply, uint256 _endedAt)
func (_FairLaunch *FairLaunchFilterer) WatchFairLaunchEnded(opts *bind.WatchOpts, sink chan<- *FairLaunchFairLaunchEnded, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _FairLaunch.contract.WatchLogs(opts, "FairLaunchEnded", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairLaunchFairLaunchEnded)
				if err := _FairLaunch.contract.UnpackLog(event, "FairLaunchEnded", log); err != nil {
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

// ParseFairLaunchEnded is a log parse operation binding the contract event 0x90423afe5c41b867c789a3a14a8084a4cd82970dfde321694c1ea4d64612df99.
//
// Solidity: event FairLaunchEnded(bytes32 indexed _poolId, uint256 _revenue, uint256 _supply, uint256 _endedAt)
func (_FairLaunch *FairLaunchFilterer) ParseFairLaunchEnded(log types.Log) (*FairLaunchFairLaunchEnded, error) {
	event := new(FairLaunchFairLaunchEnded)
	if err := _FairLaunch.contract.UnpackLog(event, "FairLaunchEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
