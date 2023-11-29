// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package core

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

// CoreMetaData contains all meta data concerning the Core contract.
var CoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"InvalidAccountId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"InvalidPermission\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"OnlyAccountTokenProxy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"PermissionNotGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionOutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedAccountId\",\"type\":\"uint128\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountLastInteraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountPermissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"permissions\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIAccountModule.AccountPermissions[]\",\"name\":\"accountPerms\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"notifyAccountTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"renouncePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyDistribution\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateralRatio\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"MarketNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"NotFundedByPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint128ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint128\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"updatedDebt\",\"type\":\"int256\"}],\"name\":\"DebtAssociated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"associateDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotCcipRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNetwork\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCcipClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCcipClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredTime\",\"type\":\"uint256\"}],\"name\":\"AccountActivityTimeoutPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"CollateralDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAvailableForDelegationD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountD18\",\"type\":\"uint256\"}],\"name\":\"InsufficentAvailableCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InsufficientAccountCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existing\",\"type\":\"uint256\"}],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parameter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"PrecisionLost\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"CollateralLockCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"CollateralLockExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"cleanExpiredLocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cleared\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"createLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getAccountAvailableCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getAccountCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDeposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAssigned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLocked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getLocks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amountD18\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"lockExpirationTime\",\"type\":\"uint64\"}],\"internalType\":\"structCollateralLock.Data[]\",\"name\":\"locks\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"CollateralConfigured\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"configureCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getCollateralConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"hideDisabled\",\"type\":\"bool\"}],\"name\":\"getCollateralConfigurations\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getCollateralPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientCcipFee\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TransferCrossChainInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCrossChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasTokenUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentDebt\",\"type\":\"int256\"}],\"name\":\"InsufficientDebt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"PoolNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"IssuanceFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UsdBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UsdMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CannotScaleEmptyMapping\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"currentCRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cratio\",\"type\":\"uint256\"}],\"name\":\"IneligibleForLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientMappedAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeVaultLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt128ToUint128\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Liquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"VaultLiquidation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"isPositionLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"isVaultLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"}],\"name\":\"liquidate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxUsd\",\"type\":\"uint256\"}],\"name\":\"liquidateVault\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountToDeposit\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarketCollateralDepositable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarketCollateralWithdrawable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"creditCapacity\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"netIssuance\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedCollateralValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportedDebt\",\"type\":\"uint256\"}],\"name\":\"MarketCollateralDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"creditCapacity\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"netIssuance\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedCollateralValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportedDebt\",\"type\":\"uint256\"}],\"name\":\"MarketCollateralWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"systemAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MaximumMarketCollateralConfigured\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"configureMaximumMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"depositMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getMarketCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralAmountD18\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getMaximumMarketCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"IncorrectMarketInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotEnoughLiquidity\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MarketRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"MarketSystemFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"creditCapacity\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"netIssuance\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedCollateralValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportedDebt\",\"type\":\"uint256\"}],\"name\":\"MarketUsdDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"creditCapacity\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"netIssuance\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedCollateralValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportedDebt\",\"type\":\"uint256\"}],\"name\":\"MarketUsdWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"SetMarketMinLiquidityRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"minDelegateTime\",\"type\":\"uint32\"}],\"name\":\"SetMinDelegateTime\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxIter\",\"type\":\"uint256\"}],\"name\":\"distributeDebtToPools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketDebtPerShare\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getMarketFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFeeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketMinDelegateTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketNetIssuance\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getMarketPoolDebtDistribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesD18\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"totalSharesD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"valuePerShareD27\",\"type\":\"int128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketPools\",\"outputs\":[{\"internalType\":\"uint128[]\",\"name\":\"inRangePoolIds\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"outRangePoolIds\",\"type\":\"uint128[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketReportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketTotalDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMinLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleManager\",\"outputs\":[{\"internalType\":\"contractIOracleManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsdToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getWithdrawableMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"isMarketCapacityLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"registerMarket\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"minDelegateTime\",\"type\":\"uint32\"}],\"name\":\"setMarketMinDelegateTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"setMinLiquidityRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolApprovedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolApprovedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PreferredPoolSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"addApprovedPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApprovedPools\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreferredPool\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"removeApprovedPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"setPreferredPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"CapacityLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"timeRemaining\",\"type\":\"uint32\"}],\"name\":\"MinDelegationTimeoutPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"PoolAlreadyExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collateralLimitD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPoolCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"PoolCollateralConfigurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"name\":\"PoolCollateralDisabledByDefaultSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"indexed\":false,\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"markets\",\"type\":\"tuple[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolConfigurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolNameUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolNominationRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolNominationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolOwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolOwnershipAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolOwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"SetMinLiquidityRatio\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"acceptPoolOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedPoolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getNominatedPoolOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPoolCollateralConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collateralLimitD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"}],\"internalType\":\"structPoolCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getPoolCollateralIssuanceRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"poolName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"nominatePoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"optionalCollateralType\",\"type\":\"address\"}],\"name\":\"rebalancePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"renouncePoolNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"renouncePoolOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"revokePoolNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"setMinLiquidityRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collateralLimitD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"}],\"internalType\":\"structPoolCollateralConfiguration.Data\",\"name\":\"newConfig\",\"type\":\"tuple\"}],\"name\":\"setPoolCollateralConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"name\":\"setPoolCollateralDisabledByDefault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"newMarketConfigurations\",\"type\":\"tuple[]\"}],\"name\":\"setPoolConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setPoolName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint32\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint32ToInt32\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint64ToInt64\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardDistributorNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardUnavailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RewardsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardsDistributorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardsDistributorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"getAvailableRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"getRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"registerRewardsDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"removeRewardsDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"updateRewards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newChainId\",\"type\":\"uint64\"}],\"name\":\"NewSupportedCrossChainNetwork\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ccipRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ccipTokenPool\",\"type\":\"address\"}],\"name\":\"configureChainlinkCrossChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleManagerAddress\",\"type\":\"address\"}],\"name\":\"configureOracleManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"v\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfigAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfigUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTrustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"v\",\"type\":\"bytes32\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"supportedNetworks\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"ccipSelectors\",\"type\":\"uint64[]\"}],\"name\":\"setSupportedCrossChainNetworks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numRegistered\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minDelegation\",\"type\":\"uint256\"}],\"name\":\"InsufficientDelegation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateralAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"}],\"name\":\"InvalidLeverage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCollateral\",\"type\":\"uint256\"}],\"name\":\"PoolCollateralLimitExceeded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"DelegationUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralAmountD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"}],\"name\":\"delegateCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"collateralizationRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionCollateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultCollateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CoreABI is the input ABI used to generate the binding from.
// Deprecated: Use CoreMetaData.ABI instead.
var CoreABI = CoreMetaData.ABI

// Core is an auto generated Go binding around an Ethereum contract.
type Core struct {
	CoreCaller     // Read-only binding to the contract
	CoreTransactor // Write-only binding to the contract
	CoreFilterer   // Log filterer for contract events
}

// CoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoreSession struct {
	Contract     *Core             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoreCallerSession struct {
	Contract *CoreCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoreTransactorSession struct {
	Contract     *CoreTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoreRaw struct {
	Contract *Core // Generic contract binding to access the raw methods on
}

// CoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoreCallerRaw struct {
	Contract *CoreCaller // Generic read-only contract binding to access the raw methods on
}

// CoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoreTransactorRaw struct {
	Contract *CoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCore creates a new instance of Core, bound to a specific deployed contract.
func NewCore(address common.Address, backend bind.ContractBackend) (*Core, error) {
	contract, err := bindCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Core{CoreCaller: CoreCaller{contract: contract}, CoreTransactor: CoreTransactor{contract: contract}, CoreFilterer: CoreFilterer{contract: contract}}, nil
}

// NewCoreCaller creates a new read-only instance of Core, bound to a specific deployed contract.
func NewCoreCaller(address common.Address, caller bind.ContractCaller) (*CoreCaller, error) {
	contract, err := bindCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoreCaller{contract: contract}, nil
}

// NewCoreTransactor creates a new write-only instance of Core, bound to a specific deployed contract.
func NewCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*CoreTransactor, error) {
	contract, err := bindCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoreTransactor{contract: contract}, nil
}

// NewCoreFilterer creates a new log filterer instance of Core, bound to a specific deployed contract.
func NewCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*CoreFilterer, error) {
	contract, err := bindCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoreFilterer{contract: contract}, nil
}

// bindCore binds a generic wrapper to an already deployed contract.
func bindCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Core *CoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Core.Contract.CoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Core *CoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.Contract.CoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Core *CoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Core.Contract.CoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Core *CoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Core.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Core *CoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Core *CoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Core.Contract.contract.Transact(opts, method, params...)
}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_Core *CoreCaller) GetAccountAvailableCollateral(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAccountAvailableCollateral", accountId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_Core *CoreSession) GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Core.Contract.GetAccountAvailableCollateral(&_Core.CallOpts, accountId, collateralType)
}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_Core *CoreCallerSession) GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Core.Contract.GetAccountAvailableCollateral(&_Core.CallOpts, accountId, collateralType)
}

// GetAccountCollateral is a free data retrieval call binding the contract method 0xef45148e.
//
// Solidity: function getAccountCollateral(uint128 accountId, address collateralType) view returns(uint256 totalDeposited, uint256 totalAssigned, uint256 totalLocked)
func (_Core *CoreCaller) GetAccountCollateral(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAccountCollateral", accountId, collateralType)

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
func (_Core *CoreSession) GetAccountCollateral(accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	return _Core.Contract.GetAccountCollateral(&_Core.CallOpts, accountId, collateralType)
}

// GetAccountCollateral is a free data retrieval call binding the contract method 0xef45148e.
//
// Solidity: function getAccountCollateral(uint128 accountId, address collateralType) view returns(uint256 totalDeposited, uint256 totalAssigned, uint256 totalLocked)
func (_Core *CoreCallerSession) GetAccountCollateral(accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	return _Core.Contract.GetAccountCollateral(&_Core.CallOpts, accountId, collateralType)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_Core *CoreCaller) GetAccountLastInteraction(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAccountLastInteraction", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_Core *CoreSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetAccountLastInteraction(&_Core.CallOpts, accountId)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_Core *CoreCallerSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetAccountLastInteraction(&_Core.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_Core *CoreCaller) GetAccountOwner(opts *bind.CallOpts, accountId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAccountOwner", accountId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_Core *CoreSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _Core.Contract.GetAccountOwner(&_Core.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_Core *CoreCallerSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _Core.Contract.GetAccountOwner(&_Core.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_Core *CoreCaller) GetAccountPermissions(opts *bind.CallOpts, accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAccountPermissions", accountId)

	if err != nil {
		return *new([]IAccountModuleAccountPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAccountModuleAccountPermissions)).(*[]IAccountModuleAccountPermissions)

	return out0, err

}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_Core *CoreSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _Core.Contract.GetAccountPermissions(&_Core.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_Core *CoreCallerSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _Core.Contract.GetAccountPermissions(&_Core.CallOpts, accountId)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_Core *CoreCaller) GetAccountTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAccountTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_Core *CoreSession) GetAccountTokenAddress() (common.Address, error) {
	return _Core.Contract.GetAccountTokenAddress(&_Core.CallOpts)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_Core *CoreCallerSession) GetAccountTokenAddress() (common.Address, error) {
	return _Core.Contract.GetAccountTokenAddress(&_Core.CallOpts)
}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_Core *CoreCaller) GetApprovedPools(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getApprovedPools")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_Core *CoreSession) GetApprovedPools() ([]*big.Int, error) {
	return _Core.Contract.GetApprovedPools(&_Core.CallOpts)
}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_Core *CoreCallerSession) GetApprovedPools() ([]*big.Int, error) {
	return _Core.Contract.GetApprovedPools(&_Core.CallOpts)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_Core *CoreCaller) GetAssociatedSystem(opts *bind.CallOpts, id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAssociatedSystem", id)

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
func (_Core *CoreSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _Core.Contract.GetAssociatedSystem(&_Core.CallOpts, id)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_Core *CoreCallerSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _Core.Contract.GetAssociatedSystem(&_Core.CallOpts, id)
}

// GetAvailableRewards is a free data retrieval call binding the contract method 0xc4b3410e.
//
// Solidity: function getAvailableRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_Core *CoreCaller) GetAvailableRewards(opts *bind.CallOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAvailableRewards", accountId, poolId, collateralType, distributor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableRewards is a free data retrieval call binding the contract method 0xc4b3410e.
//
// Solidity: function getAvailableRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_Core *CoreSession) GetAvailableRewards(accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	return _Core.Contract.GetAvailableRewards(&_Core.CallOpts, accountId, poolId, collateralType, distributor)
}

// GetAvailableRewards is a free data retrieval call binding the contract method 0xc4b3410e.
//
// Solidity: function getAvailableRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_Core *CoreCallerSession) GetAvailableRewards(accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	return _Core.Contract.GetAvailableRewards(&_Core.CallOpts, accountId, poolId, collateralType, distributor)
}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) pure returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_Core *CoreCaller) GetCollateralConfiguration(opts *bind.CallOpts, collateralType common.Address) (CollateralConfigurationData, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getCollateralConfiguration", collateralType)

	if err != nil {
		return *new(CollateralConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new(CollateralConfigurationData)).(*CollateralConfigurationData)

	return out0, err

}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) pure returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_Core *CoreSession) GetCollateralConfiguration(collateralType common.Address) (CollateralConfigurationData, error) {
	return _Core.Contract.GetCollateralConfiguration(&_Core.CallOpts, collateralType)
}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) pure returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_Core *CoreCallerSession) GetCollateralConfiguration(collateralType common.Address) (CollateralConfigurationData, error) {
	return _Core.Contract.GetCollateralConfiguration(&_Core.CallOpts, collateralType)
}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_Core *CoreCaller) GetCollateralConfigurations(opts *bind.CallOpts, hideDisabled bool) ([]CollateralConfigurationData, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getCollateralConfigurations", hideDisabled)

	if err != nil {
		return *new([]CollateralConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CollateralConfigurationData)).(*[]CollateralConfigurationData)

	return out0, err

}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_Core *CoreSession) GetCollateralConfigurations(hideDisabled bool) ([]CollateralConfigurationData, error) {
	return _Core.Contract.GetCollateralConfigurations(&_Core.CallOpts, hideDisabled)
}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_Core *CoreCallerSession) GetCollateralConfigurations(hideDisabled bool) ([]CollateralConfigurationData, error) {
	return _Core.Contract.GetCollateralConfigurations(&_Core.CallOpts, hideDisabled)
}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_Core *CoreCaller) GetCollateralPrice(opts *bind.CallOpts, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getCollateralPrice", collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_Core *CoreSession) GetCollateralPrice(collateralType common.Address) (*big.Int, error) {
	return _Core.Contract.GetCollateralPrice(&_Core.CallOpts, collateralType)
}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_Core *CoreCallerSession) GetCollateralPrice(collateralType common.Address) (*big.Int, error) {
	return _Core.Contract.GetCollateralPrice(&_Core.CallOpts, collateralType)
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_Core *CoreCaller) GetConfig(opts *bind.CallOpts, k [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getConfig", k)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_Core *CoreSession) GetConfig(k [32]byte) ([32]byte, error) {
	return _Core.Contract.GetConfig(&_Core.CallOpts, k)
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_Core *CoreCallerSession) GetConfig(k [32]byte) ([32]byte, error) {
	return _Core.Contract.GetConfig(&_Core.CallOpts, k)
}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_Core *CoreCaller) GetConfigAddress(opts *bind.CallOpts, k [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getConfigAddress", k)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_Core *CoreSession) GetConfigAddress(k [32]byte) (common.Address, error) {
	return _Core.Contract.GetConfigAddress(&_Core.CallOpts, k)
}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_Core *CoreCallerSession) GetConfigAddress(k [32]byte) (common.Address, error) {
	return _Core.Contract.GetConfigAddress(&_Core.CallOpts, k)
}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_Core *CoreCaller) GetConfigUint(opts *bind.CallOpts, k [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getConfigUint", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_Core *CoreSession) GetConfigUint(k [32]byte) (*big.Int, error) {
	return _Core.Contract.GetConfigUint(&_Core.CallOpts, k)
}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_Core *CoreCallerSession) GetConfigUint(k [32]byte) (*big.Int, error) {
	return _Core.Contract.GetConfigUint(&_Core.CallOpts, k)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_Core *CoreCaller) GetDeniers(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getDeniers", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_Core *CoreSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _Core.Contract.GetDeniers(&_Core.CallOpts, feature)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_Core *CoreCallerSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _Core.Contract.GetDeniers(&_Core.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_Core *CoreCaller) GetFeatureFlagAllowAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getFeatureFlagAllowAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_Core *CoreSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _Core.Contract.GetFeatureFlagAllowAll(&_Core.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_Core *CoreCallerSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _Core.Contract.GetFeatureFlagAllowAll(&_Core.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_Core *CoreCaller) GetFeatureFlagAllowlist(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getFeatureFlagAllowlist", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_Core *CoreSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _Core.Contract.GetFeatureFlagAllowlist(&_Core.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_Core *CoreCallerSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _Core.Contract.GetFeatureFlagAllowlist(&_Core.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_Core *CoreCaller) GetFeatureFlagDenyAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getFeatureFlagDenyAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_Core *CoreSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _Core.Contract.GetFeatureFlagDenyAll(&_Core.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_Core *CoreCallerSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _Core.Contract.GetFeatureFlagDenyAll(&_Core.CallOpts, feature)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_Core *CoreCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_Core *CoreSession) GetImplementation() (common.Address, error) {
	return _Core.Contract.GetImplementation(&_Core.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_Core *CoreCallerSession) GetImplementation() (common.Address, error) {
	return _Core.Contract.GetImplementation(&_Core.CallOpts)
}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_Core *CoreCaller) GetLocks(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getLocks", accountId, collateralType, offset, count)

	if err != nil {
		return *new([]CollateralLockData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CollateralLockData)).(*[]CollateralLockData)

	return out0, err

}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_Core *CoreSession) GetLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	return _Core.Contract.GetLocks(&_Core.CallOpts, accountId, collateralType, offset, count)
}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_Core *CoreCallerSession) GetLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	return _Core.Contract.GetLocks(&_Core.CallOpts, accountId, collateralType, offset, count)
}

// GetMarketAddress is a free data retrieval call binding the contract method 0xd24437f1.
//
// Solidity: function getMarketAddress(uint128 marketId) view returns(address)
func (_Core *CoreCaller) GetMarketAddress(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMarketAddress", marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMarketAddress is a free data retrieval call binding the contract method 0xd24437f1.
//
// Solidity: function getMarketAddress(uint128 marketId) view returns(address)
func (_Core *CoreSession) GetMarketAddress(marketId *big.Int) (common.Address, error) {
	return _Core.Contract.GetMarketAddress(&_Core.CallOpts, marketId)
}

// GetMarketAddress is a free data retrieval call binding the contract method 0xd24437f1.
//
// Solidity: function getMarketAddress(uint128 marketId) view returns(address)
func (_Core *CoreCallerSession) GetMarketAddress(marketId *big.Int) (common.Address, error) {
	return _Core.Contract.GetMarketAddress(&_Core.CallOpts, marketId)
}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_Core *CoreCaller) GetMarketCollateral(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMarketCollateral", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_Core *CoreSession) GetMarketCollateral(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMarketCollateral(&_Core.CallOpts, marketId)
}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_Core *CoreCallerSession) GetMarketCollateral(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMarketCollateral(&_Core.CallOpts, marketId)
}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_Core *CoreCaller) GetMarketCollateralAmount(opts *bind.CallOpts, marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMarketCollateralAmount", marketId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_Core *CoreSession) GetMarketCollateralAmount(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Core.Contract.GetMarketCollateralAmount(&_Core.CallOpts, marketId, collateralType)
}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_Core *CoreCallerSession) GetMarketCollateralAmount(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Core.Contract.GetMarketCollateralAmount(&_Core.CallOpts, marketId, collateralType)
}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_Core *CoreCaller) GetMarketCollateralValue(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMarketCollateralValue", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_Core *CoreSession) GetMarketCollateralValue(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMarketCollateralValue(&_Core.CallOpts, marketId)
}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_Core *CoreCallerSession) GetMarketCollateralValue(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMarketCollateralValue(&_Core.CallOpts, marketId)
}

