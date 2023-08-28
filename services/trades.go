package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) RetrieveTrades(fromBlock uint64, toBLock *uint64) ([]*models.Trade, error) {
	if fromBlock == 0 {
		fromBlock = s.spotMarketFirstBlock
	}

	opts := &bind.FilterOpts{
		Start:   fromBlock,
		End:     toBLock,
		Context: context.Background(),
	}

	iterator, err := s.perpsMarket.FilterOrderSettled(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveTrades").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "perps market")
	}

	var trades []*models.Trade

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveTrades").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "perps market")
		}
		blockNumber := iterator.Event.Raw.BlockNumber

		trade, err := s.getTrade(iterator.Event, blockNumber)
		if err != nil {
			return nil, err
		}

		trades = append(trades, trade)
	}

	return trades, nil
}

// getTrade is used to get models.Trade from given event and block number
func (s *Service) getTrade(event *perpsMarketGoerli.PerpsMarketGoerliOrderSettled, blockN uint64) (*models.Trade, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveTrades").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetTradeFromEvent(event, block.Time), nil
}
