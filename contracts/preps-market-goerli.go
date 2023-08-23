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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"InvalidAccountId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"InvalidPermission\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"OnlyAccountTokenProxy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"PermissionNotGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedAccountId\",\"type\":\"uint128\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountLastInteraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountPermissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"permissions\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIAccountModule.AccountPermissions[]\",\"name\":\"accountPerms\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"notifyAccountTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"renouncePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"InvalidMarket\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parameter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PerpsMarketNotInitialized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"globalPerpsMarketId\",\"type\":\"uint128\"}],\"name\":\"FactoryInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"marketSymbol\",\"type\":\"string\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedMarketId\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"marketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"marketSymbol\",\"type\":\"string\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeFactory\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"minimumCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"}],\"name\":\"reportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISpotMarketSystem\",\"name\":\"spotMarket\",\"type\":\"address\"}],\"name\":\"setSpotMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISynthetixSystem\",\"name\":\"synthetix\",\"type\":\"address\"}],\"name\":\"setSynthetix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountLiquidatable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateralAvailableForWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"}],\"name\":\"InvalidAmountDelta\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"MaxCollateralExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint128ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOrderExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"PriceFeedNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"SynthNotEnabledForCollateral\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"CollateralModified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAvailableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableMargin\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getOpenPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"totalPnl\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"accruedFunding\",\"type\":\"int256\"},{\"internalType\":\"int128\",\"name\":\"positionSize\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getRequiredMargins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredInitialMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMaintenanceMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAccumulatedLiquidationRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getWithdrawableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"withdrawableMargin\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"int256\",\"name\":\"amountDelta\",\"type\":\"int256\"}],\"name\":\"modifyCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"totalAccountOpenInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"totalCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"currentFundingRate\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"currentFundingVelocity\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"orderSize\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"fillPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"skew\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIPerpsMarketModule.MarketSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"indexPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"maxOpenInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"metadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"skew\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"}],\"name\":\"AcceptablePriceExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"availableMargin\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"minMargin\",\"type\":\"uint256\"}],\"name\":\"InsufficientMargin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"}],\"name\":\"InvalidSettlementStrategy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"newSideSize\",\"type\":\"int256\"}],\"name\":\"MaxOpenInterestReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSizeOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OrderCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"PreviousOrderExpired\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"commitment\",\"type\":\"tuple\"}],\"name\":\"commitOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"internalType\":\"structAsyncOrder.Data\",\"name\":\"retOrder\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"}],\"name\":\"computeOrderFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"settlementStrategyId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structAsyncOrder.OrderCommitmentRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"internalType\":\"structAsyncOrder.Data\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"}],\"name\":\"requiredMarginForOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredMargin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leftover\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarginError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deviation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tolerance\",\"type\":\"uint256\"}],\"name\":\"PriceDeviationToleranceExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"}],\"name\":\"SettlementStrategyNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementExpiration\",\"type\":\"uint256\"}],\"name\":\"SettlementWindowExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTime\",\"type\":\"uint256\"}],\"name\":\"SettlementWindowNotOpen\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"skew\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"sizeDelta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"}],\"name\":\"MarketUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fillPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"accruedFunding\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"sizeDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"newSize\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"}],\"name\":\"OrderSettled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"settlePythOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"NotEligibleForLiquidation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fullLiquidation\",\"type\":\"bool\"}],\"name\":\"AccountLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLiquidated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"currentPositionSize\",\"type\":\"int128\"}],\"name\":\"PositionLiquidated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidateFlagged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"name\":\"FundingParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maintenanceMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"name\":\"LiquidationParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedOiRatioD18\",\"type\":\"uint256\"}],\"name\":\"LockedOiRatioSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"MarketPriceDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"name\":\"MaxMarketSizeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerFeeRatio\",\"type\":\"uint256\"}],\"name\":\"OrderFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"SettlementStrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SettlementStrategyEnabled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"strategy\",\"type\":\"tuple\"}],\"name\":\"addSettlementStrategy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getFundingParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getLiquidationParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginScalarD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getLockedOiRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMaxMarketSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getOrderFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"getSettlementStrategy\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSettlementStrategy.Type\",\"name\":\"strategyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"settlementDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceVerificationContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"settlementReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDeviationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"internalType\":\"structSettlementStrategy.Data\",\"name\":\"settlementStrategy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"}],\"name\":\"setFundingParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumInitialMarginRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginScalarD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationLimitAccumulationMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSecondsInLiquidationWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumPositionMargin\",\"type\":\"uint256\"}],\"name\":\"setLiquidationParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"lockedOiRatioD18\",\"type\":\"uint256\"}],\"name\":\"setLockedOiRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketSize\",\"type\":\"uint256\"}],\"name\":\"setMaxMarketSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeRatio\",\"type\":\"uint256\"}],\"name\":\"setOrderFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setSettlementStrategyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"perpsMarketId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"updatePriceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidFeeCollector\",\"type\":\"address\"}],\"name\":\"InvalidFeeCollectorInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"InvalidReferrerShareRatio\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"name\":\"LiquidationRewardGuardsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"MaxCollateralAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"ReferrerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"newSynthDeductionPriority\",\"type\":\"uint128[]\"}],\"name\":\"SynthDeductionPrioritySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidationRewardGuards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"marketIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"}],\"name\":\"getMaxCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"getReferrerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSynthDeductionPriority\",\"outputs\":[{\"internalType\":\"uint128[]\",\"name\":\"\",\"type\":\"uint128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidationRewardUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationRewardUsd\",\"type\":\"uint256\"}],\"name\":\"setLiquidationRewardGuards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"synthMarketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"setMaxCollateralAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128[]\",\"name\":\"newSynthDeductionPriority\",\"type\":\"uint128[]\"}],\"name\":\"setSynthDeductionPriority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalGlobalCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shareRatioD18\",\"type\":\"uint256\"}],\"name\":\"updateReferrerShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// ComputeOrderFees is a free data retrieval call binding the contract method 0x98ef15a2.
//
// Solidity: function computeOrderFees(uint128 marketId, int128 sizeDelta) view returns(uint256 orderFees, uint256 fillPrice)
func (_Contracts *ContractsCaller) ComputeOrderFees(opts *bind.CallOpts, marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "computeOrderFees", marketId, sizeDelta)

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
func (_Contracts *ContractsSession) ComputeOrderFees(marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	return _Contracts.Contract.ComputeOrderFees(&_Contracts.CallOpts, marketId, sizeDelta)
}

// ComputeOrderFees is a free data retrieval call binding the contract method 0x98ef15a2.
//
// Solidity: function computeOrderFees(uint128 marketId, int128 sizeDelta) view returns(uint256 orderFees, uint256 fillPrice)
func (_Contracts *ContractsCallerSession) ComputeOrderFees(marketId *big.Int, sizeDelta *big.Int) (struct {
	OrderFees *big.Int
	FillPrice *big.Int
}, error) {
	return _Contracts.Contract.ComputeOrderFees(&_Contracts.CallOpts, marketId, sizeDelta)
}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_Contracts *ContractsCaller) CurrentFundingRate(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "currentFundingRate", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_Contracts *ContractsSession) CurrentFundingRate(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.CurrentFundingRate(&_Contracts.CallOpts, marketId)
}

// CurrentFundingRate is a free data retrieval call binding the contract method 0xd435b2a2.
//
// Solidity: function currentFundingRate(uint128 marketId) view returns(int256)
func (_Contracts *ContractsCallerSession) CurrentFundingRate(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.CurrentFundingRate(&_Contracts.CallOpts, marketId)
}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_Contracts *ContractsCaller) CurrentFundingVelocity(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "currentFundingVelocity", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_Contracts *ContractsSession) CurrentFundingVelocity(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.CurrentFundingVelocity(&_Contracts.CallOpts, marketId)
}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xf265db02.
//
// Solidity: function currentFundingVelocity(uint128 marketId) view returns(int256)
func (_Contracts *ContractsCallerSession) CurrentFundingVelocity(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.CurrentFundingVelocity(&_Contracts.CallOpts, marketId)
}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_Contracts *ContractsCaller) FillPrice(opts *bind.CallOpts, marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "fillPrice", marketId, orderSize, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_Contracts *ContractsSession) FillPrice(marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	return _Contracts.Contract.FillPrice(&_Contracts.CallOpts, marketId, orderSize, price)
}

