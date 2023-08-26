package perpsv3_Go

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/config"
	"github.com/gateway-fm/perpsv3-Go/errors"
	mock_services "github.com/gateway-fm/perpsv3-Go/mocks/service"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	testCases := []struct {
		name    string
		conf    *config.PerpsvConfig
		wantErr error
	}{
		{
			name: "goerli default config",
			conf: config.GetGoerliDefaultPerpsvConfig(),
		},
		// No default perps market address on mainnet
		//
		//{
		//	name: "optimism default config",
		//	conf: config.GetOptimismDefaultPerpsvConfig(),
		//},
		{
			name: "blank rpc",
			conf: &config.PerpsvConfig{
				RPC: "",
				ContractAddresses: &config.ContractAddresses{
					Core:       "0x76490713314fCEC173f44e99346F54c6e92a8E42",
					SpotMarket: "0x5FF4b3aacdeC86782d8c757FAa638d8790799E83",
				},
				ConnectionTimeout: time.Second * 30,
				ReadTimeout:       time.Second * 15,
			},
			wantErr: errors.BlankRPCURLErr,
		},
		{
			name: "bad rpc",
			conf: &config.PerpsvConfig{
				RPC: "not-url",
				ContractAddresses: &config.ContractAddresses{
					Core:       "0x76490713314fCEC173f44e99346F54c6e92a8E42",
					SpotMarket: "0x5FF4b3aacdeC86782d8c757FAa638d8790799E83",
				},
				ConnectionTimeout: time.Second * 30,
				ReadTimeout:       time.Second * 15,
			},
			wantErr: errors.DialPRCErr,
		},
		{
			name: "no core addr",
			conf: &config.PerpsvConfig{
				RPC: "https://rpc.goerli.optimism.gateway.fm",
				ContractAddresses: &config.ContractAddresses{
					Core:       "",
					SpotMarket: "0x5FF4b3aacdeC86782d8c757FAa638d8790799E83",
				},
				ConnectionTimeout: time.Second * 30,
				ReadTimeout:       time.Second * 15,
			},
			wantErr: errors.BlankContractAddrErr,
		},
		{
			name: "bad core addr",
			conf: &config.PerpsvConfig{
				RPC: "https://rpc.goerli.optimism.gateway.fm",
				ContractAddresses: &config.ContractAddresses{
					Core:       "not-addr",
					SpotMarket: "0x5FF4b3aacdeC86782d8c757FAa638d8790799E83",
				},
				ConnectionTimeout: time.Second * 30,
				ReadTimeout:       time.Second * 15,
			},
			wantErr: errors.InvalidContractAddrErr,
		},
		{
			name: "no spot market addr",
			conf: &config.PerpsvConfig{
				RPC: "https://rpc.goerli.optimism.gateway.fm",
				ContractAddresses: &config.ContractAddresses{
					Core:       "0x76490713314fCEC173f44e99346F54c6e92a8E42",
					SpotMarket: "",
				},
				ConnectionTimeout: time.Second * 30,
				ReadTimeout:       time.Second * 15,
			},
			wantErr: errors.BlankContractAddrErr,
		},
		{
			name: "bad spot market addr",
			conf: &config.PerpsvConfig{
				RPC: "https://rpc.goerli.optimism.gateway.fm",
				ContractAddresses: &config.ContractAddresses{
					Core:       "0x76490713314fCEC173f44e99346F54c6e92a8E42",
					SpotMarket: "not-addr",
				},
				ConnectionTimeout: time.Second * 30,
				ReadTimeout:       time.Second * 15,
			},
			wantErr: errors.InvalidContractAddrErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Create(tt.conf)

			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}

func TestPerpsv3_RetrieveTrades(t *testing.T) {
	blockN := uint64(10000)

	fillPrice := new(big.Int)
	fillPrice.SetString("26050753699652653732215", 10)

	pnl := new(big.Int)
	pnl.SetString("36901526133944874", 10)

	accruedF := new(big.Int)
	accruedF.SetString("-3064923447473062", 10)

	sizeD := new(big.Int)
	sizeD.SetString("-255300000000000000", 10)

	sizeN := new(big.Int)
	sizeN.SetString("-261000000000000000", 10)

	feesT := new(big.Int)
	feesT.SetString("4655530193664925748", 10)

	trade := &models.Trade{
		MarketID:         200,
		AccountID:        1692895537802,
		FillPrice:        fillPrice,
		PnL:              pnl,
		AccruedFunding:   accruedF,
		SizeDelta:        sizeD,
		NewSize:          sizeN,
		TotalFees:        feesT,
		ReferralFees:     big.NewInt(0),
		CollectedFees:    big.NewInt(0),
		SettlementReward: big.NewInt(0),
		TrackingCode:     [32]byte{75, 87, 69, 78, 84, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Settler:          common.HexToAddress("0x4A58e0d29558111bfDc07Dc12Ca0fF7fcD0d0d75"),
		BlockNumber:      13739029,
		BlockTimestamp:   1692906126,
		TransactionHash:  "0x16704162005c11d71c745f7392a71a5ede8eb5f042e7fa917f210748773c57bf",
	}

	testCases := []struct {
		name       string
		conf       *config.PerpsvConfig
		startBlock uint64
		endBlock   *uint64
		wantRes    []*models.Trade
		wantErr    error
	}{
		{
			name:       "no error default values",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: 0,
			endBlock:   nil,
			wantRes:    []*models.Trade{trade, trade, trade},
		},
		{
			name:       "no error custom values",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: blockN,
			endBlock:   &blockN,
			wantRes:    []*models.Trade{trade},
		},
		{
			name:       "no error custom values blank result",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: blockN,
			endBlock:   &blockN,
			wantRes:    []*models.Trade{},
		},
		{
			name:       "error",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: 0,
			endBlock:   nil,
			wantErr:    errors.FilterErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(tt.conf)
			p.service = mockService

			mockService.EXPECT().RetrieveTrades(tt.startBlock, tt.endBlock).Return(tt.wantRes, tt.wantErr)

			res, err := p.RetrieveTrades(tt.startBlock, tt.endBlock)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}
