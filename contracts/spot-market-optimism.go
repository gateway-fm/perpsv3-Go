// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMarketOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"synthImplementation\",\"type\":\"uint256\"}],\"name\":\"InvalidSynthImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OnlyMarketOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"DecayRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nominee\",\"type\":\"address\"}],\"name\":\"MarketNominationRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"MarketOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"MarketOwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"synthImplementation\",\"type\":\"address\"}],\"name\":\"SynthImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"SynthImplementationUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"buyFeedId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sellFeedId\",\"type\":\"bytes32\"}],\"name\":\"SynthPriceDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"}],\"name\":\"SynthRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"synthetix\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usdTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracleManager\",\"type\":\"address\"}],\"name\":\"SynthetixSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"acceptMarketOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"synthOwner\",\"type\":\"address\"}],\"name\":\"createSynth\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"marketOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getSynth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"synthAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getSynthImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"minimumCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateMarketOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"renounceMarketNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"reportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportedDebtAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setDecayRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"synthImplementation\",\"type\":\"address\"}],\"name\":\"setSynthImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISynthetixSystem\",\"name\":\"synthetix\",\"type\":\"address\"}],\"name\":\"setSynthetix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"buyFeedId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sellFeedId\",\"type\":\"bytes32\"}],\"name\":\"updatePriceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"upgradeSynthImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSynthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"synthAmountCharged\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxSynthAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxUsdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdAmountCharged\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxUsdAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"InsufficientAmountReceived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"InvalidMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrices\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"synthReturned\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SynthBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReturned\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SynthSold\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"buyExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxUsdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"buyExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmountCharged\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"}],\"name\":\"quoteBuyExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"}],\"name\":\"quoteBuyExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmountCharged\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"}],\"name\":\"quoteSellExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"}],\"name\":\"quoteSellExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthToBurn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minUsdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"sell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmountReceived\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"synthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"sellExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSynthAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"sellExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"synthToBurn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"IneligibleForCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InsufficientSharesAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTransaction.Type\",\"name\":\"transactionType\",\"type\":\"uint8\"}],\"name\":\"InvalidAsyncTransactionType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"asyncOrderId\",\"type\":\"uint256\"}],\"name\":\"InvalidClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InvalidCommitmentAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"}],\"name\":\"InvalidSettlementStrategy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"asyncOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"}],\"name\":\"OrderAlreadySettled\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountEscrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structAsyncOrderClaim.Data\",\"name\":\"asyncOrderClaim\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountProvided\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"OrderCommitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountProvided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"commitOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountEscrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrderClaim.Data\",\"name\":\"asyncOrderClaim\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"}],\"name\":\"getAsyncOrderClaim\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumTransaction.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountEscrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementStrategyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumSettlementAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrderClaim.Data\",\"name\":\"asyncOrderClaim\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"}],\"name\":\"InvalidSettlementStrategy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerificationResponse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"MinimumSettlementAmountNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"OutsideSettlementWindow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deviation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tolerance\",\"type\":\"uint256\"}],\"name\":\"PriceDeviationToleranceExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"}],\"name\":\"SettlementStrategyNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalOrderAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"OrderSettled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"asyncOrderId\",\"type\":\"uint128\"}],\"name\":\"settleOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"finalOrderAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"settlePythOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"finalOrderAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"SettlementStrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SettlementStrategyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumUsdExchangeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRoundingLoss\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"}],\"name\":\"addSettlementStrategy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"getSettlementStrategy\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumUsdExchangeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRoundingLoss\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"settlementStrategy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setSettlementStrategyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"InvalidCollateralType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxWrappableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToWrap\",\"type\":\"uint256\"}],\"name\":\"WrapperExceedsMaxAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUnwrapped\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesCollected\",\"type\":\"uint256\"}],\"name\":\"SynthUnwrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWrapped\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesCollected\",\"type\":\"uint256\"}],\"name\":\"SynthWrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wrapCollateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxWrappableAmount\",\"type\":\"uint256\"}],\"name\":\"WrapperSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"wrapCollateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxWrappableAmount\",\"type\":\"uint256\"}],\"name\":\"setWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"unwrapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnCollateralAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"wrapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountReceived\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToMint\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fixedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFees\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skewFees\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"wrapperFees\",\"type\":\"int256\"}],\"internalType\":\"structOrderFees.Data\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"InvalidCollateralLeverage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidFeeCollector\",\"type\":\"address\"}],\"name\":\"InvalidFeeCollectorInterface\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWrapperFees\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"asyncFixedFee\",\"type\":\"uint256\"}],\"name\":\"AsyncFixedFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"atomicFixedFee\",\"type\":\"uint256\"}],\"name\":\"AtomicFixedFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralLeverage\",\"type\":\"uint256\"}],\"name\":\"CollateralLeverageSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"name\":\"MarketSkewScaleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"utilizationFeeRate\",\"type\":\"uint256\"}],\"name\":\"MarketUtilizationFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharePercentage\",\"type\":\"uint256\"}],\"name\":\"ReferrerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fixedFeeAmount\",\"type\":\"uint256\"}],\"name\":\"TransactorFixedFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"synthMarketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"wrapFee\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"unwrapFee\",\"type\":\"int256\"}],\"name\":\"WrapperFeesSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getCollateralLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralLeverage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"transactor\",\"type\":\"address\"}],\"name\":\"getCustomTransactorFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fixedFeeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"atomicFixedFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"asyncFixedFee\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"wrapFee\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"unwrapFee\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketSkewScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMarketUtilizationFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"utilizationFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"getReferrerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharePercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"asyncFixedFee\",\"type\":\"uint256\"}],\"name\":\"setAsyncFixedFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"atomicFixedFee\",\"type\":\"uint256\"}],\"name\":\"setAtomicFixedFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralLeverage\",\"type\":\"uint256\"}],\"name\":\"setCollateralLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"transactor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fixedFeeAmount\",\"type\":\"uint256\"}],\"name\":\"setCustomTransactorFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"name\":\"setMarketSkewScale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"utilizationFeeRate\",\"type\":\"uint256\"}],\"name\":\"setMarketUtilizationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"int256\",\"name\":\"wrapFee\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"unwrapFee\",\"type\":\"int256\"}],\"name\":\"setWrapperFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sharePercentage\",\"type\":\"uint256\"}],\"name\":\"updateReferrerShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_Contracts *ContractsCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_Contracts *ContractsSession) PRECISION() (*big.Int, error) {
	return _Contracts.Contract.PRECISION(&_Contracts.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_Contracts *ContractsCallerSession) PRECISION() (*big.Int, error) {
	return _Contracts.Contract.PRECISION(&_Contracts.CallOpts)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_Contracts *ContractsCaller) GetAssociatedSystem(opts *bind.CallOpts, id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAssociatedSystem", id)

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
func (_Contracts *ContractsSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _Contracts.Contract.GetAssociatedSystem(&_Contracts.CallOpts, id)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_Contracts *ContractsCallerSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _Contracts.Contract.GetAssociatedSystem(&_Contracts.CallOpts, id)
}

// GetAsyncOrderClaim is a free data retrieval call binding the contract method 0x5381ce16.
//
// Solidity: function getAsyncOrderClaim(uint128 marketId, uint128 asyncOrderId) pure returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_Contracts *ContractsCaller) GetAsyncOrderClaim(opts *bind.CallOpts, marketId *big.Int, asyncOrderId *big.Int) (AsyncOrderClaimData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAsyncOrderClaim", marketId, asyncOrderId)

	if err != nil {
		return *new(AsyncOrderClaimData), err
	}

	out0 := *abi.ConvertType(out[0], new(AsyncOrderClaimData)).(*AsyncOrderClaimData)

	return out0, err

}

// GetAsyncOrderClaim is a free data retrieval call binding the contract method 0x5381ce16.
//
// Solidity: function getAsyncOrderClaim(uint128 marketId, uint128 asyncOrderId) pure returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_Contracts *ContractsSession) GetAsyncOrderClaim(marketId *big.Int, asyncOrderId *big.Int) (AsyncOrderClaimData, error) {
	return _Contracts.Contract.GetAsyncOrderClaim(&_Contracts.CallOpts, marketId, asyncOrderId)
}

// GetAsyncOrderClaim is a free data retrieval call binding the contract method 0x5381ce16.
//
// Solidity: function getAsyncOrderClaim(uint128 marketId, uint128 asyncOrderId) pure returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_Contracts *ContractsCallerSession) GetAsyncOrderClaim(marketId *big.Int, asyncOrderId *big.Int) (AsyncOrderClaimData, error) {
	return _Contracts.Contract.GetAsyncOrderClaim(&_Contracts.CallOpts, marketId, asyncOrderId)
}

