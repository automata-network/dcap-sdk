// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VerifiedCounter

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

// VerifiedCounterMetaData contains all meta data concerning the VerifiedCounter contract.
var VerifiedCounterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_dcapPortalAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"number\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"report\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setNumber\",\"inputs\":[{\"name\":\"newNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Report\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CALLER_NOT_DCAP_PORTAL\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INVALID_ATTESTATION_OUTPUT\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MAGIC_NUMBER_MISMATCH\",\"inputs\":[]}]",
}

// VerifiedCounterABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifiedCounterMetaData.ABI instead.
var VerifiedCounterABI = VerifiedCounterMetaData.ABI

// VerifiedCounter is an auto generated Go binding around an Ethereum contract.
type VerifiedCounter struct {
	VerifiedCounterCaller     // Read-only binding to the contract
	VerifiedCounterTransactor // Write-only binding to the contract
	VerifiedCounterFilterer   // Log filterer for contract events
}

// VerifiedCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifiedCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiedCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifiedCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiedCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifiedCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiedCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifiedCounterSession struct {
	Contract     *VerifiedCounter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifiedCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifiedCounterCallerSession struct {
	Contract *VerifiedCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VerifiedCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifiedCounterTransactorSession struct {
	Contract     *VerifiedCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VerifiedCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifiedCounterRaw struct {
	Contract *VerifiedCounter // Generic contract binding to access the raw methods on
}

// VerifiedCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifiedCounterCallerRaw struct {
	Contract *VerifiedCounterCaller // Generic read-only contract binding to access the raw methods on
}

// VerifiedCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifiedCounterTransactorRaw struct {
	Contract *VerifiedCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifiedCounter creates a new instance of VerifiedCounter, bound to a specific deployed contract.
func NewVerifiedCounter(address common.Address, backend bind.ContractBackend) (*VerifiedCounter, error) {
	contract, err := bindVerifiedCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifiedCounter{VerifiedCounterCaller: VerifiedCounterCaller{contract: contract}, VerifiedCounterTransactor: VerifiedCounterTransactor{contract: contract}, VerifiedCounterFilterer: VerifiedCounterFilterer{contract: contract}}, nil
}

// NewVerifiedCounterCaller creates a new read-only instance of VerifiedCounter, bound to a specific deployed contract.
func NewVerifiedCounterCaller(address common.Address, caller bind.ContractCaller) (*VerifiedCounterCaller, error) {
	contract, err := bindVerifiedCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifiedCounterCaller{contract: contract}, nil
}

// NewVerifiedCounterTransactor creates a new write-only instance of VerifiedCounter, bound to a specific deployed contract.
func NewVerifiedCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifiedCounterTransactor, error) {
	contract, err := bindVerifiedCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifiedCounterTransactor{contract: contract}, nil
}

// NewVerifiedCounterFilterer creates a new log filterer instance of VerifiedCounter, bound to a specific deployed contract.
func NewVerifiedCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifiedCounterFilterer, error) {
	contract, err := bindVerifiedCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifiedCounterFilterer{contract: contract}, nil
}

