package models

import "math/big"

// PriceFeedID is a permission enum
type PriceFeedID int

const (
	UNKNOWN PriceFeedID = iota
	ETH
	BTC
	LINK
	OP
	SNX
)

// priceFeedIDs is mapping PriceFeedID to its string value
var priceFeedIDs = [...]string{
	UNKNOWN: "UNKNOWN",
	ETH:     "0xca80ba6dc32e08d06f1aa886011eed1d77c77be9eb761cc10d72b7d0a2fd57a6",
	BTC:     "0xf9c0172ba10dfa4d19088d94f5bf61d3b54d5bd7483a322a982e1373ee8ea31b",
	LINK:    "0x83be4ed61dd8a3518d198098ce37240c494710a7b9d85e35d9fceac21df08994",
	OP:      "0x71334dcd37620ce3c33e3bafef04cc80dec083042e49b734315b36d1aad7991f",
	SNX:     "0xe956a4199936e913b402474cb29576066f15108121d434606a19b34036e6d5cc",
}

// String is used to return PriceFeedID string value
func (p PriceFeedID) String() string {
	return priceFeedIDs[p]
}

// GetPriceFeedIDFromMarketID is used to get PriceFeedID from given market id value
func GetPriceFeedIDFromMarketID(id *big.Int) PriceFeedID {
	switch id.Int64() {
	case int64(100):
		return ETH
	case int64(200):
		return BTC
	case int64(300):
		return LINK
	case int64(400):
		return OP
	case int64(500):
		return SNX
	default:
		return UNKNOWN
	}
}
