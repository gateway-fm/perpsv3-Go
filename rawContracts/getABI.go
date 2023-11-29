package rawContracts

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gateway-fm/perpsv3-Go/errors"
)

// gerABI is used to get contract ABI instance from given contract metadata
func getABI(meta *bind.MetaData) (*abi.ABI, error) {
	abiInstance, err := meta.GetAbi()
	if err != nil {
		logErr("getABI", fmt.Sprintln("get abi from meta err:", err.Error()))
		return nil, errors.GetInitContractErr(err)
	}

	return abiInstance, nil
}
