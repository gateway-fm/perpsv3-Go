// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package perpsMarketGoerli

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

// PerpsMarketGoerliMetaData contains all meta data concerning the PerpsMarketGoerli contract.
var PerpsMarketGoerliMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"InvalidAccountId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"InvalidPermission\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"OnlyAccountTokenProxy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"PermissionNotGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedAccountId\",\"type\":\"uint128\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountLastInteraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountPermissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"permissions\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIAccountModule.AccountPermissions[]\",\"name\":\"accountPerms\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"notifyAccountTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"renouncePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"DeniedMulticallTarget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parameter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RecursiveMulticall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"getMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"multicallThrough\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowlisted\",\"type\":\"bool\"}],\"name\":\"setAllowlistedMulticallTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"InvalidMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PerpsMarketNotInitialized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"globalPerpsMarketId\",\"type\":\"uint128\"}],\"name\":\"FactoryInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"marketSymbol\",\"type\":\"string\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedMarketId\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"marketSymbol\",\"type\":\"string\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeFactory\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"minimumCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"reportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISpotMarketSystem\",\"name\":\"spotMarket\",\"type\":\"address\"}],\"name\":\"setSpotMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISynthetixSystem\",\"name\":\"synthetix\",\"type\":\"address\"}],\"name\":\"setSynthetix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountLiquidatable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"availableUsdDenominated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredUsdDenominated\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateralAvailableForWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientSynthCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"}],\"name\":\"InvalidAmountDelta\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"MaxCollateralExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"maxCollateralsPerAccount\",\"type\":\"uint128\"}],\"name\":\"MaxCollateralsPerAccountReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint128ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOrderExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"PriceFeedNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"SynthNotEnabledForCollateral\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"CollateralModified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAvailableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableMargin\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getOpenPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"totalPnl\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"accruedFunding\",\"type\":\"int256\"},{\"internalType\":\"int128\",\"name\":\"positionSize\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getRequiredMargins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredInitialMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMaintenanceMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAccumulatedLiquidationRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getWithdrawableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"withdrawableMargin\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"}],\"name\":\"modifyCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"totalAccountOpenInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"totalCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"currentFundingRate\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"currentFundingVelocity\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"orderSize\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"fillPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"skew\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIPerpsMarketModule.MarketSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"indexPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"maxOpenInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"metadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"skew\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"}],\"name\":\"AcceptablePriceExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"availableMargin\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"minMargin\",\"type\":\"uint256\"}],\"name\":\"InsufficientMargin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"}],\"name\":\"InvalidSettlementStrategy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"newSideSize\",\"type\":\"int256\"}],\"name\":\"MaxOpenInterestReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"maxPositionsPerAccount\",\"type\":\"uint128\"}],\"name\":\"MaxPositionsPerAccountReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSizeOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OrderCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"PreviousOrderExpired\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"commitment\",\"type\":\"tuple\"}],\"name\":\"commitOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"internalType\":\"structAsyncOrder.Data\",\"name\":\"retOrder\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"}],\"name\":\"computeOrderFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"internalType\":\"structAsyncOrder.Data\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"}],\"name\":\"requiredMarginForOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredMargin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leftover\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarginError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deviation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tolerance\",\"type\":\"uint256\"}],\"name\":\"PriceDeviationToleranceExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"}],\"name\":\"SettlementStrategyNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementExpiration\",\"type\":\"uint256\"}],\"name\":\"SettlementWindowExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"}],\"name\":\"SettlementWindowNotOpen\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"skew\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"sizeDelta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"}],\"name\":\"MarketUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"accruedFunding\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"newSize\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"}],\"name\":\"OrderSettled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"settlePythOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"NotEligibleForLiquidation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fullLiquidation\",\"type\":\"bool\"}],\"name\":\"AccountLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLiquidated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"currentPositionSize\",\"type\":\"int128\"}],\"name\":\"PositionLiquidated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"canLiquidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isEligible\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidationReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidateFlagged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidationReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"name\":\"FundingParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maintenanceMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"name\":\"LiquidationParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedOiRatioD18\",\"type\":\"uint256\"}],\"name\":\"LockedOiRatioSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"MarketPriceDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLiquidationPd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"endorsedLiquidator\",\"type\":\"address\"}],\"name\":\"MaxLiquidationParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"name\":\"MaxMarketSizeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerFeeRatio\",\"type\":\"uint256\"}],\"name\":\"OrderFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"SettlementStrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SettlementStrategyEnabled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"}],\"name\":\"addSettlementStrategy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getFundingParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getLiquidationParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginScalarD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getLockedOiRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMaxLiquidationParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationPd\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"endorsedLiquidator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMaxMarketSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getOrderFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"getSettlementStrategy\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"settlementStrategy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"name\":\"setFundingParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginScalarD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"name\":\"setLiquidationParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"lockedOiRatioD18\",\"type\":\"uint256\"}],\"name\":\"setLockedOiRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationPd\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"endorsedLiquidator\",\"type\":\"address\"}],\"name\":\"setMaxLiquidationParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"name\":\"setMaxMarketSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeRatio\",\"type\":\"uint256\"}],\"name\":\"setOrderFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setSettlementStrategyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"updatePriceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidFeeCollector\",\"type\":\"address\"}],\"name\":\"InvalidFeeCollectorInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"InvalidReferrerShareRatio\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"name\":\"LiquidationRewardGuardsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"MaxCollateralAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"maxPositionsPerAccount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"maxCollateralsPerAccount\",\"type\":\"uint128\"}],\"name\":\"PerAccountCapsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"ReferrerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"newSynthDeductionPriority\",\"type\":\"uint128[]\"}],\"name\":\"SynthDeductionPrioritySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidationRewardGuards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"marketIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMaxCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPerAccountCaps\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"maxPositionsPerAccount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxCollateralsPerAccount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"getReferrerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSynthDeductionPriority\",\"outputs\":[{\"internalType\":\"uint128[]\",\"name\":\"\",\"type\":\"uint128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"name\":\"setLiquidationRewardGuards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"setMaxCollateralAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"maxPositionsPerAccount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxCollateralsPerAccount\",\"type\":\"uint128\"}],\"name\":\"setPerAccountCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128[]\",\"name\":\"newSynthDeductionPriority\",\"type\":\"uint128[]\"}],\"name\":\"setSynthDeductionPriority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalGlobalCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"updateReferrerShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PerpsMarketGoerliABI is the input ABI used to generate the binding from.
// Deprecated: Use PerpsMarketGoerliMetaData.ABI instead.
var PerpsMarketGoerliABI = PerpsMarketGoerliMetaData.ABI

// PerpsMarketGoerli is an auto generated Go binding around an Ethereum contract.
type PerpsMarketGoerli struct {
	PerpsMarketGoerliCaller     // Read-only binding to the contract
	PerpsMarketGoerliTransactor // Write-only binding to the contract
	PerpsMarketGoerliFilterer   // Log filterer for contract events
}