// GetMarketFees is a free data retrieval call binding the contract method 0xdfb83437.
//
// Solidity: function getMarketFees(uint128 , uint256 amount) view returns(uint256 depositFeeAmount, uint256 withdrawFeeAmount)
func (_Core *CoreCaller) GetMarketFees(opts *bind.CallOpts, arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMarketFees", arg0, amount)

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
func (_Core *CoreSession) GetMarketFees(arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	return _Core.Contract.GetMarketFees(&_Core.CallOpts, arg0, amount)
}

// GetMarketFees is a free data retrieval call binding the contract method 0xdfb83437.
//
// Solidity: function getMarketFees(uint128 , uint256 amount) view returns(uint256 depositFeeAmount, uint256 withdrawFeeAmount)
func (_Core *CoreCallerSession) GetMarketFees(arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	return _Core.Contract.GetMarketFees(&_Core.CallOpts, arg0, amount)
}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_Core *CoreCaller) GetMarketMinDelegateTime(opts *bind.CallOpts, marketId *big.Int) (uint32, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMarketMinDelegateTime", marketId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_Core *CoreSession) GetMarketMinDelegateTime(marketId *big.Int) (uint32, error) {
	return _Core.Contract.GetMarketMinDelegateTime(&_Core.CallOpts, marketId)
}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_Core *CoreCallerSession) GetMarketMinDelegateTime(marketId *big.Int) (uint32, error) {
	return _Core.Contract.GetMarketMinDelegateTime(&_Core.CallOpts, marketId)
}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_Core *CoreCaller) GetMarketNetIssuance(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMarketNetIssuance", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_Core *CoreSession) GetMarketNetIssuance(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMarketNetIssuance(&_Core.CallOpts, marketId)
}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_Core *CoreCallerSession) GetMarketNetIssuance(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMarketNetIssuance(&_Core.CallOpts, marketId)
}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_Core *CoreCaller) GetMarketReportedDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMarketReportedDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_Core *CoreSession) GetMarketReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMarketReportedDebt(&_Core.CallOpts, marketId)
}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_Core *CoreCallerSession) GetMarketReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMarketReportedDebt(&_Core.CallOpts, marketId)
}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_Core *CoreCaller) GetMarketTotalDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMarketTotalDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_Core *CoreSession) GetMarketTotalDebt(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMarketTotalDebt(&_Core.CallOpts, marketId)
}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_Core *CoreCallerSession) GetMarketTotalDebt(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMarketTotalDebt(&_Core.CallOpts, marketId)
}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_Core *CoreCaller) GetMaximumMarketCollateral(opts *bind.CallOpts, marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMaximumMarketCollateral", marketId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_Core *CoreSession) GetMaximumMarketCollateral(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Core.Contract.GetMaximumMarketCollateral(&_Core.CallOpts, marketId, collateralType)
}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_Core *CoreCallerSession) GetMaximumMarketCollateral(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Core.Contract.GetMaximumMarketCollateral(&_Core.CallOpts, marketId, collateralType)
}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_Core *CoreCaller) GetMinLiquidityRatio(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMinLiquidityRatio", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_Core *CoreSession) GetMinLiquidityRatio(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMinLiquidityRatio(&_Core.CallOpts, marketId)
}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_Core *CoreCallerSession) GetMinLiquidityRatio(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetMinLiquidityRatio(&_Core.CallOpts, marketId)
}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_Core *CoreCaller) GetMinLiquidityRatio0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getMinLiquidityRatio0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_Core *CoreSession) GetMinLiquidityRatio0() (*big.Int, error) {
	return _Core.Contract.GetMinLiquidityRatio0(&_Core.CallOpts)
}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_Core *CoreCallerSession) GetMinLiquidityRatio0() (*big.Int, error) {
	return _Core.Contract.GetMinLiquidityRatio0(&_Core.CallOpts)
}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_Core *CoreCaller) GetNominatedPoolOwner(opts *bind.CallOpts, poolId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getNominatedPoolOwner", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_Core *CoreSession) GetNominatedPoolOwner(poolId *big.Int) (common.Address, error) {
	return _Core.Contract.GetNominatedPoolOwner(&_Core.CallOpts, poolId)
}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_Core *CoreCallerSession) GetNominatedPoolOwner(poolId *big.Int) (common.Address, error) {
	return _Core.Contract.GetNominatedPoolOwner(&_Core.CallOpts, poolId)
}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_Core *CoreCaller) GetOracleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getOracleManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_Core *CoreSession) GetOracleManager() (common.Address, error) {
	return _Core.Contract.GetOracleManager(&_Core.CallOpts)
}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_Core *CoreCallerSession) GetOracleManager() (common.Address, error) {
	return _Core.Contract.GetOracleManager(&_Core.CallOpts)
}

// GetPoolCollateralConfiguration is a free data retrieval call binding the contract method 0xc77e51f6.
//
// Solidity: function getPoolCollateralConfiguration(uint128 poolId, address collateralType) view returns((uint256,uint256) config)
func (_Core *CoreCaller) GetPoolCollateralConfiguration(opts *bind.CallOpts, poolId *big.Int, collateralType common.Address) (PoolCollateralConfigurationData, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getPoolCollateralConfiguration", poolId, collateralType)

	if err != nil {
		return *new(PoolCollateralConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolCollateralConfigurationData)).(*PoolCollateralConfigurationData)

	return out0, err

}

// GetPoolCollateralConfiguration is a free data retrieval call binding the contract method 0xc77e51f6.
//
// Solidity: function getPoolCollateralConfiguration(uint128 poolId, address collateralType) view returns((uint256,uint256) config)
func (_Core *CoreSession) GetPoolCollateralConfiguration(poolId *big.Int, collateralType common.Address) (PoolCollateralConfigurationData, error) {
	return _Core.Contract.GetPoolCollateralConfiguration(&_Core.CallOpts, poolId, collateralType)
}

// GetPoolCollateralConfiguration is a free data retrieval call binding the contract method 0xc77e51f6.
//
// Solidity: function getPoolCollateralConfiguration(uint128 poolId, address collateralType) view returns((uint256,uint256) config)
func (_Core *CoreCallerSession) GetPoolCollateralConfiguration(poolId *big.Int, collateralType common.Address) (PoolCollateralConfigurationData, error) {
	return _Core.Contract.GetPoolCollateralConfiguration(&_Core.CallOpts, poolId, collateralType)
}

// GetPoolCollateralIssuanceRatio is a free data retrieval call binding the contract method 0xc4d2aad3.
//
// Solidity: function getPoolCollateralIssuanceRatio(uint128 poolId, address collateral) view returns(uint256)
func (_Core *CoreCaller) GetPoolCollateralIssuanceRatio(opts *bind.CallOpts, poolId *big.Int, collateral common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getPoolCollateralIssuanceRatio", poolId, collateral)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolCollateralIssuanceRatio is a free data retrieval call binding the contract method 0xc4d2aad3.
//
// Solidity: function getPoolCollateralIssuanceRatio(uint128 poolId, address collateral) view returns(uint256)
func (_Core *CoreSession) GetPoolCollateralIssuanceRatio(poolId *big.Int, collateral common.Address) (*big.Int, error) {
	return _Core.Contract.GetPoolCollateralIssuanceRatio(&_Core.CallOpts, poolId, collateral)
}

// GetPoolCollateralIssuanceRatio is a free data retrieval call binding the contract method 0xc4d2aad3.
//
// Solidity: function getPoolCollateralIssuanceRatio(uint128 poolId, address collateral) view returns(uint256)
func (_Core *CoreCallerSession) GetPoolCollateralIssuanceRatio(poolId *big.Int, collateral common.Address) (*big.Int, error) {
	return _Core.Contract.GetPoolCollateralIssuanceRatio(&_Core.CallOpts, poolId, collateral)
}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_Core *CoreCaller) GetPoolConfiguration(opts *bind.CallOpts, poolId *big.Int) ([]MarketConfigurationData, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getPoolConfiguration", poolId)

	if err != nil {
		return *new([]MarketConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketConfigurationData)).(*[]MarketConfigurationData)

	return out0, err

}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_Core *CoreSession) GetPoolConfiguration(poolId *big.Int) ([]MarketConfigurationData, error) {
	return _Core.Contract.GetPoolConfiguration(&_Core.CallOpts, poolId)
}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_Core *CoreCallerSession) GetPoolConfiguration(poolId *big.Int) ([]MarketConfigurationData, error) {
	return _Core.Contract.GetPoolConfiguration(&_Core.CallOpts, poolId)
}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_Core *CoreCaller) GetPoolName(opts *bind.CallOpts, poolId *big.Int) (string, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getPoolName", poolId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_Core *CoreSession) GetPoolName(poolId *big.Int) (string, error) {
	return _Core.Contract.GetPoolName(&_Core.CallOpts, poolId)
}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_Core *CoreCallerSession) GetPoolName(poolId *big.Int) (string, error) {
	return _Core.Contract.GetPoolName(&_Core.CallOpts, poolId)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_Core *CoreCaller) GetPoolOwner(opts *bind.CallOpts, poolId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getPoolOwner", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_Core *CoreSession) GetPoolOwner(poolId *big.Int) (common.Address, error) {
	return _Core.Contract.GetPoolOwner(&_Core.CallOpts, poolId)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_Core *CoreCallerSession) GetPoolOwner(poolId *big.Int) (common.Address, error) {
	return _Core.Contract.GetPoolOwner(&_Core.CallOpts, poolId)
}

// GetPositionCollateral is a free data retrieval call binding the contract method 0x33cc422b.
//
// Solidity: function getPositionCollateral(uint128 accountId, uint128 poolId, address collateralType) view returns(uint256 amount)
func (_Core *CoreCaller) GetPositionCollateral(opts *bind.CallOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getPositionCollateral", accountId, poolId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionCollateral is a free data retrieval call binding the contract method 0x33cc422b.
//
// Solidity: function getPositionCollateral(uint128 accountId, uint128 poolId, address collateralType) view returns(uint256 amount)
func (_Core *CoreSession) GetPositionCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Core.Contract.GetPositionCollateral(&_Core.CallOpts, accountId, poolId, collateralType)
}

// GetPositionCollateral is a free data retrieval call binding the contract method 0x33cc422b.
//
// Solidity: function getPositionCollateral(uint128 accountId, uint128 poolId, address collateralType) view returns(uint256 amount)
func (_Core *CoreCallerSession) GetPositionCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Core.Contract.GetPositionCollateral(&_Core.CallOpts, accountId, poolId, collateralType)
}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_Core *CoreCaller) GetPreferredPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getPreferredPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_Core *CoreSession) GetPreferredPool() (*big.Int, error) {
	return _Core.Contract.GetPreferredPool(&_Core.CallOpts)
}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_Core *CoreCallerSession) GetPreferredPool() (*big.Int, error) {
	return _Core.Contract.GetPreferredPool(&_Core.CallOpts)
}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_Core *CoreCaller) GetRewardRate(opts *bind.CallOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getRewardRate", poolId, collateralType, distributor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_Core *CoreSession) GetRewardRate(poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	return _Core.Contract.GetRewardRate(&_Core.CallOpts, poolId, collateralType, distributor)
}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_Core *CoreCallerSession) GetRewardRate(poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	return _Core.Contract.GetRewardRate(&_Core.CallOpts, poolId, collateralType, distributor)
}

// GetTrustedForwarder is a free data retrieval call binding the contract method 0xce1b815f.
//
// Solidity: function getTrustedForwarder() pure returns(address)
func (_Core *CoreCaller) GetTrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getTrustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTrustedForwarder is a free data retrieval call binding the contract method 0xce1b815f.
//
// Solidity: function getTrustedForwarder() pure returns(address)
func (_Core *CoreSession) GetTrustedForwarder() (common.Address, error) {
	return _Core.Contract.GetTrustedForwarder(&_Core.CallOpts)
}

// GetTrustedForwarder is a free data retrieval call binding the contract method 0xce1b815f.
//
// Solidity: function getTrustedForwarder() pure returns(address)
func (_Core *CoreCallerSession) GetTrustedForwarder() (common.Address, error) {
	return _Core.Contract.GetTrustedForwarder(&_Core.CallOpts)
}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_Core *CoreCaller) GetUsdToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getUsdToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_Core *CoreSession) GetUsdToken() (common.Address, error) {
	return _Core.Contract.GetUsdToken(&_Core.CallOpts)
}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_Core *CoreCallerSession) GetUsdToken() (common.Address, error) {
	return _Core.Contract.GetUsdToken(&_Core.CallOpts)
}

// GetVaultCollateral is a free data retrieval call binding the contract method 0x078145a8.
//
// Solidity: function getVaultCollateral(uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_Core *CoreCaller) GetVaultCollateral(opts *bind.CallOpts, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getVaultCollateral", poolId, collateralType)

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
func (_Core *CoreSession) GetVaultCollateral(poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _Core.Contract.GetVaultCollateral(&_Core.CallOpts, poolId, collateralType)
}

// GetVaultCollateral is a free data retrieval call binding the contract method 0x078145a8.
//
// Solidity: function getVaultCollateral(uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_Core *CoreCallerSession) GetVaultCollateral(poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _Core.Contract.GetVaultCollateral(&_Core.CallOpts, poolId, collateralType)
}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_Core *CoreCaller) GetWithdrawableMarketUsd(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getWithdrawableMarketUsd", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_Core *CoreSession) GetWithdrawableMarketUsd(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetWithdrawableMarketUsd(&_Core.CallOpts, marketId)
}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_Core *CoreCallerSession) GetWithdrawableMarketUsd(marketId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetWithdrawableMarketUsd(&_Core.CallOpts, marketId)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Core *CoreCaller) HasPermission(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "hasPermission", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Core *CoreSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _Core.Contract.HasPermission(&_Core.CallOpts, accountId, permission, user)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Core *CoreCallerSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _Core.Contract.HasPermission(&_Core.CallOpts, accountId, permission, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Core *CoreCaller) IsAuthorized(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "isAuthorized", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Core *CoreSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _Core.Contract.IsAuthorized(&_Core.CallOpts, accountId, permission, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_Core *CoreCallerSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _Core.Contract.IsAuthorized(&_Core.CallOpts, accountId, permission, user)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_Core *CoreCaller) IsFeatureAllowed(opts *bind.CallOpts, feature [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "isFeatureAllowed", feature, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_Core *CoreSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _Core.Contract.IsFeatureAllowed(&_Core.CallOpts, feature, account)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_Core *CoreCallerSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _Core.Contract.IsFeatureAllowed(&_Core.CallOpts, feature, account)
}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_Core *CoreCaller) IsMarketCapacityLocked(opts *bind.CallOpts, marketId *big.Int) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "isMarketCapacityLocked", marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_Core *CoreSession) IsMarketCapacityLocked(marketId *big.Int) (bool, error) {
	return _Core.Contract.IsMarketCapacityLocked(&_Core.CallOpts, marketId)
}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_Core *CoreCallerSession) IsMarketCapacityLocked(marketId *big.Int) (bool, error) {
	return _Core.Contract.IsMarketCapacityLocked(&_Core.CallOpts, marketId)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) pure returns(bool)
