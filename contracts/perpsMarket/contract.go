// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package perpsMarket

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
	Disabled                  bool
}

// PerpsMarketMetaData contains all meta data concerning the PerpsMarket contract.
var PerpsMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"InvalidAccountId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"InvalidPermission\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"OnlyAccountTokenProxy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"PermissionNotGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedAccountId\",\"type\":\"uint128\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountLastInteraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountPermissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"permissions\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIAccountModule.AccountPermissions[]\",\"name\":\"accountPerms\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"notifyAccountTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"renouncePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"DeniedMulticallTarget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RecursiveMulticall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"InvalidMarket\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parameter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PerpsMarketNotInitialized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"globalPerpsMarketId\",\"type\":\"uint128\"}],\"name\":\"FactoryInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"marketSymbol\",\"type\":\"string\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedMarketId\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"marketSymbol\",\"type\":\"string\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeFactory\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"minimumCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"reportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISpotMarketSystem\",\"name\":\"spotMarket\",\"type\":\"address\"}],\"name\":\"setSpotMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISynthetixSystem\",\"name\":\"synthetix\",\"type\":\"address\"}],\"name\":\"setSynthetix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountLiquidatable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"availableUsdDenominated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredUsdDenominated\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateralAvailableForWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientSynthCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"}],\"name\":\"InvalidAmountDelta\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"MaxCollateralExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"maxCollateralsPerAccount\",\"type\":\"uint128\"}],\"name\":\"MaxCollateralsPerAccountReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint128ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOrderExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"PriceFeedNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"SynthNotEnabledForCollateral\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"CollateralModified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAvailableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableMargin\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getOpenPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"totalPnl\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"accruedFunding\",\"type\":\"int256\"},{\"internalType\":\"int128\",\"name\":\"positionSize\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getRequiredMargins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredInitialMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMaintenanceMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAccumulatedLiquidationRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getWithdrawableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"withdrawableMargin\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"}],\"name\":\"modifyCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"totalAccountOpenInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"totalCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"currentFundingRate\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"currentFundingVelocity\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"orderSize\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"fillPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"skew\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIPerpsMarketModule.MarketSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"indexPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"maxOpenInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"metadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"skew\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"}],\"name\":\"AcceptablePriceExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"availableMargin\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"minMargin\",\"type\":\"uint256\"}],\"name\":\"InsufficientMargin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"}],\"name\":\"InvalidSettlementStrategy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"newSideSize\",\"type\":\"int256\"}],\"name\":\"MaxOpenInterestReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"maxPositionsPerAccount\",\"type\":\"uint128\"}],\"name\":\"MaxPositionsPerAccountReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSizeOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OrderCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"PreviousOrderExpired\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"commitment\",\"type\":\"tuple\"}],\"name\":\"commitOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"internalType\":\"structAsyncOrder.Data\",\"name\":\"retOrder\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"}],\"name\":\"computeOrderFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"internalType\":\"structAsyncOrder.Data\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"}],\"name\":\"requiredMarginForOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredMargin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leftover\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarginError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt128ToUint128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"}],\"name\":\"SettlementStrategyNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementExpiration\",\"type\":\"uint256\"}],\"name\":\"SettlementWindowExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"}],\"name\":\"SettlementWindowNotOpen\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"skew\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"sizeDelta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"}],\"name\":\"MarketUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"accruedFunding\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"newSize\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"}],\"name\":\"OrderSettled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"settlePythOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"NotEligibleForLiquidation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fullLiquidation\",\"type\":\"bool\"}],\"name\":\"AccountLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLiquidated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"currentPositionSize\",\"type\":\"int128\"}],\"name\":\"PositionLiquidated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"canLiquidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isEligible\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidationReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidateFlagged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidationReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"liquidationCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationInWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestLiquidationTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"name\":\"FundingParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maintenanceMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"name\":\"LiquidationParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedOiRatioD18\",\"type\":\"uint256\"}],\"name\":\"LockedOiRatioSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"MarketPriceDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLiquidationPd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"endorsedLiquidator\",\"type\":\"address\"}],\"name\":\"MaxLiquidationParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"name\":\"MaxMarketSizeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerFeeRatio\",\"type\":\"uint256\"}],\"name\":\"OrderFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"SettlementStrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SettlementStrategyEnabled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"}],\"name\":\"addSettlementStrategy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getFundingParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getLiquidationParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginScalarD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getLockedOiRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMaxLiquidationParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationPd\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"endorsedLiquidator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMaxMarketSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getOrderFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"getSettlementStrategy\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"settlementStrategy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"name\":\"setFundingParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginScalarD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"name\":\"setLiquidationParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"lockedOiRatioD18\",\"type\":\"uint256\"}],\"name\":\"setLockedOiRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationPd\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"endorsedLiquidator\",\"type\":\"address\"}],\"name\":\"setMaxLiquidationParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"name\":\"setMaxMarketSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeRatio\",\"type\":\"uint256\"}],\"name\":\"setOrderFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setSettlementStrategyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"updatePriceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidFeeCollector\",\"type\":\"address\"}],\"name\":\"InvalidFeeCollectorInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"InvalidReferrerShareRatio\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"name\":\"LiquidationRewardGuardsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"MaxCollateralAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"maxPositionsPerAccount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"maxCollateralsPerAccount\",\"type\":\"uint128\"}],\"name\":\"PerAccountCapsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"ReferrerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"newSynthDeductionPriority\",\"type\":\"uint128[]\"}],\"name\":\"SynthDeductionPrioritySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidationRewardGuards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"marketIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMaxCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPerAccountCaps\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"maxPositionsPerAccount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxCollateralsPerAccount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"getReferrerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSynthDeductionPriority\",\"outputs\":[{\"internalType\":\"uint128[]\",\"name\":\"\",\"type\":\"uint128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"name\":\"setLiquidationRewardGuards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"setMaxCollateralAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"maxPositionsPerAccount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxCollateralsPerAccount\",\"type\":\"uint128\"}],\"name\":\"setPerAccountCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128[]\",\"name\":\"newSynthDeductionPriority\",\"type\":\"uint128[]\"}],\"name\":\"setSynthDeductionPriority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalGlobalCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"updateReferrerShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PerpsMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use PerpsMarketMetaData.ABI instead.
var PerpsMarketABI = PerpsMarketMetaData.ABI

// PerpsMarket is an auto generated Go binding around an Ethereum contract.
type PerpsMarket struct {
	PerpsMarketCaller     // Read-only binding to the contract
	PerpsMarketTransactor // Write-only binding to the contract
	PerpsMarketFilterer   // Log filterer for contract events
}

// PerpsMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type PerpsMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpsMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PerpsMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpsMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PerpsMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpsMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PerpsMarketSession struct {
	Contract     *PerpsMarket      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PerpsMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PerpsMarketCallerSession struct {
	Contract *PerpsMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PerpsMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PerpsMarketTransactorSession struct {
	Contract     *PerpsMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PerpsMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type PerpsMarketRaw struct {
	Contract *PerpsMarket // Generic contract binding to access the raw methods on
}

// PerpsMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PerpsMarketCallerRaw struct {
	Contract *PerpsMarketCaller // Generic read-only contract binding to access the raw methods on
}

// PerpsMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PerpsMarketTransactorRaw struct {
	Contract *PerpsMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerpsMarket creates a new instance of PerpsMarket, bound to a specific deployed contract.
func NewPerpsMarket(address common.Address, backend bind.ContractBackend) (*PerpsMarket, error) {
	contract, err := bindPerpsMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PerpsMarket{PerpsMarketCaller: PerpsMarketCaller{contract: contract}, PerpsMarketTransactor: PerpsMarketTransactor{contract: contract}, PerpsMarketFilterer: PerpsMarketFilterer{contract: contract}}, nil
}

// NewPerpsMarketCaller creates a new read-only instance of PerpsMarket, bound to a specific deployed contract.
func NewPerpsMarketCaller(address common.Address, caller bind.ContractCaller) (*PerpsMarketCaller, error) {
	contract, err := bindPerpsMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketCaller{contract: contract}, nil
}

// NewPerpsMarketTransactor creates a new write-only instance of PerpsMarket, bound to a specific deployed contract.
func NewPerpsMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*PerpsMarketTransactor, error) {
	contract, err := bindPerpsMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketTransactor{contract: contract}, nil
}

// NewPerpsMarketFilterer creates a new log filterer instance of PerpsMarket, bound to a specific deployed contract.
func NewPerpsMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*PerpsMarketFilterer, error) {
	contract, err := bindPerpsMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketFilterer{contract: contract}, nil
}

// bindPerpsMarket binds a generic wrapper to an already deployed contract.
func bindPerpsMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PerpsMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerpsMarket *PerpsMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerpsMarket.Contract.PerpsMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerpsMarket *PerpsMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarket.Contract.PerpsMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerpsMarket *PerpsMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerpsMarket.Contract.PerpsMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerpsMarket *PerpsMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerpsMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerpsMarket *PerpsMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerpsMarket *PerpsMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerpsMarket.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_PerpsMarket *PerpsMarketCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_PerpsMarket *PerpsMarketSession) PRECISION() (*big.Int, error) {
	return _PerpsMarket.Contract.PRECISION(&_PerpsMarket.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_PerpsMarket *PerpsMarketCallerSession) PRECISION() (*big.Int, error) {
	return _PerpsMarket.Contract.PRECISION(&_PerpsMarket.CallOpts)
}

// CanLiquidate is a free data retrieval call binding the contract method 0x9b922bba.
//
// Solidity: function canLiquidate(uint128 accountId) view returns(bool isEligible)
func (_PerpsMarket *PerpsMarketCaller) CanLiquidate(opts *bind.CallOpts, accountId *big.Int) (bool, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "canLiquidate", accountId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanLiquidate is a free data retrieval call binding the contract method 0x9b922bba.
//
// Solidity: function canLiquidate(uint128 accountId) view returns(bool isEligible)
func (_PerpsMarket *PerpsMarketSession) CanLiquidate(accountId *big.Int) (bool, error) {
	return _PerpsMarket.Contract.CanLiquidate(&_PerpsMarket.CallOpts, accountId)
}

// CanLiquidate is a free data retrieval call binding the contract method 0x9b922bba.
//
// Solidity: function canLiquidate(uint128 accountId) view returns(bool isEligible)
func (_PerpsMarket *PerpsMarketCallerSession) CanLiquidate(accountId *big.Int) (bool, error) {
	return _PerpsMarket.Contract.CanLiquidate(&_PerpsMarket.CallOpts, accountId)
}

// ComputeOrderFees is a free data retrieval call binding the contract method 0x98ef15a2.
//
// Solidity: function computeOrderFees(uint128 marketId, int128 sizeDelta) view returns(uint256 orderFees, uint256 fillPrice)
func (_PerpsMarket *PerpsMarketCaller) ComputeOrderFees(opts *bind.CallOpts, marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "computeOrderFees", marketId, sizeDelta)

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
func (_PerpsMarket *PerpsMarketSession) ComputeOrderFees(marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	return _PerpsMarket.Contract.ComputeOrderFees(&_PerpsMarket.CallOpts, marketId, sizeDelta)
}

// ComputeOrderFees is a free data retrieval call binding the contract method 0x98ef15a2.
//
// Solidity: function computeOrderFees(uint128 marketId, int128 sizeDelta) view returns(uint256 orderFees, uint256 fillPrice)
func (_PerpsMarket *PerpsMarketCallerSession) ComputeOrderFees(marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	return _PerpsMarket.Contract.ComputeOrderFees(&_PerpsMarket.CallOpts, marketId, sizeDelta)
}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_PerpsMarket *PerpsMarketCaller) CurrentFundingRate(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "currentFundingRate", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_PerpsMarket *PerpsMarketSession) CurrentFundingRate(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.CurrentFundingRate(&_PerpsMarket.CallOpts, marketId)
}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_PerpsMarket *PerpsMarketCallerSession) CurrentFundingRate(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.CurrentFundingRate(&_PerpsMarket.CallOpts, marketId)
}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_PerpsMarket *PerpsMarketCaller) CurrentFundingVelocity(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "currentFundingVelocity", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_PerpsMarket *PerpsMarketSession) CurrentFundingVelocity(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.CurrentFundingVelocity(&_PerpsMarket.CallOpts, marketId)
}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_PerpsMarket *PerpsMarketCallerSession) CurrentFundingVelocity(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.CurrentFundingVelocity(&_PerpsMarket.CallOpts, marketId)
}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) FillPrice(opts *bind.CallOpts, marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "fillPrice", marketId, orderSize, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) FillPrice(marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.FillPrice(&_PerpsMarket.CallOpts, marketId, orderSize, price)
}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) FillPrice(marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.FillPrice(&_PerpsMarket.CallOpts, marketId, orderSize, price)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) GetAccountLastInteraction(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getAccountLastInteraction", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetAccountLastInteraction(&_PerpsMarket.CallOpts, accountId)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetAccountLastInteraction(&_PerpsMarket.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_PerpsMarket *PerpsMarketCaller) GetAccountOwner(opts *bind.CallOpts, accountId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getAccountOwner", accountId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_PerpsMarket *PerpsMarketSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _PerpsMarket.Contract.GetAccountOwner(&_PerpsMarket.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_PerpsMarket *PerpsMarketCallerSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _PerpsMarket.Contract.GetAccountOwner(&_PerpsMarket.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_PerpsMarket *PerpsMarketCaller) GetAccountPermissions(opts *bind.CallOpts, accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getAccountPermissions", accountId)

	if err != nil {
		return *new([]IAccountModuleAccountPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAccountModuleAccountPermissions)).(*[]IAccountModuleAccountPermissions)

	return out0, err

}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_PerpsMarket *PerpsMarketSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _PerpsMarket.Contract.GetAccountPermissions(&_PerpsMarket.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_PerpsMarket *PerpsMarketCallerSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _PerpsMarket.Contract.GetAccountPermissions(&_PerpsMarket.CallOpts, accountId)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_PerpsMarket *PerpsMarketCaller) GetAccountTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getAccountTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_PerpsMarket *PerpsMarketSession) GetAccountTokenAddress() (common.Address, error) {
	return _PerpsMarket.Contract.GetAccountTokenAddress(&_PerpsMarket.CallOpts)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_PerpsMarket *PerpsMarketCallerSession) GetAccountTokenAddress() (common.Address, error) {
	return _PerpsMarket.Contract.GetAccountTokenAddress(&_PerpsMarket.CallOpts)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_PerpsMarket *PerpsMarketCaller) GetAssociatedSystem(opts *bind.CallOpts, id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getAssociatedSystem", id)

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
func (_PerpsMarket *PerpsMarketSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _PerpsMarket.Contract.GetAssociatedSystem(&_PerpsMarket.CallOpts, id)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_PerpsMarket *PerpsMarketCallerSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _PerpsMarket.Contract.GetAssociatedSystem(&_PerpsMarket.CallOpts, id)
}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_PerpsMarket *PerpsMarketCaller) GetAvailableMargin(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getAvailableMargin", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_PerpsMarket *PerpsMarketSession) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetAvailableMargin(&_PerpsMarket.CallOpts, accountId)
}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_PerpsMarket *PerpsMarketCallerSession) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetAvailableMargin(&_PerpsMarket.CallOpts, accountId)
}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) GetCollateralAmount(opts *bind.CallOpts, accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getCollateralAmount", accountId, synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) GetCollateralAmount(accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetCollateralAmount(&_PerpsMarket.CallOpts, accountId, synthMarketId)
}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) GetCollateralAmount(accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetCollateralAmount(&_PerpsMarket.CallOpts, accountId, synthMarketId)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_PerpsMarket *PerpsMarketCaller) GetDeniers(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getDeniers", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_PerpsMarket *PerpsMarketSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _PerpsMarket.Contract.GetDeniers(&_PerpsMarket.CallOpts, feature)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_PerpsMarket *PerpsMarketCallerSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _PerpsMarket.Contract.GetDeniers(&_PerpsMarket.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_PerpsMarket *PerpsMarketCaller) GetFeatureFlagAllowAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getFeatureFlagAllowAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_PerpsMarket *PerpsMarketSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _PerpsMarket.Contract.GetFeatureFlagAllowAll(&_PerpsMarket.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_PerpsMarket *PerpsMarketCallerSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _PerpsMarket.Contract.GetFeatureFlagAllowAll(&_PerpsMarket.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_PerpsMarket *PerpsMarketCaller) GetFeatureFlagAllowlist(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getFeatureFlagAllowlist", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_PerpsMarket *PerpsMarketSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _PerpsMarket.Contract.GetFeatureFlagAllowlist(&_PerpsMarket.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_PerpsMarket *PerpsMarketCallerSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _PerpsMarket.Contract.GetFeatureFlagAllowlist(&_PerpsMarket.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_PerpsMarket *PerpsMarketCaller) GetFeatureFlagDenyAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getFeatureFlagDenyAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_PerpsMarket *PerpsMarketSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _PerpsMarket.Contract.GetFeatureFlagDenyAll(&_PerpsMarket.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_PerpsMarket *PerpsMarketCallerSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _PerpsMarket.Contract.GetFeatureFlagDenyAll(&_PerpsMarket.CallOpts, feature)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_PerpsMarket *PerpsMarketCaller) GetFeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getFeeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_PerpsMarket *PerpsMarketSession) GetFeeCollector() (common.Address, error) {
	return _PerpsMarket.Contract.GetFeeCollector(&_PerpsMarket.CallOpts)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_PerpsMarket *PerpsMarketCallerSession) GetFeeCollector() (common.Address, error) {
	return _PerpsMarket.Contract.GetFeeCollector(&_PerpsMarket.CallOpts)
}

// GetFundingParameters is a free data retrieval call binding the contract method 0x1b68d8fa.
//
// Solidity: function getFundingParameters(uint128 marketId) view returns(uint256 skewScale, uint256 maxFundingVelocity)
func (_PerpsMarket *PerpsMarketCaller) GetFundingParameters(opts *bind.CallOpts, marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getFundingParameters", marketId)

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
func (_PerpsMarket *PerpsMarketSession) GetFundingParameters(marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	return _PerpsMarket.Contract.GetFundingParameters(&_PerpsMarket.CallOpts, marketId)
}

// GetFundingParameters is a free data retrieval call binding the contract method 0x1b68d8fa.
//
// Solidity: function getFundingParameters(uint128 marketId) view returns(uint256 skewScale, uint256 maxFundingVelocity)
func (_PerpsMarket *PerpsMarketCallerSession) GetFundingParameters(marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	return _PerpsMarket.Contract.GetFundingParameters(&_PerpsMarket.CallOpts, marketId)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_PerpsMarket *PerpsMarketCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_PerpsMarket *PerpsMarketSession) GetImplementation() (common.Address, error) {
	return _PerpsMarket.Contract.GetImplementation(&_PerpsMarket.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_PerpsMarket *PerpsMarketCallerSession) GetImplementation() (common.Address, error) {
	return _PerpsMarket.Contract.GetImplementation(&_PerpsMarket.CallOpts)
}

// GetLiquidationParameters is a free data retrieval call binding the contract method 0xf94363a6.
//
// Solidity: function getLiquidationParameters(uint128 marketId) view returns(uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin)
func (_PerpsMarket *PerpsMarketCaller) GetLiquidationParameters(opts *bind.CallOpts, marketId *big.Int) (struct {
	InitialMarginRatioD18        *big.Int
	MinimumInitialMarginRatioD18 *big.Int
	MaintenanceMarginScalarD18   *big.Int
	LiquidationRewardRatioD18    *big.Int
	MinimumPositionMargin        *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getLiquidationParameters", marketId)

	outstruct := new(struct {
		InitialMarginRatioD18        *big.Int
		MinimumInitialMarginRatioD18 *big.Int
		MaintenanceMarginScalarD18   *big.Int
		LiquidationRewardRatioD18    *big.Int
		MinimumPositionMargin        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialMarginRatioD18 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MinimumInitialMarginRatioD18 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaintenanceMarginScalarD18 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LiquidationRewardRatioD18 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MinimumPositionMargin = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLiquidationParameters is a free data retrieval call binding the contract method 0xf94363a6.
//
// Solidity: function getLiquidationParameters(uint128 marketId) view returns(uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin)
func (_PerpsMarket *PerpsMarketSession) GetLiquidationParameters(marketId *big.Int) (struct {
	InitialMarginRatioD18        *big.Int
	MinimumInitialMarginRatioD18 *big.Int
	MaintenanceMarginScalarD18   *big.Int
	LiquidationRewardRatioD18    *big.Int
	MinimumPositionMargin        *big.Int
}, error) {
	return _PerpsMarket.Contract.GetLiquidationParameters(&_PerpsMarket.CallOpts, marketId)
}

// GetLiquidationParameters is a free data retrieval call binding the contract method 0xf94363a6.
//
// Solidity: function getLiquidationParameters(uint128 marketId) view returns(uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin)
func (_PerpsMarket *PerpsMarketCallerSession) GetLiquidationParameters(marketId *big.Int) (struct {
	InitialMarginRatioD18        *big.Int
	MinimumInitialMarginRatioD18 *big.Int
	MaintenanceMarginScalarD18   *big.Int
	LiquidationRewardRatioD18    *big.Int
	MinimumPositionMargin        *big.Int
}, error) {
	return _PerpsMarket.Contract.GetLiquidationParameters(&_PerpsMarket.CallOpts, marketId)
}

// GetLiquidationRewardGuards is a free data retrieval call binding the contract method 0x0b7f4b2d.
//
// Solidity: function getLiquidationRewardGuards() view returns(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd)
func (_PerpsMarket *PerpsMarketCaller) GetLiquidationRewardGuards(opts *bind.CallOpts) (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getLiquidationRewardGuards")

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
func (_PerpsMarket *PerpsMarketSession) GetLiquidationRewardGuards() (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	return _PerpsMarket.Contract.GetLiquidationRewardGuards(&_PerpsMarket.CallOpts)
}

// GetLiquidationRewardGuards is a free data retrieval call binding the contract method 0x0b7f4b2d.
//
// Solidity: function getLiquidationRewardGuards() view returns(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd)
func (_PerpsMarket *PerpsMarketCallerSession) GetLiquidationRewardGuards() (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	return _PerpsMarket.Contract.GetLiquidationRewardGuards(&_PerpsMarket.CallOpts)
}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) GetLockedOiRatio(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getLockedOiRatio", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) GetLockedOiRatio(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetLockedOiRatio(&_PerpsMarket.CallOpts, marketId)
}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) GetLockedOiRatio(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetLockedOiRatio(&_PerpsMarket.CallOpts, marketId)
}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_PerpsMarket *PerpsMarketCaller) GetMarketSummary(opts *bind.CallOpts, marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getMarketSummary", marketId)

	if err != nil {
		return *new(IPerpsMarketModuleMarketSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(IPerpsMarketModuleMarketSummary)).(*IPerpsMarketModuleMarketSummary)

	return out0, err

}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_PerpsMarket *PerpsMarketSession) GetMarketSummary(marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	return _PerpsMarket.Contract.GetMarketSummary(&_PerpsMarket.CallOpts, marketId)
}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_PerpsMarket *PerpsMarketCallerSession) GetMarketSummary(marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	return _PerpsMarket.Contract.GetMarketSummary(&_PerpsMarket.CallOpts, marketId)
}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_PerpsMarket *PerpsMarketCaller) GetMarkets(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getMarkets")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_PerpsMarket *PerpsMarketSession) GetMarkets() ([]*big.Int, error) {
	return _PerpsMarket.Contract.GetMarkets(&_PerpsMarket.CallOpts)
}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_PerpsMarket *PerpsMarketCallerSession) GetMarkets() ([]*big.Int, error) {
	return _PerpsMarket.Contract.GetMarkets(&_PerpsMarket.CallOpts)
}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) GetMaxCollateralAmount(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getMaxCollateralAmount", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) GetMaxCollateralAmount(synthMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetMaxCollateralAmount(&_PerpsMarket.CallOpts, synthMarketId)
}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) GetMaxCollateralAmount(synthMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetMaxCollateralAmount(&_PerpsMarket.CallOpts, synthMarketId)
}

