package perpsv3_Go

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/config"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/events"
	mock_events "github.com/gateway-fm/perpsv3-Go/mocks/events"
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
		AccountID:        big.NewInt(1692895537802),
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

func TestPerpsv3_RetrieveTradesLimit(t *testing.T) {
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
		AccountID:        big.NewInt(1692895537802),
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
		name    string
		conf    *config.PerpsvConfig
		limit   uint64
		wantRes []*models.Trade
		wantErr error
	}{
		{
			name:    "no error default values",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   0,
			wantRes: []*models.Trade{trade, trade, trade},
		},
		{
			name:    "no error custom values",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantRes: []*models.Trade{trade},
		},
		{
			name:    "no error custom values blank result",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantRes: []*models.Trade{},
		},
		{
			name:    "error",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantErr: errors.FilterErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(tt.conf)
			p.service = mockService

			mockService.EXPECT().RetrieveTradesLimit(tt.limit).Return(tt.wantRes, tt.wantErr)

			res, err := p.RetrieveTradesLimit(tt.limit)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_RetrieveOrders(t *testing.T) {
	blockN := uint64(10000)

	sizeDelta := new(big.Int)
	sizeDelta.SetString("-500000000000000000", 10)

	acceptablePrice := new(big.Int)
	acceptablePrice.SetString("1000000000000000000", 10)

	liquidation := &models.Order{
		MarketID:        100,
		AccountID:       big.NewInt(8714),
		OrderType:       0,
		SizeDelta:       sizeDelta,
		AcceptablePrice: acceptablePrice,
		SettlementTime:  1693269991,
		ExpirationTime:  1693270591,
		TrackingCode:    [32]byte{},
		Sender:          common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653f"),
		BlockNumber:     13920954,
		BlockTimestamp:  1693269976,
	}

	testCases := []struct {
		name       string
		conf       *config.PerpsvConfig
		startBlock uint64
		endBlock   *uint64
		wantRes    []*models.Order
		wantErr    error
	}{
		{
			name:       "no error default values",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: 0,
			endBlock:   nil,
			wantRes:    []*models.Order{liquidation, liquidation, liquidation},
		},
		{
			name:       "no error custom values",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: blockN,
			endBlock:   &blockN,
			wantRes:    []*models.Order{liquidation},
		},
		{
			name:       "no error custom values blank result",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: blockN,
			endBlock:   &blockN,
			wantRes:    []*models.Order{},
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

			mockService.EXPECT().RetrieveOrders(tt.startBlock, tt.endBlock).Return(tt.wantRes, tt.wantErr)

			res, err := p.RetrieveOrders(tt.startBlock, tt.endBlock)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_RetrieveOrdersLimit(t *testing.T) {
	sizeDelta := new(big.Int)
	sizeDelta.SetString("-500000000000000000", 10)

	acceptablePrice := new(big.Int)
	acceptablePrice.SetString("1000000000000000000", 10)

	liquidation := &models.Order{
		MarketID:        100,
		AccountID:       big.NewInt(8714),
		OrderType:       0,
		SizeDelta:       sizeDelta,
		AcceptablePrice: acceptablePrice,
		SettlementTime:  1693269991,
		ExpirationTime:  1693270591,
		TrackingCode:    [32]byte{},
		Sender:          common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653f"),
		BlockNumber:     13920954,
		BlockTimestamp:  1693269976,
	}

	testCases := []struct {
		name    string
		conf    *config.PerpsvConfig
		limit   uint64
		wantRes []*models.Order
		wantErr error
	}{
		{
			name:    "no error default values",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   0,
			wantRes: []*models.Order{liquidation, liquidation, liquidation},
		},
		{
			name:    "no error custom values",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantRes: []*models.Order{liquidation},
		},
		{
			name:    "no error custom values blank result",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantRes: []*models.Order{},
		},
		{
			name:    "error",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantErr: errors.FilterErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(tt.conf)
			p.service = mockService

			mockService.EXPECT().RetrieveOrdersLimit(tt.limit).Return(tt.wantRes, tt.wantErr)

			res, err := p.RetrieveOrdersLimit(tt.limit)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_RetrieveLiquidations(t *testing.T) {
	blockN := uint64(10000)

	liquidation := &models.Liquidation{
		MarketID:            100,
		AccountID:           big.NewInt(25),
		AmountLiquidated:    big.NewInt(200000000000000000),
		CurrentPositionSize: big.NewInt(0),
		BlockNumber:         14125002,
		BlockTimestamp:      1693678072,
	}

	testCases := []struct {
		name       string
		conf       *config.PerpsvConfig
		startBlock uint64
		endBlock   *uint64
		wantRes    []*models.Liquidation
		wantErr    error
	}{
		{
			name:       "no error default values",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: 0,
			endBlock:   nil,
			wantRes:    []*models.Liquidation{liquidation, liquidation, liquidation},
		},
		{
			name:       "no error custom values",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: blockN,
			endBlock:   &blockN,
			wantRes:    []*models.Liquidation{liquidation},
		},
		{
			name:       "no error custom values blank result",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: blockN,
			endBlock:   &blockN,
			wantRes:    []*models.Liquidation{},
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

			mockService.EXPECT().RetrieveLiquidations(tt.startBlock, tt.endBlock).Return(tt.wantRes, tt.wantErr)

			res, err := p.RetrieveLiquidations(tt.startBlock, tt.endBlock)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_RetrieveLiquidationsLimit(t *testing.T) {
	liquidation := &models.Liquidation{
		MarketID:            100,
		AccountID:           big.NewInt(25),
		AmountLiquidated:    big.NewInt(200000000000000000),
		CurrentPositionSize: big.NewInt(0),
		BlockNumber:         14125002,
		BlockTimestamp:      1693678072,
	}

	testCases := []struct {
		name    string
		conf    *config.PerpsvConfig
		limit   uint64
		wantRes []*models.Liquidation
		wantErr error
	}{
		{
			name:    "no error default values",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   0,
			wantRes: []*models.Liquidation{liquidation, liquidation, liquidation},
		},
		{
			name:    "no error custom values",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   10,
			wantRes: []*models.Liquidation{liquidation},
		},
		{
			name:    "no error custom values blank result",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   10,
			wantRes: []*models.Liquidation{},
		},
		{
			name:    "error",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   0,
			wantErr: errors.FilterErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(tt.conf)
			p.service = mockService

			mockService.EXPECT().RetrieveLiquidationsLimit(tt.limit).Return(tt.wantRes, tt.wantErr)

			res, err := p.RetrieveLiquidationsLimit(tt.limit)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_RetrieveMarketUpdates(t *testing.T) {
	blockN := uint64(10000)

	sizeDelta := new(big.Int)
	sizeDelta.SetString("-500000000000000000", 10)

	acceptablePrice := new(big.Int)
	acceptablePrice.SetString("1000000000000000000", 10)

	marketUpdate := &models.MarketUpdate{
		MarketID:               200,
		Price:                  3780527432113118208,
		Skew:                   527000000000000000,
		Size:                   1049000000000000000,
		SizeDelta:              -255300000000000000,
		CurrentFundingRate:     62031943958317,
		CurrentFundingVelocity: 47430000000000,
		BlockNumber:            13739029,
		BlockTimestamp:         1692906126,
		TransactionHash:        "0x16704162005c11d71c745f7392a71a5ede8eb5f042e7fa917f210748773c57bf",
	}

	testCases := []struct {
		name       string
		conf       *config.PerpsvConfig
		startBlock uint64
		endBlock   *uint64
		wantRes    []*models.MarketUpdate
		wantErr    error
	}{
		{
			name:       "no error default values",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: 0,
			endBlock:   nil,
			wantRes:    []*models.MarketUpdate{marketUpdate, marketUpdate, marketUpdate},
		},
		{
			name:       "no error custom values",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: blockN,
			endBlock:   &blockN,
			wantRes:    []*models.MarketUpdate{marketUpdate},
		},
		{
			name:       "no error custom values blank result",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: blockN,
			endBlock:   &blockN,
			wantRes:    []*models.MarketUpdate{},
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

			mockService.EXPECT().RetrieveMarketUpdates(tt.startBlock, tt.endBlock).Return(tt.wantRes, tt.wantErr)

			res, err := p.RetrieveMarketUpdates(tt.startBlock, tt.endBlock)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_RetrieveMarketUpdatesLimit(t *testing.T) {
	sizeDelta := new(big.Int)
	sizeDelta.SetString("-500000000000000000", 10)

	acceptablePrice := new(big.Int)
	acceptablePrice.SetString("1000000000000000000", 10)

	marketUpdate := &models.MarketUpdate{
		MarketID:               200,
		Price:                  3780527432113118208,
		Skew:                   527000000000000000,
		Size:                   1049000000000000000,
		SizeDelta:              -255300000000000000,
		CurrentFundingRate:     62031943958317,
		CurrentFundingVelocity: 47430000000000,
		BlockNumber:            13739029,
		BlockTimestamp:         1692906126,
		TransactionHash:        "0x16704162005c11d71c745f7392a71a5ede8eb5f042e7fa917f210748773c57bf",
	}

	testCases := []struct {
		name    string
		conf    *config.PerpsvConfig
		limit   uint64
		wantRes []*models.MarketUpdate
		wantErr error
	}{
		{
			name:    "no error default values",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   0,
			wantRes: []*models.MarketUpdate{marketUpdate, marketUpdate, marketUpdate},
		},
		{
			name:    "no error custom values",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantRes: []*models.MarketUpdate{marketUpdate},
		},
		{
			name:    "no error custom values blank result",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantRes: []*models.MarketUpdate{},
		},
		{
			name:    "error",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantErr: errors.FilterErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(tt.conf)
			p.service = mockService

			mockService.EXPECT().RetrieveMarketUpdatesLimit(tt.limit).Return(tt.wantRes, tt.wantErr)

			res, err := p.RetrieveMarketUpdatesLimit(tt.limit)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_RetrieveMarketUpdatesBig(t *testing.T) {
	blockN := uint64(10000)

	marketUpdate := &models.MarketUpdateBig{
		MarketID:               big.NewInt(200),
		Price:                  big.NewInt(3780527432113118208),
		Skew:                   big.NewInt(527000000000000000),
		Size:                   big.NewInt(1049000000000000000),
		SizeDelta:              big.NewInt(-255300000000000000),
		CurrentFundingRate:     big.NewInt(62031943958317),
		CurrentFundingVelocity: big.NewInt(47430000000000),
		BlockNumber:            13739029,
		BlockTimestamp:         1692906126,
		TransactionHash:        "0x16704162005c11d71c745f7392a71a5ede8eb5f042e7fa917f210748773c57bf",
	}

	testCases := []struct {
		name       string
		conf       *config.PerpsvConfig
		startBlock uint64
		endBlock   *uint64
		wantRes    []*models.MarketUpdateBig
		wantErr    error
	}{
		{
			name:       "no error default values",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: 0,
			endBlock:   nil,
			wantRes:    []*models.MarketUpdateBig{marketUpdate, marketUpdate, marketUpdate},
		},
		{
			name:       "no error custom values",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: blockN,
			endBlock:   &blockN,
			wantRes:    []*models.MarketUpdateBig{marketUpdate},
		},
		{
			name:       "no error custom values blank result",
			conf:       config.GetGoerliDefaultPerpsvConfig(),
			startBlock: blockN,
			endBlock:   &blockN,
			wantRes:    []*models.MarketUpdateBig{},
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

			mockService.EXPECT().RetrieveMarketUpdatesBig(tt.startBlock, tt.endBlock).Return(tt.wantRes, tt.wantErr)

			res, err := p.RetrieveMarketUpdatesBig(tt.startBlock, tt.endBlock)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_RetrieveMarketUpdatesBigLimit(t *testing.T) {
	marketUpdate := &models.MarketUpdateBig{
		MarketID:               big.NewInt(200),
		Price:                  big.NewInt(3780527432113118208),
		Skew:                   big.NewInt(527000000000000000),
		Size:                   big.NewInt(1049000000000000000),
		SizeDelta:              big.NewInt(-255300000000000000),
		CurrentFundingRate:     big.NewInt(62031943958317),
		CurrentFundingVelocity: big.NewInt(47430000000000),
		BlockNumber:            13739029,
		BlockTimestamp:         1692906126,
		TransactionHash:        "0x16704162005c11d71c745f7392a71a5ede8eb5f042e7fa917f210748773c57bf",
	}

	testCases := []struct {
		name    string
		conf    *config.PerpsvConfig
		limit   uint64
		wantRes []*models.MarketUpdateBig
		wantErr error
	}{
		{
			name:    "no error default values",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   0,
			wantRes: []*models.MarketUpdateBig{marketUpdate, marketUpdate, marketUpdate},
		},
		{
			name:    "no error custom values",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantRes: []*models.MarketUpdateBig{marketUpdate},
		},
		{
			name:    "no error custom values blank result",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantRes: []*models.MarketUpdateBig{},
		},
		{
			name:    "error",
			conf:    config.GetGoerliDefaultPerpsvConfig(),
			limit:   1,
			wantErr: errors.FilterErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(tt.conf)
			p.service = mockService

			mockService.EXPECT().RetrieveMarketUpdatesBigLimit(tt.limit).Return(tt.wantRes, tt.wantErr)

			res, err := p.RetrieveMarketUpdatesBigLimit(tt.limit)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_ListenTrades(t *testing.T) {
	testCases := []struct {
		name    string
		sub     *events.TradeSubscription
		wantErr error
	}{
		{
			name: "no error",
			sub:  &events.TradeSubscription{},
		},
		{
			name:    "error",
			sub:     &events.TradeSubscription{},
			wantErr: errors.ListenEventErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockEvents := mock_events.NewMockIEvents(ctrl)

			p, _ := createTest(config.GetGoerliDefaultPerpsvConfig())
			p.events = mockEvents

			mockEvents.EXPECT().ListenTrades().Return(tt.sub, tt.wantErr)

			res, err := p.ListenTrades()

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.sub, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_ListenOrders(t *testing.T) {
	testCases := []struct {
		name    string
		sub     *events.OrderSubscription
		wantErr error
	}{
		{
			name: "no error",
			sub:  &events.OrderSubscription{},
		},
		{
			name:    "error",
			sub:     &events.OrderSubscription{},
			wantErr: errors.ListenEventErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockEvents := mock_events.NewMockIEvents(ctrl)

			p, _ := createTest(config.GetGoerliDefaultPerpsvConfig())
			p.events = mockEvents

			mockEvents.EXPECT().ListenOrders().Return(tt.sub, tt.wantErr)

			res, err := p.ListenOrders()

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.sub, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_ListenMarketUpdates(t *testing.T) {
	testCases := []struct {
		name    string
		sub     *events.MarketUpdateSubscription
		wantErr error
	}{
		{
			name: "no error",
			sub:  &events.MarketUpdateSubscription{},
		},
		{
			name:    "error",
			sub:     &events.MarketUpdateSubscription{},
			wantErr: errors.ListenEventErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockEvents := mock_events.NewMockIEvents(ctrl)

			p, _ := createTest(config.GetGoerliDefaultPerpsvConfig())
			p.events = mockEvents

			mockEvents.EXPECT().ListenMarketUpdates().Return(tt.sub, tt.wantErr)

			res, err := p.ListenMarketUpdates()

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.sub, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_GetPosition(t *testing.T) {
	testCases := []struct {
		name      string
		accountID *big.Int
		marketID  *big.Int
		wantRes   *models.Position
		wantErr   error
	}{
		{
			name:      "no error",
			accountID: big.NewInt(0),
			marketID:  big.NewInt(100),
			wantRes: &models.Position{
				TotalPnl:       big.NewInt(1),
				AccruedFunding: big.NewInt(2),
				PositionSize:   big.NewInt(3),
				BlockTimestamp: uint64(time.Now().Unix()),
				BlockNumber:    uint64(4),
			},
		},
		{
			name:    "error",
			wantErr: errors.ReadContractErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(config.GetGoerliDefaultPerpsvConfig())
			p.service = mockService

			mockService.EXPECT().GetPosition(tt.accountID, tt.marketID).Return(tt.wantRes, tt.wantErr)

			res, err := p.GetPosition(tt.accountID, tt.marketID)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_GetMarketMetadata(t *testing.T) {
	testCases := []struct {
		name     string
		marketID *big.Int
		wantRes  *models.MarketMetadata
		wantErr  error
	}{
		{
			name:     "no error",
			marketID: big.NewInt(100),
			wantRes: &models.MarketMetadata{
				MarketID: big.NewInt(100),
				Name:     "Ethereum",
				Symbol:   "ETH",
			},
		},
		{
			name:     "no error",
			marketID: big.NewInt(300),
			wantErr:  errors.InvalidArgumentErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(config.GetGoerliDefaultPerpsvConfig())
			p.service = mockService

			mockService.EXPECT().GetMarketMetadata(tt.marketID).Return(tt.wantRes, tt.wantErr)

			res, err := p.GetMarketMetadata(tt.marketID)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_FormatAccount(t *testing.T) {
	testCases := []struct {
		name      string
		accountID *big.Int
		wantRes   *models.Account
		wantErr   error
	}{
		{
			name: "blank account",
		},
		{
			name:      "no error",
			accountID: big.NewInt(0),
			wantRes: &models.Account{
				ID: big.NewInt(1),
				Permissions: []*models.UserPermissions{
					{
						User:        common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653f"),
						Permissions: []models.Permission{0, 1, 2, 3},
					},
					{
						User:        common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653g"),
						Permissions: []models.Permission{},
					},
				},
				Owner:           common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653c"),
				LastInteraction: uint64(time.Now().Unix()),
			},
		},
		{
			name:    "error",
			wantErr: errors.ReadContractErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(config.GetGoerliDefaultPerpsvConfig())
			p.service = mockService

			mockService.EXPECT().FormatAccount(tt.accountID).Return(tt.wantRes, tt.wantErr)

			res, err := p.FormatAccount(tt.accountID)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_FormatAccounts(t *testing.T) {
	testCases := []struct {
		name    string
		wantRes []*models.Account
		wantErr error
	}{
		{
			name: "no accounts",
		},
		{
			name: "no error",
			wantRes: []*models.Account{
				{
					ID: big.NewInt(1),
					Permissions: []*models.UserPermissions{
						{
							User:        common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653f"),
							Permissions: []models.Permission{0, 1, 2, 3},
						},
						{
							User:        common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653g"),
							Permissions: []models.Permission{},
						},
					},
					Owner:           common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653c"),
					LastInteraction: uint64(time.Now().Unix()),
				},
				{
					ID: big.NewInt(1),
					Permissions: []*models.UserPermissions{
						{
							User:        common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653f"),
							Permissions: []models.Permission{0, 1, 2, 3},
						},
						{
							User:        common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653g"),
							Permissions: []models.Permission{},
						},
					},
					Owner:           common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653c"),
					LastInteraction: uint64(time.Now().Unix()),
				},
			},
		},
		{
			name:    "error",
			wantErr: errors.ReadContractErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(config.GetGoerliDefaultPerpsvConfig())
			p.service = mockService

			mockService.EXPECT().FormatAccounts().Return(tt.wantRes, tt.wantErr)

			res, err := p.FormatAccounts()

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_FormatAccountsLimit(t *testing.T) {
	testCases := []struct {
		name    string
		limit   uint64
		wantRes []*models.Account
		wantErr error
	}{
		{
			name:  "no accounts",
			limit: uint64(1),
		},
		{
			name:  "no error",
			limit: uint64(1),
			wantRes: []*models.Account{
				{
					ID: big.NewInt(1),
					Permissions: []*models.UserPermissions{
						{
							User:        common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653f"),
							Permissions: []models.Permission{0, 1, 2, 3},
						},
						{
							User:        common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653g"),
							Permissions: []models.Permission{},
						},
					},
					Owner:           common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653c"),
					LastInteraction: uint64(time.Now().Unix()),
				},
				{
					ID: big.NewInt(1),
					Permissions: []*models.UserPermissions{
						{
							User:        common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653f"),
							Permissions: []models.Permission{0, 1, 2, 3},
						},
						{
							User:        common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653g"),
							Permissions: []models.Permission{},
						},
					},
					Owner:           common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653c"),
					LastInteraction: uint64(time.Now().Unix()),
				},
			},
		},
		{
			name:    "error",
			limit:   uint64(1),
			wantErr: errors.ReadContractErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(config.GetGoerliDefaultPerpsvConfig())
			p.service = mockService

			mockService.EXPECT().FormatAccountsLimit(tt.limit).Return(tt.wantRes, tt.wantErr)

			res, err := p.FormatAccountsLimit(tt.limit)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}

func TestPerpsv3_GetMarketSummary(t *testing.T) {
	testCases := []struct {
		name     string
		marketID *big.Int
		summary  *models.MarketSummary
		wantErr  error
	}{
		{
			name:     "no error",
			marketID: big.NewInt(100),
			summary: &models.MarketSummary{
				MarketID:               big.NewInt(100),
				Size:                   big.NewInt(200),
				Skew:                   big.NewInt(300),
				MaxOpenInterest:        big.NewInt(400),
				CurrentFundingRate:     big.NewInt(500),
				CurrentFundingVelocity: big.NewInt(600),
				IndexPrice:             big.NewInt(700),
			},
		},
		{
			name:     "error",
			marketID: big.NewInt(100),
			wantErr:  errors.InvalidArgumentErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_services.NewMockIService(ctrl)

			p, _ := createTest(config.GetGoerliDefaultPerpsvConfig())
			p.service = mockService

			mockService.EXPECT().GetMarketSummary(tt.marketID).Return(tt.summary, tt.wantErr)

			res, err := p.GetMarketSummary(tt.marketID)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.summary, res)
			} else {
				require.ErrorIs(t, tt.wantErr, err)
			}
		})
	}
}