// FillPrice is a free data retrieval call binding the contract method 0xdeff90ef.
//
// Solidity: function fillPrice(uint128 marketId, int128 orderSize, uint256 price) view returns(uint256)
func (_Contracts *ContractsCallerSession) FillPrice(marketId *big.Int, orderSize *big.Int, price *big.Int) (*big.Int, error) {
	return _Contracts.Contract.FillPrice(&_Contracts.CallOpts, marketId, orderSize, price)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_Contracts *ContractsCaller) GetAccountLastInteraction(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAccountLastInteraction", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_Contracts *ContractsSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetAccountLastInteraction(&_Contracts.CallOpts, accountId)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetAccountLastInteraction(&_Contracts.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_Contracts *ContractsCaller) GetAccountOwner(opts *bind.CallOpts, accountId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAccountOwner", accountId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_Contracts *ContractsSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetAccountOwner(&_Contracts.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_Contracts *ContractsCallerSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetAccountOwner(&_Contracts.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_Contracts *ContractsCaller) GetAccountPermissions(opts *bind.CallOpts, accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAccountPermissions", accountId)

	if err != nil {
		return *new([]IAccountModuleAccountPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAccountModuleAccountPermissions)).(*[]IAccountModuleAccountPermissions)

	return out0, err

}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_Contracts *ContractsSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _Contracts.Contract.GetAccountPermissions(&_Contracts.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_Contracts *ContractsCallerSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _Contracts.Contract.GetAccountPermissions(&_Contracts.CallOpts, accountId)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_Contracts *ContractsCaller) GetAccountTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAccountTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_Contracts *ContractsSession) GetAccountTokenAddress() (common.Address, error) {
	return _Contracts.Contract.GetAccountTokenAddress(&_Contracts.CallOpts)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_Contracts *ContractsCallerSession) GetAccountTokenAddress() (common.Address, error) {
	return _Contracts.Contract.GetAccountTokenAddress(&_Contracts.CallOpts)
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

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_Contracts *ContractsCaller) GetAvailableMargin(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAvailableMargin", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_Contracts *ContractsSession) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetAvailableMargin(&_Contracts.CallOpts, accountId)
}

// GetAvailableMargin is a free data retrieval call binding the contract method 0x0a7dad2d.
//
// Solidity: function getAvailableMargin(uint128 accountId) view returns(int256 availableMargin)
func (_Contracts *ContractsCallerSession) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetAvailableMargin(&_Contracts.CallOpts, accountId)
}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_Contracts *ContractsCaller) GetCollateralAmount(opts *bind.CallOpts, accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCollateralAmount", accountId, synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_Contracts *ContractsSession) GetCollateralAmount(accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetCollateralAmount(&_Contracts.CallOpts, accountId, synthMarketId)
}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x5dbd5c9b.
//
// Solidity: function getCollateralAmount(uint128 accountId, uint128 synthMarketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetCollateralAmount(accountId *big.Int, synthMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetCollateralAmount(&_Contracts.CallOpts, accountId, synthMarketId)
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

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_Contracts *ContractsCaller) GetFeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getFeeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_Contracts *ContractsSession) GetFeeCollector() (common.Address, error) {
	return _Contracts.Contract.GetFeeCollector(&_Contracts.CallOpts)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x12fde4b7.
//
// Solidity: function getFeeCollector() view returns(address feeCollector)
func (_Contracts *ContractsCallerSession) GetFeeCollector() (common.Address, error) {
	return _Contracts.Contract.GetFeeCollector(&_Contracts.CallOpts)
}

// GetFundingParameters is a free data retrieval call binding the contract method 0x1b68d8fa.
//
// Solidity: function getFundingParameters(uint128 marketId) view returns(uint256 skewScale, uint256 maxFundingVelocity)
func (_Contracts *ContractsCaller) GetFundingParameters(opts *bind.CallOpts, marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getFundingParameters", marketId)

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
func (_Contracts *ContractsSession) GetFundingParameters(marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	return _Contracts.Contract.GetFundingParameters(&_Contracts.CallOpts, marketId)
}

// GetFundingParameters is a free data retrieval call binding the contract method 0x1b68d8fa.
//
// Solidity: function getFundingParameters(uint128 marketId) view returns(uint256 skewScale, uint256 maxFundingVelocity)
func (_Contracts *ContractsCallerSession) GetFundingParameters(marketId *big.Int) (struct {
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
}, error) {
	return _Contracts.Contract.GetFundingParameters(&_Contracts.CallOpts, marketId)
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

// GetLiquidationParameters is a free data retrieval call binding the contract method 0xf94363a6.
//
// Solidity: function getLiquidationParameters(uint128 marketId) view returns(uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin)
func (_Contracts *ContractsCaller) GetLiquidationParameters(opts *bind.CallOpts, marketId *big.Int) (struct {
	InitialMarginRatioD18                     *big.Int
	MinimumInitialMarginRatioD18              *big.Int
	MaintenanceMarginScalarD18                *big.Int
	LiquidationRewardRatioD18                 *big.Int
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MinimumPositionMargin                     *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLiquidationParameters", marketId)

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
func (_Contracts *ContractsSession) GetLiquidationParameters(marketId *big.Int) (struct {
	InitialMarginRatioD18                     *big.Int
	MinimumInitialMarginRatioD18              *big.Int
	MaintenanceMarginScalarD18                *big.Int
	LiquidationRewardRatioD18                 *big.Int
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MinimumPositionMargin                     *big.Int
}, error) {
	return _Contracts.Contract.GetLiquidationParameters(&_Contracts.CallOpts, marketId)
}

// GetLiquidationParameters is a free data retrieval call binding the contract method 0xf94363a6.
//
// Solidity: function getLiquidationParameters(uint128 marketId) view returns(uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin)
func (_Contracts *ContractsCallerSession) GetLiquidationParameters(marketId *big.Int) (struct {
	InitialMarginRatioD18                     *big.Int
	MinimumInitialMarginRatioD18              *big.Int
	MaintenanceMarginScalarD18                *big.Int
	LiquidationRewardRatioD18                 *big.Int
	MaxLiquidationLimitAccumulationMultiplier *big.Int
	MaxSecondsInLiquidationWindow             *big.Int
	MinimumPositionMargin                     *big.Int
}, error) {
	return _Contracts.Contract.GetLiquidationParameters(&_Contracts.CallOpts, marketId)
}

// GetLiquidationRewardGuards is a free data retrieval call binding the contract method 0x0b7f4b2d.
//
// Solidity: function getLiquidationRewardGuards() view returns(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd)
func (_Contracts *ContractsCaller) GetLiquidationRewardGuards(opts *bind.CallOpts) (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLiquidationRewardGuards")

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
func (_Contracts *ContractsSession) GetLiquidationRewardGuards() (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	return _Contracts.Contract.GetLiquidationRewardGuards(&_Contracts.CallOpts)
}

// GetLiquidationRewardGuards is a free data retrieval call binding the contract method 0x0b7f4b2d.
//
// Solidity: function getLiquidationRewardGuards() view returns(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd)
func (_Contracts *ContractsCallerSession) GetLiquidationRewardGuards() (struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
}, error) {
	return _Contracts.Contract.GetLiquidationRewardGuards(&_Contracts.CallOpts)
}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCaller) GetLockedOiRatio(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLockedOiRatio", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsSession) GetLockedOiRatio(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetLockedOiRatio(&_Contracts.CallOpts, marketId)
}

// GetLockedOiRatio is a free data retrieval call binding the contract method 0x31edc046.
//
// Solidity: function getLockedOiRatio(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetLockedOiRatio(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetLockedOiRatio(&_Contracts.CallOpts, marketId)
}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_Contracts *ContractsCaller) GetMarketSummary(opts *bind.CallOpts, marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketSummary", marketId)

	if err != nil {
		return *new(IPerpsMarketModuleMarketSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(IPerpsMarketModuleMarketSummary)).(*IPerpsMarketModuleMarketSummary)

	return out0, err

}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_Contracts *ContractsSession) GetMarketSummary(marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	return _Contracts.Contract.GetMarketSummary(&_Contracts.CallOpts, marketId)
}

// GetMarketSummary is a free data retrieval call binding the contract method 0x41c2e8bd.
//
// Solidity: function getMarketSummary(uint128 marketId) view returns((int256,uint256,uint256,int256,int256,uint256) summary)
func (_Contracts *ContractsCallerSession) GetMarketSummary(marketId *big.Int) (IPerpsMarketModuleMarketSummary, error) {
	return _Contracts.Contract.GetMarketSummary(&_Contracts.CallOpts, marketId)
}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_Contracts *ContractsCaller) GetMarkets(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarkets")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_Contracts *ContractsSession) GetMarkets() ([]*big.Int, error) {
	return _Contracts.Contract.GetMarkets(&_Contracts.CallOpts)
}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns(uint256[] marketIds)
func (_Contracts *ContractsCallerSession) GetMarkets() ([]*big.Int, error) {
	return _Contracts.Contract.GetMarkets(&_Contracts.CallOpts)
}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_Contracts *ContractsCaller) GetMaxCollateralAmount(opts *bind.CallOpts, synthMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMaxCollateralAmount", synthMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_Contracts *ContractsSession) GetMaxCollateralAmount(synthMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMaxCollateralAmount(&_Contracts.CallOpts, synthMarketId)
}

// GetMaxCollateralAmount is a free data retrieval call binding the contract method 0x4ff68ae2.
//
// Solidity: function getMaxCollateralAmount(uint128 synthMarketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetMaxCollateralAmount(synthMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMaxCollateralAmount(&_Contracts.CallOpts, synthMarketId)
}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_Contracts *ContractsCaller) GetMaxMarketSize(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMaxMarketSize", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_Contracts *ContractsSession) GetMaxMarketSize(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMaxMarketSize(&_Contracts.CallOpts, marketId)
}

// GetMaxMarketSize is a free data retrieval call binding the contract method 0x19a99bf5.
//
// Solidity: function getMaxMarketSize(uint128 marketId) view returns(uint256 maxMarketSize)
func (_Contracts *ContractsCallerSession) GetMaxMarketSize(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMaxMarketSize(&_Contracts.CallOpts, marketId)
}

// GetOpenPosition is a free data retrieval call binding the contract method 0x22a73967.
//
// Solidity: function getOpenPosition(uint128 accountId, uint128 marketId) view returns(int256 totalPnl, int256 accruedFunding, int128 positionSize)
func (_Contracts *ContractsCaller) GetOpenPosition(opts *bind.CallOpts, accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOpenPosition", accountId, marketId)

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
func (_Contracts *ContractsSession) GetOpenPosition(accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	return _Contracts.Contract.GetOpenPosition(&_Contracts.CallOpts, accountId, marketId)
}

// GetOpenPosition is a free data retrieval call binding the contract method 0x22a73967.
//
// Solidity: function getOpenPosition(uint128 accountId, uint128 marketId) view returns(int256 totalPnl, int256 accruedFunding, int128 positionSize)
func (_Contracts *ContractsCallerSession) GetOpenPosition(accountId *big.Int, marketId *big.Int) (struct {
	TotalPnl       *big.Int
	AccruedFunding *big.Int
	PositionSize   *big.Int
}, error) {
	return _Contracts.Contract.GetOpenPosition(&_Contracts.CallOpts, accountId, marketId)
}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_Contracts *ContractsCaller) GetOrder(opts *bind.CallOpts, accountId *big.Int) (AsyncOrderData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOrder", accountId)

	if err != nil {
		return *new(AsyncOrderData), err
	}

	out0 := *abi.ConvertType(out[0], new(AsyncOrderData)).(*AsyncOrderData)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_Contracts *ContractsSession) GetOrder(accountId *big.Int) (AsyncOrderData, error) {
	return _Contracts.Contract.GetOrder(&_Contracts.CallOpts, accountId)
}

// GetOrder is a free data retrieval call binding the contract method 0x117d4128.
//
// Solidity: function getOrder(uint128 accountId) view returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) order)
func (_Contracts *ContractsCallerSession) GetOrder(accountId *big.Int) (AsyncOrderData, error) {
	return _Contracts.Contract.GetOrder(&_Contracts.CallOpts, accountId)
}

// GetOrderFees is a free data retrieval call binding the contract method 0xaac23e8c.
//
// Solidity: function getOrderFees(uint128 marketId) view returns(uint256 makerFee, uint256 takerFee)
func (_Contracts *ContractsCaller) GetOrderFees(opts *bind.CallOpts, marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOrderFees", marketId)

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
func (_Contracts *ContractsSession) GetOrderFees(marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	return _Contracts.Contract.GetOrderFees(&_Contracts.CallOpts, marketId)
}

// GetOrderFees is a free data retrieval call binding the contract method 0xaac23e8c.
//
// Solidity: function getOrderFees(uint128 marketId) view returns(uint256 makerFee, uint256 takerFee)
func (_Contracts *ContractsCallerSession) GetOrderFees(marketId *big.Int) (struct {
	MakerFee *big.Int
	TakerFee *big.Int
}, error) {
	return _Contracts.Contract.GetOrderFees(&_Contracts.CallOpts, marketId)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_Contracts *ContractsCaller) GetReferrerShare(opts *bind.CallOpts, referrer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getReferrerShare", referrer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_Contracts *ContractsSession) GetReferrerShare(referrer common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetReferrerShare(&_Contracts.CallOpts, referrer)
}

// GetReferrerShare is a free data retrieval call binding the contract method 0xcae77b70.
//
// Solidity: function getReferrerShare(address referrer) view returns(uint256 shareRatioD18)
func (_Contracts *ContractsCallerSession) GetReferrerShare(referrer common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetReferrerShare(&_Contracts.CallOpts, referrer)
}

// GetRequiredMargins is a free data retrieval call binding the contract method 0x3c0f0753.
//
// Solidity: function getRequiredMargins(uint128 accountId) view returns(uint256 requiredInitialMargin, uint256 requiredMaintenanceMargin, uint256 totalAccumulatedLiquidationRewards, uint256 maxLiquidationReward)
func (_Contracts *ContractsCaller) GetRequiredMargins(opts *bind.CallOpts, accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getRequiredMargins", accountId)

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
func (_Contracts *ContractsSession) GetRequiredMargins(accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	return _Contracts.Contract.GetRequiredMargins(&_Contracts.CallOpts, accountId)
}

// GetRequiredMargins is a free data retrieval call binding the contract method 0x3c0f0753.
//
// Solidity: function getRequiredMargins(uint128 accountId) view returns(uint256 requiredInitialMargin, uint256 requiredMaintenanceMargin, uint256 totalAccumulatedLiquidationRewards, uint256 maxLiquidationReward)
func (_Contracts *ContractsCallerSession) GetRequiredMargins(accountId *big.Int) (struct {
	RequiredInitialMargin              *big.Int
	RequiredMaintenanceMargin          *big.Int
	TotalAccumulatedLiquidationRewards *big.Int
	MaxLiquidationReward               *big.Int
}, error) {
	return _Contracts.Contract.GetRequiredMargins(&_Contracts.CallOpts, accountId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) settlementStrategy)
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
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) settlementStrategy)
func (_Contracts *ContractsSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _Contracts.Contract.GetSettlementStrategy(&_Contracts.CallOpts, marketId, strategyId)
}

// GetSettlementStrategy is a free data retrieval call binding the contract method 0xf74c377f.
//
// Solidity: function getSettlementStrategy(uint128 marketId, uint256 strategyId) view returns((uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) settlementStrategy)
func (_Contracts *ContractsCallerSession) GetSettlementStrategy(marketId *big.Int, strategyId *big.Int) (SettlementStrategyData, error) {
	return _Contracts.Contract.GetSettlementStrategy(&_Contracts.CallOpts, marketId, strategyId)
}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_Contracts *ContractsCaller) GetSynthDeductionPriority(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getSynthDeductionPriority")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_Contracts *ContractsSession) GetSynthDeductionPriority() ([]*big.Int, error) {
	return _Contracts.Contract.GetSynthDeductionPriority(&_Contracts.CallOpts)
}

