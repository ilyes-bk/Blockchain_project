// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"EnergyListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"EnergyPurchased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"energies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"listEnergy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"purchaseEnergy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Energies is a free data retrieval call binding the contract method 0x61588a4e.
//
// Solidity: function energies(uint256 ) view returns(uint256 id, address seller, uint256 amount, uint256 price, bool sold)
func (_Main *MainCaller) Energies(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Seller common.Address
	Amount *big.Int
	Price  *big.Int
	Sold   bool
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "energies", arg0)

	outstruct := new(struct {
		Id     *big.Int
		Seller common.Address
		Amount *big.Int
		Price  *big.Int
		Sold   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Sold = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Energies is a free data retrieval call binding the contract method 0x61588a4e.
//
// Solidity: function energies(uint256 ) view returns(uint256 id, address seller, uint256 amount, uint256 price, bool sold)
func (_Main *MainSession) Energies(arg0 *big.Int) (struct {
	Id     *big.Int
	Seller common.Address
	Amount *big.Int
	Price  *big.Int
	Sold   bool
}, error) {
	return _Main.Contract.Energies(&_Main.CallOpts, arg0)
}

// Energies is a free data retrieval call binding the contract method 0x61588a4e.
//
// Solidity: function energies(uint256 ) view returns(uint256 id, address seller, uint256 amount, uint256 price, bool sold)
func (_Main *MainCallerSession) Energies(arg0 *big.Int) (struct {
	Id     *big.Int
	Seller common.Address
	Amount *big.Int
	Price  *big.Int
	Sold   bool
}, error) {
	return _Main.Contract.Energies(&_Main.CallOpts, arg0)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Main *MainCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Main *MainSession) NextId() (*big.Int, error) {
	return _Main.Contract.NextId(&_Main.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Main *MainCallerSession) NextId() (*big.Int, error) {
	return _Main.Contract.NextId(&_Main.CallOpts)
}

// ListEnergy is a paid mutator transaction binding the contract method 0xc8a9fb5e.
//
// Solidity: function listEnergy(uint256 amount, uint256 price) returns()
func (_Main *MainTransactor) ListEnergy(opts *bind.TransactOpts, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "listEnergy", amount, price)
}

// ListEnergy is a paid mutator transaction binding the contract method 0xc8a9fb5e.
//
// Solidity: function listEnergy(uint256 amount, uint256 price) returns()
func (_Main *MainSession) ListEnergy(amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ListEnergy(&_Main.TransactOpts, amount, price)
}

// ListEnergy is a paid mutator transaction binding the contract method 0xc8a9fb5e.
//
// Solidity: function listEnergy(uint256 amount, uint256 price) returns()
func (_Main *MainTransactorSession) ListEnergy(amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ListEnergy(&_Main.TransactOpts, amount, price)
}

// PurchaseEnergy is a paid mutator transaction binding the contract method 0x6afe2520.
//
// Solidity: function purchaseEnergy(uint256 id) payable returns()
func (_Main *MainTransactor) PurchaseEnergy(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "purchaseEnergy", id)
}

// PurchaseEnergy is a paid mutator transaction binding the contract method 0x6afe2520.
//
// Solidity: function purchaseEnergy(uint256 id) payable returns()
func (_Main *MainSession) PurchaseEnergy(id *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PurchaseEnergy(&_Main.TransactOpts, id)
}

// PurchaseEnergy is a paid mutator transaction binding the contract method 0x6afe2520.
//
// Solidity: function purchaseEnergy(uint256 id) payable returns()
func (_Main *MainTransactorSession) PurchaseEnergy(id *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PurchaseEnergy(&_Main.TransactOpts, id)
}

// MainEnergyListedIterator is returned from FilterEnergyListed and is used to iterate over the raw logs and unpacked data for EnergyListed events raised by the Main contract.
type MainEnergyListedIterator struct {
	Event *MainEnergyListed // Event containing the contract specifics and raw log

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
func (it *MainEnergyListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainEnergyListed)
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
		it.Event = new(MainEnergyListed)
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
func (it *MainEnergyListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainEnergyListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainEnergyListed represents a EnergyListed event raised by the Main contract.
type MainEnergyListed struct {
	Id     *big.Int
	Seller common.Address
	Amount *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnergyListed is a free log retrieval operation binding the contract event 0xdce237255d5bd98b57c816861c3b76c8e47e10382084f482135e29cb83a534be.
//
// Solidity: event EnergyListed(uint256 id, address seller, uint256 amount, uint256 price)
func (_Main *MainFilterer) FilterEnergyListed(opts *bind.FilterOpts) (*MainEnergyListedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "EnergyListed")
	if err != nil {
		return nil, err
	}
	return &MainEnergyListedIterator{contract: _Main.contract, event: "EnergyListed", logs: logs, sub: sub}, nil
}

// WatchEnergyListed is a free log subscription operation binding the contract event 0xdce237255d5bd98b57c816861c3b76c8e47e10382084f482135e29cb83a534be.
//
// Solidity: event EnergyListed(uint256 id, address seller, uint256 amount, uint256 price)
func (_Main *MainFilterer) WatchEnergyListed(opts *bind.WatchOpts, sink chan<- *MainEnergyListed) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "EnergyListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainEnergyListed)
				if err := _Main.contract.UnpackLog(event, "EnergyListed", log); err != nil {
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

// ParseEnergyListed is a log parse operation binding the contract event 0xdce237255d5bd98b57c816861c3b76c8e47e10382084f482135e29cb83a534be.
//
// Solidity: event EnergyListed(uint256 id, address seller, uint256 amount, uint256 price)
func (_Main *MainFilterer) ParseEnergyListed(log types.Log) (*MainEnergyListed, error) {
	event := new(MainEnergyListed)
	if err := _Main.contract.UnpackLog(event, "EnergyListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainEnergyPurchasedIterator is returned from FilterEnergyPurchased and is used to iterate over the raw logs and unpacked data for EnergyPurchased events raised by the Main contract.
type MainEnergyPurchasedIterator struct {
	Event *MainEnergyPurchased // Event containing the contract specifics and raw log

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
func (it *MainEnergyPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainEnergyPurchased)
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
		it.Event = new(MainEnergyPurchased)
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
func (it *MainEnergyPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainEnergyPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainEnergyPurchased represents a EnergyPurchased event raised by the Main contract.
type MainEnergyPurchased struct {
	Id    *big.Int
	Buyer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEnergyPurchased is a free log retrieval operation binding the contract event 0xc3fed0fd7988126357ed3e92abd50623b7e1792521cb604e436a22baee3903a3.
//
// Solidity: event EnergyPurchased(uint256 id, address buyer)
func (_Main *MainFilterer) FilterEnergyPurchased(opts *bind.FilterOpts) (*MainEnergyPurchasedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "EnergyPurchased")
	if err != nil {
		return nil, err
	}
	return &MainEnergyPurchasedIterator{contract: _Main.contract, event: "EnergyPurchased", logs: logs, sub: sub}, nil
}

// WatchEnergyPurchased is a free log subscription operation binding the contract event 0xc3fed0fd7988126357ed3e92abd50623b7e1792521cb604e436a22baee3903a3.
//
// Solidity: event EnergyPurchased(uint256 id, address buyer)
func (_Main *MainFilterer) WatchEnergyPurchased(opts *bind.WatchOpts, sink chan<- *MainEnergyPurchased) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "EnergyPurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainEnergyPurchased)
				if err := _Main.contract.UnpackLog(event, "EnergyPurchased", log); err != nil {
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

// ParseEnergyPurchased is a log parse operation binding the contract event 0xc3fed0fd7988126357ed3e92abd50623b7e1792521cb604e436a22baee3903a3.
//
// Solidity: event EnergyPurchased(uint256 id, address buyer)
func (_Main *MainFilterer) ParseEnergyPurchased(log types.Log) (*MainEnergyPurchased, error) {
	event := new(MainEnergyPurchased)
	if err := _Main.contract.UnpackLog(event, "EnergyPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
