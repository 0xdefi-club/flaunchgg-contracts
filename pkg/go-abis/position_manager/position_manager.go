// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package position_manager

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

// FeeDistributorFeeDistribution is an auto generated low-level Go binding around an user-defined struct.
type FeeDistributorFeeDistribution struct {
	SwapFee  *big.Int
	Referrer *big.Int
	Protocol *big.Int
	Active   bool
}

// HooksPermissions is an auto generated low-level Go binding around an user-defined struct.
type HooksPermissions struct {
	BeforeInitialize                bool
	AfterInitialize                 bool
	BeforeAddLiquidity              bool
	AfterAddLiquidity               bool
	BeforeRemoveLiquidity           bool
	AfterRemoveLiquidity            bool
	BeforeSwap                      bool
	AfterSwap                       bool
	BeforeDonate                    bool
	AfterDonate                     bool
	BeforeSwapReturnDelta           bool
	AfterSwapReturnDelta            bool
	AfterAddLiquidityReturnDelta    bool
	AfterRemoveLiquidityReturnDelta bool
}

// IPoolManagerModifyLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerModifyLiquidityParams struct {
	TickLower      *big.Int
	TickUpper      *big.Int
	LiquidityDelta *big.Int
	Salt           [32]byte
}

// IPoolManagerSwapParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerSwapParams struct {
	ZeroForOne        bool
	AmountSpecified   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// InternalSwapPoolClaimableFees is an auto generated low-level Go binding around an user-defined struct.
type InternalSwapPoolClaimableFees struct {
	Amount0 *big.Int
	Amount1 *big.Int
}

// PoolKey is an auto generated low-level Go binding around an user-defined struct.
type PoolKey struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

// PositionManagerConstructorParams is an auto generated low-level Go binding around an user-defined struct.
type PositionManagerConstructorParams struct {
	NativeToken          common.Address
	PoolManager          common.Address
	FeeDistribution      FeeDistributorFeeDistribution
	InitialPrice         common.Address
	ProtocolOwner        common.Address
	ProtocolFeeRecipient common.Address
	FlayGovernance       common.Address
	FeeExemptions        common.Address
}

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

