// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package spotMarketGoerli

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

// SpotMarketGoerliMetaData contains all meta data concerning the SpotMarketGoerli contract.
var SpotMarketGoerliMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMarketOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"synthImplementation\",\"type\":\"uint256\"}],\"name\":\"InvalidSynthImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OnlyMarketOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"DecayRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nominee\",\"type\":\"address\"}],\"name\":\"MarketNominationRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"MarketOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"MarketOwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"synthImplementation\",\"type\":\"address\"}],\"name\":\"SynthImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"SynthImplementationUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"buyFeedId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sellFeedId\",\"type\":\"bytes32\"}],\"name\":\"SynthPriceDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"synthTokenAddress\",\"type\":\"address\"}],\"name\":\"SynthRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"synthetix\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usdTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracleManager\",\"type\":\"address\"}],\"name\":\"SynthetixSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"acceptMarketOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"synthOwner\",\"type\":\"address\"}],\"name\":\"createSynth\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"marketOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getSynth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"synthAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getSynthImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"minimumCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateMarketOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"renounceMarketNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"renounceMarketOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"reportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportedDebtAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setDecayRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"synthImplementation\",\"type\":\"address\"}],\"name\":\"setSynthImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISynthetixSystem\",\"name\":\"synthetix\",\"type\":\"address\"}],\"name\":\"setSynthetix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"buyFeedId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sellFeedId\",\"type\":\"bytes32\"}],\"name\":\"updatePriceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"upgradeSynthImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSynthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"synthAmountCharged\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxSynthAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxUsdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdAmountCharged\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxUsdAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"InsufficientAmountReceived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"InvalidMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrices\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"synthReturned\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SynthBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReturned\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SynthSold\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"buyExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxUsdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"buyExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmountCharged\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"}],\"name\":\"quoteBuyExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"}],\"name\":\"quoteBuyExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmountCharged\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"}],\"name\":\"quoteSellExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"}],\"name\":\"quoteSellExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthToBurn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minUsdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"sell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmountReceived\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"sellExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSynthAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"sellExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthToBurn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"IneligibleForCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InsufficientSharesAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTransaction.Type\",\"name\":\"transactionType\",\"type\":\"uint8\"}],\"name\":\"InvalidAsyncTransactionType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"asyncOrderId\",\"type\":\"uint256\"}],\"name\":\"InvalidClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InvalidCommitmentAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"}],\"name\":\"InvalidSettlementStrategy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"asyncOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"}],\"name\":\"OrderAlreadySettled\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountEscrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structAsyncOrderClaim.Data\",\"name\":\"asyncOrderClaim\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountProvided\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"OrderCommitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountProvided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"commitOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountEscrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrderClaim.Data\",\"name\":\"asyncOrderClaim\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"}],\"name\":\"getAsyncOrderClaim\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountEscrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrderClaim.Data\",\"name\":\"asyncOrderClaim\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"}],\"name\":\"InvalidSettlementStrategy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerificationResponse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"MinimumSettlementAmountNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"OutsideSettlementWindow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deviation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tolerance\",\"type\":\"uint256\"}],\"name\":\"PriceDeviationToleranceExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"}],\"name\":\"SettlementStrategyNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalOrderAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"}],\"name\":\"OrderSettled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"}],\"name\":\"settleOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"finalOrderAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"settlePythOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"finalOrderAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"SettlementStrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SettlementStrategyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumUsdExchangeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRoundingLoss\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"}],\"name\":\"addSettlementStrategy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"getSettlementStrategy\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumUsdExchangeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRoundingLoss\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"settlementStrategy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setSettlementStrategyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"InvalidCollateralType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxWrappableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToWrap\",\"type\":\"uint256\"}],\"name\":\"WrapperExceedsMaxAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUnwrapped\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesCollected\",\"type\":\"uint256\"}],\"name\":\"SynthUnwrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWrapped\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesCollected\",\"type\":\"uint256\"}],\"name\":\"SynthWrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wrapCollateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxWrappableAmount\",\"type\":\"uint256\"}],\"name\":\"WrapperSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"wrapCollateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxWrappableAmount\",\"type\":\"uint256\"}],\"name\":\"setWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"unwrapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnCollateralAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"wrapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToMint\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"InvalidCollateralLeverage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidFeeCollector\",\"type\":\"address\"}],\"name\":\"InvalidFeeCollectorInterface\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWrapperFees\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"asyncFixedFee\",\"type\":\"uint256\"}],\"name\":\"AsyncFixedFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"atomicFixedFee\",\"type\":\"uint256\"}],\"name\":\"AtomicFixedFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralLeverage\",\"type\":\"uint256\"}],\"name\":\"CollateralLeverageSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"name\":\"MarketSkewScaleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"utilizationFeeRate\",\"type\":\"uint256\"}],\"name\":\"MarketUtilizationFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharePercentage\",\"type\":\"uint256\"}],\"name\":\"ReferrerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fixedFeeAmount\",\"type\":\"uint256\"}],\"name\":\"TransactorFixedFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"wrapFee\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"unwrapFee\",\"type\":\"int256\"}],\"name\":\"WrapperFeesSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getCollateralLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralLeverage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"transactor\",\"type\":\"address\"}],\"name\":\"getCustomTransactorFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fixedFeeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"atomicFixedFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"asyncFixedFee\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"wrapFee\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"unwrapFee\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketSkewScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketUtilizationFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"utilizationFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"getReferrerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharePercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"asyncFixedFee\",\"type\":\"uint256\"}],\"name\":\"setAsyncFixedFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"atomicFixedFee\",\"type\":\"uint256\"}],\"name\":\"setAtomicFixedFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralLeverage\",\"type\":\"uint256\"}],\"name\":\"setCollateralLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"transactor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fixedFeeAmount\",\"type\":\"uint256\"}],\"name\":\"setCustomTransactorFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"name\":\"setMarketSkewScale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFeeRate\",\"type\":\"uint256\"}],\"name\":\"setMarketUtilizationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"int256\",\"name\":\"wrapFee\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"unwrapFee\",\"type\":\"int256\"}],\"name\":\"setWrapperFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sharePercentage\",\"type\":\"uint256\"}],\"name\":\"updateReferrerShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SpotMarketGoerliABI is the input ABI used to generate the binding from.
// Deprecated: Use SpotMarketGoerliMetaData.ABI instead.
var SpotMarketGoerliABI = SpotMarketGoerliMetaData.ABI

// SpotMarketGoerli is an auto generated Go binding around an Ethereum contract.
type SpotMarketGoerli struct {
	SpotMarketGoerliCaller     // Read-only binding to the contract
	SpotMarketGoerliTransactor // Write-only binding to the contract
	SpotMarketGoerliFilterer   // Log filterer for contract events
}

// SpotMarketGoerliCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpotMarketGoerliCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotMarketGoerliTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpotMarketGoerliTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotMarketGoerliFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpotMarketGoerliFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotMarketGoerliSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpotMarketGoerliSession struct {
	Contract     *SpotMarketGoerli // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpotMarketGoerliCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpotMarketGoerliCallerSession struct {
	Contract *SpotMarketGoerliCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SpotMarketGoerliTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpotMarketGoerliTransactorSession struct {
	Contract     *SpotMarketGoerliTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SpotMarketGoerliRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpotMarketGoerliRaw struct {
	Contract *SpotMarketGoerli // Generic contract binding to access the raw methods on
}

// SpotMarketGoerliCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpotMarketGoerliCallerRaw struct {
	Contract *SpotMarketGoerliCaller // Generic read-only contract binding to access the raw methods on
}

// SpotMarketGoerliTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpotMarketGoerliTransactorRaw struct {
	Contract *SpotMarketGoerliTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpotMarketGoerli creates a new instance of SpotMarketGoerli, bound to a specific deployed contract.
