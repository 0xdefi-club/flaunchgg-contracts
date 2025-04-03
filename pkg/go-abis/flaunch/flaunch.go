// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flaunch

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

// FlaunchMemecoinMetadata is an auto generated low-level Go binding around an user-defined struct.
type FlaunchMemecoinMetadata struct {
	Name     string
	Symbol   string
	TokenUri string
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

// FlaunchMetaData contains all meta data concerning the Flaunch contract.
var FlaunchMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_memecoinImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_baseURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_CREATOR_ALLOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FAIR_LAUNCH_TOKENS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_SCHEDULE_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseURI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgingStatus\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_started\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"finalizeBridge\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_metadata\",\"type\":\"tuple\",\"internalType\":\"structFlaunch.MemecoinMetadata\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenUri\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"flaunch\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple\",\"internalType\":\"structPositionManager.FlaunchParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialTokenFairLaunch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"premineAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creatorFeeAllocation\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"flaunchAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialPriceParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeCalculatorParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"memecoin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"memecoinTreasury_\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"tokenId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_positionManager\",\"type\":\"address\",\"internalType\":\"contractPositionManager\"},{\"name\":\"_memecoinTreasuryImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeBridge\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"memecoin\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"memecoinImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"memecoinTreasury\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"memecoinTreasuryImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextTokenId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPositionManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isApproved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBaseURI\",\"inputs\":[{\"name\":\"_baseURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMemecoinMetadata\",\"inputs\":[{\"name\":\"_memecoin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenId\",\"inputs\":[{\"name\":\"_memecoin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isApproved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenBridged\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_memecoin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_messageSource\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenBridging\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_memecoin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccountBalanceOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BalanceQueryForZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotPositionManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotL2ToL2CrossDomainMessenger\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CreatorFeeAllocationInvalid\",\"inputs\":[{\"name\":\"_allocation\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"_maxAllocation\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidCrossDomainSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestinationChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFlaunchSchedule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialSupply\",\"inputs\":[{\"name\":\"_initialSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwnerNorApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PremineExceedsInitialAmount\",\"inputs\":[{\"name\":\"_buyAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_initialSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenAlreadyBridged\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFromIncorrectOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferToNonERC721ReceiverImplementer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownMemecoin\",\"inputs\":[]}]",
	Bin: "0x5f80546001600160a01b03191673420000000000000000000000000000000000002317905560e0604052601360a09081527f466c61756e6368204d656d6573747265616d730000000000000000000000000060c05260019061006190826101d7565b5060408051808201909152600781526608c9882aa9c86960cb1b602082015260029061008d90826101d7565b50600160045534801561009e575f5ffd5b506040516126623803806126628339810160408190526100bd91610291565b6001600160a01b03821660805260036100d682826101d7565b506100e0336100e7565b5050610360565b638b78c6d81980541561010157630dc149f05f526004601cfd5b6001600160a01b03909116801560ff1b8117909155805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061016757607f821691505b60208210810361018557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d257805f5260205f20601f840160051c810160208510156101b05750805b601f840160051c820191505b818110156101cf575f81556001016101bc565b50505b505050565b81516001600160401b038111156101f0576101f061013f565b610204816101fe8454610153565b8461018b565b6020601f821160018114610236575f831561021f5750848201515b5f19600385901b1c1916600184901b1784556101cf565b5f84815260208120601f198516915b828110156102655787850151825560209485019460019092019101610245565b508482101561028257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f5f604083850312156102a2575f5ffd5b82516001600160a01b03811681146102b8575f5ffd5b60208401519092506001600160401b038111156102d3575f5ffd5b8301601f810185136102e3575f5ffd5b80516001600160401b038111156102fc576102fc61013f565b604051601f8201601f19908116603f011681016001600160401b038111828210171561032a5761032a61013f565b604052818152828201602001871015610341575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b6080516122dc6103865f395f818161056f01528181610a140152610c6e01526122dc5ff3fe6080604052600436106101bf575f3560e01c806301ffc9a7146101c357806306fdde0314610214578063081812fc14610235578063095ea7b31461026c5780631a5de1df1461028157806323b872dd146102a057806325692962146102b35780632a12e488146102bb5780632b31d94a14610300578063359687e41461031f57806342842e0e1461035857806342966c681461036b578063485cc9551461038a57806354d1f13d146103a957806355f804b3146103b15780635b7426be146103d05780635f151475146103f45780636352211e146104095780636c0360eb1461042857806370a082311461043c578063715018a61461045b57806375794a3c14610463578063791b98bc146104785780637ca3172414610497578063829f042a146104c25780638da5cb5b146104e157806395d89b41146104f957806397cd41a81461050d578063a22cb4651461052c578063b88d4fde1461054b578063bf989c781461055e578063c87b56dd14610591578063e49c1854146105b0578063e4b47c0b146105cf578063e985e9c5146105ee578063f04e283e14610621578063f2fde38b14610634578063f38f955014610647578063fee81cf41461067e575b5f5ffd5b3480156101ce575f5ffd5b506101ff6101dd3660046119e8565b6301ffc9a760e09190911c9081146380ac58cd821417635b5e139f9091141790565b60405190151581526020015b60405180910390f35b34801561021f575f5ffd5b506102286106af565b60405161020b9190611a3d565b348015610240575f5ffd5b5061025461024f366004611a4f565b61073f565b6040516001600160a01b03909116815260200161020b565b61027f61027a366004611a7a565b61077a565b005b34801561028c575f5ffd5b50600654610254906001600160a01b031681565b61027f6102ae366004611aa4565b610789565b61027f610870565b3480156102c6575f5ffd5b506102da6102d5366004611ae2565b6108bc565b604080516001600160a01b0394851681529390921660208401529082015260600161020b565b34801561030b575f5ffd5b5061027f61031a366004611bfd565b610bae565b34801561032a575f5ffd5b506101ff610339366004611cc7565b600960209081525f928352604080842090915290825290205460ff1681565b61027f610366366004611aa4565b610dd0565b348015610376575f5ffd5b5061027f610385366004611a4f565b610dfc565b348015610395575f5ffd5b5061027f6103a4366004611ce7565b610e09565b61027f610eb1565b3480156103bc575f5ffd5b5061027f6103cb366004611d1e565b610eea565b3480156103db575f5ffd5b506103e662278d0081565b60405190815260200161020b565b3480156103ff575f5ffd5b506103e661271081565b348015610414575f5ffd5b50610254610423366004611a4f565b610efe565b348015610433575f5ffd5b50610228610f3a565b348015610447575f5ffd5b506103e6610456366004611d57565b610fc6565b61027f610ffe565b34801561046e575f5ffd5b506103e660045481565b348015610483575f5ffd5b50600554610254906001600160a01b031681565b3480156104a2575f5ffd5b506103e66104b1366004611d57565b60086020525f908152604090205481565b3480156104cd575f5ffd5b5061027f6104dc366004611db6565b611011565b3480156104ec575f5ffd5b50638b78c6d81954610254565b348015610504575f5ffd5b5061022861107f565b348015610518575f5ffd5b506103e66bdef376571332906a8800000081565b348015610537575f5ffd5b5061027f610546366004611e36565b61108e565b61027f610559366004611e66565b6110e0565b348015610569575f5ffd5b506102547f000000000000000000000000000000000000000000000000000000000000000081565b34801561059c575f5ffd5b506102286105ab366004611a4f565b61113a565b3480156105bb575f5ffd5b506102546105ca366004611a4f565b611292565b3480156105da575f5ffd5b5061027f6105e9366004611cc7565b6112ac565b3480156105f9575f5ffd5b506101ff610608366004611ce7565b601c5263052d173d60211b6008525f526030600c205490565b61027f61062f366004611d57565b6115ba565b61027f610642366004611d57565b6115f4565b348015610652575f5ffd5b50610254610661366004611a4f565b5f908152600760205260409020600101546001600160a01b031690565b348015610689575f5ffd5b506103e6610698366004611d57565b63389a75e1600c9081525f91909152602090205490565b6060600180546106be90611eb6565b80601f01602080910402602001604051908101604052809291908181526020018280546106ea90611eb6565b80156107355780601f1061070c57610100808354040283529160200191610735565b820191905f5260205f20905b81548152906001019060200180831161071857829003601f168201915b5050505050905090565b5f815f52673ec412a9852d173d60c11b601c5260205f2082018201805460601b6107705763ceea21b65f526004601cfd5b6001015492915050565b61078533838361161a565b5050565b5f818152673ec412a9852d173d60c11b3317601c52602090208101810180546001600160a01b0394851694938416938116919082861483026107da5767ceea21b6a1148100831560021b526004601cfd5b855f528160010154925082331486331417610806576030600c205461080657634b6e7f185f526004601cfd5b8215610813575f82600101555b85851818905550601c600c81812080545f190190555f84905220805460010163ffffffff811684026108545767ea553b3401336cea841560021b526004601cfd5b90558082845f5160206122875f395f51905f525f38a45b505050565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f5fa250565b6005545f90819081906001600160a01b031633146108ed57604051630a2840f360e01b815260040160405180910390fd5b6108fa62278d0042611eee565b8460e00135111561091e57604051630f0ffca560e41b815260040160405180910390fd5b6bdef376571332906a880000008460600135111561095a57604051631554010f60e01b8152606085013560048201526024015b60405180910390fd5b83606001358460800135111561099357604051639f68473d60e01b81526080850135600482015260608501356024820152604401610951565b6127106109a660e0860160c08701611f0d565b62ffffff1611156109ea576109c160e0850160c08601611f0d565b604051633707e93960e21b815262ffffff90911660048201526127106024820152604401610951565b506004805460018101909155610a0f610a0960c0860160a08701611d57565b826116b4565b610a397f00000000000000000000000000000000000000000000000000000000000000008261173d565b6001600160a01b0381165f818152600860205260409020839055909350839063a6487c53610a678780611f2f565b610a7460208a018a611f2f565b610a8160408c018c611f2f565b6040518763ffffffff1660e01b8152600401610aa296959493929190611f99565b5f604051808303815f87803b158015610ab9575f5ffd5b505af1158015610acb573d5f5f3e3d5ffd5b5050600654610ae692506001600160a01b031690508361173d565b6040805180820182526001600160a01b03878116825283811660208084019182525f8881526007909152849020925183549083166001600160a01b0319918216178455905160019093018054938316939091169290921790915560055491516340c10f1960e01b81529181166004830152680a18f07d736b90be55601d1b6024830152919450908216906340c10f19906044015f604051808303815f87803b158015610b90575f5ffd5b505af1158015610ba2573d5f5f3e3d5ffd5b50505050509193909250565b5f546001600160a01b03163314610bd7576040516265d51560e41b815260040160405180910390fd5b5f546040805163071ffbc360e31b8152905130926001600160a01b0316916338ffde189160048083019260209291908290030181865afa158015610c1d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c419190611fe1565b6001600160a01b031614610c6857604051635e11715560e11b815260040160405180910390fd5b5f610c937f00000000000000000000000000000000000000000000000000000000000000008461173d565b82516020840151604080860151905163a6487c5360e01b81529394506001600160a01b0385169363a6487c5393610cd09390929091600401611ffc565b5f604051808303815f87803b158015610ce7575f5ffd5b505af1158015610cf9573d5f5f3e3d5ffd5b505050507fd953cced6d6de41e29efa3f395f6815b35e5a0bb1bff016cb796b95e34a932548346835f5f9054906101000a90046001600160a01b03166001600160a01b031663247944626040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d70573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d94919061203e565b604051610dc3949392919093845260208401929092526001600160a01b03166040830152606082015260800190565b60405180910390a1505050565b610ddb838383610789565b813b1561086b5761086b83838360405180602001604052805f815250611750565b610e0633826117d9565b50565b610e1161188c565b63409feecd198054600382558015610e475760018160011c14303b10610e3e5763f92ee8a95f526004601cfd5b818160ff1b1b91505b50600580546001600160a01b038086166001600160a01b0319928316179092556006805492851692909116919091179055801561086b576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a1505050565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f5fa2565b610ef261188c565b60036107858282612099565b5f818152673ec412a9852d173d60c11b601c526020902081018101546001600160a01b031680610f355763ceea21b65f526004601cfd5b919050565b60038054610f4790611eb6565b80601f0160208091040260200160405190810160405280929190818152602001828054610f7390611eb6565b8015610fbe5780601f10610f9557610100808354040283529160200191610fbe565b820191905f5260205f20905b815481529060010190602001808311610fa157829003601f168201915b505050505081565b5f81610fd957638f4eb6045f526004601cfd5b673ec412a9852d173d60c11b601c52815f5263ffffffff601c600c2054169050919050565b61100661188c565b61100f5f6118a6565b565b61101961188c565b60405163051335b560e41b81526001600160a01b038616906351335b509061104b908790879087908790600401612153565b5f604051808303815f87803b158015611062575f5ffd5b505af1158015611074573d5f5f3e3d5ffd5b505050505050505050565b6060600280546106be90611eb6565b801515905081601c5263052d173d60211b600852335f52806030600c2055805f528160601b60601c337f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160205fa35050565b6110eb858585610789565b833b156111335761113385858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061175092505050565b5050505050565b606081158061114b57506004548210155b156111695760405163677510db60e11b815260040160405180910390fd5b6003805461117690611eb6565b90505f036111f8575f828152600760205260408082205481516303c130d960e41b815291516001600160a01b0390911692633c130d9092600480820193918290030181865afa1580156111cb573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526111f29190810190612184565b92915050565b6111f26003805461120890611eb6565b80601f016020809104026020016040519081016040528092919081815260200182805461123490611eb6565b801561127f5780601f106112565761010080835404028352916020019161127f565b820191905f5260205f20905b81548152906001019060200180831161126257829003601f168201915b505050505061128d846118ec565b61192e565b5f908152600760205260409020546001600160a01b031690565b4681036112cc5760405163b86ac1ef60e01b815260040160405180910390fd5b5f82815260096020908152604080832084845290915290205460ff161561130657604051635314a19160e11b815260040160405180910390fd5b5f8281526009602090815260408083208484529091528120805460ff1916600117905561133283611292565b90506001600160a01b03811661135b5760405163a4c0139760e01b815260040160405180910390fd5b5f8190505f5f9054906101000a90046001600160a01b03166001600160a01b0316637056f41f8430306001600160a01b0316632b31d94a896040518060600160405280896001600160a01b03166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa1580156113d9573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526114009190810190612184565b8152602001896001600160a01b03166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015611440573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526114679190810190612184565b8152602001896001600160a01b0316633c130d906040518163ffffffff1660e01b81526004015f60405180830381865afa1580156114a7573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526114ce9190810190612184565b90526040516114e19291906024016121f8565b60408051808303601f1901815291815260208201805160e094851b6001600160e01b03909116179052519186901b6001600160e01b031916825261152b9493925090600401612254565b6020604051808303815f875af1158015611547573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061156b919061203e565b5060408051858152602081018590526001600160a01b0384168183015290517ff2a584a2309fee3f807e426ce11328f237b9bc191989cbc6f62cc5fee1d6b46f9181900360600190a150505050565b6115c261188c565b63389a75e1600c52805f526020600c2080544211156115e857636f5e88185f526004601cfd5b5f9055610e06816118a6565b6115fc61188c565b8060601b61161157637448fbae5f526004601cfd5b610e06816118a6565b5f1960601c82811692508381169350815f5283673ec412a9852d173d60c11b17601c5260205f20820182018054821691508161165d5763ceea21b65f526004601cfd5b81851485151761168157815f526030600c205461168157634b6e7f185f526004601cfd5b6001018390558183827f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9255f38a450505050565b8160601b60601c9150805f52673ec412a9852d173d60c11b601c5260205f208101810180548060601b156116ef5763c991cbb15f526004601cfd5b831790555f829052601c600c20805460010163ffffffff811684026117235767ea553b3401336cea841560021b526004601cfd5b905580825f5f5160206122875f395f51905f528138a45050565b5f6117495f848461193a565b9392505050565b60405163150b7a028082523360208301528560601b60601c604083015283606083015260808083015282518060a08401528015611797578060c08401826020870160045afa505b60208360a48301601c86015f8a5af16117b8573d156117b8573d5f843e3d83fd5b508060e01b8251146117d15763d1a57ed65f526004601cfd5b505050505050565b5f6117e382610efe565b9050505f8181526001600160a01b03928316673ec412a9852d173d60c11b8117601c5260209091208201820180549193821691826118285763ceea21b65f526004601cfd5b825f52816001015480861484871417861517611855576030600c205461185557634b6e7f185f526004601cfd5b8015611862575f83600101555b5082189055601c600c2080545f19019055815f825f5160206122875f395f51905f528238a4505050565b638b78c6d81954331461100f576382b429005f526004601cfd5b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3811560ff1b8217905550565b60606080604051019050602081016040525f8152805f19835b928101926030600a8206018453600a900480611905575050819003601f19909101908152919050565b6060611749838361198d565b5f6c5af43d3d93803e602a57fd5bf36021528260145273602c3d8160093d39f33d3d3d3d363d3d37363d735f52816035600c86f59050806119825763301164255f526004601cfd5b5f6021529392505050565b6040518251601f19906020810182165b858101518482015282018061199d575083518184018360208301165b86810151828201528401806119b95750505f910183810160208101929092528352604090810190525092915050565b5f602082840312156119f8575f5ffd5b81356001600160e01b031981168114611749575f5ffd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6117496020830184611a0f565b5f60208284031215611a5f575f5ffd5b5035919050565b6001600160a01b0381168114610e06575f5ffd5b5f5f60408385031215611a8b575f5ffd5b8235611a9681611a66565b946020939093013593505050565b5f5f5f60608486031215611ab6575f5ffd5b8335611ac181611a66565b92506020840135611ad181611a66565b929592945050506040919091013590565b5f60208284031215611af2575f5ffd5b81356001600160401b03811115611b07575f5ffd5b82016101408185031215611749575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715611b4f57611b4f611b19565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611b7d57611b7d611b19565b604052919050565b5f6001600160401b03821115611b9d57611b9d611b19565b50601f01601f191660200190565b5f82601f830112611bba575f5ffd5b8135611bcd611bc882611b85565b611b55565b818152846020838601011115611be1575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f60408385031215611c0e575f5ffd5b8235915060208301356001600160401b03811115611c2a575f5ffd5b830160608186031215611c3b575f5ffd5b611c43611b2d565b81356001600160401b03811115611c58575f5ffd5b611c6487828501611bab565b82525060208201356001600160401b03811115611c7f575f5ffd5b611c8b87828501611bab565b60208301525060408201356001600160401b03811115611ca9575f5ffd5b611cb587828501611bab565b60408301525080925050509250929050565b5f5f60408385031215611cd8575f5ffd5b50508035926020909101359150565b5f5f60408385031215611cf8575f5ffd5b8235611d0381611a66565b91506020830135611d1381611a66565b809150509250929050565b5f60208284031215611d2e575f5ffd5b81356001600160401b03811115611d43575f5ffd5b611d4f84828501611bab565b949350505050565b5f60208284031215611d67575f5ffd5b813561174981611a66565b5f5f83601f840112611d82575f5ffd5b5081356001600160401b03811115611d98575f5ffd5b602083019150836020828501011115611daf575f5ffd5b9250929050565b5f5f5f5f5f60608688031215611dca575f5ffd5b8535611dd581611a66565b945060208601356001600160401b03811115611def575f5ffd5b611dfb88828901611d72565b90955093505060408601356001600160401b03811115611e19575f5ffd5b611e2588828901611d72565b969995985093965092949392505050565b5f5f60408385031215611e47575f5ffd5b8235611e5281611a66565b915060208301358015158114611d13575f5ffd5b5f5f5f5f5f60808688031215611e7a575f5ffd5b8535611e8581611a66565b94506020860135611e9581611a66565b93506040860135925060608601356001600160401b03811115611e19575f5ffd5b600181811c90821680611eca57607f821691505b602082108103611ee857634e487b7160e01b5f52602260045260245ffd5b50919050565b808201808211156111f257634e487b7160e01b5f52601160045260245ffd5b5f60208284031215611f1d575f5ffd5b813562ffffff81168114611749575f5ffd5b5f5f8335601e19843603018112611f44575f5ffd5b8301803591506001600160401b03821115611f5d575f5ffd5b602001915036819003821315611daf575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b606081525f611fac60608301888a611f71565b8281036020840152611fbf818789611f71565b90508281036040840152611fd4818587611f71565b9998505050505050505050565b5f60208284031215611ff1575f5ffd5b815161174981611a66565b606081525f61200e6060830186611a0f565b82810360208401526120208186611a0f565b905082810360408401526120348185611a0f565b9695505050505050565b5f6020828403121561204e575f5ffd5b5051919050565b601f82111561086b57805f5260205f20601f840160051c8101602085101561207a5750805b601f840160051c820191505b81811015611133575f8155600101612086565b81516001600160401b038111156120b2576120b2611b19565b6120c6816120c08454611eb6565b84612055565b6020601f8211600181146120f8575f83156120e15750848201515b5f19600385901b1c1916600184901b178455611133565b5f84815260208120601f198516915b828110156121275787850151825560209485019460019092019101612107565b508482101561214457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b604081525f612166604083018688611f71565b8281036020840152612179818587611f71565b979650505050505050565b5f60208284031215612194575f5ffd5b81516001600160401b038111156121a9575f5ffd5b8201601f810184136121b9575f5ffd5b80516121c7611bc882611b85565b8181528560208385010111156121db575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b828152604060208201525f82516060604084015261221960a0840182611a0f565b90506020840151603f198483030160608501526122368282611a0f565b9150506040840151603f198483030160808501526120348282611a0f565b8381526001600160a01b03831660208201526060604082018190525f9061227d90830184611a0f565b9594505050505056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa26469706673582212209e53d1e15a46f68ea5b6fe9eff74b4eb848f7f23340cbc9fd424dd96f982cc3364736f6c634300081b0033",
}

