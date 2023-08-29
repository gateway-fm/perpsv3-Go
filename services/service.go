package services

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/models"
)

// IService is a service layer interface
type IService interface {
	// RetrieveTrades is used to get logs from the "OrderSettled" event preps market contract within given block range
	RetrieveTrades(fromBlock uint64, toBLock *uint64) ([]*models.Trade, error)

	// GetPosition is used to get "Position" data struct from the latest block from the perps market with given data
	GetPosition(accountID *big.Int, marketID *big.Int) (*models.Position, error)
}

// Service is an implementation of IService interface
type Service struct {
	rpcClient             *ethclient.Client
	core                  *coreGoerli.CoreGoerli
	coreFirstBlock        uint64
	spotMarket            *spotMarketGoerli.SpotMarketGoerli
	spotMarketFirstBlock  uint64
	perpsMarket           *perpsMarketGoerli.PerpsMarketGoerli
	perpsMarketFirstBlock uint64
}

// NewService is used to get instance of Service
func NewService(
	rpc *ethclient.Client,
	core *coreGoerli.CoreGoerli,
	coreFirstBlock uint64,
	spotMarket *spotMarketGoerli.SpotMarketGoerli,
	spotMarketFirstBlock uint64,
	perpsMarket *perpsMarketGoerli.PerpsMarketGoerli,
	perpsMarketFirstBlock uint64,
) IService {
	return &Service{
		rpcClient:             rpc,
		core:                  core,
		coreFirstBlock:        coreFirstBlock,
		spotMarket:            spotMarket,
		spotMarketFirstBlock:  spotMarketFirstBlock,
		perpsMarket:           perpsMarket,
		perpsMarketFirstBlock: perpsMarketFirstBlock,
	}
}
