package services

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
)

// IService is a service layer interface
type IService interface {
	//
}

// Service is an implementation of IService interface
type Service struct {
	rpcClient  *ethclient.Client
	core       *coreGoerli.CoreGoerli
	spotMarket *spotMarketGoerli.SpotMarketGoerli
}

// NewService is used to get instance of Service
func NewService(rpc *ethclient.Client, core *coreGoerli.CoreGoerli, spotMarket *spotMarketGoerli.SpotMarketGoerli) IService {
	return &Service{
		rpcClient:  rpc,
		core:       core,
		spotMarket: spotMarket,
	}
}
