package services

import (
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
	"math/big"

	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
)

func (s *Service) GetMarketMetadata(marketID *big.Int) (*models.MarketMetadata, error) {
	if marketID == nil {
		logger.Log().WithField("layer", "Service-GetMarketMetadata").Errorf("received nil market id")
		return nil, errors.GetInvalidArgumentErr("market id cannot be nil")
	}

	res, err := s.perpsMarket.Metadata(nil, marketID)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetMarketMetadata").Errorf("error from the contract: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perpsMarket", "metadata")
	}

	if res.Name == "" || res.Symbol == "" {
		logger.Log().WithField("layer", "Service-GetMarketMetadata").Errorf("received blank name and symbol from the contract")
		return nil, errors.GetInvalidArgumentErr("market id does not exist")
	}

	return models.GetMarketMetadataFromContractResponse(marketID, res.Name, res.Symbol), nil
}