// PositionManagerMetaData contains all meta data concerning the PositionManager contract.
var PositionManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structPositionManager.ConstructorParams\",\"components\":[{\"name\":\"nativeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolManager\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"},{\"name\":\"feeDistribution\",\"type\":\"tuple\",\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"initialPrice\",\"type\":\"address\",\"internalType\":\"contractIInitialPrice\"},{\"name\":\"protocolOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"protocolFeeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"flayGovernance\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeExemptions\",\"type\":\"address\",\"internalType\":\"contractFeeExemptions\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MAX_PROTOCOL_ALLOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_DISTRIBUTE_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"actionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractTreasuryActionManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"afterAddLiquidity\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_delta\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"_feesAccrued\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"selector_\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterDonate\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"_amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"selector_\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterInitialize\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterRemoveLiquidity\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_delta\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"_feesAccrued\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"selector_\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"afterSwap\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"_params\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"_delta\",\"type\":\"int256\",\"internalType\":\"BalanceDelta\"},{\"name\":\"_hookData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"selector_\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"hookDeltaUnspecified_\",\"type\":\"int128\",\"internalType\":\"int128\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeAddLiquidity\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"selector_\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeDonate\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeInitialize\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeRemoveLiquidity\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"components\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"selector_\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeSwap\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]},{\"name\":\"_params\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.SwapParams\",\"components\":[{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]},{\"name\":\"_hookData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"selector_\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"beforeSwapDelta_\",\"type\":\"int256\",\"internalType\":\"BeforeSwapDelta\"},{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bidWall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBidWall\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"closeBidWall\",\"inputs\":[{\"name\":\"_key\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fairLaunch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractFairLaunch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fairLaunchFeeCalculator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeCalculator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeExemptions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractFeeExemptions\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplit\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"bidWall_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creator_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"protocol_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flaunch\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple\",\"internalType\":\"structPositionManager.FlaunchParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialTokenFairLaunch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"premineAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creatorFeeAllocation\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"flaunchAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeCalculatorParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"memecoin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"flaunchContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFlaunch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flaunchesAt\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"_flaunchTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flayGovernance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeCalculator\",\"inputs\":[{\"name\":\"_isFairLaunch\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFlaunchingFee\",\"inputs\":[{\"name\":\"_initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFlaunchingMarketCap\",\"inputs\":[{\"name\":\"_initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHookPermissions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHooks.Permissions\",\"components\":[{\"name\":\"beforeInitialize\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterInitialize\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeAddLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterAddLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeRemoveLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterRemoveLiquidity\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeSwap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterSwap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeDonate\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterDonate\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"beforeSwapReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterSwapReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterAddLiquidityReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"afterRemoveLiquidityReturnDelta\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getPoolFeeDistribution\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"feeDistribution_\",\"type\":\"tuple\",\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIInitialPrice\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nativeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"notifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractNotifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolFees\",\"inputs\":[{\"name\":\"_poolKey\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternalSwapPool.ClaimableFees\",\"components\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolKey\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"premineInfo\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}],\"outputs\":[{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"referralEscrow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractReferralEscrow\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setFairLaunchFeeCalculator\",\"inputs\":[{\"name\":\"_feeCalculator\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeCalculator\",\"inputs\":[{\"name\":\"_feeCalculator\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeDistribution\",\"inputs\":[{\"name\":\"_feeDistribution\",\"type\":\"tuple\",\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFlaunch\",\"inputs\":[{\"name\":\"_flaunchContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInitialPrice\",\"inputs\":[{\"name\":\"_initialPrice\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoolFeeDistribution\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"},{\"name\":\"_feeDistribution\",\"type\":\"tuple\",\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProtocolFeeDistribution\",\"inputs\":[{\"name\":\"_protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReferralEscrow\",\"inputs\":[{\"name\":\"_referralEscrow\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"unlockCallback\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_unwrap\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CreatorFeeAllocationUpdated\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_allocation\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_payee\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FairLaunchFeeCalculatorUpdated\",\"inputs\":[{\"name\":\"_feeCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeCalculatorUpdated\",\"inputs\":[{\"name\":\"_feeCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeDistributionUpdated\",\"inputs\":[{\"name\":\"_feeDistribution\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitialPriceUpdated\",\"inputs\":[{\"name\":\"_initialPrice\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolCreated\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_memecoin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_memecoinTreasury\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_currencyFlipped\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"_flaunchFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_params\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structPositionManager.FlaunchParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialTokenFairLaunch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"premineAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creatorFeeAllocation\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"flaunchAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeCalculatorParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeDistributionUpdated\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_feeDistribution\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFeeDistributor.FeeDistribution\",\"components\":[{\"name\":\"swapFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"referrer\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocol\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeesDistributed\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_donateAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_creatorAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_bidWallAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_governanceAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_protocolAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeesReceived\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeesSwapped\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"zeroForOne\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"_amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolPremine\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_premineAmount\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolScheduled\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_flaunchesAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolStateUpdated\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_sqrtPriceX96\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"_tick\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"},{\"name\":\"_protocolFee\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"},{\"name\":\"_swapFee\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"},{\"name\":\"_liquidity\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolSwap\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"flAmount0\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"flAmount1\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"flFee0\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"flFee1\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"ispAmount0\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"ispAmount1\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"ispFee0\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"ispFee1\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"uniAmount0\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"uniAmount1\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"uniFee0\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"uniFee1\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReferralEscrowUpdated\",\"inputs\":[{\"name\":\"_referralEscrow\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReferrerFeePaid\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotBidWall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotCreator\",\"inputs\":[{\"name\":\"_caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotBeInitializedDirectly\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotModifyLiquidityDuringFairLaunch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSellTokenDuringFairLaunch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HookNotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFlaunchFee\",\"inputs\":[{\"name\":\"_paid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LockFailure\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPoolManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProtocolFeeInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecipientZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferrerFeeInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SwapFeeInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenNotFlaunched\",\"inputs\":[{\"name\":\"_flaunchesAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownPool\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"internalType\":\"PoolId\"}]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161be1838038061be1883398101604081905261002e91610985565b8051604082015160808084015160c085015160208601516001600160a01b038116909352909161005d3061042f565b506001600160a01b03841660a05261007483610522565b82516004805460208601516040808801516060890151151569010000000000000000000260ff60481b1962ffffff9283166601000000000000021663ffffffff60301b1994831663010000000265ffffffffffff199096169290971691909117939093179190911693909317179055517fca7619d8d03127c99378decb6d4c002b5b8b8d60918ea28c2300a7811e7c13009061014c9085905f60808201905062ffffff835116825262ffffff602084015116602083015262ffffff604084015116604083015260608301511515606083015292915050565b60405180910390a16001600160a01b03811660c0525f61016f638b78c6d8195490565b6001600160a01b03160361018657610186826105aa565b5050506060820151600a80546001600160a01b039283166001600160a01b03199182161790915560a0840151600b805491841691831691909117905560e0840151600f8054919093169116179055508051602082015160808301516040516101ed90610857565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103905ff080158015610226573d5f5f3e3d5ffd5b50600c80546001600160a01b0319166001600160a01b0392909216919091179055602081015160405161025890610864565b6102629190610a2d565b604051809103905ff08015801561027b573d5f5f3e3d5ffd5b50600d80546001600160a01b0319166001600160a01b039290921691909117905560808101516040516102ad90610871565b6102b79190610a2d565b604051809103905ff0801580156102d0573d5f5f3e3d5ffd5b50600e80546001600160a01b0319166001600160a01b039290921691909117905560808101516040516103029061087e565b61030c9190610a2d565b604051809103905ff080158015610325573d5f5f3e3d5ffd5b50601080546001600160a01b0319166001600160a01b039283161790558151600c5460405163095ea7b360e01b81529183169263095ea7b39261037192909116905f1990600401610a41565b6020604051808303815f875af115801561038d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103b19190610a5a565b508051600d5460405163095ea7b360e01b81526001600160a01b039283169263095ea7b3926103e8929116905f1990600401610a41565b6020604051808303815f875af1158015610404573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104289190610a5a565b5050610a7a565b61051f8161051a604080516101c0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081019190915250604080516101c08101825260018082525f60208301819052928201819052606082018190526080820181905260a0820181905260c0820181905260e0820181905261010082018390526101208201819052610140820181905261016082015261018081018290526101a081019190915290565b610602565b50565b612710815f015162ffffff16111561054d5760405163520fab3d60e01b815260040160405180910390fd5b612710816020015162ffffff16111561057957604051633626bec760e01b815260040160405180910390fd5b6103e862ffffff16816040015162ffffff16111561051f5760405163f37b175b60e01b815260040160405180910390fd5b638b78c6d8198054156105c457630dc149f05f526004601cfd5b6001600160a01b03909116801560ff1b8117909155805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6106176001600160a01b038316612000610833565b81511515901515141580610647575061063b6001600160a01b038316611000610833565b15158160200151151514155b8061066e57506106626001600160a01b038316610800610833565b15158160400151151514155b8061069557506106896001600160a01b038316610400610833565b15158160600151151514155b806106bc57506106b06001600160a01b038316610200610833565b15158160800151151514155b806106e357506106d76001600160a01b038316610100610833565b15158160a00151151514155b8061070957506106fd6001600160a01b0383166080610833565b15158160c00151151514155b8061072f57506107236001600160a01b0383166040610833565b15158160e00151151514155b8061075657506107496001600160a01b0383166020610833565b1515816101000151151514155b8061077d57506107706001600160a01b0383166010610833565b1515816101200151151514155b806107a457506107976001600160a01b0383166008610833565b1515816101400151151514155b806107cb57506107be6001600160a01b0383166004610833565b1515816101600151151514155b806107f257506107e56001600160a01b0383166002610833565b1515816101800151151514155b80610819575061080c6001600160a01b0383166001610833565b1515816101a00151151514155b1561082f5761082f630732d7b560e51b83610842565b5050565b166001600160a01b0316151590565b815f526001600160a01b03811660045260245ffd5b61280b80616b8a83390190565b611bda8061939583390190565b6104d68061af6f83390190565b6109d38061b44583390190565b60405161010081016001600160401b03811182821017156108ba57634e487b7160e01b5f52604160045260245ffd5b60405290565b80516001600160a01b03811681146108d6575f5ffd5b919050565b805162ffffff811681146108d6575f5ffd5b805180151581146108d6575f5ffd5b5f6080828403121561090c575f5ffd5b604051608081016001600160401b038111828210171561093a57634e487b7160e01b5f52604160045260245ffd5b604052905080610949836108db565b8152610957602084016108db565b6020820152610968604084016108db565b6040820152610979606084016108ed565b60608201525092915050565b5f610160828403128015610997575f5ffd5b506109a061088b565b6109a9836108c0565b81526109b7602084016108c0565b60208201526109c984604085016108fc565b60408201526109da60c084016108c0565b60608201526109eb60e084016108c0565b60808201526109fd61010084016108c0565b60a0820152610a0f61012084016108c0565b60c0820152610a2161014084016108c0565b60e08201529392505050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b5f60208284031215610a6a575f5ffd5b610a73826108ed565b9392505050565b60805160a05160c051615fea610ba05f395f8181610a980152610c4901525f8181610aea01528181610e7201528181610ea801528181610ee301528181610f500152818161165001528181611779015281816117b101528181611acf01528181611eab015281816123080152818161319a015281816131e3015281816137be015281816139aa01526145fd01525f8181610a0c01528181610bcc01528181610d7a0152818161111c015281816118f201528181611e72015281816120080152818161207a015281816121c501528181612223015281816123d50152818161252a0152818161283e0152818161294701528181612c2e01528181612ca101528181612e1c01528181612e8101528181612f2501528181612f88015261305b0152615fea5ff3fe608060405260043610610272575f3560e01c806301a2cae81461027d57806309276ea4146102b25780631d0806ae146102d15780631d61a816146102f057806321d0ee70146103375780632423028c1461036f5780632569296214610390578063259982e51461039857806327e235e3146103b7578063299d532e146103f05780632a12e4881461040a57806338b1e7001461041d5780634875cbb81461043c5780634c2d94c0146104765780634ed0f0f51461049557806354d1f13d146104b457806355d1cb60146104bc578063575e24b4146105795780636c2bbe7e146105c35780636fe7e6eb14610602578063715018a61461062157806371c4ddb01461062957806375d252a5146106485780637907addf1461066757806384aa1da0146106925780638c66d04f146106b15780638da5cb5b146106d05780638f2bdf75146106e857806391dd73461461070757806394e1cf96146107335780639f063efc14610752578063a1dd2d9114610771578063a87277dd14610790578063ad49d66e146107af578063b00eb9fe146107ce578063b3b42795146107ed578063b47b2fb114610819578063b6a8b0fa1461085b578063ba3e69b714610875578063be72150514610894578063c29c945b146108b3578063c4e833ce146108dc578063d3bff717146109dc578063dc4c90d3146109fb578063dc98354e14610a2e578063ddb475d514610a4d578063dddfecf514610a87578063df81740e14610aba578063e1758bd814610ad9578063e1b4af6914610b0c578063ec876d7114610b2b578063f04e283e14610b4a578063f2fde38b14610b5d578063fee81cf414610b70578063ff3575e414610ba1575f5ffd5b3661027957005b5f5ffd5b348015610288575f5ffd5b50600f5461029c906001600160a01b031681565b6040516102a99190614bcd565b60405180910390f35b3480156102bd575f5ffd5b5060105461029c906001600160a01b031681565b3480156102dc575f5ffd5b50600a5461029c906001600160a01b031681565b3480156102fb575f5ffd5b5061032261030a366004614be1565b60126020525f90815260409020805460019091015482565b604080519283526020830191909152016102a9565b348015610342575f5ffd5b50610356610351366004614c81565b610bc0565b6040516001600160e01b031990911681526020016102a9565b34801561037a575f5ffd5b5061038e610389366004614d12565b610c3e565b005b61038e610d22565b3480156103a3575f5ffd5b506103566103b2366004614c81565b610d6e565b3480156103c2575f5ffd5b506103e26103d1366004614d2d565b60016020525f908152604090205481565b6040519081526020016102a9565b3480156103fb575f5ffd5b506103e266038d7ea4c6800081565b61029c610418366004614d48565b610ddf565b348015610428575f5ffd5b5061038e610437366004614d2d565b611516565b348015610447575f5ffd5b5061045b610456366004614d7f565b611569565b604080519384526020840192909252908201526060016102a9565b348015610481575f5ffd5b5061038e610490366004614dac565b611609565b3480156104a0575f5ffd5b5061038e6104af366004614ead565b6117ef565b61038e6118ab565b3480156104c7575f5ffd5b5061056c6104d6366004614d2d565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152506001600160a01b039081165f90815260036020908152604091829020825160a08101845281548516815260018201548086169382019390935262ffffff600160a01b84041693810193909352600160b81b909104600290810b60608401520154909116608082015290565b6040516102a99190614f1b565b348015610584575f5ffd5b50610598610593366004614f93565b6118e4565b604080516001600160e01b03199094168452602084019290925262ffffff16908201526060016102a9565b3480156105ce575f5ffd5b506105e26105dd366004614fec565b61206d565b604080516001600160e01b031990931683526020830191909152016102a9565b34801561060d575f5ffd5b5061035661061c366004615084565b6120f9565b61038e612113565b348015610634575f5ffd5b5061029c6106433660046150dd565b612126565b348015610653575f5ffd5b50600e5461029c906001600160a01b031681565b348015610672575f5ffd5b506103e2610681366004614be1565b60116020525f908152604090205481565b34801561069d575f5ffd5b5060095461029c906001600160a01b031681565b3480156106bc575f5ffd5b5061038e6106cb366004614d2d565b612165565b3480156106db575f5ffd5b50638b78c6d8195461029c565b3480156106f3575f5ffd5b5060055461029c906001600160a01b031681565b348015610712575f5ffd5b506107266107213660046150f8565b6121b8565b6040516102a99190615164565b34801561073e575f5ffd5b50600d5461029c906001600160a01b031681565b34801561075d575f5ffd5b506105e261076c366004614fec565b612216565b34801561077c575f5ffd5b5060075461029c906001600160a01b031681565b34801561079b575f5ffd5b5061038e6107aa366004614d2d565b612280565b3480156107ba575f5ffd5b5061038e6107c9366004615176565b6122d3565b3480156107d9575f5ffd5b5060065461029c906001600160a01b031681565b3480156107f8575f5ffd5b5061080c610807366004614be1565b612480565b6040516102a9919061520d565b348015610824575f5ffd5b5061083861083336600461524d565b61251d565b604080516001600160e01b03199093168352600f9190910b6020830152016102a9565b348015610866575f5ffd5b5061035661061c3660046152d7565b348015610880575f5ffd5b50600c5461029c906001600160a01b031681565b34801561089f575f5ffd5b5061038e6108ae366004614d2d565b612776565b3480156108be575f5ffd5b506108c86103e881565b60405162ffffff90911681526020016102a9565b3480156108e7575f5ffd5b506109cf604080516101c0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081019190915250604080516101c08101825260018082525f60208301819052928201819052606082018190526080820181905260a0820181905260c0820181905260e0820181905261010082018390526101208201819052610140820181905261016082015261018081018290526101a081019190915290565b6040516102a99190615330565b3480156109e7575f5ffd5b5061038e6109f6366004615451565b6127a0565b348015610a06575f5ffd5b5061029c7f000000000000000000000000000000000000000000000000000000000000000081565b348015610a39575f5ffd5b50610356610a4836600461546b565b612832565b348015610a58575f5ffd5b50610a6c610a67366004615176565b612895565b604080518251815260209283015192810192909252016102a9565b348015610a92575f5ffd5b5061029c7f000000000000000000000000000000000000000000000000000000000000000081565b348015610ac5575f5ffd5b5061038e610ad4366004614d2d565b6128e8565b348015610ae4575f5ffd5b5061029c7f000000000000000000000000000000000000000000000000000000000000000081565b348015610b17575f5ffd5b50610356610b263660046152d7565b61293b565b348015610b36575f5ffd5b506103e2610b453660046150f8565b6129c3565b61038e610b58366004614d2d565b612a36565b61038e610b6b366004614d2d565b612a73565b348015610b7b575f5ffd5b506103e2610b8a366004614d2d565b63389a75e1600c9081525f91909152602090205490565b348015610bac575f5ffd5b506103e2610bbb3660046150f8565b612a99565b5f336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c0a5760405163570c108560e11b815260040160405180910390fd5b610c2c610c26610c1f36889003880188615176565b60a0902090565b87612acb565b5063021d0ee760e41b95945050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c86576040516282b42960e81b815260040160405180910390fd5b6103e862ffffff82161115610cae5760405163f37b175b60e01b815260040160405180910390fd5b6004805462ffffff8316600160301b0262ffffff60301b199091161781556040515f516020615e755f395f51905f5291610d1791905462ffffff8082168352601882901c81166020840152603082901c16604083015260481c60ff161515606082015260800190565b60405180910390a150565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b5f336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610db85760405163570c108560e11b815260040160405180910390fd5b610dcd610c26610c1f36889003880188615176565b5063259982e560e01b95945050505050565b6009546040516305425c9160e31b81525f91829182916001600160a01b031690632a12e48890610e1390879060040161561c565b6060604051808303815f875af1158015610e2f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e53919061562e565b6040805160a08101909152929550935091506001600160a01b038085167f0000000000000000000000000000000000000000000000000000000000000000909116108015915f918190610ea65786610ec8565b7f00000000000000000000000000000000000000000000000000000000000000005b6001600160a01b0316815260200183610ee15786610f03565b7f00000000000000000000000000000000000000000000000000000000000000005b6001600160a01b0390811682525f6020830152603c604080840191909152306060909301839052600e549051632095f47360e21b815293945086821693638257d1cc93610f7a93909216907f000000000000000000000000000000000000000000000000000000000000000090879060040161566f565b5f604051808303815f87803b158015610f91575f5ffd5b505af1158015610fa3573d5f5f3e3d5ffd5b505050506001600160a01b038581165f90815260036020908152604091829020845181546001600160a01b0319908116918616919091178255918501516001820180549487015160608801519287166001600160b81b031990961695909517600160a01b62ffffff968716021762ffffff60b81b1916600160b81b9590921694909402179092556080840151600290920180549091169190921617905560a0812061105460e0880160c08901614d12565b62ffffff16156110925761106e60e0880160c08901614d12565b5f828152602081905260409020805462ffffff191662ffffff929092169190911790555b5f61109d6001612126565b90506001600160a01b03811615611118576001600160a01b03811663edf2ec79836110cc6101208c018c61569d565b6040518463ffffffff1660e01b81526004016110ea939291906156df565b5f604051808303815f87803b158015611101575f5ffd5b505af1158015611113573d5f5f3e3d5ffd5b505050505b505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636276cbbe84600a5f9054906101000a90046001600160a01b03166001600160a01b031663fc87f3f633898e806101000190611180919061569d565b6040518563ffffffff1660e01b815260040161119f94939291906156f8565b602060405180830381865afa1580156111ba573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111de9190615724565b6040518363ffffffff1660e01b81526004016111fb92919061573f565b6020604051808303815f875af1158015611217573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061123b9190615765565b90505f61124f610b456101008b018b61569d565b9050827fe9a023154a0e1bd430ba68aafea07b09c78a0e5406c3696fb3c0dd631fa53b6489888a89868f60405161128b96959493929190615780565b60405180910390a26080890135156112d85760405180604001604052806112b58b60800135612b88565b8152436020918201525f8581526012825260409020825181559101516001909101555b6060890135156113fd57600d546001600160a01b031663d37f0cb284844260e08e013511611306574261130c565b8c60e001355b6040516001600160e01b031960e086901b168152600481019390935260029190910b6024830152604482015260608c0135606482015260840160c0604051808303815f875af1158015611361573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113859190615857565b50600d5460405163095ea7b360e01b81526001600160a01b038a81169263095ea7b3926113bb92909116905f1990600401615871565b6020604051808303815f875af11580156113d7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113fb919061588a565b505b428960e001351115611460575f838152601160205260409081902060e08b013590819055905184917f0a5b59cacc8ab09e32bf89cb0477200adacf9c0237786bf21d17f44cec9981079161145391815260200190565b60405180910390a2611471565b5f8381526011602052604090204290555b80156114bc57803410156114a6576040516384eca25560e01b8152346004820152602481018290526044015b60405180910390fd5b600b546114bc906001600160a01b031682612ba6565b803411156114d7576114d7336114d283346158b9565b612ba6565b61150a83636fe7e6eb60e01b898c6040516020016114f69291906158cc565b604051602081830303815290604052612bc3565b50505050505050919050565b61151e612d17565b600780546001600160a01b0319166001600160a01b0383161790556040517f87043577396d39ef835a9eb69bb496c219cc61bdd6e718447add3c06b6cc084490610d17908390614bcd565b5f5f5f5f61157686612480565b9050806040015162ffffff165f146115b957612710816040015162ffffff16866115a091906158e4565b6115aa919061590f565b91506115b682866158b9565b94505b5f8681526020819052604090205462ffffff1680156115fd576127106115e462ffffff8316886158e4565b6115ee919061590f565b93506115fa84876158b9565b95505b85945050509250925092565b335f908152600160205260408120549081900361162557505050565b335f90815260016020526040812055811561177457604051632e1a7d4d60e01b8152600481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d906024015f604051808303815f87803b158015611699575f5ffd5b505af11580156116ab573d5f5f3e3d5ffd5b505050505f836001600160a01b0316826040515f6040518083038185875af1925050503d805f81146116f8576040519150601f19603f3d011682016040523d82523d5f602084013e6116fd565b606091505b50509050806117445760405162461bcd60e51b815260206004820152601360248201527211551208151c985b9cd9995c8811985a5b1959606a1b604482015260640161149d565b5f516020615f355f395f51905f5233855f85604051611766949392919061592e565b60405180910390a150505050565b61179f7f00000000000000000000000000000000000000000000000000000000000000008483612d31565b5f516020615f355f395f51905f5233847f0000000000000000000000000000000000000000000000000000000000000000846040516117e1949392919061592e565b60405180910390a15b505050565b6117f7612d17565b61180081612d7b565b5f8281526002602090815260409182902083518154928501518486015160608701511515600160481b0260ff60481b1962ffffff928316600160301b021663ffffffff60301b1993831663010000000265ffffffffffff1990971692909416919091179490941716179190911790555182907fe1b2af81a774e1ebee3ca7c94c1e1a0df1210b236149a2079b8bba8dbc475a289061189f90849061520d565b60405180910390a25050565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b5f8080336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146119305760405163570c108560e11b815260040160405180910390fd5b5f611943610c1f368a90038a018a615176565b5f8181526011602052604090205490915080156119fa575f82815260126020526040902060018101544314801561197e5750805460208a0151145b156119c557805460405190815283907fdfa452364a13ecc87d8b629926a27f0b82206c68e4126e138157c5d853c71d849060200160405180910390a25f60018201556119f8565b428211156119e95760405163129b67b160e21b81526004810183905260240161149d565b5f838152601160205260408120555b505b5050600d545f906001600160a01b031663e4f203f6611a21610c1f368c90038c018c615176565b6040518263ffffffff1660e01b8152600401611a3f91815260200190565b60c060405180830381865afa158015611a5a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a7e9190615857565b90508060a00151611e6b575f611a9c610c1f368b90038b018b615176565b90505f611aac60208b018b614d2d565b600d54604051621cc8fb60e81b8152600481018590526001600160a01b039283167f00000000000000000000000000000000000000000000000000000000000000008416149350911690631cc8fb0090602401602060405180830381865afa158015611b1a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b3e919061588a565b611bcd57600d545f838152600860205260409081902060010154905163e7ea93c760e01b81526001600160a01b039092169163e7ea93c791611b87918e919086906004016159dc565b60c0604051808303815f875af1158015611ba3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bc79190615857565b50611e68565b8851151581151514611bf25760405163013fa16b60e41b815260040160405180910390fd5b600d5460208a0151604051633132604760e11b81525f926001600160a01b031691636264c08e91611c29918f9187906004016159dc565b610100604051808303815f875af1158015611c46573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c6a9190615a00565b9197509094509050611c8a611c84368d90038d018d615176565b82612e03565b5f611ca28c8c8f611c9b8b600f0b90565b8e8e613017565b9050611cca8b5f516020615f555f395f51905f525f516020615f755f395f51905f528a61323b565b611cf08b5f516020615e955f395f51905f525f516020615df55f395f51905f52846132ba565b611d2d611cfd8860801d90565b611d06836132f2565b611d108a600f0b90565b611d1a9190615a35565b6001600160801b031660809190911b1790565b96505f8b6020015112158015611d4257508015155b15611dc057600d546001600160a01b031663197a866b85611d62846132f2565b611d6b90615a62565b6040516001600160e01b031960e085901b1681526004810192909252600f0b60248201526044015f604051808303815f87803b158015611da9575f5ffd5b505af1158015611dbb573d5f5f3e3d5ffd5b505050505b84608001515f03611e6557600d5f9054906101000a90046001600160a01b03166001600160a01b031663e7ea93c78d60085f8881526020019081526020015f2060010154866040518463ffffffff1660e01b8152600401611e23939291906159dc565b60c0604051808303815f875af1158015611e3f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e639190615857565b505b50505b50505b5f80611ed97f00000000000000000000000000000000000000000000000000000000000000008b8b611ea06020830183614d2d565b6001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614613312565b9092509050611ee88183615a83565b15611fe9575f5f8a602001511215611f1d57611f18611f06846132f2565b611f0f846132f2565b611d1a90615a62565b611f3b565b611f3b611f29836132f2565b611f3290615a62565b611d1a856132f2565b90505f611f4e8c8c8f611c9b86600f0b90565b9050611f768b5f516020615f955f395f51905f525f516020615ef55f395f51905f528561323b565b611f9c8b5f516020615e155f395f51905f525f516020615e355f395f51905f52846132ba565b611fe4611fa98360801d90565b611fb38960801d90565b611fbd9190615a35565b611fc6836132f2565b611fd085600f0b90565b611fda8b600f0b90565b611d109190615a35565b965050505b61202e611ffe610c1f368d90038d018d615176565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906135e0565b5050600a805462ffffff909216600160a01b0262ffffff60a01b19909216919091179055506315d7892d60e21b9b949a50929850929650505050505050565b5f80336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146120b85760405163570c108560e11b815260040160405180910390fd5b633615df3f60e11b91506120ed6120d7610c1f368b90038b018b615176565b838b89896040516020016114f693929190615a96565b97509795505050505050565b5f604051630a85dc2960e01b815260040160405180910390fd5b61211b612d17565b6121245f613692565b565b5f81801561213e57506007546001600160a01b031615155b156121545750506007546001600160a01b031690565b50506006546001600160a01b031690565b61216d612d17565b600680546001600160a01b0319166001600160a01b0383161790556040517f3e762c7e655633ce63121393b9694f9ca1883d14d18f48f1be55e5dc7a9fb6c190610d17908390614bcd565b6060336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146122035760405163570c108560e11b815260040160405180910390fd5b61220d83836136d8565b90505b92915050565b5f80336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146122615760405163570c108560e11b815260040160405180910390fd5b6327c18fbf60e21b91506120ed6120d7610c1f368b90038b018b615176565b612288612d17565b600580546001600160a01b0319166001600160a01b0383161790556040517f9c922de256a07b4d188faacda5c1abb2cae12f74f4370d5c2f11efb37a742d7090610d17908390614bcd565b600c546001600160a01b031633146122fe57604051630632844f60e31b815260040160405180910390fd5b5f60038161232c847f0000000000000000000000000000000000000000000000000000000000000000613747565b6001600160a01b03908116825260208083019390935260409182015f20825160a08101845281548316815260018201548084169582019590955262ffffff600160a01b86041693810193909352600160b81b909304600290810b6060840152909201549091166080820181905290915015806123ae575060a0822060a0822014155b156123d35760a0822060405163180b855560e01b815260040161149d91815260200190565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166348c89491836040516020016124139190614f1b565b6040516020818303038152906040526040518263ffffffff1660e01b815260040161243e9190615164565b5f604051808303815f875af1158015612459573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526117ea9190810190615ab7565b604080516080810182525f8082526020820181905291810182905260608101919091525f82815260026020526040902054600160481b900460ff166124c65760046124d4565b5f8281526002602052604090205b60408051608081018252915462ffffff80821684526301000000820481166020850152600160301b82041691830191909152600160481b900460ff161515606082015292915050565b5f80336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146125685760405163570c108560e11b815260040160405180910390fd5b5f5f6125748760801d90565b61257e88600f0b90565b90925090505f61259160208a018a6150dd565b15155f60208b013512146125a557826125a7565b815b90505f6125c68b6125bd368d90038d018d615b49565b8e858c8c613017565b9050835f516020615e555f395f51905f525d825f516020615eb55f395f51905f525d61261c6125fa368c90038c018c615b49565b5f516020615f155f395f51905f525f516020615ed55f395f51905f52846132ba565b61263361262e368d90038d018d615176565b61376d565b5f612646610c1f368e90038e018e615176565b600d54604051621cc8fb60e81b8152600481018390529192505f916126b8916001600160a01b031690631cc8fb00906024015b602060405180830381865afa158015612694573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610643919061588a565b90506001600160a01b0381161561272d57806001600160a01b03166324f961778f8f8f8f8f8f6040518763ffffffff1660e01b81526004016126ff96959493929190615ba0565b5f604051808303815f87803b158015612716575f5ffd5b505af1158015612728573d5f5f3e3d5ffd5b505050505b50612737826132f2565b63b47b2fb160e01b9750955061274c81613b39565b61276681888f8e8e6040516020016114f693929190615be6565b5050505050965096945050505050565b61277e612d17565b600980546001600160a01b0319166001600160a01b0392909216919091179055565b6127a8612d17565b6127b181612d7b565b805160048054602084015160408086015160608701511515600160481b0260ff60481b1962ffffff928316600160301b021663ffffffff60301b1994831663010000000265ffffffffffff199096169290971691909117939093179190911693909317179055515f516020615e755f395f51905f5290610d1790839061520d565b5f336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461287c5760405163570c108560e11b815260040160405180910390fd5b604051635116b0eb60e11b815260040160405180910390fd5b604080518082019091525f808252602082015260085f6128b68460a0902090565b81526020019081526020015f206040518060400160405290815f82015481526020016001820154815250509050919050565b6128f0612d17565b600a80546001600160a01b0319166001600160a01b0383161790556040517fbfad4f58ed4ffa32358e2ae423bf8a56d8adf778ad13eb9160c0fecd99ec157790610d17908390614bcd565b5f336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146129855760405163570c108560e11b815260040160405180910390fd5b5063e1b4af6960e01b6129b96129a3610c1f36899003890189615176565b828988886040516020016114f693929190615a96565b9695505050505050565b600a5460405163010b630b60e61b81525f916001600160a01b0316906342d8c2c0906129f790339087908790600401615c11565b602060405180830381865afa158015612a12573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061220d9190615c35565b612a3e612d17565b63389a75e1600c52805f526020600c208054421115612a6457636f5e88185f526004601cfd5b5f9055612a7081613692565b50565b612a7b612d17565b8060601b612a9057637448fbae5f526004601cfd5b612a7081613692565b600a54604051630550389f60e31b81525f916001600160a01b031690632a81c4f8906129f79086908690600401615c4c565b600c546001600160a01b0382811691161480612af45750600d546001600160a01b038281169116145b15612afd575050565b600d54604051621cc8fb60e81b8152600481018490526001600160a01b0390911690631cc8fb0090602401602060405180830381865afa158015612b43573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b67919061588a565b612b6f575050565b604051631d3e310160e21b815260040160405180910390fd5b805f811215612ba157612ba16393dafdf160e01b613d20565b919050565b5f385f3884865af1612bbf5763b12d13eb5f526004601cfd5b5050565b60105460405163314e79ad60e01b81526001600160a01b039091169063314e79ad90612bf790869086908690600401615c5f565b5f604051808303815f87803b158015612c0e575f5ffd5b505af1158015612c20573d5f5f3e3d5ffd5b505050505f5f5f5f612c64877f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166135e090919063ffffffff16565b92965090945092509050867f1a111a34a945a6d821c9dc87031a478ad3107acb9b19f2ee72d3aaa72b0849c985858585612cc76001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001687613d28565b604080516001600160a01b0396909616865260029490940b602086015262ffffff92831685850152911660608401526001600160801b03166080830152519081900360a00190a250505050505050565b638b78c6d819543314612124576382b429005f526004601cfd5b816014528060345263a9059cbb60601b5f5260205f604460105f875af18060015f511416612d7157803d853b151710612d71576390b8ec185f526004601cfd5b505f603452505050565b612710815f015162ffffff161115612da65760405163520fab3d60e01b815260040160405180910390fd5b612710816020015162ffffff161115612dd257604051633626bec760e01b815260040160405180910390fd5b6103e862ffffff16816040015162ffffff161115612a705760405163f37b175b60e01b815260040160405180910390fd5b5f612e0e8260801d90565b600f0b1215612e6b57612e667f000000000000000000000000000000000000000000000000000000000000000030612e468460801d90565b600f0b612e5290615c86565b85516001600160a01b03169291905f613db4565b612f0c565b5f612e768260801d90565b600f0b1315612f0c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630b0d9c09835f015130612ebd8560801d90565b600f0b6040518463ffffffff1660e01b8152600401612ede93929190615ca0565b5f604051808303815f87803b158015612ef5575f5ffd5b505af1158015612f07573d5f5f3e3d5ffd5b505050505b5f612f1782600f0b90565b600f0b1215612f7257612bbf7f000000000000000000000000000000000000000000000000000000000000000030612f4f84600f0b90565b600f0b612f5b90615c86565b60208601516001600160a01b03169291905f613db4565b5f612f7d82600f0b90565b600f0b1315612bbf577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630b0d9c09836020015130612fc585600f0b90565b600f0b6040518463ffffffff1660e01b8152600401612fe693929190615ca0565b5f604051808303815f87803b158015612ffd575f5ffd5b505af115801561300f573d5f5f3e3d5ffd5b505050505050565b5f5f865f015115155f88602001511215151461303f5761303a6020890189614d2d565b61304f565b61304f6040890160208a01614d2d565b600d54909150613167907f0000000000000000000000000000000000000000000000000000000000000000908a908a906130c1906001600160a01b0316631cc8fb006130a3610c1f36879003870187615176565b6040518263ffffffff1660e01b815260040161267991815260200190565b855f8b600f0b126130d2578a6130db565b6130db8b615a62565b6001600160801b0316600f5f9054906101000a90046001600160a01b03166001600160a01b031663f6a23d988e6040518263ffffffff1660e01b81526004016131249190614bcd565b6040805180830381865afa15801561313e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131629190615cc4565b614064565b9150815f0361317657506129b9565b5f61318489838588886141c2565b905061322f613198368b90038b018b615176565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316846001600160a01b0316146131d7575f6131e1565b6131e183866158b9565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316856001600160a01b0316146132295761322484876158b9565b61434c565b5f61434c565b50509695505050505050565b5f5f855f015115155f87602001511215151461327c5761325b83600f0b90565b61326490615a62565b61326e8460801d90565b61327790615a62565b6132a2565b6132868360801d90565b61328f90615a62565b61329984600f0b90565b6132a290615a62565b600f0b9150600f0b915081855d80845d505050505050565b5f6132c482615c86565b8551602087015191925015155f909112036132e4575f845d80835d6132eb565b80845d5f835d5b5050505050565b5f6001607f1b821061330e5761330e6393dafdf160e01b613d20565b5090565b5f8080613327610c1f36889003880188615176565b5f818152600860205260408120600181015492935091900361334a5750506135d7565b855115158515151461335d5750506135d7565b5f6133716001600160a01b038a16846135e0565b50505090505f8760200151126133dd575f826001015488602001511161339b5787602001516133a1565b82600101545b90506133cf8289604001516133c8878e6001600160a01b0316613d2890919063ffffffff16565b845f6143d8565b50909750955061347e915050565b61344481885f01516133fe576133f96401000276a36001615d22565b61341d565b61341d600173fffd8963efd1fc6a506488495d951d5263988d26615d41565b6134306001600160a01b038d1687613d28565b8a6020015161343e90615c86565b5f6143d8565b506001850154909750909550851115905061347e578385836001015461346a91906158e4565b613474919061590f565b9450816001015493505b8415801561348a575083155b15613497575050506135d7565b84825f015f8282546134a99190615a83565b9250508190555083826001015f8282546134c391906158b9565b90915550506001600160a01b038916630b0d9c0987156134ef576134ea60208b018b614d2d565b6134ff565b6134ff60408b0160208c01614d2d565b30886040518463ffffffff1660e01b815260040161351f93929190615ca0565b5f604051808303815f87803b158015613536575f5ffd5b505af1158015613548573d5f5f3e3d5ffd5b5050505061358d8930865f8a1561356e5761356960408e0160208f01614d2d565b61357b565b61357b60208e018e614d2d565b6001600160a01b031693929190613db4565b865160408051911515825260208201879052810185905283907fce97caf4fd0295de9544b52b4b9e79fe34c370bebb6fb71bc5baae9a702079689060600160405180910390a25050505b94509492505050565b5f5f5f5f5f6135ee8661453b565b604051631e2eaeaf60e01b8152600481018290529091505f906001600160a01b03891690631e2eaeaf90602401602060405180830381865afa158015613636573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061365a9190615c35565b90506001600160a01b03811695508060a01c60020b945062ffffff8160b81c16935062ffffff8160d01c169250505092959194509250565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3811560ff1b8217905550565b600c546060906001600160a01b031663ad49d66e6136f884860186615176565b6040518263ffffffff1660e01b81526004016137149190614f1b565b5f604051808303815f87803b15801561372b575f5ffd5b505af115801561373d573d5f5f3e3d5ffd5b5050505092915050565b81515f906001600160a01b0380841691161461376457825161220d565b50506020015190565b60a081205f8181526008602052604090205466038d7ea4c6800081101561379357505050565b5f82815260086020526040812081905580806137af8585611569565b919450925090505f806137e2887f0000000000000000000000000000000000000000000000000000000000000000613747565b90505f816001600160a01b03166302d05d3f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613821573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138459190615724565b90506001600160a01b03811615851561387d578061386d57613868898388614577565b61387d565b6138778688615a83565b96505f95505b8615613a3457600c54604051634e944d5760e01b8152600481018b90526001600160a01b0390911690634e944d5790602401602060405180830381865afa1580156138ca573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138ee919061588a565b80156139615750600d54604051621cc8fb60e81b8152600481018b90526001600160a01b0390911690631cc8fb0090602401602060405180830381865afa15801561393b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061395f919061588a565b155b15613a2457600c5f9054906101000a90046001600160a01b03166001600160a01b031663088a80c78b89600a60149054906101000a900460020b8e5f01516001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316146040518563ffffffff1660e01b81526004016139f29493929190615d60565b5f604051808303815f87803b158015613a09575f5ffd5b505af1158015613a1b573d5f5f3e3d5ffd5b50505050613a34565b613a2e8785615a83565b93505f96505b8315613abe5780613aae57613aa989846001600160a01b03166361d027b36040518163ffffffff1660e01b8152600401602060405180830381865afa158015613a7f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613aa39190615724565b86614577565b613abe565b613ab88486615a83565b94505f93505b8415613adc57600b54613adc908a906001600160a01b031687614577565b6040805189815260208101889052908101889052606081018590526080810186905289907fc7241a69d3660bdfe5f36bdcca3b2da1fe8844366e46adb58be95171ab0665ad9060a00160405180910390a250505050505050505050565b604080515f516020615f555f395f51905f525c81525f516020615f755f395f51905f525c60208201525f516020615e955f395f51905f525c818301525f516020615df55f395f51905f525c60608201525f516020615f955f395f51905f525c60808201525f516020615ef55f395f51905f525c60a08201525f516020615e155f395f51905f525c60c08201525f516020615e355f395f51905f525c60e08201525f516020615e555f395f51905f525c6101008201525f516020615eb55f395f51905f525c6101208201525f516020615f155f395f51905f525c6101408201525f516020615ed55f395f51905f525c610160820152905182917f35001030bac7516ed5b1daf42bbfd3eaf587acc70c90f7f00ee42d5a34d7523391908190036101800190a25f5f516020615f555f395f51905f525d5f5f516020615f755f395f51905f525d5f5f516020615e955f395f51905f525d5f5f516020615df55f395f51905f525d5f5f516020615f955f395f51905f525d5f5f516020615ef55f395f51905f525d5f5f516020615e155f395f51905f525d5f5f516020615e355f395f51905f525d5f5f516020615e555f395f51905f525d5f5f516020615eb55f395f51905f525d5f5f516020615f155f395f51905f525d5f5f516020615ed55f395f51905f525d50565b805f5260045ffd5b5f5f613d338361453b565b90505f613d41600383615a83565b604051631e2eaeaf60e01b8152600481018290529091506001600160a01b03861690631e2eaeaf90602401602060405180830381865afa158015613d87573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613dab9190615c35565b95945050505050565b8015613e3457836001600160a01b031663f5298aca84613de3886001600160a01b03166001600160a01b031690565b856040518463ffffffff1660e01b8152600401613e0293929190615a96565b5f604051808303815f87803b158015613e19575f5ffd5b505af1158015613e2b573d5f5f3e3d5ffd5b505050506132eb565b613e46856001600160a01b0316614639565b15613eb457836001600160a01b03166311da60b4836040518263ffffffff1660e01b815260040160206040518083038185885af1158015613e89573d5f5f3e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190613eae9190615c35565b506132eb565b604051632961046560e21b81526001600160a01b0385169063a584119490613ee0908890600401614bcd565b5f604051808303815f87803b158015613ef7575f5ffd5b505af1158015613f09573d5f5f3e3d5ffd5b505050506001600160a01b0383163014613f93576040516323b872dd60e01b81526001600160a01b038616906323b872dd90613f4d90869088908790600401615ca0565b6020604051808303815f875af1158015613f69573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f8d919061588a565b50614003565b60405163a9059cbb60e01b81526001600160a01b0386169063a9059cbb90613fc19087908690600401615871565b6020604051808303815f875af1158015613fdd573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614001919061588a565b505b836001600160a01b03166311da60b46040518163ffffffff1660e01b81526004016020604051808303815f875af1158015614040573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061300f9190615c35565b5f82156141b7575f614081610807610c1f368b90038b018b615176565b5190506001600160a01b0386161561410557604051634039108560e01b81526001600160a01b038716906340391085906140c3908b908b908690600401615d8f565b602060405180830381865afa1580156140de573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906141029190615dd9565b90505b8260200151801561411e5750825162ffffff8083169116105b15614127575081515b8062ffffff165f0361413957506141b7565b61271061414b62ffffff8316866158e4565b614155919061590f565b604051630b0d9c0960e01b81529092506001600160a01b038a1690630b0d9c099061418890889030908790600401615ca0565b5f604051808303815f87803b15801561419f575f5ffd5b505af11580156141b1573d5f5f3e3d5ffd5b50505050505b979650505050505050565b5f8115613dab575f6141dc610c1f36899003890189615176565b90505f6141e882612480565b6020015190508062ffffff165f03614201575050613dab565b5f61420e85870187614d2d565b90506001600160a01b03811661422657505050613dab565b60045461271090614243906301000000900462ffffff16896158e4565b61424d919061590f565b6005549094506001600160a01b03166142b5576142746001600160a01b0389168286614646565b827f3ca4a7850462c23d5854e8b3e852626beb21b37354c670c0135ab46f0c4bdc31828a876040516142a893929190615ca0565b60405180910390a2614340565b6005546142cf906001600160a01b038a8116911686614646565b60055460405163587a617560e11b8152600481018590526001600160a01b0383811660248301528a81166044830152606482018790529091169063b0f4c2ea906084015f604051808303815f87803b158015614329575f5ffd5b505af115801561433b573d5f5f3e3d5ffd5b505050505b50505095945050505050565b60a083205f818152600860205260408120805485929061436d908490615a83565b90915550505f8181526008602052604081206001018054849290614392908490615a83565b9091555050604080518481526020810184905282917fa245a9a38e8877add82f0a82c13e062ab3df16a26121977ddcca8827d46c690a910160405180910390a250505050565b5f80808062ffffff85166001600160a01b03808a16908b16101582881280156144ae575f6144118a5f0385620f424003620f42406146f4565b90508261442a576144258d8d8d6001614791565b614437565b6144378c8e8d60016147dd565b965086811061446b578b9750620f424084146144625761445d878586620f4240036148a2565b614464565b865b9450614484565b80965061447a8d8c83866148d2565b9750868a5f030394505b8261449a576144958d898d5f6147dd565b6144a6565b6144a6888e8d5f614791565b95505061452c565b816144c4576144bf8c8c8c5f6147dd565b6144d0565b6144d08b8d8c5f614791565b94508489106144e1578a96506144f3565b8894506144f08c8b8785614920565b96505b8161450a576145058c888c6001614791565b614517565b614517878d8c60016147dd565b9550614529868485620f4240036148a2565b93505b50505095509550955095915050565b6040515f9061455a908390600690602001918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b805f0361458357505050565b6001600160a01b0382166145aa576040516384bc540160e01b815260040160405180910390fd5b6001600160a01b0382165f90815260016020526040812080548392906145d1908490615a83565b92505081905550827fc95ddcaddf83340b68d0d44c01b1703f5d28d0611a3fd87e69d79ba7e2ac21d3837f00000000000000000000000000000000000000000000000000000000000000008460405161462c93929190615ca0565b60405180910390a2505050565b6001600160a01b03161590565b5f614659846001600160a01b0316614639565b15614684575f5f5f5f85875af190508061467f5761467f835f633d2cec6f60e21b614967565b6146ee565b60405163a9059cbb60e01b81526001600160a01b038416600482015282602482015260205f6044835f895af13d15601f3d1160015f511416171691505f81525f60208201525f604082015250806146ee576146ee8463a9059cbb60e01b633c9fd93960e21b614967565b50505050565b5f838302815f1985870982811083820303915050808411614713575f5ffd5b805f036147255750829004905061478a565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b5f6001600160a01b038481169086160360ff81901d90810118600160601b6001600160801b0385166147c48184846146f4565b9350845f83858409111684019350505050949350505050565b5f836001600160a01b0316856001600160a01b031611156147fc579293925b6001600160a01b0385166148165762bfc9215f526004601cfd5b600160601b600160e01b03606084901b166001600160a01b03868603168361486957866001600160a01b03166148568383896001600160a01b03166146f4565b81614863576148636158fb565b04614895565b6148956148808383896001600160a01b03166148a2565b886001600160a01b0316808204910615150190565b925050505b949350505050565b5f6148ae8484846146f4565b905081806148be576148be6158fb565b8385091561478a576001018061478a575f5ffd5b5f6001600160801b038416156001600160a01b0386161517156148fc57634f2461b85f526004601cfd5b816149135761490e85858560016149df565b613dab565b613dab8585856001614aca565b5f6001600160801b038416156001600160a01b03861615171561494a57634f2461b85f526004601cfd5b8161495b5761490e8585855f614aca565b613dab8585855f6149df565b6040516390bfb86560e01b8082526001600160a01b03851660048301526001600160e01b031984166024830152608060448301526020601f3d018190040260a0810160648401523d608484015290913d5f60a483013e60048260a4018201526001600160e01b031984168260c4018201528160e40181fd5b5f8115614a4f575f6001600160a01b03841115614a1357614a0e84600160601b876001600160801b03166146f4565b614a2a565b614a2a6001600160801b038616606086901b61590f565b9050614a47614a42826001600160a01b038916615a83565b614bac565b91505061489a565b5f6001600160a01b03841115614a7c57614a7784600160601b876001600160801b03166148a2565b614a99565b614a99606085901b6001600160801b038716808204910615150190565b9050806001600160a01b03871611614ab857634323a5555f526004601cfd5b6001600160a01b03861603905061489a565b5f825f03614ad957508361489a565b600160601b600160e01b03606085901b168215614b6b576001600160a01b03861684810290858281614b0d57614b0d6158fb565b0403614b3d57818101828110614b3b57614b3183896001600160a01b0316836148a2565b935050505061489a565b505b50614a478185614b566001600160a01b038a168361590f565b614b609190615a83565b808204910615150190565b6001600160a01b038616848102908582041481831116614b925763f5c787f15f526004601cfd5b808203614b31614a42846001600160a01b038b16846148a2565b806001600160a01b0381168114612ba157612ba16393dafdf160e01b613d20565b6001600160a01b0391909116815260200190565b5f60208284031215614bf1575f5ffd5b5035919050565b6001600160a01b0381168114612a70575f5ffd5b8035612ba181614bf8565b5f60a08284031215614c27575f5ffd5b50919050565b5f60808284031215614c27575f5ffd5b5f5f83601f840112614c4d575f5ffd5b5081356001600160401b03811115614c63575f5ffd5b602083019150836020828501011115614c7a575f5ffd5b9250929050565b5f5f5f5f5f6101608688031215614c96575f5ffd5b8535614ca181614bf8565b9450614cb08760208801614c17565b9350614cbf8760c08801614c2d565b92506101408601356001600160401b03811115614cda575f5ffd5b614ce688828901614c3d565b969995985093965092949392505050565b62ffffff81168114612a70575f5ffd5b8035612ba181614cf7565b5f60208284031215614d22575f5ffd5b813561478a81614cf7565b5f60208284031215614d3d575f5ffd5b813561478a81614bf8565b5f60208284031215614d58575f5ffd5b81356001600160401b03811115614d6d575f5ffd5b8201610140818503121561478a575f5ffd5b5f5f60408385031215614d90575f5ffd5b50508035926020909101359150565b8015158114612a70575f5ffd5b5f5f60408385031215614dbd575f5ffd5b8235614dc881614bf8565b91506020830135614dd881614d9f565b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715614e1f57614e1f614de3565b604052919050565b5f60808284031215614e37575f5ffd5b604051608081016001600160401b0381118282101715614e5957614e59614de3565b6040529050808235614e6a81614cf7565b81526020830135614e7a81614cf7565b60208201526040830135614e8d81614cf7565b60408201526060830135614ea081614d9f565b6060919091015292915050565b5f5f60a08385031215614ebe575f5ffd5b82359150614ecf8460208501614e27565b90509250929050565b80516001600160a01b03908116835260208083015182169084015260408083015162ffffff169084015260608083015160020b9084015260809182015116910152565b60a081016122108284614ed8565b5f60608284031215614f39575f5ffd5b604051606081016001600160401b0381118282101715614f5b57614f5b614de3565b6040529050808235614f6c81614d9f565b8152602083810135908201526040830135614f8681614bf8565b6040919091015292915050565b5f5f5f5f5f6101408688031215614fa8575f5ffd5b8535614fb381614bf8565b9450614fc28760208801614c17565b9350614fd18760c08801614f29565b92506101208601356001600160401b03811115614cda575f5ffd5b5f5f5f5f5f5f5f6101a0888a031215615003575f5ffd5b873561500e81614bf8565b965061501d8960208a01614c17565b955061502c8960c08a01614c2d565b9450610140880135935061016088013592506101808801356001600160401b03811115615057575f5ffd5b6150638a828b01614c3d565b989b979a50959850939692959293505050565b8060020b8114612a70575f5ffd5b5f5f5f5f6101008587031215615098575f5ffd5b84356150a381614bf8565b93506150b28660208701614c17565b925060c08501356150c281614bf8565b915060e08501356150d281615076565b939692955090935050565b5f602082840312156150ed575f5ffd5b813561478a81614d9f565b5f5f60208385031215615109575f5ffd5b82356001600160401b0381111561511e575f5ffd5b61512a85828601614c3d565b90969095509350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61220d6020830184615136565b5f60a0828403128015615187575f5ffd5b5060405160a081016001600160401b03811182821017156151aa576151aa614de3565b60405282356151b881614bf8565b815260208301356151c881614bf8565b602082015260408301356151db81614cf7565b604082015260608301356151ee81615076565b6060820152608083013561520181614bf8565b60808201529392505050565b5f60808201905062ffffff835116825262ffffff602084015116602083015262ffffff604084015116604083015260608301511515606083015292915050565b5f5f5f5f5f5f868803610160811215615264575f5ffd5b873561526f81614bf8565b965061527e8960208a01614c17565b9550606060bf1982011215615291575f5ffd5b5060c08701935061012087013592506101408701356001600160401b038111156152b9575f5ffd5b6152c589828a01614c3d565b979a9699509497509295939492505050565b5f5f5f5f5f5f61012087890312156152ed575f5ffd5b86356152f881614bf8565b95506153078860208901614c17565b945060c0870135935060e087013592506101008701356001600160401b038111156152b9575f5ffd5b8151151581526101c08101602083015161534e602084018215159052565b506040830151615362604084018215159052565b506060830151615376606084018215159052565b50608083015161538a608084018215159052565b5060a083015161539e60a084018215159052565b5060c08301516153b260c084018215159052565b5060e08301516153c660e084018215159052565b506101008301516153dc61010084018215159052565b506101208301516153f261012084018215159052565b5061014083015161540861014084018215159052565b5061016083015161541e61016084018215159052565b5061018083015161543461018084018215159052565b506101a083015161544a6101a084018215159052565b5092915050565b5f60808284031215615461575f5ffd5b61220d8383614e27565b5f5f5f60e0848603121561547d575f5ffd5b833561548881614bf8565b92506154978560208601614c17565b915060c08401356154a781614bf8565b809150509250925092565b5f5f8335601e198436030181126154c7575f5ffd5b83016020810192503590506001600160401b038111156154e5575f5ffd5b803603821315614c7a575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f61552682836154b2565b610140855261553a610140860182846154f3565b91505061554a60208401846154b2565b858303602087015261555d8382846154f3565b9250505061556e60408401846154b2565b85830360408701526155818382846154f3565b606086810135908801526080808701359088015292506155a691505060a08401614c0c565b6001600160a01b031660a08501526155c060c08401614d07565b62ffffff1660c085015260e083810135908501526155e26101008401846154b2565b8583036101008701526155f68382846154f3565b925050506156086101208401846154b2565b8583036101208701526129b98382846154f3565b602081525f61220d602083018461551b565b5f5f5f60608486031215615640575f5ffd5b835161564b81614bf8565b602085015190935061565c81614bf8565b6040949094015192959394509192915050565b6001600160a01b0385811682528481166020830152831660408201526101008101613dab6060830184614ed8565b5f5f8335601e198436030181126156b2575f5ffd5b8301803591506001600160401b038211156156cb575f5ffd5b602001915036819003821315614c7a575f5ffd5b838152604060208201525f613dab6040830184866154f3565b6001600160a01b038516815283151560208201526060604082018190525f906129b990830184866154f3565b5f60208284031215615734575f5ffd5b815161478a81614bf8565b60c0810161574d8285614ed8565b6001600160a01b039290921660a09190910152919050565b5f60208284031215615775575f5ffd5b815161478a81615076565b6001600160a01b038781168252861660208201526040810185905283151560608201526080810183905260c060a082018190525f906157c19083018461551b565b98975050505050505050565b5f60c082840312156157dd575f5ffd5b60405160c081016001600160401b03811182821017156157ff576157ff614de3565b60409081528351825260208085015190830152830151909150819061582381615076565b6040820152606083810151908201526080808401519082015260a083015161584a81614d9f565b60a0919091015292915050565b5f60c08284031215615867575f5ffd5b61220d83836157cd565b6001600160a01b03929092168252602082015260400190565b5f6020828403121561589a575f5ffd5b815161478a81614d9f565b634e487b7160e01b5f52601160045260245ffd5b81810381811115612210576122106158a5565b828152604060208201525f61489a604083018461551b565b8082028115828204841417612210576122106158a5565b634e487b7160e01b5f52601260045260245ffd5b5f8261592957634e487b7160e01b5f52601260045260245ffd5b500490565b6001600160a01b039485168152928416602084015292166040820152606081019190915260800190565b803561596381614bf8565b6001600160a01b03168252602081013561597c81614bf8565b6001600160a01b03166020830152604081013561599881614cf7565b62ffffff16604083015260608101356159b081615076565b60020b606083015260808101356159c681614bf8565b6001600160a01b03166080929092019190915250565b60e081016159ea8286615958565b8360a083015282151560c0830152949350505050565b5f5f5f6101008486031215615a13575f5ffd5b835160208501519093509150615a2c85604086016157cd565b90509250925092565b600f81810b9083900b0160016001607f1b03811360016001607f1b031982121715612210576122106158a5565b5f600f82900b6001607f1b8101615a7b57615a7b6158a5565b5f0392915050565b80820180821115612210576122106158a5565b6001600160a01b039390931683526020830191909152604082015260600190565b5f60208284031215615ac7575f5ffd5b81516001600160401b03811115615adc575f5ffd5b8201601f81018413615aec575f5ffd5b80516001600160401b03811115615b0557615b05614de3565b615b18601f8201601f1916602001614df7565b818152856020838501011115615b2c575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b5f60608284031215615b59575f5ffd5b61220d8383614f29565b8035615b6e81614d9f565b15158252602081810135908301526040810135615b8a81614bf8565b6001600160a01b03166040929092019190915250565b6001600160a01b0387168152615bb96020820187615958565b615bc660c0820186615b63565b836101208201526101606101408201525f6157c1610160830184866154f3565b6001600160a01b038416815260a08101615c036020830185615b63565b826080830152949350505050565b6001600160a01b03841681526040602082018190525f90613dab90830184866154f3565b5f60208284031215615c45575f5ffd5b5051919050565b602081525f61489a6020830184866154f3565b83815263ffffffff60e01b83166020820152606060408201525f613dab6060830184615136565b5f600160ff1b8201615c9a57615c9a6158a5565b505f0390565b6001600160a01b039384168152919092166020820152604081019190915260600190565b5f6040828403128015615cd5575f5ffd5b50604080519081016001600160401b0381118282101715615cf857615cf8614de3565b6040528251615d0681614cf7565b81526020830151615d1681614d9f565b60208201529392505050565b6001600160a01b038181168382160190811115612210576122106158a5565b6001600160a01b038281168282160390811115612210576122106158a5565b6101008101615d6f8287614ed8565b8460a08301528360020b60c083015282151560e083015295945050505050565b6101208101615d9e8286615958565b8351151560a0830152602084015160c08301526040909301516001600160a01b031660e082015262ffffff9190911661010090910152919050565b5f60208284031215615de9575f5ffd5b815161478a81614cf756fe2153348e08952a9737a728deab496289242e3884d90dd8e5fc8362f809e843141a426dc34b0367779dd37e66c4b193647bd2fd094d26b263dccbe93433d0a25bb5692db6de7b607ce10bbcb6a38fb443823969b63ffea1591751f25f9fe7059a0f457c8132f52d9098267126b744d1c04e04ccc48dce7c5b4396b38879d8b08aca7619d8d03127c99378decb6d4c002b5b8b8d60918ea28c2300a7811e7c1300aac705f00f1df487f4863cf5b7baa49be15a9dc73d477f33ca4d694d45cac7fef5955824691119aec0d1c983542055bfb83b16d3bc8d04e6cd271bbd244e894a6bc3f16992c8e44a3018b72b6d7aec54e64e1cdec74e64c5154c782ffe8aff2dee887e5283585aea7d61ae82ec91b0d9dbc7a52caa1fbc4f511651d06967058d84457b11f2602289c00904214483ea820055ea325cc54ad9b512e76c6541dfb4342e7ff505a8a0364cd0dc2ff195c315e43bce86b204846ecd36913e117b109e17cab649bf7619b4535991407dcf8f717035632cf45997a05c0f90b7baafa479170e502bdc7039c895af9c291daf6592705d439a7873ee28dbc514d2f1aa51b3859fad1dd1a8a0c97fb9e594529061ff17256ed2fd92004dfd6d47a8324651c4a26469706673582212203339e440194ab39ceb8733f51f3ab09ad7686aaf9f649795de1013015b97604764736f6c634300081b003360e060405234801561000f575f5ffd5b5060405161280b38038061280b83398101604081905261002e91610109565b3360a0526001600160a01b0383811660c052821660805267016345785d8a00005f8190556040519081527ff33554d7210318fe10b4d834a6233a2fd789ea1db37278b7e27f433b9e8659c79060200160405180910390a161008e81610096565b505050610149565b638b78c6d8198054156100b057630dc149f05f526004601cfd5b6001600160a01b03909116801560ff1b8117909155805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b80516001600160a01b0381168114610104575f5ffd5b919050565b5f5f5f6060848603121561011b575f5ffd5b610124846100ee565b9250610132602085016100ee565b9150610140604085016100ee565b90509250925092565b60805160a05160c0516125fe61020d5f395f8181610313015281816104ed0152818161060b01528181610a0701528181610aae01528181610b5401528181610c0c0152610d7601525f818161015f015281816103a50152818161051a015281816109c501528181610b8101528181610e7c01526111a601525f81816102e0015281816105a7015281816109080152818161095201528181610fa7015281816116620152818161173c015281816117a70152818161184801526118b601526125fe5ff3fe6080604052600436106100cd575f3560e01c8063088a80c7146100d157806325692962146100f25780634e944d57146100fa57806354d1f13d1461013e578063715018a614610146578063791b98bc1461014e578063870b9c171461018e5780638cebd942146101ad5780638da5cb5b1461023f578063957d1fe114610257578063ad49d66e14610291578063cfbc79fe146102b0578063dc4c90d3146102cf578063e1758bd814610302578063f04e283e14610335578063f2fde38b14610348578063fee81cf41461035b575b5f5ffd5b3480156100dc575f5ffd5b506100f06100eb366004612149565b61039a565b005b6100f061079a565b348015610105575f5ffd5b50610129610114366004612197565b5f9081526001602052604090205460ff161590565b60405190151581526020015b60405180910390f35b6100f06107e6565b6100f061081f565b348015610159575f5ffd5b506101817f000000000000000000000000000000000000000000000000000000000000000081565b60405161013591906121ae565b348015610199575f5ffd5b506100f06101a8366004612197565b610832565b3480156101b8575f5ffd5b506102096101c7366004612197565b600160208190525f918252604090912080549181015460029182015460ff80851694610100810490911693620100008204810b93600160281b909204900b9186565b6040805196151587529415156020870152600293840b94860194909452910b6060840152608083015260a082015260c001610135565b34801561024a575f5ffd5b50638b78c6d81954610181565b348015610262575f5ffd5b50610276610271366004612197565b610874565b60408051938452602084019290925290820152606001610135565b34801561029c575f5ffd5b506100f06102ab3660046121c2565b6109ba565b3480156102bb575f5ffd5b506100f06102ca3660046121dc565b610d70565b3480156102da575f5ffd5b506101817f000000000000000000000000000000000000000000000000000000000000000081565b34801561030d575f5ffd5b506101817f000000000000000000000000000000000000000000000000000000000000000081565b6100f0610343366004612212565b610f29565b6100f0610356366004612212565b610f66565b348015610366575f5ffd5b5061038c610375366004612212565b63389a75e1600c9081525f91909152602090205490565b604051908152602001610135565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103e35760405163041fb8cb60e31b815260040160405180910390fd5b82156107945760a084205f8181526001602052604081206002810180549192879261040f908490612241565b9250508190555084816001015f8282546104299190612241565b9091555050600181015460405183917fa5f00d866fe48379f7d8a625ccea2f92572b5cbd3f234591dfa4c7f845836fd09161046c91898252602082015260400190565b60405180910390a25f5481600101541015610488575050610794565b6001810180545f9182905582549091908190610100900460ff161561058b5783546104cb908a908890620100008104600290810b91600160281b9004900b610f8c565b909250905081156105865760405163a9059cbb60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb90610544907f0000000000000000000000000000000000000000000000000000000000000000908690600401612254565b6020604051808303815f875af1158015610560573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610584919061226d565b505b610599565b835461ff0019166101001784555b5f6105cd6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001687611044565b50509150508760020b8160020b131515871515036105e9578097505b6105fe8a888a6105f98888612241565b6110f6565b811561072e575f61062f8b7f0000000000000000000000000000000000000000000000000000000000000000611229565b90505f816001600160a01b03166361d027b36040518163ffffffff1660e01b8152600401602060405180830381865afa15801561066e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106929190612288565b60405163a9059cbb60e01b81529091506001600160a01b0383169063a9059cbb906106c39084908890600401612254565b6020604051808303815f875af11580156106df573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610703919061226d565b50875f5160206125a95f395f51905f528286604051610723929190612254565b60405180910390a250505b857fba69105014fb62310aea3b6cd667c51d18201a3c7112e2ead21829195f6fa3fc61075a8686612241565b875460408051928352620100008204600290810b6020850152600160281b90920490910b9082015260600160405180910390a25050505050505b50505050565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b610827611255565b6108305f61126f565b565b61083a611255565b5f8190556040518181527ff33554d7210318fe10b4d834a6233a2fd789ea1db37278b7e27f433b9e8659c79060200160405180910390a150565b5f818152600160208181526040808420815160c081018352815460ff808216151583526101008204161515948201859052620100008104600290810b94830194909452600160281b9004830b6060820152938101546080850152015460a0830152829182916108ee57608001515f935083925090506109b3565b604081015160608201515f9161093e916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691899130919066189a591dd85b1b60ca1b6112b5565b509091505f90506109786001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001688611044565b50505090506109a18161098e856040015161130c565b61099b866060015161130c565b856115c4565b60809094015190965092945091925050505b9193909250565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a035760405163041fb8cb60e31b815260040160405180910390fd5b80517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081169116145f610a408360a0902090565b5f8181526001602052604081208054929350918190610100900460ff1615610a95578254610a869087908790620100008104600290810b91600160281b9004900b610f8c565b845461ff001916855590925090505b6001830180545f918290556002850182905590610ad2887f0000000000000000000000000000000000000000000000000000000000000000611229565b90505f816001600160a01b03166361d027b36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b11573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b359190612288565b90508215610bef576040516323b872dd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd90610bad907f000000000000000000000000000000000000000000000000000000000000000090859088906004016122a3565b6020604051808303815f875af1158015610bc9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bed919061226d565b505b8415610c855760405163a9059cbb60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb90610c439084908990600401612254565b6020604051808303815f875af1158015610c5f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c83919061226d565b505b8315610d225760405163a9059cbb60e01b81526001600160a01b0383169063a9059cbb90610cb99084908890600401612254565b6020604051808303815f875af1158015610cd5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cf9919061226d565b50865f5160206125a95f395f51905f528286604051610d19929190612254565b60405180910390a25b867f98332f6bdb1b434256cc2c7b313581d76b7763b9c764bd32b25d461ff42c2bc382610d4f8689612241565b604051610d5d929190612254565b60405180910390a2505050505050505050565b610d9a827f0000000000000000000000000000000000000000000000000000000000000000611229565b6001600160a01b03166302d05d3f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dd5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610df99190612288565b6001600160a01b0316336001600160a01b031614610e2a57604051632a118c8960e01b815260040160405180910390fd5b5f60015f610e398560a0902090565b815260208101919091526040015f20805490915060ff16151582151503610e5f57505050565b8115610edf576040516356a4eb3760e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063ad49d66e90610eb190869060040161230a565b5f604051808303815f87803b158015610ec8575f5ffd5b505af1158015610eda573d5f5f3e3d5ffd5b505050505b805460ff191682151517815560a0832060405183151581527f88e359a3e31b529597eae3a8c8a107e55d6ff85edf8ca19a7368a9be35f012719060200160405180910390a2505050565b610f31611255565b63389a75e1600c52805f526020600c208054421115610f5757636f5e88185f526004601cfd5b5f9055610f638161126f565b50565b610f6e611255565b8060601b610f8357637448fbae5f526004601cfd5b610f638161126f565b5f5f5f610fdb610f9d8860a0902090565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169030888866189a591dd85b1b60ca1b6112b5565b505090505f610ff688878785610ff090612318565b3061165f565b9050866110165761100781600f0b90565b6110118260801d90565b61102a565b6110208160801d90565b61102a82600f0b90565b6001600160801b039182169a911698509650505050505050565b5f5f5f5f5f61105286611948565b604051631e2eaeaf60e01b8152600481018290529091505f906001600160a01b03891690631e2eaeaf90602401602060405180830381865afa15801561109a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110be9190612339565b90506001600160a01b03811695508060a01c60020b945062ffffff8160b81c16935062ffffff8160d01c169250505092959194509250565b5f8361110c57611107600184612350565b611117565b611117836001612375565b90505f5f5f861561116157611130600285900b5f611984565b925061113d603c84612375565b915061115a61114b8461130c565b6111548461130c565b87611a1a565b905061119d565b611170600285900b6001611984565b915061117d603c83612350565b925061119a61118b8461130c565b6111948461130c565b87611a8d565b90505b6111ca888484847f000000000000000000000000000000000000000000000000000000000000000061165f565b505f60015f6111da8b60a0902090565b815260208101919091526040015f20805467ffffffffffff000019166201000062ffffff9687160262ffffff60281b191617600160281b94909516939093029390931790915550505050505050565b81515f906001600160a01b0380841691161461124657825161124c565b82602001515b90505b92915050565b638b78c6d819543314610830576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3811560ff1b8217905550565b60408051602681018390526006810184905260038101859052858152603a600c8201205f9282018390526020820183905290829052819081906112f98a8a83611aca565b919c909b50909950975050505050505050565b60020b5f60ff82901d80830118620d89e8811115611335576113356345c3193d60e11b84611b6d565b7001fffcb933bd6fad37aa2d162d1a5940016001821602600160801b186002821615611371576ffff97272373d413259a46990580e213a0260801c5b6004821615611390576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b60088216156113af576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b60108216156113ce576fffcb9843d60f6159c9db58835c9266440260801c5b60208216156113ed576fff973b41fa98c081472e6896dfb254c00260801c5b604082161561140c576fff2ea16466c96a3843ec78b326b528610260801c5b608082161561142b576ffe5dee046a99a2a811c461f1969c30530260801c5b61010082161561144b576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b61020082161561146b576ff987a7253ac413176f2b074cf7815e540260801c5b61040082161561148b576ff3392b0822b70005940c7a398e4b70f30260801c5b6108008216156114ab576fe7159475a2c29b7443b29c7fa6e889d90260801c5b6110008216156114cb576fd097f3bdfd2022b8845ad8f792aa58250260801c5b6120008216156114eb576fa9f746462d870fdf8a65dc1f90e061e50260801c5b61400082161561150b576f70d869a156d2a1b890bb3df62baf32f70260801c5b61800082161561152b576f31be135f97d08fd981231505542fcfa60260801c5b6201000082161561154c576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b6202000082161561156c576e5d6af8dedb81196699c329225ee6040260801c5b6204000082161561158b576d2216e584f5fa1ea926041bedfe980260801c5b620800008216156115a8576b048a170391f7dc42444e8fa20260801c5b5f8413156115b4575f19045b63ffffffff0160201c9392505050565b5f5f836001600160a01b0316856001600160a01b031611156115e4579293925b846001600160a01b0316866001600160a01b03161161160f57611608858585611b7c565b9150611656565b836001600160a01b0316866001600160a01b0316101561164857611634868585611b7c565b9150611641858785611be5565b9050611656565b611653858585611be5565b90505b94509492505050565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635a6bcfda8760405180608001604052808960020b81526020018860020b815260200187600f0b815260200166189a591dd85b1b60ca1b8152506040518363ffffffff1660e01b81526004016116e192919061239a565b60408051808303815f875af11580156116fc573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061172091906123eb565b5090505f61172e8260801d90565b600f0b12156117915761178c7f0000000000000000000000000000000000000000000000000000000000000000836117668460801d90565b61176f90612318565b89516001600160a01b03169291906001600160801b03165f611c2e565b61182f565b5f61179c8260801d90565b600f0b131561182f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630b0d9c09875f0151846117e38560801d90565b6040518463ffffffff1660e01b81526004016118019392919061240d565b5f604051808303815f87803b158015611818575f5ffd5b505af115801561182a573d5f5f3e3d5ffd5b505050505b5f61183a82600f0b90565b600f0b12156118a05761189b7f00000000000000000000000000000000000000000000000000000000000000008361187284600f0b90565b61187b90612318565b60208a01516001600160a01b03169291906001600160801b03165f611c2e565b61193f565b5f6118ab82600f0b90565b600f0b131561193f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630b0d9c098760200151846118f385600f0b90565b6040518463ffffffff1660e01b81526004016119119392919061240d565b5f604051808303815f87803b158015611928575f5ffd5b505af115801561193a573d5f5f3e3d5ffd5b505050505b95945050505050565b6040515f90611967908390600690602001918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b5f620d89b319600284900b12156119a157620d89b31992506119b7565b620d89b4600284900b13156119b757620d89b492505b6119c2603c8461244d565b60020b5f036119d257508161124f565b603c6119de818561246e565b6119e891906124a6565b90505f8360020b1215611a0357611a00603c82612350565b90505b8161124f57611a13603c82612375565b905061124f565b5f826001600160a01b0316846001600160a01b03161115611a39579192915b5f611a5b856001600160a01b0316856001600160a01b0316600160601b611ef2565b9050611a82611a7d8483611a6f89896124cc565b6001600160a01b0316611ef2565b611f8e565b9150505b9392505050565b5f826001600160a01b0316846001600160a01b03161115611aac579192915b611ac2611a7d83600160601b611a6f88886124cc565b949350505050565b5f5f5f5f611ad88686611fe5565b604051631afeb18d60e11b815260048101829052600360248201529091505f906001600160a01b038916906335fd631a906044015f60405180830381865afa158015611b26573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611b4d91908101906124eb565b60208101516040820151606090920151909a919950975095505050505050565b815f528060020b60045260245ffd5b5f826001600160a01b0316846001600160a01b03161115611b9b579192915b6001600160a01b038416611bdb600160601b600160e01b03606085901b16611bc387876124cc565b6001600160a01b0316866001600160a01b0316611ef2565b611ac29190612595565b5f826001600160a01b0316846001600160a01b03161115611c04579192915b611ac26001600160801b038316611c1b86866124cc565b6001600160a01b0316600160601b611ef2565b8015611cc257836001600160a01b031663f5298aca84611c5d886001600160a01b03166001600160a01b031690565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604481018590526064015f604051808303815f87803b158015611ca7575f5ffd5b505af1158015611cb9573d5f5f3e3d5ffd5b50505050611eeb565b6001600160a01b038516611d3957836001600160a01b03166311da60b4836040518263ffffffff1660e01b815260040160206040518083038185885af1158015611d0e573d5f5f3e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190611d339190612339565b50611eeb565b604051632961046560e21b81526001600160a01b0385169063a584119490611d659088906004016121ae565b5f604051808303815f87803b158015611d7c575f5ffd5b505af1158015611d8e573d5f5f3e3d5ffd5b505050506001600160a01b0383163014611e18576040516323b872dd60e01b81526001600160a01b038616906323b872dd90611dd2908690889087906004016122a3565b6020604051808303815f875af1158015611dee573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e12919061226d565b50611e88565b60405163a9059cbb60e01b81526001600160a01b0386169063a9059cbb90611e469087908690600401612254565b6020604051808303815f875af1158015611e62573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e86919061226d565b505b836001600160a01b03166311da60b46040518163ffffffff1660e01b81526004016020604051808303815f875af1158015611ec5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ee99190612339565b505b5050505050565b5f838302815f1985870982811083820303915050808411611f11575f5ffd5b805f03611f2357508290049050611a86565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b806001600160801b0381168114611fe05760405162461bcd60e51b81526020600482015260126024820152716c6971756964697479206f766572666c6f7760701b604482015260640160405180910390fd5b919050565b5f5f611ff084611948565b90505f611ffe600683612241565b6040805160208101879052908101829052909150606001604051602081830303815290604052805190602001209250505092915050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561207157612071612035565b604052919050565b6001600160a01b0381168114610f63575f5ffd5b8035600281900b8114611fe0575f5ffd5b5f60a082840312156120ae575f5ffd5b60405160a081016001600160401b03811182821017156120d0576120d0612035565b60405290508082356120e181612079565b815260208301356120f181612079565b6020820152604083013562ffffff8116811461210b575f5ffd5b604082015261211c6060840161208d565b6060820152608083013561212f81612079565b6080919091015292915050565b8015158114610f63575f5ffd5b5f5f5f5f610100858703121561215d575f5ffd5b612167868661209e565b935060a0850135925061217c60c0860161208d565b915060e085013561218c8161213c565b939692955090935050565b5f602082840312156121a7575f5ffd5b5035919050565b6001600160a01b0391909116815260200190565b5f60a082840312156121d2575f5ffd5b61124c838361209e565b5f5f60c083850312156121ed575f5ffd5b6121f7848461209e565b915060a08301356122078161213c565b809150509250929050565b5f60208284031215612222575f5ffd5b8135611a8681612079565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561124f5761124f61222d565b6001600160a01b03929092168252602082015260400190565b5f6020828403121561227d575f5ffd5b8151611a868161213c565b5f60208284031215612298575f5ffd5b8151611a8681612079565b6001600160a01b039384168152919092166020820152604081019190915260600190565b80516001600160a01b03908116835260208083015182169084015260408083015162ffffff169084015260608083015160020b9084015260809182015116910152565b60a0810161124f82846122c7565b5f600f82900b6001607f1b81016123315761233161222d565b5f0392915050565b5f60208284031215612349575f5ffd5b5051919050565b600282810b9082900b03627fffff198112627fffff8213171561124f5761124f61222d565b600281810b9083900b01627fffff8113627fffff198212171561124f5761124f61222d565b6123a481846122c7565b8151600290810b60a08301526020830151900b60c0820152604082015160e082015260609091015161010082015261014061012082018190525f9082015261016001919050565b5f5f604083850312156123fc575f5ffd5b505080516020909101519092909150565b6001600160a01b0393841681529190921660208201526001600160801b03909116604082015260600190565b634e487b7160e01b5f52601260045260245ffd5b5f8260020b8061245f5761245f612439565b808360020b0791505092915050565b5f8160020b8360020b8061248457612484612439565b627fffff1982145f198214161561249d5761249d61222d565b90059392505050565b5f8260020b8260020b028060020b91508082146124c5576124c561222d565b5092915050565b6001600160a01b03828116828216039081111561124f5761124f61222d565b5f602082840312156124fb575f5ffd5b81516001600160401b03811115612510575f5ffd5b8201601f81018413612520575f5ffd5b80516001600160401b0381111561253957612539612035565b8060051b61254960208201612049565b91825260208184018101929081019087841115612564575f5ffd5b6020850194505b8385101561258a5784518083526020958601959093509091019061256b565b979650505050505050565b5f826125a3576125a3612439565b50049056fe6f25beea688da23574729528d9040ac00b2a9f98fe04601175b5bed75222d37ba2646970667358221220de9d4bfc6feccedaea14cf8f5936327299730ff4b5a7fd076ae086207ef0c3d664736f6c634300081b003360c060405234801561000f575f5ffd5b50604051611bda380380611bda83398101604081905261002e91610043565b6001600160a01b03166080523360a052610070565b5f60208284031215610053575f5ffd5b81516001600160a01b0381168114610069575f5ffd5b9392505050565b60805160a051611afc6100de5f395f8181610100015281816101a7015281816102da01528181610599015281816107b1015281816108850152818161093b01528181610d340152610db701525f818161015401528181610c3201528181610d130152610d960152611afc5ff3fe608060405234801561000f575f5ffd5b5060043610610081575f3560e01c8063197a866b146100855780631cc8fb001461009a5780635cb0fdf4146100c25780636264c08e146100d9578063791b98bc146100fb578063d37f0cb21461012f578063dc4c90d31461014f578063e4f203f614610176578063e7ea93c714610189575b5f5ffd5b610098610093366004611621565b61019c565b005b6100ad6100a8366004611641565b610253565b60405190151581526020015b60405180910390f35b6100cb61070881565b6040519081526020016100b9565b6100ec6100e736600461173b565b6102c5565b6040516100b9939291906117b8565b6101227f000000000000000000000000000000000000000000000000000000000000000081565b6040516100b991906117d4565b61014261013d3660046117e8565b610586565b6040516100b99190611820565b6101227f000000000000000000000000000000000000000000000000000000000000000081565b610142610184366004611641565b610738565b61014261019736600461173b565b61079e565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101e55760405163041fb8cb60e31b815260040160405180910390fd5b5f811215610221576101f681611842565b5f838152602081905260408120600301805490919061021690849061185c565b9091555061024f9050565b5f81131561024f575f828152602081905260408120600301805483929061024990849061186f565b90915550505b5050565b5f81815260208181526040808320815160c0810183528154808252600183015494820194909452600280830154900b9281019290925260038101546060830152600481015460808301526005015460ff16151560a08201529042108015906102be5750806020015142105b9392505050565b5f5f6102cf6115ec565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103185760405163041fb8cb60e31b815260040160405180910390fd5b60a086205f81815260208190526040812090879003610385576040805160c0810182528254815260018301546020820152600280840154900b91810191909152600382015460608201526004820154608082015260059091015460ff16151560a0820152915061057d9050565b5f5f5f8912156103d75761039889611842565b6002808501549193506103d091900b838a6103b7578c602001516103ba565b8c515b8b6103c6578d51610a3c565b8d60200151610a3c565b905061040a565b50600280830154899161040791900b828a156103f7578c602001516103fa565b8c515b8b156103c6578d51610a3c565b91505b8260040154811115610464575f818460040154670de0b6b3a76400006104309190611882565b61043a91906118ad565b9050670de0b6b3a764000061044f8285611882565b61045991906118ad565b925083600401549150505b5f89126104a15761049c61047782610b1e565b610480906118c0565b61048984610b1e565b6001600160801b031660809190911b1790565b6104bf565b6104bf6104ad83610b1e565b6104b683610b1e565b610489906118c0565b96506104f7886104e0576104d282610b1e565b6104db906118c0565b6104e9565b6104e983610b1e565b896104ad5761048984610b1e565b955081836003015f82825461050c919061186f565b9250508190555080836004015f828254610526919061185c565b90915550506040805160c0810182528454815260018501546020820152600280860154900b91810191909152600384015460608201526004840154608082015260059093015460ff16151560a08401525090925050505b93509350939050565b61058e6115ec565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105d75760405163041fb8cb60e31b815260040160405180910390fd5b5f6105e46107088561186f565b90506040518060c001604052808581526020018281526020018660020b81526020015f81526020018481526020015f15158152505f5f8881526020019081526020015f205f820151815f0155602082015181600101556040820151816002015f6101000a81548162ffffff021916908360020b62ffffff160217905550606082015181600301556080820151816004015560a0820151816005015f6101000a81548160ff021916908315150217905550905050857f0935055505897986a1e75d206c9f6c9ff968893b8820199e8d781008b03550a88486846040516106cb939291906118e1565b60405180910390a250505f8481526020818152604091829020825160c08101845281548152600182015492810192909252600280820154900b92820192909252600382015460608201526004820154608082015260059091015460ff16151560a08201525b949350505050565b6107406115ec565b505f9081526020818152604091829020825160c08101845281548152600182015492810192909252600280820154900b92820192909252600382015460608201526004820154608082015260059091015460ff16151560a082015290565b6107a66115ec565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107ef5760405163041fb8cb60e31b815260040160405180910390fd5b5f5f5f6107fd8760a0902090565b81526020019081526020015f2090505f5f84156108d557600280840154610835915f9161082c910b60016118f7565b60020b90610b3e565b9150610842603c836118f7565b905061085687838386600301546001610bd6565b600280840154620d89b31993506108779160019161082c918391900b61191c565b90506108d0878383896108c07f00000000000000000000000000000000000000000000000000000000000000008d602001516001600160a01b0316610e0d90919063ffffffff16565b6108ca919061185c565b5f610bd6565b610986565b6002808401546108ef9160019161082c918391900b61191c565b90506108fc603c8261191c565b915061090f87838386600301545f610bd6565b600280840154610927915f9161082c910b60016118f7565b9150620d89b49050610986878383896109757f00000000000000000000000000000000000000000000000000000000000000008d5f01516001600160a01b0316610e0d90919063ffffffff16565b61097f919061185c565b6001610bd6565b42600184810182905560058501805460ff1916909117905560a088206003850154600486015460405192937f90423afe5c41b867c789a3a14a8084a4cd82970dfde321694c1ea4d64612df99936109df939291906118e1565b60405180910390a250506040805160c0810182528254815260018301546020820152600280840154900b91810191909152600382015460608201526004820154608082015260059091015460ff16151560a0820152949350505050565b5f5f610a4786610ea1565b90506001600160801b036001600160a01b03821611610aba575f610a746001600160a01b03831680611882565b9050836001600160a01b0316856001600160a01b031610610aa357610a9e600160c01b8783611159565b610ab2565b610ab28187600160c01b611159565b925050610b15565b5f610ad36001600160a01b03831680600160401b611159565b9050836001600160a01b0316856001600160a01b031610610b0257610afd600160801b8783611159565b610b11565b610b118187600160801b611159565b9250505b50949350505050565b5f6001607f1b8210610b3a57610b3a6393dafdf160e01b6111f5565b5090565b5f620d89b319600284900b1215610b5b57620d89b3199250610b71565b620d89b4600284900b1315610b7157620d89b492505b610b7c603c84611941565b60020b5f03610b8c575081610bd0565b603c610b988185611962565b610ba2919061199a565b90505f8360020b1215610bbd57610bba603c8261191c565b90505b81610bd057610bcd603c826118f7565b90505b92915050565b5f81610bfc57610bf7610be886610ea1565b610bf186610ea1565b856111fd565b610c17565b610c17610c0886610ea1565b610c1186610ea1565b85611245565b9050806001600160801b03165f03610c2f5750610e06565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635a6bcfda8860405180608001604052808a60020b81526020018960020b8152602001610c8f876001600160801b0316610b1e565b600f0b81526020015f8152506040518363ffffffff1660e01b8152600401610cb89291906119c0565b60408051808303815f875af1158015610cd3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cf79190611a53565b5090505f610d058260801d90565b600f0b1215610d7d57610d7d7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000610d5d8460801d90565b600f0b610d6990611842565b8a516001600160a01b03169291905f6112a3565b5f610d8882600f0b90565b600f0b1215610e0357610e037f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000610de084600f0b90565b600f0b610dec90611842565b60208b01516001600160a01b03169291905f6112a3565b50505b5050505050565b5f610e20836001600160a01b0316611579565b15610e3657506001600160a01b03811631610bd0565b6040516370a0823160e01b81526001600160a01b038416906370a0823190610e629085906004016117d4565b602060405180830381865afa158015610e7d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bcd9190611a75565b60020b5f60ff82901d80830118620d89e8811115610eca57610eca6345c3193d60e11b84611586565b7001fffcb933bd6fad37aa2d162d1a5940016001821602600160801b186002821615610f06576ffff97272373d413259a46990580e213a0260801c5b6004821615610f25576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615610f44576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b6010821615610f63576fffcb9843d60f6159c9db58835c9266440260801c5b6020821615610f82576fff973b41fa98c081472e6896dfb254c00260801c5b6040821615610fa1576fff2ea16466c96a3843ec78b326b528610260801c5b6080821615610fc0576ffe5dee046a99a2a811c461f1969c30530260801c5b610100821615610fe0576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b610200821615611000576ff987a7253ac413176f2b074cf7815e540260801c5b610400821615611020576ff3392b0822b70005940c7a398e4b70f30260801c5b610800821615611040576fe7159475a2c29b7443b29c7fa6e889d90260801c5b611000821615611060576fd097f3bdfd2022b8845ad8f792aa58250260801c5b612000821615611080576fa9f746462d870fdf8a65dc1f90e061e50260801c5b6140008216156110a0576f70d869a156d2a1b890bb3df62baf32f70260801c5b6180008216156110c0576f31be135f97d08fd981231505542fcfa60260801c5b620100008216156110e1576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b62020000821615611101576e5d6af8dedb81196699c329225ee6040260801c5b62040000821615611120576d2216e584f5fa1ea926041bedfe980260801c5b6208000082161561113d576b048a170391f7dc42444e8fa20260801c5b5f841315611149575f19045b63ffffffff0160201c9392505050565b5f838302815f1985870982811083820303915050808411611178575f5ffd5b805f0361118a575082900490506102be565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b805f5260045ffd5b5f826001600160a01b0316846001600160a01b0316111561121c579192915b61073061124083600160601b6112328888611a8c565b6001600160a01b0316611159565b611595565b5f826001600160a01b0316846001600160a01b03161115611264579192915b5f611286856001600160a01b0316856001600160a01b0316600160601b611159565b905061129a61124084836112328989611a8c565b95945050505050565b801561133757836001600160a01b031663f5298aca846112d2886001600160a01b03166001600160a01b031690565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604481018590526064015f604051808303815f87803b15801561131c575f5ffd5b505af115801561132e573d5f5f3e3d5ffd5b50505050610e06565b611349856001600160a01b0316611579565b156113b757836001600160a01b03166311da60b4836040518263ffffffff1660e01b815260040160206040518083038185885af115801561138c573d5f5f3e3d5ffd5b50505050506040513d601f19601f820116820180604052508101906113b19190611a75565b50610e06565b604051632961046560e21b81526001600160a01b0385169063a5841194906113e39088906004016117d4565b5f604051808303815f87803b1580156113fa575f5ffd5b505af115801561140c573d5f5f3e3d5ffd5b505050506001600160a01b038316301461149e576040516323b872dd60e01b81526001600160a01b0384811660048301528581166024830152604482018490528616906323b872dd906064016020604051808303815f875af1158015611474573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114989190611aab565b50611510565b60405163a9059cbb60e01b81526001600160a01b0385811660048301526024820184905286169063a9059cbb906044016020604051808303815f875af11580156114ea573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061150e9190611aab565b505b836001600160a01b03166311da60b46040518163ffffffff1660e01b81526004016020604051808303815f875af115801561154d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115719190611a75565b505050505050565b6001600160a01b03161590565b815f528060020b60045260245ffd5b806001600160801b03811681146115e75760405162461bcd60e51b81526020600482015260126024820152716c6971756964697479206f766572666c6f7760701b604482015260640160405180910390fd5b919050565b6040518060c001604052805f81526020015f81526020015f60020b81526020015f81526020015f81526020015f151581525090565b5f5f60408385031215611632575f5ffd5b50508035926020909101359150565b5f60208284031215611651575f5ffd5b5035919050565b80356001600160a01b03811681146115e7575f5ffd5b803562ffffff811681146115e7575f5ffd5b8035600281900b81146115e7575f5ffd5b5f60a082840312156116a1575f5ffd5b60405160a081016001600160401b03811182821017156116cf57634e487b7160e01b5f52604160045260245ffd5b6040529050806116de83611658565b81526116ec60208401611658565b60208201526116fd6040840161166e565b604082015261170e60608401611680565b606082015261171f60808401611658565b60808201525092915050565b8015158114611738575f5ffd5b50565b5f5f5f60e0848603121561174d575f5ffd5b6117578585611691565b925060a0840135915060c084013561176e8161172b565b809150509250925092565b8051825260208101516020830152604081015160020b6040830152606081015160608301526080810151608083015260a0810151151560a08301525050565b8381526020810183905261010081016107306040830184611779565b6001600160a01b0391909116815260200190565b5f5f5f5f608085870312156117fb575f5ffd5b8435935061180b60208601611680565b93969395505050506040820135916060013590565b60c08101610bd08284611779565b634e487b7160e01b5f52601160045260245ffd5b5f600160ff1b82016118565761185661182e565b505f0390565b81810381811115610bd057610bd061182e565b80820180821115610bd057610bd061182e565b8082028115828204841417610bd057610bd061182e565b634e487b7160e01b5f52601260045260245ffd5b5f826118bb576118bb611899565b500490565b5f600f82900b6001607f1b81016118d9576118d961182e565b5f0392915050565b9283526020830191909152604082015260600190565b600281810b9083900b01627fffff8113627fffff1982121715610bd057610bd061182e565b600282810b9082900b03627fffff198112627fffff82131715610bd057610bd061182e565b5f8260020b8061195357611953611899565b808360020b0791505092915050565b5f8160020b8360020b8061197857611978611899565b627fffff1982145f19821416156119915761199161182e565b90059392505050565b5f8260020b8260020b028060020b91508082146119b9576119b961182e565b5092915050565b82516001600160a01b03908116825260208085015182169083015260408085015162ffffff169083015260608085015160020b9083015260808085015190911690820152611a3860a0820183805160020b8252602081015160020b602083015260408101516040830152606081015160608301525050565b6101406101208201525f61073061014083015f815260200190565b5f5f60408385031215611a64575f5ffd5b505080516020909101519092909150565b5f60208284031215611a85575f5ffd5b5051919050565b6001600160a01b038281168282160390811115610bd057610bd061182e565b5f60208284031215611abb575f5ffd5b81516102be8161172b56fea2646970667358221220876bb60b472c63a1e8881bba7db4d3e47798616abde1b0901c0da9f7cff936d464736f6c634300081b00336080604052348015600e575f5ffd5b506040516104d63803806104d6833981016040819052602b91608e565b6032816037565b5060b9565b638b78c6d819805415605057630dc149f05f526004601cfd5b6001600160a01b03909116801560ff1b8117909155805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b5f60208284031215609d575f5ffd5b81516001600160a01b038116811460b2575f5ffd5b9392505050565b610410806100c65f395ff3fe608060405260043610610080575f3560e01c8063256929621461008457806354d1f13d1461008e57806359a8c9cc1461009657806370d34a0e146100b5578063715018a6146100f85780638da5cb5b14610100578063d0ce51e41461012b578063f04e283e1461014a578063f2fde38b1461015d578063fee81cf414610170575b5f5ffd5b61008c6101af565b005b61008c6101fb565b3480156100a1575f5ffd5b5061008c6100b03660046103ad565b610234565b3480156100c0575f5ffd5b506100e36100cf3660046103ad565b5f6020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61008c610287565b34801561010b575f5ffd5b50638b78c6d819546040516001600160a01b0390911681526020016100ef565b348015610136575f5ffd5b5061008c6101453660046103ad565b61029a565b61008c6101583660046103ad565b6102ea565b61008c61016b3660046103ad565b610327565b34801561017b575f5ffd5b506101a161018a3660046103ad565b63389a75e1600c9081525f91909152602090205490565b6040519081526020016100ef565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b61023c61034d565b6001600160a01b0381165f81815260208190526040808220805460ff19166001179055517fc27063b13f068c906426cb21fe4cd7d9b11b516de8e4daf1a608a67a6c1a204a9190a250565b61028f61034d565b6102985f610367565b565b6102a261034d565b6001600160a01b0381165f81815260208190526040808220805460ff19169055517fc89da65ab0bb433dd2433dd17b376528a87fc0fe74b21ff53b800dd834c1290f9190a250565b6102f261034d565b63389a75e1600c52805f526020600c20805442111561031857636f5e88185f526004601cfd5b5f905561032481610367565b50565b61032f61034d565b8060601b61034457637448fbae5f526004601cfd5b61032481610367565b638b78c6d819543314610298576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3811560ff1b8217905550565b5f602082840312156103bd575f5ffd5b81356001600160a01b03811681146103d3575f5ffd5b939250505056fea2646970667358221220b28ec331c094b4db99895f3bfddbee382a97473b1b5ac53dd253dd213ac4fb9b64736f6c634300081b00336080604052348015600e575f5ffd5b506040516109d33803806109d3833981016040819052602b916084565b600280546001600160a01b031916331790556044816049565b5060af565b6001600160a01b0316638b78c6d819819055805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b5f602082840312156093575f5ffd5b81516001600160a01b038116811460a8575f5ffd5b9392505050565b610917806100bc5f395ff3fe608060405260043610610080575f3560e01c80632569296214610084578063314e79ad1461008e57806354d1f13d146100ad578063715018a6146100b55780637262561c146100bd5780638da5cb5b146100dc578063d5f9a39714610105578063f04e283e14610124578063f2fde38b14610137578063fee81cf41461014a575b5f5ffd5b61008c610189565b005b348015610099575f5ffd5b5061008c6100a836600461070d565b6101d5565b61008c61027c565b61008c6102b5565b3480156100c8575f5ffd5b5061008c6100d736600461078b565b6102c8565b3480156100e7575f5ffd5b50638b78c6d819546040516100fc91906107a4565b60405180910390f35b348015610110575f5ffd5b5061008c61011f3660046107b8565b61036f565b61008c61013236600461078b565b61044e565b61008c61014536600461078b565b610488565b348015610155575f5ffd5b5061017b61016436600461078b565b63389a75e1600c9081525f91909152602090205490565b6040519081526020016100fc565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b6002546001600160a01b031633146101eb575f5ffd5b5f6101f55f6104ae565b90505f5b818110156102745761020b5f826104bd565b6001600160a01b0316634ca1f062878787876040518563ffffffff1660e01b815260040161023c949392919061082e565b5f604051808303815f87803b158015610253575f5ffd5b505af1158015610265573d5f5f3e3d5ffd5b505050508060010190506101f9565b505050505050565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b6102bd6104cf565b6102c65f6104e9565b565b6102d06104cf565b6102da5f82610526565b1561036c576102e95f8261053a565b50806001600160a01b031663fcae44846040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610322575f5ffd5b505af1925050508015610333575060015b507f804bdbbf25e80c5635294ecf375fa16a531a4b7ffa416f74e121ae40bf4a00fe8160405161036391906107a4565b60405180910390a15b50565b6103776104cf565b6103815f8461054e565b15610449576040516323085b8560e11b81526001600160a01b03841690634610b70a906103b49085908590600401610860565b6020604051808303815f875af11580156103d0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103f4919061087b565b61041157604051633bf9fecd60e01b815260040160405180910390fd5b7fa6610d70b57486ac216657a87c0fce7f41758dd77d526daf7974f813a9b2a46a8360405161044091906107a4565b60405180910390a15b505050565b6104566104cf565b63389a75e1600c52805f526020600c20805442111561047c57636f5e88185f526004601cfd5b5f905561036c816104e9565b6104906104cf565b8060601b6104a557637448fbae5f526004601cfd5b61036c816104e9565b5f6104b7825490565b92915050565b5f6104c88383610562565b9392505050565b638b78c6d8195433146102c6576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a355565b5f6104c8836001600160a01b038416610588565b5f6104c8836001600160a01b03841661059f565b5f6104c8836001600160a01b038416610682565b5f825f0182815481106105775761057761089a565b905f5260205f200154905092915050565b5f9081526001919091016020526040902054151590565b5f8181526001830160205260408120548015610679575f6105c16001836108ae565b85549091505f906105d4906001906108ae565b9050808214610633575f865f0182815481106105f2576105f261089a565b905f5260205f200154905080875f0184815481106106125761061261089a565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080610644576106446108cd565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506104b7565b5f9150506104b7565b5f61068d8383610588565b6106c257508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556104b7565b505f6104b7565b5f5f83601f8401126106d9575f5ffd5b5081356001600160401b038111156106ef575f5ffd5b602083019150836020828501011115610706575f5ffd5b9250929050565b5f5f5f5f60608587031215610720575f5ffd5b8435935060208501356001600160e01b03198116811461073e575f5ffd5b925060408501356001600160401b03811115610758575f5ffd5b610764878288016106c9565b95989497509550505050565b80356001600160a01b0381168114610786575f5ffd5b919050565b5f6020828403121561079b575f5ffd5b6104c882610770565b6001600160a01b0391909116815260200190565b5f5f5f604084860312156107ca575f5ffd5b6107d384610770565b925060208401356001600160401b038111156107ed575f5ffd5b6107f9868287016106c9565b9497909650939450505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b84815263ffffffff60e01b84166020820152606060408201525f610856606083018486610806565b9695505050505050565b602081525f610873602083018486610806565b949350505050565b5f6020828403121561088b575f5ffd5b815180151581146104c8575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b818103818111156104b757634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220e7d3fe79eab14863996255c9403a53f57af392f2a80225ae8bdabd164fe2b2aa64736f6c634300081b0033",
}

// PositionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PositionManagerMetaData.ABI instead.
var PositionManagerABI = PositionManagerMetaData.ABI

// PositionManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PositionManagerMetaData.Bin instead.
var PositionManagerBin = PositionManagerMetaData.Bin

// DeployPositionManager deploys a new Ethereum contract, binding an instance of PositionManager to it.
func DeployPositionManager(auth *bind.TransactOpts, backend bind.ContractBackend, params PositionManagerConstructorParams) (common.Address, *types.Transaction, *PositionManager, error) {
	parsed, err := PositionManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PositionManagerBin), backend, params)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PositionManager{PositionManagerCaller: PositionManagerCaller{contract: contract}, PositionManagerTransactor: PositionManagerTransactor{contract: contract}, PositionManagerFilterer: PositionManagerFilterer{contract: contract}}, nil
}

// PositionManager is an auto generated Go binding around an Ethereum contract.
type PositionManager struct {
	PositionManagerCaller     // Read-only binding to the contract
	PositionManagerTransactor // Write-only binding to the contract
	PositionManagerFilterer   // Log filterer for contract events
}

// PositionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PositionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PositionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PositionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PositionManagerSession struct {
	Contract     *PositionManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PositionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PositionManagerCallerSession struct {
	Contract *PositionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PositionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PositionManagerTransactorSession struct {
	Contract     *PositionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PositionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PositionManagerRaw struct {
	Contract *PositionManager // Generic contract binding to access the raw methods on
}

// PositionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PositionManagerCallerRaw struct {
	Contract *PositionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PositionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PositionManagerTransactorRaw struct {
	Contract *PositionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPositionManager creates a new instance of PositionManager, bound to a specific deployed contract.
func NewPositionManager(address common.Address, backend bind.ContractBackend) (*PositionManager, error) {
	contract, err := bindPositionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PositionManager{PositionManagerCaller: PositionManagerCaller{contract: contract}, PositionManagerTransactor: PositionManagerTransactor{contract: contract}, PositionManagerFilterer: PositionManagerFilterer{contract: contract}}, nil
}

// NewPositionManagerCaller creates a new read-only instance of PositionManager, bound to a specific deployed contract.
func NewPositionManagerCaller(address common.Address, caller bind.ContractCaller) (*PositionManagerCaller, error) {
	contract, err := bindPositionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PositionManagerCaller{contract: contract}, nil
}

// NewPositionManagerTransactor creates a new write-only instance of PositionManager, bound to a specific deployed contract.
func NewPositionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PositionManagerTransactor, error) {
	contract, err := bindPositionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PositionManagerTransactor{contract: contract}, nil
}

// NewPositionManagerFilterer creates a new log filterer instance of PositionManager, bound to a specific deployed contract.
func NewPositionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PositionManagerFilterer, error) {
	contract, err := bindPositionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PositionManagerFilterer{contract: contract}, nil
}

// bindPositionManager binds a generic wrapper to an already deployed contract.
func bindPositionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PositionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionManager *PositionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionManager.Contract.PositionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionManager *PositionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManager.Contract.PositionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionManager *PositionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionManager.Contract.PositionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionManager *PositionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionManager *PositionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionManager *PositionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionManager.Contract.contract.Transact(opts, method, params...)
}