func (_Core *CoreCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) pure returns(bool)
func (_Core *CoreSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Core.Contract.IsTrustedForwarder(&_Core.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) pure returns(bool)
func (_Core *CoreCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Core.Contract.IsTrustedForwarder(&_Core.CallOpts, forwarder)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Core *CoreCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Core *CoreSession) NominatedOwner() (common.Address, error) {
	return _Core.Contract.NominatedOwner(&_Core.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Core *CoreCallerSession) NominatedOwner() (common.Address, error) {
	return _Core.Contract.NominatedOwner(&_Core.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Core *CoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Core *CoreSession) Owner() (common.Address, error) {
	return _Core.Contract.Owner(&_Core.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Core *CoreCallerSession) Owner() (common.Address, error) {
	return _Core.Contract.Owner(&_Core.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Core *CoreCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Core *CoreSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Core.Contract.SupportsInterface(&_Core.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Core *CoreCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Core.Contract.SupportsInterface(&_Core.CallOpts, interfaceId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Core *CoreTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Core *CoreSession) AcceptOwnership() (*types.Transaction, error) {
	return _Core.Contract.AcceptOwnership(&_Core.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Core *CoreTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Core.Contract.AcceptOwnership(&_Core.TransactOpts)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_Core *CoreTransactor) AcceptPoolOwnership(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "acceptPoolOwnership", poolId)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_Core *CoreSession) AcceptPoolOwnership(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.AcceptPoolOwnership(&_Core.TransactOpts, poolId)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_Core *CoreTransactorSession) AcceptPoolOwnership(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.AcceptPoolOwnership(&_Core.TransactOpts, poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_Core *CoreTransactor) AddApprovedPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "addApprovedPool", poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_Core *CoreSession) AddApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.AddApprovedPool(&_Core.TransactOpts, poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_Core *CoreTransactorSession) AddApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.AddApprovedPool(&_Core.TransactOpts, poolId)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Core *CoreTransactor) AddToFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "addToFeatureFlagAllowlist", feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Core *CoreSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Core.Contract.AddToFeatureFlagAllowlist(&_Core.TransactOpts, feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Core *CoreTransactorSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Core.Contract.AddToFeatureFlagAllowlist(&_Core.TransactOpts, feature, account)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_Core *CoreTransactor) AssociateDebt(opts *bind.TransactOpts, marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "associateDebt", marketId, poolId, collateralType, accountId, amount)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_Core *CoreSession) AssociateDebt(marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.AssociateDebt(&_Core.TransactOpts, marketId, poolId, collateralType, accountId, amount)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_Core *CoreTransactorSession) AssociateDebt(marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.AssociateDebt(&_Core.TransactOpts, marketId, poolId, collateralType, accountId, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Core *CoreTransactor) BurnUsd(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "burnUsd", accountId, poolId, collateralType, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Core *CoreSession) BurnUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.BurnUsd(&_Core.TransactOpts, accountId, poolId, collateralType, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Core *CoreTransactorSession) BurnUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.BurnUsd(&_Core.TransactOpts, accountId, poolId, collateralType, amount)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Core *CoreTransactor) CcipReceive(opts *bind.TransactOpts, message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Core *CoreSession) CcipReceive(message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _Core.Contract.CcipReceive(&_Core.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Core *CoreTransactorSession) CcipReceive(message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _Core.Contract.CcipReceive(&_Core.TransactOpts, message)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_Core *CoreTransactor) ClaimRewards(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "claimRewards", accountId, poolId, collateralType, distributor)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_Core *CoreSession) ClaimRewards(accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Core.Contract.ClaimRewards(&_Core.TransactOpts, accountId, poolId, collateralType, distributor)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_Core *CoreTransactorSession) ClaimRewards(accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Core.Contract.ClaimRewards(&_Core.TransactOpts, accountId, poolId, collateralType, distributor)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_Core *CoreTransactor) CleanExpiredLocks(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "cleanExpiredLocks", accountId, collateralType, offset, count)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_Core *CoreSession) CleanExpiredLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Core.Contract.CleanExpiredLocks(&_Core.TransactOpts, accountId, collateralType, offset, count)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_Core *CoreTransactorSession) CleanExpiredLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Core.Contract.CleanExpiredLocks(&_Core.TransactOpts, accountId, collateralType, offset, count)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_Core *CoreTransactor) ConfigureChainlinkCrossChain(opts *bind.TransactOpts, ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "configureChainlinkCrossChain", ccipRouter, ccipTokenPool)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_Core *CoreSession) ConfigureChainlinkCrossChain(ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _Core.Contract.ConfigureChainlinkCrossChain(&_Core.TransactOpts, ccipRouter, ccipTokenPool)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_Core *CoreTransactorSession) ConfigureChainlinkCrossChain(ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _Core.Contract.ConfigureChainlinkCrossChain(&_Core.TransactOpts, ccipRouter, ccipTokenPool)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_Core *CoreTransactor) ConfigureCollateral(opts *bind.TransactOpts, config CollateralConfigurationData) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "configureCollateral", config)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_Core *CoreSession) ConfigureCollateral(config CollateralConfigurationData) (*types.Transaction, error) {
	return _Core.Contract.ConfigureCollateral(&_Core.TransactOpts, config)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_Core *CoreTransactorSession) ConfigureCollateral(config CollateralConfigurationData) (*types.Transaction, error) {
	return _Core.Contract.ConfigureCollateral(&_Core.TransactOpts, config)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_Core *CoreTransactor) ConfigureMaximumMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "configureMaximumMarketCollateral", marketId, collateralType, amount)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_Core *CoreSession) ConfigureMaximumMarketCollateral(marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.ConfigureMaximumMarketCollateral(&_Core.TransactOpts, marketId, collateralType, amount)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_Core *CoreTransactorSession) ConfigureMaximumMarketCollateral(marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.ConfigureMaximumMarketCollateral(&_Core.TransactOpts, marketId, collateralType, amount)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_Core *CoreTransactor) ConfigureOracleManager(opts *bind.TransactOpts, oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "configureOracleManager", oracleManagerAddress)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_Core *CoreSession) ConfigureOracleManager(oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _Core.Contract.ConfigureOracleManager(&_Core.TransactOpts, oracleManagerAddress)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_Core *CoreTransactorSession) ConfigureOracleManager(oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _Core.Contract.ConfigureOracleManager(&_Core.TransactOpts, oracleManagerAddress)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_Core *CoreTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_Core *CoreSession) CreateAccount() (*types.Transaction, error) {
	return _Core.Contract.CreateAccount(&_Core.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_Core *CoreTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _Core.Contract.CreateAccount(&_Core.TransactOpts)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_Core *CoreTransactor) CreateAccount0(opts *bind.TransactOpts, requestedAccountId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "createAccount0", requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_Core *CoreSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.CreateAccount0(&_Core.TransactOpts, requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_Core *CoreTransactorSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.CreateAccount0(&_Core.TransactOpts, requestedAccountId)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_Core *CoreTransactor) CreateLock(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "createLock", accountId, collateralType, amount, expireTimestamp)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_Core *CoreSession) CreateLock(accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _Core.Contract.CreateLock(&_Core.TransactOpts, accountId, collateralType, amount, expireTimestamp)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_Core *CoreTransactorSession) CreateLock(accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _Core.Contract.CreateLock(&_Core.TransactOpts, accountId, collateralType, amount, expireTimestamp)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_Core *CoreTransactor) CreatePool(opts *bind.TransactOpts, requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "createPool", requestedPoolId, owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_Core *CoreSession) CreatePool(requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Core.Contract.CreatePool(&_Core.TransactOpts, requestedPoolId, owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_Core *CoreTransactorSession) CreatePool(requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Core.Contract.CreatePool(&_Core.TransactOpts, requestedPoolId, owner)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_Core *CoreTransactor) DelegateCollateral(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "delegateCollateral", accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_Core *CoreSession) DelegateCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DelegateCollateral(&_Core.TransactOpts, accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_Core *CoreTransactorSession) DelegateCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DelegateCollateral(&_Core.TransactOpts, accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreTransactor) Deposit(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "deposit", accountId, collateralType, tokenAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreSession) Deposit(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Deposit(&_Core.TransactOpts, accountId, collateralType, tokenAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreTransactorSession) Deposit(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Deposit(&_Core.TransactOpts, accountId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreTransactor) DepositMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "depositMarketCollateral", marketId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreSession) DepositMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositMarketCollateral(&_Core.TransactOpts, marketId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreTransactorSession) DepositMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositMarketCollateral(&_Core.TransactOpts, marketId, collateralType, tokenAmount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Core *CoreTransactor) DepositMarketUsd(opts *bind.TransactOpts, marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "depositMarketUsd", marketId, target, amount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Core *CoreSession) DepositMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositMarketUsd(&_Core.TransactOpts, marketId, target, amount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Core *CoreTransactorSession) DepositMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositMarketUsd(&_Core.TransactOpts, marketId, target, amount)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_Core *CoreTransactor) DistributeDebtToPools(opts *bind.TransactOpts, marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "distributeDebtToPools", marketId, maxIter)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_Core *CoreSession) DistributeDebtToPools(marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DistributeDebtToPools(&_Core.TransactOpts, marketId, maxIter)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_Core *CoreTransactorSession) DistributeDebtToPools(marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DistributeDebtToPools(&_Core.TransactOpts, marketId, maxIter)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_Core *CoreTransactor) DistributeRewards(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "distributeRewards", poolId, collateralType, amount, start, duration)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_Core *CoreSession) DistributeRewards(poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _Core.Contract.DistributeRewards(&_Core.TransactOpts, poolId, collateralType, amount, start, duration)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_Core *CoreTransactorSession) DistributeRewards(poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _Core.Contract.DistributeRewards(&_Core.TransactOpts, poolId, collateralType, amount, start, duration)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_Core *CoreTransactor) GetMarketDebtPerShare(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "getMarketDebtPerShare", marketId)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_Core *CoreSession) GetMarketDebtPerShare(marketId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.GetMarketDebtPerShare(&_Core.TransactOpts, marketId)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_Core *CoreTransactorSession) GetMarketDebtPerShare(marketId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.GetMarketDebtPerShare(&_Core.TransactOpts, marketId)
}

// GetMarketPoolDebtDistribution is a paid mutator transaction binding the contract method 0x25eeea4b.
//
// Solidity: function getMarketPoolDebtDistribution(uint128 marketId, uint128 poolId) returns(uint256 sharesD18, uint128 totalSharesD18, int128 valuePerShareD27)
func (_Core *CoreTransactor) GetMarketPoolDebtDistribution(opts *bind.TransactOpts, marketId *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "getMarketPoolDebtDistribution", marketId, poolId)
}

// GetMarketPoolDebtDistribution is a paid mutator transaction binding the contract method 0x25eeea4b.
//
// Solidity: function getMarketPoolDebtDistribution(uint128 marketId, uint128 poolId) returns(uint256 sharesD18, uint128 totalSharesD18, int128 valuePerShareD27)
func (_Core *CoreSession) GetMarketPoolDebtDistribution(marketId *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.GetMarketPoolDebtDistribution(&_Core.TransactOpts, marketId, poolId)
}

// GetMarketPoolDebtDistribution is a paid mutator transaction binding the contract method 0x25eeea4b.
//
// Solidity: function getMarketPoolDebtDistribution(uint128 marketId, uint128 poolId) returns(uint256 sharesD18, uint128 totalSharesD18, int128 valuePerShareD27)
func (_Core *CoreTransactorSession) GetMarketPoolDebtDistribution(marketId *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.GetMarketPoolDebtDistribution(&_Core.TransactOpts, marketId, poolId)
}

// GetMarketPools is a paid mutator transaction binding the contract method 0xbe0b8e6f.
//
// Solidity: function getMarketPools(uint128 marketId) returns(uint128[] inRangePoolIds, uint128[] outRangePoolIds)
func (_Core *CoreTransactor) GetMarketPools(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "getMarketPools", marketId)
}

// GetMarketPools is a paid mutator transaction binding the contract method 0xbe0b8e6f.
//
// Solidity: function getMarketPools(uint128 marketId) returns(uint128[] inRangePoolIds, uint128[] outRangePoolIds)
func (_Core *CoreSession) GetMarketPools(marketId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.GetMarketPools(&_Core.TransactOpts, marketId)
}