func NewSpotMarketGoerli(address common.Address, backend bind.ContractBackend) (*SpotMarketGoerli, error) {
	contract, err := bindSpotMarketGoerli(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerli{SpotMarketGoerliCaller: SpotMarketGoerliCaller{contract: contract}, SpotMarketGoerliTransactor: SpotMarketGoerliTransactor{contract: contract}, SpotMarketGoerliFilterer: SpotMarketGoerliFilterer{contract: contract}}, nil
}

// NewSpotMarketGoerliCaller creates a new read-only instance of SpotMarketGoerli, bound to a specific deployed contract.
func NewSpotMarketGoerliCaller(address common.Address, caller bind.ContractCaller) (*SpotMarketGoerliCaller, error) {
	contract, err := bindSpotMarketGoerli(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliCaller{contract: contract}, nil
}

// NewSpotMarketGoerliTransactor creates a new write-only instance of SpotMarketGoerli, bound to a specific deployed contract.
func NewSpotMarketGoerliTransactor(address common.Address, transactor bind.ContractTransactor) (*SpotMarketGoerliTransactor, error) {
	contract, err := bindSpotMarketGoerli(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliTransactor{contract: contract}, nil
}

// NewSpotMarketGoerliFilterer creates a new log filterer instance of SpotMarketGoerli, bound to a specific deployed contract.
func NewSpotMarketGoerliFilterer(address common.Address, filterer bind.ContractFilterer) (*SpotMarketGoerliFilterer, error) {
	contract, err := bindSpotMarketGoerli(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliFilterer{contract: contract}, nil
}

// bindSpotMarketGoerli binds a generic wrapper to an already deployed contract.
func bindSpotMarketGoerli(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SpotMarketGoerliMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SpotMarketGoerli *SpotMarketGoerliRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SpotMarketGoerli.Contract.SpotMarketGoerliCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SpotMarketGoerli *SpotMarketGoerliRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SpotMarketGoerliTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SpotMarketGoerli *SpotMarketGoerliRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SpotMarketGoerliTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SpotMarketGoerli *SpotMarketGoerliCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SpotMarketGoerli.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SpotMarketGoerli *SpotMarketGoerliTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SpotMarketGoerli *SpotMarketGoerliTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_SpotMarketGoerli *SpotMarketGoerliSession) PRECISION() (*big.Int, error) {
	return _SpotMarketGoerli.Contract.PRECISION(&_SpotMarketGoerli.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) PRECISION() (*big.Int, error) {
	return _SpotMarketGoerli.Contract.PRECISION(&_SpotMarketGoerli.CallOpts)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetAssociatedSystem(opts *bind.CallOpts, id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getAssociatedSystem", id)

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
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _SpotMarketGoerli.Contract.GetAssociatedSystem(&_SpotMarketGoerli.CallOpts, id)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _SpotMarketGoerli.Contract.GetAssociatedSystem(&_SpotMarketGoerli.CallOpts, id)
}

// GetAsyncOrderClaim is a free data retrieval call binding the contract method 0x5381ce16.
//
// Solidity: function getAsyncOrderClaim(uint128 marketId, uint128 asyncOrderId) pure returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetAsyncOrderClaim(opts *bind.CallOpts, marketId *big.Int, asyncOrderId *big.Int) (AsyncOrderClaimData, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getAsyncOrderClaim", marketId, asyncOrderId)

	if err != nil {
		return *new(AsyncOrderClaimData), err
	}

	out0 := *abi.ConvertType(out[0], new(AsyncOrderClaimData)).(*AsyncOrderClaimData)

	return out0, err

}

// GetAsyncOrderClaim is a free data retrieval call binding the contract method 0x5381ce16.
//
// Solidity: function getAsyncOrderClaim(uint128 marketId, uint128 asyncOrderId) pure returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetAsyncOrderClaim(marketId *big.Int, asyncOrderId *big.Int) (AsyncOrderClaimData, error) {
	return _SpotMarketGoerli.Contract.GetAsyncOrderClaim(&_SpotMarketGoerli.CallOpts, marketId, asyncOrderId)
}

// GetAsyncOrderClaim is a free data retrieval call binding the contract method 0x5381ce16.
//
// Solidity: function getAsyncOrderClaim(uint128 marketId, uint128 asyncOrderId) pure returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetAsyncOrderClaim(marketId *big.Int, asyncOrderId *big.Int) (AsyncOrderClaimData, error) {
	return _SpotMarketGoerli.Contract.GetAsyncOrderClaim(&_SpotMarketGoerli.CallOpts, marketId, asyncOrderId)
}

// GetCollateralLeverage is a free data retrieval call binding the contract method 0xcdfaef0f.
//
// Solidity: function getCollateralLeverage(uint128 synthMarketId) view returns(uint256 collateralLeverage)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetCollateralLeverage(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getCollateralLeverage", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralLeverage is a free data retrieval call binding the contract method 0xcdfaef0f.
//
// Solidity: function getCollateralLeverage(uint128 synthMarketId) view returns(uint256 collateralLeverage)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetCollateralLeverage(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.GetCollateralLeverage(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetCollateralLeverage is a free data retrieval call binding the contract method 0xcdfaef0f.
//
// Solidity: function getCollateralLeverage(uint128 synthMarketId) view returns(uint256 collateralLeverage)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetCollateralLeverage(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.GetCollateralLeverage(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetCustomTransactorFees is a free data retrieval call binding the contract method 0x2efaa971.
//
// Solidity: function getCustomTransactorFees(uint128 synthMarketId, address transactor) view returns(uint256 fixedFeeAmount)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetCustomTransactorFees(opts *bind.CallOpts, synthMarketId *big.Int, transactor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getCustomTransactorFees", synthMarketId, transactor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCustomTransactorFees is a free data retrieval call binding the contract method 0x2efaa971.
//
// Solidity: function getCustomTransactorFees(uint128 synthMarketId, address transactor) view returns(uint256 fixedFeeAmount)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.GetCustomTransactorFees(&_SpotMarketGoerli.CallOpts, synthMarketId, transactor)
}

// GetCustomTransactorFees is a free data retrieval call binding the contract method 0x2efaa971.
//
// Solidity: function getCustomTransactorFees(uint128 synthMarketId, address transactor) view returns(uint256 fixedFeeAmount)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.GetCustomTransactorFees(&_SpotMarketGoerli.CallOpts, synthMarketId, transactor)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetDeniers(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getDeniers", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _SpotMarketGoerli.Contract.GetDeniers(&_SpotMarketGoerli.CallOpts, feature)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _SpotMarketGoerli.Contract.GetDeniers(&_SpotMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetFeatureFlagAllowAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getFeatureFlagAllowAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _SpotMarketGoerli.Contract.GetFeatureFlagAllowAll(&_SpotMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _SpotMarketGoerli.Contract.GetFeatureFlagAllowAll(&_SpotMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetFeatureFlagAllowlist(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getFeatureFlagAllowlist", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _SpotMarketGoerli.Contract.GetFeatureFlagAllowlist(&_SpotMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _SpotMarketGoerli.Contract.GetFeatureFlagAllowlist(&_SpotMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetFeatureFlagDenyAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getFeatureFlagDenyAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _SpotMarketGoerli.Contract.GetFeatureFlagDenyAll(&_SpotMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _SpotMarketGoerli.Contract.GetFeatureFlagDenyAll(&_SpotMarketGoerli.CallOpts, feature)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x5497eb23.
//
// Solidity: function getFeeCollector(uint128 synthMarketId) view returns(address feeCollector)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetFeeCollector(opts *bind.CallOpts, synthMarketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getFeeCollector", synthMarketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeCollector is a free data retrieval call binding the contract method 0x5497eb23.
//
// Solidity: function getFeeCollector(uint128 synthMarketId) view returns(address feeCollector)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetFeeCollector(synthMarketId *big.Int) (common.Address, error) {
	return _SpotMarketGoerli.Contract.GetFeeCollector(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x5497eb23.
//
// Solidity: function getFeeCollector(uint128 synthMarketId) view returns(address feeCollector)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetFeeCollector(synthMarketId *big.Int) (common.Address, error) {
	return _SpotMarketGoerli.Contract.GetFeeCollector(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetImplementation() (common.Address, error) {
	return _SpotMarketGoerli.Contract.GetImplementation(&_SpotMarketGoerli.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetImplementation() (common.Address, error) {
	return _SpotMarketGoerli.Contract.GetImplementation(&_SpotMarketGoerli.CallOpts)
}

// GetMarketFees is a free data retrieval call binding the contract method 0x32598e61.
//
// Solidity: function getMarketFees(uint128 synthMarketId) view returns(uint256 atomicFixedFee, uint256 asyncFixedFee, int256 wrapFee, int256 unwrapFee)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetMarketFees(opts *bind.CallOpts, synthMarketId *big.Int) (struct {
	AtomicFixedFee *big.Int
	AsyncFixedFee  *big.Int
	WrapFee        *big.Int
	UnwrapFee      *big.Int
}, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getMarketFees", synthMarketId)

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
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetMarketFees(synthMarketId *big.Int) (struct {
	AtomicFixedFee *big.Int
	AsyncFixedFee  *big.Int
	WrapFee        *big.Int
	UnwrapFee      *big.Int
}, error) {
	return _SpotMarketGoerli.Contract.GetMarketFees(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetMarketFees is a free data retrieval call binding the contract method 0x32598e61.
//
// Solidity: function getMarketFees(uint128 synthMarketId) view returns(uint256 atomicFixedFee, uint256 asyncFixedFee, int256 wrapFee, int256 unwrapFee)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetMarketFees(synthMarketId *big.Int) (struct {
	AtomicFixedFee *big.Int
	AsyncFixedFee  *big.Int
	WrapFee        *big.Int
	UnwrapFee      *big.Int
}, error) {
	return _SpotMarketGoerli.Contract.GetMarketFees(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetMarketOwner is a free data retrieval call binding the contract method 0xa7b8cb9f.
//
// Solidity: function getMarketOwner(uint128 synthMarketId) view returns(address marketOwner)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetMarketOwner(opts *bind.CallOpts, synthMarketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getMarketOwner", synthMarketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMarketOwner is a free data retrieval call binding the contract method 0xa7b8cb9f.
//
// Solidity: function getMarketOwner(uint128 synthMarketId) view returns(address marketOwner)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetMarketOwner(synthMarketId *big.Int) (common.Address, error) {
	return _SpotMarketGoerli.Contract.GetMarketOwner(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetMarketOwner is a free data retrieval call binding the contract method 0xa7b8cb9f.
//
// Solidity: function getMarketOwner(uint128 synthMarketId) view returns(address marketOwner)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetMarketOwner(synthMarketId *big.Int) (common.Address, error) {
	return _SpotMarketGoerli.Contract.GetMarketOwner(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetMarketSkewScale is a free data retrieval call binding the contract method 0x8d105571.
//
// Solidity: function getMarketSkewScale(uint128 synthMarketId) view returns(uint256 skewScale)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetMarketSkewScale(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getMarketSkewScale", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketSkewScale is a free data retrieval call binding the contract method 0x8d105571.
//
// Solidity: function getMarketSkewScale(uint128 synthMarketId) view returns(uint256 skewScale)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetMarketSkewScale(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.GetMarketSkewScale(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetMarketSkewScale is a free data retrieval call binding the contract method 0x8d105571.
//
// Solidity: function getMarketSkewScale(uint128 synthMarketId) view returns(uint256 skewScale)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetMarketSkewScale(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.GetMarketSkewScale(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetMarketUtilizationFees is a free data retrieval call binding the contract method 0xf375f324.
//
// Solidity: function getMarketUtilizationFees(uint128 synthMarketId) view returns(uint256 utilizationFeeRate)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetMarketUtilizationFees(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getMarketUtilizationFees", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketUtilizationFees is a free data retrieval call binding the contract method 0xf375f324.
//
// Solidity: function getMarketUtilizationFees(uint128 synthMarketId) view returns(uint256 utilizationFeeRate)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetMarketUtilizationFees(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.GetMarketUtilizationFees(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetMarketUtilizationFees is a free data retrieval call binding the contract method 0xf375f324.
//
// Solidity: function getMarketUtilizationFees(uint128 synthMarketId) view returns(uint256 utilizationFeeRate)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetMarketUtilizationFees(synthMarketId *big.Int) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.GetMarketUtilizationFees(&_SpotMarketGoerli.CallOpts, synthMarketId)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xfa4b28ed.
//
// Solidity: function getReferrerShare(uint128 synthMarketId, address referrer) view returns(uint256 sharePercentage)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetReferrerShare(opts *bind.CallOpts, synthMarketId *big.Int, referrer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getReferrerShare", synthMarketId, referrer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReferrerShare is a free data retrieval call binding the contract method 0xfa4b28ed.
//
// Solidity: function getReferrerShare(uint128 synthMarketId, address referrer) view returns(uint256 sharePercentage)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetReferrerShare(synthMarketId *big.Int, referrer common.Address) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.GetReferrerShare(&_SpotMarketGoerli.CallOpts, synthMarketId, referrer)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xfa4b28ed.
//
// Solidity: function getReferrerShare(uint128 synthMarketId, address referrer) view returns(uint256 sharePercentage)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetReferrerShare(synthMarketId *big.Int, referrer common.Address) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.GetReferrerShare(&_SpotMarketGoerli.CallOpts, synthMarketId, referrer)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) settlementStrategy)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetSettlementStrategy(opts *bind.CallOpts, marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getSettlementStrategy", marketId, strategyId)

	if err != nil {
		return *new(SettlementStrategyData), err
	}

	out0 := *abi.ConvertType(out[0], new(SettlementStrategyData)).(*SettlementStrategyData)

	return out0, err

}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) settlementStrategy)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _SpotMarketGoerli.Contract.GetSettlementStrategy(&_SpotMarketGoerli.CallOpts, marketId, strategyId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) settlementStrategy)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _SpotMarketGoerli.Contract.GetSettlementStrategy(&_SpotMarketGoerli.CallOpts, marketId, strategyId)
}

// GetSynth is a free data retrieval call binding the contract method 0x69e0365f.
//
// Solidity: function getSynth(uint128 marketId) view returns(address synthAddress)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetSynth(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getSynth", marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSynth is a free data retrieval call binding the contract method 0x69e0365f.
//
// Solidity: function getSynth(uint128 marketId) view returns(address synthAddress)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetSynth(marketId *big.Int) (common.Address, error) {
	return _SpotMarketGoerli.Contract.GetSynth(&_SpotMarketGoerli.CallOpts, marketId)
}

// GetSynth is a free data retrieval call binding the contract method 0x69e0365f.
//
// Solidity: function getSynth(uint128 marketId) view returns(address synthAddress)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetSynth(marketId *big.Int) (common.Address, error) {
	return _SpotMarketGoerli.Contract.GetSynth(&_SpotMarketGoerli.CallOpts, marketId)
}

// GetSynthImpl is a free data retrieval call binding the contract method 0x3e0c76ca.
//
// Solidity: function getSynthImpl(uint128 marketId) view returns(address implAddress)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) GetSynthImpl(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "getSynthImpl", marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSynthImpl is a free data retrieval call binding the contract method 0x3e0c76ca.
//
// Solidity: function getSynthImpl(uint128 marketId) view returns(address implAddress)
func (_SpotMarketGoerli *SpotMarketGoerliSession) GetSynthImpl(marketId *big.Int) (common.Address, error) {
	return _SpotMarketGoerli.Contract.GetSynthImpl(&_SpotMarketGoerli.CallOpts, marketId)
}

// GetSynthImpl is a free data retrieval call binding the contract method 0x3e0c76ca.
//
// Solidity: function getSynthImpl(uint128 marketId) view returns(address implAddress)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) GetSynthImpl(marketId *big.Int) (common.Address, error) {
	return _SpotMarketGoerli.Contract.GetSynthImpl(&_SpotMarketGoerli.CallOpts, marketId)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) IsFeatureAllowed(opts *bind.CallOpts, feature [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "isFeatureAllowed", feature, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_SpotMarketGoerli *SpotMarketGoerliSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _SpotMarketGoerli.Contract.IsFeatureAllowed(&_SpotMarketGoerli.CallOpts, feature, account)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _SpotMarketGoerli.Contract.IsFeatureAllowed(&_SpotMarketGoerli.CallOpts, feature, account)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 marketId) view returns(uint256 lockedAmount)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) MinimumCredit(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "minimumCredit", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 marketId) view returns(uint256 lockedAmount)
func (_SpotMarketGoerli *SpotMarketGoerliSession) MinimumCredit(marketId *big.Int) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.MinimumCredit(&_SpotMarketGoerli.CallOpts, marketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 marketId) view returns(uint256 lockedAmount)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) MinimumCredit(marketId *big.Int) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.MinimumCredit(&_SpotMarketGoerli.CallOpts, marketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 marketId) view returns(string marketName)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) Name(opts *bind.CallOpts, marketId *big.Int) (string, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "name", marketId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 marketId) view returns(string marketName)
func (_SpotMarketGoerli *SpotMarketGoerliSession) Name(marketId *big.Int) (string, error) {
	return _SpotMarketGoerli.Contract.Name(&_SpotMarketGoerli.CallOpts, marketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 marketId) view returns(string marketName)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) Name(marketId *big.Int) (string, error) {
	return _SpotMarketGoerli.Contract.Name(&_SpotMarketGoerli.CallOpts, marketId)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SpotMarketGoerli *SpotMarketGoerliSession) NominatedOwner() (common.Address, error) {
	return _SpotMarketGoerli.Contract.NominatedOwner(&_SpotMarketGoerli.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) NominatedOwner() (common.Address, error) {
	return _SpotMarketGoerli.Contract.NominatedOwner(&_SpotMarketGoerli.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpotMarketGoerli *SpotMarketGoerliSession) Owner() (common.Address, error) {
	return _SpotMarketGoerli.Contract.Owner(&_SpotMarketGoerli.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) Owner() (common.Address, error) {
	return _SpotMarketGoerli.Contract.Owner(&_SpotMarketGoerli.CallOpts)
}

// QuoteBuyExactIn is a free data retrieval call binding the contract method 0x6b5e6ae4.
//
// Solidity: function quoteBuyExactIn(uint128 marketId, uint256 usdAmount) view returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) QuoteBuyExactIn(opts *bind.CallOpts, marketId *big.Int, usdAmount *big.Int) (struct {
	SynthAmount *big.Int
	Fees        OrderFeesData
}, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "quoteBuyExactIn", marketId, usdAmount)

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
func (_SpotMarketGoerli *SpotMarketGoerliSession) QuoteBuyExactIn(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthAmount *big.Int
	Fees        OrderFeesData
}, error) {
	return _SpotMarketGoerli.Contract.QuoteBuyExactIn(&_SpotMarketGoerli.CallOpts, marketId, usdAmount)
}

// QuoteBuyExactIn is a free data retrieval call binding the contract method 0x6b5e6ae4.
//
// Solidity: function quoteBuyExactIn(uint128 marketId, uint256 usdAmount) view returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) QuoteBuyExactIn(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthAmount *big.Int
	Fees        OrderFeesData
}, error) {
	return _SpotMarketGoerli.Contract.QuoteBuyExactIn(&_SpotMarketGoerli.CallOpts, marketId, usdAmount)
}

// QuoteBuyExactOut is a free data retrieval call binding the contract method 0x1f3f7640.
//
// Solidity: function quoteBuyExactOut(uint128 marketId, uint256 synthAmount) view returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) QuoteBuyExactOut(opts *bind.CallOpts, marketId *big.Int, synthAmount *big.Int) (struct {
	UsdAmountCharged *big.Int
	Fees             OrderFeesData
}, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "quoteBuyExactOut", marketId, synthAmount)

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
func (_SpotMarketGoerli *SpotMarketGoerliSession) QuoteBuyExactOut(marketId *big.Int, synthAmount *big.Int) (struct {
	UsdAmountCharged *big.Int
	Fees             OrderFeesData
}, error) {
	return _SpotMarketGoerli.Contract.QuoteBuyExactOut(&_SpotMarketGoerli.CallOpts, marketId, synthAmount)
}

// QuoteBuyExactOut is a free data retrieval call binding the contract method 0x1f3f7640.
//
// Solidity: function quoteBuyExactOut(uint128 marketId, uint256 synthAmount) view returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) QuoteBuyExactOut(marketId *big.Int, synthAmount *big.Int) (struct {
	UsdAmountCharged *big.Int
	Fees             OrderFeesData
}, error) {
	return _SpotMarketGoerli.Contract.QuoteBuyExactOut(&_SpotMarketGoerli.CallOpts, marketId, synthAmount)
}

// QuoteSellExactIn is a free data retrieval call binding the contract method 0xc52d1730.
//
// Solidity: function quoteSellExactIn(uint128 marketId, uint256 synthAmount) view returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) QuoteSellExactIn(opts *bind.CallOpts, marketId *big.Int, synthAmount *big.Int) (struct {
	ReturnAmount *big.Int
	Fees         OrderFeesData
}, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "quoteSellExactIn", marketId, synthAmount)

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
func (_SpotMarketGoerli *SpotMarketGoerliSession) QuoteSellExactIn(marketId *big.Int, synthAmount *big.Int) (struct {
	ReturnAmount *big.Int
	Fees         OrderFeesData
}, error) {
	return _SpotMarketGoerli.Contract.QuoteSellExactIn(&_SpotMarketGoerli.CallOpts, marketId, synthAmount)
}

// QuoteSellExactIn is a free data retrieval call binding the contract method 0xc52d1730.
//
// Solidity: function quoteSellExactIn(uint128 marketId, uint256 synthAmount) view returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) QuoteSellExactIn(marketId *big.Int, synthAmount *big.Int) (struct {
	ReturnAmount *big.Int
	Fees         OrderFeesData
}, error) {
	return _SpotMarketGoerli.Contract.QuoteSellExactIn(&_SpotMarketGoerli.CallOpts, marketId, synthAmount)
}

// QuoteSellExactOut is a free data retrieval call binding the contract method 0x2c007522.
//
// Solidity: function quoteSellExactOut(uint128 marketId, uint256 usdAmount) view returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) QuoteSellExactOut(opts *bind.CallOpts, marketId *big.Int, usdAmount *big.Int) (struct {
	SynthToBurn *big.Int
	Fees        OrderFeesData
}, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "quoteSellExactOut", marketId, usdAmount)

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
func (_SpotMarketGoerli *SpotMarketGoerliSession) QuoteSellExactOut(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthToBurn *big.Int
	Fees        OrderFeesData
}, error) {
	return _SpotMarketGoerli.Contract.QuoteSellExactOut(&_SpotMarketGoerli.CallOpts, marketId, usdAmount)
}

// QuoteSellExactOut is a free data retrieval call binding the contract method 0x2c007522.
//
// Solidity: function quoteSellExactOut(uint128 marketId, uint256 usdAmount) view returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) QuoteSellExactOut(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthToBurn *big.Int
	Fees        OrderFeesData
}, error) {
	return _SpotMarketGoerli.Contract.QuoteSellExactOut(&_SpotMarketGoerli.CallOpts, marketId, usdAmount)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 marketId) view returns(uint256 reportedDebtAmount)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) ReportedDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "reportedDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 marketId) view returns(uint256 reportedDebtAmount)
func (_SpotMarketGoerli *SpotMarketGoerliSession) ReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.ReportedDebt(&_SpotMarketGoerli.CallOpts, marketId)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 marketId) view returns(uint256 reportedDebtAmount)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) ReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _SpotMarketGoerli.Contract.ReportedDebt(&_SpotMarketGoerli.CallOpts, marketId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool isSupported)
func (_SpotMarketGoerli *SpotMarketGoerliCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SpotMarketGoerli.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool isSupported)
func (_SpotMarketGoerli *SpotMarketGoerliSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SpotMarketGoerli.Contract.SupportsInterface(&_SpotMarketGoerli.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool isSupported)
func (_SpotMarketGoerli *SpotMarketGoerliCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SpotMarketGoerli.Contract.SupportsInterface(&_SpotMarketGoerli.CallOpts, interfaceId)
}

// AcceptMarketOwnership is a paid mutator transaction binding the contract method 0x1c216a0e.
//
// Solidity: function acceptMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) AcceptMarketOwnership(opts *bind.TransactOpts, synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "acceptMarketOwnership", synthMarketId)
}

