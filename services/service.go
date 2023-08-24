package services

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
)

// IService is a service layer interface
type IService interface {
	RetrieveTrades(fromBlock uint64, toBLock *uint64) error
}

// Service is an implementation of IService interface
type Service struct {
	rpcClient            *ethclient.Client
	core                 *coreGoerli.CoreGoerli
	coreFirstBlock       uint64
	spotMarket           *spotMarketGoerli.SpotMarketGoerli
	spotMarketFirstBlock uint64
}

// NewService is used to get instance of Service
func NewService(
	rpc *ethclient.Client,
	core *coreGoerli.CoreGoerli,
	coreFirstBlock uint64,
	spotMarket *spotMarketGoerli.SpotMarketGoerli,
	spotMarketFirstBlock uint64,
) IService {
	return &Service{
		rpcClient:            rpc,
		core:                 core,
		coreFirstBlock:       coreFirstBlock,
		spotMarket:           spotMarket,
		spotMarketFirstBlock: spotMarketFirstBlock,
	}
}
