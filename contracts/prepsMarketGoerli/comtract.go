// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package prepsMarketGoerli

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

// AsyncOrderData is an auto generated low-level Go binding around an user-defined struct.
type AsyncOrderData struct {
	SettlementTime *big.Int
	Request        AsyncOrderOrderCommitmentRequest
}

// AsyncOrderOrderCommitmentRequest is an auto generated low-level Go binding around an user-defined struct.
type AsyncOrderOrderCommitmentRequest struct {
	MarketId             *big.Int
	AccountId            *big.Int
	SizeDelta            *big.Int
	SettlementStrategyId *big.Int
	AcceptablePrice      *big.Int
	TrackingCode         [32]byte
	Referrer             common.Address
}

// IAccountModuleAccountPermissions is an auto generated low-level Go binding around an user-defined struct.
type IAccountModuleAccountPermissions struct {
	User        common.Address
	Permissions [][32]byte
}

// IPerpsMarketModuleMarketSummary is an auto generated low-level Go binding around an user-defined struct.
type IPerpsMarketModuleMarketSummary struct {
	Skew                   *big.Int
	Size                   *big.Int
	MaxOpenInterest        *big.Int
	CurrentFundingRate     *big.Int
	CurrentFundingVelocity *big.Int
	IndexPrice             *big.Int
}

// SettlementStrategyData is an auto generated low-level Go binding around an user-defined struct.
type SettlementStrategyData struct {
	StrategyType              uint8
	SettlementDelay           *big.Int
	SettlementWindowDuration  *big.Int
	PriceWindowDuration       *big.Int
	PriceVerificationContract common.Address
	FeedId                    [32]byte
	Url                       string
	SettlementReward          *big.Int
	PriceDeviationTolerance   *big.Int
	Disabled                  bool
}

// PrepsMarketGoerliMetaData contains all meta data concerning the PrepsMarketGoerli contract.
var PrepsMarketGoerliMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"InvalidAccountId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"InvalidPermission\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"OnlyAccountTokenProxy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"PermissionNotGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedAccountId\",\"type\":\"uint128\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountLastInteraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountPermissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"permissions\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIAccountModule.AccountPermissions[]\",\"name\":\"accountPerms\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"notifyAccountTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"renouncePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"InvalidMarket\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parameter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PerpsMarketNotInitialized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"globalPerpsMarketId\",\"type\":\"uint128\"}],\"name\":\"FactoryInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"marketSymbol\",\"type\":\"string\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedMarketId\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"marketSymbol\",\"type\":\"string\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeFactory\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"minimumCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"reportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISpotMarketSystem\",\"name\":\"spotMarket\",\"type\":\"address\"}],\"name\":\"setSpotMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISynthetixSystem\",\"name\":\"synthetix\",\"type\":\"address\"}],\"name\":\"setSynthetix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountLiquidatable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateralAvailableForWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"}],\"name\":\"InvalidAmountDelta\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"MaxCollateralExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint128ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOrderExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"PriceFeedNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"SynthNotEnabledForCollateral\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"CollateralModified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAvailableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableMargin\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getOpenPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"totalPnl\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"accruedFunding\",\"type\":\"int256\"},{\"internalType\":\"int128\",\"name\":\"positionSize\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getRequiredMargins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredInitialMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMaintenanceMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAccumulatedLiquidationRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getWithdrawableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"withdrawableMargin\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"}],\"name\":\"modifyCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"totalAccountOpenInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"totalCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"currentFundingRate\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"currentFundingVelocity\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"orderSize\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"fillPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"skew\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIPerpsMarketModule.MarketSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"indexPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"maxOpenInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"metadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"skew\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"}],\"name\":\"AcceptablePriceExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"availableMargin\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"minMargin\",\"type\":\"uint256\"}],\"name\":\"InsufficientMargin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"}],\"name\":\"InvalidSettlementStrategy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"newSideSize\",\"type\":\"int256\"}],\"name\":\"MaxOpenInterestReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSizeOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OrderCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"PreviousOrderExpired\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"commitment\",\"type\":\"tuple\"}],\"name\":\"commitOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"internalType\":\"structAsyncOrder.Data\",\"name\":\"retOrder\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"}],\"name\":\"computeOrderFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"internalType\":\"structAsyncOrder.Data\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"}],\"name\":\"requiredMarginForOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredMargin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leftover\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarginError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deviation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tolerance\",\"type\":\"uint256\"}],\"name\":\"PriceDeviationToleranceExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"}],\"name\":\"SettlementStrategyNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementExpiration\",\"type\":\"uint256\"}],\"name\":\"SettlementWindowExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"}],\"name\":\"SettlementWindowNotOpen\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"skew\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"sizeDelta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"}],\"name\":\"MarketUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"accruedFunding\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"newSize\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"}],\"name\":\"OrderSettled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"settlePythOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"NotEligibleForLiquidation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fullLiquidation\",\"type\":\"bool\"}],\"name\":\"AccountLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLiquidated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"currentPositionSize\",\"type\":\"int128\"}],\"name\":\"PositionLiquidated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidateFlagged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"name\":\"FundingParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maintenanceMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"name\":\"LiquidationParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedOiRatioD18\",\"type\":\"uint256\"}],\"name\":\"LockedOiRatioSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"MarketPriceDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"name\":\"MaxMarketSizeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerFeeRatio\",\"type\":\"uint256\"}],\"name\":\"OrderFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"SettlementStrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SettlementStrategyEnabled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"}],\"name\":\"addSettlementStrategy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getFundingParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getLiquidationParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginScalarD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getLockedOiRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMaxMarketSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getOrderFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"getSettlementStrategy\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"settlementStrategy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"name\":\"setFundingParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginScalarD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"name\":\"setLiquidationParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"lockedOiRatioD18\",\"type\":\"uint256\"}],\"name\":\"setLockedOiRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"name\":\"setMaxMarketSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeRatio\",\"type\":\"uint256\"}],\"name\":\"setOrderFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setSettlementStrategyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"updatePriceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidFeeCollector\",\"type\":\"address\"}],\"name\":\"InvalidFeeCollectorInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"InvalidReferrerShareRatio\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"name\":\"LiquidationRewardGuardsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"MaxCollateralAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"ReferrerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"newSynthDeductionPriority\",\"type\":\"uint128[]\"}],\"name\":\"SynthDeductionPrioritySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidationRewardGuards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"marketIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMaxCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"getReferrerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSynthDeductionPriority\",\"outputs\":[{\"internalType\":\"uint128[]\",\"name\":\"\",\"type\":\"uint128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"name\":\"setLiquidationRewardGuards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"setMaxCollateralAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128[]\",\"name\":\"newSynthDeductionPriority\",\"type\":\"uint128[]\"}],\"name\":\"setSynthDeductionPriority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalGlobalCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"updateReferrerShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PrepsMarketGoerliABI is the input ABI used to generate the binding from.
// Deprecated: Use PrepsMarketGoerliMetaData.ABI instead.
var PrepsMarketGoerliABI = PrepsMarketGoerliMetaData.ABI

// PrepsMarketGoerli is an auto generated Go binding around an Ethereum contract.
type PrepsMarketGoerli struct {
	PrepsMarketGoerliCaller     // Read-only binding to the contract
	PrepsMarketGoerliTransactor // Write-only binding to the contract
	PrepsMarketGoerliFilterer   // Log filterer for contract events
}

// PrepsMarketGoerliCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrepsMarketGoerliCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrepsMarketGoerliTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrepsMarketGoerliTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrepsMarketGoerliFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrepsMarketGoerliFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrepsMarketGoerliSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrepsMarketGoerliSession struct {
	Contract     *PrepsMarketGoerli // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PrepsMarketGoerliCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrepsMarketGoerliCallerSession struct {
	Contract *PrepsMarketGoerliCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PrepsMarketGoerliTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrepsMarketGoerliTransactorSession struct {
	Contract     *PrepsMarketGoerliTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PrepsMarketGoerliRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrepsMarketGoerliRaw struct {
	Contract *PrepsMarketGoerli // Generic contract binding to access the raw methods on
}

// PrepsMarketGoerliCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrepsMarketGoerliCallerRaw struct {
	Contract *PrepsMarketGoerliCaller // Generic read-only contract binding to access the raw methods on
}

// PrepsMarketGoerliTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrepsMarketGoerliTransactorRaw struct {
	Contract *PrepsMarketGoerliTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrepsMarketGoerli creates a new instance of PrepsMarketGoerli, bound to a specific deployed contract.