// bindVerifiedCounter binds a generic wrapper to an already deployed contract.
func bindVerifiedCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifiedCounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifiedCounter *VerifiedCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifiedCounter.Contract.VerifiedCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifiedCounter *VerifiedCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiedCounter.Contract.VerifiedCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifiedCounter *VerifiedCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifiedCounter.Contract.VerifiedCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifiedCounter *VerifiedCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifiedCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifiedCounter *VerifiedCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiedCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifiedCounter *VerifiedCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifiedCounter.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_VerifiedCounter *VerifiedCounterCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VerifiedCounter.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_VerifiedCounter *VerifiedCounterSession) Number() (*big.Int, error) {
	return _VerifiedCounter.Contract.Number(&_VerifiedCounter.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_VerifiedCounter *VerifiedCounterCallerSession) Number() (*big.Int, error) {
	return _VerifiedCounter.Contract.Number(&_VerifiedCounter.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifiedCounter *VerifiedCounterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiedCounter.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifiedCounter *VerifiedCounterSession) Deposit() (*types.Transaction, error) {
	return _VerifiedCounter.Contract.Deposit(&_VerifiedCounter.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifiedCounter *VerifiedCounterTransactorSession) Deposit() (*types.Transaction, error) {
	return _VerifiedCounter.Contract.Deposit(&_VerifiedCounter.TransactOpts)
}

// Report is a paid mutator transaction binding the contract method 0x2606a10b.
//
// Solidity: function report() payable returns()
func (_VerifiedCounter *VerifiedCounterTransactor) Report(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiedCounter.contract.Transact(opts, "report")
}

// Report is a paid mutator transaction binding the contract method 0x2606a10b.
//
// Solidity: function report() payable returns()
func (_VerifiedCounter *VerifiedCounterSession) Report() (*types.Transaction, error) {
	return _VerifiedCounter.Contract.Report(&_VerifiedCounter.TransactOpts)
}

// Report is a paid mutator transaction binding the contract method 0x2606a10b.
//
// Solidity: function report() payable returns()
func (_VerifiedCounter *VerifiedCounterTransactorSession) Report() (*types.Transaction, error) {
	return _VerifiedCounter.Contract.Report(&_VerifiedCounter.TransactOpts)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_VerifiedCounter *VerifiedCounterTransactor) SetNumber(opts *bind.TransactOpts, newNumber *big.Int) (*types.Transaction, error) {
	return _VerifiedCounter.contract.Transact(opts, "setNumber", newNumber)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_VerifiedCounter *VerifiedCounterSession) SetNumber(newNumber *big.Int) (*types.Transaction, error) {
	return _VerifiedCounter.Contract.SetNumber(&_VerifiedCounter.TransactOpts, newNumber)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_VerifiedCounter *VerifiedCounterTransactorSession) SetNumber(newNumber *big.Int) (*types.Transaction, error) {
	return _VerifiedCounter.Contract.SetNumber(&_VerifiedCounter.TransactOpts, newNumber)
}

// VerifiedCounterReportIterator is returned from FilterReport and is used to iterate over the raw logs and unpacked data for Report events raised by the VerifiedCounter contract.
type VerifiedCounterReportIterator struct {
	Event *VerifiedCounterReport // Event containing the contract specifics and raw log

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
func (it *VerifiedCounterReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiedCounterReport)
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
		it.Event = new(VerifiedCounterReport)
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
func (it *VerifiedCounterReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiedCounterReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiedCounterReport represents a Report event raised by the VerifiedCounter contract.
type VerifiedCounterReport struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReport is a free log retrieval operation binding the contract event 0x15a6eee4587ae4ff553cd5b3f6f455469ccf6c51070136bd9b7c5c3ec606c9b5.
//
// Solidity: event Report(bytes arg0)
func (_VerifiedCounter *VerifiedCounterFilterer) FilterReport(opts *bind.FilterOpts) (*VerifiedCounterReportIterator, error) {

	logs, sub, err := _VerifiedCounter.contract.FilterLogs(opts, "Report")
	if err != nil {
		return nil, err
	}
	return &VerifiedCounterReportIterator{contract: _VerifiedCounter.contract, event: "Report", logs: logs, sub: sub}, nil
}

// WatchReport is a free log subscription operation binding the contract event 0x15a6eee4587ae4ff553cd5b3f6f455469ccf6c51070136bd9b7c5c3ec606c9b5.
//
// Solidity: event Report(bytes arg0)
func (_VerifiedCounter *VerifiedCounterFilterer) WatchReport(opts *bind.WatchOpts, sink chan<- *VerifiedCounterReport) (event.Subscription, error) {

	logs, sub, err := _VerifiedCounter.contract.WatchLogs(opts, "Report")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiedCounterReport)
				if err := _VerifiedCounter.contract.UnpackLog(event, "Report", log); err != nil {
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

// ParseReport is a log parse operation binding the contract event 0x15a6eee4587ae4ff553cd5b3f6f455469ccf6c51070136bd9b7c5c3ec606c9b5.
//
// Solidity: event Report(bytes arg0)
func (_VerifiedCounter *VerifiedCounterFilterer) ParseReport(log types.Log) (*VerifiedCounterReport, error) {
	event := new(VerifiedCounterReport)
	if err := _VerifiedCounter.contract.UnpackLog(event, "Report", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
