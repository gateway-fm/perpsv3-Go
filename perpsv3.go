package perpsv3_Go

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/config"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
	"github.com/gateway-fm/perpsv3-Go/services"
)

// IPerpsv3 is an interface for perpsv3 lib
type IPerpsv3 interface {
	// Close used to stop the lib work
	Close()
}

// Perpsv3 is a main perpsv3 lib object, it is implementing IPerpsv3 interface
type Perpsv3 struct {
	config    *config.PerpsvConfig
	service   services.IService
	rpcClient *ethclient.Client
}

// Create used to get Perpsv3 instance with given configuration settings
func Create(conf *config.PerpsvConfig) (IPerpsv3, error) {
	lib := &Perpsv3{
		config: conf,
	}

	if err := lib.init(); err != nil {
		return nil, err
	}

	return lib, nil
}

func (p *Perpsv3) Close() {
	p.rpcClient.Close()
}

// init used to initialize all lib dependencies
func (p *Perpsv3) init() error {
	if p.config.RPC == "" {
		return errors.BlankRPCURLErr
	}

	rpcClient, err := ethclient.Dial(p.config.RPC)
	if err != nil {
		logger.Log().WithField("layer", "Init").Errorf("error dial rpc: %v", err.Error())
		return errors.GetDialRPCErr(err)
	}

	p.rpcClient = rpcClient

	core, err := p.getGoerliCoreContract()
	if err != nil {
		return err
	}

	spotMarket, err := p.getGoerliSpotMarketContract()
	if err != nil {
		return err
	}

	p.service = services.NewService(rpcClient, core, p.config.CoreContractFirstBlock, spotMarket, p.config.CoreContractFirstBlock)

	return nil
}

// getGoerliSpotMarketContract is used to get spot market contract instance deployed on goerli test net
func (p *Perpsv3) getGoerliSpotMarketContract() (*spotMarketGoerli.SpotMarketGoerli, error) {
	if p.config.SpotMarketContractAddress != "" {
		addr, err := getAddr(p.config.SpotMarketContractAddress, "spot market")
		if err != nil {
			return nil, err
		}

		contract, err := spotMarketGoerli.NewSpotMarketGoerli(addr, p.rpcClient)
		if err != nil {
			logger.Log().WithField("layer", "Init").Errorf("error getting spot market contract: %v", err.Error())
			return nil, errors.GetInitContractErr(err)
		}

		return contract, nil
	} else {
		logger.Log().WithField("layer", "Init").Errorf("no spot market contract address")
		return nil, errors.BlankContractAddrErr
	}
}

// getGoerliCoreContract is used to get core contract instance deployed on goerli test net
func (p *Perpsv3) getGoerliCoreContract() (*coreGoerli.CoreGoerli, error) {
	if p.config.CoreContractAddress != "" {
		addr, err := getAddr(p.config.CoreContractAddress, "core")
		if err != nil {
			return nil, err
		}

		contract, err := coreGoerli.NewCoreGoerli(addr, p.rpcClient)
		if err != nil {
			logger.Log().WithField("layer", "Init").Errorf("error getting core contract: %v", err.Error())
			return nil, errors.GetInitContractErr(err)
		}

		return contract, nil
	} else {
		logger.Log().WithField("layer", "Init").Errorf("no core contract address")
		return nil, errors.BlankContractAddrErr
	}
}

// getAddr is used to get contract address as common.Address type
func getAddr(addr string, name string) (common.Address, error) {
	isAddr := common.IsHexAddress(addr)
	if !isAddr {
		logger.Log().WithField("layer", "Init").Errorf("invalid %v contract address", name)
		return common.Address{}, errors.InvalidContractAddrErr
	}

	return common.HexToAddress(addr), nil
}
