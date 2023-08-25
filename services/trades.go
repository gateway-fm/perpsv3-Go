package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) RetrieveTrades(fromBlock uint64, toBLock *uint64) ([]models.Trade, error) {
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

	var trades []models.Trade

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveTrades").Errorf("iterator error: %v", err.Error())
			return nil, errors.GetFilterErr(err, "perps market")
		}
		trade := models.Trade{}

		blockNumber := iterator.Event.Raw.BlockNumber

		block, err := s.rpcClient.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			logger.Log().WithField("layer", "Service-RetrieveTrades").Errorf(
				"get block:%v by number error: %v", blockNumber, err.Error(),
			)
			return nil, errors.GetFilterErr(err, "perps market")
		}

		trade.ParseFromEvent(iterator.Event, block.Time())
		trades = append(trades, trade)
	}

	return trades, nil
}