// MAXPROTOCOLALLOCATION is a free data retrieval call binding the contract method 0xc29c945b.
//
// Solidity: function MAX_PROTOCOL_ALLOCATION() view returns(uint24)
func (_PositionManager *PositionManagerCaller) MAXPROTOCOLALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "MAX_PROTOCOL_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPROTOCOLALLOCATION is a free data retrieval call binding the contract method 0xc29c945b.
//
// Solidity: function MAX_PROTOCOL_ALLOCATION() view returns(uint24)
func (_PositionManager *PositionManagerSession) MAXPROTOCOLALLOCATION() (*big.Int, error) {
	return _PositionManager.Contract.MAXPROTOCOLALLOCATION(&_PositionManager.CallOpts)
}

// MAXPROTOCOLALLOCATION is a free data retrieval call binding the contract method 0xc29c945b.
//
// Solidity: function MAX_PROTOCOL_ALLOCATION() view returns(uint24)
func (_PositionManager *PositionManagerCallerSession) MAXPROTOCOLALLOCATION() (*big.Int, error) {
	return _PositionManager.Contract.MAXPROTOCOLALLOCATION(&_PositionManager.CallOpts)
}

// MINDISTRIBUTETHRESHOLD is a free data retrieval call binding the contract method 0x299d532e.
//
// Solidity: function MIN_DISTRIBUTE_THRESHOLD() view returns(uint256)
func (_PositionManager *PositionManagerCaller) MINDISTRIBUTETHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "MIN_DISTRIBUTE_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDISTRIBUTETHRESHOLD is a free data retrieval call binding the contract method 0x299d532e.
//
// Solidity: function MIN_DISTRIBUTE_THRESHOLD() view returns(uint256)
func (_PositionManager *PositionManagerSession) MINDISTRIBUTETHRESHOLD() (*big.Int, error) {
	return _PositionManager.Contract.MINDISTRIBUTETHRESHOLD(&_PositionManager.CallOpts)
}

