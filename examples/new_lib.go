package main

import (
	"github.com/ethereum/go-ethereum/common"
	perpsv3_Go "github.com/gateway-fm/perpsv3-Go"
	"log"
	"math/big"
)

func main() {
	// you can use custom rpc provider url. To use default url use blank string
	rpcURL := ""

	//conf := perpsv3_Go.GetOptimismGoerliDefaultConfig(rpcURL)
	// OR you can use base andromeda chain
	conf := perpsv3_Go.GetBaseMainnetDefaultConfig(rpcURL)

	lib, err := perpsv3_Go.Create(conf)
	if err != nil {
		log.Fatal(err)
	}

	res, err := lib.GetAccountCollateralCore(big.NewInt(1), common.HexToAddress("0xC74EA762CF06C9151CE074E6A569A5945B6302E7"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	//...
	// call needed methods
	// ...

	lib.Close()
}
