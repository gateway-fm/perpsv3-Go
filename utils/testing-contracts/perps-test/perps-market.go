package perps_test

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/contracts/sUSDT"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

type TestPerpsMarket struct {
	TestAddress     common.Address
	testPK          *ecdsa.PrivateKey
	TestAccount     *big.Int
	chainID         *big.Int
	rpcClient       *ethclient.Client
	perpsMarket     *perpsMarket.PerpsMarket
	perpsMarketAddr common.Address
	snxUSDT         *sUSDT.SUSDT
}

func GetTestPerpsMarket(rpcUrl string, perpsMarketAddress string, sUSDTAddress string, chainID int) *TestPerpsMarket {
	testPKS := os.Getenv("TEST_PK")
	if testPKS == "" {
		logger.Log().WithField("layer", "TestPerpsMarket-GetTestPerpsMarket").Fatalf("no test private key in env vars")
	}

	testAccountS := os.Getenv("TEST_ACCOUNT")
	if testAccountS == "" {
		logger.Log().WithField("layer", "TestPerpsMarket-GetTestPerpsMarket").Fatalf("no test account in env vars")
	}

	rpcClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-GetTestPerpsMarket").Fatalf("dial rpc error: %v", err.Error())
	}

	perpsMarket, err := perpsMarket.NewPerpsMarket(common.HexToAddress(perpsMarketAddress), rpcClient)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-GetTestPerpsMarket").Fatalf("create perps contract err: %v", err.Error())
	}

	snxUSDT, err := sUSDT.NewSUSDT(common.HexToAddress(sUSDTAddress), rpcClient)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-GetTestPerpsMarket").Fatalf("create s-usdt contract err: %v", err.Error())
	}

	testPK, err := crypto.HexToECDSA(testPKS)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-GetTestPerpsMarket").Fatalf("create private key error: %v", err.Error())
	}

	publicKey := testPK.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Log().WithField("layer", "TestPerpsMarket-GetTestPerpsMarket").Fatalf("convert private key to public error")
	}

	testAccount := new(big.Int)
	testAccount.SetString(testAccountS, 10)

	testAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &TestPerpsMarket{
		TestAddress:     testAddress,
		testPK:          testPK,
		TestAccount:     testAccount,
		chainID:         big.NewInt(int64(chainID)),
		rpcClient:       rpcClient,
		perpsMarket:     perpsMarket,
		perpsMarketAddr: common.HexToAddress(perpsMarketAddress),
		snxUSDT:         snxUSDT,
	}
}

func (m *TestPerpsMarket) GrantPermission(idS string, permissionS string, userS string) {
	aut := m.getAut("GrantPermission")
	id := m.getBig(idS)

	var permission [32]byte
	copy(permission[:], permissionS)
	user := common.HexToAddress(userS)

	tx, err := m.perpsMarket.GrantPermission(aut, id, permission, user)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-GrantPermission").Fatalf("error grant permission: %v", err.Error())
	}

	logger.Log().WithField("layer", "TestPerpsMarket-GrantPermission").Infof("tx hash: %v", tx.Hash())
}

func (m *TestPerpsMarket) CreateAccount() {
	aut := m.getAut("CreateAccount")

	tx, err := m.perpsMarket.CreateAccount(aut)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-CreateAccount").Fatalf("error create account: %v", err.Error())
	}

	logger.Log().WithField("layer", "TestPerpsMarket-CreateAccount").Infof("tx hash: %v", tx.Hash())
}

func (m *TestPerpsMarket) SetApproval(amountS string) {
	aut := m.getAut("SetApproval")
	amount := m.getBig(amountS)

	tx, err := m.snxUSDT.Approve(aut, m.perpsMarketAddr, amount)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-SetApproval").Fatalf("error approve: %v", err.Error())
	}

	logger.Log().WithField("layer", "TestPerpsMarket-SetApproval").Infof("tx hash: %v", tx.Hash())
}

func (m *TestPerpsMarket) GetApproval() {
	allowance, err := m.snxUSDT.Allowance(nil, m.TestAddress, m.perpsMarketAddr)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-GetApproval").Fatalf("error allowance: %v", err.Error())
	}

	logger.Log().WithField("layer", "TestPerpsMarket-GetApproval").Infof("allowance: %v", allowance.String())
}

func (m *TestPerpsMarket) GetUSDTBalance() {
	balance, err := m.snxUSDT.BalanceOf(nil, m.TestAddress)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-GetUSDTBalance").Fatalf("error blanace of: %v", err.Error())
	}

	logger.Log().WithField("layer", "TestPerpsMarket-GetUSDTBalance").Infof("balance: %v", balance.String())
}

