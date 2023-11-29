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

	// TODO: clean out tests

	// POSITION

	accs, err := lib.FormatAccountsLimit(20000)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("accs", len(accs))

	for _, a := range accs {
		_, err := lib.GetPosition(a.ID, big.NewInt(100))
		if err != nil {
			log.Println("id", a.ID.String(), "market", 100)
			log.Println(err.Error())
		}

		_, err = lib.GetPosition(a.ID, big.NewInt(200))
		if err != nil {
			log.Println("id", a.ID.String(), "market", 200)
			log.Println(err.Error())
		}
	}

	//id := new(big.Int)
	//id.SetString("170141183460469231731687303715884105753", 10)
	//
	//position, err := lib.GetPosition(id, big.NewInt(100))
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(position)
	//}

	// ACCOUNTS

	//id := new(big.Int)
	//id.SetString("170141183460469231731687303715884105753", 10)
	//
	//format, err := lib.FormatAccount(id)
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(format)
	//}
	//
	//margin, err := lib.GetAvailableMargin(id)
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(margin)
	//}
	//
	//rmargin, err := lib.GetRequiredMaintenanceMargin(id)
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(rmargin)
	//}
	//
	//li, err := lib.GetAccountLastInteraction(id)
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(li)
	//}
	//
	//ao, err := lib.GetAccountOwner(id)
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(ao)
	//}

	// MARKETS

	//ids, err := lib.GetMarketIDs()
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(ids)
	//}
	//
	//summary100, err := lib.GetMarketSummary(big.NewInt(100))
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(summary100)
	//}
	//
	//summary200, err := lib.GetMarketSummary(big.NewInt(200))
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(summary200)
	//}
	//
	//meta100, err := lib.GetMarketMetadata(big.NewInt(100))
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(meta100)
	//}
	//
	//meta200, err := lib.GetMarketMetadata(big.NewInt(200))
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(meta200)
	//}
	//
	//liq100, err := lib.GetLiquidationParameters(big.NewInt(100))
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(liq100)
	//}
	//
	//liq200, err := lib.GetLiquidationParameters(big.NewInt(200))
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(liq200)
	//}
	//
	//found100, err := lib.GetFoundingRate(big.NewInt(100))
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(found100)
	//}
	//
	//found200, err := lib.GetFoundingRate(big.NewInt(200))
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(found200)
	//}

	//...
	// call needed methods
	// ...

	lib.Close()
}
