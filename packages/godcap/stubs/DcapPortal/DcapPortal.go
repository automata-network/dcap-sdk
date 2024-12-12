// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DcapPortal

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

// IDcapPortalCallback is an auto generated low-level Go binding around an user-defined struct.
type IDcapPortalCallback struct {
	Value  *big.Int
	To     common.Address
	Params []byte
}

// DcapPortalMetaData contains all meta data concerning the DcapPortal contract.
var DcapPortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_Attestationaddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"estimateFeeBaseVerifyAndAttestWithZKProof\",\"inputs\":[{\"name\":\"output\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zkCoprocessor\",\"type\":\"uint8\",\"internalType\":\"enumIDcapAttestation.ZkCoProcessorType\"},{\"name\":\"proofBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"estimateFeeBaseVerifyOnChain\",\"inputs\":[{\"name\":\"rawQuote\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyAndAttestWithZKProof\",\"inputs\":[{\"name\":\"output\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zkCoprocessor\",\"type\":\"uint8\",\"internalType\":\"enumIDcapAttestation.ZkCoProcessorType\"},{\"name\":\"proofBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callback\",\"type\":\"tuple\",\"internalType\":\"structIDcapPortal.Callback\",\"components\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"verifiedOutput\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callbackOutput\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyOnChain\",\"inputs\":[{\"name\":\"rawQuote\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callback\",\"type\":\"tuple\",\"internalType\":\"structIDcapPortal.Callback\",\"components\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"verifiedOutput\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callbackOutput\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"error\",\"name\":\"CALLBACK_FAILED\",\"inputs\":[{\"name\":\"callbackOutput\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"INSUFFICIENT_FEE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"REJECT_RECURSIVE_CALL\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VERIFICATION_FAILED\",\"inputs\":[{\"name\":\"output\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]",
}

// DcapPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use DcapPortalMetaData.ABI instead.
var DcapPortalABI = DcapPortalMetaData.ABI

// DcapPortal is an auto generated Go binding around an Ethereum contract.
type DcapPortal struct {
	DcapPortalCaller     // Read-only binding to the contract
	DcapPortalTransactor // Write-only binding to the contract
	DcapPortalFilterer   // Log filterer for contract events
}

// DcapPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type DcapPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DcapPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DcapPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DcapPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DcapPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DcapPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DcapPortalSession struct {
	Contract     *DcapPortal       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DcapPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DcapPortalCallerSession struct {
	Contract *DcapPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DcapPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DcapPortalTransactorSession struct {
	Contract     *DcapPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DcapPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type DcapPortalRaw struct {
	Contract *DcapPortal // Generic contract binding to access the raw methods on
}

// DcapPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DcapPortalCallerRaw struct {
	Contract *DcapPortalCaller // Generic read-only contract binding to access the raw methods on
}

// DcapPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DcapPortalTransactorRaw struct {
	Contract *DcapPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDcapPortal creates a new instance of DcapPortal, bound to a specific deployed contract.
func NewDcapPortal(address common.Address, backend bind.ContractBackend) (*DcapPortal, error) {
	contract, err := bindDcapPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DcapPortal{DcapPortalCaller: DcapPortalCaller{contract: contract}, DcapPortalTransactor: DcapPortalTransactor{contract: contract}, DcapPortalFilterer: DcapPortalFilterer{contract: contract}}, nil
}

// NewDcapPortalCaller creates a new read-only instance of DcapPortal, bound to a specific deployed contract.
func NewDcapPortalCaller(address common.Address, caller bind.ContractCaller) (*DcapPortalCaller, error) {
	contract, err := bindDcapPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DcapPortalCaller{contract: contract}, nil
}

// NewDcapPortalTransactor creates a new write-only instance of DcapPortal, bound to a specific deployed contract.
func NewDcapPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*DcapPortalTransactor, error) {
	contract, err := bindDcapPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DcapPortalTransactor{contract: contract}, nil
}

// NewDcapPortalFilterer creates a new log filterer instance of DcapPortal, bound to a specific deployed contract.
func NewDcapPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*DcapPortalFilterer, error) {
	contract, err := bindDcapPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DcapPortalFilterer{contract: contract}, nil
}

