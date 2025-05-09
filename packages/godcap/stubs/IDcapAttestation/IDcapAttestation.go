// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDcapAttestation

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

// IDcapAttestationMetaData contains all meta data concerning the IDcapAttestation contract.
var IDcapAttestationMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getBp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyAndAttestOnChain\",\"inputs\":[{\"name\":\"rawQuote\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"output\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyAndAttestWithZKProof\",\"inputs\":[{\"name\":\"output\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zkCoprocessor\",\"type\":\"uint8\",\"internalType\":\"enumIDcapAttestation.ZkCoProcessorType\"},{\"name\":\"proofBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"verifiedOutput\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"}]",
}

// IDcapAttestationABI is the input ABI used to generate the binding from.
// Deprecated: Use IDcapAttestationMetaData.ABI instead.
var IDcapAttestationABI = IDcapAttestationMetaData.ABI

// IDcapAttestation is an auto generated Go binding around an Ethereum contract.
type IDcapAttestation struct {
	IDcapAttestationCaller     // Read-only binding to the contract
	IDcapAttestationTransactor // Write-only binding to the contract
	IDcapAttestationFilterer   // Log filterer for contract events
}

// IDcapAttestationCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDcapAttestationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDcapAttestationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDcapAttestationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDcapAttestationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDcapAttestationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDcapAttestationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDcapAttestationSession struct {
	Contract     *IDcapAttestation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDcapAttestationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDcapAttestationCallerSession struct {
	Contract *IDcapAttestationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IDcapAttestationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDcapAttestationTransactorSession struct {
	Contract     *IDcapAttestationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IDcapAttestationRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDcapAttestationRaw struct {
	Contract *IDcapAttestation // Generic contract binding to access the raw methods on
}

// IDcapAttestationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDcapAttestationCallerRaw struct {
	Contract *IDcapAttestationCaller // Generic read-only contract binding to access the raw methods on
}

// IDcapAttestationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDcapAttestationTransactorRaw struct {
	Contract *IDcapAttestationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDcapAttestation creates a new instance of IDcapAttestation, bound to a specific deployed contract.
func NewIDcapAttestation(address common.Address, backend bind.ContractBackend) (*IDcapAttestation, error) {
	contract, err := bindIDcapAttestation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDcapAttestation{IDcapAttestationCaller: IDcapAttestationCaller{contract: contract}, IDcapAttestationTransactor: IDcapAttestationTransactor{contract: contract}, IDcapAttestationFilterer: IDcapAttestationFilterer{contract: contract}}, nil
}

// NewIDcapAttestationCaller creates a new read-only instance of IDcapAttestation, bound to a specific deployed contract.
func NewIDcapAttestationCaller(address common.Address, caller bind.ContractCaller) (*IDcapAttestationCaller, error) {
	contract, err := bindIDcapAttestation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDcapAttestationCaller{contract: contract}, nil
}

// NewIDcapAttestationTransactor creates a new write-only instance of IDcapAttestation, bound to a specific deployed contract.
func NewIDcapAttestationTransactor(address common.Address, transactor bind.ContractTransactor) (*IDcapAttestationTransactor, error) {
	contract, err := bindIDcapAttestation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDcapAttestationTransactor{contract: contract}, nil
}

// NewIDcapAttestationFilterer creates a new log filterer instance of IDcapAttestation, bound to a specific deployed contract.
func NewIDcapAttestationFilterer(address common.Address, filterer bind.ContractFilterer) (*IDcapAttestationFilterer, error) {
	contract, err := bindIDcapAttestation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDcapAttestationFilterer{contract: contract}, nil
}

// bindIDcapAttestation binds a generic wrapper to an already deployed contract.
func bindIDcapAttestation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDcapAttestationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDcapAttestation *IDcapAttestationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDcapAttestation.Contract.IDcapAttestationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDcapAttestation *IDcapAttestationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDcapAttestation.Contract.IDcapAttestationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDcapAttestation *IDcapAttestationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDcapAttestation.Contract.IDcapAttestationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDcapAttestation *IDcapAttestationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDcapAttestation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDcapAttestation *IDcapAttestationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDcapAttestation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDcapAttestation *IDcapAttestationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDcapAttestation.Contract.contract.Transact(opts, method, params...)
}

