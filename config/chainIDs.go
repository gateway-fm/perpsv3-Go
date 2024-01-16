package config

// ChainID is a chain id enum
type ChainID int

const (
	Unknown ChainID = iota
	BaseMainnet
	OptimismGoerli
	BaseAndromeda
	BaseSepolia
)

var chainIDStrings = [...]string{
	Unknown:        "Unknown",
	BaseMainnet:    "BaseMainnet",
	OptimismGoerli: "OptimismGoerli",
	BaseAndromeda:  "BaseAndromeda",
	BaseSepolia:    "BaseSepolia",
}

var chainIDNums = [...]int{
	Unknown:        0,
	BaseMainnet:    8453,
	OptimismGoerli: 420,
	BaseAndromeda:  84531,
	BaseSepolia:    84532,
}

func (i ChainID) String() string {
	return chainIDStrings[i]
}

func (i ChainID) Int() int {
	return chainIDNums[i]
}