func NewPrepsMarketGoerli(address common.Address, backend bind.ContractBackend) (*PrepsMarketGoerli, error) {
	contract, err := bindPrepsMarketGoerli(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerli{PrepsMarketGoerliCaller: PrepsMarketGoerliCaller{contract: contract}, PrepsMarketGoerliTransactor: PrepsMarketGoerliTransactor{contract: contract}, PrepsMarketGoerliFilterer: PrepsMarketGoerliFilterer{contract: contract}}, nil
}

// NewPrepsMarketGoerliCaller creates a new read-only instance of PrepsMarketGoerli, bound to a specific deployed contract.
func NewPrepsMarketGoerliCaller(address common.Address, caller bind.ContractCaller) (*PrepsMarketGoerliCaller, error) {
	contract, err := bindPrepsMarketGoerli(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliCaller{contract: contract}, nil
}

// NewPrepsMarketGoerliTransactor creates a new write-only instance of PrepsMarketGoerli, bound to a specific deployed contract.
func NewPrepsMarketGoerliTransactor(address common.Address, transactor bind.ContractTransactor) (*PrepsMarketGoerliTransactor, error) {
	contract, err := bindPrepsMarketGoerli(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliTransactor{contract: contract}, nil
}

// NewPrepsMarketGoerliFilterer creates a new log filterer instance of PrepsMarketGoerli, bound to a specific deployed contract.
func NewPrepsMarketGoerliFilterer(address common.Address, filterer bind.ContractFilterer) (*PrepsMarketGoerliFilterer, error) {
	contract, err := bindPrepsMarketGoerli(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliFilterer{contract: contract}, nil
}

// bindPrepsMarketGoerli binds a generic wrapper to an already deployed contract.
func bindPrepsMarketGoerli(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrepsMarketGoerliMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrepsMarketGoerli *PrepsMarketGoerliRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrepsMarketGoerli.Contract.PrepsMarketGoerliCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrepsMarketGoerli *PrepsMarketGoerliRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.PrepsMarketGoerliTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrepsMarketGoerli *PrepsMarketGoerliRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.PrepsMarketGoerliTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrepsMarketGoerli.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) PRECISION() (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.PRECISION(&_PrepsMarketGoerli.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) PRECISION() (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.PRECISION(&_PrepsMarketGoerli.CallOpts)
}

// ComputeOrderFees is a free data retrieval call binding the contract method 0x98ef15a2.
//
// Solidity: function computeOrderFees(uint128 marketId, int128 sizeDelta) view returns(uint256 orderFees, uint256 fillPrice)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) ComputeOrderFees(opts *bind.CallOpts, marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "computeOrderFees", marketId, sizeDelta)

	outstruct := new(struct {
		OrderFees *big.Int
		FillPrice *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderFees = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FillPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ComputeOrderFees is a free data retrieval call binding the contract method 0x98ef15a2.
//
// Solidity: function computeOrderFees(uint128 marketId, int128 sizeDelta) view returns(uint256 orderFees, uint256 fillPrice)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) ComputeOrderFees(marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.ComputeOrderFees(&_PrepsMarketGoerli.CallOpts, marketId, sizeDelta)
}

// ComputeOrderFees is a free data retrieval call binding the contract method 0x98ef15a2.
//
// Solidity: function computeOrderFees(uint128 marketId, int128 sizeDelta) view returns(uint256 orderFees, uint256 fillPrice)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) ComputeOrderFees(marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.ComputeOrderFees(&_PrepsMarketGoerli.CallOpts, marketId, sizeDelta)
}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) CurrentFundingRate(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "currentFundingRate", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) CurrentFundingRate(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.CurrentFundingRate(&_PrepsMarketGoerli.CallOpts, marketId)
}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) CurrentFundingRate(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.CurrentFundingRate(&_PrepsMarketGoerli.CallOpts, marketId)
}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) CurrentFundingVelocity(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "currentFundingVelocity", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) CurrentFundingVelocity(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.CurrentFundingVelocity(&_PrepsMarketGoerli.CallOpts, marketId)
}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) CurrentFundingVelocity(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.CurrentFundingVelocity(&_PrepsMarketGoerli.CallOpts, marketId)
}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) FillPrice(opts *bind.CallOpts, marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "fillPrice", marketId, orderSize, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) FillPrice(marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.FillPrice(&_PrepsMarketGoerli.CallOpts, marketId, orderSize, price)
}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) FillPrice(marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.FillPrice(&_PrepsMarketGoerli.CallOpts, marketId, orderSize, price)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetAccountLastInteraction(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getAccountLastInteraction", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetAccountLastInteraction(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetAccountLastInteraction(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetAccountOwner(opts *bind.CallOpts, accountId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getAccountOwner", accountId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetAccountOwner(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetAccountOwner(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetAccountPermissions(opts *bind.CallOpts, accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getAccountPermissions", accountId)

	if err != nil {
		return *new([]IAccountModuleAccountPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAccountModuleAccountPermissions)).(*[]IAccountModuleAccountPermissions)

	return out0, err

}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _PrepsMarketGoerli.Contract.GetAccountPermissions(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _PrepsMarketGoerli.Contract.GetAccountPermissions(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetAccountTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getAccountTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetAccountTokenAddress() (common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetAccountTokenAddress(&_PrepsMarketGoerli.CallOpts)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetAccountTokenAddress() (common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetAccountTokenAddress(&_PrepsMarketGoerli.CallOpts)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetAssociatedSystem(opts *bind.CallOpts, id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getAssociatedSystem", id)

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
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _PrepsMarketGoerli.Contract.GetAssociatedSystem(&_PrepsMarketGoerli.CallOpts, id)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _PrepsMarketGoerli.Contract.GetAssociatedSystem(&_PrepsMarketGoerli.CallOpts, id)
}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetAvailableMargin(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getAvailableMargin", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetAvailableMargin(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetAvailableMargin(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetCollateralAmount(opts *bind.CallOpts, accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getCollateralAmount", accountId, synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetCollateralAmount(accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetCollateralAmount(&_PrepsMarketGoerli.CallOpts, accountId, synthMarketId)
}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetCollateralAmount(accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetCollateralAmount(&_PrepsMarketGoerli.CallOpts, accountId, synthMarketId)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetDeniers(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getDeniers", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetDeniers(&_PrepsMarketGoerli.CallOpts, feature)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetDeniers(&_PrepsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetFeatureFlagAllowAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getFeatureFlagAllowAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _PrepsMarketGoerli.Contract.GetFeatureFlagAllowAll(&_PrepsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _PrepsMarketGoerli.Contract.GetFeatureFlagAllowAll(&_PrepsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetFeatureFlagAllowlist(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getFeatureFlagAllowlist", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetFeatureFlagAllowlist(&_PrepsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetFeatureFlagAllowlist(&_PrepsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetFeatureFlagDenyAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getFeatureFlagDenyAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _PrepsMarketGoerli.Contract.GetFeatureFlagDenyAll(&_PrepsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _PrepsMarketGoerli.Contract.GetFeatureFlagDenyAll(&_PrepsMarketGoerli.CallOpts, feature)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetFeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getFeeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetFeeCollector() (common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetFeeCollector(&_PrepsMarketGoerli.CallOpts)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetFeeCollector() (common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetFeeCollector(&_PrepsMarketGoerli.CallOpts)
}

// GetFundingParameters is a free data retrieval call binding the contract method 0x1b68d8fa.
//
// Solidity: function getFundingParameters(uint128 marketId) view returns(uint256 skewScale, uint256 maxFundingVelocity)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetFundingParameters(opts *bind.CallOpts, marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getFundingParameters", marketId)

	outstruct := new(struct {
		SkewScale          *big.Int
		MaxFundingVelocity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SkewScale = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxFundingVelocity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFundingParameters is a free data retrieval call binding the contract method 0x1b68d8fa.
//
// Solidity: function getFundingParameters(uint128 marketId) view returns(uint256 skewScale, uint256 maxFundingVelocity)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetFundingParameters(marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetFundingParameters(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetFundingParameters is a free data retrieval call binding the contract method 0x1b68d8fa.
//
// Solidity: function getFundingParameters(uint128 marketId) view returns(uint256 skewScale, uint256 maxFundingVelocity)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetFundingParameters(marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetFundingParameters(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetImplementation() (common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetImplementation(&_PrepsMarketGoerli.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetImplementation() (common.Address, error) {
	return _PrepsMarketGoerli.Contract.GetImplementation(&_PrepsMarketGoerli.CallOpts)
}

// GetLiquidationParameters is a free data retrieval call binding the contract method 0xf94363a6.
//
// Solidity: function getLiquidationParameters(uint128 marketId) view returns(uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetLiquidationParameters(opts *bind.CallOpts, marketId *big.Int) (struct {
	InitialMarginRatioD18                     *big.Int
	MinimumInitialMarginRatioD18              *big.Int
	MaintenanceMarginScalarD18                *big.Int
	LiquidationRewardRatioD18                 *big.Int
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MinimumPositionMargin                     *big.Int
}, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getLiquidationParameters", marketId)

	outstruct := new(struct {
		InitialMarginRatioD18                     *big.Int
		MinimumInitialMarginRatioD18              *big.Int
		MaintenanceMarginScalarD18                *big.Int
		LiquidationRewardRatioD18                 *big.Int
		MaxLiquidationLimitAccumulationMultiplier *big.Int
		MaxSecondsInLiquidationWindow             *big.Int
		MinimumPositionMargin                     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialMarginRatioD18 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MinimumInitialMarginRatioD18 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaintenanceMarginScalarD18 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LiquidationRewardRatioD18 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxLiquidationLimitAccumulationMultiplier = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MaxSecondsInLiquidationWindow = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MinimumPositionMargin = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLiquidationParameters is a free data retrieval call binding the contract method 0xf94363a6.
//
// Solidity: function getLiquidationParameters(uint128 marketId) view returns(uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetLiquidationParameters(marketId *big.Int) (struct {
	InitialMarginRatioD18                     *big.Int
	MinimumInitialMarginRatioD18              *big.Int
	MaintenanceMarginScalarD18                *big.Int
	LiquidationRewardRatioD18                 *big.Int
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MinimumPositionMargin                     *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetLiquidationParameters(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetLiquidationParameters is a free data retrieval call binding the contract method 0xf94363a6.
//
// Solidity: function getLiquidationParameters(uint128 marketId) view returns(uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetLiquidationParameters(marketId *big.Int) (struct {
	InitialMarginRatioD18                     *big.Int
	MinimumInitialMarginRatioD18              *big.Int
	MaintenanceMarginScalarD18                *big.Int
	LiquidationRewardRatioD18                 *big.Int
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MinimumPositionMargin                     *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetLiquidationParameters(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetLiquidationRewardGuards is a free data retrieval call binding the contract method 0x0b7f4b2d.
//
// Solidity: function getLiquidationRewardGuards() view returns(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetLiquidationRewardGuards(opts *bind.CallOpts) (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getLiquidationRewardGuards")

	outstruct := new(struct {
		MinLiquidationRewardUsd *big.Int
		MaxLiquidationRewardUsd *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinLiquidationRewardUsd = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxLiquidationRewardUsd = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLiquidationRewardGuards is a free data retrieval call binding the contract method 0x0b7f4b2d.
//
// Solidity: function getLiquidationRewardGuards() view returns(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetLiquidationRewardGuards() (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetLiquidationRewardGuards(&_PrepsMarketGoerli.CallOpts)
}

// GetLiquidationRewardGuards is a free data retrieval call binding the contract method 0x0b7f4b2d.
//
// Solidity: function getLiquidationRewardGuards() view returns(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetLiquidationRewardGuards() (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetLiquidationRewardGuards(&_PrepsMarketGoerli.CallOpts)
}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetLockedOiRatio(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getLockedOiRatio", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetLockedOiRatio(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetLockedOiRatio(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetLockedOiRatio(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetLockedOiRatio(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetMarketSummary(opts *bind.CallOpts, marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getMarketSummary", marketId)

	if err != nil {
		return *new(IPerpsMarketModuleMarketSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(IPerpsMarketModuleMarketSummary)).(*IPerpsMarketModuleMarketSummary)

	return out0, err

}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetMarketSummary(marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	return _PrepsMarketGoerli.Contract.GetMarketSummary(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetMarketSummary(marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	return _PrepsMarketGoerli.Contract.GetMarketSummary(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetMarkets(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getMarkets")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetMarkets() ([]*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetMarkets(&_PrepsMarketGoerli.CallOpts)
}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetMarkets() ([]*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetMarkets(&_PrepsMarketGoerli.CallOpts)
}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetMaxCollateralAmount(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getMaxCollateralAmount", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetMaxCollateralAmount(synthMarketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetMaxCollateralAmount(&_PrepsMarketGoerli.CallOpts, synthMarketId)
}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetMaxCollateralAmount(synthMarketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetMaxCollateralAmount(&_PrepsMarketGoerli.CallOpts, synthMarketId)
}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetMaxMarketSize(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getMaxMarketSize", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetMaxMarketSize(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetMaxMarketSize(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetMaxMarketSize(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetMaxMarketSize(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetOpenPosition is a free data retrieval call binding the contract method 0x22a73967.
//
// Solidity: function getOpenPosition(uint128 accountId, uint128 marketId) view returns(int256 totalPnl, int256 accruedFunding, int128 positionSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetOpenPosition(opts *bind.CallOpts, accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getOpenPosition", accountId, marketId)

	outstruct := new(struct {
		TotalPnl       *big.Int
		AccruedFunding *big.Int
		PositionSize   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalPnl = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AccruedFunding = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PositionSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOpenPosition is a free data retrieval call binding the contract method 0x22a73967.
//
// Solidity: function getOpenPosition(uint128 accountId, uint128 marketId) view returns(int256 totalPnl, int256 accruedFunding, int128 positionSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetOpenPosition(accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetOpenPosition(&_PrepsMarketGoerli.CallOpts, accountId, marketId)
}

// GetOpenPosition is a free data retrieval call binding the contract method 0x22a73967.
//
// Solidity: function getOpenPosition(uint128 accountId, uint128 marketId) view returns(int256 totalPnl, int256 accruedFunding, int128 positionSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetOpenPosition(accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetOpenPosition(&_PrepsMarketGoerli.CallOpts, accountId, marketId)
}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetOrder(opts *bind.CallOpts, accountId *big.Int) (AsyncOrderData, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getOrder", accountId)

	if err != nil {
		return *new(AsyncOrderData), err
	}

	out0 := *abi.ConvertType(out[0], new(AsyncOrderData)).(*AsyncOrderData)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetOrder(accountId *big.Int) (AsyncOrderData, error) {
	return _PrepsMarketGoerli.Contract.GetOrder(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetOrder(accountId *big.Int) (AsyncOrderData, error) {
	return _PrepsMarketGoerli.Contract.GetOrder(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetOrderFees is a free data retrieval call binding the contract method 0xaac23e8c.
//
// Solidity: function getOrderFees(uint128 marketId) view returns(uint256 makerFee, uint256 takerFee)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetOrderFees(opts *bind.CallOpts, marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getOrderFees", marketId)

	outstruct := new(struct {
		MakerFee *big.Int
		TakerFee *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MakerFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TakerFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOrderFees is a free data retrieval call binding the contract method 0xaac23e8c.
//
// Solidity: function getOrderFees(uint128 marketId) view returns(uint256 makerFee, uint256 takerFee)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetOrderFees(marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetOrderFees(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetOrderFees is a free data retrieval call binding the contract method 0xaac23e8c.
//
// Solidity: function getOrderFees(uint128 marketId) view returns(uint256 makerFee, uint256 takerFee)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetOrderFees(marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetOrderFees(&_PrepsMarketGoerli.CallOpts, marketId)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetReferrerShare(opts *bind.CallOpts, referrer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getReferrerShare", referrer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetReferrerShare(referrer common.Address) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetReferrerShare(&_PrepsMarketGoerli.CallOpts, referrer)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetReferrerShare(referrer common.Address) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetReferrerShare(&_PrepsMarketGoerli.CallOpts, referrer)
}

// GetRequiredMargins is a free data retrieval call binding the contract method 0x3c0f0753.
//
// Solidity: function getRequiredMargins(uint128 accountId) view returns(uint256 requiredInitialMargin, uint256 requiredMaintenanceMargin, uint256 totalAccumulatedLiquidationRewards, uint256 maxLiquidationReward)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetRequiredMargins(opts *bind.CallOpts, accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getRequiredMargins", accountId)

	outstruct := new(struct {
		RequiredInitialMargin              *big.Int
		RequiredMaintenanceMargin          *big.Int
		TotalAccumulatedLiquidationRewards *big.Int
		MaxLiquidationReward               *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredInitialMargin = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RequiredMaintenanceMargin = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalAccumulatedLiquidationRewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxLiquidationReward = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRequiredMargins is a free data retrieval call binding the contract method 0x3c0f0753.
//
// Solidity: function getRequiredMargins(uint128 accountId) view returns(uint256 requiredInitialMargin, uint256 requiredMaintenanceMargin, uint256 totalAccumulatedLiquidationRewards, uint256 maxLiquidationReward)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetRequiredMargins(accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetRequiredMargins(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetRequiredMargins is a free data retrieval call binding the contract method 0x3c0f0753.
//
// Solidity: function getRequiredMargins(uint128 accountId) view returns(uint256 requiredInitialMargin, uint256 requiredMaintenanceMargin, uint256 totalAccumulatedLiquidationRewards, uint256 maxLiquidationReward)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetRequiredMargins(accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	return _PrepsMarketGoerli.Contract.GetRequiredMargins(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) settlementStrategy)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetSettlementStrategy(opts *bind.CallOpts, marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getSettlementStrategy", marketId, strategyId)

	if err != nil {
		return *new(SettlementStrategyData), err
	}

	out0 := *abi.ConvertType(out[0], new(SettlementStrategyData)).(*SettlementStrategyData)

	return out0, err

}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) settlementStrategy)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _PrepsMarketGoerli.Contract.GetSettlementStrategy(&_PrepsMarketGoerli.CallOpts, marketId, strategyId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) settlementStrategy)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _PrepsMarketGoerli.Contract.GetSettlementStrategy(&_PrepsMarketGoerli.CallOpts, marketId, strategyId)
}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetSynthDeductionPriority(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getSynthDeductionPriority")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetSynthDeductionPriority() ([]*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetSynthDeductionPriority(&_PrepsMarketGoerli.CallOpts)
}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetSynthDeductionPriority() ([]*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetSynthDeductionPriority(&_PrepsMarketGoerli.CallOpts)
}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) GetWithdrawableMargin(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "getWithdrawableMargin", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GetWithdrawableMargin(accountId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetWithdrawableMargin(&_PrepsMarketGoerli.CallOpts, accountId)
}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) GetWithdrawableMargin(accountId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.GetWithdrawableMargin(&_PrepsMarketGoerli.CallOpts, accountId)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) HasPermission(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "hasPermission", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PrepsMarketGoerli.Contract.HasPermission(&_PrepsMarketGoerli.CallOpts, accountId, permission, user)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PrepsMarketGoerli.Contract.HasPermission(&_PrepsMarketGoerli.CallOpts, accountId, permission, user)
}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) IndexPrice(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "indexPrice", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) IndexPrice(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.IndexPrice(&_PrepsMarketGoerli.CallOpts, marketId)
}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) IndexPrice(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.IndexPrice(&_PrepsMarketGoerli.CallOpts, marketId)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) IsAuthorized(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "isAuthorized", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PrepsMarketGoerli.Contract.IsAuthorized(&_PrepsMarketGoerli.CallOpts, accountId, permission, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PrepsMarketGoerli.Contract.IsAuthorized(&_PrepsMarketGoerli.CallOpts, accountId, permission, user)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) IsFeatureAllowed(opts *bind.CallOpts, feature [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "isFeatureAllowed", feature, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _PrepsMarketGoerli.Contract.IsFeatureAllowed(&_PrepsMarketGoerli.CallOpts, feature, account)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _PrepsMarketGoerli.Contract.IsFeatureAllowed(&_PrepsMarketGoerli.CallOpts, feature, account)
}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) MaxOpenInterest(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "maxOpenInterest", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) MaxOpenInterest(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.MaxOpenInterest(&_PrepsMarketGoerli.CallOpts, marketId)
}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) MaxOpenInterest(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.MaxOpenInterest(&_PrepsMarketGoerli.CallOpts, marketId)
}

// Metadata is a free data retrieval call binding the contract method 0xe3bc36bf.
//
// Solidity: function metadata(uint128 marketId) view returns(string name, string symbol)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) Metadata(opts *bind.CallOpts, marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "metadata", marketId)

	outstruct := new(struct {
		Name   string
		Symbol string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Metadata is a free data retrieval call binding the contract method 0xe3bc36bf.
//
// Solidity: function metadata(uint128 marketId) view returns(string name, string symbol)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) Metadata(marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	return _PrepsMarketGoerli.Contract.Metadata(&_PrepsMarketGoerli.CallOpts, marketId)
}

// Metadata is a free data retrieval call binding the contract method 0xe3bc36bf.
//
// Solidity: function metadata(uint128 marketId) view returns(string name, string symbol)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) Metadata(marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	return _PrepsMarketGoerli.Contract.Metadata(&_PrepsMarketGoerli.CallOpts, marketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) MinimumCredit(opts *bind.CallOpts, perpsMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "minimumCredit", perpsMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) MinimumCredit(perpsMarketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.MinimumCredit(&_PrepsMarketGoerli.CallOpts, perpsMarketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) MinimumCredit(perpsMarketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.MinimumCredit(&_PrepsMarketGoerli.CallOpts, perpsMarketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) Name(opts *bind.CallOpts, perpsMarketId *big.Int) (string, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "name", perpsMarketId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) Name(perpsMarketId *big.Int) (string, error) {
	return _PrepsMarketGoerli.Contract.Name(&_PrepsMarketGoerli.CallOpts, perpsMarketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) Name(perpsMarketId *big.Int) (string, error) {
	return _PrepsMarketGoerli.Contract.Name(&_PrepsMarketGoerli.CallOpts, perpsMarketId)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) NominatedOwner() (common.Address, error) {
	return _PrepsMarketGoerli.Contract.NominatedOwner(&_PrepsMarketGoerli.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) NominatedOwner() (common.Address, error) {
	return _PrepsMarketGoerli.Contract.NominatedOwner(&_PrepsMarketGoerli.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) Owner() (common.Address, error) {
	return _PrepsMarketGoerli.Contract.Owner(&_PrepsMarketGoerli.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) Owner() (common.Address, error) {
	return _PrepsMarketGoerli.Contract.Owner(&_PrepsMarketGoerli.CallOpts)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) ReportedDebt(opts *bind.CallOpts, perpsMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "reportedDebt", perpsMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) ReportedDebt(perpsMarketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.ReportedDebt(&_PrepsMarketGoerli.CallOpts, perpsMarketId)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) ReportedDebt(perpsMarketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.ReportedDebt(&_PrepsMarketGoerli.CallOpts, perpsMarketId)
}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) RequiredMarginForOrder(opts *bind.CallOpts, accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "requiredMarginForOrder", accountId, marketId, sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) RequiredMarginForOrder(accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.RequiredMarginForOrder(&_PrepsMarketGoerli.CallOpts, accountId, marketId, sizeDelta)
}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) RequiredMarginForOrder(accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.RequiredMarginForOrder(&_PrepsMarketGoerli.CallOpts, accountId, marketId, sizeDelta)
}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) Settle(opts *bind.CallOpts, accountId *big.Int) error {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "settle", accountId)

	if err != nil {
		return err
	}

	return err

}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) Settle(accountId *big.Int) error {
	return _PrepsMarketGoerli.Contract.Settle(&_PrepsMarketGoerli.CallOpts, accountId)
}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) Settle(accountId *big.Int) error {
	return _PrepsMarketGoerli.Contract.Settle(&_PrepsMarketGoerli.CallOpts, accountId)
}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) Size(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "size", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) Size(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.Size(&_PrepsMarketGoerli.CallOpts, marketId)
}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) Size(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.Size(&_PrepsMarketGoerli.CallOpts, marketId)
}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) Skew(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "skew", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) Skew(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.Skew(&_PrepsMarketGoerli.CallOpts, marketId)
}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) Skew(marketId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.Skew(&_PrepsMarketGoerli.CallOpts, marketId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PrepsMarketGoerli.Contract.SupportsInterface(&_PrepsMarketGoerli.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PrepsMarketGoerli.Contract.SupportsInterface(&_PrepsMarketGoerli.CallOpts, interfaceId)
}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) TotalAccountOpenInterest(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "totalAccountOpenInterest", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) TotalAccountOpenInterest(accountId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.TotalAccountOpenInterest(&_PrepsMarketGoerli.CallOpts, accountId)
}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) TotalAccountOpenInterest(accountId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.TotalAccountOpenInterest(&_PrepsMarketGoerli.CallOpts, accountId)
}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) TotalCollateralValue(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "totalCollateralValue", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) TotalCollateralValue(accountId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.TotalCollateralValue(&_PrepsMarketGoerli.CallOpts, accountId)
}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) TotalCollateralValue(accountId *big.Int) (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.TotalCollateralValue(&_PrepsMarketGoerli.CallOpts, accountId)
}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_PrepsMarketGoerli *PrepsMarketGoerliCaller) TotalGlobalCollateralValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrepsMarketGoerli.contract.Call(opts, &out, "totalGlobalCollateralValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) TotalGlobalCollateralValue() (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.TotalGlobalCollateralValue(&_PrepsMarketGoerli.CallOpts)
}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_PrepsMarketGoerli *PrepsMarketGoerliCallerSession) TotalGlobalCollateralValue() (*big.Int, error) {
	return _PrepsMarketGoerli.Contract.TotalGlobalCollateralValue(&_PrepsMarketGoerli.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) AcceptOwnership() (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.AcceptOwnership(&_PrepsMarketGoerli.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.AcceptOwnership(&_PrepsMarketGoerli.TransactOpts)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0xda91187c.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) AddSettlementStrategy(opts *bind.TransactOpts, marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "addSettlementStrategy", marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0xda91187c.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.AddSettlementStrategy(&_PrepsMarketGoerli.TransactOpts, marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0xda91187c.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.AddSettlementStrategy(&_PrepsMarketGoerli.TransactOpts, marketId, strategy)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) AddToFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "addToFeatureFlagAllowlist", feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.AddToFeatureFlagAllowlist(&_PrepsMarketGoerli.TransactOpts, feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.AddToFeatureFlagAllowlist(&_PrepsMarketGoerli.TransactOpts, feature, account)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) CommitOrder(opts *bind.TransactOpts, commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "commitOrder", commitment)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) CommitOrder(commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.CommitOrder(&_PrepsMarketGoerli.TransactOpts, commitment)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) CommitOrder(commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.CommitOrder(&_PrepsMarketGoerli.TransactOpts, commitment)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) CreateAccount() (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.CreateAccount(&_PrepsMarketGoerli.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.CreateAccount(&_PrepsMarketGoerli.TransactOpts)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) CreateAccount0(opts *bind.TransactOpts, requestedAccountId *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "createAccount0", requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.CreateAccount0(&_PrepsMarketGoerli.TransactOpts, requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.CreateAccount0(&_PrepsMarketGoerli.TransactOpts, requestedAccountId)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) CreateMarket(opts *bind.TransactOpts, requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "createMarket", requestedMarketId, marketName, marketSymbol)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) CreateMarket(requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.CreateMarket(&_PrepsMarketGoerli.TransactOpts, requestedMarketId, marketName, marketSymbol)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) CreateMarket(requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.CreateMarket(&_PrepsMarketGoerli.TransactOpts, requestedMarketId, marketName, marketSymbol)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) GrantPermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "grantPermission", accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.GrantPermission(&_PrepsMarketGoerli.TransactOpts, accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.GrantPermission(&_PrepsMarketGoerli.TransactOpts, accountId, permission, user)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) InitOrUpgradeNft(opts *bind.TransactOpts, id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "initOrUpgradeNft", id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.InitOrUpgradeNft(&_PrepsMarketGoerli.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.InitOrUpgradeNft(&_PrepsMarketGoerli.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) InitOrUpgradeToken(opts *bind.TransactOpts, id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "initOrUpgradeToken", id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.InitOrUpgradeToken(&_PrepsMarketGoerli.TransactOpts, id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.InitOrUpgradeToken(&_PrepsMarketGoerli.TransactOpts, id, name, symbol, decimals, impl)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) InitializeFactory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "initializeFactory")
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) InitializeFactory() (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.InitializeFactory(&_PrepsMarketGoerli.TransactOpts)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) InitializeFactory() (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.InitializeFactory(&_PrepsMarketGoerli.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) Liquidate(opts *bind.TransactOpts, accountId *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "liquidate", accountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) Liquidate(accountId *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.Liquidate(&_PrepsMarketGoerli.TransactOpts, accountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) Liquidate(accountId *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.Liquidate(&_PrepsMarketGoerli.TransactOpts, accountId)
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) LiquidateFlagged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "liquidateFlagged")
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) LiquidateFlagged() (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.LiquidateFlagged(&_PrepsMarketGoerli.TransactOpts)
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) LiquidateFlagged() (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.LiquidateFlagged(&_PrepsMarketGoerli.TransactOpts)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) ModifyCollateral(opts *bind.TransactOpts, accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "modifyCollateral", accountId, synthMarketId, amountDelta)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) ModifyCollateral(accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.ModifyCollateral(&_PrepsMarketGoerli.TransactOpts, accountId, synthMarketId, amountDelta)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) ModifyCollateral(accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.ModifyCollateral(&_PrepsMarketGoerli.TransactOpts, accountId, synthMarketId, amountDelta)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) NominateNewOwner(opts *bind.TransactOpts, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "nominateNewOwner", newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.NominateNewOwner(&_PrepsMarketGoerli.TransactOpts, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.NominateNewOwner(&_PrepsMarketGoerli.TransactOpts, newNominatedOwner)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) NotifyAccountTransfer(opts *bind.TransactOpts, to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "notifyAccountTransfer", to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.NotifyAccountTransfer(&_PrepsMarketGoerli.TransactOpts, to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.NotifyAccountTransfer(&_PrepsMarketGoerli.TransactOpts, to, accountId)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) RegisterUnmanagedSystem(opts *bind.TransactOpts, id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "registerUnmanagedSystem", id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.RegisterUnmanagedSystem(&_PrepsMarketGoerli.TransactOpts, id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.RegisterUnmanagedSystem(&_PrepsMarketGoerli.TransactOpts, id, endpoint)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) RemoveFromFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "removeFromFeatureFlagAllowlist", feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.RemoveFromFeatureFlagAllowlist(&_PrepsMarketGoerli.TransactOpts, feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.RemoveFromFeatureFlagAllowlist(&_PrepsMarketGoerli.TransactOpts, feature, account)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) RenounceNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "renounceNomination")
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) RenounceNomination() (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.RenounceNomination(&_PrepsMarketGoerli.TransactOpts)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) RenounceNomination() (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.RenounceNomination(&_PrepsMarketGoerli.TransactOpts)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) RenouncePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "renouncePermission", accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.RenouncePermission(&_PrepsMarketGoerli.TransactOpts, accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.RenouncePermission(&_PrepsMarketGoerli.TransactOpts, accountId, permission)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) RevokePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "revokePermission", accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.RevokePermission(&_PrepsMarketGoerli.TransactOpts, accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.RevokePermission(&_PrepsMarketGoerli.TransactOpts, accountId, permission, user)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetDeniers(opts *bind.TransactOpts, feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setDeniers", feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetDeniers(&_PrepsMarketGoerli.TransactOpts, feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetDeniers(&_PrepsMarketGoerli.TransactOpts, feature, deniers)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetFeatureFlagAllowAll(opts *bind.TransactOpts, feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setFeatureFlagAllowAll", feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetFeatureFlagAllowAll(&_PrepsMarketGoerli.TransactOpts, feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetFeatureFlagAllowAll(&_PrepsMarketGoerli.TransactOpts, feature, allowAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetFeatureFlagDenyAll(opts *bind.TransactOpts, feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setFeatureFlagDenyAll", feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetFeatureFlagDenyAll(&_PrepsMarketGoerli.TransactOpts, feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetFeatureFlagDenyAll(&_PrepsMarketGoerli.TransactOpts, feature, denyAll)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetFeeCollector(opts *bind.TransactOpts, feeCollector common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setFeeCollector", feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetFeeCollector(&_PrepsMarketGoerli.TransactOpts, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetFeeCollector(&_PrepsMarketGoerli.TransactOpts, feeCollector)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetFundingParameters(opts *bind.TransactOpts, marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setFundingParameters", marketId, skewScale, maxFundingVelocity)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetFundingParameters(marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetFundingParameters(&_PrepsMarketGoerli.TransactOpts, marketId, skewScale, maxFundingVelocity)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetFundingParameters(marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetFundingParameters(&_PrepsMarketGoerli.TransactOpts, marketId, skewScale, maxFundingVelocity)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0xcb450077.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetLiquidationParameters(opts *bind.TransactOpts, marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setLiquidationParameters", marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, minimumPositionMargin)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0xcb450077.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetLiquidationParameters(marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetLiquidationParameters(&_PrepsMarketGoerli.TransactOpts, marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, minimumPositionMargin)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0xcb450077.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetLiquidationParameters(marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetLiquidationParameters(&_PrepsMarketGoerli.TransactOpts, marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, minimumPositionMargin)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetLiquidationRewardGuards(opts *bind.TransactOpts, minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setLiquidationRewardGuards", minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetLiquidationRewardGuards(minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetLiquidationRewardGuards(&_PrepsMarketGoerli.TransactOpts, minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetLiquidationRewardGuards(minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetLiquidationRewardGuards(&_PrepsMarketGoerli.TransactOpts, minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetLockedOiRatio(opts *bind.TransactOpts, marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setLockedOiRatio", marketId, lockedOiRatioD18)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetLockedOiRatio(marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetLockedOiRatio(&_PrepsMarketGoerli.TransactOpts, marketId, lockedOiRatioD18)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetLockedOiRatio(marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetLockedOiRatio(&_PrepsMarketGoerli.TransactOpts, marketId, lockedOiRatioD18)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetMaxCollateralAmount(opts *bind.TransactOpts, synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setMaxCollateralAmount", synthMarketId, collateralAmount)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetMaxCollateralAmount(synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetMaxCollateralAmount(&_PrepsMarketGoerli.TransactOpts, synthMarketId, collateralAmount)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetMaxCollateralAmount(synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetMaxCollateralAmount(&_PrepsMarketGoerli.TransactOpts, synthMarketId, collateralAmount)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetMaxMarketSize(opts *bind.TransactOpts, marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setMaxMarketSize", marketId, maxMarketSize)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetMaxMarketSize(marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetMaxMarketSize(&_PrepsMarketGoerli.TransactOpts, marketId, maxMarketSize)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetMaxMarketSize(marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetMaxMarketSize(&_PrepsMarketGoerli.TransactOpts, marketId, maxMarketSize)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetOrderFees(opts *bind.TransactOpts, marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setOrderFees", marketId, makerFeeRatio, takerFeeRatio)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetOrderFees(marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetOrderFees(&_PrepsMarketGoerli.TransactOpts, marketId, makerFeeRatio, takerFeeRatio)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetOrderFees(marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetOrderFees(&_PrepsMarketGoerli.TransactOpts, marketId, makerFeeRatio, takerFeeRatio)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetSettlementStrategyEnabled(opts *bind.TransactOpts, marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setSettlementStrategyEnabled", marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetSettlementStrategyEnabled(&_PrepsMarketGoerli.TransactOpts, marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetSettlementStrategyEnabled(&_PrepsMarketGoerli.TransactOpts, marketId, strategyId, enabled)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetSpotMarket(opts *bind.TransactOpts, spotMarket common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setSpotMarket", spotMarket)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetSpotMarket(spotMarket common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetSpotMarket(&_PrepsMarketGoerli.TransactOpts, spotMarket)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetSpotMarket(spotMarket common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetSpotMarket(&_PrepsMarketGoerli.TransactOpts, spotMarket)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetSynthDeductionPriority(opts *bind.TransactOpts, newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setSynthDeductionPriority", newSynthDeductionPriority)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetSynthDeductionPriority(newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetSynthDeductionPriority(&_PrepsMarketGoerli.TransactOpts, newSynthDeductionPriority)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetSynthDeductionPriority(newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetSynthDeductionPriority(&_PrepsMarketGoerli.TransactOpts, newSynthDeductionPriority)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SetSynthetix(opts *bind.TransactOpts, synthetix common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "setSynthetix", synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetSynthetix(&_PrepsMarketGoerli.TransactOpts, synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SetSynthetix(&_PrepsMarketGoerli.TransactOpts, synthetix)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SettlePythOrder(opts *bind.TransactOpts, result []byte, extraData []byte) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "settlePythOrder", result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SettlePythOrder(&_PrepsMarketGoerli.TransactOpts, result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SettlePythOrder(&_PrepsMarketGoerli.TransactOpts, result, extraData)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) SimulateUpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "simulateUpgradeTo", newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SimulateUpgradeTo(&_PrepsMarketGoerli.TransactOpts, newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.SimulateUpgradeTo(&_PrepsMarketGoerli.TransactOpts, newImplementation)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) UpdatePriceData(opts *bind.TransactOpts, perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "updatePriceData", perpsMarketId, feedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) UpdatePriceData(perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.UpdatePriceData(&_PrepsMarketGoerli.TransactOpts, perpsMarketId, feedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) UpdatePriceData(perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.UpdatePriceData(&_PrepsMarketGoerli.TransactOpts, perpsMarketId, feedId)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) UpdateReferrerShare(opts *bind.TransactOpts, referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "updateReferrerShare", referrer, shareRatioD18)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) UpdateReferrerShare(referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.UpdateReferrerShare(&_PrepsMarketGoerli.TransactOpts, referrer, shareRatioD18)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) UpdateReferrerShare(referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.UpdateReferrerShare(&_PrepsMarketGoerli.TransactOpts, referrer, shareRatioD18)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.UpgradeTo(&_PrepsMarketGoerli.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PrepsMarketGoerli *PrepsMarketGoerliTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PrepsMarketGoerli.Contract.UpgradeTo(&_PrepsMarketGoerli.TransactOpts, newImplementation)
}

// PrepsMarketGoerliAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliAccountCreatedIterator struct {
	Event *PrepsMarketGoerliAccountCreated // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliAccountCreated)
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
		it.Event = new(PrepsMarketGoerliAccountCreated)
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
func (it *PrepsMarketGoerliAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliAccountCreated represents a AccountCreated event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliAccountCreated struct {
	AccountId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterAccountCreated(opts *bind.FilterOpts, accountId []*big.Int, owner []common.Address) (*PrepsMarketGoerliAccountCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliAccountCreatedIterator{contract: _PrepsMarketGoerli.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliAccountCreated, accountId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliAccountCreated)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseAccountCreated(log types.Log) (*PrepsMarketGoerliAccountCreated, error) {
	event := new(PrepsMarketGoerliAccountCreated)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliAccountLiquidatedIterator is returned from FilterAccountLiquidated and is used to iterate over the raw logs and unpacked data for AccountLiquidated events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliAccountLiquidatedIterator struct {
	Event *PrepsMarketGoerliAccountLiquidated // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliAccountLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliAccountLiquidated)
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
		it.Event = new(PrepsMarketGoerliAccountLiquidated)
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
func (it *PrepsMarketGoerliAccountLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliAccountLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliAccountLiquidated represents a AccountLiquidated event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliAccountLiquidated struct {
	AccountId       *big.Int
	Reward          *big.Int
	FullLiquidation bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccountLiquidated is a free log retrieval operation binding the contract event 0x926f56b0a80cc150ad77b59cc60bc53cb907eab441d204924ec67b40ae64f5f5.
//
// Solidity: event AccountLiquidated(uint128 indexed accountId, uint256 reward, bool fullLiquidation)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterAccountLiquidated(opts *bind.FilterOpts, accountId []*big.Int) (*PrepsMarketGoerliAccountLiquidatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "AccountLiquidated", accountIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliAccountLiquidatedIterator{contract: _PrepsMarketGoerli.contract, event: "AccountLiquidated", logs: logs, sub: sub}, nil
}

// WatchAccountLiquidated is a free log subscription operation binding the contract event 0x926f56b0a80cc150ad77b59cc60bc53cb907eab441d204924ec67b40ae64f5f5.
//
// Solidity: event AccountLiquidated(uint128 indexed accountId, uint256 reward, bool fullLiquidation)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchAccountLiquidated(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliAccountLiquidated, accountId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "AccountLiquidated", accountIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliAccountLiquidated)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "AccountLiquidated", log); err != nil {
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

// ParseAccountLiquidated is a log parse operation binding the contract event 0x926f56b0a80cc150ad77b59cc60bc53cb907eab441d204924ec67b40ae64f5f5.
//
// Solidity: event AccountLiquidated(uint128 indexed accountId, uint256 reward, bool fullLiquidation)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseAccountLiquidated(log types.Log) (*PrepsMarketGoerliAccountLiquidated, error) {
	event := new(PrepsMarketGoerliAccountLiquidated)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "AccountLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliAssociatedSystemSetIterator is returned from FilterAssociatedSystemSet and is used to iterate over the raw logs and unpacked data for AssociatedSystemSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliAssociatedSystemSetIterator struct {
	Event *PrepsMarketGoerliAssociatedSystemSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliAssociatedSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliAssociatedSystemSet)
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
		it.Event = new(PrepsMarketGoerliAssociatedSystemSet)
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
func (it *PrepsMarketGoerliAssociatedSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliAssociatedSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliAssociatedSystemSet represents a AssociatedSystemSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliAssociatedSystemSet struct {
	Kind  [32]byte
	Id    [32]byte
	Proxy common.Address
	Impl  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssociatedSystemSet is a free log retrieval operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterAssociatedSystemSet(opts *bind.FilterOpts, kind [][32]byte, id [][32]byte) (*PrepsMarketGoerliAssociatedSystemSetIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliAssociatedSystemSetIterator{contract: _PrepsMarketGoerli.contract, event: "AssociatedSystemSet", logs: logs, sub: sub}, nil
}

// WatchAssociatedSystemSet is a free log subscription operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchAssociatedSystemSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliAssociatedSystemSet, kind [][32]byte, id [][32]byte) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliAssociatedSystemSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
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
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseAssociatedSystemSet(log types.Log) (*PrepsMarketGoerliAssociatedSystemSet, error) {
	event := new(PrepsMarketGoerliAssociatedSystemSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliCollateralModifiedIterator is returned from FilterCollateralModified and is used to iterate over the raw logs and unpacked data for CollateralModified events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliCollateralModifiedIterator struct {
	Event *PrepsMarketGoerliCollateralModified // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliCollateralModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliCollateralModified)
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
		it.Event = new(PrepsMarketGoerliCollateralModified)
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
func (it *PrepsMarketGoerliCollateralModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliCollateralModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliCollateralModified represents a CollateralModified event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliCollateralModified struct {
	AccountId     *big.Int
	SynthMarketId *big.Int
	AmountDelta   *big.Int
	Sender        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralModified is a free log retrieval operation binding the contract event 0x2e8360c2f3a6fc1a15aefdd0a0922bea3c898cb90d38c3a97354e7c013116064.
//
// Solidity: event CollateralModified(uint128 indexed accountId, uint128 indexed synthMarketId, int256 amountDelta, address indexed sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterCollateralModified(opts *bind.FilterOpts, accountId []*big.Int, synthMarketId []*big.Int, sender []common.Address) (*PrepsMarketGoerliCollateralModifiedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "CollateralModified", accountIdRule, synthMarketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliCollateralModifiedIterator{contract: _PrepsMarketGoerli.contract, event: "CollateralModified", logs: logs, sub: sub}, nil
}

// WatchCollateralModified is a free log subscription operation binding the contract event 0x2e8360c2f3a6fc1a15aefdd0a0922bea3c898cb90d38c3a97354e7c013116064.
//
// Solidity: event CollateralModified(uint128 indexed accountId, uint128 indexed synthMarketId, int256 amountDelta, address indexed sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchCollateralModified(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliCollateralModified, accountId []*big.Int, synthMarketId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "CollateralModified", accountIdRule, synthMarketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliCollateralModified)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "CollateralModified", log); err != nil {
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

// ParseCollateralModified is a log parse operation binding the contract event 0x2e8360c2f3a6fc1a15aefdd0a0922bea3c898cb90d38c3a97354e7c013116064.
//
// Solidity: event CollateralModified(uint128 indexed accountId, uint128 indexed synthMarketId, int256 amountDelta, address indexed sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseCollateralModified(log types.Log) (*PrepsMarketGoerliCollateralModified, error) {
	event := new(PrepsMarketGoerliCollateralModified)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "CollateralModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliFactoryInitializedIterator is returned from FilterFactoryInitialized and is used to iterate over the raw logs and unpacked data for FactoryInitialized events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFactoryInitializedIterator struct {
	Event *PrepsMarketGoerliFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliFactoryInitialized)
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
		it.Event = new(PrepsMarketGoerliFactoryInitialized)
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
func (it *PrepsMarketGoerliFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliFactoryInitialized represents a FactoryInitialized event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFactoryInitialized struct {
	GlobalPerpsMarketId *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFactoryInitialized is a free log retrieval operation binding the contract event 0xb3240229b07e26f2f02e69dda85ede86e162ccbc6d10e6aade28931e7f533339.
//
// Solidity: event FactoryInitialized(uint128 globalPerpsMarketId)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterFactoryInitialized(opts *bind.FilterOpts) (*PrepsMarketGoerliFactoryInitializedIterator, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "FactoryInitialized")
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliFactoryInitializedIterator{contract: _PrepsMarketGoerli.contract, event: "FactoryInitialized", logs: logs, sub: sub}, nil
}

// WatchFactoryInitialized is a free log subscription operation binding the contract event 0xb3240229b07e26f2f02e69dda85ede86e162ccbc6d10e6aade28931e7f533339.
//
// Solidity: event FactoryInitialized(uint128 globalPerpsMarketId)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchFactoryInitialized(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "FactoryInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliFactoryInitialized)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FactoryInitialized", log); err != nil {
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

// ParseFactoryInitialized is a log parse operation binding the contract event 0xb3240229b07e26f2f02e69dda85ede86e162ccbc6d10e6aade28931e7f533339.
//
// Solidity: event FactoryInitialized(uint128 globalPerpsMarketId)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseFactoryInitialized(log types.Log) (*PrepsMarketGoerliFactoryInitialized, error) {
	event := new(PrepsMarketGoerliFactoryInitialized)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FactoryInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliFeatureFlagAllowAllSetIterator is returned from FilterFeatureFlagAllowAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowAllSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeatureFlagAllowAllSetIterator struct {
	Event *PrepsMarketGoerliFeatureFlagAllowAllSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliFeatureFlagAllowAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliFeatureFlagAllowAllSet)
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
		it.Event = new(PrepsMarketGoerliFeatureFlagAllowAllSet)
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
func (it *PrepsMarketGoerliFeatureFlagAllowAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliFeatureFlagAllowAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliFeatureFlagAllowAllSet represents a FeatureFlagAllowAllSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeatureFlagAllowAllSet struct {
	Feature  [32]byte
	AllowAll bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowAllSet is a free log retrieval operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterFeatureFlagAllowAllSet(opts *bind.FilterOpts, feature [][32]byte) (*PrepsMarketGoerliFeatureFlagAllowAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliFeatureFlagAllowAllSetIterator{contract: _PrepsMarketGoerli.contract, event: "FeatureFlagAllowAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowAllSet is a free log subscription operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchFeatureFlagAllowAllSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliFeatureFlagAllowAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliFeatureFlagAllowAllSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
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
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseFeatureFlagAllowAllSet(log types.Log) (*PrepsMarketGoerliFeatureFlagAllowAllSet, error) {
	event := new(PrepsMarketGoerliFeatureFlagAllowAllSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliFeatureFlagAllowlistAddedIterator is returned from FilterFeatureFlagAllowlistAdded and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistAdded events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeatureFlagAllowlistAddedIterator struct {
	Event *PrepsMarketGoerliFeatureFlagAllowlistAdded // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliFeatureFlagAllowlistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliFeatureFlagAllowlistAdded)
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
		it.Event = new(PrepsMarketGoerliFeatureFlagAllowlistAdded)
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
func (it *PrepsMarketGoerliFeatureFlagAllowlistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliFeatureFlagAllowlistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliFeatureFlagAllowlistAdded represents a FeatureFlagAllowlistAdded event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeatureFlagAllowlistAdded struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistAdded is a free log retrieval operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterFeatureFlagAllowlistAdded(opts *bind.FilterOpts, feature [][32]byte) (*PrepsMarketGoerliFeatureFlagAllowlistAddedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliFeatureFlagAllowlistAddedIterator{contract: _PrepsMarketGoerli.contract, event: "FeatureFlagAllowlistAdded", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistAdded is a free log subscription operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchFeatureFlagAllowlistAdded(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliFeatureFlagAllowlistAdded, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliFeatureFlagAllowlistAdded)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
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
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseFeatureFlagAllowlistAdded(log types.Log) (*PrepsMarketGoerliFeatureFlagAllowlistAdded, error) {
	event := new(PrepsMarketGoerliFeatureFlagAllowlistAdded)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliFeatureFlagAllowlistRemovedIterator is returned from FilterFeatureFlagAllowlistRemoved and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistRemoved events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeatureFlagAllowlistRemovedIterator struct {
	Event *PrepsMarketGoerliFeatureFlagAllowlistRemoved // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliFeatureFlagAllowlistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliFeatureFlagAllowlistRemoved)
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
		it.Event = new(PrepsMarketGoerliFeatureFlagAllowlistRemoved)
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
func (it *PrepsMarketGoerliFeatureFlagAllowlistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliFeatureFlagAllowlistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliFeatureFlagAllowlistRemoved represents a FeatureFlagAllowlistRemoved event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeatureFlagAllowlistRemoved struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistRemoved is a free log retrieval operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterFeatureFlagAllowlistRemoved(opts *bind.FilterOpts, feature [][32]byte) (*PrepsMarketGoerliFeatureFlagAllowlistRemovedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliFeatureFlagAllowlistRemovedIterator{contract: _PrepsMarketGoerli.contract, event: "FeatureFlagAllowlistRemoved", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistRemoved is a free log subscription operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchFeatureFlagAllowlistRemoved(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliFeatureFlagAllowlistRemoved, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliFeatureFlagAllowlistRemoved)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
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
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseFeatureFlagAllowlistRemoved(log types.Log) (*PrepsMarketGoerliFeatureFlagAllowlistRemoved, error) {
	event := new(PrepsMarketGoerliFeatureFlagAllowlistRemoved)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliFeatureFlagDeniersResetIterator is returned from FilterFeatureFlagDeniersReset and is used to iterate over the raw logs and unpacked data for FeatureFlagDeniersReset events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeatureFlagDeniersResetIterator struct {
	Event *PrepsMarketGoerliFeatureFlagDeniersReset // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliFeatureFlagDeniersResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliFeatureFlagDeniersReset)
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
		it.Event = new(PrepsMarketGoerliFeatureFlagDeniersReset)
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
func (it *PrepsMarketGoerliFeatureFlagDeniersResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliFeatureFlagDeniersResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliFeatureFlagDeniersReset represents a FeatureFlagDeniersReset event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeatureFlagDeniersReset struct {
	Feature [32]byte
	Deniers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDeniersReset is a free log retrieval operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterFeatureFlagDeniersReset(opts *bind.FilterOpts, feature [][32]byte) (*PrepsMarketGoerliFeatureFlagDeniersResetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliFeatureFlagDeniersResetIterator{contract: _PrepsMarketGoerli.contract, event: "FeatureFlagDeniersReset", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDeniersReset is a free log subscription operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchFeatureFlagDeniersReset(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliFeatureFlagDeniersReset, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliFeatureFlagDeniersReset)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
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
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseFeatureFlagDeniersReset(log types.Log) (*PrepsMarketGoerliFeatureFlagDeniersReset, error) {
	event := new(PrepsMarketGoerliFeatureFlagDeniersReset)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliFeatureFlagDenyAllSetIterator is returned from FilterFeatureFlagDenyAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagDenyAllSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeatureFlagDenyAllSetIterator struct {
	Event *PrepsMarketGoerliFeatureFlagDenyAllSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliFeatureFlagDenyAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliFeatureFlagDenyAllSet)
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
		it.Event = new(PrepsMarketGoerliFeatureFlagDenyAllSet)
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
func (it *PrepsMarketGoerliFeatureFlagDenyAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliFeatureFlagDenyAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliFeatureFlagDenyAllSet represents a FeatureFlagDenyAllSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeatureFlagDenyAllSet struct {
	Feature [32]byte
	DenyAll bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDenyAllSet is a free log retrieval operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterFeatureFlagDenyAllSet(opts *bind.FilterOpts, feature [][32]byte) (*PrepsMarketGoerliFeatureFlagDenyAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliFeatureFlagDenyAllSetIterator{contract: _PrepsMarketGoerli.contract, event: "FeatureFlagDenyAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDenyAllSet is a free log subscription operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchFeatureFlagDenyAllSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliFeatureFlagDenyAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliFeatureFlagDenyAllSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
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
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseFeatureFlagDenyAllSet(log types.Log) (*PrepsMarketGoerliFeatureFlagDenyAllSet, error) {
	event := new(PrepsMarketGoerliFeatureFlagDenyAllSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliFeeCollectorSetIterator is returned from FilterFeeCollectorSet and is used to iterate over the raw logs and unpacked data for FeeCollectorSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeeCollectorSetIterator struct {
	Event *PrepsMarketGoerliFeeCollectorSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliFeeCollectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliFeeCollectorSet)
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
		it.Event = new(PrepsMarketGoerliFeeCollectorSet)
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
func (it *PrepsMarketGoerliFeeCollectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliFeeCollectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliFeeCollectorSet represents a FeeCollectorSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFeeCollectorSet struct {
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorSet is a free log retrieval operation binding the contract event 0x12e1d17016b94668449f97876f4a8d5cc2c19f314db337418894734037cc19d4.
//
// Solidity: event FeeCollectorSet(address feeCollector)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterFeeCollectorSet(opts *bind.FilterOpts) (*PrepsMarketGoerliFeeCollectorSetIterator, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "FeeCollectorSet")
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliFeeCollectorSetIterator{contract: _PrepsMarketGoerli.contract, event: "FeeCollectorSet", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorSet is a free log subscription operation binding the contract event 0x12e1d17016b94668449f97876f4a8d5cc2c19f314db337418894734037cc19d4.
//
// Solidity: event FeeCollectorSet(address feeCollector)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchFeeCollectorSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliFeeCollectorSet) (event.Subscription, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "FeeCollectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliFeeCollectorSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
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

// ParseFeeCollectorSet is a log parse operation binding the contract event 0x12e1d17016b94668449f97876f4a8d5cc2c19f314db337418894734037cc19d4.
//
// Solidity: event FeeCollectorSet(address feeCollector)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseFeeCollectorSet(log types.Log) (*PrepsMarketGoerliFeeCollectorSet, error) {
	event := new(PrepsMarketGoerliFeeCollectorSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliFundingParametersSetIterator is returned from FilterFundingParametersSet and is used to iterate over the raw logs and unpacked data for FundingParametersSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFundingParametersSetIterator struct {
	Event *PrepsMarketGoerliFundingParametersSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliFundingParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliFundingParametersSet)
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
		it.Event = new(PrepsMarketGoerliFundingParametersSet)
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
func (it *PrepsMarketGoerliFundingParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliFundingParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliFundingParametersSet represents a FundingParametersSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliFundingParametersSet struct {
	MarketId           *big.Int
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFundingParametersSet is a free log retrieval operation binding the contract event 0xa74afd926bbafbb9252d224a1fcd6a209f851324bd485f556786820a76e31b65.
//
// Solidity: event FundingParametersSet(uint128 indexed marketId, uint256 skewScale, uint256 maxFundingVelocity)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterFundingParametersSet(opts *bind.FilterOpts, marketId []*big.Int) (*PrepsMarketGoerliFundingParametersSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "FundingParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliFundingParametersSetIterator{contract: _PrepsMarketGoerli.contract, event: "FundingParametersSet", logs: logs, sub: sub}, nil
}

// WatchFundingParametersSet is a free log subscription operation binding the contract event 0xa74afd926bbafbb9252d224a1fcd6a209f851324bd485f556786820a76e31b65.
//
// Solidity: event FundingParametersSet(uint128 indexed marketId, uint256 skewScale, uint256 maxFundingVelocity)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchFundingParametersSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliFundingParametersSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "FundingParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliFundingParametersSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FundingParametersSet", log); err != nil {
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

// ParseFundingParametersSet is a log parse operation binding the contract event 0xa74afd926bbafbb9252d224a1fcd6a209f851324bd485f556786820a76e31b65.
//
// Solidity: event FundingParametersSet(uint128 indexed marketId, uint256 skewScale, uint256 maxFundingVelocity)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseFundingParametersSet(log types.Log) (*PrepsMarketGoerliFundingParametersSet, error) {
	event := new(PrepsMarketGoerliFundingParametersSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "FundingParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliLiquidationParametersSetIterator is returned from FilterLiquidationParametersSet and is used to iterate over the raw logs and unpacked data for LiquidationParametersSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliLiquidationParametersSetIterator struct {
	Event *PrepsMarketGoerliLiquidationParametersSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliLiquidationParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliLiquidationParametersSet)
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
		it.Event = new(PrepsMarketGoerliLiquidationParametersSet)
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
func (it *PrepsMarketGoerliLiquidationParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliLiquidationParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliLiquidationParametersSet represents a LiquidationParametersSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliLiquidationParametersSet struct {
	MarketId                                  *big.Int
	InitialMarginRatioD18                     *big.Int
	MaintenanceMarginRatioD18                 *big.Int
	MinimumInitialMarginRatioD18              *big.Int
	LiquidationRewardRatioD18                 *big.Int
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MinimumPositionMargin                     *big.Int
	Raw                                       types.Log // Blockchain specific contextual infos
}

// FilterLiquidationParametersSet is a free log retrieval operation binding the contract event 0x362c38617bd453ec04ea677dc8dba9aa5b977195f14ddcf5cef0509bb3fd9200.
//
// Solidity: event LiquidationParametersSet(uint128 indexed marketId, uint256 initialMarginRatioD18, uint256 maintenanceMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterLiquidationParametersSet(opts *bind.FilterOpts, marketId []*big.Int) (*PrepsMarketGoerliLiquidationParametersSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "LiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliLiquidationParametersSetIterator{contract: _PrepsMarketGoerli.contract, event: "LiquidationParametersSet", logs: logs, sub: sub}, nil
}

// WatchLiquidationParametersSet is a free log subscription operation binding the contract event 0x362c38617bd453ec04ea677dc8dba9aa5b977195f14ddcf5cef0509bb3fd9200.
//
// Solidity: event LiquidationParametersSet(uint128 indexed marketId, uint256 initialMarginRatioD18, uint256 maintenanceMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchLiquidationParametersSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliLiquidationParametersSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "LiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliLiquidationParametersSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "LiquidationParametersSet", log); err != nil {
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

// ParseLiquidationParametersSet is a log parse operation binding the contract event 0x362c38617bd453ec04ea677dc8dba9aa5b977195f14ddcf5cef0509bb3fd9200.
//
// Solidity: event LiquidationParametersSet(uint128 indexed marketId, uint256 initialMarginRatioD18, uint256 maintenanceMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseLiquidationParametersSet(log types.Log) (*PrepsMarketGoerliLiquidationParametersSet, error) {
	event := new(PrepsMarketGoerliLiquidationParametersSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "LiquidationParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliLiquidationRewardGuardsSetIterator is returned from FilterLiquidationRewardGuardsSet and is used to iterate over the raw logs and unpacked data for LiquidationRewardGuardsSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliLiquidationRewardGuardsSetIterator struct {
	Event *PrepsMarketGoerliLiquidationRewardGuardsSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliLiquidationRewardGuardsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliLiquidationRewardGuardsSet)
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
		it.Event = new(PrepsMarketGoerliLiquidationRewardGuardsSet)
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
func (it *PrepsMarketGoerliLiquidationRewardGuardsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliLiquidationRewardGuardsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliLiquidationRewardGuardsSet represents a LiquidationRewardGuardsSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliLiquidationRewardGuardsSet struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterLiquidationRewardGuardsSet is a free log retrieval operation binding the contract event 0xbb7f8eb17b00549c0f234a3c363ef3eb39329f98b52048416480cf62d0f0e1f5.
//
// Solidity: event LiquidationRewardGuardsSet(uint256 indexed minLiquidationRewardUsd, uint256 indexed maxLiquidationRewardUsd)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterLiquidationRewardGuardsSet(opts *bind.FilterOpts, minLiquidationRewardUsd []*big.Int, maxLiquidationRewardUsd []*big.Int) (*PrepsMarketGoerliLiquidationRewardGuardsSetIterator, error) {

	var minLiquidationRewardUsdRule []interface{}
	for _, minLiquidationRewardUsdItem := range minLiquidationRewardUsd {
		minLiquidationRewardUsdRule = append(minLiquidationRewardUsdRule, minLiquidationRewardUsdItem)
	}
	var maxLiquidationRewardUsdRule []interface{}
	for _, maxLiquidationRewardUsdItem := range maxLiquidationRewardUsd {
		maxLiquidationRewardUsdRule = append(maxLiquidationRewardUsdRule, maxLiquidationRewardUsdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "LiquidationRewardGuardsSet", minLiquidationRewardUsdRule, maxLiquidationRewardUsdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliLiquidationRewardGuardsSetIterator{contract: _PrepsMarketGoerli.contract, event: "LiquidationRewardGuardsSet", logs: logs, sub: sub}, nil
}

// WatchLiquidationRewardGuardsSet is a free log subscription operation binding the contract event 0xbb7f8eb17b00549c0f234a3c363ef3eb39329f98b52048416480cf62d0f0e1f5.
//
// Solidity: event LiquidationRewardGuardsSet(uint256 indexed minLiquidationRewardUsd, uint256 indexed maxLiquidationRewardUsd)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchLiquidationRewardGuardsSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliLiquidationRewardGuardsSet, minLiquidationRewardUsd []*big.Int, maxLiquidationRewardUsd []*big.Int) (event.Subscription, error) {

	var minLiquidationRewardUsdRule []interface{}
	for _, minLiquidationRewardUsdItem := range minLiquidationRewardUsd {
		minLiquidationRewardUsdRule = append(minLiquidationRewardUsdRule, minLiquidationRewardUsdItem)
	}
	var maxLiquidationRewardUsdRule []interface{}
	for _, maxLiquidationRewardUsdItem := range maxLiquidationRewardUsd {
		maxLiquidationRewardUsdRule = append(maxLiquidationRewardUsdRule, maxLiquidationRewardUsdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "LiquidationRewardGuardsSet", minLiquidationRewardUsdRule, maxLiquidationRewardUsdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliLiquidationRewardGuardsSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "LiquidationRewardGuardsSet", log); err != nil {
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

// ParseLiquidationRewardGuardsSet is a log parse operation binding the contract event 0xbb7f8eb17b00549c0f234a3c363ef3eb39329f98b52048416480cf62d0f0e1f5.
//
// Solidity: event LiquidationRewardGuardsSet(uint256 indexed minLiquidationRewardUsd, uint256 indexed maxLiquidationRewardUsd)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseLiquidationRewardGuardsSet(log types.Log) (*PrepsMarketGoerliLiquidationRewardGuardsSet, error) {
	event := new(PrepsMarketGoerliLiquidationRewardGuardsSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "LiquidationRewardGuardsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliLockedOiRatioSetIterator is returned from FilterLockedOiRatioSet and is used to iterate over the raw logs and unpacked data for LockedOiRatioSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliLockedOiRatioSetIterator struct {
	Event *PrepsMarketGoerliLockedOiRatioSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliLockedOiRatioSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliLockedOiRatioSet)
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
		it.Event = new(PrepsMarketGoerliLockedOiRatioSet)
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
func (it *PrepsMarketGoerliLockedOiRatioSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliLockedOiRatioSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliLockedOiRatioSet represents a LockedOiRatioSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliLockedOiRatioSet struct {
	MarketId         *big.Int
	LockedOiRatioD18 *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLockedOiRatioSet is a free log retrieval operation binding the contract event 0x1d841fd5b4c806bc5a073d637a8506e1e74d16cb18251b711cb47e133ceafc2d.
//
// Solidity: event LockedOiRatioSet(uint128 indexed marketId, uint256 lockedOiRatioD18)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterLockedOiRatioSet(opts *bind.FilterOpts, marketId []*big.Int) (*PrepsMarketGoerliLockedOiRatioSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "LockedOiRatioSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliLockedOiRatioSetIterator{contract: _PrepsMarketGoerli.contract, event: "LockedOiRatioSet", logs: logs, sub: sub}, nil
}

// WatchLockedOiRatioSet is a free log subscription operation binding the contract event 0x1d841fd5b4c806bc5a073d637a8506e1e74d16cb18251b711cb47e133ceafc2d.
//
// Solidity: event LockedOiRatioSet(uint128 indexed marketId, uint256 lockedOiRatioD18)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchLockedOiRatioSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliLockedOiRatioSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "LockedOiRatioSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliLockedOiRatioSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "LockedOiRatioSet", log); err != nil {
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

// ParseLockedOiRatioSet is a log parse operation binding the contract event 0x1d841fd5b4c806bc5a073d637a8506e1e74d16cb18251b711cb47e133ceafc2d.
//
// Solidity: event LockedOiRatioSet(uint128 indexed marketId, uint256 lockedOiRatioD18)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseLockedOiRatioSet(log types.Log) (*PrepsMarketGoerliLockedOiRatioSet, error) {
	event := new(PrepsMarketGoerliLockedOiRatioSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "LockedOiRatioSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliMarketCreatedIterator struct {
	Event *PrepsMarketGoerliMarketCreated // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliMarketCreated)
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
		it.Event = new(PrepsMarketGoerliMarketCreated)
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
func (it *PrepsMarketGoerliMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliMarketCreated represents a MarketCreated event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliMarketCreated struct {
	PerpsMarketId *big.Int
	MarketName    string
	MarketSymbol  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0x032553f94ac1323933f22650ec5b8e232babf1c47efca69383b749463116cc49.
//
// Solidity: event MarketCreated(uint128 indexed perpsMarketId, string marketName, string marketSymbol)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterMarketCreated(opts *bind.FilterOpts, perpsMarketId []*big.Int) (*PrepsMarketGoerliMarketCreatedIterator, error) {

	var perpsMarketIdRule []interface{}
	for _, perpsMarketIdItem := range perpsMarketId {
		perpsMarketIdRule = append(perpsMarketIdRule, perpsMarketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "MarketCreated", perpsMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliMarketCreatedIterator{contract: _PrepsMarketGoerli.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0x032553f94ac1323933f22650ec5b8e232babf1c47efca69383b749463116cc49.
//
// Solidity: event MarketCreated(uint128 indexed perpsMarketId, string marketName, string marketSymbol)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliMarketCreated, perpsMarketId []*big.Int) (event.Subscription, error) {

	var perpsMarketIdRule []interface{}
	for _, perpsMarketIdItem := range perpsMarketId {
		perpsMarketIdRule = append(perpsMarketIdRule, perpsMarketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "MarketCreated", perpsMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliMarketCreated)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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

// ParseMarketCreated is a log parse operation binding the contract event 0x032553f94ac1323933f22650ec5b8e232babf1c47efca69383b749463116cc49.
//
// Solidity: event MarketCreated(uint128 indexed perpsMarketId, string marketName, string marketSymbol)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseMarketCreated(log types.Log) (*PrepsMarketGoerliMarketCreated, error) {
	event := new(PrepsMarketGoerliMarketCreated)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliMarketPriceDataUpdatedIterator is returned from FilterMarketPriceDataUpdated and is used to iterate over the raw logs and unpacked data for MarketPriceDataUpdated events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliMarketPriceDataUpdatedIterator struct {
	Event *PrepsMarketGoerliMarketPriceDataUpdated // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliMarketPriceDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliMarketPriceDataUpdated)
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
		it.Event = new(PrepsMarketGoerliMarketPriceDataUpdated)
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
func (it *PrepsMarketGoerliMarketPriceDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliMarketPriceDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliMarketPriceDataUpdated represents a MarketPriceDataUpdated event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliMarketPriceDataUpdated struct {
	MarketId *big.Int
	FeedId   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketPriceDataUpdated is a free log retrieval operation binding the contract event 0xd5ec22bbdbde803d8585a5aa89d991c67dc62c6a7bc3c4fa000b4e263b6783a0.
//
// Solidity: event MarketPriceDataUpdated(uint128 indexed marketId, bytes32 feedId)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterMarketPriceDataUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*PrepsMarketGoerliMarketPriceDataUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "MarketPriceDataUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliMarketPriceDataUpdatedIterator{contract: _PrepsMarketGoerli.contract, event: "MarketPriceDataUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketPriceDataUpdated is a free log subscription operation binding the contract event 0xd5ec22bbdbde803d8585a5aa89d991c67dc62c6a7bc3c4fa000b4e263b6783a0.
//
// Solidity: event MarketPriceDataUpdated(uint128 indexed marketId, bytes32 feedId)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchMarketPriceDataUpdated(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliMarketPriceDataUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "MarketPriceDataUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliMarketPriceDataUpdated)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "MarketPriceDataUpdated", log); err != nil {
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

// ParseMarketPriceDataUpdated is a log parse operation binding the contract event 0xd5ec22bbdbde803d8585a5aa89d991c67dc62c6a7bc3c4fa000b4e263b6783a0.
//
// Solidity: event MarketPriceDataUpdated(uint128 indexed marketId, bytes32 feedId)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseMarketPriceDataUpdated(log types.Log) (*PrepsMarketGoerliMarketPriceDataUpdated, error) {
	event := new(PrepsMarketGoerliMarketPriceDataUpdated)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "MarketPriceDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliMarketUpdatedIterator is returned from FilterMarketUpdated and is used to iterate over the raw logs and unpacked data for MarketUpdated events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliMarketUpdatedIterator struct {
	Event *PrepsMarketGoerliMarketUpdated // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliMarketUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliMarketUpdated)
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
		it.Event = new(PrepsMarketGoerliMarketUpdated)
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
func (it *PrepsMarketGoerliMarketUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliMarketUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliMarketUpdated represents a MarketUpdated event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliMarketUpdated struct {
	MarketId               *big.Int
	Price                  *big.Int
	Skew                   *big.Int
	Size                   *big.Int
	SizeDelta              *big.Int
	CurrentFundingRate     *big.Int
	CurrentFundingVelocity *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterMarketUpdated is a free log retrieval operation binding the contract event 0xb317f068f30db823aeb6ac6ddbffcbb6c805f558972e6b16625ec58b83f1f3d5.
//
// Solidity: event MarketUpdated(uint128 marketId, uint256 price, int256 skew, uint256 size, int256 sizeDelta, int256 currentFundingRate, int256 currentFundingVelocity)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterMarketUpdated(opts *bind.FilterOpts) (*PrepsMarketGoerliMarketUpdatedIterator, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "MarketUpdated")
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliMarketUpdatedIterator{contract: _PrepsMarketGoerli.contract, event: "MarketUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketUpdated is a free log subscription operation binding the contract event 0xb317f068f30db823aeb6ac6ddbffcbb6c805f558972e6b16625ec58b83f1f3d5.
//
// Solidity: event MarketUpdated(uint128 marketId, uint256 price, int256 skew, uint256 size, int256 sizeDelta, int256 currentFundingRate, int256 currentFundingVelocity)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchMarketUpdated(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliMarketUpdated) (event.Subscription, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "MarketUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliMarketUpdated)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "MarketUpdated", log); err != nil {
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

// ParseMarketUpdated is a log parse operation binding the contract event 0xb317f068f30db823aeb6ac6ddbffcbb6c805f558972e6b16625ec58b83f1f3d5.
//
// Solidity: event MarketUpdated(uint128 marketId, uint256 price, int256 skew, uint256 size, int256 sizeDelta, int256 currentFundingRate, int256 currentFundingVelocity)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseMarketUpdated(log types.Log) (*PrepsMarketGoerliMarketUpdated, error) {
	event := new(PrepsMarketGoerliMarketUpdated)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "MarketUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliMaxCollateralAmountSetIterator is returned from FilterMaxCollateralAmountSet and is used to iterate over the raw logs and unpacked data for MaxCollateralAmountSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliMaxCollateralAmountSetIterator struct {
	Event *PrepsMarketGoerliMaxCollateralAmountSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliMaxCollateralAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliMaxCollateralAmountSet)
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
		it.Event = new(PrepsMarketGoerliMaxCollateralAmountSet)
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
func (it *PrepsMarketGoerliMaxCollateralAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliMaxCollateralAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliMaxCollateralAmountSet represents a MaxCollateralAmountSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliMaxCollateralAmountSet struct {
	SynthMarketId    *big.Int
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMaxCollateralAmountSet is a free log retrieval operation binding the contract event 0xa552ba0f2552ec13b62b479a41e6e7dfcef9f10aea18527e5a3e1b3963bf70d7.
//
// Solidity: event MaxCollateralAmountSet(uint128 indexed synthMarketId, uint256 collateralAmount)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterMaxCollateralAmountSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*PrepsMarketGoerliMaxCollateralAmountSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "MaxCollateralAmountSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliMaxCollateralAmountSetIterator{contract: _PrepsMarketGoerli.contract, event: "MaxCollateralAmountSet", logs: logs, sub: sub}, nil
}

// WatchMaxCollateralAmountSet is a free log subscription operation binding the contract event 0xa552ba0f2552ec13b62b479a41e6e7dfcef9f10aea18527e5a3e1b3963bf70d7.
//
// Solidity: event MaxCollateralAmountSet(uint128 indexed synthMarketId, uint256 collateralAmount)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchMaxCollateralAmountSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliMaxCollateralAmountSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "MaxCollateralAmountSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliMaxCollateralAmountSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "MaxCollateralAmountSet", log); err != nil {
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

// ParseMaxCollateralAmountSet is a log parse operation binding the contract event 0xa552ba0f2552ec13b62b479a41e6e7dfcef9f10aea18527e5a3e1b3963bf70d7.
//
// Solidity: event MaxCollateralAmountSet(uint128 indexed synthMarketId, uint256 collateralAmount)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseMaxCollateralAmountSet(log types.Log) (*PrepsMarketGoerliMaxCollateralAmountSet, error) {
	event := new(PrepsMarketGoerliMaxCollateralAmountSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "MaxCollateralAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliMaxMarketSizeSetIterator is returned from FilterMaxMarketSizeSet and is used to iterate over the raw logs and unpacked data for MaxMarketSizeSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliMaxMarketSizeSetIterator struct {
	Event *PrepsMarketGoerliMaxMarketSizeSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliMaxMarketSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliMaxMarketSizeSet)
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
		it.Event = new(PrepsMarketGoerliMaxMarketSizeSet)
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
func (it *PrepsMarketGoerliMaxMarketSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliMaxMarketSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliMaxMarketSizeSet represents a MaxMarketSizeSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliMaxMarketSizeSet struct {
	MarketId      *big.Int
	MaxMarketSize *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxMarketSizeSet is a free log retrieval operation binding the contract event 0xbd063bd3072a194b255163ab8dfd3400c4ab1cc641b920e7aaf1c11f92cd26cf.
//
// Solidity: event MaxMarketSizeSet(uint128 indexed marketId, uint256 maxMarketSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterMaxMarketSizeSet(opts *bind.FilterOpts, marketId []*big.Int) (*PrepsMarketGoerliMaxMarketSizeSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "MaxMarketSizeSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliMaxMarketSizeSetIterator{contract: _PrepsMarketGoerli.contract, event: "MaxMarketSizeSet", logs: logs, sub: sub}, nil
}

// WatchMaxMarketSizeSet is a free log subscription operation binding the contract event 0xbd063bd3072a194b255163ab8dfd3400c4ab1cc641b920e7aaf1c11f92cd26cf.
//
// Solidity: event MaxMarketSizeSet(uint128 indexed marketId, uint256 maxMarketSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchMaxMarketSizeSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliMaxMarketSizeSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "MaxMarketSizeSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliMaxMarketSizeSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "MaxMarketSizeSet", log); err != nil {
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

// ParseMaxMarketSizeSet is a log parse operation binding the contract event 0xbd063bd3072a194b255163ab8dfd3400c4ab1cc641b920e7aaf1c11f92cd26cf.
//
// Solidity: event MaxMarketSizeSet(uint128 indexed marketId, uint256 maxMarketSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseMaxMarketSizeSet(log types.Log) (*PrepsMarketGoerliMaxMarketSizeSet, error) {
	event := new(PrepsMarketGoerliMaxMarketSizeSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "MaxMarketSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliOrderCommittedIterator is returned from FilterOrderCommitted and is used to iterate over the raw logs and unpacked data for OrderCommitted events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliOrderCommittedIterator struct {
	Event *PrepsMarketGoerliOrderCommitted // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliOrderCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliOrderCommitted)
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
		it.Event = new(PrepsMarketGoerliOrderCommitted)
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
func (it *PrepsMarketGoerliOrderCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliOrderCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliOrderCommitted represents a OrderCommitted event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliOrderCommitted struct {
	MarketId        *big.Int
	AccountId       *big.Int
	OrderType       uint8
	SizeDelta       *big.Int
	AcceptablePrice *big.Int
	SettlementTime  *big.Int
	ExpirationTime  *big.Int
	TrackingCode    [32]byte
	Sender          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderCommitted is a free log retrieval operation binding the contract event 0xc74e8cf044fc959dcef2695245eb90f7e9fb5676532a71eec5b3c513d4793709.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint128 indexed accountId, uint8 orderType, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, uint256 expirationTime, bytes32 indexed trackingCode, address sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterOrderCommitted(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*PrepsMarketGoerliOrderCommittedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var trackingCodeRule []interface{}
	for _, trackingCodeItem := range trackingCode {
		trackingCodeRule = append(trackingCodeRule, trackingCodeItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "OrderCommitted", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliOrderCommittedIterator{contract: _PrepsMarketGoerli.contract, event: "OrderCommitted", logs: logs, sub: sub}, nil
}

// WatchOrderCommitted is a free log subscription operation binding the contract event 0xc74e8cf044fc959dcef2695245eb90f7e9fb5676532a71eec5b3c513d4793709.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint128 indexed accountId, uint8 orderType, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, uint256 expirationTime, bytes32 indexed trackingCode, address sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchOrderCommitted(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliOrderCommitted, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var trackingCodeRule []interface{}
	for _, trackingCodeItem := range trackingCode {
		trackingCodeRule = append(trackingCodeRule, trackingCodeItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "OrderCommitted", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliOrderCommitted)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
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

// ParseOrderCommitted is a log parse operation binding the contract event 0xc74e8cf044fc959dcef2695245eb90f7e9fb5676532a71eec5b3c513d4793709.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint128 indexed accountId, uint8 orderType, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, uint256 expirationTime, bytes32 indexed trackingCode, address sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseOrderCommitted(log types.Log) (*PrepsMarketGoerliOrderCommitted, error) {
	event := new(PrepsMarketGoerliOrderCommitted)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliOrderFeesSetIterator is returned from FilterOrderFeesSet and is used to iterate over the raw logs and unpacked data for OrderFeesSet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliOrderFeesSetIterator struct {
	Event *PrepsMarketGoerliOrderFeesSet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliOrderFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliOrderFeesSet)
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
		it.Event = new(PrepsMarketGoerliOrderFeesSet)
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
func (it *PrepsMarketGoerliOrderFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliOrderFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliOrderFeesSet represents a OrderFeesSet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliOrderFeesSet struct {
	MarketId      *big.Int
	MakerFeeRatio *big.Int
	TakerFeeRatio *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFeesSet is a free log retrieval operation binding the contract event 0x28969f156340ba9eb31589904eb174d3a4b6a37781fa6f7ad289d349d75dd1ee.
//
// Solidity: event OrderFeesSet(uint128 indexed marketId, uint256 makerFeeRatio, uint256 takerFeeRatio)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterOrderFeesSet(opts *bind.FilterOpts, marketId []*big.Int) (*PrepsMarketGoerliOrderFeesSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "OrderFeesSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliOrderFeesSetIterator{contract: _PrepsMarketGoerli.contract, event: "OrderFeesSet", logs: logs, sub: sub}, nil
}

// WatchOrderFeesSet is a free log subscription operation binding the contract event 0x28969f156340ba9eb31589904eb174d3a4b6a37781fa6f7ad289d349d75dd1ee.
//
// Solidity: event OrderFeesSet(uint128 indexed marketId, uint256 makerFeeRatio, uint256 takerFeeRatio)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchOrderFeesSet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliOrderFeesSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "OrderFeesSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliOrderFeesSet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "OrderFeesSet", log); err != nil {
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

// ParseOrderFeesSet is a log parse operation binding the contract event 0x28969f156340ba9eb31589904eb174d3a4b6a37781fa6f7ad289d349d75dd1ee.
//
// Solidity: event OrderFeesSet(uint128 indexed marketId, uint256 makerFeeRatio, uint256 takerFeeRatio)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseOrderFeesSet(log types.Log) (*PrepsMarketGoerliOrderFeesSet, error) {
	event := new(PrepsMarketGoerliOrderFeesSet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "OrderFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliOrderSettledIterator is returned from FilterOrderSettled and is used to iterate over the raw logs and unpacked data for OrderSettled events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliOrderSettledIterator struct {
	Event *PrepsMarketGoerliOrderSettled // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliOrderSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliOrderSettled)
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
		it.Event = new(PrepsMarketGoerliOrderSettled)
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
func (it *PrepsMarketGoerliOrderSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliOrderSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliOrderSettled represents a OrderSettled event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliOrderSettled struct {
	MarketId         *big.Int
	AccountId        *big.Int
	FillPrice        *big.Int
	Pnl              *big.Int
	AccruedFunding   *big.Int
	SizeDelta        *big.Int
	NewSize          *big.Int
	TotalFees        *big.Int
	ReferralFees     *big.Int
	CollectedFees    *big.Int
	SettlementReward *big.Int
	TrackingCode     [32]byte
	Settler          common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderSettled is a free log retrieval operation binding the contract event 0x460080a757ec90719fe90ab2384c0196cdeed071a9fd7ce1ada43481d96b7db5.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed accountId, uint256 fillPrice, int256 pnl, int256 accruedFunding, int128 sizeDelta, int128 newSize, uint256 totalFees, uint256 referralFees, uint256 collectedFees, uint256 settlementReward, bytes32 indexed trackingCode, address settler)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterOrderSettled(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*PrepsMarketGoerliOrderSettledIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var trackingCodeRule []interface{}
	for _, trackingCodeItem := range trackingCode {
		trackingCodeRule = append(trackingCodeRule, trackingCodeItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "OrderSettled", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliOrderSettledIterator{contract: _PrepsMarketGoerli.contract, event: "OrderSettled", logs: logs, sub: sub}, nil
}

// WatchOrderSettled is a free log subscription operation binding the contract event 0x460080a757ec90719fe90ab2384c0196cdeed071a9fd7ce1ada43481d96b7db5.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed accountId, uint256 fillPrice, int256 pnl, int256 accruedFunding, int128 sizeDelta, int128 newSize, uint256 totalFees, uint256 referralFees, uint256 collectedFees, uint256 settlementReward, bytes32 indexed trackingCode, address settler)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchOrderSettled(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliOrderSettled, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var trackingCodeRule []interface{}
	for _, trackingCodeItem := range trackingCode {
		trackingCodeRule = append(trackingCodeRule, trackingCodeItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "OrderSettled", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliOrderSettled)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "OrderSettled", log); err != nil {
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

// ParseOrderSettled is a log parse operation binding the contract event 0x460080a757ec90719fe90ab2384c0196cdeed071a9fd7ce1ada43481d96b7db5.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed accountId, uint256 fillPrice, int256 pnl, int256 accruedFunding, int128 sizeDelta, int128 newSize, uint256 totalFees, uint256 referralFees, uint256 collectedFees, uint256 settlementReward, bytes32 indexed trackingCode, address settler)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseOrderSettled(log types.Log) (*PrepsMarketGoerliOrderSettled, error) {
	event := new(PrepsMarketGoerliOrderSettled)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "OrderSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliOwnerChangedIterator struct {
	Event *PrepsMarketGoerliOwnerChanged // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliOwnerChanged)
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
		it.Event = new(PrepsMarketGoerliOwnerChanged)
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
func (it *PrepsMarketGoerliOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliOwnerChanged represents a OwnerChanged event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*PrepsMarketGoerliOwnerChangedIterator, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliOwnerChangedIterator{contract: _PrepsMarketGoerli.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliOwnerChanged)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseOwnerChanged(log types.Log) (*PrepsMarketGoerliOwnerChanged, error) {
	event := new(PrepsMarketGoerliOwnerChanged)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliOwnerNominatedIterator struct {
	Event *PrepsMarketGoerliOwnerNominated // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliOwnerNominated)
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
		it.Event = new(PrepsMarketGoerliOwnerNominated)
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
func (it *PrepsMarketGoerliOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliOwnerNominated represents a OwnerNominated event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*PrepsMarketGoerliOwnerNominatedIterator, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliOwnerNominatedIterator{contract: _PrepsMarketGoerli.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliOwnerNominated)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseOwnerNominated(log types.Log) (*PrepsMarketGoerliOwnerNominated, error) {
	event := new(PrepsMarketGoerliOwnerNominated)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliPermissionGrantedIterator is returned from FilterPermissionGranted and is used to iterate over the raw logs and unpacked data for PermissionGranted events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliPermissionGrantedIterator struct {
	Event *PrepsMarketGoerliPermissionGranted // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliPermissionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliPermissionGranted)
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
		it.Event = new(PrepsMarketGoerliPermissionGranted)
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
func (it *PrepsMarketGoerliPermissionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliPermissionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliPermissionGranted represents a PermissionGranted event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliPermissionGranted struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionGranted is a free log retrieval operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterPermissionGranted(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*PrepsMarketGoerliPermissionGrantedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var permissionRule []interface{}
	for _, permissionItem := range permission {
		permissionRule = append(permissionRule, permissionItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliPermissionGrantedIterator{contract: _PrepsMarketGoerli.contract, event: "PermissionGranted", logs: logs, sub: sub}, nil
}

// WatchPermissionGranted is a free log subscription operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchPermissionGranted(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliPermissionGranted, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var permissionRule []interface{}
	for _, permissionItem := range permission {
		permissionRule = append(permissionRule, permissionItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliPermissionGranted)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
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

// ParsePermissionGranted is a log parse operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParsePermissionGranted(log types.Log) (*PrepsMarketGoerliPermissionGranted, error) {
	event := new(PrepsMarketGoerliPermissionGranted)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliPermissionRevokedIterator is returned from FilterPermissionRevoked and is used to iterate over the raw logs and unpacked data for PermissionRevoked events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliPermissionRevokedIterator struct {
	Event *PrepsMarketGoerliPermissionRevoked // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliPermissionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliPermissionRevoked)
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
		it.Event = new(PrepsMarketGoerliPermissionRevoked)
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
func (it *PrepsMarketGoerliPermissionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliPermissionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliPermissionRevoked represents a PermissionRevoked event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliPermissionRevoked struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionRevoked is a free log retrieval operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterPermissionRevoked(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*PrepsMarketGoerliPermissionRevokedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var permissionRule []interface{}
	for _, permissionItem := range permission {
		permissionRule = append(permissionRule, permissionItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliPermissionRevokedIterator{contract: _PrepsMarketGoerli.contract, event: "PermissionRevoked", logs: logs, sub: sub}, nil
}

// WatchPermissionRevoked is a free log subscription operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchPermissionRevoked(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliPermissionRevoked, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var permissionRule []interface{}
	for _, permissionItem := range permission {
		permissionRule = append(permissionRule, permissionItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliPermissionRevoked)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
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

// ParsePermissionRevoked is a log parse operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParsePermissionRevoked(log types.Log) (*PrepsMarketGoerliPermissionRevoked, error) {
	event := new(PrepsMarketGoerliPermissionRevoked)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliPositionLiquidatedIterator is returned from FilterPositionLiquidated and is used to iterate over the raw logs and unpacked data for PositionLiquidated events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliPositionLiquidatedIterator struct {
	Event *PrepsMarketGoerliPositionLiquidated // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliPositionLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliPositionLiquidated)
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
		it.Event = new(PrepsMarketGoerliPositionLiquidated)
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
func (it *PrepsMarketGoerliPositionLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliPositionLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliPositionLiquidated represents a PositionLiquidated event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliPositionLiquidated struct {
	AccountId           *big.Int
	MarketId            *big.Int
	AmountLiquidated    *big.Int
	CurrentPositionSize *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPositionLiquidated is a free log retrieval operation binding the contract event 0xd583c0e2965aae317f4a9a6603c0c75602b9bc97dc7c5fc6446b0ba8d3ff1bb2.
//
// Solidity: event PositionLiquidated(uint128 indexed accountId, uint128 indexed marketId, uint256 amountLiquidated, int128 currentPositionSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterPositionLiquidated(opts *bind.FilterOpts, accountId []*big.Int, marketId []*big.Int) (*PrepsMarketGoerliPositionLiquidatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "PositionLiquidated", accountIdRule, marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliPositionLiquidatedIterator{contract: _PrepsMarketGoerli.contract, event: "PositionLiquidated", logs: logs, sub: sub}, nil
}

// WatchPositionLiquidated is a free log subscription operation binding the contract event 0xd583c0e2965aae317f4a9a6603c0c75602b9bc97dc7c5fc6446b0ba8d3ff1bb2.
//
// Solidity: event PositionLiquidated(uint128 indexed accountId, uint128 indexed marketId, uint256 amountLiquidated, int128 currentPositionSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchPositionLiquidated(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliPositionLiquidated, accountId []*big.Int, marketId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "PositionLiquidated", accountIdRule, marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliPositionLiquidated)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
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

// ParsePositionLiquidated is a log parse operation binding the contract event 0xd583c0e2965aae317f4a9a6603c0c75602b9bc97dc7c5fc6446b0ba8d3ff1bb2.
//
// Solidity: event PositionLiquidated(uint128 indexed accountId, uint128 indexed marketId, uint256 amountLiquidated, int128 currentPositionSize)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParsePositionLiquidated(log types.Log) (*PrepsMarketGoerliPositionLiquidated, error) {
	event := new(PrepsMarketGoerliPositionLiquidated)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliPreviousOrderExpiredIterator is returned from FilterPreviousOrderExpired and is used to iterate over the raw logs and unpacked data for PreviousOrderExpired events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliPreviousOrderExpiredIterator struct {
	Event *PrepsMarketGoerliPreviousOrderExpired // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliPreviousOrderExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliPreviousOrderExpired)
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
		it.Event = new(PrepsMarketGoerliPreviousOrderExpired)
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
func (it *PrepsMarketGoerliPreviousOrderExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliPreviousOrderExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliPreviousOrderExpired represents a PreviousOrderExpired event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliPreviousOrderExpired struct {
	MarketId        *big.Int
	AccountId       *big.Int
	SizeDelta       *big.Int
	AcceptablePrice *big.Int
	SettlementTime  *big.Int
	TrackingCode    [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPreviousOrderExpired is a free log retrieval operation binding the contract event 0x6d83c6751813f50325d75bc054621f83299659c5814d1e5fe6ac117860710dde.
//
// Solidity: event PreviousOrderExpired(uint128 indexed marketId, uint128 indexed accountId, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, bytes32 indexed trackingCode)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterPreviousOrderExpired(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*PrepsMarketGoerliPreviousOrderExpiredIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var trackingCodeRule []interface{}
	for _, trackingCodeItem := range trackingCode {
		trackingCodeRule = append(trackingCodeRule, trackingCodeItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "PreviousOrderExpired", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliPreviousOrderExpiredIterator{contract: _PrepsMarketGoerli.contract, event: "PreviousOrderExpired", logs: logs, sub: sub}, nil
}

// WatchPreviousOrderExpired is a free log subscription operation binding the contract event 0x6d83c6751813f50325d75bc054621f83299659c5814d1e5fe6ac117860710dde.
//
// Solidity: event PreviousOrderExpired(uint128 indexed marketId, uint128 indexed accountId, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, bytes32 indexed trackingCode)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchPreviousOrderExpired(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliPreviousOrderExpired, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var trackingCodeRule []interface{}
	for _, trackingCodeItem := range trackingCode {
		trackingCodeRule = append(trackingCodeRule, trackingCodeItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "PreviousOrderExpired", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliPreviousOrderExpired)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "PreviousOrderExpired", log); err != nil {
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

// ParsePreviousOrderExpired is a log parse operation binding the contract event 0x6d83c6751813f50325d75bc054621f83299659c5814d1e5fe6ac117860710dde.
//
// Solidity: event PreviousOrderExpired(uint128 indexed marketId, uint128 indexed accountId, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, bytes32 indexed trackingCode)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParsePreviousOrderExpired(log types.Log) (*PrepsMarketGoerliPreviousOrderExpired, error) {
	event := new(PrepsMarketGoerliPreviousOrderExpired)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "PreviousOrderExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliReferrerShareUpdatedIterator is returned from FilterReferrerShareUpdated and is used to iterate over the raw logs and unpacked data for ReferrerShareUpdated events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliReferrerShareUpdatedIterator struct {
	Event *PrepsMarketGoerliReferrerShareUpdated // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliReferrerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliReferrerShareUpdated)
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
		it.Event = new(PrepsMarketGoerliReferrerShareUpdated)
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
func (it *PrepsMarketGoerliReferrerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliReferrerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliReferrerShareUpdated represents a ReferrerShareUpdated event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliReferrerShareUpdated struct {
	Referrer      common.Address
	ShareRatioD18 *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReferrerShareUpdated is a free log retrieval operation binding the contract event 0xa225c555f4cd21a533ad4e01eaf30153c84ca28265d954a651410d3c1e56242c.
//
// Solidity: event ReferrerShareUpdated(address referrer, uint256 shareRatioD18)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterReferrerShareUpdated(opts *bind.FilterOpts) (*PrepsMarketGoerliReferrerShareUpdatedIterator, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "ReferrerShareUpdated")
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliReferrerShareUpdatedIterator{contract: _PrepsMarketGoerli.contract, event: "ReferrerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchReferrerShareUpdated is a free log subscription operation binding the contract event 0xa225c555f4cd21a533ad4e01eaf30153c84ca28265d954a651410d3c1e56242c.
//
// Solidity: event ReferrerShareUpdated(address referrer, uint256 shareRatioD18)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchReferrerShareUpdated(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliReferrerShareUpdated) (event.Subscription, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "ReferrerShareUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliReferrerShareUpdated)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
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

// ParseReferrerShareUpdated is a log parse operation binding the contract event 0xa225c555f4cd21a533ad4e01eaf30153c84ca28265d954a651410d3c1e56242c.
//
// Solidity: event ReferrerShareUpdated(address referrer, uint256 shareRatioD18)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseReferrerShareUpdated(log types.Log) (*PrepsMarketGoerliReferrerShareUpdated, error) {
	event := new(PrepsMarketGoerliReferrerShareUpdated)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliSettlementStrategyAddedIterator is returned from FilterSettlementStrategyAdded and is used to iterate over the raw logs and unpacked data for SettlementStrategyAdded events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliSettlementStrategyAddedIterator struct {
	Event *PrepsMarketGoerliSettlementStrategyAdded // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliSettlementStrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliSettlementStrategyAdded)
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
		it.Event = new(PrepsMarketGoerliSettlementStrategyAdded)
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
func (it *PrepsMarketGoerliSettlementStrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliSettlementStrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliSettlementStrategyAdded represents a SettlementStrategyAdded event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliSettlementStrategyAdded struct {
	MarketId   *big.Int
	Strategy   SettlementStrategyData
	StrategyId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyAdded is a free log retrieval operation binding the contract event 0x6ea79b5183a0fa7ead0ea991cd5e5fa23e2da92b703bbe111479fffcdea58827.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy, uint256 indexed strategyId)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterSettlementStrategyAdded(opts *bind.FilterOpts, marketId []*big.Int, strategyId []*big.Int) (*PrepsMarketGoerliSettlementStrategyAddedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "SettlementStrategyAdded", marketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliSettlementStrategyAddedIterator{contract: _PrepsMarketGoerli.contract, event: "SettlementStrategyAdded", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyAdded is a free log subscription operation binding the contract event 0x6ea79b5183a0fa7ead0ea991cd5e5fa23e2da92b703bbe111479fffcdea58827.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy, uint256 indexed strategyId)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchSettlementStrategyAdded(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliSettlementStrategyAdded, marketId []*big.Int, strategyId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "SettlementStrategyAdded", marketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliSettlementStrategyAdded)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
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

// ParseSettlementStrategyAdded is a log parse operation binding the contract event 0x6ea79b5183a0fa7ead0ea991cd5e5fa23e2da92b703bbe111479fffcdea58827.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy, uint256 indexed strategyId)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseSettlementStrategyAdded(log types.Log) (*PrepsMarketGoerliSettlementStrategyAdded, error) {
	event := new(PrepsMarketGoerliSettlementStrategyAdded)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliSettlementStrategyEnabledIterator is returned from FilterSettlementStrategyEnabled and is used to iterate over the raw logs and unpacked data for SettlementStrategyEnabled events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliSettlementStrategyEnabledIterator struct {
	Event *PrepsMarketGoerliSettlementStrategyEnabled // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliSettlementStrategyEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliSettlementStrategyEnabled)
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
		it.Event = new(PrepsMarketGoerliSettlementStrategyEnabled)
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
func (it *PrepsMarketGoerliSettlementStrategyEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliSettlementStrategyEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliSettlementStrategyEnabled represents a SettlementStrategyEnabled event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliSettlementStrategyEnabled struct {
	MarketId   *big.Int
	StrategyId *big.Int
	Enabled    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyEnabled is a free log retrieval operation binding the contract event 0x4803eb4e51b45b2471b0c87d8c17e0c2d8961ae080cc4a930883eab155e0fa22.
//
// Solidity: event SettlementStrategyEnabled(uint128 indexed marketId, uint256 strategyId, bool enabled)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterSettlementStrategyEnabled(opts *bind.FilterOpts, marketId []*big.Int) (*PrepsMarketGoerliSettlementStrategyEnabledIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "SettlementStrategyEnabled", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliSettlementStrategyEnabledIterator{contract: _PrepsMarketGoerli.contract, event: "SettlementStrategyEnabled", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyEnabled is a free log subscription operation binding the contract event 0x4803eb4e51b45b2471b0c87d8c17e0c2d8961ae080cc4a930883eab155e0fa22.
//
// Solidity: event SettlementStrategyEnabled(uint128 indexed marketId, uint256 strategyId, bool enabled)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchSettlementStrategyEnabled(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliSettlementStrategyEnabled, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "SettlementStrategyEnabled", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliSettlementStrategyEnabled)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "SettlementStrategyEnabled", log); err != nil {
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

// ParseSettlementStrategyEnabled is a log parse operation binding the contract event 0x4803eb4e51b45b2471b0c87d8c17e0c2d8961ae080cc4a930883eab155e0fa22.
//
// Solidity: event SettlementStrategyEnabled(uint128 indexed marketId, uint256 strategyId, bool enabled)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseSettlementStrategyEnabled(log types.Log) (*PrepsMarketGoerliSettlementStrategyEnabled, error) {
	event := new(PrepsMarketGoerliSettlementStrategyEnabled)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "SettlementStrategyEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliSynthDeductionPrioritySetIterator is returned from FilterSynthDeductionPrioritySet and is used to iterate over the raw logs and unpacked data for SynthDeductionPrioritySet events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliSynthDeductionPrioritySetIterator struct {
	Event *PrepsMarketGoerliSynthDeductionPrioritySet // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliSynthDeductionPrioritySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliSynthDeductionPrioritySet)
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
		it.Event = new(PrepsMarketGoerliSynthDeductionPrioritySet)
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
func (it *PrepsMarketGoerliSynthDeductionPrioritySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliSynthDeductionPrioritySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliSynthDeductionPrioritySet represents a SynthDeductionPrioritySet event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliSynthDeductionPrioritySet struct {
	NewSynthDeductionPriority []*big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSynthDeductionPrioritySet is a free log retrieval operation binding the contract event 0xa6beea856d32db9e9614e4af02fc7d6a3fa8337a13e4a48486c142a9a9a8ed8f.
//
// Solidity: event SynthDeductionPrioritySet(uint128[] newSynthDeductionPriority)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterSynthDeductionPrioritySet(opts *bind.FilterOpts) (*PrepsMarketGoerliSynthDeductionPrioritySetIterator, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "SynthDeductionPrioritySet")
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliSynthDeductionPrioritySetIterator{contract: _PrepsMarketGoerli.contract, event: "SynthDeductionPrioritySet", logs: logs, sub: sub}, nil
}

// WatchSynthDeductionPrioritySet is a free log subscription operation binding the contract event 0xa6beea856d32db9e9614e4af02fc7d6a3fa8337a13e4a48486c142a9a9a8ed8f.
//
// Solidity: event SynthDeductionPrioritySet(uint128[] newSynthDeductionPriority)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchSynthDeductionPrioritySet(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliSynthDeductionPrioritySet) (event.Subscription, error) {

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "SynthDeductionPrioritySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliSynthDeductionPrioritySet)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "SynthDeductionPrioritySet", log); err != nil {
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

// ParseSynthDeductionPrioritySet is a log parse operation binding the contract event 0xa6beea856d32db9e9614e4af02fc7d6a3fa8337a13e4a48486c142a9a9a8ed8f.
//
// Solidity: event SynthDeductionPrioritySet(uint128[] newSynthDeductionPriority)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseSynthDeductionPrioritySet(log types.Log) (*PrepsMarketGoerliSynthDeductionPrioritySet, error) {
	event := new(PrepsMarketGoerliSynthDeductionPrioritySet)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "SynthDeductionPrioritySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrepsMarketGoerliUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliUpgradedIterator struct {
	Event *PrepsMarketGoerliUpgraded // Event containing the contract specifics and raw log

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
func (it *PrepsMarketGoerliUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrepsMarketGoerliUpgraded)
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
		it.Event = new(PrepsMarketGoerliUpgraded)
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
func (it *PrepsMarketGoerliUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrepsMarketGoerliUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrepsMarketGoerliUpgraded represents a Upgraded event raised by the PrepsMarketGoerli contract.
type PrepsMarketGoerliUpgraded struct {
	Self           common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) FilterUpgraded(opts *bind.FilterOpts, self []common.Address) (*PrepsMarketGoerliUpgradedIterator, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.FilterLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return &PrepsMarketGoerliUpgradedIterator{contract: _PrepsMarketGoerli.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PrepsMarketGoerliUpgraded, self []common.Address) (event.Subscription, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _PrepsMarketGoerli.contract.WatchLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrepsMarketGoerliUpgraded)
				if err := _PrepsMarketGoerli.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_PrepsMarketGoerli *PrepsMarketGoerliFilterer) ParseUpgraded(log types.Log) (*PrepsMarketGoerliUpgraded, error) {
	event := new(PrepsMarketGoerliUpgraded)
	if err := _PrepsMarketGoerli.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