// GetCollateralLeverage is a free data retrieval call binding the contract method 0xcdfaef0f.
//
// Solidity: function getCollateralLeverage(uint128 synthMarketId) view returns(uint256 collateralLeverage)
func (_Contracts *ContractsCaller) GetCollateralLeverage(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCollateralLeverage", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralLeverage is a free data retrieval call binding the contract method 0xcdfaef0f.
//
// Solidity: function getCollateralLeverage(uint128 synthMarketId) view returns(uint256 collateralLeverage)
func (_Contracts *ContractsSession) GetCollateralLeverage(synthMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetCollateralLeverage(&_Contracts.CallOpts, synthMarketId)
}

// GetCollateralLeverage is a free data retrieval call binding the contract method 0xcdfaef0f.
//
// Solidity: function getCollateralLeverage(uint128 synthMarketId) view returns(uint256 collateralLeverage)
func (_Contracts *ContractsCallerSession) GetCollateralLeverage(synthMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetCollateralLeverage(&_Contracts.CallOpts, synthMarketId)
}

// GetCustomTransactorFees is a free data retrieval call binding the contract method 0x2efaa971.
//
// Solidity: function getCustomTransactorFees(uint128 synthMarketId, address transactor) view returns(uint256 fixedFeeAmount)
func (_Contracts *ContractsCaller) GetCustomTransactorFees(opts *bind.CallOpts, synthMarketId *big.Int, transactor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCustomTransactorFees", synthMarketId, transactor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCustomTransactorFees is a free data retrieval call binding the contract method 0x2efaa971.
//
// Solidity: function getCustomTransactorFees(uint128 synthMarketId, address transactor) view returns(uint256 fixedFeeAmount)
func (_Contracts *ContractsSession) GetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetCustomTransactorFees(&_Contracts.CallOpts, synthMarketId, transactor)
}

// GetCustomTransactorFees is a free data retrieval call binding the contract method 0x2efaa971.
//
// Solidity: function getCustomTransactorFees(uint128 synthMarketId, address transactor) view returns(uint256 fixedFeeAmount)
func (_Contracts *ContractsCallerSession) GetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetCustomTransactorFees(&_Contracts.CallOpts, synthMarketId, transactor)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_Contracts *ContractsCaller) GetDeniers(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getDeniers", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_Contracts *ContractsSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _Contracts.Contract.GetDeniers(&_Contracts.CallOpts, feature)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_Contracts *ContractsCallerSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _Contracts.Contract.GetDeniers(&_Contracts.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_Contracts *ContractsCaller) GetFeatureFlagAllowAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getFeatureFlagAllowAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_Contracts *ContractsSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _Contracts.Contract.GetFeatureFlagAllowAll(&_Contracts.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_Contracts *ContractsCallerSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _Contracts.Contract.GetFeatureFlagAllowAll(&_Contracts.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_Contracts *ContractsCaller) GetFeatureFlagAllowlist(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getFeatureFlagAllowlist", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_Contracts *ContractsSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _Contracts.Contract.GetFeatureFlagAllowlist(&_Contracts.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_Contracts *ContractsCallerSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _Contracts.Contract.GetFeatureFlagAllowlist(&_Contracts.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_Contracts *ContractsCaller) GetFeatureFlagDenyAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getFeatureFlagDenyAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_Contracts *ContractsSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _Contracts.Contract.GetFeatureFlagDenyAll(&_Contracts.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_Contracts *ContractsCallerSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _Contracts.Contract.GetFeatureFlagDenyAll(&_Contracts.CallOpts, feature)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x5497eb23.
//
// Solidity: function getFeeCollector(uint128 synthMarketId) view returns(address feeCollector)
func (_Contracts *ContractsCaller) GetFeeCollector(opts *bind.CallOpts, synthMarketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getFeeCollector", synthMarketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeCollector is a free data retrieval call binding the contract method 0x5497eb23.
//
// Solidity: function getFeeCollector(uint128 synthMarketId) view returns(address feeCollector)
func (_Contracts *ContractsSession) GetFeeCollector(synthMarketId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetFeeCollector(&_Contracts.CallOpts, synthMarketId)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x5497eb23.
//
// Solidity: function getFeeCollector(uint128 synthMarketId) view returns(address feeCollector)
func (_Contracts *ContractsCallerSession) GetFeeCollector(synthMarketId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetFeeCollector(&_Contracts.CallOpts, synthMarketId)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_Contracts *ContractsCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_Contracts *ContractsSession) GetImplementation() (common.Address, error) {
	return _Contracts.Contract.GetImplementation(&_Contracts.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_Contracts *ContractsCallerSession) GetImplementation() (common.Address, error) {
	return _Contracts.Contract.GetImplementation(&_Contracts.CallOpts)
}

// GetMarketFees is a free data retrieval call binding the contract method 0x32598e61.
//
// Solidity: function getMarketFees(uint128 synthMarketId) view returns(uint256 atomicFixedFee, uint256 asyncFixedFee, int256 wrapFee, int256 unwrapFee)
func (_Contracts *ContractsCaller) GetMarketFees(opts *bind.CallOpts, synthMarketId *big.Int) (struct {
	AtomicFixedFee *big.Int
	AsyncFixedFee  *big.Int
	WrapFee        *big.Int
	UnwrapFee      *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketFees", synthMarketId)

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
func (_Contracts *ContractsSession) GetMarketFees(synthMarketId *big.Int) (struct {
	AtomicFixedFee *big.Int
	AsyncFixedFee  *big.Int
	WrapFee        *big.Int
	UnwrapFee      *big.Int
}, error) {
	return _Contracts.Contract.GetMarketFees(&_Contracts.CallOpts, synthMarketId)
}

// GetMarketFees is a free data retrieval call binding the contract method 0x32598e61.
//
// Solidity: function getMarketFees(uint128 synthMarketId) view returns(uint256 atomicFixedFee, uint256 asyncFixedFee, int256 wrapFee, int256 unwrapFee)
func (_Contracts *ContractsCallerSession) GetMarketFees(synthMarketId *big.Int) (struct {
	AtomicFixedFee *big.Int
	AsyncFixedFee  *big.Int
	WrapFee        *big.Int
	UnwrapFee      *big.Int
}, error) {
	return _Contracts.Contract.GetMarketFees(&_Contracts.CallOpts, synthMarketId)
}

// GetMarketOwner is a free data retrieval call binding the contract method 0xa7b8cb9f.
//
// Solidity: function getMarketOwner(uint128 synthMarketId) view returns(address marketOwner)
func (_Contracts *ContractsCaller) GetMarketOwner(opts *bind.CallOpts, synthMarketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketOwner", synthMarketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMarketOwner is a free data retrieval call binding the contract method 0xa7b8cb9f.
//
// Solidity: function getMarketOwner(uint128 synthMarketId) view returns(address marketOwner)
func (_Contracts *ContractsSession) GetMarketOwner(synthMarketId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetMarketOwner(&_Contracts.CallOpts, synthMarketId)
}

// GetMarketOwner is a free data retrieval call binding the contract method 0xa7b8cb9f.
//
// Solidity: function getMarketOwner(uint128 synthMarketId) view returns(address marketOwner)
func (_Contracts *ContractsCallerSession) GetMarketOwner(synthMarketId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetMarketOwner(&_Contracts.CallOpts, synthMarketId)
}

// GetMarketSkewScale is a free data retrieval call binding the contract method 0x8d105571.
//
// Solidity: function getMarketSkewScale(uint128 synthMarketId) view returns(uint256 skewScale)
func (_Contracts *ContractsCaller) GetMarketSkewScale(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketSkewScale", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketSkewScale is a free data retrieval call binding the contract method 0x8d105571.
//
// Solidity: function getMarketSkewScale(uint128 synthMarketId) view returns(uint256 skewScale)
func (_Contracts *ContractsSession) GetMarketSkewScale(synthMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketSkewScale(&_Contracts.CallOpts, synthMarketId)
}

// GetMarketSkewScale is a free data retrieval call binding the contract method 0x8d105571.
//
// Solidity: function getMarketSkewScale(uint128 synthMarketId) view returns(uint256 skewScale)
func (_Contracts *ContractsCallerSession) GetMarketSkewScale(synthMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketSkewScale(&_Contracts.CallOpts, synthMarketId)
}

// GetMarketUtilizationFees is a free data retrieval call binding the contract method 0xf375f324.
//
// Solidity: function getMarketUtilizationFees(uint128 synthMarketId) view returns(uint256 utilizationFeeRate)
func (_Contracts *ContractsCaller) GetMarketUtilizationFees(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketUtilizationFees", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketUtilizationFees is a free data retrieval call binding the contract method 0xf375f324.
//
// Solidity: function getMarketUtilizationFees(uint128 synthMarketId) view returns(uint256 utilizationFeeRate)
func (_Contracts *ContractsSession) GetMarketUtilizationFees(synthMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketUtilizationFees(&_Contracts.CallOpts, synthMarketId)
}

// GetMarketUtilizationFees is a free data retrieval call binding the contract method 0xf375f324.
//
// Solidity: function getMarketUtilizationFees(uint128 synthMarketId) view returns(uint256 utilizationFeeRate)
func (_Contracts *ContractsCallerSession) GetMarketUtilizationFees(synthMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketUtilizationFees(&_Contracts.CallOpts, synthMarketId)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xfa4b28ed.
//
// Solidity: function getReferrerShare(uint128 synthMarketId, address referrer) view returns(uint256 sharePercentage)
func (_Contracts *ContractsCaller) GetReferrerShare(opts *bind.CallOpts, synthMarketId *big.Int, referrer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getReferrerShare", synthMarketId, referrer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReferrerShare is a free data retrieval call binding the contract method 0xfa4b28ed.
//
// Solidity: function getReferrerShare(uint128 synthMarketId, address referrer) view returns(uint256 sharePercentage)
func (_Contracts *ContractsSession) GetReferrerShare(synthMarketId *big.Int, referrer common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetReferrerShare(&_Contracts.CallOpts, synthMarketId, referrer)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xfa4b28ed.
//
// Solidity: function getReferrerShare(uint128 synthMarketId, address referrer) view returns(uint256 sharePercentage)
func (_Contracts *ContractsCallerSession) GetReferrerShare(synthMarketId *big.Int, referrer common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetReferrerShare(&_Contracts.CallOpts, synthMarketId, referrer)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) settlementStrategy)
func (_Contracts *ContractsCaller) GetSettlementStrategy(opts *bind.CallOpts, marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getSettlementStrategy", marketId, strategyId)

	if err != nil {
		return *new(SettlementStrategyData), err
	}

	out0 := *abi.ConvertType(out[0], new(SettlementStrategyData)).(*SettlementStrategyData)

	return out0, err

}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) settlementStrategy)
func (_Contracts *ContractsSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _Contracts.Contract.GetSettlementStrategy(&_Contracts.CallOpts, marketId, strategyId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) settlementStrategy)
func (_Contracts *ContractsCallerSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _Contracts.Contract.GetSettlementStrategy(&_Contracts.CallOpts, marketId, strategyId)
}

// GetSynth is a free data retrieval call binding the contract method 0x69e0365f.
//
// Solidity: function getSynth(uint128 marketId) view returns(address synthAddress)
func (_Contracts *ContractsCaller) GetSynth(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getSynth", marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSynth is a free data retrieval call binding the contract method 0x69e0365f.
//
// Solidity: function getSynth(uint128 marketId) view returns(address synthAddress)
func (_Contracts *ContractsSession) GetSynth(marketId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetSynth(&_Contracts.CallOpts, marketId)
}

// GetSynth is a free data retrieval call binding the contract method 0x69e0365f.
//
// Solidity: function getSynth(uint128 marketId) view returns(address synthAddress)
func (_Contracts *ContractsCallerSession) GetSynth(marketId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetSynth(&_Contracts.CallOpts, marketId)
}

// GetSynthImpl is a free data retrieval call binding the contract method 0x3e0c76ca.
//
// Solidity: function getSynthImpl(uint128 marketId) view returns(address implAddress)
func (_Contracts *ContractsCaller) GetSynthImpl(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getSynthImpl", marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSynthImpl is a free data retrieval call binding the contract method 0x3e0c76ca.
//
// Solidity: function getSynthImpl(uint128 marketId) view returns(address implAddress)
func (_Contracts *ContractsSession) GetSynthImpl(marketId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetSynthImpl(&_Contracts.CallOpts, marketId)
}

// GetSynthImpl is a free data retrieval call binding the contract method 0x3e0c76ca.
//
// Solidity: function getSynthImpl(uint128 marketId) view returns(address implAddress)
func (_Contracts *ContractsCallerSession) GetSynthImpl(marketId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetSynthImpl(&_Contracts.CallOpts, marketId)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_Contracts *ContractsCaller) IsFeatureAllowed(opts *bind.CallOpts, feature [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isFeatureAllowed", feature, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_Contracts *ContractsSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _Contracts.Contract.IsFeatureAllowed(&_Contracts.CallOpts, feature, account)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_Contracts *ContractsCallerSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _Contracts.Contract.IsFeatureAllowed(&_Contracts.CallOpts, feature, account)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 marketId) view returns(uint256 lockedAmount)
func (_Contracts *ContractsCaller) MinimumCredit(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "minimumCredit", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 marketId) view returns(uint256 lockedAmount)
func (_Contracts *ContractsSession) MinimumCredit(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.MinimumCredit(&_Contracts.CallOpts, marketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 marketId) view returns(uint256 lockedAmount)
func (_Contracts *ContractsCallerSession) MinimumCredit(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.MinimumCredit(&_Contracts.CallOpts, marketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 marketId) view returns(string marketName)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts, marketId *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name", marketId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 marketId) view returns(string marketName)
func (_Contracts *ContractsSession) Name(marketId *big.Int) (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts, marketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 marketId) view returns(string marketName)
func (_Contracts *ContractsCallerSession) Name(marketId *big.Int) (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts, marketId)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Contracts *ContractsCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Contracts *ContractsSession) NominatedOwner() (common.Address, error) {
	return _Contracts.Contract.NominatedOwner(&_Contracts.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Contracts *ContractsCallerSession) NominatedOwner() (common.Address, error) {
	return _Contracts.Contract.NominatedOwner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// QuoteBuyExactIn is a free data retrieval call binding the contract method 0x6b5e6ae4.
//
// Solidity: function quoteBuyExactIn(uint128 marketId, uint256 usdAmount) view returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsCaller) QuoteBuyExactIn(opts *bind.CallOpts, marketId *big.Int, usdAmount *big.Int) (struct {
	SynthAmount *big.Int
	Fees        OrderFeesData
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quoteBuyExactIn", marketId, usdAmount)

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
func (_Contracts *ContractsSession) QuoteBuyExactIn(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthAmount *big.Int
	Fees        OrderFeesData
}, error) {
	return _Contracts.Contract.QuoteBuyExactIn(&_Contracts.CallOpts, marketId, usdAmount)
}

// QuoteBuyExactIn is a free data retrieval call binding the contract method 0x6b5e6ae4.
//
// Solidity: function quoteBuyExactIn(uint128 marketId, uint256 usdAmount) view returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsCallerSession) QuoteBuyExactIn(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthAmount *big.Int
	Fees        OrderFeesData
}, error) {
	return _Contracts.Contract.QuoteBuyExactIn(&_Contracts.CallOpts, marketId, usdAmount)
}

// QuoteBuyExactOut is a free data retrieval call binding the contract method 0x1f3f7640.
//
// Solidity: function quoteBuyExactOut(uint128 marketId, uint256 synthAmount) view returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsCaller) QuoteBuyExactOut(opts *bind.CallOpts, marketId *big.Int, synthAmount *big.Int) (struct {
	UsdAmountCharged *big.Int
	Fees             OrderFeesData
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quoteBuyExactOut", marketId, synthAmount)

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
func (_Contracts *ContractsSession) QuoteBuyExactOut(marketId *big.Int, synthAmount *big.Int) (struct {
	UsdAmountCharged *big.Int
	Fees             OrderFeesData
}, error) {
	return _Contracts.Contract.QuoteBuyExactOut(&_Contracts.CallOpts, marketId, synthAmount)
}

// QuoteBuyExactOut is a free data retrieval call binding the contract method 0x1f3f7640.
//
// Solidity: function quoteBuyExactOut(uint128 marketId, uint256 synthAmount) view returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsCallerSession) QuoteBuyExactOut(marketId *big.Int, synthAmount *big.Int) (struct {
	UsdAmountCharged *big.Int
	Fees             OrderFeesData
}, error) {
	return _Contracts.Contract.QuoteBuyExactOut(&_Contracts.CallOpts, marketId, synthAmount)
}

// QuoteSellExactIn is a free data retrieval call binding the contract method 0xc52d1730.
//
// Solidity: function quoteSellExactIn(uint128 marketId, uint256 synthAmount) view returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsCaller) QuoteSellExactIn(opts *bind.CallOpts, marketId *big.Int, synthAmount *big.Int) (struct {
	ReturnAmount *big.Int
	Fees         OrderFeesData
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quoteSellExactIn", marketId, synthAmount)

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
func (_Contracts *ContractsSession) QuoteSellExactIn(marketId *big.Int, synthAmount *big.Int) (struct {
	ReturnAmount *big.Int
	Fees         OrderFeesData
}, error) {
	return _Contracts.Contract.QuoteSellExactIn(&_Contracts.CallOpts, marketId, synthAmount)
}

// QuoteSellExactIn is a free data retrieval call binding the contract method 0xc52d1730.
//
// Solidity: function quoteSellExactIn(uint128 marketId, uint256 synthAmount) view returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsCallerSession) QuoteSellExactIn(marketId *big.Int, synthAmount *big.Int) (struct {
	ReturnAmount *big.Int
	Fees         OrderFeesData
}, error) {
	return _Contracts.Contract.QuoteSellExactIn(&_Contracts.CallOpts, marketId, synthAmount)
}

// QuoteSellExactOut is a free data retrieval call binding the contract method 0x2c007522.
//
// Solidity: function quoteSellExactOut(uint128 marketId, uint256 usdAmount) view returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsCaller) QuoteSellExactOut(opts *bind.CallOpts, marketId *big.Int, usdAmount *big.Int) (struct {
	SynthToBurn *big.Int
	Fees        OrderFeesData
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quoteSellExactOut", marketId, usdAmount)

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
func (_Contracts *ContractsSession) QuoteSellExactOut(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthToBurn *big.Int
	Fees        OrderFeesData
}, error) {
	return _Contracts.Contract.QuoteSellExactOut(&_Contracts.CallOpts, marketId, usdAmount)
}

// QuoteSellExactOut is a free data retrieval call binding the contract method 0x2c007522.
//
// Solidity: function quoteSellExactOut(uint128 marketId, uint256 usdAmount) view returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsCallerSession) QuoteSellExactOut(marketId *big.Int, usdAmount *big.Int) (struct {
	SynthToBurn *big.Int
	Fees        OrderFeesData
}, error) {
	return _Contracts.Contract.QuoteSellExactOut(&_Contracts.CallOpts, marketId, usdAmount)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 marketId) view returns(uint256 reportedDebtAmount)
func (_Contracts *ContractsCaller) ReportedDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "reportedDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 marketId) view returns(uint256 reportedDebtAmount)
func (_Contracts *ContractsSession) ReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ReportedDebt(&_Contracts.CallOpts, marketId)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 marketId) view returns(uint256 reportedDebtAmount)
func (_Contracts *ContractsCallerSession) ReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ReportedDebt(&_Contracts.CallOpts, marketId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool isSupported)
func (_Contracts *ContractsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool isSupported)
func (_Contracts *ContractsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool isSupported)
func (_Contracts *ContractsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// AcceptMarketOwnership is a paid mutator transaction binding the contract method 0x1c216a0e.
//
// Solidity: function acceptMarketOwnership(uint128 synthMarketId) returns()
func (_Contracts *ContractsTransactor) AcceptMarketOwnership(opts *bind.TransactOpts, synthMarketId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "acceptMarketOwnership", synthMarketId)
}

// AcceptMarketOwnership is a paid mutator transaction binding the contract method 0x1c216a0e.
//
// Solidity: function acceptMarketOwnership(uint128 synthMarketId) returns()
func (_Contracts *ContractsSession) AcceptMarketOwnership(synthMarketId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AcceptMarketOwnership(&_Contracts.TransactOpts, synthMarketId)
}

// AcceptMarketOwnership is a paid mutator transaction binding the contract method 0x1c216a0e.
//
// Solidity: function acceptMarketOwnership(uint128 synthMarketId) returns()
func (_Contracts *ContractsTransactorSession) AcceptMarketOwnership(synthMarketId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AcceptMarketOwnership(&_Contracts.TransactOpts, synthMarketId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contracts *ContractsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contracts *ContractsSession) AcceptOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.AcceptOwnership(&_Contracts.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contracts *ContractsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.AcceptOwnership(&_Contracts.TransactOpts)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x97b30e6d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_Contracts *ContractsTransactor) AddSettlementStrategy(opts *bind.TransactOpts, marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addSettlementStrategy", marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x97b30e6d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_Contracts *ContractsSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _Contracts.Contract.AddSettlementStrategy(&_Contracts.TransactOpts, marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x97b30e6d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,address,bytes32,string,uint256,uint256,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_Contracts *ContractsTransactorSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _Contracts.Contract.AddSettlementStrategy(&_Contracts.TransactOpts, marketId, strategy)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Contracts *ContractsTransactor) AddToFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addToFeatureFlagAllowlist", feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Contracts *ContractsSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddToFeatureFlagAllowlist(&_Contracts.TransactOpts, feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Contracts *ContractsTransactorSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddToFeatureFlagAllowlist(&_Contracts.TransactOpts, feature, account)
}

// Buy is a paid mutator transaction binding the contract method 0x37fb3369.
//
// Solidity: function buy(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactor) Buy(opts *bind.TransactOpts, marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "buy", marketId, usdAmount, minAmountReceived, referrer)
}

// Buy is a paid mutator transaction binding the contract method 0x37fb3369.
//
// Solidity: function buy(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsSession) Buy(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Buy(&_Contracts.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// Buy is a paid mutator transaction binding the contract method 0x37fb3369.
//
// Solidity: function buy(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactorSession) Buy(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Buy(&_Contracts.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactIn is a paid mutator transaction binding the contract method 0xa12d9400.
//
// Solidity: function buyExactIn(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactor) BuyExactIn(opts *bind.TransactOpts, marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "buyExactIn", marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactIn is a paid mutator transaction binding the contract method 0xa12d9400.
//
// Solidity: function buyExactIn(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsSession) BuyExactIn(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.BuyExactIn(&_Contracts.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactIn is a paid mutator transaction binding the contract method 0xa12d9400.
//
// Solidity: function buyExactIn(uint128 marketId, uint256 usdAmount, uint256 minAmountReceived, address referrer) returns(uint256 synthAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactorSession) BuyExactIn(marketId *big.Int, usdAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.BuyExactIn(&_Contracts.TransactOpts, marketId, usdAmount, minAmountReceived, referrer)
}

// BuyExactOut is a paid mutator transaction binding the contract method 0x983220bb.
//
// Solidity: function buyExactOut(uint128 marketId, uint256 synthAmount, uint256 maxUsdAmount, address referrer) returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactor) BuyExactOut(opts *bind.TransactOpts, marketId *big.Int, synthAmount *big.Int, maxUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "buyExactOut", marketId, synthAmount, maxUsdAmount, referrer)
}

// BuyExactOut is a paid mutator transaction binding the contract method 0x983220bb.
//
// Solidity: function buyExactOut(uint128 marketId, uint256 synthAmount, uint256 maxUsdAmount, address referrer) returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsSession) BuyExactOut(marketId *big.Int, synthAmount *big.Int, maxUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.BuyExactOut(&_Contracts.TransactOpts, marketId, synthAmount, maxUsdAmount, referrer)
}

// BuyExactOut is a paid mutator transaction binding the contract method 0x983220bb.
//
// Solidity: function buyExactOut(uint128 marketId, uint256 synthAmount, uint256 maxUsdAmount, address referrer) returns(uint256 usdAmountCharged, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactorSession) BuyExactOut(marketId *big.Int, synthAmount *big.Int, maxUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.BuyExactOut(&_Contracts.TransactOpts, marketId, synthAmount, maxUsdAmount, referrer)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x4138dc53.
//
// Solidity: function cancelOrder(uint128 marketId, uint128 asyncOrderId) returns()
func (_Contracts *ContractsTransactor) CancelOrder(opts *bind.TransactOpts, marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "cancelOrder", marketId, asyncOrderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x4138dc53.
//
// Solidity: function cancelOrder(uint128 marketId, uint128 asyncOrderId) returns()
func (_Contracts *ContractsSession) CancelOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CancelOrder(&_Contracts.TransactOpts, marketId, asyncOrderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x4138dc53.
//
// Solidity: function cancelOrder(uint128 marketId, uint128 asyncOrderId) returns()
func (_Contracts *ContractsTransactorSession) CancelOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CancelOrder(&_Contracts.TransactOpts, marketId, asyncOrderId)
}

// CommitOrder is a paid mutator transaction binding the contract method 0xd393340e.
//
// Solidity: function commitOrder(uint128 marketId, uint8 orderType, uint256 amountProvided, uint256 settlementStrategyId, uint256 minimumSettlementAmount, address referrer) returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_Contracts *ContractsTransactor) CommitOrder(opts *bind.TransactOpts, marketId *big.Int, orderType uint8, amountProvided *big.Int, settlementStrategyId *big.Int, minimumSettlementAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "commitOrder", marketId, orderType, amountProvided, settlementStrategyId, minimumSettlementAmount, referrer)
}

// CommitOrder is a paid mutator transaction binding the contract method 0xd393340e.
//
// Solidity: function commitOrder(uint128 marketId, uint8 orderType, uint256 amountProvided, uint256 settlementStrategyId, uint256 minimumSettlementAmount, address referrer) returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_Contracts *ContractsSession) CommitOrder(marketId *big.Int, orderType uint8, amountProvided *big.Int, settlementStrategyId *big.Int, minimumSettlementAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.CommitOrder(&_Contracts.TransactOpts, marketId, orderType, amountProvided, settlementStrategyId, minimumSettlementAmount, referrer)
}

// CommitOrder is a paid mutator transaction binding the contract method 0xd393340e.
//
// Solidity: function commitOrder(uint128 marketId, uint8 orderType, uint256 amountProvided, uint256 settlementStrategyId, uint256 minimumSettlementAmount, address referrer) returns((uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim)
func (_Contracts *ContractsTransactorSession) CommitOrder(marketId *big.Int, orderType uint8, amountProvided *big.Int, settlementStrategyId *big.Int, minimumSettlementAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.CommitOrder(&_Contracts.TransactOpts, marketId, orderType, amountProvided, settlementStrategyId, minimumSettlementAmount, referrer)
}

// CreateSynth is a paid mutator transaction binding the contract method 0x2e535d61.
//
// Solidity: function createSynth(string tokenName, string tokenSymbol, address synthOwner) returns(uint128 synthMarketId)
func (_Contracts *ContractsTransactor) CreateSynth(opts *bind.TransactOpts, tokenName string, tokenSymbol string, synthOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createSynth", tokenName, tokenSymbol, synthOwner)
}

// CreateSynth is a paid mutator transaction binding the contract method 0x2e535d61.
//
// Solidity: function createSynth(string tokenName, string tokenSymbol, address synthOwner) returns(uint128 synthMarketId)
func (_Contracts *ContractsSession) CreateSynth(tokenName string, tokenSymbol string, synthOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.CreateSynth(&_Contracts.TransactOpts, tokenName, tokenSymbol, synthOwner)
}

// CreateSynth is a paid mutator transaction binding the contract method 0x2e535d61.
//
// Solidity: function createSynth(string tokenName, string tokenSymbol, address synthOwner) returns(uint128 synthMarketId)
func (_Contracts *ContractsTransactorSession) CreateSynth(tokenName string, tokenSymbol string, synthOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.CreateSynth(&_Contracts.TransactOpts, tokenName, tokenSymbol, synthOwner)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_Contracts *ContractsTransactor) InitOrUpgradeNft(opts *bind.TransactOpts, id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initOrUpgradeNft", id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_Contracts *ContractsSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.InitOrUpgradeNft(&_Contracts.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_Contracts *ContractsTransactorSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.InitOrUpgradeNft(&_Contracts.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_Contracts *ContractsTransactor) InitOrUpgradeToken(opts *bind.TransactOpts, id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initOrUpgradeToken", id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_Contracts *ContractsSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.InitOrUpgradeToken(&_Contracts.TransactOpts, id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_Contracts *ContractsTransactorSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.InitOrUpgradeToken(&_Contracts.TransactOpts, id, name, symbol, decimals, impl)
}

// NominateMarketOwner is a paid mutator transaction binding the contract method 0x5950864b.
//
// Solidity: function nominateMarketOwner(uint128 synthMarketId, address newNominatedOwner) returns()
func (_Contracts *ContractsTransactor) NominateMarketOwner(opts *bind.TransactOpts, synthMarketId *big.Int, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "nominateMarketOwner", synthMarketId, newNominatedOwner)
}

// NominateMarketOwner is a paid mutator transaction binding the contract method 0x5950864b.
//
// Solidity: function nominateMarketOwner(uint128 synthMarketId, address newNominatedOwner) returns()
func (_Contracts *ContractsSession) NominateMarketOwner(synthMarketId *big.Int, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.NominateMarketOwner(&_Contracts.TransactOpts, synthMarketId, newNominatedOwner)
}

// NominateMarketOwner is a paid mutator transaction binding the contract method 0x5950864b.
//
// Solidity: function nominateMarketOwner(uint128 synthMarketId, address newNominatedOwner) returns()
func (_Contracts *ContractsTransactorSession) NominateMarketOwner(synthMarketId *big.Int, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.NominateMarketOwner(&_Contracts.TransactOpts, synthMarketId, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_Contracts *ContractsTransactor) NominateNewOwner(opts *bind.TransactOpts, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "nominateNewOwner", newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_Contracts *ContractsSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.NominateNewOwner(&_Contracts.TransactOpts, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_Contracts *ContractsTransactorSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.NominateNewOwner(&_Contracts.TransactOpts, newNominatedOwner)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_Contracts *ContractsTransactor) RegisterUnmanagedSystem(opts *bind.TransactOpts, id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerUnmanagedSystem", id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_Contracts *ContractsSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterUnmanagedSystem(&_Contracts.TransactOpts, id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_Contracts *ContractsTransactorSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterUnmanagedSystem(&_Contracts.TransactOpts, id, endpoint)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Contracts *ContractsTransactor) RemoveFromFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeFromFeatureFlagAllowlist", feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Contracts *ContractsSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveFromFeatureFlagAllowlist(&_Contracts.TransactOpts, feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Contracts *ContractsTransactorSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveFromFeatureFlagAllowlist(&_Contracts.TransactOpts, feature, account)
}

// RenounceMarketNomination is a paid mutator transaction binding the contract method 0x298b26bf.
//
// Solidity: function renounceMarketNomination(uint128 synthMarketId) returns()
func (_Contracts *ContractsTransactor) RenounceMarketNomination(opts *bind.TransactOpts, synthMarketId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceMarketNomination", synthMarketId)
}

// RenounceMarketNomination is a paid mutator transaction binding the contract method 0x298b26bf.
//
// Solidity: function renounceMarketNomination(uint128 synthMarketId) returns()
func (_Contracts *ContractsSession) RenounceMarketNomination(synthMarketId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RenounceMarketNomination(&_Contracts.TransactOpts, synthMarketId)
}

// RenounceMarketNomination is a paid mutator transaction binding the contract method 0x298b26bf.
//
// Solidity: function renounceMarketNomination(uint128 synthMarketId) returns()
func (_Contracts *ContractsTransactorSession) RenounceMarketNomination(synthMarketId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RenounceMarketNomination(&_Contracts.TransactOpts, synthMarketId)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_Contracts *ContractsTransactor) RenounceNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceNomination")
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_Contracts *ContractsSession) RenounceNomination() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceNomination(&_Contracts.TransactOpts)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_Contracts *ContractsTransactorSession) RenounceNomination() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceNomination(&_Contracts.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0x4d4bfbd5.
//
// Solidity: function sell(uint128 marketId, uint256 synthAmount, uint256 minUsdAmount, address referrer) returns(uint256 usdAmountReceived, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactor) Sell(opts *bind.TransactOpts, marketId *big.Int, synthAmount *big.Int, minUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "sell", marketId, synthAmount, minUsdAmount, referrer)
}

// Sell is a paid mutator transaction binding the contract method 0x4d4bfbd5.
//
// Solidity: function sell(uint128 marketId, uint256 synthAmount, uint256 minUsdAmount, address referrer) returns(uint256 usdAmountReceived, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsSession) Sell(marketId *big.Int, synthAmount *big.Int, minUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Sell(&_Contracts.TransactOpts, marketId, synthAmount, minUsdAmount, referrer)
}

// Sell is a paid mutator transaction binding the contract method 0x4d4bfbd5.
//
// Solidity: function sell(uint128 marketId, uint256 synthAmount, uint256 minUsdAmount, address referrer) returns(uint256 usdAmountReceived, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactorSession) Sell(marketId *big.Int, synthAmount *big.Int, minUsdAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Sell(&_Contracts.TransactOpts, marketId, synthAmount, minUsdAmount, referrer)
}

// SellExactIn is a paid mutator transaction binding the contract method 0x3d1a60e4.
//
// Solidity: function sellExactIn(uint128 marketId, uint256 synthAmount, uint256 minAmountReceived, address referrer) returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactor) SellExactIn(opts *bind.TransactOpts, marketId *big.Int, synthAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "sellExactIn", marketId, synthAmount, minAmountReceived, referrer)
}

// SellExactIn is a paid mutator transaction binding the contract method 0x3d1a60e4.
//
// Solidity: function sellExactIn(uint128 marketId, uint256 synthAmount, uint256 minAmountReceived, address referrer) returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsSession) SellExactIn(marketId *big.Int, synthAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SellExactIn(&_Contracts.TransactOpts, marketId, synthAmount, minAmountReceived, referrer)
}

// SellExactIn is a paid mutator transaction binding the contract method 0x3d1a60e4.
//
// Solidity: function sellExactIn(uint128 marketId, uint256 synthAmount, uint256 minAmountReceived, address referrer) returns(uint256 returnAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactorSession) SellExactIn(marketId *big.Int, synthAmount *big.Int, minAmountReceived *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SellExactIn(&_Contracts.TransactOpts, marketId, synthAmount, minAmountReceived, referrer)
}