func (m *TestPerpsMarket) GetCollateralAmount(marketIDS string) {
	marketID := m.getBig(marketIDS)

	amount, err := m.perpsMarket.GetCollateralAmount(nil, m.TestAccount, marketID)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-GetCollateralAmount").Fatalf("get collateral amount err: %v", err.Error())
	}

	logger.Log().WithField("layer", "TestPerpsMarket-GetCollateralAmount").Infof("collateral amount: %v", amount.String())
}

//func (m *TestPerpsMarket) GetMaxCollateralAmount(marketIDS string) {
//	marketID := m.getBig(marketIDS)
//
//	amount, err := m.perpsMarket.GetMaxCollateralAmount(nil, marketID)
//	if err != nil {
//		logger.Log().WithField("layer", "TestPerpsMarket-GetMaxCollateralAmount").Fatalf("get max collateral amount err: %v", err.Error())
//	}
//
//	logger.Log().WithField("layer", "TestPerpsMarket-GetMaxCollateralAmount").Infof("max collateral amount: %v", amount.String())
//}

func (m *TestPerpsMarket) ModifyCollateral(marketIDS string, amountS string) {
	aut := m.getAut("ModifyCollateral")

	marketID := m.getBig(marketIDS)
	amount := m.getBig(amountS)

	tx, err := m.perpsMarket.ModifyCollateral(aut, m.TestAccount, marketID, amount)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-ModifyCollateral").Fatalf("modify collateral err: %v", err.Error())
	}

	logger.Log().WithField("layer", "TestPerpsMarket-ModifyCollateral").Infof("modify collateral, tx hash: %v", tx.Hash())
}

func (m *TestPerpsMarket) CommitOrder(marketIDS string, sizeS string) {
	aut := m.getAut("CommitOrder")
	marketID := m.getBig(marketIDS)
	size := m.getBig(sizeS)
	//acceptablePrice := m.getBig("50000000000000000")

	summary, err := m.perpsMarket.GetMarketSummary(nil, marketID)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-GetMarketSummary").Fatalf("market summary err: %v", err.Error())
	}

	logger.Log().WithField("layer", "TestPerpsMarket-CommitOrder").Infof("index price: %v", summary.IndexPrice.String())

	req := perpsMarket.AsyncOrderOrderCommitmentRequest{
		MarketId:             marketID,
		AccountId:            m.TestAccount,
		SizeDelta:            size,
		SettlementStrategyId: big.NewInt(2),
		AcceptablePrice:      summary.IndexPrice.Add(summary.IndexPrice, big.NewInt(10000)),
		TrackingCode:         [32]byte{},
		Referrer:             common.HexToAddress("0x0000000000000000000000000000000000000000"),
	}

	logger.Log().WithField("layer", "TestPerpsMarket-CommitOrder").Infof(
		"sending transaction with request:\n"+
			"- MarketID: %v\n"+
			"- AccountID: %v\n"+
			"- SizeDelta: %v\n"+
			"- SettlementStrategyId: %s\n"+
			"- AcceptablePrice: %v\n"+
			"- TrackingCode: 0\n"+
			"- Referrer: 0x0000000000000000000000000000000000000000",
		marketID.String(), m.TestAccount, size.String(), req.SettlementStrategyId.String(), summary.IndexPrice.Add(summary.IndexPrice, big.NewInt(10000)),
	)

	tx, err := m.perpsMarket.CommitOrder(aut, req)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-CommitOrder").Fatalf("commit order err: %v", err.Error())
	}

	logger.Log().WithField("layer", "TestPerpsMarket-CommitOrder").Infof("order commited, tx hash: %v", tx.Hash())

	if err = m.SettlePythOrder(m.TestAccount, 10, 1*time.Second); err != nil {
		logger.Log().WithField("layer", "SettlePythOrder").Fatalf(
			"error get latest block: %v", err.Error(),
		)
	}
}

func (m *TestPerpsMarket) LiquidateAccount(accountIDs string) {
	aut := m.getAut("Liquidate")

	accountID := m.getBig(accountIDs)

	tx, err := m.perpsMarket.Liquidate(aut, accountID)
	if err != nil {
		logger.Log().WithField("layer", "TestPerpsMarket-LiquidateAccount").Fatalf("liquidate account err: %v", err.Error())
	}

	logger.Log().WithField("layer", "TestPerpsMarket-LiquidateAccount").Infof("account liquidated, tx hash: %v", tx.Hash())
}

func (m *TestPerpsMarket) Close() {
	m.rpcClient.Close()
}

func (m *TestPerpsMarket) getBig(value string) *big.Int {
	res := new(big.Int)
	res.SetString(value, 10)
	return res
}

func (m *TestPerpsMarket) getAut(funcName string) *bind.TransactOpts {
	aut, err := bind.NewKeyedTransactorWithChainID(m.testPK, m.chainID)
	if err != nil {
		logger.Log().WithField("layer", fmt.Sprintf("TestPerpsMarket-%v", funcName)).Fatalf("error create keyd tx: %v", err.Error())
	}

	return aut
}
