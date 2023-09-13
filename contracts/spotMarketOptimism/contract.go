// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package spotMarketOptimism

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AsyncOrderClaimData is an auto generated low-level Go binding around an user-defined struct.
type AsyncOrderClaimData struct {
	Id                      *big.Int
	Owner                   common.Address
	OrderType               uint8
	AmountEscrowed          *big.Int
	SettlementStrategyId    *big.Int
	SettlementTime          *big.Int
	MinimumSettlementAmount *big.Int
	SettledAt               *big.Int
	Referrer                common.Address
}

// OrderFeesData is an auto generated low-level Go binding around an user-defined struct.
type OrderFeesData struct {
	FixedFees       *big.Int
	UtilizationFees *big.Int
	SkewFees        *big.Int
	WrapperFees     *big.Int
}

// SettlementStrategyData is an auto generated low-level Go binding around an user-defined struct.
type SettlementStrategyData struct {
	StrategyType              uint8
	SettlementDelay           *big.Int
	SettlementWindowDuration  *big.Int
	PriceVerificationContract common.Address
	FeedId                    [32]byte
	Url                       string
	SettlementReward          *big.Int
	PriceDeviationTolerance   *big.Int
	MinimumUsdExchangeAmount  *big.Int
	MaxRoundingLoss           *big.Int
	Disabled                  bool
}

// SpotMarketOptimismMetaData contains all meta data concerning the SpotMarketOptimism contract.
var SpotMarketOptimismMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMarketOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"synthImplementation\",\"type\":\"uint256\"}],\"name\":\"InvalidSynthImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OnlyMarketOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"DecayRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nominee\",\"type\":\"address\"}],\"name\":\"MarketNominationRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"MarketOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"MarketOwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"synthImplementation\",\"type\":\"address\"}],\"name\":\"SynthImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"SynthImplementationUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"buyFeedId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sellFeedId\",\"type\":\"bytes32\"}],\"name\":\"SynthPriceDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"synthTokenAddress\",\"type\":\"address\"}],\"name\":\"SynthRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"synthetix\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usdTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracleManager\",\"type\":\"address\"}],\"name\":\"SynthetixSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"acceptMarketOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"synthOwner\",\"type\":\"address\"}],\"name\":\"createSynth\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"marketOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getSynth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"synthAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getSynthImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"minimumCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateMarketOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"renounceMarketNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"renounceMarketOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"reportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportedDebtAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setDecayRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"synthImplementation\",\"type\":\"address\"}],\"name\":\"setSynthImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISynthetixSystem\",\"name\":\"synthetix\",\"type\":\"address\"}],\"name\":\"setSynthetix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"buyFeedId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sellFeedId\",\"type\":\"bytes32\"}],\"name\":\"updatePriceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"upgradeSynthImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSynthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"synthAmountCharged\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxSynthAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxUsdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdAmountCharged\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxUsdAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"InsufficientAmountReceived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"InvalidMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrices\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"synthReturned\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SynthBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReturned\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SynthSold\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"buyExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxUsdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"buyExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmountCharged\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"}],\"name\":\"quoteBuyExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"}],\"name\":\"quoteBuyExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmountCharged\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"}],\"name\":\"quoteSellExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"}],\"name\":\"quoteSellExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthToBurn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minUsdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"sell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmountReceived\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"sellExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSynthAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"sellExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthToBurn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"IneligibleForCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InsufficientSharesAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTransaction.Type\",\"name\":\"transactionType\",\"type\":\"uint8\"}],\"name\":\"InvalidAsyncTransactionType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"asyncOrderId\",\"type\":\"uint256\"}],\"name\":\"InvalidClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InvalidCommitmentAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"}],\"name\":\"InvalidSettlementStrategy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"asyncOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"}],\"name\":\"OrderAlreadySettled\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountEscrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structAsyncOrderClaim.Data\",\"name\":\"asyncOrderClaim\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountProvided\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"OrderCommitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountProvided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"commitOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountEscrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrderClaim.Data\",\"name\":\"asyncOrderClaim\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"}],\"name\":\"getAsyncOrderClaim\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountEscrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrderClaim.Data\",\"name\":\"asyncOrderClaim\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"}],\"name\":\"InvalidSettlementStrategy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerificationResponse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"MinimumSettlementAmountNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"OutsideSettlementWindow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deviation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tolerance\",\"type\":\"uint256\"}],\"name\":\"PriceDeviationToleranceExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"}],\"name\":\"SettlementStrategyNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalOrderAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"}],\"name\":\"OrderSettled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"}],\"name\":\"settleOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"finalOrderAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"settlePythOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"finalOrderAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"SettlementStrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SettlementStrategyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumUsdExchangeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRoundingLoss\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"}],\"name\":\"addSettlementStrategy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"getSettlementStrategy\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumUsdExchangeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRoundingLoss\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"settlementStrategy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setSettlementStrategyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"InvalidCollateralType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxWrappableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToWrap\",\"type\":\"uint256\"}],\"name\":\"WrapperExceedsMaxAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUnwrapped\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesCollected\",\"type\":\"uint256\"}],\"name\":\"SynthUnwrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWrapped\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesCollected\",\"type\":\"uint256\"}],\"name\":\"SynthWrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wrapCollateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxWrappableAmount\",\"type\":\"uint256\"}],\"name\":\"WrapperSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"wrapCollateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxWrappableAmount\",\"type\":\"uint256\"}],\"name\":\"setWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"unwrapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnCollateralAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"wrapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToMint\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"InvalidCollateralLeverage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidFeeCollector\",\"type\":\"address\"}],\"name\":\"InvalidFeeCollectorInterface\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWrapperFees\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"asyncFixedFee\",\"type\":\"uint256\"}],\"name\":\"AsyncFixedFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"atomicFixedFee\",\"type\":\"uint256\"}],\"name\":\"AtomicFixedFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralLeverage\",\"type\":\"uint256\"}],\"name\":\"CollateralLeverageSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"name\":\"MarketSkewScaleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"utilizationFeeRate\",\"type\":\"uint256\"}],\"name\":\"MarketUtilizationFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharePercentage\",\"type\":\"uint256\"}],\"name\":\"ReferrerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fixedFeeAmount\",\"type\":\"uint256\"}],\"name\":\"TransactorFixedFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"wrapFee\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"unwrapFee\",\"type\":\"int256\"}],\"name\":\"WrapperFeesSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getCollateralLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralLeverage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"transactor\",\"type\":\"address\"}],\"name\":\"getCustomTransactorFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fixedFeeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"atomicFixedFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"asyncFixedFee\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"wrapFee\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"unwrapFee\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketSkewScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketUtilizationFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"utilizationFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"getReferrerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharePercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"asyncFixedFee\",\"type\":\"uint256\"}],\"name\":\"setAsyncFixedFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"atomicFixedFee\",\"type\":\"uint256\"}],\"name\":\"setAtomicFixedFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralLeverage\",\"type\":\"uint256\"}],\"name\":\"setCollateralLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"transactor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fixedFeeAmount\",\"type\":\"uint256\"}],\"name\":\"setCustomTransactorFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"name\":\"setMarketSkewScale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFeeRate\",\"type\":\"uint256\"}],\"name\":\"setMarketUtilizationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"int256\",\"name\":\"wrapFee\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"unwrapFee\",\"type\":\"int256\"}],\"name\":\"setWrapperFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sharePercentage\",\"type\":\"uint256\"}],\"name\":\"updateReferrerShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SpotMarketOptimismABI is the input ABI used to generate the binding from.
// Deprecated: Use SpotMarketOptimismMetaData.ABI instead.
var SpotMarketOptimismABI = SpotMarketOptimismMetaData.ABI

// SpotMarketOptimism is an auto generated Go binding around an Ethereum contract.
type SpotMarketOptimism struct {
	SpotMarketOptimismCaller     // Read-only binding to the contract
	SpotMarketOptimismTransactor // Write-only binding to the contract
	SpotMarketOptimismFilterer   // Log filterer for contract events
}

// SpotMarketOptimismCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpotMarketOptimismCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotMarketOptimismTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpotMarketOptimismTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotMarketOptimismFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpotMarketOptimismFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotMarketOptimismSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpotMarketOptimismSession struct {
	Contract     *SpotMarketOptimism // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SpotMarketOptimismCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpotMarketOptimismCallerSession struct {
	Contract *SpotMarketOptimismCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SpotMarketOptimismTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpotMarketOptimismTransactorSession struct {
	Contract     *SpotMarketOptimismTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SpotMarketOptimismRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpotMarketOptimismRaw struct {
	Contract *SpotMarketOptimism // Generic contract binding to access the raw methods on
}

// SpotMarketOptimismCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpotMarketOptimismCallerRaw struct {
	Contract *SpotMarketOptimismCaller // Generic read-only contract binding to access the raw methods on
}

// SpotMarketOptimismTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpotMarketOptimismTransactorRaw struct {
	Contract *SpotMarketOptimismTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpotMarketOptimism creates a new instance of SpotMarketOptimism, bound to a specific deployed contract.