// FlaunchABI is the input ABI used to generate the binding from.
// Deprecated: Use FlaunchMetaData.ABI instead.
var FlaunchABI = FlaunchMetaData.ABI

// FlaunchBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FlaunchMetaData.Bin instead.
var FlaunchBin = FlaunchMetaData.Bin

// DeployFlaunch deploys a new Ethereum contract, binding an instance of Flaunch to it.
func DeployFlaunch(auth *bind.TransactOpts, backend bind.ContractBackend, _memecoinImplementation common.Address, _baseURI string) (common.Address, *types.Transaction, *Flaunch, error) {
	parsed, err := FlaunchMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FlaunchBin), backend, _memecoinImplementation, _baseURI)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Flaunch{FlaunchCaller: FlaunchCaller{contract: contract}, FlaunchTransactor: FlaunchTransactor{contract: contract}, FlaunchFilterer: FlaunchFilterer{contract: contract}}, nil
}

// Flaunch is an auto generated Go binding around an Ethereum contract.
type Flaunch struct {
	FlaunchCaller     // Read-only binding to the contract
	FlaunchTransactor // Write-only binding to the contract
	FlaunchFilterer   // Log filterer for contract events
}

// FlaunchCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlaunchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlaunchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlaunchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlaunchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlaunchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlaunchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlaunchSession struct {
	Contract     *Flaunch          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlaunchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlaunchCallerSession struct {
	Contract *FlaunchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FlaunchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlaunchTransactorSession struct {
	Contract     *FlaunchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FlaunchRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlaunchRaw struct {
	Contract *Flaunch // Generic contract binding to access the raw methods on
}

// FlaunchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlaunchCallerRaw struct {
	Contract *FlaunchCaller // Generic read-only contract binding to access the raw methods on
}

// FlaunchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlaunchTransactorRaw struct {
	Contract *FlaunchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlaunch creates a new instance of Flaunch, bound to a specific deployed contract.
func NewFlaunch(address common.Address, backend bind.ContractBackend) (*Flaunch, error) {
	contract, err := bindFlaunch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Flaunch{FlaunchCaller: FlaunchCaller{contract: contract}, FlaunchTransactor: FlaunchTransactor{contract: contract}, FlaunchFilterer: FlaunchFilterer{contract: contract}}, nil
}

// NewFlaunchCaller creates a new read-only instance of Flaunch, bound to a specific deployed contract.
func NewFlaunchCaller(address common.Address, caller bind.ContractCaller) (*FlaunchCaller, error) {
	contract, err := bindFlaunch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlaunchCaller{contract: contract}, nil
}

// NewFlaunchTransactor creates a new write-only instance of Flaunch, bound to a specific deployed contract.
func NewFlaunchTransactor(address common.Address, transactor bind.ContractTransactor) (*FlaunchTransactor, error) {
	contract, err := bindFlaunch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlaunchTransactor{contract: contract}, nil
}

// NewFlaunchFilterer creates a new log filterer instance of Flaunch, bound to a specific deployed contract.
func NewFlaunchFilterer(address common.Address, filterer bind.ContractFilterer) (*FlaunchFilterer, error) {
	contract, err := bindFlaunch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlaunchFilterer{contract: contract}, nil
}

// bindFlaunch binds a generic wrapper to an already deployed contract.
func bindFlaunch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlaunchMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flaunch *FlaunchRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flaunch.Contract.FlaunchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flaunch *FlaunchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flaunch.Contract.FlaunchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flaunch *FlaunchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flaunch.Contract.FlaunchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flaunch *FlaunchCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flaunch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flaunch *FlaunchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flaunch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flaunch *FlaunchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flaunch.Contract.contract.Transact(opts, method, params...)
}

