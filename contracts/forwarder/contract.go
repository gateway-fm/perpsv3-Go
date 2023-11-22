// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package forwarder

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

// MinimalForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type MinimalForwarderForwardRequest struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Gas   *big.Int
	Nonce *big.Int
	Data  []byte
}

// TrustedMulticallForwarderCall is an auto generated low-level Go binding around an user-defined struct.
type TrustedMulticallForwarderCall struct {
	Target   common.Address
	CallData []byte
}

// TrustedMulticallForwarderCall3 is an auto generated low-level Go binding around an user-defined struct.
type TrustedMulticallForwarderCall3 struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

// TrustedMulticallForwarderCall3Value is an auto generated low-level Go binding around an user-defined struct.
type TrustedMulticallForwarderCall3Value struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
}

// TrustedMulticallForwarderResult is an auto generated low-level Go binding around an user-defined struct.
type TrustedMulticallForwarderResult struct {
	Success    bool
	ReturnData []byte
}

// ForwarderMetaData contains all meta data concerning the Forwarder contract.
var ForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Call3[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Call3Value[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3Value\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"basefee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structTrustedMulticallForwarder.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use ForwarderMetaData.ABI instead.
var ForwarderABI = ForwarderMetaData.ABI

// Forwarder is an auto generated Go binding around an Ethereum contract.
type Forwarder struct {
	ForwarderCaller     // Read-only binding to the contract
	ForwarderTransactor // Write-only binding to the contract
	ForwarderFilterer   // Log filterer for contract events
}

// ForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForwarderSession struct {
	Contract     *Forwarder        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForwarderCallerSession struct {
	Contract *ForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForwarderTransactorSession struct {
	Contract     *ForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForwarderRaw struct {
	Contract *Forwarder // Generic contract binding to access the raw methods on
}

// ForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForwarderCallerRaw struct {
	Contract *ForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// ForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForwarderTransactorRaw struct {
	Contract *ForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForwarder creates a new instance of Forwarder, bound to a specific deployed contract.
func NewForwarder(address common.Address, backend bind.ContractBackend) (*Forwarder, error) {
	contract, err := bindForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Forwarder{ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

// NewForwarderCaller creates a new read-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderCaller(address common.Address, caller bind.ContractCaller) (*ForwarderCaller, error) {
	contract, err := bindForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderCaller{contract: contract}, nil
}

// NewForwarderTransactor creates a new write-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ForwarderTransactor, error) {
	contract, err := bindForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderTransactor{contract: contract}, nil
}

// NewForwarderFilterer creates a new log filterer instance of Forwarder, bound to a specific deployed contract.
func NewForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ForwarderFilterer, error) {
	contract, err := bindForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForwarderFilterer{contract: contract}, nil
}

// bindForwarder binds a generic wrapper to an already deployed contract.
func bindForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.ForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Forwarder *ForwarderCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Forwarder *ForwarderSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Forwarder.Contract.Eip712Domain(&_Forwarder.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Forwarder *ForwarderCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Forwarder.Contract.Eip712Domain(&_Forwarder.CallOpts)
}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_Forwarder *ForwarderCaller) GetBasefee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getBasefee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_Forwarder *ForwarderSession) GetBasefee() (*big.Int, error) {
	return _Forwarder.Contract.GetBasefee(&_Forwarder.CallOpts)
}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_Forwarder *ForwarderCallerSession) GetBasefee() (*big.Int, error) {
	return _Forwarder.Contract.GetBasefee(&_Forwarder.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Forwarder *ForwarderCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Forwarder *ForwarderSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Forwarder.Contract.GetBlockHash(&_Forwarder.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Forwarder *ForwarderCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Forwarder.Contract.GetBlockHash(&_Forwarder.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Forwarder *ForwarderCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Forwarder *ForwarderSession) GetBlockNumber() (*big.Int, error) {
	return _Forwarder.Contract.GetBlockNumber(&_Forwarder.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Forwarder *ForwarderCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Forwarder.Contract.GetBlockNumber(&_Forwarder.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_Forwarder *ForwarderCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_Forwarder *ForwarderSession) GetChainId() (*big.Int, error) {
	return _Forwarder.Contract.GetChainId(&_Forwarder.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_Forwarder *ForwarderCallerSession) GetChainId() (*big.Int, error) {
	return _Forwarder.Contract.GetChainId(&_Forwarder.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Forwarder *ForwarderCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Forwarder *ForwarderSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Forwarder.Contract.GetCurrentBlockCoinbase(&_Forwarder.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Forwarder *ForwarderCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Forwarder.Contract.GetCurrentBlockCoinbase(&_Forwarder.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Forwarder *ForwarderCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Forwarder *ForwarderSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Forwarder.Contract.GetCurrentBlockDifficulty(&_Forwarder.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Forwarder *ForwarderCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Forwarder.Contract.GetCurrentBlockDifficulty(&_Forwarder.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Forwarder *ForwarderCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Forwarder *ForwarderSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Forwarder.Contract.GetCurrentBlockGasLimit(&_Forwarder.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Forwarder *ForwarderCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Forwarder.Contract.GetCurrentBlockGasLimit(&_Forwarder.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Forwarder *ForwarderCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Forwarder *ForwarderSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Forwarder.Contract.GetCurrentBlockTimestamp(&_Forwarder.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Forwarder *ForwarderCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Forwarder.Contract.GetCurrentBlockTimestamp(&_Forwarder.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Forwarder *ForwarderCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Forwarder *ForwarderSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetEthBalance(&_Forwarder.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Forwarder *ForwarderCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetEthBalance(&_Forwarder.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Forwarder *ForwarderCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Forwarder *ForwarderSession) GetLastBlockHash() ([32]byte, error) {
	return _Forwarder.Contract.GetLastBlockHash(&_Forwarder.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Forwarder *ForwarderCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _Forwarder.Contract.GetLastBlockHash(&_Forwarder.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Forwarder *ForwarderCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Forwarder *ForwarderSession) GetNonce(from common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Forwarder *ForwarderCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, from)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_Forwarder *ForwarderCaller) Verify(opts *bind.CallOpts, req MinimalForwarderForwardRequest, signature []byte) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "verify", req, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_Forwarder *ForwarderSession) Verify(req MinimalForwarderForwardRequest, signature []byte) (bool, error) {
	return _Forwarder.Contract.Verify(&_Forwarder.CallOpts, req, signature)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_Forwarder *ForwarderCallerSession) Verify(req MinimalForwarderForwardRequest, signature []byte) (bool, error) {
	return _Forwarder.Contract.Verify(&_Forwarder.CallOpts, req, signature)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_Forwarder *ForwarderTransactor) Aggregate(opts *bind.TransactOpts, calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_Forwarder *ForwarderSession) Aggregate(calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.Contract.Aggregate(&_Forwarder.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_Forwarder *ForwarderTransactorSession) Aggregate(calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.Contract.Aggregate(&_Forwarder.TransactOpts, calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Forwarder *ForwarderTransactor) Aggregate3(opts *bind.TransactOpts, calls []TrustedMulticallForwarderCall3) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "aggregate3", calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Forwarder *ForwarderSession) Aggregate3(calls []TrustedMulticallForwarderCall3) (*types.Transaction, error) {
	return _Forwarder.Contract.Aggregate3(&_Forwarder.TransactOpts, calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Forwarder *ForwarderTransactorSession) Aggregate3(calls []TrustedMulticallForwarderCall3) (*types.Transaction, error) {
	return _Forwarder.Contract.Aggregate3(&_Forwarder.TransactOpts, calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Forwarder *ForwarderTransactor) Aggregate3Value(opts *bind.TransactOpts, calls []TrustedMulticallForwarderCall3Value) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "aggregate3Value", calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Forwarder *ForwarderSession) Aggregate3Value(calls []TrustedMulticallForwarderCall3Value) (*types.Transaction, error) {
	return _Forwarder.Contract.Aggregate3Value(&_Forwarder.TransactOpts, calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Forwarder *ForwarderTransactorSession) Aggregate3Value(calls []TrustedMulticallForwarderCall3Value) (*types.Transaction, error) {
	return _Forwarder.Contract.Aggregate3Value(&_Forwarder.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Forwarder *ForwarderTransactor) BlockAndAggregate(opts *bind.TransactOpts, calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "blockAndAggregate", calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Forwarder *ForwarderSession) BlockAndAggregate(calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.Contract.BlockAndAggregate(&_Forwarder.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Forwarder *ForwarderTransactorSession) BlockAndAggregate(calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.Contract.BlockAndAggregate(&_Forwarder.TransactOpts, calls)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_Forwarder *ForwarderTransactor) Execute(opts *bind.TransactOpts, req MinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "execute", req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_Forwarder *ForwarderSession) Execute(req MinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_Forwarder *ForwarderTransactorSession) Execute(req MinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, signature)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Forwarder *ForwarderTransactor) TryAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "tryAggregate", requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Forwarder *ForwarderSession) TryAggregate(requireSuccess bool, calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.Contract.TryAggregate(&_Forwarder.TransactOpts, requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_Forwarder *ForwarderTransactorSession) TryAggregate(requireSuccess bool, calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.Contract.TryAggregate(&_Forwarder.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Forwarder *ForwarderTransactor) TryBlockAndAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Forwarder *ForwarderSession) TryBlockAndAggregate(requireSuccess bool, calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.Contract.TryBlockAndAggregate(&_Forwarder.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Forwarder *ForwarderTransactorSession) TryBlockAndAggregate(requireSuccess bool, calls []TrustedMulticallForwarderCall) (*types.Transaction, error) {
	return _Forwarder.Contract.TryBlockAndAggregate(&_Forwarder.TransactOpts, requireSuccess, calls)
}

// ForwarderEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Forwarder contract.
type ForwarderEIP712DomainChangedIterator struct {
	Event *ForwarderEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *ForwarderEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderEIP712DomainChanged)
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
		it.Event = new(ForwarderEIP712DomainChanged)
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
func (it *ForwarderEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderEIP712DomainChanged represents a EIP712DomainChanged event raised by the Forwarder contract.
type ForwarderEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Forwarder *ForwarderFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ForwarderEIP712DomainChangedIterator, error) {

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ForwarderEIP712DomainChangedIterator{contract: _Forwarder.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Forwarder *ForwarderFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ForwarderEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderEIP712DomainChanged)
				if err := _Forwarder.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Forwarder *ForwarderFilterer) ParseEIP712DomainChanged(log types.Log) (*ForwarderEIP712DomainChanged, error) {
	event := new(ForwarderEIP712DomainChanged)
	if err := _Forwarder.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
