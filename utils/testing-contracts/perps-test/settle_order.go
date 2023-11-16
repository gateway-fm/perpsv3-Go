package perps_test

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// SettlePythOrder settles an order by fetching data from the Pyth price feed and submitting a transaction
func (m *TestPerpsMarket) SettlePythOrder(accountID *big.Int, maxPythTries int, pythDelay time.Duration) error {
	if accountID == nil {
		accountID = m.TestAccount
	}

	latest, err := m.rpcClient.BlockNumber(context.Background())
	if err != nil {
		logger.Log().WithField("layer", "BlockNumber").Fatalf(
			"error get latest block: %v", err.Error(),
		)
	}

	opts := &bind.CallOpts{BlockNumber: big.NewInt(int64(latest))}

	ordData, err := m.perpsMarket.GetOrder(opts, m.TestAccount)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-SettlePythOrder").Fatalf("GetOrder err: %v", err)
	}

	settlementTime := time.Unix(ordData.SettlementTime.Int64(), 2)
	if time.Now().Before(settlementTime) {
		duration := time.Until(settlementTime)
		logger.Log().WithField("layer", "TestPerpsMarket-SettlePythOrder").Infof("Waiting %v until order can be settled??? Not in this attempt!", duration)
		//time.Sleep(duration)
	} else {
		logger.Log().WithField("layer", "TestPerpsMarket-SettlePythOrder").Infof("Order is ready to be settled")
	}

	settlementStrategy := GetSettlementStrategy(int(ordData.Request.MarketId.Int64()), 2)

	settlementTimeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(settlementTimeBytes, ordData.SettlementTime.Uint64())
	settlementTimeHex := hex.EncodeToString(settlementTimeBytes)
	dataParam := fmt.Sprintf("0x%s%s", hex.EncodeToString(settlementStrategy.FeedID), settlementTimeHex)

	url := fmt.Sprintf("%s%s", settlementStrategy.URL, dataParam)
	logger.Log().WithField("layer", "TestPerpsMarket-SettlePythOrder").Infof("url=%s", url)

	var priceUpdateData PriceUpdateData
	for pythTries := 0; pythTries < maxPythTries; pythTries++ {
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != http.StatusOK {
			if pythTries >= maxPythTries {
				return fmt.Errorf("price update data not available")
			}
			log.Println("Price update data not available, waiting and retrying")
			time.Sleep(pythDelay)
			continue
		}

		defer resp.Body.Close()

		if err = json.NewDecoder(resp.Body).Decode(&priceUpdateData); err != nil {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.Log().WithField("layer", "TestPerpsMarket-SettlePythOrder").Fatalf("Error reading response body: %v", err)
			}
			bodyString := string(body)

			logger.Log().WithField("layer", "TestPerpsMarket-SettlePythOrder").Infof("res=\n\n%s", bodyString)

			return fmt.Errorf("failed to decode price update data:%s, %v", bodyString, err)
		}

		logger.Log().WithField("layer", "TestPerpsMarket-SettlePythOrder").Infof("priceUpdateData:%+v", priceUpdateData.Data)

		break
	}

	accountBytes := make([]byte, 32)
	binary.BigEndian.PutUint32(accountBytes, uint32(accountID.Uint64()))
	marketBytes := make([]byte, 32)
	binary.BigEndian.PutUint32(marketBytes, uint32(ordData.Request.MarketId.Uint64()))

	extraData := append(accountBytes, marketBytes...)

	log.Printf("extra_data to string: %s", hex.EncodeToString(extraData))

	data, err := hex.DecodeString(strings.TrimPrefix(priceUpdateData.Data, "0x"))
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-SettlePythOrder").Fatalf("Failed to decode hex string priceUpdateData.Data: %v", err)
	}
	aut := m.getAut("SettlePythOrder")

	tx, err := m.perpsMarket.SettlePythOrder(aut, data, extraData)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-SettlePythOrder").Fatalf("SettlePythOrder: %+v", err)
	}
	logger.Log().WithField("layer", "TestPerpsMarket-SettlePythOrder").Infof("tx=%+v", tx)

	return nil
}

type PriceUpdateData struct {
	Data string `json:"data"`
}

type SettlementStrategy struct {
	StrategyType              int     `json:"strategy_type"`
	SettlementDelay           int     `json:"settlement_delay"`
	SettlementWindowDuration  int     `json:"settlement_window_duration"`
	PriceWindowDuration       int     `json:"price_window_duration"`
	PriceVerificationContract string  `json:"price_verification_contract"`
	FeedID                    []byte  `json:"feed_id"`
	FeedIDx                   string  `json:"feed_id_x"`
	URL                       string  `json:"url"`
	SettlementReward          float64 `json:"settlement_reward"`
	Disabled                  bool    `json:"disabled"`
}