// MAXCREATORALLOCATION is a free data retrieval call binding the contract method 0x5f151475.
//
// Solidity: function MAX_CREATOR_ALLOCATION() view returns(uint256)
func (_Flaunch *FlaunchCaller) MAXCREATORALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "MAX_CREATOR_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCREATORALLOCATION is a free data retrieval call binding the contract method 0x5f151475.
//
// Solidity: function MAX_CREATOR_ALLOCATION() view returns(uint256)
func (_Flaunch *FlaunchSession) MAXCREATORALLOCATION() (*big.Int, error) {
	return _Flaunch.Contract.MAXCREATORALLOCATION(&_Flaunch.CallOpts)
}

// MAXCREATORALLOCATION is a free data retrieval call binding the contract method 0x5f151475.
//
// Solidity: function MAX_CREATOR_ALLOCATION() view returns(uint256)
func (_Flaunch *FlaunchCallerSession) MAXCREATORALLOCATION() (*big.Int, error) {
	return _Flaunch.Contract.MAXCREATORALLOCATION(&_Flaunch.CallOpts)
}

// MAXFAIRLAUNCHTOKENS is a free data retrieval call binding the contract method 0x97cd41a8.
//
// Solidity: function MAX_FAIR_LAUNCH_TOKENS() view returns(uint256)
func (_Flaunch *FlaunchCaller) MAXFAIRLAUNCHTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "MAX_FAIR_LAUNCH_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFAIRLAUNCHTOKENS is a free data retrieval call binding the contract method 0x97cd41a8.
//
// Solidity: function MAX_FAIR_LAUNCH_TOKENS() view returns(uint256)
func (_Flaunch *FlaunchSession) MAXFAIRLAUNCHTOKENS() (*big.Int, error) {
	return _Flaunch.Contract.MAXFAIRLAUNCHTOKENS(&_Flaunch.CallOpts)
}

