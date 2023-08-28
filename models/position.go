package models

import "math/big"

// Position
//   - TotalPnl: Represents the total profit and loss for the position.
//   - AccruedFunding: Represents the accrued funding for the position.
//   - PositionSize: Represents the size of the position.
//   - BlockNumber: Represents the block number at which the position data was fetched.
//   - BlockTimestamp: Represents the timestamp of the block at which the position data was fetched.
type Position struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
	BlockNumber    uint64
	BlockTimestamp uint64
}

// positionContract is a data struct received from contract
type positionContract struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}

// GetPositionFromContract is used to get Position struct from given contract data struct and block values
func GetPositionFromContract(position positionContract, blockN uint64, blockT uint64) *Position {
	return &Position{
		TotalPnl:       position.TotalPnl,
		AccruedFunding: position.AccruedFunding,
		PositionSize:   position.PositionSize,
		BlockNumber:    blockN,
		BlockTimestamp: blockT,
	}
}
