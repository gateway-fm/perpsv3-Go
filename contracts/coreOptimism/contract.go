// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coreOptimism

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

// CcipClientAny2EVMMessage is an auto generated low-level Go binding around an user-defined struct.
type CcipClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	TokenAmounts        []CcipClientEVMTokenAmount
}

// CcipClientEVMTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type CcipClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// CollateralConfigurationData is an auto generated low-level Go binding around an user-defined struct.
type CollateralConfigurationData struct {
	DepositingEnabled    bool
	IssuanceRatioD18     *big.Int
	LiquidationRatioD18  *big.Int
	LiquidationRewardD18 *big.Int
	OracleNodeId         [32]byte
	TokenAddress         common.Address
	MinDelegationD18     *big.Int
}

// CollateralLockData is an auto generated low-level Go binding around an user-defined struct.
type CollateralLockData struct {
	AmountD18          *big.Int
	LockExpirationTime uint64
}

// IAccountModuleAccountPermissions is an auto generated low-level Go binding around an user-defined struct.
type IAccountModuleAccountPermissions struct {
	User        common.Address
	Permissions [][32]byte
}

// ILiquidationModuleLiquidationData is an auto generated low-level Go binding around an user-defined struct.
type ILiquidationModuleLiquidationData struct {
	DebtLiquidated       *big.Int
	CollateralLiquidated *big.Int
	AmountRewarded       *big.Int
}

// MarketConfigurationData is an auto generated low-level Go binding around an user-defined struct.
type MarketConfigurationData struct {
	MarketId             *big.Int
	WeightD18            *big.Int
	MaxDebtShareValueD18 *big.Int
}

// PoolCollateralConfigurationData is an auto generated low-level Go binding around an user-defined struct.
type PoolCollateralConfigurationData struct {
	CollateralLimitD18 *big.Int
	IssuanceRatioD18   *big.Int
}

// CoreOptimismMetaData contains all meta data concerning the CoreOptimism contract.
var CoreOptimismMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"InvalidAccountId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"InvalidPermission\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"OnlyAccountTokenProxy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"PermissionNotGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionOutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedAccountId\",\"type\":\"uint128\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountLastInteraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountPermissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"permissions\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIAccountModule.AccountPermissions[]\",\"name\":\"accountPerms\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"notifyAccountTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"renouncePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyDistribution\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateralRatio\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"MarketNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"NotFundedByPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint128ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint128\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"updatedDebt\",\"type\":\"int256\"}],\"name\":\"DebtAssociated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"associateDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotCcipRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNetwork\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCcipClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCcipClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredTime\",\"type\":\"uint256\"}],\"name\":\"AccountActivityTimeoutPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"CollateralDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAvailableForDelegationD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountD18\",\"type\":\"uint256\"}],\"name\":\"InsufficentAvailableCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InsufficientAccountCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existing\",\"type\":\"uint256\"}],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parameter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"PrecisionLost\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"CollateralLockCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"CollateralLockExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"cleanExpiredLocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cleared\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"createLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getAccountAvailableCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getAccountCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDeposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAssigned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLocked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getLocks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amountD18\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"lockExpirationTime\",\"type\":\"uint64\"}],\"internalType\":\"structCollateralLock.Data[]\",\"name\":\"locks\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"CollateralConfigured\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"configureCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getCollateralConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"hideDisabled\",\"type\":\"bool\"}],\"name\":\"getCollateralConfigurations\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getCollateralPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientCcipFee\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TransferCrossChainInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCrossChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasTokenUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentDebt\",\"type\":\"int256\"}],\"name\":\"InsufficientDebt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"PoolNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"IssuanceFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UsdBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UsdMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CannotScaleEmptyMapping\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"currentCRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cratio\",\"type\":\"uint256\"}],\"name\":\"IneligibleForLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientMappedAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeVaultLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt128ToUint128\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Liquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"VaultLiquidation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"isPositionLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"isVaultLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"}],\"name\":\"liquidate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxUsd\",\"type\":\"uint256\"}],\"name\":\"liquidateVault\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountToDeposit\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarketCollateralDepositable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarketCollateralWithdrawable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MarketCollateralDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MarketCollateralWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"systemAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MaximumMarketCollateralConfigured\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"configureMaximumMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"depositMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getMarketCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralAmountD18\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getMaximumMarketCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"IncorrectMarketInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotEnoughLiquidity\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MarketRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"MarketSystemFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketUsdDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketUsdWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"SetMarketMinLiquidityRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"minDelegateTime\",\"type\":\"uint32\"}],\"name\":\"SetMinDelegateTime\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxIter\",\"type\":\"uint256\"}],\"name\":\"distributeDebtToPools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketDebtPerShare\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getMarketFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFeeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketMinDelegateTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketNetIssuance\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getMarketPoolDebtDistribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesD18\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"totalSharesD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"valuePerShareD27\",\"type\":\"int128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketPools\",\"outputs\":[{\"internalType\":\"uint128[]\",\"name\":\"inRangePoolIds\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"outRangePoolIds\",\"type\":\"uint128[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketReportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketTotalDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMinLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleManager\",\"outputs\":[{\"internalType\":\"contractIOracleManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsdToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getWithdrawableMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"isMarketCapacityLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"registerMarket\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"minDelegateTime\",\"type\":\"uint32\"}],\"name\":\"setMarketMinDelegateTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"setMinLiquidityRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"DeniedMulticallTarget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RecursiveMulticall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"getMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"multicallThrough\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowlisted\",\"type\":\"bool\"}],\"name\":\"setAllowlistedMulticallTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolApprovedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolApprovedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PreferredPoolSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"addApprovedPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApprovedPools\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreferredPool\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"removeApprovedPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"setPreferredPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"CapacityLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"timeRemaining\",\"type\":\"uint32\"}],\"name\":\"MinDelegationTimeoutPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"PoolAlreadyExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collateralLimitD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPoolCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"PoolCollateralConfigurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"name\":\"PoolCollateralDisabledByDefaultSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"indexed\":false,\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"markets\",\"type\":\"tuple[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolConfigurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolNameUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolNominationRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolNominationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolOwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolOwnershipAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"SetMinLiquidityRatio\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"acceptPoolOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedPoolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getNominatedPoolOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getPoolCollateralIssuanceRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"poolName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"nominatePoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"optionalCollateralType\",\"type\":\"address\"}],\"name\":\"rebalancePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"renouncePoolNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"revokePoolNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"setMinLiquidityRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collateralLimitD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"}],\"internalType\":\"structPoolCollateralConfiguration.Data\",\"name\":\"newConfig\",\"type\":\"tuple\"}],\"name\":\"setPoolCollateralConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"name\":\"setPoolCollateralDisabledByDefault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"newMarketConfigurations\",\"type\":\"tuple[]\"}],\"name\":\"setPoolConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setPoolName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint32\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint32ToInt32\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint64ToInt64\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardDistributorNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardUnavailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RewardsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardsDistributorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardsDistributorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"getRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"registerRewardsDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"removeRewardsDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"updateRewards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newChainId\",\"type\":\"uint64\"}],\"name\":\"NewSupportedCrossChainNetwork\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ccipRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ccipTokenPool\",\"type\":\"address\"}],\"name\":\"configureChainlinkCrossChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleManagerAddress\",\"type\":\"address\"}],\"name\":\"configureOracleManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"v\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfigAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfigUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"v\",\"type\":\"bytes32\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"supportedNetworks\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"ccipSelectors\",\"type\":\"uint64[]\"}],\"name\":\"setSupportedCrossChainNetworks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numRegistered\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minDelegation\",\"type\":\"uint256\"}],\"name\":\"InsufficientDelegation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateralAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"}],\"name\":\"InvalidLeverage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCollateral\",\"type\":\"uint256\"}],\"name\":\"PoolCollateralLimitExceeded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"DelegationUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralAmountD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"}],\"name\":\"delegateCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"collateralizationRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionCollateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultCollateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CoreOptimismABI is the input ABI used to generate the binding from.
// Deprecated: Use CoreOptimismMetaData.ABI instead.
var CoreOptimismABI = CoreOptimismMetaData.ABI

// CoreOptimism is an auto generated Go binding around an Ethereum contract.
type CoreOptimism struct {
	CoreOptimismCaller     // Read-only binding to the contract
	CoreOptimismTransactor // Write-only binding to the contract
	CoreOptimismFilterer   // Log filterer for contract events
}

// CoreOptimismCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoreOptimismCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreOptimismTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoreOptimismTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreOptimismFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoreOptimismFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreOptimismSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoreOptimismSession struct {
	Contract     *CoreOptimism     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreOptimismCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoreOptimismCallerSession struct {
	Contract *CoreOptimismCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CoreOptimismTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoreOptimismTransactorSession struct {
	Contract     *CoreOptimismTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CoreOptimismRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoreOptimismRaw struct {
	Contract *CoreOptimism // Generic contract binding to access the raw methods on
}

// CoreOptimismCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoreOptimismCallerRaw struct {
	Contract *CoreOptimismCaller // Generic read-only contract binding to access the raw methods on
}

// CoreOptimismTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoreOptimismTransactorRaw struct {
	Contract *CoreOptimismTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoreOptimism creates a new instance of CoreOptimism, bound to a specific deployed contract.
func NewCoreOptimism(address common.Address, backend bind.ContractBackend) (*CoreOptimism, error) {
	contract, err := bindCoreOptimism(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoreOptimism{CoreOptimismCaller: CoreOptimismCaller{contract: contract}, CoreOptimismTransactor: CoreOptimismTransactor{contract: contract}, CoreOptimismFilterer: CoreOptimismFilterer{contract: contract}}, nil
}

// NewCoreOptimismCaller creates a new read-only instance of CoreOptimism, bound to a specific deployed contract.
func NewCoreOptimismCaller(address common.Address, caller bind.ContractCaller) (*CoreOptimismCaller, error) {
	contract, err := bindCoreOptimism(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismCaller{contract: contract}, nil
}

// NewCoreOptimismTransactor creates a new write-only instance of CoreOptimism, bound to a specific deployed contract.
func NewCoreOptimismTransactor(address common.Address, transactor bind.ContractTransactor) (*CoreOptimismTransactor, error) {
	contract, err := bindCoreOptimism(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismTransactor{contract: contract}, nil
}

// NewCoreOptimismFilterer creates a new log filterer instance of CoreOptimism, bound to a specific deployed contract.
func NewCoreOptimismFilterer(address common.Address, filterer bind.ContractFilterer) (*CoreOptimismFilterer, error) {
	contract, err := bindCoreOptimism(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismFilterer{contract: contract}, nil
}

// bindCoreOptimism binds a generic wrapper to an already deployed contract.
func bindCoreOptimism(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoreOptimismMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoreOptimism *CoreOptimismRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoreOptimism.Contract.CoreOptimismCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoreOptimism *CoreOptimismRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CoreOptimismTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoreOptimism *CoreOptimismRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CoreOptimismTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoreOptimism *CoreOptimismCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoreOptimism.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoreOptimism *CoreOptimismTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreOptimism.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoreOptimism *CoreOptimismTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoreOptimism.Contract.contract.Transact(opts, method, params...)
}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetAccountAvailableCollateral(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getAccountAvailableCollateral", accountId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetAccountAvailableCollateral(&_CoreOptimism.CallOpts, accountId, collateralType)
}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetAccountAvailableCollateral(&_CoreOptimism.CallOpts, accountId, collateralType)
}

// GetAccountCollateral is a free data retrieval call binding the contract method 0xef45148e.
//
// Solidity: function getAccountCollateral(uint128 accountId, address collateralType) view returns(uint256 totalDeposited, uint256 totalAssigned, uint256 totalLocked)
func (_CoreOptimism *CoreOptimismCaller) GetAccountCollateral(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getAccountCollateral", accountId, collateralType)

	outstruct := new(struct {
		TotalDeposited *big.Int
		TotalAssigned  *big.Int
		TotalLocked    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalDeposited = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalAssigned = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalLocked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAccountCollateral is a free data retrieval call binding the contract method 0xef45148e.
//
// Solidity: function getAccountCollateral(uint128 accountId, address collateralType) view returns(uint256 totalDeposited, uint256 totalAssigned, uint256 totalLocked)
func (_CoreOptimism *CoreOptimismSession) GetAccountCollateral(accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	return _CoreOptimism.Contract.GetAccountCollateral(&_CoreOptimism.CallOpts, accountId, collateralType)
}

// GetAccountCollateral is a free data retrieval call binding the contract method 0xef45148e.
//
// Solidity: function getAccountCollateral(uint128 accountId, address collateralType) view returns(uint256 totalDeposited, uint256 totalAssigned, uint256 totalLocked)
func (_CoreOptimism *CoreOptimismCallerSession) GetAccountCollateral(accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	return _CoreOptimism.Contract.GetAccountCollateral(&_CoreOptimism.CallOpts, accountId, collateralType)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetAccountLastInteraction(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getAccountLastInteraction", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetAccountLastInteraction(&_CoreOptimism.CallOpts, accountId)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetAccountLastInteraction(&_CoreOptimism.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_CoreOptimism *CoreOptimismCaller) GetAccountOwner(opts *bind.CallOpts, accountId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getAccountOwner", accountId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_CoreOptimism *CoreOptimismSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _CoreOptimism.Contract.GetAccountOwner(&_CoreOptimism.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _CoreOptimism.Contract.GetAccountOwner(&_CoreOptimism.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_CoreOptimism *CoreOptimismCaller) GetAccountPermissions(opts *bind.CallOpts, accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getAccountPermissions", accountId)

	if err != nil {
		return *new([]IAccountModuleAccountPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAccountModuleAccountPermissions)).(*[]IAccountModuleAccountPermissions)

	return out0, err

}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_CoreOptimism *CoreOptimismSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _CoreOptimism.Contract.GetAccountPermissions(&_CoreOptimism.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_CoreOptimism *CoreOptimismCallerSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _CoreOptimism.Contract.GetAccountPermissions(&_CoreOptimism.CallOpts, accountId)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_CoreOptimism *CoreOptimismCaller) GetAccountTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getAccountTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_CoreOptimism *CoreOptimismSession) GetAccountTokenAddress() (common.Address, error) {
	return _CoreOptimism.Contract.GetAccountTokenAddress(&_CoreOptimism.CallOpts)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) GetAccountTokenAddress() (common.Address, error) {
	return _CoreOptimism.Contract.GetAccountTokenAddress(&_CoreOptimism.CallOpts)
}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_CoreOptimism *CoreOptimismCaller) GetApprovedPools(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getApprovedPools")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_CoreOptimism *CoreOptimismSession) GetApprovedPools() ([]*big.Int, error) {
	return _CoreOptimism.Contract.GetApprovedPools(&_CoreOptimism.CallOpts)
}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_CoreOptimism *CoreOptimismCallerSession) GetApprovedPools() ([]*big.Int, error) {
	return _CoreOptimism.Contract.GetApprovedPools(&_CoreOptimism.CallOpts)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_CoreOptimism *CoreOptimismCaller) GetAssociatedSystem(opts *bind.CallOpts, id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getAssociatedSystem", id)

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
func (_CoreOptimism *CoreOptimismSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _CoreOptimism.Contract.GetAssociatedSystem(&_CoreOptimism.CallOpts, id)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_CoreOptimism *CoreOptimismCallerSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _CoreOptimism.Contract.GetAssociatedSystem(&_CoreOptimism.CallOpts, id)
}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) pure returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_CoreOptimism *CoreOptimismCaller) GetCollateralConfiguration(opts *bind.CallOpts, collateralType common.Address) (CollateralConfigurationData, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getCollateralConfiguration", collateralType)

	if err != nil {
		return *new(CollateralConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new(CollateralConfigurationData)).(*CollateralConfigurationData)

	return out0, err

}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) pure returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_CoreOptimism *CoreOptimismSession) GetCollateralConfiguration(collateralType common.Address) (CollateralConfigurationData, error) {
	return _CoreOptimism.Contract.GetCollateralConfiguration(&_CoreOptimism.CallOpts, collateralType)
}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) pure returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_CoreOptimism *CoreOptimismCallerSession) GetCollateralConfiguration(collateralType common.Address) (CollateralConfigurationData, error) {
	return _CoreOptimism.Contract.GetCollateralConfiguration(&_CoreOptimism.CallOpts, collateralType)
}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_CoreOptimism *CoreOptimismCaller) GetCollateralConfigurations(opts *bind.CallOpts, hideDisabled bool) ([]CollateralConfigurationData, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getCollateralConfigurations", hideDisabled)

	if err != nil {
		return *new([]CollateralConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CollateralConfigurationData)).(*[]CollateralConfigurationData)

	return out0, err

}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_CoreOptimism *CoreOptimismSession) GetCollateralConfigurations(hideDisabled bool) ([]CollateralConfigurationData, error) {
	return _CoreOptimism.Contract.GetCollateralConfigurations(&_CoreOptimism.CallOpts, hideDisabled)
}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_CoreOptimism *CoreOptimismCallerSession) GetCollateralConfigurations(hideDisabled bool) ([]CollateralConfigurationData, error) {
	return _CoreOptimism.Contract.GetCollateralConfigurations(&_CoreOptimism.CallOpts, hideDisabled)
}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetCollateralPrice(opts *bind.CallOpts, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getCollateralPrice", collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetCollateralPrice(collateralType common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetCollateralPrice(&_CoreOptimism.CallOpts, collateralType)
}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetCollateralPrice(collateralType common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetCollateralPrice(&_CoreOptimism.CallOpts, collateralType)
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_CoreOptimism *CoreOptimismCaller) GetConfig(opts *bind.CallOpts, k [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getConfig", k)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_CoreOptimism *CoreOptimismSession) GetConfig(k [32]byte) ([32]byte, error) {
	return _CoreOptimism.Contract.GetConfig(&_CoreOptimism.CallOpts, k)
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_CoreOptimism *CoreOptimismCallerSession) GetConfig(k [32]byte) ([32]byte, error) {
	return _CoreOptimism.Contract.GetConfig(&_CoreOptimism.CallOpts, k)
}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_CoreOptimism *CoreOptimismCaller) GetConfigAddress(opts *bind.CallOpts, k [32]byte) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getConfigAddress", k)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_CoreOptimism *CoreOptimismSession) GetConfigAddress(k [32]byte) (common.Address, error) {
	return _CoreOptimism.Contract.GetConfigAddress(&_CoreOptimism.CallOpts, k)
}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_CoreOptimism *CoreOptimismCallerSession) GetConfigAddress(k [32]byte) (common.Address, error) {
	return _CoreOptimism.Contract.GetConfigAddress(&_CoreOptimism.CallOpts, k)
}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_CoreOptimism *CoreOptimismCaller) GetConfigUint(opts *bind.CallOpts, k [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getConfigUint", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_CoreOptimism *CoreOptimismSession) GetConfigUint(k [32]byte) (*big.Int, error) {
	return _CoreOptimism.Contract.GetConfigUint(&_CoreOptimism.CallOpts, k)
}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_CoreOptimism *CoreOptimismCallerSession) GetConfigUint(k [32]byte) (*big.Int, error) {
	return _CoreOptimism.Contract.GetConfigUint(&_CoreOptimism.CallOpts, k)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_CoreOptimism *CoreOptimismCaller) GetDeniers(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getDeniers", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_CoreOptimism *CoreOptimismSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _CoreOptimism.Contract.GetDeniers(&_CoreOptimism.CallOpts, feature)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_CoreOptimism *CoreOptimismCallerSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _CoreOptimism.Contract.GetDeniers(&_CoreOptimism.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_CoreOptimism *CoreOptimismCaller) GetFeatureFlagAllowAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getFeatureFlagAllowAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_CoreOptimism *CoreOptimismSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _CoreOptimism.Contract.GetFeatureFlagAllowAll(&_CoreOptimism.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_CoreOptimism *CoreOptimismCallerSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _CoreOptimism.Contract.GetFeatureFlagAllowAll(&_CoreOptimism.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_CoreOptimism *CoreOptimismCaller) GetFeatureFlagAllowlist(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getFeatureFlagAllowlist", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_CoreOptimism *CoreOptimismSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _CoreOptimism.Contract.GetFeatureFlagAllowlist(&_CoreOptimism.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_CoreOptimism *CoreOptimismCallerSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _CoreOptimism.Contract.GetFeatureFlagAllowlist(&_CoreOptimism.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_CoreOptimism *CoreOptimismCaller) GetFeatureFlagDenyAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getFeatureFlagDenyAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_CoreOptimism *CoreOptimismSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _CoreOptimism.Contract.GetFeatureFlagDenyAll(&_CoreOptimism.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_CoreOptimism *CoreOptimismCallerSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _CoreOptimism.Contract.GetFeatureFlagDenyAll(&_CoreOptimism.CallOpts, feature)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_CoreOptimism *CoreOptimismCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_CoreOptimism *CoreOptimismSession) GetImplementation() (common.Address, error) {
	return _CoreOptimism.Contract.GetImplementation(&_CoreOptimism.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) GetImplementation() (common.Address, error) {
	return _CoreOptimism.Contract.GetImplementation(&_CoreOptimism.CallOpts)
}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_CoreOptimism *CoreOptimismCaller) GetLocks(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getLocks", accountId, collateralType, offset, count)

	if err != nil {
		return *new([]CollateralLockData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CollateralLockData)).(*[]CollateralLockData)

	return out0, err

}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_CoreOptimism *CoreOptimismSession) GetLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	return _CoreOptimism.Contract.GetLocks(&_CoreOptimism.CallOpts, accountId, collateralType, offset, count)
}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_CoreOptimism *CoreOptimismCallerSession) GetLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	return _CoreOptimism.Contract.GetLocks(&_CoreOptimism.CallOpts, accountId, collateralType, offset, count)
}

