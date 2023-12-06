package services

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
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

	return s.getPosition(opts, accountID, marketID, block)
}

// DEPRECATED
//func (s *Service) getPositionMultiCall(opts *bind.CallOpts, accountID *big.Int, marketID *big.Int, block *types.Header) (res *models.Position, err error) {
//	getOpenPositionCallData, err := s.rawPerpsContract.GetCallDataOpenPosition(marketID, accountID)
//	if err != nil {
//		return res, err
//	}
//
//	callPostion := forwarder.TrustedMulticallForwarderCall3Value{
//		Target:       s.rawPerpsContract.Address(),
//		AllowFailure: false,
//		Value:        big.NewInt(0),
//		CallData:     getOpenPositionCallData,
//	}
//
//	feedID := models.GetPriceFeedIDFromMarketID(marketID)
//	if feedID == models.UNKNOWN {
//		logger.Log().WithField("layer", "getPositionMultiCall").Errorf(
//			"market ud: %v not supported on andromeda net", marketID.String(),
//		)
//		return res, errors.GetReadContractErr(fmt.Errorf("market %v not supported", marketID.String()), "rawForwarder", "Aggregate3Value")
//	}
//
//	fulfillOracleQueryCallData, err := s.rawERC7412.GetCallFulfillOracleQuery(feedID.String())
//	if err != nil {
//		return res, err
//	}
//
//	callFulfill := forwarder.TrustedMulticallForwarderCall3Value{
//		Target:       s.rawERC7412.Address(),
//		AllowFailure: false,
//		Value:        big.NewInt(1),
//		CallData:     fulfillOracleQueryCallData,
//	}
//
//	call, err := s.rawForwarder.Aggregate3Value([]forwarder.TrustedMulticallForwarderCall3Value{callFulfill, callPostion})
//	if err != nil {
//		return res, err
//	}
//
//	if len(call) != 2 {
//		logger.Log().WithField("layer", "getPositionMultiCall").Errorf("received %v from rawForwarder contract, expected 2", len(call))
//		return res, errors.GetReadContractErr(fmt.Errorf("invalid call"), "rawForwarder", "Aggregate3Value")
//	}
//
//	if !call[0].Success {
//		logger.Log().WithField("layer", "getPositionMultiCall").Error("call to perps unsuccessful")
//		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to erc7412"), "rawForwarder", "Aggregate3Value")
//	}
//
//	if !call[1].Success {
//		logger.Log().WithField("layer", "getPositionMultiCall").Error("call to erc7412 unsuccessful")
//		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to perps market"), "rawForwarder", "Aggregate3Value")
//	}
//
//	positionContract, err := s.rawPerpsContract.UnpackOpenPosition(call[1].ReturnData)
//	if err != nil {
//		return res, err
//	}
//
//	return models.GetPositionFromContract(positionContract, block.Number.Uint64(), block.Time), nil
//}

func (s *Service) getPosition(opts *bind.CallOpts, accountID *big.Int, marketID *big.Int, block *types.Header) (*models.Position, error) {
	positionContract, err := s.perpsMarket.GetOpenPosition(opts, accountID, marketID)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetPositions").Errorf(
			"contract getOpenPosition with accountID: %v, marketID: %v error: %v", accountID, marketID, err.Error(),
		)
		return nil, errors.GetReadContractErr(err, "PerpsMarket", "getOpenPosition")
	}

	return models.GetPositionFromContract(positionContract, block.Number.Uint64(), block.Time), nil
}