// bindDcapPortal binds a generic wrapper to an already deployed contract.
func bindDcapPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DcapPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DcapPortal *DcapPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DcapPortal.Contract.DcapPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DcapPortal *DcapPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DcapPortal.Contract.DcapPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DcapPortal *DcapPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DcapPortal.Contract.DcapPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DcapPortal *DcapPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DcapPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DcapPortal *DcapPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DcapPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DcapPortal *DcapPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DcapPortal.Contract.contract.Transact(opts, method, params...)
}

// EstimateFeeBaseVerifyAndAttestWithZKProof is a paid mutator transaction binding the contract method 0x80541893.
//
// Solidity: function estimateFeeBaseVerifyAndAttestWithZKProof(bytes output, uint8 zkCoprocessor, bytes proofBytes) payable returns(uint256)
func (_DcapPortal *DcapPortalTransactor) EstimateFeeBaseVerifyAndAttestWithZKProof(opts *bind.TransactOpts, output []byte, zkCoprocessor uint8, proofBytes []byte) (*types.Transaction, error) {
	return _DcapPortal.contract.Transact(opts, "estimateFeeBaseVerifyAndAttestWithZKProof", output, zkCoprocessor, proofBytes)
}

// EstimateFeeBaseVerifyAndAttestWithZKProof is a paid mutator transaction binding the contract method 0x80541893.
//
// Solidity: function estimateFeeBaseVerifyAndAttestWithZKProof(bytes output, uint8 zkCoprocessor, bytes proofBytes) payable returns(uint256)
func (_DcapPortal *DcapPortalSession) EstimateFeeBaseVerifyAndAttestWithZKProof(output []byte, zkCoprocessor uint8, proofBytes []byte) (*types.Transaction, error) {
	return _DcapPortal.Contract.EstimateFeeBaseVerifyAndAttestWithZKProof(&_DcapPortal.TransactOpts, output, zkCoprocessor, proofBytes)
}

// EstimateFeeBaseVerifyAndAttestWithZKProof is a paid mutator transaction binding the contract method 0x80541893.
//
// Solidity: function estimateFeeBaseVerifyAndAttestWithZKProof(bytes output, uint8 zkCoprocessor, bytes proofBytes) payable returns(uint256)
func (_DcapPortal *DcapPortalTransactorSession) EstimateFeeBaseVerifyAndAttestWithZKProof(output []byte, zkCoprocessor uint8, proofBytes []byte) (*types.Transaction, error) {
	return _DcapPortal.Contract.EstimateFeeBaseVerifyAndAttestWithZKProof(&_DcapPortal.TransactOpts, output, zkCoprocessor, proofBytes)
}

// EstimateFeeBaseVerifyOnChain is a paid mutator transaction binding the contract method 0xe713bfa4.
//
// Solidity: function estimateFeeBaseVerifyOnChain(bytes rawQuote) payable returns(uint256)
func (_DcapPortal *DcapPortalTransactor) EstimateFeeBaseVerifyOnChain(opts *bind.TransactOpts, rawQuote []byte) (*types.Transaction, error) {
	return _DcapPortal.contract.Transact(opts, "estimateFeeBaseVerifyOnChain", rawQuote)
}

// EstimateFeeBaseVerifyOnChain is a paid mutator transaction binding the contract method 0xe713bfa4.
//
// Solidity: function estimateFeeBaseVerifyOnChain(bytes rawQuote) payable returns(uint256)
func (_DcapPortal *DcapPortalSession) EstimateFeeBaseVerifyOnChain(rawQuote []byte) (*types.Transaction, error) {
	return _DcapPortal.Contract.EstimateFeeBaseVerifyOnChain(&_DcapPortal.TransactOpts, rawQuote)
}

// EstimateFeeBaseVerifyOnChain is a paid mutator transaction binding the contract method 0xe713bfa4.
//
// Solidity: function estimateFeeBaseVerifyOnChain(bytes rawQuote) payable returns(uint256)
func (_DcapPortal *DcapPortalTransactorSession) EstimateFeeBaseVerifyOnChain(rawQuote []byte) (*types.Transaction, error) {
	return _DcapPortal.Contract.EstimateFeeBaseVerifyOnChain(&_DcapPortal.TransactOpts, rawQuote)
}