// AcceptMarketOwnership is a paid mutator transaction binding the contract method 0x1c216a0e.
//
// Solidity: function acceptMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) AcceptMarketOwnership(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.AcceptMarketOwnership(&_SpotMarketGoerli.TransactOpts, synthMarketId)
}

// AcceptMarketOwnership is a paid mutator transaction binding the contract method 0x1c216a0e.
//
// Solidity: function acceptMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) AcceptMarketOwnership(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.AcceptMarketOwnership(&_SpotMarketGoerli.TransactOpts, synthMarketId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) AcceptOwnership() (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.AcceptOwnership(&_SpotMarketGoerli.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.AcceptOwnership(&_SpotMarketGoerli.TransactOpts)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x97b30e6d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) AddSettlementStrategy(opts *bind.TransactOpts, marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "addSettlementStrategy", marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x97b30e6d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_SpotMarketGoerli *SpotMarketGoerliSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.AddSettlementStrategy(&_SpotMarketGoerli.TransactOpts, marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x97b30e6d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.AddSettlementStrategy(&_SpotMarketGoerli.TransactOpts, marketId, strategy)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) AddToFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "addToFeatureFlagAllowlist", feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.AddToFeatureFlagAllowlist(&_SpotMarketGoerli.TransactOpts, feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.AddToFeatureFlagAllowlist(&_SpotMarketGoerli.TransactOpts, feature, account)
}

// Buy is a paid mutator transaction binding the contract method 0x37fb3369.
//
// Solidity: function buy(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) Buy(opts *bind.TransactOpts, marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "buy", marketId, usdAmount, minAmountReceived, referrer)
}

// Buy is a paid mutator transaction binding the contract method 0x37fb3369.
//
// Solidity: function buy(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliSession) Buy(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.Buy(&_SpotMarketGoerli.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// Buy is a paid mutator transaction binding the contract method 0x37fb3369.
//
// Solidity: function buy(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) Buy(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.Buy(&_SpotMarketGoerli.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactIn is a paid mutator transaction binding the contract method 0xa12d9400.
//
// Solidity: function buyExactIn(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) BuyExactIn(opts *bind.TransactOpts, marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "buyExactIn", marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactIn is a paid mutator transaction binding the contract method 0xa12d9400.
//
// Solidity: function buyExactIn(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliSession) BuyExactIn(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.BuyExactIn(&_SpotMarketGoerli.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactIn is a paid mutator transaction binding the contract method 0xa12d9400.
//
// Solidity: function buyExactIn(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) BuyExactIn(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.BuyExactIn(&_SpotMarketGoerli.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactOut is a paid mutator transaction binding the contract method 0x983220bb.
//
// Solidity: function buyExactOut(uint128 marketId, uint256 synthAmount, uint256 maxUsdAmount, address referrer) returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) BuyExactOut(opts *bind.TransactOpts, marketId *big.Int, synthAmount *big.Int, maxUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "buyExactOut", marketId, synthAmount, maxUsdAmount, referrer)
}

// BuyExactOut is a paid mutator transaction binding the contract method 0x983220bb.
//
// Solidity: function buyExactOut(uint128 marketId, uint256 synthAmount, uint256 maxUsdAmount, address referrer) returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliSession) BuyExactOut(marketId *big.Int, synthAmount *big.Int, maxUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.BuyExactOut(&_SpotMarketGoerli.TransactOpts, marketId, synthAmount, maxUsdAmount, referrer)
}

// BuyExactOut is a paid mutator transaction binding the contract method 0x983220bb.
//
// Solidity: function buyExactOut(uint128 marketId, uint256 synthAmount, uint256 maxUsdAmount, address referrer) returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) BuyExactOut(marketId *big.Int, synthAmount *big.Int, maxUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.BuyExactOut(&_SpotMarketGoerli.TransactOpts, marketId, synthAmount, maxUsdAmount, referrer)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x4138dc53.
//
// Solidity: function cancelOrder(uint128 marketId, uint128 asyncOrderId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) CancelOrder(opts *bind.TransactOpts, marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "cancelOrder", marketId, asyncOrderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x4138dc53.
//
// Solidity: function cancelOrder(uint128 marketId, uint128 asyncOrderId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) CancelOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.CancelOrder(&_SpotMarketGoerli.TransactOpts, marketId, asyncOrderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x4138dc53.
//
// Solidity: function cancelOrder(uint128 marketId, uint128 asyncOrderId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) CancelOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.CancelOrder(&_SpotMarketGoerli.TransactOpts, marketId, asyncOrderId)
}

// CommitOrder is a paid mutator transaction binding the contract method 0xd393340e.
//
// Solidity: function commitOrder(uint128 marketId, uint8 orderType, uint256 amountProvided, uint256 settlementStrategyId, uint256 minimumSettlementAmount, address referrer) returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) CommitOrder(opts *bind.TransactOpts, marketId *big.Int, orderType uint8, amountProvided *big.Int, settlementStrategyId *big.Int, minimumSettlementAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "commitOrder", marketId, orderType, amountProvided, settlementStrategyId, minimumSettlementAmount, referrer)
}

// CommitOrder is a paid mutator transaction binding the contract method 0xd393340e.
//
// Solidity: function commitOrder(uint128 marketId, uint8 orderType, uint256 amountProvided, uint256 settlementStrategyId, uint256 minimumSettlementAmount, address referrer) returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketGoerli *SpotMarketGoerliSession) CommitOrder(marketId *big.Int, orderType uint8, amountProvided *big.Int, settlementStrategyId *big.Int, minimumSettlementAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.CommitOrder(&_SpotMarketGoerli.TransactOpts, marketId, orderType, amountProvided, settlementStrategyId, minimumSettlementAmount, referrer)
}

// CommitOrder is a paid mutator transaction binding the contract method 0xd393340e.
//
// Solidity: function commitOrder(uint128 marketId, uint8 orderType, uint256 amountProvided, uint256 settlementStrategyId, uint256 minimumSettlementAmount, address referrer) returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) CommitOrder(marketId *big.Int, orderType uint8, amountProvided *big.Int, settlementStrategyId *big.Int, minimumSettlementAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.CommitOrder(&_SpotMarketGoerli.TransactOpts, marketId, orderType, amountProvided, settlementStrategyId, minimumSettlementAmount, referrer)
}

// CreateSynth is a paid mutator transaction binding the contract method 0x2e535d61.
//
// Solidity: function createSynth(string tokenName, string tokenSymbol, address synthOwner) returns(uint128 synthMarketId)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) CreateSynth(opts *bind.TransactOpts, tokenName string, tokenSymbol string, synthOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "createSynth", tokenName, tokenSymbol, synthOwner)
}

// CreateSynth is a paid mutator transaction binding the contract method 0x2e535d61.
//
// Solidity: function createSynth(string tokenName, string tokenSymbol, address synthOwner) returns(uint128 synthMarketId)
func (_SpotMarketGoerli *SpotMarketGoerliSession) CreateSynth(tokenName string, tokenSymbol string, synthOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.CreateSynth(&_SpotMarketGoerli.TransactOpts, tokenName, tokenSymbol, synthOwner)
}

// CreateSynth is a paid mutator transaction binding the contract method 0x2e535d61.
//
// Solidity: function createSynth(string tokenName, string tokenSymbol, address synthOwner) returns(uint128 synthMarketId)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) CreateSynth(tokenName string, tokenSymbol string, synthOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.CreateSynth(&_SpotMarketGoerli.TransactOpts, tokenName, tokenSymbol, synthOwner)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) InitOrUpgradeNft(opts *bind.TransactOpts, id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "initOrUpgradeNft", id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.InitOrUpgradeNft(&_SpotMarketGoerli.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.InitOrUpgradeNft(&_SpotMarketGoerli.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) InitOrUpgradeToken(opts *bind.TransactOpts, id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "initOrUpgradeToken", id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.InitOrUpgradeToken(&_SpotMarketGoerli.TransactOpts, id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.InitOrUpgradeToken(&_SpotMarketGoerli.TransactOpts, id, name, symbol, decimals, impl)
}

// NominateMarketOwner is a paid mutator transaction binding the contract method 0x5950864b.
//
// Solidity: function nominateMarketOwner(uint128 synthMarketId, address newNominatedOwner) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) NominateMarketOwner(opts *bind.TransactOpts, synthMarketId *big.Int, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "nominateMarketOwner", synthMarketId, newNominatedOwner)
}

// NominateMarketOwner is a paid mutator transaction binding the contract method 0x5950864b.
//
// Solidity: function nominateMarketOwner(uint128 synthMarketId, address newNominatedOwner) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) NominateMarketOwner(synthMarketId *big.Int, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.NominateMarketOwner(&_SpotMarketGoerli.TransactOpts, synthMarketId, newNominatedOwner)
}

// NominateMarketOwner is a paid mutator transaction binding the contract method 0x5950864b.
//
// Solidity: function nominateMarketOwner(uint128 synthMarketId, address newNominatedOwner) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) NominateMarketOwner(synthMarketId *big.Int, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.NominateMarketOwner(&_SpotMarketGoerli.TransactOpts, synthMarketId, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) NominateNewOwner(opts *bind.TransactOpts, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "nominateNewOwner", newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.NominateNewOwner(&_SpotMarketGoerli.TransactOpts, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.NominateNewOwner(&_SpotMarketGoerli.TransactOpts, newNominatedOwner)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) RegisterUnmanagedSystem(opts *bind.TransactOpts, id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "registerUnmanagedSystem", id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.RegisterUnmanagedSystem(&_SpotMarketGoerli.TransactOpts, id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.RegisterUnmanagedSystem(&_SpotMarketGoerli.TransactOpts, id, endpoint)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) RemoveFromFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "removeFromFeatureFlagAllowlist", feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.RemoveFromFeatureFlagAllowlist(&_SpotMarketGoerli.TransactOpts, feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.RemoveFromFeatureFlagAllowlist(&_SpotMarketGoerli.TransactOpts, feature, account)
}

// RenounceMarketNomination is a paid mutator transaction binding the contract method 0x298b26bf.
//
// Solidity: function renounceMarketNomination(uint128 synthMarketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) RenounceMarketNomination(opts *bind.TransactOpts, synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "renounceMarketNomination", synthMarketId)
}

// RenounceMarketNomination is a paid mutator transaction binding the contract method 0x298b26bf.
//
// Solidity: function renounceMarketNomination(uint128 synthMarketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) RenounceMarketNomination(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.RenounceMarketNomination(&_SpotMarketGoerli.TransactOpts, synthMarketId)
}

// RenounceMarketNomination is a paid mutator transaction binding the contract method 0x298b26bf.
//
// Solidity: function renounceMarketNomination(uint128 synthMarketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) RenounceMarketNomination(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.RenounceMarketNomination(&_SpotMarketGoerli.TransactOpts, synthMarketId)
}

// RenounceMarketOwnership is a paid mutator transaction binding the contract method 0xbd1cdfb5.
//
// Solidity: function renounceMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) RenounceMarketOwnership(opts *bind.TransactOpts, synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "renounceMarketOwnership", synthMarketId)
}

// RenounceMarketOwnership is a paid mutator transaction binding the contract method 0xbd1cdfb5.
//
// Solidity: function renounceMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) RenounceMarketOwnership(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.RenounceMarketOwnership(&_SpotMarketGoerli.TransactOpts, synthMarketId)
}

// RenounceMarketOwnership is a paid mutator transaction binding the contract method 0xbd1cdfb5.
//
// Solidity: function renounceMarketOwnership(uint128 synthMarketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) RenounceMarketOwnership(synthMarketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.RenounceMarketOwnership(&_SpotMarketGoerli.TransactOpts, synthMarketId)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) RenounceNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "renounceNomination")
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) RenounceNomination() (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.RenounceNomination(&_SpotMarketGoerli.TransactOpts)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) RenounceNomination() (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.RenounceNomination(&_SpotMarketGoerli.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0x4d4bfbd5.
//
// Solidity: function sell(uint128 marketId, uint256 synthAmount, uint256 minUsdAmount, address referrer) returns(uint256 usdAmountReceived, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) Sell(opts *bind.TransactOpts, marketId *big.Int, synthAmount *big.Int, minUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "sell", marketId, synthAmount, minUsdAmount, referrer)
}

// Sell is a paid mutator transaction binding the contract method 0x4d4bfbd5.
//
// Solidity: function sell(uint128 marketId, uint256 synthAmount, uint256 minUsdAmount, address referrer) returns(uint256 usdAmountReceived, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliSession) Sell(marketId *big.Int, synthAmount *big.Int, minUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.Sell(&_SpotMarketGoerli.TransactOpts, marketId, synthAmount, minUsdAmount, referrer)
}

// Sell is a paid mutator transaction binding the contract method 0x4d4bfbd5.
//
// Solidity: function sell(uint128 marketId, uint256 synthAmount, uint256 minUsdAmount, address referrer) returns(uint256 usdAmountReceived, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) Sell(marketId *big.Int, synthAmount *big.Int, minUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.Sell(&_SpotMarketGoerli.TransactOpts, marketId, synthAmount, minUsdAmount, referrer)
}

// SellExactIn is a paid mutator transaction binding the contract method 0x3d1a60e4.
//
// Solidity: function sellExactIn(uint128 marketId, uint256 synthAmount, uint256 minAmountReceived, address referrer) returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SellExactIn(opts *bind.TransactOpts, marketId *big.Int, synthAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "sellExactIn", marketId, synthAmount, minAmountReceived, referrer)
}

// SellExactIn is a paid mutator transaction binding the contract method 0x3d1a60e4.
//
// Solidity: function sellExactIn(uint128 marketId, uint256 synthAmount, uint256 minAmountReceived, address referrer) returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliSession) SellExactIn(marketId *big.Int, synthAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SellExactIn(&_SpotMarketGoerli.TransactOpts, marketId, synthAmount, minAmountReceived, referrer)
}

// SellExactIn is a paid mutator transaction binding the contract method 0x3d1a60e4.
//
// Solidity: function sellExactIn(uint128 marketId, uint256 synthAmount, uint256 minAmountReceived, address referrer) returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SellExactIn(marketId *big.Int, synthAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SellExactIn(&_SpotMarketGoerli.TransactOpts, marketId, synthAmount, minAmountReceived, referrer)
}

// SellExactOut is a paid mutator transaction binding the contract method 0x4ce94d9d.
//
// Solidity: function sellExactOut(uint128 marketId, uint256 usdAmount, uint256 maxSynthAmount, address referrer) returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SellExactOut(opts *bind.TransactOpts, marketId *big.Int, usdAmount *big.Int, maxSynthAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "sellExactOut", marketId, usdAmount, maxSynthAmount, referrer)
}

// SellExactOut is a paid mutator transaction binding the contract method 0x4ce94d9d.
//
// Solidity: function sellExactOut(uint128 marketId, uint256 usdAmount, uint256 maxSynthAmount, address referrer) returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliSession) SellExactOut(marketId *big.Int, usdAmount *big.Int, maxSynthAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SellExactOut(&_SpotMarketGoerli.TransactOpts, marketId, usdAmount, maxSynthAmount, referrer)
}

// SellExactOut is a paid mutator transaction binding the contract method 0x4ce94d9d.
//
// Solidity: function sellExactOut(uint128 marketId, uint256 usdAmount, uint256 maxSynthAmount, address referrer) returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SellExactOut(marketId *big.Int, usdAmount *big.Int, maxSynthAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SellExactOut(&_SpotMarketGoerli.TransactOpts, marketId, usdAmount, maxSynthAmount, referrer)
}

// SetAsyncFixedFee is a paid mutator transaction binding the contract method 0x61dcca86.
//
// Solidity: function setAsyncFixedFee(uint128 synthMarketId, uint256 asyncFixedFee) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetAsyncFixedFee(opts *bind.TransactOpts, synthMarketId *big.Int, asyncFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setAsyncFixedFee", synthMarketId, asyncFixedFee)
}