func NewSpotMarketOptimism(address common.Address, backend bind.ContractBackend) (*SpotMarketOptimism, error) {
	contract, err := bindSpotMarketOptimism(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimism{SpotMarketOptimismCaller: SpotMarketOptimismCaller{contract: contract}, SpotMarketOptimismTransactor: SpotMarketOptimismTransactor{contract: contract}, SpotMarketOptimismFilterer: SpotMarketOptimismFilterer{contract: contract}}, nil
}

// NewSpotMarketOptimismCaller creates a new read-only instance of SpotMarketOptimism, bound to a specific deployed contract.
func NewSpotMarketOptimismCaller(address common.Address, caller bind.ContractCaller) (*SpotMarketOptimismCaller, error) {
	contract, err := bindSpotMarketOptimism(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismCaller{contract: contract}, nil
}

// NewSpotMarketOptimismTransactor creates a new write-only instance of SpotMarketOptimism, bound to a specific deployed contract.
func NewSpotMarketOptimismTransactor(address common.Address, transactor bind.ContractTransactor) (*SpotMarketOptimismTransactor, error) {
	contract, err := bindSpotMarketOptimism(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismTransactor{contract: contract}, nil
}

// NewSpotMarketOptimismFilterer creates a new log filterer instance of SpotMarketOptimism, bound to a specific deployed contract.
func NewSpotMarketOptimismFilterer(address common.Address, filterer bind.ContractFilterer) (*SpotMarketOptimismFilterer, error) {
	contract, err := bindSpotMarketOptimism(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismFilterer{contract: contract}, nil
}

// bindSpotMarketOptimism binds a generic wrapper to an already deployed contract.
func bindSpotMarketOptimism(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SpotMarketOptimismMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SpotMarketOptimism *SpotMarketOptimismRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SpotMarketOptimism.Contract.SpotMarketOptimismCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SpotMarketOptimism *SpotMarketOptimismRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SpotMarketOptimismTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SpotMarketOptimism *SpotMarketOptimismRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SpotMarketOptimismTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SpotMarketOptimism *SpotMarketOptimismCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SpotMarketOptimism.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SpotMarketOptimism *SpotMarketOptimismTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SpotMarketOptimism *SpotMarketOptimismTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_SpotMarketOptimism *SpotMarketOptimismSession) PRECISION() (*big.Int, error) {
	return _SpotMarketOptimism.Contract.PRECISION(&_SpotMarketOptimism.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) PRECISION() (*big.Int, error) {
	return _SpotMarketOptimism.Contract.PRECISION(&_SpotMarketOptimism.CallOpts)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetAssociatedSystem(opts *bind.CallOpts, id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getAssociatedSystem", id)

	outstruct := new(struct {
		Addr common.Address
		Kind [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Kind = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _SpotMarketOptimism.Contract.GetAssociatedSystem(&_SpotMarketOptimism.CallOpts, id)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _SpotMarketOptimism.Contract.GetAssociatedSystem(&_SpotMarketOptimism.CallOpts, id)
}

// GetAsyncOrderClaim is a free data retrieval call binding the contract method 0x5381ce16.
//
// Solidity: function getAsyncOrderClaim(uint128 marketId, uint128 asyncOrderId) pure returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetAsyncOrderClaim(opts *bind.CallOpts, marketId *big.Int, asyncOrderId *big.Int) (AsyncOrderClaimData, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getAsyncOrderClaim", marketId, asyncOrderId)

	if err != nil {
		return *new(AsyncOrderClaimData), err
	}

	out0 := *abi.ConvertType(out[0], new(AsyncOrderClaimData)).(*AsyncOrderClaimData)

	return out0, err

}

// GetAsyncOrderClaim is a free data retrieval call binding the contract method 0x5381ce16.
//
// Solidity: function getAsyncOrderClaim(uint128 marketId, uint128 asyncOrderId) pure returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetAsyncOrderClaim(marketId *big.Int, asyncOrderId *big.Int) (AsyncOrderClaimData, error) {
	return _SpotMarketOptimism.Contract.GetAsyncOrderClaim(&_SpotMarketOptimism.CallOpts, marketId, asyncOrderId)
}

// GetAsyncOrderClaim is a free data retrieval call binding the contract method 0x5381ce16.
//
// Solidity: function getAsyncOrderClaim(uint128 marketId, uint128 asyncOrderId) pure returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetAsyncOrderClaim(marketId *big.Int, asyncOrderId *big.Int) (AsyncOrderClaimData, error) {
	return _SpotMarketOptimism.Contract.GetAsyncOrderClaim(&_SpotMarketOptimism.CallOpts, marketId, asyncOrderId)
}

// GetCollateralLeverage is a free data retrieval call binding the contract method 0xcdfaef0f.
//
// Solidity: function getCollateralLeverage(uint128 synthMarketId) view returns(uint256 collateralLeverage)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetCollateralLeverage(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getCollateralLeverage", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralLeverage is a free data retrieval call binding the contract method 0xcdfaef0f.
//
// Solidity: function getCollateralLeverage(uint128 synthMarketId) view returns(uint256 collateralLeverage)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetCollateralLeverage(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.GetCollateralLeverage(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetCollateralLeverage is a free data retrieval call binding the contract method 0xcdfaef0f.
//
// Solidity: function getCollateralLeverage(uint128 synthMarketId) view returns(uint256 collateralLeverage)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetCollateralLeverage(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.GetCollateralLeverage(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetCustomTransactorFees is a free data retrieval call binding the contract method 0x2efaa971.
//
// Solidity: function getCustomTransactorFees(uint128 synthMarketId, address transactor) view returns(uint256 fixedFeeAmount)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetCustomTransactorFees(opts *bind.CallOpts, synthMarketId *big.Int, transactor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getCustomTransactorFees", synthMarketId, transactor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCustomTransactorFees is a free data retrieval call binding the contract method 0x2efaa971.
//
// Solidity: function getCustomTransactorFees(uint128 synthMarketId, address transactor) view returns(uint256 fixedFeeAmount)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.GetCustomTransactorFees(&_SpotMarketOptimism.CallOpts, synthMarketId, transactor)
}

// GetCustomTransactorFees is a free data retrieval call binding the contract method 0x2efaa971.
//
// Solidity: function getCustomTransactorFees(uint128 synthMarketId, address transactor) view returns(uint256 fixedFeeAmount)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.GetCustomTransactorFees(&_SpotMarketOptimism.CallOpts, synthMarketId, transactor)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetDeniers(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getDeniers", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _SpotMarketOptimism.Contract.GetDeniers(&_SpotMarketOptimism.CallOpts, feature)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _SpotMarketOptimism.Contract.GetDeniers(&_SpotMarketOptimism.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetFeatureFlagAllowAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getFeatureFlagAllowAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _SpotMarketOptimism.Contract.GetFeatureFlagAllowAll(&_SpotMarketOptimism.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _SpotMarketOptimism.Contract.GetFeatureFlagAllowAll(&_SpotMarketOptimism.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetFeatureFlagAllowlist(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getFeatureFlagAllowlist", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _SpotMarketOptimism.Contract.GetFeatureFlagAllowlist(&_SpotMarketOptimism.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _SpotMarketOptimism.Contract.GetFeatureFlagAllowlist(&_SpotMarketOptimism.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetFeatureFlagDenyAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getFeatureFlagDenyAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _SpotMarketOptimism.Contract.GetFeatureFlagDenyAll(&_SpotMarketOptimism.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _SpotMarketOptimism.Contract.GetFeatureFlagDenyAll(&_SpotMarketOptimism.CallOpts, feature)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x5497eb23.
//
// Solidity: function getFeeCollector(uint128 synthMarketId) view returns(address feeCollector)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetFeeCollector(opts *bind.CallOpts, synthMarketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getFeeCollector", synthMarketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeCollector is a free data retrieval call binding the contract method 0x5497eb23.
//
// Solidity: function getFeeCollector(uint128 synthMarketId) view returns(address feeCollector)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetFeeCollector(synthMarketId *big.Int) (common.Address, error) {
	return _SpotMarketOptimism.Contract.GetFeeCollector(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x5497eb23.
//
// Solidity: function getFeeCollector(uint128 synthMarketId) view returns(address feeCollector)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetFeeCollector(synthMarketId *big.Int) (common.Address, error) {
	return _SpotMarketOptimism.Contract.GetFeeCollector(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetImplementation() (common.Address, error) {
	return _SpotMarketOptimism.Contract.GetImplementation(&_SpotMarketOptimism.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetImplementation() (common.Address, error) {
	return _SpotMarketOptimism.Contract.GetImplementation(&_SpotMarketOptimism.CallOpts)
}

// GetMarketFees is a free data retrieval call binding the contract method 0x32598e61.
//
// Solidity: function getMarketFees(uint128 synthMarketId) view returns(uint256 atomicFixedFee, uint256 asyncFixedFee, int256 wrapFee, int256 unwrapFee)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetMarketFees(opts *bind.CallOpts, synthMarketId *big.Int) (struct {
	AtomicFixedFee *big.Int
	AsyncFixedFee  *big.Int
	WrapFee        *big.Int
	UnwrapFee      *big.Int
}, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getMarketFees", synthMarketId)

	outstruct := new(struct {
		AtomicFixedFee *big.Int
		AsyncFixedFee  *big.Int
		WrapFee        *big.Int
		UnwrapFee      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AtomicFixedFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AsyncFixedFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WrapFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnwrapFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMarketFees is a free data retrieval call binding the contract method 0x32598e61.
//
// Solidity: function getMarketFees(uint128 synthMarketId) view returns(uint256 atomicFixedFee, uint256 asyncFixedFee, int256 wrapFee, int256 unwrapFee)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetMarketFees(synthMarketId *big.Int) (struct {
	AtomicFixedFee *big.Int
	AsyncFixedFee  *big.Int
	WrapFee        *big.Int
	UnwrapFee      *big.Int
}, error) {
	return _SpotMarketOptimism.Contract.GetMarketFees(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetMarketFees is a free data retrieval call binding the contract method 0x32598e61.
//
// Solidity: function getMarketFees(uint128 synthMarketId) view returns(uint256 atomicFixedFee, uint256 asyncFixedFee, int256 wrapFee, int256 unwrapFee)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetMarketFees(synthMarketId *big.Int) (struct {
	AtomicFixedFee *big.Int
	AsyncFixedFee  *big.Int
	WrapFee        *big.Int
	UnwrapFee      *big.Int
}, error) {
	return _SpotMarketOptimism.Contract.GetMarketFees(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetMarketOwner is a free data retrieval call binding the contract method 0xa7b8cb9f.
//
// Solidity: function getMarketOwner(uint128 synthMarketId) view returns(address marketOwner)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetMarketOwner(opts *bind.CallOpts, synthMarketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getMarketOwner", synthMarketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMarketOwner is a free data retrieval call binding the contract method 0xa7b8cb9f.
//
// Solidity: function getMarketOwner(uint128 synthMarketId) view returns(address marketOwner)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetMarketOwner(synthMarketId *big.Int) (common.Address, error) {
	return _SpotMarketOptimism.Contract.GetMarketOwner(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetMarketOwner is a free data retrieval call binding the contract method 0xa7b8cb9f.
//
// Solidity: function getMarketOwner(uint128 synthMarketId) view returns(address marketOwner)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetMarketOwner(synthMarketId *big.Int) (common.Address, error) {
	return _SpotMarketOptimism.Contract.GetMarketOwner(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetMarketSkewScale is a free data retrieval call binding the contract method 0x8d105571.
//
// Solidity: function getMarketSkewScale(uint128 synthMarketId) view returns(uint256 skewScale)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetMarketSkewScale(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getMarketSkewScale", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketSkewScale is a free data retrieval call binding the contract method 0x8d105571.
//
// Solidity: function getMarketSkewScale(uint128 synthMarketId) view returns(uint256 skewScale)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetMarketSkewScale(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.GetMarketSkewScale(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetMarketSkewScale is a free data retrieval call binding the contract method 0x8d105571.
//
// Solidity: function getMarketSkewScale(uint128 synthMarketId) view returns(uint256 skewScale)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetMarketSkewScale(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.GetMarketSkewScale(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetMarketUtilizationFees is a free data retrieval call binding the contract method 0xf375f324.
//
// Solidity: function getMarketUtilizationFees(uint128 synthMarketId) view returns(uint256 utilizationFeeRate)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetMarketUtilizationFees(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getMarketUtilizationFees", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketUtilizationFees is a free data retrieval call binding the contract method 0xf375f324.
//
// Solidity: function getMarketUtilizationFees(uint128 synthMarketId) view returns(uint256 utilizationFeeRate)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetMarketUtilizationFees(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.GetMarketUtilizationFees(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetMarketUtilizationFees is a free data retrieval call binding the contract method 0xf375f324.
//
// Solidity: function getMarketUtilizationFees(uint128 synthMarketId) view returns(uint256 utilizationFeeRate)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetMarketUtilizationFees(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.GetMarketUtilizationFees(&_SpotMarketOptimism.CallOpts, synthMarketId)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xfa4b28ed.
//
// Solidity: function getReferrerShare(uint128 synthMarketId, address referrer) view returns(uint256 sharePercentage)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetReferrerShare(opts *bind.CallOpts, synthMarketId *big.Int, referrer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getReferrerShare", synthMarketId, referrer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReferrerShare is a free data retrieval call binding the contract method 0xfa4b28ed.
//
// Solidity: function getReferrerShare(uint128 synthMarketId, address referrer) view returns(uint256 sharePercentage)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetReferrerShare(synthMarketId *big.Int, referrer common.Address) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.GetReferrerShare(&_SpotMarketOptimism.CallOpts, synthMarketId, referrer)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xfa4b28ed.
//
// Solidity: function getReferrerShare(uint128 synthMarketId, address referrer) view returns(uint256 sharePercentage)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetReferrerShare(synthMarketId *big.Int, referrer common.Address) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.GetReferrerShare(&_SpotMarketOptimism.CallOpts, synthMarketId, referrer)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) settlementStrategy)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetSettlementStrategy(opts *bind.CallOpts, marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getSettlementStrategy", marketId, strategyId)

	if err != nil {
		return *new(SettlementStrategyData), err
	}

	out0 := *abi.ConvertType(out[0], new(SettlementStrategyData)).(*SettlementStrategyData)

	return out0, err

}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) settlementStrategy)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _SpotMarketOptimism.Contract.GetSettlementStrategy(&_SpotMarketOptimism.CallOpts, marketId, strategyId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) settlementStrategy)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _SpotMarketOptimism.Contract.GetSettlementStrategy(&_SpotMarketOptimism.CallOpts, marketId, strategyId)
}

// GetSynth is a free data retrieval call binding the contract method 0x69e0365f.
//
// Solidity: function getSynth(uint128 marketId) view returns(address synthAddress)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetSynth(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getSynth", marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSynth is a free data retrieval call binding the contract method 0x69e0365f.
//
// Solidity: function getSynth(uint128 marketId) view returns(address synthAddress)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetSynth(marketId *big.Int) (common.Address, error) {
	return _SpotMarketOptimism.Contract.GetSynth(&_SpotMarketOptimism.CallOpts, marketId)
}

// GetSynth is a free data retrieval call binding the contract method 0x69e0365f.
//
// Solidity: function getSynth(uint128 marketId) view returns(address synthAddress)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetSynth(marketId *big.Int) (common.Address, error) {
	return _SpotMarketOptimism.Contract.GetSynth(&_SpotMarketOptimism.CallOpts, marketId)
}

// GetSynthImpl is a free data retrieval call binding the contract method 0x3e0c76ca.
//
// Solidity: function getSynthImpl(uint128 marketId) view returns(address implAddress)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) GetSynthImpl(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "getSynthImpl", marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSynthImpl is a free data retrieval call binding the contract method 0x3e0c76ca.
//
// Solidity: function getSynthImpl(uint128 marketId) view returns(address implAddress)
func (_SpotMarketOptimism *SpotMarketOptimismSession) GetSynthImpl(marketId *big.Int) (common.Address, error) {
	return _SpotMarketOptimism.Contract.GetSynthImpl(&_SpotMarketOptimism.CallOpts, marketId)
}

// GetSynthImpl is a free data retrieval call binding the contract method 0x3e0c76ca.
//
// Solidity: function getSynthImpl(uint128 marketId) view returns(address implAddress)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) GetSynthImpl(marketId *big.Int) (common.Address, error) {
	return _SpotMarketOptimism.Contract.GetSynthImpl(&_SpotMarketOptimism.CallOpts, marketId)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) IsFeatureAllowed(opts *bind.CallOpts, feature [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "isFeatureAllowed", feature, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_SpotMarketOptimism *SpotMarketOptimismSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _SpotMarketOptimism.Contract.IsFeatureAllowed(&_SpotMarketOptimism.CallOpts, feature, account)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _SpotMarketOptimism.Contract.IsFeatureAllowed(&_SpotMarketOptimism.CallOpts, feature, account)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 marketId) view returns(uint256 lockedAmount)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) MinimumCredit(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "minimumCredit", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 marketId) view returns(uint256 lockedAmount)
func (_SpotMarketOptimism *SpotMarketOptimismSession) MinimumCredit(marketId *big.Int) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.MinimumCredit(&_SpotMarketOptimism.CallOpts, marketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 marketId) view returns(uint256 lockedAmount)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) MinimumCredit(marketId *big.Int) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.MinimumCredit(&_SpotMarketOptimism.CallOpts, marketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 marketId) view returns(string marketName)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) Name(opts *bind.CallOpts, marketId *big.Int) (string, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "name", marketId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 marketId) view returns(string marketName)
func (_SpotMarketOptimism *SpotMarketOptimismSession) Name(marketId *big.Int) (string, error) {
	return _SpotMarketOptimism.Contract.Name(&_SpotMarketOptimism.CallOpts, marketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 marketId) view returns(string marketName)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) Name(marketId *big.Int) (string, error) {
	return _SpotMarketOptimism.Contract.Name(&_SpotMarketOptimism.CallOpts, marketId)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SpotMarketOptimism *SpotMarketOptimismSession) NominatedOwner() (common.Address, error) {
	return _SpotMarketOptimism.Contract.NominatedOwner(&_SpotMarketOptimism.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) NominatedOwner() (common.Address, error) {
	return _SpotMarketOptimism.Contract.NominatedOwner(&_SpotMarketOptimism.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpotMarketOptimism *SpotMarketOptimismSession) Owner() (common.Address, error) {
	return _SpotMarketOptimism.Contract.Owner(&_SpotMarketOptimism.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) Owner() (common.Address, error) {
	return _SpotMarketOptimism.Contract.Owner(&_SpotMarketOptimism.CallOpts)
}

// QuoteBuyExactIn is a free data retrieval call binding the contract method 0x6b5e6ae4.
//
// Solidity: function quoteBuyExactIn(uint128 marketId, uint256 usdAmount) view returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) QuoteBuyExactIn(opts *bind.CallOpts, marketId *big.Int, usdAmount *big.Int) (struct {
	SynthAmount *big.Int
	Fees        OrderFeesData
}, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "quoteBuyExactIn", marketId, usdAmount)

	outstruct := new(struct {
		SynthAmount *big.Int
		Fees        OrderFeesData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SynthAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fees = *abi.ConvertType(out[1], new(OrderFeesData)).(*OrderFeesData)

	return *outstruct, err

}

// QuoteBuyExactIn is a free data retrieval call binding the contract method 0x6b5e6ae4.
//
// Solidity: function quoteBuyExactIn(uint128 marketId, uint256 usdAmount) view returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) QuoteBuyExactIn(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthAmount *big.Int
	Fees        OrderFeesData
}, error) {
	return _SpotMarketOptimism.Contract.QuoteBuyExactIn(&_SpotMarketOptimism.CallOpts, marketId, usdAmount)
}

// QuoteBuyExactIn is a free data retrieval call binding the contract method 0x6b5e6ae4.
//
// Solidity: function quoteBuyExactIn(uint128 marketId, uint256 usdAmount) view returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) QuoteBuyExactIn(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthAmount *big.Int
	Fees        OrderFeesData
}, error) {
	return _SpotMarketOptimism.Contract.QuoteBuyExactIn(&_SpotMarketOptimism.CallOpts, marketId, usdAmount)
}

// QuoteBuyExactOut is a free data retrieval call binding the contract method 0x1f3f7640.
//
// Solidity: function quoteBuyExactOut(uint128 marketId, uint256 synthAmount) view returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) QuoteBuyExactOut(opts *bind.CallOpts, marketId *big.Int, synthAmount *big.Int) (struct {
	UsdAmountCharged *big.Int
	Fees             OrderFeesData
}, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "quoteBuyExactOut", marketId, synthAmount)

	outstruct := new(struct {
		UsdAmountCharged *big.Int
		Fees             OrderFeesData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UsdAmountCharged = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fees = *abi.ConvertType(out[1], new(OrderFeesData)).(*OrderFeesData)

	return *outstruct, err

}

// QuoteBuyExactOut is a free data retrieval call binding the contract method 0x1f3f7640.
//
// Solidity: function quoteBuyExactOut(uint128 marketId, uint256 synthAmount) view returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) QuoteBuyExactOut(marketId *big.Int, synthAmount *big.Int) (struct {
	UsdAmountCharged *big.Int
	Fees             OrderFeesData
}, error) {
	return _SpotMarketOptimism.Contract.QuoteBuyExactOut(&_SpotMarketOptimism.CallOpts, marketId, synthAmount)
}

// QuoteBuyExactOut is a free data retrieval call binding the contract method 0x1f3f7640.
//
// Solidity: function quoteBuyExactOut(uint128 marketId, uint256 synthAmount) view returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) QuoteBuyExactOut(marketId *big.Int, synthAmount *big.Int) (struct {
	UsdAmountCharged *big.Int
	Fees             OrderFeesData
}, error) {
	return _SpotMarketOptimism.Contract.QuoteBuyExactOut(&_SpotMarketOptimism.CallOpts, marketId, synthAmount)
}

// QuoteSellExactIn is a free data retrieval call binding the contract method 0xc52d1730.
//
// Solidity: function quoteSellExactIn(uint128 marketId, uint256 synthAmount) view returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) QuoteSellExactIn(opts *bind.CallOpts, marketId *big.Int, synthAmount *big.Int) (struct {
	ReturnAmount *big.Int
	Fees         OrderFeesData
}, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "quoteSellExactIn", marketId, synthAmount)

	outstruct := new(struct {
		ReturnAmount *big.Int
		Fees         OrderFeesData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReturnAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fees = *abi.ConvertType(out[1], new(OrderFeesData)).(*OrderFeesData)

	return *outstruct, err

}

// QuoteSellExactIn is a free data retrieval call binding the contract method 0xc52d1730.
//
// Solidity: function quoteSellExactIn(uint128 marketId, uint256 synthAmount) view returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) QuoteSellExactIn(marketId *big.Int, synthAmount *big.Int) (struct {
	ReturnAmount *big.Int
	Fees         OrderFeesData
}, error) {
	return _SpotMarketOptimism.Contract.QuoteSellExactIn(&_SpotMarketOptimism.CallOpts, marketId, synthAmount)
}

// QuoteSellExactIn is a free data retrieval call binding the contract method 0xc52d1730.
//
// Solidity: function quoteSellExactIn(uint128 marketId, uint256 synthAmount) view returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) QuoteSellExactIn(marketId *big.Int, synthAmount *big.Int) (struct {
	ReturnAmount *big.Int
	Fees         OrderFeesData
}, error) {
	return _SpotMarketOptimism.Contract.QuoteSellExactIn(&_SpotMarketOptimism.CallOpts, marketId, synthAmount)
}

// QuoteSellExactOut is a free data retrieval call binding the contract method 0x2c007522.
//
// Solidity: function quoteSellExactOut(uint128 marketId, uint256 usdAmount) view returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) QuoteSellExactOut(opts *bind.CallOpts, marketId *big.Int, usdAmount *big.Int) (struct {
	SynthToBurn *big.Int
	Fees        OrderFeesData
}, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "quoteSellExactOut", marketId, usdAmount)

	outstruct := new(struct {
		SynthToBurn *big.Int
		Fees        OrderFeesData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SynthToBurn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fees = *abi.ConvertType(out[1], new(OrderFeesData)).(*OrderFeesData)

	return *outstruct, err

}

// QuoteSellExactOut is a free data retrieval call binding the contract method 0x2c007522.
//
// Solidity: function quoteSellExactOut(uint128 marketId, uint256 usdAmount) view returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) QuoteSellExactOut(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthToBurn *big.Int
	Fees        OrderFeesData
}, error) {
	return _SpotMarketOptimism.Contract.QuoteSellExactOut(&_SpotMarketOptimism.CallOpts, marketId, usdAmount)
}

// QuoteSellExactOut is a free data retrieval call binding the contract method 0x2c007522.
//
// Solidity: function quoteSellExactOut(uint128 marketId, uint256 usdAmount) view returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) QuoteSellExactOut(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthToBurn *big.Int
	Fees        OrderFeesData
}, error) {
	return _SpotMarketOptimism.Contract.QuoteSellExactOut(&_SpotMarketOptimism.CallOpts, marketId, usdAmount)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 marketId) view returns(uint256 reportedDebtAmount)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) ReportedDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "reportedDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 marketId) view returns(uint256 reportedDebtAmount)
func (_SpotMarketOptimism *SpotMarketOptimismSession) ReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.ReportedDebt(&_SpotMarketOptimism.CallOpts, marketId)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 marketId) view returns(uint256 reportedDebtAmount)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) ReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _SpotMarketOptimism.Contract.ReportedDebt(&_SpotMarketOptimism.CallOpts, marketId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool isSupported)
func (_SpotMarketOptimism *SpotMarketOptimismCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SpotMarketOptimism.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool isSupported)
func (_SpotMarketOptimism *SpotMarketOptimismSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SpotMarketOptimism.Contract.SupportsInterface(&_SpotMarketOptimism.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool isSupported)
func (_SpotMarketOptimism *SpotMarketOptimismCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SpotMarketOptimism.Contract.SupportsInterface(&_SpotMarketOptimism.CallOpts, interfaceId)
}

// AcceptMarketOwnership is a paid mutator transaction binding the contract method 0x1c216a0e.
//
// Solidity: function acceptMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) AcceptMarketOwnership(opts *bind.TransactOpts, synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "acceptMarketOwnership", synthMarketId)
}

