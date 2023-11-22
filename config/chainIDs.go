package config

// ChainID is a chain id enum
type ChainID int

const (
	OptimismGoerli ChainID = iota
	BaseAndromeda
)

var chainIDStrings = [...]string{
	OptimismGoerli: "OptimismGoerli",
	BaseAndromeda:  "BaseAndromeda",
}

var chainIDNums = [...]int{
	OptimismGoerli: 420,
	BaseAndromeda:  84513,
}

func (i ChainID) String() string {
	return chainIDStrings[i]
}

func (i ChainID) Int() int {
	return chainIDNums[i]
}