// SetAsyncFixedFee is a paid mutator transaction binding the contract method 0x61dcca86.
//
// Solidity: function setAsyncFixedFee(uint128 synthMarketId, uint256 asyncFixedFee) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetAsyncFixedFee(synthMarketId *big.Int, asyncFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetAsyncFixedFee(&_SpotMarketGoerli.TransactOpts, synthMarketId, asyncFixedFee)
}

// SetAsyncFixedFee is a paid mutator transaction binding the contract method 0x61dcca86.
//
// Solidity: function setAsyncFixedFee(uint128 synthMarketId, uint256 asyncFixedFee) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetAsyncFixedFee(synthMarketId *big.Int, asyncFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetAsyncFixedFee(&_SpotMarketGoerli.TransactOpts, synthMarketId, asyncFixedFee)
}

// SetAtomicFixedFee is a paid mutator transaction binding the contract method 0x480be91f.
//
// Solidity: function setAtomicFixedFee(uint128 synthMarketId, uint256 atomicFixedFee) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetAtomicFixedFee(opts *bind.TransactOpts, synthMarketId *big.Int, atomicFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setAtomicFixedFee", synthMarketId, atomicFixedFee)
}

// SetAtomicFixedFee is a paid mutator transaction binding the contract method 0x480be91f.
//
// Solidity: function setAtomicFixedFee(uint128 synthMarketId, uint256 atomicFixedFee) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetAtomicFixedFee(synthMarketId *big.Int, atomicFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetAtomicFixedFee(&_SpotMarketGoerli.TransactOpts, synthMarketId, atomicFixedFee)
}

// SetAtomicFixedFee is a paid mutator transaction binding the contract method 0x480be91f.
//
// Solidity: function setAtomicFixedFee(uint128 synthMarketId, uint256 atomicFixedFee) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetAtomicFixedFee(synthMarketId *big.Int, atomicFixedFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetAtomicFixedFee(&_SpotMarketGoerli.TransactOpts, synthMarketId, atomicFixedFee)
}

// SetCollateralLeverage is a paid mutator transaction binding the contract method 0x21f7f58f.
//
// Solidity: function setCollateralLeverage(uint128 synthMarketId, uint256 collateralLeverage) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetCollateralLeverage(opts *bind.TransactOpts, synthMarketId *big.Int, collateralLeverage *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setCollateralLeverage", synthMarketId, collateralLeverage)
}

// SetCollateralLeverage is a paid mutator transaction binding the contract method 0x21f7f58f.
//
// Solidity: function setCollateralLeverage(uint128 synthMarketId, uint256 collateralLeverage) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetCollateralLeverage(synthMarketId *big.Int, collateralLeverage *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetCollateralLeverage(&_SpotMarketGoerli.TransactOpts, synthMarketId, collateralLeverage)
}

// SetCollateralLeverage is a paid mutator transaction binding the contract method 0x21f7f58f.
//
// Solidity: function setCollateralLeverage(uint128 synthMarketId, uint256 collateralLeverage) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetCollateralLeverage(synthMarketId *big.Int, collateralLeverage *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetCollateralLeverage(&_SpotMarketGoerli.TransactOpts, synthMarketId, collateralLeverage)
}

// SetCustomTransactorFees is a paid mutator transaction binding the contract method 0x95fcd547.
//
// Solidity: function setCustomTransactorFees(uint128 synthMarketId, address transactor, uint256 fixedFeeAmount) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetCustomTransactorFees(opts *bind.TransactOpts, synthMarketId *big.Int, transactor common.Address, fixedFeeAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setCustomTransactorFees", synthMarketId, transactor, fixedFeeAmount)
}

// SetCustomTransactorFees is a paid mutator transaction binding the contract method 0x95fcd547.
//
// Solidity: function setCustomTransactorFees(uint128 synthMarketId, address transactor, uint256 fixedFeeAmount) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address, fixedFeeAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetCustomTransactorFees(&_SpotMarketGoerli.TransactOpts, synthMarketId, transactor, fixedFeeAmount)
}

// SetCustomTransactorFees is a paid mutator transaction binding the contract method 0x95fcd547.
//
// Solidity: function setCustomTransactorFees(uint128 synthMarketId, address transactor, uint256 fixedFeeAmount) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address, fixedFeeAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetCustomTransactorFees(&_SpotMarketGoerli.TransactOpts, synthMarketId, transactor, fixedFeeAmount)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xec846bac.
//
// Solidity: function setDecayRate(uint128 marketId, uint256 rate) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetDecayRate(opts *bind.TransactOpts, marketId *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setDecayRate", marketId, rate)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xec846bac.
//
// Solidity: function setDecayRate(uint128 marketId, uint256 rate) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetDecayRate(marketId *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetDecayRate(&_SpotMarketGoerli.TransactOpts, marketId, rate)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xec846bac.
//
// Solidity: function setDecayRate(uint128 marketId, uint256 rate) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetDecayRate(marketId *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetDecayRate(&_SpotMarketGoerli.TransactOpts, marketId, rate)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetDeniers(opts *bind.TransactOpts, feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setDeniers", feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetDeniers(&_SpotMarketGoerli.TransactOpts, feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetDeniers(&_SpotMarketGoerli.TransactOpts, feature, deniers)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetFeatureFlagAllowAll(opts *bind.TransactOpts, feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setFeatureFlagAllowAll", feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetFeatureFlagAllowAll(&_SpotMarketGoerli.TransactOpts, feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetFeatureFlagAllowAll(&_SpotMarketGoerli.TransactOpts, feature, allowAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetFeatureFlagDenyAll(opts *bind.TransactOpts, feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setFeatureFlagDenyAll", feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetFeatureFlagDenyAll(&_SpotMarketGoerli.TransactOpts, feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetFeatureFlagDenyAll(&_SpotMarketGoerli.TransactOpts, feature, denyAll)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x025f6120.
//
// Solidity: function setFeeCollector(uint128 synthMarketId, address feeCollector) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetFeeCollector(opts *bind.TransactOpts, synthMarketId *big.Int, feeCollector common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setFeeCollector", synthMarketId, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x025f6120.
//
// Solidity: function setFeeCollector(uint128 synthMarketId, address feeCollector) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetFeeCollector(synthMarketId *big.Int, feeCollector common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetFeeCollector(&_SpotMarketGoerli.TransactOpts, synthMarketId, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x025f6120.
//
// Solidity: function setFeeCollector(uint128 synthMarketId, address feeCollector) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetFeeCollector(synthMarketId *big.Int, feeCollector common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetFeeCollector(&_SpotMarketGoerli.TransactOpts, synthMarketId, feeCollector)
}

// SetMarketSkewScale is a paid mutator transaction binding the contract method 0x9a40f8cb.
//
// Solidity: function setMarketSkewScale(uint128 synthMarketId, uint256 skewScale) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetMarketSkewScale(opts *bind.TransactOpts, synthMarketId *big.Int, skewScale *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setMarketSkewScale", synthMarketId, skewScale)
}

// SetMarketSkewScale is a paid mutator transaction binding the contract method 0x9a40f8cb.
//
// Solidity: function setMarketSkewScale(uint128 synthMarketId, uint256 skewScale) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetMarketSkewScale(synthMarketId *big.Int, skewScale *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetMarketSkewScale(&_SpotMarketGoerli.TransactOpts, synthMarketId, skewScale)
}

// SetMarketSkewScale is a paid mutator transaction binding the contract method 0x9a40f8cb.
//
// Solidity: function setMarketSkewScale(uint128 synthMarketId, uint256 skewScale) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetMarketSkewScale(synthMarketId *big.Int, skewScale *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetMarketSkewScale(&_SpotMarketGoerli.TransactOpts, synthMarketId, skewScale)
}

// SetMarketUtilizationFees is a paid mutator transaction binding the contract method 0x45f2601c.
//
// Solidity: function setMarketUtilizationFees(uint128 synthMarketId, uint256 utilizationFeeRate) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetMarketUtilizationFees(opts *bind.TransactOpts, synthMarketId *big.Int, utilizationFeeRate *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setMarketUtilizationFees", synthMarketId, utilizationFeeRate)
}

// SetMarketUtilizationFees is a paid mutator transaction binding the contract method 0x45f2601c.
//
// Solidity: function setMarketUtilizationFees(uint128 synthMarketId, uint256 utilizationFeeRate) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetMarketUtilizationFees(synthMarketId *big.Int, utilizationFeeRate *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetMarketUtilizationFees(&_SpotMarketGoerli.TransactOpts, synthMarketId, utilizationFeeRate)
}

// SetMarketUtilizationFees is a paid mutator transaction binding the contract method 0x45f2601c.
//
// Solidity: function setMarketUtilizationFees(uint128 synthMarketId, uint256 utilizationFeeRate) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetMarketUtilizationFees(synthMarketId *big.Int, utilizationFeeRate *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetMarketUtilizationFees(&_SpotMarketGoerli.TransactOpts, synthMarketId, utilizationFeeRate)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetSettlementStrategyEnabled(opts *bind.TransactOpts, marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setSettlementStrategyEnabled", marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetSettlementStrategyEnabled(&_SpotMarketGoerli.TransactOpts, marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetSettlementStrategyEnabled(&_SpotMarketGoerli.TransactOpts, marketId, strategyId, enabled)
}

// SetSynthImplementation is a paid mutator transaction binding the contract method 0xec04ceb1.
//
// Solidity: function setSynthImplementation(address synthImplementation) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetSynthImplementation(opts *bind.TransactOpts, synthImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setSynthImplementation", synthImplementation)
}

// SetSynthImplementation is a paid mutator transaction binding the contract method 0xec04ceb1.
//
// Solidity: function setSynthImplementation(address synthImplementation) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetSynthImplementation(synthImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetSynthImplementation(&_SpotMarketGoerli.TransactOpts, synthImplementation)
}

// SetSynthImplementation is a paid mutator transaction binding the contract method 0xec04ceb1.
//
// Solidity: function setSynthImplementation(address synthImplementation) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetSynthImplementation(synthImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetSynthImplementation(&_SpotMarketGoerli.TransactOpts, synthImplementation)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetSynthetix(opts *bind.TransactOpts, synthetix common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setSynthetix", synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetSynthetix(&_SpotMarketGoerli.TransactOpts, synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetSynthetix(&_SpotMarketGoerli.TransactOpts, synthetix)
}

// SetWrapper is a paid mutator transaction binding the contract method 0x673a21e5.
//
// Solidity: function setWrapper(uint128 marketId, address wrapCollateralType, uint256 maxWrappableAmount) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetWrapper(opts *bind.TransactOpts, marketId *big.Int, wrapCollateralType common.Address, maxWrappableAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setWrapper", marketId, wrapCollateralType, maxWrappableAmount)
}

// SetWrapper is a paid mutator transaction binding the contract method 0x673a21e5.
//
// Solidity: function setWrapper(uint128 marketId, address wrapCollateralType, uint256 maxWrappableAmount) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetWrapper(marketId *big.Int, wrapCollateralType common.Address, maxWrappableAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetWrapper(&_SpotMarketGoerli.TransactOpts, marketId, wrapCollateralType, maxWrappableAmount)
}

// SetWrapper is a paid mutator transaction binding the contract method 0x673a21e5.
//
// Solidity: function setWrapper(uint128 marketId, address wrapCollateralType, uint256 maxWrappableAmount) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetWrapper(marketId *big.Int, wrapCollateralType common.Address, maxWrappableAmount *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetWrapper(&_SpotMarketGoerli.TransactOpts, marketId, wrapCollateralType, maxWrappableAmount)
}

// SetWrapperFees is a paid mutator transaction binding the contract method 0x6539b1c3.
//
// Solidity: function setWrapperFees(uint128 synthMarketId, int256 wrapFee, int256 unwrapFee) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SetWrapperFees(opts *bind.TransactOpts, synthMarketId *big.Int, wrapFee *big.Int, unwrapFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "setWrapperFees", synthMarketId, wrapFee, unwrapFee)
}

// SetWrapperFees is a paid mutator transaction binding the contract method 0x6539b1c3.
//
// Solidity: function setWrapperFees(uint128 synthMarketId, int256 wrapFee, int256 unwrapFee) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SetWrapperFees(synthMarketId *big.Int, wrapFee *big.Int, unwrapFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetWrapperFees(&_SpotMarketGoerli.TransactOpts, synthMarketId, wrapFee, unwrapFee)
}

// SetWrapperFees is a paid mutator transaction binding the contract method 0x6539b1c3.
//
// Solidity: function setWrapperFees(uint128 synthMarketId, int256 wrapFee, int256 unwrapFee) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SetWrapperFees(synthMarketId *big.Int, wrapFee *big.Int, unwrapFee *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SetWrapperFees(&_SpotMarketGoerli.TransactOpts, synthMarketId, wrapFee, unwrapFee)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x9444ac48.
//
// Solidity: function settleOrder(uint128 marketId, uint128 asyncOrderId) returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SettleOrder(opts *bind.TransactOpts, marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "settleOrder", marketId, asyncOrderId)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x9444ac48.
//
// Solidity: function settleOrder(uint128 marketId, uint128 asyncOrderId) returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliSession) SettleOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SettleOrder(&_SpotMarketGoerli.TransactOpts, marketId, asyncOrderId)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x9444ac48.
//
// Solidity: function settleOrder(uint128 marketId, uint128 asyncOrderId) returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SettleOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SettleOrder(&_SpotMarketGoerli.TransactOpts, marketId, asyncOrderId)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SettlePythOrder(opts *bind.TransactOpts, result []byte, extraData []byte) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "settlePythOrder", result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SettlePythOrder(&_SpotMarketGoerli.TransactOpts, result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SettlePythOrder(&_SpotMarketGoerli.TransactOpts, result, extraData)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) SimulateUpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "simulateUpgradeTo", newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SimulateUpgradeTo(&_SpotMarketGoerli.TransactOpts, newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.SimulateUpgradeTo(&_SpotMarketGoerli.TransactOpts, newImplementation)
}

// Unwrap is a paid mutator transaction binding the contract method 0x784dad9e.
//
// Solidity: function unwrap(uint128 marketId, uint256 unwrapAmount, uint256 minAmountReceived) returns(uint256 returnCollateralAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) Unwrap(opts *bind.TransactOpts, marketId *big.Int, unwrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "unwrap", marketId, unwrapAmount, minAmountReceived)
}

// Unwrap is a paid mutator transaction binding the contract method 0x784dad9e.
//
// Solidity: function unwrap(uint128 marketId, uint256 unwrapAmount, uint256 minAmountReceived) returns(uint256 returnCollateralAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliSession) Unwrap(marketId *big.Int, unwrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.Unwrap(&_SpotMarketGoerli.TransactOpts, marketId, unwrapAmount, minAmountReceived)
}

// Unwrap is a paid mutator transaction binding the contract method 0x784dad9e.
//
// Solidity: function unwrap(uint128 marketId, uint256 unwrapAmount, uint256 minAmountReceived) returns(uint256 returnCollateralAmount, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) Unwrap(marketId *big.Int, unwrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.Unwrap(&_SpotMarketGoerli.TransactOpts, marketId, unwrapAmount, minAmountReceived)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x6a0e2085.
//
// Solidity: function updatePriceData(uint128 synthMarketId, bytes32 buyFeedId, bytes32 sellFeedId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) UpdatePriceData(opts *bind.TransactOpts, synthMarketId *big.Int, buyFeedId [32]byte, sellFeedId [32]byte) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "updatePriceData", synthMarketId, buyFeedId, sellFeedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x6a0e2085.
//
// Solidity: function updatePriceData(uint128 synthMarketId, bytes32 buyFeedId, bytes32 sellFeedId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) UpdatePriceData(synthMarketId *big.Int, buyFeedId [32]byte, sellFeedId [32]byte) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.UpdatePriceData(&_SpotMarketGoerli.TransactOpts, synthMarketId, buyFeedId, sellFeedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x6a0e2085.
//
// Solidity: function updatePriceData(uint128 synthMarketId, bytes32 buyFeedId, bytes32 sellFeedId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) UpdatePriceData(synthMarketId *big.Int, buyFeedId [32]byte, sellFeedId [32]byte) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.UpdatePriceData(&_SpotMarketGoerli.TransactOpts, synthMarketId, buyFeedId, sellFeedId)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6ad77077.
//
// Solidity: function updateReferrerShare(uint128 synthMarketId, address referrer, uint256 sharePercentage) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) UpdateReferrerShare(opts *bind.TransactOpts, synthMarketId *big.Int, referrer common.Address, sharePercentage *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "updateReferrerShare", synthMarketId, referrer, sharePercentage)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6ad77077.
//
// Solidity: function updateReferrerShare(uint128 synthMarketId, address referrer, uint256 sharePercentage) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) UpdateReferrerShare(synthMarketId *big.Int, referrer common.Address, sharePercentage *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.UpdateReferrerShare(&_SpotMarketGoerli.TransactOpts, synthMarketId, referrer, sharePercentage)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6ad77077.
//
// Solidity: function updateReferrerShare(uint128 synthMarketId, address referrer, uint256 sharePercentage) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) UpdateReferrerShare(synthMarketId *big.Int, referrer common.Address, sharePercentage *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.UpdateReferrerShare(&_SpotMarketGoerli.TransactOpts, synthMarketId, referrer, sharePercentage)
}

