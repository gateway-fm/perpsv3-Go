package services

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gateway-fm/perpsv3-Go/errors"
)

func (s *Service) RetrieveTrades(fromBlock uint64, toBLock *uint64) error {
	if fromBlock == 0 {
		fromBlock = s.spotMarketFirstBlock
	}

	opts := &bind.FilterOpts{
		Start:   fromBlock,
		End:     toBLock,
		Context: context.Background(),
	}

	iterator, err := s.spotMarket.FilterOrderSettled(opts, nil, nil, nil)
	if err != nil {
		return errors.GetFilterErr(err, "spot market")
	}

	for iterator.Next() {

		log.Println("iterator next")

		if iterator.Error() != nil {
			return errors.GetFilterErr(err, "spot market")
		}

		log.Println(iterator.Event)
	}

	//owner, _ := s.spotMarket.Owner(nil)
	//log.Println(owner)

	return nil
}
