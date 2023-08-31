// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sUSDTGoerli

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

// SUSDTGoerliMetaData contains all meta data concerning the SUSDTGoerli contract.
var SUSDTGoerliMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existing\",\"type\":\"uint256\"}],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existing\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parameter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnWithAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SUSDTGoerliABI is the input ABI used to generate the binding from.
// Deprecated: Use SUSDTGoerliMetaData.ABI instead.
var SUSDTGoerliABI = SUSDTGoerliMetaData.ABI

// SUSDTGoerli is an auto generated Go binding around an Ethereum contract.
type SUSDTGoerli struct {
	SUSDTGoerliCaller     // Read-only binding to the contract
	SUSDTGoerliTransactor // Write-only binding to the contract
	SUSDTGoerliFilterer   // Log filterer for contract events
}

// SUSDTGoerliCaller is an auto generated read-only Go binding around an Ethereum contract.
type SUSDTGoerliCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SUSDTGoerliTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SUSDTGoerliTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SUSDTGoerliFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SUSDTGoerliFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SUSDTGoerliSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SUSDTGoerliSession struct {
	Contract     *SUSDTGoerli      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SUSDTGoerliCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SUSDTGoerliCallerSession struct {
	Contract *SUSDTGoerliCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SUSDTGoerliTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SUSDTGoerliTransactorSession struct {
	Contract     *SUSDTGoerliTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SUSDTGoerliRaw is an auto generated low-level Go binding around an Ethereum contract.
type SUSDTGoerliRaw struct {
	Contract *SUSDTGoerli // Generic contract binding to access the raw methods on
}

// SUSDTGoerliCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SUSDTGoerliCallerRaw struct {
	Contract *SUSDTGoerliCaller // Generic read-only contract binding to access the raw methods on
}

// SUSDTGoerliTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SUSDTGoerliTransactorRaw struct {
	Contract *SUSDTGoerliTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSUSDTGoerli creates a new instance of SUSDTGoerli, bound to a specific deployed contract.
func NewSUSDTGoerli(address common.Address, backend bind.ContractBackend) (*SUSDTGoerli, error) {
	contract, err := bindSUSDTGoerli(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SUSDTGoerli{SUSDTGoerliCaller: SUSDTGoerliCaller{contract: contract}, SUSDTGoerliTransactor: SUSDTGoerliTransactor{contract: contract}, SUSDTGoerliFilterer: SUSDTGoerliFilterer{contract: contract}}, nil
}

// NewSUSDTGoerliCaller creates a new read-only instance of SUSDTGoerli, bound to a specific deployed contract.
func NewSUSDTGoerliCaller(address common.Address, caller bind.ContractCaller) (*SUSDTGoerliCaller, error) {
	contract, err := bindSUSDTGoerli(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SUSDTGoerliCaller{contract: contract}, nil
}

// NewSUSDTGoerliTransactor creates a new write-only instance of SUSDTGoerli, bound to a specific deployed contract.
func NewSUSDTGoerliTransactor(address common.Address, transactor bind.ContractTransactor) (*SUSDTGoerliTransactor, error) {
	contract, err := bindSUSDTGoerli(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SUSDTGoerliTransactor{contract: contract}, nil
}

// NewSUSDTGoerliFilterer creates a new log filterer instance of SUSDTGoerli, bound to a specific deployed contract.
func NewSUSDTGoerliFilterer(address common.Address, filterer bind.ContractFilterer) (*SUSDTGoerliFilterer, error) {
	contract, err := bindSUSDTGoerli(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SUSDTGoerliFilterer{contract: contract}, nil
}

// bindSUSDTGoerli binds a generic wrapper to an already deployed contract.
func bindSUSDTGoerli(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SUSDTGoerliMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SUSDTGoerli *SUSDTGoerliRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SUSDTGoerli.Contract.SUSDTGoerliCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SUSDTGoerli *SUSDTGoerliRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.SUSDTGoerliTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SUSDTGoerli *SUSDTGoerliRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.SUSDTGoerliTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SUSDTGoerli *SUSDTGoerliCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SUSDTGoerli.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SUSDTGoerli *SUSDTGoerliTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SUSDTGoerli *SUSDTGoerliTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SUSDTGoerli *SUSDTGoerliCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SUSDTGoerli *SUSDTGoerliSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SUSDTGoerli.Contract.Allowance(&_SUSDTGoerli.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SUSDTGoerli.Contract.Allowance(&_SUSDTGoerli.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SUSDTGoerli *SUSDTGoerliCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SUSDTGoerli *SUSDTGoerliSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SUSDTGoerli.Contract.BalanceOf(&_SUSDTGoerli.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SUSDTGoerli.Contract.BalanceOf(&_SUSDTGoerli.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SUSDTGoerli *SUSDTGoerliCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SUSDTGoerli *SUSDTGoerliSession) Decimals() (uint8, error) {
	return _SUSDTGoerli.Contract.Decimals(&_SUSDTGoerli.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) Decimals() (uint8, error) {
	return _SUSDTGoerli.Contract.Decimals(&_SUSDTGoerli.CallOpts)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_SUSDTGoerli *SUSDTGoerliCaller) GetAssociatedSystem(opts *bind.CallOpts, id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "getAssociatedSystem", id)

	outstruct := new(struct {
		Addr common.Address
		Kind [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Kind = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_SUSDTGoerli *SUSDTGoerliSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _SUSDTGoerli.Contract.GetAssociatedSystem(&_SUSDTGoerli.CallOpts, id)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _SUSDTGoerli.Contract.GetAssociatedSystem(&_SUSDTGoerli.CallOpts, id)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SUSDTGoerli *SUSDTGoerliCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SUSDTGoerli *SUSDTGoerliSession) GetImplementation() (common.Address, error) {
	return _SUSDTGoerli.Contract.GetImplementation(&_SUSDTGoerli.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) GetImplementation() (common.Address, error) {
	return _SUSDTGoerli.Contract.GetImplementation(&_SUSDTGoerli.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SUSDTGoerli *SUSDTGoerliCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SUSDTGoerli *SUSDTGoerliSession) IsInitialized() (bool, error) {
	return _SUSDTGoerli.Contract.IsInitialized(&_SUSDTGoerli.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) IsInitialized() (bool, error) {
	return _SUSDTGoerli.Contract.IsInitialized(&_SUSDTGoerli.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SUSDTGoerli *SUSDTGoerliCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SUSDTGoerli *SUSDTGoerliSession) Name() (string, error) {
	return _SUSDTGoerli.Contract.Name(&_SUSDTGoerli.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) Name() (string, error) {
	return _SUSDTGoerli.Contract.Name(&_SUSDTGoerli.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SUSDTGoerli *SUSDTGoerliCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SUSDTGoerli *SUSDTGoerliSession) NominatedOwner() (common.Address, error) {
	return _SUSDTGoerli.Contract.NominatedOwner(&_SUSDTGoerli.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) NominatedOwner() (common.Address, error) {
	return _SUSDTGoerli.Contract.NominatedOwner(&_SUSDTGoerli.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SUSDTGoerli *SUSDTGoerliCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SUSDTGoerli *SUSDTGoerliSession) Owner() (common.Address, error) {
	return _SUSDTGoerli.Contract.Owner(&_SUSDTGoerli.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) Owner() (common.Address, error) {
	return _SUSDTGoerli.Contract.Owner(&_SUSDTGoerli.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SUSDTGoerli *SUSDTGoerliCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SUSDTGoerli *SUSDTGoerliSession) Symbol() (string, error) {
	return _SUSDTGoerli.Contract.Symbol(&_SUSDTGoerli.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) Symbol() (string, error) {
	return _SUSDTGoerli.Contract.Symbol(&_SUSDTGoerli.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SUSDTGoerli *SUSDTGoerliCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SUSDTGoerli.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SUSDTGoerli *SUSDTGoerliSession) TotalSupply() (*big.Int, error) {
	return _SUSDTGoerli.Contract.TotalSupply(&_SUSDTGoerli.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SUSDTGoerli *SUSDTGoerliCallerSession) TotalSupply() (*big.Int, error) {
	return _SUSDTGoerli.Contract.TotalSupply(&_SUSDTGoerli.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SUSDTGoerli *SUSDTGoerliSession) AcceptOwnership() (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.AcceptOwnership(&_SUSDTGoerli.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.AcceptOwnership(&_SUSDTGoerli.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Approve(&_SUSDTGoerli.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Approve(&_SUSDTGoerli.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Burn(&_SUSDTGoerli.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Burn(&_SUSDTGoerli.TransactOpts, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address target, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) Burn0(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "burn0", target, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address target, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) Burn0(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Burn0(&_SUSDTGoerli.TransactOpts, target, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address target, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) Burn0(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Burn0(&_SUSDTGoerli.TransactOpts, target, amount)
}

// BurnWithAllowance is a paid mutator transaction binding the contract method 0xaaa15fd1.
//
// Solidity: function burnWithAllowance(address from, address spender, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) BurnWithAllowance(opts *bind.TransactOpts, from common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "burnWithAllowance", from, spender, amount)
}

// BurnWithAllowance is a paid mutator transaction binding the contract method 0xaaa15fd1.
//
// Solidity: function burnWithAllowance(address from, address spender, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) BurnWithAllowance(from common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.BurnWithAllowance(&_SUSDTGoerli.TransactOpts, from, spender, amount)
}

// BurnWithAllowance is a paid mutator transaction binding the contract method 0xaaa15fd1.
//
// Solidity: function burnWithAllowance(address from, address spender, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) BurnWithAllowance(from common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.BurnWithAllowance(&_SUSDTGoerli.TransactOpts, from, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.DecreaseAllowance(&_SUSDTGoerli.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.DecreaseAllowance(&_SUSDTGoerli.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.IncreaseAllowance(&_SUSDTGoerli.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.IncreaseAllowance(&_SUSDTGoerli.TransactOpts, spender, addedValue)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) InitOrUpgradeNft(opts *bind.TransactOpts, id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "initOrUpgradeNft", id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.InitOrUpgradeNft(&_SUSDTGoerli.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.InitOrUpgradeNft(&_SUSDTGoerli.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) InitOrUpgradeToken(opts *bind.TransactOpts, id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "initOrUpgradeToken", id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.InitOrUpgradeToken(&_SUSDTGoerli.TransactOpts, id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.InitOrUpgradeToken(&_SUSDTGoerli.TransactOpts, id, name, symbol, decimals, impl)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) Initialize(opts *bind.TransactOpts, tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "initialize", tokenName, tokenSymbol, tokenDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) Initialize(tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Initialize(&_SUSDTGoerli.TransactOpts, tokenName, tokenSymbol, tokenDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) Initialize(tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Initialize(&_SUSDTGoerli.TransactOpts, tokenName, tokenSymbol, tokenDecimals)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address target, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) Mint(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "mint", target, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address target, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) Mint(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Mint(&_SUSDTGoerli.TransactOpts, target, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address target, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) Mint(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Mint(&_SUSDTGoerli.TransactOpts, target, amount)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) NominateNewOwner(opts *bind.TransactOpts, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "nominateNewOwner", newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.NominateNewOwner(&_SUSDTGoerli.TransactOpts, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.NominateNewOwner(&_SUSDTGoerli.TransactOpts, newNominatedOwner)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) RegisterUnmanagedSystem(opts *bind.TransactOpts, id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "registerUnmanagedSystem", id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.RegisterUnmanagedSystem(&_SUSDTGoerli.TransactOpts, id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.RegisterUnmanagedSystem(&_SUSDTGoerli.TransactOpts, id, endpoint)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) RenounceNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "renounceNomination")
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_SUSDTGoerli *SUSDTGoerliSession) RenounceNomination() (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.RenounceNomination(&_SUSDTGoerli.TransactOpts)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) RenounceNomination() (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.RenounceNomination(&_SUSDTGoerli.TransactOpts)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address from, address spender, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) SetAllowance(opts *bind.TransactOpts, from common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "setAllowance", from, spender, amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address from, address spender, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) SetAllowance(from common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.SetAllowance(&_SUSDTGoerli.TransactOpts, from, spender, amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address from, address spender, uint256 amount) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) SetAllowance(from common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.SetAllowance(&_SUSDTGoerli.TransactOpts, from, spender, amount)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) SimulateUpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "simulateUpgradeTo", newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.SimulateUpgradeTo(&_SUSDTGoerli.TransactOpts, newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.SimulateUpgradeTo(&_SUSDTGoerli.TransactOpts, newImplementation)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Transfer(&_SUSDTGoerli.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.Transfer(&_SUSDTGoerli.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.TransferFrom(&_SUSDTGoerli.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.TransferFrom(&_SUSDTGoerli.TransactOpts, from, to, amount)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SUSDTGoerli *SUSDTGoerliSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.UpgradeTo(&_SUSDTGoerli.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SUSDTGoerli *SUSDTGoerliTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SUSDTGoerli.Contract.UpgradeTo(&_SUSDTGoerli.TransactOpts, newImplementation)
}

// SUSDTGoerliApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SUSDTGoerli contract.
type SUSDTGoerliApprovalIterator struct {
	Event *SUSDTGoerliApproval // Event containing the contract specifics and raw log

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
func (it *SUSDTGoerliApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SUSDTGoerliApproval)
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
		it.Event = new(SUSDTGoerliApproval)
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
func (it *SUSDTGoerliApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SUSDTGoerliApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SUSDTGoerliApproval represents a Approval event raised by the SUSDTGoerli contract.
type SUSDTGoerliApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_SUSDTGoerli *SUSDTGoerliFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SUSDTGoerliApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SUSDTGoerli.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SUSDTGoerliApprovalIterator{contract: _SUSDTGoerli.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_SUSDTGoerli *SUSDTGoerliFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SUSDTGoerliApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SUSDTGoerli.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SUSDTGoerliApproval)
				if err := _SUSDTGoerli.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_SUSDTGoerli *SUSDTGoerliFilterer) ParseApproval(log types.Log) (*SUSDTGoerliApproval, error) {
	event := new(SUSDTGoerliApproval)
	if err := _SUSDTGoerli.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SUSDTGoerliAssociatedSystemSetIterator is returned from FilterAssociatedSystemSet and is used to iterate over the raw logs and unpacked data for AssociatedSystemSet events raised by the SUSDTGoerli contract.
type SUSDTGoerliAssociatedSystemSetIterator struct {
	Event *SUSDTGoerliAssociatedSystemSet // Event containing the contract specifics and raw log

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
func (it *SUSDTGoerliAssociatedSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SUSDTGoerliAssociatedSystemSet)
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
		it.Event = new(SUSDTGoerliAssociatedSystemSet)
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
func (it *SUSDTGoerliAssociatedSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SUSDTGoerliAssociatedSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SUSDTGoerliAssociatedSystemSet represents a AssociatedSystemSet event raised by the SUSDTGoerli contract.
type SUSDTGoerliAssociatedSystemSet struct {
	Kind  [32]byte
	Id    [32]byte
	Proxy common.Address
	Impl  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssociatedSystemSet is a free log retrieval operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_SUSDTGoerli *SUSDTGoerliFilterer) FilterAssociatedSystemSet(opts *bind.FilterOpts, kind [][32]byte, id [][32]byte) (*SUSDTGoerliAssociatedSystemSetIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SUSDTGoerli.contract.FilterLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return &SUSDTGoerliAssociatedSystemSetIterator{contract: _SUSDTGoerli.contract, event: "AssociatedSystemSet", logs: logs, sub: sub}, nil
}

// WatchAssociatedSystemSet is a free log subscription operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_SUSDTGoerli *SUSDTGoerliFilterer) WatchAssociatedSystemSet(opts *bind.WatchOpts, sink chan<- *SUSDTGoerliAssociatedSystemSet, kind [][32]byte, id [][32]byte) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SUSDTGoerli.contract.WatchLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SUSDTGoerliAssociatedSystemSet)
				if err := _SUSDTGoerli.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
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

// ParseAssociatedSystemSet is a log parse operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_SUSDTGoerli *SUSDTGoerliFilterer) ParseAssociatedSystemSet(log types.Log) (*SUSDTGoerliAssociatedSystemSet, error) {
	event := new(SUSDTGoerliAssociatedSystemSet)
	if err := _SUSDTGoerli.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SUSDTGoerliOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the SUSDTGoerli contract.
type SUSDTGoerliOwnerChangedIterator struct {
	Event *SUSDTGoerliOwnerChanged // Event containing the contract specifics and raw log

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
func (it *SUSDTGoerliOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SUSDTGoerliOwnerChanged)
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
		it.Event = new(SUSDTGoerliOwnerChanged)
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
func (it *SUSDTGoerliOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SUSDTGoerliOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SUSDTGoerliOwnerChanged represents a OwnerChanged event raised by the SUSDTGoerli contract.
type SUSDTGoerliOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SUSDTGoerli *SUSDTGoerliFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SUSDTGoerliOwnerChangedIterator, error) {

	logs, sub, err := _SUSDTGoerli.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SUSDTGoerliOwnerChangedIterator{contract: _SUSDTGoerli.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SUSDTGoerli *SUSDTGoerliFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SUSDTGoerliOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _SUSDTGoerli.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SUSDTGoerliOwnerChanged)
				if err := _SUSDTGoerli.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SUSDTGoerli *SUSDTGoerliFilterer) ParseOwnerChanged(log types.Log) (*SUSDTGoerliOwnerChanged, error) {
	event := new(SUSDTGoerliOwnerChanged)
	if err := _SUSDTGoerli.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SUSDTGoerliOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the SUSDTGoerli contract.
type SUSDTGoerliOwnerNominatedIterator struct {
	Event *SUSDTGoerliOwnerNominated // Event containing the contract specifics and raw log

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
func (it *SUSDTGoerliOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SUSDTGoerliOwnerNominated)
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
		it.Event = new(SUSDTGoerliOwnerNominated)
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
func (it *SUSDTGoerliOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SUSDTGoerliOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SUSDTGoerliOwnerNominated represents a OwnerNominated event raised by the SUSDTGoerli contract.
type SUSDTGoerliOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SUSDTGoerli *SUSDTGoerliFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*SUSDTGoerliOwnerNominatedIterator, error) {

	logs, sub, err := _SUSDTGoerli.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &SUSDTGoerliOwnerNominatedIterator{contract: _SUSDTGoerli.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SUSDTGoerli *SUSDTGoerliFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *SUSDTGoerliOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _SUSDTGoerli.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SUSDTGoerliOwnerNominated)
				if err := _SUSDTGoerli.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SUSDTGoerli *SUSDTGoerliFilterer) ParseOwnerNominated(log types.Log) (*SUSDTGoerliOwnerNominated, error) {
	event := new(SUSDTGoerliOwnerNominated)
	if err := _SUSDTGoerli.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SUSDTGoerliTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SUSDTGoerli contract.
type SUSDTGoerliTransferIterator struct {
	Event *SUSDTGoerliTransfer // Event containing the contract specifics and raw log

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
func (it *SUSDTGoerliTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SUSDTGoerliTransfer)
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
		it.Event = new(SUSDTGoerliTransfer)
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
func (it *SUSDTGoerliTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SUSDTGoerliTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SUSDTGoerliTransfer represents a Transfer event raised by the SUSDTGoerli contract.
type SUSDTGoerliTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_SUSDTGoerli *SUSDTGoerliFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SUSDTGoerliTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SUSDTGoerli.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SUSDTGoerliTransferIterator{contract: _SUSDTGoerli.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_SUSDTGoerli *SUSDTGoerliFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SUSDTGoerliTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SUSDTGoerli.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SUSDTGoerliTransfer)
				if err := _SUSDTGoerli.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_SUSDTGoerli *SUSDTGoerliFilterer) ParseTransfer(log types.Log) (*SUSDTGoerliTransfer, error) {
	event := new(SUSDTGoerliTransfer)
	if err := _SUSDTGoerli.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SUSDTGoerliUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SUSDTGoerli contract.
type SUSDTGoerliUpgradedIterator struct {
	Event *SUSDTGoerliUpgraded // Event containing the contract specifics and raw log

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
func (it *SUSDTGoerliUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SUSDTGoerliUpgraded)
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
		it.Event = new(SUSDTGoerliUpgraded)
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
func (it *SUSDTGoerliUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SUSDTGoerliUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SUSDTGoerliUpgraded represents a Upgraded event raised by the SUSDTGoerli contract.
type SUSDTGoerliUpgraded struct {
	Self           common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_SUSDTGoerli *SUSDTGoerliFilterer) FilterUpgraded(opts *bind.FilterOpts, self []common.Address) (*SUSDTGoerliUpgradedIterator, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _SUSDTGoerli.contract.FilterLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return &SUSDTGoerliUpgradedIterator{contract: _SUSDTGoerli.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_SUSDTGoerli *SUSDTGoerliFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SUSDTGoerliUpgraded, self []common.Address) (event.Subscription, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _SUSDTGoerli.contract.WatchLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SUSDTGoerliUpgraded)
				if err := _SUSDTGoerli.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_SUSDTGoerli *SUSDTGoerliFilterer) ParseUpgraded(log types.Log) (*SUSDTGoerliUpgraded, error) {
	event := new(SUSDTGoerliUpgraded)
	if err := _SUSDTGoerli.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