// MINDISTRIBUTETHRESHOLD is a free data retrieval call binding the contract method 0x299d532e.
//
// Solidity: function MIN_DISTRIBUTE_THRESHOLD() view returns(uint256)
func (_PositionManager *PositionManagerCallerSession) MINDISTRIBUTETHRESHOLD() (*big.Int, error) {
	return _PositionManager.Contract.MINDISTRIBUTETHRESHOLD(&_PositionManager.CallOpts)
}

// ActionManager is a free data retrieval call binding the contract method 0x75d252a5.
//
// Solidity: function actionManager() view returns(address)
func (_PositionManager *PositionManagerCaller) ActionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "actionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActionManager is a free data retrieval call binding the contract method 0x75d252a5.
//
// Solidity: function actionManager() view returns(address)
func (_PositionManager *PositionManagerSession) ActionManager() (common.Address, error) {
	return _PositionManager.Contract.ActionManager(&_PositionManager.CallOpts)
}

// ActionManager is a free data retrieval call binding the contract method 0x75d252a5.
//
// Solidity: function actionManager() view returns(address)
func (_PositionManager *PositionManagerCallerSession) ActionManager() (common.Address, error) {
	return _PositionManager.Contract.ActionManager(&_PositionManager.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address _recipient) view returns(uint256 _amount)
func (_PositionManager *PositionManagerCaller) Balances(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "balances", _recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address _recipient) view returns(uint256 _amount)
func (_PositionManager *PositionManagerSession) Balances(_recipient common.Address) (*big.Int, error) {
	return _PositionManager.Contract.Balances(&_PositionManager.CallOpts, _recipient)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address _recipient) view returns(uint256 _amount)
func (_PositionManager *PositionManagerCallerSession) Balances(_recipient common.Address) (*big.Int, error) {
	return _PositionManager.Contract.Balances(&_PositionManager.CallOpts, _recipient)
}

// BeforeAddLiquidity is a free data retrieval call binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , bytes ) view returns(bytes4 selector_)
func (_PositionManager *PositionManagerCaller) BeforeAddLiquidity(opts *bind.CallOpts, _sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "beforeAddLiquidity", _sender, _key, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// BeforeAddLiquidity is a free data retrieval call binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , bytes ) view returns(bytes4 selector_)
func (_PositionManager *PositionManagerSession) BeforeAddLiquidity(_sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	return _PositionManager.Contract.BeforeAddLiquidity(&_PositionManager.CallOpts, _sender, _key, arg2, arg3)
}

// BeforeAddLiquidity is a free data retrieval call binding the contract method 0x259982e5.
//
// Solidity: function beforeAddLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , bytes ) view returns(bytes4 selector_)
func (_PositionManager *PositionManagerCallerSession) BeforeAddLiquidity(_sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	return _PositionManager.Contract.BeforeAddLiquidity(&_PositionManager.CallOpts, _sender, _key, arg2, arg3)
}

// BeforeInitialize is a free data retrieval call binding the contract method 0xdc98354e.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 ) view returns(bytes4)
func (_PositionManager *PositionManagerCaller) BeforeInitialize(opts *bind.CallOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int) ([4]byte, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "beforeInitialize", arg0, arg1, arg2)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// BeforeInitialize is a free data retrieval call binding the contract method 0xdc98354e.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 ) view returns(bytes4)
func (_PositionManager *PositionManagerSession) BeforeInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int) ([4]byte, error) {
	return _PositionManager.Contract.BeforeInitialize(&_PositionManager.CallOpts, arg0, arg1, arg2)
}

// BeforeInitialize is a free data retrieval call binding the contract method 0xdc98354e.
//
// Solidity: function beforeInitialize(address , (address,address,uint24,int24,address) , uint160 ) view returns(bytes4)
func (_PositionManager *PositionManagerCallerSession) BeforeInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int) ([4]byte, error) {
	return _PositionManager.Contract.BeforeInitialize(&_PositionManager.CallOpts, arg0, arg1, arg2)
}

// BeforeRemoveLiquidity is a free data retrieval call binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , bytes ) view returns(bytes4 selector_)
func (_PositionManager *PositionManagerCaller) BeforeRemoveLiquidity(opts *bind.CallOpts, _sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "beforeRemoveLiquidity", _sender, _key, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// BeforeRemoveLiquidity is a free data retrieval call binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , bytes ) view returns(bytes4 selector_)
func (_PositionManager *PositionManagerSession) BeforeRemoveLiquidity(_sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	return _PositionManager.Contract.BeforeRemoveLiquidity(&_PositionManager.CallOpts, _sender, _key, arg2, arg3)
}

// BeforeRemoveLiquidity is a free data retrieval call binding the contract method 0x21d0ee70.
//
// Solidity: function beforeRemoveLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , bytes ) view returns(bytes4 selector_)
func (_PositionManager *PositionManagerCallerSession) BeforeRemoveLiquidity(_sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, arg3 []byte) ([4]byte, error) {
	return _PositionManager.Contract.BeforeRemoveLiquidity(&_PositionManager.CallOpts, _sender, _key, arg2, arg3)
}

// BidWall is a free data retrieval call binding the contract method 0xba3e69b7.
//
// Solidity: function bidWall() view returns(address)
func (_PositionManager *PositionManagerCaller) BidWall(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "bidWall")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BidWall is a free data retrieval call binding the contract method 0xba3e69b7.
//
// Solidity: function bidWall() view returns(address)
func (_PositionManager *PositionManagerSession) BidWall() (common.Address, error) {
	return _PositionManager.Contract.BidWall(&_PositionManager.CallOpts)
}

// BidWall is a free data retrieval call binding the contract method 0xba3e69b7.
//
// Solidity: function bidWall() view returns(address)
func (_PositionManager *PositionManagerCallerSession) BidWall() (common.Address, error) {
	return _PositionManager.Contract.BidWall(&_PositionManager.CallOpts)
}

// FairLaunch is a free data retrieval call binding the contract method 0x94e1cf96.
//
// Solidity: function fairLaunch() view returns(address)
func (_PositionManager *PositionManagerCaller) FairLaunch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "fairLaunch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FairLaunch is a free data retrieval call binding the contract method 0x94e1cf96.
//
// Solidity: function fairLaunch() view returns(address)
func (_PositionManager *PositionManagerSession) FairLaunch() (common.Address, error) {
	return _PositionManager.Contract.FairLaunch(&_PositionManager.CallOpts)
}

// FairLaunch is a free data retrieval call binding the contract method 0x94e1cf96.
//
// Solidity: function fairLaunch() view returns(address)
func (_PositionManager *PositionManagerCallerSession) FairLaunch() (common.Address, error) {
	return _PositionManager.Contract.FairLaunch(&_PositionManager.CallOpts)
}

// FairLaunchFeeCalculator is a free data retrieval call binding the contract method 0xa1dd2d91.
//
// Solidity: function fairLaunchFeeCalculator() view returns(address)
func (_PositionManager *PositionManagerCaller) FairLaunchFeeCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "fairLaunchFeeCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FairLaunchFeeCalculator is a free data retrieval call binding the contract method 0xa1dd2d91.
//
// Solidity: function fairLaunchFeeCalculator() view returns(address)
func (_PositionManager *PositionManagerSession) FairLaunchFeeCalculator() (common.Address, error) {
	return _PositionManager.Contract.FairLaunchFeeCalculator(&_PositionManager.CallOpts)
}

// FairLaunchFeeCalculator is a free data retrieval call binding the contract method 0xa1dd2d91.
//
// Solidity: function fairLaunchFeeCalculator() view returns(address)
func (_PositionManager *PositionManagerCallerSession) FairLaunchFeeCalculator() (common.Address, error) {
	return _PositionManager.Contract.FairLaunchFeeCalculator(&_PositionManager.CallOpts)
}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_PositionManager *PositionManagerCaller) FeeCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "feeCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_PositionManager *PositionManagerSession) FeeCalculator() (common.Address, error) {
	return _PositionManager.Contract.FeeCalculator(&_PositionManager.CallOpts)
}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_PositionManager *PositionManagerCallerSession) FeeCalculator() (common.Address, error) {
	return _PositionManager.Contract.FeeCalculator(&_PositionManager.CallOpts)
}

// FeeExemptions is a free data retrieval call binding the contract method 0x01a2cae8.
//
// Solidity: function feeExemptions() view returns(address)
func (_PositionManager *PositionManagerCaller) FeeExemptions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "feeExemptions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeExemptions is a free data retrieval call binding the contract method 0x01a2cae8.
//
// Solidity: function feeExemptions() view returns(address)
func (_PositionManager *PositionManagerSession) FeeExemptions() (common.Address, error) {
	return _PositionManager.Contract.FeeExemptions(&_PositionManager.CallOpts)
}

// FeeExemptions is a free data retrieval call binding the contract method 0x01a2cae8.
//
// Solidity: function feeExemptions() view returns(address)
func (_PositionManager *PositionManagerCallerSession) FeeExemptions() (common.Address, error) {
	return _PositionManager.Contract.FeeExemptions(&_PositionManager.CallOpts)
}

// FeeSplit is a free data retrieval call binding the contract method 0x4875cbb8.
//
// Solidity: function feeSplit(bytes32 _poolId, uint256 _amount) view returns(uint256 bidWall_, uint256 creator_, uint256 protocol_)
func (_PositionManager *PositionManagerCaller) FeeSplit(opts *bind.CallOpts, _poolId [32]byte, _amount *big.Int) (struct {
	BidWall  *big.Int
	Creator  *big.Int
	Protocol *big.Int
}, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "feeSplit", _poolId, _amount)

	outstruct := new(struct {
		BidWall  *big.Int
		Creator  *big.Int
		Protocol *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BidWall = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Protocol = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FeeSplit is a free data retrieval call binding the contract method 0x4875cbb8.
//
// Solidity: function feeSplit(bytes32 _poolId, uint256 _amount) view returns(uint256 bidWall_, uint256 creator_, uint256 protocol_)
func (_PositionManager *PositionManagerSession) FeeSplit(_poolId [32]byte, _amount *big.Int) (struct {
	BidWall  *big.Int
	Creator  *big.Int
	Protocol *big.Int
}, error) {
	return _PositionManager.Contract.FeeSplit(&_PositionManager.CallOpts, _poolId, _amount)
}

// FeeSplit is a free data retrieval call binding the contract method 0x4875cbb8.
//
// Solidity: function feeSplit(bytes32 _poolId, uint256 _amount) view returns(uint256 bidWall_, uint256 creator_, uint256 protocol_)
func (_PositionManager *PositionManagerCallerSession) FeeSplit(_poolId [32]byte, _amount *big.Int) (struct {
	BidWall  *big.Int
	Creator  *big.Int
	Protocol *big.Int
}, error) {
	return _PositionManager.Contract.FeeSplit(&_PositionManager.CallOpts, _poolId, _amount)
}

// FlaunchContract is a free data retrieval call binding the contract method 0x84aa1da0.
//
// Solidity: function flaunchContract() view returns(address)
func (_PositionManager *PositionManagerCaller) FlaunchContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "flaunchContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlaunchContract is a free data retrieval call binding the contract method 0x84aa1da0.
//
// Solidity: function flaunchContract() view returns(address)
func (_PositionManager *PositionManagerSession) FlaunchContract() (common.Address, error) {
	return _PositionManager.Contract.FlaunchContract(&_PositionManager.CallOpts)
}

// FlaunchContract is a free data retrieval call binding the contract method 0x84aa1da0.
//
// Solidity: function flaunchContract() view returns(address)
func (_PositionManager *PositionManagerCallerSession) FlaunchContract() (common.Address, error) {
	return _PositionManager.Contract.FlaunchContract(&_PositionManager.CallOpts)
}

// FlaunchesAt is a free data retrieval call binding the contract method 0x7907addf.
//
// Solidity: function flaunchesAt(bytes32 _poolId) view returns(uint256 _flaunchTime)
func (_PositionManager *PositionManagerCaller) FlaunchesAt(opts *bind.CallOpts, _poolId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "flaunchesAt", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlaunchesAt is a free data retrieval call binding the contract method 0x7907addf.
//
// Solidity: function flaunchesAt(bytes32 _poolId) view returns(uint256 _flaunchTime)
func (_PositionManager *PositionManagerSession) FlaunchesAt(_poolId [32]byte) (*big.Int, error) {
	return _PositionManager.Contract.FlaunchesAt(&_PositionManager.CallOpts, _poolId)
}

// FlaunchesAt is a free data retrieval call binding the contract method 0x7907addf.
//
// Solidity: function flaunchesAt(bytes32 _poolId) view returns(uint256 _flaunchTime)
func (_PositionManager *PositionManagerCallerSession) FlaunchesAt(_poolId [32]byte) (*big.Int, error) {
	return _PositionManager.Contract.FlaunchesAt(&_PositionManager.CallOpts, _poolId)
}

// FlayGovernance is a free data retrieval call binding the contract method 0xdddfecf5.
//
// Solidity: function flayGovernance() view returns(address)
func (_PositionManager *PositionManagerCaller) FlayGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "flayGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlayGovernance is a free data retrieval call binding the contract method 0xdddfecf5.
//
// Solidity: function flayGovernance() view returns(address)
func (_PositionManager *PositionManagerSession) FlayGovernance() (common.Address, error) {
	return _PositionManager.Contract.FlayGovernance(&_PositionManager.CallOpts)
}

// FlayGovernance is a free data retrieval call binding the contract method 0xdddfecf5.
//
// Solidity: function flayGovernance() view returns(address)
func (_PositionManager *PositionManagerCallerSession) FlayGovernance() (common.Address, error) {
	return _PositionManager.Contract.FlayGovernance(&_PositionManager.CallOpts)
}

// GetFeeCalculator is a free data retrieval call binding the contract method 0x71c4ddb0.
//
// Solidity: function getFeeCalculator(bool _isFairLaunch) view returns(address)
func (_PositionManager *PositionManagerCaller) GetFeeCalculator(opts *bind.CallOpts, _isFairLaunch bool) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "getFeeCalculator", _isFairLaunch)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeCalculator is a free data retrieval call binding the contract method 0x71c4ddb0.
//
// Solidity: function getFeeCalculator(bool _isFairLaunch) view returns(address)
func (_PositionManager *PositionManagerSession) GetFeeCalculator(_isFairLaunch bool) (common.Address, error) {
	return _PositionManager.Contract.GetFeeCalculator(&_PositionManager.CallOpts, _isFairLaunch)
}

// GetFeeCalculator is a free data retrieval call binding the contract method 0x71c4ddb0.
//
// Solidity: function getFeeCalculator(bool _isFairLaunch) view returns(address)
func (_PositionManager *PositionManagerCallerSession) GetFeeCalculator(_isFairLaunch bool) (common.Address, error) {
	return _PositionManager.Contract.GetFeeCalculator(&_PositionManager.CallOpts, _isFairLaunch)
}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0xec876d71.
//
// Solidity: function getFlaunchingFee(bytes _initialPriceParams) view returns(uint256)
func (_PositionManager *PositionManagerCaller) GetFlaunchingFee(opts *bind.CallOpts, _initialPriceParams []byte) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "getFlaunchingFee", _initialPriceParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0xec876d71.
//
// Solidity: function getFlaunchingFee(bytes _initialPriceParams) view returns(uint256)
func (_PositionManager *PositionManagerSession) GetFlaunchingFee(_initialPriceParams []byte) (*big.Int, error) {
	return _PositionManager.Contract.GetFlaunchingFee(&_PositionManager.CallOpts, _initialPriceParams)
}

// GetFlaunchingFee is a free data retrieval call binding the contract method 0xec876d71.
//
// Solidity: function getFlaunchingFee(bytes _initialPriceParams) view returns(uint256)
func (_PositionManager *PositionManagerCallerSession) GetFlaunchingFee(_initialPriceParams []byte) (*big.Int, error) {
	return _PositionManager.Contract.GetFlaunchingFee(&_PositionManager.CallOpts, _initialPriceParams)
}

// GetFlaunchingMarketCap is a free data retrieval call binding the contract method 0xff3575e4.
//
// Solidity: function getFlaunchingMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_PositionManager *PositionManagerCaller) GetFlaunchingMarketCap(opts *bind.CallOpts, _initialPriceParams []byte) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "getFlaunchingMarketCap", _initialPriceParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFlaunchingMarketCap is a free data retrieval call binding the contract method 0xff3575e4.
//
// Solidity: function getFlaunchingMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_PositionManager *PositionManagerSession) GetFlaunchingMarketCap(_initialPriceParams []byte) (*big.Int, error) {
	return _PositionManager.Contract.GetFlaunchingMarketCap(&_PositionManager.CallOpts, _initialPriceParams)
}

// GetFlaunchingMarketCap is a free data retrieval call binding the contract method 0xff3575e4.
//
// Solidity: function getFlaunchingMarketCap(bytes _initialPriceParams) view returns(uint256)
func (_PositionManager *PositionManagerCallerSession) GetFlaunchingMarketCap(_initialPriceParams []byte) (*big.Int, error) {
	return _PositionManager.Contract.GetFlaunchingMarketCap(&_PositionManager.CallOpts, _initialPriceParams)
}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_PositionManager *PositionManagerCaller) GetHookPermissions(opts *bind.CallOpts) (HooksPermissions, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "getHookPermissions")

	if err != nil {
		return *new(HooksPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new(HooksPermissions)).(*HooksPermissions)

	return out0, err

}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_PositionManager *PositionManagerSession) GetHookPermissions() (HooksPermissions, error) {
	return _PositionManager.Contract.GetHookPermissions(&_PositionManager.CallOpts)
}

// GetHookPermissions is a free data retrieval call binding the contract method 0xc4e833ce.
//
// Solidity: function getHookPermissions() pure returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_PositionManager *PositionManagerCallerSession) GetHookPermissions() (HooksPermissions, error) {
	return _PositionManager.Contract.GetHookPermissions(&_PositionManager.CallOpts)
}

// GetPoolFeeDistribution is a free data retrieval call binding the contract method 0xb3b42795.
//
// Solidity: function getPoolFeeDistribution(bytes32 _poolId) view returns((uint24,uint24,uint24,bool) feeDistribution_)
func (_PositionManager *PositionManagerCaller) GetPoolFeeDistribution(opts *bind.CallOpts, _poolId [32]byte) (FeeDistributorFeeDistribution, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "getPoolFeeDistribution", _poolId)

	if err != nil {
		return *new(FeeDistributorFeeDistribution), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeDistributorFeeDistribution)).(*FeeDistributorFeeDistribution)

	return out0, err

}

// GetPoolFeeDistribution is a free data retrieval call binding the contract method 0xb3b42795.
//
// Solidity: function getPoolFeeDistribution(bytes32 _poolId) view returns((uint24,uint24,uint24,bool) feeDistribution_)
func (_PositionManager *PositionManagerSession) GetPoolFeeDistribution(_poolId [32]byte) (FeeDistributorFeeDistribution, error) {
	return _PositionManager.Contract.GetPoolFeeDistribution(&_PositionManager.CallOpts, _poolId)
}

