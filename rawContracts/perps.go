package rawContracts

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"math/big"

	"github.com/gateway-fm/perpsv3-Go/errors"
)

type IRawPerpsContract interface {
	GetCallDataMarketSummary(marketID *big.Int) ([]byte, error)
	UnpackGetMarketSummary(value []byte) (res *perpsMarket.IPerpsMarketModuleMarketSummary, err error)
	Address() common.Address
}

type Perps struct {
	abi      *abi.ABI
	address  common.Address
	provider *ethclient.Client
	contract *bind.BoundContract
}

func NewPerps(address common.Address, provider *ethclient.Client) (IRawPerpsContract, error) {
	c := &Perps{}

	abiInstance, err := getABI("./contracts/84531-andromeda-PerpsMarket.json")
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

func (p *Perps) Address() common.Address {
	return p.address
}

func (p *Perps) GetCallDataMarketSummary(marketID *big.Int) ([]byte, error) {
	callDataSummary, err := p.abi.Pack("getMarketSummary", marketID)
	if err != nil {
		logErr("GetCallDataMarketSummary", fmt.Sprintln("abi pack getMarketSummary err:", err.Error()))
		return nil, errors.GetReadContractErr(err, "PerpsRaw", "getMarketSummary")
	}

	return callDataSummary, nil
}

func (p *Perps) UnpackGetMarketSummary(value []byte) (res *perpsMarket.IPerpsMarketModuleMarketSummary, err error) {
	unpackedSummary, err := p.abi.Unpack("getMarketSummary", value)
	if err != nil {
		logErr("UnpackGetMarketSummary", fmt.Sprintln("abi unpack getMarketSummart err:", err.Error()))
		return res, errors.GetReadContractErr(err, "PerpsRaw", "UnpackGetMarketSummary")
	}

	return abi.ConvertType(unpackedSummary[0], new(perpsMarket.IPerpsMarketModuleMarketSummary)).(*perpsMarket.IPerpsMarketModuleMarketSummary), nil
}