var market100Strategy0 = SettlementStrategy{
	StrategyType:              0,
	SettlementDelay:           15,
	SettlementWindowDuration:  600,
	PriceWindowDuration:       180,
	PriceVerificationContract: "0xff1a0f4744e8582df1ae09d5611b887b6a12925c",
	FeedIDx:                   "ca80ba6dc32e08d06f1aa886011eed1d77c77be9eb761cc10d72b7d0a2fd57a6",
	URL:                       "https://api.synthetix.io/pyth-testnet/api/get_vaa_ccip?data=",
	SettlementReward:          0.0,
	Disabled:                  true,
}
var market200Strategy0 = SettlementStrategy{
	StrategyType:              0,
	SettlementDelay:           15,
	SettlementWindowDuration:  60,
	PriceWindowDuration:       180,
	PriceVerificationContract: "0xff1a0f4744e8582df1ae09d5611b887b6a12925c",
	FeedIDx:                   "f9c0172ba10dfa4d19088d94f5bf61d3b54d5bd7483a322a982e1373ee8ea31b",
	URL:                       "https://api.synthetix.io/pyth-testnet/api/get_vaa_ccip?data=",
	SettlementReward:          0.0,
	Disabled:                  true,
}
var market100Strategy2 = SettlementStrategy{
	StrategyType:              0,
	SettlementDelay:           15,
	SettlementWindowDuration:  60,
	PriceWindowDuration:       600,
	PriceVerificationContract: "0xff1a0f4744e8582df1ae09d5611b887b6a12925c",
	FeedIDx:                   "ca80ba6dc32e08d06f1aa886011eed1d77c77be9eb761cc10d72b7d0a2fd57a6",
	URL:                       "https://api.synthetix.io/pyth-testnet/api/get_vaa_ccip?data=",
	SettlementReward:          0.0,
	Disabled:                  false,
}
var market200Strategy2 = SettlementStrategy{
	StrategyType:              0,
	SettlementDelay:           15,
	SettlementWindowDuration:  60,
	PriceWindowDuration:       600,
	PriceVerificationContract: "0xff1a0f4744e8582df1ae09d5611b887b6a12925c",
	FeedIDx:                   "f9c0172ba10dfa4d19088d94f5bf61d3b54d5bd7483a322a982e1373ee8ea31b",
	URL:                       "https://api.synthetix.io/pyth-testnet/api/get_vaa_ccip?data=",
	SettlementReward:          0.0,
	Disabled:                  false,
}

func GetSettlementStrategy(market, strategy int) *SettlementStrategy {
	var res *SettlementStrategy
	switch {
	case market == 100 && strategy == 0:
		res = &market100Strategy0
	case market == 200 && strategy == 0:
		res = &market200Strategy0
	case market == 100 && strategy == 2:
		res = &market100Strategy2
	case market == 200 && strategy == 2:
		res = &market200Strategy2
	default:
		// Return nil or a default strategy if no match is found
		return nil
	}
	var err error
	res.FeedID, err = hex.DecodeString(res.FeedIDx)
	if err != nil {
		log.Fatalf("Failed to decode feed ID %s: %v", res.FeedIDx, err)
	}

	return res
}

/*
market 100, strategy 0
{'strategy_type': 0,
'settlement_delay': 15,
'settlement_window_duration': 600,
'price_window_duration': 180,
'price_verification_contract': '0xff1a0f4744e8582df1ae09d5611b887b6a12925c',
'feed_id': b'\xca\x80\xbam\xc3.\x08\xd0o\x1a\xa8\x86\x01\x1e\xed\x1dw\xc7{\xe9\xebv\x1c\xc1\rr\xb7\xd0\xa2\xfdW\xa6',
'url': 'https://api.synthetix.io/pyth-testnet/api/get_vaa_ccip?data=',
'settlement_reward': 0.0,
'disabled': True}

market 200, strategy 0
{'strategy_type': 0,
'settlement_delay': 15,
'settlement_window_duration': 60,
'price_window_duration': 180,
'price_verification_contract': '0xff1a0f4744e8582df1ae09d5611b887b6a12925c',
'feed_id': b'\xf9\xc0\x17+\xa1\r\xfaM\x19\x08\x8d\x94\xf5\xbfa\xd3\xb5M[\xd7H:2*\x98.\x13s\xee\x8e\xa3\x1b',
'url': 'https://api.synthetix.io/pyth-testnet/api/get_vaa_ccip?data=',
'settlement_reward': 0.0,
'disabled': True}
*/
/*
market 100, strategy 2
{'strategy_type': 0,
 'settlement_delay': 15,
 'settlement_window_duration': 60,
 'price_window_duration': 600,
 'price_verification_contract': '0xff1a0f4744e8582df1ae09d5611b887b6a12925c',
 'feed_id': b'\xca\x80\xbam\xc3.\x08\xd0o\x1a\xa8\x86\x01\x1e\xed\x1dw\xc7{\xe9\xebv\x1c\xc1\rr\xb7\xd0\xa2\xfdW\xa6',
 'url': 'https://api.synthetix.io/pyth-testnet/api/get_vaa_ccip?data=',
 'settlement_reward': 0.0,
 'disabled': False}

market 200, strategy 2
{'strategy_type': 0,
 'settlement_delay': 15,
 'settlement_window_duration': 60,
 'price_window_duration': 600,
 'price_verification_contract': '0xff1a0f4744e8582df1ae09d5611b887b6a12925c',
 'feed_id': b'\xf9\xc0\x17+\xa1\r\xfaM\x19\x08\x8d\x94\xf5\xbfa\xd3\xb5M[\xd7H:2*\x98.\x13s\xee\x8e\xa3\x1b',
 'url': 'https://api.synthetix.io/pyth-testnet/api/get_vaa_ccip?data=',
 'settlement_reward': 0.0,
 'disabled': False}
*/
