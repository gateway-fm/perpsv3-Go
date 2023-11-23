package main

import (
	"log"
	"math/big"

	perpsv3_Go "github.com/gateway-fm/perpsv3-Go"
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

	// TODO: clean out tests
	res, err := lib.GetMarketSummary(big.NewInt(200))
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(res)
	}
	res, err = lib.GetMarketSummary(big.NewInt(100))
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(res)
	}
	//res, err := lib.GetMarketSummary(big.NewInt(200))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println(res)

	//...
	// call needed methods
	// ...

	lib.Close()
}
