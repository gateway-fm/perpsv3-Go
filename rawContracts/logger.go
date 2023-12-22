package rawContracts

import (
	"fmt"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func logWarn(method string, msg string) {
	logger.Log().WithField("layer", fmt.Sprintf("RawContracts-%v", method)).Warn(msg)
}

func logErr(method string, msg string) {
	logger.Log().WithField("layer", fmt.Sprintf("RawContracts-%v", method)).Error(msg)
}