// GetMaxLiquidationParameters is a free data retrieval call binding the contract method 0x5443e33e.
//
// Solidity: function getMaxLiquidationParameters(uint128 marketId) view returns(uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator)
func (_PerpsMarket *PerpsMarketCaller) GetMaxLiquidationParameters(opts *bind.CallOpts, marketId *big.Int) (struct {
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MaxLiquidationPd                          *big.Int
	EndorsedLiquidator                        common.Address
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getMaxLiquidationParameters", marketId)

	outstruct := new(struct {
		MaxLiquidationLimitAccumulationMultiplier *big.Int
		MaxSecondsInLiquidationWindow             *big.Int
		MaxLiquidationPd                          *big.Int
		EndorsedLiquidator                        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxLiquidationLimitAccumulationMultiplier = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxSecondsInLiquidationWindow = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxLiquidationPd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndorsedLiquidator = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetMaxLiquidationParameters is a free data retrieval call binding the contract method 0x5443e33e.
//
// Solidity: function getMaxLiquidationParameters(uint128 marketId) view returns(uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator)
func (_PerpsMarket *PerpsMarketSession) GetMaxLiquidationParameters(marketId *big.Int) (struct {
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MaxLiquidationPd                          *big.Int
	EndorsedLiquidator                        common.Address
}, error) {
	return _PerpsMarket.Contract.GetMaxLiquidationParameters(&_PerpsMarket.CallOpts, marketId)
}

// GetMaxLiquidationParameters is a free data retrieval call binding the contract method 0x5443e33e.
//
// Solidity: function getMaxLiquidationParameters(uint128 marketId) view returns(uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator)
func (_PerpsMarket *PerpsMarketCallerSession) GetMaxLiquidationParameters(marketId *big.Int) (struct {
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MaxLiquidationPd                          *big.Int
	EndorsedLiquidator                        common.Address
}, error) {
	return _PerpsMarket.Contract.GetMaxLiquidationParameters(&_PerpsMarket.CallOpts, marketId)
}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_PerpsMarket *PerpsMarketCaller) GetMaxMarketSize(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getMaxMarketSize", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_PerpsMarket *PerpsMarketSession) GetMaxMarketSize(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetMaxMarketSize(&_PerpsMarket.CallOpts, marketId)
}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_PerpsMarket *PerpsMarketCallerSession) GetMaxMarketSize(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetMaxMarketSize(&_PerpsMarket.CallOpts, marketId)
}

// GetOpenPosition is a free data retrieval call binding the contract method 0x22a73967.
//
// Solidity: function getOpenPosition(uint128 accountId, uint128 marketId) view returns(int256 totalPnl, int256 accruedFunding, int128 positionSize)
func (_PerpsMarket *PerpsMarketCaller) GetOpenPosition(opts *bind.CallOpts, accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getOpenPosition", accountId, marketId)

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
func (_PerpsMarket *PerpsMarketSession) GetOpenPosition(accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	return _PerpsMarket.Contract.GetOpenPosition(&_PerpsMarket.CallOpts, accountId, marketId)
}

// GetOpenPosition is a free data retrieval call binding the contract method 0x22a73967.
//
// Solidity: function getOpenPosition(uint128 accountId, uint128 marketId) view returns(int256 totalPnl, int256 accruedFunding, int128 positionSize)
func (_PerpsMarket *PerpsMarketCallerSession) GetOpenPosition(accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	return _PerpsMarket.Contract.GetOpenPosition(&_PerpsMarket.CallOpts, accountId, marketId)
}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_PerpsMarket *PerpsMarketCaller) GetOrder(opts *bind.CallOpts, accountId *big.Int) (AsyncOrderData, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getOrder", accountId)

	if err != nil {
		return *new(AsyncOrderData), err
	}

	out0 := *abi.ConvertType(out[0], new(AsyncOrderData)).(*AsyncOrderData)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_PerpsMarket *PerpsMarketSession) GetOrder(accountId *big.Int) (AsyncOrderData, error) {
	return _PerpsMarket.Contract.GetOrder(&_PerpsMarket.CallOpts, accountId)
}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_PerpsMarket *PerpsMarketCallerSession) GetOrder(accountId *big.Int) (AsyncOrderData, error) {
	return _PerpsMarket.Contract.GetOrder(&_PerpsMarket.CallOpts, accountId)
}

// GetOrderFees is a free data retrieval call binding the contract method 0xaac23e8c.
//
// Solidity: function getOrderFees(uint128 marketId) view returns(uint256 makerFee, uint256 takerFee)
func (_PerpsMarket *PerpsMarketCaller) GetOrderFees(opts *bind.CallOpts, marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getOrderFees", marketId)

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
func (_PerpsMarket *PerpsMarketSession) GetOrderFees(marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	return _PerpsMarket.Contract.GetOrderFees(&_PerpsMarket.CallOpts, marketId)
}

// GetOrderFees is a free data retrieval call binding the contract method 0xaac23e8c.
//
// Solidity: function getOrderFees(uint128 marketId) view returns(uint256 makerFee, uint256 takerFee)
func (_PerpsMarket *PerpsMarketCallerSession) GetOrderFees(marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	return _PerpsMarket.Contract.GetOrderFees(&_PerpsMarket.CallOpts, marketId)
}

// GetPerAccountCaps is a free data retrieval call binding the contract method 0x774f7b07.
//
// Solidity: function getPerAccountCaps() view returns(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount)
func (_PerpsMarket *PerpsMarketCaller) GetPerAccountCaps(opts *bind.CallOpts) (struct {
	MaxPositionsPerAccount   *big.Int
	MaxCollateralsPerAccount *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getPerAccountCaps")

	outstruct := new(struct {
		MaxPositionsPerAccount   *big.Int
		MaxCollateralsPerAccount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxPositionsPerAccount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxCollateralsPerAccount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPerAccountCaps is a free data retrieval call binding the contract method 0x774f7b07.
//
// Solidity: function getPerAccountCaps() view returns(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount)
func (_PerpsMarket *PerpsMarketSession) GetPerAccountCaps() (struct {
	MaxPositionsPerAccount   *big.Int
	MaxCollateralsPerAccount *big.Int
}, error) {
	return _PerpsMarket.Contract.GetPerAccountCaps(&_PerpsMarket.CallOpts)
}

// GetPerAccountCaps is a free data retrieval call binding the contract method 0x774f7b07.
//
// Solidity: function getPerAccountCaps() view returns(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount)
func (_PerpsMarket *PerpsMarketCallerSession) GetPerAccountCaps() (struct {
	MaxPositionsPerAccount   *big.Int
	MaxCollateralsPerAccount *big.Int
}, error) {
	return _PerpsMarket.Contract.GetPerAccountCaps(&_PerpsMarket.CallOpts)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_PerpsMarket *PerpsMarketCaller) GetReferrerShare(opts *bind.CallOpts, referrer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getReferrerShare", referrer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_PerpsMarket *PerpsMarketSession) GetReferrerShare(referrer common.Address) (*big.Int, error) {
	return _PerpsMarket.Contract.GetReferrerShare(&_PerpsMarket.CallOpts, referrer)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_PerpsMarket *PerpsMarketCallerSession) GetReferrerShare(referrer common.Address) (*big.Int, error) {
	return _PerpsMarket.Contract.GetReferrerShare(&_PerpsMarket.CallOpts, referrer)
}

// GetRequiredMargins is a free data retrieval call binding the contract method 0x3c0f0753.
//
// Solidity: function getRequiredMargins(uint128 accountId) view returns(uint256 requiredInitialMargin, uint256 requiredMaintenanceMargin, uint256 totalAccumulatedLiquidationRewards, uint256 maxLiquidationReward)
func (_PerpsMarket *PerpsMarketCaller) GetRequiredMargins(opts *bind.CallOpts, accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getRequiredMargins", accountId)

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
func (_PerpsMarket *PerpsMarketSession) GetRequiredMargins(accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	return _PerpsMarket.Contract.GetRequiredMargins(&_PerpsMarket.CallOpts, accountId)
}

// GetRequiredMargins is a free data retrieval call binding the contract method 0x3c0f0753.
//
// Solidity: function getRequiredMargins(uint128 accountId) view returns(uint256 requiredInitialMargin, uint256 requiredMaintenanceMargin, uint256 totalAccumulatedLiquidationRewards, uint256 maxLiquidationReward)
func (_PerpsMarket *PerpsMarketCallerSession) GetRequiredMargins(accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	return _PerpsMarket.Contract.GetRequiredMargins(&_PerpsMarket.CallOpts, accountId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,bool) settlementStrategy)
func (_PerpsMarket *PerpsMarketCaller) GetSettlementStrategy(opts *bind.CallOpts, marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getSettlementStrategy", marketId, strategyId)

	if err != nil {
		return *new(SettlementStrategyData), err
	}

	out0 := *abi.ConvertType(out[0], new(SettlementStrategyData)).(*SettlementStrategyData)

	return out0, err

}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,bool) settlementStrategy)
func (_PerpsMarket *PerpsMarketSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _PerpsMarket.Contract.GetSettlementStrategy(&_PerpsMarket.CallOpts, marketId, strategyId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,bool) settlementStrategy)
func (_PerpsMarket *PerpsMarketCallerSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _PerpsMarket.Contract.GetSettlementStrategy(&_PerpsMarket.CallOpts, marketId, strategyId)
}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_PerpsMarket *PerpsMarketCaller) GetSynthDeductionPriority(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getSynthDeductionPriority")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_PerpsMarket *PerpsMarketSession) GetSynthDeductionPriority() ([]*big.Int, error) {
	return _PerpsMarket.Contract.GetSynthDeductionPriority(&_PerpsMarket.CallOpts)
}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_PerpsMarket *PerpsMarketCallerSession) GetSynthDeductionPriority() ([]*big.Int, error) {
	return _PerpsMarket.Contract.GetSynthDeductionPriority(&_PerpsMarket.CallOpts)
}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_PerpsMarket *PerpsMarketCaller) GetWithdrawableMargin(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "getWithdrawableMargin", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_PerpsMarket *PerpsMarketSession) GetWithdrawableMargin(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetWithdrawableMargin(&_PerpsMarket.CallOpts, accountId)
}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_PerpsMarket *PerpsMarketCallerSession) GetWithdrawableMargin(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.GetWithdrawableMargin(&_PerpsMarket.CallOpts, accountId)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarket *PerpsMarketCaller) HasPermission(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "hasPermission", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarket *PerpsMarketSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PerpsMarket.Contract.HasPermission(&_PerpsMarket.CallOpts, accountId, permission, user)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarket *PerpsMarketCallerSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PerpsMarket.Contract.HasPermission(&_PerpsMarket.CallOpts, accountId, permission, user)
}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) IndexPrice(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "indexPrice", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) IndexPrice(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.IndexPrice(&_PerpsMarket.CallOpts, marketId)
}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) IndexPrice(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.IndexPrice(&_PerpsMarket.CallOpts, marketId)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarket *PerpsMarketCaller) IsAuthorized(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "isAuthorized", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarket *PerpsMarketSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PerpsMarket.Contract.IsAuthorized(&_PerpsMarket.CallOpts, accountId, permission, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarket *PerpsMarketCallerSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PerpsMarket.Contract.IsAuthorized(&_PerpsMarket.CallOpts, accountId, permission, user)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_PerpsMarket *PerpsMarketCaller) IsFeatureAllowed(opts *bind.CallOpts, feature [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "isFeatureAllowed", feature, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_PerpsMarket *PerpsMarketSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _PerpsMarket.Contract.IsFeatureAllowed(&_PerpsMarket.CallOpts, feature, account)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_PerpsMarket *PerpsMarketCallerSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _PerpsMarket.Contract.IsFeatureAllowed(&_PerpsMarket.CallOpts, feature, account)
}

// LiquidationCapacity is a free data retrieval call binding the contract method 0xbb36f896.
//
// Solidity: function liquidationCapacity(uint128 marketId) view returns(uint256 capacity, uint256 maxLiquidationInWindow, uint256 latestLiquidationTimestamp)
func (_PerpsMarket *PerpsMarketCaller) LiquidationCapacity(opts *bind.CallOpts, marketId *big.Int) (struct {
	Capacity                   *big.Int
	MaxLiquidationInWindow     *big.Int
	LatestLiquidationTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "liquidationCapacity", marketId)

	outstruct := new(struct {
		Capacity                   *big.Int
		MaxLiquidationInWindow     *big.Int
		LatestLiquidationTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Capacity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxLiquidationInWindow = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LatestLiquidationTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LiquidationCapacity is a free data retrieval call binding the contract method 0xbb36f896.
//
// Solidity: function liquidationCapacity(uint128 marketId) view returns(uint256 capacity, uint256 maxLiquidationInWindow, uint256 latestLiquidationTimestamp)
func (_PerpsMarket *PerpsMarketSession) LiquidationCapacity(marketId *big.Int) (struct {
	Capacity                   *big.Int
	MaxLiquidationInWindow     *big.Int
	LatestLiquidationTimestamp *big.Int
}, error) {
	return _PerpsMarket.Contract.LiquidationCapacity(&_PerpsMarket.CallOpts, marketId)
}

// LiquidationCapacity is a free data retrieval call binding the contract method 0xbb36f896.
//
// Solidity: function liquidationCapacity(uint128 marketId) view returns(uint256 capacity, uint256 maxLiquidationInWindow, uint256 latestLiquidationTimestamp)
func (_PerpsMarket *PerpsMarketCallerSession) LiquidationCapacity(marketId *big.Int) (struct {
	Capacity                   *big.Int
	MaxLiquidationInWindow     *big.Int
	LatestLiquidationTimestamp *big.Int
}, error) {
	return _PerpsMarket.Contract.LiquidationCapacity(&_PerpsMarket.CallOpts, marketId)
}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) MaxOpenInterest(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "maxOpenInterest", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) MaxOpenInterest(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.MaxOpenInterest(&_PerpsMarket.CallOpts, marketId)
}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) MaxOpenInterest(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.MaxOpenInterest(&_PerpsMarket.CallOpts, marketId)
}

// Metadata is a free data retrieval call binding the contract method 0xe3bc36bf.
//
// Solidity: function metadata(uint128 marketId) view returns(string name, string symbol)
func (_PerpsMarket *PerpsMarketCaller) Metadata(opts *bind.CallOpts, marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "metadata", marketId)

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
func (_PerpsMarket *PerpsMarketSession) Metadata(marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	return _PerpsMarket.Contract.Metadata(&_PerpsMarket.CallOpts, marketId)
}

// Metadata is a free data retrieval call binding the contract method 0xe3bc36bf.
//
// Solidity: function metadata(uint128 marketId) view returns(string name, string symbol)
func (_PerpsMarket *PerpsMarketCallerSession) Metadata(marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	return _PerpsMarket.Contract.Metadata(&_PerpsMarket.CallOpts, marketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) MinimumCredit(opts *bind.CallOpts, perpsMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "minimumCredit", perpsMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) MinimumCredit(perpsMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.MinimumCredit(&_PerpsMarket.CallOpts, perpsMarketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) MinimumCredit(perpsMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.MinimumCredit(&_PerpsMarket.CallOpts, perpsMarketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_PerpsMarket *PerpsMarketCaller) Name(opts *bind.CallOpts, perpsMarketId *big.Int) (string, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "name", perpsMarketId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_PerpsMarket *PerpsMarketSession) Name(perpsMarketId *big.Int) (string, error) {
	return _PerpsMarket.Contract.Name(&_PerpsMarket.CallOpts, perpsMarketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_PerpsMarket *PerpsMarketCallerSession) Name(perpsMarketId *big.Int) (string, error) {
	return _PerpsMarket.Contract.Name(&_PerpsMarket.CallOpts, perpsMarketId)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_PerpsMarket *PerpsMarketCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_PerpsMarket *PerpsMarketSession) NominatedOwner() (common.Address, error) {
	return _PerpsMarket.Contract.NominatedOwner(&_PerpsMarket.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_PerpsMarket *PerpsMarketCallerSession) NominatedOwner() (common.Address, error) {
	return _PerpsMarket.Contract.NominatedOwner(&_PerpsMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpsMarket *PerpsMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpsMarket *PerpsMarketSession) Owner() (common.Address, error) {
	return _PerpsMarket.Contract.Owner(&_PerpsMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpsMarket *PerpsMarketCallerSession) Owner() (common.Address, error) {
	return _PerpsMarket.Contract.Owner(&_PerpsMarket.CallOpts)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) ReportedDebt(opts *bind.CallOpts, perpsMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "reportedDebt", perpsMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) ReportedDebt(perpsMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.ReportedDebt(&_PerpsMarket.CallOpts, perpsMarketId)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) ReportedDebt(perpsMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.ReportedDebt(&_PerpsMarket.CallOpts, perpsMarketId)
}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_PerpsMarket *PerpsMarketCaller) RequiredMarginForOrder(opts *bind.CallOpts, accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "requiredMarginForOrder", accountId, marketId, sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_PerpsMarket *PerpsMarketSession) RequiredMarginForOrder(accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.RequiredMarginForOrder(&_PerpsMarket.CallOpts, accountId, marketId, sizeDelta)
}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_PerpsMarket *PerpsMarketCallerSession) RequiredMarginForOrder(accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.RequiredMarginForOrder(&_PerpsMarket.CallOpts, accountId, marketId, sizeDelta)
}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_PerpsMarket *PerpsMarketCaller) Settle(opts *bind.CallOpts, accountId *big.Int) error {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "settle", accountId)

	if err != nil {
		return err
	}

	return err

}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_PerpsMarket *PerpsMarketSession) Settle(accountId *big.Int) error {
	return _PerpsMarket.Contract.Settle(&_PerpsMarket.CallOpts, accountId)
}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_PerpsMarket *PerpsMarketCallerSession) Settle(accountId *big.Int) error {
	return _PerpsMarket.Contract.Settle(&_PerpsMarket.CallOpts, accountId)
}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) Size(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "size", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) Size(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.Size(&_PerpsMarket.CallOpts, marketId)
}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) Size(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.Size(&_PerpsMarket.CallOpts, marketId)
}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_PerpsMarket *PerpsMarketCaller) Skew(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "skew", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_PerpsMarket *PerpsMarketSession) Skew(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.Skew(&_PerpsMarket.CallOpts, marketId)
}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_PerpsMarket *PerpsMarketCallerSession) Skew(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.Skew(&_PerpsMarket.CallOpts, marketId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PerpsMarket *PerpsMarketCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PerpsMarket *PerpsMarketSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PerpsMarket.Contract.SupportsInterface(&_PerpsMarket.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PerpsMarket *PerpsMarketCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PerpsMarket.Contract.SupportsInterface(&_PerpsMarket.CallOpts, interfaceId)
}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) TotalAccountOpenInterest(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "totalAccountOpenInterest", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) TotalAccountOpenInterest(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.TotalAccountOpenInterest(&_PerpsMarket.CallOpts, accountId)
}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) TotalAccountOpenInterest(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.TotalAccountOpenInterest(&_PerpsMarket.CallOpts, accountId)
}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCaller) TotalCollateralValue(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "totalCollateralValue", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_PerpsMarket *PerpsMarketSession) TotalCollateralValue(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.TotalCollateralValue(&_PerpsMarket.CallOpts, accountId)
}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_PerpsMarket *PerpsMarketCallerSession) TotalCollateralValue(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarket.Contract.TotalCollateralValue(&_PerpsMarket.CallOpts, accountId)
}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_PerpsMarket *PerpsMarketCaller) TotalGlobalCollateralValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarket.contract.Call(opts, &out, "totalGlobalCollateralValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_PerpsMarket *PerpsMarketSession) TotalGlobalCollateralValue() (*big.Int, error) {
	return _PerpsMarket.Contract.TotalGlobalCollateralValue(&_PerpsMarket.CallOpts)
}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_PerpsMarket *PerpsMarketCallerSession) TotalGlobalCollateralValue() (*big.Int, error) {
	return _PerpsMarket.Contract.TotalGlobalCollateralValue(&_PerpsMarket.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PerpsMarket *PerpsMarketTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PerpsMarket *PerpsMarketSession) AcceptOwnership() (*types.Transaction, error) {
	return _PerpsMarket.Contract.AcceptOwnership(&_PerpsMarket.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PerpsMarket *PerpsMarketTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PerpsMarket.Contract.AcceptOwnership(&_PerpsMarket.TransactOpts)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x6bd5626d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,bool) strategy) returns(uint256 strategyId)
func (_PerpsMarket *PerpsMarketTransactor) AddSettlementStrategy(opts *bind.TransactOpts, marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "addSettlementStrategy", marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x6bd5626d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,bool) strategy) returns(uint256 strategyId)
func (_PerpsMarket *PerpsMarketSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _PerpsMarket.Contract.AddSettlementStrategy(&_PerpsMarket.TransactOpts, marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0x6bd5626d.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,bool) strategy) returns(uint256 strategyId)
func (_PerpsMarket *PerpsMarketTransactorSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _PerpsMarket.Contract.AddSettlementStrategy(&_PerpsMarket.TransactOpts, marketId, strategy)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarket *PerpsMarketTransactor) AddToFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "addToFeatureFlagAllowlist", feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarket *PerpsMarketSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.AddToFeatureFlagAllowlist(&_PerpsMarket.TransactOpts, feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.AddToFeatureFlagAllowlist(&_PerpsMarket.TransactOpts, feature, account)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_PerpsMarket *PerpsMarketTransactor) CommitOrder(opts *bind.TransactOpts, commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "commitOrder", commitment)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_PerpsMarket *PerpsMarketSession) CommitOrder(commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _PerpsMarket.Contract.CommitOrder(&_PerpsMarket.TransactOpts, commitment)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_PerpsMarket *PerpsMarketTransactorSession) CommitOrder(commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _PerpsMarket.Contract.CommitOrder(&_PerpsMarket.TransactOpts, commitment)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_PerpsMarket *PerpsMarketTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_PerpsMarket *PerpsMarketSession) CreateAccount() (*types.Transaction, error) {
	return _PerpsMarket.Contract.CreateAccount(&_PerpsMarket.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_PerpsMarket *PerpsMarketTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _PerpsMarket.Contract.CreateAccount(&_PerpsMarket.TransactOpts)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_PerpsMarket *PerpsMarketTransactor) CreateAccount0(opts *bind.TransactOpts, requestedAccountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "createAccount0", requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_PerpsMarket *PerpsMarketSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.CreateAccount0(&_PerpsMarket.TransactOpts, requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.CreateAccount0(&_PerpsMarket.TransactOpts, requestedAccountId)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_PerpsMarket *PerpsMarketTransactor) CreateMarket(opts *bind.TransactOpts, requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "createMarket", requestedMarketId, marketName, marketSymbol)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_PerpsMarket *PerpsMarketSession) CreateMarket(requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _PerpsMarket.Contract.CreateMarket(&_PerpsMarket.TransactOpts, requestedMarketId, marketName, marketSymbol)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_PerpsMarket *PerpsMarketTransactorSession) CreateMarket(requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _PerpsMarket.Contract.CreateMarket(&_PerpsMarket.TransactOpts, requestedMarketId, marketName, marketSymbol)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarket *PerpsMarketTransactor) GrantPermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "grantPermission", accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarket *PerpsMarketSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.GrantPermission(&_PerpsMarket.TransactOpts, accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.GrantPermission(&_PerpsMarket.TransactOpts, accountId, permission, user)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_PerpsMarket *PerpsMarketTransactor) InitOrUpgradeNft(opts *bind.TransactOpts, id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "initOrUpgradeNft", id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_PerpsMarket *PerpsMarketSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.InitOrUpgradeNft(&_PerpsMarket.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.InitOrUpgradeNft(&_PerpsMarket.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_PerpsMarket *PerpsMarketTransactor) InitOrUpgradeToken(opts *bind.TransactOpts, id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "initOrUpgradeToken", id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_PerpsMarket *PerpsMarketSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.InitOrUpgradeToken(&_PerpsMarket.TransactOpts, id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.InitOrUpgradeToken(&_PerpsMarket.TransactOpts, id, name, symbol, decimals, impl)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_PerpsMarket *PerpsMarketTransactor) InitializeFactory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "initializeFactory")
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_PerpsMarket *PerpsMarketSession) InitializeFactory() (*types.Transaction, error) {
	return _PerpsMarket.Contract.InitializeFactory(&_PerpsMarket.TransactOpts)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_PerpsMarket *PerpsMarketTransactorSession) InitializeFactory() (*types.Transaction, error) {
	return _PerpsMarket.Contract.InitializeFactory(&_PerpsMarket.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns(uint256 liquidationReward)
func (_PerpsMarket *PerpsMarketTransactor) Liquidate(opts *bind.TransactOpts, accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "liquidate", accountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns(uint256 liquidationReward)
func (_PerpsMarket *PerpsMarketSession) Liquidate(accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.Liquidate(&_PerpsMarket.TransactOpts, accountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns(uint256 liquidationReward)
func (_PerpsMarket *PerpsMarketTransactorSession) Liquidate(accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.Liquidate(&_PerpsMarket.TransactOpts, accountId)
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns(uint256 liquidationReward)
func (_PerpsMarket *PerpsMarketTransactor) LiquidateFlagged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "liquidateFlagged")
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns(uint256 liquidationReward)
func (_PerpsMarket *PerpsMarketSession) LiquidateFlagged() (*types.Transaction, error) {
	return _PerpsMarket.Contract.LiquidateFlagged(&_PerpsMarket.TransactOpts)
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns(uint256 liquidationReward)
func (_PerpsMarket *PerpsMarketTransactorSession) LiquidateFlagged() (*types.Transaction, error) {
	return _PerpsMarket.Contract.LiquidateFlagged(&_PerpsMarket.TransactOpts)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_PerpsMarket *PerpsMarketTransactor) ModifyCollateral(opts *bind.TransactOpts, accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "modifyCollateral", accountId, synthMarketId, amountDelta)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_PerpsMarket *PerpsMarketSession) ModifyCollateral(accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.ModifyCollateral(&_PerpsMarket.TransactOpts, accountId, synthMarketId, amountDelta)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) ModifyCollateral(accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.ModifyCollateral(&_PerpsMarket.TransactOpts, accountId, synthMarketId, amountDelta)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_PerpsMarket *PerpsMarketTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_PerpsMarket *PerpsMarketSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _PerpsMarket.Contract.Multicall(&_PerpsMarket.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_PerpsMarket *PerpsMarketTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _PerpsMarket.Contract.Multicall(&_PerpsMarket.TransactOpts, data)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_PerpsMarket *PerpsMarketTransactor) NominateNewOwner(opts *bind.TransactOpts, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "nominateNewOwner", newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_PerpsMarket *PerpsMarketSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.NominateNewOwner(&_PerpsMarket.TransactOpts, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.NominateNewOwner(&_PerpsMarket.TransactOpts, newNominatedOwner)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_PerpsMarket *PerpsMarketTransactor) NotifyAccountTransfer(opts *bind.TransactOpts, to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "notifyAccountTransfer", to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_PerpsMarket *PerpsMarketSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.NotifyAccountTransfer(&_PerpsMarket.TransactOpts, to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.NotifyAccountTransfer(&_PerpsMarket.TransactOpts, to, accountId)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_PerpsMarket *PerpsMarketTransactor) RegisterUnmanagedSystem(opts *bind.TransactOpts, id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "registerUnmanagedSystem", id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_PerpsMarket *PerpsMarketSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.RegisterUnmanagedSystem(&_PerpsMarket.TransactOpts, id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.RegisterUnmanagedSystem(&_PerpsMarket.TransactOpts, id, endpoint)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarket *PerpsMarketTransactor) RemoveFromFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "removeFromFeatureFlagAllowlist", feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarket *PerpsMarketSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.RemoveFromFeatureFlagAllowlist(&_PerpsMarket.TransactOpts, feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.RemoveFromFeatureFlagAllowlist(&_PerpsMarket.TransactOpts, feature, account)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_PerpsMarket *PerpsMarketTransactor) RenounceNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "renounceNomination")
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_PerpsMarket *PerpsMarketSession) RenounceNomination() (*types.Transaction, error) {
	return _PerpsMarket.Contract.RenounceNomination(&_PerpsMarket.TransactOpts)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_PerpsMarket *PerpsMarketTransactorSession) RenounceNomination() (*types.Transaction, error) {
	return _PerpsMarket.Contract.RenounceNomination(&_PerpsMarket.TransactOpts)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_PerpsMarket *PerpsMarketTransactor) RenouncePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "renouncePermission", accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_PerpsMarket *PerpsMarketSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _PerpsMarket.Contract.RenouncePermission(&_PerpsMarket.TransactOpts, accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _PerpsMarket.Contract.RenouncePermission(&_PerpsMarket.TransactOpts, accountId, permission)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarket *PerpsMarketTransactor) RevokePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "revokePermission", accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarket *PerpsMarketSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.RevokePermission(&_PerpsMarket.TransactOpts, accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.RevokePermission(&_PerpsMarket.TransactOpts, accountId, permission, user)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetDeniers(opts *bind.TransactOpts, feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setDeniers", feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_PerpsMarket *PerpsMarketSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetDeniers(&_PerpsMarket.TransactOpts, feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetDeniers(&_PerpsMarket.TransactOpts, feature, deniers)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetFeatureFlagAllowAll(opts *bind.TransactOpts, feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setFeatureFlagAllowAll", feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_PerpsMarket *PerpsMarketSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetFeatureFlagAllowAll(&_PerpsMarket.TransactOpts, feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetFeatureFlagAllowAll(&_PerpsMarket.TransactOpts, feature, allowAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetFeatureFlagDenyAll(opts *bind.TransactOpts, feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setFeatureFlagDenyAll", feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_PerpsMarket *PerpsMarketSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetFeatureFlagDenyAll(&_PerpsMarket.TransactOpts, feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetFeatureFlagDenyAll(&_PerpsMarket.TransactOpts, feature, denyAll)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetFeeCollector(opts *bind.TransactOpts, feeCollector common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setFeeCollector", feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_PerpsMarket *PerpsMarketSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetFeeCollector(&_PerpsMarket.TransactOpts, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetFeeCollector(&_PerpsMarket.TransactOpts, feeCollector)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetFundingParameters(opts *bind.TransactOpts, marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setFundingParameters", marketId, skewScale, maxFundingVelocity)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_PerpsMarket *PerpsMarketSession) SetFundingParameters(marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetFundingParameters(&_PerpsMarket.TransactOpts, marketId, skewScale, maxFundingVelocity)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetFundingParameters(marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetFundingParameters(&_PerpsMarket.TransactOpts, marketId, skewScale, maxFundingVelocity)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0x25e5409e.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetLiquidationParameters(opts *bind.TransactOpts, marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setLiquidationParameters", marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, minimumPositionMargin)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0x25e5409e.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin) returns()
func (_PerpsMarket *PerpsMarketSession) SetLiquidationParameters(marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetLiquidationParameters(&_PerpsMarket.TransactOpts, marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, minimumPositionMargin)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0x25e5409e.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetLiquidationParameters(marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetLiquidationParameters(&_PerpsMarket.TransactOpts, marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, minimumPositionMargin)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetLiquidationRewardGuards(opts *bind.TransactOpts, minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setLiquidationRewardGuards", minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_PerpsMarket *PerpsMarketSession) SetLiquidationRewardGuards(minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetLiquidationRewardGuards(&_PerpsMarket.TransactOpts, minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetLiquidationRewardGuards(minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetLiquidationRewardGuards(&_PerpsMarket.TransactOpts, minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetLockedOiRatio(opts *bind.TransactOpts, marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setLockedOiRatio", marketId, lockedOiRatioD18)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_PerpsMarket *PerpsMarketSession) SetLockedOiRatio(marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetLockedOiRatio(&_PerpsMarket.TransactOpts, marketId, lockedOiRatioD18)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetLockedOiRatio(marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetLockedOiRatio(&_PerpsMarket.TransactOpts, marketId, lockedOiRatioD18)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetMaxCollateralAmount(opts *bind.TransactOpts, synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setMaxCollateralAmount", synthMarketId, collateralAmount)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_PerpsMarket *PerpsMarketSession) SetMaxCollateralAmount(synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetMaxCollateralAmount(&_PerpsMarket.TransactOpts, synthMarketId, collateralAmount)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetMaxCollateralAmount(synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetMaxCollateralAmount(&_PerpsMarket.TransactOpts, synthMarketId, collateralAmount)
}

// SetMaxLiquidationParameters is a paid mutator transaction binding the contract method 0xc7f8a94f.
//
// Solidity: function setMaxLiquidationParameters(uint128 marketId, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetMaxLiquidationParameters(opts *bind.TransactOpts, marketId *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, maxLiquidationPd *big.Int, endorsedLiquidator common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setMaxLiquidationParameters", marketId, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, maxLiquidationPd, endorsedLiquidator)
}

// SetMaxLiquidationParameters is a paid mutator transaction binding the contract method 0xc7f8a94f.
//
// Solidity: function setMaxLiquidationParameters(uint128 marketId, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator) returns()
func (_PerpsMarket *PerpsMarketSession) SetMaxLiquidationParameters(marketId *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, maxLiquidationPd *big.Int, endorsedLiquidator common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetMaxLiquidationParameters(&_PerpsMarket.TransactOpts, marketId, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, maxLiquidationPd, endorsedLiquidator)
}

// SetMaxLiquidationParameters is a paid mutator transaction binding the contract method 0xc7f8a94f.
//
// Solidity: function setMaxLiquidationParameters(uint128 marketId, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetMaxLiquidationParameters(marketId *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, maxLiquidationPd *big.Int, endorsedLiquidator common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetMaxLiquidationParameters(&_PerpsMarket.TransactOpts, marketId, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, maxLiquidationPd, endorsedLiquidator)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetMaxMarketSize(opts *bind.TransactOpts, marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setMaxMarketSize", marketId, maxMarketSize)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_PerpsMarket *PerpsMarketSession) SetMaxMarketSize(marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetMaxMarketSize(&_PerpsMarket.TransactOpts, marketId, maxMarketSize)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetMaxMarketSize(marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetMaxMarketSize(&_PerpsMarket.TransactOpts, marketId, maxMarketSize)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetOrderFees(opts *bind.TransactOpts, marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setOrderFees", marketId, makerFeeRatio, takerFeeRatio)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_PerpsMarket *PerpsMarketSession) SetOrderFees(marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetOrderFees(&_PerpsMarket.TransactOpts, marketId, makerFeeRatio, takerFeeRatio)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetOrderFees(marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetOrderFees(&_PerpsMarket.TransactOpts, marketId, makerFeeRatio, takerFeeRatio)
}

// SetPerAccountCaps is a paid mutator transaction binding the contract method 0xfa0e70a7.
//
// Solidity: function setPerAccountCaps(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetPerAccountCaps(opts *bind.TransactOpts, maxPositionsPerAccount *big.Int, maxCollateralsPerAccount *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setPerAccountCaps", maxPositionsPerAccount, maxCollateralsPerAccount)
}

// SetPerAccountCaps is a paid mutator transaction binding the contract method 0xfa0e70a7.
//
// Solidity: function setPerAccountCaps(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount) returns()
func (_PerpsMarket *PerpsMarketSession) SetPerAccountCaps(maxPositionsPerAccount *big.Int, maxCollateralsPerAccount *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetPerAccountCaps(&_PerpsMarket.TransactOpts, maxPositionsPerAccount, maxCollateralsPerAccount)
}

// SetPerAccountCaps is a paid mutator transaction binding the contract method 0xfa0e70a7.
//
// Solidity: function setPerAccountCaps(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetPerAccountCaps(maxPositionsPerAccount *big.Int, maxCollateralsPerAccount *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetPerAccountCaps(&_PerpsMarket.TransactOpts, maxPositionsPerAccount, maxCollateralsPerAccount)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetSettlementStrategyEnabled(opts *bind.TransactOpts, marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setSettlementStrategyEnabled", marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_PerpsMarket *PerpsMarketSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetSettlementStrategyEnabled(&_PerpsMarket.TransactOpts, marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetSettlementStrategyEnabled(&_PerpsMarket.TransactOpts, marketId, strategyId, enabled)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetSpotMarket(opts *bind.TransactOpts, spotMarket common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setSpotMarket", spotMarket)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_PerpsMarket *PerpsMarketSession) SetSpotMarket(spotMarket common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetSpotMarket(&_PerpsMarket.TransactOpts, spotMarket)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetSpotMarket(spotMarket common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetSpotMarket(&_PerpsMarket.TransactOpts, spotMarket)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetSynthDeductionPriority(opts *bind.TransactOpts, newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setSynthDeductionPriority", newSynthDeductionPriority)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_PerpsMarket *PerpsMarketSession) SetSynthDeductionPriority(newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetSynthDeductionPriority(&_PerpsMarket.TransactOpts, newSynthDeductionPriority)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetSynthDeductionPriority(newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetSynthDeductionPriority(&_PerpsMarket.TransactOpts, newSynthDeductionPriority)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_PerpsMarket *PerpsMarketTransactor) SetSynthetix(opts *bind.TransactOpts, synthetix common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "setSynthetix", synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_PerpsMarket *PerpsMarketSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetSynthetix(&_PerpsMarket.TransactOpts, synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SetSynthetix(&_PerpsMarket.TransactOpts, synthetix)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_PerpsMarket *PerpsMarketTransactor) SettlePythOrder(opts *bind.TransactOpts, result []byte, extraData []byte) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "settlePythOrder", result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_PerpsMarket *PerpsMarketSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SettlePythOrder(&_PerpsMarket.TransactOpts, result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SettlePythOrder(&_PerpsMarket.TransactOpts, result, extraData)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_PerpsMarket *PerpsMarketTransactor) SimulateUpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "simulateUpgradeTo", newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_PerpsMarket *PerpsMarketSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SimulateUpgradeTo(&_PerpsMarket.TransactOpts, newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.SimulateUpgradeTo(&_PerpsMarket.TransactOpts, newImplementation)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_PerpsMarket *PerpsMarketTransactor) UpdatePriceData(opts *bind.TransactOpts, perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "updatePriceData", perpsMarketId, feedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_PerpsMarket *PerpsMarketSession) UpdatePriceData(perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _PerpsMarket.Contract.UpdatePriceData(&_PerpsMarket.TransactOpts, perpsMarketId, feedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) UpdatePriceData(perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _PerpsMarket.Contract.UpdatePriceData(&_PerpsMarket.TransactOpts, perpsMarketId, feedId)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_PerpsMarket *PerpsMarketTransactor) UpdateReferrerShare(opts *bind.TransactOpts, referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "updateReferrerShare", referrer, shareRatioD18)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_PerpsMarket *PerpsMarketSession) UpdateReferrerShare(referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.UpdateReferrerShare(&_PerpsMarket.TransactOpts, referrer, shareRatioD18)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) UpdateReferrerShare(referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarket.Contract.UpdateReferrerShare(&_PerpsMarket.TransactOpts, referrer, shareRatioD18)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PerpsMarket *PerpsMarketTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarket.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PerpsMarket *PerpsMarketSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.UpgradeTo(&_PerpsMarket.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PerpsMarket *PerpsMarketTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarket.Contract.UpgradeTo(&_PerpsMarket.TransactOpts, newImplementation)
}

// PerpsMarketAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the PerpsMarket contract.
type PerpsMarketAccountCreatedIterator struct {
	Event *PerpsMarketAccountCreated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketAccountCreated)
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
		it.Event = new(PerpsMarketAccountCreated)
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
func (it *PerpsMarketAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketAccountCreated represents a AccountCreated event raised by the PerpsMarket contract.
type PerpsMarketAccountCreated struct {
	AccountId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_PerpsMarket *PerpsMarketFilterer) FilterAccountCreated(opts *bind.FilterOpts, accountId []*big.Int, owner []common.Address) (*PerpsMarketAccountCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketAccountCreatedIterator{contract: _PerpsMarket.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_PerpsMarket *PerpsMarketFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *PerpsMarketAccountCreated, accountId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketAccountCreated)
				if err := _PerpsMarket.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseAccountCreated(log types.Log) (*PerpsMarketAccountCreated, error) {
	event := new(PerpsMarketAccountCreated)
	if err := _PerpsMarket.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketAccountLiquidatedIterator is returned from FilterAccountLiquidated and is used to iterate over the raw logs and unpacked data for AccountLiquidated events raised by the PerpsMarket contract.
type PerpsMarketAccountLiquidatedIterator struct {
	Event *PerpsMarketAccountLiquidated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketAccountLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketAccountLiquidated)
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
		it.Event = new(PerpsMarketAccountLiquidated)
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
func (it *PerpsMarketAccountLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketAccountLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketAccountLiquidated represents a AccountLiquidated event raised by the PerpsMarket contract.
type PerpsMarketAccountLiquidated struct {
	AccountId       *big.Int
	Reward          *big.Int
	FullLiquidation bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccountLiquidated is a free log retrieval operation binding the contract event 0x926f56b0a80cc150ad77b59cc60bc53cb907eab441d204924ec67b40ae64f5f5.
//
// Solidity: event AccountLiquidated(uint128 indexed accountId, uint256 reward, bool fullLiquidation)
func (_PerpsMarket *PerpsMarketFilterer) FilterAccountLiquidated(opts *bind.FilterOpts, accountId []*big.Int) (*PerpsMarketAccountLiquidatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "AccountLiquidated", accountIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketAccountLiquidatedIterator{contract: _PerpsMarket.contract, event: "AccountLiquidated", logs: logs, sub: sub}, nil
}

// WatchAccountLiquidated is a free log subscription operation binding the contract event 0x926f56b0a80cc150ad77b59cc60bc53cb907eab441d204924ec67b40ae64f5f5.
//
// Solidity: event AccountLiquidated(uint128 indexed accountId, uint256 reward, bool fullLiquidation)
func (_PerpsMarket *PerpsMarketFilterer) WatchAccountLiquidated(opts *bind.WatchOpts, sink chan<- *PerpsMarketAccountLiquidated, accountId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "AccountLiquidated", accountIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketAccountLiquidated)
				if err := _PerpsMarket.contract.UnpackLog(event, "AccountLiquidated", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseAccountLiquidated(log types.Log) (*PerpsMarketAccountLiquidated, error) {
	event := new(PerpsMarketAccountLiquidated)
	if err := _PerpsMarket.contract.UnpackLog(event, "AccountLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketAssociatedSystemSetIterator is returned from FilterAssociatedSystemSet and is used to iterate over the raw logs and unpacked data for AssociatedSystemSet events raised by the PerpsMarket contract.
type PerpsMarketAssociatedSystemSetIterator struct {
	Event *PerpsMarketAssociatedSystemSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketAssociatedSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketAssociatedSystemSet)
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
		it.Event = new(PerpsMarketAssociatedSystemSet)
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
func (it *PerpsMarketAssociatedSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketAssociatedSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketAssociatedSystemSet represents a AssociatedSystemSet event raised by the PerpsMarket contract.
type PerpsMarketAssociatedSystemSet struct {
	Kind  [32]byte
	Id    [32]byte
	Proxy common.Address
	Impl  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssociatedSystemSet is a free log retrieval operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_PerpsMarket *PerpsMarketFilterer) FilterAssociatedSystemSet(opts *bind.FilterOpts, kind [][32]byte, id [][32]byte) (*PerpsMarketAssociatedSystemSetIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketAssociatedSystemSetIterator{contract: _PerpsMarket.contract, event: "AssociatedSystemSet", logs: logs, sub: sub}, nil
}

// WatchAssociatedSystemSet is a free log subscription operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_PerpsMarket *PerpsMarketFilterer) WatchAssociatedSystemSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketAssociatedSystemSet, kind [][32]byte, id [][32]byte) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketAssociatedSystemSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseAssociatedSystemSet(log types.Log) (*PerpsMarketAssociatedSystemSet, error) {
	event := new(PerpsMarketAssociatedSystemSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketCollateralModifiedIterator is returned from FilterCollateralModified and is used to iterate over the raw logs and unpacked data for CollateralModified events raised by the PerpsMarket contract.
type PerpsMarketCollateralModifiedIterator struct {
	Event *PerpsMarketCollateralModified // Event containing the contract specifics and raw log

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
func (it *PerpsMarketCollateralModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketCollateralModified)
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
		it.Event = new(PerpsMarketCollateralModified)
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
func (it *PerpsMarketCollateralModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketCollateralModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketCollateralModified represents a CollateralModified event raised by the PerpsMarket contract.
type PerpsMarketCollateralModified struct {
	AccountId     *big.Int
	SynthMarketId *big.Int
	AmountDelta   *big.Int
	Sender        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralModified is a free log retrieval operation binding the contract event 0x2e8360c2f3a6fc1a15aefdd0a0922bea3c898cb90d38c3a97354e7c013116064.
//
// Solidity: event CollateralModified(uint128 indexed accountId, uint128 indexed synthMarketId, int256 amountDelta, address indexed sender)
func (_PerpsMarket *PerpsMarketFilterer) FilterCollateralModified(opts *bind.FilterOpts, accountId []*big.Int, synthMarketId []*big.Int, sender []common.Address) (*PerpsMarketCollateralModifiedIterator, error) {

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

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "CollateralModified", accountIdRule, synthMarketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketCollateralModifiedIterator{contract: _PerpsMarket.contract, event: "CollateralModified", logs: logs, sub: sub}, nil
}

// WatchCollateralModified is a free log subscription operation binding the contract event 0x2e8360c2f3a6fc1a15aefdd0a0922bea3c898cb90d38c3a97354e7c013116064.
//
// Solidity: event CollateralModified(uint128 indexed accountId, uint128 indexed synthMarketId, int256 amountDelta, address indexed sender)
func (_PerpsMarket *PerpsMarketFilterer) WatchCollateralModified(opts *bind.WatchOpts, sink chan<- *PerpsMarketCollateralModified, accountId []*big.Int, synthMarketId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "CollateralModified", accountIdRule, synthMarketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketCollateralModified)
				if err := _PerpsMarket.contract.UnpackLog(event, "CollateralModified", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseCollateralModified(log types.Log) (*PerpsMarketCollateralModified, error) {
	event := new(PerpsMarketCollateralModified)
	if err := _PerpsMarket.contract.UnpackLog(event, "CollateralModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketFactoryInitializedIterator is returned from FilterFactoryInitialized and is used to iterate over the raw logs and unpacked data for FactoryInitialized events raised by the PerpsMarket contract.
type PerpsMarketFactoryInitializedIterator struct {
	Event *PerpsMarketFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *PerpsMarketFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketFactoryInitialized)
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
		it.Event = new(PerpsMarketFactoryInitialized)
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
func (it *PerpsMarketFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketFactoryInitialized represents a FactoryInitialized event raised by the PerpsMarket contract.
type PerpsMarketFactoryInitialized struct {
	GlobalPerpsMarketId *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFactoryInitialized is a free log retrieval operation binding the contract event 0xb3240229b07e26f2f02e69dda85ede86e162ccbc6d10e6aade28931e7f533339.
//
// Solidity: event FactoryInitialized(uint128 globalPerpsMarketId)
func (_PerpsMarket *PerpsMarketFilterer) FilterFactoryInitialized(opts *bind.FilterOpts) (*PerpsMarketFactoryInitializedIterator, error) {

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "FactoryInitialized")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketFactoryInitializedIterator{contract: _PerpsMarket.contract, event: "FactoryInitialized", logs: logs, sub: sub}, nil
}

// WatchFactoryInitialized is a free log subscription operation binding the contract event 0xb3240229b07e26f2f02e69dda85ede86e162ccbc6d10e6aade28931e7f533339.
//
// Solidity: event FactoryInitialized(uint128 globalPerpsMarketId)
func (_PerpsMarket *PerpsMarketFilterer) WatchFactoryInitialized(opts *bind.WatchOpts, sink chan<- *PerpsMarketFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "FactoryInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketFactoryInitialized)
				if err := _PerpsMarket.contract.UnpackLog(event, "FactoryInitialized", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseFactoryInitialized(log types.Log) (*PerpsMarketFactoryInitialized, error) {
	event := new(PerpsMarketFactoryInitialized)
	if err := _PerpsMarket.contract.UnpackLog(event, "FactoryInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketFeatureFlagAllowAllSetIterator is returned from FilterFeatureFlagAllowAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowAllSet events raised by the PerpsMarket contract.
type PerpsMarketFeatureFlagAllowAllSetIterator struct {
	Event *PerpsMarketFeatureFlagAllowAllSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketFeatureFlagAllowAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketFeatureFlagAllowAllSet)
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
		it.Event = new(PerpsMarketFeatureFlagAllowAllSet)
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
func (it *PerpsMarketFeatureFlagAllowAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketFeatureFlagAllowAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketFeatureFlagAllowAllSet represents a FeatureFlagAllowAllSet event raised by the PerpsMarket contract.
type PerpsMarketFeatureFlagAllowAllSet struct {
	Feature  [32]byte
	AllowAll bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowAllSet is a free log retrieval operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_PerpsMarket *PerpsMarketFilterer) FilterFeatureFlagAllowAllSet(opts *bind.FilterOpts, feature [][32]byte) (*PerpsMarketFeatureFlagAllowAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketFeatureFlagAllowAllSetIterator{contract: _PerpsMarket.contract, event: "FeatureFlagAllowAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowAllSet is a free log subscription operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_PerpsMarket *PerpsMarketFilterer) WatchFeatureFlagAllowAllSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketFeatureFlagAllowAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketFeatureFlagAllowAllSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseFeatureFlagAllowAllSet(log types.Log) (*PerpsMarketFeatureFlagAllowAllSet, error) {
	event := new(PerpsMarketFeatureFlagAllowAllSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketFeatureFlagAllowlistAddedIterator is returned from FilterFeatureFlagAllowlistAdded and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistAdded events raised by the PerpsMarket contract.
type PerpsMarketFeatureFlagAllowlistAddedIterator struct {
	Event *PerpsMarketFeatureFlagAllowlistAdded // Event containing the contract specifics and raw log

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
func (it *PerpsMarketFeatureFlagAllowlistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketFeatureFlagAllowlistAdded)
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
		it.Event = new(PerpsMarketFeatureFlagAllowlistAdded)
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
func (it *PerpsMarketFeatureFlagAllowlistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketFeatureFlagAllowlistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketFeatureFlagAllowlistAdded represents a FeatureFlagAllowlistAdded event raised by the PerpsMarket contract.
type PerpsMarketFeatureFlagAllowlistAdded struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistAdded is a free log retrieval operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_PerpsMarket *PerpsMarketFilterer) FilterFeatureFlagAllowlistAdded(opts *bind.FilterOpts, feature [][32]byte) (*PerpsMarketFeatureFlagAllowlistAddedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketFeatureFlagAllowlistAddedIterator{contract: _PerpsMarket.contract, event: "FeatureFlagAllowlistAdded", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistAdded is a free log subscription operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_PerpsMarket *PerpsMarketFilterer) WatchFeatureFlagAllowlistAdded(opts *bind.WatchOpts, sink chan<- *PerpsMarketFeatureFlagAllowlistAdded, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketFeatureFlagAllowlistAdded)
				if err := _PerpsMarket.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseFeatureFlagAllowlistAdded(log types.Log) (*PerpsMarketFeatureFlagAllowlistAdded, error) {
	event := new(PerpsMarketFeatureFlagAllowlistAdded)
	if err := _PerpsMarket.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketFeatureFlagAllowlistRemovedIterator is returned from FilterFeatureFlagAllowlistRemoved and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistRemoved events raised by the PerpsMarket contract.
type PerpsMarketFeatureFlagAllowlistRemovedIterator struct {
	Event *PerpsMarketFeatureFlagAllowlistRemoved // Event containing the contract specifics and raw log

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
func (it *PerpsMarketFeatureFlagAllowlistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketFeatureFlagAllowlistRemoved)
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
		it.Event = new(PerpsMarketFeatureFlagAllowlistRemoved)
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
func (it *PerpsMarketFeatureFlagAllowlistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketFeatureFlagAllowlistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketFeatureFlagAllowlistRemoved represents a FeatureFlagAllowlistRemoved event raised by the PerpsMarket contract.
type PerpsMarketFeatureFlagAllowlistRemoved struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistRemoved is a free log retrieval operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_PerpsMarket *PerpsMarketFilterer) FilterFeatureFlagAllowlistRemoved(opts *bind.FilterOpts, feature [][32]byte) (*PerpsMarketFeatureFlagAllowlistRemovedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketFeatureFlagAllowlistRemovedIterator{contract: _PerpsMarket.contract, event: "FeatureFlagAllowlistRemoved", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistRemoved is a free log subscription operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_PerpsMarket *PerpsMarketFilterer) WatchFeatureFlagAllowlistRemoved(opts *bind.WatchOpts, sink chan<- *PerpsMarketFeatureFlagAllowlistRemoved, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketFeatureFlagAllowlistRemoved)
				if err := _PerpsMarket.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseFeatureFlagAllowlistRemoved(log types.Log) (*PerpsMarketFeatureFlagAllowlistRemoved, error) {
	event := new(PerpsMarketFeatureFlagAllowlistRemoved)
	if err := _PerpsMarket.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketFeatureFlagDeniersResetIterator is returned from FilterFeatureFlagDeniersReset and is used to iterate over the raw logs and unpacked data for FeatureFlagDeniersReset events raised by the PerpsMarket contract.
type PerpsMarketFeatureFlagDeniersResetIterator struct {
	Event *PerpsMarketFeatureFlagDeniersReset // Event containing the contract specifics and raw log

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
func (it *PerpsMarketFeatureFlagDeniersResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketFeatureFlagDeniersReset)
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
		it.Event = new(PerpsMarketFeatureFlagDeniersReset)
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
func (it *PerpsMarketFeatureFlagDeniersResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketFeatureFlagDeniersResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketFeatureFlagDeniersReset represents a FeatureFlagDeniersReset event raised by the PerpsMarket contract.
type PerpsMarketFeatureFlagDeniersReset struct {
	Feature [32]byte
	Deniers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDeniersReset is a free log retrieval operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_PerpsMarket *PerpsMarketFilterer) FilterFeatureFlagDeniersReset(opts *bind.FilterOpts, feature [][32]byte) (*PerpsMarketFeatureFlagDeniersResetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketFeatureFlagDeniersResetIterator{contract: _PerpsMarket.contract, event: "FeatureFlagDeniersReset", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDeniersReset is a free log subscription operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_PerpsMarket *PerpsMarketFilterer) WatchFeatureFlagDeniersReset(opts *bind.WatchOpts, sink chan<- *PerpsMarketFeatureFlagDeniersReset, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketFeatureFlagDeniersReset)
				if err := _PerpsMarket.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseFeatureFlagDeniersReset(log types.Log) (*PerpsMarketFeatureFlagDeniersReset, error) {
	event := new(PerpsMarketFeatureFlagDeniersReset)
	if err := _PerpsMarket.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketFeatureFlagDenyAllSetIterator is returned from FilterFeatureFlagDenyAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagDenyAllSet events raised by the PerpsMarket contract.
type PerpsMarketFeatureFlagDenyAllSetIterator struct {
	Event *PerpsMarketFeatureFlagDenyAllSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketFeatureFlagDenyAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketFeatureFlagDenyAllSet)
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
		it.Event = new(PerpsMarketFeatureFlagDenyAllSet)
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
func (it *PerpsMarketFeatureFlagDenyAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketFeatureFlagDenyAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketFeatureFlagDenyAllSet represents a FeatureFlagDenyAllSet event raised by the PerpsMarket contract.
type PerpsMarketFeatureFlagDenyAllSet struct {
	Feature [32]byte
	DenyAll bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDenyAllSet is a free log retrieval operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_PerpsMarket *PerpsMarketFilterer) FilterFeatureFlagDenyAllSet(opts *bind.FilterOpts, feature [][32]byte) (*PerpsMarketFeatureFlagDenyAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketFeatureFlagDenyAllSetIterator{contract: _PerpsMarket.contract, event: "FeatureFlagDenyAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDenyAllSet is a free log subscription operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_PerpsMarket *PerpsMarketFilterer) WatchFeatureFlagDenyAllSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketFeatureFlagDenyAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketFeatureFlagDenyAllSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseFeatureFlagDenyAllSet(log types.Log) (*PerpsMarketFeatureFlagDenyAllSet, error) {
	event := new(PerpsMarketFeatureFlagDenyAllSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketFeeCollectorSetIterator is returned from FilterFeeCollectorSet and is used to iterate over the raw logs and unpacked data for FeeCollectorSet events raised by the PerpsMarket contract.
type PerpsMarketFeeCollectorSetIterator struct {
	Event *PerpsMarketFeeCollectorSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketFeeCollectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketFeeCollectorSet)
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
		it.Event = new(PerpsMarketFeeCollectorSet)
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
func (it *PerpsMarketFeeCollectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketFeeCollectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketFeeCollectorSet represents a FeeCollectorSet event raised by the PerpsMarket contract.
type PerpsMarketFeeCollectorSet struct {
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorSet is a free log retrieval operation binding the contract event 0x12e1d17016b94668449f97876f4a8d5cc2c19f314db337418894734037cc19d4.
//
// Solidity: event FeeCollectorSet(address feeCollector)
func (_PerpsMarket *PerpsMarketFilterer) FilterFeeCollectorSet(opts *bind.FilterOpts) (*PerpsMarketFeeCollectorSetIterator, error) {

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "FeeCollectorSet")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketFeeCollectorSetIterator{contract: _PerpsMarket.contract, event: "FeeCollectorSet", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorSet is a free log subscription operation binding the contract event 0x12e1d17016b94668449f97876f4a8d5cc2c19f314db337418894734037cc19d4.
//
// Solidity: event FeeCollectorSet(address feeCollector)
func (_PerpsMarket *PerpsMarketFilterer) WatchFeeCollectorSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketFeeCollectorSet) (event.Subscription, error) {

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "FeeCollectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketFeeCollectorSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseFeeCollectorSet(log types.Log) (*PerpsMarketFeeCollectorSet, error) {
	event := new(PerpsMarketFeeCollectorSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketFundingParametersSetIterator is returned from FilterFundingParametersSet and is used to iterate over the raw logs and unpacked data for FundingParametersSet events raised by the PerpsMarket contract.
type PerpsMarketFundingParametersSetIterator struct {
	Event *PerpsMarketFundingParametersSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketFundingParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketFundingParametersSet)
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
		it.Event = new(PerpsMarketFundingParametersSet)
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
func (it *PerpsMarketFundingParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketFundingParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketFundingParametersSet represents a FundingParametersSet event raised by the PerpsMarket contract.
type PerpsMarketFundingParametersSet struct {
	MarketId           *big.Int
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFundingParametersSet is a free log retrieval operation binding the contract event 0xa74afd926bbafbb9252d224a1fcd6a209f851324bd485f556786820a76e31b65.
//
// Solidity: event FundingParametersSet(uint128 indexed marketId, uint256 skewScale, uint256 maxFundingVelocity)
func (_PerpsMarket *PerpsMarketFilterer) FilterFundingParametersSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketFundingParametersSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "FundingParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketFundingParametersSetIterator{contract: _PerpsMarket.contract, event: "FundingParametersSet", logs: logs, sub: sub}, nil
}

// WatchFundingParametersSet is a free log subscription operation binding the contract event 0xa74afd926bbafbb9252d224a1fcd6a209f851324bd485f556786820a76e31b65.
//
// Solidity: event FundingParametersSet(uint128 indexed marketId, uint256 skewScale, uint256 maxFundingVelocity)
func (_PerpsMarket *PerpsMarketFilterer) WatchFundingParametersSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketFundingParametersSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "FundingParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketFundingParametersSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "FundingParametersSet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseFundingParametersSet(log types.Log) (*PerpsMarketFundingParametersSet, error) {
	event := new(PerpsMarketFundingParametersSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "FundingParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketLiquidationParametersSetIterator is returned from FilterLiquidationParametersSet and is used to iterate over the raw logs and unpacked data for LiquidationParametersSet events raised by the PerpsMarket contract.
type PerpsMarketLiquidationParametersSetIterator struct {
	Event *PerpsMarketLiquidationParametersSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketLiquidationParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketLiquidationParametersSet)
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
		it.Event = new(PerpsMarketLiquidationParametersSet)
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
func (it *PerpsMarketLiquidationParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketLiquidationParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketLiquidationParametersSet represents a LiquidationParametersSet event raised by the PerpsMarket contract.
type PerpsMarketLiquidationParametersSet struct {
	MarketId                     *big.Int
	InitialMarginRatioD18        *big.Int
	MaintenanceMarginRatioD18    *big.Int
	MinimumInitialMarginRatioD18 *big.Int
	LiquidationRewardRatioD18    *big.Int
	MinimumPositionMargin        *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterLiquidationParametersSet is a free log retrieval operation binding the contract event 0xa0c87f048ec4f5924e50d554aa4a6e65a935f133a2114e5222590c1690e1a7b8.
//
// Solidity: event LiquidationParametersSet(uint128 indexed marketId, uint256 initialMarginRatioD18, uint256 maintenanceMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin)
func (_PerpsMarket *PerpsMarketFilterer) FilterLiquidationParametersSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketLiquidationParametersSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "LiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketLiquidationParametersSetIterator{contract: _PerpsMarket.contract, event: "LiquidationParametersSet", logs: logs, sub: sub}, nil
}

// WatchLiquidationParametersSet is a free log subscription operation binding the contract event 0xa0c87f048ec4f5924e50d554aa4a6e65a935f133a2114e5222590c1690e1a7b8.
//
// Solidity: event LiquidationParametersSet(uint128 indexed marketId, uint256 initialMarginRatioD18, uint256 maintenanceMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin)
func (_PerpsMarket *PerpsMarketFilterer) WatchLiquidationParametersSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketLiquidationParametersSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "LiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketLiquidationParametersSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "LiquidationParametersSet", log); err != nil {
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

// ParseLiquidationParametersSet is a log parse operation binding the contract event 0xa0c87f048ec4f5924e50d554aa4a6e65a935f133a2114e5222590c1690e1a7b8.
//
// Solidity: event LiquidationParametersSet(uint128 indexed marketId, uint256 initialMarginRatioD18, uint256 maintenanceMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin)
func (_PerpsMarket *PerpsMarketFilterer) ParseLiquidationParametersSet(log types.Log) (*PerpsMarketLiquidationParametersSet, error) {
	event := new(PerpsMarketLiquidationParametersSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "LiquidationParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketLiquidationRewardGuardsSetIterator is returned from FilterLiquidationRewardGuardsSet and is used to iterate over the raw logs and unpacked data for LiquidationRewardGuardsSet events raised by the PerpsMarket contract.
type PerpsMarketLiquidationRewardGuardsSetIterator struct {
	Event *PerpsMarketLiquidationRewardGuardsSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketLiquidationRewardGuardsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketLiquidationRewardGuardsSet)
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
		it.Event = new(PerpsMarketLiquidationRewardGuardsSet)
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
func (it *PerpsMarketLiquidationRewardGuardsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketLiquidationRewardGuardsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketLiquidationRewardGuardsSet represents a LiquidationRewardGuardsSet event raised by the PerpsMarket contract.
type PerpsMarketLiquidationRewardGuardsSet struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterLiquidationRewardGuardsSet is a free log retrieval operation binding the contract event 0xbb7f8eb17b00549c0f234a3c363ef3eb39329f98b52048416480cf62d0f0e1f5.
//
// Solidity: event LiquidationRewardGuardsSet(uint256 indexed minLiquidationRewardUsd, uint256 indexed maxLiquidationRewardUsd)
func (_PerpsMarket *PerpsMarketFilterer) FilterLiquidationRewardGuardsSet(opts *bind.FilterOpts, minLiquidationRewardUsd []*big.Int, maxLiquidationRewardUsd []*big.Int) (*PerpsMarketLiquidationRewardGuardsSetIterator, error) {

	var minLiquidationRewardUsdRule []interface{}
	for _, minLiquidationRewardUsdItem := range minLiquidationRewardUsd {
		minLiquidationRewardUsdRule = append(minLiquidationRewardUsdRule, minLiquidationRewardUsdItem)
	}
	var maxLiquidationRewardUsdRule []interface{}
	for _, maxLiquidationRewardUsdItem := range maxLiquidationRewardUsd {
		maxLiquidationRewardUsdRule = append(maxLiquidationRewardUsdRule, maxLiquidationRewardUsdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "LiquidationRewardGuardsSet", minLiquidationRewardUsdRule, maxLiquidationRewardUsdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketLiquidationRewardGuardsSetIterator{contract: _PerpsMarket.contract, event: "LiquidationRewardGuardsSet", logs: logs, sub: sub}, nil
}

// WatchLiquidationRewardGuardsSet is a free log subscription operation binding the contract event 0xbb7f8eb17b00549c0f234a3c363ef3eb39329f98b52048416480cf62d0f0e1f5.
//
// Solidity: event LiquidationRewardGuardsSet(uint256 indexed minLiquidationRewardUsd, uint256 indexed maxLiquidationRewardUsd)
func (_PerpsMarket *PerpsMarketFilterer) WatchLiquidationRewardGuardsSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketLiquidationRewardGuardsSet, minLiquidationRewardUsd []*big.Int, maxLiquidationRewardUsd []*big.Int) (event.Subscription, error) {

	var minLiquidationRewardUsdRule []interface{}
	for _, minLiquidationRewardUsdItem := range minLiquidationRewardUsd {
		minLiquidationRewardUsdRule = append(minLiquidationRewardUsdRule, minLiquidationRewardUsdItem)
	}
	var maxLiquidationRewardUsdRule []interface{}
	for _, maxLiquidationRewardUsdItem := range maxLiquidationRewardUsd {
		maxLiquidationRewardUsdRule = append(maxLiquidationRewardUsdRule, maxLiquidationRewardUsdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "LiquidationRewardGuardsSet", minLiquidationRewardUsdRule, maxLiquidationRewardUsdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketLiquidationRewardGuardsSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "LiquidationRewardGuardsSet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseLiquidationRewardGuardsSet(log types.Log) (*PerpsMarketLiquidationRewardGuardsSet, error) {
	event := new(PerpsMarketLiquidationRewardGuardsSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "LiquidationRewardGuardsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketLockedOiRatioSetIterator is returned from FilterLockedOiRatioSet and is used to iterate over the raw logs and unpacked data for LockedOiRatioSet events raised by the PerpsMarket contract.
type PerpsMarketLockedOiRatioSetIterator struct {
	Event *PerpsMarketLockedOiRatioSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketLockedOiRatioSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketLockedOiRatioSet)
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
		it.Event = new(PerpsMarketLockedOiRatioSet)
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
func (it *PerpsMarketLockedOiRatioSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketLockedOiRatioSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketLockedOiRatioSet represents a LockedOiRatioSet event raised by the PerpsMarket contract.
type PerpsMarketLockedOiRatioSet struct {
	MarketId         *big.Int
	LockedOiRatioD18 *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLockedOiRatioSet is a free log retrieval operation binding the contract event 0x1d841fd5b4c806bc5a073d637a8506e1e74d16cb18251b711cb47e133ceafc2d.
//
// Solidity: event LockedOiRatioSet(uint128 indexed marketId, uint256 lockedOiRatioD18)
func (_PerpsMarket *PerpsMarketFilterer) FilterLockedOiRatioSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketLockedOiRatioSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "LockedOiRatioSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketLockedOiRatioSetIterator{contract: _PerpsMarket.contract, event: "LockedOiRatioSet", logs: logs, sub: sub}, nil
}

// WatchLockedOiRatioSet is a free log subscription operation binding the contract event 0x1d841fd5b4c806bc5a073d637a8506e1e74d16cb18251b711cb47e133ceafc2d.
//
// Solidity: event LockedOiRatioSet(uint128 indexed marketId, uint256 lockedOiRatioD18)
func (_PerpsMarket *PerpsMarketFilterer) WatchLockedOiRatioSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketLockedOiRatioSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "LockedOiRatioSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketLockedOiRatioSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "LockedOiRatioSet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseLockedOiRatioSet(log types.Log) (*PerpsMarketLockedOiRatioSet, error) {
	event := new(PerpsMarketLockedOiRatioSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "LockedOiRatioSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the PerpsMarket contract.
type PerpsMarketMarketCreatedIterator struct {
	Event *PerpsMarketMarketCreated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketMarketCreated)
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
		it.Event = new(PerpsMarketMarketCreated)
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
func (it *PerpsMarketMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketMarketCreated represents a MarketCreated event raised by the PerpsMarket contract.
type PerpsMarketMarketCreated struct {
	PerpsMarketId *big.Int
	MarketName    string
	MarketSymbol  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0x032553f94ac1323933f22650ec5b8e232babf1c47efca69383b749463116cc49.
//
// Solidity: event MarketCreated(uint128 indexed perpsMarketId, string marketName, string marketSymbol)
func (_PerpsMarket *PerpsMarketFilterer) FilterMarketCreated(opts *bind.FilterOpts, perpsMarketId []*big.Int) (*PerpsMarketMarketCreatedIterator, error) {

	var perpsMarketIdRule []interface{}
	for _, perpsMarketIdItem := range perpsMarketId {
		perpsMarketIdRule = append(perpsMarketIdRule, perpsMarketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "MarketCreated", perpsMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketMarketCreatedIterator{contract: _PerpsMarket.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0x032553f94ac1323933f22650ec5b8e232babf1c47efca69383b749463116cc49.
//
// Solidity: event MarketCreated(uint128 indexed perpsMarketId, string marketName, string marketSymbol)
func (_PerpsMarket *PerpsMarketFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *PerpsMarketMarketCreated, perpsMarketId []*big.Int) (event.Subscription, error) {

	var perpsMarketIdRule []interface{}
	for _, perpsMarketIdItem := range perpsMarketId {
		perpsMarketIdRule = append(perpsMarketIdRule, perpsMarketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "MarketCreated", perpsMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketMarketCreated)
				if err := _PerpsMarket.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseMarketCreated(log types.Log) (*PerpsMarketMarketCreated, error) {
	event := new(PerpsMarketMarketCreated)
	if err := _PerpsMarket.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketMarketPriceDataUpdatedIterator is returned from FilterMarketPriceDataUpdated and is used to iterate over the raw logs and unpacked data for MarketPriceDataUpdated events raised by the PerpsMarket contract.
type PerpsMarketMarketPriceDataUpdatedIterator struct {
	Event *PerpsMarketMarketPriceDataUpdated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketMarketPriceDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketMarketPriceDataUpdated)
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
		it.Event = new(PerpsMarketMarketPriceDataUpdated)
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
func (it *PerpsMarketMarketPriceDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketMarketPriceDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketMarketPriceDataUpdated represents a MarketPriceDataUpdated event raised by the PerpsMarket contract.
type PerpsMarketMarketPriceDataUpdated struct {
	MarketId *big.Int
	FeedId   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketPriceDataUpdated is a free log retrieval operation binding the contract event 0xd5ec22bbdbde803d8585a5aa89d991c67dc62c6a7bc3c4fa000b4e263b6783a0.
//
// Solidity: event MarketPriceDataUpdated(uint128 indexed marketId, bytes32 feedId)
func (_PerpsMarket *PerpsMarketFilterer) FilterMarketPriceDataUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketMarketPriceDataUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "MarketPriceDataUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketMarketPriceDataUpdatedIterator{contract: _PerpsMarket.contract, event: "MarketPriceDataUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketPriceDataUpdated is a free log subscription operation binding the contract event 0xd5ec22bbdbde803d8585a5aa89d991c67dc62c6a7bc3c4fa000b4e263b6783a0.
//
// Solidity: event MarketPriceDataUpdated(uint128 indexed marketId, bytes32 feedId)
func (_PerpsMarket *PerpsMarketFilterer) WatchMarketPriceDataUpdated(opts *bind.WatchOpts, sink chan<- *PerpsMarketMarketPriceDataUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "MarketPriceDataUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketMarketPriceDataUpdated)
				if err := _PerpsMarket.contract.UnpackLog(event, "MarketPriceDataUpdated", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseMarketPriceDataUpdated(log types.Log) (*PerpsMarketMarketPriceDataUpdated, error) {
	event := new(PerpsMarketMarketPriceDataUpdated)
	if err := _PerpsMarket.contract.UnpackLog(event, "MarketPriceDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketMarketUpdatedIterator is returned from FilterMarketUpdated and is used to iterate over the raw logs and unpacked data for MarketUpdated events raised by the PerpsMarket contract.
type PerpsMarketMarketUpdatedIterator struct {
	Event *PerpsMarketMarketUpdated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketMarketUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketMarketUpdated)
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
		it.Event = new(PerpsMarketMarketUpdated)
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
func (it *PerpsMarketMarketUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketMarketUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketMarketUpdated represents a MarketUpdated event raised by the PerpsMarket contract.
type PerpsMarketMarketUpdated struct {
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
func (_PerpsMarket *PerpsMarketFilterer) FilterMarketUpdated(opts *bind.FilterOpts) (*PerpsMarketMarketUpdatedIterator, error) {

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "MarketUpdated")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketMarketUpdatedIterator{contract: _PerpsMarket.contract, event: "MarketUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketUpdated is a free log subscription operation binding the contract event 0xb317f068f30db823aeb6ac6ddbffcbb6c805f558972e6b16625ec58b83f1f3d5.
//
// Solidity: event MarketUpdated(uint128 marketId, uint256 price, int256 skew, uint256 size, int256 sizeDelta, int256 currentFundingRate, int256 currentFundingVelocity)
func (_PerpsMarket *PerpsMarketFilterer) WatchMarketUpdated(opts *bind.WatchOpts, sink chan<- *PerpsMarketMarketUpdated) (event.Subscription, error) {

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "MarketUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketMarketUpdated)
				if err := _PerpsMarket.contract.UnpackLog(event, "MarketUpdated", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseMarketUpdated(log types.Log) (*PerpsMarketMarketUpdated, error) {
	event := new(PerpsMarketMarketUpdated)
	if err := _PerpsMarket.contract.UnpackLog(event, "MarketUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketMaxCollateralAmountSetIterator is returned from FilterMaxCollateralAmountSet and is used to iterate over the raw logs and unpacked data for MaxCollateralAmountSet events raised by the PerpsMarket contract.
type PerpsMarketMaxCollateralAmountSetIterator struct {
	Event *PerpsMarketMaxCollateralAmountSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketMaxCollateralAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketMaxCollateralAmountSet)
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
		it.Event = new(PerpsMarketMaxCollateralAmountSet)
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
func (it *PerpsMarketMaxCollateralAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketMaxCollateralAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketMaxCollateralAmountSet represents a MaxCollateralAmountSet event raised by the PerpsMarket contract.
type PerpsMarketMaxCollateralAmountSet struct {
	SynthMarketId    *big.Int
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMaxCollateralAmountSet is a free log retrieval operation binding the contract event 0xa552ba0f2552ec13b62b479a41e6e7dfcef9f10aea18527e5a3e1b3963bf70d7.
//
// Solidity: event MaxCollateralAmountSet(uint128 indexed synthMarketId, uint256 collateralAmount)
func (_PerpsMarket *PerpsMarketFilterer) FilterMaxCollateralAmountSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*PerpsMarketMaxCollateralAmountSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "MaxCollateralAmountSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketMaxCollateralAmountSetIterator{contract: _PerpsMarket.contract, event: "MaxCollateralAmountSet", logs: logs, sub: sub}, nil
}

// WatchMaxCollateralAmountSet is a free log subscription operation binding the contract event 0xa552ba0f2552ec13b62b479a41e6e7dfcef9f10aea18527e5a3e1b3963bf70d7.
//
// Solidity: event MaxCollateralAmountSet(uint128 indexed synthMarketId, uint256 collateralAmount)
func (_PerpsMarket *PerpsMarketFilterer) WatchMaxCollateralAmountSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketMaxCollateralAmountSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "MaxCollateralAmountSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketMaxCollateralAmountSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "MaxCollateralAmountSet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseMaxCollateralAmountSet(log types.Log) (*PerpsMarketMaxCollateralAmountSet, error) {
	event := new(PerpsMarketMaxCollateralAmountSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "MaxCollateralAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketMaxLiquidationParametersSetIterator is returned from FilterMaxLiquidationParametersSet and is used to iterate over the raw logs and unpacked data for MaxLiquidationParametersSet events raised by the PerpsMarket contract.
type PerpsMarketMaxLiquidationParametersSetIterator struct {
	Event *PerpsMarketMaxLiquidationParametersSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketMaxLiquidationParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketMaxLiquidationParametersSet)
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
		it.Event = new(PerpsMarketMaxLiquidationParametersSet)
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
func (it *PerpsMarketMaxLiquidationParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketMaxLiquidationParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketMaxLiquidationParametersSet represents a MaxLiquidationParametersSet event raised by the PerpsMarket contract.
type PerpsMarketMaxLiquidationParametersSet struct {
	MarketId                                  *big.Int
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MaxLiquidationPd                          *big.Int
	EndorsedLiquidator                        common.Address
	Raw                                       types.Log // Blockchain specific contextual infos
}

// FilterMaxLiquidationParametersSet is a free log retrieval operation binding the contract event 0x9012ce377b7043d153d2cba3376efe7e1742af5acb7e38897362f392a7dc89ed.
//
// Solidity: event MaxLiquidationParametersSet(uint128 indexed marketId, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator)
func (_PerpsMarket *PerpsMarketFilterer) FilterMaxLiquidationParametersSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketMaxLiquidationParametersSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "MaxLiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketMaxLiquidationParametersSetIterator{contract: _PerpsMarket.contract, event: "MaxLiquidationParametersSet", logs: logs, sub: sub}, nil
}

// WatchMaxLiquidationParametersSet is a free log subscription operation binding the contract event 0x9012ce377b7043d153d2cba3376efe7e1742af5acb7e38897362f392a7dc89ed.
//
// Solidity: event MaxLiquidationParametersSet(uint128 indexed marketId, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator)
func (_PerpsMarket *PerpsMarketFilterer) WatchMaxLiquidationParametersSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketMaxLiquidationParametersSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "MaxLiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketMaxLiquidationParametersSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "MaxLiquidationParametersSet", log); err != nil {
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

// ParseMaxLiquidationParametersSet is a log parse operation binding the contract event 0x9012ce377b7043d153d2cba3376efe7e1742af5acb7e38897362f392a7dc89ed.
//
// Solidity: event MaxLiquidationParametersSet(uint128 indexed marketId, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator)
func (_PerpsMarket *PerpsMarketFilterer) ParseMaxLiquidationParametersSet(log types.Log) (*PerpsMarketMaxLiquidationParametersSet, error) {
	event := new(PerpsMarketMaxLiquidationParametersSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "MaxLiquidationParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketMaxMarketSizeSetIterator is returned from FilterMaxMarketSizeSet and is used to iterate over the raw logs and unpacked data for MaxMarketSizeSet events raised by the PerpsMarket contract.
type PerpsMarketMaxMarketSizeSetIterator struct {
	Event *PerpsMarketMaxMarketSizeSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketMaxMarketSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketMaxMarketSizeSet)
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
		it.Event = new(PerpsMarketMaxMarketSizeSet)
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
func (it *PerpsMarketMaxMarketSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketMaxMarketSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketMaxMarketSizeSet represents a MaxMarketSizeSet event raised by the PerpsMarket contract.
type PerpsMarketMaxMarketSizeSet struct {
	MarketId      *big.Int
	MaxMarketSize *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxMarketSizeSet is a free log retrieval operation binding the contract event 0xbd063bd3072a194b255163ab8dfd3400c4ab1cc641b920e7aaf1c11f92cd26cf.
//
// Solidity: event MaxMarketSizeSet(uint128 indexed marketId, uint256 maxMarketSize)
func (_PerpsMarket *PerpsMarketFilterer) FilterMaxMarketSizeSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketMaxMarketSizeSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "MaxMarketSizeSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketMaxMarketSizeSetIterator{contract: _PerpsMarket.contract, event: "MaxMarketSizeSet", logs: logs, sub: sub}, nil
}

// WatchMaxMarketSizeSet is a free log subscription operation binding the contract event 0xbd063bd3072a194b255163ab8dfd3400c4ab1cc641b920e7aaf1c11f92cd26cf.
//
// Solidity: event MaxMarketSizeSet(uint128 indexed marketId, uint256 maxMarketSize)
func (_PerpsMarket *PerpsMarketFilterer) WatchMaxMarketSizeSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketMaxMarketSizeSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "MaxMarketSizeSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketMaxMarketSizeSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "MaxMarketSizeSet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseMaxMarketSizeSet(log types.Log) (*PerpsMarketMaxMarketSizeSet, error) {
	event := new(PerpsMarketMaxMarketSizeSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "MaxMarketSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketOrderCommittedIterator is returned from FilterOrderCommitted and is used to iterate over the raw logs and unpacked data for OrderCommitted events raised by the PerpsMarket contract.
type PerpsMarketOrderCommittedIterator struct {
	Event *PerpsMarketOrderCommitted // Event containing the contract specifics and raw log

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
func (it *PerpsMarketOrderCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketOrderCommitted)
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
		it.Event = new(PerpsMarketOrderCommitted)
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
func (it *PerpsMarketOrderCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketOrderCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketOrderCommitted represents a OrderCommitted event raised by the PerpsMarket contract.
type PerpsMarketOrderCommitted struct {
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
func (_PerpsMarket *PerpsMarketFilterer) FilterOrderCommitted(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*PerpsMarketOrderCommittedIterator, error) {

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

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "OrderCommitted", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketOrderCommittedIterator{contract: _PerpsMarket.contract, event: "OrderCommitted", logs: logs, sub: sub}, nil
}

// WatchOrderCommitted is a free log subscription operation binding the contract event 0xc74e8cf044fc959dcef2695245eb90f7e9fb5676532a71eec5b3c513d4793709.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint128 indexed accountId, uint8 orderType, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, uint256 expirationTime, bytes32 indexed trackingCode, address sender)
func (_PerpsMarket *PerpsMarketFilterer) WatchOrderCommitted(opts *bind.WatchOpts, sink chan<- *PerpsMarketOrderCommitted, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "OrderCommitted", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketOrderCommitted)
				if err := _PerpsMarket.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseOrderCommitted(log types.Log) (*PerpsMarketOrderCommitted, error) {
	event := new(PerpsMarketOrderCommitted)
	if err := _PerpsMarket.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketOrderFeesSetIterator is returned from FilterOrderFeesSet and is used to iterate over the raw logs and unpacked data for OrderFeesSet events raised by the PerpsMarket contract.
type PerpsMarketOrderFeesSetIterator struct {
	Event *PerpsMarketOrderFeesSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketOrderFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketOrderFeesSet)
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
		it.Event = new(PerpsMarketOrderFeesSet)
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
func (it *PerpsMarketOrderFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketOrderFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketOrderFeesSet represents a OrderFeesSet event raised by the PerpsMarket contract.
type PerpsMarketOrderFeesSet struct {
	MarketId      *big.Int
	MakerFeeRatio *big.Int
	TakerFeeRatio *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFeesSet is a free log retrieval operation binding the contract event 0x28969f156340ba9eb31589904eb174d3a4b6a37781fa6f7ad289d349d75dd1ee.
//
// Solidity: event OrderFeesSet(uint128 indexed marketId, uint256 makerFeeRatio, uint256 takerFeeRatio)
func (_PerpsMarket *PerpsMarketFilterer) FilterOrderFeesSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketOrderFeesSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "OrderFeesSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketOrderFeesSetIterator{contract: _PerpsMarket.contract, event: "OrderFeesSet", logs: logs, sub: sub}, nil
}

// WatchOrderFeesSet is a free log subscription operation binding the contract event 0x28969f156340ba9eb31589904eb174d3a4b6a37781fa6f7ad289d349d75dd1ee.
//
// Solidity: event OrderFeesSet(uint128 indexed marketId, uint256 makerFeeRatio, uint256 takerFeeRatio)
func (_PerpsMarket *PerpsMarketFilterer) WatchOrderFeesSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketOrderFeesSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "OrderFeesSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketOrderFeesSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "OrderFeesSet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseOrderFeesSet(log types.Log) (*PerpsMarketOrderFeesSet, error) {
	event := new(PerpsMarketOrderFeesSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "OrderFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketOrderSettledIterator is returned from FilterOrderSettled and is used to iterate over the raw logs and unpacked data for OrderSettled events raised by the PerpsMarket contract.
type PerpsMarketOrderSettledIterator struct {
	Event *PerpsMarketOrderSettled // Event containing the contract specifics and raw log

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
func (it *PerpsMarketOrderSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketOrderSettled)
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
		it.Event = new(PerpsMarketOrderSettled)
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
func (it *PerpsMarketOrderSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketOrderSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketOrderSettled represents a OrderSettled event raised by the PerpsMarket contract.
type PerpsMarketOrderSettled struct {
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
func (_PerpsMarket *PerpsMarketFilterer) FilterOrderSettled(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*PerpsMarketOrderSettledIterator, error) {

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

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "OrderSettled", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketOrderSettledIterator{contract: _PerpsMarket.contract, event: "OrderSettled", logs: logs, sub: sub}, nil
}

// WatchOrderSettled is a free log subscription operation binding the contract event 0x460080a757ec90719fe90ab2384c0196cdeed071a9fd7ce1ada43481d96b7db5.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed accountId, uint256 fillPrice, int256 pnl, int256 accruedFunding, int128 sizeDelta, int128 newSize, uint256 totalFees, uint256 referralFees, uint256 collectedFees, uint256 settlementReward, bytes32 indexed trackingCode, address settler)
func (_PerpsMarket *PerpsMarketFilterer) WatchOrderSettled(opts *bind.WatchOpts, sink chan<- *PerpsMarketOrderSettled, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "OrderSettled", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketOrderSettled)
				if err := _PerpsMarket.contract.UnpackLog(event, "OrderSettled", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseOrderSettled(log types.Log) (*PerpsMarketOrderSettled, error) {
	event := new(PerpsMarketOrderSettled)
	if err := _PerpsMarket.contract.UnpackLog(event, "OrderSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the PerpsMarket contract.
type PerpsMarketOwnerChangedIterator struct {
	Event *PerpsMarketOwnerChanged // Event containing the contract specifics and raw log

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
func (it *PerpsMarketOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketOwnerChanged)
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
		it.Event = new(PerpsMarketOwnerChanged)
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
func (it *PerpsMarketOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketOwnerChanged represents a OwnerChanged event raised by the PerpsMarket contract.
type PerpsMarketOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_PerpsMarket *PerpsMarketFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*PerpsMarketOwnerChangedIterator, error) {

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketOwnerChangedIterator{contract: _PerpsMarket.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_PerpsMarket *PerpsMarketFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *PerpsMarketOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketOwnerChanged)
				if err := _PerpsMarket.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseOwnerChanged(log types.Log) (*PerpsMarketOwnerChanged, error) {
	event := new(PerpsMarketOwnerChanged)
	if err := _PerpsMarket.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the PerpsMarket contract.
type PerpsMarketOwnerNominatedIterator struct {
	Event *PerpsMarketOwnerNominated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketOwnerNominated)
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
		it.Event = new(PerpsMarketOwnerNominated)
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
func (it *PerpsMarketOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketOwnerNominated represents a OwnerNominated event raised by the PerpsMarket contract.
type PerpsMarketOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_PerpsMarket *PerpsMarketFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*PerpsMarketOwnerNominatedIterator, error) {

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketOwnerNominatedIterator{contract: _PerpsMarket.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_PerpsMarket *PerpsMarketFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *PerpsMarketOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketOwnerNominated)
				if err := _PerpsMarket.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseOwnerNominated(log types.Log) (*PerpsMarketOwnerNominated, error) {
	event := new(PerpsMarketOwnerNominated)
	if err := _PerpsMarket.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketPerAccountCapsSetIterator is returned from FilterPerAccountCapsSet and is used to iterate over the raw logs and unpacked data for PerAccountCapsSet events raised by the PerpsMarket contract.
type PerpsMarketPerAccountCapsSetIterator struct {
	Event *PerpsMarketPerAccountCapsSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketPerAccountCapsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketPerAccountCapsSet)
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
		it.Event = new(PerpsMarketPerAccountCapsSet)
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
func (it *PerpsMarketPerAccountCapsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketPerAccountCapsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketPerAccountCapsSet represents a PerAccountCapsSet event raised by the PerpsMarket contract.
type PerpsMarketPerAccountCapsSet struct {
	MaxPositionsPerAccount   *big.Int
	MaxCollateralsPerAccount *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterPerAccountCapsSet is a free log retrieval operation binding the contract event 0x3448c6d1990f2d48253b91394193cd11ce49f1653f2d03934af6d17195ffe34e.
//
// Solidity: event PerAccountCapsSet(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount)
func (_PerpsMarket *PerpsMarketFilterer) FilterPerAccountCapsSet(opts *bind.FilterOpts) (*PerpsMarketPerAccountCapsSetIterator, error) {

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "PerAccountCapsSet")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketPerAccountCapsSetIterator{contract: _PerpsMarket.contract, event: "PerAccountCapsSet", logs: logs, sub: sub}, nil
}

// WatchPerAccountCapsSet is a free log subscription operation binding the contract event 0x3448c6d1990f2d48253b91394193cd11ce49f1653f2d03934af6d17195ffe34e.
//
// Solidity: event PerAccountCapsSet(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount)
func (_PerpsMarket *PerpsMarketFilterer) WatchPerAccountCapsSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketPerAccountCapsSet) (event.Subscription, error) {

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "PerAccountCapsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketPerAccountCapsSet)
				if err := _PerpsMarket.contract.UnpackLog(event, "PerAccountCapsSet", log); err != nil {
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

// ParsePerAccountCapsSet is a log parse operation binding the contract event 0x3448c6d1990f2d48253b91394193cd11ce49f1653f2d03934af6d17195ffe34e.
//
// Solidity: event PerAccountCapsSet(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount)
func (_PerpsMarket *PerpsMarketFilterer) ParsePerAccountCapsSet(log types.Log) (*PerpsMarketPerAccountCapsSet, error) {
	event := new(PerpsMarketPerAccountCapsSet)
	if err := _PerpsMarket.contract.UnpackLog(event, "PerAccountCapsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketPermissionGrantedIterator is returned from FilterPermissionGranted and is used to iterate over the raw logs and unpacked data for PermissionGranted events raised by the PerpsMarket contract.
type PerpsMarketPermissionGrantedIterator struct {
	Event *PerpsMarketPermissionGranted // Event containing the contract specifics and raw log

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
func (it *PerpsMarketPermissionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketPermissionGranted)
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
		it.Event = new(PerpsMarketPermissionGranted)
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
func (it *PerpsMarketPermissionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketPermissionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketPermissionGranted represents a PermissionGranted event raised by the PerpsMarket contract.
type PerpsMarketPermissionGranted struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionGranted is a free log retrieval operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PerpsMarket *PerpsMarketFilterer) FilterPermissionGranted(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*PerpsMarketPermissionGrantedIterator, error) {

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

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketPermissionGrantedIterator{contract: _PerpsMarket.contract, event: "PermissionGranted", logs: logs, sub: sub}, nil
}

// WatchPermissionGranted is a free log subscription operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PerpsMarket *PerpsMarketFilterer) WatchPermissionGranted(opts *bind.WatchOpts, sink chan<- *PerpsMarketPermissionGranted, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketPermissionGranted)
				if err := _PerpsMarket.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParsePermissionGranted(log types.Log) (*PerpsMarketPermissionGranted, error) {
	event := new(PerpsMarketPermissionGranted)
	if err := _PerpsMarket.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketPermissionRevokedIterator is returned from FilterPermissionRevoked and is used to iterate over the raw logs and unpacked data for PermissionRevoked events raised by the PerpsMarket contract.
type PerpsMarketPermissionRevokedIterator struct {
	Event *PerpsMarketPermissionRevoked // Event containing the contract specifics and raw log

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
func (it *PerpsMarketPermissionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketPermissionRevoked)
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
		it.Event = new(PerpsMarketPermissionRevoked)
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
func (it *PerpsMarketPermissionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketPermissionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketPermissionRevoked represents a PermissionRevoked event raised by the PerpsMarket contract.
type PerpsMarketPermissionRevoked struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionRevoked is a free log retrieval operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PerpsMarket *PerpsMarketFilterer) FilterPermissionRevoked(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*PerpsMarketPermissionRevokedIterator, error) {

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

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketPermissionRevokedIterator{contract: _PerpsMarket.contract, event: "PermissionRevoked", logs: logs, sub: sub}, nil
}

// WatchPermissionRevoked is a free log subscription operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PerpsMarket *PerpsMarketFilterer) WatchPermissionRevoked(opts *bind.WatchOpts, sink chan<- *PerpsMarketPermissionRevoked, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketPermissionRevoked)
				if err := _PerpsMarket.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParsePermissionRevoked(log types.Log) (*PerpsMarketPermissionRevoked, error) {
	event := new(PerpsMarketPermissionRevoked)
	if err := _PerpsMarket.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketPositionLiquidatedIterator is returned from FilterPositionLiquidated and is used to iterate over the raw logs and unpacked data for PositionLiquidated events raised by the PerpsMarket contract.
type PerpsMarketPositionLiquidatedIterator struct {
	Event *PerpsMarketPositionLiquidated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketPositionLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketPositionLiquidated)
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
		it.Event = new(PerpsMarketPositionLiquidated)
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
func (it *PerpsMarketPositionLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketPositionLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketPositionLiquidated represents a PositionLiquidated event raised by the PerpsMarket contract.
type PerpsMarketPositionLiquidated struct {
	AccountId           *big.Int
	MarketId            *big.Int
	AmountLiquidated    *big.Int
	CurrentPositionSize *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPositionLiquidated is a free log retrieval operation binding the contract event 0xd583c0e2965aae317f4a9a6603c0c75602b9bc97dc7c5fc6446b0ba8d3ff1bb2.
//
// Solidity: event PositionLiquidated(uint128 indexed accountId, uint128 indexed marketId, uint256 amountLiquidated, int128 currentPositionSize)
func (_PerpsMarket *PerpsMarketFilterer) FilterPositionLiquidated(opts *bind.FilterOpts, accountId []*big.Int, marketId []*big.Int) (*PerpsMarketPositionLiquidatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "PositionLiquidated", accountIdRule, marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketPositionLiquidatedIterator{contract: _PerpsMarket.contract, event: "PositionLiquidated", logs: logs, sub: sub}, nil
}

// WatchPositionLiquidated is a free log subscription operation binding the contract event 0xd583c0e2965aae317f4a9a6603c0c75602b9bc97dc7c5fc6446b0ba8d3ff1bb2.
//
// Solidity: event PositionLiquidated(uint128 indexed accountId, uint128 indexed marketId, uint256 amountLiquidated, int128 currentPositionSize)
func (_PerpsMarket *PerpsMarketFilterer) WatchPositionLiquidated(opts *bind.WatchOpts, sink chan<- *PerpsMarketPositionLiquidated, accountId []*big.Int, marketId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "PositionLiquidated", accountIdRule, marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketPositionLiquidated)
				if err := _PerpsMarket.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParsePositionLiquidated(log types.Log) (*PerpsMarketPositionLiquidated, error) {
	event := new(PerpsMarketPositionLiquidated)
	if err := _PerpsMarket.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketPreviousOrderExpiredIterator is returned from FilterPreviousOrderExpired and is used to iterate over the raw logs and unpacked data for PreviousOrderExpired events raised by the PerpsMarket contract.
type PerpsMarketPreviousOrderExpiredIterator struct {
	Event *PerpsMarketPreviousOrderExpired // Event containing the contract specifics and raw log

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
func (it *PerpsMarketPreviousOrderExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketPreviousOrderExpired)
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
		it.Event = new(PerpsMarketPreviousOrderExpired)
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
func (it *PerpsMarketPreviousOrderExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketPreviousOrderExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketPreviousOrderExpired represents a PreviousOrderExpired event raised by the PerpsMarket contract.
type PerpsMarketPreviousOrderExpired struct {
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
func (_PerpsMarket *PerpsMarketFilterer) FilterPreviousOrderExpired(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*PerpsMarketPreviousOrderExpiredIterator, error) {

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

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "PreviousOrderExpired", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketPreviousOrderExpiredIterator{contract: _PerpsMarket.contract, event: "PreviousOrderExpired", logs: logs, sub: sub}, nil
}

// WatchPreviousOrderExpired is a free log subscription operation binding the contract event 0x6d83c6751813f50325d75bc054621f83299659c5814d1e5fe6ac117860710dde.
//
// Solidity: event PreviousOrderExpired(uint128 indexed marketId, uint128 indexed accountId, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, bytes32 indexed trackingCode)
func (_PerpsMarket *PerpsMarketFilterer) WatchPreviousOrderExpired(opts *bind.WatchOpts, sink chan<- *PerpsMarketPreviousOrderExpired, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "PreviousOrderExpired", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketPreviousOrderExpired)
				if err := _PerpsMarket.contract.UnpackLog(event, "PreviousOrderExpired", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParsePreviousOrderExpired(log types.Log) (*PerpsMarketPreviousOrderExpired, error) {
	event := new(PerpsMarketPreviousOrderExpired)
	if err := _PerpsMarket.contract.UnpackLog(event, "PreviousOrderExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketReferrerShareUpdatedIterator is returned from FilterReferrerShareUpdated and is used to iterate over the raw logs and unpacked data for ReferrerShareUpdated events raised by the PerpsMarket contract.
type PerpsMarketReferrerShareUpdatedIterator struct {
	Event *PerpsMarketReferrerShareUpdated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketReferrerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketReferrerShareUpdated)
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
		it.Event = new(PerpsMarketReferrerShareUpdated)
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
func (it *PerpsMarketReferrerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketReferrerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketReferrerShareUpdated represents a ReferrerShareUpdated event raised by the PerpsMarket contract.
type PerpsMarketReferrerShareUpdated struct {
	Referrer      common.Address
	ShareRatioD18 *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReferrerShareUpdated is a free log retrieval operation binding the contract event 0xa225c555f4cd21a533ad4e01eaf30153c84ca28265d954a651410d3c1e56242c.
//
// Solidity: event ReferrerShareUpdated(address referrer, uint256 shareRatioD18)
func (_PerpsMarket *PerpsMarketFilterer) FilterReferrerShareUpdated(opts *bind.FilterOpts) (*PerpsMarketReferrerShareUpdatedIterator, error) {

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "ReferrerShareUpdated")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketReferrerShareUpdatedIterator{contract: _PerpsMarket.contract, event: "ReferrerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchReferrerShareUpdated is a free log subscription operation binding the contract event 0xa225c555f4cd21a533ad4e01eaf30153c84ca28265d954a651410d3c1e56242c.
//
// Solidity: event ReferrerShareUpdated(address referrer, uint256 shareRatioD18)
func (_PerpsMarket *PerpsMarketFilterer) WatchReferrerShareUpdated(opts *bind.WatchOpts, sink chan<- *PerpsMarketReferrerShareUpdated) (event.Subscription, error) {

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "ReferrerShareUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketReferrerShareUpdated)
				if err := _PerpsMarket.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseReferrerShareUpdated(log types.Log) (*PerpsMarketReferrerShareUpdated, error) {
	event := new(PerpsMarketReferrerShareUpdated)
	if err := _PerpsMarket.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketSettlementStrategyAddedIterator is returned from FilterSettlementStrategyAdded and is used to iterate over the raw logs and unpacked data for SettlementStrategyAdded events raised by the PerpsMarket contract.
type PerpsMarketSettlementStrategyAddedIterator struct {
	Event *PerpsMarketSettlementStrategyAdded // Event containing the contract specifics and raw log

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
func (it *PerpsMarketSettlementStrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketSettlementStrategyAdded)
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
		it.Event = new(PerpsMarketSettlementStrategyAdded)
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
func (it *PerpsMarketSettlementStrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketSettlementStrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketSettlementStrategyAdded represents a SettlementStrategyAdded event raised by the PerpsMarket contract.
type PerpsMarketSettlementStrategyAdded struct {
	MarketId   *big.Int
	Strategy   SettlementStrategyData
	StrategyId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyAdded is a free log retrieval operation binding the contract event 0xba2ece29eb53e3cc8cbe3bb5e85c899c5b6a9b84edff3e0503309f22a05d48ef.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,bool) strategy, uint256 indexed strategyId)
func (_PerpsMarket *PerpsMarketFilterer) FilterSettlementStrategyAdded(opts *bind.FilterOpts, marketId []*big.Int, strategyId []*big.Int) (*PerpsMarketSettlementStrategyAddedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "SettlementStrategyAdded", marketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketSettlementStrategyAddedIterator{contract: _PerpsMarket.contract, event: "SettlementStrategyAdded", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyAdded is a free log subscription operation binding the contract event 0xba2ece29eb53e3cc8cbe3bb5e85c899c5b6a9b84edff3e0503309f22a05d48ef.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,bool) strategy, uint256 indexed strategyId)
func (_PerpsMarket *PerpsMarketFilterer) WatchSettlementStrategyAdded(opts *bind.WatchOpts, sink chan<- *PerpsMarketSettlementStrategyAdded, marketId []*big.Int, strategyId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "SettlementStrategyAdded", marketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketSettlementStrategyAdded)
				if err := _PerpsMarket.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
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

// ParseSettlementStrategyAdded is a log parse operation binding the contract event 0xba2ece29eb53e3cc8cbe3bb5e85c899c5b6a9b84edff3e0503309f22a05d48ef.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,bool) strategy, uint256 indexed strategyId)
func (_PerpsMarket *PerpsMarketFilterer) ParseSettlementStrategyAdded(log types.Log) (*PerpsMarketSettlementStrategyAdded, error) {
	event := new(PerpsMarketSettlementStrategyAdded)
	if err := _PerpsMarket.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketSettlementStrategyEnabledIterator is returned from FilterSettlementStrategyEnabled and is used to iterate over the raw logs and unpacked data for SettlementStrategyEnabled events raised by the PerpsMarket contract.
type PerpsMarketSettlementStrategyEnabledIterator struct {
	Event *PerpsMarketSettlementStrategyEnabled // Event containing the contract specifics and raw log

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
func (it *PerpsMarketSettlementStrategyEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketSettlementStrategyEnabled)
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
		it.Event = new(PerpsMarketSettlementStrategyEnabled)
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
func (it *PerpsMarketSettlementStrategyEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketSettlementStrategyEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketSettlementStrategyEnabled represents a SettlementStrategyEnabled event raised by the PerpsMarket contract.
type PerpsMarketSettlementStrategyEnabled struct {
	MarketId   *big.Int
	StrategyId *big.Int
	Enabled    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyEnabled is a free log retrieval operation binding the contract event 0x4803eb4e51b45b2471b0c87d8c17e0c2d8961ae080cc4a930883eab155e0fa22.
//
// Solidity: event SettlementStrategyEnabled(uint128 indexed marketId, uint256 strategyId, bool enabled)
func (_PerpsMarket *PerpsMarketFilterer) FilterSettlementStrategyEnabled(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketSettlementStrategyEnabledIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "SettlementStrategyEnabled", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketSettlementStrategyEnabledIterator{contract: _PerpsMarket.contract, event: "SettlementStrategyEnabled", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyEnabled is a free log subscription operation binding the contract event 0x4803eb4e51b45b2471b0c87d8c17e0c2d8961ae080cc4a930883eab155e0fa22.
//
// Solidity: event SettlementStrategyEnabled(uint128 indexed marketId, uint256 strategyId, bool enabled)
func (_PerpsMarket *PerpsMarketFilterer) WatchSettlementStrategyEnabled(opts *bind.WatchOpts, sink chan<- *PerpsMarketSettlementStrategyEnabled, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "SettlementStrategyEnabled", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketSettlementStrategyEnabled)
				if err := _PerpsMarket.contract.UnpackLog(event, "SettlementStrategyEnabled", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseSettlementStrategyEnabled(log types.Log) (*PerpsMarketSettlementStrategyEnabled, error) {
	event := new(PerpsMarketSettlementStrategyEnabled)
	if err := _PerpsMarket.contract.UnpackLog(event, "SettlementStrategyEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketSynthDeductionPrioritySetIterator is returned from FilterSynthDeductionPrioritySet and is used to iterate over the raw logs and unpacked data for SynthDeductionPrioritySet events raised by the PerpsMarket contract.
type PerpsMarketSynthDeductionPrioritySetIterator struct {
	Event *PerpsMarketSynthDeductionPrioritySet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketSynthDeductionPrioritySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketSynthDeductionPrioritySet)
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
		it.Event = new(PerpsMarketSynthDeductionPrioritySet)
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
func (it *PerpsMarketSynthDeductionPrioritySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketSynthDeductionPrioritySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketSynthDeductionPrioritySet represents a SynthDeductionPrioritySet event raised by the PerpsMarket contract.
type PerpsMarketSynthDeductionPrioritySet struct {
	NewSynthDeductionPriority []*big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSynthDeductionPrioritySet is a free log retrieval operation binding the contract event 0xa6beea856d32db9e9614e4af02fc7d6a3fa8337a13e4a48486c142a9a9a8ed8f.
//
// Solidity: event SynthDeductionPrioritySet(uint128[] newSynthDeductionPriority)
func (_PerpsMarket *PerpsMarketFilterer) FilterSynthDeductionPrioritySet(opts *bind.FilterOpts) (*PerpsMarketSynthDeductionPrioritySetIterator, error) {

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "SynthDeductionPrioritySet")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketSynthDeductionPrioritySetIterator{contract: _PerpsMarket.contract, event: "SynthDeductionPrioritySet", logs: logs, sub: sub}, nil
}

// WatchSynthDeductionPrioritySet is a free log subscription operation binding the contract event 0xa6beea856d32db9e9614e4af02fc7d6a3fa8337a13e4a48486c142a9a9a8ed8f.
//
// Solidity: event SynthDeductionPrioritySet(uint128[] newSynthDeductionPriority)
func (_PerpsMarket *PerpsMarketFilterer) WatchSynthDeductionPrioritySet(opts *bind.WatchOpts, sink chan<- *PerpsMarketSynthDeductionPrioritySet) (event.Subscription, error) {

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "SynthDeductionPrioritySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketSynthDeductionPrioritySet)
				if err := _PerpsMarket.contract.UnpackLog(event, "SynthDeductionPrioritySet", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseSynthDeductionPrioritySet(log types.Log) (*PerpsMarketSynthDeductionPrioritySet, error) {
	event := new(PerpsMarketSynthDeductionPrioritySet)
	if err := _PerpsMarket.contract.UnpackLog(event, "SynthDeductionPrioritySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PerpsMarket contract.
type PerpsMarketUpgradedIterator struct {
	Event *PerpsMarketUpgraded // Event containing the contract specifics and raw log

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
func (it *PerpsMarketUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketUpgraded)
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
		it.Event = new(PerpsMarketUpgraded)
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
func (it *PerpsMarketUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketUpgraded represents a Upgraded event raised by the PerpsMarket contract.
type PerpsMarketUpgraded struct {
	Self           common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_PerpsMarket *PerpsMarketFilterer) FilterUpgraded(opts *bind.FilterOpts, self []common.Address) (*PerpsMarketUpgradedIterator, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _PerpsMarket.contract.FilterLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketUpgradedIterator{contract: _PerpsMarket.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_PerpsMarket *PerpsMarketFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PerpsMarketUpgraded, self []common.Address) (event.Subscription, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _PerpsMarket.contract.WatchLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketUpgraded)
				if err := _PerpsMarket.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_PerpsMarket *PerpsMarketFilterer) ParseUpgraded(log types.Log) (*PerpsMarketUpgraded, error) {
	event := new(PerpsMarketUpgraded)
	if err := _PerpsMarket.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
