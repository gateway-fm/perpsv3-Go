package main

import (
	perpsv3_Go "github.com/gateway-fm/perpsv3-Go"
	"log"
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

	//accs, err := lib.FormatAccountsLimit(20000)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("accs", len(accs))
	//
	//good100 := 0
	//good200 := 0
	//good300 := 0
	//good400 := 0
	//good500 := 0
	//
	//for _, a := range accs {
	//	_, err := lib.GetPosition(a.ID, big.NewInt(100))
	//	if err != nil {
	//		log.Println("id", a.ID.String(), "market", 100)
	//		log.Println(err.Error())
	//	} else {
	//		good100++
	//	}
	//
	//	_, err = lib.GetPosition(a.ID, big.NewInt(200))
	//	if err != nil {
	//		log.Println("id", a.ID.String(), "market", 200)
	//		log.Println(err.Error())
	//	} else {
	//		good200++
	//	}
	//
	//	_, err = lib.GetPosition(a.ID, big.NewInt(300))
	//	if err != nil {
	//		log.Println("id", a.ID.String(), "market", 300)
	//		log.Println(err.Error())
	//	} else {
	//		good300++
	//	}
	//
	//	_, err = lib.GetPosition(a.ID, big.NewInt(400))
	//	if err != nil {
	//		log.Println("id", a.ID.String(), "market", 400)
	//		log.Println(err.Error())
	//	} else {
	//		good400++
	//	}
	//
	//	_, err = lib.GetPosition(a.ID, big.NewInt(500))
	//	if err != nil {
	//		log.Println("id", a.ID.String(), "market", 500)
	//		log.Println(err.Error())
	//	} else {
	//		good500++
	//	}
	//}
	//
	//log.Println("good100:", good100)
	//log.Println("good200:", good200)
	//log.Println("good300:", good300)
	//log.Println("good400:", good400)
	//log.Println("good500:", good500)

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

	ids, err := lib.GetMarketIDs()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(ids)
	}

	for _, i := range ids {
		log.Println("market:", i.String())

		summaryi, err := lib.GetMarketSummary(i)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("summary", summaryi)
		}

		metai, err := lib.GetMarketMetadata(i)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("meta:", metai)
		}

		liqi, err := lib.GetLiquidationParameters(i)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("liquidation:", liqi)
		}

		foundi, err := lib.GetFoundingRate(i)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("founding rate:", foundi)
		}

	}
	//...
	// call needed methods
	// ...

	lib.Close()
}