// PerpsMarketGoerliCaller is an auto generated read-only Go binding around an Ethereum contract.
type PerpsMarketGoerliCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpsMarketGoerliTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PerpsMarketGoerliTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpsMarketGoerliFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PerpsMarketGoerliFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpsMarketGoerliSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PerpsMarketGoerliSession struct {
	Contract     *PerpsMarketGoerli // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PerpsMarketGoerliCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PerpsMarketGoerliCallerSession struct {
	Contract *PerpsMarketGoerliCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PerpsMarketGoerliTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PerpsMarketGoerliTransactorSession struct {
	Contract     *PerpsMarketGoerliTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PerpsMarketGoerliRaw is an auto generated low-level Go binding around an Ethereum contract.
type PerpsMarketGoerliRaw struct {
	Contract *PerpsMarketGoerli // Generic contract binding to access the raw methods on
}

// PerpsMarketGoerliCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PerpsMarketGoerliCallerRaw struct {
	Contract *PerpsMarketGoerliCaller // Generic read-only contract binding to access the raw methods on
}

// PerpsMarketGoerliTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PerpsMarketGoerliTransactorRaw struct {
	Contract *PerpsMarketGoerliTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerpsMarketGoerli creates a new instance of PerpsMarketGoerli, bound to a specific deployed contract.
func NewPerpsMarketGoerli(address common.Address, backend bind.ContractBackend) (*PerpsMarketGoerli, error) {
	contract, err := bindPerpsMarketGoerli(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerli{PerpsMarketGoerliCaller: PerpsMarketGoerliCaller{contract: contract}, PerpsMarketGoerliTransactor: PerpsMarketGoerliTransactor{contract: contract}, PerpsMarketGoerliFilterer: PerpsMarketGoerliFilterer{contract: contract}}, nil
}

// NewPerpsMarketGoerliCaller creates a new read-only instance of PerpsMarketGoerli, bound to a specific deployed contract.
func NewPerpsMarketGoerliCaller(address common.Address, caller bind.ContractCaller) (*PerpsMarketGoerliCaller, error) {
	contract, err := bindPerpsMarketGoerli(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliCaller{contract: contract}, nil
}

// NewPerpsMarketGoerliTransactor creates a new write-only instance of PerpsMarketGoerli, bound to a specific deployed contract.
func NewPerpsMarketGoerliTransactor(address common.Address, transactor bind.ContractTransactor) (*PerpsMarketGoerliTransactor, error) {
	contract, err := bindPerpsMarketGoerli(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliTransactor{contract: contract}, nil
}

// NewPerpsMarketGoerliFilterer creates a new log filterer instance of PerpsMarketGoerli, bound to a specific deployed contract.
func NewPerpsMarketGoerliFilterer(address common.Address, filterer bind.ContractFilterer) (*PerpsMarketGoerliFilterer, error) {
	contract, err := bindPerpsMarketGoerli(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliFilterer{contract: contract}, nil
}

// bindPerpsMarketGoerli binds a generic wrapper to an already deployed contract.
func bindPerpsMarketGoerli(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PerpsMarketGoerliMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerpsMarketGoerli *PerpsMarketGoerliRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerpsMarketGoerli.Contract.PerpsMarketGoerliCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerpsMarketGoerli *PerpsMarketGoerliRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.PerpsMarketGoerliTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerpsMarketGoerli *PerpsMarketGoerliRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.PerpsMarketGoerliTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerpsMarketGoerli.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) PRECISION() (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.PRECISION(&_PerpsMarketGoerli.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) PRECISION() (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.PRECISION(&_PerpsMarketGoerli.CallOpts)
}

// CanLiquidate is a free data retrieval call binding the contract method 0x9b922bba.
//
// Solidity: function canLiquidate(uint128 accountId) view returns(bool isEligible)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) CanLiquidate(opts *bind.CallOpts, accountId *big.Int) (bool, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "canLiquidate", accountId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanLiquidate is a free data retrieval call binding the contract method 0x9b922bba.
//
// Solidity: function canLiquidate(uint128 accountId) view returns(bool isEligible)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) CanLiquidate(accountId *big.Int) (bool, error) {
	return _PerpsMarketGoerli.Contract.CanLiquidate(&_PerpsMarketGoerli.CallOpts, accountId)
}

// CanLiquidate is a free data retrieval call binding the contract method 0x9b922bba.
//
// Solidity: function canLiquidate(uint128 accountId) view returns(bool isEligible)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) CanLiquidate(accountId *big.Int) (bool, error) {
	return _PerpsMarketGoerli.Contract.CanLiquidate(&_PerpsMarketGoerli.CallOpts, accountId)
}

// ComputeOrderFees is a free data retrieval call binding the contract method 0x98ef15a2.
//
// Solidity: function computeOrderFees(uint128 marketId, int128 sizeDelta) view returns(uint256 orderFees, uint256 fillPrice)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) ComputeOrderFees(opts *bind.CallOpts, marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "computeOrderFees", marketId, sizeDelta)

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) ComputeOrderFees(marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.ComputeOrderFees(&_PerpsMarketGoerli.CallOpts, marketId, sizeDelta)
}

// ComputeOrderFees is a free data retrieval call binding the contract method 0x98ef15a2.
//
// Solidity: function computeOrderFees(uint128 marketId, int128 sizeDelta) view returns(uint256 orderFees, uint256 fillPrice)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) ComputeOrderFees(marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.ComputeOrderFees(&_PerpsMarketGoerli.CallOpts, marketId, sizeDelta)
}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) CurrentFundingRate(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "currentFundingRate", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) CurrentFundingRate(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.CurrentFundingRate(&_PerpsMarketGoerli.CallOpts, marketId)
}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) CurrentFundingRate(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.CurrentFundingRate(&_PerpsMarketGoerli.CallOpts, marketId)
}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) CurrentFundingVelocity(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "currentFundingVelocity", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) CurrentFundingVelocity(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.CurrentFundingVelocity(&_PerpsMarketGoerli.CallOpts, marketId)
}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) CurrentFundingVelocity(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.CurrentFundingVelocity(&_PerpsMarketGoerli.CallOpts, marketId)
}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) FillPrice(opts *bind.CallOpts, marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "fillPrice", marketId, orderSize, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) FillPrice(marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.FillPrice(&_PerpsMarketGoerli.CallOpts, marketId, orderSize, price)
}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) FillPrice(marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.FillPrice(&_PerpsMarketGoerli.CallOpts, marketId, orderSize, price)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetAccountLastInteraction(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getAccountLastInteraction", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetAccountLastInteraction(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetAccountLastInteraction(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetAccountOwner(opts *bind.CallOpts, accountId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getAccountOwner", accountId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetAccountOwner(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetAccountOwner(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetAccountPermissions(opts *bind.CallOpts, accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getAccountPermissions", accountId)

	if err != nil {
		return *new([]IAccountModuleAccountPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAccountModuleAccountPermissions)).(*[]IAccountModuleAccountPermissions)

	return out0, err

}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _PerpsMarketGoerli.Contract.GetAccountPermissions(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _PerpsMarketGoerli.Contract.GetAccountPermissions(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetAccountTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getAccountTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetAccountTokenAddress() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetAccountTokenAddress(&_PerpsMarketGoerli.CallOpts)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetAccountTokenAddress() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetAccountTokenAddress(&_PerpsMarketGoerli.CallOpts)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetAssociatedSystem(opts *bind.CallOpts, id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getAssociatedSystem", id)

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _PerpsMarketGoerli.Contract.GetAssociatedSystem(&_PerpsMarketGoerli.CallOpts, id)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _PerpsMarketGoerli.Contract.GetAssociatedSystem(&_PerpsMarketGoerli.CallOpts, id)
}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetAvailableMargin(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getAvailableMargin", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetAvailableMargin(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetAvailableMargin(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetCollateralAmount(opts *bind.CallOpts, accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getCollateralAmount", accountId, synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetCollateralAmount(accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetCollateralAmount(&_PerpsMarketGoerli.CallOpts, accountId, synthMarketId)
}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetCollateralAmount(accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetCollateralAmount(&_PerpsMarketGoerli.CallOpts, accountId, synthMarketId)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetDeniers(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getDeniers", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetDeniers(&_PerpsMarketGoerli.CallOpts, feature)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetDeniers(&_PerpsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetFeatureFlagAllowAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getFeatureFlagAllowAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _PerpsMarketGoerli.Contract.GetFeatureFlagAllowAll(&_PerpsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _PerpsMarketGoerli.Contract.GetFeatureFlagAllowAll(&_PerpsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetFeatureFlagAllowlist(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getFeatureFlagAllowlist", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetFeatureFlagAllowlist(&_PerpsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetFeatureFlagAllowlist(&_PerpsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetFeatureFlagDenyAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getFeatureFlagDenyAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _PerpsMarketGoerli.Contract.GetFeatureFlagDenyAll(&_PerpsMarketGoerli.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _PerpsMarketGoerli.Contract.GetFeatureFlagDenyAll(&_PerpsMarketGoerli.CallOpts, feature)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetFeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getFeeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetFeeCollector() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetFeeCollector(&_PerpsMarketGoerli.CallOpts)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetFeeCollector() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetFeeCollector(&_PerpsMarketGoerli.CallOpts)
}

// GetFundingParameters is a free data retrieval call binding the contract method 0x1b68d8fa.
//
// Solidity: function getFundingParameters(uint128 marketId) view returns(uint256 skewScale, uint256 maxFundingVelocity)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetFundingParameters(opts *bind.CallOpts, marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getFundingParameters", marketId)

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetFundingParameters(marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetFundingParameters(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetFundingParameters is a free data retrieval call binding the contract method 0x1b68d8fa.
//
// Solidity: function getFundingParameters(uint128 marketId) view returns(uint256 skewScale, uint256 maxFundingVelocity)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetFundingParameters(marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetFundingParameters(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetImplementation() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetImplementation(&_PerpsMarketGoerli.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetImplementation() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetImplementation(&_PerpsMarketGoerli.CallOpts)
}

// GetLiquidationParameters is a free data retrieval call binding the contract method 0xf94363a6.
//
// Solidity: function getLiquidationParameters(uint128 marketId) view returns(uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetLiquidationParameters(opts *bind.CallOpts, marketId *big.Int) (struct {
	InitialMarginRatioD18        *big.Int
	MinimumInitialMarginRatioD18 *big.Int
	MaintenanceMarginScalarD18   *big.Int
	LiquidationRewardRatioD18    *big.Int
	MinimumPositionMargin        *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getLiquidationParameters", marketId)

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetLiquidationParameters(marketId *big.Int) (struct {
	InitialMarginRatioD18        *big.Int
	MinimumInitialMarginRatioD18 *big.Int
	MaintenanceMarginScalarD18   *big.Int
	LiquidationRewardRatioD18    *big.Int
	MinimumPositionMargin        *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetLiquidationParameters(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetLiquidationParameters is a free data retrieval call binding the contract method 0xf94363a6.
//
// Solidity: function getLiquidationParameters(uint128 marketId) view returns(uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetLiquidationParameters(marketId *big.Int) (struct {
	InitialMarginRatioD18        *big.Int
	MinimumInitialMarginRatioD18 *big.Int
	MaintenanceMarginScalarD18   *big.Int
	LiquidationRewardRatioD18    *big.Int
	MinimumPositionMargin        *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetLiquidationParameters(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetLiquidationRewardGuards is a free data retrieval call binding the contract method 0x0b7f4b2d.
//
// Solidity: function getLiquidationRewardGuards() view returns(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetLiquidationRewardGuards(opts *bind.CallOpts) (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getLiquidationRewardGuards")

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetLiquidationRewardGuards() (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetLiquidationRewardGuards(&_PerpsMarketGoerli.CallOpts)
}

// GetLiquidationRewardGuards is a free data retrieval call binding the contract method 0x0b7f4b2d.
//
// Solidity: function getLiquidationRewardGuards() view returns(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetLiquidationRewardGuards() (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetLiquidationRewardGuards(&_PerpsMarketGoerli.CallOpts)
}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetLockedOiRatio(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getLockedOiRatio", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetLockedOiRatio(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetLockedOiRatio(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetLockedOiRatio(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetLockedOiRatio(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetMarketSummary(opts *bind.CallOpts, marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getMarketSummary", marketId)

	if err != nil {
		return *new(IPerpsMarketModuleMarketSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(IPerpsMarketModuleMarketSummary)).(*IPerpsMarketModuleMarketSummary)

	return out0, err

}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetMarketSummary(marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	return _PerpsMarketGoerli.Contract.GetMarketSummary(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetMarketSummary(marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	return _PerpsMarketGoerli.Contract.GetMarketSummary(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetMarkets(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getMarkets")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetMarkets() ([]*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetMarkets(&_PerpsMarketGoerli.CallOpts)
}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetMarkets() ([]*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetMarkets(&_PerpsMarketGoerli.CallOpts)
}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetMaxCollateralAmount(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getMaxCollateralAmount", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetMaxCollateralAmount(synthMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetMaxCollateralAmount(&_PerpsMarketGoerli.CallOpts, synthMarketId)
}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetMaxCollateralAmount(synthMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetMaxCollateralAmount(&_PerpsMarketGoerli.CallOpts, synthMarketId)
}

// GetMaxLiquidationParameters is a free data retrieval call binding the contract method 0x5443e33e.
//
// Solidity: function getMaxLiquidationParameters(uint128 marketId) view returns(uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetMaxLiquidationParameters(opts *bind.CallOpts, marketId *big.Int) (struct {
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MaxLiquidationPd                          *big.Int
	EndorsedLiquidator                        common.Address
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getMaxLiquidationParameters", marketId)

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetMaxLiquidationParameters(marketId *big.Int) (struct {
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MaxLiquidationPd                          *big.Int
	EndorsedLiquidator                        common.Address
}, error) {
	return _PerpsMarketGoerli.Contract.GetMaxLiquidationParameters(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetMaxLiquidationParameters is a free data retrieval call binding the contract method 0x5443e33e.
//
// Solidity: function getMaxLiquidationParameters(uint128 marketId) view returns(uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetMaxLiquidationParameters(marketId *big.Int) (struct {
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MaxLiquidationPd                          *big.Int
	EndorsedLiquidator                        common.Address
}, error) {
	return _PerpsMarketGoerli.Contract.GetMaxLiquidationParameters(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetMaxMarketSize(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getMaxMarketSize", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetMaxMarketSize(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetMaxMarketSize(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetMaxMarketSize(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetMaxMarketSize(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetMessageSender is a free data retrieval call binding the contract method 0x76167a89.
//
// Solidity: function getMessageSender() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMessageSender is a free data retrieval call binding the contract method 0x76167a89.
//
// Solidity: function getMessageSender() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetMessageSender() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetMessageSender(&_PerpsMarketGoerli.CallOpts)
}

// GetMessageSender is a free data retrieval call binding the contract method 0x76167a89.
//
// Solidity: function getMessageSender() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetMessageSender() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.GetMessageSender(&_PerpsMarketGoerli.CallOpts)
}

// GetOpenPosition is a free data retrieval call binding the contract method 0x22a73967.
//
// Solidity: function getOpenPosition(uint128 accountId, uint128 marketId) view returns(int256 totalPnl, int256 accruedFunding, int128 positionSize)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetOpenPosition(opts *bind.CallOpts, accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getOpenPosition", accountId, marketId)

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetOpenPosition(accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetOpenPosition(&_PerpsMarketGoerli.CallOpts, accountId, marketId)
}

// GetOpenPosition is a free data retrieval call binding the contract method 0x22a73967.
//
// Solidity: function getOpenPosition(uint128 accountId, uint128 marketId) view returns(int256 totalPnl, int256 accruedFunding, int128 positionSize)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetOpenPosition(accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetOpenPosition(&_PerpsMarketGoerli.CallOpts, accountId, marketId)
}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetOrder(opts *bind.CallOpts, accountId *big.Int) (AsyncOrderData, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getOrder", accountId)

	if err != nil {
		return *new(AsyncOrderData), err
	}

	out0 := *abi.ConvertType(out[0], new(AsyncOrderData)).(*AsyncOrderData)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetOrder(accountId *big.Int) (AsyncOrderData, error) {
	return _PerpsMarketGoerli.Contract.GetOrder(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetOrder(accountId *big.Int) (AsyncOrderData, error) {
	return _PerpsMarketGoerli.Contract.GetOrder(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetOrderFees is a free data retrieval call binding the contract method 0xaac23e8c.
//
// Solidity: function getOrderFees(uint128 marketId) view returns(uint256 makerFee, uint256 takerFee)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetOrderFees(opts *bind.CallOpts, marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getOrderFees", marketId)

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetOrderFees(marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetOrderFees(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetOrderFees is a free data retrieval call binding the contract method 0xaac23e8c.
//
// Solidity: function getOrderFees(uint128 marketId) view returns(uint256 makerFee, uint256 takerFee)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetOrderFees(marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetOrderFees(&_PerpsMarketGoerli.CallOpts, marketId)
}

// GetPerAccountCaps is a free data retrieval call binding the contract method 0x774f7b07.
//
// Solidity: function getPerAccountCaps() view returns(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetPerAccountCaps(opts *bind.CallOpts) (struct {
	MaxPositionsPerAccount   *big.Int
	MaxCollateralsPerAccount *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getPerAccountCaps")

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetPerAccountCaps() (struct {
	MaxPositionsPerAccount   *big.Int
	MaxCollateralsPerAccount *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetPerAccountCaps(&_PerpsMarketGoerli.CallOpts)
}

// GetPerAccountCaps is a free data retrieval call binding the contract method 0x774f7b07.
//
// Solidity: function getPerAccountCaps() view returns(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetPerAccountCaps() (struct {
	MaxPositionsPerAccount   *big.Int
	MaxCollateralsPerAccount *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetPerAccountCaps(&_PerpsMarketGoerli.CallOpts)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetReferrerShare(opts *bind.CallOpts, referrer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getReferrerShare", referrer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetReferrerShare(referrer common.Address) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetReferrerShare(&_PerpsMarketGoerli.CallOpts, referrer)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetReferrerShare(referrer common.Address) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetReferrerShare(&_PerpsMarketGoerli.CallOpts, referrer)
}

// GetRequiredMargins is a free data retrieval call binding the contract method 0x3c0f0753.
//
// Solidity: function getRequiredMargins(uint128 accountId) view returns(uint256 requiredInitialMargin, uint256 requiredMaintenanceMargin, uint256 totalAccumulatedLiquidationRewards, uint256 maxLiquidationReward)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetRequiredMargins(opts *bind.CallOpts, accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getRequiredMargins", accountId)

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetRequiredMargins(accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetRequiredMargins(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetRequiredMargins is a free data retrieval call binding the contract method 0x3c0f0753.
//
// Solidity: function getRequiredMargins(uint128 accountId) view returns(uint256 requiredInitialMargin, uint256 requiredMaintenanceMargin, uint256 totalAccumulatedLiquidationRewards, uint256 maxLiquidationReward)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetRequiredMargins(accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	return _PerpsMarketGoerli.Contract.GetRequiredMargins(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) settlementStrategy)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetSettlementStrategy(opts *bind.CallOpts, marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getSettlementStrategy", marketId, strategyId)

	if err != nil {
		return *new(SettlementStrategyData), err
	}

	out0 := *abi.ConvertType(out[0], new(SettlementStrategyData)).(*SettlementStrategyData)

	return out0, err

}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) settlementStrategy)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _PerpsMarketGoerli.Contract.GetSettlementStrategy(&_PerpsMarketGoerli.CallOpts, marketId, strategyId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) settlementStrategy)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _PerpsMarketGoerli.Contract.GetSettlementStrategy(&_PerpsMarketGoerli.CallOpts, marketId, strategyId)
}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetSynthDeductionPriority(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getSynthDeductionPriority")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetSynthDeductionPriority() ([]*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetSynthDeductionPriority(&_PerpsMarketGoerli.CallOpts)
}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetSynthDeductionPriority() ([]*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetSynthDeductionPriority(&_PerpsMarketGoerli.CallOpts)
}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) GetWithdrawableMargin(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "getWithdrawableMargin", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GetWithdrawableMargin(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetWithdrawableMargin(&_PerpsMarketGoerli.CallOpts, accountId)
}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) GetWithdrawableMargin(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.GetWithdrawableMargin(&_PerpsMarketGoerli.CallOpts, accountId)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) HasPermission(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "hasPermission", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PerpsMarketGoerli.Contract.HasPermission(&_PerpsMarketGoerli.CallOpts, accountId, permission, user)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PerpsMarketGoerli.Contract.HasPermission(&_PerpsMarketGoerli.CallOpts, accountId, permission, user)
}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) IndexPrice(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "indexPrice", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) IndexPrice(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.IndexPrice(&_PerpsMarketGoerli.CallOpts, marketId)
}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) IndexPrice(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.IndexPrice(&_PerpsMarketGoerli.CallOpts, marketId)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) IsAuthorized(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "isAuthorized", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PerpsMarketGoerli.Contract.IsAuthorized(&_PerpsMarketGoerli.CallOpts, accountId, permission, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _PerpsMarketGoerli.Contract.IsAuthorized(&_PerpsMarketGoerli.CallOpts, accountId, permission, user)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) IsFeatureAllowed(opts *bind.CallOpts, feature [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "isFeatureAllowed", feature, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _PerpsMarketGoerli.Contract.IsFeatureAllowed(&_PerpsMarketGoerli.CallOpts, feature, account)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _PerpsMarketGoerli.Contract.IsFeatureAllowed(&_PerpsMarketGoerli.CallOpts, feature, account)
}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) MaxOpenInterest(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "maxOpenInterest", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) MaxOpenInterest(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.MaxOpenInterest(&_PerpsMarketGoerli.CallOpts, marketId)
}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) MaxOpenInterest(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.MaxOpenInterest(&_PerpsMarketGoerli.CallOpts, marketId)
}

// Metadata is a free data retrieval call binding the contract method 0xe3bc36bf.
//
// Solidity: function metadata(uint128 marketId) view returns(string name, string symbol)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) Metadata(opts *bind.CallOpts, marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "metadata", marketId)

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
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) Metadata(marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	return _PerpsMarketGoerli.Contract.Metadata(&_PerpsMarketGoerli.CallOpts, marketId)
}

// Metadata is a free data retrieval call binding the contract method 0xe3bc36bf.
//
// Solidity: function metadata(uint128 marketId) view returns(string name, string symbol)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) Metadata(marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	return _PerpsMarketGoerli.Contract.Metadata(&_PerpsMarketGoerli.CallOpts, marketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) MinimumCredit(opts *bind.CallOpts, perpsMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "minimumCredit", perpsMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) MinimumCredit(perpsMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.MinimumCredit(&_PerpsMarketGoerli.CallOpts, perpsMarketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) MinimumCredit(perpsMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.MinimumCredit(&_PerpsMarketGoerli.CallOpts, perpsMarketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) Name(opts *bind.CallOpts, perpsMarketId *big.Int) (string, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "name", perpsMarketId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) Name(perpsMarketId *big.Int) (string, error) {
	return _PerpsMarketGoerli.Contract.Name(&_PerpsMarketGoerli.CallOpts, perpsMarketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) Name(perpsMarketId *big.Int) (string, error) {
	return _PerpsMarketGoerli.Contract.Name(&_PerpsMarketGoerli.CallOpts, perpsMarketId)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) NominatedOwner() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.NominatedOwner(&_PerpsMarketGoerli.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) NominatedOwner() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.NominatedOwner(&_PerpsMarketGoerli.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) Owner() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.Owner(&_PerpsMarketGoerli.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) Owner() (common.Address, error) {
	return _PerpsMarketGoerli.Contract.Owner(&_PerpsMarketGoerli.CallOpts)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) ReportedDebt(opts *bind.CallOpts, perpsMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "reportedDebt", perpsMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) ReportedDebt(perpsMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.ReportedDebt(&_PerpsMarketGoerli.CallOpts, perpsMarketId)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) ReportedDebt(perpsMarketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.ReportedDebt(&_PerpsMarketGoerli.CallOpts, perpsMarketId)
}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) RequiredMarginForOrder(opts *bind.CallOpts, accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "requiredMarginForOrder", accountId, marketId, sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) RequiredMarginForOrder(accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.RequiredMarginForOrder(&_PerpsMarketGoerli.CallOpts, accountId, marketId, sizeDelta)
}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) RequiredMarginForOrder(accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.RequiredMarginForOrder(&_PerpsMarketGoerli.CallOpts, accountId, marketId, sizeDelta)
}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) Settle(opts *bind.CallOpts, accountId *big.Int) error {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "settle", accountId)

	if err != nil {
		return err
	}

	return err

}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) Settle(accountId *big.Int) error {
	return _PerpsMarketGoerli.Contract.Settle(&_PerpsMarketGoerli.CallOpts, accountId)
}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) Settle(accountId *big.Int) error {
	return _PerpsMarketGoerli.Contract.Settle(&_PerpsMarketGoerli.CallOpts, accountId)
}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) Size(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "size", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) Size(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.Size(&_PerpsMarketGoerli.CallOpts, marketId)
}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) Size(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.Size(&_PerpsMarketGoerli.CallOpts, marketId)
}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) Skew(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "skew", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) Skew(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.Skew(&_PerpsMarketGoerli.CallOpts, marketId)
}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) Skew(marketId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.Skew(&_PerpsMarketGoerli.CallOpts, marketId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PerpsMarketGoerli.Contract.SupportsInterface(&_PerpsMarketGoerli.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PerpsMarketGoerli.Contract.SupportsInterface(&_PerpsMarketGoerli.CallOpts, interfaceId)
}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) TotalAccountOpenInterest(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "totalAccountOpenInterest", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) TotalAccountOpenInterest(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.TotalAccountOpenInterest(&_PerpsMarketGoerli.CallOpts, accountId)
}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) TotalAccountOpenInterest(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.TotalAccountOpenInterest(&_PerpsMarketGoerli.CallOpts, accountId)
}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) TotalCollateralValue(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "totalCollateralValue", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) TotalCollateralValue(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.TotalCollateralValue(&_PerpsMarketGoerli.CallOpts, accountId)
}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) TotalCollateralValue(accountId *big.Int) (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.TotalCollateralValue(&_PerpsMarketGoerli.CallOpts, accountId)
}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_PerpsMarketGoerli *PerpsMarketGoerliCaller) TotalGlobalCollateralValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpsMarketGoerli.contract.Call(opts, &out, "totalGlobalCollateralValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) TotalGlobalCollateralValue() (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.TotalGlobalCollateralValue(&_PerpsMarketGoerli.CallOpts)
}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_PerpsMarketGoerli *PerpsMarketGoerliCallerSession) TotalGlobalCollateralValue() (*big.Int, error) {
	return _PerpsMarketGoerli.Contract.TotalGlobalCollateralValue(&_PerpsMarketGoerli.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) AcceptOwnership() (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.AcceptOwnership(&_PerpsMarketGoerli.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.AcceptOwnership(&_PerpsMarketGoerli.TransactOpts)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0xda91187c.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) AddSettlementStrategy(opts *bind.TransactOpts, marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "addSettlementStrategy", marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0xda91187c.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.AddSettlementStrategy(&_PerpsMarketGoerli.TransactOpts, marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0xda91187c.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.AddSettlementStrategy(&_PerpsMarketGoerli.TransactOpts, marketId, strategy)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) AddToFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "addToFeatureFlagAllowlist", feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.AddToFeatureFlagAllowlist(&_PerpsMarketGoerli.TransactOpts, feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.AddToFeatureFlagAllowlist(&_PerpsMarketGoerli.TransactOpts, feature, account)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) CommitOrder(opts *bind.TransactOpts, commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "commitOrder", commitment)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) CommitOrder(commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.CommitOrder(&_PerpsMarketGoerli.TransactOpts, commitment)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) CommitOrder(commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.CommitOrder(&_PerpsMarketGoerli.TransactOpts, commitment)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) CreateAccount() (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.CreateAccount(&_PerpsMarketGoerli.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.CreateAccount(&_PerpsMarketGoerli.TransactOpts)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) CreateAccount0(opts *bind.TransactOpts, requestedAccountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "createAccount0", requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.CreateAccount0(&_PerpsMarketGoerli.TransactOpts, requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.CreateAccount0(&_PerpsMarketGoerli.TransactOpts, requestedAccountId)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) CreateMarket(opts *bind.TransactOpts, requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "createMarket", requestedMarketId, marketName, marketSymbol)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) CreateMarket(requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.CreateMarket(&_PerpsMarketGoerli.TransactOpts, requestedMarketId, marketName, marketSymbol)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) CreateMarket(requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.CreateMarket(&_PerpsMarketGoerli.TransactOpts, requestedMarketId, marketName, marketSymbol)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) GrantPermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "grantPermission", accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.GrantPermission(&_PerpsMarketGoerli.TransactOpts, accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.GrantPermission(&_PerpsMarketGoerli.TransactOpts, accountId, permission, user)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) InitOrUpgradeNft(opts *bind.TransactOpts, id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "initOrUpgradeNft", id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.InitOrUpgradeNft(&_PerpsMarketGoerli.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.InitOrUpgradeNft(&_PerpsMarketGoerli.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) InitOrUpgradeToken(opts *bind.TransactOpts, id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "initOrUpgradeToken", id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.InitOrUpgradeToken(&_PerpsMarketGoerli.TransactOpts, id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.InitOrUpgradeToken(&_PerpsMarketGoerli.TransactOpts, id, name, symbol, decimals, impl)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) InitializeFactory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "initializeFactory")
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) InitializeFactory() (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.InitializeFactory(&_PerpsMarketGoerli.TransactOpts)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) InitializeFactory() (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.InitializeFactory(&_PerpsMarketGoerli.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns(uint256 liquidationReward)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) Liquidate(opts *bind.TransactOpts, accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "liquidate", accountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns(uint256 liquidationReward)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) Liquidate(accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.Liquidate(&_PerpsMarketGoerli.TransactOpts, accountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns(uint256 liquidationReward)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) Liquidate(accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.Liquidate(&_PerpsMarketGoerli.TransactOpts, accountId)
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns(uint256 liquidationReward)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) LiquidateFlagged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "liquidateFlagged")
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns(uint256 liquidationReward)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) LiquidateFlagged() (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.LiquidateFlagged(&_PerpsMarketGoerli.TransactOpts)
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns(uint256 liquidationReward)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) LiquidateFlagged() (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.LiquidateFlagged(&_PerpsMarketGoerli.TransactOpts)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) ModifyCollateral(opts *bind.TransactOpts, accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "modifyCollateral", accountId, synthMarketId, amountDelta)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) ModifyCollateral(accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.ModifyCollateral(&_PerpsMarketGoerli.TransactOpts, accountId, synthMarketId, amountDelta)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) ModifyCollateral(accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.ModifyCollateral(&_PerpsMarketGoerli.TransactOpts, accountId, synthMarketId, amountDelta)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.Multicall(&_PerpsMarketGoerli.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.Multicall(&_PerpsMarketGoerli.TransactOpts, data)
}

// MulticallThrough is a paid mutator transaction binding the contract method 0xc0694da0.
//
// Solidity: function multicallThrough(address[] to, bytes[] data, uint256[] values) payable returns(bytes[] results)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) MulticallThrough(opts *bind.TransactOpts, to []common.Address, data [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "multicallThrough", to, data, values)
}

// MulticallThrough is a paid mutator transaction binding the contract method 0xc0694da0.
//
// Solidity: function multicallThrough(address[] to, bytes[] data, uint256[] values) payable returns(bytes[] results)
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) MulticallThrough(to []common.Address, data [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.MulticallThrough(&_PerpsMarketGoerli.TransactOpts, to, data, values)
}

// MulticallThrough is a paid mutator transaction binding the contract method 0xc0694da0.
//
// Solidity: function multicallThrough(address[] to, bytes[] data, uint256[] values) payable returns(bytes[] results)
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) MulticallThrough(to []common.Address, data [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.MulticallThrough(&_PerpsMarketGoerli.TransactOpts, to, data, values)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) NominateNewOwner(opts *bind.TransactOpts, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "nominateNewOwner", newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.NominateNewOwner(&_PerpsMarketGoerli.TransactOpts, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.NominateNewOwner(&_PerpsMarketGoerli.TransactOpts, newNominatedOwner)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) NotifyAccountTransfer(opts *bind.TransactOpts, to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "notifyAccountTransfer", to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.NotifyAccountTransfer(&_PerpsMarketGoerli.TransactOpts, to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.NotifyAccountTransfer(&_PerpsMarketGoerli.TransactOpts, to, accountId)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) RegisterUnmanagedSystem(opts *bind.TransactOpts, id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "registerUnmanagedSystem", id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.RegisterUnmanagedSystem(&_PerpsMarketGoerli.TransactOpts, id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.RegisterUnmanagedSystem(&_PerpsMarketGoerli.TransactOpts, id, endpoint)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) RemoveFromFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "removeFromFeatureFlagAllowlist", feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.RemoveFromFeatureFlagAllowlist(&_PerpsMarketGoerli.TransactOpts, feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.RemoveFromFeatureFlagAllowlist(&_PerpsMarketGoerli.TransactOpts, feature, account)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) RenounceNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "renounceNomination")
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) RenounceNomination() (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.RenounceNomination(&_PerpsMarketGoerli.TransactOpts)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) RenounceNomination() (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.RenounceNomination(&_PerpsMarketGoerli.TransactOpts)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) RenouncePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "renouncePermission", accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.RenouncePermission(&_PerpsMarketGoerli.TransactOpts, accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.RenouncePermission(&_PerpsMarketGoerli.TransactOpts, accountId, permission)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) RevokePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "revokePermission", accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.RevokePermission(&_PerpsMarketGoerli.TransactOpts, accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.RevokePermission(&_PerpsMarketGoerli.TransactOpts, accountId, permission, user)
}

// SetAllowlistedMulticallTarget is a paid mutator transaction binding the contract method 0x0a6c7553.
//
// Solidity: function setAllowlistedMulticallTarget(address target, bool allowlisted) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetAllowlistedMulticallTarget(opts *bind.TransactOpts, target common.Address, allowlisted bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setAllowlistedMulticallTarget", target, allowlisted)
}

// SetAllowlistedMulticallTarget is a paid mutator transaction binding the contract method 0x0a6c7553.
//
// Solidity: function setAllowlistedMulticallTarget(address target, bool allowlisted) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetAllowlistedMulticallTarget(target common.Address, allowlisted bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetAllowlistedMulticallTarget(&_PerpsMarketGoerli.TransactOpts, target, allowlisted)
}

// SetAllowlistedMulticallTarget is a paid mutator transaction binding the contract method 0x0a6c7553.
//
// Solidity: function setAllowlistedMulticallTarget(address target, bool allowlisted) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetAllowlistedMulticallTarget(target common.Address, allowlisted bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetAllowlistedMulticallTarget(&_PerpsMarketGoerli.TransactOpts, target, allowlisted)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetDeniers(opts *bind.TransactOpts, feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setDeniers", feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetDeniers(&_PerpsMarketGoerli.TransactOpts, feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetDeniers(&_PerpsMarketGoerli.TransactOpts, feature, deniers)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetFeatureFlagAllowAll(opts *bind.TransactOpts, feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setFeatureFlagAllowAll", feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetFeatureFlagAllowAll(&_PerpsMarketGoerli.TransactOpts, feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetFeatureFlagAllowAll(&_PerpsMarketGoerli.TransactOpts, feature, allowAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetFeatureFlagDenyAll(opts *bind.TransactOpts, feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setFeatureFlagDenyAll", feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetFeatureFlagDenyAll(&_PerpsMarketGoerli.TransactOpts, feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetFeatureFlagDenyAll(&_PerpsMarketGoerli.TransactOpts, feature, denyAll)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetFeeCollector(opts *bind.TransactOpts, feeCollector common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setFeeCollector", feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetFeeCollector(&_PerpsMarketGoerli.TransactOpts, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetFeeCollector(&_PerpsMarketGoerli.TransactOpts, feeCollector)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetFundingParameters(opts *bind.TransactOpts, marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setFundingParameters", marketId, skewScale, maxFundingVelocity)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetFundingParameters(marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetFundingParameters(&_PerpsMarketGoerli.TransactOpts, marketId, skewScale, maxFundingVelocity)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetFundingParameters(marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetFundingParameters(&_PerpsMarketGoerli.TransactOpts, marketId, skewScale, maxFundingVelocity)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0x25e5409e.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetLiquidationParameters(opts *bind.TransactOpts, marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setLiquidationParameters", marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, minimumPositionMargin)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0x25e5409e.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetLiquidationParameters(marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetLiquidationParameters(&_PerpsMarketGoerli.TransactOpts, marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, minimumPositionMargin)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0x25e5409e.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetLiquidationParameters(marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetLiquidationParameters(&_PerpsMarketGoerli.TransactOpts, marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, minimumPositionMargin)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetLiquidationRewardGuards(opts *bind.TransactOpts, minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setLiquidationRewardGuards", minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetLiquidationRewardGuards(minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetLiquidationRewardGuards(&_PerpsMarketGoerli.TransactOpts, minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetLiquidationRewardGuards(minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetLiquidationRewardGuards(&_PerpsMarketGoerli.TransactOpts, minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetLockedOiRatio(opts *bind.TransactOpts, marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setLockedOiRatio", marketId, lockedOiRatioD18)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetLockedOiRatio(marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetLockedOiRatio(&_PerpsMarketGoerli.TransactOpts, marketId, lockedOiRatioD18)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetLockedOiRatio(marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetLockedOiRatio(&_PerpsMarketGoerli.TransactOpts, marketId, lockedOiRatioD18)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetMaxCollateralAmount(opts *bind.TransactOpts, synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setMaxCollateralAmount", synthMarketId, collateralAmount)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetMaxCollateralAmount(synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetMaxCollateralAmount(&_PerpsMarketGoerli.TransactOpts, synthMarketId, collateralAmount)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetMaxCollateralAmount(synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetMaxCollateralAmount(&_PerpsMarketGoerli.TransactOpts, synthMarketId, collateralAmount)
}

// SetMaxLiquidationParameters is a paid mutator transaction binding the contract method 0xc7f8a94f.
//
// Solidity: function setMaxLiquidationParameters(uint128 marketId, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetMaxLiquidationParameters(opts *bind.TransactOpts, marketId *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, maxLiquidationPd *big.Int, endorsedLiquidator common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setMaxLiquidationParameters", marketId, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, maxLiquidationPd, endorsedLiquidator)
}

// SetMaxLiquidationParameters is a paid mutator transaction binding the contract method 0xc7f8a94f.
//
// Solidity: function setMaxLiquidationParameters(uint128 marketId, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetMaxLiquidationParameters(marketId *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, maxLiquidationPd *big.Int, endorsedLiquidator common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetMaxLiquidationParameters(&_PerpsMarketGoerli.TransactOpts, marketId, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, maxLiquidationPd, endorsedLiquidator)
}

// SetMaxLiquidationParameters is a paid mutator transaction binding the contract method 0xc7f8a94f.
//
// Solidity: function setMaxLiquidationParameters(uint128 marketId, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetMaxLiquidationParameters(marketId *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, maxLiquidationPd *big.Int, endorsedLiquidator common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetMaxLiquidationParameters(&_PerpsMarketGoerli.TransactOpts, marketId, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, maxLiquidationPd, endorsedLiquidator)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetMaxMarketSize(opts *bind.TransactOpts, marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setMaxMarketSize", marketId, maxMarketSize)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetMaxMarketSize(marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetMaxMarketSize(&_PerpsMarketGoerli.TransactOpts, marketId, maxMarketSize)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetMaxMarketSize(marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetMaxMarketSize(&_PerpsMarketGoerli.TransactOpts, marketId, maxMarketSize)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetOrderFees(opts *bind.TransactOpts, marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setOrderFees", marketId, makerFeeRatio, takerFeeRatio)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetOrderFees(marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetOrderFees(&_PerpsMarketGoerli.TransactOpts, marketId, makerFeeRatio, takerFeeRatio)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetOrderFees(marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetOrderFees(&_PerpsMarketGoerli.TransactOpts, marketId, makerFeeRatio, takerFeeRatio)
}

// SetPerAccountCaps is a paid mutator transaction binding the contract method 0xfa0e70a7.
//
// Solidity: function setPerAccountCaps(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetPerAccountCaps(opts *bind.TransactOpts, maxPositionsPerAccount *big.Int, maxCollateralsPerAccount *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setPerAccountCaps", maxPositionsPerAccount, maxCollateralsPerAccount)
}

// SetPerAccountCaps is a paid mutator transaction binding the contract method 0xfa0e70a7.
//
// Solidity: function setPerAccountCaps(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetPerAccountCaps(maxPositionsPerAccount *big.Int, maxCollateralsPerAccount *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetPerAccountCaps(&_PerpsMarketGoerli.TransactOpts, maxPositionsPerAccount, maxCollateralsPerAccount)
}

// SetPerAccountCaps is a paid mutator transaction binding the contract method 0xfa0e70a7.
//
// Solidity: function setPerAccountCaps(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetPerAccountCaps(maxPositionsPerAccount *big.Int, maxCollateralsPerAccount *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetPerAccountCaps(&_PerpsMarketGoerli.TransactOpts, maxPositionsPerAccount, maxCollateralsPerAccount)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetSettlementStrategyEnabled(opts *bind.TransactOpts, marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setSettlementStrategyEnabled", marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetSettlementStrategyEnabled(&_PerpsMarketGoerli.TransactOpts, marketId, strategyId, enabled)
}