// AcceptMarketOwnership is a paid mutator transaction binding the contract method 0x1c216a0e.
//
// Solidity: function acceptMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) AcceptMarketOwnership(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.AcceptMarketOwnership(&_SpotMarketOptimism.TransactOpts, synthMarketId)
}

// AcceptMarketOwnership is a paid mutator transaction binding the contract method 0x1c216a0e.
//
// Solidity: function acceptMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) AcceptMarketOwnership(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.AcceptMarketOwnership(&_SpotMarketOptimism.TransactOpts, synthMarketId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) AcceptOwnership() (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.AcceptOwnership(&_SpotMarketOptimism.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.AcceptOwnership(&_SpotMarketOptimism.TransactOpts)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x97b30e6d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) AddSettlementStrategy(opts *bind.TransactOpts, marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "addSettlementStrategy", marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x97b30e6d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_SpotMarketOptimism *SpotMarketOptimismSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.AddSettlementStrategy(&_SpotMarketOptimism.TransactOpts, marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x97b30e6d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.AddSettlementStrategy(&_SpotMarketOptimism.TransactOpts, marketId, strategy)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) AddToFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "addToFeatureFlagAllowlist", feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.AddToFeatureFlagAllowlist(&_SpotMarketOptimism.TransactOpts, feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.AddToFeatureFlagAllowlist(&_SpotMarketOptimism.TransactOpts, feature, account)
}

// Buy is a paid mutator transaction binding the contract method 0x37fb3369.
//
// Solidity: function buy(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) Buy(opts *bind.TransactOpts, marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "buy", marketId, usdAmount, minAmountReceived, referrer)
}

// Buy is a paid mutator transaction binding the contract method 0x37fb3369.
//
// Solidity: function buy(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) Buy(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.Buy(&_SpotMarketOptimism.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// Buy is a paid mutator transaction binding the contract method 0x37fb3369.
//
// Solidity: function buy(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) Buy(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.Buy(&_SpotMarketOptimism.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactIn is a paid mutator transaction binding the contract method 0xa12d9400.
//
// Solidity: function buyExactIn(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) BuyExactIn(opts *bind.TransactOpts, marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "buyExactIn", marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactIn is a paid mutator transaction binding the contract method 0xa12d9400.
//
// Solidity: function buyExactIn(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) BuyExactIn(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.BuyExactIn(&_SpotMarketOptimism.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactIn is a paid mutator transaction binding the contract method 0xa12d9400.
//
// Solidity: function buyExactIn(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) BuyExactIn(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.BuyExactIn(&_SpotMarketOptimism.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactOut is a paid mutator transaction binding the contract method 0x983220bb.
//
// Solidity: function buyExactOut(uint128 marketId, uint256 synthAmount, uint256 maxUsdAmount, address referrer) returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) BuyExactOut(opts *bind.TransactOpts, marketId *big.Int, synthAmount *big.Int, maxUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "buyExactOut", marketId, synthAmount, maxUsdAmount, referrer)
}

// BuyExactOut is a paid mutator transaction binding the contract method 0x983220bb.
//
// Solidity: function buyExactOut(uint128 marketId, uint256 synthAmount, uint256 maxUsdAmount, address referrer) returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) BuyExactOut(marketId *big.Int, synthAmount *big.Int, maxUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.BuyExactOut(&_SpotMarketOptimism.TransactOpts, marketId, synthAmount, maxUsdAmount, referrer)
}

// BuyExactOut is a paid mutator transaction binding the contract method 0x983220bb.
//
// Solidity: function buyExactOut(uint128 marketId, uint256 synthAmount, uint256 maxUsdAmount, address referrer) returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) BuyExactOut(marketId *big.Int, synthAmount *big.Int, maxUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.BuyExactOut(&_SpotMarketOptimism.TransactOpts, marketId, synthAmount, maxUsdAmount, referrer)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x4138dc53.
//
// Solidity: function cancelOrder(uint128 marketId, uint128 asyncOrderId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) CancelOrder(opts *bind.TransactOpts, marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "cancelOrder", marketId, asyncOrderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x4138dc53.
//
// Solidity: function cancelOrder(uint128 marketId, uint128 asyncOrderId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) CancelOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.CancelOrder(&_SpotMarketOptimism.TransactOpts, marketId, asyncOrderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x4138dc53.
//
// Solidity: function cancelOrder(uint128 marketId, uint128 asyncOrderId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) CancelOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.CancelOrder(&_SpotMarketOptimism.TransactOpts, marketId, asyncOrderId)
}

// CommitOrder is a paid mutator transaction binding the contract method 0xd393340e.
//
// Solidity: function commitOrder(uint128 marketId, uint8 orderType, uint256 amountProvided, uint256 settlementStrategyId, uint256 minimumSettlementAmount, address referrer) returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) CommitOrder(opts *bind.TransactOpts, marketId *big.Int, orderType uint8, amountProvided *big.Int, settlementStrategyId *big.Int, minimumSettlementAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "commitOrder", marketId, orderType, amountProvided, settlementStrategyId, minimumSettlementAmount, referrer)
}

// CommitOrder is a paid mutator transaction binding the contract method 0xd393340e.
//
// Solidity: function commitOrder(uint128 marketId, uint8 orderType, uint256 amountProvided, uint256 settlementStrategyId, uint256 minimumSettlementAmount, address referrer) returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketOptimism *SpotMarketOptimismSession) CommitOrder(marketId *big.Int, orderType uint8, amountProvided *big.Int, settlementStrategyId *big.Int, minimumSettlementAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.CommitOrder(&_SpotMarketOptimism.TransactOpts, marketId, orderType, amountProvided, settlementStrategyId, minimumSettlementAmount, referrer)
}

// CommitOrder is a paid mutator transaction binding the contract method 0xd393340e.
//
// Solidity: function commitOrder(uint128 marketId, uint8 orderType, uint256 amountProvided, uint256 settlementStrategyId, uint256 minimumSettlementAmount, address referrer) returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) CommitOrder(marketId *big.Int, orderType uint8, amountProvided *big.Int, settlementStrategyId *big.Int, minimumSettlementAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.CommitOrder(&_SpotMarketOptimism.TransactOpts, marketId, orderType, amountProvided, settlementStrategyId, minimumSettlementAmount, referrer)
}

// CreateSynth is a paid mutator transaction binding the contract method 0x2e535d61.
//
// Solidity: function createSynth(string tokenName, string tokenSymbol, address synthOwner) returns(uint128 synthMarketId)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) CreateSynth(opts *bind.TransactOpts, tokenName string, tokenSymbol string, synthOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "createSynth", tokenName, tokenSymbol, synthOwner)
}

// CreateSynth is a paid mutator transaction binding the contract method 0x2e535d61.
//
// Solidity: function createSynth(string tokenName, string tokenSymbol, address synthOwner) returns(uint128 synthMarketId)
func (_SpotMarketOptimism *SpotMarketOptimismSession) CreateSynth(tokenName string, tokenSymbol string, synthOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.CreateSynth(&_SpotMarketOptimism.TransactOpts, tokenName, tokenSymbol, synthOwner)
}

// CreateSynth is a paid mutator transaction binding the contract method 0x2e535d61.
//
// Solidity: function createSynth(string tokenName, string tokenSymbol, address synthOwner) returns(uint128 synthMarketId)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) CreateSynth(tokenName string, tokenSymbol string, synthOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.CreateSynth(&_SpotMarketOptimism.TransactOpts, tokenName, tokenSymbol, synthOwner)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) InitOrUpgradeNft(opts *bind.TransactOpts, id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "initOrUpgradeNft", id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.InitOrUpgradeNft(&_SpotMarketOptimism.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.InitOrUpgradeNft(&_SpotMarketOptimism.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) InitOrUpgradeToken(opts *bind.TransactOpts, id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "initOrUpgradeToken", id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.InitOrUpgradeToken(&_SpotMarketOptimism.TransactOpts, id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.InitOrUpgradeToken(&_SpotMarketOptimism.TransactOpts, id, name, symbol, decimals, impl)
}

// NominateMarketOwner is a paid mutator transaction binding the contract method 0x5950864b.
//
// Solidity: function nominateMarketOwner(uint128 synthMarketId, address newNominatedOwner) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) NominateMarketOwner(opts *bind.TransactOpts, synthMarketId *big.Int, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "nominateMarketOwner", synthMarketId, newNominatedOwner)
}

// NominateMarketOwner is a paid mutator transaction binding the contract method 0x5950864b.
//
// Solidity: function nominateMarketOwner(uint128 synthMarketId, address newNominatedOwner) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) NominateMarketOwner(synthMarketId *big.Int, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.NominateMarketOwner(&_SpotMarketOptimism.TransactOpts, synthMarketId, newNominatedOwner)
}

// NominateMarketOwner is a paid mutator transaction binding the contract method 0x5950864b.
//
// Solidity: function nominateMarketOwner(uint128 synthMarketId, address newNominatedOwner) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) NominateMarketOwner(synthMarketId *big.Int, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.NominateMarketOwner(&_SpotMarketOptimism.TransactOpts, synthMarketId, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) NominateNewOwner(opts *bind.TransactOpts, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "nominateNewOwner", newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.NominateNewOwner(&_SpotMarketOptimism.TransactOpts, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.NominateNewOwner(&_SpotMarketOptimism.TransactOpts, newNominatedOwner)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) RegisterUnmanagedSystem(opts *bind.TransactOpts, id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "registerUnmanagedSystem", id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.RegisterUnmanagedSystem(&_SpotMarketOptimism.TransactOpts, id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.RegisterUnmanagedSystem(&_SpotMarketOptimism.TransactOpts, id, endpoint)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) RemoveFromFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "removeFromFeatureFlagAllowlist", feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.RemoveFromFeatureFlagAllowlist(&_SpotMarketOptimism.TransactOpts, feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.RemoveFromFeatureFlagAllowlist(&_SpotMarketOptimism.TransactOpts, feature, account)
}

// RenounceMarketNomination is a paid mutator transaction binding the contract method 0x298b26bf.
//
// Solidity: function renounceMarketNomination(uint128 synthMarketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) RenounceMarketNomination(opts *bind.TransactOpts, synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "renounceMarketNomination", synthMarketId)
}

// RenounceMarketNomination is a paid mutator transaction binding the contract method 0x298b26bf.
//
// Solidity: function renounceMarketNomination(uint128 synthMarketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) RenounceMarketNomination(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.RenounceMarketNomination(&_SpotMarketOptimism.TransactOpts, synthMarketId)
}

// RenounceMarketNomination is a paid mutator transaction binding the contract method 0x298b26bf.
//
// Solidity: function renounceMarketNomination(uint128 synthMarketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) RenounceMarketNomination(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.RenounceMarketNomination(&_SpotMarketOptimism.TransactOpts, synthMarketId)
}

// RenounceMarketOwnership is a paid mutator transaction binding the contract method 0xbd1cdfb5.
//
// Solidity: function renounceMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) RenounceMarketOwnership(opts *bind.TransactOpts, synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "renounceMarketOwnership", synthMarketId)
}

// RenounceMarketOwnership is a paid mutator transaction binding the contract method 0xbd1cdfb5.
//
// Solidity: function renounceMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) RenounceMarketOwnership(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.RenounceMarketOwnership(&_SpotMarketOptimism.TransactOpts, synthMarketId)
}

// RenounceMarketOwnership is a paid mutator transaction binding the contract method 0xbd1cdfb5.
//
// Solidity: function renounceMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) RenounceMarketOwnership(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.RenounceMarketOwnership(&_SpotMarketOptimism.TransactOpts, synthMarketId)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) RenounceNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "renounceNomination")
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) RenounceNomination() (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.RenounceNomination(&_SpotMarketOptimism.TransactOpts)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) RenounceNomination() (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.RenounceNomination(&_SpotMarketOptimism.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0x4d4bfbd5.
//
// Solidity: function sell(uint128 marketId, uint256 synthAmount, uint256 minUsdAmount, address referrer) returns(uint256 usdAmountReceived, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) Sell(opts *bind.TransactOpts, marketId *big.Int, synthAmount *big.Int, minUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "sell", marketId, synthAmount, minUsdAmount, referrer)
}

// Sell is a paid mutator transaction binding the contract method 0x4d4bfbd5.
//
// Solidity: function sell(uint128 marketId, uint256 synthAmount, uint256 minUsdAmount, address referrer) returns(uint256 usdAmountReceived, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) Sell(marketId *big.Int, synthAmount *big.Int, minUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.Sell(&_SpotMarketOptimism.TransactOpts, marketId, synthAmount, minUsdAmount, referrer)
}

// Sell is a paid mutator transaction binding the contract method 0x4d4bfbd5.
//
// Solidity: function sell(uint128 marketId, uint256 synthAmount, uint256 minUsdAmount, address referrer) returns(uint256 usdAmountReceived, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) Sell(marketId *big.Int, synthAmount *big.Int, minUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.Sell(&_SpotMarketOptimism.TransactOpts, marketId, synthAmount, minUsdAmount, referrer)
}

// SellExactIn is a paid mutator transaction binding the contract method 0x3d1a60e4.
//
// Solidity: function sellExactIn(uint128 marketId, uint256 synthAmount, uint256 minAmountReceived, address referrer) returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SellExactIn(opts *bind.TransactOpts, marketId *big.Int, synthAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "sellExactIn", marketId, synthAmount, minAmountReceived, referrer)
}

// SellExactIn is a paid mutator transaction binding the contract method 0x3d1a60e4.
//
// Solidity: function sellExactIn(uint128 marketId, uint256 synthAmount, uint256 minAmountReceived, address referrer) returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) SellExactIn(marketId *big.Int, synthAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SellExactIn(&_SpotMarketOptimism.TransactOpts, marketId, synthAmount, minAmountReceived, referrer)
}

// SellExactIn is a paid mutator transaction binding the contract method 0x3d1a60e4.
//
// Solidity: function sellExactIn(uint128 marketId, uint256 synthAmount, uint256 minAmountReceived, address referrer) returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SellExactIn(marketId *big.Int, synthAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SellExactIn(&_SpotMarketOptimism.TransactOpts, marketId, synthAmount, minAmountReceived, referrer)
}

// SellExactOut is a paid mutator transaction binding the contract method 0x4ce94d9d.
//
// Solidity: function sellExactOut(uint128 marketId, uint256 usdAmount, uint256 maxSynthAmount, address referrer) returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SellExactOut(opts *bind.TransactOpts, marketId *big.Int, usdAmount *big.Int, maxSynthAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "sellExactOut", marketId, usdAmount, maxSynthAmount, referrer)
}

// SellExactOut is a paid mutator transaction binding the contract method 0x4ce94d9d.
//
// Solidity: function sellExactOut(uint128 marketId, uint256 usdAmount, uint256 maxSynthAmount, address referrer) returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) SellExactOut(marketId *big.Int, usdAmount *big.Int, maxSynthAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SellExactOut(&_SpotMarketOptimism.TransactOpts, marketId, usdAmount, maxSynthAmount, referrer)
}

// SellExactOut is a paid mutator transaction binding the contract method 0x4ce94d9d.
//
// Solidity: function sellExactOut(uint128 marketId, uint256 usdAmount, uint256 maxSynthAmount, address referrer) returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SellExactOut(marketId *big.Int, usdAmount *big.Int, maxSynthAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SellExactOut(&_SpotMarketOptimism.TransactOpts, marketId, usdAmount, maxSynthAmount, referrer)
}

// SetAsyncFixedFee is a paid mutator transaction binding the contract method 0x61dcca86.
//
// Solidity: function setAsyncFixedFee(uint128 synthMarketId, uint256 asyncFixedFee) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetAsyncFixedFee(opts *bind.TransactOpts, synthMarketId *big.Int, asyncFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setAsyncFixedFee", synthMarketId, asyncFixedFee)
}

// SetAsyncFixedFee is a paid mutator transaction binding the contract method 0x61dcca86.
//
// Solidity: function setAsyncFixedFee(uint128 synthMarketId, uint256 asyncFixedFee) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetAsyncFixedFee(synthMarketId *big.Int, asyncFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetAsyncFixedFee(&_SpotMarketOptimism.TransactOpts, synthMarketId, asyncFixedFee)
}

// SetAsyncFixedFee is a paid mutator transaction binding the contract method 0x61dcca86.
//
// Solidity: function setAsyncFixedFee(uint128 synthMarketId, uint256 asyncFixedFee) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetAsyncFixedFee(synthMarketId *big.Int, asyncFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetAsyncFixedFee(&_SpotMarketOptimism.TransactOpts, synthMarketId, asyncFixedFee)
}

// SetAtomicFixedFee is a paid mutator transaction binding the contract method 0x480be91f.
//
// Solidity: function setAtomicFixedFee(uint128 synthMarketId, uint256 atomicFixedFee) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetAtomicFixedFee(opts *bind.TransactOpts, synthMarketId *big.Int, atomicFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setAtomicFixedFee", synthMarketId, atomicFixedFee)
}