// GetBp is a free data retrieval call binding the contract method 0x6655dddc.
//
// Solidity: function getBp() view returns(uint16)
func (_IDcapAttestation *IDcapAttestationCaller) GetBp(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _IDcapAttestation.contract.Call(opts, &out, "getBp")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetBp is a free data retrieval call binding the contract method 0x6655dddc.
//
// Solidity: function getBp() view returns(uint16)
func (_IDcapAttestation *IDcapAttestationSession) GetBp() (uint16, error) {
	return _IDcapAttestation.Contract.GetBp(&_IDcapAttestation.CallOpts)
}

// GetBp is a free data retrieval call binding the contract method 0x6655dddc.
//
// Solidity: function getBp() view returns(uint16)
func (_IDcapAttestation *IDcapAttestationCallerSession) GetBp() (uint16, error) {
	return _IDcapAttestation.Contract.GetBp(&_IDcapAttestation.CallOpts)
}

// VerifyAndAttestOnChain is a paid mutator transaction binding the contract method 0x38d8480a.
//
// Solidity: function verifyAndAttestOnChain(bytes rawQuote) payable returns(bool success, bytes output)
func (_IDcapAttestation *IDcapAttestationTransactor) VerifyAndAttestOnChain(opts *bind.TransactOpts, rawQuote []byte) (*types.Transaction, error) {
	return _IDcapAttestation.contract.Transact(opts, "verifyAndAttestOnChain", rawQuote)
}

// VerifyAndAttestOnChain is a paid mutator transaction binding the contract method 0x38d8480a.
//
// Solidity: function verifyAndAttestOnChain(bytes rawQuote) payable returns(bool success, bytes output)
func (_IDcapAttestation *IDcapAttestationSession) VerifyAndAttestOnChain(rawQuote []byte) (*types.Transaction, error) {
	return _IDcapAttestation.Contract.VerifyAndAttestOnChain(&_IDcapAttestation.TransactOpts, rawQuote)
}

// VerifyAndAttestOnChain is a paid mutator transaction binding the contract method 0x38d8480a.
//
// Solidity: function verifyAndAttestOnChain(bytes rawQuote) payable returns(bool success, bytes output)
func (_IDcapAttestation *IDcapAttestationTransactorSession) VerifyAndAttestOnChain(rawQuote []byte) (*types.Transaction, error) {
	return _IDcapAttestation.Contract.VerifyAndAttestOnChain(&_IDcapAttestation.TransactOpts, rawQuote)
}

// VerifyAndAttestWithZKProof is a paid mutator transaction binding the contract method 0x57859ce0.
//
// Solidity: function verifyAndAttestWithZKProof(bytes output, uint8 zkCoprocessor, bytes proofBytes) payable returns(bool success, bytes verifiedOutput)
func (_IDcapAttestation *IDcapAttestationTransactor) VerifyAndAttestWithZKProof(opts *bind.TransactOpts, output []byte, zkCoprocessor uint8, proofBytes []byte) (*types.Transaction, error) {
	return _IDcapAttestation.contract.Transact(opts, "verifyAndAttestWithZKProof", output, zkCoprocessor, proofBytes)
}

// VerifyAndAttestWithZKProof is a paid mutator transaction binding the contract method 0x57859ce0.
//
// Solidity: function verifyAndAttestWithZKProof(bytes output, uint8 zkCoprocessor, bytes proofBytes) payable returns(bool success, bytes verifiedOutput)
func (_IDcapAttestation *IDcapAttestationSession) VerifyAndAttestWithZKProof(output []byte, zkCoprocessor uint8, proofBytes []byte) (*types.Transaction, error) {
	return _IDcapAttestation.Contract.VerifyAndAttestWithZKProof(&_IDcapAttestation.TransactOpts, output, zkCoprocessor, proofBytes)
}

// VerifyAndAttestWithZKProof is a paid mutator transaction binding the contract method 0x57859ce0.
//
// Solidity: function verifyAndAttestWithZKProof(bytes output, uint8 zkCoprocessor, bytes proofBytes) payable returns(bool success, bytes verifiedOutput)
func (_IDcapAttestation *IDcapAttestationTransactorSession) VerifyAndAttestWithZKProof(output []byte, zkCoprocessor uint8, proofBytes []byte) (*types.Transaction, error) {
	return _IDcapAttestation.Contract.VerifyAndAttestWithZKProof(&_IDcapAttestation.TransactOpts, output, zkCoprocessor, proofBytes)
}