// SetSettlementStrategyEnabled is a paid mutator transaction binding the contract method 0x7f73a891.
//
// Solidity: function setSettlementStrategyEnabled(uint128 marketId, uint256 strategyId, bool enabled) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetSettlementStrategyEnabled(marketId *big.Int, strategyId *big.Int, enabled bool) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetSettlementStrategyEnabled(&_PerpsMarketGoerli.TransactOpts, marketId, strategyId, enabled)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetSpotMarket(opts *bind.TransactOpts, spotMarket common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setSpotMarket", spotMarket)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetSpotMarket(spotMarket common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetSpotMarket(&_PerpsMarketGoerli.TransactOpts, spotMarket)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetSpotMarket(spotMarket common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetSpotMarket(&_PerpsMarketGoerli.TransactOpts, spotMarket)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetSynthDeductionPriority(opts *bind.TransactOpts, newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setSynthDeductionPriority", newSynthDeductionPriority)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetSynthDeductionPriority(newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetSynthDeductionPriority(&_PerpsMarketGoerli.TransactOpts, newSynthDeductionPriority)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetSynthDeductionPriority(newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetSynthDeductionPriority(&_PerpsMarketGoerli.TransactOpts, newSynthDeductionPriority)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SetSynthetix(opts *bind.TransactOpts, synthetix common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "setSynthetix", synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetSynthetix(&_PerpsMarketGoerli.TransactOpts, synthetix)
}