// MAXFAIRLAUNCHTOKENS is a free data retrieval call binding the contract method 0x97cd41a8.
//
// Solidity: function MAX_FAIR_LAUNCH_TOKENS() view returns(uint256)
func (_Flaunch *FlaunchCallerSession) MAXFAIRLAUNCHTOKENS() (*big.Int, error) {
	return _Flaunch.Contract.MAXFAIRLAUNCHTOKENS(&_Flaunch.CallOpts)
}

// MAXSCHEDULEDURATION is a free data retrieval call binding the contract method 0x5b7426be.
//
// Solidity: function MAX_SCHEDULE_DURATION() view returns(uint256)
func (_Flaunch *FlaunchCaller) MAXSCHEDULEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "MAX_SCHEDULE_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSCHEDULEDURATION is a free data retrieval call binding the contract method 0x5b7426be.
//
// Solidity: function MAX_SCHEDULE_DURATION() view returns(uint256)
func (_Flaunch *FlaunchSession) MAXSCHEDULEDURATION() (*big.Int, error) {
	return _Flaunch.Contract.MAXSCHEDULEDURATION(&_Flaunch.CallOpts)
}

// MAXSCHEDULEDURATION is a free data retrieval call binding the contract method 0x5b7426be.
//
// Solidity: function MAX_SCHEDULE_DURATION() view returns(uint256)
func (_Flaunch *FlaunchCallerSession) MAXSCHEDULEDURATION() (*big.Int, error) {
	return _Flaunch.Contract.MAXSCHEDULEDURATION(&_Flaunch.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_Flaunch *FlaunchCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_Flaunch *FlaunchSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Flaunch.Contract.BalanceOf(&_Flaunch.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_Flaunch *FlaunchCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Flaunch.Contract.BalanceOf(&_Flaunch.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Flaunch *FlaunchCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Flaunch *FlaunchSession) BaseURI() (string, error) {
	return _Flaunch.Contract.BaseURI(&_Flaunch.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Flaunch *FlaunchCallerSession) BaseURI() (string, error) {
	return _Flaunch.Contract.BaseURI(&_Flaunch.CallOpts)
}

// BridgingStatus is a free data retrieval call binding the contract method 0x359687e4.
//
// Solidity: function bridgingStatus(uint256 _tokenId, uint256 _chainId) view returns(bool _started)
func (_Flaunch *FlaunchCaller) BridgingStatus(opts *bind.CallOpts, _tokenId *big.Int, _chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "bridgingStatus", _tokenId, _chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgingStatus is a free data retrieval call binding the contract method 0x359687e4.
//
// Solidity: function bridgingStatus(uint256 _tokenId, uint256 _chainId) view returns(bool _started)
func (_Flaunch *FlaunchSession) BridgingStatus(_tokenId *big.Int, _chainId *big.Int) (bool, error) {
	return _Flaunch.Contract.BridgingStatus(&_Flaunch.CallOpts, _tokenId, _chainId)
}

// BridgingStatus is a free data retrieval call binding the contract method 0x359687e4.
//
// Solidity: function bridgingStatus(uint256 _tokenId, uint256 _chainId) view returns(bool _started)
func (_Flaunch *FlaunchCallerSession) BridgingStatus(_tokenId *big.Int, _chainId *big.Int) (bool, error) {
	return _Flaunch.Contract.BridgingStatus(&_Flaunch.CallOpts, _tokenId, _chainId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address result)
func (_Flaunch *FlaunchCaller) GetApproved(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "getApproved", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address result)
func (_Flaunch *FlaunchSession) GetApproved(id *big.Int) (common.Address, error) {
	return _Flaunch.Contract.GetApproved(&_Flaunch.CallOpts, id)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address result)
func (_Flaunch *FlaunchCallerSession) GetApproved(id *big.Int) (common.Address, error) {
	return _Flaunch.Contract.GetApproved(&_Flaunch.CallOpts, id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool result)
func (_Flaunch *FlaunchCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool result)
func (_Flaunch *FlaunchSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Flaunch.Contract.IsApprovedForAll(&_Flaunch.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool result)
func (_Flaunch *FlaunchCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Flaunch.Contract.IsApprovedForAll(&_Flaunch.CallOpts, owner, operator)
}

// Memecoin is a free data retrieval call binding the contract method 0xe49c1854.
//
// Solidity: function memecoin(uint256 _tokenId) view returns(address)
func (_Flaunch *FlaunchCaller) Memecoin(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "memecoin", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Memecoin is a free data retrieval call binding the contract method 0xe49c1854.
//
// Solidity: function memecoin(uint256 _tokenId) view returns(address)
func (_Flaunch *FlaunchSession) Memecoin(_tokenId *big.Int) (common.Address, error) {
	return _Flaunch.Contract.Memecoin(&_Flaunch.CallOpts, _tokenId)
}

// Memecoin is a free data retrieval call binding the contract method 0xe49c1854.
//
// Solidity: function memecoin(uint256 _tokenId) view returns(address)
func (_Flaunch *FlaunchCallerSession) Memecoin(_tokenId *big.Int) (common.Address, error) {
	return _Flaunch.Contract.Memecoin(&_Flaunch.CallOpts, _tokenId)
}

// MemecoinImplementation is a free data retrieval call binding the contract method 0xbf989c78.
//
// Solidity: function memecoinImplementation() view returns(address)
func (_Flaunch *FlaunchCaller) MemecoinImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "memecoinImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MemecoinImplementation is a free data retrieval call binding the contract method 0xbf989c78.
//
// Solidity: function memecoinImplementation() view returns(address)
func (_Flaunch *FlaunchSession) MemecoinImplementation() (common.Address, error) {
	return _Flaunch.Contract.MemecoinImplementation(&_Flaunch.CallOpts)
}

// MemecoinImplementation is a free data retrieval call binding the contract method 0xbf989c78.
//
// Solidity: function memecoinImplementation() view returns(address)
func (_Flaunch *FlaunchCallerSession) MemecoinImplementation() (common.Address, error) {
	return _Flaunch.Contract.MemecoinImplementation(&_Flaunch.CallOpts)
}

// MemecoinTreasury is a free data retrieval call binding the contract method 0xf38f9550.
//
// Solidity: function memecoinTreasury(uint256 _tokenId) view returns(address)
func (_Flaunch *FlaunchCaller) MemecoinTreasury(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "memecoinTreasury", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MemecoinTreasury is a free data retrieval call binding the contract method 0xf38f9550.
//
// Solidity: function memecoinTreasury(uint256 _tokenId) view returns(address)
func (_Flaunch *FlaunchSession) MemecoinTreasury(_tokenId *big.Int) (common.Address, error) {
	return _Flaunch.Contract.MemecoinTreasury(&_Flaunch.CallOpts, _tokenId)
}

// MemecoinTreasury is a free data retrieval call binding the contract method 0xf38f9550.
//
// Solidity: function memecoinTreasury(uint256 _tokenId) view returns(address)
func (_Flaunch *FlaunchCallerSession) MemecoinTreasury(_tokenId *big.Int) (common.Address, error) {
	return _Flaunch.Contract.MemecoinTreasury(&_Flaunch.CallOpts, _tokenId)
}

// MemecoinTreasuryImplementation is a free data retrieval call binding the contract method 0x1a5de1df.
//
// Solidity: function memecoinTreasuryImplementation() view returns(address)
func (_Flaunch *FlaunchCaller) MemecoinTreasuryImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "memecoinTreasuryImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MemecoinTreasuryImplementation is a free data retrieval call binding the contract method 0x1a5de1df.
//
// Solidity: function memecoinTreasuryImplementation() view returns(address)
func (_Flaunch *FlaunchSession) MemecoinTreasuryImplementation() (common.Address, error) {
	return _Flaunch.Contract.MemecoinTreasuryImplementation(&_Flaunch.CallOpts)
}

// MemecoinTreasuryImplementation is a free data retrieval call binding the contract method 0x1a5de1df.
//
// Solidity: function memecoinTreasuryImplementation() view returns(address)
func (_Flaunch *FlaunchCallerSession) MemecoinTreasuryImplementation() (common.Address, error) {
	return _Flaunch.Contract.MemecoinTreasuryImplementation(&_Flaunch.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Flaunch *FlaunchCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Flaunch *FlaunchSession) Name() (string, error) {
	return _Flaunch.Contract.Name(&_Flaunch.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Flaunch *FlaunchCallerSession) Name() (string, error) {
	return _Flaunch.Contract.Name(&_Flaunch.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Flaunch *FlaunchCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Flaunch *FlaunchSession) NextTokenId() (*big.Int, error) {
	return _Flaunch.Contract.NextTokenId(&_Flaunch.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Flaunch *FlaunchCallerSession) NextTokenId() (*big.Int, error) {
	return _Flaunch.Contract.NextTokenId(&_Flaunch.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Flaunch *FlaunchCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Flaunch *FlaunchSession) Owner() (common.Address, error) {
	return _Flaunch.Contract.Owner(&_Flaunch.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Flaunch *FlaunchCallerSession) Owner() (common.Address, error) {
	return _Flaunch.Contract.Owner(&_Flaunch.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address result)
func (_Flaunch *FlaunchCaller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address result)
func (_Flaunch *FlaunchSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Flaunch.Contract.OwnerOf(&_Flaunch.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address result)
func (_Flaunch *FlaunchCallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Flaunch.Contract.OwnerOf(&_Flaunch.CallOpts, id)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Flaunch *FlaunchCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Flaunch *FlaunchSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Flaunch.Contract.OwnershipHandoverExpiresAt(&_Flaunch.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Flaunch *FlaunchCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Flaunch.Contract.OwnershipHandoverExpiresAt(&_Flaunch.CallOpts, pendingOwner)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_Flaunch *FlaunchCaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_Flaunch *FlaunchSession) PositionManager() (common.Address, error) {
	return _Flaunch.Contract.PositionManager(&_Flaunch.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_Flaunch *FlaunchCallerSession) PositionManager() (common.Address, error) {
	return _Flaunch.Contract.PositionManager(&_Flaunch.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool result)
func (_Flaunch *FlaunchCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool result)
func (_Flaunch *FlaunchSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Flaunch.Contract.SupportsInterface(&_Flaunch.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool result)
func (_Flaunch *FlaunchCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Flaunch.Contract.SupportsInterface(&_Flaunch.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Flaunch *FlaunchCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Flaunch *FlaunchSession) Symbol() (string, error) {
	return _Flaunch.Contract.Symbol(&_Flaunch.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Flaunch *FlaunchCallerSession) Symbol() (string, error) {
	return _Flaunch.Contract.Symbol(&_Flaunch.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x7ca31724.
//
// Solidity: function tokenId(address _memecoin) view returns(uint256 _tokenId)
func (_Flaunch *FlaunchCaller) TokenId(opts *bind.CallOpts, _memecoin common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "tokenId", _memecoin)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x7ca31724.
//
// Solidity: function tokenId(address _memecoin) view returns(uint256 _tokenId)
func (_Flaunch *FlaunchSession) TokenId(_memecoin common.Address) (*big.Int, error) {
	return _Flaunch.Contract.TokenId(&_Flaunch.CallOpts, _memecoin)
}

// TokenId is a free data retrieval call binding the contract method 0x7ca31724.
//
// Solidity: function tokenId(address _memecoin) view returns(uint256 _tokenId)
func (_Flaunch *FlaunchCallerSession) TokenId(_memecoin common.Address) (*big.Int, error) {
	return _Flaunch.Contract.TokenId(&_Flaunch.CallOpts, _memecoin)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Flaunch *FlaunchCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Flaunch.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Flaunch *FlaunchSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Flaunch.Contract.TokenURI(&_Flaunch.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Flaunch *FlaunchCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Flaunch.Contract.TokenURI(&_Flaunch.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address account, uint256 id) payable returns()
func (_Flaunch *FlaunchTransactor) Approve(opts *bind.TransactOpts, account common.Address, id *big.Int) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "approve", account, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address account, uint256 id) payable returns()
func (_Flaunch *FlaunchSession) Approve(account common.Address, id *big.Int) (*types.Transaction, error) {
	return _Flaunch.Contract.Approve(&_Flaunch.TransactOpts, account, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address account, uint256 id) payable returns()
func (_Flaunch *FlaunchTransactorSession) Approve(account common.Address, id *big.Int) (*types.Transaction, error) {
	return _Flaunch.Contract.Approve(&_Flaunch.TransactOpts, account, id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_Flaunch *FlaunchTransactor) Burn(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "burn", _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_Flaunch *FlaunchSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _Flaunch.Contract.Burn(&_Flaunch.TransactOpts, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_Flaunch *FlaunchTransactorSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _Flaunch.Contract.Burn(&_Flaunch.TransactOpts, _tokenId)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Flaunch *FlaunchTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Flaunch *FlaunchSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Flaunch.Contract.CancelOwnershipHandover(&_Flaunch.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Flaunch *FlaunchTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Flaunch.Contract.CancelOwnershipHandover(&_Flaunch.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Flaunch *FlaunchTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Flaunch *FlaunchSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Flaunch.Contract.CompleteOwnershipHandover(&_Flaunch.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Flaunch *FlaunchTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Flaunch.Contract.CompleteOwnershipHandover(&_Flaunch.TransactOpts, pendingOwner)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x2b31d94a.
//
// Solidity: function finalizeBridge(uint256 _tokenId, (string,string,string) _metadata) returns()
func (_Flaunch *FlaunchTransactor) FinalizeBridge(opts *bind.TransactOpts, _tokenId *big.Int, _metadata FlaunchMemecoinMetadata) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "finalizeBridge", _tokenId, _metadata)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x2b31d94a.
//
// Solidity: function finalizeBridge(uint256 _tokenId, (string,string,string) _metadata) returns()
func (_Flaunch *FlaunchSession) FinalizeBridge(_tokenId *big.Int, _metadata FlaunchMemecoinMetadata) (*types.Transaction, error) {
	return _Flaunch.Contract.FinalizeBridge(&_Flaunch.TransactOpts, _tokenId, _metadata)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x2b31d94a.
//
// Solidity: function finalizeBridge(uint256 _tokenId, (string,string,string) _metadata) returns()
func (_Flaunch *FlaunchTransactorSession) FinalizeBridge(_tokenId *big.Int, _metadata FlaunchMemecoinMetadata) (*types.Transaction, error) {
	return _Flaunch.Contract.FinalizeBridge(&_Flaunch.TransactOpts, _tokenId, _metadata)
}

// Flaunch is a paid mutator transaction binding the contract method 0x2a12e488.
//
// Solidity: function flaunch((string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params) returns(address memecoin_, address memecoinTreasury_, uint256 tokenId_)
func (_Flaunch *FlaunchTransactor) Flaunch(opts *bind.TransactOpts, _params PositionManagerFlaunchParams) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "flaunch", _params)
}

// Flaunch is a paid mutator transaction binding the contract method 0x2a12e488.
//
// Solidity: function flaunch((string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params) returns(address memecoin_, address memecoinTreasury_, uint256 tokenId_)
func (_Flaunch *FlaunchSession) Flaunch(_params PositionManagerFlaunchParams) (*types.Transaction, error) {
	return _Flaunch.Contract.Flaunch(&_Flaunch.TransactOpts, _params)
}

// Flaunch is a paid mutator transaction binding the contract method 0x2a12e488.
//
// Solidity: function flaunch((string,string,string,uint256,uint256,address,uint24,uint256,bytes,bytes) _params) returns(address memecoin_, address memecoinTreasury_, uint256 tokenId_)
func (_Flaunch *FlaunchTransactorSession) Flaunch(_params PositionManagerFlaunchParams) (*types.Transaction, error) {
	return _Flaunch.Contract.Flaunch(&_Flaunch.TransactOpts, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _positionManager, address _memecoinTreasuryImplementation) returns()
func (_Flaunch *FlaunchTransactor) Initialize(opts *bind.TransactOpts, _positionManager common.Address, _memecoinTreasuryImplementation common.Address) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "initialize", _positionManager, _memecoinTreasuryImplementation)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _positionManager, address _memecoinTreasuryImplementation) returns()
func (_Flaunch *FlaunchSession) Initialize(_positionManager common.Address, _memecoinTreasuryImplementation common.Address) (*types.Transaction, error) {
	return _Flaunch.Contract.Initialize(&_Flaunch.TransactOpts, _positionManager, _memecoinTreasuryImplementation)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _positionManager, address _memecoinTreasuryImplementation) returns()
func (_Flaunch *FlaunchTransactorSession) Initialize(_positionManager common.Address, _memecoinTreasuryImplementation common.Address) (*types.Transaction, error) {
	return _Flaunch.Contract.Initialize(&_Flaunch.TransactOpts, _positionManager, _memecoinTreasuryImplementation)
}

// InitializeBridge is a paid mutator transaction binding the contract method 0xe4b47c0b.
//
// Solidity: function initializeBridge(uint256 _tokenId, uint256 _chainId) returns()
func (_Flaunch *FlaunchTransactor) InitializeBridge(opts *bind.TransactOpts, _tokenId *big.Int, _chainId *big.Int) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "initializeBridge", _tokenId, _chainId)
}

// InitializeBridge is a paid mutator transaction binding the contract method 0xe4b47c0b.
//
// Solidity: function initializeBridge(uint256 _tokenId, uint256 _chainId) returns()
func (_Flaunch *FlaunchSession) InitializeBridge(_tokenId *big.Int, _chainId *big.Int) (*types.Transaction, error) {
	return _Flaunch.Contract.InitializeBridge(&_Flaunch.TransactOpts, _tokenId, _chainId)
}

// InitializeBridge is a paid mutator transaction binding the contract method 0xe4b47c0b.
//
// Solidity: function initializeBridge(uint256 _tokenId, uint256 _chainId) returns()
func (_Flaunch *FlaunchTransactorSession) InitializeBridge(_tokenId *big.Int, _chainId *big.Int) (*types.Transaction, error) {
	return _Flaunch.Contract.InitializeBridge(&_Flaunch.TransactOpts, _tokenId, _chainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Flaunch *FlaunchTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Flaunch *FlaunchSession) RenounceOwnership() (*types.Transaction, error) {
	return _Flaunch.Contract.RenounceOwnership(&_Flaunch.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Flaunch *FlaunchTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Flaunch.Contract.RenounceOwnership(&_Flaunch.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Flaunch *FlaunchTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Flaunch *FlaunchSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Flaunch.Contract.RequestOwnershipHandover(&_Flaunch.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Flaunch *FlaunchTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Flaunch.Contract.RequestOwnershipHandover(&_Flaunch.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) payable returns()
func (_Flaunch *FlaunchTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) payable returns()
func (_Flaunch *FlaunchSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Flaunch.Contract.SafeTransferFrom(&_Flaunch.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) payable returns()
func (_Flaunch *FlaunchTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Flaunch.Contract.SafeTransferFrom(&_Flaunch.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) payable returns()
func (_Flaunch *FlaunchTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) payable returns()
func (_Flaunch *FlaunchSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Flaunch.Contract.SafeTransferFrom0(&_Flaunch.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) payable returns()
func (_Flaunch *FlaunchTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Flaunch.Contract.SafeTransferFrom0(&_Flaunch.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool isApproved) returns()
func (_Flaunch *FlaunchTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, isApproved bool) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "setApprovalForAll", operator, isApproved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool isApproved) returns()
func (_Flaunch *FlaunchSession) SetApprovalForAll(operator common.Address, isApproved bool) (*types.Transaction, error) {
	return _Flaunch.Contract.SetApprovalForAll(&_Flaunch.TransactOpts, operator, isApproved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool isApproved) returns()
func (_Flaunch *FlaunchTransactorSession) SetApprovalForAll(operator common.Address, isApproved bool) (*types.Transaction, error) {
	return _Flaunch.Contract.SetApprovalForAll(&_Flaunch.TransactOpts, operator, isApproved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Flaunch *FlaunchTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Flaunch *FlaunchSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Flaunch.Contract.SetBaseURI(&_Flaunch.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Flaunch *FlaunchTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Flaunch.Contract.SetBaseURI(&_Flaunch.TransactOpts, _baseURI)
}

// SetMemecoinMetadata is a paid mutator transaction binding the contract method 0x829f042a.
//
// Solidity: function setMemecoinMetadata(address _memecoin, string name_, string symbol_) returns()
func (_Flaunch *FlaunchTransactor) SetMemecoinMetadata(opts *bind.TransactOpts, _memecoin common.Address, name_ string, symbol_ string) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "setMemecoinMetadata", _memecoin, name_, symbol_)
}

// SetMemecoinMetadata is a paid mutator transaction binding the contract method 0x829f042a.
//
// Solidity: function setMemecoinMetadata(address _memecoin, string name_, string symbol_) returns()
func (_Flaunch *FlaunchSession) SetMemecoinMetadata(_memecoin common.Address, name_ string, symbol_ string) (*types.Transaction, error) {
	return _Flaunch.Contract.SetMemecoinMetadata(&_Flaunch.TransactOpts, _memecoin, name_, symbol_)
}

// SetMemecoinMetadata is a paid mutator transaction binding the contract method 0x829f042a.
//
// Solidity: function setMemecoinMetadata(address _memecoin, string name_, string symbol_) returns()
func (_Flaunch *FlaunchTransactorSession) SetMemecoinMetadata(_memecoin common.Address, name_ string, symbol_ string) (*types.Transaction, error) {
	return _Flaunch.Contract.SetMemecoinMetadata(&_Flaunch.TransactOpts, _memecoin, name_, symbol_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) payable returns()
func (_Flaunch *FlaunchTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) payable returns()
func (_Flaunch *FlaunchSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Flaunch.Contract.TransferFrom(&_Flaunch.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) payable returns()
func (_Flaunch *FlaunchTransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Flaunch.Contract.TransferFrom(&_Flaunch.TransactOpts, from, to, id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Flaunch *FlaunchTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Flaunch.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Flaunch *FlaunchSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Flaunch.Contract.TransferOwnership(&_Flaunch.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Flaunch *FlaunchTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Flaunch.Contract.TransferOwnership(&_Flaunch.TransactOpts, newOwner)
}

// FlaunchApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Flaunch contract.
type FlaunchApprovalIterator struct {
	Event *FlaunchApproval // Event containing the contract specifics and raw log

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
func (it *FlaunchApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlaunchApproval)
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
		it.Event = new(FlaunchApproval)
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
func (it *FlaunchApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlaunchApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlaunchApproval represents a Approval event raised by the Flaunch contract.
type FlaunchApproval struct {
	Owner   common.Address
	Account common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed account, uint256 indexed id)
func (_Flaunch *FlaunchFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, account []common.Address, id []*big.Int) (*FlaunchApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Flaunch.contract.FilterLogs(opts, "Approval", ownerRule, accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return &FlaunchApprovalIterator{contract: _Flaunch.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed account, uint256 indexed id)
func (_Flaunch *FlaunchFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FlaunchApproval, owner []common.Address, account []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Flaunch.contract.WatchLogs(opts, "Approval", ownerRule, accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlaunchApproval)
				if err := _Flaunch.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed account, uint256 indexed id)
func (_Flaunch *FlaunchFilterer) ParseApproval(log types.Log) (*FlaunchApproval, error) {
	event := new(FlaunchApproval)
	if err := _Flaunch.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlaunchApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Flaunch contract.
type FlaunchApprovalForAllIterator struct {
	Event *FlaunchApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FlaunchApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlaunchApprovalForAll)
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
		it.Event = new(FlaunchApprovalForAll)
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
func (it *FlaunchApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlaunchApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlaunchApprovalForAll represents a ApprovalForAll event raised by the Flaunch contract.
type FlaunchApprovalForAll struct {
	Owner      common.Address
	Operator   common.Address
	IsApproved bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool isApproved)
func (_Flaunch *FlaunchFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*FlaunchApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Flaunch.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FlaunchApprovalForAllIterator{contract: _Flaunch.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool isApproved)
func (_Flaunch *FlaunchFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FlaunchApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Flaunch.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlaunchApprovalForAll)
				if err := _Flaunch.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool isApproved)
func (_Flaunch *FlaunchFilterer) ParseApprovalForAll(log types.Log) (*FlaunchApprovalForAll, error) {
	event := new(FlaunchApprovalForAll)
	if err := _Flaunch.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlaunchInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Flaunch contract.
type FlaunchInitializedIterator struct {
	Event *FlaunchInitialized // Event containing the contract specifics and raw log

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
func (it *FlaunchInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlaunchInitialized)
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
		it.Event = new(FlaunchInitialized)
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
func (it *FlaunchInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlaunchInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlaunchInitialized represents a Initialized event raised by the Flaunch contract.
type FlaunchInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Flaunch *FlaunchFilterer) FilterInitialized(opts *bind.FilterOpts) (*FlaunchInitializedIterator, error) {

	logs, sub, err := _Flaunch.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FlaunchInitializedIterator{contract: _Flaunch.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Flaunch *FlaunchFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FlaunchInitialized) (event.Subscription, error) {

	logs, sub, err := _Flaunch.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlaunchInitialized)
				if err := _Flaunch.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Flaunch *FlaunchFilterer) ParseInitialized(log types.Log) (*FlaunchInitialized, error) {
	event := new(FlaunchInitialized)
	if err := _Flaunch.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlaunchOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the Flaunch contract.
type FlaunchOwnershipHandoverCanceledIterator struct {
	Event *FlaunchOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *FlaunchOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlaunchOwnershipHandoverCanceled)
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
		it.Event = new(FlaunchOwnershipHandoverCanceled)
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
func (it *FlaunchOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlaunchOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlaunchOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Flaunch contract.
type FlaunchOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Flaunch *FlaunchFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*FlaunchOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Flaunch.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FlaunchOwnershipHandoverCanceledIterator{contract: _Flaunch.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Flaunch *FlaunchFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *FlaunchOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Flaunch.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlaunchOwnershipHandoverCanceled)
				if err := _Flaunch.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_Flaunch *FlaunchFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*FlaunchOwnershipHandoverCanceled, error) {
	event := new(FlaunchOwnershipHandoverCanceled)
	if err := _Flaunch.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlaunchOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the Flaunch contract.
type FlaunchOwnershipHandoverRequestedIterator struct {
	Event *FlaunchOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *FlaunchOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlaunchOwnershipHandoverRequested)
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
		it.Event = new(FlaunchOwnershipHandoverRequested)
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
func (it *FlaunchOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlaunchOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlaunchOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Flaunch contract.
type FlaunchOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Flaunch *FlaunchFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*FlaunchOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Flaunch.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FlaunchOwnershipHandoverRequestedIterator{contract: _Flaunch.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Flaunch *FlaunchFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *FlaunchOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Flaunch.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlaunchOwnershipHandoverRequested)
				if err := _Flaunch.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_Flaunch *FlaunchFilterer) ParseOwnershipHandoverRequested(log types.Log) (*FlaunchOwnershipHandoverRequested, error) {
	event := new(FlaunchOwnershipHandoverRequested)
	if err := _Flaunch.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlaunchOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Flaunch contract.
type FlaunchOwnershipTransferredIterator struct {
	Event *FlaunchOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FlaunchOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlaunchOwnershipTransferred)
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
		it.Event = new(FlaunchOwnershipTransferred)
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
func (it *FlaunchOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlaunchOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlaunchOwnershipTransferred represents a OwnershipTransferred event raised by the Flaunch contract.
type FlaunchOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Flaunch *FlaunchFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*FlaunchOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Flaunch.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FlaunchOwnershipTransferredIterator{contract: _Flaunch.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Flaunch *FlaunchFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FlaunchOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Flaunch.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlaunchOwnershipTransferred)
				if err := _Flaunch.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Flaunch *FlaunchFilterer) ParseOwnershipTransferred(log types.Log) (*FlaunchOwnershipTransferred, error) {
	event := new(FlaunchOwnershipTransferred)
	if err := _Flaunch.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlaunchTokenBridgedIterator is returned from FilterTokenBridged and is used to iterate over the raw logs and unpacked data for TokenBridged events raised by the Flaunch contract.
type FlaunchTokenBridgedIterator struct {
	Event *FlaunchTokenBridged // Event containing the contract specifics and raw log

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
func (it *FlaunchTokenBridgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlaunchTokenBridged)
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
		it.Event = new(FlaunchTokenBridged)
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
func (it *FlaunchTokenBridgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlaunchTokenBridgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlaunchTokenBridged represents a TokenBridged event raised by the Flaunch contract.
type FlaunchTokenBridged struct {
	TokenId       *big.Int
	ChainId       *big.Int
	Memecoin      common.Address
	MessageSource *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBridged is a free log retrieval operation binding the contract event 0xd953cced6d6de41e29efa3f395f6815b35e5a0bb1bff016cb796b95e34a93254.
//
// Solidity: event TokenBridged(uint256 _tokenId, uint256 _chainId, address _memecoin, uint256 _messageSource)
func (_Flaunch *FlaunchFilterer) FilterTokenBridged(opts *bind.FilterOpts) (*FlaunchTokenBridgedIterator, error) {

	logs, sub, err := _Flaunch.contract.FilterLogs(opts, "TokenBridged")
	if err != nil {
		return nil, err
	}
	return &FlaunchTokenBridgedIterator{contract: _Flaunch.contract, event: "TokenBridged", logs: logs, sub: sub}, nil
}

// WatchTokenBridged is a free log subscription operation binding the contract event 0xd953cced6d6de41e29efa3f395f6815b35e5a0bb1bff016cb796b95e34a93254.
//
// Solidity: event TokenBridged(uint256 _tokenId, uint256 _chainId, address _memecoin, uint256 _messageSource)
func (_Flaunch *FlaunchFilterer) WatchTokenBridged(opts *bind.WatchOpts, sink chan<- *FlaunchTokenBridged) (event.Subscription, error) {

	logs, sub, err := _Flaunch.contract.WatchLogs(opts, "TokenBridged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlaunchTokenBridged)
				if err := _Flaunch.contract.UnpackLog(event, "TokenBridged", log); err != nil {
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

// ParseTokenBridged is a log parse operation binding the contract event 0xd953cced6d6de41e29efa3f395f6815b35e5a0bb1bff016cb796b95e34a93254.
//
// Solidity: event TokenBridged(uint256 _tokenId, uint256 _chainId, address _memecoin, uint256 _messageSource)
func (_Flaunch *FlaunchFilterer) ParseTokenBridged(log types.Log) (*FlaunchTokenBridged, error) {
	event := new(FlaunchTokenBridged)
	if err := _Flaunch.contract.UnpackLog(event, "TokenBridged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlaunchTokenBridgingIterator is returned from FilterTokenBridging and is used to iterate over the raw logs and unpacked data for TokenBridging events raised by the Flaunch contract.
type FlaunchTokenBridgingIterator struct {
	Event *FlaunchTokenBridging // Event containing the contract specifics and raw log

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
func (it *FlaunchTokenBridgingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlaunchTokenBridging)
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
		it.Event = new(FlaunchTokenBridging)
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
func (it *FlaunchTokenBridgingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlaunchTokenBridgingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlaunchTokenBridging represents a TokenBridging event raised by the Flaunch contract.
type FlaunchTokenBridging struct {
	TokenId  *big.Int
	ChainId  *big.Int
	Memecoin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenBridging is a free log retrieval operation binding the contract event 0xf2a584a2309fee3f807e426ce11328f237b9bc191989cbc6f62cc5fee1d6b46f.
//
// Solidity: event TokenBridging(uint256 _tokenId, uint256 _chainId, address _memecoin)
func (_Flaunch *FlaunchFilterer) FilterTokenBridging(opts *bind.FilterOpts) (*FlaunchTokenBridgingIterator, error) {

	logs, sub, err := _Flaunch.contract.FilterLogs(opts, "TokenBridging")
	if err != nil {
		return nil, err
	}
	return &FlaunchTokenBridgingIterator{contract: _Flaunch.contract, event: "TokenBridging", logs: logs, sub: sub}, nil
}

// WatchTokenBridging is a free log subscription operation binding the contract event 0xf2a584a2309fee3f807e426ce11328f237b9bc191989cbc6f62cc5fee1d6b46f.
//
// Solidity: event TokenBridging(uint256 _tokenId, uint256 _chainId, address _memecoin)
func (_Flaunch *FlaunchFilterer) WatchTokenBridging(opts *bind.WatchOpts, sink chan<- *FlaunchTokenBridging) (event.Subscription, error) {

	logs, sub, err := _Flaunch.contract.WatchLogs(opts, "TokenBridging")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlaunchTokenBridging)
				if err := _Flaunch.contract.UnpackLog(event, "TokenBridging", log); err != nil {
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

// ParseTokenBridging is a log parse operation binding the contract event 0xf2a584a2309fee3f807e426ce11328f237b9bc191989cbc6f62cc5fee1d6b46f.
//
// Solidity: event TokenBridging(uint256 _tokenId, uint256 _chainId, address _memecoin)
func (_Flaunch *FlaunchFilterer) ParseTokenBridging(log types.Log) (*FlaunchTokenBridging, error) {
	event := new(FlaunchTokenBridging)
	if err := _Flaunch.contract.UnpackLog(event, "TokenBridging", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlaunchTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Flaunch contract.
type FlaunchTransferIterator struct {
	Event *FlaunchTransfer // Event containing the contract specifics and raw log

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
func (it *FlaunchTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlaunchTransfer)
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
		it.Event = new(FlaunchTransfer)
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
func (it *FlaunchTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlaunchTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlaunchTransfer represents a Transfer event raised by the Flaunch contract.
type FlaunchTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Flaunch *FlaunchFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*FlaunchTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Flaunch.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &FlaunchTransferIterator{contract: _Flaunch.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Flaunch *FlaunchFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FlaunchTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Flaunch.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlaunchTransfer)
				if err := _Flaunch.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Flaunch *FlaunchFilterer) ParseTransfer(log types.Log) (*FlaunchTransfer, error) {
	event := new(FlaunchTransfer)
	if err := _Flaunch.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