// GetSynthDeductionPriority is a free data retrieval call binding the contract method 0xfea84a3f.
//
// Solidity: function getSynthDeductionPriority() view returns(uint128[])
func (_Contracts *ContractsCallerSession) GetSynthDeductionPriority() ([]*big.Int, error) {
	return _Contracts.Contract.GetSynthDeductionPriority(&_Contracts.CallOpts)
}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_Contracts *ContractsCaller) GetWithdrawableMargin(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getWithdrawableMargin", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_Contracts *ContractsSession) GetWithdrawableMargin(accountId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetWithdrawableMargin(&_Contracts.CallOpts, accountId)
}

// GetWithdrawableMargin is a free data retrieval call binding the contract method 0x04aa363e.
//
// Solidity: function getWithdrawableMargin(uint128 accountId) view returns(int256 withdrawableMargin)
func (_Contracts *ContractsCallerSession) GetWithdrawableMargin(accountId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetWithdrawableMargin(&_Contracts.CallOpts, accountId)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Contracts *ContractsCaller) HasPermission(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "hasPermission", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Contracts *ContractsSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _Contracts.Contract.HasPermission(&_Contracts.CallOpts, accountId, permission, user)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Contracts *ContractsCallerSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _Contracts.Contract.HasPermission(&_Contracts.CallOpts, accountId, permission, user)
}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCaller) IndexPrice(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "indexPrice", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsSession) IndexPrice(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.IndexPrice(&_Contracts.CallOpts, marketId)
}

// IndexPrice is a free data retrieval call binding the contract method 0x4f778fb4.
//
// Solidity: function indexPrice(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) IndexPrice(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.IndexPrice(&_Contracts.CallOpts, marketId)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Contracts *ContractsCaller) IsAuthorized(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isAuthorized", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Contracts *ContractsSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _Contracts.Contract.IsAuthorized(&_Contracts.CallOpts, accountId, permission, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Contracts *ContractsCallerSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _Contracts.Contract.IsAuthorized(&_Contracts.CallOpts, accountId, permission, user)
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

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCaller) MaxOpenInterest(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "maxOpenInterest", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsSession) MaxOpenInterest(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.MaxOpenInterest(&_Contracts.CallOpts, marketId)
}

// MaxOpenInterest is a free data retrieval call binding the contract method 0x0e7cace9.
//
// Solidity: function maxOpenInterest(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) MaxOpenInterest(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.MaxOpenInterest(&_Contracts.CallOpts, marketId)
}

// Metadata is a free data retrieval call binding the contract method 0xe3bc36bf.
//
// Solidity: function metadata(uint128 marketId) view returns(string name, string symbol)
func (_Contracts *ContractsCaller) Metadata(opts *bind.CallOpts, marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "metadata", marketId)

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
func (_Contracts *ContractsSession) Metadata(marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	return _Contracts.Contract.Metadata(&_Contracts.CallOpts, marketId)
}

// Metadata is a free data retrieval call binding the contract method 0xe3bc36bf.
//
// Solidity: function metadata(uint128 marketId) view returns(string name, string symbol)
func (_Contracts *ContractsCallerSession) Metadata(marketId *big.Int) (struct {
	Name   string
	Symbol string
}, error) {
	return _Contracts.Contract.Metadata(&_Contracts.CallOpts, marketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_Contracts *ContractsCaller) MinimumCredit(opts *bind.CallOpts, perpsMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "minimumCredit", perpsMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_Contracts *ContractsSession) MinimumCredit(perpsMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.MinimumCredit(&_Contracts.CallOpts, perpsMarketId)
}

// MinimumCredit is a free data retrieval call binding the contract method 0xafe79200.
//
// Solidity: function minimumCredit(uint128 perpsMarketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) MinimumCredit(perpsMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.MinimumCredit(&_Contracts.CallOpts, perpsMarketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts, perpsMarketId *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name", perpsMarketId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_Contracts *ContractsSession) Name(perpsMarketId *big.Int) (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts, perpsMarketId)
}

// Name is a free data retrieval call binding the contract method 0xc624440a.
//
// Solidity: function name(uint128 perpsMarketId) view returns(string)
func (_Contracts *ContractsCallerSession) Name(perpsMarketId *big.Int) (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts, perpsMarketId)
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

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_Contracts *ContractsCaller) ReportedDebt(opts *bind.CallOpts, perpsMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "reportedDebt", perpsMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_Contracts *ContractsSession) ReportedDebt(perpsMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ReportedDebt(&_Contracts.CallOpts, perpsMarketId)
}

// ReportedDebt is a free data retrieval call binding the contract method 0xbcec0d0f.
//
// Solidity: function reportedDebt(uint128 perpsMarketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) ReportedDebt(perpsMarketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ReportedDebt(&_Contracts.CallOpts, perpsMarketId)
}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_Contracts *ContractsCaller) RequiredMarginForOrder(opts *bind.CallOpts, accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "requiredMarginForOrder", accountId, marketId, sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_Contracts *ContractsSession) RequiredMarginForOrder(accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	return _Contracts.Contract.RequiredMarginForOrder(&_Contracts.CallOpts, accountId, marketId, sizeDelta)
}

// RequiredMarginForOrder is a free data retrieval call binding the contract method 0xb8830a25.
//
// Solidity: function requiredMarginForOrder(uint128 accountId, uint128 marketId, int128 sizeDelta) view returns(uint256 requiredMargin)
func (_Contracts *ContractsCallerSession) RequiredMarginForOrder(accountId *big.Int, marketId *big.Int, sizeDelta *big.Int) (*big.Int, error) {
	return _Contracts.Contract.RequiredMarginForOrder(&_Contracts.CallOpts, accountId, marketId, sizeDelta)
}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_Contracts *ContractsCaller) Settle(opts *bind.CallOpts, accountId *big.Int) error {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "settle", accountId)

	if err != nil {
		return err
	}

	return err

}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_Contracts *ContractsSession) Settle(accountId *big.Int) error {
	return _Contracts.Contract.Settle(&_Contracts.CallOpts, accountId)
}

// Settle is a free data retrieval call binding the contract method 0x895e3bed.
//
// Solidity: function settle(uint128 accountId) view returns()
func (_Contracts *ContractsCallerSession) Settle(accountId *big.Int) error {
	return _Contracts.Contract.Settle(&_Contracts.CallOpts, accountId)
}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCaller) Size(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "size", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsSession) Size(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Size(&_Contracts.CallOpts, marketId)
}

// Size is a free data retrieval call binding the contract method 0x2b267635.
//
// Solidity: function size(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) Size(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Size(&_Contracts.CallOpts, marketId)
}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_Contracts *ContractsCaller) Skew(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "skew", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_Contracts *ContractsSession) Skew(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Skew(&_Contracts.CallOpts, marketId)
}

// Skew is a free data retrieval call binding the contract method 0x83a7db27.
//
// Solidity: function skew(uint128 marketId) view returns(int256)
func (_Contracts *ContractsCallerSession) Skew(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Skew(&_Contracts.CallOpts, marketId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
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
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_Contracts *ContractsCaller) TotalAccountOpenInterest(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalAccountOpenInterest", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_Contracts *ContractsSession) TotalAccountOpenInterest(accountId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TotalAccountOpenInterest(&_Contracts.CallOpts, accountId)
}

// TotalAccountOpenInterest is a free data retrieval call binding the contract method 0x2daf43bc.
//
// Solidity: function totalAccountOpenInterest(uint128 accountId) view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalAccountOpenInterest(accountId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TotalAccountOpenInterest(&_Contracts.CallOpts, accountId)
}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_Contracts *ContractsCaller) TotalCollateralValue(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalCollateralValue", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_Contracts *ContractsSession) TotalCollateralValue(accountId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TotalCollateralValue(&_Contracts.CallOpts, accountId)
}

// TotalCollateralValue is a free data retrieval call binding the contract method 0xb568ae42.
//
// Solidity: function totalCollateralValue(uint128 accountId) view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalCollateralValue(accountId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TotalCollateralValue(&_Contracts.CallOpts, accountId)
}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_Contracts *ContractsCaller) TotalGlobalCollateralValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalGlobalCollateralValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_Contracts *ContractsSession) TotalGlobalCollateralValue() (*big.Int, error) {
	return _Contracts.Contract.TotalGlobalCollateralValue(&_Contracts.CallOpts)
}

// TotalGlobalCollateralValue is a free data retrieval call binding the contract method 0x65c5a0fe.
//
// Solidity: function totalGlobalCollateralValue() view returns(uint256 totalCollateralValue)
func (_Contracts *ContractsCallerSession) TotalGlobalCollateralValue() (*big.Int, error) {
	return _Contracts.Contract.TotalGlobalCollateralValue(&_Contracts.CallOpts)
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

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0xda91187c.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_Contracts *ContractsTransactor) AddSettlementStrategy(opts *bind.TransactOpts, marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addSettlementStrategy", marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0xda91187c.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy) returns(uint256 strategyId)
func (_Contracts *ContractsSession) AddSettlementStrategy(marketId *big.Int, strategy SettlementStrategyData) (*types.Transaction, error) {
	return _Contracts.Contract.AddSettlementStrategy(&_Contracts.TransactOpts, marketId, strategy)
}

// AddSettlementStrategy is a paid mutator transaction binding the contract method 0xda91187c.
//
// Solidity: function addSettlementStrategy(uint128 marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy) returns(uint256 strategyId)
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

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_Contracts *ContractsTransactor) CommitOrder(opts *bind.TransactOpts, commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "commitOrder", commitment)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_Contracts *ContractsSession) CommitOrder(commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _Contracts.Contract.CommitOrder(&_Contracts.TransactOpts, commitment)
}