// SetSynthetix is a paid mutator transaction binding the contract method 0xfec9f9da.
//
// Solidity: function setSynthetix(address synthetix) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SetSynthetix(synthetix common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SetSynthetix(&_PerpsMarketGoerli.TransactOpts, synthetix)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SettlePythOrder(opts *bind.TransactOpts, result []byte, extraData []byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "settlePythOrder", result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SettlePythOrder(&_PerpsMarketGoerli.TransactOpts, result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SettlePythOrder(&_PerpsMarketGoerli.TransactOpts, result, extraData)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) SimulateUpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "simulateUpgradeTo", newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SimulateUpgradeTo(&_PerpsMarketGoerli.TransactOpts, newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.SimulateUpgradeTo(&_PerpsMarketGoerli.TransactOpts, newImplementation)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) UpdatePriceData(opts *bind.TransactOpts, perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "updatePriceData", perpsMarketId, feedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) UpdatePriceData(perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.UpdatePriceData(&_PerpsMarketGoerli.TransactOpts, perpsMarketId, feedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) UpdatePriceData(perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.UpdatePriceData(&_PerpsMarketGoerli.TransactOpts, perpsMarketId, feedId)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) UpdateReferrerShare(opts *bind.TransactOpts, referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "updateReferrerShare", referrer, shareRatioD18)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) UpdateReferrerShare(referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.UpdateReferrerShare(&_PerpsMarketGoerli.TransactOpts, referrer, shareRatioD18)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) UpdateReferrerShare(referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.UpdateReferrerShare(&_PerpsMarketGoerli.TransactOpts, referrer, shareRatioD18)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.UpgradeTo(&_PerpsMarketGoerli.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PerpsMarketGoerli *PerpsMarketGoerliTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PerpsMarketGoerli.Contract.UpgradeTo(&_PerpsMarketGoerli.TransactOpts, newImplementation)
}

// PerpsMarketGoerliAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliAccountCreatedIterator struct {
	Event *PerpsMarketGoerliAccountCreated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliAccountCreated)
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
		it.Event = new(PerpsMarketGoerliAccountCreated)
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
func (it *PerpsMarketGoerliAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliAccountCreated represents a AccountCreated event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliAccountCreated struct {
	AccountId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterAccountCreated(opts *bind.FilterOpts, accountId []*big.Int, owner []common.Address) (*PerpsMarketGoerliAccountCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliAccountCreatedIterator{contract: _PerpsMarketGoerli.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliAccountCreated, accountId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliAccountCreated)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseAccountCreated(log types.Log) (*PerpsMarketGoerliAccountCreated, error) {
	event := new(PerpsMarketGoerliAccountCreated)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliAccountLiquidatedIterator is returned from FilterAccountLiquidated and is used to iterate over the raw logs and unpacked data for AccountLiquidated events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliAccountLiquidatedIterator struct {
	Event *PerpsMarketGoerliAccountLiquidated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliAccountLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliAccountLiquidated)
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
		it.Event = new(PerpsMarketGoerliAccountLiquidated)
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
func (it *PerpsMarketGoerliAccountLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliAccountLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliAccountLiquidated represents a AccountLiquidated event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliAccountLiquidated struct {
	AccountId       *big.Int
	Reward          *big.Int
	FullLiquidation bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccountLiquidated is a free log retrieval operation binding the contract event 0x926f56b0a80cc150ad77b59cc60bc53cb907eab441d204924ec67b40ae64f5f5.
//
// Solidity: event AccountLiquidated(uint128 indexed accountId, uint256 reward, bool fullLiquidation)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterAccountLiquidated(opts *bind.FilterOpts, accountId []*big.Int) (*PerpsMarketGoerliAccountLiquidatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "AccountLiquidated", accountIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliAccountLiquidatedIterator{contract: _PerpsMarketGoerli.contract, event: "AccountLiquidated", logs: logs, sub: sub}, nil
}

// WatchAccountLiquidated is a free log subscription operation binding the contract event 0x926f56b0a80cc150ad77b59cc60bc53cb907eab441d204924ec67b40ae64f5f5.
//
// Solidity: event AccountLiquidated(uint128 indexed accountId, uint256 reward, bool fullLiquidation)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchAccountLiquidated(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliAccountLiquidated, accountId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "AccountLiquidated", accountIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliAccountLiquidated)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "AccountLiquidated", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseAccountLiquidated(log types.Log) (*PerpsMarketGoerliAccountLiquidated, error) {
	event := new(PerpsMarketGoerliAccountLiquidated)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "AccountLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliAssociatedSystemSetIterator is returned from FilterAssociatedSystemSet and is used to iterate over the raw logs and unpacked data for AssociatedSystemSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliAssociatedSystemSetIterator struct {
	Event *PerpsMarketGoerliAssociatedSystemSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliAssociatedSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliAssociatedSystemSet)
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
		it.Event = new(PerpsMarketGoerliAssociatedSystemSet)
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
func (it *PerpsMarketGoerliAssociatedSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliAssociatedSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliAssociatedSystemSet represents a AssociatedSystemSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliAssociatedSystemSet struct {
	Kind  [32]byte
	Id    [32]byte
	Proxy common.Address
	Impl  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssociatedSystemSet is a free log retrieval operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterAssociatedSystemSet(opts *bind.FilterOpts, kind [][32]byte, id [][32]byte) (*PerpsMarketGoerliAssociatedSystemSetIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliAssociatedSystemSetIterator{contract: _PerpsMarketGoerli.contract, event: "AssociatedSystemSet", logs: logs, sub: sub}, nil
}

// WatchAssociatedSystemSet is a free log subscription operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchAssociatedSystemSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliAssociatedSystemSet, kind [][32]byte, id [][32]byte) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliAssociatedSystemSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseAssociatedSystemSet(log types.Log) (*PerpsMarketGoerliAssociatedSystemSet, error) {
	event := new(PerpsMarketGoerliAssociatedSystemSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliCollateralModifiedIterator is returned from FilterCollateralModified and is used to iterate over the raw logs and unpacked data for CollateralModified events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliCollateralModifiedIterator struct {
	Event *PerpsMarketGoerliCollateralModified // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliCollateralModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliCollateralModified)
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
		it.Event = new(PerpsMarketGoerliCollateralModified)
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
func (it *PerpsMarketGoerliCollateralModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliCollateralModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliCollateralModified represents a CollateralModified event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliCollateralModified struct {
	AccountId     *big.Int
	SynthMarketId *big.Int
	AmountDelta   *big.Int
	Sender        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralModified is a free log retrieval operation binding the contract event 0x2e8360c2f3a6fc1a15aefdd0a0922bea3c898cb90d38c3a97354e7c013116064.
//
// Solidity: event CollateralModified(uint128 indexed accountId, uint128 indexed synthMarketId, int256 amountDelta, address indexed sender)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterCollateralModified(opts *bind.FilterOpts, accountId []*big.Int, synthMarketId []*big.Int, sender []common.Address) (*PerpsMarketGoerliCollateralModifiedIterator, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "CollateralModified", accountIdRule, synthMarketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliCollateralModifiedIterator{contract: _PerpsMarketGoerli.contract, event: "CollateralModified", logs: logs, sub: sub}, nil
}

// WatchCollateralModified is a free log subscription operation binding the contract event 0x2e8360c2f3a6fc1a15aefdd0a0922bea3c898cb90d38c3a97354e7c013116064.
//
// Solidity: event CollateralModified(uint128 indexed accountId, uint128 indexed synthMarketId, int256 amountDelta, address indexed sender)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchCollateralModified(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliCollateralModified, accountId []*big.Int, synthMarketId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "CollateralModified", accountIdRule, synthMarketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliCollateralModified)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "CollateralModified", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseCollateralModified(log types.Log) (*PerpsMarketGoerliCollateralModified, error) {
	event := new(PerpsMarketGoerliCollateralModified)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "CollateralModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliFactoryInitializedIterator is returned from FilterFactoryInitialized and is used to iterate over the raw logs and unpacked data for FactoryInitialized events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFactoryInitializedIterator struct {
	Event *PerpsMarketGoerliFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliFactoryInitialized)
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
		it.Event = new(PerpsMarketGoerliFactoryInitialized)
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
func (it *PerpsMarketGoerliFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliFactoryInitialized represents a FactoryInitialized event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFactoryInitialized struct {
	GlobalPerpsMarketId *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFactoryInitialized is a free log retrieval operation binding the contract event 0xb3240229b07e26f2f02e69dda85ede86e162ccbc6d10e6aade28931e7f533339.
//
// Solidity: event FactoryInitialized(uint128 globalPerpsMarketId)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterFactoryInitialized(opts *bind.FilterOpts) (*PerpsMarketGoerliFactoryInitializedIterator, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "FactoryInitialized")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliFactoryInitializedIterator{contract: _PerpsMarketGoerli.contract, event: "FactoryInitialized", logs: logs, sub: sub}, nil
}

// WatchFactoryInitialized is a free log subscription operation binding the contract event 0xb3240229b07e26f2f02e69dda85ede86e162ccbc6d10e6aade28931e7f533339.
//
// Solidity: event FactoryInitialized(uint128 globalPerpsMarketId)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchFactoryInitialized(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "FactoryInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliFactoryInitialized)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FactoryInitialized", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseFactoryInitialized(log types.Log) (*PerpsMarketGoerliFactoryInitialized, error) {
	event := new(PerpsMarketGoerliFactoryInitialized)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FactoryInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliFeatureFlagAllowAllSetIterator is returned from FilterFeatureFlagAllowAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowAllSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeatureFlagAllowAllSetIterator struct {
	Event *PerpsMarketGoerliFeatureFlagAllowAllSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliFeatureFlagAllowAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliFeatureFlagAllowAllSet)
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
		it.Event = new(PerpsMarketGoerliFeatureFlagAllowAllSet)
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
func (it *PerpsMarketGoerliFeatureFlagAllowAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliFeatureFlagAllowAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliFeatureFlagAllowAllSet represents a FeatureFlagAllowAllSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeatureFlagAllowAllSet struct {
	Feature  [32]byte
	AllowAll bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowAllSet is a free log retrieval operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterFeatureFlagAllowAllSet(opts *bind.FilterOpts, feature [][32]byte) (*PerpsMarketGoerliFeatureFlagAllowAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliFeatureFlagAllowAllSetIterator{contract: _PerpsMarketGoerli.contract, event: "FeatureFlagAllowAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowAllSet is a free log subscription operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchFeatureFlagAllowAllSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliFeatureFlagAllowAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliFeatureFlagAllowAllSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseFeatureFlagAllowAllSet(log types.Log) (*PerpsMarketGoerliFeatureFlagAllowAllSet, error) {
	event := new(PerpsMarketGoerliFeatureFlagAllowAllSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliFeatureFlagAllowlistAddedIterator is returned from FilterFeatureFlagAllowlistAdded and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistAdded events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeatureFlagAllowlistAddedIterator struct {
	Event *PerpsMarketGoerliFeatureFlagAllowlistAdded // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliFeatureFlagAllowlistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliFeatureFlagAllowlistAdded)
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
		it.Event = new(PerpsMarketGoerliFeatureFlagAllowlistAdded)
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
func (it *PerpsMarketGoerliFeatureFlagAllowlistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliFeatureFlagAllowlistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliFeatureFlagAllowlistAdded represents a FeatureFlagAllowlistAdded event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeatureFlagAllowlistAdded struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistAdded is a free log retrieval operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterFeatureFlagAllowlistAdded(opts *bind.FilterOpts, feature [][32]byte) (*PerpsMarketGoerliFeatureFlagAllowlistAddedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliFeatureFlagAllowlistAddedIterator{contract: _PerpsMarketGoerli.contract, event: "FeatureFlagAllowlistAdded", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistAdded is a free log subscription operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchFeatureFlagAllowlistAdded(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliFeatureFlagAllowlistAdded, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliFeatureFlagAllowlistAdded)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseFeatureFlagAllowlistAdded(log types.Log) (*PerpsMarketGoerliFeatureFlagAllowlistAdded, error) {
	event := new(PerpsMarketGoerliFeatureFlagAllowlistAdded)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliFeatureFlagAllowlistRemovedIterator is returned from FilterFeatureFlagAllowlistRemoved and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistRemoved events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeatureFlagAllowlistRemovedIterator struct {
	Event *PerpsMarketGoerliFeatureFlagAllowlistRemoved // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliFeatureFlagAllowlistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliFeatureFlagAllowlistRemoved)
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
		it.Event = new(PerpsMarketGoerliFeatureFlagAllowlistRemoved)
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
func (it *PerpsMarketGoerliFeatureFlagAllowlistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliFeatureFlagAllowlistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliFeatureFlagAllowlistRemoved represents a FeatureFlagAllowlistRemoved event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeatureFlagAllowlistRemoved struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistRemoved is a free log retrieval operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterFeatureFlagAllowlistRemoved(opts *bind.FilterOpts, feature [][32]byte) (*PerpsMarketGoerliFeatureFlagAllowlistRemovedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliFeatureFlagAllowlistRemovedIterator{contract: _PerpsMarketGoerli.contract, event: "FeatureFlagAllowlistRemoved", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistRemoved is a free log subscription operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchFeatureFlagAllowlistRemoved(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliFeatureFlagAllowlistRemoved, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliFeatureFlagAllowlistRemoved)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseFeatureFlagAllowlistRemoved(log types.Log) (*PerpsMarketGoerliFeatureFlagAllowlistRemoved, error) {
	event := new(PerpsMarketGoerliFeatureFlagAllowlistRemoved)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliFeatureFlagDeniersResetIterator is returned from FilterFeatureFlagDeniersReset and is used to iterate over the raw logs and unpacked data for FeatureFlagDeniersReset events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeatureFlagDeniersResetIterator struct {
	Event *PerpsMarketGoerliFeatureFlagDeniersReset // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliFeatureFlagDeniersResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliFeatureFlagDeniersReset)
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
		it.Event = new(PerpsMarketGoerliFeatureFlagDeniersReset)
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
func (it *PerpsMarketGoerliFeatureFlagDeniersResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliFeatureFlagDeniersResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliFeatureFlagDeniersReset represents a FeatureFlagDeniersReset event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeatureFlagDeniersReset struct {
	Feature [32]byte
	Deniers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDeniersReset is a free log retrieval operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterFeatureFlagDeniersReset(opts *bind.FilterOpts, feature [][32]byte) (*PerpsMarketGoerliFeatureFlagDeniersResetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliFeatureFlagDeniersResetIterator{contract: _PerpsMarketGoerli.contract, event: "FeatureFlagDeniersReset", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDeniersReset is a free log subscription operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchFeatureFlagDeniersReset(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliFeatureFlagDeniersReset, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliFeatureFlagDeniersReset)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseFeatureFlagDeniersReset(log types.Log) (*PerpsMarketGoerliFeatureFlagDeniersReset, error) {
	event := new(PerpsMarketGoerliFeatureFlagDeniersReset)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliFeatureFlagDenyAllSetIterator is returned from FilterFeatureFlagDenyAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagDenyAllSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeatureFlagDenyAllSetIterator struct {
	Event *PerpsMarketGoerliFeatureFlagDenyAllSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliFeatureFlagDenyAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliFeatureFlagDenyAllSet)
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
		it.Event = new(PerpsMarketGoerliFeatureFlagDenyAllSet)
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
func (it *PerpsMarketGoerliFeatureFlagDenyAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliFeatureFlagDenyAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliFeatureFlagDenyAllSet represents a FeatureFlagDenyAllSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeatureFlagDenyAllSet struct {
	Feature [32]byte
	DenyAll bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDenyAllSet is a free log retrieval operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterFeatureFlagDenyAllSet(opts *bind.FilterOpts, feature [][32]byte) (*PerpsMarketGoerliFeatureFlagDenyAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliFeatureFlagDenyAllSetIterator{contract: _PerpsMarketGoerli.contract, event: "FeatureFlagDenyAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDenyAllSet is a free log subscription operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchFeatureFlagDenyAllSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliFeatureFlagDenyAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliFeatureFlagDenyAllSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseFeatureFlagDenyAllSet(log types.Log) (*PerpsMarketGoerliFeatureFlagDenyAllSet, error) {
	event := new(PerpsMarketGoerliFeatureFlagDenyAllSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliFeeCollectorSetIterator is returned from FilterFeeCollectorSet and is used to iterate over the raw logs and unpacked data for FeeCollectorSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeeCollectorSetIterator struct {
	Event *PerpsMarketGoerliFeeCollectorSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliFeeCollectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliFeeCollectorSet)
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
		it.Event = new(PerpsMarketGoerliFeeCollectorSet)
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
func (it *PerpsMarketGoerliFeeCollectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliFeeCollectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliFeeCollectorSet represents a FeeCollectorSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFeeCollectorSet struct {
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorSet is a free log retrieval operation binding the contract event 0x12e1d17016b94668449f97876f4a8d5cc2c19f314db337418894734037cc19d4.
//
// Solidity: event FeeCollectorSet(address feeCollector)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterFeeCollectorSet(opts *bind.FilterOpts) (*PerpsMarketGoerliFeeCollectorSetIterator, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "FeeCollectorSet")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliFeeCollectorSetIterator{contract: _PerpsMarketGoerli.contract, event: "FeeCollectorSet", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorSet is a free log subscription operation binding the contract event 0x12e1d17016b94668449f97876f4a8d5cc2c19f314db337418894734037cc19d4.
//
// Solidity: event FeeCollectorSet(address feeCollector)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchFeeCollectorSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliFeeCollectorSet) (event.Subscription, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "FeeCollectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliFeeCollectorSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseFeeCollectorSet(log types.Log) (*PerpsMarketGoerliFeeCollectorSet, error) {
	event := new(PerpsMarketGoerliFeeCollectorSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliFundingParametersSetIterator is returned from FilterFundingParametersSet and is used to iterate over the raw logs and unpacked data for FundingParametersSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFundingParametersSetIterator struct {
	Event *PerpsMarketGoerliFundingParametersSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliFundingParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliFundingParametersSet)
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
		it.Event = new(PerpsMarketGoerliFundingParametersSet)
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
func (it *PerpsMarketGoerliFundingParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliFundingParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliFundingParametersSet represents a FundingParametersSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliFundingParametersSet struct {
	MarketId           *big.Int
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFundingParametersSet is a free log retrieval operation binding the contract event 0xa74afd926bbafbb9252d224a1fcd6a209f851324bd485f556786820a76e31b65.
//
// Solidity: event FundingParametersSet(uint128 indexed marketId, uint256 skewScale, uint256 maxFundingVelocity)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterFundingParametersSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketGoerliFundingParametersSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "FundingParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliFundingParametersSetIterator{contract: _PerpsMarketGoerli.contract, event: "FundingParametersSet", logs: logs, sub: sub}, nil
}

// WatchFundingParametersSet is a free log subscription operation binding the contract event 0xa74afd926bbafbb9252d224a1fcd6a209f851324bd485f556786820a76e31b65.
//
// Solidity: event FundingParametersSet(uint128 indexed marketId, uint256 skewScale, uint256 maxFundingVelocity)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchFundingParametersSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliFundingParametersSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "FundingParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliFundingParametersSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FundingParametersSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseFundingParametersSet(log types.Log) (*PerpsMarketGoerliFundingParametersSet, error) {
	event := new(PerpsMarketGoerliFundingParametersSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "FundingParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliLiquidationParametersSetIterator is returned from FilterLiquidationParametersSet and is used to iterate over the raw logs and unpacked data for LiquidationParametersSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliLiquidationParametersSetIterator struct {
	Event *PerpsMarketGoerliLiquidationParametersSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliLiquidationParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliLiquidationParametersSet)
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
		it.Event = new(PerpsMarketGoerliLiquidationParametersSet)
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
func (it *PerpsMarketGoerliLiquidationParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliLiquidationParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliLiquidationParametersSet represents a LiquidationParametersSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliLiquidationParametersSet struct {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterLiquidationParametersSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketGoerliLiquidationParametersSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "LiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliLiquidationParametersSetIterator{contract: _PerpsMarketGoerli.contract, event: "LiquidationParametersSet", logs: logs, sub: sub}, nil
}

// WatchLiquidationParametersSet is a free log subscription operation binding the contract event 0xa0c87f048ec4f5924e50d554aa4a6e65a935f133a2114e5222590c1690e1a7b8.
//
// Solidity: event LiquidationParametersSet(uint128 indexed marketId, uint256 initialMarginRatioD18, uint256 maintenanceMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 liquidationRewardRatioD18, uint256 minimumPositionMargin)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchLiquidationParametersSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliLiquidationParametersSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "LiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliLiquidationParametersSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "LiquidationParametersSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseLiquidationParametersSet(log types.Log) (*PerpsMarketGoerliLiquidationParametersSet, error) {
	event := new(PerpsMarketGoerliLiquidationParametersSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "LiquidationParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliLiquidationRewardGuardsSetIterator is returned from FilterLiquidationRewardGuardsSet and is used to iterate over the raw logs and unpacked data for LiquidationRewardGuardsSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliLiquidationRewardGuardsSetIterator struct {
	Event *PerpsMarketGoerliLiquidationRewardGuardsSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliLiquidationRewardGuardsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliLiquidationRewardGuardsSet)
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
		it.Event = new(PerpsMarketGoerliLiquidationRewardGuardsSet)
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
func (it *PerpsMarketGoerliLiquidationRewardGuardsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliLiquidationRewardGuardsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliLiquidationRewardGuardsSet represents a LiquidationRewardGuardsSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliLiquidationRewardGuardsSet struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterLiquidationRewardGuardsSet is a free log retrieval operation binding the contract event 0xbb7f8eb17b00549c0f234a3c363ef3eb39329f98b52048416480cf62d0f0e1f5.
//
// Solidity: event LiquidationRewardGuardsSet(uint256 indexed minLiquidationRewardUsd, uint256 indexed maxLiquidationRewardUsd)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterLiquidationRewardGuardsSet(opts *bind.FilterOpts, minLiquidationRewardUsd []*big.Int, maxLiquidationRewardUsd []*big.Int) (*PerpsMarketGoerliLiquidationRewardGuardsSetIterator, error) {

	var minLiquidationRewardUsdRule []interface{}
	for _, minLiquidationRewardUsdItem := range minLiquidationRewardUsd {
		minLiquidationRewardUsdRule = append(minLiquidationRewardUsdRule, minLiquidationRewardUsdItem)
	}
	var maxLiquidationRewardUsdRule []interface{}
	for _, maxLiquidationRewardUsdItem := range maxLiquidationRewardUsd {
		maxLiquidationRewardUsdRule = append(maxLiquidationRewardUsdRule, maxLiquidationRewardUsdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "LiquidationRewardGuardsSet", minLiquidationRewardUsdRule, maxLiquidationRewardUsdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliLiquidationRewardGuardsSetIterator{contract: _PerpsMarketGoerli.contract, event: "LiquidationRewardGuardsSet", logs: logs, sub: sub}, nil
}

// WatchLiquidationRewardGuardsSet is a free log subscription operation binding the contract event 0xbb7f8eb17b00549c0f234a3c363ef3eb39329f98b52048416480cf62d0f0e1f5.
//
// Solidity: event LiquidationRewardGuardsSet(uint256 indexed minLiquidationRewardUsd, uint256 indexed maxLiquidationRewardUsd)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchLiquidationRewardGuardsSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliLiquidationRewardGuardsSet, minLiquidationRewardUsd []*big.Int, maxLiquidationRewardUsd []*big.Int) (event.Subscription, error) {

	var minLiquidationRewardUsdRule []interface{}
	for _, minLiquidationRewardUsdItem := range minLiquidationRewardUsd {
		minLiquidationRewardUsdRule = append(minLiquidationRewardUsdRule, minLiquidationRewardUsdItem)
	}
	var maxLiquidationRewardUsdRule []interface{}
	for _, maxLiquidationRewardUsdItem := range maxLiquidationRewardUsd {
		maxLiquidationRewardUsdRule = append(maxLiquidationRewardUsdRule, maxLiquidationRewardUsdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "LiquidationRewardGuardsSet", minLiquidationRewardUsdRule, maxLiquidationRewardUsdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliLiquidationRewardGuardsSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "LiquidationRewardGuardsSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseLiquidationRewardGuardsSet(log types.Log) (*PerpsMarketGoerliLiquidationRewardGuardsSet, error) {
	event := new(PerpsMarketGoerliLiquidationRewardGuardsSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "LiquidationRewardGuardsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliLockedOiRatioSetIterator is returned from FilterLockedOiRatioSet and is used to iterate over the raw logs and unpacked data for LockedOiRatioSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliLockedOiRatioSetIterator struct {
	Event *PerpsMarketGoerliLockedOiRatioSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliLockedOiRatioSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliLockedOiRatioSet)
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
		it.Event = new(PerpsMarketGoerliLockedOiRatioSet)
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
func (it *PerpsMarketGoerliLockedOiRatioSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliLockedOiRatioSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliLockedOiRatioSet represents a LockedOiRatioSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliLockedOiRatioSet struct {
	MarketId         *big.Int
	LockedOiRatioD18 *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLockedOiRatioSet is a free log retrieval operation binding the contract event 0x1d841fd5b4c806bc5a073d637a8506e1e74d16cb18251b711cb47e133ceafc2d.
//
// Solidity: event LockedOiRatioSet(uint128 indexed marketId, uint256 lockedOiRatioD18)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterLockedOiRatioSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketGoerliLockedOiRatioSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "LockedOiRatioSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliLockedOiRatioSetIterator{contract: _PerpsMarketGoerli.contract, event: "LockedOiRatioSet", logs: logs, sub: sub}, nil
}

// WatchLockedOiRatioSet is a free log subscription operation binding the contract event 0x1d841fd5b4c806bc5a073d637a8506e1e74d16cb18251b711cb47e133ceafc2d.
//
// Solidity: event LockedOiRatioSet(uint128 indexed marketId, uint256 lockedOiRatioD18)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchLockedOiRatioSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliLockedOiRatioSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "LockedOiRatioSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliLockedOiRatioSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "LockedOiRatioSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseLockedOiRatioSet(log types.Log) (*PerpsMarketGoerliLockedOiRatioSet, error) {
	event := new(PerpsMarketGoerliLockedOiRatioSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "LockedOiRatioSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMarketCreatedIterator struct {
	Event *PerpsMarketGoerliMarketCreated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliMarketCreated)
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
		it.Event = new(PerpsMarketGoerliMarketCreated)
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
func (it *PerpsMarketGoerliMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliMarketCreated represents a MarketCreated event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMarketCreated struct {
	PerpsMarketId *big.Int
	MarketName    string
	MarketSymbol  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0x032553f94ac1323933f22650ec5b8e232babf1c47efca69383b749463116cc49.
//
// Solidity: event MarketCreated(uint128 indexed perpsMarketId, string marketName, string marketSymbol)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterMarketCreated(opts *bind.FilterOpts, perpsMarketId []*big.Int) (*PerpsMarketGoerliMarketCreatedIterator, error) {

	var perpsMarketIdRule []interface{}
	for _, perpsMarketIdItem := range perpsMarketId {
		perpsMarketIdRule = append(perpsMarketIdRule, perpsMarketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "MarketCreated", perpsMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliMarketCreatedIterator{contract: _PerpsMarketGoerli.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0x032553f94ac1323933f22650ec5b8e232babf1c47efca69383b749463116cc49.
//
// Solidity: event MarketCreated(uint128 indexed perpsMarketId, string marketName, string marketSymbol)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliMarketCreated, perpsMarketId []*big.Int) (event.Subscription, error) {

	var perpsMarketIdRule []interface{}
	for _, perpsMarketIdItem := range perpsMarketId {
		perpsMarketIdRule = append(perpsMarketIdRule, perpsMarketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "MarketCreated", perpsMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliMarketCreated)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseMarketCreated(log types.Log) (*PerpsMarketGoerliMarketCreated, error) {
	event := new(PerpsMarketGoerliMarketCreated)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliMarketPriceDataUpdatedIterator is returned from FilterMarketPriceDataUpdated and is used to iterate over the raw logs and unpacked data for MarketPriceDataUpdated events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMarketPriceDataUpdatedIterator struct {
	Event *PerpsMarketGoerliMarketPriceDataUpdated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliMarketPriceDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliMarketPriceDataUpdated)
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
		it.Event = new(PerpsMarketGoerliMarketPriceDataUpdated)
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
func (it *PerpsMarketGoerliMarketPriceDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliMarketPriceDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliMarketPriceDataUpdated represents a MarketPriceDataUpdated event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMarketPriceDataUpdated struct {
	MarketId *big.Int
	FeedId   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketPriceDataUpdated is a free log retrieval operation binding the contract event 0xd5ec22bbdbde803d8585a5aa89d991c67dc62c6a7bc3c4fa000b4e263b6783a0.
//
// Solidity: event MarketPriceDataUpdated(uint128 indexed marketId, bytes32 feedId)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterMarketPriceDataUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketGoerliMarketPriceDataUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "MarketPriceDataUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliMarketPriceDataUpdatedIterator{contract: _PerpsMarketGoerli.contract, event: "MarketPriceDataUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketPriceDataUpdated is a free log subscription operation binding the contract event 0xd5ec22bbdbde803d8585a5aa89d991c67dc62c6a7bc3c4fa000b4e263b6783a0.
//
// Solidity: event MarketPriceDataUpdated(uint128 indexed marketId, bytes32 feedId)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchMarketPriceDataUpdated(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliMarketPriceDataUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "MarketPriceDataUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliMarketPriceDataUpdated)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MarketPriceDataUpdated", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseMarketPriceDataUpdated(log types.Log) (*PerpsMarketGoerliMarketPriceDataUpdated, error) {
	event := new(PerpsMarketGoerliMarketPriceDataUpdated)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MarketPriceDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliMarketUpdatedIterator is returned from FilterMarketUpdated and is used to iterate over the raw logs and unpacked data for MarketUpdated events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMarketUpdatedIterator struct {
	Event *PerpsMarketGoerliMarketUpdated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliMarketUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliMarketUpdated)
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
		it.Event = new(PerpsMarketGoerliMarketUpdated)
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
func (it *PerpsMarketGoerliMarketUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliMarketUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliMarketUpdated represents a MarketUpdated event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMarketUpdated struct {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterMarketUpdated(opts *bind.FilterOpts) (*PerpsMarketGoerliMarketUpdatedIterator, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "MarketUpdated")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliMarketUpdatedIterator{contract: _PerpsMarketGoerli.contract, event: "MarketUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketUpdated is a free log subscription operation binding the contract event 0xb317f068f30db823aeb6ac6ddbffcbb6c805f558972e6b16625ec58b83f1f3d5.
//
// Solidity: event MarketUpdated(uint128 marketId, uint256 price, int256 skew, uint256 size, int256 sizeDelta, int256 currentFundingRate, int256 currentFundingVelocity)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchMarketUpdated(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliMarketUpdated) (event.Subscription, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "MarketUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliMarketUpdated)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MarketUpdated", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseMarketUpdated(log types.Log) (*PerpsMarketGoerliMarketUpdated, error) {
	event := new(PerpsMarketGoerliMarketUpdated)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MarketUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliMaxCollateralAmountSetIterator is returned from FilterMaxCollateralAmountSet and is used to iterate over the raw logs and unpacked data for MaxCollateralAmountSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMaxCollateralAmountSetIterator struct {
	Event *PerpsMarketGoerliMaxCollateralAmountSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliMaxCollateralAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliMaxCollateralAmountSet)
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
		it.Event = new(PerpsMarketGoerliMaxCollateralAmountSet)
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
func (it *PerpsMarketGoerliMaxCollateralAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliMaxCollateralAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliMaxCollateralAmountSet represents a MaxCollateralAmountSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMaxCollateralAmountSet struct {
	SynthMarketId    *big.Int
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMaxCollateralAmountSet is a free log retrieval operation binding the contract event 0xa552ba0f2552ec13b62b479a41e6e7dfcef9f10aea18527e5a3e1b3963bf70d7.
//
// Solidity: event MaxCollateralAmountSet(uint128 indexed synthMarketId, uint256 collateralAmount)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterMaxCollateralAmountSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*PerpsMarketGoerliMaxCollateralAmountSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "MaxCollateralAmountSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliMaxCollateralAmountSetIterator{contract: _PerpsMarketGoerli.contract, event: "MaxCollateralAmountSet", logs: logs, sub: sub}, nil
}

// WatchMaxCollateralAmountSet is a free log subscription operation binding the contract event 0xa552ba0f2552ec13b62b479a41e6e7dfcef9f10aea18527e5a3e1b3963bf70d7.
//
// Solidity: event MaxCollateralAmountSet(uint128 indexed synthMarketId, uint256 collateralAmount)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchMaxCollateralAmountSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliMaxCollateralAmountSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "MaxCollateralAmountSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliMaxCollateralAmountSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MaxCollateralAmountSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseMaxCollateralAmountSet(log types.Log) (*PerpsMarketGoerliMaxCollateralAmountSet, error) {
	event := new(PerpsMarketGoerliMaxCollateralAmountSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MaxCollateralAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliMaxLiquidationParametersSetIterator is returned from FilterMaxLiquidationParametersSet and is used to iterate over the raw logs and unpacked data for MaxLiquidationParametersSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMaxLiquidationParametersSetIterator struct {
	Event *PerpsMarketGoerliMaxLiquidationParametersSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliMaxLiquidationParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliMaxLiquidationParametersSet)
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
		it.Event = new(PerpsMarketGoerliMaxLiquidationParametersSet)
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
func (it *PerpsMarketGoerliMaxLiquidationParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliMaxLiquidationParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliMaxLiquidationParametersSet represents a MaxLiquidationParametersSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMaxLiquidationParametersSet struct {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterMaxLiquidationParametersSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketGoerliMaxLiquidationParametersSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "MaxLiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliMaxLiquidationParametersSetIterator{contract: _PerpsMarketGoerli.contract, event: "MaxLiquidationParametersSet", logs: logs, sub: sub}, nil
}

// WatchMaxLiquidationParametersSet is a free log subscription operation binding the contract event 0x9012ce377b7043d153d2cba3376efe7e1742af5acb7e38897362f392a7dc89ed.
//
// Solidity: event MaxLiquidationParametersSet(uint128 indexed marketId, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 maxLiquidationPd, address endorsedLiquidator)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchMaxLiquidationParametersSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliMaxLiquidationParametersSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "MaxLiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliMaxLiquidationParametersSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MaxLiquidationParametersSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseMaxLiquidationParametersSet(log types.Log) (*PerpsMarketGoerliMaxLiquidationParametersSet, error) {
	event := new(PerpsMarketGoerliMaxLiquidationParametersSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MaxLiquidationParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliMaxMarketSizeSetIterator is returned from FilterMaxMarketSizeSet and is used to iterate over the raw logs and unpacked data for MaxMarketSizeSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMaxMarketSizeSetIterator struct {
	Event *PerpsMarketGoerliMaxMarketSizeSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliMaxMarketSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliMaxMarketSizeSet)
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
		it.Event = new(PerpsMarketGoerliMaxMarketSizeSet)
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
func (it *PerpsMarketGoerliMaxMarketSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliMaxMarketSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliMaxMarketSizeSet represents a MaxMarketSizeSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliMaxMarketSizeSet struct {
	MarketId      *big.Int
	MaxMarketSize *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxMarketSizeSet is a free log retrieval operation binding the contract event 0xbd063bd3072a194b255163ab8dfd3400c4ab1cc641b920e7aaf1c11f92cd26cf.
//
// Solidity: event MaxMarketSizeSet(uint128 indexed marketId, uint256 maxMarketSize)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterMaxMarketSizeSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketGoerliMaxMarketSizeSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "MaxMarketSizeSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliMaxMarketSizeSetIterator{contract: _PerpsMarketGoerli.contract, event: "MaxMarketSizeSet", logs: logs, sub: sub}, nil
}

// WatchMaxMarketSizeSet is a free log subscription operation binding the contract event 0xbd063bd3072a194b255163ab8dfd3400c4ab1cc641b920e7aaf1c11f92cd26cf.
//
// Solidity: event MaxMarketSizeSet(uint128 indexed marketId, uint256 maxMarketSize)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchMaxMarketSizeSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliMaxMarketSizeSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "MaxMarketSizeSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliMaxMarketSizeSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MaxMarketSizeSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseMaxMarketSizeSet(log types.Log) (*PerpsMarketGoerliMaxMarketSizeSet, error) {
	event := new(PerpsMarketGoerliMaxMarketSizeSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "MaxMarketSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliOrderCommittedIterator is returned from FilterOrderCommitted and is used to iterate over the raw logs and unpacked data for OrderCommitted events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliOrderCommittedIterator struct {
	Event *PerpsMarketGoerliOrderCommitted // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliOrderCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliOrderCommitted)
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
		it.Event = new(PerpsMarketGoerliOrderCommitted)
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
func (it *PerpsMarketGoerliOrderCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliOrderCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliOrderCommitted represents a OrderCommitted event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliOrderCommitted struct {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterOrderCommitted(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*PerpsMarketGoerliOrderCommittedIterator, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "OrderCommitted", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliOrderCommittedIterator{contract: _PerpsMarketGoerli.contract, event: "OrderCommitted", logs: logs, sub: sub}, nil
}

// WatchOrderCommitted is a free log subscription operation binding the contract event 0xc74e8cf044fc959dcef2695245eb90f7e9fb5676532a71eec5b3c513d4793709.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint128 indexed accountId, uint8 orderType, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, uint256 expirationTime, bytes32 indexed trackingCode, address sender)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchOrderCommitted(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliOrderCommitted, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "OrderCommitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliOrderCommitted)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseOrderCommitted(log types.Log) (*PerpsMarketGoerliOrderCommitted, error) {
	event := new(PerpsMarketGoerliOrderCommitted)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliOrderFeesSetIterator is returned from FilterOrderFeesSet and is used to iterate over the raw logs and unpacked data for OrderFeesSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliOrderFeesSetIterator struct {
	Event *PerpsMarketGoerliOrderFeesSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliOrderFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliOrderFeesSet)
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
		it.Event = new(PerpsMarketGoerliOrderFeesSet)
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
func (it *PerpsMarketGoerliOrderFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliOrderFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliOrderFeesSet represents a OrderFeesSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliOrderFeesSet struct {
	MarketId      *big.Int
	MakerFeeRatio *big.Int
	TakerFeeRatio *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFeesSet is a free log retrieval operation binding the contract event 0x28969f156340ba9eb31589904eb174d3a4b6a37781fa6f7ad289d349d75dd1ee.
//
// Solidity: event OrderFeesSet(uint128 indexed marketId, uint256 makerFeeRatio, uint256 takerFeeRatio)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterOrderFeesSet(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketGoerliOrderFeesSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "OrderFeesSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliOrderFeesSetIterator{contract: _PerpsMarketGoerli.contract, event: "OrderFeesSet", logs: logs, sub: sub}, nil
}

// WatchOrderFeesSet is a free log subscription operation binding the contract event 0x28969f156340ba9eb31589904eb174d3a4b6a37781fa6f7ad289d349d75dd1ee.
//
// Solidity: event OrderFeesSet(uint128 indexed marketId, uint256 makerFeeRatio, uint256 takerFeeRatio)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchOrderFeesSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliOrderFeesSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "OrderFeesSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliOrderFeesSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "OrderFeesSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseOrderFeesSet(log types.Log) (*PerpsMarketGoerliOrderFeesSet, error) {
	event := new(PerpsMarketGoerliOrderFeesSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "OrderFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliOrderSettledIterator is returned from FilterOrderSettled and is used to iterate over the raw logs and unpacked data for OrderSettled events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliOrderSettledIterator struct {
	Event *PerpsMarketGoerliOrderSettled // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliOrderSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliOrderSettled)
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
		it.Event = new(PerpsMarketGoerliOrderSettled)
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
func (it *PerpsMarketGoerliOrderSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliOrderSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliOrderSettled represents a OrderSettled event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliOrderSettled struct {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterOrderSettled(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*PerpsMarketGoerliOrderSettledIterator, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "OrderSettled", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliOrderSettledIterator{contract: _PerpsMarketGoerli.contract, event: "OrderSettled", logs: logs, sub: sub}, nil
}

// WatchOrderSettled is a free log subscription operation binding the contract event 0x460080a757ec90719fe90ab2384c0196cdeed071a9fd7ce1ada43481d96b7db5.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed accountId, uint256 fillPrice, int256 pnl, int256 accruedFunding, int128 sizeDelta, int128 newSize, uint256 totalFees, uint256 referralFees, uint256 collectedFees, uint256 settlementReward, bytes32 indexed trackingCode, address settler)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchOrderSettled(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliOrderSettled, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "OrderSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliOrderSettled)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "OrderSettled", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseOrderSettled(log types.Log) (*PerpsMarketGoerliOrderSettled, error) {
	event := new(PerpsMarketGoerliOrderSettled)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "OrderSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliOwnerChangedIterator struct {
	Event *PerpsMarketGoerliOwnerChanged // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliOwnerChanged)
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
		it.Event = new(PerpsMarketGoerliOwnerChanged)
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
func (it *PerpsMarketGoerliOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliOwnerChanged represents a OwnerChanged event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*PerpsMarketGoerliOwnerChangedIterator, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliOwnerChangedIterator{contract: _PerpsMarketGoerli.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliOwnerChanged)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseOwnerChanged(log types.Log) (*PerpsMarketGoerliOwnerChanged, error) {
	event := new(PerpsMarketGoerliOwnerChanged)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliOwnerNominatedIterator struct {
	Event *PerpsMarketGoerliOwnerNominated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliOwnerNominated)
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
		it.Event = new(PerpsMarketGoerliOwnerNominated)
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
func (it *PerpsMarketGoerliOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliOwnerNominated represents a OwnerNominated event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*PerpsMarketGoerliOwnerNominatedIterator, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliOwnerNominatedIterator{contract: _PerpsMarketGoerli.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliOwnerNominated)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseOwnerNominated(log types.Log) (*PerpsMarketGoerliOwnerNominated, error) {
	event := new(PerpsMarketGoerliOwnerNominated)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliPerAccountCapsSetIterator is returned from FilterPerAccountCapsSet and is used to iterate over the raw logs and unpacked data for PerAccountCapsSet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliPerAccountCapsSetIterator struct {
	Event *PerpsMarketGoerliPerAccountCapsSet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliPerAccountCapsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliPerAccountCapsSet)
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
		it.Event = new(PerpsMarketGoerliPerAccountCapsSet)
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
func (it *PerpsMarketGoerliPerAccountCapsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliPerAccountCapsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliPerAccountCapsSet represents a PerAccountCapsSet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliPerAccountCapsSet struct {
	MaxPositionsPerAccount   *big.Int
	MaxCollateralsPerAccount *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterPerAccountCapsSet is a free log retrieval operation binding the contract event 0x3448c6d1990f2d48253b91394193cd11ce49f1653f2d03934af6d17195ffe34e.
//
// Solidity: event PerAccountCapsSet(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterPerAccountCapsSet(opts *bind.FilterOpts) (*PerpsMarketGoerliPerAccountCapsSetIterator, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "PerAccountCapsSet")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliPerAccountCapsSetIterator{contract: _PerpsMarketGoerli.contract, event: "PerAccountCapsSet", logs: logs, sub: sub}, nil
}

// WatchPerAccountCapsSet is a free log subscription operation binding the contract event 0x3448c6d1990f2d48253b91394193cd11ce49f1653f2d03934af6d17195ffe34e.
//
// Solidity: event PerAccountCapsSet(uint128 maxPositionsPerAccount, uint128 maxCollateralsPerAccount)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchPerAccountCapsSet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliPerAccountCapsSet) (event.Subscription, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "PerAccountCapsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliPerAccountCapsSet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "PerAccountCapsSet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParsePerAccountCapsSet(log types.Log) (*PerpsMarketGoerliPerAccountCapsSet, error) {
	event := new(PerpsMarketGoerliPerAccountCapsSet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "PerAccountCapsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliPermissionGrantedIterator is returned from FilterPermissionGranted and is used to iterate over the raw logs and unpacked data for PermissionGranted events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliPermissionGrantedIterator struct {
	Event *PerpsMarketGoerliPermissionGranted // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliPermissionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliPermissionGranted)
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
		it.Event = new(PerpsMarketGoerliPermissionGranted)
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
func (it *PerpsMarketGoerliPermissionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliPermissionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliPermissionGranted represents a PermissionGranted event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliPermissionGranted struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionGranted is a free log retrieval operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterPermissionGranted(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*PerpsMarketGoerliPermissionGrantedIterator, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliPermissionGrantedIterator{contract: _PerpsMarketGoerli.contract, event: "PermissionGranted", logs: logs, sub: sub}, nil
}

// WatchPermissionGranted is a free log subscription operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchPermissionGranted(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliPermissionGranted, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliPermissionGranted)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParsePermissionGranted(log types.Log) (*PerpsMarketGoerliPermissionGranted, error) {
	event := new(PerpsMarketGoerliPermissionGranted)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliPermissionRevokedIterator is returned from FilterPermissionRevoked and is used to iterate over the raw logs and unpacked data for PermissionRevoked events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliPermissionRevokedIterator struct {
	Event *PerpsMarketGoerliPermissionRevoked // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliPermissionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliPermissionRevoked)
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
		it.Event = new(PerpsMarketGoerliPermissionRevoked)
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
func (it *PerpsMarketGoerliPermissionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliPermissionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliPermissionRevoked represents a PermissionRevoked event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliPermissionRevoked struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionRevoked is a free log retrieval operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterPermissionRevoked(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*PerpsMarketGoerliPermissionRevokedIterator, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliPermissionRevokedIterator{contract: _PerpsMarketGoerli.contract, event: "PermissionRevoked", logs: logs, sub: sub}, nil
}

// WatchPermissionRevoked is a free log subscription operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchPermissionRevoked(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliPermissionRevoked, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliPermissionRevoked)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParsePermissionRevoked(log types.Log) (*PerpsMarketGoerliPermissionRevoked, error) {
	event := new(PerpsMarketGoerliPermissionRevoked)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliPositionLiquidatedIterator is returned from FilterPositionLiquidated and is used to iterate over the raw logs and unpacked data for PositionLiquidated events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliPositionLiquidatedIterator struct {
	Event *PerpsMarketGoerliPositionLiquidated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliPositionLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliPositionLiquidated)
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
		it.Event = new(PerpsMarketGoerliPositionLiquidated)
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
func (it *PerpsMarketGoerliPositionLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliPositionLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliPositionLiquidated represents a PositionLiquidated event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliPositionLiquidated struct {
	AccountId           *big.Int
	MarketId            *big.Int
	AmountLiquidated    *big.Int
	CurrentPositionSize *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPositionLiquidated is a free log retrieval operation binding the contract event 0xd583c0e2965aae317f4a9a6603c0c75602b9bc97dc7c5fc6446b0ba8d3ff1bb2.
//
// Solidity: event PositionLiquidated(uint128 indexed accountId, uint128 indexed marketId, uint256 amountLiquidated, int128 currentPositionSize)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterPositionLiquidated(opts *bind.FilterOpts, accountId []*big.Int, marketId []*big.Int) (*PerpsMarketGoerliPositionLiquidatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "PositionLiquidated", accountIdRule, marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliPositionLiquidatedIterator{contract: _PerpsMarketGoerli.contract, event: "PositionLiquidated", logs: logs, sub: sub}, nil
}

// WatchPositionLiquidated is a free log subscription operation binding the contract event 0xd583c0e2965aae317f4a9a6603c0c75602b9bc97dc7c5fc6446b0ba8d3ff1bb2.
//
// Solidity: event PositionLiquidated(uint128 indexed accountId, uint128 indexed marketId, uint256 amountLiquidated, int128 currentPositionSize)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchPositionLiquidated(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliPositionLiquidated, accountId []*big.Int, marketId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "PositionLiquidated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliPositionLiquidated)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParsePositionLiquidated(log types.Log) (*PerpsMarketGoerliPositionLiquidated, error) {
	event := new(PerpsMarketGoerliPositionLiquidated)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliPreviousOrderExpiredIterator is returned from FilterPreviousOrderExpired and is used to iterate over the raw logs and unpacked data for PreviousOrderExpired events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliPreviousOrderExpiredIterator struct {
	Event *PerpsMarketGoerliPreviousOrderExpired // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliPreviousOrderExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliPreviousOrderExpired)
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
		it.Event = new(PerpsMarketGoerliPreviousOrderExpired)
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
func (it *PerpsMarketGoerliPreviousOrderExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliPreviousOrderExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliPreviousOrderExpired represents a PreviousOrderExpired event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliPreviousOrderExpired struct {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterPreviousOrderExpired(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*PerpsMarketGoerliPreviousOrderExpiredIterator, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "PreviousOrderExpired", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliPreviousOrderExpiredIterator{contract: _PerpsMarketGoerli.contract, event: "PreviousOrderExpired", logs: logs, sub: sub}, nil
}

// WatchPreviousOrderExpired is a free log subscription operation binding the contract event 0x6d83c6751813f50325d75bc054621f83299659c5814d1e5fe6ac117860710dde.
//
// Solidity: event PreviousOrderExpired(uint128 indexed marketId, uint128 indexed accountId, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, bytes32 indexed trackingCode)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchPreviousOrderExpired(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliPreviousOrderExpired, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "PreviousOrderExpired", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliPreviousOrderExpired)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "PreviousOrderExpired", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParsePreviousOrderExpired(log types.Log) (*PerpsMarketGoerliPreviousOrderExpired, error) {
	event := new(PerpsMarketGoerliPreviousOrderExpired)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "PreviousOrderExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliReferrerShareUpdatedIterator is returned from FilterReferrerShareUpdated and is used to iterate over the raw logs and unpacked data for ReferrerShareUpdated events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliReferrerShareUpdatedIterator struct {
	Event *PerpsMarketGoerliReferrerShareUpdated // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliReferrerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliReferrerShareUpdated)
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
		it.Event = new(PerpsMarketGoerliReferrerShareUpdated)
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
func (it *PerpsMarketGoerliReferrerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliReferrerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliReferrerShareUpdated represents a ReferrerShareUpdated event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliReferrerShareUpdated struct {
	Referrer      common.Address
	ShareRatioD18 *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReferrerShareUpdated is a free log retrieval operation binding the contract event 0xa225c555f4cd21a533ad4e01eaf30153c84ca28265d954a651410d3c1e56242c.
//
// Solidity: event ReferrerShareUpdated(address referrer, uint256 shareRatioD18)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterReferrerShareUpdated(opts *bind.FilterOpts) (*PerpsMarketGoerliReferrerShareUpdatedIterator, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "ReferrerShareUpdated")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliReferrerShareUpdatedIterator{contract: _PerpsMarketGoerli.contract, event: "ReferrerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchReferrerShareUpdated is a free log subscription operation binding the contract event 0xa225c555f4cd21a533ad4e01eaf30153c84ca28265d954a651410d3c1e56242c.
//
// Solidity: event ReferrerShareUpdated(address referrer, uint256 shareRatioD18)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchReferrerShareUpdated(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliReferrerShareUpdated) (event.Subscription, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "ReferrerShareUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliReferrerShareUpdated)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseReferrerShareUpdated(log types.Log) (*PerpsMarketGoerliReferrerShareUpdated, error) {
	event := new(PerpsMarketGoerliReferrerShareUpdated)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "ReferrerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliSettlementStrategyAddedIterator is returned from FilterSettlementStrategyAdded and is used to iterate over the raw logs and unpacked data for SettlementStrategyAdded events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliSettlementStrategyAddedIterator struct {
	Event *PerpsMarketGoerliSettlementStrategyAdded // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliSettlementStrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliSettlementStrategyAdded)
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
		it.Event = new(PerpsMarketGoerliSettlementStrategyAdded)
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
func (it *PerpsMarketGoerliSettlementStrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliSettlementStrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliSettlementStrategyAdded represents a SettlementStrategyAdded event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliSettlementStrategyAdded struct {
	MarketId   *big.Int
	Strategy   SettlementStrategyData
	StrategyId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyAdded is a free log retrieval operation binding the contract event 0x6ea79b5183a0fa7ead0ea991cd5e5fa23e2da92b703bbe111479fffcdea58827.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy, uint256 indexed strategyId)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterSettlementStrategyAdded(opts *bind.FilterOpts, marketId []*big.Int, strategyId []*big.Int) (*PerpsMarketGoerliSettlementStrategyAddedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "SettlementStrategyAdded", marketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliSettlementStrategyAddedIterator{contract: _PerpsMarketGoerli.contract, event: "SettlementStrategyAdded", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyAdded is a free log subscription operation binding the contract event 0x6ea79b5183a0fa7ead0ea991cd5e5fa23e2da92b703bbe111479fffcdea58827.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy, uint256 indexed strategyId)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchSettlementStrategyAdded(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliSettlementStrategyAdded, marketId []*big.Int, strategyId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "SettlementStrategyAdded", marketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliSettlementStrategyAdded)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseSettlementStrategyAdded(log types.Log) (*PerpsMarketGoerliSettlementStrategyAdded, error) {
	event := new(PerpsMarketGoerliSettlementStrategyAdded)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliSettlementStrategyEnabledIterator is returned from FilterSettlementStrategyEnabled and is used to iterate over the raw logs and unpacked data for SettlementStrategyEnabled events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliSettlementStrategyEnabledIterator struct {
	Event *PerpsMarketGoerliSettlementStrategyEnabled // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliSettlementStrategyEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliSettlementStrategyEnabled)
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
		it.Event = new(PerpsMarketGoerliSettlementStrategyEnabled)
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
func (it *PerpsMarketGoerliSettlementStrategyEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliSettlementStrategyEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliSettlementStrategyEnabled represents a SettlementStrategyEnabled event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliSettlementStrategyEnabled struct {
	MarketId   *big.Int
	StrategyId *big.Int
	Enabled    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyEnabled is a free log retrieval operation binding the contract event 0x4803eb4e51b45b2471b0c87d8c17e0c2d8961ae080cc4a930883eab155e0fa22.
//
// Solidity: event SettlementStrategyEnabled(uint128 indexed marketId, uint256 strategyId, bool enabled)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterSettlementStrategyEnabled(opts *bind.FilterOpts, marketId []*big.Int) (*PerpsMarketGoerliSettlementStrategyEnabledIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "SettlementStrategyEnabled", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliSettlementStrategyEnabledIterator{contract: _PerpsMarketGoerli.contract, event: "SettlementStrategyEnabled", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyEnabled is a free log subscription operation binding the contract event 0x4803eb4e51b45b2471b0c87d8c17e0c2d8961ae080cc4a930883eab155e0fa22.
//
// Solidity: event SettlementStrategyEnabled(uint128 indexed marketId, uint256 strategyId, bool enabled)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchSettlementStrategyEnabled(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliSettlementStrategyEnabled, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "SettlementStrategyEnabled", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliSettlementStrategyEnabled)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "SettlementStrategyEnabled", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseSettlementStrategyEnabled(log types.Log) (*PerpsMarketGoerliSettlementStrategyEnabled, error) {
	event := new(PerpsMarketGoerliSettlementStrategyEnabled)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "SettlementStrategyEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliSynthDeductionPrioritySetIterator is returned from FilterSynthDeductionPrioritySet and is used to iterate over the raw logs and unpacked data for SynthDeductionPrioritySet events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliSynthDeductionPrioritySetIterator struct {
	Event *PerpsMarketGoerliSynthDeductionPrioritySet // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliSynthDeductionPrioritySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliSynthDeductionPrioritySet)
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
		it.Event = new(PerpsMarketGoerliSynthDeductionPrioritySet)
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
func (it *PerpsMarketGoerliSynthDeductionPrioritySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliSynthDeductionPrioritySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliSynthDeductionPrioritySet represents a SynthDeductionPrioritySet event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliSynthDeductionPrioritySet struct {
	NewSynthDeductionPriority []*big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSynthDeductionPrioritySet is a free log retrieval operation binding the contract event 0xa6beea856d32db9e9614e4af02fc7d6a3fa8337a13e4a48486c142a9a9a8ed8f.
//
// Solidity: event SynthDeductionPrioritySet(uint128[] newSynthDeductionPriority)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterSynthDeductionPrioritySet(opts *bind.FilterOpts) (*PerpsMarketGoerliSynthDeductionPrioritySetIterator, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "SynthDeductionPrioritySet")
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliSynthDeductionPrioritySetIterator{contract: _PerpsMarketGoerli.contract, event: "SynthDeductionPrioritySet", logs: logs, sub: sub}, nil
}

// WatchSynthDeductionPrioritySet is a free log subscription operation binding the contract event 0xa6beea856d32db9e9614e4af02fc7d6a3fa8337a13e4a48486c142a9a9a8ed8f.
//
// Solidity: event SynthDeductionPrioritySet(uint128[] newSynthDeductionPriority)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchSynthDeductionPrioritySet(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliSynthDeductionPrioritySet) (event.Subscription, error) {

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "SynthDeductionPrioritySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliSynthDeductionPrioritySet)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "SynthDeductionPrioritySet", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseSynthDeductionPrioritySet(log types.Log) (*PerpsMarketGoerliSynthDeductionPrioritySet, error) {
	event := new(PerpsMarketGoerliSynthDeductionPrioritySet)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "SynthDeductionPrioritySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpsMarketGoerliUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliUpgradedIterator struct {
	Event *PerpsMarketGoerliUpgraded // Event containing the contract specifics and raw log

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
func (it *PerpsMarketGoerliUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpsMarketGoerliUpgraded)
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
		it.Event = new(PerpsMarketGoerliUpgraded)
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
func (it *PerpsMarketGoerliUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpsMarketGoerliUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpsMarketGoerliUpgraded represents a Upgraded event raised by the PerpsMarketGoerli contract.
type PerpsMarketGoerliUpgraded struct {
	Self           common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) FilterUpgraded(opts *bind.FilterOpts, self []common.Address) (*PerpsMarketGoerliUpgradedIterator, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.FilterLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return &PerpsMarketGoerliUpgradedIterator{contract: _PerpsMarketGoerli.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PerpsMarketGoerliUpgraded, self []common.Address) (event.Subscription, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _PerpsMarketGoerli.contract.WatchLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpsMarketGoerliUpgraded)
				if err := _PerpsMarketGoerli.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_PerpsMarketGoerli *PerpsMarketGoerliFilterer) ParseUpgraded(log types.Log) (*PerpsMarketGoerliUpgraded, error) {
	event := new(PerpsMarketGoerliUpgraded)
	if err := _PerpsMarketGoerli.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
