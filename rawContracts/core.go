package rawContracts

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"math/big"
)

// IRawCoreContract is an interface for the raw implementation of the core data contracts
type IRawCoreContract interface {
	UnpackVaultDebt(value []byte) (*big.Int, error)
	GetCallDataVaultDebt(poolID *big.Int, collateralType common.Address) ([]byte, error)
	// Address is used to get perps contract address
	Address() common.Address
}

// Core is a raw core contract implementation
type Core struct {
	abi      *abi.ABI
	address  common.Address
	provider *ethclient.Client
	contract *bind.BoundContract
}

// NewCore is used to get new Core instance
func NewCore(address common.Address, provider *ethclient.Client) (IRawCoreContract, error) {
	c := &Core{}

	abiInstance, err := getABI(core.CoreMetaData)
	if err != nil {
		return nil, err
	}

	c.abi = abiInstance
	c.provider = provider
	c.address = address
	c.contract = bind.NewBoundContract(
		address,
		*abiInstance,
		provider,
		provider,
		provider,
	)

	return c, nil
}

func (p *Core) GetCallDataVaultDebt(poolID *big.Int, collateralType common.Address) ([]byte, error) {
	callDataVaultDebt, err := p.abi.Pack("getVaultDebt", poolID, collateralType)
	if err != nil {
		logErr("GetCallDataVaultDebt", fmt.Sprintln("abi pack getVaultDebt err:", err.Error()))
		return nil, errors.GetReadContractErr(err, "CoreRaw", "GetCallDataVaultDebt")
	}

	return callDataVaultDebt, nil
}

func (p *Core) UnpackVaultDebt(value []byte) (*big.Int, error) {
	unpackDebt, err := p.abi.Unpack("getVaultDebt", value)
	if err != nil {
		logErr("UnpackVaultDebt", fmt.Sprintln("abi unpack getVaultDebt err:", err.Error()))
		return nil, errors.GetReadContractErr(err, "CoreRaw", "UnpackVaultDebt")
	}

	return abi.ConvertType(unpackDebt[0], new(big.Int)).(*big.Int), nil
}

func (p *Core) Address() common.Address {
	return p.address
}
