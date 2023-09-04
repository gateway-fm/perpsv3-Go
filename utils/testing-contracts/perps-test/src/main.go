package main

import (
	"os"

	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
	perps_test "github.com/gateway-fm/perpsv3-Go/utils/testing-contracts/perps-test"
)

func main() {
	rpcURL := os.Getenv("TEST_RPC")
	if rpcURL == "" {
		logger.Log().WithField("layer", "TestPerpsMarket-Main").Fatalf("no test rpc url events in env vars")
	}

	perps := perps_test.GetTestPerpsMarket(
		rpcURL,
		"0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b",
		"0xe487Ad4291019b33e2230F8E2FB1fb6490325260",
		420,
	)

	args := os.Args

	if len(args) > 1 {

		switch args[1] {
		case "CommitOrder":
			// Market ID and Size in wei
			perps.CommitOrder(args[2], args[3])
			perps.Close()
		case "ModifyCollateral":
			// Market ID and Amount in wei
			perps.ModifyCollateral(args[2], args[3])
			perps.Close()
		case "GetCollateralAmount":
			// Market ID
			perps.GetCollateralAmount(args[2])
			perps.Close()
		case "GetMaxCollateralAmount":
			// Market ID
			perps.GetMaxCollateralAmount(args[2])
			perps.Close()
		case "SetApproval":
			// Amount
			perps.SetApproval(args[2])
			perps.Close()
		case "GetApproval":
			perps.GetApproval()
			perps.Close()
		case "CreateAccount":
			perps.CreateAccount()
			perps.Close()
		case "GrantPermission":
			// Account ID, Permission and User
			perps.GrantPermission(args[2], args[3], args[4])
			perps.Close()
		default:
			perps.Close()
			logger.Log().WithField("layer", "TestPerpsMarket-Main").Fatalf("unknown command")
		}

	} else {
		perps.Close()
		logger.Log().WithField("layer", "TestPerpsMarket-Main").Fatalf("no args")
	}
}
