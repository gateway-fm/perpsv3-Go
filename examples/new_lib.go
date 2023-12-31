package main

import (
	perpsv3_Go "github.com/gateway-fm/perpsv3-Go"
	"log"
	"math/big"
)

func main() {
	// you can use custom rpc provider url. To use default url use blank string
	rpcURL := ""

	//conf := perpsv3_Go.GetOptimismGoerliDefaultConfig(rpcURL)
	// OR you can use base andromeda chain
	conf := perpsv3_Go.GetBaseAndromedaDefaultConfig(rpcURL)

	lib, err := perpsv3_Go.Create(conf)
	if err != nil {
		log.Fatal(err)
	}

	res, err := lib.GetAvailableMargin(big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	//...
	// call needed methods
	// ...

	lib.Close()
}