// SellExactOut is a paid mutator transaction binding the contract method 0x4ce94d9d.
//
// Solidity: function sellExactOut(uint128 marketId, uint256 usdAmount, uint256 maxSynthAmount, address referrer) returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactor) SellExactOut(opts *bind.TransactOpts, marketId *big.Int, usdAmount *big.Int, maxSynthAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "sellExactOut", marketId, usdAmount, maxSynthAmount, referrer)
}

// SellExactOut is a paid mutator transaction binding the contract method 0x4ce94d9d.
//
// Solidity: function sellExactOut(uint128 marketId, uint256 usdAmount, uint256 maxSynthAmount, address referrer) returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsSession) SellExactOut(marketId *big.Int, usdAmount *big.Int, maxSynthAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SellExactOut(&_Contracts.TransactOpts, marketId, usdAmount, maxSynthAmount, referrer)
}

// SellExactOut is a paid mutator transaction binding the contract method 0x4ce94d9d.
//
// Solidity: function sellExactOut(uint128 marketId, uint256 usdAmount, uint256 maxSynthAmount, address referrer) returns(uint256 synthToBurn, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactorSession) SellExactOut(marketId *big.Int, usdAmount *big.Int, maxSynthAmount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SellExactOut(&_Contracts.TransactOpts, marketId, usdAmount, maxSynthAmount, referrer)
}

// SetAsyncFixedFee is a paid mutator transaction binding the contract method 0x61dcca86.
//
// Solidity: function setAsyncFixedFee(uint128 synthMarketId, uint256 asyncFixedFee) returns()
func (_Contracts *ContractsTransactor) SetAsyncFixedFee(opts *bind.TransactOpts, synthMarketId *big.Int, asyncFixedFee *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setAsyncFixedFee", synthMarketId, asyncFixedFee)
}

// SetAsyncFixedFee is a paid mutator transaction binding the contract method 0x61dcca86.
//
// Solidity: function setAsyncFixedFee(uint128 synthMarketId, uint256 asyncFixedFee) returns()
func (_Contracts *ContractsSession) SetAsyncFixedFee(synthMarketId *big.Int, asyncFixedFee *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetAsyncFixedFee(&_Contracts.TransactOpts, synthMarketId, asyncFixedFee)
}

// SetAsyncFixedFee is a paid mutator transaction binding the contract method 0x61dcca86.
//
// Solidity: function setAsyncFixedFee(uint128 synthMarketId, uint256 asyncFixedFee) returns()
func (_Contracts *ContractsTransactorSession) SetAsyncFixedFee(synthMarketId *big.Int, asyncFixedFee *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetAsyncFixedFee(&_Contracts.TransactOpts, synthMarketId, asyncFixedFee)
}

// SetAtomicFixedFee is a paid mutator transaction binding the contract method 0x480be91f.
//
// Solidity: function setAtomicFixedFee(uint128 synthMarketId, uint256 atomicFixedFee) returns()
func (_Contracts *ContractsTransactor) SetAtomicFixedFee(opts *bind.TransactOpts, synthMarketId *big.Int, atomicFixedFee *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setAtomicFixedFee", synthMarketId, atomicFixedFee)
}

// SetAtomicFixedFee is a paid mutator transaction binding the contract method 0x480be91f.
//
// Solidity: function setAtomicFixedFee(uint128 synthMarketId, uint256 atomicFixedFee) returns()
func (_Contracts *ContractsSession) SetAtomicFixedFee(synthMarketId *big.Int, atomicFixedFee *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetAtomicFixedFee(&_Contracts.TransactOpts, synthMarketId, atomicFixedFee)
}

// SetAtomicFixedFee is a paid mutator transaction binding the contract method 0x480be91f.
//
// Solidity: function setAtomicFixedFee(uint128 synthMarketId, uint256 atomicFixedFee) returns()
func (_Contracts *ContractsTransactorSession) SetAtomicFixedFee(synthMarketId *big.Int, atomicFixedFee *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetAtomicFixedFee(&_Contracts.TransactOpts, synthMarketId, atomicFixedFee)
}

// SetCollateralLeverage is a paid mutator transaction binding the contract method 0x21f7f58f.
//
// Solidity: function setCollateralLeverage(uint128 synthMarketId, uint256 collateralLeverage) returns()
func (_Contracts *ContractsTransactor) SetCollateralLeverage(opts *bind.TransactOpts, synthMarketId *big.Int, collateralLeverage *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setCollateralLeverage", synthMarketId, collateralLeverage)
}

// SetCollateralLeverage is a paid mutator transaction binding the contract method 0x21f7f58f.
//
// Solidity: function setCollateralLeverage(uint128 synthMarketId, uint256 collateralLeverage) returns()
func (_Contracts *ContractsSession) SetCollateralLeverage(synthMarketId *big.Int, collateralLeverage *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetCollateralLeverage(&_Contracts.TransactOpts, synthMarketId, collateralLeverage)
}

// SetCollateralLeverage is a paid mutator transaction binding the contract method 0x21f7f58f.
//
// Solidity: function setCollateralLeverage(uint128 synthMarketId, uint256 collateralLeverage) returns()
func (_Contracts *ContractsTransactorSession) SetCollateralLeverage(synthMarketId *big.Int, collateralLeverage *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetCollateralLeverage(&_Contracts.TransactOpts, synthMarketId, collateralLeverage)
}

// SetCustomTransactorFees is a paid mutator transaction binding the contract method 0x95fcd547.
//
// Solidity: function setCustomTransactorFees(uint128 synthMarketId, address transactor, uint256 fixedFeeAmount) returns()
func (_Contracts *ContractsTransactor) SetCustomTransactorFees(opts *bind.TransactOpts, synthMarketId *big.Int, transactor common.Address, fixedFeeAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setCustomTransactorFees", synthMarketId, transactor, fixedFeeAmount)
}

// SetCustomTransactorFees is a paid mutator transaction binding the contract method 0x95fcd547.
//
// Solidity: function setCustomTransactorFees(uint128 synthMarketId, address transactor, uint256 fixedFeeAmount) returns()
func (_Contracts *ContractsSession) SetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address, fixedFeeAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetCustomTransactorFees(&_Contracts.TransactOpts, synthMarketId, transactor, fixedFeeAmount)
}

