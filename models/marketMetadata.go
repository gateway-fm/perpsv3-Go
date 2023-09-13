package models

import (
	"math/big"
)

// MarketMetadata is a market metadata model
//   - MarketID is a market ID value
//   - Name is a market name value
//   - Symbol is a market symbol value for example 'ETH'
type MarketMetadata struct {
	MarketID *big.Int
	Name     string
	Symbol   string
}

// GetMarketMetadataFromContractResponse is used to get MarketMetadata model from given values
func GetMarketMetadataFromContractResponse(id *big.Int, name string, symbol string) *MarketMetadata {
	return &MarketMetadata{
		MarketID: id,
		Name:     name,
		Symbol:   symbol,
	}
}