// SetAtomicFixedFee is a paid mutator transaction binding the contract method 0x480be91f.
//
// Solidity: function setAtomicFixedFee(uint128 synthMarketId, uint256 atomicFixedFee) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetAtomicFixedFee(synthMarketId *big.Int, atomicFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetAtomicFixedFee(&_SpotMarketOptimism.TransactOpts, synthMarketId, atomicFixedFee)
}

// SetAtomicFixedFee is a paid mutator transaction binding the contract method 0x480be91f.
//
// Solidity: function setAtomicFixedFee(uint128 synthMarketId, uint256 atomicFixedFee) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetAtomicFixedFee(synthMarketId *big.Int, atomicFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetAtomicFixedFee(&_SpotMarketOptimism.TransactOpts, synthMarketId, atomicFixedFee)
}

// SetCollateralLeverage is a paid mutator transaction binding the contract method 0x21f7f58f.
//
// Solidity: function setCollateralLeverage(uint128 synthMarketId, uint256 collateralLeverage) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetCollateralLeverage(opts *bind.TransactOpts, synthMarketId *big.Int, collateralLeverage *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setCollateralLeverage", synthMarketId, collateralLeverage)
}

// SetCollateralLeverage is a paid mutator transaction binding the contract method 0x21f7f58f.
//
// Solidity: function setCollateralLeverage(uint128 synthMarketId, uint256 collateralLeverage) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetCollateralLeverage(synthMarketId *big.Int, collateralLeverage *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetCollateralLeverage(&_SpotMarketOptimism.TransactOpts, synthMarketId, collateralLeverage)
}

// SetCollateralLeverage is a paid mutator transaction binding the contract method 0x21f7f58f.
//
// Solidity: function setCollateralLeverage(uint128 synthMarketId, uint256 collateralLeverage) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetCollateralLeverage(synthMarketId *big.Int, collateralLeverage *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetCollateralLeverage(&_SpotMarketOptimism.TransactOpts, synthMarketId, collateralLeverage)
}

// SetCustomTransactorFees is a paid mutator transaction binding the contract method 0x95fcd547.
//
// Solidity: function setCustomTransactorFees(uint128 synthMarketId, address transactor, uint256 fixedFeeAmount) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetCustomTransactorFees(opts *bind.TransactOpts, synthMarketId *big.Int, transactor common.Address, fixedFeeAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setCustomTransactorFees", synthMarketId, transactor, fixedFeeAmount)
}

// SetCustomTransactorFees is a paid mutator transaction binding the contract method 0x95fcd547.
//
// Solidity: function setCustomTransactorFees(uint128 synthMarketId, address transactor, uint256 fixedFeeAmount) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address, fixedFeeAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetCustomTransactorFees(&_SpotMarketOptimism.TransactOpts, synthMarketId, transactor, fixedFeeAmount)
}

// SetCustomTransactorFees is a paid mutator transaction binding the contract method 0x95fcd547.
//
// Solidity: function setCustomTransactorFees(uint128 synthMarketId, address transactor, uint256 fixedFeeAmount) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address, fixedFeeAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetCustomTransactorFees(&_SpotMarketOptimism.TransactOpts, synthMarketId, transactor, fixedFeeAmount)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xec846bac.
//
// Solidity: function setDecayRate(uint128 marketId, uint256 rate) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetDecayRate(opts *bind.TransactOpts, marketId *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setDecayRate", marketId, rate)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xec846bac.
//
// Solidity: function setDecayRate(uint128 marketId, uint256 rate) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetDecayRate(marketId *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetDecayRate(&_SpotMarketOptimism.TransactOpts, marketId, rate)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xec846bac.
//
// Solidity: function setDecayRate(uint128 marketId, uint256 rate) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetDecayRate(marketId *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetDecayRate(&_SpotMarketOptimism.TransactOpts, marketId, rate)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetDeniers(opts *bind.TransactOpts, feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setDeniers", feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetDeniers(&_SpotMarketOptimism.TransactOpts, feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetDeniers(&_SpotMarketOptimism.TransactOpts, feature, deniers)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetFeatureFlagAllowAll(opts *bind.TransactOpts, feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setFeatureFlagAllowAll", feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetFeatureFlagAllowAll(&_SpotMarketOptimism.TransactOpts, feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetFeatureFlagAllowAll(&_SpotMarketOptimism.TransactOpts, feature, allowAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetFeatureFlagDenyAll(opts *bind.TransactOpts, feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setFeatureFlagDenyAll", feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetFeatureFlagDenyAll(&_SpotMarketOptimism.TransactOpts, feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetFeatureFlagDenyAll(&_SpotMarketOptimism.TransactOpts, feature, denyAll)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x025f6120.
//
// Solidity: function setFeeCollector(uint128 synthMarketId, address feeCollector) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetFeeCollector(opts *bind.TransactOpts, synthMarketId *big.Int, feeCollector common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setFeeCollector", synthMarketId, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x025f6120.
//
// Solidity: function setFeeCollector(uint128 synthMarketId, address feeCollector) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetFeeCollector(synthMarketId *big.Int, feeCollector common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetFeeCollector(&_SpotMarketOptimism.TransactOpts, synthMarketId, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x025f6120.
//
// Solidity: function setFeeCollector(uint128 synthMarketId, address feeCollector) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetFeeCollector(synthMarketId *big.Int, feeCollector common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetFeeCollector(&_SpotMarketOptimism.TransactOpts, synthMarketId, feeCollector)
}

// SetMarketSkewScale is a paid mutator transaction binding the contract method 0x9a40f8cb.
//
// Solidity: function setMarketSkewScale(uint128 synthMarketId, uint256 skewScale) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetMarketSkewScale(opts *bind.TransactOpts, synthMarketId *big.Int, skewScale *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setMarketSkewScale", synthMarketId, skewScale)
}

// SetMarketSkewScale is a paid mutator transaction binding the contract method 0x9a40f8cb.
//
// Solidity: function setMarketSkewScale(uint128 synthMarketId, uint256 skewScale) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetMarketSkewScale(synthMarketId *big.Int, skewScale *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetMarketSkewScale(&_SpotMarketOptimism.TransactOpts, synthMarketId, skewScale)
}

// SetMarketSkewScale is a paid mutator transaction binding the contract method 0x9a40f8cb.
//
// Solidity: function setMarketSkewScale(uint128 synthMarketId, uint256 skewScale) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetMarketSkewScale(synthMarketId *big.Int, skewScale *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetMarketSkewScale(&_SpotMarketOptimism.TransactOpts, synthMarketId, skewScale)
}

// SetMarketUtilizationFees is a paid mutator transaction binding the contract method 0x45f2601c.
//
// Solidity: function setMarketUtilizationFees(uint128 synthMarketId, uint256 utilizationFeeRate) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetMarketUtilizationFees(opts *bind.TransactOpts, synthMarketId *big.Int, utilizationFeeRate *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setMarketUtilizationFees", synthMarketId, utilizationFeeRate)
}

// SetMarketUtilizationFees is a paid mutator transaction binding the contract method 0x45f2601c.
//
// Solidity: function setMarketUtilizationFees(uint128 synthMarketId, uint256 utilizationFeeRate) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetMarketUtilizationFees(synthMarketId *big.Int, utilizationFeeRate *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetMarketUtilizationFees(&_SpotMarketOptimism.TransactOpts, synthMarketId, utilizationFeeRate)
}

// SetMarketUtilizationFees is a paid mutator transaction binding the contract method 0x45f2601c.
//
// Solidity: function setMarketUtilizationFees(uint128 synthMarketId, uint256 utilizationFeeRate) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetMarketUtilizationFees(synthMarketId *big.Int, utilizationFeeRate *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetMarketUtilizationFees(&_SpotMarketOptimism.TransactOpts, synthMarketId, utilizationFeeRate)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetSettlementStrategyEnabled(opts *bind.TransactOpts, marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setSettlementStrategyEnabled", marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetSettlementStrategyEnabled(&_SpotMarketOptimism.TransactOpts, marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetSettlementStrategyEnabled(&_SpotMarketOptimism.TransactOpts, marketId, strategyId, enabled)
}

// SetSynthImplementation is a paid mutator transaction binding the contract method 0xec04ceb1.
//
// Solidity: function setSynthImplementation(address synthImplementation) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetSynthImplementation(opts *bind.TransactOpts, synthImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setSynthImplementation", synthImplementation)
}

// SetSynthImplementation is a paid mutator transaction binding the contract method 0xec04ceb1.
//
// Solidity: function setSynthImplementation(address synthImplementation) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetSynthImplementation(synthImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetSynthImplementation(&_SpotMarketOptimism.TransactOpts, synthImplementation)
}

// SetSynthImplementation is a paid mutator transaction binding the contract method 0xec04ceb1.
//
// Solidity: function setSynthImplementation(address synthImplementation) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetSynthImplementation(synthImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetSynthImplementation(&_SpotMarketOptimism.TransactOpts, synthImplementation)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetSynthetix(opts *bind.TransactOpts, synthetix common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setSynthetix", synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetSynthetix(&_SpotMarketOptimism.TransactOpts, synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetSynthetix(&_SpotMarketOptimism.TransactOpts, synthetix)
}

// SetWrapper is a paid mutator transaction binding the contract method 0x673a21e5.
//
// Solidity: function setWrapper(uint128 marketId, address wrapCollateralType, uint256 maxWrappableAmount) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetWrapper(opts *bind.TransactOpts, marketId *big.Int, wrapCollateralType common.Address, maxWrappableAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setWrapper", marketId, wrapCollateralType, maxWrappableAmount)
}

// SetWrapper is a paid mutator transaction binding the contract method 0x673a21e5.
//
// Solidity: function setWrapper(uint128 marketId, address wrapCollateralType, uint256 maxWrappableAmount) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetWrapper(marketId *big.Int, wrapCollateralType common.Address, maxWrappableAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetWrapper(&_SpotMarketOptimism.TransactOpts, marketId, wrapCollateralType, maxWrappableAmount)
}

// SetWrapper is a paid mutator transaction binding the contract method 0x673a21e5.
//
// Solidity: function setWrapper(uint128 marketId, address wrapCollateralType, uint256 maxWrappableAmount) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetWrapper(marketId *big.Int, wrapCollateralType common.Address, maxWrappableAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetWrapper(&_SpotMarketOptimism.TransactOpts, marketId, wrapCollateralType, maxWrappableAmount)
}

// SetWrapperFees is a paid mutator transaction binding the contract method 0x6539b1c3.
//
// Solidity: function setWrapperFees(uint128 synthMarketId, int256 wrapFee, int256 unwrapFee) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SetWrapperFees(opts *bind.TransactOpts, synthMarketId *big.Int, wrapFee *big.Int, unwrapFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "setWrapperFees", synthMarketId, wrapFee, unwrapFee)
}

// SetWrapperFees is a paid mutator transaction binding the contract method 0x6539b1c3.
//
// Solidity: function setWrapperFees(uint128 synthMarketId, int256 wrapFee, int256 unwrapFee) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SetWrapperFees(synthMarketId *big.Int, wrapFee *big.Int, unwrapFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetWrapperFees(&_SpotMarketOptimism.TransactOpts, synthMarketId, wrapFee, unwrapFee)
}

// SetWrapperFees is a paid mutator transaction binding the contract method 0x6539b1c3.
//
// Solidity: function setWrapperFees(uint128 synthMarketId, int256 wrapFee, int256 unwrapFee) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SetWrapperFees(synthMarketId *big.Int, wrapFee *big.Int, unwrapFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SetWrapperFees(&_SpotMarketOptimism.TransactOpts, synthMarketId, wrapFee, unwrapFee)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x9444ac48.
//
// Solidity: function settleOrder(uint128 marketId, uint128 asyncOrderId) returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SettleOrder(opts *bind.TransactOpts, marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "settleOrder", marketId, asyncOrderId)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x9444ac48.
//
// Solidity: function settleOrder(uint128 marketId, uint128 asyncOrderId) returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) SettleOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SettleOrder(&_SpotMarketOptimism.TransactOpts, marketId, asyncOrderId)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x9444ac48.
//
// Solidity: function settleOrder(uint128 marketId, uint128 asyncOrderId) returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SettleOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SettleOrder(&_SpotMarketOptimism.TransactOpts, marketId, asyncOrderId)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SettlePythOrder(opts *bind.TransactOpts, result []byte, extraData []byte) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "settlePythOrder", result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SettlePythOrder(&_SpotMarketOptimism.TransactOpts, result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SettlePythOrder(&_SpotMarketOptimism.TransactOpts, result, extraData)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) SimulateUpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "simulateUpgradeTo", newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SimulateUpgradeTo(&_SpotMarketOptimism.TransactOpts, newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.SimulateUpgradeTo(&_SpotMarketOptimism.TransactOpts, newImplementation)
}

// Unwrap is a paid mutator transaction binding the contract method 0x784dad9e.
//
// Solidity: function unwrap(uint128 marketId, uint256 unwrapAmount, uint256 minAmountReceived) returns(uint256 returnCollateralAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) Unwrap(opts *bind.TransactOpts, marketId *big.Int, unwrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "unwrap", marketId, unwrapAmount, minAmountReceived)
}

// Unwrap is a paid mutator transaction binding the contract method 0x784dad9e.
//
// Solidity: function unwrap(uint128 marketId, uint256 unwrapAmount, uint256 minAmountReceived) returns(uint256 returnCollateralAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) Unwrap(marketId *big.Int, unwrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.Unwrap(&_SpotMarketOptimism.TransactOpts, marketId, unwrapAmount, minAmountReceived)
}

// Unwrap is a paid mutator transaction binding the contract method 0x784dad9e.
//
// Solidity: function unwrap(uint128 marketId, uint256 unwrapAmount, uint256 minAmountReceived) returns(uint256 returnCollateralAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) Unwrap(marketId *big.Int, unwrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.Unwrap(&_SpotMarketOptimism.TransactOpts, marketId, unwrapAmount, minAmountReceived)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x6a0e2085.
//
// Solidity: function updatePriceData(uint128 synthMarketId, bytes32 buyFeedId, bytes32 sellFeedId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) UpdatePriceData(opts *bind.TransactOpts, synthMarketId *big.Int, buyFeedId [32]byte, sellFeedId [32]byte) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "updatePriceData", synthMarketId, buyFeedId, sellFeedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x6a0e2085.
//
// Solidity: function updatePriceData(uint128 synthMarketId, bytes32 buyFeedId, bytes32 sellFeedId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) UpdatePriceData(synthMarketId *big.Int, buyFeedId [32]byte, sellFeedId [32]byte) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.UpdatePriceData(&_SpotMarketOptimism.TransactOpts, synthMarketId, buyFeedId, sellFeedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x6a0e2085.
//
// Solidity: function updatePriceData(uint128 synthMarketId, bytes32 buyFeedId, bytes32 sellFeedId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) UpdatePriceData(synthMarketId *big.Int, buyFeedId [32]byte, sellFeedId [32]byte) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.UpdatePriceData(&_SpotMarketOptimism.TransactOpts, synthMarketId, buyFeedId, sellFeedId)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6ad77077.
//
// Solidity: function updateReferrerShare(uint128 synthMarketId, address referrer, uint256 sharePercentage) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) UpdateReferrerShare(opts *bind.TransactOpts, synthMarketId *big.Int, referrer common.Address, sharePercentage *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "updateReferrerShare", synthMarketId, referrer, sharePercentage)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6ad77077.
//
// Solidity: function updateReferrerShare(uint128 synthMarketId, address referrer, uint256 sharePercentage) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) UpdateReferrerShare(synthMarketId *big.Int, referrer common.Address, sharePercentage *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.UpdateReferrerShare(&_SpotMarketOptimism.TransactOpts, synthMarketId, referrer, sharePercentage)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6ad77077.
//
// Solidity: function updateReferrerShare(uint128 synthMarketId, address referrer, uint256 sharePercentage) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) UpdateReferrerShare(synthMarketId *big.Int, referrer common.Address, sharePercentage *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.UpdateReferrerShare(&_SpotMarketOptimism.TransactOpts, synthMarketId, referrer, sharePercentage)
}

// UpgradeSynthImpl is a paid mutator transaction binding the contract method 0xc99d0cdd.
//
// Solidity: function upgradeSynthImpl(uint128 marketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) UpgradeSynthImpl(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "upgradeSynthImpl", marketId)
}

// UpgradeSynthImpl is a paid mutator transaction binding the contract method 0xc99d0cdd.
//
// Solidity: function upgradeSynthImpl(uint128 marketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) UpgradeSynthImpl(marketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.UpgradeSynthImpl(&_SpotMarketOptimism.TransactOpts, marketId)
}

// UpgradeSynthImpl is a paid mutator transaction binding the contract method 0xc99d0cdd.
//
// Solidity: function upgradeSynthImpl(uint128 marketId) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) UpgradeSynthImpl(marketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.UpgradeSynthImpl(&_SpotMarketOptimism.TransactOpts, marketId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SpotMarketOptimism *SpotMarketOptimismSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.UpgradeTo(&_SpotMarketOptimism.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.UpgradeTo(&_SpotMarketOptimism.TransactOpts, newImplementation)
}

// Wrap is a paid mutator transaction binding the contract method 0xd7ce770c.
//
// Solidity: function wrap(uint128 marketId, uint256 wrapAmount, uint256 minAmountReceived) returns(uint256 amountToMint, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactor) Wrap(opts *bind.TransactOpts, marketId *big.Int, wrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.contract.Transact(opts, "wrap", marketId, wrapAmount, minAmountReceived)
}