// SetCustomTransactorFees is a paid mutator transaction binding the contract method 0x95fcd547.
//
// Solidity: function setCustomTransactorFees(uint128 synthMarketId, address transactor, uint256 fixedFeeAmount) returns()
func (_Contracts *ContractsTransactorSession) SetCustomTransactorFees(synthMarketId *big.Int, transactor common.Address, fixedFeeAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetCustomTransactorFees(&_Contracts.TransactOpts, synthMarketId, transactor, fixedFeeAmount)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xec846bac.
//
// Solidity: function setDecayRate(uint128 marketId, uint256 rate) returns()
func (_Contracts *ContractsTransactor) SetDecayRate(opts *bind.TransactOpts, marketId *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setDecayRate", marketId, rate)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xec846bac.
//
// Solidity: function setDecayRate(uint128 marketId, uint256 rate) returns()
func (_Contracts *ContractsSession) SetDecayRate(marketId *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetDecayRate(&_Contracts.TransactOpts, marketId, rate)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xec846bac.
//
// Solidity: function setDecayRate(uint128 marketId, uint256 rate) returns()
func (_Contracts *ContractsTransactorSession) SetDecayRate(marketId *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetDecayRate(&_Contracts.TransactOpts, marketId, rate)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_Contracts *ContractsTransactor) SetDeniers(opts *bind.TransactOpts, feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setDeniers", feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_Contracts *ContractsSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetDeniers(&_Contracts.TransactOpts, feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_Contracts *ContractsTransactorSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetDeniers(&_Contracts.TransactOpts, feature, deniers)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_Contracts *ContractsTransactor) SetFeatureFlagAllowAll(opts *bind.TransactOpts, feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setFeatureFlagAllowAll", feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_Contracts *ContractsSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetFeatureFlagAllowAll(&_Contracts.TransactOpts, feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_Contracts *ContractsTransactorSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetFeatureFlagAllowAll(&_Contracts.TransactOpts, feature, allowAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_Contracts *ContractsTransactor) SetFeatureFlagDenyAll(opts *bind.TransactOpts, feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setFeatureFlagDenyAll", feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_Contracts *ContractsSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetFeatureFlagDenyAll(&_Contracts.TransactOpts, feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_Contracts *ContractsTransactorSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetFeatureFlagDenyAll(&_Contracts.TransactOpts, feature, denyAll)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x025f6120.
//
// Solidity: function setFeeCollector(uint128 synthMarketId, address feeCollector) returns()
func (_Contracts *ContractsTransactor) SetFeeCollector(opts *bind.TransactOpts, synthMarketId *big.Int, feeCollector common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setFeeCollector", synthMarketId, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x025f6120.
//
// Solidity: function setFeeCollector(uint128 synthMarketId, address feeCollector) returns()
func (_Contracts *ContractsSession) SetFeeCollector(synthMarketId *big.Int, feeCollector common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetFeeCollector(&_Contracts.TransactOpts, synthMarketId, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x025f6120.
//
// Solidity: function setFeeCollector(uint128 synthMarketId, address feeCollector) returns()
func (_Contracts *ContractsTransactorSession) SetFeeCollector(synthMarketId *big.Int, feeCollector common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetFeeCollector(&_Contracts.TransactOpts, synthMarketId, feeCollector)
}

// SetMarketSkewScale is a paid mutator transaction binding the contract method 0x9a40f8cb.
//
// Solidity: function setMarketSkewScale(uint128 synthMarketId, uint256 skewScale) returns()
func (_Contracts *ContractsTransactor) SetMarketSkewScale(opts *bind.TransactOpts, synthMarketId *big.Int, skewScale *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMarketSkewScale", synthMarketId, skewScale)
}

// SetMarketSkewScale is a paid mutator transaction binding the contract method 0x9a40f8cb.
//
// Solidity: function setMarketSkewScale(uint128 synthMarketId, uint256 skewScale) returns()
func (_Contracts *ContractsSession) SetMarketSkewScale(synthMarketId *big.Int, skewScale *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMarketSkewScale(&_Contracts.TransactOpts, synthMarketId, skewScale)
}

// SetMarketSkewScale is a paid mutator transaction binding the contract method 0x9a40f8cb.
//
// Solidity: function setMarketSkewScale(uint128 synthMarketId, uint256 skewScale) returns()
func (_Contracts *ContractsTransactorSession) SetMarketSkewScale(synthMarketId *big.Int, skewScale *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMarketSkewScale(&_Contracts.TransactOpts, synthMarketId, skewScale)
}

// SetMarketUtilizationFees is a paid mutator transaction binding the contract method 0x45f2601c.
//
// Solidity: function setMarketUtilizationFees(uint128 synthMarketId, uint256 utilizationFeeRate) returns()
func (_Contracts *ContractsTransactor) SetMarketUtilizationFees(opts *bind.TransactOpts, synthMarketId *big.Int, utilizationFeeRate *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMarketUtilizationFees", synthMarketId, utilizationFeeRate)
}

// SetMarketUtilizationFees is a paid mutator transaction binding the contract method 0x45f2601c.
//
// Solidity: function setMarketUtilizationFees(uint128 synthMarketId, uint256 utilizationFeeRate) returns()
func (_Contracts *ContractsSession) SetMarketUtilizationFees(synthMarketId *big.Int, utilizationFeeRate *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMarketUtilizationFees(&_Contracts.TransactOpts, synthMarketId, utilizationFeeRate)
}

// SetMarketUtilizationFees is a paid mutator transaction binding the contract method 0x45f2601c.
//
// Solidity: function setMarketUtilizationFees(uint128 synthMarketId, uint256 utilizationFeeRate) returns()
func (_Contracts *ContractsTransactorSession) SetMarketUtilizationFees(synthMarketId *big.Int, utilizationFeeRate *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMarketUtilizationFees(&_Contracts.TransactOpts, synthMarketId, utilizationFeeRate)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_Contracts *ContractsTransactor) SetSettlementStrategyEnabled(opts *bind.TransactOpts, marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setSettlementStrategyEnabled", marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_Contracts *ContractsSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetSettlementStrategyEnabled(&_Contracts.TransactOpts, marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_Contracts *ContractsTransactorSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetSettlementStrategyEnabled(&_Contracts.TransactOpts, marketId, strategyId, enabled)
}

// SetSynthImplementation is a paid mutator transaction binding the contract method 0xec04ceb1.
//
// Solidity: function setSynthImplementation(address synthImplementation) returns()
func (_Contracts *ContractsTransactor) SetSynthImplementation(opts *bind.TransactOpts, synthImplementation common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setSynthImplementation", synthImplementation)
}

// SetSynthImplementation is a paid mutator transaction binding the contract method 0xec04ceb1.
//
// Solidity: function setSynthImplementation(address synthImplementation) returns()
func (_Contracts *ContractsSession) SetSynthImplementation(synthImplementation common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetSynthImplementation(&_Contracts.TransactOpts, synthImplementation)
}

// SetSynthImplementation is a paid mutator transaction binding the contract method 0xec04ceb1.
//
// Solidity: function setSynthImplementation(address synthImplementation) returns()
func (_Contracts *ContractsTransactorSession) SetSynthImplementation(synthImplementation common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetSynthImplementation(&_Contracts.TransactOpts, synthImplementation)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_Contracts *ContractsTransactor) SetSynthetix(opts *bind.TransactOpts, synthetix common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setSynthetix", synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_Contracts *ContractsSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetSynthetix(&_Contracts.TransactOpts, synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_Contracts *ContractsTransactorSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetSynthetix(&_Contracts.TransactOpts, synthetix)
}

// SetWrapper is a paid mutator transaction binding the contract method 0x673a21e5.
//
// Solidity: function setWrapper(uint128 marketId, address wrapCollateralType, uint256 maxWrappableAmount) returns()
func (_Contracts *ContractsTransactor) SetWrapper(opts *bind.TransactOpts, marketId *big.Int, wrapCollateralType common.Address, maxWrappableAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setWrapper", marketId, wrapCollateralType, maxWrappableAmount)
}

// SetWrapper is a paid mutator transaction binding the contract method 0x673a21e5.
//
// Solidity: function setWrapper(uint128 marketId, address wrapCollateralType, uint256 maxWrappableAmount) returns()
func (_Contracts *ContractsSession) SetWrapper(marketId *big.Int, wrapCollateralType common.Address, maxWrappableAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetWrapper(&_Contracts.TransactOpts, marketId, wrapCollateralType, maxWrappableAmount)
}

// SetWrapper is a paid mutator transaction binding the contract method 0x673a21e5.
//
// Solidity: function setWrapper(uint128 marketId, address wrapCollateralType, uint256 maxWrappableAmount) returns()
func (_Contracts *ContractsTransactorSession) SetWrapper(marketId *big.Int, wrapCollateralType common.Address, maxWrappableAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetWrapper(&_Contracts.TransactOpts, marketId, wrapCollateralType, maxWrappableAmount)
}

// SetWrapperFees is a paid mutator transaction binding the contract method 0x6539b1c3.
//
// Solidity: function setWrapperFees(uint128 synthMarketId, int256 wrapFee, int256 unwrapFee) returns()
func (_Contracts *ContractsTransactor) SetWrapperFees(opts *bind.TransactOpts, synthMarketId *big.Int, wrapFee *big.Int, unwrapFee *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setWrapperFees", synthMarketId, wrapFee, unwrapFee)
}

// SetWrapperFees is a paid mutator transaction binding the contract method 0x6539b1c3.
//
// Solidity: function setWrapperFees(uint128 synthMarketId, int256 wrapFee, int256 unwrapFee) returns()
func (_Contracts *ContractsSession) SetWrapperFees(synthMarketId *big.Int, wrapFee *big.Int, unwrapFee *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetWrapperFees(&_Contracts.TransactOpts, synthMarketId, wrapFee, unwrapFee)
}

// SetWrapperFees is a paid mutator transaction binding the contract method 0x6539b1c3.
//
// Solidity: function setWrapperFees(uint128 synthMarketId, int256 wrapFee, int256 unwrapFee) returns()
func (_Contracts *ContractsTransactorSession) SetWrapperFees(synthMarketId *big.Int, wrapFee *big.Int, unwrapFee *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetWrapperFees(&_Contracts.TransactOpts, synthMarketId, wrapFee, unwrapFee)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x9444ac48.
//
// Solidity: function settleOrder(uint128 marketId, uint128 asyncOrderId) returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactor) SettleOrder(opts *bind.TransactOpts, marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "settleOrder", marketId, asyncOrderId)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x9444ac48.
//
// Solidity: function settleOrder(uint128 marketId, uint128 asyncOrderId) returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsSession) SettleOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SettleOrder(&_Contracts.TransactOpts, marketId, asyncOrderId)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x9444ac48.
//
// Solidity: function settleOrder(uint128 marketId, uint128 asyncOrderId) returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactorSession) SettleOrder(marketId *big.Int, asyncOrderId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SettleOrder(&_Contracts.TransactOpts, marketId, asyncOrderId)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactor) SettlePythOrder(opts *bind.TransactOpts, result []byte, extraData []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "settlePythOrder", result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SettlePythOrder(&_Contracts.TransactOpts, result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns(uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactorSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SettlePythOrder(&_Contracts.TransactOpts, result, extraData)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_Contracts *ContractsTransactor) SimulateUpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "simulateUpgradeTo", newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_Contracts *ContractsSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SimulateUpgradeTo(&_Contracts.TransactOpts, newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_Contracts *ContractsTransactorSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SimulateUpgradeTo(&_Contracts.TransactOpts, newImplementation)
}

// Unwrap is a paid mutator transaction binding the contract method 0x784dad9e.
//
// Solidity: function unwrap(uint128 marketId, uint256 unwrapAmount, uint256 minAmountReceived) returns(uint256 returnCollateralAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactor) Unwrap(opts *bind.TransactOpts, marketId *big.Int, unwrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unwrap", marketId, unwrapAmount, minAmountReceived)
}

// Unwrap is a paid mutator transaction binding the contract method 0x784dad9e.
//
// Solidity: function unwrap(uint128 marketId, uint256 unwrapAmount, uint256 minAmountReceived) returns(uint256 returnCollateralAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsSession) Unwrap(marketId *big.Int, unwrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Unwrap(&_Contracts.TransactOpts, marketId, unwrapAmount, minAmountReceived)
}

// Unwrap is a paid mutator transaction binding the contract method 0x784dad9e.
//
// Solidity: function unwrap(uint128 marketId, uint256 unwrapAmount, uint256 minAmountReceived) returns(uint256 returnCollateralAmount, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactorSession) Unwrap(marketId *big.Int, unwrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Unwrap(&_Contracts.TransactOpts, marketId, unwrapAmount, minAmountReceived)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x6a0e2085.
//
// Solidity: function updatePriceData(uint128 synthMarketId, bytes32 buyFeedId, bytes32 sellFeedId) returns()
func (_Contracts *ContractsTransactor) UpdatePriceData(opts *bind.TransactOpts, synthMarketId *big.Int, buyFeedId [32]byte, sellFeedId [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updatePriceData", synthMarketId, buyFeedId, sellFeedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x6a0e2085.
//
// Solidity: function updatePriceData(uint128 synthMarketId, bytes32 buyFeedId, bytes32 sellFeedId) returns()
func (_Contracts *ContractsSession) UpdatePriceData(synthMarketId *big.Int, buyFeedId [32]byte, sellFeedId [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpdatePriceData(&_Contracts.TransactOpts, synthMarketId, buyFeedId, sellFeedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x6a0e2085.
//
// Solidity: function updatePriceData(uint128 synthMarketId, bytes32 buyFeedId, bytes32 sellFeedId) returns()
func (_Contracts *ContractsTransactorSession) UpdatePriceData(synthMarketId *big.Int, buyFeedId [32]byte, sellFeedId [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpdatePriceData(&_Contracts.TransactOpts, synthMarketId, buyFeedId, sellFeedId)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6ad77077.
//
// Solidity: function updateReferrerShare(uint128 synthMarketId, address referrer, uint256 sharePercentage) returns()
func (_Contracts *ContractsTransactor) UpdateReferrerShare(opts *bind.TransactOpts, synthMarketId *big.Int, referrer common.Address, sharePercentage *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateReferrerShare", synthMarketId, referrer, sharePercentage)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6ad77077.
//
// Solidity: function updateReferrerShare(uint128 synthMarketId, address referrer, uint256 sharePercentage) returns()
func (_Contracts *ContractsSession) UpdateReferrerShare(synthMarketId *big.Int, referrer common.Address, sharePercentage *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateReferrerShare(&_Contracts.TransactOpts, synthMarketId, referrer, sharePercentage)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6ad77077.
//
// Solidity: function updateReferrerShare(uint128 synthMarketId, address referrer, uint256 sharePercentage) returns()
func (_Contracts *ContractsTransactorSession) UpdateReferrerShare(synthMarketId *big.Int, referrer common.Address, sharePercentage *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateReferrerShare(&_Contracts.TransactOpts, synthMarketId, referrer, sharePercentage)
}

// UpgradeSynthImpl is a paid mutator transaction binding the contract method 0xc99d0cdd.
//
// Solidity: function upgradeSynthImpl(uint128 marketId) returns()
func (_Contracts *ContractsTransactor) UpgradeSynthImpl(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "upgradeSynthImpl", marketId)
}

// UpgradeSynthImpl is a paid mutator transaction binding the contract method 0xc99d0cdd.
//
// Solidity: function upgradeSynthImpl(uint128 marketId) returns()
func (_Contracts *ContractsSession) UpgradeSynthImpl(marketId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpgradeSynthImpl(&_Contracts.TransactOpts, marketId)
}

// UpgradeSynthImpl is a paid mutator transaction binding the contract method 0xc99d0cdd.
//
// Solidity: function upgradeSynthImpl(uint128 marketId) returns()
func (_Contracts *ContractsTransactorSession) UpgradeSynthImpl(marketId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpgradeSynthImpl(&_Contracts.TransactOpts, marketId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Contracts *ContractsTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Contracts *ContractsSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UpgradeTo(&_Contracts.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Contracts *ContractsTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UpgradeTo(&_Contracts.TransactOpts, newImplementation)
}

// Wrap is a paid mutator transaction binding the contract method 0xd7ce770c.
//
// Solidity: function wrap(uint128 marketId, uint256 wrapAmount, uint256 minAmountReceived) returns(uint256 amountToMint, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactor) Wrap(opts *bind.TransactOpts, marketId *big.Int, wrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "wrap", marketId, wrapAmount, minAmountReceived)
}

// Wrap is a paid mutator transaction binding the contract method 0xd7ce770c.
//
// Solidity: function wrap(uint128 marketId, uint256 wrapAmount, uint256 minAmountReceived) returns(uint256 amountToMint, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsSession) Wrap(marketId *big.Int, wrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Wrap(&_Contracts.TransactOpts, marketId, wrapAmount, minAmountReceived)
}

// Wrap is a paid mutator transaction binding the contract method 0xd7ce770c.
//
// Solidity: function wrap(uint128 marketId, uint256 wrapAmount, uint256 minAmountReceived) returns(uint256 amountToMint, (uint256,uint256,int256,int256) fees)
func (_Contracts *ContractsTransactorSession) Wrap(marketId *big.Int, wrapAmount *big.Int, minAmountReceived *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Wrap(&_Contracts.TransactOpts, marketId, wrapAmount, minAmountReceived)
}

// ContractsAssociatedSystemSetIterator is returned from FilterAssociatedSystemSet and is used to iterate over the raw logs and unpacked data for AssociatedSystemSet events raised by the Contracts contract.
type ContractsAssociatedSystemSetIterator struct {
	Event *ContractsAssociatedSystemSet // Event containing the contract specifics and raw log

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
func (it *ContractsAssociatedSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAssociatedSystemSet)
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
		it.Event = new(ContractsAssociatedSystemSet)
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
func (it *ContractsAssociatedSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAssociatedSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAssociatedSystemSet represents a AssociatedSystemSet event raised by the Contracts contract.
type ContractsAssociatedSystemSet struct {
	Kind  [32]byte
	Id    [32]byte
	Proxy common.Address
	Impl  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssociatedSystemSet is a free log retrieval operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_Contracts *ContractsFilterer) FilterAssociatedSystemSet(opts *bind.FilterOpts, kind [][32]byte, id [][32]byte) (*ContractsAssociatedSystemSetIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAssociatedSystemSetIterator{contract: _Contracts.contract, event: "AssociatedSystemSet", logs: logs, sub: sub}, nil
}

// WatchAssociatedSystemSet is a free log subscription operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_Contracts *ContractsFilterer) WatchAssociatedSystemSet(opts *bind.WatchOpts, sink chan<- *ContractsAssociatedSystemSet, kind [][32]byte, id [][32]byte) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAssociatedSystemSet)
				if err := _Contracts.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseAssociatedSystemSet(log types.Log) (*ContractsAssociatedSystemSet, error) {
	event := new(ContractsAssociatedSystemSet)
	if err := _Contracts.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAsyncFixedFeeSetIterator is returned from FilterAsyncFixedFeeSet and is used to iterate over the raw logs and unpacked data for AsyncFixedFeeSet events raised by the Contracts contract.
type ContractsAsyncFixedFeeSetIterator struct {
	Event *ContractsAsyncFixedFeeSet // Event containing the contract specifics and raw log

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
func (it *ContractsAsyncFixedFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAsyncFixedFeeSet)
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
		it.Event = new(ContractsAsyncFixedFeeSet)
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
func (it *ContractsAsyncFixedFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAsyncFixedFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAsyncFixedFeeSet represents a AsyncFixedFeeSet event raised by the Contracts contract.
type ContractsAsyncFixedFeeSet struct {
	SynthMarketId *big.Int
	AsyncFixedFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAsyncFixedFeeSet is a free log retrieval operation binding the contract event 0xaef59d229d195c5b8372221acbf4041b926fb1616a95f93e44379e4f30d57bfe.
//
// Solidity: event AsyncFixedFeeSet(uint256 indexed synthMarketId, uint256 asyncFixedFee)
func (_Contracts *ContractsFilterer) FilterAsyncFixedFeeSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsAsyncFixedFeeSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AsyncFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAsyncFixedFeeSetIterator{contract: _Contracts.contract, event: "AsyncFixedFeeSet", logs: logs, sub: sub}, nil
}

// WatchAsyncFixedFeeSet is a free log subscription operation binding the contract event 0xaef59d229d195c5b8372221acbf4041b926fb1616a95f93e44379e4f30d57bfe.
//
// Solidity: event AsyncFixedFeeSet(uint256 indexed synthMarketId, uint256 asyncFixedFee)
func (_Contracts *ContractsFilterer) WatchAsyncFixedFeeSet(opts *bind.WatchOpts, sink chan<- *ContractsAsyncFixedFeeSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AsyncFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAsyncFixedFeeSet)
				if err := _Contracts.contract.UnpackLog(event, "AsyncFixedFeeSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseAsyncFixedFeeSet(log types.Log) (*ContractsAsyncFixedFeeSet, error) {
	event := new(ContractsAsyncFixedFeeSet)
	if err := _Contracts.contract.UnpackLog(event, "AsyncFixedFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAtomicFixedFeeSetIterator is returned from FilterAtomicFixedFeeSet and is used to iterate over the raw logs and unpacked data for AtomicFixedFeeSet events raised by the Contracts contract.
type ContractsAtomicFixedFeeSetIterator struct {
	Event *ContractsAtomicFixedFeeSet // Event containing the contract specifics and raw log

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
func (it *ContractsAtomicFixedFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAtomicFixedFeeSet)
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
		it.Event = new(ContractsAtomicFixedFeeSet)
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
func (it *ContractsAtomicFixedFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAtomicFixedFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAtomicFixedFeeSet represents a AtomicFixedFeeSet event raised by the Contracts contract.
type ContractsAtomicFixedFeeSet struct {
	SynthMarketId  *big.Int
	AtomicFixedFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAtomicFixedFeeSet is a free log retrieval operation binding the contract event 0x6b0a526b06b2f30ba2d5b063c2ef81547512216d37c540b86039e3a19f1d2b3f.
//
// Solidity: event AtomicFixedFeeSet(uint256 indexed synthMarketId, uint256 atomicFixedFee)
func (_Contracts *ContractsFilterer) FilterAtomicFixedFeeSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsAtomicFixedFeeSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AtomicFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAtomicFixedFeeSetIterator{contract: _Contracts.contract, event: "AtomicFixedFeeSet", logs: logs, sub: sub}, nil
}

// WatchAtomicFixedFeeSet is a free log subscription operation binding the contract event 0x6b0a526b06b2f30ba2d5b063c2ef81547512216d37c540b86039e3a19f1d2b3f.
//
// Solidity: event AtomicFixedFeeSet(uint256 indexed synthMarketId, uint256 atomicFixedFee)
func (_Contracts *ContractsFilterer) WatchAtomicFixedFeeSet(opts *bind.WatchOpts, sink chan<- *ContractsAtomicFixedFeeSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AtomicFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAtomicFixedFeeSet)
				if err := _Contracts.contract.UnpackLog(event, "AtomicFixedFeeSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseAtomicFixedFeeSet(log types.Log) (*ContractsAtomicFixedFeeSet, error) {
	event := new(ContractsAtomicFixedFeeSet)
	if err := _Contracts.contract.UnpackLog(event, "AtomicFixedFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCollateralLeverageSetIterator is returned from FilterCollateralLeverageSet and is used to iterate over the raw logs and unpacked data for CollateralLeverageSet events raised by the Contracts contract.
type ContractsCollateralLeverageSetIterator struct {
	Event *ContractsCollateralLeverageSet // Event containing the contract specifics and raw log

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
func (it *ContractsCollateralLeverageSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCollateralLeverageSet)
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
		it.Event = new(ContractsCollateralLeverageSet)
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
func (it *ContractsCollateralLeverageSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCollateralLeverageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCollateralLeverageSet represents a CollateralLeverageSet event raised by the Contracts contract.
type ContractsCollateralLeverageSet struct {
	SynthMarketId      *big.Int
	CollateralLeverage *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCollateralLeverageSet is a free log retrieval operation binding the contract event 0x83113c0e8d811f7e7017948357e945a1a8a6b6fc0d76c8512ffdd6f6766e8a13.
//
// Solidity: event CollateralLeverageSet(uint256 indexed synthMarketId, uint256 collateralLeverage)
func (_Contracts *ContractsFilterer) FilterCollateralLeverageSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsCollateralLeverageSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CollateralLeverageSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCollateralLeverageSetIterator{contract: _Contracts.contract, event: "CollateralLeverageSet", logs: logs, sub: sub}, nil
}

// WatchCollateralLeverageSet is a free log subscription operation binding the contract event 0x83113c0e8d811f7e7017948357e945a1a8a6b6fc0d76c8512ffdd6f6766e8a13.
//
// Solidity: event CollateralLeverageSet(uint256 indexed synthMarketId, uint256 collateralLeverage)
func (_Contracts *ContractsFilterer) WatchCollateralLeverageSet(opts *bind.WatchOpts, sink chan<- *ContractsCollateralLeverageSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CollateralLeverageSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCollateralLeverageSet)
				if err := _Contracts.contract.UnpackLog(event, "CollateralLeverageSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseCollateralLeverageSet(log types.Log) (*ContractsCollateralLeverageSet, error) {
	event := new(ContractsCollateralLeverageSet)
	if err := _Contracts.contract.UnpackLog(event, "CollateralLeverageSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDecayRateUpdatedIterator is returned from FilterDecayRateUpdated and is used to iterate over the raw logs and unpacked data for DecayRateUpdated events raised by the Contracts contract.
type ContractsDecayRateUpdatedIterator struct {
	Event *ContractsDecayRateUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsDecayRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDecayRateUpdated)
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
		it.Event = new(ContractsDecayRateUpdated)
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
func (it *ContractsDecayRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDecayRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDecayRateUpdated represents a DecayRateUpdated event raised by the Contracts contract.
type ContractsDecayRateUpdated struct {
	MarketId *big.Int
	Rate     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDecayRateUpdated is a free log retrieval operation binding the contract event 0x8a1ed33cd71d1533eadd1d4eb0ea2ae64da7d343cb2a932fdf135f345264e2b5.
//
// Solidity: event DecayRateUpdated(uint128 indexed marketId, uint256 rate)
func (_Contracts *ContractsFilterer) FilterDecayRateUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsDecayRateUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DecayRateUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsDecayRateUpdatedIterator{contract: _Contracts.contract, event: "DecayRateUpdated", logs: logs, sub: sub}, nil
}

// WatchDecayRateUpdated is a free log subscription operation binding the contract event 0x8a1ed33cd71d1533eadd1d4eb0ea2ae64da7d343cb2a932fdf135f345264e2b5.
//
// Solidity: event DecayRateUpdated(uint128 indexed marketId, uint256 rate)
func (_Contracts *ContractsFilterer) WatchDecayRateUpdated(opts *bind.WatchOpts, sink chan<- *ContractsDecayRateUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DecayRateUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDecayRateUpdated)
				if err := _Contracts.contract.UnpackLog(event, "DecayRateUpdated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseDecayRateUpdated(log types.Log) (*ContractsDecayRateUpdated, error) {
	event := new(ContractsDecayRateUpdated)
	if err := _Contracts.contract.UnpackLog(event, "DecayRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFeatureFlagAllowAllSetIterator is returned from FilterFeatureFlagAllowAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowAllSet events raised by the Contracts contract.
type ContractsFeatureFlagAllowAllSetIterator struct {
	Event *ContractsFeatureFlagAllowAllSet // Event containing the contract specifics and raw log

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
func (it *ContractsFeatureFlagAllowAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFeatureFlagAllowAllSet)
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
		it.Event = new(ContractsFeatureFlagAllowAllSet)
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
func (it *ContractsFeatureFlagAllowAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFeatureFlagAllowAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFeatureFlagAllowAllSet represents a FeatureFlagAllowAllSet event raised by the Contracts contract.
type ContractsFeatureFlagAllowAllSet struct {
	Feature  [32]byte
	AllowAll bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowAllSet is a free log retrieval operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_Contracts *ContractsFilterer) FilterFeatureFlagAllowAllSet(opts *bind.FilterOpts, feature [][32]byte) (*ContractsFeatureFlagAllowAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &ContractsFeatureFlagAllowAllSetIterator{contract: _Contracts.contract, event: "FeatureFlagAllowAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowAllSet is a free log subscription operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_Contracts *ContractsFilterer) WatchFeatureFlagAllowAllSet(opts *bind.WatchOpts, sink chan<- *ContractsFeatureFlagAllowAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFeatureFlagAllowAllSet)
				if err := _Contracts.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseFeatureFlagAllowAllSet(log types.Log) (*ContractsFeatureFlagAllowAllSet, error) {
	event := new(ContractsFeatureFlagAllowAllSet)
	if err := _Contracts.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFeatureFlagAllowlistAddedIterator is returned from FilterFeatureFlagAllowlistAdded and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistAdded events raised by the Contracts contract.
type ContractsFeatureFlagAllowlistAddedIterator struct {
	Event *ContractsFeatureFlagAllowlistAdded // Event containing the contract specifics and raw log

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
func (it *ContractsFeatureFlagAllowlistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFeatureFlagAllowlistAdded)
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
		it.Event = new(ContractsFeatureFlagAllowlistAdded)
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
func (it *ContractsFeatureFlagAllowlistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFeatureFlagAllowlistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFeatureFlagAllowlistAdded represents a FeatureFlagAllowlistAdded event raised by the Contracts contract.
type ContractsFeatureFlagAllowlistAdded struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistAdded is a free log retrieval operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_Contracts *ContractsFilterer) FilterFeatureFlagAllowlistAdded(opts *bind.FilterOpts, feature [][32]byte) (*ContractsFeatureFlagAllowlistAddedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return &ContractsFeatureFlagAllowlistAddedIterator{contract: _Contracts.contract, event: "FeatureFlagAllowlistAdded", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistAdded is a free log subscription operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_Contracts *ContractsFilterer) WatchFeatureFlagAllowlistAdded(opts *bind.WatchOpts, sink chan<- *ContractsFeatureFlagAllowlistAdded, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFeatureFlagAllowlistAdded)
				if err := _Contracts.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseFeatureFlagAllowlistAdded(log types.Log) (*ContractsFeatureFlagAllowlistAdded, error) {
	event := new(ContractsFeatureFlagAllowlistAdded)
	if err := _Contracts.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFeatureFlagAllowlistRemovedIterator is returned from FilterFeatureFlagAllowlistRemoved and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistRemoved events raised by the Contracts contract.
type ContractsFeatureFlagAllowlistRemovedIterator struct {
	Event *ContractsFeatureFlagAllowlistRemoved // Event containing the contract specifics and raw log

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
func (it *ContractsFeatureFlagAllowlistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFeatureFlagAllowlistRemoved)
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
		it.Event = new(ContractsFeatureFlagAllowlistRemoved)
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
func (it *ContractsFeatureFlagAllowlistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFeatureFlagAllowlistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFeatureFlagAllowlistRemoved represents a FeatureFlagAllowlistRemoved event raised by the Contracts contract.
type ContractsFeatureFlagAllowlistRemoved struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistRemoved is a free log retrieval operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_Contracts *ContractsFilterer) FilterFeatureFlagAllowlistRemoved(opts *bind.FilterOpts, feature [][32]byte) (*ContractsFeatureFlagAllowlistRemovedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return &ContractsFeatureFlagAllowlistRemovedIterator{contract: _Contracts.contract, event: "FeatureFlagAllowlistRemoved", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistRemoved is a free log subscription operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_Contracts *ContractsFilterer) WatchFeatureFlagAllowlistRemoved(opts *bind.WatchOpts, sink chan<- *ContractsFeatureFlagAllowlistRemoved, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFeatureFlagAllowlistRemoved)
				if err := _Contracts.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseFeatureFlagAllowlistRemoved(log types.Log) (*ContractsFeatureFlagAllowlistRemoved, error) {
	event := new(ContractsFeatureFlagAllowlistRemoved)
	if err := _Contracts.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFeatureFlagDeniersResetIterator is returned from FilterFeatureFlagDeniersReset and is used to iterate over the raw logs and unpacked data for FeatureFlagDeniersReset events raised by the Contracts contract.
type ContractsFeatureFlagDeniersResetIterator struct {
	Event *ContractsFeatureFlagDeniersReset // Event containing the contract specifics and raw log

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
func (it *ContractsFeatureFlagDeniersResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFeatureFlagDeniersReset)
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
		it.Event = new(ContractsFeatureFlagDeniersReset)
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
func (it *ContractsFeatureFlagDeniersResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFeatureFlagDeniersResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFeatureFlagDeniersReset represents a FeatureFlagDeniersReset event raised by the Contracts contract.
type ContractsFeatureFlagDeniersReset struct {
	Feature [32]byte
	Deniers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDeniersReset is a free log retrieval operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_Contracts *ContractsFilterer) FilterFeatureFlagDeniersReset(opts *bind.FilterOpts, feature [][32]byte) (*ContractsFeatureFlagDeniersResetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return &ContractsFeatureFlagDeniersResetIterator{contract: _Contracts.contract, event: "FeatureFlagDeniersReset", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDeniersReset is a free log subscription operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_Contracts *ContractsFilterer) WatchFeatureFlagDeniersReset(opts *bind.WatchOpts, sink chan<- *ContractsFeatureFlagDeniersReset, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFeatureFlagDeniersReset)
				if err := _Contracts.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseFeatureFlagDeniersReset(log types.Log) (*ContractsFeatureFlagDeniersReset, error) {
	event := new(ContractsFeatureFlagDeniersReset)
	if err := _Contracts.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFeatureFlagDenyAllSetIterator is returned from FilterFeatureFlagDenyAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagDenyAllSet events raised by the Contracts contract.
type ContractsFeatureFlagDenyAllSetIterator struct {
	Event *ContractsFeatureFlagDenyAllSet // Event containing the contract specifics and raw log

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
func (it *ContractsFeatureFlagDenyAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFeatureFlagDenyAllSet)
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
		it.Event = new(ContractsFeatureFlagDenyAllSet)
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
func (it *ContractsFeatureFlagDenyAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFeatureFlagDenyAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFeatureFlagDenyAllSet represents a FeatureFlagDenyAllSet event raised by the Contracts contract.
type ContractsFeatureFlagDenyAllSet struct {
	Feature [32]byte
	DenyAll bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDenyAllSet is a free log retrieval operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_Contracts *ContractsFilterer) FilterFeatureFlagDenyAllSet(opts *bind.FilterOpts, feature [][32]byte) (*ContractsFeatureFlagDenyAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &ContractsFeatureFlagDenyAllSetIterator{contract: _Contracts.contract, event: "FeatureFlagDenyAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDenyAllSet is a free log subscription operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_Contracts *ContractsFilterer) WatchFeatureFlagDenyAllSet(opts *bind.WatchOpts, sink chan<- *ContractsFeatureFlagDenyAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFeatureFlagDenyAllSet)
				if err := _Contracts.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseFeatureFlagDenyAllSet(log types.Log) (*ContractsFeatureFlagDenyAllSet, error) {
	event := new(ContractsFeatureFlagDenyAllSet)
	if err := _Contracts.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFeeCollectorSetIterator is returned from FilterFeeCollectorSet and is used to iterate over the raw logs and unpacked data for FeeCollectorSet events raised by the Contracts contract.
type ContractsFeeCollectorSetIterator struct {
	Event *ContractsFeeCollectorSet // Event containing the contract specifics and raw log

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
func (it *ContractsFeeCollectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFeeCollectorSet)
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
		it.Event = new(ContractsFeeCollectorSet)
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
func (it *ContractsFeeCollectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFeeCollectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFeeCollectorSet represents a FeeCollectorSet event raised by the Contracts contract.
type ContractsFeeCollectorSet struct {
	SynthMarketId *big.Int
	FeeCollector  common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorSet is a free log retrieval operation binding the contract event 0x9559a9b7e14bf53553c17859be245a108350185ec859ec690012b13b820b7ef4.
//
// Solidity: event FeeCollectorSet(uint256 indexed synthMarketId, address feeCollector)
func (_Contracts *ContractsFilterer) FilterFeeCollectorSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsFeeCollectorSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FeeCollectorSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsFeeCollectorSetIterator{contract: _Contracts.contract, event: "FeeCollectorSet", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorSet is a free log subscription operation binding the contract event 0x9559a9b7e14bf53553c17859be245a108350185ec859ec690012b13b820b7ef4.
//
// Solidity: event FeeCollectorSet(uint256 indexed synthMarketId, address feeCollector)
func (_Contracts *ContractsFilterer) WatchFeeCollectorSet(opts *bind.WatchOpts, sink chan<- *ContractsFeeCollectorSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FeeCollectorSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFeeCollectorSet)
				if err := _Contracts.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseFeeCollectorSet(log types.Log) (*ContractsFeeCollectorSet, error) {
	event := new(ContractsFeeCollectorSet)
	if err := _Contracts.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketNominationRenouncedIterator is returned from FilterMarketNominationRenounced and is used to iterate over the raw logs and unpacked data for MarketNominationRenounced events raised by the Contracts contract.
type ContractsMarketNominationRenouncedIterator struct {
	Event *ContractsMarketNominationRenounced // Event containing the contract specifics and raw log

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
func (it *ContractsMarketNominationRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketNominationRenounced)
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
		it.Event = new(ContractsMarketNominationRenounced)
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
func (it *ContractsMarketNominationRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketNominationRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketNominationRenounced represents a MarketNominationRenounced event raised by the Contracts contract.
type ContractsMarketNominationRenounced struct {
	MarketId *big.Int
	Nominee  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketNominationRenounced is a free log retrieval operation binding the contract event 0xf5b87e3c7e0caa8e0d233591fff85e764ebc73d5c7027cce729fd4beab04c2b6.
//
// Solidity: event MarketNominationRenounced(uint128 indexed marketId, address nominee)
func (_Contracts *ContractsFilterer) FilterMarketNominationRenounced(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsMarketNominationRenouncedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketNominationRenounced", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketNominationRenouncedIterator{contract: _Contracts.contract, event: "MarketNominationRenounced", logs: logs, sub: sub}, nil
}

// WatchMarketNominationRenounced is a free log subscription operation binding the contract event 0xf5b87e3c7e0caa8e0d233591fff85e764ebc73d5c7027cce729fd4beab04c2b6.
//
// Solidity: event MarketNominationRenounced(uint128 indexed marketId, address nominee)
func (_Contracts *ContractsFilterer) WatchMarketNominationRenounced(opts *bind.WatchOpts, sink chan<- *ContractsMarketNominationRenounced, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketNominationRenounced", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketNominationRenounced)
				if err := _Contracts.contract.UnpackLog(event, "MarketNominationRenounced", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseMarketNominationRenounced(log types.Log) (*ContractsMarketNominationRenounced, error) {
	event := new(ContractsMarketNominationRenounced)
	if err := _Contracts.contract.UnpackLog(event, "MarketNominationRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketOwnerChangedIterator is returned from FilterMarketOwnerChanged and is used to iterate over the raw logs and unpacked data for MarketOwnerChanged events raised by the Contracts contract.
type ContractsMarketOwnerChangedIterator struct {
	Event *ContractsMarketOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ContractsMarketOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketOwnerChanged)
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
		it.Event = new(ContractsMarketOwnerChanged)
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
func (it *ContractsMarketOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketOwnerChanged represents a MarketOwnerChanged event raised by the Contracts contract.
type ContractsMarketOwnerChanged struct {
	MarketId *big.Int
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketOwnerChanged is a free log retrieval operation binding the contract event 0xe73c996387b656d1e0ea2866c854ed68122ce4e4eea51d6af012938b3c7ff52f.
//
// Solidity: event MarketOwnerChanged(uint128 indexed marketId, address oldOwner, address newOwner)
func (_Contracts *ContractsFilterer) FilterMarketOwnerChanged(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsMarketOwnerChangedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketOwnerChanged", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketOwnerChangedIterator{contract: _Contracts.contract, event: "MarketOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchMarketOwnerChanged is a free log subscription operation binding the contract event 0xe73c996387b656d1e0ea2866c854ed68122ce4e4eea51d6af012938b3c7ff52f.
//
// Solidity: event MarketOwnerChanged(uint128 indexed marketId, address oldOwner, address newOwner)
func (_Contracts *ContractsFilterer) WatchMarketOwnerChanged(opts *bind.WatchOpts, sink chan<- *ContractsMarketOwnerChanged, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketOwnerChanged", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketOwnerChanged)
				if err := _Contracts.contract.UnpackLog(event, "MarketOwnerChanged", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseMarketOwnerChanged(log types.Log) (*ContractsMarketOwnerChanged, error) {
	event := new(ContractsMarketOwnerChanged)
	if err := _Contracts.contract.UnpackLog(event, "MarketOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketOwnerNominatedIterator is returned from FilterMarketOwnerNominated and is used to iterate over the raw logs and unpacked data for MarketOwnerNominated events raised by the Contracts contract.
type ContractsMarketOwnerNominatedIterator struct {
	Event *ContractsMarketOwnerNominated // Event containing the contract specifics and raw log

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
func (it *ContractsMarketOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketOwnerNominated)
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
		it.Event = new(ContractsMarketOwnerNominated)
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
func (it *ContractsMarketOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketOwnerNominated represents a MarketOwnerNominated event raised by the Contracts contract.
type ContractsMarketOwnerNominated struct {
	MarketId *big.Int
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketOwnerNominated is a free log retrieval operation binding the contract event 0x54605d90ee82d9b4318b0b4b479f3966976e44b94c6aff221c04f5294f85016b.
//
// Solidity: event MarketOwnerNominated(uint128 indexed marketId, address newOwner)
func (_Contracts *ContractsFilterer) FilterMarketOwnerNominated(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsMarketOwnerNominatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketOwnerNominated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketOwnerNominatedIterator{contract: _Contracts.contract, event: "MarketOwnerNominated", logs: logs, sub: sub}, nil
}

// WatchMarketOwnerNominated is a free log subscription operation binding the contract event 0x54605d90ee82d9b4318b0b4b479f3966976e44b94c6aff221c04f5294f85016b.
//
// Solidity: event MarketOwnerNominated(uint128 indexed marketId, address newOwner)
func (_Contracts *ContractsFilterer) WatchMarketOwnerNominated(opts *bind.WatchOpts, sink chan<- *ContractsMarketOwnerNominated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketOwnerNominated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketOwnerNominated)
				if err := _Contracts.contract.UnpackLog(event, "MarketOwnerNominated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseMarketOwnerNominated(log types.Log) (*ContractsMarketOwnerNominated, error) {
	event := new(ContractsMarketOwnerNominated)
	if err := _Contracts.contract.UnpackLog(event, "MarketOwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketSkewScaleSetIterator is returned from FilterMarketSkewScaleSet and is used to iterate over the raw logs and unpacked data for MarketSkewScaleSet events raised by the Contracts contract.
type ContractsMarketSkewScaleSetIterator struct {
	Event *ContractsMarketSkewScaleSet // Event containing the contract specifics and raw log

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
func (it *ContractsMarketSkewScaleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketSkewScaleSet)
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
		it.Event = new(ContractsMarketSkewScaleSet)
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
func (it *ContractsMarketSkewScaleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketSkewScaleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketSkewScaleSet represents a MarketSkewScaleSet event raised by the Contracts contract.
type ContractsMarketSkewScaleSet struct {
	SynthMarketId *big.Int
	SkewScale     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarketSkewScaleSet is a free log retrieval operation binding the contract event 0x786fdfd5a0e146d8f4878876a8439ff01436e5969b14682e12d07d7e926b157c.
//
// Solidity: event MarketSkewScaleSet(uint256 indexed synthMarketId, uint256 skewScale)
func (_Contracts *ContractsFilterer) FilterMarketSkewScaleSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsMarketSkewScaleSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketSkewScaleSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketSkewScaleSetIterator{contract: _Contracts.contract, event: "MarketSkewScaleSet", logs: logs, sub: sub}, nil
}

// WatchMarketSkewScaleSet is a free log subscription operation binding the contract event 0x786fdfd5a0e146d8f4878876a8439ff01436e5969b14682e12d07d7e926b157c.
//
// Solidity: event MarketSkewScaleSet(uint256 indexed synthMarketId, uint256 skewScale)
func (_Contracts *ContractsFilterer) WatchMarketSkewScaleSet(opts *bind.WatchOpts, sink chan<- *ContractsMarketSkewScaleSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketSkewScaleSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketSkewScaleSet)
				if err := _Contracts.contract.UnpackLog(event, "MarketSkewScaleSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseMarketSkewScaleSet(log types.Log) (*ContractsMarketSkewScaleSet, error) {
	event := new(ContractsMarketSkewScaleSet)
	if err := _Contracts.contract.UnpackLog(event, "MarketSkewScaleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketUtilizationFeesSetIterator is returned from FilterMarketUtilizationFeesSet and is used to iterate over the raw logs and unpacked data for MarketUtilizationFeesSet events raised by the Contracts contract.
type ContractsMarketUtilizationFeesSetIterator struct {
	Event *ContractsMarketUtilizationFeesSet // Event containing the contract specifics and raw log

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
func (it *ContractsMarketUtilizationFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketUtilizationFeesSet)
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
		it.Event = new(ContractsMarketUtilizationFeesSet)
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
func (it *ContractsMarketUtilizationFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketUtilizationFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketUtilizationFeesSet represents a MarketUtilizationFeesSet event raised by the Contracts contract.
type ContractsMarketUtilizationFeesSet struct {
	SynthMarketId      *big.Int
	UtilizationFeeRate *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketUtilizationFeesSet is a free log retrieval operation binding the contract event 0x83dfad56ac61e49feb43345b8c73b6d45eb121decc66b1709ca0413b31c64f63.
//
// Solidity: event MarketUtilizationFeesSet(uint256 indexed synthMarketId, uint256 utilizationFeeRate)
func (_Contracts *ContractsFilterer) FilterMarketUtilizationFeesSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsMarketUtilizationFeesSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketUtilizationFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketUtilizationFeesSetIterator{contract: _Contracts.contract, event: "MarketUtilizationFeesSet", logs: logs, sub: sub}, nil
}

// WatchMarketUtilizationFeesSet is a free log subscription operation binding the contract event 0x83dfad56ac61e49feb43345b8c73b6d45eb121decc66b1709ca0413b31c64f63.
//
// Solidity: event MarketUtilizationFeesSet(uint256 indexed synthMarketId, uint256 utilizationFeeRate)
func (_Contracts *ContractsFilterer) WatchMarketUtilizationFeesSet(opts *bind.WatchOpts, sink chan<- *ContractsMarketUtilizationFeesSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketUtilizationFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketUtilizationFeesSet)
				if err := _Contracts.contract.UnpackLog(event, "MarketUtilizationFeesSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseMarketUtilizationFeesSet(log types.Log) (*ContractsMarketUtilizationFeesSet, error) {
	event := new(ContractsMarketUtilizationFeesSet)
	if err := _Contracts.contract.UnpackLog(event, "MarketUtilizationFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Contracts contract.
type ContractsOrderCancelledIterator struct {
	Event *ContractsOrderCancelled // Event containing the contract specifics and raw log

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
func (it *ContractsOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOrderCancelled)
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
		it.Event = new(ContractsOrderCancelled)
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
func (it *ContractsOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOrderCancelled represents a OrderCancelled event raised by the Contracts contract.
type ContractsOrderCancelled struct {
	MarketId        *big.Int
	AsyncOrderId    *big.Int
	AsyncOrderClaim AsyncOrderClaimData
	Sender          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xa57ffc5057d10e88a0dd7713d14130f9638d680af94401e7f1f4fba44a3ad5e2.
//
// Solidity: event OrderCancelled(uint128 indexed marketId, uint128 indexed asyncOrderId, (uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim, address indexed sender)
func (_Contracts *ContractsFilterer) FilterOrderCancelled(opts *bind.FilterOpts, marketId []*big.Int, asyncOrderId []*big.Int, sender []common.Address) (*ContractsOrderCancelledIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OrderCancelled", marketIdRule, asyncOrderIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOrderCancelledIterator{contract: _Contracts.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xa57ffc5057d10e88a0dd7713d14130f9638d680af94401e7f1f4fba44a3ad5e2.
//
// Solidity: event OrderCancelled(uint128 indexed marketId, uint128 indexed asyncOrderId, (uint128,address,uint8,uint256,uint256,uint256,uint256,uint256,address) asyncOrderClaim, address indexed sender)
func (_Contracts *ContractsFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *ContractsOrderCancelled, marketId []*big.Int, asyncOrderId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OrderCancelled", marketIdRule, asyncOrderIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOrderCancelled)
				if err := _Contracts.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOrderCancelled(log types.Log) (*ContractsOrderCancelled, error) {
	event := new(ContractsOrderCancelled)
	if err := _Contracts.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOrderCommittedIterator is returned from FilterOrderCommitted and is used to iterate over the raw logs and unpacked data for OrderCommitted events raised by the Contracts contract.
type ContractsOrderCommittedIterator struct {
	Event *ContractsOrderCommitted // Event containing the contract specifics and raw log

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
func (it *ContractsOrderCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOrderCommitted)
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
		it.Event = new(ContractsOrderCommitted)
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
func (it *ContractsOrderCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOrderCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOrderCommitted represents a OrderCommitted event raised by the Contracts contract.
type ContractsOrderCommitted struct {
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
func (_Contracts *ContractsFilterer) FilterOrderCommitted(opts *bind.FilterOpts, marketId []*big.Int, orderType []uint8, sender []common.Address) (*ContractsOrderCommittedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OrderCommitted", marketIdRule, orderTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOrderCommittedIterator{contract: _Contracts.contract, event: "OrderCommitted", logs: logs, sub: sub}, nil
}

// WatchOrderCommitted is a free log subscription operation binding the contract event 0xb26c216bf0a127dddc2431e4d8ca845513c8e6fb80e754296e7afab1ce92722f.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint8 indexed orderType, uint256 amountProvided, uint128 asyncOrderId, address indexed sender, address referrer)
func (_Contracts *ContractsFilterer) WatchOrderCommitted(opts *bind.WatchOpts, sink chan<- *ContractsOrderCommitted, marketId []*big.Int, orderType []uint8, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OrderCommitted", marketIdRule, orderTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOrderCommitted)
				if err := _Contracts.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOrderCommitted(log types.Log) (*ContractsOrderCommitted, error) {
	event := new(ContractsOrderCommitted)
	if err := _Contracts.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOrderSettledIterator is returned from FilterOrderSettled and is used to iterate over the raw logs and unpacked data for OrderSettled events raised by the Contracts contract.
type ContractsOrderSettledIterator struct {
	Event *ContractsOrderSettled // Event containing the contract specifics and raw log

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
func (it *ContractsOrderSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOrderSettled)
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
		it.Event = new(ContractsOrderSettled)
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
func (it *ContractsOrderSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOrderSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOrderSettled represents a OrderSettled event raised by the Contracts contract.
type ContractsOrderSettled struct {
	MarketId         *big.Int
	AsyncOrderId     *big.Int
	FinalOrderAmount *big.Int
	Fees             OrderFeesData
	CollectedFees    *big.Int
	Settler          common.Address
	Price            *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderSettled is a free log retrieval operation binding the contract event 0x5765b80e91d8a49fe1392afb8431d4ab816fb14e1fe2caabba6aa8a6772d7eb0.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed asyncOrderId, uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address indexed settler, uint256 price)
func (_Contracts *ContractsFilterer) FilterOrderSettled(opts *bind.FilterOpts, marketId []*big.Int, asyncOrderId []*big.Int, settler []common.Address) (*ContractsOrderSettledIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OrderSettled", marketIdRule, asyncOrderIdRule, settlerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOrderSettledIterator{contract: _Contracts.contract, event: "OrderSettled", logs: logs, sub: sub}, nil
}

// WatchOrderSettled is a free log subscription operation binding the contract event 0x5765b80e91d8a49fe1392afb8431d4ab816fb14e1fe2caabba6aa8a6772d7eb0.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed asyncOrderId, uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address indexed settler, uint256 price)
func (_Contracts *ContractsFilterer) WatchOrderSettled(opts *bind.WatchOpts, sink chan<- *ContractsOrderSettled, marketId []*big.Int, asyncOrderId []*big.Int, settler []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OrderSettled", marketIdRule, asyncOrderIdRule, settlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOrderSettled)
				if err := _Contracts.contract.UnpackLog(event, "OrderSettled", log); err != nil {
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

// ParseOrderSettled is a log parse operation binding the contract event 0x5765b80e91d8a49fe1392afb8431d4ab816fb14e1fe2caabba6aa8a6772d7eb0.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed asyncOrderId, uint256 finalOrderAmount, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address indexed settler, uint256 price)
func (_Contracts *ContractsFilterer) ParseOrderSettled(log types.Log) (*ContractsOrderSettled, error) {
	event := new(ContractsOrderSettled)
	if err := _Contracts.contract.UnpackLog(event, "OrderSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Contracts contract.
type ContractsOwnerChangedIterator struct {
	Event *ContractsOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ContractsOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnerChanged)
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
		it.Event = new(ContractsOwnerChanged)
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
func (it *ContractsOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnerChanged represents a OwnerChanged event raised by the Contracts contract.
type ContractsOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Contracts *ContractsFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*ContractsOwnerChangedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &ContractsOwnerChangedIterator{contract: _Contracts.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Contracts *ContractsFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ContractsOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnerChanged)
				if err := _Contracts.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOwnerChanged(log types.Log) (*ContractsOwnerChanged, error) {
	event := new(ContractsOwnerChanged)
	if err := _Contracts.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Contracts contract.
type ContractsOwnerNominatedIterator struct {
	Event *ContractsOwnerNominated // Event containing the contract specifics and raw log

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
func (it *ContractsOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnerNominated)
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
		it.Event = new(ContractsOwnerNominated)
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
func (it *ContractsOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnerNominated represents a OwnerNominated event raised by the Contracts contract.
type ContractsOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Contracts *ContractsFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*ContractsOwnerNominatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &ContractsOwnerNominatedIterator{contract: _Contracts.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Contracts *ContractsFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *ContractsOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnerNominated)
				if err := _Contracts.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOwnerNominated(log types.Log) (*ContractsOwnerNominated, error) {
	event := new(ContractsOwnerNominated)
	if err := _Contracts.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsReferrerShareUpdatedIterator is returned from FilterReferrerShareUpdated and is used to iterate over the raw logs and unpacked data for ReferrerShareUpdated events raised by the Contracts contract.
type ContractsReferrerShareUpdatedIterator struct {
	Event *ContractsReferrerShareUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsReferrerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsReferrerShareUpdated)
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
		it.Event = new(ContractsReferrerShareUpdated)
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
func (it *ContractsReferrerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsReferrerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsReferrerShareUpdated represents a ReferrerShareUpdated event raised by the Contracts contract.
type ContractsReferrerShareUpdated struct {
	MarketId        *big.Int
	Referrer        common.Address
	SharePercentage *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReferrerShareUpdated is a free log retrieval operation binding the contract event 0xd2a3339bb3c610e9030023c1cb3e89374fe0ebd7e37faee9b3d343f33e9df2fb.
//
// Solidity: event ReferrerShareUpdated(uint128 indexed marketId, address referrer, uint256 sharePercentage)
func (_Contracts *ContractsFilterer) FilterReferrerShareUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsReferrerShareUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ReferrerShareUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsReferrerShareUpdatedIterator{contract: _Contracts.contract, event: "ReferrerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchReferrerShareUpdated is a free log subscription operation binding the contract event 0xd2a3339bb3c610e9030023c1cb3e89374fe0ebd7e37faee9b3d343f33e9df2fb.
//
// Solidity: event ReferrerShareUpdated(uint128 indexed marketId, address referrer, uint256 sharePercentage)
func (_Contracts *ContractsFilterer) WatchReferrerShareUpdated(opts *bind.WatchOpts, sink chan<- *ContractsReferrerShareUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ReferrerShareUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsReferrerShareUpdated)
				if err := _Contracts.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseReferrerShareUpdated(log types.Log) (*ContractsReferrerShareUpdated, error) {
	event := new(ContractsReferrerShareUpdated)
	if err := _Contracts.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSettlementStrategyAddedIterator is returned from FilterSettlementStrategyAdded and is used to iterate over the raw logs and unpacked data for SettlementStrategyAdded events raised by the Contracts contract.
type ContractsSettlementStrategyAddedIterator struct {
	Event *ContractsSettlementStrategyAdded // Event containing the contract specifics and raw log

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
func (it *ContractsSettlementStrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSettlementStrategyAdded)
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
		it.Event = new(ContractsSettlementStrategyAdded)
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
func (it *ContractsSettlementStrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSettlementStrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSettlementStrategyAdded represents a SettlementStrategyAdded event raised by the Contracts contract.
type ContractsSettlementStrategyAdded struct {
	SynthMarketId *big.Int
	StrategyId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyAdded is a free log retrieval operation binding the contract event 0x91750a8e3d84ed1bccdc79dcecf63cc1b6f83b5bf8293bf86628cbb248e487f1.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed synthMarketId, uint256 indexed strategyId)
func (_Contracts *ContractsFilterer) FilterSettlementStrategyAdded(opts *bind.FilterOpts, synthMarketId []*big.Int, strategyId []*big.Int) (*ContractsSettlementStrategyAddedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SettlementStrategyAdded", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSettlementStrategyAddedIterator{contract: _Contracts.contract, event: "SettlementStrategyAdded", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyAdded is a free log subscription operation binding the contract event 0x91750a8e3d84ed1bccdc79dcecf63cc1b6f83b5bf8293bf86628cbb248e487f1.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed synthMarketId, uint256 indexed strategyId)
func (_Contracts *ContractsFilterer) WatchSettlementStrategyAdded(opts *bind.WatchOpts, sink chan<- *ContractsSettlementStrategyAdded, synthMarketId []*big.Int, strategyId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SettlementStrategyAdded", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSettlementStrategyAdded)
				if err := _Contracts.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSettlementStrategyAdded(log types.Log) (*ContractsSettlementStrategyAdded, error) {
	event := new(ContractsSettlementStrategyAdded)
	if err := _Contracts.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSettlementStrategyUpdatedIterator is returned from FilterSettlementStrategyUpdated and is used to iterate over the raw logs and unpacked data for SettlementStrategyUpdated events raised by the Contracts contract.
type ContractsSettlementStrategyUpdatedIterator struct {
	Event *ContractsSettlementStrategyUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsSettlementStrategyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSettlementStrategyUpdated)
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
		it.Event = new(ContractsSettlementStrategyUpdated)
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
func (it *ContractsSettlementStrategyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSettlementStrategyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSettlementStrategyUpdated represents a SettlementStrategyUpdated event raised by the Contracts contract.
type ContractsSettlementStrategyUpdated struct {
	SynthMarketId *big.Int
	StrategyId    *big.Int
	Enabled       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyUpdated is a free log retrieval operation binding the contract event 0xb0187d91c1ffe66740e5f7619aab4d868a85ac3698b026bed3f420d871b6c048.
//
// Solidity: event SettlementStrategyUpdated(uint128 indexed synthMarketId, uint256 indexed strategyId, bool enabled)
func (_Contracts *ContractsFilterer) FilterSettlementStrategyUpdated(opts *bind.FilterOpts, synthMarketId []*big.Int, strategyId []*big.Int) (*ContractsSettlementStrategyUpdatedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SettlementStrategyUpdated", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSettlementStrategyUpdatedIterator{contract: _Contracts.contract, event: "SettlementStrategyUpdated", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyUpdated is a free log subscription operation binding the contract event 0xb0187d91c1ffe66740e5f7619aab4d868a85ac3698b026bed3f420d871b6c048.
//
// Solidity: event SettlementStrategyUpdated(uint128 indexed synthMarketId, uint256 indexed strategyId, bool enabled)
func (_Contracts *ContractsFilterer) WatchSettlementStrategyUpdated(opts *bind.WatchOpts, sink chan<- *ContractsSettlementStrategyUpdated, synthMarketId []*big.Int, strategyId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SettlementStrategyUpdated", synthMarketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSettlementStrategyUpdated)
				if err := _Contracts.contract.UnpackLog(event, "SettlementStrategyUpdated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSettlementStrategyUpdated(log types.Log) (*ContractsSettlementStrategyUpdated, error) {
	event := new(ContractsSettlementStrategyUpdated)
	if err := _Contracts.contract.UnpackLog(event, "SettlementStrategyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSynthBoughtIterator is returned from FilterSynthBought and is used to iterate over the raw logs and unpacked data for SynthBought events raised by the Contracts contract.
type ContractsSynthBoughtIterator struct {
	Event *ContractsSynthBought // Event containing the contract specifics and raw log

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
func (it *ContractsSynthBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSynthBought)
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
		it.Event = new(ContractsSynthBought)
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
func (it *ContractsSynthBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSynthBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSynthBought represents a SynthBought event raised by the Contracts contract.
type ContractsSynthBought struct {
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
func (_Contracts *ContractsFilterer) FilterSynthBought(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsSynthBoughtIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SynthBought", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSynthBoughtIterator{contract: _Contracts.contract, event: "SynthBought", logs: logs, sub: sub}, nil
}

// WatchSynthBought is a free log subscription operation binding the contract event 0xac82d63e679c7d862613aa8b5ccd94f9adc4986763ab14bb3351ab9092ef1303.
//
// Solidity: event SynthBought(uint256 indexed synthMarketId, uint256 synthReturned, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address referrer, uint256 price)
func (_Contracts *ContractsFilterer) WatchSynthBought(opts *bind.WatchOpts, sink chan<- *ContractsSynthBought, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SynthBought", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSynthBought)
				if err := _Contracts.contract.UnpackLog(event, "SynthBought", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSynthBought(log types.Log) (*ContractsSynthBought, error) {
	event := new(ContractsSynthBought)
	if err := _Contracts.contract.UnpackLog(event, "SynthBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSynthImplementationSetIterator is returned from FilterSynthImplementationSet and is used to iterate over the raw logs and unpacked data for SynthImplementationSet events raised by the Contracts contract.
type ContractsSynthImplementationSetIterator struct {
	Event *ContractsSynthImplementationSet // Event containing the contract specifics and raw log

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
func (it *ContractsSynthImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSynthImplementationSet)
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
		it.Event = new(ContractsSynthImplementationSet)
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
func (it *ContractsSynthImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSynthImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSynthImplementationSet represents a SynthImplementationSet event raised by the Contracts contract.
type ContractsSynthImplementationSet struct {
	SynthImplementation common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSynthImplementationSet is a free log retrieval operation binding the contract event 0xafffc48b3243eba10d901f21ba761ad741f18a23feed86ca425df4974d3314b0.
//
// Solidity: event SynthImplementationSet(address synthImplementation)
func (_Contracts *ContractsFilterer) FilterSynthImplementationSet(opts *bind.FilterOpts) (*ContractsSynthImplementationSetIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SynthImplementationSet")
	if err != nil {
		return nil, err
	}
	return &ContractsSynthImplementationSetIterator{contract: _Contracts.contract, event: "SynthImplementationSet", logs: logs, sub: sub}, nil
}

// WatchSynthImplementationSet is a free log subscription operation binding the contract event 0xafffc48b3243eba10d901f21ba761ad741f18a23feed86ca425df4974d3314b0.
//
// Solidity: event SynthImplementationSet(address synthImplementation)
func (_Contracts *ContractsFilterer) WatchSynthImplementationSet(opts *bind.WatchOpts, sink chan<- *ContractsSynthImplementationSet) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SynthImplementationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSynthImplementationSet)
				if err := _Contracts.contract.UnpackLog(event, "SynthImplementationSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSynthImplementationSet(log types.Log) (*ContractsSynthImplementationSet, error) {
	event := new(ContractsSynthImplementationSet)
	if err := _Contracts.contract.UnpackLog(event, "SynthImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSynthImplementationUpgradedIterator is returned from FilterSynthImplementationUpgraded and is used to iterate over the raw logs and unpacked data for SynthImplementationUpgraded events raised by the Contracts contract.
type ContractsSynthImplementationUpgradedIterator struct {
	Event *ContractsSynthImplementationUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractsSynthImplementationUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSynthImplementationUpgraded)
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
		it.Event = new(ContractsSynthImplementationUpgraded)
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
func (it *ContractsSynthImplementationUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSynthImplementationUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSynthImplementationUpgraded represents a SynthImplementationUpgraded event raised by the Contracts contract.
type ContractsSynthImplementationUpgraded struct {
	SynthMarketId  *big.Int
	Proxy          common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSynthImplementationUpgraded is a free log retrieval operation binding the contract event 0xa7badace8b24daeeb3981497a506d3812b3dc6f147ef3f78bc0e2cc664c50330.
//
// Solidity: event SynthImplementationUpgraded(uint256 indexed synthMarketId, address indexed proxy, address implementation)
func (_Contracts *ContractsFilterer) FilterSynthImplementationUpgraded(opts *bind.FilterOpts, synthMarketId []*big.Int, proxy []common.Address) (*ContractsSynthImplementationUpgradedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SynthImplementationUpgraded", synthMarketIdRule, proxyRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSynthImplementationUpgradedIterator{contract: _Contracts.contract, event: "SynthImplementationUpgraded", logs: logs, sub: sub}, nil
}

// WatchSynthImplementationUpgraded is a free log subscription operation binding the contract event 0xa7badace8b24daeeb3981497a506d3812b3dc6f147ef3f78bc0e2cc664c50330.
//
// Solidity: event SynthImplementationUpgraded(uint256 indexed synthMarketId, address indexed proxy, address implementation)
func (_Contracts *ContractsFilterer) WatchSynthImplementationUpgraded(opts *bind.WatchOpts, sink chan<- *ContractsSynthImplementationUpgraded, synthMarketId []*big.Int, proxy []common.Address) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SynthImplementationUpgraded", synthMarketIdRule, proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSynthImplementationUpgraded)
				if err := _Contracts.contract.UnpackLog(event, "SynthImplementationUpgraded", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSynthImplementationUpgraded(log types.Log) (*ContractsSynthImplementationUpgraded, error) {
	event := new(ContractsSynthImplementationUpgraded)
	if err := _Contracts.contract.UnpackLog(event, "SynthImplementationUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSynthPriceDataUpdatedIterator is returned from FilterSynthPriceDataUpdated and is used to iterate over the raw logs and unpacked data for SynthPriceDataUpdated events raised by the Contracts contract.
type ContractsSynthPriceDataUpdatedIterator struct {
	Event *ContractsSynthPriceDataUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsSynthPriceDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSynthPriceDataUpdated)
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
		it.Event = new(ContractsSynthPriceDataUpdated)
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
func (it *ContractsSynthPriceDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSynthPriceDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSynthPriceDataUpdated represents a SynthPriceDataUpdated event raised by the Contracts contract.
type ContractsSynthPriceDataUpdated struct {
	SynthMarketId *big.Int
	BuyFeedId     [32]byte
	SellFeedId    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSynthPriceDataUpdated is a free log retrieval operation binding the contract event 0x30309eafe094276dfad42e7ac4676aebcf5ea8480c86708ad2a2c9c0335c912e.
//
// Solidity: event SynthPriceDataUpdated(uint256 indexed synthMarketId, bytes32 indexed buyFeedId, bytes32 indexed sellFeedId)
func (_Contracts *ContractsFilterer) FilterSynthPriceDataUpdated(opts *bind.FilterOpts, synthMarketId []*big.Int, buyFeedId [][32]byte, sellFeedId [][32]byte) (*ContractsSynthPriceDataUpdatedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SynthPriceDataUpdated", synthMarketIdRule, buyFeedIdRule, sellFeedIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSynthPriceDataUpdatedIterator{contract: _Contracts.contract, event: "SynthPriceDataUpdated", logs: logs, sub: sub}, nil
}

// WatchSynthPriceDataUpdated is a free log subscription operation binding the contract event 0x30309eafe094276dfad42e7ac4676aebcf5ea8480c86708ad2a2c9c0335c912e.
//
// Solidity: event SynthPriceDataUpdated(uint256 indexed synthMarketId, bytes32 indexed buyFeedId, bytes32 indexed sellFeedId)
func (_Contracts *ContractsFilterer) WatchSynthPriceDataUpdated(opts *bind.WatchOpts, sink chan<- *ContractsSynthPriceDataUpdated, synthMarketId []*big.Int, buyFeedId [][32]byte, sellFeedId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SynthPriceDataUpdated", synthMarketIdRule, buyFeedIdRule, sellFeedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSynthPriceDataUpdated)
				if err := _Contracts.contract.UnpackLog(event, "SynthPriceDataUpdated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSynthPriceDataUpdated(log types.Log) (*ContractsSynthPriceDataUpdated, error) {
	event := new(ContractsSynthPriceDataUpdated)
	if err := _Contracts.contract.UnpackLog(event, "SynthPriceDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSynthRegisteredIterator is returned from FilterSynthRegistered and is used to iterate over the raw logs and unpacked data for SynthRegistered events raised by the Contracts contract.
type ContractsSynthRegisteredIterator struct {
	Event *ContractsSynthRegistered // Event containing the contract specifics and raw log

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
func (it *ContractsSynthRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSynthRegistered)
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
		it.Event = new(ContractsSynthRegistered)
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
func (it *ContractsSynthRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSynthRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSynthRegistered represents a SynthRegistered event raised by the Contracts contract.
type ContractsSynthRegistered struct {
	SynthMarketId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSynthRegistered is a free log retrieval operation binding the contract event 0xfebb2ba957958f85fc4768ead2486373c87540fa3662ddfa3f6c8b564b00b1b1.
//
// Solidity: event SynthRegistered(uint256 indexed synthMarketId)
func (_Contracts *ContractsFilterer) FilterSynthRegistered(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsSynthRegisteredIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SynthRegistered", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSynthRegisteredIterator{contract: _Contracts.contract, event: "SynthRegistered", logs: logs, sub: sub}, nil
}

// WatchSynthRegistered is a free log subscription operation binding the contract event 0xfebb2ba957958f85fc4768ead2486373c87540fa3662ddfa3f6c8b564b00b1b1.
//
// Solidity: event SynthRegistered(uint256 indexed synthMarketId)
func (_Contracts *ContractsFilterer) WatchSynthRegistered(opts *bind.WatchOpts, sink chan<- *ContractsSynthRegistered, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SynthRegistered", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSynthRegistered)
				if err := _Contracts.contract.UnpackLog(event, "SynthRegistered", log); err != nil {
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

// ParseSynthRegistered is a log parse operation binding the contract event 0xfebb2ba957958f85fc4768ead2486373c87540fa3662ddfa3f6c8b564b00b1b1.
//
// Solidity: event SynthRegistered(uint256 indexed synthMarketId)
func (_Contracts *ContractsFilterer) ParseSynthRegistered(log types.Log) (*ContractsSynthRegistered, error) {
	event := new(ContractsSynthRegistered)
	if err := _Contracts.contract.UnpackLog(event, "SynthRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSynthSoldIterator is returned from FilterSynthSold and is used to iterate over the raw logs and unpacked data for SynthSold events raised by the Contracts contract.
type ContractsSynthSoldIterator struct {
	Event *ContractsSynthSold // Event containing the contract specifics and raw log

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
func (it *ContractsSynthSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSynthSold)
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
		it.Event = new(ContractsSynthSold)
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
func (it *ContractsSynthSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSynthSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSynthSold represents a SynthSold event raised by the Contracts contract.
type ContractsSynthSold struct {
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
func (_Contracts *ContractsFilterer) FilterSynthSold(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsSynthSoldIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SynthSold", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSynthSoldIterator{contract: _Contracts.contract, event: "SynthSold", logs: logs, sub: sub}, nil
}

// WatchSynthSold is a free log subscription operation binding the contract event 0x61fa4bb370a2f18a502b3bcf1d0755e53371d58791fa42766aa6386bbefb594a.
//
// Solidity: event SynthSold(uint256 indexed synthMarketId, uint256 amountReturned, (uint256,uint256,int256,int256) fees, uint256 collectedFees, address referrer, uint256 price)
func (_Contracts *ContractsFilterer) WatchSynthSold(opts *bind.WatchOpts, sink chan<- *ContractsSynthSold, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SynthSold", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSynthSold)
				if err := _Contracts.contract.UnpackLog(event, "SynthSold", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSynthSold(log types.Log) (*ContractsSynthSold, error) {
	event := new(ContractsSynthSold)
	if err := _Contracts.contract.UnpackLog(event, "SynthSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSynthUnwrappedIterator is returned from FilterSynthUnwrapped and is used to iterate over the raw logs and unpacked data for SynthUnwrapped events raised by the Contracts contract.
type ContractsSynthUnwrappedIterator struct {
	Event *ContractsSynthUnwrapped // Event containing the contract specifics and raw log

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
func (it *ContractsSynthUnwrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSynthUnwrapped)
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
		it.Event = new(ContractsSynthUnwrapped)
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
func (it *ContractsSynthUnwrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSynthUnwrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSynthUnwrapped represents a SynthUnwrapped event raised by the Contracts contract.
type ContractsSynthUnwrapped struct {
	SynthMarketId   *big.Int
	AmountUnwrapped *big.Int
	Fees            OrderFeesData
	FeesCollected   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSynthUnwrapped is a free log retrieval operation binding the contract event 0xa1dd74fb936c7942732e4355961ca6944ca6c6744121ace0d9a1203d664231b3.
//
// Solidity: event SynthUnwrapped(uint256 indexed synthMarketId, uint256 amountUnwrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_Contracts *ContractsFilterer) FilterSynthUnwrapped(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsSynthUnwrappedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SynthUnwrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSynthUnwrappedIterator{contract: _Contracts.contract, event: "SynthUnwrapped", logs: logs, sub: sub}, nil
}

// WatchSynthUnwrapped is a free log subscription operation binding the contract event 0xa1dd74fb936c7942732e4355961ca6944ca6c6744121ace0d9a1203d664231b3.
//
// Solidity: event SynthUnwrapped(uint256 indexed synthMarketId, uint256 amountUnwrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_Contracts *ContractsFilterer) WatchSynthUnwrapped(opts *bind.WatchOpts, sink chan<- *ContractsSynthUnwrapped, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SynthUnwrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSynthUnwrapped)
				if err := _Contracts.contract.UnpackLog(event, "SynthUnwrapped", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSynthUnwrapped(log types.Log) (*ContractsSynthUnwrapped, error) {
	event := new(ContractsSynthUnwrapped)
	if err := _Contracts.contract.UnpackLog(event, "SynthUnwrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSynthWrappedIterator is returned from FilterSynthWrapped and is used to iterate over the raw logs and unpacked data for SynthWrapped events raised by the Contracts contract.
type ContractsSynthWrappedIterator struct {
	Event *ContractsSynthWrapped // Event containing the contract specifics and raw log

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
func (it *ContractsSynthWrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSynthWrapped)
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
		it.Event = new(ContractsSynthWrapped)
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
func (it *ContractsSynthWrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSynthWrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSynthWrapped represents a SynthWrapped event raised by the Contracts contract.
type ContractsSynthWrapped struct {
	SynthMarketId *big.Int
	AmountWrapped *big.Int
	Fees          OrderFeesData
	FeesCollected *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSynthWrapped is a free log retrieval operation binding the contract event 0xea50ab2f37d37692441c4a16317c1287bb410a3a616a16c49c9ca76d415667ff.
//
// Solidity: event SynthWrapped(uint256 indexed synthMarketId, uint256 amountWrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_Contracts *ContractsFilterer) FilterSynthWrapped(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsSynthWrappedIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SynthWrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSynthWrappedIterator{contract: _Contracts.contract, event: "SynthWrapped", logs: logs, sub: sub}, nil
}

// WatchSynthWrapped is a free log subscription operation binding the contract event 0xea50ab2f37d37692441c4a16317c1287bb410a3a616a16c49c9ca76d415667ff.
//
// Solidity: event SynthWrapped(uint256 indexed synthMarketId, uint256 amountWrapped, (uint256,uint256,int256,int256) fees, uint256 feesCollected)
func (_Contracts *ContractsFilterer) WatchSynthWrapped(opts *bind.WatchOpts, sink chan<- *ContractsSynthWrapped, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SynthWrapped", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSynthWrapped)
				if err := _Contracts.contract.UnpackLog(event, "SynthWrapped", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSynthWrapped(log types.Log) (*ContractsSynthWrapped, error) {
	event := new(ContractsSynthWrapped)
	if err := _Contracts.contract.UnpackLog(event, "SynthWrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSynthetixSystemSetIterator is returned from FilterSynthetixSystemSet and is used to iterate over the raw logs and unpacked data for SynthetixSystemSet events raised by the Contracts contract.
type ContractsSynthetixSystemSetIterator struct {
	Event *ContractsSynthetixSystemSet // Event containing the contract specifics and raw log

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
func (it *ContractsSynthetixSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSynthetixSystemSet)
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
		it.Event = new(ContractsSynthetixSystemSet)
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
func (it *ContractsSynthetixSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSynthetixSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSynthetixSystemSet represents a SynthetixSystemSet event raised by the Contracts contract.
type ContractsSynthetixSystemSet struct {
	Synthetix       common.Address
	UsdTokenAddress common.Address
	OracleManager   common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSynthetixSystemSet is a free log retrieval operation binding the contract event 0x52dccf7e2653d5ae8cf1d18c5499fd530f01920181d81afdf6bf489d897e24aa.
//
// Solidity: event SynthetixSystemSet(address synthetix, address usdTokenAddress, address oracleManager)
func (_Contracts *ContractsFilterer) FilterSynthetixSystemSet(opts *bind.FilterOpts) (*ContractsSynthetixSystemSetIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SynthetixSystemSet")
	if err != nil {
		return nil, err
	}
	return &ContractsSynthetixSystemSetIterator{contract: _Contracts.contract, event: "SynthetixSystemSet", logs: logs, sub: sub}, nil
}

// WatchSynthetixSystemSet is a free log subscription operation binding the contract event 0x52dccf7e2653d5ae8cf1d18c5499fd530f01920181d81afdf6bf489d897e24aa.
//
// Solidity: event SynthetixSystemSet(address synthetix, address usdTokenAddress, address oracleManager)
func (_Contracts *ContractsFilterer) WatchSynthetixSystemSet(opts *bind.WatchOpts, sink chan<- *ContractsSynthetixSystemSet) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SynthetixSystemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSynthetixSystemSet)
				if err := _Contracts.contract.UnpackLog(event, "SynthetixSystemSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSynthetixSystemSet(log types.Log) (*ContractsSynthetixSystemSet, error) {
	event := new(ContractsSynthetixSystemSet)
	if err := _Contracts.contract.UnpackLog(event, "SynthetixSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransactorFixedFeeSetIterator is returned from FilterTransactorFixedFeeSet and is used to iterate over the raw logs and unpacked data for TransactorFixedFeeSet events raised by the Contracts contract.
type ContractsTransactorFixedFeeSetIterator struct {
	Event *ContractsTransactorFixedFeeSet // Event containing the contract specifics and raw log

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
func (it *ContractsTransactorFixedFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransactorFixedFeeSet)
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
		it.Event = new(ContractsTransactorFixedFeeSet)
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
func (it *ContractsTransactorFixedFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransactorFixedFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransactorFixedFeeSet represents a TransactorFixedFeeSet event raised by the Contracts contract.
type ContractsTransactorFixedFeeSet struct {
	SynthMarketId  *big.Int
	Transactor     common.Address
	FixedFeeAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransactorFixedFeeSet is a free log retrieval operation binding the contract event 0xeed7c7ebc4a7e7456a5b14f961bbe55d026f35a2a2b52d1ad43fe11c348df24a.
//
// Solidity: event TransactorFixedFeeSet(uint256 indexed synthMarketId, address transactor, uint256 fixedFeeAmount)
func (_Contracts *ContractsFilterer) FilterTransactorFixedFeeSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsTransactorFixedFeeSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "TransactorFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactorFixedFeeSetIterator{contract: _Contracts.contract, event: "TransactorFixedFeeSet", logs: logs, sub: sub}, nil
}

// WatchTransactorFixedFeeSet is a free log subscription operation binding the contract event 0xeed7c7ebc4a7e7456a5b14f961bbe55d026f35a2a2b52d1ad43fe11c348df24a.
//
// Solidity: event TransactorFixedFeeSet(uint256 indexed synthMarketId, address transactor, uint256 fixedFeeAmount)
func (_Contracts *ContractsFilterer) WatchTransactorFixedFeeSet(opts *bind.WatchOpts, sink chan<- *ContractsTransactorFixedFeeSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "TransactorFixedFeeSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransactorFixedFeeSet)
				if err := _Contracts.contract.UnpackLog(event, "TransactorFixedFeeSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseTransactorFixedFeeSet(log types.Log) (*ContractsTransactorFixedFeeSet, error) {
	event := new(ContractsTransactorFixedFeeSet)
	if err := _Contracts.contract.UnpackLog(event, "TransactorFixedFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Contracts contract.
type ContractsUpgradedIterator struct {
	Event *ContractsUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUpgraded)
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
		it.Event = new(ContractsUpgraded)
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
func (it *ContractsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUpgraded represents a Upgraded event raised by the Contracts contract.
type ContractsUpgraded struct {
	Self           common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_Contracts *ContractsFilterer) FilterUpgraded(opts *bind.FilterOpts, self []common.Address) (*ContractsUpgradedIterator, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return &ContractsUpgradedIterator{contract: _Contracts.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_Contracts *ContractsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractsUpgraded, self []common.Address) (event.Subscription, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUpgraded)
				if err := _Contracts.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseUpgraded(log types.Log) (*ContractsUpgraded, error) {
	event := new(ContractsUpgraded)
	if err := _Contracts.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsWrapperFeesSetIterator is returned from FilterWrapperFeesSet and is used to iterate over the raw logs and unpacked data for WrapperFeesSet events raised by the Contracts contract.
type ContractsWrapperFeesSetIterator struct {
	Event *ContractsWrapperFeesSet // Event containing the contract specifics and raw log

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
func (it *ContractsWrapperFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsWrapperFeesSet)
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
		it.Event = new(ContractsWrapperFeesSet)
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
func (it *ContractsWrapperFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsWrapperFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsWrapperFeesSet represents a WrapperFeesSet event raised by the Contracts contract.
type ContractsWrapperFeesSet struct {
	SynthMarketId *big.Int
	WrapFee       *big.Int
	UnwrapFee     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWrapperFeesSet is a free log retrieval operation binding the contract event 0x84c5bc20d6e52e92afe6ebc9d85d3e4d35de276ba3f05cae640db053f5b861b8.
//
// Solidity: event WrapperFeesSet(uint256 indexed synthMarketId, int256 wrapFee, int256 unwrapFee)
func (_Contracts *ContractsFilterer) FilterWrapperFeesSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsWrapperFeesSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "WrapperFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsWrapperFeesSetIterator{contract: _Contracts.contract, event: "WrapperFeesSet", logs: logs, sub: sub}, nil
}

// WatchWrapperFeesSet is a free log subscription operation binding the contract event 0x84c5bc20d6e52e92afe6ebc9d85d3e4d35de276ba3f05cae640db053f5b861b8.
//
// Solidity: event WrapperFeesSet(uint256 indexed synthMarketId, int256 wrapFee, int256 unwrapFee)
func (_Contracts *ContractsFilterer) WatchWrapperFeesSet(opts *bind.WatchOpts, sink chan<- *ContractsWrapperFeesSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "WrapperFeesSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsWrapperFeesSet)
				if err := _Contracts.contract.UnpackLog(event, "WrapperFeesSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseWrapperFeesSet(log types.Log) (*ContractsWrapperFeesSet, error) {
	event := new(ContractsWrapperFeesSet)
	if err := _Contracts.contract.UnpackLog(event, "WrapperFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsWrapperSetIterator is returned from FilterWrapperSet and is used to iterate over the raw logs and unpacked data for WrapperSet events raised by the Contracts contract.
type ContractsWrapperSetIterator struct {
	Event *ContractsWrapperSet // Event containing the contract specifics and raw log

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
func (it *ContractsWrapperSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsWrapperSet)
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
		it.Event = new(ContractsWrapperSet)
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
func (it *ContractsWrapperSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsWrapperSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsWrapperSet represents a WrapperSet event raised by the Contracts contract.
type ContractsWrapperSet struct {
	SynthMarketId      *big.Int
	WrapCollateralType common.Address
	MaxWrappableAmount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWrapperSet is a free log retrieval operation binding the contract event 0xf6b8d296783aecfc5d372dff3e3e802ab63338637f9a2f3e2aae1e745c148def.
//
// Solidity: event WrapperSet(uint256 indexed synthMarketId, address indexed wrapCollateralType, uint256 maxWrappableAmount)
func (_Contracts *ContractsFilterer) FilterWrapperSet(opts *bind.FilterOpts, synthMarketId []*big.Int, wrapCollateralType []common.Address) (*ContractsWrapperSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var wrapCollateralTypeRule []interface{}
	for _, wrapCollateralTypeItem := range wrapCollateralType {
		wrapCollateralTypeRule = append(wrapCollateralTypeRule, wrapCollateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "WrapperSet", synthMarketIdRule, wrapCollateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsWrapperSetIterator{contract: _Contracts.contract, event: "WrapperSet", logs: logs, sub: sub}, nil
}

// WatchWrapperSet is a free log subscription operation binding the contract event 0xf6b8d296783aecfc5d372dff3e3e802ab63338637f9a2f3e2aae1e745c148def.
//
// Solidity: event WrapperSet(uint256 indexed synthMarketId, address indexed wrapCollateralType, uint256 maxWrappableAmount)
func (_Contracts *ContractsFilterer) WatchWrapperSet(opts *bind.WatchOpts, sink chan<- *ContractsWrapperSet, synthMarketId []*big.Int, wrapCollateralType []common.Address) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}
	var wrapCollateralTypeRule []interface{}
	for _, wrapCollateralTypeItem := range wrapCollateralType {
		wrapCollateralTypeRule = append(wrapCollateralTypeRule, wrapCollateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "WrapperSet", synthMarketIdRule, wrapCollateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsWrapperSet)
				if err := _Contracts.contract.UnpackLog(event, "WrapperSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseWrapperSet(log types.Log) (*ContractsWrapperSet, error) {
	event := new(ContractsWrapperSet)
	if err := _Contracts.contract.UnpackLog(event, "WrapperSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
