// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Flaunch} from '../src/contracts/Flaunch.sol';
import {Memecoin} from '../src/contracts/Memecoin.sol';
import {PositionManager} from '../src/contracts/PositionManager.sol';
import {BidWall} from '../src/contracts/bidwall/BidWall.sol';

import {StaticFeeCalculator} from '../src/contracts/fees/StaticFeeCalculator.sol';
import {FairLaunch} from '../src/contracts/hooks/FairLaunch.sol';
import {FeeDistributor} from '../src/contracts/hooks/FeeDistributor.sol';
import {FeeExemptions} from '../src/contracts/hooks/FeeExemptions.sol';
import {Notifier} from '../src/contracts/hooks/Notifier.sol';
import {MarketCappedPriceV3} from '../src/contracts/price/MarketCappedPriceV3.sol';

import {ReferralEscrow} from '../src/contracts/referrals/ReferralEscrow.sol';
import {PreventNoFairLaunch} from '../src/contracts/subscribers/PreventNoFairLaunch.sol';
import {TreasuryActionManager} from '../src/contracts/treasury/ActionManager.sol';
import {MemecoinTreasury} from '../src/contracts/treasury/MemecoinTreasury.sol';
import {BurnTokensAction} from '../src/contracts/treasury/actions/BurnTokens.sol';
import {BuyBackAction} from '../src/contracts/treasury/actions/BuyBack.sol';

import {FastFlaunchZap} from '../src/contracts/zaps/FastFlaunchZap.sol';
import {FlaunchPremineZap} from '../src/contracts/zaps/FlaunchPremineZap.sol';
import {PoolSwap} from '../src/contracts/zaps/PoolSwap.sol';
import {IFeeCalculator} from '../src/interfaces/IFeeCalculator.sol';

import {HookMiner} from './HookMiner.sol';

import {IPoolManager} from '@uniswap/v4-core/src/interfaces/IPoolManager.sol';
import {Hooks} from '@uniswap/v4-core/src/libraries/Hooks.sol';

import {Script} from 'forge-std/Script.sol';

import {stdJson} from 'forge-std/StdJson.sol';
import {console} from 'forge-std/console.sol';