// Wrap is a paid mutator transaction binding the contract method 0xd7ce770c.
//
// Solidity: function wrap(uint128 marketId, uint256 wrapAmount, uint256 minAmountReceived) returns(uint256 amountToMint, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismSession) Wrap(marketId *big.Int, wrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.Wrap(&_SpotMarketOptimism.TransactOpts, marketId, wrapAmount, minAmountReceived)
}

// Wrap is a paid mutator transaction binding the contract method 0xd7ce770c.
//
// Solidity: function wrap(uint128 marketId, uint256 wrapAmount, uint256 minAmountReceived) returns(uint256 amountToMint, (uint256,uint256,int256,int256) fees)
func (_SpotMarketOptimism *SpotMarketOptimismTransactorSession) Wrap(marketId *big.Int, wrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketOptimism.Contract.Wrap(&_SpotMarketOptimism.TransactOpts, marketId, wrapAmount, minAmountReceived)
}

// SpotMarketOptimismAssociatedSystemSetIterator is returned from FilterAssociatedSystemSet and is used to iterate over the raw logs and unpacked data for AssociatedSystemSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismAssociatedSystemSetIterator struct {
	Event *SpotMarketOptimismAssociatedSystemSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismAssociatedSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismAssociatedSystemSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismAssociatedSystemSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismAssociatedSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismAssociatedSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismAssociatedSystemSet represents a AssociatedSystemSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismAssociatedSystemSet struct {
	Kind  [32]byte
	Id    [32]byte
	Proxy common.Address
	Impl  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssociatedSystemSet is a free log retrieval operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterAssociatedSystemSet(opts *bind.FilterOpts, kind [][32]byte, id [][32]byte) (*SpotMarketOptimismAssociatedSystemSetIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismAssociatedSystemSetIterator{contract: _SpotMarketOptimism.contract, event: "AssociatedSystemSet", logs: logs, sub: sub}, nil
}

// WatchAssociatedSystemSet is a free log subscription operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchAssociatedSystemSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismAssociatedSystemSet, kind [][32]byte, id [][32]byte) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismAssociatedSystemSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssociatedSystemSet is a log parse operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseAssociatedSystemSet(log types.Log) (*SpotMarketOptimismAssociatedSystemSet, error) {
	event := new(SpotMarketOptimismAssociatedSystemSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismAsyncFixedFeeSetIterator is returned from FilterAsyncFixedFeeSet and is used to iterate over the raw logs and unpacked data for AsyncFixedFeeSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismAsyncFixedFeeSetIterator struct {
	Event *SpotMarketOptimismAsyncFixedFeeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismAsyncFixedFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismAsyncFixedFeeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismAsyncFixedFeeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismAsyncFixedFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismAsyncFixedFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismAsyncFixedFeeSet represents a AsyncFixedFeeSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismAsyncFixedFeeSet struct {
	SynthMarketId *big.Int
	AsyncFixedFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAsyncFixedFeeSet is a free log retrieval operation binding the contract event 0xaef59d229d195c5b8372221acbf4041b926fb1616a95f93e44379e4f30d57bfe.
//
// Solidity: event AsyncFixedFeeSet(uint256 indexed synthMarketId, uint256 asyncFixedFee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterAsyncFixedFeeSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismAsyncFixedFeeSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "AsyncFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismAsyncFixedFeeSetIterator{contract: _SpotMarketOptimism.contract, event: "AsyncFixedFeeSet", logs: logs, sub: sub}, nil
}

// WatchAsyncFixedFeeSet is a free log subscription operation binding the contract event 0xaef59d229d195c5b8372221acbf4041b926fb1616a95f93e44379e4f30d57bfe.
//
// Solidity: event AsyncFixedFeeSet(uint256 indexed synthMarketId, uint256 asyncFixedFee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchAsyncFixedFeeSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismAsyncFixedFeeSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "AsyncFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismAsyncFixedFeeSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "AsyncFixedFeeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAsyncFixedFeeSet is a log parse operation binding the contract event 0xaef59d229d195c5b8372221acbf4041b926fb1616a95f93e44379e4f30d57bfe.
//
// Solidity: event AsyncFixedFeeSet(uint256 indexed synthMarketId, uint256 asyncFixedFee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseAsyncFixedFeeSet(log types.Log) (*SpotMarketOptimismAsyncFixedFeeSet, error) {
	event := new(SpotMarketOptimismAsyncFixedFeeSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "AsyncFixedFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismAtomicFixedFeeSetIterator is returned from FilterAtomicFixedFeeSet and is used to iterate over the raw logs and unpacked data for AtomicFixedFeeSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismAtomicFixedFeeSetIterator struct {
	Event *SpotMarketOptimismAtomicFixedFeeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismAtomicFixedFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismAtomicFixedFeeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismAtomicFixedFeeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismAtomicFixedFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismAtomicFixedFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismAtomicFixedFeeSet represents a AtomicFixedFeeSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismAtomicFixedFeeSet struct {
	SynthMarketId  *big.Int
	AtomicFixedFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAtomicFixedFeeSet is a free log retrieval operation binding the contract event 0x6b0a526b06b2f30ba2d5b063c2ef81547512216d37c540b86039e3a19f1d2b3f.
//
// Solidity: event AtomicFixedFeeSet(uint256 indexed synthMarketId, uint256 atomicFixedFee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterAtomicFixedFeeSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismAtomicFixedFeeSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "AtomicFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismAtomicFixedFeeSetIterator{contract: _SpotMarketOptimism.contract, event: "AtomicFixedFeeSet", logs: logs, sub: sub}, nil
}

// WatchAtomicFixedFeeSet is a free log subscription operation binding the contract event 0x6b0a526b06b2f30ba2d5b063c2ef81547512216d37c540b86039e3a19f1d2b3f.
//
// Solidity: event AtomicFixedFeeSet(uint256 indexed synthMarketId, uint256 atomicFixedFee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchAtomicFixedFeeSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismAtomicFixedFeeSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "AtomicFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismAtomicFixedFeeSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "AtomicFixedFeeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAtomicFixedFeeSet is a log parse operation binding the contract event 0x6b0a526b06b2f30ba2d5b063c2ef81547512216d37c540b86039e3a19f1d2b3f.
//
// Solidity: event AtomicFixedFeeSet(uint256 indexed synthMarketId, uint256 atomicFixedFee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseAtomicFixedFeeSet(log types.Log) (*SpotMarketOptimismAtomicFixedFeeSet, error) {
	event := new(SpotMarketOptimismAtomicFixedFeeSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "AtomicFixedFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismCollateralLeverageSetIterator is returned from FilterCollateralLeverageSet and is used to iterate over the raw logs and unpacked data for CollateralLeverageSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismCollateralLeverageSetIterator struct {
	Event *SpotMarketOptimismCollateralLeverageSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismCollateralLeverageSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismCollateralLeverageSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismCollateralLeverageSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismCollateralLeverageSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismCollateralLeverageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismCollateralLeverageSet represents a CollateralLeverageSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismCollateralLeverageSet struct {
	SynthMarketId      *big.Int
	CollateralLeverage *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCollateralLeverageSet is a free log retrieval operation binding the contract event 0x83113c0e8d811f7e7017948357e945a1a8a6b6fc0d76c8512ffdd6f6766e8a13.
//
// Solidity: event CollateralLeverageSet(uint256 indexed synthMarketId, uint256 collateralLeverage)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterCollateralLeverageSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismCollateralLeverageSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "CollateralLeverageSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismCollateralLeverageSetIterator{contract: _SpotMarketOptimism.contract, event: "CollateralLeverageSet", logs: logs, sub: sub}, nil
}

// WatchCollateralLeverageSet is a free log subscription operation binding the contract event 0x83113c0e8d811f7e7017948357e945a1a8a6b6fc0d76c8512ffdd6f6766e8a13.
//
// Solidity: event CollateralLeverageSet(uint256 indexed synthMarketId, uint256 collateralLeverage)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchCollateralLeverageSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismCollateralLeverageSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "CollateralLeverageSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismCollateralLeverageSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "CollateralLeverageSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralLeverageSet is a log parse operation binding the contract event 0x83113c0e8d811f7e7017948357e945a1a8a6b6fc0d76c8512ffdd6f6766e8a13.
//
// Solidity: event CollateralLeverageSet(uint256 indexed synthMarketId, uint256 collateralLeverage)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseCollateralLeverageSet(log types.Log) (*SpotMarketOptimismCollateralLeverageSet, error) {
	event := new(SpotMarketOptimismCollateralLeverageSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "CollateralLeverageSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismDecayRateUpdatedIterator is returned from FilterDecayRateUpdated and is used to iterate over the raw logs and unpacked data for DecayRateUpdated events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismDecayRateUpdatedIterator struct {
	Event *SpotMarketOptimismDecayRateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismDecayRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismDecayRateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismDecayRateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismDecayRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismDecayRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismDecayRateUpdated represents a DecayRateUpdated event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismDecayRateUpdated struct {
	MarketId *big.Int
	Rate     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDecayRateUpdated is a free log retrieval operation binding the contract event 0x8a1ed33cd71d1533eadd1d4eb0ea2ae64da7d343cb2a932fdf135f345264e2b5.
//
// Solidity: event DecayRateUpdated(uint128 indexed marketId, uint256 rate)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterDecayRateUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*SpotMarketOptimismDecayRateUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "DecayRateUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismDecayRateUpdatedIterator{contract: _SpotMarketOptimism.contract, event: "DecayRateUpdated", logs: logs, sub: sub}, nil
}

// WatchDecayRateUpdated is a free log subscription operation binding the contract event 0x8a1ed33cd71d1533eadd1d4eb0ea2ae64da7d343cb2a932fdf135f345264e2b5.
//
// Solidity: event DecayRateUpdated(uint128 indexed marketId, uint256 rate)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchDecayRateUpdated(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismDecayRateUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "DecayRateUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismDecayRateUpdated)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "DecayRateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecayRateUpdated is a log parse operation binding the contract event 0x8a1ed33cd71d1533eadd1d4eb0ea2ae64da7d343cb2a932fdf135f345264e2b5.
//
// Solidity: event DecayRateUpdated(uint128 indexed marketId, uint256 rate)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseDecayRateUpdated(log types.Log) (*SpotMarketOptimismDecayRateUpdated, error) {
	event := new(SpotMarketOptimismDecayRateUpdated)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "DecayRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismFeatureFlagAllowAllSetIterator is returned from FilterFeatureFlagAllowAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowAllSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeatureFlagAllowAllSetIterator struct {
	Event *SpotMarketOptimismFeatureFlagAllowAllSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismFeatureFlagAllowAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismFeatureFlagAllowAllSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismFeatureFlagAllowAllSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismFeatureFlagAllowAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismFeatureFlagAllowAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismFeatureFlagAllowAllSet represents a FeatureFlagAllowAllSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeatureFlagAllowAllSet struct {
	Feature  [32]byte
	AllowAll bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowAllSet is a free log retrieval operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterFeatureFlagAllowAllSet(opts *bind.FilterOpts, feature [][32]byte) (*SpotMarketOptimismFeatureFlagAllowAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismFeatureFlagAllowAllSetIterator{contract: _SpotMarketOptimism.contract, event: "FeatureFlagAllowAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowAllSet is a free log subscription operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchFeatureFlagAllowAllSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismFeatureFlagAllowAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismFeatureFlagAllowAllSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeatureFlagAllowAllSet is a log parse operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseFeatureFlagAllowAllSet(log types.Log) (*SpotMarketOptimismFeatureFlagAllowAllSet, error) {
	event := new(SpotMarketOptimismFeatureFlagAllowAllSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismFeatureFlagAllowlistAddedIterator is returned from FilterFeatureFlagAllowlistAdded and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistAdded events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeatureFlagAllowlistAddedIterator struct {
	Event *SpotMarketOptimismFeatureFlagAllowlistAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismFeatureFlagAllowlistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismFeatureFlagAllowlistAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismFeatureFlagAllowlistAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismFeatureFlagAllowlistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismFeatureFlagAllowlistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismFeatureFlagAllowlistAdded represents a FeatureFlagAllowlistAdded event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeatureFlagAllowlistAdded struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistAdded is a free log retrieval operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterFeatureFlagAllowlistAdded(opts *bind.FilterOpts, feature [][32]byte) (*SpotMarketOptimismFeatureFlagAllowlistAddedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismFeatureFlagAllowlistAddedIterator{contract: _SpotMarketOptimism.contract, event: "FeatureFlagAllowlistAdded", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistAdded is a free log subscription operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchFeatureFlagAllowlistAdded(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismFeatureFlagAllowlistAdded, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismFeatureFlagAllowlistAdded)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeatureFlagAllowlistAdded is a log parse operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseFeatureFlagAllowlistAdded(log types.Log) (*SpotMarketOptimismFeatureFlagAllowlistAdded, error) {
	event := new(SpotMarketOptimismFeatureFlagAllowlistAdded)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismFeatureFlagAllowlistRemovedIterator is returned from FilterFeatureFlagAllowlistRemoved and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistRemoved events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeatureFlagAllowlistRemovedIterator struct {
	Event *SpotMarketOptimismFeatureFlagAllowlistRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismFeatureFlagAllowlistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismFeatureFlagAllowlistRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismFeatureFlagAllowlistRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismFeatureFlagAllowlistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismFeatureFlagAllowlistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismFeatureFlagAllowlistRemoved represents a FeatureFlagAllowlistRemoved event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeatureFlagAllowlistRemoved struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistRemoved is a free log retrieval operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterFeatureFlagAllowlistRemoved(opts *bind.FilterOpts, feature [][32]byte) (*SpotMarketOptimismFeatureFlagAllowlistRemovedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismFeatureFlagAllowlistRemovedIterator{contract: _SpotMarketOptimism.contract, event: "FeatureFlagAllowlistRemoved", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistRemoved is a free log subscription operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchFeatureFlagAllowlistRemoved(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismFeatureFlagAllowlistRemoved, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismFeatureFlagAllowlistRemoved)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeatureFlagAllowlistRemoved is a log parse operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseFeatureFlagAllowlistRemoved(log types.Log) (*SpotMarketOptimismFeatureFlagAllowlistRemoved, error) {
	event := new(SpotMarketOptimismFeatureFlagAllowlistRemoved)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismFeatureFlagDeniersResetIterator is returned from FilterFeatureFlagDeniersReset and is used to iterate over the raw logs and unpacked data for FeatureFlagDeniersReset events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeatureFlagDeniersResetIterator struct {
	Event *SpotMarketOptimismFeatureFlagDeniersReset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismFeatureFlagDeniersResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismFeatureFlagDeniersReset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismFeatureFlagDeniersReset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismFeatureFlagDeniersResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismFeatureFlagDeniersResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismFeatureFlagDeniersReset represents a FeatureFlagDeniersReset event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeatureFlagDeniersReset struct {
	Feature [32]byte
	Deniers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDeniersReset is a free log retrieval operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterFeatureFlagDeniersReset(opts *bind.FilterOpts, feature [][32]byte) (*SpotMarketOptimismFeatureFlagDeniersResetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismFeatureFlagDeniersResetIterator{contract: _SpotMarketOptimism.contract, event: "FeatureFlagDeniersReset", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDeniersReset is a free log subscription operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchFeatureFlagDeniersReset(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismFeatureFlagDeniersReset, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismFeatureFlagDeniersReset)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeatureFlagDeniersReset is a log parse operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseFeatureFlagDeniersReset(log types.Log) (*SpotMarketOptimismFeatureFlagDeniersReset, error) {
	event := new(SpotMarketOptimismFeatureFlagDeniersReset)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismFeatureFlagDenyAllSetIterator is returned from FilterFeatureFlagDenyAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagDenyAllSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeatureFlagDenyAllSetIterator struct {
	Event *SpotMarketOptimismFeatureFlagDenyAllSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismFeatureFlagDenyAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismFeatureFlagDenyAllSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismFeatureFlagDenyAllSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismFeatureFlagDenyAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismFeatureFlagDenyAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismFeatureFlagDenyAllSet represents a FeatureFlagDenyAllSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeatureFlagDenyAllSet struct {
	Feature [32]byte
	DenyAll bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDenyAllSet is a free log retrieval operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterFeatureFlagDenyAllSet(opts *bind.FilterOpts, feature [][32]byte) (*SpotMarketOptimismFeatureFlagDenyAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismFeatureFlagDenyAllSetIterator{contract: _SpotMarketOptimism.contract, event: "FeatureFlagDenyAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDenyAllSet is a free log subscription operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchFeatureFlagDenyAllSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismFeatureFlagDenyAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismFeatureFlagDenyAllSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeatureFlagDenyAllSet is a log parse operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseFeatureFlagDenyAllSet(log types.Log) (*SpotMarketOptimismFeatureFlagDenyAllSet, error) {
	event := new(SpotMarketOptimismFeatureFlagDenyAllSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismFeeCollectorSetIterator is returned from FilterFeeCollectorSet and is used to iterate over the raw logs and unpacked data for FeeCollectorSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeeCollectorSetIterator struct {
	Event *SpotMarketOptimismFeeCollectorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismFeeCollectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismFeeCollectorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismFeeCollectorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismFeeCollectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismFeeCollectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismFeeCollectorSet represents a FeeCollectorSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismFeeCollectorSet struct {
	SynthMarketId *big.Int
	FeeCollector  common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorSet is a free log retrieval operation binding the contract event 0x9559a9b7e14bf53553c17859be245a108350185ec859ec690012b13b820b7ef4.
//
// Solidity: event FeeCollectorSet(uint256 indexed synthMarketId, address feeCollector)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterFeeCollectorSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismFeeCollectorSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "FeeCollectorSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismFeeCollectorSetIterator{contract: _SpotMarketOptimism.contract, event: "FeeCollectorSet", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorSet is a free log subscription operation binding the contract event 0x9559a9b7e14bf53553c17859be245a108350185ec859ec690012b13b820b7ef4.
//
// Solidity: event FeeCollectorSet(uint256 indexed synthMarketId, address feeCollector)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchFeeCollectorSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismFeeCollectorSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "FeeCollectorSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismFeeCollectorSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeCollectorSet is a log parse operation binding the contract event 0x9559a9b7e14bf53553c17859be245a108350185ec859ec690012b13b820b7ef4.
//
// Solidity: event FeeCollectorSet(uint256 indexed synthMarketId, address feeCollector)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseFeeCollectorSet(log types.Log) (*SpotMarketOptimismFeeCollectorSet, error) {
	event := new(SpotMarketOptimismFeeCollectorSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismMarketNominationRenouncedIterator is returned from FilterMarketNominationRenounced and is used to iterate over the raw logs and unpacked data for MarketNominationRenounced events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismMarketNominationRenouncedIterator struct {
	Event *SpotMarketOptimismMarketNominationRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismMarketNominationRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismMarketNominationRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismMarketNominationRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismMarketNominationRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismMarketNominationRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismMarketNominationRenounced represents a MarketNominationRenounced event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismMarketNominationRenounced struct {
	MarketId *big.Int
	Nominee  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketNominationRenounced is a free log retrieval operation binding the contract event 0xf5b87e3c7e0caa8e0d233591fff85e764ebc73d5c7027cce729fd4beab04c2b6.
//
// Solidity: event MarketNominationRenounced(uint128 indexed marketId, address nominee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterMarketNominationRenounced(opts *bind.FilterOpts, marketId []*big.Int) (*SpotMarketOptimismMarketNominationRenouncedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "MarketNominationRenounced", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismMarketNominationRenouncedIterator{contract: _SpotMarketOptimism.contract, event: "MarketNominationRenounced", logs: logs, sub: sub}, nil
}

// WatchMarketNominationRenounced is a free log subscription operation binding the contract event 0xf5b87e3c7e0caa8e0d233591fff85e764ebc73d5c7027cce729fd4beab04c2b6.
//
// Solidity: event MarketNominationRenounced(uint128 indexed marketId, address nominee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchMarketNominationRenounced(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismMarketNominationRenounced, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "MarketNominationRenounced", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismMarketNominationRenounced)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "MarketNominationRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketNominationRenounced is a log parse operation binding the contract event 0xf5b87e3c7e0caa8e0d233591fff85e764ebc73d5c7027cce729fd4beab04c2b6.
//
// Solidity: event MarketNominationRenounced(uint128 indexed marketId, address nominee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseMarketNominationRenounced(log types.Log) (*SpotMarketOptimismMarketNominationRenounced, error) {
	event := new(SpotMarketOptimismMarketNominationRenounced)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "MarketNominationRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismMarketOwnerChangedIterator is returned from FilterMarketOwnerChanged and is used to iterate over the raw logs and unpacked data for MarketOwnerChanged events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismMarketOwnerChangedIterator struct {
	Event *SpotMarketOptimismMarketOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismMarketOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismMarketOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismMarketOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismMarketOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismMarketOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismMarketOwnerChanged represents a MarketOwnerChanged event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismMarketOwnerChanged struct {
	MarketId *big.Int
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketOwnerChanged is a free log retrieval operation binding the contract event 0xe73c996387b656d1e0ea2866c854ed68122ce4e4eea51d6af012938b3c7ff52f.
//
// Solidity: event MarketOwnerChanged(uint128 indexed marketId, address oldOwner, address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterMarketOwnerChanged(opts *bind.FilterOpts, marketId []*big.Int) (*SpotMarketOptimismMarketOwnerChangedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "MarketOwnerChanged", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismMarketOwnerChangedIterator{contract: _SpotMarketOptimism.contract, event: "MarketOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchMarketOwnerChanged is a free log subscription operation binding the contract event 0xe73c996387b656d1e0ea2866c854ed68122ce4e4eea51d6af012938b3c7ff52f.
//
// Solidity: event MarketOwnerChanged(uint128 indexed marketId, address oldOwner, address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchMarketOwnerChanged(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismMarketOwnerChanged, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "MarketOwnerChanged", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismMarketOwnerChanged)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "MarketOwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketOwnerChanged is a log parse operation binding the contract event 0xe73c996387b656d1e0ea2866c854ed68122ce4e4eea51d6af012938b3c7ff52f.
//
// Solidity: event MarketOwnerChanged(uint128 indexed marketId, address oldOwner, address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseMarketOwnerChanged(log types.Log) (*SpotMarketOptimismMarketOwnerChanged, error) {
	event := new(SpotMarketOptimismMarketOwnerChanged)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "MarketOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismMarketOwnerNominatedIterator is returned from FilterMarketOwnerNominated and is used to iterate over the raw logs and unpacked data for MarketOwnerNominated events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismMarketOwnerNominatedIterator struct {
	Event *SpotMarketOptimismMarketOwnerNominated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismMarketOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismMarketOwnerNominated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismMarketOwnerNominated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismMarketOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismMarketOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismMarketOwnerNominated represents a MarketOwnerNominated event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismMarketOwnerNominated struct {
	MarketId *big.Int
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketOwnerNominated is a free log retrieval operation binding the contract event 0x54605d90ee82d9b4318b0b4b479f3966976e44b94c6aff221c04f5294f85016b.
//
// Solidity: event MarketOwnerNominated(uint128 indexed marketId, address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterMarketOwnerNominated(opts *bind.FilterOpts, marketId []*big.Int) (*SpotMarketOptimismMarketOwnerNominatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "MarketOwnerNominated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismMarketOwnerNominatedIterator{contract: _SpotMarketOptimism.contract, event: "MarketOwnerNominated", logs: logs, sub: sub}, nil
}

// WatchMarketOwnerNominated is a free log subscription operation binding the contract event 0x54605d90ee82d9b4318b0b4b479f3966976e44b94c6aff221c04f5294f85016b.
//
// Solidity: event MarketOwnerNominated(uint128 indexed marketId, address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchMarketOwnerNominated(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismMarketOwnerNominated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "MarketOwnerNominated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismMarketOwnerNominated)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "MarketOwnerNominated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketOwnerNominated is a log parse operation binding the contract event 0x54605d90ee82d9b4318b0b4b479f3966976e44b94c6aff221c04f5294f85016b.
//
// Solidity: event MarketOwnerNominated(uint128 indexed marketId, address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseMarketOwnerNominated(log types.Log) (*SpotMarketOptimismMarketOwnerNominated, error) {
	event := new(SpotMarketOptimismMarketOwnerNominated)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "MarketOwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismMarketSkewScaleSetIterator is returned from FilterMarketSkewScaleSet and is used to iterate over the raw logs and unpacked data for MarketSkewScaleSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismMarketSkewScaleSetIterator struct {
	Event *SpotMarketOptimismMarketSkewScaleSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismMarketSkewScaleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismMarketSkewScaleSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismMarketSkewScaleSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismMarketSkewScaleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismMarketSkewScaleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismMarketSkewScaleSet represents a MarketSkewScaleSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismMarketSkewScaleSet struct {
	SynthMarketId *big.Int
	SkewScale     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarketSkewScaleSet is a free log retrieval operation binding the contract event 0x786fdfd5a0e146d8f4878876a8439ff01436e5969b14682e12d07d7e926b157c.
//
// Solidity: event MarketSkewScaleSet(uint256 indexed synthMarketId, uint256 skewScale)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterMarketSkewScaleSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismMarketSkewScaleSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "MarketSkewScaleSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismMarketSkewScaleSetIterator{contract: _SpotMarketOptimism.contract, event: "MarketSkewScaleSet", logs: logs, sub: sub}, nil
}

// WatchMarketSkewScaleSet is a free log subscription operation binding the contract event 0x786fdfd5a0e146d8f4878876a8439ff01436e5969b14682e12d07d7e926b157c.
//
// Solidity: event MarketSkewScaleSet(uint256 indexed synthMarketId, uint256 skewScale)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchMarketSkewScaleSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismMarketSkewScaleSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "MarketSkewScaleSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismMarketSkewScaleSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "MarketSkewScaleSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketSkewScaleSet is a log parse operation binding the contract event 0x786fdfd5a0e146d8f4878876a8439ff01436e5969b14682e12d07d7e926b157c.
//
// Solidity: event MarketSkewScaleSet(uint256 indexed synthMarketId, uint256 skewScale)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseMarketSkewScaleSet(log types.Log) (*SpotMarketOptimismMarketSkewScaleSet, error) {
	event := new(SpotMarketOptimismMarketSkewScaleSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "MarketSkewScaleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismMarketUtilizationFeesSetIterator is returned from FilterMarketUtilizationFeesSet and is used to iterate over the raw logs and unpacked data for MarketUtilizationFeesSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismMarketUtilizationFeesSetIterator struct {
	Event *SpotMarketOptimismMarketUtilizationFeesSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismMarketUtilizationFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismMarketUtilizationFeesSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismMarketUtilizationFeesSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismMarketUtilizationFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismMarketUtilizationFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismMarketUtilizationFeesSet represents a MarketUtilizationFeesSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismMarketUtilizationFeesSet struct {
	SynthMarketId      *big.Int
	UtilizationFeeRate *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketUtilizationFeesSet is a free log retrieval operation binding the contract event 0x83dfad56ac61e49feb43345b8c73b6d45eb121decc66b1709ca0413b31c64f63.
//
// Solidity: event MarketUtilizationFeesSet(uint256 indexed synthMarketId, uint256 utilizationFeeRate)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterMarketUtilizationFeesSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismMarketUtilizationFeesSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "MarketUtilizationFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismMarketUtilizationFeesSetIterator{contract: _SpotMarketOptimism.contract, event: "MarketUtilizationFeesSet", logs: logs, sub: sub}, nil
}

// WatchMarketUtilizationFeesSet is a free log subscription operation binding the contract event 0x83dfad56ac61e49feb43345b8c73b6d45eb121decc66b1709ca0413b31c64f63.
//
// Solidity: event MarketUtilizationFeesSet(uint256 indexed synthMarketId, uint256 utilizationFeeRate)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchMarketUtilizationFeesSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismMarketUtilizationFeesSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "MarketUtilizationFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismMarketUtilizationFeesSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "MarketUtilizationFeesSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketUtilizationFeesSet is a log parse operation binding the contract event 0x83dfad56ac61e49feb43345b8c73b6d45eb121decc66b1709ca0413b31c64f63.
//
// Solidity: event MarketUtilizationFeesSet(uint256 indexed synthMarketId, uint256 utilizationFeeRate)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseMarketUtilizationFeesSet(log types.Log) (*SpotMarketOptimismMarketUtilizationFeesSet, error) {
	event := new(SpotMarketOptimismMarketUtilizationFeesSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "MarketUtilizationFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismOrderCancelledIterator struct {
	Event *SpotMarketOptimismOrderCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismOrderCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismOrderCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismOrderCancelled represents a OrderCancelled event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismOrderCancelled struct {
	MarketId        *big.Int
	AsyncOrderId    *big.Int
	AsyncOrderClaim AsyncOrderClaimData
	Sender          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xa57ffc5057d10e88a0dd7713d14130f9638d680af94401e7f1f4fba44a3ad5e2.
//
// Solidity: event OrderCancelled(uint128 indexed marketId, uint128 indexed asyncOrderId, (uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim, address indexed sender)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterOrderCancelled(opts *bind.FilterOpts, marketId []*big.Int, asyncOrderId []*big.Int, sender []common.Address) (*SpotMarketOptimismOrderCancelledIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var asyncOrderIdRule []interface{}
	for _, asyncOrderIdItem := range asyncOrderId {
		asyncOrderIdRule = append(asyncOrderIdRule, asyncOrderIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "OrderCancelled", marketIdRule, asyncOrderIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismOrderCancelledIterator{contract: _SpotMarketOptimism.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xa57ffc5057d10e88a0dd7713d14130f9638d680af94401e7f1f4fba44a3ad5e2.
//
// Solidity: event OrderCancelled(uint128 indexed marketId, uint128 indexed asyncOrderId, (uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim, address indexed sender)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismOrderCancelled, marketId []*big.Int, asyncOrderId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var asyncOrderIdRule []interface{}
	for _, asyncOrderIdItem := range asyncOrderId {
		asyncOrderIdRule = append(asyncOrderIdRule, asyncOrderIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "OrderCancelled", marketIdRule, asyncOrderIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismOrderCancelled)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderCancelled is a log parse operation binding the contract event 0xa57ffc5057d10e88a0dd7713d14130f9638d680af94401e7f1f4fba44a3ad5e2.
//
// Solidity: event OrderCancelled(uint128 indexed marketId, uint128 indexed asyncOrderId, (uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim, address indexed sender)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseOrderCancelled(log types.Log) (*SpotMarketOptimismOrderCancelled, error) {
	event := new(SpotMarketOptimismOrderCancelled)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismOrderCommittedIterator is returned from FilterOrderCommitted and is used to iterate over the raw logs and unpacked data for OrderCommitted events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismOrderCommittedIterator struct {
	Event *SpotMarketOptimismOrderCommitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismOrderCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismOrderCommitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismOrderCommitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismOrderCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismOrderCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismOrderCommitted represents a OrderCommitted event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismOrderCommitted struct {
	MarketId       *big.Int
	OrderType      uint8
	AmountProvided *big.Int
	AsyncOrderId   *big.Int
	Sender         common.Address
	Referrer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOrderCommitted is a free log retrieval operation binding the contract event 0xb26c216bf0a127dddc2431e4d8ca845513c8e6fb80e754296e7afab1ce92722f.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint8 indexed orderType, uint256 amountProvided, uint128 asyncOrderId, address indexed sender, address referrer)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterOrderCommitted(opts *bind.FilterOpts, marketId []*big.Int, orderType []uint8, sender []common.Address) (*SpotMarketOptimismOrderCommittedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var orderTypeRule []interface{}
	for _, orderTypeItem := range orderType {
		orderTypeRule = append(orderTypeRule, orderTypeItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "OrderCommitted", marketIdRule, orderTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismOrderCommittedIterator{contract: _SpotMarketOptimism.contract, event: "OrderCommitted", logs: logs, sub: sub}, nil
}

// WatchOrderCommitted is a free log subscription operation binding the contract event 0xb26c216bf0a127dddc2431e4d8ca845513c8e6fb80e754296e7afab1ce92722f.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint8 indexed orderType, uint256 amountProvided, uint128 asyncOrderId, address indexed sender, address referrer)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchOrderCommitted(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismOrderCommitted, marketId []*big.Int, orderType []uint8, sender []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var orderTypeRule []interface{}
	for _, orderTypeItem := range orderType {
		orderTypeRule = append(orderTypeRule, orderTypeItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "OrderCommitted", marketIdRule, orderTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismOrderCommitted)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderCommitted is a log parse operation binding the contract event 0xb26c216bf0a127dddc2431e4d8ca845513c8e6fb80e754296e7afab1ce92722f.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint8 indexed orderType, uint256 amountProvided, uint128 asyncOrderId, address indexed sender, address referrer)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseOrderCommitted(log types.Log) (*SpotMarketOptimismOrderCommitted, error) {
	event := new(SpotMarketOptimismOrderCommitted)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismOrderSettledIterator is returned from FilterOrderSettled and is used to iterate over the raw logs and unpacked data for OrderSettled events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismOrderSettledIterator struct {
	Event *SpotMarketOptimismOrderSettled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismOrderSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismOrderSettled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismOrderSettled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismOrderSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismOrderSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismOrderSettled represents a OrderSettled event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismOrderSettled struct {
	MarketId         *big.Int
	AsyncOrderId     *big.Int
	FinalOrderAmount *big.Int
	Fees             OrderFeesData
	CollectedFees    *big.Int
	Settler          common.Address
	Price            *big.Int
	OrderType        uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderSettled is a free log retrieval operation binding the contract event 0x6c5e8ff282d52fb9f532408e86d4afc62fc1f89c749a8ddca7a6f34c0439a183.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed asyncOrderId, uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address indexed settler, uint256 price, uint8 orderType)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterOrderSettled(opts *bind.FilterOpts, marketId []*big.Int, asyncOrderId []*big.Int, settler []common.Address) (*SpotMarketOptimismOrderSettledIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var asyncOrderIdRule []interface{}
	for _, asyncOrderIdItem := range asyncOrderId {
		asyncOrderIdRule = append(asyncOrderIdRule, asyncOrderIdItem)
	}

	var settlerRule []interface{}
	for _, settlerItem := range settler {
		settlerRule = append(settlerRule, settlerItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "OrderSettled", marketIdRule, asyncOrderIdRule, settlerRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismOrderSettledIterator{contract: _SpotMarketOptimism.contract, event: "OrderSettled", logs: logs, sub: sub}, nil
}

// WatchOrderSettled is a free log subscription operation binding the contract event 0x6c5e8ff282d52fb9f532408e86d4afc62fc1f89c749a8ddca7a6f34c0439a183.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed asyncOrderId, uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address indexed settler, uint256 price, uint8 orderType)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchOrderSettled(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismOrderSettled, marketId []*big.Int, asyncOrderId []*big.Int, settler []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var asyncOrderIdRule []interface{}
	for _, asyncOrderIdItem := range asyncOrderId {
		asyncOrderIdRule = append(asyncOrderIdRule, asyncOrderIdItem)
	}

	var settlerRule []interface{}
	for _, settlerItem := range settler {
		settlerRule = append(settlerRule, settlerItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "OrderSettled", marketIdRule, asyncOrderIdRule, settlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismOrderSettled)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "OrderSettled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderSettled is a log parse operation binding the contract event 0x6c5e8ff282d52fb9f532408e86d4afc62fc1f89c749a8ddca7a6f34c0439a183.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed asyncOrderId, uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address indexed settler, uint256 price, uint8 orderType)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseOrderSettled(log types.Log) (*SpotMarketOptimismOrderSettled, error) {
	event := new(SpotMarketOptimismOrderSettled)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "OrderSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismOwnerChangedIterator struct {
	Event *SpotMarketOptimismOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismOwnerChanged represents a OwnerChanged event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SpotMarketOptimismOwnerChangedIterator, error) {

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismOwnerChangedIterator{contract: _SpotMarketOptimism.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismOwnerChanged)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseOwnerChanged(log types.Log) (*SpotMarketOptimismOwnerChanged, error) {
	event := new(SpotMarketOptimismOwnerChanged)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismOwnerNominatedIterator struct {
	Event *SpotMarketOptimismOwnerNominated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismOwnerNominated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismOwnerNominated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismOwnerNominated represents a OwnerNominated event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*SpotMarketOptimismOwnerNominatedIterator, error) {

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismOwnerNominatedIterator{contract: _SpotMarketOptimism.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismOwnerNominated)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseOwnerNominated(log types.Log) (*SpotMarketOptimismOwnerNominated, error) {
	event := new(SpotMarketOptimismOwnerNominated)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismReferrerShareUpdatedIterator is returned from FilterReferrerShareUpdated and is used to iterate over the raw logs and unpacked data for ReferrerShareUpdated events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismReferrerShareUpdatedIterator struct {
	Event *SpotMarketOptimismReferrerShareUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismReferrerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismReferrerShareUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismReferrerShareUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismReferrerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismReferrerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismReferrerShareUpdated represents a ReferrerShareUpdated event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismReferrerShareUpdated struct {
	MarketId        *big.Int
	Referrer        common.Address
	SharePercentage *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReferrerShareUpdated is a free log retrieval operation binding the contract event 0xd2a3339bb3c610e9030023c1cb3e89374fe0ebd7e37faee9b3d343f33e9df2fb.
//
// Solidity: event ReferrerShareUpdated(uint128 indexed marketId, address referrer, uint256 sharePercentage)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterReferrerShareUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*SpotMarketOptimismReferrerShareUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "ReferrerShareUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismReferrerShareUpdatedIterator{contract: _SpotMarketOptimism.contract, event: "ReferrerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchReferrerShareUpdated is a free log subscription operation binding the contract event 0xd2a3339bb3c610e9030023c1cb3e89374fe0ebd7e37faee9b3d343f33e9df2fb.
//
// Solidity: event ReferrerShareUpdated(uint128 indexed marketId, address referrer, uint256 sharePercentage)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchReferrerShareUpdated(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismReferrerShareUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "ReferrerShareUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismReferrerShareUpdated)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReferrerShareUpdated is a log parse operation binding the contract event 0xd2a3339bb3c610e9030023c1cb3e89374fe0ebd7e37faee9b3d343f33e9df2fb.
//
// Solidity: event ReferrerShareUpdated(uint128 indexed marketId, address referrer, uint256 sharePercentage)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseReferrerShareUpdated(log types.Log) (*SpotMarketOptimismReferrerShareUpdated, error) {
	event := new(SpotMarketOptimismReferrerShareUpdated)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSettlementStrategyAddedIterator is returned from FilterSettlementStrategyAdded and is used to iterate over the raw logs and unpacked data for SettlementStrategyAdded events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSettlementStrategyAddedIterator struct {
	Event *SpotMarketOptimismSettlementStrategyAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSettlementStrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSettlementStrategyAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSettlementStrategyAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSettlementStrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSettlementStrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSettlementStrategyAdded represents a SettlementStrategyAdded event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSettlementStrategyAdded struct {
	SynthMarketId *big.Int
	StrategyId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyAdded is a free log retrieval operation binding the contract event 0x91750a8e3d84ed1bccdc79dcecf63cc1b6f83b5bf8293bf86628cbb248e487f1.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed synthMarketId, uint256 indexed strategyId)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSettlementStrategyAdded(opts *bind.FilterOpts, synthMarketId []*big.Int, strategyId []*big.Int) (*SpotMarketOptimismSettlementStrategyAddedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SettlementStrategyAdded", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSettlementStrategyAddedIterator{contract: _SpotMarketOptimism.contract, event: "SettlementStrategyAdded", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyAdded is a free log subscription operation binding the contract event 0x91750a8e3d84ed1bccdc79dcecf63cc1b6f83b5bf8293bf86628cbb248e487f1.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed synthMarketId, uint256 indexed strategyId)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSettlementStrategyAdded(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSettlementStrategyAdded, synthMarketId []*big.Int, strategyId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SettlementStrategyAdded", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSettlementStrategyAdded)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettlementStrategyAdded is a log parse operation binding the contract event 0x91750a8e3d84ed1bccdc79dcecf63cc1b6f83b5bf8293bf86628cbb248e487f1.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed synthMarketId, uint256 indexed strategyId)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSettlementStrategyAdded(log types.Log) (*SpotMarketOptimismSettlementStrategyAdded, error) {
	event := new(SpotMarketOptimismSettlementStrategyAdded)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSettlementStrategyUpdatedIterator is returned from FilterSettlementStrategyUpdated and is used to iterate over the raw logs and unpacked data for SettlementStrategyUpdated events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSettlementStrategyUpdatedIterator struct {
	Event *SpotMarketOptimismSettlementStrategyUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSettlementStrategyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSettlementStrategyUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSettlementStrategyUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSettlementStrategyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSettlementStrategyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSettlementStrategyUpdated represents a SettlementStrategyUpdated event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSettlementStrategyUpdated struct {
	SynthMarketId *big.Int
	StrategyId    *big.Int
	Enabled       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyUpdated is a free log retrieval operation binding the contract event 0xb0187d91c1ffe66740e5f7619aab4d868a85ac3698b026bed3f420d871b6c048.
//
// Solidity: event SettlementStrategyUpdated(uint128 indexed synthMarketId, uint256 indexed strategyId, bool enabled)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSettlementStrategyUpdated(opts *bind.FilterOpts, synthMarketId []*big.Int, strategyId []*big.Int) (*SpotMarketOptimismSettlementStrategyUpdatedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SettlementStrategyUpdated", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSettlementStrategyUpdatedIterator{contract: _SpotMarketOptimism.contract, event: "SettlementStrategyUpdated", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyUpdated is a free log subscription operation binding the contract event 0xb0187d91c1ffe66740e5f7619aab4d868a85ac3698b026bed3f420d871b6c048.
//
// Solidity: event SettlementStrategyUpdated(uint128 indexed synthMarketId, uint256 indexed strategyId, bool enabled)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSettlementStrategyUpdated(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSettlementStrategyUpdated, synthMarketId []*big.Int, strategyId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SettlementStrategyUpdated", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSettlementStrategyUpdated)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SettlementStrategyUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettlementStrategyUpdated is a log parse operation binding the contract event 0xb0187d91c1ffe66740e5f7619aab4d868a85ac3698b026bed3f420d871b6c048.
//
// Solidity: event SettlementStrategyUpdated(uint128 indexed synthMarketId, uint256 indexed strategyId, bool enabled)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSettlementStrategyUpdated(log types.Log) (*SpotMarketOptimismSettlementStrategyUpdated, error) {
	event := new(SpotMarketOptimismSettlementStrategyUpdated)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SettlementStrategyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSynthBoughtIterator is returned from FilterSynthBought and is used to iterate over the raw logs and unpacked data for SynthBought events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthBoughtIterator struct {
	Event *SpotMarketOptimismSynthBought // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSynthBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSynthBought)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSynthBought)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSynthBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSynthBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSynthBought represents a SynthBought event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthBought struct {
	SynthMarketId *big.Int
	SynthReturned *big.Int
	Fees          OrderFeesData
	CollectedFees *big.Int
	Referrer      common.Address
	Price         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSynthBought is a free log retrieval operation binding the contract event 0xac82d63e679c7d862613aa8b5ccd94f9adc4986763ab14bb3351ab9092ef1303.
//
// Solidity: event SynthBought(uint256 indexed synthMarketId, uint256 synthReturned, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address referrer, uint256 price)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSynthBought(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismSynthBoughtIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SynthBought", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSynthBoughtIterator{contract: _SpotMarketOptimism.contract, event: "SynthBought", logs: logs, sub: sub}, nil
}

// WatchSynthBought is a free log subscription operation binding the contract event 0xac82d63e679c7d862613aa8b5ccd94f9adc4986763ab14bb3351ab9092ef1303.
//
// Solidity: event SynthBought(uint256 indexed synthMarketId, uint256 synthReturned, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address referrer, uint256 price)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSynthBought(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSynthBought, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SynthBought", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSynthBought)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthBought", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSynthBought is a log parse operation binding the contract event 0xac82d63e679c7d862613aa8b5ccd94f9adc4986763ab14bb3351ab9092ef1303.
//
// Solidity: event SynthBought(uint256 indexed synthMarketId, uint256 synthReturned, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address referrer, uint256 price)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSynthBought(log types.Log) (*SpotMarketOptimismSynthBought, error) {
	event := new(SpotMarketOptimismSynthBought)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSynthImplementationSetIterator is returned from FilterSynthImplementationSet and is used to iterate over the raw logs and unpacked data for SynthImplementationSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthImplementationSetIterator struct {
	Event *SpotMarketOptimismSynthImplementationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSynthImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSynthImplementationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSynthImplementationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSynthImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSynthImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSynthImplementationSet represents a SynthImplementationSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthImplementationSet struct {
	SynthImplementation common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSynthImplementationSet is a free log retrieval operation binding the contract event 0xafffc48b3243eba10d901f21ba761ad741f18a23feed86ca425df4974d3314b0.
//
// Solidity: event SynthImplementationSet(address synthImplementation)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSynthImplementationSet(opts *bind.FilterOpts) (*SpotMarketOptimismSynthImplementationSetIterator, error) {

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SynthImplementationSet")
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSynthImplementationSetIterator{contract: _SpotMarketOptimism.contract, event: "SynthImplementationSet", logs: logs, sub: sub}, nil
}

// WatchSynthImplementationSet is a free log subscription operation binding the contract event 0xafffc48b3243eba10d901f21ba761ad741f18a23feed86ca425df4974d3314b0.
//
// Solidity: event SynthImplementationSet(address synthImplementation)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSynthImplementationSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSynthImplementationSet) (event.Subscription, error) {

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SynthImplementationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSynthImplementationSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthImplementationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSynthImplementationSet is a log parse operation binding the contract event 0xafffc48b3243eba10d901f21ba761ad741f18a23feed86ca425df4974d3314b0.
//
// Solidity: event SynthImplementationSet(address synthImplementation)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSynthImplementationSet(log types.Log) (*SpotMarketOptimismSynthImplementationSet, error) {
	event := new(SpotMarketOptimismSynthImplementationSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSynthImplementationUpgradedIterator is returned from FilterSynthImplementationUpgraded and is used to iterate over the raw logs and unpacked data for SynthImplementationUpgraded events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthImplementationUpgradedIterator struct {
	Event *SpotMarketOptimismSynthImplementationUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSynthImplementationUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSynthImplementationUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSynthImplementationUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSynthImplementationUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSynthImplementationUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSynthImplementationUpgraded represents a SynthImplementationUpgraded event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthImplementationUpgraded struct {
	SynthMarketId  *big.Int
	Proxy          common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSynthImplementationUpgraded is a free log retrieval operation binding the contract event 0xa7badace8b24daeeb3981497a506d3812b3dc6f147ef3f78bc0e2cc664c50330.
//
// Solidity: event SynthImplementationUpgraded(uint256 indexed synthMarketId, address indexed proxy, address implementation)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSynthImplementationUpgraded(opts *bind.FilterOpts, synthMarketId []*big.Int, proxy []common.Address) (*SpotMarketOptimismSynthImplementationUpgradedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SynthImplementationUpgraded", synthMarketIdRule, proxyRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSynthImplementationUpgradedIterator{contract: _SpotMarketOptimism.contract, event: "SynthImplementationUpgraded", logs: logs, sub: sub}, nil
}

// WatchSynthImplementationUpgraded is a free log subscription operation binding the contract event 0xa7badace8b24daeeb3981497a506d3812b3dc6f147ef3f78bc0e2cc664c50330.
//
// Solidity: event SynthImplementationUpgraded(uint256 indexed synthMarketId, address indexed proxy, address implementation)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSynthImplementationUpgraded(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSynthImplementationUpgraded, synthMarketId []*big.Int, proxy []common.Address) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SynthImplementationUpgraded", synthMarketIdRule, proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSynthImplementationUpgraded)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthImplementationUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSynthImplementationUpgraded is a log parse operation binding the contract event 0xa7badace8b24daeeb3981497a506d3812b3dc6f147ef3f78bc0e2cc664c50330.
//
// Solidity: event SynthImplementationUpgraded(uint256 indexed synthMarketId, address indexed proxy, address implementation)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSynthImplementationUpgraded(log types.Log) (*SpotMarketOptimismSynthImplementationUpgraded, error) {
	event := new(SpotMarketOptimismSynthImplementationUpgraded)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthImplementationUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSynthPriceDataUpdatedIterator is returned from FilterSynthPriceDataUpdated and is used to iterate over the raw logs and unpacked data for SynthPriceDataUpdated events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthPriceDataUpdatedIterator struct {
	Event *SpotMarketOptimismSynthPriceDataUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSynthPriceDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSynthPriceDataUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSynthPriceDataUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSynthPriceDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSynthPriceDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSynthPriceDataUpdated represents a SynthPriceDataUpdated event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthPriceDataUpdated struct {
	SynthMarketId *big.Int
	BuyFeedId     [32]byte
	SellFeedId    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSynthPriceDataUpdated is a free log retrieval operation binding the contract event 0x30309eafe094276dfad42e7ac4676aebcf5ea8480c86708ad2a2c9c0335c912e.
//
// Solidity: event SynthPriceDataUpdated(uint256 indexed synthMarketId, bytes32 indexed buyFeedId, bytes32 indexed sellFeedId)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSynthPriceDataUpdated(opts *bind.FilterOpts, synthMarketId []*big.Int, buyFeedId [][32]byte, sellFeedId [][32]byte) (*SpotMarketOptimismSynthPriceDataUpdatedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var buyFeedIdRule []interface{}
	for _, buyFeedIdItem := range buyFeedId {
		buyFeedIdRule = append(buyFeedIdRule, buyFeedIdItem)
	}
	var sellFeedIdRule []interface{}
	for _, sellFeedIdItem := range sellFeedId {
		sellFeedIdRule = append(sellFeedIdRule, sellFeedIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SynthPriceDataUpdated", synthMarketIdRule, buyFeedIdRule, sellFeedIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSynthPriceDataUpdatedIterator{contract: _SpotMarketOptimism.contract, event: "SynthPriceDataUpdated", logs: logs, sub: sub}, nil
}

// WatchSynthPriceDataUpdated is a free log subscription operation binding the contract event 0x30309eafe094276dfad42e7ac4676aebcf5ea8480c86708ad2a2c9c0335c912e.
//
// Solidity: event SynthPriceDataUpdated(uint256 indexed synthMarketId, bytes32 indexed buyFeedId, bytes32 indexed sellFeedId)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSynthPriceDataUpdated(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSynthPriceDataUpdated, synthMarketId []*big.Int, buyFeedId [][32]byte, sellFeedId [][32]byte) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var buyFeedIdRule []interface{}
	for _, buyFeedIdItem := range buyFeedId {
		buyFeedIdRule = append(buyFeedIdRule, buyFeedIdItem)
	}
	var sellFeedIdRule []interface{}
	for _, sellFeedIdItem := range sellFeedId {
		sellFeedIdRule = append(sellFeedIdRule, sellFeedIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SynthPriceDataUpdated", synthMarketIdRule, buyFeedIdRule, sellFeedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSynthPriceDataUpdated)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthPriceDataUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSynthPriceDataUpdated is a log parse operation binding the contract event 0x30309eafe094276dfad42e7ac4676aebcf5ea8480c86708ad2a2c9c0335c912e.
//
// Solidity: event SynthPriceDataUpdated(uint256 indexed synthMarketId, bytes32 indexed buyFeedId, bytes32 indexed sellFeedId)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSynthPriceDataUpdated(log types.Log) (*SpotMarketOptimismSynthPriceDataUpdated, error) {
	event := new(SpotMarketOptimismSynthPriceDataUpdated)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthPriceDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSynthRegisteredIterator is returned from FilterSynthRegistered and is used to iterate over the raw logs and unpacked data for SynthRegistered events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthRegisteredIterator struct {
	Event *SpotMarketOptimismSynthRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSynthRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSynthRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSynthRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSynthRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSynthRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSynthRegistered represents a SynthRegistered event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthRegistered struct {
	SynthMarketId     *big.Int
	SynthTokenAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSynthRegistered is a free log retrieval operation binding the contract event 0x04b07b1c236913894927915a3dd91c8e250223fc668ceabdc47074c5a77e2cb9.
//
// Solidity: event SynthRegistered(uint256 indexed synthMarketId, address synthTokenAddress)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSynthRegistered(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismSynthRegisteredIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SynthRegistered", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSynthRegisteredIterator{contract: _SpotMarketOptimism.contract, event: "SynthRegistered", logs: logs, sub: sub}, nil
}

// WatchSynthRegistered is a free log subscription operation binding the contract event 0x04b07b1c236913894927915a3dd91c8e250223fc668ceabdc47074c5a77e2cb9.
//
// Solidity: event SynthRegistered(uint256 indexed synthMarketId, address synthTokenAddress)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSynthRegistered(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSynthRegistered, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SynthRegistered", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSynthRegistered)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSynthRegistered is a log parse operation binding the contract event 0x04b07b1c236913894927915a3dd91c8e250223fc668ceabdc47074c5a77e2cb9.
//
// Solidity: event SynthRegistered(uint256 indexed synthMarketId, address synthTokenAddress)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSynthRegistered(log types.Log) (*SpotMarketOptimismSynthRegistered, error) {
	event := new(SpotMarketOptimismSynthRegistered)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSynthSoldIterator is returned from FilterSynthSold and is used to iterate over the raw logs and unpacked data for SynthSold events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthSoldIterator struct {
	Event *SpotMarketOptimismSynthSold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSynthSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSynthSold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSynthSold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSynthSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSynthSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSynthSold represents a SynthSold event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthSold struct {
	SynthMarketId  *big.Int
	AmountReturned *big.Int
	Fees           OrderFeesData
	CollectedFees  *big.Int
	Referrer       common.Address
	Price          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSynthSold is a free log retrieval operation binding the contract event 0x61fa4bb370a2f18a502b3bcf1d0755e53371d58791fa42766aa6386bbefb594a.
//
// Solidity: event SynthSold(uint256 indexed synthMarketId, uint256 amountReturned, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address referrer, uint256 price)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSynthSold(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismSynthSoldIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SynthSold", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSynthSoldIterator{contract: _SpotMarketOptimism.contract, event: "SynthSold", logs: logs, sub: sub}, nil
}

// WatchSynthSold is a free log subscription operation binding the contract event 0x61fa4bb370a2f18a502b3bcf1d0755e53371d58791fa42766aa6386bbefb594a.
//
// Solidity: event SynthSold(uint256 indexed synthMarketId, uint256 amountReturned, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address referrer, uint256 price)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSynthSold(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSynthSold, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SynthSold", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSynthSold)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthSold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSynthSold is a log parse operation binding the contract event 0x61fa4bb370a2f18a502b3bcf1d0755e53371d58791fa42766aa6386bbefb594a.
//
// Solidity: event SynthSold(uint256 indexed synthMarketId, uint256 amountReturned, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address referrer, uint256 price)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSynthSold(log types.Log) (*SpotMarketOptimismSynthSold, error) {
	event := new(SpotMarketOptimismSynthSold)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSynthUnwrappedIterator is returned from FilterSynthUnwrapped and is used to iterate over the raw logs and unpacked data for SynthUnwrapped events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthUnwrappedIterator struct {
	Event *SpotMarketOptimismSynthUnwrapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSynthUnwrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSynthUnwrapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSynthUnwrapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSynthUnwrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSynthUnwrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSynthUnwrapped represents a SynthUnwrapped event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthUnwrapped struct {
	SynthMarketId   *big.Int
	AmountUnwrapped *big.Int
	Fees            OrderFeesData
	FeesCollected   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSynthUnwrapped is a free log retrieval operation binding the contract event 0xa1dd74fb936c7942732e4355961ca6944ca6c6744121ace0d9a1203d664231b3.
//
// Solidity: event SynthUnwrapped(uint256 indexed synthMarketId, uint256 amountUnwrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSynthUnwrapped(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismSynthUnwrappedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SynthUnwrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSynthUnwrappedIterator{contract: _SpotMarketOptimism.contract, event: "SynthUnwrapped", logs: logs, sub: sub}, nil
}

// WatchSynthUnwrapped is a free log subscription operation binding the contract event 0xa1dd74fb936c7942732e4355961ca6944ca6c6744121ace0d9a1203d664231b3.
//
// Solidity: event SynthUnwrapped(uint256 indexed synthMarketId, uint256 amountUnwrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSynthUnwrapped(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSynthUnwrapped, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SynthUnwrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSynthUnwrapped)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthUnwrapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSynthUnwrapped is a log parse operation binding the contract event 0xa1dd74fb936c7942732e4355961ca6944ca6c6744121ace0d9a1203d664231b3.
//
// Solidity: event SynthUnwrapped(uint256 indexed synthMarketId, uint256 amountUnwrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSynthUnwrapped(log types.Log) (*SpotMarketOptimismSynthUnwrapped, error) {
	event := new(SpotMarketOptimismSynthUnwrapped)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthUnwrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSynthWrappedIterator is returned from FilterSynthWrapped and is used to iterate over the raw logs and unpacked data for SynthWrapped events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthWrappedIterator struct {
	Event *SpotMarketOptimismSynthWrapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSynthWrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSynthWrapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSynthWrapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSynthWrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSynthWrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSynthWrapped represents a SynthWrapped event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthWrapped struct {
	SynthMarketId *big.Int
	AmountWrapped *big.Int
	Fees          OrderFeesData
	FeesCollected *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSynthWrapped is a free log retrieval operation binding the contract event 0xea50ab2f37d37692441c4a16317c1287bb410a3a616a16c49c9ca76d415667ff.
//
// Solidity: event SynthWrapped(uint256 indexed synthMarketId, uint256 amountWrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSynthWrapped(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismSynthWrappedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SynthWrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSynthWrappedIterator{contract: _SpotMarketOptimism.contract, event: "SynthWrapped", logs: logs, sub: sub}, nil
}

// WatchSynthWrapped is a free log subscription operation binding the contract event 0xea50ab2f37d37692441c4a16317c1287bb410a3a616a16c49c9ca76d415667ff.
//
// Solidity: event SynthWrapped(uint256 indexed synthMarketId, uint256 amountWrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSynthWrapped(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSynthWrapped, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SynthWrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSynthWrapped)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthWrapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSynthWrapped is a log parse operation binding the contract event 0xea50ab2f37d37692441c4a16317c1287bb410a3a616a16c49c9ca76d415667ff.
//
// Solidity: event SynthWrapped(uint256 indexed synthMarketId, uint256 amountWrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSynthWrapped(log types.Log) (*SpotMarketOptimismSynthWrapped, error) {
	event := new(SpotMarketOptimismSynthWrapped)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthWrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismSynthetixSystemSetIterator is returned from FilterSynthetixSystemSet and is used to iterate over the raw logs and unpacked data for SynthetixSystemSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthetixSystemSetIterator struct {
	Event *SpotMarketOptimismSynthetixSystemSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismSynthetixSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismSynthetixSystemSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismSynthetixSystemSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismSynthetixSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismSynthetixSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismSynthetixSystemSet represents a SynthetixSystemSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismSynthetixSystemSet struct {
	Synthetix       common.Address
	UsdTokenAddress common.Address
	OracleManager   common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSynthetixSystemSet is a free log retrieval operation binding the contract event 0x52dccf7e2653d5ae8cf1d18c5499fd530f01920181d81afdf6bf489d897e24aa.
//
// Solidity: event SynthetixSystemSet(address synthetix, address usdTokenAddress, address oracleManager)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterSynthetixSystemSet(opts *bind.FilterOpts) (*SpotMarketOptimismSynthetixSystemSetIterator, error) {

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "SynthetixSystemSet")
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismSynthetixSystemSetIterator{contract: _SpotMarketOptimism.contract, event: "SynthetixSystemSet", logs: logs, sub: sub}, nil
}

// WatchSynthetixSystemSet is a free log subscription operation binding the contract event 0x52dccf7e2653d5ae8cf1d18c5499fd530f01920181d81afdf6bf489d897e24aa.
//
// Solidity: event SynthetixSystemSet(address synthetix, address usdTokenAddress, address oracleManager)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchSynthetixSystemSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismSynthetixSystemSet) (event.Subscription, error) {

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "SynthetixSystemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismSynthetixSystemSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthetixSystemSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSynthetixSystemSet is a log parse operation binding the contract event 0x52dccf7e2653d5ae8cf1d18c5499fd530f01920181d81afdf6bf489d897e24aa.
//
// Solidity: event SynthetixSystemSet(address synthetix, address usdTokenAddress, address oracleManager)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseSynthetixSystemSet(log types.Log) (*SpotMarketOptimismSynthetixSystemSet, error) {
	event := new(SpotMarketOptimismSynthetixSystemSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "SynthetixSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismTransactorFixedFeeSetIterator is returned from FilterTransactorFixedFeeSet and is used to iterate over the raw logs and unpacked data for TransactorFixedFeeSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismTransactorFixedFeeSetIterator struct {
	Event *SpotMarketOptimismTransactorFixedFeeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismTransactorFixedFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismTransactorFixedFeeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismTransactorFixedFeeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismTransactorFixedFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismTransactorFixedFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismTransactorFixedFeeSet represents a TransactorFixedFeeSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismTransactorFixedFeeSet struct {
	SynthMarketId  *big.Int
	Transactor     common.Address
	FixedFeeAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransactorFixedFeeSet is a free log retrieval operation binding the contract event 0xeed7c7ebc4a7e7456a5b14f961bbe55d026f35a2a2b52d1ad43fe11c348df24a.
//
// Solidity: event TransactorFixedFeeSet(uint256 indexed synthMarketId, address transactor, uint256 fixedFeeAmount)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterTransactorFixedFeeSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismTransactorFixedFeeSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "TransactorFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismTransactorFixedFeeSetIterator{contract: _SpotMarketOptimism.contract, event: "TransactorFixedFeeSet", logs: logs, sub: sub}, nil
}

// WatchTransactorFixedFeeSet is a free log subscription operation binding the contract event 0xeed7c7ebc4a7e7456a5b14f961bbe55d026f35a2a2b52d1ad43fe11c348df24a.
//
// Solidity: event TransactorFixedFeeSet(uint256 indexed synthMarketId, address transactor, uint256 fixedFeeAmount)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchTransactorFixedFeeSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismTransactorFixedFeeSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "TransactorFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismTransactorFixedFeeSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "TransactorFixedFeeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransactorFixedFeeSet is a log parse operation binding the contract event 0xeed7c7ebc4a7e7456a5b14f961bbe55d026f35a2a2b52d1ad43fe11c348df24a.
//
// Solidity: event TransactorFixedFeeSet(uint256 indexed synthMarketId, address transactor, uint256 fixedFeeAmount)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseTransactorFixedFeeSet(log types.Log) (*SpotMarketOptimismTransactorFixedFeeSet, error) {
	event := new(SpotMarketOptimismTransactorFixedFeeSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "TransactorFixedFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismUpgradedIterator struct {
	Event *SpotMarketOptimismUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismUpgraded represents a Upgraded event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismUpgraded struct {
	Self           common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterUpgraded(opts *bind.FilterOpts, self []common.Address) (*SpotMarketOptimismUpgradedIterator, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismUpgradedIterator{contract: _SpotMarketOptimism.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismUpgraded, self []common.Address) (event.Subscription, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismUpgraded)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseUpgraded(log types.Log) (*SpotMarketOptimismUpgraded, error) {
	event := new(SpotMarketOptimismUpgraded)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismWrapperFeesSetIterator is returned from FilterWrapperFeesSet and is used to iterate over the raw logs and unpacked data for WrapperFeesSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismWrapperFeesSetIterator struct {
	Event *SpotMarketOptimismWrapperFeesSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismWrapperFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismWrapperFeesSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismWrapperFeesSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismWrapperFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismWrapperFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismWrapperFeesSet represents a WrapperFeesSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismWrapperFeesSet struct {
	SynthMarketId *big.Int
	WrapFee       *big.Int
	UnwrapFee     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWrapperFeesSet is a free log retrieval operation binding the contract event 0x84c5bc20d6e52e92afe6ebc9d85d3e4d35de276ba3f05cae640db053f5b861b8.
//
// Solidity: event WrapperFeesSet(uint256 indexed synthMarketId, int256 wrapFee, int256 unwrapFee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterWrapperFeesSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketOptimismWrapperFeesSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "WrapperFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismWrapperFeesSetIterator{contract: _SpotMarketOptimism.contract, event: "WrapperFeesSet", logs: logs, sub: sub}, nil
}

// WatchWrapperFeesSet is a free log subscription operation binding the contract event 0x84c5bc20d6e52e92afe6ebc9d85d3e4d35de276ba3f05cae640db053f5b861b8.
//
// Solidity: event WrapperFeesSet(uint256 indexed synthMarketId, int256 wrapFee, int256 unwrapFee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchWrapperFeesSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismWrapperFeesSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "WrapperFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismWrapperFeesSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "WrapperFeesSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWrapperFeesSet is a log parse operation binding the contract event 0x84c5bc20d6e52e92afe6ebc9d85d3e4d35de276ba3f05cae640db053f5b861b8.
//
// Solidity: event WrapperFeesSet(uint256 indexed synthMarketId, int256 wrapFee, int256 unwrapFee)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseWrapperFeesSet(log types.Log) (*SpotMarketOptimismWrapperFeesSet, error) {
	event := new(SpotMarketOptimismWrapperFeesSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "WrapperFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketOptimismWrapperSetIterator is returned from FilterWrapperSet and is used to iterate over the raw logs and unpacked data for WrapperSet events raised by the SpotMarketOptimism contract.
type SpotMarketOptimismWrapperSetIterator struct {
	Event *SpotMarketOptimismWrapperSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SpotMarketOptimismWrapperSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketOptimismWrapperSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SpotMarketOptimismWrapperSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SpotMarketOptimismWrapperSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketOptimismWrapperSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketOptimismWrapperSet represents a WrapperSet event raised by the SpotMarketOptimism contract.
type SpotMarketOptimismWrapperSet struct {
	SynthMarketId      *big.Int
	WrapCollateralType common.Address
	MaxWrappableAmount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWrapperSet is a free log retrieval operation binding the contract event 0xf6b8d296783aecfc5d372dff3e3e802ab63338637f9a2f3e2aae1e745c148def.
//
// Solidity: event WrapperSet(uint256 indexed synthMarketId, address indexed wrapCollateralType, uint256 maxWrappableAmount)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) FilterWrapperSet(opts *bind.FilterOpts, synthMarketId []*big.Int, wrapCollateralType []common.Address) (*SpotMarketOptimismWrapperSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var wrapCollateralTypeRule []interface{}
	for _, wrapCollateralTypeItem := range wrapCollateralType {
		wrapCollateralTypeRule = append(wrapCollateralTypeRule, wrapCollateralTypeItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.FilterLogs(opts, "WrapperSet", synthMarketIdRule, wrapCollateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketOptimismWrapperSetIterator{contract: _SpotMarketOptimism.contract, event: "WrapperSet", logs: logs, sub: sub}, nil
}

// WatchWrapperSet is a free log subscription operation binding the contract event 0xf6b8d296783aecfc5d372dff3e3e802ab63338637f9a2f3e2aae1e745c148def.
//
// Solidity: event WrapperSet(uint256 indexed synthMarketId, address indexed wrapCollateralType, uint256 maxWrappableAmount)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) WatchWrapperSet(opts *bind.WatchOpts, sink chan<- *SpotMarketOptimismWrapperSet, synthMarketId []*big.Int, wrapCollateralType []common.Address) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var wrapCollateralTypeRule []interface{}
	for _, wrapCollateralTypeItem := range wrapCollateralType {
		wrapCollateralTypeRule = append(wrapCollateralTypeRule, wrapCollateralTypeItem)
	}

	logs, sub, err := _SpotMarketOptimism.contract.WatchLogs(opts, "WrapperSet", synthMarketIdRule, wrapCollateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketOptimismWrapperSet)
				if err := _SpotMarketOptimism.contract.UnpackLog(event, "WrapperSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWrapperSet is a log parse operation binding the contract event 0xf6b8d296783aecfc5d372dff3e3e802ab63338637f9a2f3e2aae1e745c148def.
//
// Solidity: event WrapperSet(uint256 indexed synthMarketId, address indexed wrapCollateralType, uint256 maxWrappableAmount)
func (_SpotMarketOptimism *SpotMarketOptimismFilterer) ParseWrapperSet(log types.Log) (*SpotMarketOptimismWrapperSet, error) {
	event := new(SpotMarketOptimismWrapperSet)
	if err := _SpotMarketOptimism.contract.UnpackLog(event, "WrapperSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
