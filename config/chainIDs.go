package config

// ChainID is a chain id enum
type ChainID int

const (
	Unknown ChainID = iota
	OptimismGoerli
	BaseAndromeda
)

var chainIDStrings = [...]string{
	Unknown:        "Unknown",
	OptimismGoerli: "OptimismGoerli",
	BaseAndromeda:  "BaseAndromeda",
}

var chainIDNums = [...]int{
	Unknown:        0,
	OptimismGoerli: 420,
	BaseAndromeda:  84513,
}

func (i ChainID) String() string {
	return chainIDStrings[i]
}

func (i ChainID) Int() int {
	return chainIDNums[i]
}