// GetMarketAddress is a free data retrieval call binding the contract method 0xd24437f1.
//
// Solidity: function getMarketAddress(uint128 marketId) view returns(address)
func (_CoreOptimism *CoreOptimismCaller) GetMarketAddress(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMarketAddress", marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMarketAddress is a free data retrieval call binding the contract method 0xd24437f1.
//
// Solidity: function getMarketAddress(uint128 marketId) view returns(address)
func (_CoreOptimism *CoreOptimismSession) GetMarketAddress(marketId *big.Int) (common.Address, error) {
	return _CoreOptimism.Contract.GetMarketAddress(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketAddress is a free data retrieval call binding the contract method 0xd24437f1.
//
// Solidity: function getMarketAddress(uint128 marketId) view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) GetMarketAddress(marketId *big.Int) (common.Address, error) {
	return _CoreOptimism.Contract.GetMarketAddress(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetMarketCollateral(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMarketCollateral", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetMarketCollateral(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketCollateral(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetMarketCollateral(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketCollateral(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_CoreOptimism *CoreOptimismCaller) GetMarketCollateralAmount(opts *bind.CallOpts, marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMarketCollateralAmount", marketId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_CoreOptimism *CoreOptimismSession) GetMarketCollateralAmount(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketCollateralAmount(&_CoreOptimism.CallOpts, marketId, collateralType)
}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_CoreOptimism *CoreOptimismCallerSession) GetMarketCollateralAmount(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketCollateralAmount(&_CoreOptimism.CallOpts, marketId, collateralType)
}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetMarketCollateralValue(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMarketCollateralValue", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetMarketCollateralValue(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketCollateralValue(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetMarketCollateralValue(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketCollateralValue(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketFees is a free data retrieval call binding the contract method 0xdfb83437.
//
// Solidity: function getMarketFees(uint128 , uint256 amount) view returns(uint256 depositFeeAmount, uint256 withdrawFeeAmount)
func (_CoreOptimism *CoreOptimismCaller) GetMarketFees(opts *bind.CallOpts, arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMarketFees", arg0, amount)

	outstruct := new(struct {
		DepositFeeAmount  *big.Int
		WithdrawFeeAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositFeeAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawFeeAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMarketFees is a free data retrieval call binding the contract method 0xdfb83437.
//
// Solidity: function getMarketFees(uint128 , uint256 amount) view returns(uint256 depositFeeAmount, uint256 withdrawFeeAmount)
func (_CoreOptimism *CoreOptimismSession) GetMarketFees(arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	return _CoreOptimism.Contract.GetMarketFees(&_CoreOptimism.CallOpts, arg0, amount)
}

// GetMarketFees is a free data retrieval call binding the contract method 0xdfb83437.
//
// Solidity: function getMarketFees(uint128 , uint256 amount) view returns(uint256 depositFeeAmount, uint256 withdrawFeeAmount)
func (_CoreOptimism *CoreOptimismCallerSession) GetMarketFees(arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	return _CoreOptimism.Contract.GetMarketFees(&_CoreOptimism.CallOpts, arg0, amount)
}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_CoreOptimism *CoreOptimismCaller) GetMarketMinDelegateTime(opts *bind.CallOpts, marketId *big.Int) (uint32, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMarketMinDelegateTime", marketId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_CoreOptimism *CoreOptimismSession) GetMarketMinDelegateTime(marketId *big.Int) (uint32, error) {
	return _CoreOptimism.Contract.GetMarketMinDelegateTime(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_CoreOptimism *CoreOptimismCallerSession) GetMarketMinDelegateTime(marketId *big.Int) (uint32, error) {
	return _CoreOptimism.Contract.GetMarketMinDelegateTime(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_CoreOptimism *CoreOptimismCaller) GetMarketNetIssuance(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMarketNetIssuance", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_CoreOptimism *CoreOptimismSession) GetMarketNetIssuance(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketNetIssuance(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_CoreOptimism *CoreOptimismCallerSession) GetMarketNetIssuance(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketNetIssuance(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetMarketReportedDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMarketReportedDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetMarketReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketReportedDebt(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetMarketReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketReportedDebt(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_CoreOptimism *CoreOptimismCaller) GetMarketTotalDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMarketTotalDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_CoreOptimism *CoreOptimismSession) GetMarketTotalDebt(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketTotalDebt(&_CoreOptimism.CallOpts, marketId)
}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_CoreOptimism *CoreOptimismCallerSession) GetMarketTotalDebt(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMarketTotalDebt(&_CoreOptimism.CallOpts, marketId)
}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetMaximumMarketCollateral(opts *bind.CallOpts, marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMaximumMarketCollateral", marketId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetMaximumMarketCollateral(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMaximumMarketCollateral(&_CoreOptimism.CallOpts, marketId, collateralType)
}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetMaximumMarketCollateral(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMaximumMarketCollateral(&_CoreOptimism.CallOpts, marketId, collateralType)
}

// GetMessageSender is a free data retrieval call binding the contract method 0x76167a89.
//
// Solidity: function getMessageSender() view returns(address)
func (_CoreOptimism *CoreOptimismCaller) GetMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMessageSender is a free data retrieval call binding the contract method 0x76167a89.
//
// Solidity: function getMessageSender() view returns(address)
func (_CoreOptimism *CoreOptimismSession) GetMessageSender() (common.Address, error) {
	return _CoreOptimism.Contract.GetMessageSender(&_CoreOptimism.CallOpts)
}

// GetMessageSender is a free data retrieval call binding the contract method 0x76167a89.
//
// Solidity: function getMessageSender() view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) GetMessageSender() (common.Address, error) {
	return _CoreOptimism.Contract.GetMessageSender(&_CoreOptimism.CallOpts)
}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetMinLiquidityRatio(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMinLiquidityRatio", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetMinLiquidityRatio(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMinLiquidityRatio(&_CoreOptimism.CallOpts, marketId)
}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetMinLiquidityRatio(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetMinLiquidityRatio(&_CoreOptimism.CallOpts, marketId)
}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetMinLiquidityRatio0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getMinLiquidityRatio0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetMinLiquidityRatio0() (*big.Int, error) {
	return _CoreOptimism.Contract.GetMinLiquidityRatio0(&_CoreOptimism.CallOpts)
}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetMinLiquidityRatio0() (*big.Int, error) {
	return _CoreOptimism.Contract.GetMinLiquidityRatio0(&_CoreOptimism.CallOpts)
}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_CoreOptimism *CoreOptimismCaller) GetNominatedPoolOwner(opts *bind.CallOpts, poolId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getNominatedPoolOwner", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_CoreOptimism *CoreOptimismSession) GetNominatedPoolOwner(poolId *big.Int) (common.Address, error) {
	return _CoreOptimism.Contract.GetNominatedPoolOwner(&_CoreOptimism.CallOpts, poolId)
}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) GetNominatedPoolOwner(poolId *big.Int) (common.Address, error) {
	return _CoreOptimism.Contract.GetNominatedPoolOwner(&_CoreOptimism.CallOpts, poolId)
}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_CoreOptimism *CoreOptimismCaller) GetOracleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getOracleManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_CoreOptimism *CoreOptimismSession) GetOracleManager() (common.Address, error) {
	return _CoreOptimism.Contract.GetOracleManager(&_CoreOptimism.CallOpts)
}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) GetOracleManager() (common.Address, error) {
	return _CoreOptimism.Contract.GetOracleManager(&_CoreOptimism.CallOpts)
}

// GetPoolCollateralIssuanceRatio is a free data retrieval call binding the contract method 0xc4d2aad3.
//
// Solidity: function getPoolCollateralIssuanceRatio(uint128 poolId, address collateral) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetPoolCollateralIssuanceRatio(opts *bind.CallOpts, poolId *big.Int, collateral common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getPoolCollateralIssuanceRatio", poolId, collateral)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolCollateralIssuanceRatio is a free data retrieval call binding the contract method 0xc4d2aad3.
//
// Solidity: function getPoolCollateralIssuanceRatio(uint128 poolId, address collateral) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetPoolCollateralIssuanceRatio(poolId *big.Int, collateral common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetPoolCollateralIssuanceRatio(&_CoreOptimism.CallOpts, poolId, collateral)
}

// GetPoolCollateralIssuanceRatio is a free data retrieval call binding the contract method 0xc4d2aad3.
//
// Solidity: function getPoolCollateralIssuanceRatio(uint128 poolId, address collateral) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetPoolCollateralIssuanceRatio(poolId *big.Int, collateral common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetPoolCollateralIssuanceRatio(&_CoreOptimism.CallOpts, poolId, collateral)
}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_CoreOptimism *CoreOptimismCaller) GetPoolConfiguration(opts *bind.CallOpts, poolId *big.Int) ([]MarketConfigurationData, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getPoolConfiguration", poolId)

	if err != nil {
		return *new([]MarketConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketConfigurationData)).(*[]MarketConfigurationData)

	return out0, err

}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_CoreOptimism *CoreOptimismSession) GetPoolConfiguration(poolId *big.Int) ([]MarketConfigurationData, error) {
	return _CoreOptimism.Contract.GetPoolConfiguration(&_CoreOptimism.CallOpts, poolId)
}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_CoreOptimism *CoreOptimismCallerSession) GetPoolConfiguration(poolId *big.Int) ([]MarketConfigurationData, error) {
	return _CoreOptimism.Contract.GetPoolConfiguration(&_CoreOptimism.CallOpts, poolId)
}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_CoreOptimism *CoreOptimismCaller) GetPoolName(opts *bind.CallOpts, poolId *big.Int) (string, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getPoolName", poolId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_CoreOptimism *CoreOptimismSession) GetPoolName(poolId *big.Int) (string, error) {
	return _CoreOptimism.Contract.GetPoolName(&_CoreOptimism.CallOpts, poolId)
}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_CoreOptimism *CoreOptimismCallerSession) GetPoolName(poolId *big.Int) (string, error) {
	return _CoreOptimism.Contract.GetPoolName(&_CoreOptimism.CallOpts, poolId)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_CoreOptimism *CoreOptimismCaller) GetPoolOwner(opts *bind.CallOpts, poolId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getPoolOwner", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_CoreOptimism *CoreOptimismSession) GetPoolOwner(poolId *big.Int) (common.Address, error) {
	return _CoreOptimism.Contract.GetPoolOwner(&_CoreOptimism.CallOpts, poolId)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) GetPoolOwner(poolId *big.Int) (common.Address, error) {
	return _CoreOptimism.Contract.GetPoolOwner(&_CoreOptimism.CallOpts, poolId)
}

// GetPositionCollateral is a free data retrieval call binding the contract method 0x33cc422b.
//
// Solidity: function getPositionCollateral(uint128 accountId, uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_CoreOptimism *CoreOptimismCaller) GetPositionCollateral(opts *bind.CallOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getPositionCollateral", accountId, poolId, collateralType)

	outstruct := new(struct {
		Amount *big.Int
		Value  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPositionCollateral is a free data retrieval call binding the contract method 0x33cc422b.
//
// Solidity: function getPositionCollateral(uint128 accountId, uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_CoreOptimism *CoreOptimismSession) GetPositionCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _CoreOptimism.Contract.GetPositionCollateral(&_CoreOptimism.CallOpts, accountId, poolId, collateralType)
}

// GetPositionCollateral is a free data retrieval call binding the contract method 0x33cc422b.
//
// Solidity: function getPositionCollateral(uint128 accountId, uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_CoreOptimism *CoreOptimismCallerSession) GetPositionCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _CoreOptimism.Contract.GetPositionCollateral(&_CoreOptimism.CallOpts, accountId, poolId, collateralType)
}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_CoreOptimism *CoreOptimismCaller) GetPreferredPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getPreferredPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_CoreOptimism *CoreOptimismSession) GetPreferredPool() (*big.Int, error) {
	return _CoreOptimism.Contract.GetPreferredPool(&_CoreOptimism.CallOpts)
}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_CoreOptimism *CoreOptimismCallerSession) GetPreferredPool() (*big.Int, error) {
	return _CoreOptimism.Contract.GetPreferredPool(&_CoreOptimism.CallOpts)
}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetRewardRate(opts *bind.CallOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getRewardRate", poolId, collateralType, distributor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetRewardRate(poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetRewardRate(&_CoreOptimism.CallOpts, poolId, collateralType, distributor)
}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetRewardRate(poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	return _CoreOptimism.Contract.GetRewardRate(&_CoreOptimism.CallOpts, poolId, collateralType, distributor)
}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_CoreOptimism *CoreOptimismCaller) GetUsdToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getUsdToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_CoreOptimism *CoreOptimismSession) GetUsdToken() (common.Address, error) {
	return _CoreOptimism.Contract.GetUsdToken(&_CoreOptimism.CallOpts)
}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) GetUsdToken() (common.Address, error) {
	return _CoreOptimism.Contract.GetUsdToken(&_CoreOptimism.CallOpts)
}

// GetVaultCollateral is a free data retrieval call binding the contract method 0x078145a8.
//
// Solidity: function getVaultCollateral(uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_CoreOptimism *CoreOptimismCaller) GetVaultCollateral(opts *bind.CallOpts, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getVaultCollateral", poolId, collateralType)

	outstruct := new(struct {
		Amount *big.Int
		Value  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVaultCollateral is a free data retrieval call binding the contract method 0x078145a8.
//
// Solidity: function getVaultCollateral(uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_CoreOptimism *CoreOptimismSession) GetVaultCollateral(poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _CoreOptimism.Contract.GetVaultCollateral(&_CoreOptimism.CallOpts, poolId, collateralType)
}

// GetVaultCollateral is a free data retrieval call binding the contract method 0x078145a8.
//
// Solidity: function getVaultCollateral(uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_CoreOptimism *CoreOptimismCallerSession) GetVaultCollateral(poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _CoreOptimism.Contract.GetVaultCollateral(&_CoreOptimism.CallOpts, poolId, collateralType)
}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCaller) GetWithdrawableMarketUsd(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "getWithdrawableMarketUsd", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetWithdrawableMarketUsd(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetWithdrawableMarketUsd(&_CoreOptimism.CallOpts, marketId)
}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_CoreOptimism *CoreOptimismCallerSession) GetWithdrawableMarketUsd(marketId *big.Int) (*big.Int, error) {
	return _CoreOptimism.Contract.GetWithdrawableMarketUsd(&_CoreOptimism.CallOpts, marketId)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreOptimism *CoreOptimismCaller) HasPermission(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "hasPermission", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreOptimism *CoreOptimismSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _CoreOptimism.Contract.HasPermission(&_CoreOptimism.CallOpts, accountId, permission, user)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreOptimism *CoreOptimismCallerSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _CoreOptimism.Contract.HasPermission(&_CoreOptimism.CallOpts, accountId, permission, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreOptimism *CoreOptimismCaller) IsAuthorized(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "isAuthorized", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreOptimism *CoreOptimismSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _CoreOptimism.Contract.IsAuthorized(&_CoreOptimism.CallOpts, accountId, permission, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreOptimism *CoreOptimismCallerSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _CoreOptimism.Contract.IsAuthorized(&_CoreOptimism.CallOpts, accountId, permission, user)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_CoreOptimism *CoreOptimismCaller) IsFeatureAllowed(opts *bind.CallOpts, feature [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "isFeatureAllowed", feature, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_CoreOptimism *CoreOptimismSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _CoreOptimism.Contract.IsFeatureAllowed(&_CoreOptimism.CallOpts, feature, account)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_CoreOptimism *CoreOptimismCallerSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _CoreOptimism.Contract.IsFeatureAllowed(&_CoreOptimism.CallOpts, feature, account)
}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_CoreOptimism *CoreOptimismCaller) IsMarketCapacityLocked(opts *bind.CallOpts, marketId *big.Int) (bool, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "isMarketCapacityLocked", marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_CoreOptimism *CoreOptimismSession) IsMarketCapacityLocked(marketId *big.Int) (bool, error) {
	return _CoreOptimism.Contract.IsMarketCapacityLocked(&_CoreOptimism.CallOpts, marketId)
}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_CoreOptimism *CoreOptimismCallerSession) IsMarketCapacityLocked(marketId *big.Int) (bool, error) {
	return _CoreOptimism.Contract.IsMarketCapacityLocked(&_CoreOptimism.CallOpts, marketId)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_CoreOptimism *CoreOptimismCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_CoreOptimism *CoreOptimismSession) NominatedOwner() (common.Address, error) {
	return _CoreOptimism.Contract.NominatedOwner(&_CoreOptimism.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) NominatedOwner() (common.Address, error) {
	return _CoreOptimism.Contract.NominatedOwner(&_CoreOptimism.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CoreOptimism *CoreOptimismCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CoreOptimism *CoreOptimismSession) Owner() (common.Address, error) {
	return _CoreOptimism.Contract.Owner(&_CoreOptimism.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CoreOptimism *CoreOptimismCallerSession) Owner() (common.Address, error) {
	return _CoreOptimism.Contract.Owner(&_CoreOptimism.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoreOptimism *CoreOptimismCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CoreOptimism.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoreOptimism *CoreOptimismSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CoreOptimism.Contract.SupportsInterface(&_CoreOptimism.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoreOptimism *CoreOptimismCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CoreOptimism.Contract.SupportsInterface(&_CoreOptimism.CallOpts, interfaceId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_CoreOptimism *CoreOptimismTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_CoreOptimism *CoreOptimismSession) AcceptOwnership() (*types.Transaction, error) {
	return _CoreOptimism.Contract.AcceptOwnership(&_CoreOptimism.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_CoreOptimism *CoreOptimismTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CoreOptimism.Contract.AcceptOwnership(&_CoreOptimism.TransactOpts)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactor) AcceptPoolOwnership(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "acceptPoolOwnership", poolId)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismSession) AcceptPoolOwnership(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.AcceptPoolOwnership(&_CoreOptimism.TransactOpts, poolId)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) AcceptPoolOwnership(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.AcceptPoolOwnership(&_CoreOptimism.TransactOpts, poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactor) AddApprovedPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "addApprovedPool", poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismSession) AddApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.AddApprovedPool(&_CoreOptimism.TransactOpts, poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) AddApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.AddApprovedPool(&_CoreOptimism.TransactOpts, poolId)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreOptimism *CoreOptimismTransactor) AddToFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "addToFeatureFlagAllowlist", feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreOptimism *CoreOptimismSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.AddToFeatureFlagAllowlist(&_CoreOptimism.TransactOpts, feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.AddToFeatureFlagAllowlist(&_CoreOptimism.TransactOpts, feature, account)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_CoreOptimism *CoreOptimismTransactor) AssociateDebt(opts *bind.TransactOpts, marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "associateDebt", marketId, poolId, collateralType, accountId, amount)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_CoreOptimism *CoreOptimismSession) AssociateDebt(marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.AssociateDebt(&_CoreOptimism.TransactOpts, marketId, poolId, collateralType, accountId, amount)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_CoreOptimism *CoreOptimismTransactorSession) AssociateDebt(marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.AssociateDebt(&_CoreOptimism.TransactOpts, marketId, poolId, collateralType, accountId, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreOptimism *CoreOptimismTransactor) BurnUsd(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "burnUsd", accountId, poolId, collateralType, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreOptimism *CoreOptimismSession) BurnUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.BurnUsd(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) BurnUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.BurnUsd(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType, amount)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_CoreOptimism *CoreOptimismTransactor) CcipReceive(opts *bind.TransactOpts, message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_CoreOptimism *CoreOptimismSession) CcipReceive(message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CcipReceive(&_CoreOptimism.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) CcipReceive(message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CcipReceive(&_CoreOptimism.TransactOpts, message)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_CoreOptimism *CoreOptimismTransactor) ClaimRewards(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "claimRewards", accountId, poolId, collateralType, distributor)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_CoreOptimism *CoreOptimismSession) ClaimRewards(accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.ClaimRewards(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType, distributor)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_CoreOptimism *CoreOptimismTransactorSession) ClaimRewards(accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.ClaimRewards(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType, distributor)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_CoreOptimism *CoreOptimismTransactor) CleanExpiredLocks(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "cleanExpiredLocks", accountId, collateralType, offset, count)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_CoreOptimism *CoreOptimismSession) CleanExpiredLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CleanExpiredLocks(&_CoreOptimism.TransactOpts, accountId, collateralType, offset, count)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_CoreOptimism *CoreOptimismTransactorSession) CleanExpiredLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CleanExpiredLocks(&_CoreOptimism.TransactOpts, accountId, collateralType, offset, count)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_CoreOptimism *CoreOptimismTransactor) ConfigureChainlinkCrossChain(opts *bind.TransactOpts, ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "configureChainlinkCrossChain", ccipRouter, ccipTokenPool)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_CoreOptimism *CoreOptimismSession) ConfigureChainlinkCrossChain(ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.ConfigureChainlinkCrossChain(&_CoreOptimism.TransactOpts, ccipRouter, ccipTokenPool)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) ConfigureChainlinkCrossChain(ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.ConfigureChainlinkCrossChain(&_CoreOptimism.TransactOpts, ccipRouter, ccipTokenPool)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_CoreOptimism *CoreOptimismTransactor) ConfigureCollateral(opts *bind.TransactOpts, config CollateralConfigurationData) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "configureCollateral", config)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_CoreOptimism *CoreOptimismSession) ConfigureCollateral(config CollateralConfigurationData) (*types.Transaction, error) {
	return _CoreOptimism.Contract.ConfigureCollateral(&_CoreOptimism.TransactOpts, config)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) ConfigureCollateral(config CollateralConfigurationData) (*types.Transaction, error) {
	return _CoreOptimism.Contract.ConfigureCollateral(&_CoreOptimism.TransactOpts, config)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_CoreOptimism *CoreOptimismTransactor) ConfigureMaximumMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "configureMaximumMarketCollateral", marketId, collateralType, amount)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_CoreOptimism *CoreOptimismSession) ConfigureMaximumMarketCollateral(marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.ConfigureMaximumMarketCollateral(&_CoreOptimism.TransactOpts, marketId, collateralType, amount)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) ConfigureMaximumMarketCollateral(marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.ConfigureMaximumMarketCollateral(&_CoreOptimism.TransactOpts, marketId, collateralType, amount)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_CoreOptimism *CoreOptimismTransactor) ConfigureOracleManager(opts *bind.TransactOpts, oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "configureOracleManager", oracleManagerAddress)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_CoreOptimism *CoreOptimismSession) ConfigureOracleManager(oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.ConfigureOracleManager(&_CoreOptimism.TransactOpts, oracleManagerAddress)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) ConfigureOracleManager(oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.ConfigureOracleManager(&_CoreOptimism.TransactOpts, oracleManagerAddress)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_CoreOptimism *CoreOptimismTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_CoreOptimism *CoreOptimismSession) CreateAccount() (*types.Transaction, error) {
	return _CoreOptimism.Contract.CreateAccount(&_CoreOptimism.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_CoreOptimism *CoreOptimismTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _CoreOptimism.Contract.CreateAccount(&_CoreOptimism.TransactOpts)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_CoreOptimism *CoreOptimismTransactor) CreateAccount0(opts *bind.TransactOpts, requestedAccountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "createAccount0", requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_CoreOptimism *CoreOptimismSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CreateAccount0(&_CoreOptimism.TransactOpts, requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CreateAccount0(&_CoreOptimism.TransactOpts, requestedAccountId)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_CoreOptimism *CoreOptimismTransactor) CreateLock(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "createLock", accountId, collateralType, amount, expireTimestamp)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_CoreOptimism *CoreOptimismSession) CreateLock(accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CreateLock(&_CoreOptimism.TransactOpts, accountId, collateralType, amount, expireTimestamp)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) CreateLock(accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CreateLock(&_CoreOptimism.TransactOpts, accountId, collateralType, amount, expireTimestamp)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_CoreOptimism *CoreOptimismTransactor) CreatePool(opts *bind.TransactOpts, requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "createPool", requestedPoolId, owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_CoreOptimism *CoreOptimismSession) CreatePool(requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CreatePool(&_CoreOptimism.TransactOpts, requestedPoolId, owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) CreatePool(requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.CreatePool(&_CoreOptimism.TransactOpts, requestedPoolId, owner)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_CoreOptimism *CoreOptimismTransactor) DelegateCollateral(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "delegateCollateral", accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_CoreOptimism *CoreOptimismSession) DelegateCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.DelegateCollateral(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) DelegateCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.DelegateCollateral(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismTransactor) Deposit(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "deposit", accountId, collateralType, tokenAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismSession) Deposit(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.Deposit(&_CoreOptimism.TransactOpts, accountId, collateralType, tokenAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) Deposit(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.Deposit(&_CoreOptimism.TransactOpts, accountId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismTransactor) DepositMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "depositMarketCollateral", marketId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismSession) DepositMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.DepositMarketCollateral(&_CoreOptimism.TransactOpts, marketId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) DepositMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.DepositMarketCollateral(&_CoreOptimism.TransactOpts, marketId, collateralType, tokenAmount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreOptimism *CoreOptimismTransactor) DepositMarketUsd(opts *bind.TransactOpts, marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "depositMarketUsd", marketId, target, amount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreOptimism *CoreOptimismSession) DepositMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.DepositMarketUsd(&_CoreOptimism.TransactOpts, marketId, target, amount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreOptimism *CoreOptimismTransactorSession) DepositMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.DepositMarketUsd(&_CoreOptimism.TransactOpts, marketId, target, amount)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_CoreOptimism *CoreOptimismTransactor) DistributeDebtToPools(opts *bind.TransactOpts, marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "distributeDebtToPools", marketId, maxIter)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_CoreOptimism *CoreOptimismSession) DistributeDebtToPools(marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.DistributeDebtToPools(&_CoreOptimism.TransactOpts, marketId, maxIter)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_CoreOptimism *CoreOptimismTransactorSession) DistributeDebtToPools(marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.DistributeDebtToPools(&_CoreOptimism.TransactOpts, marketId, maxIter)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_CoreOptimism *CoreOptimismTransactor) DistributeRewards(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "distributeRewards", poolId, collateralType, amount, start, duration)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_CoreOptimism *CoreOptimismSession) DistributeRewards(poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _CoreOptimism.Contract.DistributeRewards(&_CoreOptimism.TransactOpts, poolId, collateralType, amount, start, duration)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) DistributeRewards(poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _CoreOptimism.Contract.DistributeRewards(&_CoreOptimism.TransactOpts, poolId, collateralType, amount, start, duration)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_CoreOptimism *CoreOptimismTransactor) GetMarketDebtPerShare(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "getMarketDebtPerShare", marketId)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_CoreOptimism *CoreOptimismSession) GetMarketDebtPerShare(marketId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetMarketDebtPerShare(&_CoreOptimism.TransactOpts, marketId)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_CoreOptimism *CoreOptimismTransactorSession) GetMarketDebtPerShare(marketId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetMarketDebtPerShare(&_CoreOptimism.TransactOpts, marketId)
}

