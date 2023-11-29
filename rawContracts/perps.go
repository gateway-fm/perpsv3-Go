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

// IRawPerpsContract is an interface for the raw implementation of the perps data contracts
type IRawPerpsContract interface {
	// GetCallDataOpenPosition is used to get calldata for the getOpenPosition method
	GetCallDataOpenPosition(marketID *big.Int, accountID *big.Int) ([]byte, error)
	// GetCallDataMarketSummary is used to get calldata for the getMarketSummary method
	GetCallDataMarketSummary(marketID *big.Int) ([]byte, error)
	// UnpackGetMarketSummary is used to unpack outputs for the getMarketSummary method
	UnpackGetMarketSummary(value []byte) (res *perpsMarket.IPerpsMarketModuleMarketSummary, err error)
	// UnpackOpenPosition is used to unpack outputs for the getOpenPosition method
	UnpackOpenPosition(value []byte) (res struct {
		TotalPnl       *big.Int
		AccruedFunding *big.Int
		PositionSize   *big.Int
	}, err error)
	// Address is used to get perps contract address
	Address() common.Address
}

// Perps is a raw perps contract implementation
type Perps struct {
	abi      *abi.ABI
	address  common.Address
	provider *ethclient.Client
	contract *bind.BoundContract
}

// NewPerps is used to get new Perps instance
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

func (p *Perps) GetCallDataOpenPosition(marketID *big.Int, accountID *big.Int) ([]byte, error) {
	callOpenPosition, err := p.abi.Pack("getOpenPosition", accountID, marketID)
	if err != nil {
		logErr("GetCallDataOpenPosition", fmt.Sprintln("abi pack getOpenPosition err:", err.Error()))
		return nil, errors.GetReadContractErr(err, "PerpsRaw", "getOpenPosition")
	}

	return callOpenPosition, nil
}

func (p *Perps) UnpackGetMarketSummary(value []byte) (res *perpsMarket.IPerpsMarketModuleMarketSummary, err error) {
	unpackedSummary, err := p.abi.Unpack("getMarketSummary", value)
	if err != nil {
		logErr("UnpackGetMarketSummary", fmt.Sprintln("abi unpack getMarketSummart err:", err.Error()))
		return res, errors.GetReadContractErr(err, "PerpsRaw", "UnpackGetMarketSummary")
	}

	return abi.ConvertType(unpackedSummary[0], new(perpsMarket.IPerpsMarketModuleMarketSummary)).(*perpsMarket.IPerpsMarketModuleMarketSummary), nil
}

func (p *Perps) UnpackOpenPosition(value []byte) (res struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, err error) {
	unpackedOpenPosition, err := p.abi.Unpack("getOpenPosition", value)
	if err != nil {
		logErr("UnpackOpenPosition", fmt.Sprintln("abi unpack getOpenPosition err:", err.Error()))
		return res, errors.GetReadContractErr(err, "PerpsRaw", "UnpackOpenPosition")
	}

	outstruct := new(struct {
		TotalPnl       *big.Int
		AccruedFunding *big.Int
		PositionSize   *big.Int
	})

	outstruct.TotalPnl = *abi.ConvertType(unpackedOpenPosition[0], new(*big.Int)).(**big.Int)
	outstruct.AccruedFunding = *abi.ConvertType(unpackedOpenPosition[1], new(*big.Int)).(**big.Int)
	outstruct.PositionSize = *abi.ConvertType(unpackedOpenPosition[2], new(*big.Int)).(**big.Int)

	return *outstruct, err
}
