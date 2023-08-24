package services

import (
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
)

func TestService_RetrieveTrades(t *testing.T) {
	rpcClient, _ := ethclient.Dial("https://rpc.goerli.optimism.gateway.fm")

	core, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)

	s := NewService(rpcClient, core, 11664658, spot, 10875051)

	sBlock := uint64(13670059)
	lBlock := uint64(13670059)

	err := s.RetrieveTrades(sBlock, &lBlock)

	if err != nil {
		log.Println("err: ", err.Error())
	}
}