// GetMarketPoolDebtDistribution is a paid mutator transaction binding the contract method 0x25eeea4b.
//
// Solidity: function getMarketPoolDebtDistribution(uint128 marketId, uint128 poolId) returns(uint256 sharesD18, uint128 totalSharesD18, int128 valuePerShareD27)
func (_CoreOptimism *CoreOptimismTransactor) GetMarketPoolDebtDistribution(opts *bind.TransactOpts, marketId *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "getMarketPoolDebtDistribution", marketId, poolId)
}

// GetMarketPoolDebtDistribution is a paid mutator transaction binding the contract method 0x25eeea4b.
//
// Solidity: function getMarketPoolDebtDistribution(uint128 marketId, uint128 poolId) returns(uint256 sharesD18, uint128 totalSharesD18, int128 valuePerShareD27)
func (_CoreOptimism *CoreOptimismSession) GetMarketPoolDebtDistribution(marketId *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetMarketPoolDebtDistribution(&_CoreOptimism.TransactOpts, marketId, poolId)
}

// GetMarketPoolDebtDistribution is a paid mutator transaction binding the contract method 0x25eeea4b.
//
// Solidity: function getMarketPoolDebtDistribution(uint128 marketId, uint128 poolId) returns(uint256 sharesD18, uint128 totalSharesD18, int128 valuePerShareD27)
func (_CoreOptimism *CoreOptimismTransactorSession) GetMarketPoolDebtDistribution(marketId *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetMarketPoolDebtDistribution(&_CoreOptimism.TransactOpts, marketId, poolId)
}

// GetMarketPools is a paid mutator transaction binding the contract method 0xbe0b8e6f.
//
// Solidity: function getMarketPools(uint128 marketId) returns(uint128[] inRangePoolIds, uint128[] outRangePoolIds)
func (_CoreOptimism *CoreOptimismTransactor) GetMarketPools(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "getMarketPools", marketId)
}

// GetMarketPools is a paid mutator transaction binding the contract method 0xbe0b8e6f.
//
// Solidity: function getMarketPools(uint128 marketId) returns(uint128[] inRangePoolIds, uint128[] outRangePoolIds)
func (_CoreOptimism *CoreOptimismSession) GetMarketPools(marketId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetMarketPools(&_CoreOptimism.TransactOpts, marketId)
}

// GetMarketPools is a paid mutator transaction binding the contract method 0xbe0b8e6f.
//
// Solidity: function getMarketPools(uint128 marketId) returns(uint128[] inRangePoolIds, uint128[] outRangePoolIds)
func (_CoreOptimism *CoreOptimismTransactorSession) GetMarketPools(marketId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetMarketPools(&_CoreOptimism.TransactOpts, marketId)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_CoreOptimism *CoreOptimismTransactor) GetPosition(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "getPosition", accountId, poolId, collateralType)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_CoreOptimism *CoreOptimismSession) GetPosition(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetPosition(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_CoreOptimism *CoreOptimismTransactorSession) GetPosition(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetPosition(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_CoreOptimism *CoreOptimismTransactor) GetPositionCollateralRatio(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "getPositionCollateralRatio", accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetPositionCollateralRatio(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetPositionCollateralRatio(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_CoreOptimism *CoreOptimismTransactorSession) GetPositionCollateralRatio(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetPositionCollateralRatio(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_CoreOptimism *CoreOptimismTransactor) GetPositionDebt(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "getPositionDebt", accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_CoreOptimism *CoreOptimismSession) GetPositionDebt(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetPositionDebt(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_CoreOptimism *CoreOptimismTransactorSession) GetPositionDebt(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetPositionDebt(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_CoreOptimism *CoreOptimismTransactor) GetVaultCollateralRatio(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "getVaultCollateralRatio", poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_CoreOptimism *CoreOptimismSession) GetVaultCollateralRatio(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetVaultCollateralRatio(&_CoreOptimism.TransactOpts, poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_CoreOptimism *CoreOptimismTransactorSession) GetVaultCollateralRatio(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetVaultCollateralRatio(&_CoreOptimism.TransactOpts, poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_CoreOptimism *CoreOptimismTransactor) GetVaultDebt(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "getVaultDebt", poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_CoreOptimism *CoreOptimismSession) GetVaultDebt(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetVaultDebt(&_CoreOptimism.TransactOpts, poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_CoreOptimism *CoreOptimismTransactorSession) GetVaultDebt(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GetVaultDebt(&_CoreOptimism.TransactOpts, poolId, collateralType)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreOptimism *CoreOptimismTransactor) GrantPermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "grantPermission", accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreOptimism *CoreOptimismSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GrantPermission(&_CoreOptimism.TransactOpts, accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.GrantPermission(&_CoreOptimism.TransactOpts, accountId, permission, user)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_CoreOptimism *CoreOptimismTransactor) InitOrUpgradeNft(opts *bind.TransactOpts, id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "initOrUpgradeNft", id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_CoreOptimism *CoreOptimismSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.InitOrUpgradeNft(&_CoreOptimism.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.InitOrUpgradeNft(&_CoreOptimism.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_CoreOptimism *CoreOptimismTransactor) InitOrUpgradeToken(opts *bind.TransactOpts, id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "initOrUpgradeToken", id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_CoreOptimism *CoreOptimismSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.InitOrUpgradeToken(&_CoreOptimism.TransactOpts, id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.InitOrUpgradeToken(&_CoreOptimism.TransactOpts, id, name, symbol, decimals, impl)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_CoreOptimism *CoreOptimismTransactor) IsPositionLiquidatable(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "isPositionLiquidatable", accountId, poolId, collateralType)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_CoreOptimism *CoreOptimismSession) IsPositionLiquidatable(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.IsPositionLiquidatable(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_CoreOptimism *CoreOptimismTransactorSession) IsPositionLiquidatable(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.IsPositionLiquidatable(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_CoreOptimism *CoreOptimismTransactor) IsVaultLiquidatable(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "isVaultLiquidatable", poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_CoreOptimism *CoreOptimismSession) IsVaultLiquidatable(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.IsVaultLiquidatable(&_CoreOptimism.TransactOpts, poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_CoreOptimism *CoreOptimismTransactorSession) IsVaultLiquidatable(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.IsVaultLiquidatable(&_CoreOptimism.TransactOpts, poolId, collateralType)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_CoreOptimism *CoreOptimismTransactor) Liquidate(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "liquidate", accountId, poolId, collateralType, liquidateAsAccountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_CoreOptimism *CoreOptimismSession) Liquidate(accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.Liquidate(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType, liquidateAsAccountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_CoreOptimism *CoreOptimismTransactorSession) Liquidate(accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.Liquidate(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType, liquidateAsAccountId)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_CoreOptimism *CoreOptimismTransactor) LiquidateVault(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "liquidateVault", poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_CoreOptimism *CoreOptimismSession) LiquidateVault(poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.LiquidateVault(&_CoreOptimism.TransactOpts, poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_CoreOptimism *CoreOptimismTransactorSession) LiquidateVault(poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.LiquidateVault(&_CoreOptimism.TransactOpts, poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreOptimism *CoreOptimismTransactor) MintUsd(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "mintUsd", accountId, poolId, collateralType, amount)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreOptimism *CoreOptimismSession) MintUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.MintUsd(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType, amount)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) MintUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.MintUsd(&_CoreOptimism.TransactOpts, accountId, poolId, collateralType, amount)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_CoreOptimism *CoreOptimismTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_CoreOptimism *CoreOptimismSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _CoreOptimism.Contract.Multicall(&_CoreOptimism.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_CoreOptimism *CoreOptimismTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _CoreOptimism.Contract.Multicall(&_CoreOptimism.TransactOpts, data)
}

// MulticallThrough is a paid mutator transaction binding the contract method 0xc0694da0.
//
// Solidity: function multicallThrough(address[] to, bytes[] data, uint256[] values) payable returns(bytes[] results)
func (_CoreOptimism *CoreOptimismTransactor) MulticallThrough(opts *bind.TransactOpts, to []common.Address, data [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "multicallThrough", to, data, values)
}

// MulticallThrough is a paid mutator transaction binding the contract method 0xc0694da0.
//
// Solidity: function multicallThrough(address[] to, bytes[] data, uint256[] values) payable returns(bytes[] results)
func (_CoreOptimism *CoreOptimismSession) MulticallThrough(to []common.Address, data [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.MulticallThrough(&_CoreOptimism.TransactOpts, to, data, values)
}

// MulticallThrough is a paid mutator transaction binding the contract method 0xc0694da0.
//
// Solidity: function multicallThrough(address[] to, bytes[] data, uint256[] values) payable returns(bytes[] results)
func (_CoreOptimism *CoreOptimismTransactorSession) MulticallThrough(to []common.Address, data [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.MulticallThrough(&_CoreOptimism.TransactOpts, to, data, values)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_CoreOptimism *CoreOptimismTransactor) NominateNewOwner(opts *bind.TransactOpts, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "nominateNewOwner", newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_CoreOptimism *CoreOptimismSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.NominateNewOwner(&_CoreOptimism.TransactOpts, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.NominateNewOwner(&_CoreOptimism.TransactOpts, newNominatedOwner)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactor) NominatePoolOwner(opts *bind.TransactOpts, nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "nominatePoolOwner", nominatedOwner, poolId)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismSession) NominatePoolOwner(nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.NominatePoolOwner(&_CoreOptimism.TransactOpts, nominatedOwner, poolId)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) NominatePoolOwner(nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.NominatePoolOwner(&_CoreOptimism.TransactOpts, nominatedOwner, poolId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_CoreOptimism *CoreOptimismTransactor) NotifyAccountTransfer(opts *bind.TransactOpts, to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "notifyAccountTransfer", to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_CoreOptimism *CoreOptimismSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.NotifyAccountTransfer(&_CoreOptimism.TransactOpts, to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.NotifyAccountTransfer(&_CoreOptimism.TransactOpts, to, accountId)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_CoreOptimism *CoreOptimismTransactor) RebalancePool(opts *bind.TransactOpts, poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "rebalancePool", poolId, optionalCollateralType)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_CoreOptimism *CoreOptimismSession) RebalancePool(poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RebalancePool(&_CoreOptimism.TransactOpts, poolId, optionalCollateralType)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RebalancePool(poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RebalancePool(&_CoreOptimism.TransactOpts, poolId, optionalCollateralType)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_CoreOptimism *CoreOptimismTransactor) RegisterMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "registerMarket", market)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_CoreOptimism *CoreOptimismSession) RegisterMarket(market common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RegisterMarket(&_CoreOptimism.TransactOpts, market)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_CoreOptimism *CoreOptimismTransactorSession) RegisterMarket(market common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RegisterMarket(&_CoreOptimism.TransactOpts, market)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreOptimism *CoreOptimismTransactor) RegisterRewardsDistributor(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "registerRewardsDistributor", poolId, collateralType, distributor)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreOptimism *CoreOptimismSession) RegisterRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RegisterRewardsDistributor(&_CoreOptimism.TransactOpts, poolId, collateralType, distributor)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RegisterRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RegisterRewardsDistributor(&_CoreOptimism.TransactOpts, poolId, collateralType, distributor)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_CoreOptimism *CoreOptimismTransactor) RegisterUnmanagedSystem(opts *bind.TransactOpts, id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "registerUnmanagedSystem", id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_CoreOptimism *CoreOptimismSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RegisterUnmanagedSystem(&_CoreOptimism.TransactOpts, id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RegisterUnmanagedSystem(&_CoreOptimism.TransactOpts, id, endpoint)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactor) RemoveApprovedPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "removeApprovedPool", poolId)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismSession) RemoveApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RemoveApprovedPool(&_CoreOptimism.TransactOpts, poolId)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RemoveApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RemoveApprovedPool(&_CoreOptimism.TransactOpts, poolId)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreOptimism *CoreOptimismTransactor) RemoveFromFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "removeFromFeatureFlagAllowlist", feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreOptimism *CoreOptimismSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RemoveFromFeatureFlagAllowlist(&_CoreOptimism.TransactOpts, feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RemoveFromFeatureFlagAllowlist(&_CoreOptimism.TransactOpts, feature, account)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreOptimism *CoreOptimismTransactor) RemoveRewardsDistributor(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "removeRewardsDistributor", poolId, collateralType, distributor)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreOptimism *CoreOptimismSession) RemoveRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RemoveRewardsDistributor(&_CoreOptimism.TransactOpts, poolId, collateralType, distributor)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RemoveRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RemoveRewardsDistributor(&_CoreOptimism.TransactOpts, poolId, collateralType, distributor)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_CoreOptimism *CoreOptimismTransactor) RenounceNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "renounceNomination")
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_CoreOptimism *CoreOptimismSession) RenounceNomination() (*types.Transaction, error) {
	return _CoreOptimism.Contract.RenounceNomination(&_CoreOptimism.TransactOpts)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RenounceNomination() (*types.Transaction, error) {
	return _CoreOptimism.Contract.RenounceNomination(&_CoreOptimism.TransactOpts)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_CoreOptimism *CoreOptimismTransactor) RenouncePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "renouncePermission", accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_CoreOptimism *CoreOptimismSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RenouncePermission(&_CoreOptimism.TransactOpts, accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RenouncePermission(&_CoreOptimism.TransactOpts, accountId, permission)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactor) RenouncePoolNomination(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "renouncePoolNomination", poolId)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismSession) RenouncePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RenouncePoolNomination(&_CoreOptimism.TransactOpts, poolId)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RenouncePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RenouncePoolNomination(&_CoreOptimism.TransactOpts, poolId)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreOptimism *CoreOptimismTransactor) RevokePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "revokePermission", accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreOptimism *CoreOptimismSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RevokePermission(&_CoreOptimism.TransactOpts, accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RevokePermission(&_CoreOptimism.TransactOpts, accountId, permission, user)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactor) RevokePoolNomination(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "revokePoolNomination", poolId)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismSession) RevokePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RevokePoolNomination(&_CoreOptimism.TransactOpts, poolId)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) RevokePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.RevokePoolNomination(&_CoreOptimism.TransactOpts, poolId)
}

// SetAllowlistedMulticallTarget is a paid mutator transaction binding the contract method 0x0a6c7553.
//
// Solidity: function setAllowlistedMulticallTarget(address target, bool allowlisted) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetAllowlistedMulticallTarget(opts *bind.TransactOpts, target common.Address, allowlisted bool) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setAllowlistedMulticallTarget", target, allowlisted)
}

// SetAllowlistedMulticallTarget is a paid mutator transaction binding the contract method 0x0a6c7553.
//
// Solidity: function setAllowlistedMulticallTarget(address target, bool allowlisted) returns()
func (_CoreOptimism *CoreOptimismSession) SetAllowlistedMulticallTarget(target common.Address, allowlisted bool) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetAllowlistedMulticallTarget(&_CoreOptimism.TransactOpts, target, allowlisted)
}

// SetAllowlistedMulticallTarget is a paid mutator transaction binding the contract method 0x0a6c7553.
//
// Solidity: function setAllowlistedMulticallTarget(address target, bool allowlisted) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetAllowlistedMulticallTarget(target common.Address, allowlisted bool) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetAllowlistedMulticallTarget(&_CoreOptimism.TransactOpts, target, allowlisted)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetConfig(opts *bind.TransactOpts, k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setConfig", k, v)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_CoreOptimism *CoreOptimismSession) SetConfig(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetConfig(&_CoreOptimism.TransactOpts, k, v)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetConfig(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetConfig(&_CoreOptimism.TransactOpts, k, v)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetDeniers(opts *bind.TransactOpts, feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setDeniers", feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_CoreOptimism *CoreOptimismSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetDeniers(&_CoreOptimism.TransactOpts, feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetDeniers(&_CoreOptimism.TransactOpts, feature, deniers)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetFeatureFlagAllowAll(opts *bind.TransactOpts, feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setFeatureFlagAllowAll", feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_CoreOptimism *CoreOptimismSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetFeatureFlagAllowAll(&_CoreOptimism.TransactOpts, feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetFeatureFlagAllowAll(&_CoreOptimism.TransactOpts, feature, allowAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetFeatureFlagDenyAll(opts *bind.TransactOpts, feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setFeatureFlagDenyAll", feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_CoreOptimism *CoreOptimismSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetFeatureFlagDenyAll(&_CoreOptimism.TransactOpts, feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetFeatureFlagDenyAll(&_CoreOptimism.TransactOpts, feature, denyAll)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetMarketMinDelegateTime(opts *bind.TransactOpts, marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setMarketMinDelegateTime", marketId, minDelegateTime)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_CoreOptimism *CoreOptimismSession) SetMarketMinDelegateTime(marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetMarketMinDelegateTime(&_CoreOptimism.TransactOpts, marketId, minDelegateTime)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetMarketMinDelegateTime(marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetMarketMinDelegateTime(&_CoreOptimism.TransactOpts, marketId, minDelegateTime)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetMinLiquidityRatio(opts *bind.TransactOpts, marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setMinLiquidityRatio", marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_CoreOptimism *CoreOptimismSession) SetMinLiquidityRatio(marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetMinLiquidityRatio(&_CoreOptimism.TransactOpts, marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetMinLiquidityRatio(marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetMinLiquidityRatio(&_CoreOptimism.TransactOpts, marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetMinLiquidityRatio0(opts *bind.TransactOpts, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setMinLiquidityRatio0", minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_CoreOptimism *CoreOptimismSession) SetMinLiquidityRatio0(minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetMinLiquidityRatio0(&_CoreOptimism.TransactOpts, minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetMinLiquidityRatio0(minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetMinLiquidityRatio0(&_CoreOptimism.TransactOpts, minLiquidityRatio)
}

// SetPoolCollateralConfiguration is a paid mutator transaction binding the contract method 0x5a4aabb1.
//
// Solidity: function setPoolCollateralConfiguration(uint128 poolId, address collateralType, (uint256,uint256) newConfig) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetPoolCollateralConfiguration(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, newConfig PoolCollateralConfigurationData) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setPoolCollateralConfiguration", poolId, collateralType, newConfig)
}

// SetPoolCollateralConfiguration is a paid mutator transaction binding the contract method 0x5a4aabb1.
//
// Solidity: function setPoolCollateralConfiguration(uint128 poolId, address collateralType, (uint256,uint256) newConfig) returns()
func (_CoreOptimism *CoreOptimismSession) SetPoolCollateralConfiguration(poolId *big.Int, collateralType common.Address, newConfig PoolCollateralConfigurationData) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetPoolCollateralConfiguration(&_CoreOptimism.TransactOpts, poolId, collateralType, newConfig)
}

// SetPoolCollateralConfiguration is a paid mutator transaction binding the contract method 0x5a4aabb1.
//
// Solidity: function setPoolCollateralConfiguration(uint128 poolId, address collateralType, (uint256,uint256) newConfig) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetPoolCollateralConfiguration(poolId *big.Int, collateralType common.Address, newConfig PoolCollateralConfigurationData) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetPoolCollateralConfiguration(&_CoreOptimism.TransactOpts, poolId, collateralType, newConfig)
}

// SetPoolCollateralDisabledByDefault is a paid mutator transaction binding the contract method 0x4c6568b1.
//
// Solidity: function setPoolCollateralDisabledByDefault(uint128 poolId, bool disabled) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetPoolCollateralDisabledByDefault(opts *bind.TransactOpts, poolId *big.Int, disabled bool) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setPoolCollateralDisabledByDefault", poolId, disabled)
}

// SetPoolCollateralDisabledByDefault is a paid mutator transaction binding the contract method 0x4c6568b1.
//
// Solidity: function setPoolCollateralDisabledByDefault(uint128 poolId, bool disabled) returns()
func (_CoreOptimism *CoreOptimismSession) SetPoolCollateralDisabledByDefault(poolId *big.Int, disabled bool) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetPoolCollateralDisabledByDefault(&_CoreOptimism.TransactOpts, poolId, disabled)
}

