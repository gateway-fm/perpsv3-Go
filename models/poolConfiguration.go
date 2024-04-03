package models

import (
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"math/big"
)

type PoolConfiguration struct {
	MarketConfigurations []*MarketConfiguration
}

type MarketConfiguration struct {
	MarketID             *big.Int
	WeightD18            *big.Int
	MaxDebtShareValueD18 *big.Int
}

func GetPoolConfigurationFromContractData(configs []core.MarketConfigurationData) *PoolConfiguration {
	p := &PoolConfiguration{}

	for _, m := range configs {
		p.MarketConfigurations = append(p.MarketConfigurations, GetMarketConfigurationFromContractData(m))
	}

	return p
}

func GetMarketConfigurationFromContractData(config core.MarketConfigurationData) *MarketConfiguration {
	m := &MarketConfiguration{}

	m.MarketID = config.MarketId
	m.WeightD18 = config.WeightD18
	m.MaxDebtShareValueD18 = config.MaxDebtShareValueD18

	return m
}
