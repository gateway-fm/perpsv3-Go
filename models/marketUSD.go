package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
)

type MarketUSDDeposited struct {
	MarketId                 *big.Int
	Target                   common.Address
	Amount                   *big.Int
	Market                   common.Address
	CreditCapacity           *big.Int
	NetIssuance              *big.Int
	DepositedCollateralValue *big.Int
	ReportedDebt             *big.Int
	BlockNumber              uint64
	BlockTimestamp           uint64
	TransactionHash          string
}

type MarketUSDWithdrawn struct {
	MarketId                 *big.Int
	Target                   common.Address
	Amount                   *big.Int
	Market                   common.Address
	CreditCapacity           *big.Int
	NetIssuance              *big.Int
	DepositedCollateralValue *big.Int
	ReportedDebt             *big.Int
	BlockNumber              uint64
	BlockTimestamp           uint64
	TransactionHash          string
}

func GetMarketUSDDepositedFromEvent(event *core.CoreMarketUsdDeposited, time uint64) *MarketUSDDeposited {
	m := &MarketUSDDeposited{}

	m.MarketId = event.MarketId
	m.Target = event.Target
	m.Amount = event.Amount
	m.Market = event.Market
	m.CreditCapacity = event.CreditCapacity
	m.NetIssuance = event.NetIssuance
	m.DepositedCollateralValue = event.DepositedCollateralValue
	m.ReportedDebt = event.ReportedDebt
	m.BlockNumber = event.Raw.BlockNumber
	m.BlockTimestamp = time
	m.TransactionHash = event.Raw.TxHash.Hex()

	return m
}

func GetMarketUSDWithdrawnFromEvent(event *core.CoreMarketUsdWithdrawn, time uint64) *MarketUSDWithdrawn {
	m := &MarketUSDWithdrawn{}

	m.MarketId = event.MarketId
	m.Target = event.Target
	m.Amount = event.Amount
	m.Market = event.Market
	m.CreditCapacity = event.CreditCapacity
	m.NetIssuance = event.NetIssuance
	m.DepositedCollateralValue = event.DepositedCollateralValue
	m.ReportedDebt = event.ReportedDebt
	m.BlockNumber = event.Raw.BlockNumber
	m.BlockTimestamp = time
	m.TransactionHash = event.Raw.TxHash.Hex()

	return m
}