// SetPoolCollateralDisabledByDefault is a paid mutator transaction binding the contract method 0x4c6568b1.
//
// Solidity: function setPoolCollateralDisabledByDefault(uint128 poolId, bool disabled) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetPoolCollateralDisabledByDefault(poolId *big.Int, disabled bool) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetPoolCollateralDisabledByDefault(&_CoreOptimism.TransactOpts, poolId, disabled)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetPoolConfiguration(opts *bind.TransactOpts, poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setPoolConfiguration", poolId, newMarketConfigurations)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_CoreOptimism *CoreOptimismSession) SetPoolConfiguration(poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetPoolConfiguration(&_CoreOptimism.TransactOpts, poolId, newMarketConfigurations)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetPoolConfiguration(poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetPoolConfiguration(&_CoreOptimism.TransactOpts, poolId, newMarketConfigurations)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetPoolName(opts *bind.TransactOpts, poolId *big.Int, name string) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setPoolName", poolId, name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_CoreOptimism *CoreOptimismSession) SetPoolName(poolId *big.Int, name string) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetPoolName(&_CoreOptimism.TransactOpts, poolId, name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetPoolName(poolId *big.Int, name string) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetPoolName(&_CoreOptimism.TransactOpts, poolId, name)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactor) SetPreferredPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setPreferredPool", poolId)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismSession) SetPreferredPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetPreferredPool(&_CoreOptimism.TransactOpts, poolId)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SetPreferredPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetPreferredPool(&_CoreOptimism.TransactOpts, poolId)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_CoreOptimism *CoreOptimismTransactor) SetSupportedCrossChainNetworks(opts *bind.TransactOpts, supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "setSupportedCrossChainNetworks", supportedNetworks, ccipSelectors)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_CoreOptimism *CoreOptimismSession) SetSupportedCrossChainNetworks(supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetSupportedCrossChainNetworks(&_CoreOptimism.TransactOpts, supportedNetworks, ccipSelectors)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_CoreOptimism *CoreOptimismTransactorSession) SetSupportedCrossChainNetworks(supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SetSupportedCrossChainNetworks(&_CoreOptimism.TransactOpts, supportedNetworks, ccipSelectors)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_CoreOptimism *CoreOptimismTransactor) SimulateUpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "simulateUpgradeTo", newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_CoreOptimism *CoreOptimismSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SimulateUpgradeTo(&_CoreOptimism.TransactOpts, newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.SimulateUpgradeTo(&_CoreOptimism.TransactOpts, newImplementation)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_CoreOptimism *CoreOptimismTransactor) TransferCrossChain(opts *bind.TransactOpts, destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "transferCrossChain", destChainId, amount)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_CoreOptimism *CoreOptimismSession) TransferCrossChain(destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.TransferCrossChain(&_CoreOptimism.TransactOpts, destChainId, amount)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_CoreOptimism *CoreOptimismTransactorSession) TransferCrossChain(destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.TransferCrossChain(&_CoreOptimism.TransactOpts, destChainId, amount)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_CoreOptimism *CoreOptimismTransactor) UpdateRewards(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "updateRewards", poolId, collateralType, accountId)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_CoreOptimism *CoreOptimismSession) UpdateRewards(poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.UpdateRewards(&_CoreOptimism.TransactOpts, poolId, collateralType, accountId)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_CoreOptimism *CoreOptimismTransactorSession) UpdateRewards(poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.UpdateRewards(&_CoreOptimism.TransactOpts, poolId, collateralType, accountId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_CoreOptimism *CoreOptimismTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_CoreOptimism *CoreOptimismSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.UpgradeTo(&_CoreOptimism.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _CoreOptimism.Contract.UpgradeTo(&_CoreOptimism.TransactOpts, newImplementation)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismTransactor) Withdraw(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "withdraw", accountId, collateralType, tokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismSession) Withdraw(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.Withdraw(&_CoreOptimism.TransactOpts, accountId, collateralType, tokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) Withdraw(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.Withdraw(&_CoreOptimism.TransactOpts, accountId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismTransactor) WithdrawMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "withdrawMarketCollateral", marketId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismSession) WithdrawMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.WithdrawMarketCollateral(&_CoreOptimism.TransactOpts, marketId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreOptimism *CoreOptimismTransactorSession) WithdrawMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.WithdrawMarketCollateral(&_CoreOptimism.TransactOpts, marketId, collateralType, tokenAmount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreOptimism *CoreOptimismTransactor) WithdrawMarketUsd(opts *bind.TransactOpts, marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.contract.Transact(opts, "withdrawMarketUsd", marketId, target, amount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreOptimism *CoreOptimismSession) WithdrawMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.WithdrawMarketUsd(&_CoreOptimism.TransactOpts, marketId, target, amount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreOptimism *CoreOptimismTransactorSession) WithdrawMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreOptimism.Contract.WithdrawMarketUsd(&_CoreOptimism.TransactOpts, marketId, target, amount)
}

// CoreOptimismAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the CoreOptimism contract.
type CoreOptimismAccountCreatedIterator struct {
	Event *CoreOptimismAccountCreated // Event containing the contract specifics and raw log

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
func (it *CoreOptimismAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismAccountCreated)
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
		it.Event = new(CoreOptimismAccountCreated)
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
func (it *CoreOptimismAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismAccountCreated represents a AccountCreated event raised by the CoreOptimism contract.
type CoreOptimismAccountCreated struct {
	AccountId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) FilterAccountCreated(opts *bind.FilterOpts, accountId []*big.Int, owner []common.Address) (*CoreOptimismAccountCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismAccountCreatedIterator{contract: _CoreOptimism.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *CoreOptimismAccountCreated, accountId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismAccountCreated)
				if err := _CoreOptimism.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParseAccountCreated(log types.Log) (*CoreOptimismAccountCreated, error) {
	event := new(CoreOptimismAccountCreated)
	if err := _CoreOptimism.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismAssociatedSystemSetIterator is returned from FilterAssociatedSystemSet and is used to iterate over the raw logs and unpacked data for AssociatedSystemSet events raised by the CoreOptimism contract.
type CoreOptimismAssociatedSystemSetIterator struct {
	Event *CoreOptimismAssociatedSystemSet // Event containing the contract specifics and raw log

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
func (it *CoreOptimismAssociatedSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismAssociatedSystemSet)
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
		it.Event = new(CoreOptimismAssociatedSystemSet)
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
func (it *CoreOptimismAssociatedSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismAssociatedSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismAssociatedSystemSet represents a AssociatedSystemSet event raised by the CoreOptimism contract.
type CoreOptimismAssociatedSystemSet struct {
	Kind  [32]byte
	Id    [32]byte
	Proxy common.Address
	Impl  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssociatedSystemSet is a free log retrieval operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_CoreOptimism *CoreOptimismFilterer) FilterAssociatedSystemSet(opts *bind.FilterOpts, kind [][32]byte, id [][32]byte) (*CoreOptimismAssociatedSystemSetIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismAssociatedSystemSetIterator{contract: _CoreOptimism.contract, event: "AssociatedSystemSet", logs: logs, sub: sub}, nil
}

// WatchAssociatedSystemSet is a free log subscription operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_CoreOptimism *CoreOptimismFilterer) WatchAssociatedSystemSet(opts *bind.WatchOpts, sink chan<- *CoreOptimismAssociatedSystemSet, kind [][32]byte, id [][32]byte) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismAssociatedSystemSet)
				if err := _CoreOptimism.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParseAssociatedSystemSet(log types.Log) (*CoreOptimismAssociatedSystemSet, error) {
	event := new(CoreOptimismAssociatedSystemSet)
	if err := _CoreOptimism.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismCollateralConfiguredIterator is returned from FilterCollateralConfigured and is used to iterate over the raw logs and unpacked data for CollateralConfigured events raised by the CoreOptimism contract.
type CoreOptimismCollateralConfiguredIterator struct {
	Event *CoreOptimismCollateralConfigured // Event containing the contract specifics and raw log

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
func (it *CoreOptimismCollateralConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismCollateralConfigured)
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
		it.Event = new(CoreOptimismCollateralConfigured)
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
func (it *CoreOptimismCollateralConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismCollateralConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismCollateralConfigured represents a CollateralConfigured event raised by the CoreOptimism contract.
type CoreOptimismCollateralConfigured struct {
	CollateralType common.Address
	Config         CollateralConfigurationData
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCollateralConfigured is a free log retrieval operation binding the contract event 0xefc23317f58afd6b22480bd22174cc7da0913bce25c03d9859216dacddebe6fe.
//
// Solidity: event CollateralConfigured(address indexed collateralType, (bool,uint256,uint256,uint256,bytes32,address,uint256) config)
func (_CoreOptimism *CoreOptimismFilterer) FilterCollateralConfigured(opts *bind.FilterOpts, collateralType []common.Address) (*CoreOptimismCollateralConfiguredIterator, error) {

	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "CollateralConfigured", collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismCollateralConfiguredIterator{contract: _CoreOptimism.contract, event: "CollateralConfigured", logs: logs, sub: sub}, nil
}

// WatchCollateralConfigured is a free log subscription operation binding the contract event 0xefc23317f58afd6b22480bd22174cc7da0913bce25c03d9859216dacddebe6fe.
//
// Solidity: event CollateralConfigured(address indexed collateralType, (bool,uint256,uint256,uint256,bytes32,address,uint256) config)
func (_CoreOptimism *CoreOptimismFilterer) WatchCollateralConfigured(opts *bind.WatchOpts, sink chan<- *CoreOptimismCollateralConfigured, collateralType []common.Address) (event.Subscription, error) {

	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "CollateralConfigured", collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismCollateralConfigured)
				if err := _CoreOptimism.contract.UnpackLog(event, "CollateralConfigured", log); err != nil {
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

// ParseCollateralConfigured is a log parse operation binding the contract event 0xefc23317f58afd6b22480bd22174cc7da0913bce25c03d9859216dacddebe6fe.
//
// Solidity: event CollateralConfigured(address indexed collateralType, (bool,uint256,uint256,uint256,bytes32,address,uint256) config)
func (_CoreOptimism *CoreOptimismFilterer) ParseCollateralConfigured(log types.Log) (*CoreOptimismCollateralConfigured, error) {
	event := new(CoreOptimismCollateralConfigured)
	if err := _CoreOptimism.contract.UnpackLog(event, "CollateralConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismCollateralLockCreatedIterator is returned from FilterCollateralLockCreated and is used to iterate over the raw logs and unpacked data for CollateralLockCreated events raised by the CoreOptimism contract.
type CoreOptimismCollateralLockCreatedIterator struct {
	Event *CoreOptimismCollateralLockCreated // Event containing the contract specifics and raw log

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
func (it *CoreOptimismCollateralLockCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismCollateralLockCreated)
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
		it.Event = new(CoreOptimismCollateralLockCreated)
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
func (it *CoreOptimismCollateralLockCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismCollateralLockCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismCollateralLockCreated represents a CollateralLockCreated event raised by the CoreOptimism contract.
type CoreOptimismCollateralLockCreated struct {
	AccountId       *big.Int
	CollateralType  common.Address
	TokenAmount     *big.Int
	ExpireTimestamp uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollateralLockCreated is a free log retrieval operation binding the contract event 0x8a78446a234220d3ee3f46aa6ea1ea5bc438bd153398ebcbd171843744b452a6.
//
// Solidity: event CollateralLockCreated(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_CoreOptimism *CoreOptimismFilterer) FilterCollateralLockCreated(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address) (*CoreOptimismCollateralLockCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "CollateralLockCreated", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismCollateralLockCreatedIterator{contract: _CoreOptimism.contract, event: "CollateralLockCreated", logs: logs, sub: sub}, nil
}

// WatchCollateralLockCreated is a free log subscription operation binding the contract event 0x8a78446a234220d3ee3f46aa6ea1ea5bc438bd153398ebcbd171843744b452a6.
//
// Solidity: event CollateralLockCreated(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_CoreOptimism *CoreOptimismFilterer) WatchCollateralLockCreated(opts *bind.WatchOpts, sink chan<- *CoreOptimismCollateralLockCreated, accountId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "CollateralLockCreated", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismCollateralLockCreated)
				if err := _CoreOptimism.contract.UnpackLog(event, "CollateralLockCreated", log); err != nil {
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

// ParseCollateralLockCreated is a log parse operation binding the contract event 0x8a78446a234220d3ee3f46aa6ea1ea5bc438bd153398ebcbd171843744b452a6.
//
// Solidity: event CollateralLockCreated(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_CoreOptimism *CoreOptimismFilterer) ParseCollateralLockCreated(log types.Log) (*CoreOptimismCollateralLockCreated, error) {
	event := new(CoreOptimismCollateralLockCreated)
	if err := _CoreOptimism.contract.UnpackLog(event, "CollateralLockCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismCollateralLockExpiredIterator is returned from FilterCollateralLockExpired and is used to iterate over the raw logs and unpacked data for CollateralLockExpired events raised by the CoreOptimism contract.
type CoreOptimismCollateralLockExpiredIterator struct {
	Event *CoreOptimismCollateralLockExpired // Event containing the contract specifics and raw log

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
func (it *CoreOptimismCollateralLockExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismCollateralLockExpired)
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
		it.Event = new(CoreOptimismCollateralLockExpired)
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
func (it *CoreOptimismCollateralLockExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismCollateralLockExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismCollateralLockExpired represents a CollateralLockExpired event raised by the CoreOptimism contract.
type CoreOptimismCollateralLockExpired struct {
	AccountId       *big.Int
	CollateralType  common.Address
	TokenAmount     *big.Int
	ExpireTimestamp uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollateralLockExpired is a free log retrieval operation binding the contract event 0xc5e2c7afc4e54d998977e11260a0bfc0ad5678a3a8b6628162f9d4e642d7f160.
//
// Solidity: event CollateralLockExpired(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_CoreOptimism *CoreOptimismFilterer) FilterCollateralLockExpired(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address) (*CoreOptimismCollateralLockExpiredIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "CollateralLockExpired", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismCollateralLockExpiredIterator{contract: _CoreOptimism.contract, event: "CollateralLockExpired", logs: logs, sub: sub}, nil
}

// WatchCollateralLockExpired is a free log subscription operation binding the contract event 0xc5e2c7afc4e54d998977e11260a0bfc0ad5678a3a8b6628162f9d4e642d7f160.
//
// Solidity: event CollateralLockExpired(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_CoreOptimism *CoreOptimismFilterer) WatchCollateralLockExpired(opts *bind.WatchOpts, sink chan<- *CoreOptimismCollateralLockExpired, accountId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "CollateralLockExpired", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismCollateralLockExpired)
				if err := _CoreOptimism.contract.UnpackLog(event, "CollateralLockExpired", log); err != nil {
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

// ParseCollateralLockExpired is a log parse operation binding the contract event 0xc5e2c7afc4e54d998977e11260a0bfc0ad5678a3a8b6628162f9d4e642d7f160.
//
// Solidity: event CollateralLockExpired(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_CoreOptimism *CoreOptimismFilterer) ParseCollateralLockExpired(log types.Log) (*CoreOptimismCollateralLockExpired, error) {
	event := new(CoreOptimismCollateralLockExpired)
	if err := _CoreOptimism.contract.UnpackLog(event, "CollateralLockExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismDebtAssociatedIterator is returned from FilterDebtAssociated and is used to iterate over the raw logs and unpacked data for DebtAssociated events raised by the CoreOptimism contract.
type CoreOptimismDebtAssociatedIterator struct {
	Event *CoreOptimismDebtAssociated // Event containing the contract specifics and raw log

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
func (it *CoreOptimismDebtAssociatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismDebtAssociated)
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
		it.Event = new(CoreOptimismDebtAssociated)
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
func (it *CoreOptimismDebtAssociatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismDebtAssociatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismDebtAssociated represents a DebtAssociated event raised by the CoreOptimism contract.
type CoreOptimismDebtAssociated struct {
	MarketId       *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	AccountId      *big.Int
	Amount         *big.Int
	UpdatedDebt    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDebtAssociated is a free log retrieval operation binding the contract event 0xb03bc7530b5650601d2c8bc86c81a45b6b50b0099defd3f17a6bdb48660180ad.
//
// Solidity: event DebtAssociated(uint128 indexed marketId, uint128 indexed poolId, address indexed collateralType, uint128 accountId, uint256 amount, int256 updatedDebt)
func (_CoreOptimism *CoreOptimismFilterer) FilterDebtAssociated(opts *bind.FilterOpts, marketId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*CoreOptimismDebtAssociatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "DebtAssociated", marketIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismDebtAssociatedIterator{contract: _CoreOptimism.contract, event: "DebtAssociated", logs: logs, sub: sub}, nil
}

// WatchDebtAssociated is a free log subscription operation binding the contract event 0xb03bc7530b5650601d2c8bc86c81a45b6b50b0099defd3f17a6bdb48660180ad.
//
// Solidity: event DebtAssociated(uint128 indexed marketId, uint128 indexed poolId, address indexed collateralType, uint128 accountId, uint256 amount, int256 updatedDebt)
func (_CoreOptimism *CoreOptimismFilterer) WatchDebtAssociated(opts *bind.WatchOpts, sink chan<- *CoreOptimismDebtAssociated, marketId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "DebtAssociated", marketIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismDebtAssociated)
				if err := _CoreOptimism.contract.UnpackLog(event, "DebtAssociated", log); err != nil {
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

// ParseDebtAssociated is a log parse operation binding the contract event 0xb03bc7530b5650601d2c8bc86c81a45b6b50b0099defd3f17a6bdb48660180ad.
//
// Solidity: event DebtAssociated(uint128 indexed marketId, uint128 indexed poolId, address indexed collateralType, uint128 accountId, uint256 amount, int256 updatedDebt)
func (_CoreOptimism *CoreOptimismFilterer) ParseDebtAssociated(log types.Log) (*CoreOptimismDebtAssociated, error) {
	event := new(CoreOptimismDebtAssociated)
	if err := _CoreOptimism.contract.UnpackLog(event, "DebtAssociated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismDelegationUpdatedIterator is returned from FilterDelegationUpdated and is used to iterate over the raw logs and unpacked data for DelegationUpdated events raised by the CoreOptimism contract.
type CoreOptimismDelegationUpdatedIterator struct {
	Event *CoreOptimismDelegationUpdated // Event containing the contract specifics and raw log

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
func (it *CoreOptimismDelegationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismDelegationUpdated)
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
		it.Event = new(CoreOptimismDelegationUpdated)
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
func (it *CoreOptimismDelegationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismDelegationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismDelegationUpdated represents a DelegationUpdated event raised by the CoreOptimism contract.
type CoreOptimismDelegationUpdated struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	Amount         *big.Int
	Leverage       *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegationUpdated is a free log retrieval operation binding the contract event 0x7b12dd38f18c0ff77ae702f6da13fbbcb28f53f807ecc7d39ee8d8b1ea8295ad.
//
// Solidity: event DelegationUpdated(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, uint256 leverage, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterDelegationUpdated(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*CoreOptimismDelegationUpdatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "DelegationUpdated", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismDelegationUpdatedIterator{contract: _CoreOptimism.contract, event: "DelegationUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegationUpdated is a free log subscription operation binding the contract event 0x7b12dd38f18c0ff77ae702f6da13fbbcb28f53f807ecc7d39ee8d8b1ea8295ad.
//
// Solidity: event DelegationUpdated(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, uint256 leverage, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchDelegationUpdated(opts *bind.WatchOpts, sink chan<- *CoreOptimismDelegationUpdated, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "DelegationUpdated", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismDelegationUpdated)
				if err := _CoreOptimism.contract.UnpackLog(event, "DelegationUpdated", log); err != nil {
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

// ParseDelegationUpdated is a log parse operation binding the contract event 0x7b12dd38f18c0ff77ae702f6da13fbbcb28f53f807ecc7d39ee8d8b1ea8295ad.
//
// Solidity: event DelegationUpdated(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, uint256 leverage, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseDelegationUpdated(log types.Log) (*CoreOptimismDelegationUpdated, error) {
	event := new(CoreOptimismDelegationUpdated)
	if err := _CoreOptimism.contract.UnpackLog(event, "DelegationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the CoreOptimism contract.
type CoreOptimismDepositedIterator struct {
	Event *CoreOptimismDeposited // Event containing the contract specifics and raw log

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
func (it *CoreOptimismDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismDeposited)
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
		it.Event = new(CoreOptimismDeposited)
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
func (it *CoreOptimismDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismDeposited represents a Deposited event raised by the CoreOptimism contract.
type CoreOptimismDeposited struct {
	AccountId      *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xd92122e67326e9313bfae33ccb1fccf5194584c6bf93a8529a6b006d8c6e24a9.
//
// Solidity: event Deposited(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterDeposited(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreOptimismDepositedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "Deposited", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismDepositedIterator{contract: _CoreOptimism.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xd92122e67326e9313bfae33ccb1fccf5194584c6bf93a8529a6b006d8c6e24a9.
//
// Solidity: event Deposited(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *CoreOptimismDeposited, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "Deposited", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismDeposited)
				if err := _CoreOptimism.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xd92122e67326e9313bfae33ccb1fccf5194584c6bf93a8529a6b006d8c6e24a9.
//
// Solidity: event Deposited(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseDeposited(log types.Log) (*CoreOptimismDeposited, error) {
	event := new(CoreOptimismDeposited)
	if err := _CoreOptimism.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismFeatureFlagAllowAllSetIterator is returned from FilterFeatureFlagAllowAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowAllSet events raised by the CoreOptimism contract.
type CoreOptimismFeatureFlagAllowAllSetIterator struct {
	Event *CoreOptimismFeatureFlagAllowAllSet // Event containing the contract specifics and raw log

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
func (it *CoreOptimismFeatureFlagAllowAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismFeatureFlagAllowAllSet)
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
		it.Event = new(CoreOptimismFeatureFlagAllowAllSet)
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
func (it *CoreOptimismFeatureFlagAllowAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismFeatureFlagAllowAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismFeatureFlagAllowAllSet represents a FeatureFlagAllowAllSet event raised by the CoreOptimism contract.
type CoreOptimismFeatureFlagAllowAllSet struct {
	Feature  [32]byte
	AllowAll bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowAllSet is a free log retrieval operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_CoreOptimism *CoreOptimismFilterer) FilterFeatureFlagAllowAllSet(opts *bind.FilterOpts, feature [][32]byte) (*CoreOptimismFeatureFlagAllowAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismFeatureFlagAllowAllSetIterator{contract: _CoreOptimism.contract, event: "FeatureFlagAllowAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowAllSet is a free log subscription operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_CoreOptimism *CoreOptimismFilterer) WatchFeatureFlagAllowAllSet(opts *bind.WatchOpts, sink chan<- *CoreOptimismFeatureFlagAllowAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismFeatureFlagAllowAllSet)
				if err := _CoreOptimism.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParseFeatureFlagAllowAllSet(log types.Log) (*CoreOptimismFeatureFlagAllowAllSet, error) {
	event := new(CoreOptimismFeatureFlagAllowAllSet)
	if err := _CoreOptimism.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismFeatureFlagAllowlistAddedIterator is returned from FilterFeatureFlagAllowlistAdded and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistAdded events raised by the CoreOptimism contract.
type CoreOptimismFeatureFlagAllowlistAddedIterator struct {
	Event *CoreOptimismFeatureFlagAllowlistAdded // Event containing the contract specifics and raw log

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
func (it *CoreOptimismFeatureFlagAllowlistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismFeatureFlagAllowlistAdded)
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
		it.Event = new(CoreOptimismFeatureFlagAllowlistAdded)
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
func (it *CoreOptimismFeatureFlagAllowlistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismFeatureFlagAllowlistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismFeatureFlagAllowlistAdded represents a FeatureFlagAllowlistAdded event raised by the CoreOptimism contract.
type CoreOptimismFeatureFlagAllowlistAdded struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistAdded is a free log retrieval operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_CoreOptimism *CoreOptimismFilterer) FilterFeatureFlagAllowlistAdded(opts *bind.FilterOpts, feature [][32]byte) (*CoreOptimismFeatureFlagAllowlistAddedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismFeatureFlagAllowlistAddedIterator{contract: _CoreOptimism.contract, event: "FeatureFlagAllowlistAdded", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistAdded is a free log subscription operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_CoreOptimism *CoreOptimismFilterer) WatchFeatureFlagAllowlistAdded(opts *bind.WatchOpts, sink chan<- *CoreOptimismFeatureFlagAllowlistAdded, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismFeatureFlagAllowlistAdded)
				if err := _CoreOptimism.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParseFeatureFlagAllowlistAdded(log types.Log) (*CoreOptimismFeatureFlagAllowlistAdded, error) {
	event := new(CoreOptimismFeatureFlagAllowlistAdded)
	if err := _CoreOptimism.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismFeatureFlagAllowlistRemovedIterator is returned from FilterFeatureFlagAllowlistRemoved and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistRemoved events raised by the CoreOptimism contract.
type CoreOptimismFeatureFlagAllowlistRemovedIterator struct {
	Event *CoreOptimismFeatureFlagAllowlistRemoved // Event containing the contract specifics and raw log

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
func (it *CoreOptimismFeatureFlagAllowlistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismFeatureFlagAllowlistRemoved)
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
		it.Event = new(CoreOptimismFeatureFlagAllowlistRemoved)
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
func (it *CoreOptimismFeatureFlagAllowlistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismFeatureFlagAllowlistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismFeatureFlagAllowlistRemoved represents a FeatureFlagAllowlistRemoved event raised by the CoreOptimism contract.
type CoreOptimismFeatureFlagAllowlistRemoved struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistRemoved is a free log retrieval operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_CoreOptimism *CoreOptimismFilterer) FilterFeatureFlagAllowlistRemoved(opts *bind.FilterOpts, feature [][32]byte) (*CoreOptimismFeatureFlagAllowlistRemovedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismFeatureFlagAllowlistRemovedIterator{contract: _CoreOptimism.contract, event: "FeatureFlagAllowlistRemoved", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistRemoved is a free log subscription operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_CoreOptimism *CoreOptimismFilterer) WatchFeatureFlagAllowlistRemoved(opts *bind.WatchOpts, sink chan<- *CoreOptimismFeatureFlagAllowlistRemoved, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismFeatureFlagAllowlistRemoved)
				if err := _CoreOptimism.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParseFeatureFlagAllowlistRemoved(log types.Log) (*CoreOptimismFeatureFlagAllowlistRemoved, error) {
	event := new(CoreOptimismFeatureFlagAllowlistRemoved)
	if err := _CoreOptimism.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismFeatureFlagDeniersResetIterator is returned from FilterFeatureFlagDeniersReset and is used to iterate over the raw logs and unpacked data for FeatureFlagDeniersReset events raised by the CoreOptimism contract.
type CoreOptimismFeatureFlagDeniersResetIterator struct {
	Event *CoreOptimismFeatureFlagDeniersReset // Event containing the contract specifics and raw log

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
func (it *CoreOptimismFeatureFlagDeniersResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismFeatureFlagDeniersReset)
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
		it.Event = new(CoreOptimismFeatureFlagDeniersReset)
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
func (it *CoreOptimismFeatureFlagDeniersResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismFeatureFlagDeniersResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismFeatureFlagDeniersReset represents a FeatureFlagDeniersReset event raised by the CoreOptimism contract.
type CoreOptimismFeatureFlagDeniersReset struct {
	Feature [32]byte
	Deniers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDeniersReset is a free log retrieval operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_CoreOptimism *CoreOptimismFilterer) FilterFeatureFlagDeniersReset(opts *bind.FilterOpts, feature [][32]byte) (*CoreOptimismFeatureFlagDeniersResetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismFeatureFlagDeniersResetIterator{contract: _CoreOptimism.contract, event: "FeatureFlagDeniersReset", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDeniersReset is a free log subscription operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_CoreOptimism *CoreOptimismFilterer) WatchFeatureFlagDeniersReset(opts *bind.WatchOpts, sink chan<- *CoreOptimismFeatureFlagDeniersReset, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismFeatureFlagDeniersReset)
				if err := _CoreOptimism.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParseFeatureFlagDeniersReset(log types.Log) (*CoreOptimismFeatureFlagDeniersReset, error) {
	event := new(CoreOptimismFeatureFlagDeniersReset)
	if err := _CoreOptimism.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismFeatureFlagDenyAllSetIterator is returned from FilterFeatureFlagDenyAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagDenyAllSet events raised by the CoreOptimism contract.
type CoreOptimismFeatureFlagDenyAllSetIterator struct {
	Event *CoreOptimismFeatureFlagDenyAllSet // Event containing the contract specifics and raw log

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
func (it *CoreOptimismFeatureFlagDenyAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismFeatureFlagDenyAllSet)
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
		it.Event = new(CoreOptimismFeatureFlagDenyAllSet)
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
func (it *CoreOptimismFeatureFlagDenyAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismFeatureFlagDenyAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismFeatureFlagDenyAllSet represents a FeatureFlagDenyAllSet event raised by the CoreOptimism contract.
type CoreOptimismFeatureFlagDenyAllSet struct {
	Feature [32]byte
	DenyAll bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDenyAllSet is a free log retrieval operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_CoreOptimism *CoreOptimismFilterer) FilterFeatureFlagDenyAllSet(opts *bind.FilterOpts, feature [][32]byte) (*CoreOptimismFeatureFlagDenyAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismFeatureFlagDenyAllSetIterator{contract: _CoreOptimism.contract, event: "FeatureFlagDenyAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDenyAllSet is a free log subscription operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_CoreOptimism *CoreOptimismFilterer) WatchFeatureFlagDenyAllSet(opts *bind.WatchOpts, sink chan<- *CoreOptimismFeatureFlagDenyAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismFeatureFlagDenyAllSet)
				if err := _CoreOptimism.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParseFeatureFlagDenyAllSet(log types.Log) (*CoreOptimismFeatureFlagDenyAllSet, error) {
	event := new(CoreOptimismFeatureFlagDenyAllSet)
	if err := _CoreOptimism.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismIssuanceFeePaidIterator is returned from FilterIssuanceFeePaid and is used to iterate over the raw logs and unpacked data for IssuanceFeePaid events raised by the CoreOptimism contract.
type CoreOptimismIssuanceFeePaidIterator struct {
	Event *CoreOptimismIssuanceFeePaid // Event containing the contract specifics and raw log

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
func (it *CoreOptimismIssuanceFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismIssuanceFeePaid)
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
		it.Event = new(CoreOptimismIssuanceFeePaid)
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
func (it *CoreOptimismIssuanceFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismIssuanceFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismIssuanceFeePaid represents a IssuanceFeePaid event raised by the CoreOptimism contract.
type CoreOptimismIssuanceFeePaid struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	FeeAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIssuanceFeePaid is a free log retrieval operation binding the contract event 0x28d0fb10e1c8ce51490a16fb3b40bf23f8064f7c624d3a3852ad66683a25995d.
//
// Solidity: event IssuanceFeePaid(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 feeAmount)
func (_CoreOptimism *CoreOptimismFilterer) FilterIssuanceFeePaid(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int) (*CoreOptimismIssuanceFeePaidIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "IssuanceFeePaid", accountIdRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismIssuanceFeePaidIterator{contract: _CoreOptimism.contract, event: "IssuanceFeePaid", logs: logs, sub: sub}, nil
}

// WatchIssuanceFeePaid is a free log subscription operation binding the contract event 0x28d0fb10e1c8ce51490a16fb3b40bf23f8064f7c624d3a3852ad66683a25995d.
//
// Solidity: event IssuanceFeePaid(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 feeAmount)
func (_CoreOptimism *CoreOptimismFilterer) WatchIssuanceFeePaid(opts *bind.WatchOpts, sink chan<- *CoreOptimismIssuanceFeePaid, accountId []*big.Int, poolId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "IssuanceFeePaid", accountIdRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismIssuanceFeePaid)
				if err := _CoreOptimism.contract.UnpackLog(event, "IssuanceFeePaid", log); err != nil {
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

// ParseIssuanceFeePaid is a log parse operation binding the contract event 0x28d0fb10e1c8ce51490a16fb3b40bf23f8064f7c624d3a3852ad66683a25995d.
//
// Solidity: event IssuanceFeePaid(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 feeAmount)
func (_CoreOptimism *CoreOptimismFilterer) ParseIssuanceFeePaid(log types.Log) (*CoreOptimismIssuanceFeePaid, error) {
	event := new(CoreOptimismIssuanceFeePaid)
	if err := _CoreOptimism.contract.UnpackLog(event, "IssuanceFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismLiquidationIterator is returned from FilterLiquidation and is used to iterate over the raw logs and unpacked data for Liquidation events raised by the CoreOptimism contract.
type CoreOptimismLiquidationIterator struct {
	Event *CoreOptimismLiquidation // Event containing the contract specifics and raw log

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
func (it *CoreOptimismLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismLiquidation)
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
		it.Event = new(CoreOptimismLiquidation)
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
func (it *CoreOptimismLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismLiquidation represents a Liquidation event raised by the CoreOptimism contract.
type CoreOptimismLiquidation struct {
	AccountId            *big.Int
	PoolId               *big.Int
	CollateralType       common.Address
	LiquidationData      ILiquidationModuleLiquidationData
	LiquidateAsAccountId *big.Int
	Sender               common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLiquidation is a free log retrieval operation binding the contract event 0xe6c1b26644f880854bf954d4186be9f0b2d06d50fa0484b596e79d409c07a5fd.
//
// Solidity: event Liquidation(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterLiquidation(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*CoreOptimismLiquidationIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "Liquidation", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismLiquidationIterator{contract: _CoreOptimism.contract, event: "Liquidation", logs: logs, sub: sub}, nil
}

// WatchLiquidation is a free log subscription operation binding the contract event 0xe6c1b26644f880854bf954d4186be9f0b2d06d50fa0484b596e79d409c07a5fd.
//
// Solidity: event Liquidation(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchLiquidation(opts *bind.WatchOpts, sink chan<- *CoreOptimismLiquidation, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "Liquidation", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismLiquidation)
				if err := _CoreOptimism.contract.UnpackLog(event, "Liquidation", log); err != nil {
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

// ParseLiquidation is a log parse operation binding the contract event 0xe6c1b26644f880854bf954d4186be9f0b2d06d50fa0484b596e79d409c07a5fd.
//
// Solidity: event Liquidation(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseLiquidation(log types.Log) (*CoreOptimismLiquidation, error) {
	event := new(CoreOptimismLiquidation)
	if err := _CoreOptimism.contract.UnpackLog(event, "Liquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismMarketCollateralDepositedIterator is returned from FilterMarketCollateralDeposited and is used to iterate over the raw logs and unpacked data for MarketCollateralDeposited events raised by the CoreOptimism contract.
type CoreOptimismMarketCollateralDepositedIterator struct {
	Event *CoreOptimismMarketCollateralDeposited // Event containing the contract specifics and raw log

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
func (it *CoreOptimismMarketCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismMarketCollateralDeposited)
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
		it.Event = new(CoreOptimismMarketCollateralDeposited)
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
func (it *CoreOptimismMarketCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismMarketCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismMarketCollateralDeposited represents a MarketCollateralDeposited event raised by the CoreOptimism contract.
type CoreOptimismMarketCollateralDeposited struct {
	MarketId       *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketCollateralDeposited is a free log retrieval operation binding the contract event 0x6affa35c1a0ba6198bd74578c93e3a0eb32a96af0af26ce8adf06b933212ca1b.
//
// Solidity: event MarketCollateralDeposited(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterMarketCollateralDeposited(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreOptimismMarketCollateralDepositedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "MarketCollateralDeposited", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismMarketCollateralDepositedIterator{contract: _CoreOptimism.contract, event: "MarketCollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchMarketCollateralDeposited is a free log subscription operation binding the contract event 0x6affa35c1a0ba6198bd74578c93e3a0eb32a96af0af26ce8adf06b933212ca1b.
//
// Solidity: event MarketCollateralDeposited(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchMarketCollateralDeposited(opts *bind.WatchOpts, sink chan<- *CoreOptimismMarketCollateralDeposited, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "MarketCollateralDeposited", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismMarketCollateralDeposited)
				if err := _CoreOptimism.contract.UnpackLog(event, "MarketCollateralDeposited", log); err != nil {
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

// ParseMarketCollateralDeposited is a log parse operation binding the contract event 0x6affa35c1a0ba6198bd74578c93e3a0eb32a96af0af26ce8adf06b933212ca1b.
//
// Solidity: event MarketCollateralDeposited(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseMarketCollateralDeposited(log types.Log) (*CoreOptimismMarketCollateralDeposited, error) {
	event := new(CoreOptimismMarketCollateralDeposited)
	if err := _CoreOptimism.contract.UnpackLog(event, "MarketCollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismMarketCollateralWithdrawnIterator is returned from FilterMarketCollateralWithdrawn and is used to iterate over the raw logs and unpacked data for MarketCollateralWithdrawn events raised by the CoreOptimism contract.
type CoreOptimismMarketCollateralWithdrawnIterator struct {
	Event *CoreOptimismMarketCollateralWithdrawn // Event containing the contract specifics and raw log

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
func (it *CoreOptimismMarketCollateralWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismMarketCollateralWithdrawn)
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
		it.Event = new(CoreOptimismMarketCollateralWithdrawn)
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
func (it *CoreOptimismMarketCollateralWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismMarketCollateralWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismMarketCollateralWithdrawn represents a MarketCollateralWithdrawn event raised by the CoreOptimism contract.
type CoreOptimismMarketCollateralWithdrawn struct {
	MarketId       *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketCollateralWithdrawn is a free log retrieval operation binding the contract event 0x31fd88ca973dd7ee814313c50c2aaa7cd2d2142634582332b1e2671e6d97fd9a.
//
// Solidity: event MarketCollateralWithdrawn(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterMarketCollateralWithdrawn(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreOptimismMarketCollateralWithdrawnIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "MarketCollateralWithdrawn", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismMarketCollateralWithdrawnIterator{contract: _CoreOptimism.contract, event: "MarketCollateralWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMarketCollateralWithdrawn is a free log subscription operation binding the contract event 0x31fd88ca973dd7ee814313c50c2aaa7cd2d2142634582332b1e2671e6d97fd9a.
//
// Solidity: event MarketCollateralWithdrawn(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchMarketCollateralWithdrawn(opts *bind.WatchOpts, sink chan<- *CoreOptimismMarketCollateralWithdrawn, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "MarketCollateralWithdrawn", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismMarketCollateralWithdrawn)
				if err := _CoreOptimism.contract.UnpackLog(event, "MarketCollateralWithdrawn", log); err != nil {
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

// ParseMarketCollateralWithdrawn is a log parse operation binding the contract event 0x31fd88ca973dd7ee814313c50c2aaa7cd2d2142634582332b1e2671e6d97fd9a.
//
// Solidity: event MarketCollateralWithdrawn(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseMarketCollateralWithdrawn(log types.Log) (*CoreOptimismMarketCollateralWithdrawn, error) {
	event := new(CoreOptimismMarketCollateralWithdrawn)
	if err := _CoreOptimism.contract.UnpackLog(event, "MarketCollateralWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismMarketRegisteredIterator is returned from FilterMarketRegistered and is used to iterate over the raw logs and unpacked data for MarketRegistered events raised by the CoreOptimism contract.
type CoreOptimismMarketRegisteredIterator struct {
	Event *CoreOptimismMarketRegistered // Event containing the contract specifics and raw log

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
func (it *CoreOptimismMarketRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismMarketRegistered)
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
		it.Event = new(CoreOptimismMarketRegistered)
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
func (it *CoreOptimismMarketRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismMarketRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismMarketRegistered represents a MarketRegistered event raised by the CoreOptimism contract.
type CoreOptimismMarketRegistered struct {
	Market   common.Address
	MarketId *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketRegistered is a free log retrieval operation binding the contract event 0xeb87361ace8c1947e121293eb214f68d889d9cf273c48246b38c3cbf185748d0.
//
// Solidity: event MarketRegistered(address indexed market, uint128 indexed marketId, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterMarketRegistered(opts *bind.FilterOpts, market []common.Address, marketId []*big.Int, sender []common.Address) (*CoreOptimismMarketRegisteredIterator, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "MarketRegistered", marketRule, marketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismMarketRegisteredIterator{contract: _CoreOptimism.contract, event: "MarketRegistered", logs: logs, sub: sub}, nil
}

// WatchMarketRegistered is a free log subscription operation binding the contract event 0xeb87361ace8c1947e121293eb214f68d889d9cf273c48246b38c3cbf185748d0.
//
// Solidity: event MarketRegistered(address indexed market, uint128 indexed marketId, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchMarketRegistered(opts *bind.WatchOpts, sink chan<- *CoreOptimismMarketRegistered, market []common.Address, marketId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "MarketRegistered", marketRule, marketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismMarketRegistered)
				if err := _CoreOptimism.contract.UnpackLog(event, "MarketRegistered", log); err != nil {
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

// ParseMarketRegistered is a log parse operation binding the contract event 0xeb87361ace8c1947e121293eb214f68d889d9cf273c48246b38c3cbf185748d0.
//
// Solidity: event MarketRegistered(address indexed market, uint128 indexed marketId, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseMarketRegistered(log types.Log) (*CoreOptimismMarketRegistered, error) {
	event := new(CoreOptimismMarketRegistered)
	if err := _CoreOptimism.contract.UnpackLog(event, "MarketRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismMarketSystemFeePaidIterator is returned from FilterMarketSystemFeePaid and is used to iterate over the raw logs and unpacked data for MarketSystemFeePaid events raised by the CoreOptimism contract.
type CoreOptimismMarketSystemFeePaidIterator struct {
	Event *CoreOptimismMarketSystemFeePaid // Event containing the contract specifics and raw log

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
func (it *CoreOptimismMarketSystemFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismMarketSystemFeePaid)
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
		it.Event = new(CoreOptimismMarketSystemFeePaid)
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
func (it *CoreOptimismMarketSystemFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismMarketSystemFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismMarketSystemFeePaid represents a MarketSystemFeePaid event raised by the CoreOptimism contract.
type CoreOptimismMarketSystemFeePaid struct {
	MarketId  *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMarketSystemFeePaid is a free log retrieval operation binding the contract event 0x8b69fed8aed97ef9572216662359ece45fa52f2b5ff44a78b7ec3c5ef05153f8.
//
// Solidity: event MarketSystemFeePaid(uint128 indexed marketId, uint256 feeAmount)
func (_CoreOptimism *CoreOptimismFilterer) FilterMarketSystemFeePaid(opts *bind.FilterOpts, marketId []*big.Int) (*CoreOptimismMarketSystemFeePaidIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "MarketSystemFeePaid", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismMarketSystemFeePaidIterator{contract: _CoreOptimism.contract, event: "MarketSystemFeePaid", logs: logs, sub: sub}, nil
}

// WatchMarketSystemFeePaid is a free log subscription operation binding the contract event 0x8b69fed8aed97ef9572216662359ece45fa52f2b5ff44a78b7ec3c5ef05153f8.
//
// Solidity: event MarketSystemFeePaid(uint128 indexed marketId, uint256 feeAmount)
func (_CoreOptimism *CoreOptimismFilterer) WatchMarketSystemFeePaid(opts *bind.WatchOpts, sink chan<- *CoreOptimismMarketSystemFeePaid, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "MarketSystemFeePaid", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismMarketSystemFeePaid)
				if err := _CoreOptimism.contract.UnpackLog(event, "MarketSystemFeePaid", log); err != nil {
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

// ParseMarketSystemFeePaid is a log parse operation binding the contract event 0x8b69fed8aed97ef9572216662359ece45fa52f2b5ff44a78b7ec3c5ef05153f8.
//
// Solidity: event MarketSystemFeePaid(uint128 indexed marketId, uint256 feeAmount)
func (_CoreOptimism *CoreOptimismFilterer) ParseMarketSystemFeePaid(log types.Log) (*CoreOptimismMarketSystemFeePaid, error) {
	event := new(CoreOptimismMarketSystemFeePaid)
	if err := _CoreOptimism.contract.UnpackLog(event, "MarketSystemFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismMarketUsdDepositedIterator is returned from FilterMarketUsdDeposited and is used to iterate over the raw logs and unpacked data for MarketUsdDeposited events raised by the CoreOptimism contract.
type CoreOptimismMarketUsdDepositedIterator struct {
	Event *CoreOptimismMarketUsdDeposited // Event containing the contract specifics and raw log

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
func (it *CoreOptimismMarketUsdDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismMarketUsdDeposited)
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
		it.Event = new(CoreOptimismMarketUsdDeposited)
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
func (it *CoreOptimismMarketUsdDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismMarketUsdDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismMarketUsdDeposited represents a MarketUsdDeposited event raised by the CoreOptimism contract.
type CoreOptimismMarketUsdDeposited struct {
	MarketId *big.Int
	Target   common.Address
	Amount   *big.Int
	Market   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketUsdDeposited is a free log retrieval operation binding the contract event 0xcbdbec7c74466ca0e764de5f6bca5bef859e31a7d6284be22ac6b4beabddf835.
//
// Solidity: event MarketUsdDeposited(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_CoreOptimism *CoreOptimismFilterer) FilterMarketUsdDeposited(opts *bind.FilterOpts, marketId []*big.Int, target []common.Address, market []common.Address) (*CoreOptimismMarketUsdDepositedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "MarketUsdDeposited", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismMarketUsdDepositedIterator{contract: _CoreOptimism.contract, event: "MarketUsdDeposited", logs: logs, sub: sub}, nil
}

// WatchMarketUsdDeposited is a free log subscription operation binding the contract event 0xcbdbec7c74466ca0e764de5f6bca5bef859e31a7d6284be22ac6b4beabddf835.
//
// Solidity: event MarketUsdDeposited(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_CoreOptimism *CoreOptimismFilterer) WatchMarketUsdDeposited(opts *bind.WatchOpts, sink chan<- *CoreOptimismMarketUsdDeposited, marketId []*big.Int, target []common.Address, market []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "MarketUsdDeposited", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismMarketUsdDeposited)
				if err := _CoreOptimism.contract.UnpackLog(event, "MarketUsdDeposited", log); err != nil {
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

// ParseMarketUsdDeposited is a log parse operation binding the contract event 0xcbdbec7c74466ca0e764de5f6bca5bef859e31a7d6284be22ac6b4beabddf835.
//
// Solidity: event MarketUsdDeposited(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_CoreOptimism *CoreOptimismFilterer) ParseMarketUsdDeposited(log types.Log) (*CoreOptimismMarketUsdDeposited, error) {
	event := new(CoreOptimismMarketUsdDeposited)
	if err := _CoreOptimism.contract.UnpackLog(event, "MarketUsdDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismMarketUsdWithdrawnIterator is returned from FilterMarketUsdWithdrawn and is used to iterate over the raw logs and unpacked data for MarketUsdWithdrawn events raised by the CoreOptimism contract.
type CoreOptimismMarketUsdWithdrawnIterator struct {
	Event *CoreOptimismMarketUsdWithdrawn // Event containing the contract specifics and raw log

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
func (it *CoreOptimismMarketUsdWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismMarketUsdWithdrawn)
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
		it.Event = new(CoreOptimismMarketUsdWithdrawn)
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
func (it *CoreOptimismMarketUsdWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismMarketUsdWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismMarketUsdWithdrawn represents a MarketUsdWithdrawn event raised by the CoreOptimism contract.
type CoreOptimismMarketUsdWithdrawn struct {
	MarketId *big.Int
	Target   common.Address
	Amount   *big.Int
	Market   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketUsdWithdrawn is a free log retrieval operation binding the contract event 0x8669e4526579b1f1da479be2f96d1c7953fff439a4f98362a1c9ca07fe42ba41.
//
// Solidity: event MarketUsdWithdrawn(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_CoreOptimism *CoreOptimismFilterer) FilterMarketUsdWithdrawn(opts *bind.FilterOpts, marketId []*big.Int, target []common.Address, market []common.Address) (*CoreOptimismMarketUsdWithdrawnIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "MarketUsdWithdrawn", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismMarketUsdWithdrawnIterator{contract: _CoreOptimism.contract, event: "MarketUsdWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMarketUsdWithdrawn is a free log subscription operation binding the contract event 0x8669e4526579b1f1da479be2f96d1c7953fff439a4f98362a1c9ca07fe42ba41.
//
// Solidity: event MarketUsdWithdrawn(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_CoreOptimism *CoreOptimismFilterer) WatchMarketUsdWithdrawn(opts *bind.WatchOpts, sink chan<- *CoreOptimismMarketUsdWithdrawn, marketId []*big.Int, target []common.Address, market []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "MarketUsdWithdrawn", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismMarketUsdWithdrawn)
				if err := _CoreOptimism.contract.UnpackLog(event, "MarketUsdWithdrawn", log); err != nil {
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

// ParseMarketUsdWithdrawn is a log parse operation binding the contract event 0x8669e4526579b1f1da479be2f96d1c7953fff439a4f98362a1c9ca07fe42ba41.
//
// Solidity: event MarketUsdWithdrawn(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_CoreOptimism *CoreOptimismFilterer) ParseMarketUsdWithdrawn(log types.Log) (*CoreOptimismMarketUsdWithdrawn, error) {
	event := new(CoreOptimismMarketUsdWithdrawn)
	if err := _CoreOptimism.contract.UnpackLog(event, "MarketUsdWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismMaximumMarketCollateralConfiguredIterator is returned from FilterMaximumMarketCollateralConfigured and is used to iterate over the raw logs and unpacked data for MaximumMarketCollateralConfigured events raised by the CoreOptimism contract.
type CoreOptimismMaximumMarketCollateralConfiguredIterator struct {
	Event *CoreOptimismMaximumMarketCollateralConfigured // Event containing the contract specifics and raw log

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
func (it *CoreOptimismMaximumMarketCollateralConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismMaximumMarketCollateralConfigured)
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
		it.Event = new(CoreOptimismMaximumMarketCollateralConfigured)
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
func (it *CoreOptimismMaximumMarketCollateralConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismMaximumMarketCollateralConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismMaximumMarketCollateralConfigured represents a MaximumMarketCollateralConfigured event raised by the CoreOptimism contract.
type CoreOptimismMaximumMarketCollateralConfigured struct {
	MarketId       *big.Int
	CollateralType common.Address
	SystemAmount   *big.Int
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMaximumMarketCollateralConfigured is a free log retrieval operation binding the contract event 0x499c8fcfbc4341c37dcf444c890d42ef888d46aa586f97ceb20577c2635c8075.
//
// Solidity: event MaximumMarketCollateralConfigured(uint128 indexed marketId, address indexed collateralType, uint256 systemAmount, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) FilterMaximumMarketCollateralConfigured(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, owner []common.Address) (*CoreOptimismMaximumMarketCollateralConfiguredIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "MaximumMarketCollateralConfigured", marketIdRule, collateralTypeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismMaximumMarketCollateralConfiguredIterator{contract: _CoreOptimism.contract, event: "MaximumMarketCollateralConfigured", logs: logs, sub: sub}, nil
}

// WatchMaximumMarketCollateralConfigured is a free log subscription operation binding the contract event 0x499c8fcfbc4341c37dcf444c890d42ef888d46aa586f97ceb20577c2635c8075.
//
// Solidity: event MaximumMarketCollateralConfigured(uint128 indexed marketId, address indexed collateralType, uint256 systemAmount, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) WatchMaximumMarketCollateralConfigured(opts *bind.WatchOpts, sink chan<- *CoreOptimismMaximumMarketCollateralConfigured, marketId []*big.Int, collateralType []common.Address, owner []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "MaximumMarketCollateralConfigured", marketIdRule, collateralTypeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismMaximumMarketCollateralConfigured)
				if err := _CoreOptimism.contract.UnpackLog(event, "MaximumMarketCollateralConfigured", log); err != nil {
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

// ParseMaximumMarketCollateralConfigured is a log parse operation binding the contract event 0x499c8fcfbc4341c37dcf444c890d42ef888d46aa586f97ceb20577c2635c8075.
//
// Solidity: event MaximumMarketCollateralConfigured(uint128 indexed marketId, address indexed collateralType, uint256 systemAmount, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) ParseMaximumMarketCollateralConfigured(log types.Log) (*CoreOptimismMaximumMarketCollateralConfigured, error) {
	event := new(CoreOptimismMaximumMarketCollateralConfigured)
	if err := _CoreOptimism.contract.UnpackLog(event, "MaximumMarketCollateralConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismNewSupportedCrossChainNetworkIterator is returned from FilterNewSupportedCrossChainNetwork and is used to iterate over the raw logs and unpacked data for NewSupportedCrossChainNetwork events raised by the CoreOptimism contract.
type CoreOptimismNewSupportedCrossChainNetworkIterator struct {
	Event *CoreOptimismNewSupportedCrossChainNetwork // Event containing the contract specifics and raw log

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
func (it *CoreOptimismNewSupportedCrossChainNetworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismNewSupportedCrossChainNetwork)
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
		it.Event = new(CoreOptimismNewSupportedCrossChainNetwork)
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
func (it *CoreOptimismNewSupportedCrossChainNetworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismNewSupportedCrossChainNetworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismNewSupportedCrossChainNetwork represents a NewSupportedCrossChainNetwork event raised by the CoreOptimism contract.
type CoreOptimismNewSupportedCrossChainNetwork struct {
	NewChainId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewSupportedCrossChainNetwork is a free log retrieval operation binding the contract event 0x1874eb2a5288e478dcedf1d33291bd7293eeef5946ec516d2ef54a364b3f63d8.
//
// Solidity: event NewSupportedCrossChainNetwork(uint64 newChainId)
func (_CoreOptimism *CoreOptimismFilterer) FilterNewSupportedCrossChainNetwork(opts *bind.FilterOpts) (*CoreOptimismNewSupportedCrossChainNetworkIterator, error) {

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "NewSupportedCrossChainNetwork")
	if err != nil {
		return nil, err
	}
	return &CoreOptimismNewSupportedCrossChainNetworkIterator{contract: _CoreOptimism.contract, event: "NewSupportedCrossChainNetwork", logs: logs, sub: sub}, nil
}

// WatchNewSupportedCrossChainNetwork is a free log subscription operation binding the contract event 0x1874eb2a5288e478dcedf1d33291bd7293eeef5946ec516d2ef54a364b3f63d8.
//
// Solidity: event NewSupportedCrossChainNetwork(uint64 newChainId)
func (_CoreOptimism *CoreOptimismFilterer) WatchNewSupportedCrossChainNetwork(opts *bind.WatchOpts, sink chan<- *CoreOptimismNewSupportedCrossChainNetwork) (event.Subscription, error) {

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "NewSupportedCrossChainNetwork")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismNewSupportedCrossChainNetwork)
				if err := _CoreOptimism.contract.UnpackLog(event, "NewSupportedCrossChainNetwork", log); err != nil {
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

// ParseNewSupportedCrossChainNetwork is a log parse operation binding the contract event 0x1874eb2a5288e478dcedf1d33291bd7293eeef5946ec516d2ef54a364b3f63d8.
//
// Solidity: event NewSupportedCrossChainNetwork(uint64 newChainId)
func (_CoreOptimism *CoreOptimismFilterer) ParseNewSupportedCrossChainNetwork(log types.Log) (*CoreOptimismNewSupportedCrossChainNetwork, error) {
	event := new(CoreOptimismNewSupportedCrossChainNetwork)
	if err := _CoreOptimism.contract.UnpackLog(event, "NewSupportedCrossChainNetwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the CoreOptimism contract.
type CoreOptimismOwnerChangedIterator struct {
	Event *CoreOptimismOwnerChanged // Event containing the contract specifics and raw log

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
func (it *CoreOptimismOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismOwnerChanged)
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
		it.Event = new(CoreOptimismOwnerChanged)
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
func (it *CoreOptimismOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismOwnerChanged represents a OwnerChanged event raised by the CoreOptimism contract.
type CoreOptimismOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_CoreOptimism *CoreOptimismFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*CoreOptimismOwnerChangedIterator, error) {

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &CoreOptimismOwnerChangedIterator{contract: _CoreOptimism.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_CoreOptimism *CoreOptimismFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *CoreOptimismOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismOwnerChanged)
				if err := _CoreOptimism.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParseOwnerChanged(log types.Log) (*CoreOptimismOwnerChanged, error) {
	event := new(CoreOptimismOwnerChanged)
	if err := _CoreOptimism.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the CoreOptimism contract.
type CoreOptimismOwnerNominatedIterator struct {
	Event *CoreOptimismOwnerNominated // Event containing the contract specifics and raw log

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
func (it *CoreOptimismOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismOwnerNominated)
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
		it.Event = new(CoreOptimismOwnerNominated)
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
func (it *CoreOptimismOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismOwnerNominated represents a OwnerNominated event raised by the CoreOptimism contract.
type CoreOptimismOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_CoreOptimism *CoreOptimismFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*CoreOptimismOwnerNominatedIterator, error) {

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &CoreOptimismOwnerNominatedIterator{contract: _CoreOptimism.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_CoreOptimism *CoreOptimismFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *CoreOptimismOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismOwnerNominated)
				if err := _CoreOptimism.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParseOwnerNominated(log types.Log) (*CoreOptimismOwnerNominated, error) {
	event := new(CoreOptimismOwnerNominated)
	if err := _CoreOptimism.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPermissionGrantedIterator is returned from FilterPermissionGranted and is used to iterate over the raw logs and unpacked data for PermissionGranted events raised by the CoreOptimism contract.
type CoreOptimismPermissionGrantedIterator struct {
	Event *CoreOptimismPermissionGranted // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPermissionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPermissionGranted)
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
		it.Event = new(CoreOptimismPermissionGranted)
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
func (it *CoreOptimismPermissionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPermissionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPermissionGranted represents a PermissionGranted event raised by the CoreOptimism contract.
type CoreOptimismPermissionGranted struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionGranted is a free log retrieval operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterPermissionGranted(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*CoreOptimismPermissionGrantedIterator, error) {

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

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPermissionGrantedIterator{contract: _CoreOptimism.contract, event: "PermissionGranted", logs: logs, sub: sub}, nil
}

// WatchPermissionGranted is a free log subscription operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchPermissionGranted(opts *bind.WatchOpts, sink chan<- *CoreOptimismPermissionGranted, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPermissionGranted)
				if err := _CoreOptimism.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParsePermissionGranted(log types.Log) (*CoreOptimismPermissionGranted, error) {
	event := new(CoreOptimismPermissionGranted)
	if err := _CoreOptimism.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPermissionRevokedIterator is returned from FilterPermissionRevoked and is used to iterate over the raw logs and unpacked data for PermissionRevoked events raised by the CoreOptimism contract.
type CoreOptimismPermissionRevokedIterator struct {
	Event *CoreOptimismPermissionRevoked // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPermissionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPermissionRevoked)
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
		it.Event = new(CoreOptimismPermissionRevoked)
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
func (it *CoreOptimismPermissionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPermissionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPermissionRevoked represents a PermissionRevoked event raised by the CoreOptimism contract.
type CoreOptimismPermissionRevoked struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionRevoked is a free log retrieval operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterPermissionRevoked(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*CoreOptimismPermissionRevokedIterator, error) {

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

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPermissionRevokedIterator{contract: _CoreOptimism.contract, event: "PermissionRevoked", logs: logs, sub: sub}, nil
}

// WatchPermissionRevoked is a free log subscription operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchPermissionRevoked(opts *bind.WatchOpts, sink chan<- *CoreOptimismPermissionRevoked, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPermissionRevoked)
				if err := _CoreOptimism.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParsePermissionRevoked(log types.Log) (*CoreOptimismPermissionRevoked, error) {
	event := new(CoreOptimismPermissionRevoked)
	if err := _CoreOptimism.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolApprovedAddedIterator is returned from FilterPoolApprovedAdded and is used to iterate over the raw logs and unpacked data for PoolApprovedAdded events raised by the CoreOptimism contract.
type CoreOptimismPoolApprovedAddedIterator struct {
	Event *CoreOptimismPoolApprovedAdded // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolApprovedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolApprovedAdded)
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
		it.Event = new(CoreOptimismPoolApprovedAdded)
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
func (it *CoreOptimismPoolApprovedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolApprovedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolApprovedAdded represents a PoolApprovedAdded event raised by the CoreOptimism contract.
type CoreOptimismPoolApprovedAdded struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolApprovedAdded is a free log retrieval operation binding the contract event 0x7d5bdf4e8c44e0b5a8249bf03c2a1febd848cc7f580efd7b1703301c5b1a9e4e.
//
// Solidity: event PoolApprovedAdded(uint256 poolId)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolApprovedAdded(opts *bind.FilterOpts) (*CoreOptimismPoolApprovedAddedIterator, error) {

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolApprovedAdded")
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolApprovedAddedIterator{contract: _CoreOptimism.contract, event: "PoolApprovedAdded", logs: logs, sub: sub}, nil
}

// WatchPoolApprovedAdded is a free log subscription operation binding the contract event 0x7d5bdf4e8c44e0b5a8249bf03c2a1febd848cc7f580efd7b1703301c5b1a9e4e.
//
// Solidity: event PoolApprovedAdded(uint256 poolId)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolApprovedAdded(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolApprovedAdded) (event.Subscription, error) {

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolApprovedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolApprovedAdded)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolApprovedAdded", log); err != nil {
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

// ParsePoolApprovedAdded is a log parse operation binding the contract event 0x7d5bdf4e8c44e0b5a8249bf03c2a1febd848cc7f580efd7b1703301c5b1a9e4e.
//
// Solidity: event PoolApprovedAdded(uint256 poolId)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolApprovedAdded(log types.Log) (*CoreOptimismPoolApprovedAdded, error) {
	event := new(CoreOptimismPoolApprovedAdded)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolApprovedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolApprovedRemovedIterator is returned from FilterPoolApprovedRemoved and is used to iterate over the raw logs and unpacked data for PoolApprovedRemoved events raised by the CoreOptimism contract.
type CoreOptimismPoolApprovedRemovedIterator struct {
	Event *CoreOptimismPoolApprovedRemoved // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolApprovedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolApprovedRemoved)
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
		it.Event = new(CoreOptimismPoolApprovedRemoved)
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
func (it *CoreOptimismPoolApprovedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolApprovedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolApprovedRemoved represents a PoolApprovedRemoved event raised by the CoreOptimism contract.
type CoreOptimismPoolApprovedRemoved struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolApprovedRemoved is a free log retrieval operation binding the contract event 0xc1567ee9983f306f073ea7d59a7fb5680ce07985f8b49cc50d00a3a9f748d3c2.
//
// Solidity: event PoolApprovedRemoved(uint256 poolId)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolApprovedRemoved(opts *bind.FilterOpts) (*CoreOptimismPoolApprovedRemovedIterator, error) {

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolApprovedRemoved")
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolApprovedRemovedIterator{contract: _CoreOptimism.contract, event: "PoolApprovedRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolApprovedRemoved is a free log subscription operation binding the contract event 0xc1567ee9983f306f073ea7d59a7fb5680ce07985f8b49cc50d00a3a9f748d3c2.
//
// Solidity: event PoolApprovedRemoved(uint256 poolId)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolApprovedRemoved(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolApprovedRemoved) (event.Subscription, error) {

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolApprovedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolApprovedRemoved)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolApprovedRemoved", log); err != nil {
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

// ParsePoolApprovedRemoved is a log parse operation binding the contract event 0xc1567ee9983f306f073ea7d59a7fb5680ce07985f8b49cc50d00a3a9f748d3c2.
//
// Solidity: event PoolApprovedRemoved(uint256 poolId)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolApprovedRemoved(log types.Log) (*CoreOptimismPoolApprovedRemoved, error) {
	event := new(CoreOptimismPoolApprovedRemoved)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolApprovedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolCollateralConfigurationUpdatedIterator is returned from FilterPoolCollateralConfigurationUpdated and is used to iterate over the raw logs and unpacked data for PoolCollateralConfigurationUpdated events raised by the CoreOptimism contract.
type CoreOptimismPoolCollateralConfigurationUpdatedIterator struct {
	Event *CoreOptimismPoolCollateralConfigurationUpdated // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolCollateralConfigurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolCollateralConfigurationUpdated)
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
		it.Event = new(CoreOptimismPoolCollateralConfigurationUpdated)
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
func (it *CoreOptimismPoolCollateralConfigurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolCollateralConfigurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolCollateralConfigurationUpdated represents a PoolCollateralConfigurationUpdated event raised by the CoreOptimism contract.
type CoreOptimismPoolCollateralConfigurationUpdated struct {
	PoolId         *big.Int
	CollateralType common.Address
	Config         PoolCollateralConfigurationData
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolCollateralConfigurationUpdated is a free log retrieval operation binding the contract event 0x5ebb5c59166ab9735b293a159ee2129e61d16b526867763f25557a275a2aad92.
//
// Solidity: event PoolCollateralConfigurationUpdated(uint128 indexed poolId, address collateralType, (uint256,uint256) config)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolCollateralConfigurationUpdated(opts *bind.FilterOpts, poolId []*big.Int) (*CoreOptimismPoolCollateralConfigurationUpdatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolCollateralConfigurationUpdated", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolCollateralConfigurationUpdatedIterator{contract: _CoreOptimism.contract, event: "PoolCollateralConfigurationUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolCollateralConfigurationUpdated is a free log subscription operation binding the contract event 0x5ebb5c59166ab9735b293a159ee2129e61d16b526867763f25557a275a2aad92.
//
// Solidity: event PoolCollateralConfigurationUpdated(uint128 indexed poolId, address collateralType, (uint256,uint256) config)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolCollateralConfigurationUpdated(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolCollateralConfigurationUpdated, poolId []*big.Int) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolCollateralConfigurationUpdated", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolCollateralConfigurationUpdated)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolCollateralConfigurationUpdated", log); err != nil {
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

// ParsePoolCollateralConfigurationUpdated is a log parse operation binding the contract event 0x5ebb5c59166ab9735b293a159ee2129e61d16b526867763f25557a275a2aad92.
//
// Solidity: event PoolCollateralConfigurationUpdated(uint128 indexed poolId, address collateralType, (uint256,uint256) config)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolCollateralConfigurationUpdated(log types.Log) (*CoreOptimismPoolCollateralConfigurationUpdated, error) {
	event := new(CoreOptimismPoolCollateralConfigurationUpdated)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolCollateralConfigurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolCollateralDisabledByDefaultSetIterator is returned from FilterPoolCollateralDisabledByDefaultSet and is used to iterate over the raw logs and unpacked data for PoolCollateralDisabledByDefaultSet events raised by the CoreOptimism contract.
type CoreOptimismPoolCollateralDisabledByDefaultSetIterator struct {
	Event *CoreOptimismPoolCollateralDisabledByDefaultSet // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolCollateralDisabledByDefaultSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolCollateralDisabledByDefaultSet)
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
		it.Event = new(CoreOptimismPoolCollateralDisabledByDefaultSet)
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
func (it *CoreOptimismPoolCollateralDisabledByDefaultSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolCollateralDisabledByDefaultSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolCollateralDisabledByDefaultSet represents a PoolCollateralDisabledByDefaultSet event raised by the CoreOptimism contract.
type CoreOptimismPoolCollateralDisabledByDefaultSet struct {
	PoolId   *big.Int
	Disabled bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPoolCollateralDisabledByDefaultSet is a free log retrieval operation binding the contract event 0xe0ed98ef42e6a4a881ae0d3c4459c9ed06a36a2144e02efc11823c6cae515bf2.
//
// Solidity: event PoolCollateralDisabledByDefaultSet(uint128 poolId, bool disabled)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolCollateralDisabledByDefaultSet(opts *bind.FilterOpts) (*CoreOptimismPoolCollateralDisabledByDefaultSetIterator, error) {

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolCollateralDisabledByDefaultSet")
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolCollateralDisabledByDefaultSetIterator{contract: _CoreOptimism.contract, event: "PoolCollateralDisabledByDefaultSet", logs: logs, sub: sub}, nil
}

// WatchPoolCollateralDisabledByDefaultSet is a free log subscription operation binding the contract event 0xe0ed98ef42e6a4a881ae0d3c4459c9ed06a36a2144e02efc11823c6cae515bf2.
//
// Solidity: event PoolCollateralDisabledByDefaultSet(uint128 poolId, bool disabled)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolCollateralDisabledByDefaultSet(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolCollateralDisabledByDefaultSet) (event.Subscription, error) {

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolCollateralDisabledByDefaultSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolCollateralDisabledByDefaultSet)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolCollateralDisabledByDefaultSet", log); err != nil {
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

// ParsePoolCollateralDisabledByDefaultSet is a log parse operation binding the contract event 0xe0ed98ef42e6a4a881ae0d3c4459c9ed06a36a2144e02efc11823c6cae515bf2.
//
// Solidity: event PoolCollateralDisabledByDefaultSet(uint128 poolId, bool disabled)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolCollateralDisabledByDefaultSet(log types.Log) (*CoreOptimismPoolCollateralDisabledByDefaultSet, error) {
	event := new(CoreOptimismPoolCollateralDisabledByDefaultSet)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolCollateralDisabledByDefaultSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolConfigurationSetIterator is returned from FilterPoolConfigurationSet and is used to iterate over the raw logs and unpacked data for PoolConfigurationSet events raised by the CoreOptimism contract.
type CoreOptimismPoolConfigurationSetIterator struct {
	Event *CoreOptimismPoolConfigurationSet // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolConfigurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolConfigurationSet)
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
		it.Event = new(CoreOptimismPoolConfigurationSet)
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
func (it *CoreOptimismPoolConfigurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolConfigurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolConfigurationSet represents a PoolConfigurationSet event raised by the CoreOptimism contract.
type CoreOptimismPoolConfigurationSet struct {
	PoolId  *big.Int
	Markets []MarketConfigurationData
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolConfigurationSet is a free log retrieval operation binding the contract event 0xdd812c2e47943d98e6c66b2b9872d1f9270b8523c82eb60ad5c8d580a614081c.
//
// Solidity: event PoolConfigurationSet(uint128 indexed poolId, (uint128,uint128,int128)[] markets, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolConfigurationSet(opts *bind.FilterOpts, poolId []*big.Int, sender []common.Address) (*CoreOptimismPoolConfigurationSetIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolConfigurationSet", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolConfigurationSetIterator{contract: _CoreOptimism.contract, event: "PoolConfigurationSet", logs: logs, sub: sub}, nil
}

// WatchPoolConfigurationSet is a free log subscription operation binding the contract event 0xdd812c2e47943d98e6c66b2b9872d1f9270b8523c82eb60ad5c8d580a614081c.
//
// Solidity: event PoolConfigurationSet(uint128 indexed poolId, (uint128,uint128,int128)[] markets, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolConfigurationSet(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolConfigurationSet, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolConfigurationSet", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolConfigurationSet)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolConfigurationSet", log); err != nil {
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

// ParsePoolConfigurationSet is a log parse operation binding the contract event 0xdd812c2e47943d98e6c66b2b9872d1f9270b8523c82eb60ad5c8d580a614081c.
//
// Solidity: event PoolConfigurationSet(uint128 indexed poolId, (uint128,uint128,int128)[] markets, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolConfigurationSet(log types.Log) (*CoreOptimismPoolConfigurationSet, error) {
	event := new(CoreOptimismPoolConfigurationSet)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolConfigurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the CoreOptimism contract.
type CoreOptimismPoolCreatedIterator struct {
	Event *CoreOptimismPoolCreated // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolCreated)
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
		it.Event = new(CoreOptimismPoolCreated)
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
func (it *CoreOptimismPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolCreated represents a PoolCreated event raised by the CoreOptimism contract.
type CoreOptimismPoolCreated struct {
	PoolId *big.Int
	Owner  common.Address
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xb1517ad708e5f9a104c30d3f1ff749d55833b1d03bf472013c29888e741cf340.
//
// Solidity: event PoolCreated(uint128 indexed poolId, address indexed owner, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolCreated(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address, sender []common.Address) (*CoreOptimismPoolCreatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolCreated", poolIdRule, ownerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolCreatedIterator{contract: _CoreOptimism.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xb1517ad708e5f9a104c30d3f1ff749d55833b1d03bf472013c29888e741cf340.
//
// Solidity: event PoolCreated(uint128 indexed poolId, address indexed owner, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolCreated, poolId []*big.Int, owner []common.Address, sender []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolCreated", poolIdRule, ownerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolCreated)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0xb1517ad708e5f9a104c30d3f1ff749d55833b1d03bf472013c29888e741cf340.
//
// Solidity: event PoolCreated(uint128 indexed poolId, address indexed owner, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolCreated(log types.Log) (*CoreOptimismPoolCreated, error) {
	event := new(CoreOptimismPoolCreated)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolNameUpdatedIterator is returned from FilterPoolNameUpdated and is used to iterate over the raw logs and unpacked data for PoolNameUpdated events raised by the CoreOptimism contract.
type CoreOptimismPoolNameUpdatedIterator struct {
	Event *CoreOptimismPoolNameUpdated // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolNameUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolNameUpdated)
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
		it.Event = new(CoreOptimismPoolNameUpdated)
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
func (it *CoreOptimismPoolNameUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolNameUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolNameUpdated represents a PoolNameUpdated event raised by the CoreOptimism contract.
type CoreOptimismPoolNameUpdated struct {
	PoolId *big.Int
	Name   string
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNameUpdated is a free log retrieval operation binding the contract event 0x63b42abaf7e145a993f20bc64259f45d09c43d18838ab0bca078b15093ac55f4.
//
// Solidity: event PoolNameUpdated(uint128 indexed poolId, string name, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolNameUpdated(opts *bind.FilterOpts, poolId []*big.Int, sender []common.Address) (*CoreOptimismPoolNameUpdatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolNameUpdated", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolNameUpdatedIterator{contract: _CoreOptimism.contract, event: "PoolNameUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolNameUpdated is a free log subscription operation binding the contract event 0x63b42abaf7e145a993f20bc64259f45d09c43d18838ab0bca078b15093ac55f4.
//
// Solidity: event PoolNameUpdated(uint128 indexed poolId, string name, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolNameUpdated(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolNameUpdated, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolNameUpdated", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolNameUpdated)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolNameUpdated", log); err != nil {
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

// ParsePoolNameUpdated is a log parse operation binding the contract event 0x63b42abaf7e145a993f20bc64259f45d09c43d18838ab0bca078b15093ac55f4.
//
// Solidity: event PoolNameUpdated(uint128 indexed poolId, string name, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolNameUpdated(log types.Log) (*CoreOptimismPoolNameUpdated, error) {
	event := new(CoreOptimismPoolNameUpdated)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolNameUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolNominationRenouncedIterator is returned from FilterPoolNominationRenounced and is used to iterate over the raw logs and unpacked data for PoolNominationRenounced events raised by the CoreOptimism contract.
type CoreOptimismPoolNominationRenouncedIterator struct {
	Event *CoreOptimismPoolNominationRenounced // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolNominationRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolNominationRenounced)
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
		it.Event = new(CoreOptimismPoolNominationRenounced)
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
func (it *CoreOptimismPoolNominationRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolNominationRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolNominationRenounced represents a PoolNominationRenounced event raised by the CoreOptimism contract.
type CoreOptimismPoolNominationRenounced struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNominationRenounced is a free log retrieval operation binding the contract event 0x28301da5fb0feefb138efa6310af4547a74f415d62616f90519436dc169c3ae0.
//
// Solidity: event PoolNominationRenounced(uint128 indexed poolId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolNominationRenounced(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*CoreOptimismPoolNominationRenouncedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolNominationRenounced", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolNominationRenouncedIterator{contract: _CoreOptimism.contract, event: "PoolNominationRenounced", logs: logs, sub: sub}, nil
}

// WatchPoolNominationRenounced is a free log subscription operation binding the contract event 0x28301da5fb0feefb138efa6310af4547a74f415d62616f90519436dc169c3ae0.
//
// Solidity: event PoolNominationRenounced(uint128 indexed poolId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolNominationRenounced(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolNominationRenounced, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolNominationRenounced", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolNominationRenounced)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolNominationRenounced", log); err != nil {
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

// ParsePoolNominationRenounced is a log parse operation binding the contract event 0x28301da5fb0feefb138efa6310af4547a74f415d62616f90519436dc169c3ae0.
//
// Solidity: event PoolNominationRenounced(uint128 indexed poolId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolNominationRenounced(log types.Log) (*CoreOptimismPoolNominationRenounced, error) {
	event := new(CoreOptimismPoolNominationRenounced)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolNominationRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolNominationRevokedIterator is returned from FilterPoolNominationRevoked and is used to iterate over the raw logs and unpacked data for PoolNominationRevoked events raised by the CoreOptimism contract.
type CoreOptimismPoolNominationRevokedIterator struct {
	Event *CoreOptimismPoolNominationRevoked // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolNominationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolNominationRevoked)
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
		it.Event = new(CoreOptimismPoolNominationRevoked)
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
func (it *CoreOptimismPoolNominationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolNominationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolNominationRevoked represents a PoolNominationRevoked event raised by the CoreOptimism contract.
type CoreOptimismPoolNominationRevoked struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNominationRevoked is a free log retrieval operation binding the contract event 0xa20a605599b6da4a06e0662f1284c442a576bc452b77a38c8c55805cb82a1865.
//
// Solidity: event PoolNominationRevoked(uint128 indexed poolId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolNominationRevoked(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*CoreOptimismPoolNominationRevokedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolNominationRevoked", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolNominationRevokedIterator{contract: _CoreOptimism.contract, event: "PoolNominationRevoked", logs: logs, sub: sub}, nil
}

// WatchPoolNominationRevoked is a free log subscription operation binding the contract event 0xa20a605599b6da4a06e0662f1284c442a576bc452b77a38c8c55805cb82a1865.
//
// Solidity: event PoolNominationRevoked(uint128 indexed poolId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolNominationRevoked(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolNominationRevoked, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolNominationRevoked", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolNominationRevoked)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolNominationRevoked", log); err != nil {
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

// ParsePoolNominationRevoked is a log parse operation binding the contract event 0xa20a605599b6da4a06e0662f1284c442a576bc452b77a38c8c55805cb82a1865.
//
// Solidity: event PoolNominationRevoked(uint128 indexed poolId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolNominationRevoked(log types.Log) (*CoreOptimismPoolNominationRevoked, error) {
	event := new(CoreOptimismPoolNominationRevoked)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolNominationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolOwnerNominatedIterator is returned from FilterPoolOwnerNominated and is used to iterate over the raw logs and unpacked data for PoolOwnerNominated events raised by the CoreOptimism contract.
type CoreOptimismPoolOwnerNominatedIterator struct {
	Event *CoreOptimismPoolOwnerNominated // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolOwnerNominated)
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
		it.Event = new(CoreOptimismPoolOwnerNominated)
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
func (it *CoreOptimismPoolOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolOwnerNominated represents a PoolOwnerNominated event raised by the CoreOptimism contract.
type CoreOptimismPoolOwnerNominated struct {
	PoolId         *big.Int
	NominatedOwner common.Address
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolOwnerNominated is a free log retrieval operation binding the contract event 0x55d98f82a53fb5776e9ea48d624ab9cb015b51a45249b1ed8425fc857c82f4f8.
//
// Solidity: event PoolOwnerNominated(uint128 indexed poolId, address indexed nominatedOwner, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolOwnerNominated(opts *bind.FilterOpts, poolId []*big.Int, nominatedOwner []common.Address, owner []common.Address) (*CoreOptimismPoolOwnerNominatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var nominatedOwnerRule []interface{}
	for _, nominatedOwnerItem := range nominatedOwner {
		nominatedOwnerRule = append(nominatedOwnerRule, nominatedOwnerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolOwnerNominated", poolIdRule, nominatedOwnerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolOwnerNominatedIterator{contract: _CoreOptimism.contract, event: "PoolOwnerNominated", logs: logs, sub: sub}, nil
}

// WatchPoolOwnerNominated is a free log subscription operation binding the contract event 0x55d98f82a53fb5776e9ea48d624ab9cb015b51a45249b1ed8425fc857c82f4f8.
//
// Solidity: event PoolOwnerNominated(uint128 indexed poolId, address indexed nominatedOwner, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolOwnerNominated(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolOwnerNominated, poolId []*big.Int, nominatedOwner []common.Address, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var nominatedOwnerRule []interface{}
	for _, nominatedOwnerItem := range nominatedOwner {
		nominatedOwnerRule = append(nominatedOwnerRule, nominatedOwnerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolOwnerNominated", poolIdRule, nominatedOwnerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolOwnerNominated)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolOwnerNominated", log); err != nil {
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

// ParsePoolOwnerNominated is a log parse operation binding the contract event 0x55d98f82a53fb5776e9ea48d624ab9cb015b51a45249b1ed8425fc857c82f4f8.
//
// Solidity: event PoolOwnerNominated(uint128 indexed poolId, address indexed nominatedOwner, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolOwnerNominated(log types.Log) (*CoreOptimismPoolOwnerNominated, error) {
	event := new(CoreOptimismPoolOwnerNominated)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolOwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPoolOwnershipAcceptedIterator is returned from FilterPoolOwnershipAccepted and is used to iterate over the raw logs and unpacked data for PoolOwnershipAccepted events raised by the CoreOptimism contract.
type CoreOptimismPoolOwnershipAcceptedIterator struct {
	Event *CoreOptimismPoolOwnershipAccepted // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPoolOwnershipAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPoolOwnershipAccepted)
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
		it.Event = new(CoreOptimismPoolOwnershipAccepted)
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
func (it *CoreOptimismPoolOwnershipAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPoolOwnershipAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPoolOwnershipAccepted represents a PoolOwnershipAccepted event raised by the CoreOptimism contract.
type CoreOptimismPoolOwnershipAccepted struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolOwnershipAccepted is a free log retrieval operation binding the contract event 0x4f86f2ce8b08e27d0e470f4269b71c3bbc68407d51a2e692f6573236074ebc5a.
//
// Solidity: event PoolOwnershipAccepted(uint128 indexed poolId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) FilterPoolOwnershipAccepted(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*CoreOptimismPoolOwnershipAcceptedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PoolOwnershipAccepted", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPoolOwnershipAcceptedIterator{contract: _CoreOptimism.contract, event: "PoolOwnershipAccepted", logs: logs, sub: sub}, nil
}

// WatchPoolOwnershipAccepted is a free log subscription operation binding the contract event 0x4f86f2ce8b08e27d0e470f4269b71c3bbc68407d51a2e692f6573236074ebc5a.
//
// Solidity: event PoolOwnershipAccepted(uint128 indexed poolId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) WatchPoolOwnershipAccepted(opts *bind.WatchOpts, sink chan<- *CoreOptimismPoolOwnershipAccepted, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PoolOwnershipAccepted", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPoolOwnershipAccepted)
				if err := _CoreOptimism.contract.UnpackLog(event, "PoolOwnershipAccepted", log); err != nil {
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

// ParsePoolOwnershipAccepted is a log parse operation binding the contract event 0x4f86f2ce8b08e27d0e470f4269b71c3bbc68407d51a2e692f6573236074ebc5a.
//
// Solidity: event PoolOwnershipAccepted(uint128 indexed poolId, address indexed owner)
func (_CoreOptimism *CoreOptimismFilterer) ParsePoolOwnershipAccepted(log types.Log) (*CoreOptimismPoolOwnershipAccepted, error) {
	event := new(CoreOptimismPoolOwnershipAccepted)
	if err := _CoreOptimism.contract.UnpackLog(event, "PoolOwnershipAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismPreferredPoolSetIterator is returned from FilterPreferredPoolSet and is used to iterate over the raw logs and unpacked data for PreferredPoolSet events raised by the CoreOptimism contract.
type CoreOptimismPreferredPoolSetIterator struct {
	Event *CoreOptimismPreferredPoolSet // Event containing the contract specifics and raw log

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
func (it *CoreOptimismPreferredPoolSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismPreferredPoolSet)
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
		it.Event = new(CoreOptimismPreferredPoolSet)
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
func (it *CoreOptimismPreferredPoolSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismPreferredPoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismPreferredPoolSet represents a PreferredPoolSet event raised by the CoreOptimism contract.
type CoreOptimismPreferredPoolSet struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPreferredPoolSet is a free log retrieval operation binding the contract event 0x7e7cb4726e710dc12fad41f158c37a4a71af3a6f053b8b13670d35c710139a56.
//
// Solidity: event PreferredPoolSet(uint256 poolId)
func (_CoreOptimism *CoreOptimismFilterer) FilterPreferredPoolSet(opts *bind.FilterOpts) (*CoreOptimismPreferredPoolSetIterator, error) {

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "PreferredPoolSet")
	if err != nil {
		return nil, err
	}
	return &CoreOptimismPreferredPoolSetIterator{contract: _CoreOptimism.contract, event: "PreferredPoolSet", logs: logs, sub: sub}, nil
}

// WatchPreferredPoolSet is a free log subscription operation binding the contract event 0x7e7cb4726e710dc12fad41f158c37a4a71af3a6f053b8b13670d35c710139a56.
//
// Solidity: event PreferredPoolSet(uint256 poolId)
func (_CoreOptimism *CoreOptimismFilterer) WatchPreferredPoolSet(opts *bind.WatchOpts, sink chan<- *CoreOptimismPreferredPoolSet) (event.Subscription, error) {

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "PreferredPoolSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismPreferredPoolSet)
				if err := _CoreOptimism.contract.UnpackLog(event, "PreferredPoolSet", log); err != nil {
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

// ParsePreferredPoolSet is a log parse operation binding the contract event 0x7e7cb4726e710dc12fad41f158c37a4a71af3a6f053b8b13670d35c710139a56.
//
// Solidity: event PreferredPoolSet(uint256 poolId)
func (_CoreOptimism *CoreOptimismFilterer) ParsePreferredPoolSet(log types.Log) (*CoreOptimismPreferredPoolSet, error) {
	event := new(CoreOptimismPreferredPoolSet)
	if err := _CoreOptimism.contract.UnpackLog(event, "PreferredPoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the CoreOptimism contract.
type CoreOptimismRewardsClaimedIterator struct {
	Event *CoreOptimismRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *CoreOptimismRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismRewardsClaimed)
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
		it.Event = new(CoreOptimismRewardsClaimed)
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
func (it *CoreOptimismRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismRewardsClaimed represents a RewardsClaimed event raised by the CoreOptimism contract.
type CoreOptimismRewardsClaimed struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xa4a60be4203e7975e54ab5314c7e9e18aba9ad71e8da714d8de987f4f05410f2.
//
// Solidity: event RewardsClaimed(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount)
func (_CoreOptimism *CoreOptimismFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*CoreOptimismRewardsClaimedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "RewardsClaimed", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismRewardsClaimedIterator{contract: _CoreOptimism.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xa4a60be4203e7975e54ab5314c7e9e18aba9ad71e8da714d8de987f4f05410f2.
//
// Solidity: event RewardsClaimed(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount)
func (_CoreOptimism *CoreOptimismFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *CoreOptimismRewardsClaimed, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "RewardsClaimed", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismRewardsClaimed)
				if err := _CoreOptimism.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0xa4a60be4203e7975e54ab5314c7e9e18aba9ad71e8da714d8de987f4f05410f2.
//
// Solidity: event RewardsClaimed(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount)
func (_CoreOptimism *CoreOptimismFilterer) ParseRewardsClaimed(log types.Log) (*CoreOptimismRewardsClaimed, error) {
	event := new(CoreOptimismRewardsClaimed)
	if err := _CoreOptimism.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismRewardsDistributedIterator is returned from FilterRewardsDistributed and is used to iterate over the raw logs and unpacked data for RewardsDistributed events raised by the CoreOptimism contract.
type CoreOptimismRewardsDistributedIterator struct {
	Event *CoreOptimismRewardsDistributed // Event containing the contract specifics and raw log

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
func (it *CoreOptimismRewardsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismRewardsDistributed)
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
		it.Event = new(CoreOptimismRewardsDistributed)
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
func (it *CoreOptimismRewardsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismRewardsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismRewardsDistributed represents a RewardsDistributed event raised by the CoreOptimism contract.
type CoreOptimismRewardsDistributed struct {
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Amount         *big.Int
	Start          *big.Int
	Duration       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributed is a free log retrieval operation binding the contract event 0x19ced31d71d1db45f99d5a8d3a7616fe9d78828df58f2a28feb68c9f9ab876ca.
//
// Solidity: event RewardsDistributed(uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount, uint256 start, uint256 duration)
func (_CoreOptimism *CoreOptimismFilterer) FilterRewardsDistributed(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address) (*CoreOptimismRewardsDistributedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "RewardsDistributed", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismRewardsDistributedIterator{contract: _CoreOptimism.contract, event: "RewardsDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributed is a free log subscription operation binding the contract event 0x19ced31d71d1db45f99d5a8d3a7616fe9d78828df58f2a28feb68c9f9ab876ca.
//
// Solidity: event RewardsDistributed(uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount, uint256 start, uint256 duration)
func (_CoreOptimism *CoreOptimismFilterer) WatchRewardsDistributed(opts *bind.WatchOpts, sink chan<- *CoreOptimismRewardsDistributed, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "RewardsDistributed", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismRewardsDistributed)
				if err := _CoreOptimism.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
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

// ParseRewardsDistributed is a log parse operation binding the contract event 0x19ced31d71d1db45f99d5a8d3a7616fe9d78828df58f2a28feb68c9f9ab876ca.
//
// Solidity: event RewardsDistributed(uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount, uint256 start, uint256 duration)
func (_CoreOptimism *CoreOptimismFilterer) ParseRewardsDistributed(log types.Log) (*CoreOptimismRewardsDistributed, error) {
	event := new(CoreOptimismRewardsDistributed)
	if err := _CoreOptimism.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismRewardsDistributorRegisteredIterator is returned from FilterRewardsDistributorRegistered and is used to iterate over the raw logs and unpacked data for RewardsDistributorRegistered events raised by the CoreOptimism contract.
type CoreOptimismRewardsDistributorRegisteredIterator struct {
	Event *CoreOptimismRewardsDistributorRegistered // Event containing the contract specifics and raw log

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
func (it *CoreOptimismRewardsDistributorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismRewardsDistributorRegistered)
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
		it.Event = new(CoreOptimismRewardsDistributorRegistered)
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
func (it *CoreOptimismRewardsDistributorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismRewardsDistributorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismRewardsDistributorRegistered represents a RewardsDistributorRegistered event raised by the CoreOptimism contract.
type CoreOptimismRewardsDistributorRegistered struct {
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributorRegistered is a free log retrieval operation binding the contract event 0x9d3609c05a83dc93a5b355d62c2b37dfde8f0833b1184d4d05c6f51cd46b6e5b.
//
// Solidity: event RewardsDistributorRegistered(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_CoreOptimism *CoreOptimismFilterer) FilterRewardsDistributorRegistered(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (*CoreOptimismRewardsDistributorRegisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}
	var distributorRule []interface{}
	for _, distributorItem := range distributor {
		distributorRule = append(distributorRule, distributorItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "RewardsDistributorRegistered", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismRewardsDistributorRegisteredIterator{contract: _CoreOptimism.contract, event: "RewardsDistributorRegistered", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributorRegistered is a free log subscription operation binding the contract event 0x9d3609c05a83dc93a5b355d62c2b37dfde8f0833b1184d4d05c6f51cd46b6e5b.
//
// Solidity: event RewardsDistributorRegistered(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_CoreOptimism *CoreOptimismFilterer) WatchRewardsDistributorRegistered(opts *bind.WatchOpts, sink chan<- *CoreOptimismRewardsDistributorRegistered, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}
	var distributorRule []interface{}
	for _, distributorItem := range distributor {
		distributorRule = append(distributorRule, distributorItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "RewardsDistributorRegistered", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismRewardsDistributorRegistered)
				if err := _CoreOptimism.contract.UnpackLog(event, "RewardsDistributorRegistered", log); err != nil {
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

// ParseRewardsDistributorRegistered is a log parse operation binding the contract event 0x9d3609c05a83dc93a5b355d62c2b37dfde8f0833b1184d4d05c6f51cd46b6e5b.
//
// Solidity: event RewardsDistributorRegistered(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_CoreOptimism *CoreOptimismFilterer) ParseRewardsDistributorRegistered(log types.Log) (*CoreOptimismRewardsDistributorRegistered, error) {
	event := new(CoreOptimismRewardsDistributorRegistered)
	if err := _CoreOptimism.contract.UnpackLog(event, "RewardsDistributorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismRewardsDistributorRemovedIterator is returned from FilterRewardsDistributorRemoved and is used to iterate over the raw logs and unpacked data for RewardsDistributorRemoved events raised by the CoreOptimism contract.
type CoreOptimismRewardsDistributorRemovedIterator struct {
	Event *CoreOptimismRewardsDistributorRemoved // Event containing the contract specifics and raw log

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
func (it *CoreOptimismRewardsDistributorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismRewardsDistributorRemoved)
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
		it.Event = new(CoreOptimismRewardsDistributorRemoved)
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
func (it *CoreOptimismRewardsDistributorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismRewardsDistributorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismRewardsDistributorRemoved represents a RewardsDistributorRemoved event raised by the CoreOptimism contract.
type CoreOptimismRewardsDistributorRemoved struct {
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributorRemoved is a free log retrieval operation binding the contract event 0x375c4507f463c55a506be95e2cfd3cfdc0610be055087eac6049588a1bcfacba.
//
// Solidity: event RewardsDistributorRemoved(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_CoreOptimism *CoreOptimismFilterer) FilterRewardsDistributorRemoved(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (*CoreOptimismRewardsDistributorRemovedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}
	var distributorRule []interface{}
	for _, distributorItem := range distributor {
		distributorRule = append(distributorRule, distributorItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "RewardsDistributorRemoved", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismRewardsDistributorRemovedIterator{contract: _CoreOptimism.contract, event: "RewardsDistributorRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributorRemoved is a free log subscription operation binding the contract event 0x375c4507f463c55a506be95e2cfd3cfdc0610be055087eac6049588a1bcfacba.
//
// Solidity: event RewardsDistributorRemoved(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_CoreOptimism *CoreOptimismFilterer) WatchRewardsDistributorRemoved(opts *bind.WatchOpts, sink chan<- *CoreOptimismRewardsDistributorRemoved, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}
	var distributorRule []interface{}
	for _, distributorItem := range distributor {
		distributorRule = append(distributorRule, distributorItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "RewardsDistributorRemoved", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismRewardsDistributorRemoved)
				if err := _CoreOptimism.contract.UnpackLog(event, "RewardsDistributorRemoved", log); err != nil {
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

// ParseRewardsDistributorRemoved is a log parse operation binding the contract event 0x375c4507f463c55a506be95e2cfd3cfdc0610be055087eac6049588a1bcfacba.
//
// Solidity: event RewardsDistributorRemoved(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_CoreOptimism *CoreOptimismFilterer) ParseRewardsDistributorRemoved(log types.Log) (*CoreOptimismRewardsDistributorRemoved, error) {
	event := new(CoreOptimismRewardsDistributorRemoved)
	if err := _CoreOptimism.contract.UnpackLog(event, "RewardsDistributorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismSetMarketMinLiquidityRatioIterator is returned from FilterSetMarketMinLiquidityRatio and is used to iterate over the raw logs and unpacked data for SetMarketMinLiquidityRatio events raised by the CoreOptimism contract.
type CoreOptimismSetMarketMinLiquidityRatioIterator struct {
	Event *CoreOptimismSetMarketMinLiquidityRatio // Event containing the contract specifics and raw log

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
func (it *CoreOptimismSetMarketMinLiquidityRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismSetMarketMinLiquidityRatio)
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
		it.Event = new(CoreOptimismSetMarketMinLiquidityRatio)
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
func (it *CoreOptimismSetMarketMinLiquidityRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismSetMarketMinLiquidityRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismSetMarketMinLiquidityRatio represents a SetMarketMinLiquidityRatio event raised by the CoreOptimism contract.
type CoreOptimismSetMarketMinLiquidityRatio struct {
	MarketId          *big.Int
	MinLiquidityRatio *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMarketMinLiquidityRatio is a free log retrieval operation binding the contract event 0x563eb723f21b3e87ec8932cfb4ffa64d1b68c42053c28d6b4db019a40f6daf47.
//
// Solidity: event SetMarketMinLiquidityRatio(uint128 indexed marketId, uint256 minLiquidityRatio)
func (_CoreOptimism *CoreOptimismFilterer) FilterSetMarketMinLiquidityRatio(opts *bind.FilterOpts, marketId []*big.Int) (*CoreOptimismSetMarketMinLiquidityRatioIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "SetMarketMinLiquidityRatio", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismSetMarketMinLiquidityRatioIterator{contract: _CoreOptimism.contract, event: "SetMarketMinLiquidityRatio", logs: logs, sub: sub}, nil
}

// WatchSetMarketMinLiquidityRatio is a free log subscription operation binding the contract event 0x563eb723f21b3e87ec8932cfb4ffa64d1b68c42053c28d6b4db019a40f6daf47.
//
// Solidity: event SetMarketMinLiquidityRatio(uint128 indexed marketId, uint256 minLiquidityRatio)
func (_CoreOptimism *CoreOptimismFilterer) WatchSetMarketMinLiquidityRatio(opts *bind.WatchOpts, sink chan<- *CoreOptimismSetMarketMinLiquidityRatio, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "SetMarketMinLiquidityRatio", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismSetMarketMinLiquidityRatio)
				if err := _CoreOptimism.contract.UnpackLog(event, "SetMarketMinLiquidityRatio", log); err != nil {
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

// ParseSetMarketMinLiquidityRatio is a log parse operation binding the contract event 0x563eb723f21b3e87ec8932cfb4ffa64d1b68c42053c28d6b4db019a40f6daf47.
//
// Solidity: event SetMarketMinLiquidityRatio(uint128 indexed marketId, uint256 minLiquidityRatio)
func (_CoreOptimism *CoreOptimismFilterer) ParseSetMarketMinLiquidityRatio(log types.Log) (*CoreOptimismSetMarketMinLiquidityRatio, error) {
	event := new(CoreOptimismSetMarketMinLiquidityRatio)
	if err := _CoreOptimism.contract.UnpackLog(event, "SetMarketMinLiquidityRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismSetMinDelegateTimeIterator is returned from FilterSetMinDelegateTime and is used to iterate over the raw logs and unpacked data for SetMinDelegateTime events raised by the CoreOptimism contract.
type CoreOptimismSetMinDelegateTimeIterator struct {
	Event *CoreOptimismSetMinDelegateTime // Event containing the contract specifics and raw log

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
func (it *CoreOptimismSetMinDelegateTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismSetMinDelegateTime)
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
		it.Event = new(CoreOptimismSetMinDelegateTime)
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
func (it *CoreOptimismSetMinDelegateTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismSetMinDelegateTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismSetMinDelegateTime represents a SetMinDelegateTime event raised by the CoreOptimism contract.
type CoreOptimismSetMinDelegateTime struct {
	MarketId        *big.Int
	MinDelegateTime uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetMinDelegateTime is a free log retrieval operation binding the contract event 0x6942a68d151863c1fed3c0c4c5f3258af738218527147ac69290ab23ca7d26c6.
//
// Solidity: event SetMinDelegateTime(uint128 indexed marketId, uint32 minDelegateTime)
func (_CoreOptimism *CoreOptimismFilterer) FilterSetMinDelegateTime(opts *bind.FilterOpts, marketId []*big.Int) (*CoreOptimismSetMinDelegateTimeIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "SetMinDelegateTime", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismSetMinDelegateTimeIterator{contract: _CoreOptimism.contract, event: "SetMinDelegateTime", logs: logs, sub: sub}, nil
}

// WatchSetMinDelegateTime is a free log subscription operation binding the contract event 0x6942a68d151863c1fed3c0c4c5f3258af738218527147ac69290ab23ca7d26c6.
//
// Solidity: event SetMinDelegateTime(uint128 indexed marketId, uint32 minDelegateTime)
func (_CoreOptimism *CoreOptimismFilterer) WatchSetMinDelegateTime(opts *bind.WatchOpts, sink chan<- *CoreOptimismSetMinDelegateTime, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "SetMinDelegateTime", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismSetMinDelegateTime)
				if err := _CoreOptimism.contract.UnpackLog(event, "SetMinDelegateTime", log); err != nil {
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

// ParseSetMinDelegateTime is a log parse operation binding the contract event 0x6942a68d151863c1fed3c0c4c5f3258af738218527147ac69290ab23ca7d26c6.
//
// Solidity: event SetMinDelegateTime(uint128 indexed marketId, uint32 minDelegateTime)
func (_CoreOptimism *CoreOptimismFilterer) ParseSetMinDelegateTime(log types.Log) (*CoreOptimismSetMinDelegateTime, error) {
	event := new(CoreOptimismSetMinDelegateTime)
	if err := _CoreOptimism.contract.UnpackLog(event, "SetMinDelegateTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismSetMinLiquidityRatioIterator is returned from FilterSetMinLiquidityRatio and is used to iterate over the raw logs and unpacked data for SetMinLiquidityRatio events raised by the CoreOptimism contract.
type CoreOptimismSetMinLiquidityRatioIterator struct {
	Event *CoreOptimismSetMinLiquidityRatio // Event containing the contract specifics and raw log

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
func (it *CoreOptimismSetMinLiquidityRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismSetMinLiquidityRatio)
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
		it.Event = new(CoreOptimismSetMinLiquidityRatio)
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
func (it *CoreOptimismSetMinLiquidityRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismSetMinLiquidityRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismSetMinLiquidityRatio represents a SetMinLiquidityRatio event raised by the CoreOptimism contract.
type CoreOptimismSetMinLiquidityRatio struct {
	MinLiquidityRatio *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMinLiquidityRatio is a free log retrieval operation binding the contract event 0x66fd484d9868d1faddc8fef1f3faed0ed25eb4e6acde49dd1f2cbf0fba903635.
//
// Solidity: event SetMinLiquidityRatio(uint256 minLiquidityRatio)
func (_CoreOptimism *CoreOptimismFilterer) FilterSetMinLiquidityRatio(opts *bind.FilterOpts) (*CoreOptimismSetMinLiquidityRatioIterator, error) {

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "SetMinLiquidityRatio")
	if err != nil {
		return nil, err
	}
	return &CoreOptimismSetMinLiquidityRatioIterator{contract: _CoreOptimism.contract, event: "SetMinLiquidityRatio", logs: logs, sub: sub}, nil
}

// WatchSetMinLiquidityRatio is a free log subscription operation binding the contract event 0x66fd484d9868d1faddc8fef1f3faed0ed25eb4e6acde49dd1f2cbf0fba903635.
//
// Solidity: event SetMinLiquidityRatio(uint256 minLiquidityRatio)
func (_CoreOptimism *CoreOptimismFilterer) WatchSetMinLiquidityRatio(opts *bind.WatchOpts, sink chan<- *CoreOptimismSetMinLiquidityRatio) (event.Subscription, error) {

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "SetMinLiquidityRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismSetMinLiquidityRatio)
				if err := _CoreOptimism.contract.UnpackLog(event, "SetMinLiquidityRatio", log); err != nil {
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

// ParseSetMinLiquidityRatio is a log parse operation binding the contract event 0x66fd484d9868d1faddc8fef1f3faed0ed25eb4e6acde49dd1f2cbf0fba903635.
//
// Solidity: event SetMinLiquidityRatio(uint256 minLiquidityRatio)
func (_CoreOptimism *CoreOptimismFilterer) ParseSetMinLiquidityRatio(log types.Log) (*CoreOptimismSetMinLiquidityRatio, error) {
	event := new(CoreOptimismSetMinLiquidityRatio)
	if err := _CoreOptimism.contract.UnpackLog(event, "SetMinLiquidityRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismTransferCrossChainInitiatedIterator is returned from FilterTransferCrossChainInitiated and is used to iterate over the raw logs and unpacked data for TransferCrossChainInitiated events raised by the CoreOptimism contract.
type CoreOptimismTransferCrossChainInitiatedIterator struct {
	Event *CoreOptimismTransferCrossChainInitiated // Event containing the contract specifics and raw log

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
func (it *CoreOptimismTransferCrossChainInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismTransferCrossChainInitiated)
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
		it.Event = new(CoreOptimismTransferCrossChainInitiated)
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
func (it *CoreOptimismTransferCrossChainInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismTransferCrossChainInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismTransferCrossChainInitiated represents a TransferCrossChainInitiated event raised by the CoreOptimism contract.
type CoreOptimismTransferCrossChainInitiated struct {
	DestChainId uint64
	Amount      *big.Int
	Sender      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferCrossChainInitiated is a free log retrieval operation binding the contract event 0xb87c3097d7f9145a4915e8e434f04a1b7b91646d8a6e66a5cdab25caccb644c4.
//
// Solidity: event TransferCrossChainInitiated(uint64 indexed destChainId, uint256 indexed amount, address sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterTransferCrossChainInitiated(opts *bind.FilterOpts, destChainId []uint64, amount []*big.Int) (*CoreOptimismTransferCrossChainInitiatedIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "TransferCrossChainInitiated", destChainIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismTransferCrossChainInitiatedIterator{contract: _CoreOptimism.contract, event: "TransferCrossChainInitiated", logs: logs, sub: sub}, nil
}

// WatchTransferCrossChainInitiated is a free log subscription operation binding the contract event 0xb87c3097d7f9145a4915e8e434f04a1b7b91646d8a6e66a5cdab25caccb644c4.
//
// Solidity: event TransferCrossChainInitiated(uint64 indexed destChainId, uint256 indexed amount, address sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchTransferCrossChainInitiated(opts *bind.WatchOpts, sink chan<- *CoreOptimismTransferCrossChainInitiated, destChainId []uint64, amount []*big.Int) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "TransferCrossChainInitiated", destChainIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismTransferCrossChainInitiated)
				if err := _CoreOptimism.contract.UnpackLog(event, "TransferCrossChainInitiated", log); err != nil {
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

// ParseTransferCrossChainInitiated is a log parse operation binding the contract event 0xb87c3097d7f9145a4915e8e434f04a1b7b91646d8a6e66a5cdab25caccb644c4.
//
// Solidity: event TransferCrossChainInitiated(uint64 indexed destChainId, uint256 indexed amount, address sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseTransferCrossChainInitiated(log types.Log) (*CoreOptimismTransferCrossChainInitiated, error) {
	event := new(CoreOptimismTransferCrossChainInitiated)
	if err := _CoreOptimism.contract.UnpackLog(event, "TransferCrossChainInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CoreOptimism contract.
type CoreOptimismUpgradedIterator struct {
	Event *CoreOptimismUpgraded // Event containing the contract specifics and raw log

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
func (it *CoreOptimismUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismUpgraded)
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
		it.Event = new(CoreOptimismUpgraded)
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
func (it *CoreOptimismUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismUpgraded represents a Upgraded event raised by the CoreOptimism contract.
type CoreOptimismUpgraded struct {
	Self           common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_CoreOptimism *CoreOptimismFilterer) FilterUpgraded(opts *bind.FilterOpts, self []common.Address) (*CoreOptimismUpgradedIterator, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismUpgradedIterator{contract: _CoreOptimism.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_CoreOptimism *CoreOptimismFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CoreOptimismUpgraded, self []common.Address) (event.Subscription, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismUpgraded)
				if err := _CoreOptimism.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CoreOptimism *CoreOptimismFilterer) ParseUpgraded(log types.Log) (*CoreOptimismUpgraded, error) {
	event := new(CoreOptimismUpgraded)
	if err := _CoreOptimism.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismUsdBurnedIterator is returned from FilterUsdBurned and is used to iterate over the raw logs and unpacked data for UsdBurned events raised by the CoreOptimism contract.
type CoreOptimismUsdBurnedIterator struct {
	Event *CoreOptimismUsdBurned // Event containing the contract specifics and raw log

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
func (it *CoreOptimismUsdBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismUsdBurned)
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
		it.Event = new(CoreOptimismUsdBurned)
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
func (it *CoreOptimismUsdBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismUsdBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismUsdBurned represents a UsdBurned event raised by the CoreOptimism contract.
type CoreOptimismUsdBurned struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	Amount         *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUsdBurned is a free log retrieval operation binding the contract event 0x6b0230f0abe9188cbdbb1c816a4c5e213758b5b743d8a94af056280cc1e7aeb1.
//
// Solidity: event UsdBurned(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterUsdBurned(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*CoreOptimismUsdBurnedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "UsdBurned", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismUsdBurnedIterator{contract: _CoreOptimism.contract, event: "UsdBurned", logs: logs, sub: sub}, nil
}

// WatchUsdBurned is a free log subscription operation binding the contract event 0x6b0230f0abe9188cbdbb1c816a4c5e213758b5b743d8a94af056280cc1e7aeb1.
//
// Solidity: event UsdBurned(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchUsdBurned(opts *bind.WatchOpts, sink chan<- *CoreOptimismUsdBurned, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "UsdBurned", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismUsdBurned)
				if err := _CoreOptimism.contract.UnpackLog(event, "UsdBurned", log); err != nil {
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

// ParseUsdBurned is a log parse operation binding the contract event 0x6b0230f0abe9188cbdbb1c816a4c5e213758b5b743d8a94af056280cc1e7aeb1.
//
// Solidity: event UsdBurned(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseUsdBurned(log types.Log) (*CoreOptimismUsdBurned, error) {
	event := new(CoreOptimismUsdBurned)
	if err := _CoreOptimism.contract.UnpackLog(event, "UsdBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismUsdMintedIterator is returned from FilterUsdMinted and is used to iterate over the raw logs and unpacked data for UsdMinted events raised by the CoreOptimism contract.
type CoreOptimismUsdMintedIterator struct {
	Event *CoreOptimismUsdMinted // Event containing the contract specifics and raw log

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
func (it *CoreOptimismUsdMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismUsdMinted)
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
		it.Event = new(CoreOptimismUsdMinted)
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
func (it *CoreOptimismUsdMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismUsdMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismUsdMinted represents a UsdMinted event raised by the CoreOptimism contract.
type CoreOptimismUsdMinted struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	Amount         *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUsdMinted is a free log retrieval operation binding the contract event 0x2100f67dc9a5917400f799bb13194553e3f74c19a207c56350d2c223ac9c36c9.
//
// Solidity: event UsdMinted(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterUsdMinted(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*CoreOptimismUsdMintedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "UsdMinted", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismUsdMintedIterator{contract: _CoreOptimism.contract, event: "UsdMinted", logs: logs, sub: sub}, nil
}

// WatchUsdMinted is a free log subscription operation binding the contract event 0x2100f67dc9a5917400f799bb13194553e3f74c19a207c56350d2c223ac9c36c9.
//
// Solidity: event UsdMinted(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchUsdMinted(opts *bind.WatchOpts, sink chan<- *CoreOptimismUsdMinted, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "UsdMinted", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismUsdMinted)
				if err := _CoreOptimism.contract.UnpackLog(event, "UsdMinted", log); err != nil {
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

// ParseUsdMinted is a log parse operation binding the contract event 0x2100f67dc9a5917400f799bb13194553e3f74c19a207c56350d2c223ac9c36c9.
//
// Solidity: event UsdMinted(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseUsdMinted(log types.Log) (*CoreOptimismUsdMinted, error) {
	event := new(CoreOptimismUsdMinted)
	if err := _CoreOptimism.contract.UnpackLog(event, "UsdMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismVaultLiquidationIterator is returned from FilterVaultLiquidation and is used to iterate over the raw logs and unpacked data for VaultLiquidation events raised by the CoreOptimism contract.
type CoreOptimismVaultLiquidationIterator struct {
	Event *CoreOptimismVaultLiquidation // Event containing the contract specifics and raw log

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
func (it *CoreOptimismVaultLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismVaultLiquidation)
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
		it.Event = new(CoreOptimismVaultLiquidation)
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
func (it *CoreOptimismVaultLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismVaultLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismVaultLiquidation represents a VaultLiquidation event raised by the CoreOptimism contract.
type CoreOptimismVaultLiquidation struct {
	PoolId               *big.Int
	CollateralType       common.Address
	LiquidationData      ILiquidationModuleLiquidationData
	LiquidateAsAccountId *big.Int
	Sender               common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterVaultLiquidation is a free log retrieval operation binding the contract event 0x1834a7cc9d14f9bfa482df5c0404dadd1b8ec123b41f082e76ae28a3b2ea68d5.
//
// Solidity: event VaultLiquidation(uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterVaultLiquidation(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address) (*CoreOptimismVaultLiquidationIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "VaultLiquidation", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismVaultLiquidationIterator{contract: _CoreOptimism.contract, event: "VaultLiquidation", logs: logs, sub: sub}, nil
}

// WatchVaultLiquidation is a free log subscription operation binding the contract event 0x1834a7cc9d14f9bfa482df5c0404dadd1b8ec123b41f082e76ae28a3b2ea68d5.
//
// Solidity: event VaultLiquidation(uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchVaultLiquidation(opts *bind.WatchOpts, sink chan<- *CoreOptimismVaultLiquidation, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "VaultLiquidation", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismVaultLiquidation)
				if err := _CoreOptimism.contract.UnpackLog(event, "VaultLiquidation", log); err != nil {
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

// ParseVaultLiquidation is a log parse operation binding the contract event 0x1834a7cc9d14f9bfa482df5c0404dadd1b8ec123b41f082e76ae28a3b2ea68d5.
//
// Solidity: event VaultLiquidation(uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseVaultLiquidation(log types.Log) (*CoreOptimismVaultLiquidation, error) {
	event := new(CoreOptimismVaultLiquidation)
	if err := _CoreOptimism.contract.UnpackLog(event, "VaultLiquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOptimismWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the CoreOptimism contract.
type CoreOptimismWithdrawnIterator struct {
	Event *CoreOptimismWithdrawn // Event containing the contract specifics and raw log

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
func (it *CoreOptimismWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOptimismWithdrawn)
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
		it.Event = new(CoreOptimismWithdrawn)
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
func (it *CoreOptimismWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOptimismWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOptimismWithdrawn represents a Withdrawn event raised by the CoreOptimism contract.
type CoreOptimismWithdrawn struct {
	AccountId      *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x8b5f9d7ce522936589c630db08c0fa2405b21c4a5ff8ef19899900172736ba38.
//
// Solidity: event Withdrawn(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) FilterWithdrawn(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreOptimismWithdrawnIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.FilterLogs(opts, "Withdrawn", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreOptimismWithdrawnIterator{contract: _CoreOptimism.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x8b5f9d7ce522936589c630db08c0fa2405b21c4a5ff8ef19899900172736ba38.
//
// Solidity: event Withdrawn(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CoreOptimismWithdrawn, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreOptimism.contract.WatchLogs(opts, "Withdrawn", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOptimismWithdrawn)
				if err := _CoreOptimism.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x8b5f9d7ce522936589c630db08c0fa2405b21c4a5ff8ef19899900172736ba38.
//
// Solidity: event Withdrawn(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreOptimism *CoreOptimismFilterer) ParseWithdrawn(log types.Log) (*CoreOptimismWithdrawn, error) {
	event := new(CoreOptimismWithdrawn)
	if err := _CoreOptimism.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
