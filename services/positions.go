package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) GetPosition(accountID *big.Int, marketID *big.Int) (*models.Position, error) {
	latest, err := s.rpcClient.BlockNumber(context.Background())
	if err != nil {
		logger.Log().WithField("layer", "Service-GetPositions").Errorf(
			"error get latest block: %v", err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "BlockNumber")
	}

	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(latest)))
	if err != nil {
		logger.Log().WithField("layer", "Service-GetPositions").Errorf(
			"get block by numer: %v error: %v", latest, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	opts := &bind.CallOpts{BlockNumber: big.NewInt(int64(latest))}

	positionContract, err := s.perpsMarket.GetOpenPosition(opts, accountID, marketID)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetPositions").Errorf(
			"contract getOpenPosition with accountID: %v, marketID: %v error: %v", accountID, marketID, err.Error(),
		)
		return nil, errors.GetReadContractErr(err, "PerpsMarket", "getOpenPosition")
	}

	return models.GetPositionFromContract(positionContract, block.Number.Uint64(), block.Time), nil
}