// CommitOrder is a paid mutator transaction binding the contract method 0x9f978860.
//
// Solidity: function commitOrder((uint128,uint128,int128,uint128,uint256,bytes32,address) commitment) returns((uint256,(uint128,uint128,int128,uint128,uint256,bytes32,address)) retOrder, uint256 fees)
func (_Contracts *ContractsTransactorSession) CommitOrder(commitment AsyncOrderOrderCommitmentRequest) (*types.Transaction, error) {
	return _Contracts.Contract.CommitOrder(&_Contracts.TransactOpts, commitment)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_Contracts *ContractsTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_Contracts *ContractsSession) CreateAccount() (*types.Transaction, error) {
	return _Contracts.Contract.CreateAccount(&_Contracts.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_Contracts *ContractsTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _Contracts.Contract.CreateAccount(&_Contracts.TransactOpts)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_Contracts *ContractsTransactor) CreateAccount0(opts *bind.TransactOpts, requestedAccountId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createAccount0", requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_Contracts *ContractsSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateAccount0(&_Contracts.TransactOpts, requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_Contracts *ContractsTransactorSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateAccount0(&_Contracts.TransactOpts, requestedAccountId)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_Contracts *ContractsTransactor) CreateMarket(opts *bind.TransactOpts, requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createMarket", requestedMarketId, marketName, marketSymbol)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_Contracts *ContractsSession) CreateMarket(requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _Contracts.Contract.CreateMarket(&_Contracts.TransactOpts, requestedMarketId, marketName, marketSymbol)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7e947ea4.
//
// Solidity: function createMarket(uint128 requestedMarketId, string marketName, string marketSymbol) returns(uint128)
func (_Contracts *ContractsTransactorSession) CreateMarket(requestedMarketId *big.Int, marketName string, marketSymbol string) (*types.Transaction, error) {
	return _Contracts.Contract.CreateMarket(&_Contracts.TransactOpts, requestedMarketId, marketName, marketSymbol)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Contracts *ContractsTransactor) GrantPermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "grantPermission", accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Contracts *ContractsSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GrantPermission(&_Contracts.TransactOpts, accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Contracts *ContractsTransactorSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GrantPermission(&_Contracts.TransactOpts, accountId, permission, user)
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

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_Contracts *ContractsTransactor) InitializeFactory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initializeFactory")
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_Contracts *ContractsSession) InitializeFactory() (*types.Transaction, error) {
	return _Contracts.Contract.InitializeFactory(&_Contracts.TransactOpts)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0xd95de76b.
//
// Solidity: function initializeFactory() returns(uint128)
func (_Contracts *ContractsTransactorSession) InitializeFactory() (*types.Transaction, error) {
	return _Contracts.Contract.InitializeFactory(&_Contracts.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns()
func (_Contracts *ContractsTransactor) Liquidate(opts *bind.TransactOpts, accountId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "liquidate", accountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns()
func (_Contracts *ContractsSession) Liquidate(accountId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Liquidate(&_Contracts.TransactOpts, accountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x048577de.
//
// Solidity: function liquidate(uint128 accountId) returns()
func (_Contracts *ContractsTransactorSession) Liquidate(accountId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Liquidate(&_Contracts.TransactOpts, accountId)
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns()
func (_Contracts *ContractsTransactor) LiquidateFlagged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "liquidateFlagged")
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns()
func (_Contracts *ContractsSession) LiquidateFlagged() (*types.Transaction, error) {
	return _Contracts.Contract.LiquidateFlagged(&_Contracts.TransactOpts)
}

// LiquidateFlagged is a paid mutator transaction binding the contract method 0x1d6d458c.
//
// Solidity: function liquidateFlagged() returns()
func (_Contracts *ContractsTransactorSession) LiquidateFlagged() (*types.Transaction, error) {
	return _Contracts.Contract.LiquidateFlagged(&_Contracts.TransactOpts)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_Contracts *ContractsTransactor) ModifyCollateral(opts *bind.TransactOpts, accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "modifyCollateral", accountId, synthMarketId, amountDelta)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_Contracts *ContractsSession) ModifyCollateral(accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ModifyCollateral(&_Contracts.TransactOpts, accountId, synthMarketId, amountDelta)
}

// ModifyCollateral is a paid mutator transaction binding the contract method 0xbb58672c.
//
// Solidity: function modifyCollateral(uint128 accountId, uint128 synthMarketId, int256 amountDelta) returns()
func (_Contracts *ContractsTransactorSession) ModifyCollateral(accountId *big.Int, synthMarketId *big.Int, amountDelta *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ModifyCollateral(&_Contracts.TransactOpts, accountId, synthMarketId, amountDelta)
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

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_Contracts *ContractsTransactor) NotifyAccountTransfer(opts *bind.TransactOpts, to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "notifyAccountTransfer", to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_Contracts *ContractsSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NotifyAccountTransfer(&_Contracts.TransactOpts, to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_Contracts *ContractsTransactorSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NotifyAccountTransfer(&_Contracts.TransactOpts, to, accountId)
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

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_Contracts *ContractsTransactor) RenouncePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renouncePermission", accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_Contracts *ContractsSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.RenouncePermission(&_Contracts.TransactOpts, accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_Contracts *ContractsTransactorSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.RenouncePermission(&_Contracts.TransactOpts, accountId, permission)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Contracts *ContractsTransactor) RevokePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revokePermission", accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Contracts *ContractsSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokePermission(&_Contracts.TransactOpts, accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Contracts *ContractsTransactorSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokePermission(&_Contracts.TransactOpts, accountId, permission, user)
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

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_Contracts *ContractsTransactor) SetFeeCollector(opts *bind.TransactOpts, feeCollector common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setFeeCollector", feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_Contracts *ContractsSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetFeeCollector(&_Contracts.TransactOpts, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_Contracts *ContractsTransactorSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetFeeCollector(&_Contracts.TransactOpts, feeCollector)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_Contracts *ContractsTransactor) SetFundingParameters(opts *bind.TransactOpts, marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setFundingParameters", marketId, skewScale, maxFundingVelocity)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_Contracts *ContractsSession) SetFundingParameters(marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetFundingParameters(&_Contracts.TransactOpts, marketId, skewScale, maxFundingVelocity)
}

// SetFundingParameters is a paid mutator transaction binding the contract method 0xc2382277.
//
// Solidity: function setFundingParameters(uint128 marketId, uint256 skewScale, uint256 maxFundingVelocity) returns()
func (_Contracts *ContractsTransactorSession) SetFundingParameters(marketId *big.Int, skewScale *big.Int, maxFundingVelocity *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetFundingParameters(&_Contracts.TransactOpts, marketId, skewScale, maxFundingVelocity)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0xcb450077.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin) returns()
func (_Contracts *ContractsTransactor) SetLiquidationParameters(opts *bind.TransactOpts, marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setLiquidationParameters", marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, minimumPositionMargin)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0xcb450077.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin) returns()
func (_Contracts *ContractsSession) SetLiquidationParameters(marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetLiquidationParameters(&_Contracts.TransactOpts, marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, minimumPositionMargin)
}

// SetLiquidationParameters is a paid mutator transaction binding the contract method 0xcb450077.
//
// Solidity: function setLiquidationParameters(uint128 marketId, uint256 initialMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 maintenanceMarginScalarD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin) returns()
func (_Contracts *ContractsTransactorSession) SetLiquidationParameters(marketId *big.Int, initialMarginRatioD18 *big.Int, minimumInitialMarginRatioD18 *big.Int, maintenanceMarginScalarD18 *big.Int, liquidationRewardRatioD18 *big.Int, maxLiquidationLimitAccumulationMultiplier *big.Int, maxSecondsInLiquidationWindow *big.Int, minimumPositionMargin *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetLiquidationParameters(&_Contracts.TransactOpts, marketId, initialMarginRatioD18, minimumInitialMarginRatioD18, maintenanceMarginScalarD18, liquidationRewardRatioD18, maxLiquidationLimitAccumulationMultiplier, maxSecondsInLiquidationWindow, minimumPositionMargin)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_Contracts *ContractsTransactor) SetLiquidationRewardGuards(opts *bind.TransactOpts, minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setLiquidationRewardGuards", minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_Contracts *ContractsSession) SetLiquidationRewardGuards(minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetLiquidationRewardGuards(&_Contracts.TransactOpts, minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLiquidationRewardGuards is a paid mutator transaction binding the contract method 0xdbc593a9.
//
// Solidity: function setLiquidationRewardGuards(uint256 minLiquidationRewardUsd, uint256 maxLiquidationRewardUsd) returns()
func (_Contracts *ContractsTransactorSession) SetLiquidationRewardGuards(minLiquidationRewardUsd *big.Int, maxLiquidationRewardUsd *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetLiquidationRewardGuards(&_Contracts.TransactOpts, minLiquidationRewardUsd, maxLiquidationRewardUsd)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_Contracts *ContractsTransactor) SetLockedOiRatio(opts *bind.TransactOpts, marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setLockedOiRatio", marketId, lockedOiRatioD18)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_Contracts *ContractsSession) SetLockedOiRatio(marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetLockedOiRatio(&_Contracts.TransactOpts, marketId, lockedOiRatioD18)
}

// SetLockedOiRatio is a paid mutator transaction binding the contract method 0x033723d9.
//
// Solidity: function setLockedOiRatio(uint128 marketId, uint256 lockedOiRatioD18) returns()
func (_Contracts *ContractsTransactorSession) SetLockedOiRatio(marketId *big.Int, lockedOiRatioD18 *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetLockedOiRatio(&_Contracts.TransactOpts, marketId, lockedOiRatioD18)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_Contracts *ContractsTransactor) SetMaxCollateralAmount(opts *bind.TransactOpts, synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMaxCollateralAmount", synthMarketId, collateralAmount)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_Contracts *ContractsSession) SetMaxCollateralAmount(synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMaxCollateralAmount(&_Contracts.TransactOpts, synthMarketId, collateralAmount)
}

// SetMaxCollateralAmount is a paid mutator transaction binding the contract method 0x6cded665.
//
// Solidity: function setMaxCollateralAmount(uint128 synthMarketId, uint256 collateralAmount) returns()
func (_Contracts *ContractsTransactorSession) SetMaxCollateralAmount(synthMarketId *big.Int, collateralAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMaxCollateralAmount(&_Contracts.TransactOpts, synthMarketId, collateralAmount)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_Contracts *ContractsTransactor) SetMaxMarketSize(opts *bind.TransactOpts, marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMaxMarketSize", marketId, maxMarketSize)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_Contracts *ContractsSession) SetMaxMarketSize(marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMaxMarketSize(&_Contracts.TransactOpts, marketId, maxMarketSize)
}

// SetMaxMarketSize is a paid mutator transaction binding the contract method 0x404a68aa.
//
// Solidity: function setMaxMarketSize(uint128 marketId, uint256 maxMarketSize) returns()
func (_Contracts *ContractsTransactorSession) SetMaxMarketSize(marketId *big.Int, maxMarketSize *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMaxMarketSize(&_Contracts.TransactOpts, marketId, maxMarketSize)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_Contracts *ContractsTransactor) SetOrderFees(opts *bind.TransactOpts, marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setOrderFees", marketId, makerFeeRatio, takerFeeRatio)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_Contracts *ContractsSession) SetOrderFees(marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetOrderFees(&_Contracts.TransactOpts, marketId, makerFeeRatio, takerFeeRatio)
}

// SetOrderFees is a paid mutator transaction binding the contract method 0xf842fa86.
//
// Solidity: function setOrderFees(uint128 marketId, uint256 makerFeeRatio, uint256 takerFeeRatio) returns()
func (_Contracts *ContractsTransactorSession) SetOrderFees(marketId *big.Int, makerFeeRatio *big.Int, takerFeeRatio *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetOrderFees(&_Contracts.TransactOpts, marketId, makerFeeRatio, takerFeeRatio)
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

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_Contracts *ContractsTransactor) SetSpotMarket(opts *bind.TransactOpts, spotMarket common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setSpotMarket", spotMarket)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_Contracts *ContractsSession) SetSpotMarket(spotMarket common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetSpotMarket(&_Contracts.TransactOpts, spotMarket)
}

// SetSpotMarket is a paid mutator transaction binding the contract method 0x92d48a4e.
//
// Solidity: function setSpotMarket(address spotMarket) returns()
func (_Contracts *ContractsTransactorSession) SetSpotMarket(spotMarket common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetSpotMarket(&_Contracts.TransactOpts, spotMarket)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_Contracts *ContractsTransactor) SetSynthDeductionPriority(opts *bind.TransactOpts, newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setSynthDeductionPriority", newSynthDeductionPriority)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_Contracts *ContractsSession) SetSynthDeductionPriority(newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetSynthDeductionPriority(&_Contracts.TransactOpts, newSynthDeductionPriority)
}

// SetSynthDeductionPriority is a paid mutator transaction binding the contract method 0x6aba84a7.
//
// Solidity: function setSynthDeductionPriority(uint128[] newSynthDeductionPriority) returns()
func (_Contracts *ContractsTransactorSession) SetSynthDeductionPriority(newSynthDeductionPriority []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetSynthDeductionPriority(&_Contracts.TransactOpts, newSynthDeductionPriority)
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

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_Contracts *ContractsTransactor) SettlePythOrder(opts *bind.TransactOpts, result []byte, extraData []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "settlePythOrder", result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
func (_Contracts *ContractsSession) SettlePythOrder(result []byte, extraData []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SettlePythOrder(&_Contracts.TransactOpts, result, extraData)
}

// SettlePythOrder is a paid mutator transaction binding the contract method 0x8a0345c6.
//
// Solidity: function settlePythOrder(bytes result, bytes extraData) payable returns()
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

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_Contracts *ContractsTransactor) UpdatePriceData(opts *bind.TransactOpts, perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updatePriceData", perpsMarketId, feedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_Contracts *ContractsSession) UpdatePriceData(perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpdatePriceData(&_Contracts.TransactOpts, perpsMarketId, feedId)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x2d73137b.
//
// Solidity: function updatePriceData(uint128 perpsMarketId, bytes32 feedId) returns()
func (_Contracts *ContractsTransactorSession) UpdatePriceData(perpsMarketId *big.Int, feedId [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpdatePriceData(&_Contracts.TransactOpts, perpsMarketId, feedId)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_Contracts *ContractsTransactor) UpdateReferrerShare(opts *bind.TransactOpts, referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateReferrerShare", referrer, shareRatioD18)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_Contracts *ContractsSession) UpdateReferrerShare(referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateReferrerShare(&_Contracts.TransactOpts, referrer, shareRatioD18)
}

// UpdateReferrerShare is a paid mutator transaction binding the contract method 0x6809fb4d.
//
// Solidity: function updateReferrerShare(address referrer, uint256 shareRatioD18) returns()
func (_Contracts *ContractsTransactorSession) UpdateReferrerShare(referrer common.Address, shareRatioD18 *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateReferrerShare(&_Contracts.TransactOpts, referrer, shareRatioD18)
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

// ContractsAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Contracts contract.
type ContractsAccountCreatedIterator struct {
	Event *ContractsAccountCreated // Event containing the contract specifics and raw log

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
func (it *ContractsAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAccountCreated)
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
		it.Event = new(ContractsAccountCreated)
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
func (it *ContractsAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAccountCreated represents a AccountCreated event raised by the Contracts contract.
type ContractsAccountCreated struct {
	AccountId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_Contracts *ContractsFilterer) FilterAccountCreated(opts *bind.FilterOpts, accountId []*big.Int, owner []common.Address) (*ContractsAccountCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAccountCreatedIterator{contract: _Contracts.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_Contracts *ContractsFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *ContractsAccountCreated, accountId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAccountCreated)
				if err := _Contracts.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseAccountCreated(log types.Log) (*ContractsAccountCreated, error) {
	event := new(ContractsAccountCreated)
	if err := _Contracts.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAccountLiquidatedIterator is returned from FilterAccountLiquidated and is used to iterate over the raw logs and unpacked data for AccountLiquidated events raised by the Contracts contract.
type ContractsAccountLiquidatedIterator struct {
	Event *ContractsAccountLiquidated // Event containing the contract specifics and raw log

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
func (it *ContractsAccountLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAccountLiquidated)
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
		it.Event = new(ContractsAccountLiquidated)
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
func (it *ContractsAccountLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAccountLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAccountLiquidated represents a AccountLiquidated event raised by the Contracts contract.
type ContractsAccountLiquidated struct {
	AccountId       *big.Int
	Reward          *big.Int
	FullLiquidation bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccountLiquidated is a free log retrieval operation binding the contract event 0x926f56b0a80cc150ad77b59cc60bc53cb907eab441d204924ec67b40ae64f5f5.
//
// Solidity: event AccountLiquidated(uint128 indexed accountId, uint256 reward, bool fullLiquidation)
func (_Contracts *ContractsFilterer) FilterAccountLiquidated(opts *bind.FilterOpts, accountId []*big.Int) (*ContractsAccountLiquidatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AccountLiquidated", accountIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAccountLiquidatedIterator{contract: _Contracts.contract, event: "AccountLiquidated", logs: logs, sub: sub}, nil
}

// WatchAccountLiquidated is a free log subscription operation binding the contract event 0x926f56b0a80cc150ad77b59cc60bc53cb907eab441d204924ec67b40ae64f5f5.
//
// Solidity: event AccountLiquidated(uint128 indexed accountId, uint256 reward, bool fullLiquidation)
func (_Contracts *ContractsFilterer) WatchAccountLiquidated(opts *bind.WatchOpts, sink chan<- *ContractsAccountLiquidated, accountId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AccountLiquidated", accountIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAccountLiquidated)
				if err := _Contracts.contract.UnpackLog(event, "AccountLiquidated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseAccountLiquidated(log types.Log) (*ContractsAccountLiquidated, error) {
	event := new(ContractsAccountLiquidated)
	if err := _Contracts.contract.UnpackLog(event, "AccountLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// ContractsCollateralModifiedIterator is returned from FilterCollateralModified and is used to iterate over the raw logs and unpacked data for CollateralModified events raised by the Contracts contract.
type ContractsCollateralModifiedIterator struct {
	Event *ContractsCollateralModified // Event containing the contract specifics and raw log

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
func (it *ContractsCollateralModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCollateralModified)
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
		it.Event = new(ContractsCollateralModified)
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
func (it *ContractsCollateralModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCollateralModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCollateralModified represents a CollateralModified event raised by the Contracts contract.
type ContractsCollateralModified struct {
	AccountId     *big.Int
	SynthMarketId *big.Int
	AmountDelta   *big.Int
	Sender        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralModified is a free log retrieval operation binding the contract event 0x2e8360c2f3a6fc1a15aefdd0a0922bea3c898cb90d38c3a97354e7c013116064.
//
// Solidity: event CollateralModified(uint128 indexed accountId, uint128 indexed synthMarketId, int256 amountDelta, address indexed sender)
func (_Contracts *ContractsFilterer) FilterCollateralModified(opts *bind.FilterOpts, accountId []*big.Int, synthMarketId []*big.Int, sender []common.Address) (*ContractsCollateralModifiedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CollateralModified", accountIdRule, synthMarketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCollateralModifiedIterator{contract: _Contracts.contract, event: "CollateralModified", logs: logs, sub: sub}, nil
}

// WatchCollateralModified is a free log subscription operation binding the contract event 0x2e8360c2f3a6fc1a15aefdd0a0922bea3c898cb90d38c3a97354e7c013116064.
//
// Solidity: event CollateralModified(uint128 indexed accountId, uint128 indexed synthMarketId, int256 amountDelta, address indexed sender)
func (_Contracts *ContractsFilterer) WatchCollateralModified(opts *bind.WatchOpts, sink chan<- *ContractsCollateralModified, accountId []*big.Int, synthMarketId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CollateralModified", accountIdRule, synthMarketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCollateralModified)
				if err := _Contracts.contract.UnpackLog(event, "CollateralModified", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseCollateralModified(log types.Log) (*ContractsCollateralModified, error) {
	event := new(ContractsCollateralModified)
	if err := _Contracts.contract.UnpackLog(event, "CollateralModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFactoryInitializedIterator is returned from FilterFactoryInitialized and is used to iterate over the raw logs and unpacked data for FactoryInitialized events raised by the Contracts contract.
type ContractsFactoryInitializedIterator struct {
	Event *ContractsFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractsFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFactoryInitialized)
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
		it.Event = new(ContractsFactoryInitialized)
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
func (it *ContractsFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFactoryInitialized represents a FactoryInitialized event raised by the Contracts contract.
type ContractsFactoryInitialized struct {
	GlobalPerpsMarketId *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFactoryInitialized is a free log retrieval operation binding the contract event 0xb3240229b07e26f2f02e69dda85ede86e162ccbc6d10e6aade28931e7f533339.
//
// Solidity: event FactoryInitialized(uint128 globalPerpsMarketId)
func (_Contracts *ContractsFilterer) FilterFactoryInitialized(opts *bind.FilterOpts) (*ContractsFactoryInitializedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FactoryInitialized")
	if err != nil {
		return nil, err
	}
	return &ContractsFactoryInitializedIterator{contract: _Contracts.contract, event: "FactoryInitialized", logs: logs, sub: sub}, nil
}

// WatchFactoryInitialized is a free log subscription operation binding the contract event 0xb3240229b07e26f2f02e69dda85ede86e162ccbc6d10e6aade28931e7f533339.
//
// Solidity: event FactoryInitialized(uint128 globalPerpsMarketId)
func (_Contracts *ContractsFilterer) WatchFactoryInitialized(opts *bind.WatchOpts, sink chan<- *ContractsFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FactoryInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFactoryInitialized)
				if err := _Contracts.contract.UnpackLog(event, "FactoryInitialized", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseFactoryInitialized(log types.Log) (*ContractsFactoryInitialized, error) {
	event := new(ContractsFactoryInitialized)
	if err := _Contracts.contract.UnpackLog(event, "FactoryInitialized", log); err != nil {
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
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorSet is a free log retrieval operation binding the contract event 0x12e1d17016b94668449f97876f4a8d5cc2c19f314db337418894734037cc19d4.
//
// Solidity: event FeeCollectorSet(address feeCollector)
func (_Contracts *ContractsFilterer) FilterFeeCollectorSet(opts *bind.FilterOpts) (*ContractsFeeCollectorSetIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FeeCollectorSet")
	if err != nil {
		return nil, err
	}
	return &ContractsFeeCollectorSetIterator{contract: _Contracts.contract, event: "FeeCollectorSet", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorSet is a free log subscription operation binding the contract event 0x12e1d17016b94668449f97876f4a8d5cc2c19f314db337418894734037cc19d4.
//
// Solidity: event FeeCollectorSet(address feeCollector)
func (_Contracts *ContractsFilterer) WatchFeeCollectorSet(opts *bind.WatchOpts, sink chan<- *ContractsFeeCollectorSet) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FeeCollectorSet")
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

// ParseFeeCollectorSet is a log parse operation binding the contract event 0x12e1d17016b94668449f97876f4a8d5cc2c19f314db337418894734037cc19d4.
//
// Solidity: event FeeCollectorSet(address feeCollector)
func (_Contracts *ContractsFilterer) ParseFeeCollectorSet(log types.Log) (*ContractsFeeCollectorSet, error) {
	event := new(ContractsFeeCollectorSet)
	if err := _Contracts.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFundingParametersSetIterator is returned from FilterFundingParametersSet and is used to iterate over the raw logs and unpacked data for FundingParametersSet events raised by the Contracts contract.
type ContractsFundingParametersSetIterator struct {
	Event *ContractsFundingParametersSet // Event containing the contract specifics and raw log

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
func (it *ContractsFundingParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFundingParametersSet)
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
		it.Event = new(ContractsFundingParametersSet)
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
func (it *ContractsFundingParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFundingParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFundingParametersSet represents a FundingParametersSet event raised by the Contracts contract.
type ContractsFundingParametersSet struct {
	MarketId           *big.Int
	SkewScale          *big.Int
	MaxFundingVelocity *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFundingParametersSet is a free log retrieval operation binding the contract event 0xa74afd926bbafbb9252d224a1fcd6a209f851324bd485f556786820a76e31b65.
//
// Solidity: event FundingParametersSet(uint128 indexed marketId, uint256 skewScale, uint256 maxFundingVelocity)
func (_Contracts *ContractsFilterer) FilterFundingParametersSet(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsFundingParametersSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FundingParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsFundingParametersSetIterator{contract: _Contracts.contract, event: "FundingParametersSet", logs: logs, sub: sub}, nil
}

// WatchFundingParametersSet is a free log subscription operation binding the contract event 0xa74afd926bbafbb9252d224a1fcd6a209f851324bd485f556786820a76e31b65.
//
// Solidity: event FundingParametersSet(uint128 indexed marketId, uint256 skewScale, uint256 maxFundingVelocity)
func (_Contracts *ContractsFilterer) WatchFundingParametersSet(opts *bind.WatchOpts, sink chan<- *ContractsFundingParametersSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FundingParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFundingParametersSet)
				if err := _Contracts.contract.UnpackLog(event, "FundingParametersSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseFundingParametersSet(log types.Log) (*ContractsFundingParametersSet, error) {
	event := new(ContractsFundingParametersSet)
	if err := _Contracts.contract.UnpackLog(event, "FundingParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLiquidationParametersSetIterator is returned from FilterLiquidationParametersSet and is used to iterate over the raw logs and unpacked data for LiquidationParametersSet events raised by the Contracts contract.
type ContractsLiquidationParametersSetIterator struct {
	Event *ContractsLiquidationParametersSet // Event containing the contract specifics and raw log

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
func (it *ContractsLiquidationParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLiquidationParametersSet)
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
		it.Event = new(ContractsLiquidationParametersSet)
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
func (it *ContractsLiquidationParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLiquidationParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLiquidationParametersSet represents a LiquidationParametersSet event raised by the Contracts contract.
type ContractsLiquidationParametersSet struct {
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
func (_Contracts *ContractsFilterer) FilterLiquidationParametersSet(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsLiquidationParametersSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLiquidationParametersSetIterator{contract: _Contracts.contract, event: "LiquidationParametersSet", logs: logs, sub: sub}, nil
}

// WatchLiquidationParametersSet is a free log subscription operation binding the contract event 0x362c38617bd453ec04ea677dc8dba9aa5b977195f14ddcf5cef0509bb3fd9200.
//
// Solidity: event LiquidationParametersSet(uint128 indexed marketId, uint256 initialMarginRatioD18, uint256 maintenanceMarginRatioD18, uint256 minimumInitialMarginRatioD18, uint256 liquidationRewardRatioD18, uint256 maxLiquidationLimitAccumulationMultiplier, uint256 maxSecondsInLiquidationWindow, uint256 minimumPositionMargin)
func (_Contracts *ContractsFilterer) WatchLiquidationParametersSet(opts *bind.WatchOpts, sink chan<- *ContractsLiquidationParametersSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LiquidationParametersSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLiquidationParametersSet)
				if err := _Contracts.contract.UnpackLog(event, "LiquidationParametersSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseLiquidationParametersSet(log types.Log) (*ContractsLiquidationParametersSet, error) {
	event := new(ContractsLiquidationParametersSet)
	if err := _Contracts.contract.UnpackLog(event, "LiquidationParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLiquidationRewardGuardsSetIterator is returned from FilterLiquidationRewardGuardsSet and is used to iterate over the raw logs and unpacked data for LiquidationRewardGuardsSet events raised by the Contracts contract.
type ContractsLiquidationRewardGuardsSetIterator struct {
	Event *ContractsLiquidationRewardGuardsSet // Event containing the contract specifics and raw log

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
func (it *ContractsLiquidationRewardGuardsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLiquidationRewardGuardsSet)
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
		it.Event = new(ContractsLiquidationRewardGuardsSet)
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
func (it *ContractsLiquidationRewardGuardsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLiquidationRewardGuardsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLiquidationRewardGuardsSet represents a LiquidationRewardGuardsSet event raised by the Contracts contract.
type ContractsLiquidationRewardGuardsSet struct {
	MinLiquidationRewardUsd *big.Int
	MaxLiquidationRewardUsd *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterLiquidationRewardGuardsSet is a free log retrieval operation binding the contract event 0xbb7f8eb17b00549c0f234a3c363ef3eb39329f98b52048416480cf62d0f0e1f5.
//
// Solidity: event LiquidationRewardGuardsSet(uint256 indexed minLiquidationRewardUsd, uint256 indexed maxLiquidationRewardUsd)
func (_Contracts *ContractsFilterer) FilterLiquidationRewardGuardsSet(opts *bind.FilterOpts, minLiquidationRewardUsd []*big.Int, maxLiquidationRewardUsd []*big.Int) (*ContractsLiquidationRewardGuardsSetIterator, error) {

	var minLiquidationRewardUsdRule []interface{}
	for _, minLiquidationRewardUsdItem := range minLiquidationRewardUsd {
		minLiquidationRewardUsdRule = append(minLiquidationRewardUsdRule, minLiquidationRewardUsdItem)
	}
	var maxLiquidationRewardUsdRule []interface{}
	for _, maxLiquidationRewardUsdItem := range maxLiquidationRewardUsd {
		maxLiquidationRewardUsdRule = append(maxLiquidationRewardUsdRule, maxLiquidationRewardUsdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LiquidationRewardGuardsSet", minLiquidationRewardUsdRule, maxLiquidationRewardUsdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLiquidationRewardGuardsSetIterator{contract: _Contracts.contract, event: "LiquidationRewardGuardsSet", logs: logs, sub: sub}, nil
}

// WatchLiquidationRewardGuardsSet is a free log subscription operation binding the contract event 0xbb7f8eb17b00549c0f234a3c363ef3eb39329f98b52048416480cf62d0f0e1f5.
//
// Solidity: event LiquidationRewardGuardsSet(uint256 indexed minLiquidationRewardUsd, uint256 indexed maxLiquidationRewardUsd)
func (_Contracts *ContractsFilterer) WatchLiquidationRewardGuardsSet(opts *bind.WatchOpts, sink chan<- *ContractsLiquidationRewardGuardsSet, minLiquidationRewardUsd []*big.Int, maxLiquidationRewardUsd []*big.Int) (event.Subscription, error) {

	var minLiquidationRewardUsdRule []interface{}
	for _, minLiquidationRewardUsdItem := range minLiquidationRewardUsd {
		minLiquidationRewardUsdRule = append(minLiquidationRewardUsdRule, minLiquidationRewardUsdItem)
	}
	var maxLiquidationRewardUsdRule []interface{}
	for _, maxLiquidationRewardUsdItem := range maxLiquidationRewardUsd {
		maxLiquidationRewardUsdRule = append(maxLiquidationRewardUsdRule, maxLiquidationRewardUsdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LiquidationRewardGuardsSet", minLiquidationRewardUsdRule, maxLiquidationRewardUsdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLiquidationRewardGuardsSet)
				if err := _Contracts.contract.UnpackLog(event, "LiquidationRewardGuardsSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseLiquidationRewardGuardsSet(log types.Log) (*ContractsLiquidationRewardGuardsSet, error) {
	event := new(ContractsLiquidationRewardGuardsSet)
	if err := _Contracts.contract.UnpackLog(event, "LiquidationRewardGuardsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLockedOiRatioSetIterator is returned from FilterLockedOiRatioSet and is used to iterate over the raw logs and unpacked data for LockedOiRatioSet events raised by the Contracts contract.
type ContractsLockedOiRatioSetIterator struct {
	Event *ContractsLockedOiRatioSet // Event containing the contract specifics and raw log

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
func (it *ContractsLockedOiRatioSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLockedOiRatioSet)
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
		it.Event = new(ContractsLockedOiRatioSet)
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
func (it *ContractsLockedOiRatioSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLockedOiRatioSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLockedOiRatioSet represents a LockedOiRatioSet event raised by the Contracts contract.
type ContractsLockedOiRatioSet struct {
	MarketId         *big.Int
	LockedOiRatioD18 *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLockedOiRatioSet is a free log retrieval operation binding the contract event 0x1d841fd5b4c806bc5a073d637a8506e1e74d16cb18251b711cb47e133ceafc2d.
//
// Solidity: event LockedOiRatioSet(uint128 indexed marketId, uint256 lockedOiRatioD18)
func (_Contracts *ContractsFilterer) FilterLockedOiRatioSet(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsLockedOiRatioSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LockedOiRatioSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLockedOiRatioSetIterator{contract: _Contracts.contract, event: "LockedOiRatioSet", logs: logs, sub: sub}, nil
}

// WatchLockedOiRatioSet is a free log subscription operation binding the contract event 0x1d841fd5b4c806bc5a073d637a8506e1e74d16cb18251b711cb47e133ceafc2d.
//
// Solidity: event LockedOiRatioSet(uint128 indexed marketId, uint256 lockedOiRatioD18)
func (_Contracts *ContractsFilterer) WatchLockedOiRatioSet(opts *bind.WatchOpts, sink chan<- *ContractsLockedOiRatioSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LockedOiRatioSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLockedOiRatioSet)
				if err := _Contracts.contract.UnpackLog(event, "LockedOiRatioSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseLockedOiRatioSet(log types.Log) (*ContractsLockedOiRatioSet, error) {
	event := new(ContractsLockedOiRatioSet)
	if err := _Contracts.contract.UnpackLog(event, "LockedOiRatioSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the Contracts contract.
type ContractsMarketCreatedIterator struct {
	Event *ContractsMarketCreated // Event containing the contract specifics and raw log

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
func (it *ContractsMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketCreated)
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
		it.Event = new(ContractsMarketCreated)
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
func (it *ContractsMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketCreated represents a MarketCreated event raised by the Contracts contract.
type ContractsMarketCreated struct {
	PerpsMarketId *big.Int
	MarketName    string
	MarketSymbol  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0x032553f94ac1323933f22650ec5b8e232babf1c47efca69383b749463116cc49.
//
// Solidity: event MarketCreated(uint128 indexed perpsMarketId, string marketName, string marketSymbol)
func (_Contracts *ContractsFilterer) FilterMarketCreated(opts *bind.FilterOpts, perpsMarketId []*big.Int) (*ContractsMarketCreatedIterator, error) {

	var perpsMarketIdRule []interface{}
	for _, perpsMarketIdItem := range perpsMarketId {
		perpsMarketIdRule = append(perpsMarketIdRule, perpsMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketCreated", perpsMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketCreatedIterator{contract: _Contracts.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0x032553f94ac1323933f22650ec5b8e232babf1c47efca69383b749463116cc49.
//
// Solidity: event MarketCreated(uint128 indexed perpsMarketId, string marketName, string marketSymbol)
func (_Contracts *ContractsFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *ContractsMarketCreated, perpsMarketId []*big.Int) (event.Subscription, error) {

	var perpsMarketIdRule []interface{}
	for _, perpsMarketIdItem := range perpsMarketId {
		perpsMarketIdRule = append(perpsMarketIdRule, perpsMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketCreated", perpsMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketCreated)
				if err := _Contracts.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseMarketCreated(log types.Log) (*ContractsMarketCreated, error) {
	event := new(ContractsMarketCreated)
	if err := _Contracts.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketPriceDataUpdatedIterator is returned from FilterMarketPriceDataUpdated and is used to iterate over the raw logs and unpacked data for MarketPriceDataUpdated events raised by the Contracts contract.
type ContractsMarketPriceDataUpdatedIterator struct {
	Event *ContractsMarketPriceDataUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsMarketPriceDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketPriceDataUpdated)
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
		it.Event = new(ContractsMarketPriceDataUpdated)
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
func (it *ContractsMarketPriceDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketPriceDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketPriceDataUpdated represents a MarketPriceDataUpdated event raised by the Contracts contract.
type ContractsMarketPriceDataUpdated struct {
	MarketId *big.Int
	FeedId   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketPriceDataUpdated is a free log retrieval operation binding the contract event 0xd5ec22bbdbde803d8585a5aa89d991c67dc62c6a7bc3c4fa000b4e263b6783a0.
//
// Solidity: event MarketPriceDataUpdated(uint128 indexed marketId, bytes32 feedId)
func (_Contracts *ContractsFilterer) FilterMarketPriceDataUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsMarketPriceDataUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketPriceDataUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketPriceDataUpdatedIterator{contract: _Contracts.contract, event: "MarketPriceDataUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketPriceDataUpdated is a free log subscription operation binding the contract event 0xd5ec22bbdbde803d8585a5aa89d991c67dc62c6a7bc3c4fa000b4e263b6783a0.
//
// Solidity: event MarketPriceDataUpdated(uint128 indexed marketId, bytes32 feedId)
func (_Contracts *ContractsFilterer) WatchMarketPriceDataUpdated(opts *bind.WatchOpts, sink chan<- *ContractsMarketPriceDataUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketPriceDataUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketPriceDataUpdated)
				if err := _Contracts.contract.UnpackLog(event, "MarketPriceDataUpdated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseMarketPriceDataUpdated(log types.Log) (*ContractsMarketPriceDataUpdated, error) {
	event := new(ContractsMarketPriceDataUpdated)
	if err := _Contracts.contract.UnpackLog(event, "MarketPriceDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketUpdatedIterator is returned from FilterMarketUpdated and is used to iterate over the raw logs and unpacked data for MarketUpdated events raised by the Contracts contract.
type ContractsMarketUpdatedIterator struct {
	Event *ContractsMarketUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsMarketUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketUpdated)
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
		it.Event = new(ContractsMarketUpdated)
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
func (it *ContractsMarketUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketUpdated represents a MarketUpdated event raised by the Contracts contract.
type ContractsMarketUpdated struct {
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
func (_Contracts *ContractsFilterer) FilterMarketUpdated(opts *bind.FilterOpts) (*ContractsMarketUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsMarketUpdatedIterator{contract: _Contracts.contract, event: "MarketUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketUpdated is a free log subscription operation binding the contract event 0xb317f068f30db823aeb6ac6ddbffcbb6c805f558972e6b16625ec58b83f1f3d5.
//
// Solidity: event MarketUpdated(uint128 marketId, uint256 price, int256 skew, uint256 size, int256 sizeDelta, int256 currentFundingRate, int256 currentFundingVelocity)
func (_Contracts *ContractsFilterer) WatchMarketUpdated(opts *bind.WatchOpts, sink chan<- *ContractsMarketUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketUpdated)
				if err := _Contracts.contract.UnpackLog(event, "MarketUpdated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseMarketUpdated(log types.Log) (*ContractsMarketUpdated, error) {
	event := new(ContractsMarketUpdated)
	if err := _Contracts.contract.UnpackLog(event, "MarketUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMaxCollateralAmountSetIterator is returned from FilterMaxCollateralAmountSet and is used to iterate over the raw logs and unpacked data for MaxCollateralAmountSet events raised by the Contracts contract.
type ContractsMaxCollateralAmountSetIterator struct {
	Event *ContractsMaxCollateralAmountSet // Event containing the contract specifics and raw log

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
func (it *ContractsMaxCollateralAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMaxCollateralAmountSet)
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
		it.Event = new(ContractsMaxCollateralAmountSet)
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
func (it *ContractsMaxCollateralAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMaxCollateralAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMaxCollateralAmountSet represents a MaxCollateralAmountSet event raised by the Contracts contract.
type ContractsMaxCollateralAmountSet struct {
	SynthMarketId    *big.Int
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMaxCollateralAmountSet is a free log retrieval operation binding the contract event 0xa552ba0f2552ec13b62b479a41e6e7dfcef9f10aea18527e5a3e1b3963bf70d7.
//
// Solidity: event MaxCollateralAmountSet(uint128 indexed synthMarketId, uint256 collateralAmount)
func (_Contracts *ContractsFilterer) FilterMaxCollateralAmountSet(opts *bind.FilterOpts, synthMarketId []*big.Int) (*ContractsMaxCollateralAmountSetIterator, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MaxCollateralAmountSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMaxCollateralAmountSetIterator{contract: _Contracts.contract, event: "MaxCollateralAmountSet", logs: logs, sub: sub}, nil
}

// WatchMaxCollateralAmountSet is a free log subscription operation binding the contract event 0xa552ba0f2552ec13b62b479a41e6e7dfcef9f10aea18527e5a3e1b3963bf70d7.
//
// Solidity: event MaxCollateralAmountSet(uint128 indexed synthMarketId, uint256 collateralAmount)
func (_Contracts *ContractsFilterer) WatchMaxCollateralAmountSet(opts *bind.WatchOpts, sink chan<- *ContractsMaxCollateralAmountSet, synthMarketId []*big.Int) (event.Subscription, error) {

	var synthMarketIdRule []interface{}
	for _, synthMarketIdItem := range synthMarketId {
		synthMarketIdRule = append(synthMarketIdRule, synthMarketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MaxCollateralAmountSet", synthMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMaxCollateralAmountSet)
				if err := _Contracts.contract.UnpackLog(event, "MaxCollateralAmountSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseMaxCollateralAmountSet(log types.Log) (*ContractsMaxCollateralAmountSet, error) {
	event := new(ContractsMaxCollateralAmountSet)
	if err := _Contracts.contract.UnpackLog(event, "MaxCollateralAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMaxMarketSizeSetIterator is returned from FilterMaxMarketSizeSet and is used to iterate over the raw logs and unpacked data for MaxMarketSizeSet events raised by the Contracts contract.
type ContractsMaxMarketSizeSetIterator struct {
	Event *ContractsMaxMarketSizeSet // Event containing the contract specifics and raw log

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
func (it *ContractsMaxMarketSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMaxMarketSizeSet)
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
		it.Event = new(ContractsMaxMarketSizeSet)
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
func (it *ContractsMaxMarketSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMaxMarketSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMaxMarketSizeSet represents a MaxMarketSizeSet event raised by the Contracts contract.
type ContractsMaxMarketSizeSet struct {
	MarketId      *big.Int
	MaxMarketSize *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxMarketSizeSet is a free log retrieval operation binding the contract event 0xbd063bd3072a194b255163ab8dfd3400c4ab1cc641b920e7aaf1c11f92cd26cf.
//
// Solidity: event MaxMarketSizeSet(uint128 indexed marketId, uint256 maxMarketSize)
func (_Contracts *ContractsFilterer) FilterMaxMarketSizeSet(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsMaxMarketSizeSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MaxMarketSizeSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMaxMarketSizeSetIterator{contract: _Contracts.contract, event: "MaxMarketSizeSet", logs: logs, sub: sub}, nil
}

// WatchMaxMarketSizeSet is a free log subscription operation binding the contract event 0xbd063bd3072a194b255163ab8dfd3400c4ab1cc641b920e7aaf1c11f92cd26cf.
//
// Solidity: event MaxMarketSizeSet(uint128 indexed marketId, uint256 maxMarketSize)
func (_Contracts *ContractsFilterer) WatchMaxMarketSizeSet(opts *bind.WatchOpts, sink chan<- *ContractsMaxMarketSizeSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MaxMarketSizeSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMaxMarketSizeSet)
				if err := _Contracts.contract.UnpackLog(event, "MaxMarketSizeSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseMaxMarketSizeSet(log types.Log) (*ContractsMaxMarketSizeSet, error) {
	event := new(ContractsMaxMarketSizeSet)
	if err := _Contracts.contract.UnpackLog(event, "MaxMarketSizeSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) FilterOrderCommitted(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*ContractsOrderCommittedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OrderCommitted", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOrderCommittedIterator{contract: _Contracts.contract, event: "OrderCommitted", logs: logs, sub: sub}, nil
}

// WatchOrderCommitted is a free log subscription operation binding the contract event 0xc74e8cf044fc959dcef2695245eb90f7e9fb5676532a71eec5b3c513d4793709.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint128 indexed accountId, uint8 orderType, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, uint256 expirationTime, bytes32 indexed trackingCode, address sender)
func (_Contracts *ContractsFilterer) WatchOrderCommitted(opts *bind.WatchOpts, sink chan<- *ContractsOrderCommitted, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OrderCommitted", marketIdRule, accountIdRule, trackingCodeRule)
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

// ParseOrderCommitted is a log parse operation binding the contract event 0xc74e8cf044fc959dcef2695245eb90f7e9fb5676532a71eec5b3c513d4793709.
//
// Solidity: event OrderCommitted(uint128 indexed marketId, uint128 indexed accountId, uint8 orderType, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, uint256 expirationTime, bytes32 indexed trackingCode, address sender)
func (_Contracts *ContractsFilterer) ParseOrderCommitted(log types.Log) (*ContractsOrderCommitted, error) {
	event := new(ContractsOrderCommitted)
	if err := _Contracts.contract.UnpackLog(event, "OrderCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOrderFeesSetIterator is returned from FilterOrderFeesSet and is used to iterate over the raw logs and unpacked data for OrderFeesSet events raised by the Contracts contract.
type ContractsOrderFeesSetIterator struct {
	Event *ContractsOrderFeesSet // Event containing the contract specifics and raw log

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
func (it *ContractsOrderFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOrderFeesSet)
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
		it.Event = new(ContractsOrderFeesSet)
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
func (it *ContractsOrderFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOrderFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOrderFeesSet represents a OrderFeesSet event raised by the Contracts contract.
type ContractsOrderFeesSet struct {
	MarketId      *big.Int
	MakerFeeRatio *big.Int
	TakerFeeRatio *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFeesSet is a free log retrieval operation binding the contract event 0x28969f156340ba9eb31589904eb174d3a4b6a37781fa6f7ad289d349d75dd1ee.
//
// Solidity: event OrderFeesSet(uint128 indexed marketId, uint256 makerFeeRatio, uint256 takerFeeRatio)
func (_Contracts *ContractsFilterer) FilterOrderFeesSet(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsOrderFeesSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OrderFeesSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOrderFeesSetIterator{contract: _Contracts.contract, event: "OrderFeesSet", logs: logs, sub: sub}, nil
}

// WatchOrderFeesSet is a free log subscription operation binding the contract event 0x28969f156340ba9eb31589904eb174d3a4b6a37781fa6f7ad289d349d75dd1ee.
//
// Solidity: event OrderFeesSet(uint128 indexed marketId, uint256 makerFeeRatio, uint256 takerFeeRatio)
func (_Contracts *ContractsFilterer) WatchOrderFeesSet(opts *bind.WatchOpts, sink chan<- *ContractsOrderFeesSet, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OrderFeesSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOrderFeesSet)
				if err := _Contracts.contract.UnpackLog(event, "OrderFeesSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOrderFeesSet(log types.Log) (*ContractsOrderFeesSet, error) {
	event := new(ContractsOrderFeesSet)
	if err := _Contracts.contract.UnpackLog(event, "OrderFeesSet", log); err != nil {
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
func (_Contracts *ContractsFilterer) FilterOrderSettled(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*ContractsOrderSettledIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OrderSettled", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOrderSettledIterator{contract: _Contracts.contract, event: "OrderSettled", logs: logs, sub: sub}, nil
}

// WatchOrderSettled is a free log subscription operation binding the contract event 0x460080a757ec90719fe90ab2384c0196cdeed071a9fd7ce1ada43481d96b7db5.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed accountId, uint256 fillPrice, int256 pnl, int256 accruedFunding, int128 sizeDelta, int128 newSize, uint256 totalFees, uint256 referralFees, uint256 collectedFees, uint256 settlementReward, bytes32 indexed trackingCode, address settler)
func (_Contracts *ContractsFilterer) WatchOrderSettled(opts *bind.WatchOpts, sink chan<- *ContractsOrderSettled, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OrderSettled", marketIdRule, accountIdRule, trackingCodeRule)
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

// ParseOrderSettled is a log parse operation binding the contract event 0x460080a757ec90719fe90ab2384c0196cdeed071a9fd7ce1ada43481d96b7db5.
//
// Solidity: event OrderSettled(uint128 indexed marketId, uint128 indexed accountId, uint256 fillPrice, int256 pnl, int256 accruedFunding, int128 sizeDelta, int128 newSize, uint256 totalFees, uint256 referralFees, uint256 collectedFees, uint256 settlementReward, bytes32 indexed trackingCode, address settler)
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

// ContractsPermissionGrantedIterator is returned from FilterPermissionGranted and is used to iterate over the raw logs and unpacked data for PermissionGranted events raised by the Contracts contract.
type ContractsPermissionGrantedIterator struct {
	Event *ContractsPermissionGranted // Event containing the contract specifics and raw log

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
func (it *ContractsPermissionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPermissionGranted)
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
		it.Event = new(ContractsPermissionGranted)
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
func (it *ContractsPermissionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPermissionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPermissionGranted represents a PermissionGranted event raised by the Contracts contract.
type ContractsPermissionGranted struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionGranted is a free log retrieval operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_Contracts *ContractsFilterer) FilterPermissionGranted(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*ContractsPermissionGrantedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPermissionGrantedIterator{contract: _Contracts.contract, event: "PermissionGranted", logs: logs, sub: sub}, nil
}

// WatchPermissionGranted is a free log subscription operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_Contracts *ContractsFilterer) WatchPermissionGranted(opts *bind.WatchOpts, sink chan<- *ContractsPermissionGranted, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPermissionGranted)
				if err := _Contracts.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParsePermissionGranted(log types.Log) (*ContractsPermissionGranted, error) {
	event := new(ContractsPermissionGranted)
	if err := _Contracts.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPermissionRevokedIterator is returned from FilterPermissionRevoked and is used to iterate over the raw logs and unpacked data for PermissionRevoked events raised by the Contracts contract.
type ContractsPermissionRevokedIterator struct {
	Event *ContractsPermissionRevoked // Event containing the contract specifics and raw log

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
func (it *ContractsPermissionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPermissionRevoked)
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
		it.Event = new(ContractsPermissionRevoked)
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
func (it *ContractsPermissionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPermissionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPermissionRevoked represents a PermissionRevoked event raised by the Contracts contract.
type ContractsPermissionRevoked struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionRevoked is a free log retrieval operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_Contracts *ContractsFilterer) FilterPermissionRevoked(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*ContractsPermissionRevokedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPermissionRevokedIterator{contract: _Contracts.contract, event: "PermissionRevoked", logs: logs, sub: sub}, nil
}

// WatchPermissionRevoked is a free log subscription operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_Contracts *ContractsFilterer) WatchPermissionRevoked(opts *bind.WatchOpts, sink chan<- *ContractsPermissionRevoked, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPermissionRevoked)
				if err := _Contracts.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParsePermissionRevoked(log types.Log) (*ContractsPermissionRevoked, error) {
	event := new(ContractsPermissionRevoked)
	if err := _Contracts.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPositionLiquidatedIterator is returned from FilterPositionLiquidated and is used to iterate over the raw logs and unpacked data for PositionLiquidated events raised by the Contracts contract.
type ContractsPositionLiquidatedIterator struct {
	Event *ContractsPositionLiquidated // Event containing the contract specifics and raw log

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
func (it *ContractsPositionLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPositionLiquidated)
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
		it.Event = new(ContractsPositionLiquidated)
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
func (it *ContractsPositionLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPositionLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPositionLiquidated represents a PositionLiquidated event raised by the Contracts contract.
type ContractsPositionLiquidated struct {
	AccountId           *big.Int
	MarketId            *big.Int
	AmountLiquidated    *big.Int
	CurrentPositionSize *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPositionLiquidated is a free log retrieval operation binding the contract event 0xd583c0e2965aae317f4a9a6603c0c75602b9bc97dc7c5fc6446b0ba8d3ff1bb2.
//
// Solidity: event PositionLiquidated(uint128 indexed accountId, uint128 indexed marketId, uint256 amountLiquidated, int128 currentPositionSize)
func (_Contracts *ContractsFilterer) FilterPositionLiquidated(opts *bind.FilterOpts, accountId []*big.Int, marketId []*big.Int) (*ContractsPositionLiquidatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PositionLiquidated", accountIdRule, marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPositionLiquidatedIterator{contract: _Contracts.contract, event: "PositionLiquidated", logs: logs, sub: sub}, nil
}

// WatchPositionLiquidated is a free log subscription operation binding the contract event 0xd583c0e2965aae317f4a9a6603c0c75602b9bc97dc7c5fc6446b0ba8d3ff1bb2.
//
// Solidity: event PositionLiquidated(uint128 indexed accountId, uint128 indexed marketId, uint256 amountLiquidated, int128 currentPositionSize)
func (_Contracts *ContractsFilterer) WatchPositionLiquidated(opts *bind.WatchOpts, sink chan<- *ContractsPositionLiquidated, accountId []*big.Int, marketId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PositionLiquidated", accountIdRule, marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPositionLiquidated)
				if err := _Contracts.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParsePositionLiquidated(log types.Log) (*ContractsPositionLiquidated, error) {
	event := new(ContractsPositionLiquidated)
	if err := _Contracts.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPreviousOrderExpiredIterator is returned from FilterPreviousOrderExpired and is used to iterate over the raw logs and unpacked data for PreviousOrderExpired events raised by the Contracts contract.
type ContractsPreviousOrderExpiredIterator struct {
	Event *ContractsPreviousOrderExpired // Event containing the contract specifics and raw log

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
func (it *ContractsPreviousOrderExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPreviousOrderExpired)
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
		it.Event = new(ContractsPreviousOrderExpired)
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
func (it *ContractsPreviousOrderExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPreviousOrderExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPreviousOrderExpired represents a PreviousOrderExpired event raised by the Contracts contract.
type ContractsPreviousOrderExpired struct {
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
func (_Contracts *ContractsFilterer) FilterPreviousOrderExpired(opts *bind.FilterOpts, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (*ContractsPreviousOrderExpiredIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PreviousOrderExpired", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPreviousOrderExpiredIterator{contract: _Contracts.contract, event: "PreviousOrderExpired", logs: logs, sub: sub}, nil
}

// WatchPreviousOrderExpired is a free log subscription operation binding the contract event 0x6d83c6751813f50325d75bc054621f83299659c5814d1e5fe6ac117860710dde.
//
// Solidity: event PreviousOrderExpired(uint128 indexed marketId, uint128 indexed accountId, int128 sizeDelta, uint256 acceptablePrice, uint256 settlementTime, bytes32 indexed trackingCode)
func (_Contracts *ContractsFilterer) WatchPreviousOrderExpired(opts *bind.WatchOpts, sink chan<- *ContractsPreviousOrderExpired, marketId []*big.Int, accountId []*big.Int, trackingCode [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PreviousOrderExpired", marketIdRule, accountIdRule, trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPreviousOrderExpired)
				if err := _Contracts.contract.UnpackLog(event, "PreviousOrderExpired", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParsePreviousOrderExpired(log types.Log) (*ContractsPreviousOrderExpired, error) {
	event := new(ContractsPreviousOrderExpired)
	if err := _Contracts.contract.UnpackLog(event, "PreviousOrderExpired", log); err != nil {
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
	Referrer      common.Address
	ShareRatioD18 *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReferrerShareUpdated is a free log retrieval operation binding the contract event 0xa225c555f4cd21a533ad4e01eaf30153c84ca28265d954a651410d3c1e56242c.
//
// Solidity: event ReferrerShareUpdated(address referrer, uint256 shareRatioD18)
func (_Contracts *ContractsFilterer) FilterReferrerShareUpdated(opts *bind.FilterOpts) (*ContractsReferrerShareUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ReferrerShareUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsReferrerShareUpdatedIterator{contract: _Contracts.contract, event: "ReferrerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchReferrerShareUpdated is a free log subscription operation binding the contract event 0xa225c555f4cd21a533ad4e01eaf30153c84ca28265d954a651410d3c1e56242c.
//
// Solidity: event ReferrerShareUpdated(address referrer, uint256 shareRatioD18)
func (_Contracts *ContractsFilterer) WatchReferrerShareUpdated(opts *bind.WatchOpts, sink chan<- *ContractsReferrerShareUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ReferrerShareUpdated")
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

// ParseReferrerShareUpdated is a log parse operation binding the contract event 0xa225c555f4cd21a533ad4e01eaf30153c84ca28265d954a651410d3c1e56242c.
//
// Solidity: event ReferrerShareUpdated(address referrer, uint256 shareRatioD18)
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
	MarketId   *big.Int
	Strategy   SettlementStrategyData
	StrategyId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyAdded is a free log retrieval operation binding the contract event 0x6ea79b5183a0fa7ead0ea991cd5e5fa23e2da92b703bbe111479fffcdea58827.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy, uint256 indexed strategyId)
func (_Contracts *ContractsFilterer) FilterSettlementStrategyAdded(opts *bind.FilterOpts, marketId []*big.Int, strategyId []*big.Int) (*ContractsSettlementStrategyAddedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SettlementStrategyAdded", marketIdRule, strategyIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSettlementStrategyAddedIterator{contract: _Contracts.contract, event: "SettlementStrategyAdded", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyAdded is a free log subscription operation binding the contract event 0x6ea79b5183a0fa7ead0ea991cd5e5fa23e2da92b703bbe111479fffcdea58827.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy, uint256 indexed strategyId)
func (_Contracts *ContractsFilterer) WatchSettlementStrategyAdded(opts *bind.WatchOpts, sink chan<- *ContractsSettlementStrategyAdded, marketId []*big.Int, strategyId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	var strategyIdRule []interface{}
	for _, strategyIdItem := range strategyId {
		strategyIdRule = append(strategyIdRule, strategyIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SettlementStrategyAdded", marketIdRule, strategyIdRule)
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

// ParseSettlementStrategyAdded is a log parse operation binding the contract event 0x6ea79b5183a0fa7ead0ea991cd5e5fa23e2da92b703bbe111479fffcdea58827.
//
// Solidity: event SettlementStrategyAdded(uint128 indexed marketId, (uint8,uint256,uint256,uint256,address,bytes32,string,uint256,uint256,bool) strategy, uint256 indexed strategyId)
func (_Contracts *ContractsFilterer) ParseSettlementStrategyAdded(log types.Log) (*ContractsSettlementStrategyAdded, error) {
	event := new(ContractsSettlementStrategyAdded)
	if err := _Contracts.contract.UnpackLog(event, "SettlementStrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSettlementStrategyEnabledIterator is returned from FilterSettlementStrategyEnabled and is used to iterate over the raw logs and unpacked data for SettlementStrategyEnabled events raised by the Contracts contract.
type ContractsSettlementStrategyEnabledIterator struct {
	Event *ContractsSettlementStrategyEnabled // Event containing the contract specifics and raw log

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
func (it *ContractsSettlementStrategyEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSettlementStrategyEnabled)
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
		it.Event = new(ContractsSettlementStrategyEnabled)
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
func (it *ContractsSettlementStrategyEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSettlementStrategyEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSettlementStrategyEnabled represents a SettlementStrategyEnabled event raised by the Contracts contract.
type ContractsSettlementStrategyEnabled struct {
	MarketId   *big.Int
	StrategyId *big.Int
	Enabled    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettlementStrategyEnabled is a free log retrieval operation binding the contract event 0x4803eb4e51b45b2471b0c87d8c17e0c2d8961ae080cc4a930883eab155e0fa22.
//
// Solidity: event SettlementStrategyEnabled(uint128 indexed marketId, uint256 strategyId, bool enabled)
func (_Contracts *ContractsFilterer) FilterSettlementStrategyEnabled(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsSettlementStrategyEnabledIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SettlementStrategyEnabled", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSettlementStrategyEnabledIterator{contract: _Contracts.contract, event: "SettlementStrategyEnabled", logs: logs, sub: sub}, nil
}

// WatchSettlementStrategyEnabled is a free log subscription operation binding the contract event 0x4803eb4e51b45b2471b0c87d8c17e0c2d8961ae080cc4a930883eab155e0fa22.
//
// Solidity: event SettlementStrategyEnabled(uint128 indexed marketId, uint256 strategyId, bool enabled)
func (_Contracts *ContractsFilterer) WatchSettlementStrategyEnabled(opts *bind.WatchOpts, sink chan<- *ContractsSettlementStrategyEnabled, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SettlementStrategyEnabled", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSettlementStrategyEnabled)
				if err := _Contracts.contract.UnpackLog(event, "SettlementStrategyEnabled", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSettlementStrategyEnabled(log types.Log) (*ContractsSettlementStrategyEnabled, error) {
	event := new(ContractsSettlementStrategyEnabled)
	if err := _Contracts.contract.UnpackLog(event, "SettlementStrategyEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSynthDeductionPrioritySetIterator is returned from FilterSynthDeductionPrioritySet and is used to iterate over the raw logs and unpacked data for SynthDeductionPrioritySet events raised by the Contracts contract.
type ContractsSynthDeductionPrioritySetIterator struct {
	Event *ContractsSynthDeductionPrioritySet // Event containing the contract specifics and raw log

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
func (it *ContractsSynthDeductionPrioritySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSynthDeductionPrioritySet)
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
		it.Event = new(ContractsSynthDeductionPrioritySet)
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
func (it *ContractsSynthDeductionPrioritySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSynthDeductionPrioritySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSynthDeductionPrioritySet represents a SynthDeductionPrioritySet event raised by the Contracts contract.
type ContractsSynthDeductionPrioritySet struct {
	NewSynthDeductionPriority []*big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSynthDeductionPrioritySet is a free log retrieval operation binding the contract event 0xa6beea856d32db9e9614e4af02fc7d6a3fa8337a13e4a48486c142a9a9a8ed8f.
//
// Solidity: event SynthDeductionPrioritySet(uint128[] newSynthDeductionPriority)
func (_Contracts *ContractsFilterer) FilterSynthDeductionPrioritySet(opts *bind.FilterOpts) (*ContractsSynthDeductionPrioritySetIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SynthDeductionPrioritySet")
	if err != nil {
		return nil, err
	}
	return &ContractsSynthDeductionPrioritySetIterator{contract: _Contracts.contract, event: "SynthDeductionPrioritySet", logs: logs, sub: sub}, nil
}

// WatchSynthDeductionPrioritySet is a free log subscription operation binding the contract event 0xa6beea856d32db9e9614e4af02fc7d6a3fa8337a13e4a48486c142a9a9a8ed8f.
//
// Solidity: event SynthDeductionPrioritySet(uint128[] newSynthDeductionPriority)
func (_Contracts *ContractsFilterer) WatchSynthDeductionPrioritySet(opts *bind.WatchOpts, sink chan<- *ContractsSynthDeductionPrioritySet) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SynthDeductionPrioritySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSynthDeductionPrioritySet)
				if err := _Contracts.contract.UnpackLog(event, "SynthDeductionPrioritySet", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseSynthDeductionPrioritySet(log types.Log) (*ContractsSynthDeductionPrioritySet, error) {
	event := new(ContractsSynthDeductionPrioritySet)
	if err := _Contracts.contract.UnpackLog(event, "SynthDeductionPrioritySet", log); err != nil {
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
