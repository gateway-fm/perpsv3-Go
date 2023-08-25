package perpsv3_Go

import (
	"testing"
	"time"

	"github.com/gateway-fm/perpsv3-Go/config"
	"github.com/gateway-fm/perpsv3-Go/errors"
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
