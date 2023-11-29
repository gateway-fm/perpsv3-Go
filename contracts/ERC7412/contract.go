// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc7412

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

// Erc7412MetaData contains all meta data concerning the Erc7412 contract.
var Erc7412MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeRequired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"oracleQuery\",\"type\":\"bytes\"}],\"name\":\"OracleDataRequired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signedOffchainData\",\"type\":\"bytes\"}],\"name\":\"fulfillOracleQuery\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"oracleId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Erc7412ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc7412MetaData.ABI instead.
var Erc7412ABI = Erc7412MetaData.ABI

// Erc7412 is an auto generated Go binding around an Ethereum contract.
type Erc7412 struct {
	Erc7412Caller     // Read-only binding to the contract
	Erc7412Transactor // Write-only binding to the contract
	Erc7412Filterer   // Log filterer for contract events
}

// Erc7412Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc7412Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc7412Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc7412Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc7412Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc7412Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc7412Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc7412Session struct {
	Contract     *Erc7412          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc7412CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc7412CallerSession struct {
	Contract *Erc7412Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Erc7412TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc7412TransactorSession struct {
	Contract     *Erc7412Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Erc7412Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc7412Raw struct {
	Contract *Erc7412 // Generic contract binding to access the raw methods on
}

// Erc7412CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc7412CallerRaw struct {
	Contract *Erc7412Caller // Generic read-only contract binding to access the raw methods on
}

// Erc7412TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc7412TransactorRaw struct {
	Contract *Erc7412Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc7412 creates a new instance of Erc7412, bound to a specific deployed contract.
func NewErc7412(address common.Address, backend bind.ContractBackend) (*Erc7412, error) {
	contract, err := bindErc7412(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc7412{Erc7412Caller: Erc7412Caller{contract: contract}, Erc7412Transactor: Erc7412Transactor{contract: contract}, Erc7412Filterer: Erc7412Filterer{contract: contract}}, nil
}

// NewErc7412Caller creates a new read-only instance of Erc7412, bound to a specific deployed contract.
func NewErc7412Caller(address common.Address, caller bind.ContractCaller) (*Erc7412Caller, error) {
	contract, err := bindErc7412(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc7412Caller{contract: contract}, nil
}

// NewErc7412Transactor creates a new write-only instance of Erc7412, bound to a specific deployed contract.
func NewErc7412Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc7412Transactor, error) {
	contract, err := bindErc7412(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc7412Transactor{contract: contract}, nil
}

// NewErc7412Filterer creates a new log filterer instance of Erc7412, bound to a specific deployed contract.
func NewErc7412Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc7412Filterer, error) {
	contract, err := bindErc7412(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc7412Filterer{contract: contract}, nil
}

// bindErc7412 binds a generic wrapper to an already deployed contract.
func bindErc7412(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc7412MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc7412 *Erc7412Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc7412.Contract.Erc7412Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc7412 *Erc7412Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc7412.Contract.Erc7412Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc7412 *Erc7412Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc7412.Contract.Erc7412Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc7412 *Erc7412CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc7412.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc7412 *Erc7412TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc7412.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc7412 *Erc7412TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc7412.Contract.contract.Transact(opts, method, params...)
}

// OracleId is a free data retrieval call binding the contract method 0x2121778d.
//
// Solidity: function oracleId() view returns(bytes32 oracleId)
func (_Erc7412 *Erc7412Caller) OracleId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc7412.contract.Call(opts, &out, "oracleId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OracleId is a free data retrieval call binding the contract method 0x2121778d.
//
// Solidity: function oracleId() view returns(bytes32 oracleId)
func (_Erc7412 *Erc7412Session) OracleId() ([32]byte, error) {
	return _Erc7412.Contract.OracleId(&_Erc7412.CallOpts)
}

// OracleId is a free data retrieval call binding the contract method 0x2121778d.
//
// Solidity: function oracleId() view returns(bytes32 oracleId)
func (_Erc7412 *Erc7412CallerSession) OracleId() ([32]byte, error) {
	return _Erc7412.Contract.OracleId(&_Erc7412.CallOpts)
}

// FulfillOracleQuery is a paid mutator transaction binding the contract method 0x14b95956.
//
// Solidity: function fulfillOracleQuery(bytes signedOffchainData) payable returns()
func (_Erc7412 *Erc7412Transactor) FulfillOracleQuery(opts *bind.TransactOpts, signedOffchainData []byte) (*types.Transaction, error) {
	return _Erc7412.contract.Transact(opts, "fulfillOracleQuery", signedOffchainData)
}

// FulfillOracleQuery is a paid mutator transaction binding the contract method 0x14b95956.
//
// Solidity: function fulfillOracleQuery(bytes signedOffchainData) payable returns()
func (_Erc7412 *Erc7412Session) FulfillOracleQuery(signedOffchainData []byte) (*types.Transaction, error) {
	return _Erc7412.Contract.FulfillOracleQuery(&_Erc7412.TransactOpts, signedOffchainData)
}

// FulfillOracleQuery is a paid mutator transaction binding the contract method 0x14b95956.
//
// Solidity: function fulfillOracleQuery(bytes signedOffchainData) payable returns()
func (_Erc7412 *Erc7412TransactorSession) FulfillOracleQuery(signedOffchainData []byte) (*types.Transaction, error) {
	return _Erc7412.Contract.FulfillOracleQuery(&_Erc7412.TransactOpts, signedOffchainData)
}