// GetPoolFeeDistribution is a free data retrieval call binding the contract method 0xb3b42795.
//
// Solidity: function getPoolFeeDistribution(bytes32 _poolId) view returns((uint24,uint24,uint24,bool) feeDistribution_)
func (_PositionManager *PositionManagerCallerSession) GetPoolFeeDistribution(_poolId [32]byte) (FeeDistributorFeeDistribution, error) {
	return _PositionManager.Contract.GetPoolFeeDistribution(&_PositionManager.CallOpts, _poolId)
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(address)
func (_PositionManager *PositionManagerCaller) InitialPrice(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "initialPrice")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(address)
func (_PositionManager *PositionManagerSession) InitialPrice() (common.Address, error) {
	return _PositionManager.Contract.InitialPrice(&_PositionManager.CallOpts)
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(address)
func (_PositionManager *PositionManagerCallerSession) InitialPrice() (common.Address, error) {
	return _PositionManager.Contract.InitialPrice(&_PositionManager.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_PositionManager *PositionManagerCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_PositionManager *PositionManagerSession) NativeToken() (common.Address, error) {
	return _PositionManager.Contract.NativeToken(&_PositionManager.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_PositionManager *PositionManagerCallerSession) NativeToken() (common.Address, error) {
	return _PositionManager.Contract.NativeToken(&_PositionManager.CallOpts)
}

// Notifier is a free data retrieval call binding the contract method 0x09276ea4.
//
// Solidity: function notifier() view returns(address)
func (_PositionManager *PositionManagerCaller) Notifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "notifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Notifier is a free data retrieval call binding the contract method 0x09276ea4.
//
// Solidity: function notifier() view returns(address)
func (_PositionManager *PositionManagerSession) Notifier() (common.Address, error) {
	return _PositionManager.Contract.Notifier(&_PositionManager.CallOpts)
}

// Notifier is a free data retrieval call binding the contract method 0x09276ea4.
//
// Solidity: function notifier() view returns(address)
func (_PositionManager *PositionManagerCallerSession) Notifier() (common.Address, error) {
	return _PositionManager.Contract.Notifier(&_PositionManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_PositionManager *PositionManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_PositionManager *PositionManagerSession) Owner() (common.Address, error) {
	return _PositionManager.Contract.Owner(&_PositionManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_PositionManager *PositionManagerCallerSession) Owner() (common.Address, error) {
	return _PositionManager.Contract.Owner(&_PositionManager.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_PositionManager *PositionManagerCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_PositionManager *PositionManagerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _PositionManager.Contract.OwnershipHandoverExpiresAt(&_PositionManager.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_PositionManager *PositionManagerCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _PositionManager.Contract.OwnershipHandoverExpiresAt(&_PositionManager.CallOpts, pendingOwner)
}

// PoolFees is a free data retrieval call binding the contract method 0xddb475d5.
//
// Solidity: function poolFees((address,address,uint24,int24,address) _poolKey) view returns((uint256,uint256))
func (_PositionManager *PositionManagerCaller) PoolFees(opts *bind.CallOpts, _poolKey PoolKey) (InternalSwapPoolClaimableFees, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "poolFees", _poolKey)

	if err != nil {
		return *new(InternalSwapPoolClaimableFees), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalSwapPoolClaimableFees)).(*InternalSwapPoolClaimableFees)

	return out0, err

}

// PoolFees is a free data retrieval call binding the contract method 0xddb475d5.
//
// Solidity: function poolFees((address,address,uint24,int24,address) _poolKey) view returns((uint256,uint256))
func (_PositionManager *PositionManagerSession) PoolFees(_poolKey PoolKey) (InternalSwapPoolClaimableFees, error) {
	return _PositionManager.Contract.PoolFees(&_PositionManager.CallOpts, _poolKey)
}

// PoolFees is a free data retrieval call binding the contract method 0xddb475d5.
//
// Solidity: function poolFees((address,address,uint24,int24,address) _poolKey) view returns((uint256,uint256))
func (_PositionManager *PositionManagerCallerSession) PoolFees(_poolKey PoolKey) (InternalSwapPoolClaimableFees, error) {
	return _PositionManager.Contract.PoolFees(&_PositionManager.CallOpts, _poolKey)
}

// PoolKey is a free data retrieval call binding the contract method 0x55d1cb60.
//
// Solidity: function poolKey(address _token) view returns((address,address,uint24,int24,address))
func (_PositionManager *PositionManagerCaller) PoolKey(opts *bind.CallOpts, _token common.Address) (PoolKey, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "poolKey", _token)

	if err != nil {
		return *new(PoolKey), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolKey)).(*PoolKey)

	return out0, err

}

// PoolKey is a free data retrieval call binding the contract method 0x55d1cb60.
//
// Solidity: function poolKey(address _token) view returns((address,address,uint24,int24,address))
func (_PositionManager *PositionManagerSession) PoolKey(_token common.Address) (PoolKey, error) {
	return _PositionManager.Contract.PoolKey(&_PositionManager.CallOpts, _token)
}

// PoolKey is a free data retrieval call binding the contract method 0x55d1cb60.
//
// Solidity: function poolKey(address _token) view returns((address,address,uint24,int24,address))
func (_PositionManager *PositionManagerCallerSession) PoolKey(_token common.Address) (PoolKey, error) {
	return _PositionManager.Contract.PoolKey(&_PositionManager.CallOpts, _token)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_PositionManager *PositionManagerCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_PositionManager *PositionManagerSession) PoolManager() (common.Address, error) {
	return _PositionManager.Contract.PoolManager(&_PositionManager.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_PositionManager *PositionManagerCallerSession) PoolManager() (common.Address, error) {
	return _PositionManager.Contract.PoolManager(&_PositionManager.CallOpts)
}

// PremineInfo is a free data retrieval call binding the contract method 0x1d61a816.
//
// Solidity: function premineInfo(bytes32 _poolId) view returns(int256 amountSpecified, uint256 blockNumber)
func (_PositionManager *PositionManagerCaller) PremineInfo(opts *bind.CallOpts, _poolId [32]byte) (struct {
	AmountSpecified *big.Int
	BlockNumber     *big.Int
}, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "premineInfo", _poolId)

	outstruct := new(struct {
		AmountSpecified *big.Int
		BlockNumber     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountSpecified = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PremineInfo is a free data retrieval call binding the contract method 0x1d61a816.
//
// Solidity: function premineInfo(bytes32 _poolId) view returns(int256 amountSpecified, uint256 blockNumber)
func (_PositionManager *PositionManagerSession) PremineInfo(_poolId [32]byte) (struct {
	AmountSpecified *big.Int
	BlockNumber     *big.Int
}, error) {
	return _PositionManager.Contract.PremineInfo(&_PositionManager.CallOpts, _poolId)
}

// PremineInfo is a free data retrieval call binding the contract method 0x1d61a816.
//
// Solidity: function premineInfo(bytes32 _poolId) view returns(int256 amountSpecified, uint256 blockNumber)
func (_PositionManager *PositionManagerCallerSession) PremineInfo(_poolId [32]byte) (struct {
	AmountSpecified *big.Int
	BlockNumber     *big.Int
}, error) {
	return _PositionManager.Contract.PremineInfo(&_PositionManager.CallOpts, _poolId)
}

// ReferralEscrow is a free data retrieval call binding the contract method 0x8f2bdf75.
//
// Solidity: function referralEscrow() view returns(address)
func (_PositionManager *PositionManagerCaller) ReferralEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "referralEscrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReferralEscrow is a free data retrieval call binding the contract method 0x8f2bdf75.
//
// Solidity: function referralEscrow() view returns(address)
func (_PositionManager *PositionManagerSession) ReferralEscrow() (common.Address, error) {
	return _PositionManager.Contract.ReferralEscrow(&_PositionManager.CallOpts)
}

// ReferralEscrow is a free data retrieval call binding the contract method 0x8f2bdf75.
//
// Solidity: function referralEscrow() view returns(address)
func (_PositionManager *PositionManagerCallerSession) ReferralEscrow() (common.Address, error) {
	return _PositionManager.Contract.ReferralEscrow(&_PositionManager.CallOpts)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x9f063efc.
//
// Solidity: function afterAddLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , int256 _delta, int256 _feesAccrued, bytes ) returns(bytes4 selector_, int256)
func (_PositionManager *PositionManagerTransactor) AfterAddLiquidity(opts *bind.TransactOpts, _sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, _delta *big.Int, _feesAccrued *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "afterAddLiquidity", _sender, _key, arg2, _delta, _feesAccrued, arg5)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x9f063efc.
//
// Solidity: function afterAddLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , int256 _delta, int256 _feesAccrued, bytes ) returns(bytes4 selector_, int256)
func (_PositionManager *PositionManagerSession) AfterAddLiquidity(_sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, _delta *big.Int, _feesAccrued *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.AfterAddLiquidity(&_PositionManager.TransactOpts, _sender, _key, arg2, _delta, _feesAccrued, arg5)
}

// AfterAddLiquidity is a paid mutator transaction binding the contract method 0x9f063efc.
//
// Solidity: function afterAddLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , int256 _delta, int256 _feesAccrued, bytes ) returns(bytes4 selector_, int256)
func (_PositionManager *PositionManagerTransactorSession) AfterAddLiquidity(_sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, _delta *big.Int, _feesAccrued *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.AfterAddLiquidity(&_PositionManager.TransactOpts, _sender, _key, arg2, _delta, _feesAccrued, arg5)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address _sender, (address,address,uint24,int24,address) _key, uint256 _amount0, uint256 _amount1, bytes ) returns(bytes4 selector_)
func (_PositionManager *PositionManagerTransactor) AfterDonate(opts *bind.TransactOpts, _sender common.Address, _key PoolKey, _amount0 *big.Int, _amount1 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "afterDonate", _sender, _key, _amount0, _amount1, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address _sender, (address,address,uint24,int24,address) _key, uint256 _amount0, uint256 _amount1, bytes ) returns(bytes4 selector_)
func (_PositionManager *PositionManagerSession) AfterDonate(_sender common.Address, _key PoolKey, _amount0 *big.Int, _amount1 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.AfterDonate(&_PositionManager.TransactOpts, _sender, _key, _amount0, _amount1, arg4)
}

// AfterDonate is a paid mutator transaction binding the contract method 0xe1b4af69.
//
// Solidity: function afterDonate(address _sender, (address,address,uint24,int24,address) _key, uint256 _amount0, uint256 _amount1, bytes ) returns(bytes4 selector_)
func (_PositionManager *PositionManagerTransactorSession) AfterDonate(_sender common.Address, _key PoolKey, _amount0 *big.Int, _amount1 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.AfterDonate(&_PositionManager.TransactOpts, _sender, _key, _amount0, _amount1, arg4)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0x6fe7e6eb.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 ) returns(bytes4)
func (_PositionManager *PositionManagerTransactor) AfterInitialize(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "afterInitialize", arg0, arg1, arg2, arg3)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0x6fe7e6eb.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 ) returns(bytes4)
func (_PositionManager *PositionManagerSession) AfterInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.AfterInitialize(&_PositionManager.TransactOpts, arg0, arg1, arg2, arg3)
}

// AfterInitialize is a paid mutator transaction binding the contract method 0x6fe7e6eb.
//
// Solidity: function afterInitialize(address , (address,address,uint24,int24,address) , uint160 , int24 ) returns(bytes4)
func (_PositionManager *PositionManagerTransactorSession) AfterInitialize(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.AfterInitialize(&_PositionManager.TransactOpts, arg0, arg1, arg2, arg3)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x6c2bbe7e.
//
// Solidity: function afterRemoveLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , int256 _delta, int256 _feesAccrued, bytes ) returns(bytes4 selector_, int256)
func (_PositionManager *PositionManagerTransactor) AfterRemoveLiquidity(opts *bind.TransactOpts, _sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, _delta *big.Int, _feesAccrued *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "afterRemoveLiquidity", _sender, _key, arg2, _delta, _feesAccrued, arg5)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x6c2bbe7e.
//
// Solidity: function afterRemoveLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , int256 _delta, int256 _feesAccrued, bytes ) returns(bytes4 selector_, int256)
func (_PositionManager *PositionManagerSession) AfterRemoveLiquidity(_sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, _delta *big.Int, _feesAccrued *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.AfterRemoveLiquidity(&_PositionManager.TransactOpts, _sender, _key, arg2, _delta, _feesAccrued, arg5)
}

// AfterRemoveLiquidity is a paid mutator transaction binding the contract method 0x6c2bbe7e.
//
// Solidity: function afterRemoveLiquidity(address _sender, (address,address,uint24,int24,address) _key, (int24,int24,int256,bytes32) , int256 _delta, int256 _feesAccrued, bytes ) returns(bytes4 selector_, int256)
func (_PositionManager *PositionManagerTransactorSession) AfterRemoveLiquidity(_sender common.Address, _key PoolKey, arg2 IPoolManagerModifyLiquidityParams, _delta *big.Int, _feesAccrued *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.AfterRemoveLiquidity(&_PositionManager.TransactOpts, _sender, _key, arg2, _delta, _feesAccrued, arg5)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address _sender, (address,address,uint24,int24,address) _key, (bool,int256,uint160) _params, int256 _delta, bytes _hookData) returns(bytes4 selector_, int128 hookDeltaUnspecified_)
func (_PositionManager *PositionManagerTransactor) AfterSwap(opts *bind.TransactOpts, _sender common.Address, _key PoolKey, _params IPoolManagerSwapParams, _delta *big.Int, _hookData []byte) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "afterSwap", _sender, _key, _params, _delta, _hookData)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address _sender, (address,address,uint24,int24,address) _key, (bool,int256,uint160) _params, int256 _delta, bytes _hookData) returns(bytes4 selector_, int128 hookDeltaUnspecified_)
func (_PositionManager *PositionManagerSession) AfterSwap(_sender common.Address, _key PoolKey, _params IPoolManagerSwapParams, _delta *big.Int, _hookData []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.AfterSwap(&_PositionManager.TransactOpts, _sender, _key, _params, _delta, _hookData)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xb47b2fb1.
//
// Solidity: function afterSwap(address _sender, (address,address,uint24,int24,address) _key, (bool,int256,uint160) _params, int256 _delta, bytes _hookData) returns(bytes4 selector_, int128 hookDeltaUnspecified_)
func (_PositionManager *PositionManagerTransactorSession) AfterSwap(_sender common.Address, _key PoolKey, _params IPoolManagerSwapParams, _delta *big.Int, _hookData []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.AfterSwap(&_PositionManager.TransactOpts, _sender, _key, _params, _delta, _hookData)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_PositionManager *PositionManagerTransactor) BeforeDonate(opts *bind.TransactOpts, arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "beforeDonate", arg0, arg1, arg2, arg3, arg4)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_PositionManager *PositionManagerSession) BeforeDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.BeforeDonate(&_PositionManager.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// BeforeDonate is a paid mutator transaction binding the contract method 0xb6a8b0fa.
//
// Solidity: function beforeDonate(address , (address,address,uint24,int24,address) , uint256 , uint256 , bytes ) returns(bytes4)
func (_PositionManager *PositionManagerTransactorSession) BeforeDonate(arg0 common.Address, arg1 PoolKey, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.BeforeDonate(&_PositionManager.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address _sender, (address,address,uint24,int24,address) _key, (bool,int256,uint160) _params, bytes _hookData) returns(bytes4 selector_, int256 beforeSwapDelta_, uint24)
func (_PositionManager *PositionManagerTransactor) BeforeSwap(opts *bind.TransactOpts, _sender common.Address, _key PoolKey, _params IPoolManagerSwapParams, _hookData []byte) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "beforeSwap", _sender, _key, _params, _hookData)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address _sender, (address,address,uint24,int24,address) _key, (bool,int256,uint160) _params, bytes _hookData) returns(bytes4 selector_, int256 beforeSwapDelta_, uint24)
func (_PositionManager *PositionManagerSession) BeforeSwap(_sender common.Address, _key PoolKey, _params IPoolManagerSwapParams, _hookData []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.BeforeSwap(&_PositionManager.TransactOpts, _sender, _key, _params, _hookData)
}

// BeforeSwap is a paid mutator transaction binding the contract method 0x575e24b4.
//
// Solidity: function beforeSwap(address _sender, (address,address,uint24,int24,address) _key, (bool,int256,uint160) _params, bytes _hookData) returns(bytes4 selector_, int256 beforeSwapDelta_, uint24)
func (_PositionManager *PositionManagerTransactorSession) BeforeSwap(_sender common.Address, _key PoolKey, _params IPoolManagerSwapParams, _hookData []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.BeforeSwap(&_PositionManager.TransactOpts, _sender, _key, _params, _hookData)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_PositionManager *PositionManagerTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_PositionManager *PositionManagerSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _PositionManager.Contract.CancelOwnershipHandover(&_PositionManager.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_PositionManager *PositionManagerTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _PositionManager.Contract.CancelOwnershipHandover(&_PositionManager.TransactOpts)
}

// CloseBidWall is a paid mutator transaction binding the contract method 0xad49d66e.
//
// Solidity: function closeBidWall((address,address,uint24,int24,address) _key) returns()
func (_PositionManager *PositionManagerTransactor) CloseBidWall(opts *bind.TransactOpts, _key PoolKey) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "closeBidWall", _key)
}

// CloseBidWall is a paid mutator transaction binding the contract method 0xad49d66e.
//
// Solidity: function closeBidWall((address,address,uint24,int24,address) _key) returns()
func (_PositionManager *PositionManagerSession) CloseBidWall(_key PoolKey) (*types.Transaction, error) {
	return _PositionManager.Contract.CloseBidWall(&_PositionManager.TransactOpts, _key)
}

// CloseBidWall is a paid mutator transaction binding the contract method 0xad49d66e.
//
// Solidity: function closeBidWall((address,address,uint24,int24,address) _key) returns()
func (_PositionManager *PositionManagerTransactorSession) CloseBidWall(_key PoolKey) (*types.Transaction, error) {
	return _PositionManager.Contract.CloseBidWall(&_PositionManager.TransactOpts, _key)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_PositionManager *PositionManagerTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_PositionManager *PositionManagerSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.CompleteOwnershipHandover(&_PositionManager.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_PositionManager *PositionManagerTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.CompleteOwnershipHandover(&_PositionManager.TransactOpts, pendingOwner)
}

// Flaunch is a paid mutator transaction binding the contract method 0x2a12e488.
//
// Solidity: function flaunch((string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params) payable returns(address memecoin_)
func (_PositionManager *PositionManagerTransactor) Flaunch(opts *bind.TransactOpts, _params PositionManagerFlaunchParams) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "flaunch", _params)
}

// Flaunch is a paid mutator transaction binding the contract method 0x2a12e488.
//
// Solidity: function flaunch((string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params) payable returns(address memecoin_)
func (_PositionManager *PositionManagerSession) Flaunch(_params PositionManagerFlaunchParams) (*types.Transaction, error) {
	return _PositionManager.Contract.Flaunch(&_PositionManager.TransactOpts, _params)
}

// Flaunch is a paid mutator transaction binding the contract method 0x2a12e488.
//
// Solidity: function flaunch((string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params) payable returns(address memecoin_)
func (_PositionManager *PositionManagerTransactorSession) Flaunch(_params PositionManagerFlaunchParams) (*types.Transaction, error) {
	return _PositionManager.Contract.Flaunch(&_PositionManager.TransactOpts, _params)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_PositionManager *PositionManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_PositionManager *PositionManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PositionManager.Contract.RenounceOwnership(&_PositionManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_PositionManager *PositionManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PositionManager.Contract.RenounceOwnership(&_PositionManager.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_PositionManager *PositionManagerTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_PositionManager *PositionManagerSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _PositionManager.Contract.RequestOwnershipHandover(&_PositionManager.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_PositionManager *PositionManagerTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _PositionManager.Contract.RequestOwnershipHandover(&_PositionManager.TransactOpts)
}

// SetFairLaunchFeeCalculator is a paid mutator transaction binding the contract method 0x38b1e700.
//
// Solidity: function setFairLaunchFeeCalculator(address _feeCalculator) returns()
func (_PositionManager *PositionManagerTransactor) SetFairLaunchFeeCalculator(opts *bind.TransactOpts, _feeCalculator common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setFairLaunchFeeCalculator", _feeCalculator)
}

// SetFairLaunchFeeCalculator is a paid mutator transaction binding the contract method 0x38b1e700.
//
// Solidity: function setFairLaunchFeeCalculator(address _feeCalculator) returns()
func (_PositionManager *PositionManagerSession) SetFairLaunchFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetFairLaunchFeeCalculator(&_PositionManager.TransactOpts, _feeCalculator)
}

// SetFairLaunchFeeCalculator is a paid mutator transaction binding the contract method 0x38b1e700.
//
// Solidity: function setFairLaunchFeeCalculator(address _feeCalculator) returns()
func (_PositionManager *PositionManagerTransactorSession) SetFairLaunchFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetFairLaunchFeeCalculator(&_PositionManager.TransactOpts, _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_PositionManager *PositionManagerTransactor) SetFeeCalculator(opts *bind.TransactOpts, _feeCalculator common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setFeeCalculator", _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_PositionManager *PositionManagerSession) SetFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetFeeCalculator(&_PositionManager.TransactOpts, _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_PositionManager *PositionManagerTransactorSession) SetFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetFeeCalculator(&_PositionManager.TransactOpts, _feeCalculator)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0xd3bff717.
//
// Solidity: function setFeeDistribution((uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_PositionManager *PositionManagerTransactor) SetFeeDistribution(opts *bind.TransactOpts, _feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setFeeDistribution", _feeDistribution)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0xd3bff717.
//
// Solidity: function setFeeDistribution((uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_PositionManager *PositionManagerSession) SetFeeDistribution(_feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _PositionManager.Contract.SetFeeDistribution(&_PositionManager.TransactOpts, _feeDistribution)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0xd3bff717.
//
// Solidity: function setFeeDistribution((uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_PositionManager *PositionManagerTransactorSession) SetFeeDistribution(_feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _PositionManager.Contract.SetFeeDistribution(&_PositionManager.TransactOpts, _feeDistribution)
}

// SetFlaunch is a paid mutator transaction binding the contract method 0xbe721505.
//
// Solidity: function setFlaunch(address _flaunchContract) returns()
func (_PositionManager *PositionManagerTransactor) SetFlaunch(opts *bind.TransactOpts, _flaunchContract common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setFlaunch", _flaunchContract)
}

// SetFlaunch is a paid mutator transaction binding the contract method 0xbe721505.
//
// Solidity: function setFlaunch(address _flaunchContract) returns()
func (_PositionManager *PositionManagerSession) SetFlaunch(_flaunchContract common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetFlaunch(&_PositionManager.TransactOpts, _flaunchContract)
}

// SetFlaunch is a paid mutator transaction binding the contract method 0xbe721505.
//
// Solidity: function setFlaunch(address _flaunchContract) returns()
func (_PositionManager *PositionManagerTransactorSession) SetFlaunch(_flaunchContract common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetFlaunch(&_PositionManager.TransactOpts, _flaunchContract)
}

// SetInitialPrice is a paid mutator transaction binding the contract method 0xdf81740e.
//
// Solidity: function setInitialPrice(address _initialPrice) returns()
func (_PositionManager *PositionManagerTransactor) SetInitialPrice(opts *bind.TransactOpts, _initialPrice common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setInitialPrice", _initialPrice)
}

// SetInitialPrice is a paid mutator transaction binding the contract method 0xdf81740e.
//
// Solidity: function setInitialPrice(address _initialPrice) returns()
func (_PositionManager *PositionManagerSession) SetInitialPrice(_initialPrice common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetInitialPrice(&_PositionManager.TransactOpts, _initialPrice)
}

// SetInitialPrice is a paid mutator transaction binding the contract method 0xdf81740e.
//
// Solidity: function setInitialPrice(address _initialPrice) returns()
func (_PositionManager *PositionManagerTransactorSession) SetInitialPrice(_initialPrice common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetInitialPrice(&_PositionManager.TransactOpts, _initialPrice)
}

// SetPoolFeeDistribution is a paid mutator transaction binding the contract method 0x4ed0f0f5.
//
// Solidity: function setPoolFeeDistribution(bytes32 _poolId, (uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_PositionManager *PositionManagerTransactor) SetPoolFeeDistribution(opts *bind.TransactOpts, _poolId [32]byte, _feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setPoolFeeDistribution", _poolId, _feeDistribution)
}

// SetPoolFeeDistribution is a paid mutator transaction binding the contract method 0x4ed0f0f5.
//
// Solidity: function setPoolFeeDistribution(bytes32 _poolId, (uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_PositionManager *PositionManagerSession) SetPoolFeeDistribution(_poolId [32]byte, _feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _PositionManager.Contract.SetPoolFeeDistribution(&_PositionManager.TransactOpts, _poolId, _feeDistribution)
}

// SetPoolFeeDistribution is a paid mutator transaction binding the contract method 0x4ed0f0f5.
//
// Solidity: function setPoolFeeDistribution(bytes32 _poolId, (uint24,uint24,uint24,bool) _feeDistribution) returns()
func (_PositionManager *PositionManagerTransactorSession) SetPoolFeeDistribution(_poolId [32]byte, _feeDistribution FeeDistributorFeeDistribution) (*types.Transaction, error) {
	return _PositionManager.Contract.SetPoolFeeDistribution(&_PositionManager.TransactOpts, _poolId, _feeDistribution)
}

// SetProtocolFeeDistribution is a paid mutator transaction binding the contract method 0x2423028c.
//
// Solidity: function setProtocolFeeDistribution(uint24 _protocol) returns()
func (_PositionManager *PositionManagerTransactor) SetProtocolFeeDistribution(opts *bind.TransactOpts, _protocol *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setProtocolFeeDistribution", _protocol)
}

// SetProtocolFeeDistribution is a paid mutator transaction binding the contract method 0x2423028c.
//
// Solidity: function setProtocolFeeDistribution(uint24 _protocol) returns()
func (_PositionManager *PositionManagerSession) SetProtocolFeeDistribution(_protocol *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SetProtocolFeeDistribution(&_PositionManager.TransactOpts, _protocol)
}

// SetProtocolFeeDistribution is a paid mutator transaction binding the contract method 0x2423028c.
//
// Solidity: function setProtocolFeeDistribution(uint24 _protocol) returns()
func (_PositionManager *PositionManagerTransactorSession) SetProtocolFeeDistribution(_protocol *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SetProtocolFeeDistribution(&_PositionManager.TransactOpts, _protocol)
}

// SetReferralEscrow is a paid mutator transaction binding the contract method 0xa87277dd.
//
// Solidity: function setReferralEscrow(address _referralEscrow) returns()
func (_PositionManager *PositionManagerTransactor) SetReferralEscrow(opts *bind.TransactOpts, _referralEscrow common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setReferralEscrow", _referralEscrow)
}

// SetReferralEscrow is a paid mutator transaction binding the contract method 0xa87277dd.
//
// Solidity: function setReferralEscrow(address _referralEscrow) returns()
func (_PositionManager *PositionManagerSession) SetReferralEscrow(_referralEscrow common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetReferralEscrow(&_PositionManager.TransactOpts, _referralEscrow)
}

// SetReferralEscrow is a paid mutator transaction binding the contract method 0xa87277dd.
//
// Solidity: function setReferralEscrow(address _referralEscrow) returns()
func (_PositionManager *PositionManagerTransactorSession) SetReferralEscrow(_referralEscrow common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetReferralEscrow(&_PositionManager.TransactOpts, _referralEscrow)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_PositionManager *PositionManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_PositionManager *PositionManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.TransferOwnership(&_PositionManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_PositionManager *PositionManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.TransferOwnership(&_PositionManager.TransactOpts, newOwner)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_PositionManager *PositionManagerTransactor) UnlockCallback(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "unlockCallback", data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_PositionManager *PositionManagerSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.UnlockCallback(&_PositionManager.TransactOpts, data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_PositionManager *PositionManagerTransactorSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _PositionManager.Contract.UnlockCallback(&_PositionManager.TransactOpts, data)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x4c2d94c0.
//
// Solidity: function withdrawFees(address _recipient, bool _unwrap) returns()
func (_PositionManager *PositionManagerTransactor) WithdrawFees(opts *bind.TransactOpts, _recipient common.Address, _unwrap bool) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "withdrawFees", _recipient, _unwrap)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x4c2d94c0.
//
// Solidity: function withdrawFees(address _recipient, bool _unwrap) returns()
func (_PositionManager *PositionManagerSession) WithdrawFees(_recipient common.Address, _unwrap bool) (*types.Transaction, error) {
	return _PositionManager.Contract.WithdrawFees(&_PositionManager.TransactOpts, _recipient, _unwrap)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x4c2d94c0.
//
// Solidity: function withdrawFees(address _recipient, bool _unwrap) returns()
func (_PositionManager *PositionManagerTransactorSession) WithdrawFees(_recipient common.Address, _unwrap bool) (*types.Transaction, error) {
	return _PositionManager.Contract.WithdrawFees(&_PositionManager.TransactOpts, _recipient, _unwrap)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PositionManager *PositionManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PositionManager *PositionManagerSession) Receive() (*types.Transaction, error) {
	return _PositionManager.Contract.Receive(&_PositionManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PositionManager *PositionManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _PositionManager.Contract.Receive(&_PositionManager.TransactOpts)
}

// PositionManagerCreatorFeeAllocationUpdatedIterator is returned from FilterCreatorFeeAllocationUpdated and is used to iterate over the raw logs and unpacked data for CreatorFeeAllocationUpdated events raised by the PositionManager contract.
type PositionManagerCreatorFeeAllocationUpdatedIterator struct {
	Event *PositionManagerCreatorFeeAllocationUpdated // Event containing the contract specifics and raw log

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
func (it *PositionManagerCreatorFeeAllocationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerCreatorFeeAllocationUpdated)
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
		it.Event = new(PositionManagerCreatorFeeAllocationUpdated)
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
func (it *PositionManagerCreatorFeeAllocationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerCreatorFeeAllocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerCreatorFeeAllocationUpdated represents a CreatorFeeAllocationUpdated event raised by the PositionManager contract.
type PositionManagerCreatorFeeAllocationUpdated struct {
	PoolId     [32]byte
	Allocation *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatorFeeAllocationUpdated is a free log retrieval operation binding the contract event 0x6d9755b9ab4c1e27104da15a9b4589ffe57a1222d13193f2f1354ffde5fee10e.
//
// Solidity: event CreatorFeeAllocationUpdated(bytes32 indexed _poolId, uint24 _allocation)
func (_PositionManager *PositionManagerFilterer) FilterCreatorFeeAllocationUpdated(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerCreatorFeeAllocationUpdatedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "CreatorFeeAllocationUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerCreatorFeeAllocationUpdatedIterator{contract: _PositionManager.contract, event: "CreatorFeeAllocationUpdated", logs: logs, sub: sub}, nil
}

// WatchCreatorFeeAllocationUpdated is a free log subscription operation binding the contract event 0x6d9755b9ab4c1e27104da15a9b4589ffe57a1222d13193f2f1354ffde5fee10e.
//
// Solidity: event CreatorFeeAllocationUpdated(bytes32 indexed _poolId, uint24 _allocation)
func (_PositionManager *PositionManagerFilterer) WatchCreatorFeeAllocationUpdated(opts *bind.WatchOpts, sink chan<- *PositionManagerCreatorFeeAllocationUpdated, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "CreatorFeeAllocationUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerCreatorFeeAllocationUpdated)
				if err := _PositionManager.contract.UnpackLog(event, "CreatorFeeAllocationUpdated", log); err != nil {
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

// ParseCreatorFeeAllocationUpdated is a log parse operation binding the contract event 0x6d9755b9ab4c1e27104da15a9b4589ffe57a1222d13193f2f1354ffde5fee10e.
//
// Solidity: event CreatorFeeAllocationUpdated(bytes32 indexed _poolId, uint24 _allocation)
func (_PositionManager *PositionManagerFilterer) ParseCreatorFeeAllocationUpdated(log types.Log) (*PositionManagerCreatorFeeAllocationUpdated, error) {
	event := new(PositionManagerCreatorFeeAllocationUpdated)
	if err := _PositionManager.contract.UnpackLog(event, "CreatorFeeAllocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PositionManager contract.
type PositionManagerDepositIterator struct {
	Event *PositionManagerDeposit // Event containing the contract specifics and raw log

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
func (it *PositionManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerDeposit)
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
		it.Event = new(PositionManagerDeposit)
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
func (it *PositionManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerDeposit represents a Deposit event raised by the PositionManager contract.
type PositionManagerDeposit struct {
	PoolId [32]byte
	Payee  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xc95ddcaddf83340b68d0d44c01b1703f5d28d0611a3fd87e69d79ba7e2ac21d3.
//
// Solidity: event Deposit(bytes32 indexed _poolId, address _payee, address _token, uint256 _amount)
func (_PositionManager *PositionManagerFilterer) FilterDeposit(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerDepositIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "Deposit", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerDepositIterator{contract: _PositionManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xc95ddcaddf83340b68d0d44c01b1703f5d28d0611a3fd87e69d79ba7e2ac21d3.
//
// Solidity: event Deposit(bytes32 indexed _poolId, address _payee, address _token, uint256 _amount)
func (_PositionManager *PositionManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PositionManagerDeposit, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "Deposit", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerDeposit)
				if err := _PositionManager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xc95ddcaddf83340b68d0d44c01b1703f5d28d0611a3fd87e69d79ba7e2ac21d3.
//
// Solidity: event Deposit(bytes32 indexed _poolId, address _payee, address _token, uint256 _amount)
func (_PositionManager *PositionManagerFilterer) ParseDeposit(log types.Log) (*PositionManagerDeposit, error) {
	event := new(PositionManagerDeposit)
	if err := _PositionManager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerFairLaunchFeeCalculatorUpdatedIterator is returned from FilterFairLaunchFeeCalculatorUpdated and is used to iterate over the raw logs and unpacked data for FairLaunchFeeCalculatorUpdated events raised by the PositionManager contract.
type PositionManagerFairLaunchFeeCalculatorUpdatedIterator struct {
	Event *PositionManagerFairLaunchFeeCalculatorUpdated // Event containing the contract specifics and raw log

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
func (it *PositionManagerFairLaunchFeeCalculatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerFairLaunchFeeCalculatorUpdated)
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
		it.Event = new(PositionManagerFairLaunchFeeCalculatorUpdated)
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
func (it *PositionManagerFairLaunchFeeCalculatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerFairLaunchFeeCalculatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerFairLaunchFeeCalculatorUpdated represents a FairLaunchFeeCalculatorUpdated event raised by the PositionManager contract.
type PositionManagerFairLaunchFeeCalculatorUpdated struct {
	FeeCalculator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFairLaunchFeeCalculatorUpdated is a free log retrieval operation binding the contract event 0x87043577396d39ef835a9eb69bb496c219cc61bdd6e718447add3c06b6cc0844.
//
// Solidity: event FairLaunchFeeCalculatorUpdated(address _feeCalculator)
func (_PositionManager *PositionManagerFilterer) FilterFairLaunchFeeCalculatorUpdated(opts *bind.FilterOpts) (*PositionManagerFairLaunchFeeCalculatorUpdatedIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "FairLaunchFeeCalculatorUpdated")
	if err != nil {
		return nil, err
	}
	return &PositionManagerFairLaunchFeeCalculatorUpdatedIterator{contract: _PositionManager.contract, event: "FairLaunchFeeCalculatorUpdated", logs: logs, sub: sub}, nil
}

// WatchFairLaunchFeeCalculatorUpdated is a free log subscription operation binding the contract event 0x87043577396d39ef835a9eb69bb496c219cc61bdd6e718447add3c06b6cc0844.
//
// Solidity: event FairLaunchFeeCalculatorUpdated(address _feeCalculator)
func (_PositionManager *PositionManagerFilterer) WatchFairLaunchFeeCalculatorUpdated(opts *bind.WatchOpts, sink chan<- *PositionManagerFairLaunchFeeCalculatorUpdated) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "FairLaunchFeeCalculatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerFairLaunchFeeCalculatorUpdated)
				if err := _PositionManager.contract.UnpackLog(event, "FairLaunchFeeCalculatorUpdated", log); err != nil {
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

// ParseFairLaunchFeeCalculatorUpdated is a log parse operation binding the contract event 0x87043577396d39ef835a9eb69bb496c219cc61bdd6e718447add3c06b6cc0844.
//
// Solidity: event FairLaunchFeeCalculatorUpdated(address _feeCalculator)
func (_PositionManager *PositionManagerFilterer) ParseFairLaunchFeeCalculatorUpdated(log types.Log) (*PositionManagerFairLaunchFeeCalculatorUpdated, error) {
	event := new(PositionManagerFairLaunchFeeCalculatorUpdated)
	if err := _PositionManager.contract.UnpackLog(event, "FairLaunchFeeCalculatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerFeeCalculatorUpdatedIterator is returned from FilterFeeCalculatorUpdated and is used to iterate over the raw logs and unpacked data for FeeCalculatorUpdated events raised by the PositionManager contract.
type PositionManagerFeeCalculatorUpdatedIterator struct {
	Event *PositionManagerFeeCalculatorUpdated // Event containing the contract specifics and raw log

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
func (it *PositionManagerFeeCalculatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerFeeCalculatorUpdated)
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
		it.Event = new(PositionManagerFeeCalculatorUpdated)
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
func (it *PositionManagerFeeCalculatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerFeeCalculatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerFeeCalculatorUpdated represents a FeeCalculatorUpdated event raised by the PositionManager contract.
type PositionManagerFeeCalculatorUpdated struct {
	FeeCalculator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeCalculatorUpdated is a free log retrieval operation binding the contract event 0x3e762c7e655633ce63121393b9694f9ca1883d14d18f48f1be55e5dc7a9fb6c1.
//
// Solidity: event FeeCalculatorUpdated(address _feeCalculator)
func (_PositionManager *PositionManagerFilterer) FilterFeeCalculatorUpdated(opts *bind.FilterOpts) (*PositionManagerFeeCalculatorUpdatedIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "FeeCalculatorUpdated")
	if err != nil {
		return nil, err
	}
	return &PositionManagerFeeCalculatorUpdatedIterator{contract: _PositionManager.contract, event: "FeeCalculatorUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeCalculatorUpdated is a free log subscription operation binding the contract event 0x3e762c7e655633ce63121393b9694f9ca1883d14d18f48f1be55e5dc7a9fb6c1.
//
// Solidity: event FeeCalculatorUpdated(address _feeCalculator)
func (_PositionManager *PositionManagerFilterer) WatchFeeCalculatorUpdated(opts *bind.WatchOpts, sink chan<- *PositionManagerFeeCalculatorUpdated) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "FeeCalculatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerFeeCalculatorUpdated)
				if err := _PositionManager.contract.UnpackLog(event, "FeeCalculatorUpdated", log); err != nil {
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

// ParseFeeCalculatorUpdated is a log parse operation binding the contract event 0x3e762c7e655633ce63121393b9694f9ca1883d14d18f48f1be55e5dc7a9fb6c1.
//
// Solidity: event FeeCalculatorUpdated(address _feeCalculator)
func (_PositionManager *PositionManagerFilterer) ParseFeeCalculatorUpdated(log types.Log) (*PositionManagerFeeCalculatorUpdated, error) {
	event := new(PositionManagerFeeCalculatorUpdated)
	if err := _PositionManager.contract.UnpackLog(event, "FeeCalculatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerFeeDistributionUpdatedIterator is returned from FilterFeeDistributionUpdated and is used to iterate over the raw logs and unpacked data for FeeDistributionUpdated events raised by the PositionManager contract.
type PositionManagerFeeDistributionUpdatedIterator struct {
	Event *PositionManagerFeeDistributionUpdated // Event containing the contract specifics and raw log

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
func (it *PositionManagerFeeDistributionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerFeeDistributionUpdated)
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
		it.Event = new(PositionManagerFeeDistributionUpdated)
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
func (it *PositionManagerFeeDistributionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerFeeDistributionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerFeeDistributionUpdated represents a FeeDistributionUpdated event raised by the PositionManager contract.
type PositionManagerFeeDistributionUpdated struct {
	FeeDistribution FeeDistributorFeeDistribution
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributionUpdated is a free log retrieval operation binding the contract event 0xca7619d8d03127c99378decb6d4c002b5b8b8d60918ea28c2300a7811e7c1300.
//
// Solidity: event FeeDistributionUpdated((uint24,uint24,uint24,bool) _feeDistribution)
func (_PositionManager *PositionManagerFilterer) FilterFeeDistributionUpdated(opts *bind.FilterOpts) (*PositionManagerFeeDistributionUpdatedIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "FeeDistributionUpdated")
	if err != nil {
		return nil, err
	}
	return &PositionManagerFeeDistributionUpdatedIterator{contract: _PositionManager.contract, event: "FeeDistributionUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeDistributionUpdated is a free log subscription operation binding the contract event 0xca7619d8d03127c99378decb6d4c002b5b8b8d60918ea28c2300a7811e7c1300.
//
// Solidity: event FeeDistributionUpdated((uint24,uint24,uint24,bool) _feeDistribution)
func (_PositionManager *PositionManagerFilterer) WatchFeeDistributionUpdated(opts *bind.WatchOpts, sink chan<- *PositionManagerFeeDistributionUpdated) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "FeeDistributionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerFeeDistributionUpdated)
				if err := _PositionManager.contract.UnpackLog(event, "FeeDistributionUpdated", log); err != nil {
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

// ParseFeeDistributionUpdated is a log parse operation binding the contract event 0xca7619d8d03127c99378decb6d4c002b5b8b8d60918ea28c2300a7811e7c1300.
//
// Solidity: event FeeDistributionUpdated((uint24,uint24,uint24,bool) _feeDistribution)
func (_PositionManager *PositionManagerFilterer) ParseFeeDistributionUpdated(log types.Log) (*PositionManagerFeeDistributionUpdated, error) {
	event := new(PositionManagerFeeDistributionUpdated)
	if err := _PositionManager.contract.UnpackLog(event, "FeeDistributionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerInitialPriceUpdatedIterator is returned from FilterInitialPriceUpdated and is used to iterate over the raw logs and unpacked data for InitialPriceUpdated events raised by the PositionManager contract.
type PositionManagerInitialPriceUpdatedIterator struct {
	Event *PositionManagerInitialPriceUpdated // Event containing the contract specifics and raw log

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
func (it *PositionManagerInitialPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerInitialPriceUpdated)
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
		it.Event = new(PositionManagerInitialPriceUpdated)
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
func (it *PositionManagerInitialPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerInitialPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerInitialPriceUpdated represents a InitialPriceUpdated event raised by the PositionManager contract.
type PositionManagerInitialPriceUpdated struct {
	InitialPrice common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialPriceUpdated is a free log retrieval operation binding the contract event 0xbfad4f58ed4ffa32358e2ae423bf8a56d8adf778ad13eb9160c0fecd99ec1577.
//
// Solidity: event InitialPriceUpdated(address _initialPrice)
func (_PositionManager *PositionManagerFilterer) FilterInitialPriceUpdated(opts *bind.FilterOpts) (*PositionManagerInitialPriceUpdatedIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "InitialPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &PositionManagerInitialPriceUpdatedIterator{contract: _PositionManager.contract, event: "InitialPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchInitialPriceUpdated is a free log subscription operation binding the contract event 0xbfad4f58ed4ffa32358e2ae423bf8a56d8adf778ad13eb9160c0fecd99ec1577.
//
// Solidity: event InitialPriceUpdated(address _initialPrice)
func (_PositionManager *PositionManagerFilterer) WatchInitialPriceUpdated(opts *bind.WatchOpts, sink chan<- *PositionManagerInitialPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "InitialPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerInitialPriceUpdated)
				if err := _PositionManager.contract.UnpackLog(event, "InitialPriceUpdated", log); err != nil {
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

// ParseInitialPriceUpdated is a log parse operation binding the contract event 0xbfad4f58ed4ffa32358e2ae423bf8a56d8adf778ad13eb9160c0fecd99ec1577.
//
// Solidity: event InitialPriceUpdated(address _initialPrice)
func (_PositionManager *PositionManagerFilterer) ParseInitialPriceUpdated(log types.Log) (*PositionManagerInitialPriceUpdated, error) {
	event := new(PositionManagerInitialPriceUpdated)
	if err := _PositionManager.contract.UnpackLog(event, "InitialPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the PositionManager contract.
type PositionManagerOwnershipHandoverCanceledIterator struct {
	Event *PositionManagerOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *PositionManagerOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerOwnershipHandoverCanceled)
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
		it.Event = new(PositionManagerOwnershipHandoverCanceled)
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
func (it *PositionManagerOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the PositionManager contract.
type PositionManagerOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_PositionManager *PositionManagerFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*PositionManagerOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerOwnershipHandoverCanceledIterator{contract: _PositionManager.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_PositionManager *PositionManagerFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *PositionManagerOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerOwnershipHandoverCanceled)
				if err := _PositionManager.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_PositionManager *PositionManagerFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*PositionManagerOwnershipHandoverCanceled, error) {
	event := new(PositionManagerOwnershipHandoverCanceled)
	if err := _PositionManager.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the PositionManager contract.
type PositionManagerOwnershipHandoverRequestedIterator struct {
	Event *PositionManagerOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *PositionManagerOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerOwnershipHandoverRequested)
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
		it.Event = new(PositionManagerOwnershipHandoverRequested)
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
func (it *PositionManagerOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the PositionManager contract.
type PositionManagerOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_PositionManager *PositionManagerFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*PositionManagerOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerOwnershipHandoverRequestedIterator{contract: _PositionManager.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_PositionManager *PositionManagerFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *PositionManagerOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerOwnershipHandoverRequested)
				if err := _PositionManager.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_PositionManager *PositionManagerFilterer) ParseOwnershipHandoverRequested(log types.Log) (*PositionManagerOwnershipHandoverRequested, error) {
	event := new(PositionManagerOwnershipHandoverRequested)
	if err := _PositionManager.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PositionManager contract.
type PositionManagerOwnershipTransferredIterator struct {
	Event *PositionManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PositionManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerOwnershipTransferred)
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
		it.Event = new(PositionManagerOwnershipTransferred)
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
func (it *PositionManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerOwnershipTransferred represents a OwnershipTransferred event raised by the PositionManager contract.
type PositionManagerOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PositionManager *PositionManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PositionManagerOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerOwnershipTransferredIterator{contract: _PositionManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PositionManager *PositionManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PositionManagerOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerOwnershipTransferred)
				if err := _PositionManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PositionManager *PositionManagerFilterer) ParseOwnershipTransferred(log types.Log) (*PositionManagerOwnershipTransferred, error) {
	event := new(PositionManagerOwnershipTransferred)
	if err := _PositionManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the PositionManager contract.
type PositionManagerPoolCreatedIterator struct {
	Event *PositionManagerPoolCreated // Event containing the contract specifics and raw log

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
func (it *PositionManagerPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerPoolCreated)
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
		it.Event = new(PositionManagerPoolCreated)
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
func (it *PositionManagerPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerPoolCreated represents a PoolCreated event raised by the PositionManager contract.
type PositionManagerPoolCreated struct {
	PoolId           [32]byte
	Memecoin         common.Address
	MemecoinTreasury common.Address
	TokenId          *big.Int
	CurrencyFlipped  bool
	FlaunchFee       *big.Int
	Params           PositionManagerFlaunchParams
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xe9a023154a0e1bd430ba68aafea07b09c78a0e5406c3696fb3c0dd631fa53b64.
//
// Solidity: event PoolCreated(bytes32 indexed _poolId, address _memecoin, address _memecoinTreasury, uint256 _tokenId, bool _currencyFlipped, uint256 _flaunchFee, (string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params)
func (_PositionManager *PositionManagerFilterer) FilterPoolCreated(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerPoolCreatedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "PoolCreated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerPoolCreatedIterator{contract: _PositionManager.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xe9a023154a0e1bd430ba68aafea07b09c78a0e5406c3696fb3c0dd631fa53b64.
//
// Solidity: event PoolCreated(bytes32 indexed _poolId, address _memecoin, address _memecoinTreasury, uint256 _tokenId, bool _currencyFlipped, uint256 _flaunchFee, (string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params)
func (_PositionManager *PositionManagerFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *PositionManagerPoolCreated, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "PoolCreated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerPoolCreated)
				if err := _PositionManager.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0xe9a023154a0e1bd430ba68aafea07b09c78a0e5406c3696fb3c0dd631fa53b64.
//
// Solidity: event PoolCreated(bytes32 indexed _poolId, address _memecoin, address _memecoinTreasury, uint256 _tokenId, bool _currencyFlipped, uint256 _flaunchFee, (string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params)
func (_PositionManager *PositionManagerFilterer) ParsePoolCreated(log types.Log) (*PositionManagerPoolCreated, error) {
	event := new(PositionManagerPoolCreated)
	if err := _PositionManager.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerPoolFeeDistributionUpdatedIterator is returned from FilterPoolFeeDistributionUpdated and is used to iterate over the raw logs and unpacked data for PoolFeeDistributionUpdated events raised by the PositionManager contract.
type PositionManagerPoolFeeDistributionUpdatedIterator struct {
	Event *PositionManagerPoolFeeDistributionUpdated // Event containing the contract specifics and raw log

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
func (it *PositionManagerPoolFeeDistributionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerPoolFeeDistributionUpdated)
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
		it.Event = new(PositionManagerPoolFeeDistributionUpdated)
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
func (it *PositionManagerPoolFeeDistributionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerPoolFeeDistributionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerPoolFeeDistributionUpdated represents a PoolFeeDistributionUpdated event raised by the PositionManager contract.
type PositionManagerPoolFeeDistributionUpdated struct {
	PoolId          [32]byte
	FeeDistribution FeeDistributorFeeDistribution
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPoolFeeDistributionUpdated is a free log retrieval operation binding the contract event 0xe1b2af81a774e1ebee3ca7c94c1e1a0df1210b236149a2079b8bba8dbc475a28.
//
// Solidity: event PoolFeeDistributionUpdated(bytes32 indexed _poolId, (uint24,uint24,uint24,bool) _feeDistribution)
func (_PositionManager *PositionManagerFilterer) FilterPoolFeeDistributionUpdated(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerPoolFeeDistributionUpdatedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "PoolFeeDistributionUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerPoolFeeDistributionUpdatedIterator{contract: _PositionManager.contract, event: "PoolFeeDistributionUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolFeeDistributionUpdated is a free log subscription operation binding the contract event 0xe1b2af81a774e1ebee3ca7c94c1e1a0df1210b236149a2079b8bba8dbc475a28.
//
// Solidity: event PoolFeeDistributionUpdated(bytes32 indexed _poolId, (uint24,uint24,uint24,bool) _feeDistribution)
func (_PositionManager *PositionManagerFilterer) WatchPoolFeeDistributionUpdated(opts *bind.WatchOpts, sink chan<- *PositionManagerPoolFeeDistributionUpdated, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "PoolFeeDistributionUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerPoolFeeDistributionUpdated)
				if err := _PositionManager.contract.UnpackLog(event, "PoolFeeDistributionUpdated", log); err != nil {
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

// ParsePoolFeeDistributionUpdated is a log parse operation binding the contract event 0xe1b2af81a774e1ebee3ca7c94c1e1a0df1210b236149a2079b8bba8dbc475a28.
//
// Solidity: event PoolFeeDistributionUpdated(bytes32 indexed _poolId, (uint24,uint24,uint24,bool) _feeDistribution)
func (_PositionManager *PositionManagerFilterer) ParsePoolFeeDistributionUpdated(log types.Log) (*PositionManagerPoolFeeDistributionUpdated, error) {
	event := new(PositionManagerPoolFeeDistributionUpdated)
	if err := _PositionManager.contract.UnpackLog(event, "PoolFeeDistributionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerPoolFeesDistributedIterator is returned from FilterPoolFeesDistributed and is used to iterate over the raw logs and unpacked data for PoolFeesDistributed events raised by the PositionManager contract.
type PositionManagerPoolFeesDistributedIterator struct {
	Event *PositionManagerPoolFeesDistributed // Event containing the contract specifics and raw log

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
func (it *PositionManagerPoolFeesDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerPoolFeesDistributed)
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
		it.Event = new(PositionManagerPoolFeesDistributed)
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
func (it *PositionManagerPoolFeesDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerPoolFeesDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerPoolFeesDistributed represents a PoolFeesDistributed event raised by the PositionManager contract.
type PositionManagerPoolFeesDistributed struct {
	PoolId           [32]byte
	DonateAmount     *big.Int
	CreatorAmount    *big.Int
	BidWallAmount    *big.Int
	GovernanceAmount *big.Int
	ProtocolAmount   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPoolFeesDistributed is a free log retrieval operation binding the contract event 0xc7241a69d3660bdfe5f36bdcca3b2da1fe8844366e46adb58be95171ab0665ad.
//
// Solidity: event PoolFeesDistributed(bytes32 indexed _poolId, uint256 _donateAmount, uint256 _creatorAmount, uint256 _bidWallAmount, uint256 _governanceAmount, uint256 _protocolAmount)
func (_PositionManager *PositionManagerFilterer) FilterPoolFeesDistributed(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerPoolFeesDistributedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "PoolFeesDistributed", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerPoolFeesDistributedIterator{contract: _PositionManager.contract, event: "PoolFeesDistributed", logs: logs, sub: sub}, nil
}

// WatchPoolFeesDistributed is a free log subscription operation binding the contract event 0xc7241a69d3660bdfe5f36bdcca3b2da1fe8844366e46adb58be95171ab0665ad.
//
// Solidity: event PoolFeesDistributed(bytes32 indexed _poolId, uint256 _donateAmount, uint256 _creatorAmount, uint256 _bidWallAmount, uint256 _governanceAmount, uint256 _protocolAmount)
func (_PositionManager *PositionManagerFilterer) WatchPoolFeesDistributed(opts *bind.WatchOpts, sink chan<- *PositionManagerPoolFeesDistributed, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "PoolFeesDistributed", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerPoolFeesDistributed)
				if err := _PositionManager.contract.UnpackLog(event, "PoolFeesDistributed", log); err != nil {
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

// ParsePoolFeesDistributed is a log parse operation binding the contract event 0xc7241a69d3660bdfe5f36bdcca3b2da1fe8844366e46adb58be95171ab0665ad.
//
// Solidity: event PoolFeesDistributed(bytes32 indexed _poolId, uint256 _donateAmount, uint256 _creatorAmount, uint256 _bidWallAmount, uint256 _governanceAmount, uint256 _protocolAmount)
func (_PositionManager *PositionManagerFilterer) ParsePoolFeesDistributed(log types.Log) (*PositionManagerPoolFeesDistributed, error) {
	event := new(PositionManagerPoolFeesDistributed)
	if err := _PositionManager.contract.UnpackLog(event, "PoolFeesDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerPoolFeesReceivedIterator is returned from FilterPoolFeesReceived and is used to iterate over the raw logs and unpacked data for PoolFeesReceived events raised by the PositionManager contract.
type PositionManagerPoolFeesReceivedIterator struct {
	Event *PositionManagerPoolFeesReceived // Event containing the contract specifics and raw log

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
func (it *PositionManagerPoolFeesReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerPoolFeesReceived)
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
		it.Event = new(PositionManagerPoolFeesReceived)
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
func (it *PositionManagerPoolFeesReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerPoolFeesReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerPoolFeesReceived represents a PoolFeesReceived event raised by the PositionManager contract.
type PositionManagerPoolFeesReceived struct {
	PoolId  [32]byte
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolFeesReceived is a free log retrieval operation binding the contract event 0xa245a9a38e8877add82f0a82c13e062ab3df16a26121977ddcca8827d46c690a.
//
// Solidity: event PoolFeesReceived(bytes32 indexed _poolId, uint256 _amount0, uint256 _amount1)
func (_PositionManager *PositionManagerFilterer) FilterPoolFeesReceived(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerPoolFeesReceivedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "PoolFeesReceived", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerPoolFeesReceivedIterator{contract: _PositionManager.contract, event: "PoolFeesReceived", logs: logs, sub: sub}, nil
}

// WatchPoolFeesReceived is a free log subscription operation binding the contract event 0xa245a9a38e8877add82f0a82c13e062ab3df16a26121977ddcca8827d46c690a.
//
// Solidity: event PoolFeesReceived(bytes32 indexed _poolId, uint256 _amount0, uint256 _amount1)
func (_PositionManager *PositionManagerFilterer) WatchPoolFeesReceived(opts *bind.WatchOpts, sink chan<- *PositionManagerPoolFeesReceived, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "PoolFeesReceived", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerPoolFeesReceived)
				if err := _PositionManager.contract.UnpackLog(event, "PoolFeesReceived", log); err != nil {
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

// ParsePoolFeesReceived is a log parse operation binding the contract event 0xa245a9a38e8877add82f0a82c13e062ab3df16a26121977ddcca8827d46c690a.
//
// Solidity: event PoolFeesReceived(bytes32 indexed _poolId, uint256 _amount0, uint256 _amount1)
func (_PositionManager *PositionManagerFilterer) ParsePoolFeesReceived(log types.Log) (*PositionManagerPoolFeesReceived, error) {
	event := new(PositionManagerPoolFeesReceived)
	if err := _PositionManager.contract.UnpackLog(event, "PoolFeesReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerPoolFeesSwappedIterator is returned from FilterPoolFeesSwapped and is used to iterate over the raw logs and unpacked data for PoolFeesSwapped events raised by the PositionManager contract.
type PositionManagerPoolFeesSwappedIterator struct {
	Event *PositionManagerPoolFeesSwapped // Event containing the contract specifics and raw log

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
func (it *PositionManagerPoolFeesSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerPoolFeesSwapped)
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
		it.Event = new(PositionManagerPoolFeesSwapped)
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
func (it *PositionManagerPoolFeesSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerPoolFeesSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerPoolFeesSwapped represents a PoolFeesSwapped event raised by the PositionManager contract.
type PositionManagerPoolFeesSwapped struct {
	PoolId     [32]byte
	ZeroForOne bool
	Amount0    *big.Int
	Amount1    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolFeesSwapped is a free log retrieval operation binding the contract event 0xce97caf4fd0295de9544b52b4b9e79fe34c370bebb6fb71bc5baae9a70207968.
//
// Solidity: event PoolFeesSwapped(bytes32 indexed _poolId, bool zeroForOne, uint256 _amount0, uint256 _amount1)
func (_PositionManager *PositionManagerFilterer) FilterPoolFeesSwapped(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerPoolFeesSwappedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "PoolFeesSwapped", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerPoolFeesSwappedIterator{contract: _PositionManager.contract, event: "PoolFeesSwapped", logs: logs, sub: sub}, nil
}

// WatchPoolFeesSwapped is a free log subscription operation binding the contract event 0xce97caf4fd0295de9544b52b4b9e79fe34c370bebb6fb71bc5baae9a70207968.
//
// Solidity: event PoolFeesSwapped(bytes32 indexed _poolId, bool zeroForOne, uint256 _amount0, uint256 _amount1)
func (_PositionManager *PositionManagerFilterer) WatchPoolFeesSwapped(opts *bind.WatchOpts, sink chan<- *PositionManagerPoolFeesSwapped, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "PoolFeesSwapped", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerPoolFeesSwapped)
				if err := _PositionManager.contract.UnpackLog(event, "PoolFeesSwapped", log); err != nil {
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

// ParsePoolFeesSwapped is a log parse operation binding the contract event 0xce97caf4fd0295de9544b52b4b9e79fe34c370bebb6fb71bc5baae9a70207968.
//
// Solidity: event PoolFeesSwapped(bytes32 indexed _poolId, bool zeroForOne, uint256 _amount0, uint256 _amount1)
func (_PositionManager *PositionManagerFilterer) ParsePoolFeesSwapped(log types.Log) (*PositionManagerPoolFeesSwapped, error) {
	event := new(PositionManagerPoolFeesSwapped)
	if err := _PositionManager.contract.UnpackLog(event, "PoolFeesSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerPoolPremineIterator is returned from FilterPoolPremine and is used to iterate over the raw logs and unpacked data for PoolPremine events raised by the PositionManager contract.
type PositionManagerPoolPremineIterator struct {
	Event *PositionManagerPoolPremine // Event containing the contract specifics and raw log

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
func (it *PositionManagerPoolPremineIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerPoolPremine)
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
		it.Event = new(PositionManagerPoolPremine)
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
func (it *PositionManagerPoolPremineIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerPoolPremineIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerPoolPremine represents a PoolPremine event raised by the PositionManager contract.
type PositionManagerPoolPremine struct {
	PoolId        [32]byte
	PremineAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPoolPremine is a free log retrieval operation binding the contract event 0xdfa452364a13ecc87d8b629926a27f0b82206c68e4126e138157c5d853c71d84.
//
// Solidity: event PoolPremine(bytes32 indexed _poolId, int256 _premineAmount)
func (_PositionManager *PositionManagerFilterer) FilterPoolPremine(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerPoolPremineIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "PoolPremine", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerPoolPremineIterator{contract: _PositionManager.contract, event: "PoolPremine", logs: logs, sub: sub}, nil
}

// WatchPoolPremine is a free log subscription operation binding the contract event 0xdfa452364a13ecc87d8b629926a27f0b82206c68e4126e138157c5d853c71d84.
//
// Solidity: event PoolPremine(bytes32 indexed _poolId, int256 _premineAmount)
func (_PositionManager *PositionManagerFilterer) WatchPoolPremine(opts *bind.WatchOpts, sink chan<- *PositionManagerPoolPremine, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "PoolPremine", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerPoolPremine)
				if err := _PositionManager.contract.UnpackLog(event, "PoolPremine", log); err != nil {
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

// ParsePoolPremine is a log parse operation binding the contract event 0xdfa452364a13ecc87d8b629926a27f0b82206c68e4126e138157c5d853c71d84.
//
// Solidity: event PoolPremine(bytes32 indexed _poolId, int256 _premineAmount)
func (_PositionManager *PositionManagerFilterer) ParsePoolPremine(log types.Log) (*PositionManagerPoolPremine, error) {
	event := new(PositionManagerPoolPremine)
	if err := _PositionManager.contract.UnpackLog(event, "PoolPremine", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerPoolScheduledIterator is returned from FilterPoolScheduled and is used to iterate over the raw logs and unpacked data for PoolScheduled events raised by the PositionManager contract.
type PositionManagerPoolScheduledIterator struct {
	Event *PositionManagerPoolScheduled // Event containing the contract specifics and raw log

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
func (it *PositionManagerPoolScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerPoolScheduled)
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
		it.Event = new(PositionManagerPoolScheduled)
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
func (it *PositionManagerPoolScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerPoolScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerPoolScheduled represents a PoolScheduled event raised by the PositionManager contract.
type PositionManagerPoolScheduled struct {
	PoolId      [32]byte
	FlaunchesAt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolScheduled is a free log retrieval operation binding the contract event 0x0a5b59cacc8ab09e32bf89cb0477200adacf9c0237786bf21d17f44cec998107.
//
// Solidity: event PoolScheduled(bytes32 indexed _poolId, uint256 _flaunchesAt)
func (_PositionManager *PositionManagerFilterer) FilterPoolScheduled(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerPoolScheduledIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "PoolScheduled", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerPoolScheduledIterator{contract: _PositionManager.contract, event: "PoolScheduled", logs: logs, sub: sub}, nil
}

// WatchPoolScheduled is a free log subscription operation binding the contract event 0x0a5b59cacc8ab09e32bf89cb0477200adacf9c0237786bf21d17f44cec998107.
//
// Solidity: event PoolScheduled(bytes32 indexed _poolId, uint256 _flaunchesAt)
func (_PositionManager *PositionManagerFilterer) WatchPoolScheduled(opts *bind.WatchOpts, sink chan<- *PositionManagerPoolScheduled, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "PoolScheduled", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerPoolScheduled)
				if err := _PositionManager.contract.UnpackLog(event, "PoolScheduled", log); err != nil {
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

// ParsePoolScheduled is a log parse operation binding the contract event 0x0a5b59cacc8ab09e32bf89cb0477200adacf9c0237786bf21d17f44cec998107.
//
// Solidity: event PoolScheduled(bytes32 indexed _poolId, uint256 _flaunchesAt)
func (_PositionManager *PositionManagerFilterer) ParsePoolScheduled(log types.Log) (*PositionManagerPoolScheduled, error) {
	event := new(PositionManagerPoolScheduled)
	if err := _PositionManager.contract.UnpackLog(event, "PoolScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerPoolStateUpdatedIterator is returned from FilterPoolStateUpdated and is used to iterate over the raw logs and unpacked data for PoolStateUpdated events raised by the PositionManager contract.
type PositionManagerPoolStateUpdatedIterator struct {
	Event *PositionManagerPoolStateUpdated // Event containing the contract specifics and raw log

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
func (it *PositionManagerPoolStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerPoolStateUpdated)
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
		it.Event = new(PositionManagerPoolStateUpdated)
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
func (it *PositionManagerPoolStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerPoolStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerPoolStateUpdated represents a PoolStateUpdated event raised by the PositionManager contract.
type PositionManagerPoolStateUpdated struct {
	PoolId       [32]byte
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	ProtocolFee  *big.Int
	SwapFee      *big.Int
	Liquidity    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolStateUpdated is a free log retrieval operation binding the contract event 0x1a111a34a945a6d821c9dc87031a478ad3107acb9b19f2ee72d3aaa72b0849c9.
//
// Solidity: event PoolStateUpdated(bytes32 indexed _poolId, uint160 _sqrtPriceX96, int24 _tick, uint24 _protocolFee, uint24 _swapFee, uint128 _liquidity)
func (_PositionManager *PositionManagerFilterer) FilterPoolStateUpdated(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerPoolStateUpdatedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "PoolStateUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerPoolStateUpdatedIterator{contract: _PositionManager.contract, event: "PoolStateUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolStateUpdated is a free log subscription operation binding the contract event 0x1a111a34a945a6d821c9dc87031a478ad3107acb9b19f2ee72d3aaa72b0849c9.
//
// Solidity: event PoolStateUpdated(bytes32 indexed _poolId, uint160 _sqrtPriceX96, int24 _tick, uint24 _protocolFee, uint24 _swapFee, uint128 _liquidity)
func (_PositionManager *PositionManagerFilterer) WatchPoolStateUpdated(opts *bind.WatchOpts, sink chan<- *PositionManagerPoolStateUpdated, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "PoolStateUpdated", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerPoolStateUpdated)
				if err := _PositionManager.contract.UnpackLog(event, "PoolStateUpdated", log); err != nil {
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

// ParsePoolStateUpdated is a log parse operation binding the contract event 0x1a111a34a945a6d821c9dc87031a478ad3107acb9b19f2ee72d3aaa72b0849c9.
//
// Solidity: event PoolStateUpdated(bytes32 indexed _poolId, uint160 _sqrtPriceX96, int24 _tick, uint24 _protocolFee, uint24 _swapFee, uint128 _liquidity)
func (_PositionManager *PositionManagerFilterer) ParsePoolStateUpdated(log types.Log) (*PositionManagerPoolStateUpdated, error) {
	event := new(PositionManagerPoolStateUpdated)
	if err := _PositionManager.contract.UnpackLog(event, "PoolStateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerPoolSwapIterator is returned from FilterPoolSwap and is used to iterate over the raw logs and unpacked data for PoolSwap events raised by the PositionManager contract.
type PositionManagerPoolSwapIterator struct {
	Event *PositionManagerPoolSwap // Event containing the contract specifics and raw log

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
func (it *PositionManagerPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerPoolSwap)
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
		it.Event = new(PositionManagerPoolSwap)
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
func (it *PositionManagerPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerPoolSwap represents a PoolSwap event raised by the PositionManager contract.
type PositionManagerPoolSwap struct {
	PoolId     [32]byte
	FlAmount0  *big.Int
	FlAmount1  *big.Int
	FlFee0     *big.Int
	FlFee1     *big.Int
	IspAmount0 *big.Int
	IspAmount1 *big.Int
	IspFee0    *big.Int
	IspFee1    *big.Int
	UniAmount0 *big.Int
	UniAmount1 *big.Int
	UniFee0    *big.Int
	UniFee1    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolSwap is a free log retrieval operation binding the contract event 0x35001030bac7516ed5b1daf42bbfd3eaf587acc70c90f7f00ee42d5a34d75233.
//
// Solidity: event PoolSwap(bytes32 indexed poolId, int256 flAmount0, int256 flAmount1, int256 flFee0, int256 flFee1, int256 ispAmount0, int256 ispAmount1, int256 ispFee0, int256 ispFee1, int256 uniAmount0, int256 uniAmount1, int256 uniFee0, int256 uniFee1)
func (_PositionManager *PositionManagerFilterer) FilterPoolSwap(opts *bind.FilterOpts, poolId [][32]byte) (*PositionManagerPoolSwapIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "PoolSwap", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerPoolSwapIterator{contract: _PositionManager.contract, event: "PoolSwap", logs: logs, sub: sub}, nil
}

// WatchPoolSwap is a free log subscription operation binding the contract event 0x35001030bac7516ed5b1daf42bbfd3eaf587acc70c90f7f00ee42d5a34d75233.
//
// Solidity: event PoolSwap(bytes32 indexed poolId, int256 flAmount0, int256 flAmount1, int256 flFee0, int256 flFee1, int256 ispAmount0, int256 ispAmount1, int256 ispFee0, int256 ispFee1, int256 uniAmount0, int256 uniAmount1, int256 uniFee0, int256 uniFee1)
func (_PositionManager *PositionManagerFilterer) WatchPoolSwap(opts *bind.WatchOpts, sink chan<- *PositionManagerPoolSwap, poolId [][32]byte) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "PoolSwap", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerPoolSwap)
				if err := _PositionManager.contract.UnpackLog(event, "PoolSwap", log); err != nil {
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

// ParsePoolSwap is a log parse operation binding the contract event 0x35001030bac7516ed5b1daf42bbfd3eaf587acc70c90f7f00ee42d5a34d75233.
//
// Solidity: event PoolSwap(bytes32 indexed poolId, int256 flAmount0, int256 flAmount1, int256 flFee0, int256 flFee1, int256 ispAmount0, int256 ispAmount1, int256 ispFee0, int256 ispFee1, int256 uniAmount0, int256 uniAmount1, int256 uniFee0, int256 uniFee1)
func (_PositionManager *PositionManagerFilterer) ParsePoolSwap(log types.Log) (*PositionManagerPoolSwap, error) {
	event := new(PositionManagerPoolSwap)
	if err := _PositionManager.contract.UnpackLog(event, "PoolSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerReferralEscrowUpdatedIterator is returned from FilterReferralEscrowUpdated and is used to iterate over the raw logs and unpacked data for ReferralEscrowUpdated events raised by the PositionManager contract.
type PositionManagerReferralEscrowUpdatedIterator struct {
	Event *PositionManagerReferralEscrowUpdated // Event containing the contract specifics and raw log

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
func (it *PositionManagerReferralEscrowUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerReferralEscrowUpdated)
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
		it.Event = new(PositionManagerReferralEscrowUpdated)
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
func (it *PositionManagerReferralEscrowUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerReferralEscrowUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerReferralEscrowUpdated represents a ReferralEscrowUpdated event raised by the PositionManager contract.
type PositionManagerReferralEscrowUpdated struct {
	ReferralEscrow common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReferralEscrowUpdated is a free log retrieval operation binding the contract event 0x9c922de256a07b4d188faacda5c1abb2cae12f74f4370d5c2f11efb37a742d70.
//
// Solidity: event ReferralEscrowUpdated(address _referralEscrow)
func (_PositionManager *PositionManagerFilterer) FilterReferralEscrowUpdated(opts *bind.FilterOpts) (*PositionManagerReferralEscrowUpdatedIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "ReferralEscrowUpdated")
	if err != nil {
		return nil, err
	}
	return &PositionManagerReferralEscrowUpdatedIterator{contract: _PositionManager.contract, event: "ReferralEscrowUpdated", logs: logs, sub: sub}, nil
}

// WatchReferralEscrowUpdated is a free log subscription operation binding the contract event 0x9c922de256a07b4d188faacda5c1abb2cae12f74f4370d5c2f11efb37a742d70.
//
// Solidity: event ReferralEscrowUpdated(address _referralEscrow)
func (_PositionManager *PositionManagerFilterer) WatchReferralEscrowUpdated(opts *bind.WatchOpts, sink chan<- *PositionManagerReferralEscrowUpdated) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "ReferralEscrowUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerReferralEscrowUpdated)
				if err := _PositionManager.contract.UnpackLog(event, "ReferralEscrowUpdated", log); err != nil {
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

// ParseReferralEscrowUpdated is a log parse operation binding the contract event 0x9c922de256a07b4d188faacda5c1abb2cae12f74f4370d5c2f11efb37a742d70.
//
// Solidity: event ReferralEscrowUpdated(address _referralEscrow)
func (_PositionManager *PositionManagerFilterer) ParseReferralEscrowUpdated(log types.Log) (*PositionManagerReferralEscrowUpdated, error) {
	event := new(PositionManagerReferralEscrowUpdated)
	if err := _PositionManager.contract.UnpackLog(event, "ReferralEscrowUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerReferrerFeePaidIterator is returned from FilterReferrerFeePaid and is used to iterate over the raw logs and unpacked data for ReferrerFeePaid events raised by the PositionManager contract.
type PositionManagerReferrerFeePaidIterator struct {
	Event *PositionManagerReferrerFeePaid // Event containing the contract specifics and raw log

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
func (it *PositionManagerReferrerFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerReferrerFeePaid)
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
		it.Event = new(PositionManagerReferrerFeePaid)
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
func (it *PositionManagerReferrerFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerReferrerFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerReferrerFeePaid represents a ReferrerFeePaid event raised by the PositionManager contract.
type PositionManagerReferrerFeePaid struct {
	PoolId    [32]byte
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReferrerFeePaid is a free log retrieval operation binding the contract event 0x3ca4a7850462c23d5854e8b3e852626beb21b37354c670c0135ab46f0c4bdc31.
//
// Solidity: event ReferrerFeePaid(bytes32 indexed _poolId, address _recipient, address _token, uint256 _amount)
func (_PositionManager *PositionManagerFilterer) FilterReferrerFeePaid(opts *bind.FilterOpts, _poolId [][32]byte) (*PositionManagerReferrerFeePaidIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "ReferrerFeePaid", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerReferrerFeePaidIterator{contract: _PositionManager.contract, event: "ReferrerFeePaid", logs: logs, sub: sub}, nil
}

// WatchReferrerFeePaid is a free log subscription operation binding the contract event 0x3ca4a7850462c23d5854e8b3e852626beb21b37354c670c0135ab46f0c4bdc31.
//
// Solidity: event ReferrerFeePaid(bytes32 indexed _poolId, address _recipient, address _token, uint256 _amount)
func (_PositionManager *PositionManagerFilterer) WatchReferrerFeePaid(opts *bind.WatchOpts, sink chan<- *PositionManagerReferrerFeePaid, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "ReferrerFeePaid", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerReferrerFeePaid)
				if err := _PositionManager.contract.UnpackLog(event, "ReferrerFeePaid", log); err != nil {
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

// ParseReferrerFeePaid is a log parse operation binding the contract event 0x3ca4a7850462c23d5854e8b3e852626beb21b37354c670c0135ab46f0c4bdc31.
//
// Solidity: event ReferrerFeePaid(bytes32 indexed _poolId, address _recipient, address _token, uint256 _amount)
func (_PositionManager *PositionManagerFilterer) ParseReferrerFeePaid(log types.Log) (*PositionManagerReferrerFeePaid, error) {
	event := new(PositionManagerReferrerFeePaid)
	if err := _PositionManager.contract.UnpackLog(event, "ReferrerFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the PositionManager contract.
type PositionManagerWithdrawalIterator struct {
	Event *PositionManagerWithdrawal // Event containing the contract specifics and raw log

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
func (it *PositionManagerWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerWithdrawal)
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
		it.Event = new(PositionManagerWithdrawal)
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
func (it *PositionManagerWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerWithdrawal represents a Withdrawal event raised by the PositionManager contract.
type PositionManagerWithdrawal struct {
	Sender    common.Address
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x342e7ff505a8a0364cd0dc2ff195c315e43bce86b204846ecd36913e117b109e.
//
// Solidity: event Withdrawal(address _sender, address _recipient, address _token, uint256 _amount)
func (_PositionManager *PositionManagerFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*PositionManagerWithdrawalIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &PositionManagerWithdrawalIterator{contract: _PositionManager.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x342e7ff505a8a0364cd0dc2ff195c315e43bce86b204846ecd36913e117b109e.
//
// Solidity: event Withdrawal(address _sender, address _recipient, address _token, uint256 _amount)
func (_PositionManager *PositionManagerFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *PositionManagerWithdrawal) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerWithdrawal)
				if err := _PositionManager.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x342e7ff505a8a0364cd0dc2ff195c315e43bce86b204846ecd36913e117b109e.
//
// Solidity: event Withdrawal(address _sender, address _recipient, address _token, uint256 _amount)
func (_PositionManager *PositionManagerFilterer) ParseWithdrawal(log types.Log) (*PositionManagerWithdrawal, error) {
	event := new(PositionManagerWithdrawal)
	if err := _PositionManager.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