contract DeployScript is Script {
    using stdJson for string;

    uint privateKey = vm.envUint(string.concat('PRIVATE_KEY_', vm.toString(block.chainid)));
    address deployer = vm.addr(privateKey);

    address WETH_Contract;
    address USDC_Contract;
    address nativeToken_Contract;
    address uni_USDC_WETH_Pool;
    address v4_PoolManager_Contract;
    uint swapFeeThreshold = 100;
    string baseURI = 'https://flaunch.xyz/';

    FeeExemptions feeExemptions;
    MarketCappedPriceV3 marketCappedPrice;
    TreasuryActionManager actionManager;
    BidWall bidWall;
    FairLaunch fairLaunch;
    PositionManager positionManager;
    Flaunch flaunch;
    StaticFeeCalculator staticFeeCalculator;
    BuyBackAction buyBackAction;
    BurnTokensAction burnTokensAction;
    PoolSwap poolSwap;
    FlaunchPremineZap flaunchPremineZap;
    ReferralEscrow referralEscrow;
    FastFlaunchZap fastFlaunchZap;
    MemecoinTreasury memecoinTreasury;
    Memecoin memecoin;
    Notifier notifier;
    PreventNoFairLaunch preventNoFairLaunch;

    PositionManager.ConstructorParams positionManagerParams;

    FeeDistributor.FeeDistribution feeDistribution;

    function setUp() public {
        string memory root = vm.projectRoot();
        string memory path = string.concat(root, '/script/configs/', vm.toString(block.chainid), '.json');
        string memory configs = vm.readFile(path);
        WETH_Contract = configs.readAddress('.WETH_Contract');
        USDC_Contract = configs.readAddress('.USDC_Contract');
        nativeToken_Contract = configs.readAddress('.nativeToken_Contract');
        v4_PoolManager_Contract = configs.readAddress('.v4_PoolManager_Contract');
        uni_USDC_WETH_Pool = configs.readAddress('.uni_USDC_WETH_Pool');

        feeDistribution = FeeDistributor.FeeDistribution({
            swapFee: uint24(configs.readUint('.swapFee')),
            referrer: uint24(configs.readUint('.referrerFee')),
            protocol: uint24(configs.readUint('.protocolFee')),
            active: true
        });
    }

    function run() public {
        vm.startBroadcast(privateKey);
        feeExemptions = new FeeExemptions(deployer);
        marketCappedPrice = new MarketCappedPriceV3(deployer, WETH_Contract, USDC_Contract);
        actionManager = new TreasuryActionManager(deployer);
        bidWall = new BidWall(nativeToken_Contract, v4_PoolManager_Contract, deployer);
        fairLaunch = new FairLaunch(IPoolManager(v4_PoolManager_Contract));
        notifier = new Notifier(deployer);
        memecoin = new Memecoin();
        memecoinTreasury = new MemecoinTreasury();
        staticFeeCalculator = new StaticFeeCalculator();
        flaunch = new Flaunch(address(memecoin), baseURI);
        poolSwap = new PoolSwap(IPoolManager(v4_PoolManager_Contract));
        flaunchPremineZap = new FlaunchPremineZap(positionManager, address(flaunch), nativeToken_Contract, poolSwap);
        buyBackAction = new BuyBackAction(nativeToken_Contract, address(poolSwap));
        burnTokensAction = new BurnTokensAction(nativeToken_Contract);
        referralEscrow = new ReferralEscrow(nativeToken_Contract, address(positionManager));
        fastFlaunchZap = new FastFlaunchZap(positionManager);
        preventNoFairLaunch = new PreventNoFairLaunch(address(notifier));

        positionManagerParams = PositionManager.ConstructorParams({
            nativeToken: nativeToken_Contract,
            poolManager: IPoolManager(v4_PoolManager_Contract),
            feeDistribution: feeDistribution,
            initialPrice: marketCappedPrice,
            protocolOwner: deployer,
            protocolFeeRecipient: deployer,
            flayGovernance: deployer,
            feeExemptions: feeExemptions
        });
        positionManager = deployPositionManager();

        flaunch.initialize(positionManager, address(memecoinTreasury));
        positionManager.setFlaunch(address(flaunch));
        positionManager.setFeeCalculator(IFeeCalculator(staticFeeCalculator));
        positionManager.setReferralEscrow(payable(address(referralEscrow)));

        marketCappedPrice.setPool(uni_USDC_WETH_Pool);
        bidWall.setSwapFeeThreshold(swapFeeThreshold);
        actionManager.approveAction(address(buyBackAction));
        actionManager.approveAction(address(burnTokensAction));
        notifier.subscribe(address(preventNoFairLaunch), bytes(''));
        referralEscrow.setPoolSwap(address(poolSwap));

        vm.stopBroadcast();
    }

    function deployPositionManager() public returns (PositionManager pm) {
        uint160 flags = uint160(
            Hooks.BEFORE_INITIALIZE_FLAG | Hooks.BEFORE_ADD_LIQUIDITY_FLAG | Hooks.AFTER_ADD_LIQUIDITY_FLAG
                | Hooks.BEFORE_REMOVE_LIQUIDITY_FLAG | Hooks.AFTER_REMOVE_LIQUIDITY_FLAG | Hooks.BEFORE_SWAP_FLAG
                | Hooks.AFTER_SWAP_FLAG | Hooks.AFTER_DONATE_FLAG | Hooks.BEFORE_SWAP_RETURNS_DELTA_FLAG
                | Hooks.AFTER_SWAP_RETURNS_DELTA_FLAG
        );
        (address hookAddress, bytes32 salt) = HookMiner.find(
            0x4e59b44847b379578588920cA78FbF26c0B4956C,
            flags,
            type(PositionManager).creationCode,
            abi.encode(positionManagerParams)
        );
        pm = new PositionManager{salt: salt}(positionManagerParams);
        require(address(pm) == hookAddress, 'Invalid hook address');
        return pm;
    }
}