// UpgradeSynthImpl is a paid mutator transaction binding the contract method 0xc99d0cdd.
//
// Solidity: function upgradeSynthImpl(uint128 marketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) UpgradeSynthImpl(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "upgradeSynthImpl", marketId)
}

// UpgradeSynthImpl is a paid mutator transaction binding the contract method 0xc99d0cdd.
//
// Solidity: function upgradeSynthImpl(uint128 marketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) UpgradeSynthImpl(marketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.UpgradeSynthImpl(&_SpotMarketGoerli.TransactOpts, marketId)
}

// UpgradeSynthImpl is a paid mutator transaction binding the contract method 0xc99d0cdd.
//
// Solidity: function upgradeSynthImpl(uint128 marketId) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) UpgradeSynthImpl(marketId *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.UpgradeSynthImpl(&_SpotMarketGoerli.TransactOpts, marketId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SpotMarketGoerli *SpotMarketGoerliSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.UpgradeTo(&_SpotMarketGoerli.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.UpgradeTo(&_SpotMarketGoerli.TransactOpts, newImplementation)
}

// Wrap is a paid mutator transaction binding the contract method 0xd7ce770c.
//
// Solidity: function wrap(uint128 marketId, uint256 wrapAmount, uint256 minAmountReceived) returns(uint256 amountToMint, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactor) Wrap(opts *bind.TransactOpts, marketId *big.Int, wrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.contract.Transact(opts, "wrap", marketId, wrapAmount, minAmountReceived)
}

// Wrap is a paid mutator transaction binding the contract method 0xd7ce770c.
//
// Solidity: function wrap(uint128 marketId, uint256 wrapAmount, uint256 minAmountReceived) returns(uint256 amountToMint, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliSession) Wrap(marketId *big.Int, wrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.Wrap(&_SpotMarketGoerli.TransactOpts, marketId, wrapAmount, minAmountReceived)
}

// Wrap is a paid mutator transaction binding the contract method 0xd7ce770c.
//
// Solidity: function wrap(uint128 marketId, uint256 wrapAmount, uint256 minAmountReceived) returns(uint256 amountToMint, (uint256,uint256,int256,int256) fees)
func (_SpotMarketGoerli *SpotMarketGoerliTransactorSession) Wrap(marketId *big.Int, wrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _SpotMarketGoerli.Contract.Wrap(&_SpotMarketGoerli.TransactOpts, marketId, wrapAmount, minAmountReceived)
}