// GetMarketPools is a paid mutator transaction binding the contract method 0xbe0b8e6f.
//
// Solidity: function getMarketPools(uint128 marketId) returns(uint128[] inRangePoolIds, uint128[] outRangePoolIds)
func (_Core *CoreTransactorSession) GetMarketPools(marketId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.GetMarketPools(&_Core.TransactOpts, marketId)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_Core *CoreTransactor) GetPosition(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "getPosition", accountId, poolId, collateralType)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_Core *CoreSession) GetPosition(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.GetPosition(&_Core.TransactOpts, accountId, poolId, collateralType)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_Core *CoreTransactorSession) GetPosition(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.GetPosition(&_Core.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_Core *CoreTransactor) GetPositionCollateralRatio(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "getPositionCollateralRatio", accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_Core *CoreSession) GetPositionCollateralRatio(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.GetPositionCollateralRatio(&_Core.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_Core *CoreTransactorSession) GetPositionCollateralRatio(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.GetPositionCollateralRatio(&_Core.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_Core *CoreTransactor) GetPositionDebt(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "getPositionDebt", accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_Core *CoreSession) GetPositionDebt(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.GetPositionDebt(&_Core.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_Core *CoreTransactorSession) GetPositionDebt(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.GetPositionDebt(&_Core.TransactOpts, accountId, poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_Core *CoreTransactor) GetVaultCollateralRatio(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "getVaultCollateralRatio", poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_Core *CoreSession) GetVaultCollateralRatio(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.GetVaultCollateralRatio(&_Core.TransactOpts, poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_Core *CoreTransactorSession) GetVaultCollateralRatio(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.GetVaultCollateralRatio(&_Core.TransactOpts, poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_Core *CoreTransactor) GetVaultDebt(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "getVaultDebt", poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_Core *CoreSession) GetVaultDebt(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.GetVaultDebt(&_Core.TransactOpts, poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_Core *CoreTransactorSession) GetVaultDebt(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.GetVaultDebt(&_Core.TransactOpts, poolId, collateralType)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Core *CoreTransactor) GrantPermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "grantPermission", accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Core *CoreSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Core.Contract.GrantPermission(&_Core.TransactOpts, accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Core *CoreTransactorSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Core.Contract.GrantPermission(&_Core.TransactOpts, accountId, permission, user)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_Core *CoreTransactor) InitOrUpgradeNft(opts *bind.TransactOpts, id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "initOrUpgradeNft", id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_Core *CoreSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _Core.Contract.InitOrUpgradeNft(&_Core.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_Core *CoreTransactorSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _Core.Contract.InitOrUpgradeNft(&_Core.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_Core *CoreTransactor) InitOrUpgradeToken(opts *bind.TransactOpts, id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "initOrUpgradeToken", id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_Core *CoreSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _Core.Contract.InitOrUpgradeToken(&_Core.TransactOpts, id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_Core *CoreTransactorSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _Core.Contract.InitOrUpgradeToken(&_Core.TransactOpts, id, name, symbol, decimals, impl)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_Core *CoreTransactor) IsPositionLiquidatable(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "isPositionLiquidatable", accountId, poolId, collateralType)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_Core *CoreSession) IsPositionLiquidatable(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsPositionLiquidatable(&_Core.TransactOpts, accountId, poolId, collateralType)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_Core *CoreTransactorSession) IsPositionLiquidatable(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsPositionLiquidatable(&_Core.TransactOpts, accountId, poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_Core *CoreTransactor) IsVaultLiquidatable(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "isVaultLiquidatable", poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_Core *CoreSession) IsVaultLiquidatable(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsVaultLiquidatable(&_Core.TransactOpts, poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_Core *CoreTransactorSession) IsVaultLiquidatable(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsVaultLiquidatable(&_Core.TransactOpts, poolId, collateralType)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_Core *CoreTransactor) Liquidate(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "liquidate", accountId, poolId, collateralType, liquidateAsAccountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_Core *CoreSession) Liquidate(accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Liquidate(&_Core.TransactOpts, accountId, poolId, collateralType, liquidateAsAccountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_Core *CoreTransactorSession) Liquidate(accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Liquidate(&_Core.TransactOpts, accountId, poolId, collateralType, liquidateAsAccountId)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_Core *CoreTransactor) LiquidateVault(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "liquidateVault", poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_Core *CoreSession) LiquidateVault(poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _Core.Contract.LiquidateVault(&_Core.TransactOpts, poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_Core *CoreTransactorSession) LiquidateVault(poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _Core.Contract.LiquidateVault(&_Core.TransactOpts, poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Core *CoreTransactor) MintUsd(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "mintUsd", accountId, poolId, collateralType, amount)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Core *CoreSession) MintUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.MintUsd(&_Core.TransactOpts, accountId, poolId, collateralType, amount)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Core *CoreTransactorSession) MintUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.MintUsd(&_Core.TransactOpts, accountId, poolId, collateralType, amount)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_Core *CoreTransactor) NominateNewOwner(opts *bind.TransactOpts, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "nominateNewOwner", newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_Core *CoreSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _Core.Contract.NominateNewOwner(&_Core.TransactOpts, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_Core *CoreTransactorSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _Core.Contract.NominateNewOwner(&_Core.TransactOpts, newNominatedOwner)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_Core *CoreTransactor) NominatePoolOwner(opts *bind.TransactOpts, nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "nominatePoolOwner", nominatedOwner, poolId)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_Core *CoreSession) NominatePoolOwner(nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.NominatePoolOwner(&_Core.TransactOpts, nominatedOwner, poolId)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_Core *CoreTransactorSession) NominatePoolOwner(nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.NominatePoolOwner(&_Core.TransactOpts, nominatedOwner, poolId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_Core *CoreTransactor) NotifyAccountTransfer(opts *bind.TransactOpts, to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "notifyAccountTransfer", to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_Core *CoreSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.NotifyAccountTransfer(&_Core.TransactOpts, to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_Core *CoreTransactorSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.NotifyAccountTransfer(&_Core.TransactOpts, to, accountId)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_Core *CoreTransactor) RebalancePool(opts *bind.TransactOpts, poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "rebalancePool", poolId, optionalCollateralType)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_Core *CoreSession) RebalancePool(poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.RebalancePool(&_Core.TransactOpts, poolId, optionalCollateralType)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_Core *CoreTransactorSession) RebalancePool(poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _Core.Contract.RebalancePool(&_Core.TransactOpts, poolId, optionalCollateralType)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_Core *CoreTransactor) RegisterMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerMarket", market)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_Core *CoreSession) RegisterMarket(market common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterMarket(&_Core.TransactOpts, market)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_Core *CoreTransactorSession) RegisterMarket(market common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterMarket(&_Core.TransactOpts, market)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Core *CoreTransactor) RegisterRewardsDistributor(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerRewardsDistributor", poolId, collateralType, distributor)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Core *CoreSession) RegisterRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterRewardsDistributor(&_Core.TransactOpts, poolId, collateralType, distributor)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Core *CoreTransactorSession) RegisterRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterRewardsDistributor(&_Core.TransactOpts, poolId, collateralType, distributor)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_Core *CoreTransactor) RegisterUnmanagedSystem(opts *bind.TransactOpts, id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerUnmanagedSystem", id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_Core *CoreSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterUnmanagedSystem(&_Core.TransactOpts, id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_Core *CoreTransactorSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterUnmanagedSystem(&_Core.TransactOpts, id, endpoint)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_Core *CoreTransactor) RemoveApprovedPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "removeApprovedPool", poolId)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_Core *CoreSession) RemoveApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RemoveApprovedPool(&_Core.TransactOpts, poolId)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_Core *CoreTransactorSession) RemoveApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RemoveApprovedPool(&_Core.TransactOpts, poolId)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Core *CoreTransactor) RemoveFromFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "removeFromFeatureFlagAllowlist", feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Core *CoreSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Core.Contract.RemoveFromFeatureFlagAllowlist(&_Core.TransactOpts, feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_Core *CoreTransactorSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _Core.Contract.RemoveFromFeatureFlagAllowlist(&_Core.TransactOpts, feature, account)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Core *CoreTransactor) RemoveRewardsDistributor(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "removeRewardsDistributor", poolId, collateralType, distributor)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Core *CoreSession) RemoveRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Core.Contract.RemoveRewardsDistributor(&_Core.TransactOpts, poolId, collateralType, distributor)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Core *CoreTransactorSession) RemoveRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Core.Contract.RemoveRewardsDistributor(&_Core.TransactOpts, poolId, collateralType, distributor)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_Core *CoreTransactor) RenounceNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "renounceNomination")
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_Core *CoreSession) RenounceNomination() (*types.Transaction, error) {
	return _Core.Contract.RenounceNomination(&_Core.TransactOpts)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_Core *CoreTransactorSession) RenounceNomination() (*types.Transaction, error) {
	return _Core.Contract.RenounceNomination(&_Core.TransactOpts)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_Core *CoreTransactor) RenouncePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "renouncePermission", accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_Core *CoreSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _Core.Contract.RenouncePermission(&_Core.TransactOpts, accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_Core *CoreTransactorSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _Core.Contract.RenouncePermission(&_Core.TransactOpts, accountId, permission)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_Core *CoreTransactor) RenouncePoolNomination(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "renouncePoolNomination", poolId)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_Core *CoreSession) RenouncePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RenouncePoolNomination(&_Core.TransactOpts, poolId)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_Core *CoreTransactorSession) RenouncePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RenouncePoolNomination(&_Core.TransactOpts, poolId)
}

// RenouncePoolOwnership is a paid mutator transaction binding the contract method 0x7cc14a92.
//
// Solidity: function renouncePoolOwnership(uint128 poolId) returns()
func (_Core *CoreTransactor) RenouncePoolOwnership(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "renouncePoolOwnership", poolId)
}

// RenouncePoolOwnership is a paid mutator transaction binding the contract method 0x7cc14a92.
//
// Solidity: function renouncePoolOwnership(uint128 poolId) returns()
func (_Core *CoreSession) RenouncePoolOwnership(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RenouncePoolOwnership(&_Core.TransactOpts, poolId)
}

// RenouncePoolOwnership is a paid mutator transaction binding the contract method 0x7cc14a92.
//
// Solidity: function renouncePoolOwnership(uint128 poolId) returns()
func (_Core *CoreTransactorSession) RenouncePoolOwnership(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RenouncePoolOwnership(&_Core.TransactOpts, poolId)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Core *CoreTransactor) RevokePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "revokePermission", accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Core *CoreSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Core.Contract.RevokePermission(&_Core.TransactOpts, accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_Core *CoreTransactorSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _Core.Contract.RevokePermission(&_Core.TransactOpts, accountId, permission, user)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_Core *CoreTransactor) RevokePoolNomination(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "revokePoolNomination", poolId)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_Core *CoreSession) RevokePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RevokePoolNomination(&_Core.TransactOpts, poolId)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_Core *CoreTransactorSession) RevokePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RevokePoolNomination(&_Core.TransactOpts, poolId)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_Core *CoreTransactor) SetConfig(opts *bind.TransactOpts, k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setConfig", k, v)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_Core *CoreSession) SetConfig(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Core.Contract.SetConfig(&_Core.TransactOpts, k, v)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_Core *CoreTransactorSession) SetConfig(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Core.Contract.SetConfig(&_Core.TransactOpts, k, v)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_Core *CoreTransactor) SetDeniers(opts *bind.TransactOpts, feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setDeniers", feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_Core *CoreSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _Core.Contract.SetDeniers(&_Core.TransactOpts, feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_Core *CoreTransactorSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _Core.Contract.SetDeniers(&_Core.TransactOpts, feature, deniers)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_Core *CoreTransactor) SetFeatureFlagAllowAll(opts *bind.TransactOpts, feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setFeatureFlagAllowAll", feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_Core *CoreSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _Core.Contract.SetFeatureFlagAllowAll(&_Core.TransactOpts, feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_Core *CoreTransactorSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _Core.Contract.SetFeatureFlagAllowAll(&_Core.TransactOpts, feature, allowAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_Core *CoreTransactor) SetFeatureFlagDenyAll(opts *bind.TransactOpts, feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setFeatureFlagDenyAll", feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_Core *CoreSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _Core.Contract.SetFeatureFlagDenyAll(&_Core.TransactOpts, feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_Core *CoreTransactorSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _Core.Contract.SetFeatureFlagDenyAll(&_Core.TransactOpts, feature, denyAll)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_Core *CoreTransactor) SetMarketMinDelegateTime(opts *bind.TransactOpts, marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setMarketMinDelegateTime", marketId, minDelegateTime)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_Core *CoreSession) SetMarketMinDelegateTime(marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _Core.Contract.SetMarketMinDelegateTime(&_Core.TransactOpts, marketId, minDelegateTime)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_Core *CoreTransactorSession) SetMarketMinDelegateTime(marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _Core.Contract.SetMarketMinDelegateTime(&_Core.TransactOpts, marketId, minDelegateTime)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_Core *CoreTransactor) SetMinLiquidityRatio(opts *bind.TransactOpts, marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setMinLiquidityRatio", marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_Core *CoreSession) SetMinLiquidityRatio(marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SetMinLiquidityRatio(&_Core.TransactOpts, marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_Core *CoreTransactorSession) SetMinLiquidityRatio(marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SetMinLiquidityRatio(&_Core.TransactOpts, marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_Core *CoreTransactor) SetMinLiquidityRatio0(opts *bind.TransactOpts, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setMinLiquidityRatio0", minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_Core *CoreSession) SetMinLiquidityRatio0(minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SetMinLiquidityRatio0(&_Core.TransactOpts, minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_Core *CoreTransactorSession) SetMinLiquidityRatio0(minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SetMinLiquidityRatio0(&_Core.TransactOpts, minLiquidityRatio)
}

// SetPoolCollateralConfiguration is a paid mutator transaction binding the contract method 0x5a4aabb1.
//
// Solidity: function setPoolCollateralConfiguration(uint128 poolId, address collateralType, (uint256,uint256) newConfig) returns()
func (_Core *CoreTransactor) SetPoolCollateralConfiguration(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, newConfig PoolCollateralConfigurationData) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setPoolCollateralConfiguration", poolId, collateralType, newConfig)
}

// SetPoolCollateralConfiguration is a paid mutator transaction binding the contract method 0x5a4aabb1.
//
// Solidity: function setPoolCollateralConfiguration(uint128 poolId, address collateralType, (uint256,uint256) newConfig) returns()
func (_Core *CoreSession) SetPoolCollateralConfiguration(poolId *big.Int, collateralType common.Address, newConfig PoolCollateralConfigurationData) (*types.Transaction, error) {
	return _Core.Contract.SetPoolCollateralConfiguration(&_Core.TransactOpts, poolId, collateralType, newConfig)
}

// SetPoolCollateralConfiguration is a paid mutator transaction binding the contract method 0x5a4aabb1.
//
// Solidity: function setPoolCollateralConfiguration(uint128 poolId, address collateralType, (uint256,uint256) newConfig) returns()
func (_Core *CoreTransactorSession) SetPoolCollateralConfiguration(poolId *big.Int, collateralType common.Address, newConfig PoolCollateralConfigurationData) (*types.Transaction, error) {
	return _Core.Contract.SetPoolCollateralConfiguration(&_Core.TransactOpts, poolId, collateralType, newConfig)
}

// SetPoolCollateralDisabledByDefault is a paid mutator transaction binding the contract method 0x4c6568b1.
//
// Solidity: function setPoolCollateralDisabledByDefault(uint128 poolId, bool disabled) returns()
func (_Core *CoreTransactor) SetPoolCollateralDisabledByDefault(opts *bind.TransactOpts, poolId *big.Int, disabled bool) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setPoolCollateralDisabledByDefault", poolId, disabled)
}

// SetPoolCollateralDisabledByDefault is a paid mutator transaction binding the contract method 0x4c6568b1.
//
// Solidity: function setPoolCollateralDisabledByDefault(uint128 poolId, bool disabled) returns()
func (_Core *CoreSession) SetPoolCollateralDisabledByDefault(poolId *big.Int, disabled bool) (*types.Transaction, error) {
	return _Core.Contract.SetPoolCollateralDisabledByDefault(&_Core.TransactOpts, poolId, disabled)
}

// SetPoolCollateralDisabledByDefault is a paid mutator transaction binding the contract method 0x4c6568b1.
//
// Solidity: function setPoolCollateralDisabledByDefault(uint128 poolId, bool disabled) returns()
func (_Core *CoreTransactorSession) SetPoolCollateralDisabledByDefault(poolId *big.Int, disabled bool) (*types.Transaction, error) {
	return _Core.Contract.SetPoolCollateralDisabledByDefault(&_Core.TransactOpts, poolId, disabled)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_Core *CoreTransactor) SetPoolConfiguration(opts *bind.TransactOpts, poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setPoolConfiguration", poolId, newMarketConfigurations)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_Core *CoreSession) SetPoolConfiguration(poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _Core.Contract.SetPoolConfiguration(&_Core.TransactOpts, poolId, newMarketConfigurations)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_Core *CoreTransactorSession) SetPoolConfiguration(poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _Core.Contract.SetPoolConfiguration(&_Core.TransactOpts, poolId, newMarketConfigurations)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_Core *CoreTransactor) SetPoolName(opts *bind.TransactOpts, poolId *big.Int, name string) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setPoolName", poolId, name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_Core *CoreSession) SetPoolName(poolId *big.Int, name string) (*types.Transaction, error) {
	return _Core.Contract.SetPoolName(&_Core.TransactOpts, poolId, name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_Core *CoreTransactorSession) SetPoolName(poolId *big.Int, name string) (*types.Transaction, error) {
	return _Core.Contract.SetPoolName(&_Core.TransactOpts, poolId, name)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_Core *CoreTransactor) SetPreferredPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setPreferredPool", poolId)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_Core *CoreSession) SetPreferredPool(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SetPreferredPool(&_Core.TransactOpts, poolId)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_Core *CoreTransactorSession) SetPreferredPool(poolId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SetPreferredPool(&_Core.TransactOpts, poolId)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_Core *CoreTransactor) SetSupportedCrossChainNetworks(opts *bind.TransactOpts, supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setSupportedCrossChainNetworks", supportedNetworks, ccipSelectors)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_Core *CoreSession) SetSupportedCrossChainNetworks(supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _Core.Contract.SetSupportedCrossChainNetworks(&_Core.TransactOpts, supportedNetworks, ccipSelectors)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_Core *CoreTransactorSession) SetSupportedCrossChainNetworks(supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _Core.Contract.SetSupportedCrossChainNetworks(&_Core.TransactOpts, supportedNetworks, ccipSelectors)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_Core *CoreTransactor) SimulateUpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "simulateUpgradeTo", newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_Core *CoreSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Core.Contract.SimulateUpgradeTo(&_Core.TransactOpts, newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_Core *CoreTransactorSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Core.Contract.SimulateUpgradeTo(&_Core.TransactOpts, newImplementation)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_Core *CoreTransactor) TransferCrossChain(opts *bind.TransactOpts, destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "transferCrossChain", destChainId, amount)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_Core *CoreSession) TransferCrossChain(destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.TransferCrossChain(&_Core.TransactOpts, destChainId, amount)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_Core *CoreTransactorSession) TransferCrossChain(destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.TransferCrossChain(&_Core.TransactOpts, destChainId, amount)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_Core *CoreTransactor) UpdateRewards(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "updateRewards", poolId, collateralType, accountId)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_Core *CoreSession) UpdateRewards(poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.UpdateRewards(&_Core.TransactOpts, poolId, collateralType, accountId)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_Core *CoreTransactorSession) UpdateRewards(poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.UpdateRewards(&_Core.TransactOpts, poolId, collateralType, accountId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Core *CoreTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Core *CoreSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Core.Contract.UpgradeTo(&_Core.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Core *CoreTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Core.Contract.UpgradeTo(&_Core.TransactOpts, newImplementation)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreTransactor) Withdraw(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "withdraw", accountId, collateralType, tokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreSession) Withdraw(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Withdraw(&_Core.TransactOpts, accountId, collateralType, tokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreTransactorSession) Withdraw(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Withdraw(&_Core.TransactOpts, accountId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreTransactor) WithdrawMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "withdrawMarketCollateral", marketId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreSession) WithdrawMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.WithdrawMarketCollateral(&_Core.TransactOpts, marketId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Core *CoreTransactorSession) WithdrawMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.WithdrawMarketCollateral(&_Core.TransactOpts, marketId, collateralType, tokenAmount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Core *CoreTransactor) WithdrawMarketUsd(opts *bind.TransactOpts, marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "withdrawMarketUsd", marketId, target, amount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Core *CoreSession) WithdrawMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.WithdrawMarketUsd(&_Core.TransactOpts, marketId, target, amount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Core *CoreTransactorSession) WithdrawMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.WithdrawMarketUsd(&_Core.TransactOpts, marketId, target, amount)
}

// CoreAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Core contract.
type CoreAccountCreatedIterator struct {
	Event *CoreAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreAccountCreated represents a AccountCreated event raised by the Core contract.
type CoreAccountCreated struct {
	AccountId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_Core *CoreFilterer) FilterAccountCreated(opts *bind.FilterOpts, accountId []*big.Int, owner []common.Address) (*CoreAccountCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreAccountCreatedIterator{contract: _Core.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_Core *CoreFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *CoreAccountCreated, accountId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreAccountCreated)
				if err := _Core.contract.UnpackLog(event, "AccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseAccountCreated(log types.Log) (*CoreAccountCreated, error) {
	event := new(CoreAccountCreated)
	if err := _Core.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreAssociatedSystemSetIterator is returned from FilterAssociatedSystemSet and is used to iterate over the raw logs and unpacked data for AssociatedSystemSet events raised by the Core contract.
type CoreAssociatedSystemSetIterator struct {
	Event *CoreAssociatedSystemSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreAssociatedSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreAssociatedSystemSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreAssociatedSystemSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreAssociatedSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreAssociatedSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreAssociatedSystemSet represents a AssociatedSystemSet event raised by the Core contract.
type CoreAssociatedSystemSet struct {
	Kind  [32]byte
	Id    [32]byte
	Proxy common.Address
	Impl  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssociatedSystemSet is a free log retrieval operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_Core *CoreFilterer) FilterAssociatedSystemSet(opts *bind.FilterOpts, kind [][32]byte, id [][32]byte) (*CoreAssociatedSystemSetIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return &CoreAssociatedSystemSetIterator{contract: _Core.contract, event: "AssociatedSystemSet", logs: logs, sub: sub}, nil
}

// WatchAssociatedSystemSet is a free log subscription operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_Core *CoreFilterer) WatchAssociatedSystemSet(opts *bind.WatchOpts, sink chan<- *CoreAssociatedSystemSet, kind [][32]byte, id [][32]byte) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreAssociatedSystemSet)
				if err := _Core.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseAssociatedSystemSet(log types.Log) (*CoreAssociatedSystemSet, error) {
	event := new(CoreAssociatedSystemSet)
	if err := _Core.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreCollateralConfiguredIterator is returned from FilterCollateralConfigured and is used to iterate over the raw logs and unpacked data for CollateralConfigured events raised by the Core contract.
type CoreCollateralConfiguredIterator struct {
	Event *CoreCollateralConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreCollateralConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreCollateralConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreCollateralConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreCollateralConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreCollateralConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreCollateralConfigured represents a CollateralConfigured event raised by the Core contract.
type CoreCollateralConfigured struct {
	CollateralType common.Address
	Config         CollateralConfigurationData
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCollateralConfigured is a free log retrieval operation binding the contract event 0xefc23317f58afd6b22480bd22174cc7da0913bce25c03d9859216dacddebe6fe.
//
// Solidity: event CollateralConfigured(address indexed collateralType, (bool,uint256,uint256,uint256,bytes32,address,uint256) config)
func (_Core *CoreFilterer) FilterCollateralConfigured(opts *bind.FilterOpts, collateralType []common.Address) (*CoreCollateralConfiguredIterator, error) {

	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "CollateralConfigured", collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreCollateralConfiguredIterator{contract: _Core.contract, event: "CollateralConfigured", logs: logs, sub: sub}, nil
}

// WatchCollateralConfigured is a free log subscription operation binding the contract event 0xefc23317f58afd6b22480bd22174cc7da0913bce25c03d9859216dacddebe6fe.
//
// Solidity: event CollateralConfigured(address indexed collateralType, (bool,uint256,uint256,uint256,bytes32,address,uint256) config)
func (_Core *CoreFilterer) WatchCollateralConfigured(opts *bind.WatchOpts, sink chan<- *CoreCollateralConfigured, collateralType []common.Address) (event.Subscription, error) {

	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "CollateralConfigured", collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreCollateralConfigured)
				if err := _Core.contract.UnpackLog(event, "CollateralConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseCollateralConfigured(log types.Log) (*CoreCollateralConfigured, error) {
	event := new(CoreCollateralConfigured)
	if err := _Core.contract.UnpackLog(event, "CollateralConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreCollateralLockCreatedIterator is returned from FilterCollateralLockCreated and is used to iterate over the raw logs and unpacked data for CollateralLockCreated events raised by the Core contract.
type CoreCollateralLockCreatedIterator struct {
	Event *CoreCollateralLockCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreCollateralLockCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreCollateralLockCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreCollateralLockCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreCollateralLockCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreCollateralLockCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreCollateralLockCreated represents a CollateralLockCreated event raised by the Core contract.
type CoreCollateralLockCreated struct {
	AccountId       *big.Int
	CollateralType  common.Address
	TokenAmount     *big.Int
	ExpireTimestamp uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollateralLockCreated is a free log retrieval operation binding the contract event 0x8a78446a234220d3ee3f46aa6ea1ea5bc438bd153398ebcbd171843744b452a6.
//
// Solidity: event CollateralLockCreated(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_Core *CoreFilterer) FilterCollateralLockCreated(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address) (*CoreCollateralLockCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "CollateralLockCreated", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreCollateralLockCreatedIterator{contract: _Core.contract, event: "CollateralLockCreated", logs: logs, sub: sub}, nil
}

// WatchCollateralLockCreated is a free log subscription operation binding the contract event 0x8a78446a234220d3ee3f46aa6ea1ea5bc438bd153398ebcbd171843744b452a6.
//
// Solidity: event CollateralLockCreated(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_Core *CoreFilterer) WatchCollateralLockCreated(opts *bind.WatchOpts, sink chan<- *CoreCollateralLockCreated, accountId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "CollateralLockCreated", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreCollateralLockCreated)
				if err := _Core.contract.UnpackLog(event, "CollateralLockCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseCollateralLockCreated(log types.Log) (*CoreCollateralLockCreated, error) {
	event := new(CoreCollateralLockCreated)
	if err := _Core.contract.UnpackLog(event, "CollateralLockCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreCollateralLockExpiredIterator is returned from FilterCollateralLockExpired and is used to iterate over the raw logs and unpacked data for CollateralLockExpired events raised by the Core contract.
type CoreCollateralLockExpiredIterator struct {
	Event *CoreCollateralLockExpired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreCollateralLockExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreCollateralLockExpired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreCollateralLockExpired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreCollateralLockExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreCollateralLockExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreCollateralLockExpired represents a CollateralLockExpired event raised by the Core contract.
type CoreCollateralLockExpired struct {
	AccountId       *big.Int
	CollateralType  common.Address
	TokenAmount     *big.Int
	ExpireTimestamp uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollateralLockExpired is a free log retrieval operation binding the contract event 0xc5e2c7afc4e54d998977e11260a0bfc0ad5678a3a8b6628162f9d4e642d7f160.
//
// Solidity: event CollateralLockExpired(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_Core *CoreFilterer) FilterCollateralLockExpired(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address) (*CoreCollateralLockExpiredIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "CollateralLockExpired", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreCollateralLockExpiredIterator{contract: _Core.contract, event: "CollateralLockExpired", logs: logs, sub: sub}, nil
}

// WatchCollateralLockExpired is a free log subscription operation binding the contract event 0xc5e2c7afc4e54d998977e11260a0bfc0ad5678a3a8b6628162f9d4e642d7f160.
//
// Solidity: event CollateralLockExpired(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_Core *CoreFilterer) WatchCollateralLockExpired(opts *bind.WatchOpts, sink chan<- *CoreCollateralLockExpired, accountId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "CollateralLockExpired", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreCollateralLockExpired)
				if err := _Core.contract.UnpackLog(event, "CollateralLockExpired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseCollateralLockExpired(log types.Log) (*CoreCollateralLockExpired, error) {
	event := new(CoreCollateralLockExpired)
	if err := _Core.contract.UnpackLog(event, "CollateralLockExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreDebtAssociatedIterator is returned from FilterDebtAssociated and is used to iterate over the raw logs and unpacked data for DebtAssociated events raised by the Core contract.
type CoreDebtAssociatedIterator struct {
	Event *CoreDebtAssociated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreDebtAssociatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreDebtAssociated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreDebtAssociated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreDebtAssociatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreDebtAssociatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreDebtAssociated represents a DebtAssociated event raised by the Core contract.
type CoreDebtAssociated struct {
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
func (_Core *CoreFilterer) FilterDebtAssociated(opts *bind.FilterOpts, marketId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*CoreDebtAssociatedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "DebtAssociated", marketIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreDebtAssociatedIterator{contract: _Core.contract, event: "DebtAssociated", logs: logs, sub: sub}, nil
}

// WatchDebtAssociated is a free log subscription operation binding the contract event 0xb03bc7530b5650601d2c8bc86c81a45b6b50b0099defd3f17a6bdb48660180ad.
//
// Solidity: event DebtAssociated(uint128 indexed marketId, uint128 indexed poolId, address indexed collateralType, uint128 accountId, uint256 amount, int256 updatedDebt)
func (_Core *CoreFilterer) WatchDebtAssociated(opts *bind.WatchOpts, sink chan<- *CoreDebtAssociated, marketId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "DebtAssociated", marketIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreDebtAssociated)
				if err := _Core.contract.UnpackLog(event, "DebtAssociated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseDebtAssociated(log types.Log) (*CoreDebtAssociated, error) {
	event := new(CoreDebtAssociated)
	if err := _Core.contract.UnpackLog(event, "DebtAssociated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreDelegationUpdatedIterator is returned from FilterDelegationUpdated and is used to iterate over the raw logs and unpacked data for DelegationUpdated events raised by the Core contract.
type CoreDelegationUpdatedIterator struct {
	Event *CoreDelegationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreDelegationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreDelegationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreDelegationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreDelegationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreDelegationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreDelegationUpdated represents a DelegationUpdated event raised by the Core contract.
type CoreDelegationUpdated struct {
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
func (_Core *CoreFilterer) FilterDelegationUpdated(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*CoreDelegationUpdatedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "DelegationUpdated", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreDelegationUpdatedIterator{contract: _Core.contract, event: "DelegationUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegationUpdated is a free log subscription operation binding the contract event 0x7b12dd38f18c0ff77ae702f6da13fbbcb28f53f807ecc7d39ee8d8b1ea8295ad.
//
// Solidity: event DelegationUpdated(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, uint256 leverage, address indexed sender)
func (_Core *CoreFilterer) WatchDelegationUpdated(opts *bind.WatchOpts, sink chan<- *CoreDelegationUpdated, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "DelegationUpdated", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreDelegationUpdated)
				if err := _Core.contract.UnpackLog(event, "DelegationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseDelegationUpdated(log types.Log) (*CoreDelegationUpdated, error) {
	event := new(CoreDelegationUpdated)
	if err := _Core.contract.UnpackLog(event, "DelegationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Core contract.
type CoreDepositedIterator struct {
	Event *CoreDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreDeposited represents a Deposited event raised by the Core contract.
type CoreDeposited struct {
	AccountId      *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xd92122e67326e9313bfae33ccb1fccf5194584c6bf93a8529a6b006d8c6e24a9.
//
// Solidity: event Deposited(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Core *CoreFilterer) FilterDeposited(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreDepositedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "Deposited", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreDepositedIterator{contract: _Core.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xd92122e67326e9313bfae33ccb1fccf5194584c6bf93a8529a6b006d8c6e24a9.
//
// Solidity: event Deposited(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Core *CoreFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *CoreDeposited, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "Deposited", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreDeposited)
				if err := _Core.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseDeposited(log types.Log) (*CoreDeposited, error) {
	event := new(CoreDeposited)
	if err := _Core.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreFeatureFlagAllowAllSetIterator is returned from FilterFeatureFlagAllowAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowAllSet events raised by the Core contract.
type CoreFeatureFlagAllowAllSetIterator struct {
	Event *CoreFeatureFlagAllowAllSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreFeatureFlagAllowAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreFeatureFlagAllowAllSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreFeatureFlagAllowAllSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreFeatureFlagAllowAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreFeatureFlagAllowAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreFeatureFlagAllowAllSet represents a FeatureFlagAllowAllSet event raised by the Core contract.
type CoreFeatureFlagAllowAllSet struct {
	Feature  [32]byte
	AllowAll bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowAllSet is a free log retrieval operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_Core *CoreFilterer) FilterFeatureFlagAllowAllSet(opts *bind.FilterOpts, feature [][32]byte) (*CoreFeatureFlagAllowAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreFeatureFlagAllowAllSetIterator{contract: _Core.contract, event: "FeatureFlagAllowAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowAllSet is a free log subscription operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_Core *CoreFilterer) WatchFeatureFlagAllowAllSet(opts *bind.WatchOpts, sink chan<- *CoreFeatureFlagAllowAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreFeatureFlagAllowAllSet)
				if err := _Core.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseFeatureFlagAllowAllSet(log types.Log) (*CoreFeatureFlagAllowAllSet, error) {
	event := new(CoreFeatureFlagAllowAllSet)
	if err := _Core.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreFeatureFlagAllowlistAddedIterator is returned from FilterFeatureFlagAllowlistAdded and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistAdded events raised by the Core contract.
type CoreFeatureFlagAllowlistAddedIterator struct {
	Event *CoreFeatureFlagAllowlistAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreFeatureFlagAllowlistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreFeatureFlagAllowlistAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreFeatureFlagAllowlistAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreFeatureFlagAllowlistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreFeatureFlagAllowlistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreFeatureFlagAllowlistAdded represents a FeatureFlagAllowlistAdded event raised by the Core contract.
type CoreFeatureFlagAllowlistAdded struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistAdded is a free log retrieval operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_Core *CoreFilterer) FilterFeatureFlagAllowlistAdded(opts *bind.FilterOpts, feature [][32]byte) (*CoreFeatureFlagAllowlistAddedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreFeatureFlagAllowlistAddedIterator{contract: _Core.contract, event: "FeatureFlagAllowlistAdded", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistAdded is a free log subscription operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_Core *CoreFilterer) WatchFeatureFlagAllowlistAdded(opts *bind.WatchOpts, sink chan<- *CoreFeatureFlagAllowlistAdded, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreFeatureFlagAllowlistAdded)
				if err := _Core.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseFeatureFlagAllowlistAdded(log types.Log) (*CoreFeatureFlagAllowlistAdded, error) {
	event := new(CoreFeatureFlagAllowlistAdded)
	if err := _Core.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreFeatureFlagAllowlistRemovedIterator is returned from FilterFeatureFlagAllowlistRemoved and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistRemoved events raised by the Core contract.
type CoreFeatureFlagAllowlistRemovedIterator struct {
	Event *CoreFeatureFlagAllowlistRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreFeatureFlagAllowlistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreFeatureFlagAllowlistRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreFeatureFlagAllowlistRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreFeatureFlagAllowlistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreFeatureFlagAllowlistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreFeatureFlagAllowlistRemoved represents a FeatureFlagAllowlistRemoved event raised by the Core contract.
type CoreFeatureFlagAllowlistRemoved struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistRemoved is a free log retrieval operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_Core *CoreFilterer) FilterFeatureFlagAllowlistRemoved(opts *bind.FilterOpts, feature [][32]byte) (*CoreFeatureFlagAllowlistRemovedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreFeatureFlagAllowlistRemovedIterator{contract: _Core.contract, event: "FeatureFlagAllowlistRemoved", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistRemoved is a free log subscription operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_Core *CoreFilterer) WatchFeatureFlagAllowlistRemoved(opts *bind.WatchOpts, sink chan<- *CoreFeatureFlagAllowlistRemoved, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreFeatureFlagAllowlistRemoved)
				if err := _Core.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseFeatureFlagAllowlistRemoved(log types.Log) (*CoreFeatureFlagAllowlistRemoved, error) {
	event := new(CoreFeatureFlagAllowlistRemoved)
	if err := _Core.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreFeatureFlagDeniersResetIterator is returned from FilterFeatureFlagDeniersReset and is used to iterate over the raw logs and unpacked data for FeatureFlagDeniersReset events raised by the Core contract.
type CoreFeatureFlagDeniersResetIterator struct {
	Event *CoreFeatureFlagDeniersReset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreFeatureFlagDeniersResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreFeatureFlagDeniersReset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreFeatureFlagDeniersReset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreFeatureFlagDeniersResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreFeatureFlagDeniersResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreFeatureFlagDeniersReset represents a FeatureFlagDeniersReset event raised by the Core contract.
type CoreFeatureFlagDeniersReset struct {
	Feature [32]byte
	Deniers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDeniersReset is a free log retrieval operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_Core *CoreFilterer) FilterFeatureFlagDeniersReset(opts *bind.FilterOpts, feature [][32]byte) (*CoreFeatureFlagDeniersResetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreFeatureFlagDeniersResetIterator{contract: _Core.contract, event: "FeatureFlagDeniersReset", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDeniersReset is a free log subscription operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_Core *CoreFilterer) WatchFeatureFlagDeniersReset(opts *bind.WatchOpts, sink chan<- *CoreFeatureFlagDeniersReset, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreFeatureFlagDeniersReset)
				if err := _Core.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseFeatureFlagDeniersReset(log types.Log) (*CoreFeatureFlagDeniersReset, error) {
	event := new(CoreFeatureFlagDeniersReset)
	if err := _Core.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreFeatureFlagDenyAllSetIterator is returned from FilterFeatureFlagDenyAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagDenyAllSet events raised by the Core contract.
type CoreFeatureFlagDenyAllSetIterator struct {
	Event *CoreFeatureFlagDenyAllSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreFeatureFlagDenyAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreFeatureFlagDenyAllSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreFeatureFlagDenyAllSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreFeatureFlagDenyAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreFeatureFlagDenyAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreFeatureFlagDenyAllSet represents a FeatureFlagDenyAllSet event raised by the Core contract.
type CoreFeatureFlagDenyAllSet struct {
	Feature [32]byte
	DenyAll bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDenyAllSet is a free log retrieval operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_Core *CoreFilterer) FilterFeatureFlagDenyAllSet(opts *bind.FilterOpts, feature [][32]byte) (*CoreFeatureFlagDenyAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreFeatureFlagDenyAllSetIterator{contract: _Core.contract, event: "FeatureFlagDenyAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDenyAllSet is a free log subscription operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_Core *CoreFilterer) WatchFeatureFlagDenyAllSet(opts *bind.WatchOpts, sink chan<- *CoreFeatureFlagDenyAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreFeatureFlagDenyAllSet)
				if err := _Core.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseFeatureFlagDenyAllSet(log types.Log) (*CoreFeatureFlagDenyAllSet, error) {
	event := new(CoreFeatureFlagDenyAllSet)
	if err := _Core.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreIssuanceFeePaidIterator is returned from FilterIssuanceFeePaid and is used to iterate over the raw logs and unpacked data for IssuanceFeePaid events raised by the Core contract.
type CoreIssuanceFeePaidIterator struct {
	Event *CoreIssuanceFeePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreIssuanceFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreIssuanceFeePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreIssuanceFeePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreIssuanceFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreIssuanceFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreIssuanceFeePaid represents a IssuanceFeePaid event raised by the Core contract.
type CoreIssuanceFeePaid struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	FeeAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIssuanceFeePaid is a free log retrieval operation binding the contract event 0x28d0fb10e1c8ce51490a16fb3b40bf23f8064f7c624d3a3852ad66683a25995d.
//
// Solidity: event IssuanceFeePaid(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 feeAmount)
func (_Core *CoreFilterer) FilterIssuanceFeePaid(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int) (*CoreIssuanceFeePaidIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "IssuanceFeePaid", accountIdRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreIssuanceFeePaidIterator{contract: _Core.contract, event: "IssuanceFeePaid", logs: logs, sub: sub}, nil
}

// WatchIssuanceFeePaid is a free log subscription operation binding the contract event 0x28d0fb10e1c8ce51490a16fb3b40bf23f8064f7c624d3a3852ad66683a25995d.
//
// Solidity: event IssuanceFeePaid(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 feeAmount)
func (_Core *CoreFilterer) WatchIssuanceFeePaid(opts *bind.WatchOpts, sink chan<- *CoreIssuanceFeePaid, accountId []*big.Int, poolId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "IssuanceFeePaid", accountIdRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreIssuanceFeePaid)
				if err := _Core.contract.UnpackLog(event, "IssuanceFeePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseIssuanceFeePaid(log types.Log) (*CoreIssuanceFeePaid, error) {
	event := new(CoreIssuanceFeePaid)
	if err := _Core.contract.UnpackLog(event, "IssuanceFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLiquidationIterator is returned from FilterLiquidation and is used to iterate over the raw logs and unpacked data for Liquidation events raised by the Core contract.
type CoreLiquidationIterator struct {
	Event *CoreLiquidation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLiquidation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLiquidation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLiquidation represents a Liquidation event raised by the Core contract.
type CoreLiquidation struct {
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
func (_Core *CoreFilterer) FilterLiquidation(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*CoreLiquidationIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "Liquidation", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreLiquidationIterator{contract: _Core.contract, event: "Liquidation", logs: logs, sub: sub}, nil
}

// WatchLiquidation is a free log subscription operation binding the contract event 0xe6c1b26644f880854bf954d4186be9f0b2d06d50fa0484b596e79d409c07a5fd.
//
// Solidity: event Liquidation(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_Core *CoreFilterer) WatchLiquidation(opts *bind.WatchOpts, sink chan<- *CoreLiquidation, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "Liquidation", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLiquidation)
				if err := _Core.contract.UnpackLog(event, "Liquidation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseLiquidation(log types.Log) (*CoreLiquidation, error) {
	event := new(CoreLiquidation)
	if err := _Core.contract.UnpackLog(event, "Liquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreMarketCollateralDepositedIterator is returned from FilterMarketCollateralDeposited and is used to iterate over the raw logs and unpacked data for MarketCollateralDeposited events raised by the Core contract.
type CoreMarketCollateralDepositedIterator struct {
	Event *CoreMarketCollateralDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreMarketCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreMarketCollateralDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreMarketCollateralDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreMarketCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreMarketCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreMarketCollateralDeposited represents a MarketCollateralDeposited event raised by the Core contract.
type CoreMarketCollateralDeposited struct {
	MarketId                 *big.Int
	CollateralType           common.Address
	TokenAmount              *big.Int
	Sender                   common.Address
	CreditCapacity           *big.Int
	NetIssuance              *big.Int
	DepositedCollateralValue *big.Int
	ReportedDebt             *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterMarketCollateralDeposited is a free log retrieval operation binding the contract event 0x0268c0025d1310f8cbf9a431c755af708633271b9b5902855857297267d6f41b.
//
// Solidity: event MarketCollateralDeposited(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) FilterMarketCollateralDeposited(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreMarketCollateralDepositedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "MarketCollateralDeposited", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreMarketCollateralDepositedIterator{contract: _Core.contract, event: "MarketCollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchMarketCollateralDeposited is a free log subscription operation binding the contract event 0x0268c0025d1310f8cbf9a431c755af708633271b9b5902855857297267d6f41b.
//
// Solidity: event MarketCollateralDeposited(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) WatchMarketCollateralDeposited(opts *bind.WatchOpts, sink chan<- *CoreMarketCollateralDeposited, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "MarketCollateralDeposited", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreMarketCollateralDeposited)
				if err := _Core.contract.UnpackLog(event, "MarketCollateralDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketCollateralDeposited is a log parse operation binding the contract event 0x0268c0025d1310f8cbf9a431c755af708633271b9b5902855857297267d6f41b.
//
// Solidity: event MarketCollateralDeposited(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) ParseMarketCollateralDeposited(log types.Log) (*CoreMarketCollateralDeposited, error) {
	event := new(CoreMarketCollateralDeposited)
	if err := _Core.contract.UnpackLog(event, "MarketCollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreMarketCollateralWithdrawnIterator is returned from FilterMarketCollateralWithdrawn and is used to iterate over the raw logs and unpacked data for MarketCollateralWithdrawn events raised by the Core contract.
type CoreMarketCollateralWithdrawnIterator struct {
	Event *CoreMarketCollateralWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreMarketCollateralWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreMarketCollateralWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreMarketCollateralWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreMarketCollateralWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreMarketCollateralWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreMarketCollateralWithdrawn represents a MarketCollateralWithdrawn event raised by the Core contract.
type CoreMarketCollateralWithdrawn struct {
	MarketId                 *big.Int
	CollateralType           common.Address
	TokenAmount              *big.Int
	Sender                   common.Address
	CreditCapacity           *big.Int
	NetIssuance              *big.Int
	DepositedCollateralValue *big.Int
	ReportedDebt             *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterMarketCollateralWithdrawn is a free log retrieval operation binding the contract event 0x88eb4cc1feb3af3a3e45798dc1d42ec34ef453093ffe0c56fc36e27abd2cc4d7.
//
// Solidity: event MarketCollateralWithdrawn(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) FilterMarketCollateralWithdrawn(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreMarketCollateralWithdrawnIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "MarketCollateralWithdrawn", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreMarketCollateralWithdrawnIterator{contract: _Core.contract, event: "MarketCollateralWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMarketCollateralWithdrawn is a free log subscription operation binding the contract event 0x88eb4cc1feb3af3a3e45798dc1d42ec34ef453093ffe0c56fc36e27abd2cc4d7.
//
// Solidity: event MarketCollateralWithdrawn(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) WatchMarketCollateralWithdrawn(opts *bind.WatchOpts, sink chan<- *CoreMarketCollateralWithdrawn, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "MarketCollateralWithdrawn", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreMarketCollateralWithdrawn)
				if err := _Core.contract.UnpackLog(event, "MarketCollateralWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketCollateralWithdrawn is a log parse operation binding the contract event 0x88eb4cc1feb3af3a3e45798dc1d42ec34ef453093ffe0c56fc36e27abd2cc4d7.
//
// Solidity: event MarketCollateralWithdrawn(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) ParseMarketCollateralWithdrawn(log types.Log) (*CoreMarketCollateralWithdrawn, error) {
	event := new(CoreMarketCollateralWithdrawn)
	if err := _Core.contract.UnpackLog(event, "MarketCollateralWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreMarketRegisteredIterator is returned from FilterMarketRegistered and is used to iterate over the raw logs and unpacked data for MarketRegistered events raised by the Core contract.
type CoreMarketRegisteredIterator struct {
	Event *CoreMarketRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreMarketRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreMarketRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreMarketRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreMarketRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreMarketRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreMarketRegistered represents a MarketRegistered event raised by the Core contract.
type CoreMarketRegistered struct {
	Market   common.Address
	MarketId *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketRegistered is a free log retrieval operation binding the contract event 0xeb87361ace8c1947e121293eb214f68d889d9cf273c48246b38c3cbf185748d0.
//
// Solidity: event MarketRegistered(address indexed market, uint128 indexed marketId, address indexed sender)
func (_Core *CoreFilterer) FilterMarketRegistered(opts *bind.FilterOpts, market []common.Address, marketId []*big.Int, sender []common.Address) (*CoreMarketRegisteredIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "MarketRegistered", marketRule, marketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreMarketRegisteredIterator{contract: _Core.contract, event: "MarketRegistered", logs: logs, sub: sub}, nil
}

// WatchMarketRegistered is a free log subscription operation binding the contract event 0xeb87361ace8c1947e121293eb214f68d889d9cf273c48246b38c3cbf185748d0.
//
// Solidity: event MarketRegistered(address indexed market, uint128 indexed marketId, address indexed sender)
func (_Core *CoreFilterer) WatchMarketRegistered(opts *bind.WatchOpts, sink chan<- *CoreMarketRegistered, market []common.Address, marketId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "MarketRegistered", marketRule, marketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreMarketRegistered)
				if err := _Core.contract.UnpackLog(event, "MarketRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseMarketRegistered(log types.Log) (*CoreMarketRegistered, error) {
	event := new(CoreMarketRegistered)
	if err := _Core.contract.UnpackLog(event, "MarketRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreMarketSystemFeePaidIterator is returned from FilterMarketSystemFeePaid and is used to iterate over the raw logs and unpacked data for MarketSystemFeePaid events raised by the Core contract.
type CoreMarketSystemFeePaidIterator struct {
	Event *CoreMarketSystemFeePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreMarketSystemFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreMarketSystemFeePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreMarketSystemFeePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreMarketSystemFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreMarketSystemFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreMarketSystemFeePaid represents a MarketSystemFeePaid event raised by the Core contract.
type CoreMarketSystemFeePaid struct {
	MarketId  *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMarketSystemFeePaid is a free log retrieval operation binding the contract event 0x8b69fed8aed97ef9572216662359ece45fa52f2b5ff44a78b7ec3c5ef05153f8.
//
// Solidity: event MarketSystemFeePaid(uint128 indexed marketId, uint256 feeAmount)
func (_Core *CoreFilterer) FilterMarketSystemFeePaid(opts *bind.FilterOpts, marketId []*big.Int) (*CoreMarketSystemFeePaidIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "MarketSystemFeePaid", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreMarketSystemFeePaidIterator{contract: _Core.contract, event: "MarketSystemFeePaid", logs: logs, sub: sub}, nil
}

// WatchMarketSystemFeePaid is a free log subscription operation binding the contract event 0x8b69fed8aed97ef9572216662359ece45fa52f2b5ff44a78b7ec3c5ef05153f8.
//
// Solidity: event MarketSystemFeePaid(uint128 indexed marketId, uint256 feeAmount)
func (_Core *CoreFilterer) WatchMarketSystemFeePaid(opts *bind.WatchOpts, sink chan<- *CoreMarketSystemFeePaid, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "MarketSystemFeePaid", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreMarketSystemFeePaid)
				if err := _Core.contract.UnpackLog(event, "MarketSystemFeePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseMarketSystemFeePaid(log types.Log) (*CoreMarketSystemFeePaid, error) {
	event := new(CoreMarketSystemFeePaid)
	if err := _Core.contract.UnpackLog(event, "MarketSystemFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreMarketUsdDepositedIterator is returned from FilterMarketUsdDeposited and is used to iterate over the raw logs and unpacked data for MarketUsdDeposited events raised by the Core contract.
type CoreMarketUsdDepositedIterator struct {
	Event *CoreMarketUsdDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreMarketUsdDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreMarketUsdDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreMarketUsdDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreMarketUsdDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreMarketUsdDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreMarketUsdDeposited represents a MarketUsdDeposited event raised by the Core contract.
type CoreMarketUsdDeposited struct {
	MarketId                 *big.Int
	Target                   common.Address
	Amount                   *big.Int
	Market                   common.Address
	CreditCapacity           *big.Int
	NetIssuance              *big.Int
	DepositedCollateralValue *big.Int
	ReportedDebt             *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterMarketUsdDeposited is a free log retrieval operation binding the contract event 0xf92dddcee1f7e80b04d1141159985b52e0dc3b69d26b9f67b714bf9ff70b4b85.
//
// Solidity: event MarketUsdDeposited(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) FilterMarketUsdDeposited(opts *bind.FilterOpts, marketId []*big.Int, target []common.Address, market []common.Address) (*CoreMarketUsdDepositedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "MarketUsdDeposited", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &CoreMarketUsdDepositedIterator{contract: _Core.contract, event: "MarketUsdDeposited", logs: logs, sub: sub}, nil
}

// WatchMarketUsdDeposited is a free log subscription operation binding the contract event 0xf92dddcee1f7e80b04d1141159985b52e0dc3b69d26b9f67b714bf9ff70b4b85.
//
// Solidity: event MarketUsdDeposited(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) WatchMarketUsdDeposited(opts *bind.WatchOpts, sink chan<- *CoreMarketUsdDeposited, marketId []*big.Int, target []common.Address, market []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "MarketUsdDeposited", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreMarketUsdDeposited)
				if err := _Core.contract.UnpackLog(event, "MarketUsdDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketUsdDeposited is a log parse operation binding the contract event 0xf92dddcee1f7e80b04d1141159985b52e0dc3b69d26b9f67b714bf9ff70b4b85.
//
// Solidity: event MarketUsdDeposited(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) ParseMarketUsdDeposited(log types.Log) (*CoreMarketUsdDeposited, error) {
	event := new(CoreMarketUsdDeposited)
	if err := _Core.contract.UnpackLog(event, "MarketUsdDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreMarketUsdWithdrawnIterator is returned from FilterMarketUsdWithdrawn and is used to iterate over the raw logs and unpacked data for MarketUsdWithdrawn events raised by the Core contract.
type CoreMarketUsdWithdrawnIterator struct {
	Event *CoreMarketUsdWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreMarketUsdWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreMarketUsdWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreMarketUsdWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreMarketUsdWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreMarketUsdWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreMarketUsdWithdrawn represents a MarketUsdWithdrawn event raised by the Core contract.
type CoreMarketUsdWithdrawn struct {
	MarketId                 *big.Int
	Target                   common.Address
	Amount                   *big.Int
	Market                   common.Address
	CreditCapacity           *big.Int
	NetIssuance              *big.Int
	DepositedCollateralValue *big.Int
	ReportedDebt             *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterMarketUsdWithdrawn is a free log retrieval operation binding the contract event 0x53e76a481a00b96d9b07448b333af94f26fbac5059ca153986bf2c98001c3f2c.
//
// Solidity: event MarketUsdWithdrawn(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) FilterMarketUsdWithdrawn(opts *bind.FilterOpts, marketId []*big.Int, target []common.Address, market []common.Address) (*CoreMarketUsdWithdrawnIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "MarketUsdWithdrawn", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &CoreMarketUsdWithdrawnIterator{contract: _Core.contract, event: "MarketUsdWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMarketUsdWithdrawn is a free log subscription operation binding the contract event 0x53e76a481a00b96d9b07448b333af94f26fbac5059ca153986bf2c98001c3f2c.
//
// Solidity: event MarketUsdWithdrawn(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) WatchMarketUsdWithdrawn(opts *bind.WatchOpts, sink chan<- *CoreMarketUsdWithdrawn, marketId []*big.Int, target []common.Address, market []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "MarketUsdWithdrawn", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreMarketUsdWithdrawn)
				if err := _Core.contract.UnpackLog(event, "MarketUsdWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketUsdWithdrawn is a log parse operation binding the contract event 0x53e76a481a00b96d9b07448b333af94f26fbac5059ca153986bf2c98001c3f2c.
//
// Solidity: event MarketUsdWithdrawn(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market, int128 creditCapacity, int128 netIssuance, uint256 depositedCollateralValue, uint256 reportedDebt)
func (_Core *CoreFilterer) ParseMarketUsdWithdrawn(log types.Log) (*CoreMarketUsdWithdrawn, error) {
	event := new(CoreMarketUsdWithdrawn)
	if err := _Core.contract.UnpackLog(event, "MarketUsdWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreMaximumMarketCollateralConfiguredIterator is returned from FilterMaximumMarketCollateralConfigured and is used to iterate over the raw logs and unpacked data for MaximumMarketCollateralConfigured events raised by the Core contract.
type CoreMaximumMarketCollateralConfiguredIterator struct {
	Event *CoreMaximumMarketCollateralConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreMaximumMarketCollateralConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreMaximumMarketCollateralConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreMaximumMarketCollateralConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreMaximumMarketCollateralConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreMaximumMarketCollateralConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreMaximumMarketCollateralConfigured represents a MaximumMarketCollateralConfigured event raised by the Core contract.
type CoreMaximumMarketCollateralConfigured struct {
	MarketId       *big.Int
	CollateralType common.Address
	SystemAmount   *big.Int
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMaximumMarketCollateralConfigured is a free log retrieval operation binding the contract event 0x499c8fcfbc4341c37dcf444c890d42ef888d46aa586f97ceb20577c2635c8075.
//
// Solidity: event MaximumMarketCollateralConfigured(uint128 indexed marketId, address indexed collateralType, uint256 systemAmount, address indexed owner)
func (_Core *CoreFilterer) FilterMaximumMarketCollateralConfigured(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, owner []common.Address) (*CoreMaximumMarketCollateralConfiguredIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "MaximumMarketCollateralConfigured", marketIdRule, collateralTypeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreMaximumMarketCollateralConfiguredIterator{contract: _Core.contract, event: "MaximumMarketCollateralConfigured", logs: logs, sub: sub}, nil
}

// WatchMaximumMarketCollateralConfigured is a free log subscription operation binding the contract event 0x499c8fcfbc4341c37dcf444c890d42ef888d46aa586f97ceb20577c2635c8075.
//
// Solidity: event MaximumMarketCollateralConfigured(uint128 indexed marketId, address indexed collateralType, uint256 systemAmount, address indexed owner)
func (_Core *CoreFilterer) WatchMaximumMarketCollateralConfigured(opts *bind.WatchOpts, sink chan<- *CoreMaximumMarketCollateralConfigured, marketId []*big.Int, collateralType []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "MaximumMarketCollateralConfigured", marketIdRule, collateralTypeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreMaximumMarketCollateralConfigured)
				if err := _Core.contract.UnpackLog(event, "MaximumMarketCollateralConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseMaximumMarketCollateralConfigured(log types.Log) (*CoreMaximumMarketCollateralConfigured, error) {
	event := new(CoreMaximumMarketCollateralConfigured)
	if err := _Core.contract.UnpackLog(event, "MaximumMarketCollateralConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreNewSupportedCrossChainNetworkIterator is returned from FilterNewSupportedCrossChainNetwork and is used to iterate over the raw logs and unpacked data for NewSupportedCrossChainNetwork events raised by the Core contract.
type CoreNewSupportedCrossChainNetworkIterator struct {
	Event *CoreNewSupportedCrossChainNetwork // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreNewSupportedCrossChainNetworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreNewSupportedCrossChainNetwork)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreNewSupportedCrossChainNetwork)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreNewSupportedCrossChainNetworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreNewSupportedCrossChainNetworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreNewSupportedCrossChainNetwork represents a NewSupportedCrossChainNetwork event raised by the Core contract.
type CoreNewSupportedCrossChainNetwork struct {
	NewChainId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewSupportedCrossChainNetwork is a free log retrieval operation binding the contract event 0x1874eb2a5288e478dcedf1d33291bd7293eeef5946ec516d2ef54a364b3f63d8.
//
// Solidity: event NewSupportedCrossChainNetwork(uint64 newChainId)
func (_Core *CoreFilterer) FilterNewSupportedCrossChainNetwork(opts *bind.FilterOpts) (*CoreNewSupportedCrossChainNetworkIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "NewSupportedCrossChainNetwork")
	if err != nil {
		return nil, err
	}
	return &CoreNewSupportedCrossChainNetworkIterator{contract: _Core.contract, event: "NewSupportedCrossChainNetwork", logs: logs, sub: sub}, nil
}

// WatchNewSupportedCrossChainNetwork is a free log subscription operation binding the contract event 0x1874eb2a5288e478dcedf1d33291bd7293eeef5946ec516d2ef54a364b3f63d8.
//
// Solidity: event NewSupportedCrossChainNetwork(uint64 newChainId)
func (_Core *CoreFilterer) WatchNewSupportedCrossChainNetwork(opts *bind.WatchOpts, sink chan<- *CoreNewSupportedCrossChainNetwork) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "NewSupportedCrossChainNetwork")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreNewSupportedCrossChainNetwork)
				if err := _Core.contract.UnpackLog(event, "NewSupportedCrossChainNetwork", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseNewSupportedCrossChainNetwork(log types.Log) (*CoreNewSupportedCrossChainNetwork, error) {
	event := new(CoreNewSupportedCrossChainNetwork)
	if err := _Core.contract.UnpackLog(event, "NewSupportedCrossChainNetwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Core contract.
type CoreOwnerChangedIterator struct {
	Event *CoreOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOwnerChanged represents a OwnerChanged event raised by the Core contract.
type CoreOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Core *CoreFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*CoreOwnerChangedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &CoreOwnerChangedIterator{contract: _Core.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Core *CoreFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *CoreOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOwnerChanged)
				if err := _Core.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseOwnerChanged(log types.Log) (*CoreOwnerChanged, error) {
	event := new(CoreOwnerChanged)
	if err := _Core.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Core contract.
type CoreOwnerNominatedIterator struct {
	Event *CoreOwnerNominated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreOwnerNominated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreOwnerNominated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreOwnerNominated represents a OwnerNominated event raised by the Core contract.
type CoreOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Core *CoreFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*CoreOwnerNominatedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &CoreOwnerNominatedIterator{contract: _Core.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Core *CoreFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *CoreOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreOwnerNominated)
				if err := _Core.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseOwnerNominated(log types.Log) (*CoreOwnerNominated, error) {
	event := new(CoreOwnerNominated)
	if err := _Core.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePermissionGrantedIterator is returned from FilterPermissionGranted and is used to iterate over the raw logs and unpacked data for PermissionGranted events raised by the Core contract.
type CorePermissionGrantedIterator struct {
	Event *CorePermissionGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePermissionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePermissionGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePermissionGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePermissionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePermissionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePermissionGranted represents a PermissionGranted event raised by the Core contract.
type CorePermissionGranted struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionGranted is a free log retrieval operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_Core *CoreFilterer) FilterPermissionGranted(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*CorePermissionGrantedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CorePermissionGrantedIterator{contract: _Core.contract, event: "PermissionGranted", logs: logs, sub: sub}, nil
}

// WatchPermissionGranted is a free log subscription operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_Core *CoreFilterer) WatchPermissionGranted(opts *bind.WatchOpts, sink chan<- *CorePermissionGranted, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePermissionGranted)
				if err := _Core.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePermissionGranted(log types.Log) (*CorePermissionGranted, error) {
	event := new(CorePermissionGranted)
	if err := _Core.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePermissionRevokedIterator is returned from FilterPermissionRevoked and is used to iterate over the raw logs and unpacked data for PermissionRevoked events raised by the Core contract.
type CorePermissionRevokedIterator struct {
	Event *CorePermissionRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePermissionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePermissionRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePermissionRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePermissionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePermissionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePermissionRevoked represents a PermissionRevoked event raised by the Core contract.
type CorePermissionRevoked struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionRevoked is a free log retrieval operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_Core *CoreFilterer) FilterPermissionRevoked(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*CorePermissionRevokedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CorePermissionRevokedIterator{contract: _Core.contract, event: "PermissionRevoked", logs: logs, sub: sub}, nil
}

// WatchPermissionRevoked is a free log subscription operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_Core *CoreFilterer) WatchPermissionRevoked(opts *bind.WatchOpts, sink chan<- *CorePermissionRevoked, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePermissionRevoked)
				if err := _Core.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePermissionRevoked(log types.Log) (*CorePermissionRevoked, error) {
	event := new(CorePermissionRevoked)
	if err := _Core.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolApprovedAddedIterator is returned from FilterPoolApprovedAdded and is used to iterate over the raw logs and unpacked data for PoolApprovedAdded events raised by the Core contract.
type CorePoolApprovedAddedIterator struct {
	Event *CorePoolApprovedAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolApprovedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolApprovedAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolApprovedAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolApprovedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolApprovedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolApprovedAdded represents a PoolApprovedAdded event raised by the Core contract.
type CorePoolApprovedAdded struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolApprovedAdded is a free log retrieval operation binding the contract event 0x7d5bdf4e8c44e0b5a8249bf03c2a1febd848cc7f580efd7b1703301c5b1a9e4e.
//
// Solidity: event PoolApprovedAdded(uint256 poolId)
func (_Core *CoreFilterer) FilterPoolApprovedAdded(opts *bind.FilterOpts) (*CorePoolApprovedAddedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolApprovedAdded")
	if err != nil {
		return nil, err
	}
	return &CorePoolApprovedAddedIterator{contract: _Core.contract, event: "PoolApprovedAdded", logs: logs, sub: sub}, nil
}

// WatchPoolApprovedAdded is a free log subscription operation binding the contract event 0x7d5bdf4e8c44e0b5a8249bf03c2a1febd848cc7f580efd7b1703301c5b1a9e4e.
//
// Solidity: event PoolApprovedAdded(uint256 poolId)
func (_Core *CoreFilterer) WatchPoolApprovedAdded(opts *bind.WatchOpts, sink chan<- *CorePoolApprovedAdded) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolApprovedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolApprovedAdded)
				if err := _Core.contract.UnpackLog(event, "PoolApprovedAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolApprovedAdded(log types.Log) (*CorePoolApprovedAdded, error) {
	event := new(CorePoolApprovedAdded)
	if err := _Core.contract.UnpackLog(event, "PoolApprovedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolApprovedRemovedIterator is returned from FilterPoolApprovedRemoved and is used to iterate over the raw logs and unpacked data for PoolApprovedRemoved events raised by the Core contract.
type CorePoolApprovedRemovedIterator struct {
	Event *CorePoolApprovedRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolApprovedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolApprovedRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolApprovedRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolApprovedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolApprovedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolApprovedRemoved represents a PoolApprovedRemoved event raised by the Core contract.
type CorePoolApprovedRemoved struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolApprovedRemoved is a free log retrieval operation binding the contract event 0xc1567ee9983f306f073ea7d59a7fb5680ce07985f8b49cc50d00a3a9f748d3c2.
//
// Solidity: event PoolApprovedRemoved(uint256 poolId)
func (_Core *CoreFilterer) FilterPoolApprovedRemoved(opts *bind.FilterOpts) (*CorePoolApprovedRemovedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolApprovedRemoved")
	if err != nil {
		return nil, err
	}
	return &CorePoolApprovedRemovedIterator{contract: _Core.contract, event: "PoolApprovedRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolApprovedRemoved is a free log subscription operation binding the contract event 0xc1567ee9983f306f073ea7d59a7fb5680ce07985f8b49cc50d00a3a9f748d3c2.
//
// Solidity: event PoolApprovedRemoved(uint256 poolId)
func (_Core *CoreFilterer) WatchPoolApprovedRemoved(opts *bind.WatchOpts, sink chan<- *CorePoolApprovedRemoved) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolApprovedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolApprovedRemoved)
				if err := _Core.contract.UnpackLog(event, "PoolApprovedRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolApprovedRemoved(log types.Log) (*CorePoolApprovedRemoved, error) {
	event := new(CorePoolApprovedRemoved)
	if err := _Core.contract.UnpackLog(event, "PoolApprovedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolCollateralConfigurationUpdatedIterator is returned from FilterPoolCollateralConfigurationUpdated and is used to iterate over the raw logs and unpacked data for PoolCollateralConfigurationUpdated events raised by the Core contract.
type CorePoolCollateralConfigurationUpdatedIterator struct {
	Event *CorePoolCollateralConfigurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolCollateralConfigurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolCollateralConfigurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolCollateralConfigurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolCollateralConfigurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolCollateralConfigurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolCollateralConfigurationUpdated represents a PoolCollateralConfigurationUpdated event raised by the Core contract.
type CorePoolCollateralConfigurationUpdated struct {
	PoolId         *big.Int
	CollateralType common.Address
	Config         PoolCollateralConfigurationData
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolCollateralConfigurationUpdated is a free log retrieval operation binding the contract event 0x5ebb5c59166ab9735b293a159ee2129e61d16b526867763f25557a275a2aad92.
//
// Solidity: event PoolCollateralConfigurationUpdated(uint128 indexed poolId, address collateralType, (uint256,uint256) config)
func (_Core *CoreFilterer) FilterPoolCollateralConfigurationUpdated(opts *bind.FilterOpts, poolId []*big.Int) (*CorePoolCollateralConfigurationUpdatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolCollateralConfigurationUpdated", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CorePoolCollateralConfigurationUpdatedIterator{contract: _Core.contract, event: "PoolCollateralConfigurationUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolCollateralConfigurationUpdated is a free log subscription operation binding the contract event 0x5ebb5c59166ab9735b293a159ee2129e61d16b526867763f25557a275a2aad92.
//
// Solidity: event PoolCollateralConfigurationUpdated(uint128 indexed poolId, address collateralType, (uint256,uint256) config)
func (_Core *CoreFilterer) WatchPoolCollateralConfigurationUpdated(opts *bind.WatchOpts, sink chan<- *CorePoolCollateralConfigurationUpdated, poolId []*big.Int) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolCollateralConfigurationUpdated", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolCollateralConfigurationUpdated)
				if err := _Core.contract.UnpackLog(event, "PoolCollateralConfigurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolCollateralConfigurationUpdated(log types.Log) (*CorePoolCollateralConfigurationUpdated, error) {
	event := new(CorePoolCollateralConfigurationUpdated)
	if err := _Core.contract.UnpackLog(event, "PoolCollateralConfigurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolCollateralDisabledByDefaultSetIterator is returned from FilterPoolCollateralDisabledByDefaultSet and is used to iterate over the raw logs and unpacked data for PoolCollateralDisabledByDefaultSet events raised by the Core contract.
type CorePoolCollateralDisabledByDefaultSetIterator struct {
	Event *CorePoolCollateralDisabledByDefaultSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolCollateralDisabledByDefaultSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolCollateralDisabledByDefaultSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolCollateralDisabledByDefaultSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolCollateralDisabledByDefaultSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolCollateralDisabledByDefaultSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolCollateralDisabledByDefaultSet represents a PoolCollateralDisabledByDefaultSet event raised by the Core contract.
type CorePoolCollateralDisabledByDefaultSet struct {
	PoolId   *big.Int
	Disabled bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPoolCollateralDisabledByDefaultSet is a free log retrieval operation binding the contract event 0xe0ed98ef42e6a4a881ae0d3c4459c9ed06a36a2144e02efc11823c6cae515bf2.
//
// Solidity: event PoolCollateralDisabledByDefaultSet(uint128 poolId, bool disabled)
func (_Core *CoreFilterer) FilterPoolCollateralDisabledByDefaultSet(opts *bind.FilterOpts) (*CorePoolCollateralDisabledByDefaultSetIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolCollateralDisabledByDefaultSet")
	if err != nil {
		return nil, err
	}
	return &CorePoolCollateralDisabledByDefaultSetIterator{contract: _Core.contract, event: "PoolCollateralDisabledByDefaultSet", logs: logs, sub: sub}, nil
}

// WatchPoolCollateralDisabledByDefaultSet is a free log subscription operation binding the contract event 0xe0ed98ef42e6a4a881ae0d3c4459c9ed06a36a2144e02efc11823c6cae515bf2.
//
// Solidity: event PoolCollateralDisabledByDefaultSet(uint128 poolId, bool disabled)
func (_Core *CoreFilterer) WatchPoolCollateralDisabledByDefaultSet(opts *bind.WatchOpts, sink chan<- *CorePoolCollateralDisabledByDefaultSet) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolCollateralDisabledByDefaultSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolCollateralDisabledByDefaultSet)
				if err := _Core.contract.UnpackLog(event, "PoolCollateralDisabledByDefaultSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolCollateralDisabledByDefaultSet(log types.Log) (*CorePoolCollateralDisabledByDefaultSet, error) {
	event := new(CorePoolCollateralDisabledByDefaultSet)
	if err := _Core.contract.UnpackLog(event, "PoolCollateralDisabledByDefaultSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolConfigurationSetIterator is returned from FilterPoolConfigurationSet and is used to iterate over the raw logs and unpacked data for PoolConfigurationSet events raised by the Core contract.
type CorePoolConfigurationSetIterator struct {
	Event *CorePoolConfigurationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolConfigurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolConfigurationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolConfigurationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolConfigurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolConfigurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolConfigurationSet represents a PoolConfigurationSet event raised by the Core contract.
type CorePoolConfigurationSet struct {
	PoolId  *big.Int
	Markets []MarketConfigurationData
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolConfigurationSet is a free log retrieval operation binding the contract event 0xdd812c2e47943d98e6c66b2b9872d1f9270b8523c82eb60ad5c8d580a614081c.
//
// Solidity: event PoolConfigurationSet(uint128 indexed poolId, (uint128,uint128,int128)[] markets, address indexed sender)
func (_Core *CoreFilterer) FilterPoolConfigurationSet(opts *bind.FilterOpts, poolId []*big.Int, sender []common.Address) (*CorePoolConfigurationSetIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolConfigurationSet", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CorePoolConfigurationSetIterator{contract: _Core.contract, event: "PoolConfigurationSet", logs: logs, sub: sub}, nil
}

// WatchPoolConfigurationSet is a free log subscription operation binding the contract event 0xdd812c2e47943d98e6c66b2b9872d1f9270b8523c82eb60ad5c8d580a614081c.
//
// Solidity: event PoolConfigurationSet(uint128 indexed poolId, (uint128,uint128,int128)[] markets, address indexed sender)
func (_Core *CoreFilterer) WatchPoolConfigurationSet(opts *bind.WatchOpts, sink chan<- *CorePoolConfigurationSet, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolConfigurationSet", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolConfigurationSet)
				if err := _Core.contract.UnpackLog(event, "PoolConfigurationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolConfigurationSet(log types.Log) (*CorePoolConfigurationSet, error) {
	event := new(CorePoolConfigurationSet)
	if err := _Core.contract.UnpackLog(event, "PoolConfigurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the Core contract.
type CorePoolCreatedIterator struct {
	Event *CorePoolCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolCreated represents a PoolCreated event raised by the Core contract.
type CorePoolCreated struct {
	PoolId *big.Int
	Owner  common.Address
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xb1517ad708e5f9a104c30d3f1ff749d55833b1d03bf472013c29888e741cf340.
//
// Solidity: event PoolCreated(uint128 indexed poolId, address indexed owner, address indexed sender)
func (_Core *CoreFilterer) FilterPoolCreated(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address, sender []common.Address) (*CorePoolCreatedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolCreated", poolIdRule, ownerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CorePoolCreatedIterator{contract: _Core.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xb1517ad708e5f9a104c30d3f1ff749d55833b1d03bf472013c29888e741cf340.
//
// Solidity: event PoolCreated(uint128 indexed poolId, address indexed owner, address indexed sender)
func (_Core *CoreFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *CorePoolCreated, poolId []*big.Int, owner []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolCreated", poolIdRule, ownerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolCreated)
				if err := _Core.contract.UnpackLog(event, "PoolCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolCreated(log types.Log) (*CorePoolCreated, error) {
	event := new(CorePoolCreated)
	if err := _Core.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolNameUpdatedIterator is returned from FilterPoolNameUpdated and is used to iterate over the raw logs and unpacked data for PoolNameUpdated events raised by the Core contract.
type CorePoolNameUpdatedIterator struct {
	Event *CorePoolNameUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolNameUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolNameUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolNameUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolNameUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolNameUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolNameUpdated represents a PoolNameUpdated event raised by the Core contract.
type CorePoolNameUpdated struct {
	PoolId *big.Int
	Name   string
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNameUpdated is a free log retrieval operation binding the contract event 0x63b42abaf7e145a993f20bc64259f45d09c43d18838ab0bca078b15093ac55f4.
//
// Solidity: event PoolNameUpdated(uint128 indexed poolId, string name, address indexed sender)
func (_Core *CoreFilterer) FilterPoolNameUpdated(opts *bind.FilterOpts, poolId []*big.Int, sender []common.Address) (*CorePoolNameUpdatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolNameUpdated", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CorePoolNameUpdatedIterator{contract: _Core.contract, event: "PoolNameUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolNameUpdated is a free log subscription operation binding the contract event 0x63b42abaf7e145a993f20bc64259f45d09c43d18838ab0bca078b15093ac55f4.
//
// Solidity: event PoolNameUpdated(uint128 indexed poolId, string name, address indexed sender)
func (_Core *CoreFilterer) WatchPoolNameUpdated(opts *bind.WatchOpts, sink chan<- *CorePoolNameUpdated, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolNameUpdated", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolNameUpdated)
				if err := _Core.contract.UnpackLog(event, "PoolNameUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolNameUpdated(log types.Log) (*CorePoolNameUpdated, error) {
	event := new(CorePoolNameUpdated)
	if err := _Core.contract.UnpackLog(event, "PoolNameUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolNominationRenouncedIterator is returned from FilterPoolNominationRenounced and is used to iterate over the raw logs and unpacked data for PoolNominationRenounced events raised by the Core contract.
type CorePoolNominationRenouncedIterator struct {
	Event *CorePoolNominationRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolNominationRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolNominationRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolNominationRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolNominationRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolNominationRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolNominationRenounced represents a PoolNominationRenounced event raised by the Core contract.
type CorePoolNominationRenounced struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNominationRenounced is a free log retrieval operation binding the contract event 0x28301da5fb0feefb138efa6310af4547a74f415d62616f90519436dc169c3ae0.
//
// Solidity: event PoolNominationRenounced(uint128 indexed poolId, address indexed owner)
func (_Core *CoreFilterer) FilterPoolNominationRenounced(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*CorePoolNominationRenouncedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolNominationRenounced", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CorePoolNominationRenouncedIterator{contract: _Core.contract, event: "PoolNominationRenounced", logs: logs, sub: sub}, nil
}

// WatchPoolNominationRenounced is a free log subscription operation binding the contract event 0x28301da5fb0feefb138efa6310af4547a74f415d62616f90519436dc169c3ae0.
//
// Solidity: event PoolNominationRenounced(uint128 indexed poolId, address indexed owner)
func (_Core *CoreFilterer) WatchPoolNominationRenounced(opts *bind.WatchOpts, sink chan<- *CorePoolNominationRenounced, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolNominationRenounced", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolNominationRenounced)
				if err := _Core.contract.UnpackLog(event, "PoolNominationRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolNominationRenounced(log types.Log) (*CorePoolNominationRenounced, error) {
	event := new(CorePoolNominationRenounced)
	if err := _Core.contract.UnpackLog(event, "PoolNominationRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolNominationRevokedIterator is returned from FilterPoolNominationRevoked and is used to iterate over the raw logs and unpacked data for PoolNominationRevoked events raised by the Core contract.
type CorePoolNominationRevokedIterator struct {
	Event *CorePoolNominationRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolNominationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolNominationRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolNominationRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolNominationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolNominationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolNominationRevoked represents a PoolNominationRevoked event raised by the Core contract.
type CorePoolNominationRevoked struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNominationRevoked is a free log retrieval operation binding the contract event 0xa20a605599b6da4a06e0662f1284c442a576bc452b77a38c8c55805cb82a1865.
//
// Solidity: event PoolNominationRevoked(uint128 indexed poolId, address indexed owner)
func (_Core *CoreFilterer) FilterPoolNominationRevoked(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*CorePoolNominationRevokedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolNominationRevoked", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CorePoolNominationRevokedIterator{contract: _Core.contract, event: "PoolNominationRevoked", logs: logs, sub: sub}, nil
}

// WatchPoolNominationRevoked is a free log subscription operation binding the contract event 0xa20a605599b6da4a06e0662f1284c442a576bc452b77a38c8c55805cb82a1865.
//
// Solidity: event PoolNominationRevoked(uint128 indexed poolId, address indexed owner)
func (_Core *CoreFilterer) WatchPoolNominationRevoked(opts *bind.WatchOpts, sink chan<- *CorePoolNominationRevoked, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolNominationRevoked", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolNominationRevoked)
				if err := _Core.contract.UnpackLog(event, "PoolNominationRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolNominationRevoked(log types.Log) (*CorePoolNominationRevoked, error) {
	event := new(CorePoolNominationRevoked)
	if err := _Core.contract.UnpackLog(event, "PoolNominationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolOwnerNominatedIterator is returned from FilterPoolOwnerNominated and is used to iterate over the raw logs and unpacked data for PoolOwnerNominated events raised by the Core contract.
type CorePoolOwnerNominatedIterator struct {
	Event *CorePoolOwnerNominated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolOwnerNominated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolOwnerNominated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolOwnerNominated represents a PoolOwnerNominated event raised by the Core contract.
type CorePoolOwnerNominated struct {
	PoolId         *big.Int
	NominatedOwner common.Address
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolOwnerNominated is a free log retrieval operation binding the contract event 0x55d98f82a53fb5776e9ea48d624ab9cb015b51a45249b1ed8425fc857c82f4f8.
//
// Solidity: event PoolOwnerNominated(uint128 indexed poolId, address indexed nominatedOwner, address indexed owner)
func (_Core *CoreFilterer) FilterPoolOwnerNominated(opts *bind.FilterOpts, poolId []*big.Int, nominatedOwner []common.Address, owner []common.Address) (*CorePoolOwnerNominatedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolOwnerNominated", poolIdRule, nominatedOwnerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CorePoolOwnerNominatedIterator{contract: _Core.contract, event: "PoolOwnerNominated", logs: logs, sub: sub}, nil
}

// WatchPoolOwnerNominated is a free log subscription operation binding the contract event 0x55d98f82a53fb5776e9ea48d624ab9cb015b51a45249b1ed8425fc857c82f4f8.
//
// Solidity: event PoolOwnerNominated(uint128 indexed poolId, address indexed nominatedOwner, address indexed owner)
func (_Core *CoreFilterer) WatchPoolOwnerNominated(opts *bind.WatchOpts, sink chan<- *CorePoolOwnerNominated, poolId []*big.Int, nominatedOwner []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolOwnerNominated", poolIdRule, nominatedOwnerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolOwnerNominated)
				if err := _Core.contract.UnpackLog(event, "PoolOwnerNominated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolOwnerNominated(log types.Log) (*CorePoolOwnerNominated, error) {
	event := new(CorePoolOwnerNominated)
	if err := _Core.contract.UnpackLog(event, "PoolOwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolOwnershipAcceptedIterator is returned from FilterPoolOwnershipAccepted and is used to iterate over the raw logs and unpacked data for PoolOwnershipAccepted events raised by the Core contract.
type CorePoolOwnershipAcceptedIterator struct {
	Event *CorePoolOwnershipAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolOwnershipAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolOwnershipAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolOwnershipAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolOwnershipAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolOwnershipAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolOwnershipAccepted represents a PoolOwnershipAccepted event raised by the Core contract.
type CorePoolOwnershipAccepted struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolOwnershipAccepted is a free log retrieval operation binding the contract event 0x4f86f2ce8b08e27d0e470f4269b71c3bbc68407d51a2e692f6573236074ebc5a.
//
// Solidity: event PoolOwnershipAccepted(uint128 indexed poolId, address indexed owner)
func (_Core *CoreFilterer) FilterPoolOwnershipAccepted(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*CorePoolOwnershipAcceptedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolOwnershipAccepted", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CorePoolOwnershipAcceptedIterator{contract: _Core.contract, event: "PoolOwnershipAccepted", logs: logs, sub: sub}, nil
}

// WatchPoolOwnershipAccepted is a free log subscription operation binding the contract event 0x4f86f2ce8b08e27d0e470f4269b71c3bbc68407d51a2e692f6573236074ebc5a.
//
// Solidity: event PoolOwnershipAccepted(uint128 indexed poolId, address indexed owner)
func (_Core *CoreFilterer) WatchPoolOwnershipAccepted(opts *bind.WatchOpts, sink chan<- *CorePoolOwnershipAccepted, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolOwnershipAccepted", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolOwnershipAccepted)
				if err := _Core.contract.UnpackLog(event, "PoolOwnershipAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePoolOwnershipAccepted(log types.Log) (*CorePoolOwnershipAccepted, error) {
	event := new(CorePoolOwnershipAccepted)
	if err := _Core.contract.UnpackLog(event, "PoolOwnershipAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePoolOwnershipRenouncedIterator is returned from FilterPoolOwnershipRenounced and is used to iterate over the raw logs and unpacked data for PoolOwnershipRenounced events raised by the Core contract.
type CorePoolOwnershipRenouncedIterator struct {
	Event *CorePoolOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePoolOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePoolOwnershipRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePoolOwnershipRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePoolOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePoolOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePoolOwnershipRenounced represents a PoolOwnershipRenounced event raised by the Core contract.
type CorePoolOwnershipRenounced struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolOwnershipRenounced is a free log retrieval operation binding the contract event 0x0d1df5c898ce9334fe91f342f5c07b0eea630d388f90b4e07e85753d61252734.
//
// Solidity: event PoolOwnershipRenounced(uint128 indexed poolId, address indexed owner)
func (_Core *CoreFilterer) FilterPoolOwnershipRenounced(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*CorePoolOwnershipRenouncedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "PoolOwnershipRenounced", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CorePoolOwnershipRenouncedIterator{contract: _Core.contract, event: "PoolOwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchPoolOwnershipRenounced is a free log subscription operation binding the contract event 0x0d1df5c898ce9334fe91f342f5c07b0eea630d388f90b4e07e85753d61252734.
//
// Solidity: event PoolOwnershipRenounced(uint128 indexed poolId, address indexed owner)
func (_Core *CoreFilterer) WatchPoolOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *CorePoolOwnershipRenounced, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "PoolOwnershipRenounced", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePoolOwnershipRenounced)
				if err := _Core.contract.UnpackLog(event, "PoolOwnershipRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolOwnershipRenounced is a log parse operation binding the contract event 0x0d1df5c898ce9334fe91f342f5c07b0eea630d388f90b4e07e85753d61252734.
//
// Solidity: event PoolOwnershipRenounced(uint128 indexed poolId, address indexed owner)
func (_Core *CoreFilterer) ParsePoolOwnershipRenounced(log types.Log) (*CorePoolOwnershipRenounced, error) {
	event := new(CorePoolOwnershipRenounced)
	if err := _Core.contract.UnpackLog(event, "PoolOwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CorePreferredPoolSetIterator is returned from FilterPreferredPoolSet and is used to iterate over the raw logs and unpacked data for PreferredPoolSet events raised by the Core contract.
type CorePreferredPoolSetIterator struct {
	Event *CorePreferredPoolSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CorePreferredPoolSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CorePreferredPoolSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CorePreferredPoolSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CorePreferredPoolSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CorePreferredPoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CorePreferredPoolSet represents a PreferredPoolSet event raised by the Core contract.
type CorePreferredPoolSet struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPreferredPoolSet is a free log retrieval operation binding the contract event 0x7e7cb4726e710dc12fad41f158c37a4a71af3a6f053b8b13670d35c710139a56.
//
// Solidity: event PreferredPoolSet(uint256 poolId)
func (_Core *CoreFilterer) FilterPreferredPoolSet(opts *bind.FilterOpts) (*CorePreferredPoolSetIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "PreferredPoolSet")
	if err != nil {
		return nil, err
	}
	return &CorePreferredPoolSetIterator{contract: _Core.contract, event: "PreferredPoolSet", logs: logs, sub: sub}, nil
}

// WatchPreferredPoolSet is a free log subscription operation binding the contract event 0x7e7cb4726e710dc12fad41f158c37a4a71af3a6f053b8b13670d35c710139a56.
//
// Solidity: event PreferredPoolSet(uint256 poolId)
func (_Core *CoreFilterer) WatchPreferredPoolSet(opts *bind.WatchOpts, sink chan<- *CorePreferredPoolSet) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "PreferredPoolSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CorePreferredPoolSet)
				if err := _Core.contract.UnpackLog(event, "PreferredPoolSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParsePreferredPoolSet(log types.Log) (*CorePreferredPoolSet, error) {
	event := new(CorePreferredPoolSet)
	if err := _Core.contract.UnpackLog(event, "PreferredPoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Core contract.
type CoreRewardsClaimedIterator struct {
	Event *CoreRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRewardsClaimed represents a RewardsClaimed event raised by the Core contract.
type CoreRewardsClaimed struct {
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
func (_Core *CoreFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*CoreRewardsClaimedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "RewardsClaimed", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreRewardsClaimedIterator{contract: _Core.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xa4a60be4203e7975e54ab5314c7e9e18aba9ad71e8da714d8de987f4f05410f2.
//
// Solidity: event RewardsClaimed(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount)
func (_Core *CoreFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *CoreRewardsClaimed, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "RewardsClaimed", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRewardsClaimed)
				if err := _Core.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseRewardsClaimed(log types.Log) (*CoreRewardsClaimed, error) {
	event := new(CoreRewardsClaimed)
	if err := _Core.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRewardsDistributedIterator is returned from FilterRewardsDistributed and is used to iterate over the raw logs and unpacked data for RewardsDistributed events raised by the Core contract.
type CoreRewardsDistributedIterator struct {
	Event *CoreRewardsDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreRewardsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRewardsDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreRewardsDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreRewardsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRewardsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRewardsDistributed represents a RewardsDistributed event raised by the Core contract.
type CoreRewardsDistributed struct {
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
func (_Core *CoreFilterer) FilterRewardsDistributed(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address) (*CoreRewardsDistributedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "RewardsDistributed", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreRewardsDistributedIterator{contract: _Core.contract, event: "RewardsDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributed is a free log subscription operation binding the contract event 0x19ced31d71d1db45f99d5a8d3a7616fe9d78828df58f2a28feb68c9f9ab876ca.
//
// Solidity: event RewardsDistributed(uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount, uint256 start, uint256 duration)
func (_Core *CoreFilterer) WatchRewardsDistributed(opts *bind.WatchOpts, sink chan<- *CoreRewardsDistributed, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "RewardsDistributed", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRewardsDistributed)
				if err := _Core.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseRewardsDistributed(log types.Log) (*CoreRewardsDistributed, error) {
	event := new(CoreRewardsDistributed)
	if err := _Core.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRewardsDistributorRegisteredIterator is returned from FilterRewardsDistributorRegistered and is used to iterate over the raw logs and unpacked data for RewardsDistributorRegistered events raised by the Core contract.
type CoreRewardsDistributorRegisteredIterator struct {
	Event *CoreRewardsDistributorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreRewardsDistributorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRewardsDistributorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreRewardsDistributorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreRewardsDistributorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRewardsDistributorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRewardsDistributorRegistered represents a RewardsDistributorRegistered event raised by the Core contract.
type CoreRewardsDistributorRegistered struct {
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributorRegistered is a free log retrieval operation binding the contract event 0x9d3609c05a83dc93a5b355d62c2b37dfde8f0833b1184d4d05c6f51cd46b6e5b.
//
// Solidity: event RewardsDistributorRegistered(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_Core *CoreFilterer) FilterRewardsDistributorRegistered(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (*CoreRewardsDistributorRegisteredIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "RewardsDistributorRegistered", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return &CoreRewardsDistributorRegisteredIterator{contract: _Core.contract, event: "RewardsDistributorRegistered", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributorRegistered is a free log subscription operation binding the contract event 0x9d3609c05a83dc93a5b355d62c2b37dfde8f0833b1184d4d05c6f51cd46b6e5b.
//
// Solidity: event RewardsDistributorRegistered(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_Core *CoreFilterer) WatchRewardsDistributorRegistered(opts *bind.WatchOpts, sink chan<- *CoreRewardsDistributorRegistered, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "RewardsDistributorRegistered", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRewardsDistributorRegistered)
				if err := _Core.contract.UnpackLog(event, "RewardsDistributorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseRewardsDistributorRegistered(log types.Log) (*CoreRewardsDistributorRegistered, error) {
	event := new(CoreRewardsDistributorRegistered)
	if err := _Core.contract.UnpackLog(event, "RewardsDistributorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRewardsDistributorRemovedIterator is returned from FilterRewardsDistributorRemoved and is used to iterate over the raw logs and unpacked data for RewardsDistributorRemoved events raised by the Core contract.
type CoreRewardsDistributorRemovedIterator struct {
	Event *CoreRewardsDistributorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreRewardsDistributorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRewardsDistributorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreRewardsDistributorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreRewardsDistributorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRewardsDistributorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRewardsDistributorRemoved represents a RewardsDistributorRemoved event raised by the Core contract.
type CoreRewardsDistributorRemoved struct {
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributorRemoved is a free log retrieval operation binding the contract event 0x375c4507f463c55a506be95e2cfd3cfdc0610be055087eac6049588a1bcfacba.
//
// Solidity: event RewardsDistributorRemoved(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_Core *CoreFilterer) FilterRewardsDistributorRemoved(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (*CoreRewardsDistributorRemovedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "RewardsDistributorRemoved", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return &CoreRewardsDistributorRemovedIterator{contract: _Core.contract, event: "RewardsDistributorRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributorRemoved is a free log subscription operation binding the contract event 0x375c4507f463c55a506be95e2cfd3cfdc0610be055087eac6049588a1bcfacba.
//
// Solidity: event RewardsDistributorRemoved(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_Core *CoreFilterer) WatchRewardsDistributorRemoved(opts *bind.WatchOpts, sink chan<- *CoreRewardsDistributorRemoved, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "RewardsDistributorRemoved", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRewardsDistributorRemoved)
				if err := _Core.contract.UnpackLog(event, "RewardsDistributorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseRewardsDistributorRemoved(log types.Log) (*CoreRewardsDistributorRemoved, error) {
	event := new(CoreRewardsDistributorRemoved)
	if err := _Core.contract.UnpackLog(event, "RewardsDistributorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreSetMarketMinLiquidityRatioIterator is returned from FilterSetMarketMinLiquidityRatio and is used to iterate over the raw logs and unpacked data for SetMarketMinLiquidityRatio events raised by the Core contract.
type CoreSetMarketMinLiquidityRatioIterator struct {
	Event *CoreSetMarketMinLiquidityRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreSetMarketMinLiquidityRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreSetMarketMinLiquidityRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreSetMarketMinLiquidityRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreSetMarketMinLiquidityRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreSetMarketMinLiquidityRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreSetMarketMinLiquidityRatio represents a SetMarketMinLiquidityRatio event raised by the Core contract.
type CoreSetMarketMinLiquidityRatio struct {
	MarketId          *big.Int
	MinLiquidityRatio *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMarketMinLiquidityRatio is a free log retrieval operation binding the contract event 0x563eb723f21b3e87ec8932cfb4ffa64d1b68c42053c28d6b4db019a40f6daf47.
//
// Solidity: event SetMarketMinLiquidityRatio(uint128 indexed marketId, uint256 minLiquidityRatio)
func (_Core *CoreFilterer) FilterSetMarketMinLiquidityRatio(opts *bind.FilterOpts, marketId []*big.Int) (*CoreSetMarketMinLiquidityRatioIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "SetMarketMinLiquidityRatio", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreSetMarketMinLiquidityRatioIterator{contract: _Core.contract, event: "SetMarketMinLiquidityRatio", logs: logs, sub: sub}, nil
}

// WatchSetMarketMinLiquidityRatio is a free log subscription operation binding the contract event 0x563eb723f21b3e87ec8932cfb4ffa64d1b68c42053c28d6b4db019a40f6daf47.
//
// Solidity: event SetMarketMinLiquidityRatio(uint128 indexed marketId, uint256 minLiquidityRatio)
func (_Core *CoreFilterer) WatchSetMarketMinLiquidityRatio(opts *bind.WatchOpts, sink chan<- *CoreSetMarketMinLiquidityRatio, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "SetMarketMinLiquidityRatio", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreSetMarketMinLiquidityRatio)
				if err := _Core.contract.UnpackLog(event, "SetMarketMinLiquidityRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseSetMarketMinLiquidityRatio(log types.Log) (*CoreSetMarketMinLiquidityRatio, error) {
	event := new(CoreSetMarketMinLiquidityRatio)
	if err := _Core.contract.UnpackLog(event, "SetMarketMinLiquidityRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreSetMinDelegateTimeIterator is returned from FilterSetMinDelegateTime and is used to iterate over the raw logs and unpacked data for SetMinDelegateTime events raised by the Core contract.
type CoreSetMinDelegateTimeIterator struct {
	Event *CoreSetMinDelegateTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreSetMinDelegateTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreSetMinDelegateTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreSetMinDelegateTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreSetMinDelegateTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreSetMinDelegateTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreSetMinDelegateTime represents a SetMinDelegateTime event raised by the Core contract.
type CoreSetMinDelegateTime struct {
	MarketId        *big.Int
	MinDelegateTime uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetMinDelegateTime is a free log retrieval operation binding the contract event 0x6942a68d151863c1fed3c0c4c5f3258af738218527147ac69290ab23ca7d26c6.
//
// Solidity: event SetMinDelegateTime(uint128 indexed marketId, uint32 minDelegateTime)
func (_Core *CoreFilterer) FilterSetMinDelegateTime(opts *bind.FilterOpts, marketId []*big.Int) (*CoreSetMinDelegateTimeIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "SetMinDelegateTime", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreSetMinDelegateTimeIterator{contract: _Core.contract, event: "SetMinDelegateTime", logs: logs, sub: sub}, nil
}

// WatchSetMinDelegateTime is a free log subscription operation binding the contract event 0x6942a68d151863c1fed3c0c4c5f3258af738218527147ac69290ab23ca7d26c6.
//
// Solidity: event SetMinDelegateTime(uint128 indexed marketId, uint32 minDelegateTime)
func (_Core *CoreFilterer) WatchSetMinDelegateTime(opts *bind.WatchOpts, sink chan<- *CoreSetMinDelegateTime, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "SetMinDelegateTime", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreSetMinDelegateTime)
				if err := _Core.contract.UnpackLog(event, "SetMinDelegateTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseSetMinDelegateTime(log types.Log) (*CoreSetMinDelegateTime, error) {
	event := new(CoreSetMinDelegateTime)
	if err := _Core.contract.UnpackLog(event, "SetMinDelegateTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreSetMinLiquidityRatioIterator is returned from FilterSetMinLiquidityRatio and is used to iterate over the raw logs and unpacked data for SetMinLiquidityRatio events raised by the Core contract.
type CoreSetMinLiquidityRatioIterator struct {
	Event *CoreSetMinLiquidityRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreSetMinLiquidityRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreSetMinLiquidityRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreSetMinLiquidityRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreSetMinLiquidityRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreSetMinLiquidityRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreSetMinLiquidityRatio represents a SetMinLiquidityRatio event raised by the Core contract.
type CoreSetMinLiquidityRatio struct {
	MinLiquidityRatio *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMinLiquidityRatio is a free log retrieval operation binding the contract event 0x66fd484d9868d1faddc8fef1f3faed0ed25eb4e6acde49dd1f2cbf0fba903635.
//
// Solidity: event SetMinLiquidityRatio(uint256 minLiquidityRatio)
func (_Core *CoreFilterer) FilterSetMinLiquidityRatio(opts *bind.FilterOpts) (*CoreSetMinLiquidityRatioIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "SetMinLiquidityRatio")
	if err != nil {
		return nil, err
	}
	return &CoreSetMinLiquidityRatioIterator{contract: _Core.contract, event: "SetMinLiquidityRatio", logs: logs, sub: sub}, nil
}

// WatchSetMinLiquidityRatio is a free log subscription operation binding the contract event 0x66fd484d9868d1faddc8fef1f3faed0ed25eb4e6acde49dd1f2cbf0fba903635.
//
// Solidity: event SetMinLiquidityRatio(uint256 minLiquidityRatio)
func (_Core *CoreFilterer) WatchSetMinLiquidityRatio(opts *bind.WatchOpts, sink chan<- *CoreSetMinLiquidityRatio) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "SetMinLiquidityRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreSetMinLiquidityRatio)
				if err := _Core.contract.UnpackLog(event, "SetMinLiquidityRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseSetMinLiquidityRatio(log types.Log) (*CoreSetMinLiquidityRatio, error) {
	event := new(CoreSetMinLiquidityRatio)
	if err := _Core.contract.UnpackLog(event, "SetMinLiquidityRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreTransferCrossChainInitiatedIterator is returned from FilterTransferCrossChainInitiated and is used to iterate over the raw logs and unpacked data for TransferCrossChainInitiated events raised by the Core contract.
type CoreTransferCrossChainInitiatedIterator struct {
	Event *CoreTransferCrossChainInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreTransferCrossChainInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreTransferCrossChainInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreTransferCrossChainInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreTransferCrossChainInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreTransferCrossChainInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreTransferCrossChainInitiated represents a TransferCrossChainInitiated event raised by the Core contract.
type CoreTransferCrossChainInitiated struct {
	DestChainId uint64
	Amount      *big.Int
	Sender      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferCrossChainInitiated is a free log retrieval operation binding the contract event 0xb87c3097d7f9145a4915e8e434f04a1b7b91646d8a6e66a5cdab25caccb644c4.
//
// Solidity: event TransferCrossChainInitiated(uint64 indexed destChainId, uint256 indexed amount, address sender)
func (_Core *CoreFilterer) FilterTransferCrossChainInitiated(opts *bind.FilterOpts, destChainId []uint64, amount []*big.Int) (*CoreTransferCrossChainInitiatedIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "TransferCrossChainInitiated", destChainIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &CoreTransferCrossChainInitiatedIterator{contract: _Core.contract, event: "TransferCrossChainInitiated", logs: logs, sub: sub}, nil
}

// WatchTransferCrossChainInitiated is a free log subscription operation binding the contract event 0xb87c3097d7f9145a4915e8e434f04a1b7b91646d8a6e66a5cdab25caccb644c4.
//
// Solidity: event TransferCrossChainInitiated(uint64 indexed destChainId, uint256 indexed amount, address sender)
func (_Core *CoreFilterer) WatchTransferCrossChainInitiated(opts *bind.WatchOpts, sink chan<- *CoreTransferCrossChainInitiated, destChainId []uint64, amount []*big.Int) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "TransferCrossChainInitiated", destChainIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreTransferCrossChainInitiated)
				if err := _Core.contract.UnpackLog(event, "TransferCrossChainInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseTransferCrossChainInitiated(log types.Log) (*CoreTransferCrossChainInitiated, error) {
	event := new(CoreTransferCrossChainInitiated)
	if err := _Core.contract.UnpackLog(event, "TransferCrossChainInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Core contract.
type CoreUpgradedIterator struct {
	Event *CoreUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreUpgraded represents a Upgraded event raised by the Core contract.
type CoreUpgraded struct {
	Self           common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_Core *CoreFilterer) FilterUpgraded(opts *bind.FilterOpts, self []common.Address) (*CoreUpgradedIterator, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return &CoreUpgradedIterator{contract: _Core.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_Core *CoreFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CoreUpgraded, self []common.Address) (event.Subscription, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreUpgraded)
				if err := _Core.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseUpgraded(log types.Log) (*CoreUpgraded, error) {
	event := new(CoreUpgraded)
	if err := _Core.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreUsdBurnedIterator is returned from FilterUsdBurned and is used to iterate over the raw logs and unpacked data for UsdBurned events raised by the Core contract.
type CoreUsdBurnedIterator struct {
	Event *CoreUsdBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreUsdBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreUsdBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreUsdBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreUsdBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreUsdBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreUsdBurned represents a UsdBurned event raised by the Core contract.
type CoreUsdBurned struct {
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
func (_Core *CoreFilterer) FilterUsdBurned(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*CoreUsdBurnedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "UsdBurned", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreUsdBurnedIterator{contract: _Core.contract, event: "UsdBurned", logs: logs, sub: sub}, nil
}

// WatchUsdBurned is a free log subscription operation binding the contract event 0x6b0230f0abe9188cbdbb1c816a4c5e213758b5b743d8a94af056280cc1e7aeb1.
//
// Solidity: event UsdBurned(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_Core *CoreFilterer) WatchUsdBurned(opts *bind.WatchOpts, sink chan<- *CoreUsdBurned, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "UsdBurned", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreUsdBurned)
				if err := _Core.contract.UnpackLog(event, "UsdBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseUsdBurned(log types.Log) (*CoreUsdBurned, error) {
	event := new(CoreUsdBurned)
	if err := _Core.contract.UnpackLog(event, "UsdBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreUsdMintedIterator is returned from FilterUsdMinted and is used to iterate over the raw logs and unpacked data for UsdMinted events raised by the Core contract.
type CoreUsdMintedIterator struct {
	Event *CoreUsdMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreUsdMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreUsdMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreUsdMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreUsdMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreUsdMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreUsdMinted represents a UsdMinted event raised by the Core contract.
type CoreUsdMinted struct {
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
func (_Core *CoreFilterer) FilterUsdMinted(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*CoreUsdMintedIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "UsdMinted", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreUsdMintedIterator{contract: _Core.contract, event: "UsdMinted", logs: logs, sub: sub}, nil
}

// WatchUsdMinted is a free log subscription operation binding the contract event 0x2100f67dc9a5917400f799bb13194553e3f74c19a207c56350d2c223ac9c36c9.
//
// Solidity: event UsdMinted(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_Core *CoreFilterer) WatchUsdMinted(opts *bind.WatchOpts, sink chan<- *CoreUsdMinted, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "UsdMinted", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreUsdMinted)
				if err := _Core.contract.UnpackLog(event, "UsdMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseUsdMinted(log types.Log) (*CoreUsdMinted, error) {
	event := new(CoreUsdMinted)
	if err := _Core.contract.UnpackLog(event, "UsdMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreVaultLiquidationIterator is returned from FilterVaultLiquidation and is used to iterate over the raw logs and unpacked data for VaultLiquidation events raised by the Core contract.
type CoreVaultLiquidationIterator struct {
	Event *CoreVaultLiquidation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreVaultLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreVaultLiquidation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreVaultLiquidation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreVaultLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreVaultLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreVaultLiquidation represents a VaultLiquidation event raised by the Core contract.
type CoreVaultLiquidation struct {
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
func (_Core *CoreFilterer) FilterVaultLiquidation(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address) (*CoreVaultLiquidationIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "VaultLiquidation", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreVaultLiquidationIterator{contract: _Core.contract, event: "VaultLiquidation", logs: logs, sub: sub}, nil
}

// WatchVaultLiquidation is a free log subscription operation binding the contract event 0x1834a7cc9d14f9bfa482df5c0404dadd1b8ec123b41f082e76ae28a3b2ea68d5.
//
// Solidity: event VaultLiquidation(uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_Core *CoreFilterer) WatchVaultLiquidation(opts *bind.WatchOpts, sink chan<- *CoreVaultLiquidation, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "VaultLiquidation", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreVaultLiquidation)
				if err := _Core.contract.UnpackLog(event, "VaultLiquidation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseVaultLiquidation(log types.Log) (*CoreVaultLiquidation, error) {
	event := new(CoreVaultLiquidation)
	if err := _Core.contract.UnpackLog(event, "VaultLiquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Core contract.
type CoreWithdrawnIterator struct {
	Event *CoreWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreWithdrawn represents a Withdrawn event raised by the Core contract.
type CoreWithdrawn struct {
	AccountId      *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x8b5f9d7ce522936589c630db08c0fa2405b21c4a5ff8ef19899900172736ba38.
//
// Solidity: event Withdrawn(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Core *CoreFilterer) FilterWithdrawn(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreWithdrawnIterator, error) {

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

	logs, sub, err := _Core.contract.FilterLogs(opts, "Withdrawn", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreWithdrawnIterator{contract: _Core.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x8b5f9d7ce522936589c630db08c0fa2405b21c4a5ff8ef19899900172736ba38.
//
// Solidity: event Withdrawn(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Core *CoreFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CoreWithdrawn, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Core.contract.WatchLogs(opts, "Withdrawn", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreWithdrawn)
				if err := _Core.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Core *CoreFilterer) ParseWithdrawn(log types.Log) (*CoreWithdrawn, error) {
	event := new(CoreWithdrawn)
	if err := _Core.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
