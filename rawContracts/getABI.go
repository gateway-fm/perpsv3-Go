package rawContracts

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"io"
	"os"
)

func getABI(abiFilePath string) (*abi.ABI, error) {
	abiFile, err := os.Open(abiFilePath)
	if err != nil {
		logErr("getABI", fmt.Sprintln("open abi file err:", err.Error()))
		return nil, errors.GetInitContractErr(err)
	}
	defer func() {
		if err := abiFile.Close(); err != nil {
			logWarn("getABI", fmt.Sprintln("err close abi file:", err.Error()))
		}
	}()

	b, err := io.ReadAll(abiFile)
	if err != nil {
		logErr("getABI", fmt.Sprintln("read abi file err:", err.Error()))
		return nil, errors.GetInitContractErr(err)
	}

	meta := &bind.MetaData{ABI: string(b)}

	abiInstance, err := meta.GetAbi()
	if err != nil {
		logErr("getABI", fmt.Sprintln("get abi from meta err:", err.Error()))
		return nil, errors.GetInitContractErr(err)
	}

	return abiInstance, nil
}