// VerifyAndAttestWithZKProof is a paid mutator transaction binding the contract method 0x20b43de7.
//
// Solidity: function verifyAndAttestWithZKProof(bytes output, uint8 zkCoprocessor, bytes proofBytes, (uint256,address,bytes) callback) payable returns(bytes verifiedOutput, bytes callbackOutput)
func (_DcapPortal *DcapPortalTransactor) VerifyAndAttestWithZKProof(opts *bind.TransactOpts, output []byte, zkCoprocessor uint8, proofBytes []byte, callback IDcapPortalCallback) (*types.Transaction, error) {
	return _DcapPortal.contract.Transact(opts, "verifyAndAttestWithZKProof", output, zkCoprocessor, proofBytes, callback)
}

// VerifyAndAttestWithZKProof is a paid mutator transaction binding the contract method 0x20b43de7.
//
// Solidity: function verifyAndAttestWithZKProof(bytes output, uint8 zkCoprocessor, bytes proofBytes, (uint256,address,bytes) callback) payable returns(bytes verifiedOutput, bytes callbackOutput)
func (_DcapPortal *DcapPortalSession) VerifyAndAttestWithZKProof(output []byte, zkCoprocessor uint8, proofBytes []byte, callback IDcapPortalCallback) (*types.Transaction, error) {
	return _DcapPortal.Contract.VerifyAndAttestWithZKProof(&_DcapPortal.TransactOpts, output, zkCoprocessor, proofBytes, callback)
}

// VerifyAndAttestWithZKProof is a paid mutator transaction binding the contract method 0x20b43de7.
//
// Solidity: function verifyAndAttestWithZKProof(bytes output, uint8 zkCoprocessor, bytes proofBytes, (uint256,address,bytes) callback) payable returns(bytes verifiedOutput, bytes callbackOutput)
func (_DcapPortal *DcapPortalTransactorSession) VerifyAndAttestWithZKProof(output []byte, zkCoprocessor uint8, proofBytes []byte, callback IDcapPortalCallback) (*types.Transaction, error) {
	return _DcapPortal.Contract.VerifyAndAttestWithZKProof(&_DcapPortal.TransactOpts, output, zkCoprocessor, proofBytes, callback)
}

// VerifyOnChain is a paid mutator transaction binding the contract method 0xff962b20.
//
// Solidity: function verifyOnChain(bytes rawQuote, (uint256,address,bytes) callback) payable returns(bytes verifiedOutput, bytes callbackOutput)
func (_DcapPortal *DcapPortalTransactor) VerifyOnChain(opts *bind.TransactOpts, rawQuote []byte, callback IDcapPortalCallback) (*types.Transaction, error) {
	return _DcapPortal.contract.Transact(opts, "verifyOnChain", rawQuote, callback)
}

// VerifyOnChain is a paid mutator transaction binding the contract method 0xff962b20.
//
// Solidity: function verifyOnChain(bytes rawQuote, (uint256,address,bytes) callback) payable returns(bytes verifiedOutput, bytes callbackOutput)
func (_DcapPortal *DcapPortalSession) VerifyOnChain(rawQuote []byte, callback IDcapPortalCallback) (*types.Transaction, error) {
	return _DcapPortal.Contract.VerifyOnChain(&_DcapPortal.TransactOpts, rawQuote, callback)
}

// VerifyOnChain is a paid mutator transaction binding the contract method 0xff962b20.
//
// Solidity: function verifyOnChain(bytes rawQuote, (uint256,address,bytes) callback) payable returns(bytes verifiedOutput, bytes callbackOutput)
func (_DcapPortal *DcapPortalTransactorSession) VerifyOnChain(rawQuote []byte, callback IDcapPortalCallback) (*types.Transaction, error) {
	return _DcapPortal.Contract.VerifyOnChain(&_DcapPortal.TransactOpts, rawQuote, callback)
}
