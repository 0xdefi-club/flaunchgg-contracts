// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internal_swap_pool

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

// InternalSwapPoolMetaData contains all meta data concerning the InternalSwapPool contract.
var InternalSwapPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"poolFees\",\"inputs\":[{\"name\":\"_poolKey\",\"type\":\"tuple\",\"internalType\":\"structPoolKey\",\"components\":[{\"name\":\"currency0\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"currency1\",\"type\":\"address\",\"internalType\":\"Currency\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"hooks\",\"type\":\"address\",\"internalType\":\"contractIHooks\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInternalSwapPool.ClaimableFees\",\"components\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"PoolFeesDistributed\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_donateAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_creatorAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_bidWallAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_governanceAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_protocolAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeesReceived\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"_amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeesSwapped\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"PoolId\"},{\"name\":\"zeroForOne\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"_amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// InternalSwapPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use InternalSwapPoolMetaData.ABI instead.
var InternalSwapPoolABI = InternalSwapPoolMetaData.ABI

// InternalSwapPool is an auto generated Go binding around an Ethereum contract.
type InternalSwapPool struct {
	InternalSwapPoolCaller     // Read-only binding to the contract
	InternalSwapPoolTransactor // Write-only binding to the contract
	InternalSwapPoolFilterer   // Log filterer for contract events
}

// InternalSwapPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type InternalSwapPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InternalSwapPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InternalSwapPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InternalSwapPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InternalSwapPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InternalSwapPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InternalSwapPoolSession struct {
	Contract     *InternalSwapPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InternalSwapPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InternalSwapPoolCallerSession struct {
	Contract *InternalSwapPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// InternalSwapPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InternalSwapPoolTransactorSession struct {
	Contract     *InternalSwapPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// InternalSwapPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type InternalSwapPoolRaw struct {
	Contract *InternalSwapPool // Generic contract binding to access the raw methods on
}

// InternalSwapPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InternalSwapPoolCallerRaw struct {
	Contract *InternalSwapPoolCaller // Generic read-only contract binding to access the raw methods on
}

// InternalSwapPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InternalSwapPoolTransactorRaw struct {
	Contract *InternalSwapPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInternalSwapPool creates a new instance of InternalSwapPool, bound to a specific deployed contract.
func NewInternalSwapPool(address common.Address, backend bind.ContractBackend) (*InternalSwapPool, error) {
	contract, err := bindInternalSwapPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InternalSwapPool{InternalSwapPoolCaller: InternalSwapPoolCaller{contract: contract}, InternalSwapPoolTransactor: InternalSwapPoolTransactor{contract: contract}, InternalSwapPoolFilterer: InternalSwapPoolFilterer{contract: contract}}, nil
}

// NewInternalSwapPoolCaller creates a new read-only instance of InternalSwapPool, bound to a specific deployed contract.
func NewInternalSwapPoolCaller(address common.Address, caller bind.ContractCaller) (*InternalSwapPoolCaller, error) {
	contract, err := bindInternalSwapPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InternalSwapPoolCaller{contract: contract}, nil
}

// NewInternalSwapPoolTransactor creates a new write-only instance of InternalSwapPool, bound to a specific deployed contract.
func NewInternalSwapPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*InternalSwapPoolTransactor, error) {
	contract, err := bindInternalSwapPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InternalSwapPoolTransactor{contract: contract}, nil
}

// NewInternalSwapPoolFilterer creates a new log filterer instance of InternalSwapPool, bound to a specific deployed contract.
func NewInternalSwapPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*InternalSwapPoolFilterer, error) {
	contract, err := bindInternalSwapPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InternalSwapPoolFilterer{contract: contract}, nil
}

// bindInternalSwapPool binds a generic wrapper to an already deployed contract.
func bindInternalSwapPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InternalSwapPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InternalSwapPool *InternalSwapPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InternalSwapPool.Contract.InternalSwapPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InternalSwapPool *InternalSwapPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InternalSwapPool.Contract.InternalSwapPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InternalSwapPool *InternalSwapPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InternalSwapPool.Contract.InternalSwapPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InternalSwapPool *InternalSwapPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InternalSwapPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InternalSwapPool *InternalSwapPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InternalSwapPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InternalSwapPool *InternalSwapPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InternalSwapPool.Contract.contract.Transact(opts, method, params...)
}

// PoolFees is a free data retrieval call binding the contract method 0xddb475d5.
//
// Solidity: function poolFees((address,address,uint24,int24,address) _poolKey) view returns((uint256,uint256))
func (_InternalSwapPool *InternalSwapPoolCaller) PoolFees(opts *bind.CallOpts, _poolKey PoolKey) (InternalSwapPoolClaimableFees, error) {
	var out []interface{}
	err := _InternalSwapPool.contract.Call(opts, &out, "poolFees", _poolKey)

	if err != nil {
		return *new(InternalSwapPoolClaimableFees), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalSwapPoolClaimableFees)).(*InternalSwapPoolClaimableFees)

	return out0, err

}

