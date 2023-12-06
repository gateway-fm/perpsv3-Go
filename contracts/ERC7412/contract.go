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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pythAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeRequired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"updateType\",\"type\":\"uint8\"}],\"name\":\"NotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"oracleQuery\",\"type\":\"bytes\"}],\"name\":\"OracleDataRequired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signedOffchainData\",\"type\":\"bytes\"}],\"name\":\"fulfillOracleQuery\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"priceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"requestedTime\",\"type\":\"uint64\"}],\"name\":\"getBenchmarkPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"priceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"stalenessTolerance\",\"type\":\"uint256\"}],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pythAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// GetBenchmarkPrice is a free data retrieval call binding the contract method 0x8f93ca56.
//
// Solidity: function getBenchmarkPrice(bytes32 priceId, uint64 requestedTime) view returns(int256)
func (_Erc7412 *Erc7412Caller) GetBenchmarkPrice(opts *bind.CallOpts, priceId [32]byte, requestedTime uint64) (*big.Int, error) {
	var out []interface{}
	err := _Erc7412.contract.Call(opts, &out, "getBenchmarkPrice", priceId, requestedTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBenchmarkPrice is a free data retrieval call binding the contract method 0x8f93ca56.
//
// Solidity: function getBenchmarkPrice(bytes32 priceId, uint64 requestedTime) view returns(int256)
func (_Erc7412 *Erc7412Session) GetBenchmarkPrice(priceId [32]byte, requestedTime uint64) (*big.Int, error) {
	return _Erc7412.Contract.GetBenchmarkPrice(&_Erc7412.CallOpts, priceId, requestedTime)
}

// GetBenchmarkPrice is a free data retrieval call binding the contract method 0x8f93ca56.
//
// Solidity: function getBenchmarkPrice(bytes32 priceId, uint64 requestedTime) view returns(int256)
func (_Erc7412 *Erc7412CallerSession) GetBenchmarkPrice(priceId [32]byte, requestedTime uint64) (*big.Int, error) {
	return _Erc7412.Contract.GetBenchmarkPrice(&_Erc7412.CallOpts, priceId, requestedTime)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x079bba7f.
//
// Solidity: function getLatestPrice(bytes32 priceId, uint256 stalenessTolerance) view returns(int256)
func (_Erc7412 *Erc7412Caller) GetLatestPrice(opts *bind.CallOpts, priceId [32]byte, stalenessTolerance *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc7412.contract.Call(opts, &out, "getLatestPrice", priceId, stalenessTolerance)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x079bba7f.
//
// Solidity: function getLatestPrice(bytes32 priceId, uint256 stalenessTolerance) view returns(int256)
func (_Erc7412 *Erc7412Session) GetLatestPrice(priceId [32]byte, stalenessTolerance *big.Int) (*big.Int, error) {
	return _Erc7412.Contract.GetLatestPrice(&_Erc7412.CallOpts, priceId, stalenessTolerance)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x079bba7f.
//
// Solidity: function getLatestPrice(bytes32 priceId, uint256 stalenessTolerance) view returns(int256)
func (_Erc7412 *Erc7412CallerSession) GetLatestPrice(priceId [32]byte, stalenessTolerance *big.Int) (*big.Int, error) {
	return _Erc7412.Contract.GetLatestPrice(&_Erc7412.CallOpts, priceId, stalenessTolerance)
}

// OracleId is a free data retrieval call binding the contract method 0x2121778d.
//
// Solidity: function oracleId() pure returns(bytes32)
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
// Solidity: function oracleId() pure returns(bytes32)
func (_Erc7412 *Erc7412Session) OracleId() ([32]byte, error) {
	return _Erc7412.Contract.OracleId(&_Erc7412.CallOpts)
}

// OracleId is a free data retrieval call binding the contract method 0x2121778d.
//
// Solidity: function oracleId() pure returns(bytes32)
func (_Erc7412 *Erc7412CallerSession) OracleId() ([32]byte, error) {
	return _Erc7412.Contract.OracleId(&_Erc7412.CallOpts)
}

// PythAddress is a free data retrieval call binding the contract method 0x65f92bac.
//
// Solidity: function pythAddress() view returns(address)
func (_Erc7412 *Erc7412Caller) PythAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc7412.contract.Call(opts, &out, "pythAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PythAddress is a free data retrieval call binding the contract method 0x65f92bac.
//
// Solidity: function pythAddress() view returns(address)
func (_Erc7412 *Erc7412Session) PythAddress() (common.Address, error) {
	return _Erc7412.Contract.PythAddress(&_Erc7412.CallOpts)
}

// PythAddress is a free data retrieval call binding the contract method 0x65f92bac.
//
// Solidity: function pythAddress() view returns(address)
func (_Erc7412 *Erc7412CallerSession) PythAddress() (common.Address, error) {
	return _Erc7412.Contract.PythAddress(&_Erc7412.CallOpts)
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Erc7412 *Erc7412Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Erc7412.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Erc7412 *Erc7412Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Erc7412.Contract.Fallback(&_Erc7412.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Erc7412 *Erc7412TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Erc7412.Contract.Fallback(&_Erc7412.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Erc7412 *Erc7412Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc7412.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Erc7412 *Erc7412Session) Receive() (*types.Transaction, error) {
	return _Erc7412.Contract.Receive(&_Erc7412.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Erc7412 *Erc7412TransactorSession) Receive() (*types.Transaction, error) {
	return _Erc7412.Contract.Receive(&_Erc7412.TransactOpts)
}