// SpotMarketGoerliAssociatedSystemSetIterator is returned from FilterAssociatedSystemSet and is used to iterate over the raw logs and unpacked data for AssociatedSystemSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliAssociatedSystemSetIterator struct {
	Event *SpotMarketGoerliAssociatedSystemSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliAssociatedSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliAssociatedSystemSet)
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
		it.Event = new(SpotMarketGoerliAssociatedSystemSet)
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
func (it *SpotMarketGoerliAssociatedSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliAssociatedSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliAssociatedSystemSet represents a AssociatedSystemSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliAssociatedSystemSet struct {
	Kind  [32]byte
	Id    [32]byte
	Proxy common.Address
	Impl  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssociatedSystemSet is a free log retrieval operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterAssociatedSystemSet(opts *bind.FilterOpts, kind [][32]byte, id [][32]byte) (*SpotMarketGoerliAssociatedSystemSetIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliAssociatedSystemSetIterator{contract: _SpotMarketGoerli.contract, event: "AssociatedSystemSet", logs: logs, sub: sub}, nil
}

// WatchAssociatedSystemSet is a free log subscription operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchAssociatedSystemSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliAssociatedSystemSet, kind [][32]byte, id [][32]byte) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliAssociatedSystemSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseAssociatedSystemSet(log types.Log) (*SpotMarketGoerliAssociatedSystemSet, error) {
	event := new(SpotMarketGoerliAssociatedSystemSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliAsyncFixedFeeSetIterator is returned from FilterAsyncFixedFeeSet and is used to iterate over the raw logs and unpacked data for AsyncFixedFeeSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliAsyncFixedFeeSetIterator struct {
	Event *SpotMarketGoerliAsyncFixedFeeSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliAsyncFixedFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliAsyncFixedFeeSet)
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
		it.Event = new(SpotMarketGoerliAsyncFixedFeeSet)
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
func (it *SpotMarketGoerliAsyncFixedFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliAsyncFixedFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliAsyncFixedFeeSet represents a AsyncFixedFeeSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliAsyncFixedFeeSet struct {
	SynthMarketId *big.Int
	AsyncFixedFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAsyncFixedFeeSet is a free log retrieval operation binding the contract event 0xaef59d229d195c5b8372221acbf4041b926fb1616a95f93e44379e4f30d57bfe.
//
// Solidity: event AsyncFixedFeeSet(uint256 indexed synthMarketId, uint256 asyncFixedFee)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterAsyncFixedFeeSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliAsyncFixedFeeSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "AsyncFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliAsyncFixedFeeSetIterator{contract: _SpotMarketGoerli.contract, event: "AsyncFixedFeeSet", logs: logs, sub: sub}, nil
}

// WatchAsyncFixedFeeSet is a free log subscription operation binding the contract event 0xaef59d229d195c5b8372221acbf4041b926fb1616a95f93e44379e4f30d57bfe.
//
// Solidity: event AsyncFixedFeeSet(uint256 indexed synthMarketId, uint256 asyncFixedFee)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchAsyncFixedFeeSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliAsyncFixedFeeSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "AsyncFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliAsyncFixedFeeSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "AsyncFixedFeeSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseAsyncFixedFeeSet(log types.Log) (*SpotMarketGoerliAsyncFixedFeeSet, error) {
	event := new(SpotMarketGoerliAsyncFixedFeeSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "AsyncFixedFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliAtomicFixedFeeSetIterator is returned from FilterAtomicFixedFeeSet and is used to iterate over the raw logs and unpacked data for AtomicFixedFeeSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliAtomicFixedFeeSetIterator struct {
	Event *SpotMarketGoerliAtomicFixedFeeSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliAtomicFixedFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliAtomicFixedFeeSet)
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
		it.Event = new(SpotMarketGoerliAtomicFixedFeeSet)
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
func (it *SpotMarketGoerliAtomicFixedFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliAtomicFixedFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliAtomicFixedFeeSet represents a AtomicFixedFeeSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliAtomicFixedFeeSet struct {
	SynthMarketId  *big.Int
	AtomicFixedFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAtomicFixedFeeSet is a free log retrieval operation binding the contract event 0x6b0a526b06b2f30ba2d5b063c2ef81547512216d37c540b86039e3a19f1d2b3f.
//
// Solidity: event AtomicFixedFeeSet(uint256 indexed synthMarketId, uint256 atomicFixedFee)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterAtomicFixedFeeSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliAtomicFixedFeeSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "AtomicFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliAtomicFixedFeeSetIterator{contract: _SpotMarketGoerli.contract, event: "AtomicFixedFeeSet", logs: logs, sub: sub}, nil
}

// WatchAtomicFixedFeeSet is a free log subscription operation binding the contract event 0x6b0a526b06b2f30ba2d5b063c2ef81547512216d37c540b86039e3a19f1d2b3f.
//
// Solidity: event AtomicFixedFeeSet(uint256 indexed synthMarketId, uint256 atomicFixedFee)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchAtomicFixedFeeSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliAtomicFixedFeeSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "AtomicFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliAtomicFixedFeeSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "AtomicFixedFeeSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseAtomicFixedFeeSet(log types.Log) (*SpotMarketGoerliAtomicFixedFeeSet, error) {
	event := new(SpotMarketGoerliAtomicFixedFeeSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "AtomicFixedFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliCollateralLeverageSetIterator is returned from FilterCollateralLeverageSet and is used to iterate over the raw logs and unpacked data for CollateralLeverageSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliCollateralLeverageSetIterator struct {
	Event *SpotMarketGoerliCollateralLeverageSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliCollateralLeverageSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliCollateralLeverageSet)
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
		it.Event = new(SpotMarketGoerliCollateralLeverageSet)
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
func (it *SpotMarketGoerliCollateralLeverageSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliCollateralLeverageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliCollateralLeverageSet represents a CollateralLeverageSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliCollateralLeverageSet struct {
	SynthMarketId      *big.Int
	CollateralLeverage *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCollateralLeverageSet is a free log retrieval operation binding the contract event 0x83113c0e8d811f7e7017948357e945a1a8a6b6fc0d76c8512ffdd6f6766e8a13.
//
// Solidity: event CollateralLeverageSet(uint256 indexed synthMarketId, uint256 collateralLeverage)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterCollateralLeverageSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliCollateralLeverageSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "CollateralLeverageSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliCollateralLeverageSetIterator{contract: _SpotMarketGoerli.contract, event: "CollateralLeverageSet", logs: logs, sub: sub}, nil
}

// WatchCollateralLeverageSet is a free log subscription operation binding the contract event 0x83113c0e8d811f7e7017948357e945a1a8a6b6fc0d76c8512ffdd6f6766e8a13.
//
// Solidity: event CollateralLeverageSet(uint256 indexed synthMarketId, uint256 collateralLeverage)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchCollateralLeverageSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliCollateralLeverageSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "CollateralLeverageSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliCollateralLeverageSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "CollateralLeverageSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseCollateralLeverageSet(log types.Log) (*SpotMarketGoerliCollateralLeverageSet, error) {
	event := new(SpotMarketGoerliCollateralLeverageSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "CollateralLeverageSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliDecayRateUpdatedIterator is returned from FilterDecayRateUpdated and is used to iterate over the raw logs and unpacked data for DecayRateUpdated events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliDecayRateUpdatedIterator struct {
	Event *SpotMarketGoerliDecayRateUpdated // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliDecayRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliDecayRateUpdated)
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
		it.Event = new(SpotMarketGoerliDecayRateUpdated)
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
func (it *SpotMarketGoerliDecayRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliDecayRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliDecayRateUpdated represents a DecayRateUpdated event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliDecayRateUpdated struct {
	MarketId *big.Int
	Rate     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDecayRateUpdated is a free log retrieval operation binding the contract event 0x8a1ed33cd71d1533eadd1d4eb0ea2ae64da7d343cb2a932fdf135f345264e2b5.
//
// Solidity: event DecayRateUpdated(uint128 indexed marketId, uint256 rate)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterDecayRateUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*SpotMarketGoerliDecayRateUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "DecayRateUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliDecayRateUpdatedIterator{contract: _SpotMarketGoerli.contract, event: "DecayRateUpdated", logs: logs, sub: sub}, nil
}

// WatchDecayRateUpdated is a free log subscription operation binding the contract event 0x8a1ed33cd71d1533eadd1d4eb0ea2ae64da7d343cb2a932fdf135f345264e2b5.
//
// Solidity: event DecayRateUpdated(uint128 indexed marketId, uint256 rate)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchDecayRateUpdated(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliDecayRateUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "DecayRateUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliDecayRateUpdated)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "DecayRateUpdated", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseDecayRateUpdated(log types.Log) (*SpotMarketGoerliDecayRateUpdated, error) {
	event := new(SpotMarketGoerliDecayRateUpdated)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "DecayRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliFeatureFlagAllowAllSetIterator is returned from FilterFeatureFlagAllowAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowAllSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeatureFlagAllowAllSetIterator struct {
	Event *SpotMarketGoerliFeatureFlagAllowAllSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliFeatureFlagAllowAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliFeatureFlagAllowAllSet)
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
		it.Event = new(SpotMarketGoerliFeatureFlagAllowAllSet)
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
func (it *SpotMarketGoerliFeatureFlagAllowAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliFeatureFlagAllowAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliFeatureFlagAllowAllSet represents a FeatureFlagAllowAllSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeatureFlagAllowAllSet struct {
	Feature  [32]byte
	AllowAll bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowAllSet is a free log retrieval operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterFeatureFlagAllowAllSet(opts *bind.FilterOpts, feature [][32]byte) (*SpotMarketGoerliFeatureFlagAllowAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliFeatureFlagAllowAllSetIterator{contract: _SpotMarketGoerli.contract, event: "FeatureFlagAllowAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowAllSet is a free log subscription operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchFeatureFlagAllowAllSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliFeatureFlagAllowAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliFeatureFlagAllowAllSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseFeatureFlagAllowAllSet(log types.Log) (*SpotMarketGoerliFeatureFlagAllowAllSet, error) {
	event := new(SpotMarketGoerliFeatureFlagAllowAllSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliFeatureFlagAllowlistAddedIterator is returned from FilterFeatureFlagAllowlistAdded and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistAdded events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeatureFlagAllowlistAddedIterator struct {
	Event *SpotMarketGoerliFeatureFlagAllowlistAdded // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliFeatureFlagAllowlistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliFeatureFlagAllowlistAdded)
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
		it.Event = new(SpotMarketGoerliFeatureFlagAllowlistAdded)
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
func (it *SpotMarketGoerliFeatureFlagAllowlistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliFeatureFlagAllowlistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliFeatureFlagAllowlistAdded represents a FeatureFlagAllowlistAdded event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeatureFlagAllowlistAdded struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistAdded is a free log retrieval operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterFeatureFlagAllowlistAdded(opts *bind.FilterOpts, feature [][32]byte) (*SpotMarketGoerliFeatureFlagAllowlistAddedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliFeatureFlagAllowlistAddedIterator{contract: _SpotMarketGoerli.contract, event: "FeatureFlagAllowlistAdded", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistAdded is a free log subscription operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchFeatureFlagAllowlistAdded(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliFeatureFlagAllowlistAdded, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliFeatureFlagAllowlistAdded)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseFeatureFlagAllowlistAdded(log types.Log) (*SpotMarketGoerliFeatureFlagAllowlistAdded, error) {
	event := new(SpotMarketGoerliFeatureFlagAllowlistAdded)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliFeatureFlagAllowlistRemovedIterator is returned from FilterFeatureFlagAllowlistRemoved and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistRemoved events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeatureFlagAllowlistRemovedIterator struct {
	Event *SpotMarketGoerliFeatureFlagAllowlistRemoved // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliFeatureFlagAllowlistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliFeatureFlagAllowlistRemoved)
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
		it.Event = new(SpotMarketGoerliFeatureFlagAllowlistRemoved)
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
func (it *SpotMarketGoerliFeatureFlagAllowlistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliFeatureFlagAllowlistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliFeatureFlagAllowlistRemoved represents a FeatureFlagAllowlistRemoved event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeatureFlagAllowlistRemoved struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistRemoved is a free log retrieval operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterFeatureFlagAllowlistRemoved(opts *bind.FilterOpts, feature [][32]byte) (*SpotMarketGoerliFeatureFlagAllowlistRemovedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliFeatureFlagAllowlistRemovedIterator{contract: _SpotMarketGoerli.contract, event: "FeatureFlagAllowlistRemoved", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistRemoved is a free log subscription operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchFeatureFlagAllowlistRemoved(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliFeatureFlagAllowlistRemoved, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliFeatureFlagAllowlistRemoved)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseFeatureFlagAllowlistRemoved(log types.Log) (*SpotMarketGoerliFeatureFlagAllowlistRemoved, error) {
	event := new(SpotMarketGoerliFeatureFlagAllowlistRemoved)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliFeatureFlagDeniersResetIterator is returned from FilterFeatureFlagDeniersReset and is used to iterate over the raw logs and unpacked data for FeatureFlagDeniersReset events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeatureFlagDeniersResetIterator struct {
	Event *SpotMarketGoerliFeatureFlagDeniersReset // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliFeatureFlagDeniersResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliFeatureFlagDeniersReset)
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
		it.Event = new(SpotMarketGoerliFeatureFlagDeniersReset)
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
func (it *SpotMarketGoerliFeatureFlagDeniersResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliFeatureFlagDeniersResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliFeatureFlagDeniersReset represents a FeatureFlagDeniersReset event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeatureFlagDeniersReset struct {
	Feature [32]byte
	Deniers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDeniersReset is a free log retrieval operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterFeatureFlagDeniersReset(opts *bind.FilterOpts, feature [][32]byte) (*SpotMarketGoerliFeatureFlagDeniersResetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliFeatureFlagDeniersResetIterator{contract: _SpotMarketGoerli.contract, event: "FeatureFlagDeniersReset", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDeniersReset is a free log subscription operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchFeatureFlagDeniersReset(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliFeatureFlagDeniersReset, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliFeatureFlagDeniersReset)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseFeatureFlagDeniersReset(log types.Log) (*SpotMarketGoerliFeatureFlagDeniersReset, error) {
	event := new(SpotMarketGoerliFeatureFlagDeniersReset)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliFeatureFlagDenyAllSetIterator is returned from FilterFeatureFlagDenyAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagDenyAllSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeatureFlagDenyAllSetIterator struct {
	Event *SpotMarketGoerliFeatureFlagDenyAllSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliFeatureFlagDenyAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliFeatureFlagDenyAllSet)
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
		it.Event = new(SpotMarketGoerliFeatureFlagDenyAllSet)
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
func (it *SpotMarketGoerliFeatureFlagDenyAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliFeatureFlagDenyAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliFeatureFlagDenyAllSet represents a FeatureFlagDenyAllSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeatureFlagDenyAllSet struct {
	Feature [32]byte
	DenyAll bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDenyAllSet is a free log retrieval operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterFeatureFlagDenyAllSet(opts *bind.FilterOpts, feature [][32]byte) (*SpotMarketGoerliFeatureFlagDenyAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliFeatureFlagDenyAllSetIterator{contract: _SpotMarketGoerli.contract, event: "FeatureFlagDenyAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDenyAllSet is a free log subscription operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchFeatureFlagDenyAllSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliFeatureFlagDenyAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliFeatureFlagDenyAllSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseFeatureFlagDenyAllSet(log types.Log) (*SpotMarketGoerliFeatureFlagDenyAllSet, error) {
	event := new(SpotMarketGoerliFeatureFlagDenyAllSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliFeeCollectorSetIterator is returned from FilterFeeCollectorSet and is used to iterate over the raw logs and unpacked data for FeeCollectorSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeeCollectorSetIterator struct {
	Event *SpotMarketGoerliFeeCollectorSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliFeeCollectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliFeeCollectorSet)
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
		it.Event = new(SpotMarketGoerliFeeCollectorSet)
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
func (it *SpotMarketGoerliFeeCollectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliFeeCollectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliFeeCollectorSet represents a FeeCollectorSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliFeeCollectorSet struct {
	SynthMarketId *big.Int
	FeeCollector  common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorSet is a free log retrieval operation binding the contract event 0x9559a9b7e14bf53553c17859be245a108350185ec859ec690012b13b820b7ef4.
//
// Solidity: event FeeCollectorSet(uint256 indexed synthMarketId, address feeCollector)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterFeeCollectorSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliFeeCollectorSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "FeeCollectorSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliFeeCollectorSetIterator{contract: _SpotMarketGoerli.contract, event: "FeeCollectorSet", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorSet is a free log subscription operation binding the contract event 0x9559a9b7e14bf53553c17859be245a108350185ec859ec690012b13b820b7ef4.
//
// Solidity: event FeeCollectorSet(uint256 indexed synthMarketId, address feeCollector)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchFeeCollectorSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliFeeCollectorSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "FeeCollectorSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliFeeCollectorSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseFeeCollectorSet(log types.Log) (*SpotMarketGoerliFeeCollectorSet, error) {
	event := new(SpotMarketGoerliFeeCollectorSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliMarketNominationRenouncedIterator is returned from FilterMarketNominationRenounced and is used to iterate over the raw logs and unpacked data for MarketNominationRenounced events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliMarketNominationRenouncedIterator struct {
	Event *SpotMarketGoerliMarketNominationRenounced // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliMarketNominationRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliMarketNominationRenounced)
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
		it.Event = new(SpotMarketGoerliMarketNominationRenounced)
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
func (it *SpotMarketGoerliMarketNominationRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliMarketNominationRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliMarketNominationRenounced represents a MarketNominationRenounced event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliMarketNominationRenounced struct {
	MarketId *big.Int
	Nominee  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketNominationRenounced is a free log retrieval operation binding the contract event 0xf5b87e3c7e0caa8e0d233591fff85e764ebc73d5c7027cce729fd4beab04c2b6.
//
// Solidity: event MarketNominationRenounced(uint128 indexed marketId, address nominee)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterMarketNominationRenounced(opts *bind.FilterOpts, marketId []*big.Int) (*SpotMarketGoerliMarketNominationRenouncedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "MarketNominationRenounced", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliMarketNominationRenouncedIterator{contract: _SpotMarketGoerli.contract, event: "MarketNominationRenounced", logs: logs, sub: sub}, nil
}

// WatchMarketNominationRenounced is a free log subscription operation binding the contract event 0xf5b87e3c7e0caa8e0d233591fff85e764ebc73d5c7027cce729fd4beab04c2b6.
//
// Solidity: event MarketNominationRenounced(uint128 indexed marketId, address nominee)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchMarketNominationRenounced(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliMarketNominationRenounced, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "MarketNominationRenounced", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliMarketNominationRenounced)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "MarketNominationRenounced", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseMarketNominationRenounced(log types.Log) (*SpotMarketGoerliMarketNominationRenounced, error) {
	event := new(SpotMarketGoerliMarketNominationRenounced)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "MarketNominationRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliMarketOwnerChangedIterator is returned from FilterMarketOwnerChanged and is used to iterate over the raw logs and unpacked data for MarketOwnerChanged events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliMarketOwnerChangedIterator struct {
	Event *SpotMarketGoerliMarketOwnerChanged // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliMarketOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliMarketOwnerChanged)
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
		it.Event = new(SpotMarketGoerliMarketOwnerChanged)
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
func (it *SpotMarketGoerliMarketOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliMarketOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliMarketOwnerChanged represents a MarketOwnerChanged event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliMarketOwnerChanged struct {
	MarketId *big.Int
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketOwnerChanged is a free log retrieval operation binding the contract event 0xe73c996387b656d1e0ea2866c854ed68122ce4e4eea51d6af012938b3c7ff52f.
//
// Solidity: event MarketOwnerChanged(uint128 indexed marketId, address oldOwner, address newOwner)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterMarketOwnerChanged(opts *bind.FilterOpts, marketId []*big.Int) (*SpotMarketGoerliMarketOwnerChangedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "MarketOwnerChanged", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliMarketOwnerChangedIterator{contract: _SpotMarketGoerli.contract, event: "MarketOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchMarketOwnerChanged is a free log subscription operation binding the contract event 0xe73c996387b656d1e0ea2866c854ed68122ce4e4eea51d6af012938b3c7ff52f.
//
// Solidity: event MarketOwnerChanged(uint128 indexed marketId, address oldOwner, address newOwner)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchMarketOwnerChanged(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliMarketOwnerChanged, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "MarketOwnerChanged", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliMarketOwnerChanged)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "MarketOwnerChanged", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseMarketOwnerChanged(log types.Log) (*SpotMarketGoerliMarketOwnerChanged, error) {
	event := new(SpotMarketGoerliMarketOwnerChanged)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "MarketOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliMarketOwnerNominatedIterator is returned from FilterMarketOwnerNominated and is used to iterate over the raw logs and unpacked data for MarketOwnerNominated events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliMarketOwnerNominatedIterator struct {
	Event *SpotMarketGoerliMarketOwnerNominated // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliMarketOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliMarketOwnerNominated)
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
		it.Event = new(SpotMarketGoerliMarketOwnerNominated)
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
func (it *SpotMarketGoerliMarketOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliMarketOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliMarketOwnerNominated represents a MarketOwnerNominated event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliMarketOwnerNominated struct {
	MarketId *big.Int
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketOwnerNominated is a free log retrieval operation binding the contract event 0x54605d90ee82d9b4318b0b4b479f3966976e44b94c6aff221c04f5294f85016b.
//
// Solidity: event MarketOwnerNominated(uint128 indexed marketId, address newOwner)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterMarketOwnerNominated(opts *bind.FilterOpts, marketId []*big.Int) (*SpotMarketGoerliMarketOwnerNominatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "MarketOwnerNominated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliMarketOwnerNominatedIterator{contract: _SpotMarketGoerli.contract, event: "MarketOwnerNominated", logs: logs, sub: sub}, nil
}

// WatchMarketOwnerNominated is a free log subscription operation binding the contract event 0x54605d90ee82d9b4318b0b4b479f3966976e44b94c6aff221c04f5294f85016b.
//
// Solidity: event MarketOwnerNominated(uint128 indexed marketId, address newOwner)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchMarketOwnerNominated(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliMarketOwnerNominated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "MarketOwnerNominated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliMarketOwnerNominated)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "MarketOwnerNominated", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseMarketOwnerNominated(log types.Log) (*SpotMarketGoerliMarketOwnerNominated, error) {
	event := new(SpotMarketGoerliMarketOwnerNominated)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "MarketOwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliMarketSkewScaleSetIterator is returned from FilterMarketSkewScaleSet and is used to iterate over the raw logs and unpacked data for MarketSkewScaleSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliMarketSkewScaleSetIterator struct {
	Event *SpotMarketGoerliMarketSkewScaleSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliMarketSkewScaleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliMarketSkewScaleSet)
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
		it.Event = new(SpotMarketGoerliMarketSkewScaleSet)
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
func (it *SpotMarketGoerliMarketSkewScaleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliMarketSkewScaleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliMarketSkewScaleSet represents a MarketSkewScaleSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliMarketSkewScaleSet struct {
	SynthMarketId *big.Int
	SkewScale     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarketSkewScaleSet is a free log retrieval operation binding the contract event 0x786fdfd5a0e146d8f4878876a8439ff01436e5969b14682e12d07d7e926b157c.
//
// Solidity: event MarketSkewScaleSet(uint256 indexed synthMarketId, uint256 skewScale)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterMarketSkewScaleSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliMarketSkewScaleSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "MarketSkewScaleSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliMarketSkewScaleSetIterator{contract: _SpotMarketGoerli.contract, event: "MarketSkewScaleSet", logs: logs, sub: sub}, nil
}

// WatchMarketSkewScaleSet is a free log subscription operation binding the contract event 0x786fdfd5a0e146d8f4878876a8439ff01436e5969b14682e12d07d7e926b157c.
//
// Solidity: event MarketSkewScaleSet(uint256 indexed synthMarketId, uint256 skewScale)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchMarketSkewScaleSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliMarketSkewScaleSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "MarketSkewScaleSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliMarketSkewScaleSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "MarketSkewScaleSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseMarketSkewScaleSet(log types.Log) (*SpotMarketGoerliMarketSkewScaleSet, error) {
	event := new(SpotMarketGoerliMarketSkewScaleSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "MarketSkewScaleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliMarketUtilizationFeesSetIterator is returned from FilterMarketUtilizationFeesSet and is used to iterate over the raw logs and unpacked data for MarketUtilizationFeesSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliMarketUtilizationFeesSetIterator struct {
	Event *SpotMarketGoerliMarketUtilizationFeesSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliMarketUtilizationFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliMarketUtilizationFeesSet)
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
		it.Event = new(SpotMarketGoerliMarketUtilizationFeesSet)
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
func (it *SpotMarketGoerliMarketUtilizationFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliMarketUtilizationFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliMarketUtilizationFeesSet represents a MarketUtilizationFeesSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliMarketUtilizationFeesSet struct {
	SynthMarketId      *big.Int
	UtilizationFeeRate *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketUtilizationFeesSet is a free log retrieval operation binding the contract event 0x83dfad56ac61e49feb43345b8c73b6d45eb121decc66b1709ca0413b31c64f63.
//
// Solidity: event MarketUtilizationFeesSet(uint256 indexed synthMarketId, uint256 utilizationFeeRate)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterMarketUtilizationFeesSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliMarketUtilizationFeesSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "MarketUtilizationFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliMarketUtilizationFeesSetIterator{contract: _SpotMarketGoerli.contract, event: "MarketUtilizationFeesSet", logs: logs, sub: sub}, nil
}

// WatchMarketUtilizationFeesSet is a free log subscription operation binding the contract event 0x83dfad56ac61e49feb43345b8c73b6d45eb121decc66b1709ca0413b31c64f63.
//
// Solidity: event MarketUtilizationFeesSet(uint256 indexed synthMarketId, uint256 utilizationFeeRate)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchMarketUtilizationFeesSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliMarketUtilizationFeesSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "MarketUtilizationFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliMarketUtilizationFeesSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "MarketUtilizationFeesSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseMarketUtilizationFeesSet(log types.Log) (*SpotMarketGoerliMarketUtilizationFeesSet, error) {
	event := new(SpotMarketGoerliMarketUtilizationFeesSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "MarketUtilizationFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliOrderCancelledIterator struct {
	Event *SpotMarketGoerliOrderCancelled // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliOrderCancelled)
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
		it.Event = new(SpotMarketGoerliOrderCancelled)
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
func (it *SpotMarketGoerliOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliOrderCancelled represents a OrderCancelled event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliOrderCancelled struct {
	MarketId        *big.Int
	AsyncOrderId    *big.Int
	AsyncOrderClaim AsyncOrderClaimData
	Sender          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xa57ffc5057d10e88a0dd7713d14130f9638d680af94401e7f1f4fba44a3ad5e2.
//
// Solidity: event OrderCancelled(uint128 indexed marketId, uint128 indexed asyncOrderId, (uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim, address indexed sender)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterOrderCancelled(opts *bind.FilterOpts, marketId []*big.Int, asyncOrderId []*big.Int, sender []common.Address) (*SpotMarketGoerliOrderCancelledIterator, error) {

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

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "OrderCancelled", marketIdRule, asyncOrderIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliOrderCancelledIterator{contract: _SpotMarketGoerli.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xa57ffc5057d10e88a0dd7713d14130f9638d680af94401e7f1f4fba44a3ad5e2.
//
// Solidity: event OrderCancelled(uint128 indexed marketId, uint128 indexed asyncOrderId, (uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim, address indexed sender)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliOrderCancelled, marketId []*big.Int, asyncOrderId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "OrderCancelled", marketIdRule, asyncOrderIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliOrderCancelled)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseOrderCancelled(log types.Log) (*SpotMarketGoerliOrderCancelled, error) {
	event := new(SpotMarketGoerliOrderCancelled)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliOrderCommittedIterator is returned from FilterOrderCommitted and is used to iterate over the raw logs and unpacked data for OrderCommitted events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliOrderCommittedIterator struct {
	Event *SpotMarketGoerliOrderCommitted // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliOrderCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliOrderCommitted)
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
		it.Event = new(SpotMarketGoerliOrderCommitted)
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
func (it *SpotMarketGoerliOrderCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliOrderCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliOrderCommitted represents a OrderCommitted event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliOrderCommitted struct {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterOrderCommitted(opts *bind.FilterOpts, marketId []*big.Int, orderType []uint8, sender []common.Address) (*SpotMarketGoerliOrderCommittedIterator, error) {

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

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "OrderCommitted", marketIdRule, orderTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliOrderCommittedIterator{contract: _SpotMarketGoerli.contract, event: "OrderCommitted", logs: logs, sub: sub}, nil
}

// WatchOrderCommitted is a free log subscription operation binding the contract event 0xb26c216bf0a127dddc2431e4d8ca845513c8e6fb80e754296e7afab1ce92722f.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint8 indexed orderType, uint256 amountProvided, uint128 asyncOrderId, address indexed sender, address referrer)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchOrderCommitted(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliOrderCommitted, marketId []*big.Int, orderType []uint8, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "OrderCommitted", marketIdRule, orderTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliOrderCommitted)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseOrderCommitted(log types.Log) (*SpotMarketGoerliOrderCommitted, error) {
	event := new(SpotMarketGoerliOrderCommitted)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliOrderSettledIterator is returned from FilterOrderSettled and is used to iterate over the raw logs and unpacked data for OrderSettled events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliOrderSettledIterator struct {
	Event *SpotMarketGoerliOrderSettled // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliOrderSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliOrderSettled)
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
		it.Event = new(SpotMarketGoerliOrderSettled)
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
func (it *SpotMarketGoerliOrderSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliOrderSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliOrderSettled represents a OrderSettled event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliOrderSettled struct {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterOrderSettled(opts *bind.FilterOpts, marketId []*big.Int, asyncOrderId []*big.Int, settler []common.Address) (*SpotMarketGoerliOrderSettledIterator, error) {

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

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "OrderSettled", marketIdRule, asyncOrderIdRule, settlerRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliOrderSettledIterator{contract: _SpotMarketGoerli.contract, event: "OrderSettled", logs: logs, sub: sub}, nil
}

// WatchOrderSettled is a free log subscription operation binding the contract event 0x6c5e8ff282d52fb9f532408e86d4afc62fc1f89c749a8ddca7a6f34c0439a183.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed asyncOrderId, uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address indexed settler, uint256 price, uint8 orderType)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchOrderSettled(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliOrderSettled, marketId []*big.Int, asyncOrderId []*big.Int, settler []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "OrderSettled", marketIdRule, asyncOrderIdRule, settlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliOrderSettled)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "OrderSettled", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseOrderSettled(log types.Log) (*SpotMarketGoerliOrderSettled, error) {
	event := new(SpotMarketGoerliOrderSettled)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "OrderSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliOwnerChangedIterator struct {
	Event *SpotMarketGoerliOwnerChanged // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliOwnerChanged)
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
		it.Event = new(SpotMarketGoerliOwnerChanged)
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
func (it *SpotMarketGoerliOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliOwnerChanged represents a OwnerChanged event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SpotMarketGoerliOwnerChangedIterator, error) {

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliOwnerChangedIterator{contract: _SpotMarketGoerli.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliOwnerChanged)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseOwnerChanged(log types.Log) (*SpotMarketGoerliOwnerChanged, error) {
	event := new(SpotMarketGoerliOwnerChanged)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliOwnerNominatedIterator struct {
	Event *SpotMarketGoerliOwnerNominated // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliOwnerNominated)
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
		it.Event = new(SpotMarketGoerliOwnerNominated)
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
func (it *SpotMarketGoerliOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliOwnerNominated represents a OwnerNominated event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*SpotMarketGoerliOwnerNominatedIterator, error) {

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliOwnerNominatedIterator{contract: _SpotMarketGoerli.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliOwnerNominated)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseOwnerNominated(log types.Log) (*SpotMarketGoerliOwnerNominated, error) {
	event := new(SpotMarketGoerliOwnerNominated)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliReferrerShareUpdatedIterator is returned from FilterReferrerShareUpdated and is used to iterate over the raw logs and unpacked data for ReferrerShareUpdated events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliReferrerShareUpdatedIterator struct {
	Event *SpotMarketGoerliReferrerShareUpdated // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliReferrerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliReferrerShareUpdated)
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
		it.Event = new(SpotMarketGoerliReferrerShareUpdated)
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
func (it *SpotMarketGoerliReferrerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliReferrerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliReferrerShareUpdated represents a ReferrerShareUpdated event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliReferrerShareUpdated struct {
	MarketId        *big.Int
	Referrer        common.Address
	SharePercentage *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReferrerShareUpdated is a free log retrieval operation binding the contract event 0xd2a3339bb3c610e9030023c1cb3e89374fe0ebd7e37faee9b3d343f33e9df2fb.
//
// Solidity: event ReferrerShareUpdated(uint128 indexed marketId, address referrer, uint256 sharePercentage)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterReferrerShareUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*SpotMarketGoerliReferrerShareUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "ReferrerShareUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliReferrerShareUpdatedIterator{contract: _SpotMarketGoerli.contract, event: "ReferrerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchReferrerShareUpdated is a free log subscription operation binding the contract event 0xd2a3339bb3c610e9030023c1cb3e89374fe0ebd7e37faee9b3d343f33e9df2fb.
//
// Solidity: event ReferrerShareUpdated(uint128 indexed marketId, address referrer, uint256 sharePercentage)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchReferrerShareUpdated(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliReferrerShareUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "ReferrerShareUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliReferrerShareUpdated)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseReferrerShareUpdated(log types.Log) (*SpotMarketGoerliReferrerShareUpdated, error) {
	event := new(SpotMarketGoerliReferrerShareUpdated)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSettlementStrategyAddedIterator is returned from FilterSettlementStrategyAdded and is used to iterate over the raw logs and unpacked data for SettlementStrategyAdded events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSettlementStrategyAddedIterator struct {
	Event *SpotMarketGoerliSettlementStrategyAdded // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSettlementStrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSettlementStrategyAdded)
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
		it.Event = new(SpotMarketGoerliSettlementStrategyAdded)
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
func (it *SpotMarketGoerliSettlementStrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSettlementStrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSettlementStrategyAdded represents a SettlementStrategyAdded event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSettlementStrategyAdded struct {
	SynthMarketId *big.Int
	StrategyId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyAdded is a free log retrieval operation binding the contract event 0x91750a8e3d84ed1bccdc79dcecf63cc1b6f83b5bf8293bf86628cbb248e487f1.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed synthMarketId, uint256 indexed strategyId)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSettlementStrategyAdded(opts *bind.FilterOpts, synthMarketId []*big.Int, strategyId []*big.Int) (*SpotMarketGoerliSettlementStrategyAddedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SettlementStrategyAdded", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSettlementStrategyAddedIterator{contract: _SpotMarketGoerli.contract, event: "SettlementStrategyAdded", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyAdded is a free log subscription operation binding the contract event 0x91750a8e3d84ed1bccdc79dcecf63cc1b6f83b5bf8293bf86628cbb248e487f1.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed synthMarketId, uint256 indexed strategyId)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSettlementStrategyAdded(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSettlementStrategyAdded, synthMarketId []*big.Int, strategyId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SettlementStrategyAdded", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSettlementStrategyAdded)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSettlementStrategyAdded(log types.Log) (*SpotMarketGoerliSettlementStrategyAdded, error) {
	event := new(SpotMarketGoerliSettlementStrategyAdded)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSettlementStrategyUpdatedIterator is returned from FilterSettlementStrategyUpdated and is used to iterate over the raw logs and unpacked data for SettlementStrategyUpdated events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSettlementStrategyUpdatedIterator struct {
	Event *SpotMarketGoerliSettlementStrategyUpdated // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSettlementStrategyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSettlementStrategyUpdated)
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
		it.Event = new(SpotMarketGoerliSettlementStrategyUpdated)
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
func (it *SpotMarketGoerliSettlementStrategyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSettlementStrategyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSettlementStrategyUpdated represents a SettlementStrategyUpdated event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSettlementStrategyUpdated struct {
	SynthMarketId *big.Int
	StrategyId    *big.Int
	Enabled       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyUpdated is a free log retrieval operation binding the contract event 0xb0187d91c1ffe66740e5f7619aab4d868a85ac3698b026bed3f420d871b6c048.
//
// Solidity: event SettlementStrategyUpdated(uint128 indexed synthMarketId, uint256 indexed strategyId, bool enabled)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSettlementStrategyUpdated(opts *bind.FilterOpts, synthMarketId []*big.Int, strategyId []*big.Int) (*SpotMarketGoerliSettlementStrategyUpdatedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SettlementStrategyUpdated", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSettlementStrategyUpdatedIterator{contract: _SpotMarketGoerli.contract, event: "SettlementStrategyUpdated", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyUpdated is a free log subscription operation binding the contract event 0xb0187d91c1ffe66740e5f7619aab4d868a85ac3698b026bed3f420d871b6c048.
//
// Solidity: event SettlementStrategyUpdated(uint128 indexed synthMarketId, uint256 indexed strategyId, bool enabled)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSettlementStrategyUpdated(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSettlementStrategyUpdated, synthMarketId []*big.Int, strategyId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SettlementStrategyUpdated", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSettlementStrategyUpdated)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SettlementStrategyUpdated", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSettlementStrategyUpdated(log types.Log) (*SpotMarketGoerliSettlementStrategyUpdated, error) {
	event := new(SpotMarketGoerliSettlementStrategyUpdated)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SettlementStrategyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSynthBoughtIterator is returned from FilterSynthBought and is used to iterate over the raw logs and unpacked data for SynthBought events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthBoughtIterator struct {
	Event *SpotMarketGoerliSynthBought // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSynthBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSynthBought)
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
		it.Event = new(SpotMarketGoerliSynthBought)
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
func (it *SpotMarketGoerliSynthBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSynthBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSynthBought represents a SynthBought event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthBought struct {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSynthBought(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliSynthBoughtIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SynthBought", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSynthBoughtIterator{contract: _SpotMarketGoerli.contract, event: "SynthBought", logs: logs, sub: sub}, nil
}

// WatchSynthBought is a free log subscription operation binding the contract event 0xac82d63e679c7d862613aa8b5ccd94f9adc4986763ab14bb3351ab9092ef1303.
//
// Solidity: event SynthBought(uint256 indexed synthMarketId, uint256 synthReturned, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address referrer, uint256 price)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSynthBought(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSynthBought, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SynthBought", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSynthBought)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthBought", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSynthBought(log types.Log) (*SpotMarketGoerliSynthBought, error) {
	event := new(SpotMarketGoerliSynthBought)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSynthImplementationSetIterator is returned from FilterSynthImplementationSet and is used to iterate over the raw logs and unpacked data for SynthImplementationSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthImplementationSetIterator struct {
	Event *SpotMarketGoerliSynthImplementationSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSynthImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSynthImplementationSet)
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
		it.Event = new(SpotMarketGoerliSynthImplementationSet)
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
func (it *SpotMarketGoerliSynthImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSynthImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSynthImplementationSet represents a SynthImplementationSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthImplementationSet struct {
	SynthImplementation common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSynthImplementationSet is a free log retrieval operation binding the contract event 0xafffc48b3243eba10d901f21ba761ad741f18a23feed86ca425df4974d3314b0.
//
// Solidity: event SynthImplementationSet(address synthImplementation)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSynthImplementationSet(opts *bind.FilterOpts) (*SpotMarketGoerliSynthImplementationSetIterator, error) {

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SynthImplementationSet")
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSynthImplementationSetIterator{contract: _SpotMarketGoerli.contract, event: "SynthImplementationSet", logs: logs, sub: sub}, nil
}

// WatchSynthImplementationSet is a free log subscription operation binding the contract event 0xafffc48b3243eba10d901f21ba761ad741f18a23feed86ca425df4974d3314b0.
//
// Solidity: event SynthImplementationSet(address synthImplementation)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSynthImplementationSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSynthImplementationSet) (event.Subscription, error) {

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SynthImplementationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSynthImplementationSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthImplementationSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSynthImplementationSet(log types.Log) (*SpotMarketGoerliSynthImplementationSet, error) {
	event := new(SpotMarketGoerliSynthImplementationSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSynthImplementationUpgradedIterator is returned from FilterSynthImplementationUpgraded and is used to iterate over the raw logs and unpacked data for SynthImplementationUpgraded events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthImplementationUpgradedIterator struct {
	Event *SpotMarketGoerliSynthImplementationUpgraded // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSynthImplementationUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSynthImplementationUpgraded)
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
		it.Event = new(SpotMarketGoerliSynthImplementationUpgraded)
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
func (it *SpotMarketGoerliSynthImplementationUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSynthImplementationUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSynthImplementationUpgraded represents a SynthImplementationUpgraded event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthImplementationUpgraded struct {
	SynthMarketId  *big.Int
	Proxy          common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSynthImplementationUpgraded is a free log retrieval operation binding the contract event 0xa7badace8b24daeeb3981497a506d3812b3dc6f147ef3f78bc0e2cc664c50330.
//
// Solidity: event SynthImplementationUpgraded(uint256 indexed synthMarketId, address indexed proxy, address implementation)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSynthImplementationUpgraded(opts *bind.FilterOpts, synthMarketId []*big.Int, proxy []common.Address) (*SpotMarketGoerliSynthImplementationUpgradedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SynthImplementationUpgraded", synthMarketIdRule, proxyRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSynthImplementationUpgradedIterator{contract: _SpotMarketGoerli.contract, event: "SynthImplementationUpgraded", logs: logs, sub: sub}, nil
}

// WatchSynthImplementationUpgraded is a free log subscription operation binding the contract event 0xa7badace8b24daeeb3981497a506d3812b3dc6f147ef3f78bc0e2cc664c50330.
//
// Solidity: event SynthImplementationUpgraded(uint256 indexed synthMarketId, address indexed proxy, address implementation)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSynthImplementationUpgraded(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSynthImplementationUpgraded, synthMarketId []*big.Int, proxy []common.Address) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SynthImplementationUpgraded", synthMarketIdRule, proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSynthImplementationUpgraded)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthImplementationUpgraded", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSynthImplementationUpgraded(log types.Log) (*SpotMarketGoerliSynthImplementationUpgraded, error) {
	event := new(SpotMarketGoerliSynthImplementationUpgraded)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthImplementationUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSynthPriceDataUpdatedIterator is returned from FilterSynthPriceDataUpdated and is used to iterate over the raw logs and unpacked data for SynthPriceDataUpdated events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthPriceDataUpdatedIterator struct {
	Event *SpotMarketGoerliSynthPriceDataUpdated // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSynthPriceDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSynthPriceDataUpdated)
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
		it.Event = new(SpotMarketGoerliSynthPriceDataUpdated)
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
func (it *SpotMarketGoerliSynthPriceDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSynthPriceDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSynthPriceDataUpdated represents a SynthPriceDataUpdated event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthPriceDataUpdated struct {
	SynthMarketId *big.Int
	BuyFeedId     [32]byte
	SellFeedId    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSynthPriceDataUpdated is a free log retrieval operation binding the contract event 0x30309eafe094276dfad42e7ac4676aebcf5ea8480c86708ad2a2c9c0335c912e.
//
// Solidity: event SynthPriceDataUpdated(uint256 indexed synthMarketId, bytes32 indexed buyFeedId, bytes32 indexed sellFeedId)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSynthPriceDataUpdated(opts *bind.FilterOpts, synthMarketId []*big.Int, buyFeedId [][32]byte, sellFeedId [][32]byte) (*SpotMarketGoerliSynthPriceDataUpdatedIterator, error) {

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

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SynthPriceDataUpdated", synthMarketIdRule, buyFeedIdRule, sellFeedIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSynthPriceDataUpdatedIterator{contract: _SpotMarketGoerli.contract, event: "SynthPriceDataUpdated", logs: logs, sub: sub}, nil
}

// WatchSynthPriceDataUpdated is a free log subscription operation binding the contract event 0x30309eafe094276dfad42e7ac4676aebcf5ea8480c86708ad2a2c9c0335c912e.
//
// Solidity: event SynthPriceDataUpdated(uint256 indexed synthMarketId, bytes32 indexed buyFeedId, bytes32 indexed sellFeedId)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSynthPriceDataUpdated(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSynthPriceDataUpdated, synthMarketId []*big.Int, buyFeedId [][32]byte, sellFeedId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SynthPriceDataUpdated", synthMarketIdRule, buyFeedIdRule, sellFeedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSynthPriceDataUpdated)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthPriceDataUpdated", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSynthPriceDataUpdated(log types.Log) (*SpotMarketGoerliSynthPriceDataUpdated, error) {
	event := new(SpotMarketGoerliSynthPriceDataUpdated)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthPriceDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSynthRegisteredIterator is returned from FilterSynthRegistered and is used to iterate over the raw logs and unpacked data for SynthRegistered events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthRegisteredIterator struct {
	Event *SpotMarketGoerliSynthRegistered // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSynthRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSynthRegistered)
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
		it.Event = new(SpotMarketGoerliSynthRegistered)
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
func (it *SpotMarketGoerliSynthRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSynthRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSynthRegistered represents a SynthRegistered event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthRegistered struct {
	SynthMarketId     *big.Int
	SynthTokenAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSynthRegistered is a free log retrieval operation binding the contract event 0x04b07b1c236913894927915a3dd91c8e250223fc668ceabdc47074c5a77e2cb9.
//
// Solidity: event SynthRegistered(uint256 indexed synthMarketId, address synthTokenAddress)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSynthRegistered(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliSynthRegisteredIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SynthRegistered", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSynthRegisteredIterator{contract: _SpotMarketGoerli.contract, event: "SynthRegistered", logs: logs, sub: sub}, nil
}

// WatchSynthRegistered is a free log subscription operation binding the contract event 0x04b07b1c236913894927915a3dd91c8e250223fc668ceabdc47074c5a77e2cb9.
//
// Solidity: event SynthRegistered(uint256 indexed synthMarketId, address synthTokenAddress)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSynthRegistered(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSynthRegistered, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SynthRegistered", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSynthRegistered)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthRegistered", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSynthRegistered(log types.Log) (*SpotMarketGoerliSynthRegistered, error) {
	event := new(SpotMarketGoerliSynthRegistered)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSynthSoldIterator is returned from FilterSynthSold and is used to iterate over the raw logs and unpacked data for SynthSold events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthSoldIterator struct {
	Event *SpotMarketGoerliSynthSold // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSynthSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSynthSold)
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
		it.Event = new(SpotMarketGoerliSynthSold)
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
func (it *SpotMarketGoerliSynthSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSynthSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSynthSold represents a SynthSold event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthSold struct {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSynthSold(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliSynthSoldIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SynthSold", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSynthSoldIterator{contract: _SpotMarketGoerli.contract, event: "SynthSold", logs: logs, sub: sub}, nil
}

// WatchSynthSold is a free log subscription operation binding the contract event 0x61fa4bb370a2f18a502b3bcf1d0755e53371d58791fa42766aa6386bbefb594a.
//
// Solidity: event SynthSold(uint256 indexed synthMarketId, uint256 amountReturned, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address referrer, uint256 price)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSynthSold(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSynthSold, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SynthSold", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSynthSold)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthSold", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSynthSold(log types.Log) (*SpotMarketGoerliSynthSold, error) {
	event := new(SpotMarketGoerliSynthSold)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSynthUnwrappedIterator is returned from FilterSynthUnwrapped and is used to iterate over the raw logs and unpacked data for SynthUnwrapped events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthUnwrappedIterator struct {
	Event *SpotMarketGoerliSynthUnwrapped // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSynthUnwrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSynthUnwrapped)
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
		it.Event = new(SpotMarketGoerliSynthUnwrapped)
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
func (it *SpotMarketGoerliSynthUnwrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSynthUnwrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSynthUnwrapped represents a SynthUnwrapped event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthUnwrapped struct {
	SynthMarketId   *big.Int
	AmountUnwrapped *big.Int
	Fees            OrderFeesData
	FeesCollected   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSynthUnwrapped is a free log retrieval operation binding the contract event 0xa1dd74fb936c7942732e4355961ca6944ca6c6744121ace0d9a1203d664231b3.
//
// Solidity: event SynthUnwrapped(uint256 indexed synthMarketId, uint256 amountUnwrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSynthUnwrapped(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliSynthUnwrappedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SynthUnwrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSynthUnwrappedIterator{contract: _SpotMarketGoerli.contract, event: "SynthUnwrapped", logs: logs, sub: sub}, nil
}

// WatchSynthUnwrapped is a free log subscription operation binding the contract event 0xa1dd74fb936c7942732e4355961ca6944ca6c6744121ace0d9a1203d664231b3.
//
// Solidity: event SynthUnwrapped(uint256 indexed synthMarketId, uint256 amountUnwrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSynthUnwrapped(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSynthUnwrapped, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SynthUnwrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSynthUnwrapped)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthUnwrapped", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSynthUnwrapped(log types.Log) (*SpotMarketGoerliSynthUnwrapped, error) {
	event := new(SpotMarketGoerliSynthUnwrapped)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthUnwrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSynthWrappedIterator is returned from FilterSynthWrapped and is used to iterate over the raw logs and unpacked data for SynthWrapped events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthWrappedIterator struct {
	Event *SpotMarketGoerliSynthWrapped // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSynthWrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSynthWrapped)
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
		it.Event = new(SpotMarketGoerliSynthWrapped)
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
func (it *SpotMarketGoerliSynthWrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSynthWrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSynthWrapped represents a SynthWrapped event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthWrapped struct {
	SynthMarketId *big.Int
	AmountWrapped *big.Int
	Fees          OrderFeesData
	FeesCollected *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSynthWrapped is a free log retrieval operation binding the contract event 0xea50ab2f37d37692441c4a16317c1287bb410a3a616a16c49c9ca76d415667ff.
//
// Solidity: event SynthWrapped(uint256 indexed synthMarketId, uint256 amountWrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSynthWrapped(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliSynthWrappedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SynthWrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSynthWrappedIterator{contract: _SpotMarketGoerli.contract, event: "SynthWrapped", logs: logs, sub: sub}, nil
}

// WatchSynthWrapped is a free log subscription operation binding the contract event 0xea50ab2f37d37692441c4a16317c1287bb410a3a616a16c49c9ca76d415667ff.
//
// Solidity: event SynthWrapped(uint256 indexed synthMarketId, uint256 amountWrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSynthWrapped(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSynthWrapped, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SynthWrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSynthWrapped)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthWrapped", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSynthWrapped(log types.Log) (*SpotMarketGoerliSynthWrapped, error) {
	event := new(SpotMarketGoerliSynthWrapped)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthWrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliSynthetixSystemSetIterator is returned from FilterSynthetixSystemSet and is used to iterate over the raw logs and unpacked data for SynthetixSystemSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthetixSystemSetIterator struct {
	Event *SpotMarketGoerliSynthetixSystemSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliSynthetixSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliSynthetixSystemSet)
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
		it.Event = new(SpotMarketGoerliSynthetixSystemSet)
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
func (it *SpotMarketGoerliSynthetixSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliSynthetixSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliSynthetixSystemSet represents a SynthetixSystemSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliSynthetixSystemSet struct {
	Synthetix       common.Address
	UsdTokenAddress common.Address
	OracleManager   common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSynthetixSystemSet is a free log retrieval operation binding the contract event 0x52dccf7e2653d5ae8cf1d18c5499fd530f01920181d81afdf6bf489d897e24aa.
//
// Solidity: event SynthetixSystemSet(address synthetix, address usdTokenAddress, address oracleManager)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterSynthetixSystemSet(opts *bind.FilterOpts) (*SpotMarketGoerliSynthetixSystemSetIterator, error) {

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "SynthetixSystemSet")
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliSynthetixSystemSetIterator{contract: _SpotMarketGoerli.contract, event: "SynthetixSystemSet", logs: logs, sub: sub}, nil
}

// WatchSynthetixSystemSet is a free log subscription operation binding the contract event 0x52dccf7e2653d5ae8cf1d18c5499fd530f01920181d81afdf6bf489d897e24aa.
//
// Solidity: event SynthetixSystemSet(address synthetix, address usdTokenAddress, address oracleManager)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchSynthetixSystemSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliSynthetixSystemSet) (event.Subscription, error) {

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "SynthetixSystemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliSynthetixSystemSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthetixSystemSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseSynthetixSystemSet(log types.Log) (*SpotMarketGoerliSynthetixSystemSet, error) {
	event := new(SpotMarketGoerliSynthetixSystemSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "SynthetixSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliTransactorFixedFeeSetIterator is returned from FilterTransactorFixedFeeSet and is used to iterate over the raw logs and unpacked data for TransactorFixedFeeSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliTransactorFixedFeeSetIterator struct {
	Event *SpotMarketGoerliTransactorFixedFeeSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliTransactorFixedFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliTransactorFixedFeeSet)
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
		it.Event = new(SpotMarketGoerliTransactorFixedFeeSet)
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
func (it *SpotMarketGoerliTransactorFixedFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliTransactorFixedFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliTransactorFixedFeeSet represents a TransactorFixedFeeSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliTransactorFixedFeeSet struct {
	SynthMarketId  *big.Int
	Transactor     common.Address
	FixedFeeAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransactorFixedFeeSet is a free log retrieval operation binding the contract event 0xeed7c7ebc4a7e7456a5b14f961bbe55d026f35a2a2b52d1ad43fe11c348df24a.
//
// Solidity: event TransactorFixedFeeSet(uint256 indexed synthMarketId, address transactor, uint256 fixedFeeAmount)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterTransactorFixedFeeSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliTransactorFixedFeeSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "TransactorFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliTransactorFixedFeeSetIterator{contract: _SpotMarketGoerli.contract, event: "TransactorFixedFeeSet", logs: logs, sub: sub}, nil
}

// WatchTransactorFixedFeeSet is a free log subscription operation binding the contract event 0xeed7c7ebc4a7e7456a5b14f961bbe55d026f35a2a2b52d1ad43fe11c348df24a.
//
// Solidity: event TransactorFixedFeeSet(uint256 indexed synthMarketId, address transactor, uint256 fixedFeeAmount)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchTransactorFixedFeeSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliTransactorFixedFeeSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "TransactorFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliTransactorFixedFeeSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "TransactorFixedFeeSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseTransactorFixedFeeSet(log types.Log) (*SpotMarketGoerliTransactorFixedFeeSet, error) {
	event := new(SpotMarketGoerliTransactorFixedFeeSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "TransactorFixedFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliUpgradedIterator struct {
	Event *SpotMarketGoerliUpgraded // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliUpgraded)
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
		it.Event = new(SpotMarketGoerliUpgraded)
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
func (it *SpotMarketGoerliUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliUpgraded represents a Upgraded event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliUpgraded struct {
	Self           common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterUpgraded(opts *bind.FilterOpts, self []common.Address) (*SpotMarketGoerliUpgradedIterator, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliUpgradedIterator{contract: _SpotMarketGoerli.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliUpgraded, self []common.Address) (event.Subscription, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliUpgraded)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseUpgraded(log types.Log) (*SpotMarketGoerliUpgraded, error) {
	event := new(SpotMarketGoerliUpgraded)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliWrapperFeesSetIterator is returned from FilterWrapperFeesSet and is used to iterate over the raw logs and unpacked data for WrapperFeesSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliWrapperFeesSetIterator struct {
	Event *SpotMarketGoerliWrapperFeesSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliWrapperFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliWrapperFeesSet)
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
		it.Event = new(SpotMarketGoerliWrapperFeesSet)
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
func (it *SpotMarketGoerliWrapperFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliWrapperFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliWrapperFeesSet represents a WrapperFeesSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliWrapperFeesSet struct {
	SynthMarketId *big.Int
	WrapFee       *big.Int
	UnwrapFee     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWrapperFeesSet is a free log retrieval operation binding the contract event 0x84c5bc20d6e52e92afe6ebc9d85d3e4d35de276ba3f05cae640db053f5b861b8.
//
// Solidity: event WrapperFeesSet(uint256 indexed synthMarketId, int256 wrapFee, int256 unwrapFee)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterWrapperFeesSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*SpotMarketGoerliWrapperFeesSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "WrapperFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliWrapperFeesSetIterator{contract: _SpotMarketGoerli.contract, event: "WrapperFeesSet", logs: logs, sub: sub}, nil
}

// WatchWrapperFeesSet is a free log subscription operation binding the contract event 0x84c5bc20d6e52e92afe6ebc9d85d3e4d35de276ba3f05cae640db053f5b861b8.
//
// Solidity: event WrapperFeesSet(uint256 indexed synthMarketId, int256 wrapFee, int256 unwrapFee)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchWrapperFeesSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliWrapperFeesSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "WrapperFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliWrapperFeesSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "WrapperFeesSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseWrapperFeesSet(log types.Log) (*SpotMarketGoerliWrapperFeesSet, error) {
	event := new(SpotMarketGoerliWrapperFeesSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "WrapperFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotMarketGoerliWrapperSetIterator is returned from FilterWrapperSet and is used to iterate over the raw logs and unpacked data for WrapperSet events raised by the SpotMarketGoerli contract.
type SpotMarketGoerliWrapperSetIterator struct {
	Event *SpotMarketGoerliWrapperSet // Event containing the contract specifics and raw log

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
func (it *SpotMarketGoerliWrapperSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotMarketGoerliWrapperSet)
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
		it.Event = new(SpotMarketGoerliWrapperSet)
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
func (it *SpotMarketGoerliWrapperSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotMarketGoerliWrapperSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotMarketGoerliWrapperSet represents a WrapperSet event raised by the SpotMarketGoerli contract.
type SpotMarketGoerliWrapperSet struct {
	SynthMarketId      *big.Int
	WrapCollateralType common.Address
	MaxWrappableAmount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWrapperSet is a free log retrieval operation binding the contract event 0xf6b8d296783aecfc5d372dff3e3e802ab63338637f9a2f3e2aae1e745c148def.
//
// Solidity: event WrapperSet(uint256 indexed synthMarketId, address indexed wrapCollateralType, uint256 maxWrappableAmount)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) FilterWrapperSet(opts *bind.FilterOpts, synthMarketId []*big.Int, wrapCollateralType []common.Address) (*SpotMarketGoerliWrapperSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var wrapCollateralTypeRule []interface{}
	for _, wrapCollateralTypeItem := range wrapCollateralType {
		wrapCollateralTypeRule = append(wrapCollateralTypeRule, wrapCollateralTypeItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.FilterLogs(opts, "WrapperSet", synthMarketIdRule, wrapCollateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &SpotMarketGoerliWrapperSetIterator{contract: _SpotMarketGoerli.contract, event: "WrapperSet", logs: logs, sub: sub}, nil
}

// WatchWrapperSet is a free log subscription operation binding the contract event 0xf6b8d296783aecfc5d372dff3e3e802ab63338637f9a2f3e2aae1e745c148def.
//
// Solidity: event WrapperSet(uint256 indexed synthMarketId, address indexed wrapCollateralType, uint256 maxWrappableAmount)
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) WatchWrapperSet(opts *bind.WatchOpts, sink chan<- *SpotMarketGoerliWrapperSet, synthMarketId []*big.Int, wrapCollateralType []common.Address) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var wrapCollateralTypeRule []interface{}
	for _, wrapCollateralTypeItem := range wrapCollateralType {
		wrapCollateralTypeRule = append(wrapCollateralTypeRule, wrapCollateralTypeItem)
	}

	logs, sub, err := _SpotMarketGoerli.contract.WatchLogs(opts, "WrapperSet", synthMarketIdRule, wrapCollateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotMarketGoerliWrapperSet)
				if err := _SpotMarketGoerli.contract.UnpackLog(event, "WrapperSet", log); err != nil {
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
func (_SpotMarketGoerli *SpotMarketGoerliFilterer) ParseWrapperSet(log types.Log) (*SpotMarketGoerliWrapperSet, error) {
	event := new(SpotMarketGoerliWrapperSet)
	if err := _SpotMarketGoerli.contract.UnpackLog(event, "WrapperSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