// PoolFees is a free data retrieval call binding the contract method 0xddb475d5.
//
// Solidity: function poolFees((address,address,uint24,int24,address) _poolKey) view returns((uint256,uint256))
func (_InternalSwapPool *InternalSwapPoolSession) PoolFees(_poolKey PoolKey) (InternalSwapPoolClaimableFees, error) {
	return _InternalSwapPool.Contract.PoolFees(&_InternalSwapPool.CallOpts, _poolKey)
}

// PoolFees is a free data retrieval call binding the contract method 0xddb475d5.
//
// Solidity: function poolFees((address,address,uint24,int24,address) _poolKey) view returns((uint256,uint256))
func (_InternalSwapPool *InternalSwapPoolCallerSession) PoolFees(_poolKey PoolKey) (InternalSwapPoolClaimableFees, error) {
	return _InternalSwapPool.Contract.PoolFees(&_InternalSwapPool.CallOpts, _poolKey)
}

// InternalSwapPoolPoolFeesDistributedIterator is returned from FilterPoolFeesDistributed and is used to iterate over the raw logs and unpacked data for PoolFeesDistributed events raised by the InternalSwapPool contract.
type InternalSwapPoolPoolFeesDistributedIterator struct {
	Event *InternalSwapPoolPoolFeesDistributed // Event containing the contract specifics and raw log

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
func (it *InternalSwapPoolPoolFeesDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InternalSwapPoolPoolFeesDistributed)
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
		it.Event = new(InternalSwapPoolPoolFeesDistributed)
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
func (it *InternalSwapPoolPoolFeesDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InternalSwapPoolPoolFeesDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InternalSwapPoolPoolFeesDistributed represents a PoolFeesDistributed event raised by the InternalSwapPool contract.
type InternalSwapPoolPoolFeesDistributed struct {
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
func (_InternalSwapPool *InternalSwapPoolFilterer) FilterPoolFeesDistributed(opts *bind.FilterOpts, _poolId [][32]byte) (*InternalSwapPoolPoolFeesDistributedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _InternalSwapPool.contract.FilterLogs(opts, "PoolFeesDistributed", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &InternalSwapPoolPoolFeesDistributedIterator{contract: _InternalSwapPool.contract, event: "PoolFeesDistributed", logs: logs, sub: sub}, nil
}

// WatchPoolFeesDistributed is a free log subscription operation binding the contract event 0xc7241a69d3660bdfe5f36bdcca3b2da1fe8844366e46adb58be95171ab0665ad.
//
// Solidity: event PoolFeesDistributed(bytes32 indexed _poolId, uint256 _donateAmount, uint256 _creatorAmount, uint256 _bidWallAmount, uint256 _governanceAmount, uint256 _protocolAmount)
func (_InternalSwapPool *InternalSwapPoolFilterer) WatchPoolFeesDistributed(opts *bind.WatchOpts, sink chan<- *InternalSwapPoolPoolFeesDistributed, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _InternalSwapPool.contract.WatchLogs(opts, "PoolFeesDistributed", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InternalSwapPoolPoolFeesDistributed)
				if err := _InternalSwapPool.contract.UnpackLog(event, "PoolFeesDistributed", log); err != nil {
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
func (_InternalSwapPool *InternalSwapPoolFilterer) ParsePoolFeesDistributed(log types.Log) (*InternalSwapPoolPoolFeesDistributed, error) {
	event := new(InternalSwapPoolPoolFeesDistributed)
	if err := _InternalSwapPool.contract.UnpackLog(event, "PoolFeesDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InternalSwapPoolPoolFeesReceivedIterator is returned from FilterPoolFeesReceived and is used to iterate over the raw logs and unpacked data for PoolFeesReceived events raised by the InternalSwapPool contract.
type InternalSwapPoolPoolFeesReceivedIterator struct {
	Event *InternalSwapPoolPoolFeesReceived // Event containing the contract specifics and raw log

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
func (it *InternalSwapPoolPoolFeesReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InternalSwapPoolPoolFeesReceived)
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
		it.Event = new(InternalSwapPoolPoolFeesReceived)
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
func (it *InternalSwapPoolPoolFeesReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InternalSwapPoolPoolFeesReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InternalSwapPoolPoolFeesReceived represents a PoolFeesReceived event raised by the InternalSwapPool contract.
type InternalSwapPoolPoolFeesReceived struct {
	PoolId  [32]byte
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolFeesReceived is a free log retrieval operation binding the contract event 0xa245a9a38e8877add82f0a82c13e062ab3df16a26121977ddcca8827d46c690a.
//
// Solidity: event PoolFeesReceived(bytes32 indexed _poolId, uint256 _amount0, uint256 _amount1)
func (_InternalSwapPool *InternalSwapPoolFilterer) FilterPoolFeesReceived(opts *bind.FilterOpts, _poolId [][32]byte) (*InternalSwapPoolPoolFeesReceivedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _InternalSwapPool.contract.FilterLogs(opts, "PoolFeesReceived", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &InternalSwapPoolPoolFeesReceivedIterator{contract: _InternalSwapPool.contract, event: "PoolFeesReceived", logs: logs, sub: sub}, nil
}

// WatchPoolFeesReceived is a free log subscription operation binding the contract event 0xa245a9a38e8877add82f0a82c13e062ab3df16a26121977ddcca8827d46c690a.
//
// Solidity: event PoolFeesReceived(bytes32 indexed _poolId, uint256 _amount0, uint256 _amount1)
func (_InternalSwapPool *InternalSwapPoolFilterer) WatchPoolFeesReceived(opts *bind.WatchOpts, sink chan<- *InternalSwapPoolPoolFeesReceived, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _InternalSwapPool.contract.WatchLogs(opts, "PoolFeesReceived", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InternalSwapPoolPoolFeesReceived)
				if err := _InternalSwapPool.contract.UnpackLog(event, "PoolFeesReceived", log); err != nil {
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
func (_InternalSwapPool *InternalSwapPoolFilterer) ParsePoolFeesReceived(log types.Log) (*InternalSwapPoolPoolFeesReceived, error) {
	event := new(InternalSwapPoolPoolFeesReceived)
	if err := _InternalSwapPool.contract.UnpackLog(event, "PoolFeesReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InternalSwapPoolPoolFeesSwappedIterator is returned from FilterPoolFeesSwapped and is used to iterate over the raw logs and unpacked data for PoolFeesSwapped events raised by the InternalSwapPool contract.
type InternalSwapPoolPoolFeesSwappedIterator struct {
	Event *InternalSwapPoolPoolFeesSwapped // Event containing the contract specifics and raw log

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
func (it *InternalSwapPoolPoolFeesSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InternalSwapPoolPoolFeesSwapped)
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
		it.Event = new(InternalSwapPoolPoolFeesSwapped)
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
func (it *InternalSwapPoolPoolFeesSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InternalSwapPoolPoolFeesSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InternalSwapPoolPoolFeesSwapped represents a PoolFeesSwapped event raised by the InternalSwapPool contract.
type InternalSwapPoolPoolFeesSwapped struct {
	PoolId     [32]byte
	ZeroForOne bool
	Amount0    *big.Int
	Amount1    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolFeesSwapped is a free log retrieval operation binding the contract event 0xce97caf4fd0295de9544b52b4b9e79fe34c370bebb6fb71bc5baae9a70207968.
//
// Solidity: event PoolFeesSwapped(bytes32 indexed _poolId, bool zeroForOne, uint256 _amount0, uint256 _amount1)
func (_InternalSwapPool *InternalSwapPoolFilterer) FilterPoolFeesSwapped(opts *bind.FilterOpts, _poolId [][32]byte) (*InternalSwapPoolPoolFeesSwappedIterator, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _InternalSwapPool.contract.FilterLogs(opts, "PoolFeesSwapped", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return &InternalSwapPoolPoolFeesSwappedIterator{contract: _InternalSwapPool.contract, event: "PoolFeesSwapped", logs: logs, sub: sub}, nil
}

// WatchPoolFeesSwapped is a free log subscription operation binding the contract event 0xce97caf4fd0295de9544b52b4b9e79fe34c370bebb6fb71bc5baae9a70207968.
//
// Solidity: event PoolFeesSwapped(bytes32 indexed _poolId, bool zeroForOne, uint256 _amount0, uint256 _amount1)
func (_InternalSwapPool *InternalSwapPoolFilterer) WatchPoolFeesSwapped(opts *bind.WatchOpts, sink chan<- *InternalSwapPoolPoolFeesSwapped, _poolId [][32]byte) (event.Subscription, error) {

	var _poolIdRule []interface{}
	for _, _poolIdItem := range _poolId {
		_poolIdRule = append(_poolIdRule, _poolIdItem)
	}

	logs, sub, err := _InternalSwapPool.contract.WatchLogs(opts, "PoolFeesSwapped", _poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InternalSwapPoolPoolFeesSwapped)
				if err := _InternalSwapPool.contract.UnpackLog(event, "PoolFeesSwapped", log); err != nil {
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
func (_InternalSwapPool *InternalSwapPoolFilterer) ParsePoolFeesSwapped(log types.Log) (*InternalSwapPoolPoolFeesSwapped, error) {
	event := new(InternalSwapPoolPoolFeesSwapped)
	if err := _InternalSwapPool.contract.UnpackLog(event, "PoolFeesSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
